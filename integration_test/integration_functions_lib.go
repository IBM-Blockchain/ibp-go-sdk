package integration

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"github.com/IBM-Blockchain/ibp-go-sdk/blockchainv3"
	"github.com/hyperledger/fabric-ca/api"
	"github.com/hyperledger/fabric-ca/lib"
	catls "github.com/hyperledger/fabric-ca/lib/tls"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	LogPrefix string
	Logger *log.Logger
)

func CreateCA(service *blockchainv3.BlockchainV3, displayName string) (string, string, error) {
	Logger.Println(LogPrefix + "creating a CA")
	var identities []blockchainv3.ConfigCARegistryIdentitiesItem
	svc, err := service.NewConfigCARegistryIdentitiesItem("admin", "adminpw", "client")
	if err != nil {
		Logger.Println(LogPrefix+"**ERROR** - problem with NewConfigCARegistryIdentitiesItem: ", err)
		return "", "", err
	}
	roles := "*"
	svc.Attrs = &blockchainv3.IdentityAttrs{
		HfRegistrarRoles:      &roles,
		HfRegistrarAttributes: &roles,
	}
	identities = append(identities, *svc)

	registry, err := service.NewConfigCARegistry(-1, identities)
	if err != nil {
		Logger.Println(LogPrefix+"**ERROR** - problem with NewConfigCARegistry: ", err)
		return "", "", err
	}
	caConfigCreate, err := service.NewConfigCACreate(registry)
	if err != nil {
		Logger.Println(LogPrefix+"**ERROR** - problem with NewConfigCACreate: ", err)
		return "", "", err
	}
	configOverride, err := service.NewCreateCaBodyConfigOverride(caConfigCreate)
	if err != nil {
		Logger.Println(LogPrefix+"**ERROR** - problem with NewCreateCaBodyConfigOverride: ", err)
		return "", "", err
	}
	opts := service.NewCreateCaOptions(displayName, configOverride)
	result, _, err := service.CreateCa(opts)
	if err != nil {
		Logger.Println(LogPrefix+"**ERROR** - problem creating CA: ", err)
		return "", "", err
	}
	Logger.Println(LogPrefix + "**SUCCESS** - CA created")
	Logger.Println(LogPrefix+"[DEBUG] CA's api url: ", *result.ApiURL)
	Logger.Println(LogPrefix+"[DEBUG] CA's ID: ", *result.ID)
	Logger.Println(LogPrefix+"[DEBUG] CA's DepComponentID: ", *result.DepComponentID)
	// as a last step, we'll wait on the CA to come up before allowing anything else to happen
	err = waitForCaToComeUp(*result.ApiURL)
	return *result.Msp.Component.TlsCert, *result.ApiURL, err
}

func GetDecodedTlsCert(ec string) ([]byte, error) {
	// decode the base64 string in the CA's MSP
	tlsCert, err := base64.StdEncoding.DecodeString(ec)
	if err != nil {
		Logger.Println(LogPrefix + "error copying the cert", err)
		return nil, err
	}
	return tlsCert, nil
}

func WriteFileToLocalDirectory(filename string, tlsCert []byte) error {
	// convert the tls cert from the newly created CA into a format that can be used to create a PEM file
	Logger.Println(LogPrefix + "creating pem file locally from the tls cert passed in")

	f, err := os.Create(filename)
	if err != nil {
		Logger.Println(LogPrefix + "**ERROR** - problem creating "+filename, err)
		return err
	}

	defer f.Close()

	// write out the decoded PEM to the file
	_, err = f.Write(tlsCert)
	if err != nil {
		Logger.Println(LogPrefix + "**ERROR** - problem writing out the decoded PEM file: ", err)
		return err
	}
	if err := f.Sync(); err != nil {
		Logger.Println(LogPrefix + "**ERROR** - problem during file sync: ", err)
		return err
	}
	return nil
}

func CreateClient(tlsCertFilePath, apiURL string) *lib.Client {
	// create a client config and enable it with TLS (using the tls cert from the CA's MSP)
	Logger.Println(LogPrefix + "creating the config to enroll the CA")
	cfg := &lib.ClientConfig{
		TLS: catls.ClientTLSConfig{
			Enabled:   true,
			CertFiles: []string{tlsCertFilePath},
		},
		CAName: "Org1 CA",
		URL:    apiURL,
	}

	// use the config to create the client
	client := &lib.Client{
		HomeDir: ".",
		Config:  cfg,
	}
	Logger.Println("**SUCCESS** - client created")
	return client
}

func EnrollCA(client *lib.Client) (*lib.EnrollmentResponse, error) {
	// use the client to enroll the CA
	Logger.Println(LogPrefix + "enrolling the CA admin")

	// create CA Enrollment request and enroll the CA
	req := &api.EnrollmentRequest{
		Type:   "x509",
		Name:   "admin",
		Secret: "adminpw",
	}
	resp, err := client.Enroll(req)
	if err != nil {
		Logger.Println(LogPrefix+"**ERROR** - failed to enroll with CA", err)
		return nil, err
	}
	Logger.Println(LogPrefix + "**SUCCESS** - CA enrolled without error")
	return resp, nil
}

func RegisterAndEnrollAdmin(enrollResp *lib.EnrollmentResponse, name, secret string, retries *int) (*lib.Identity, error) {
	Logger.Println(LogPrefix+"registering and enrolling ", name)
	req := &api.RegistrationRequest{Name: name, Secret: secret, Type: "admin"} // registers user with the name
	identity, err := enrollResp.Identity.RegisterAndEnroll(req)

	if err != nil {
		errorAsString := err.Error()
		if removeIdentityIfRegistered(name, errorAsString, enrollResp, retries) {
			return RegisterAndEnrollAdmin(enrollResp, name, secret, retries)
		}
		Logger.Println(LogPrefix+"**ERROR** - problem registering and enrolling "+name, err)
		return nil, err
	}
	Logger.Println(LogPrefix + "**SUCCESS** - " + name + " registered")
	return identity, nil
}

func RegisterAdmin(enrollResp *lib.EnrollmentResponse, identityType, name, secret string, retries *int) error {
	Logger.Println(LogPrefix + "registering peer admin")
	regReq := &api.RegistrationRequest{Name: name, Secret: secret, Type: identityType} // registers user with the name
	_, err := enrollResp.Identity.Register(regReq)
	if err != nil {
		errorAsString := err.Error()
		if removeIdentityIfRegistered(name, errorAsString, enrollResp, retries) {
			return RegisterAdmin(enrollResp, identityType, name, secret, retries)
		}
		Logger.Println(LogPrefix+"**ERROR** - problem registering "+name, err)
	}
	Logger.Println(LogPrefix + "**SUCCESS** - " + name + " admin was registered")
	return nil
}

func CreateOrImportMSP(tlsCert []byte, identity *lib.Identity, service *blockchainv3.BlockchainV3, displayName, mspID string) error {
	Logger.Println(LogPrefix + "creating/importing the msp definition for " + identity.GetName())
	tlsRootCerts := []string{string(tlsCert)}
	admins := []string{string(identity.GetECert().Cert())} // registers using the identity
	//mspID := strings.ToLower(strings.Join(strings.Fields(displayName), ""))
	Logger.Println(LogPrefix+"The MSP ID is: ", mspID)
	importMspOpts := service.NewImportMspOptions(mspID, displayName, tlsRootCerts)
	importMspOpts.SetAdmins(admins[:])
	_, _, err := service.ImportMsp(importMspOpts)
	if err != nil {
		Logger.Println(LogPrefix+"**ERROR** - problem importing MSP: ", err)
		return err
	}
	Logger.Println(LogPrefix + "**SUCCESS** - created/imported MSP definition")
	return nil
}

func CreateCryptoObject(apiUrl, enrollID, enrollSecret string, tlsCert []byte, identity *lib.Identity,
	service *blockchainv3.BlockchainV3) (*blockchainv3.CryptoObject, error) {
	Logger.Println(LogPrefix+"[DEBUG] - inside createCryptoObject - api url: ", apiUrl)
	caName := "ca"
	tlsName := "tlsca"
	caTlsCert := base64.StdEncoding.EncodeToString(tlsCert)

	parsedUrl, err := url.Parse(apiUrl)
	if err != nil {
		Logger.Println(LogPrefix+"**ERROR** - problem parsing the api url while creating an instance of cryptoObjectEnrollmentsCa: ", err)
		return nil, err
	}
	hostname := strings.Split(parsedUrl.Host, ":")[0]
	port, err := strconv.ParseFloat(parsedUrl.Port(), 64)
	if err != nil {
		Logger.Println(LogPrefix + "**ERROR** - problem getting the port from the url. url: ", apiUrl)
		return nil, err
	}

	// - create the arguments
	cryptoObjectEnrollmentCa := &blockchainv3.CryptoObjectEnrollmentCa{
		Host:         &hostname,
		Port:         &port,
		Name:         &caName,
		TlsCert:      &caTlsCert,
		EnrollID:     &enrollID,
		EnrollSecret: &enrollSecret,
	}

	cryptoObjectEnrollmentTlsca := &blockchainv3.CryptoObjectEnrollmentTlsca{
		Host:         &hostname,
		Port:         &port,
		Name:         &tlsName,
		TlsCert:      &caTlsCert,
		EnrollID:     &enrollID,
		EnrollSecret: &enrollSecret,
		CsrHosts:     nil,
	}

	cryptoEnrollmentComponent := &blockchainv3.CryptoEnrollmentComponent{
		Admincerts: []string{base64.StdEncoding.EncodeToString(identity.GetECert().Cert())}}

	// enroll the crypto object
	cryptoObjectEnrollment, err := service.NewCryptoObjectEnrollment(cryptoEnrollmentComponent, cryptoObjectEnrollmentCa, cryptoObjectEnrollmentTlsca)
	if err != nil {
		Logger.Println(LogPrefix+"**ERROR** - problem enrolling the crypto object: ", err)
		return nil, err
	}

	cryptoObject := &blockchainv3.CryptoObject{Enrollment: cryptoObjectEnrollment}
	return cryptoObject, nil
}

func CreatePeer(service *blockchainv3.BlockchainV3, cryptoObject *blockchainv3.CryptoObject) error {
	opts := service.NewCreatePeerOptions("org1msp", "Peer Org1", cryptoObject)
	_, _, err := service.CreatePeer(opts)
	if err != nil {
		Logger.Println(LogPrefix+"**ERROR** - problem creating the peer: ", err)
		return err
	}
	Logger.Println(LogPrefix + "**SUCCESS** - Peer Org1 created")
	return nil
}

func CreateOrderer(service *blockchainv3.BlockchainV3, cryptoObjectSlice []blockchainv3.CryptoObject) error {
	opts := service.NewCreateOrdererOptions("raft", "osmsp", "Ordering Service MSP", cryptoObjectSlice)
	_, _, err := service.CreateOrderer(opts)
	if err != nil {
		Logger.Println(LogPrefix+"**ERROR** - problem creating the orderer", err)
		return err
	}
	Logger.Println(LogPrefix + "**SUCCESS** - Ordering Org1 created")
	return nil
}

//----------------------------------------------------------------------------------------------
// Helper/Aux functions
//----------------------------------------------------------------------------------------------

func waitForCaToComeUp(apiUrl string) error {
	Logger.Println(LogPrefix + "waiting for the CA to come up")
	// first, set the tls config to allow unsafe responses - WARNING - DO NOT DO THIS IN PRODUCTION
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	start := time.Now()

	client := http.Client{Timeout: 5 * time.Second} // create a client to handle the Get requests - this allows us the timeout

	deadline := start.Add(10 * 60 * time.Second) // ten minutes

	withinDeadline := false // final check at the end

	for !time.Now().After(deadline) { // make sure we're not past our five minute deadline
		Logger.Println(LogPrefix + "CA's cainfo polled ")
		resp, err := client.Get(apiUrl + "/cainfo")
		if err != nil {
			if os.IsTimeout(err) {
				continue
			} else {
				Logger.Println(LogPrefix+"**ERROR** - problem reaching CA - Not a timeout: ", err)
				return err
			}
		} else if resp.StatusCode != 200 {
			Logger.Println(LogPrefix+"**ERROR** - problem received a status code other than 200 while polling the CA's cainfo: ", resp)
			return err
		} else {
			elapsedTime := time.Since(start)
			log.Printf(LogPrefix+"CA came up - elapsed time was %v: ", elapsedTime)
			withinDeadline = true
			break
		}
	}
	if !withinDeadline {
		Logger.Println(LogPrefix+"**ERROR** - problem - timed out waiting for the CA to come up. current wait time is seat at ", deadline)
		err := errors.New("timed out waiting for the CA to come up. current wait time is seat at \", deadline")
		return err
	}
	return nil
}

func removeIdentity(orgName string, enrollResp *lib.EnrollmentResponse) {
	rr := &api.RemoveIdentityRequest{
		ID:    orgName,
		Force: true,
	}
	ir, err := enrollResp.Identity.RemoveIdentity(rr)
	if err != nil {
		Logger.Println(LogPrefix+"**ERROR** - problem removing identity for ", orgName) // use log.Println here so it won't stop the script during cleanup
	}
	Logger.Println(LogPrefix+"**SUCCESS** - the identity for "+orgName+"was deleted. Response: ", ir)
}

func removeIdentityIfRegistered(name, errorAsString string, enrollResp *lib.EnrollmentResponse, retries *int) bool {
	wasRemoved := false
	if strings.Contains(errorAsString, "is already registered") && *retries < 3 { // already registered then remove the registration and try again
		Logger.Println(LogPrefix + "the identity " + name + " was already registered. trying again to remove it")
		*retries++
		removeIdentity(name, enrollResp) // dsh comment out this line to run it without it ever removing anything
		Logger.Println(LogPrefix + "**SUCCESS** - " + name + " identity was removed")
		wasRemoved = true
	}
	return wasRemoved
}
