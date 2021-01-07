package integration

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"github.com/IBM-Blockchain/ibp-go-sdk/blockchainv3"
	"github.com/IBM/go-sdk-core/v4/core"
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
	Logger    *log.Logger
)

// set default logger first in case the caller fails to set a logger
func init() {
	LogPrefix = "[IT_TEST] "
	Logger = log.New(os.Stdout, "", 0)
	Logger.SetPrefix(time.Now().Format("2006-01-02 15:04:05.000 MST " + LogPrefix))
}

func CreateCA(service *blockchainv3.BlockchainV3, displayName, username, password string) (string, string, error) {
	Logger.Println("creating a CA")
	var identities []blockchainv3.ConfigCARegistryIdentitiesItem
	svc, err := service.NewConfigCARegistryIdentitiesItem(username, password, "client")
	if err != nil {
		Logger.Println("**ERROR** - problem with NewConfigCARegistryIdentitiesItem: ", err)
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
		Logger.Println("**ERROR** - problem with NewConfigCARegistry: ", err)
		return "", "", err
	}
	caConfigCreate, err := service.NewConfigCACreate(registry)
	if err != nil {
		Logger.Println("**ERROR** - problem with NewConfigCACreate: ", err)
		return "", "", err
	}
	configOverride, err := service.NewCreateCaBodyConfigOverride(caConfigCreate)
	if err != nil {
		Logger.Println("**ERROR** - problem with NewCreateCaBodyConfigOverride: ", err)
		return "", "", err
	}
	opts := service.NewCreateCaOptions(displayName, configOverride)
	result, _, err := service.CreateCa(opts)
	if err != nil {
		Logger.Println("**ERROR** - problem creating CA: ", err)
		return "", "", err
	}
	Logger.Println("**SUCCESS** - CA created")

	// as a last step, we'll wait on the CA to come up before allowing anything else to happen
	err = waitForCaToComeUp(*result.ApiURL)
	return *result.Msp.Component.TlsCert, *result.ApiURL, err
}

// decode the base64 string in the CA's MSP
func GetDecodedTlsCert(ec string) ([]byte, error) {
	tlsCert, err := base64.StdEncoding.DecodeString(ec)
	if err != nil {
		Logger.Println("error copying the cert", err)
		return nil, err
	}
	return tlsCert, nil
}

// convert the tls cert from the newly created CA into a format that can be used to create a PEM file
func WriteFileToLocalDirectory(filename string, tlsCert []byte) error {
	Logger.Println("creating pem file locally from the tls cert passed in")

	f, err := os.Create(filename)
	if err != nil {
		Logger.Println("**ERROR** - problem creating "+filename, err)
		return err
	}

	defer f.Close()

	// write out the decoded PEM to the file
	_, err = f.Write(tlsCert)
	if err != nil {
		Logger.Println("**ERROR** - problem writing out the decoded PEM file: ", err)
		return err
	}
	if err := f.Sync(); err != nil {
		Logger.Println("**ERROR** - problem during file sync: ", err)
		return err
	}
	return nil
}

// create a client config and enable it with TLS (using the tls cert from the CA's MSP)
func CreateClient(tlsCertFilePath, apiURL, name string) *lib.Client {
	Logger.Println("creating the config to enroll the CA")
	cfg := &lib.ClientConfig{
		TLS: catls.ClientTLSConfig{
			Enabled:   true,
			CertFiles: []string{tlsCertFilePath},
		},
		CAName: name,
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

func EnrollCA(client *lib.Client, name, secret string) (*lib.EnrollmentResponse, error) {
	Logger.Println("enrolling the CA admin")

	// create CA Enrollment request and enroll the CA
	req := &api.EnrollmentRequest{
		Type:   "x509",
		Name:   name,
		Secret: secret,
	}
	resp, err := client.Enroll(req)
	if err != nil {
		Logger.Println("**ERROR** - failed to enroll with CA", err)
		return nil, err
	}
	Logger.Println("**SUCCESS** - CA enrolled without error")
	return resp, nil
}

func RegisterAndEnrollAdmin(enrollResp *lib.EnrollmentResponse, name, secret string, retries *int) (*lib.Identity, error) {
	Logger.Println("registering and enrolling ", name)
	req := &api.RegistrationRequest{Name: name, Secret: secret, Type: "admin"} // registers user with the name
	identity, err := enrollResp.Identity.RegisterAndEnroll(req)

	if err != nil {
		errorAsString := err.Error()
		if removeIdentityIfRegistered(name, errorAsString, enrollResp, retries) {
			return RegisterAndEnrollAdmin(enrollResp, name, secret, retries)
		}
		Logger.Println("**ERROR** - problem registering and enrolling "+name, err)
		return nil, err
	}
	Logger.Println("**SUCCESS** - " + name + " registered")
	return identity, nil
}

func RegisterAdmin(enrollResp *lib.EnrollmentResponse, identityType, name, secret string, retries *int) error {
	Logger.Println("registering admin for " + name)
	regReq := &api.RegistrationRequest{Name: name, Secret: secret, Type: identityType} // registers user with the name
	_, err := enrollResp.Identity.Register(regReq)
	if err != nil {
		errorAsString := err.Error()
		if removeIdentityIfRegistered(name, errorAsString, enrollResp, retries) {
			return RegisterAdmin(enrollResp, identityType, name, secret, retries)
		}
		Logger.Println("**ERROR** - problem registering "+name, err)
	}
	Logger.Println("**SUCCESS** - " + name + " admin was registered")
	return nil
}

func CreateOrImportMSP(tlsCert []byte, identity *lib.Identity, service *blockchainv3.BlockchainV3, displayName, mspID string) error {
	Logger.Println("creating/importing the msp definition for " + identity.GetName())
	tlsRootCerts := []string{string(tlsCert)}
	admins := []string{string(identity.GetECert().Cert())} // registers using the identity
	Logger.Println("The MSP ID is: ", mspID)
	importMspOpts := service.NewImportMspOptions(mspID, displayName, tlsRootCerts)
	importMspOpts.SetAdmins(admins[:])
	_, _, err := service.ImportMsp(importMspOpts)
	if err != nil {
		Logger.Println("**ERROR** - problem importing MSP: ", err)
		return err
	}
	Logger.Println("**SUCCESS** - created/imported MSP definition")
	return nil
}

func CreateCryptoObject(apiUrl, enrollID, enrollSecret string, tlsCert []byte, identity *lib.Identity,
	service *blockchainv3.BlockchainV3) (*blockchainv3.CryptoObject, error) {
	caName := "ca"
	tlsName := "tlsca"
	caTlsCert := base64.StdEncoding.EncodeToString(tlsCert)

	parsedUrl, err := url.Parse(apiUrl)
	if err != nil {
		Logger.Println("**ERROR** - problem parsing the api url while creating an instance of cryptoObjectEnrollmentsCa: ", err)
		return nil, err
	}
	hostname := strings.Split(parsedUrl.Host, ":")[0]
	port, err := strconv.ParseFloat(parsedUrl.Port(), 64)
	if err != nil {
		Logger.Println("**ERROR** - problem getting the port from the url. url: ", apiUrl)
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
		Logger.Println("**ERROR** - problem enrolling the crypto object: ", err)
		return nil, err
	}

	cryptoObject := &blockchainv3.CryptoObject{Enrollment: cryptoObjectEnrollment}
	return cryptoObject, nil
}

func CreatePeer(service *blockchainv3.BlockchainV3, cryptoObject *blockchainv3.CryptoObject, mspID, displayName string) error {
	opts := service.NewCreatePeerOptions(mspID, displayName, cryptoObject)
	_, _, err := service.CreatePeer(opts)
	if err != nil {
		Logger.Println("**ERROR** - problem creating the peer: ", err)
		return err
	}
	Logger.Println("**SUCCESS** - " + displayName + " created")
	return nil
}

func CreateOrderer(service *blockchainv3.BlockchainV3, cryptoObjectSlice []blockchainv3.CryptoObject, mspID, displayName string) error {
	opts := service.NewCreateOrdererOptions("raft", mspID, displayName, cryptoObjectSlice)
	_, _, err := service.CreateOrderer(opts)
	if err != nil {
		Logger.Println("**ERROR** - problem creating the orderer", err)
		return err
	}
	Logger.Println("**SUCCESS** - " + displayName + " created")
	return nil
}

func GetComponentData(service *blockchainv3.BlockchainV3, id string) (int, error) {
	opts := service.NewGetComponentOptions(id)
	_, detailedResponse, err := service.GetComponent(opts)
	if err != nil {
		Logger.Println("**ERROR** - problem getting component data", err)
		return 0, err
	}
	Logger.Println("Retrieved component data's detailed response status code: ", detailedResponse.StatusCode)
	return detailedResponse.StatusCode, err
}

func ImportCA(service *blockchainv3.BlockchainV3, displayName, apiUrl string, tlsCert []byte) (int, error) {
	caName := "ca"
	tlsName := "tlsca"
	caTlsCertString := base64.StdEncoding.EncodeToString(tlsCert)
	caTlsCert := []string{caTlsCertString}

	ca := &blockchainv3.ImportCaBodyMspCa{
		Name:      &caName,
		RootCerts: caTlsCert,
	}

	tlsca := &blockchainv3.ImportCaBodyMspTlsca{
		Name:      &tlsName,
		RootCerts: caTlsCert,
	}

	component := &blockchainv3.ImportCaBodyMspComponent{
		TlsCert: &caTlsCertString,
	}

	importCaBodyMsp := &blockchainv3.ImportCaBodyMsp{
		Ca:        ca,
		Tlsca:     tlsca,
		Component: component,
	}
	opts := service.NewImportCaOptions(displayName, apiUrl, importCaBodyMsp)
	_, detailedResponse, err := service.ImportCa(opts)
	if err != nil {
		Logger.Println("**ERROR**")
		return 0, err
	}
	Logger.Println("response statusCode:", detailedResponse.StatusCode)
	return detailedResponse.StatusCode, err
}

func RemoveImportedComponent(service *blockchainv3.BlockchainV3, id string) (int, error) {
	// Remove imported component
	opts := service.NewRemoveComponentOptions(id)
	_, detailedResponse, err := service.RemoveComponent(opts)
	if err != nil {
		Logger.Println("**ERROR**")
		return 0, err
	}
	Logger.Println("response statusCode:", detailedResponse.StatusCode)
	return detailedResponse.StatusCode, nil
}

func DeleteComponent(service *blockchainv3.BlockchainV3, id string) (int, error) {
	opts := service.NewDeleteComponentOptions(id)
	_, detailedResponse, err := service.DeleteComponent(opts)
	if err != nil {
		Logger.Println("**ERROR** - problem deleteing ", id)
		return 0, err
	}
	return detailedResponse.StatusCode, err
}

func UpdateCA(service *blockchainv3.BlockchainV3, id string, tlsCert []byte) (int, error) {
	origins := []string{"us-south"}
	cors, err := service.NewConfigCACors(true, origins)
	if err != nil {
		Logger.Println("**ERROR** - problem creating new configCACors", err)
		return 0, err
	}
	isTrue := true
	crlSizeLimit := 1025.0

	CATls, err := service.NewConfigCATls(string(tlsCert), string(tlsCert))
	if err != nil {
		Logger.Println("**ERROR** - problem creating new configCATls", err)
		return 0, nil
	}
	caConfigUpdate := blockchainv3.ConfigCAUpdate{
		Cors:         cors,
		Debug:        &isTrue,
		Crlsizelimit: &crlSizeLimit,
		Tls:          CATls,
	}
	configOverride, err := service.NewUpdateCaBodyConfigOverride(&caConfigUpdate)
	if err != nil {
		Logger.Println("**ERROR** - problem creating new updateCABodyConfigOverride", err)
		return 0, nil
	}
	opts := service.NewUpdateCaOptions(id)
	opts.SetConfigOverride(configOverride)
	//opts.SetResources(caBodyResources)
	_, detailedResponse, err := service.UpdateCa(opts)
	if err != nil {
		return 0, err
	}
	return detailedResponse.StatusCode, err
}

func EditDataAboutCA(service *blockchainv3.BlockchainV3, id string) (int, error) {
	tags := [4]string{"fabric-ca", "ibm_sass", "blue_team", "dev"}
	opts := service.NewEditCaOptions(id)
	opts.SetCaName("My Ca Edited")
	opts.SetTags(tags[:])
	_, detailedResponse, err := service.EditCa(opts)
	if err != nil {
		Logger.Println("**ERROR** - problem editating data about CA", err)
		return 0, err
	}
	Logger.Println("**SUCCESS** - edited data about a CA")
	return detailedResponse.StatusCode, err
}

func SubmitActionToCa(service *blockchainv3.BlockchainV3, id string) (int, error) {
	restart := true
	opts := &blockchainv3.CaActionOptions{
		ID: &id,
		Restart: &restart,
	}

	// Restart CA
	_, detailedResponse, err := service.CaAction(opts)
	if err != nil {
		Logger.Println("**ERROR** problem restarting CA (SubmitActionToCA API)", err)
		return 0, err
	}
	Logger.Println("**SUCCESS** - restarted CA (SubmitActionToCA)")
	return detailedResponse.StatusCode, err
}

func ImportAPeer(service *blockchainv3.BlockchainV3, displayName, grpcwpUrl, mspID string, tlsCert []byte) (int, error) {
	caName := "ca"
	rootCerts := []string{string(tlsCert)}
	ca := &blockchainv3.MspCryptoFieldCa{
		Name:      &caName,
		RootCerts: rootCerts,
	}

	tlsca := &blockchainv3.MspCryptoFieldTlsca{
		Name:      &caName,
		RootCerts: rootCerts,
	}

	component := &blockchainv3.MspCryptoFieldComponent{
		TlsCert:    &rootCerts[0],
		Ecert:      &rootCerts[0],
		AdminCerts: rootCerts,
	}

	// Create msp field
	msp := &blockchainv3.MspCryptoField{
		Ca: ca,                  // MspCryptoFieldCa
		Tlsca: tlsca,            // MspCryptoFieldTlsca
		Component: component,     // MspCryptoFieldComponent
	}

	// Import Peer
	opts := service.NewImportPeerOptions(
		displayName,
		grpcwpUrl,
		msp,
		mspID,
	)
	_, detailedResponse, err := service.ImportPeer(opts)
	if err != nil {
		Logger.Println("**ERROR** - problem importing a peer", err)
		return 0, err
	}
	Logger.Println("**SUCCESS** - imported a peer")
	return detailedResponse.StatusCode, err
}

func ConstructImportCABodyMsp() {
	// Construct an instance of the ImportCaBodyMspCa model
	importCaBodyMspCaModel := new(blockchainv3.ImportCaBodyMspCa)
	importCaBodyMspCaModel.Name = core.StringPtr("org1CA")
	importCaBodyMspCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

	// Construct an instance of the ImportCaBodyMspComponent model
	importCaBodyMspComponentModel := new(blockchainv3.ImportCaBodyMspComponent)
	importCaBodyMspComponentModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")

	// Construct an instance of the ImportCaBodyMspTlsca model
	importCaBodyMspTlscaModel := new(blockchainv3.ImportCaBodyMspTlsca)
	importCaBodyMspTlscaModel.Name = core.StringPtr("org1tlsCA")
	importCaBodyMspTlscaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

	// Construct an instance of the ImportCaBodyMsp model
	importCaBodyMspModel := new(blockchainv3.ImportCaBodyMsp)
	importCaBodyMspModel.Ca = importCaBodyMspCaModel
	importCaBodyMspModel.Tlsca = importCaBodyMspTlscaModel
	importCaBodyMspModel.Component = importCaBodyMspComponentModel
}

//----------------------------------------------------------------------------------------------
// Helper/Aux functions
//----------------------------------------------------------------------------------------------

func waitForCaToComeUp(apiUrl string) error {
	Logger.Println("waiting for the CA to come up")
	// first, set the tls config to allow unsafe responses - WARNING - DO NOT DO THIS IN PRODUCTION
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	start := time.Now()

	client := http.Client{Timeout: 5 * time.Second} // create a client to handle the Get requests - this allows us the timeout

	deadline := start.Add(10 * 60 * time.Second) // ten minutes

	withinDeadline := false // final check at the end

	for !time.Now().After(deadline) { // make sure we're not past our deadline
		Logger.Println("CA's cainfo polled ")
		resp, err := client.Get(apiUrl + "/cainfo")
		if err != nil {
			if os.IsTimeout(err) {
				time.Sleep(2 * time.Second) // let the polling loop rest a little
				continue
			} else {
				Logger.Println("**ERROR** - problem reaching CA - Not a timeout: ", err)
				return err
			}
		} else if resp.StatusCode != 200 {
			Logger.Println("**ERROR** - problem received a status code other than 200 while polling the CA's cainfo: ", resp)
			return err
		} else {
			elapsedTime := time.Since(start)
			log.Printf("CA came up - elapsed time was %v: ", elapsedTime)
			withinDeadline = true
			break
		}
	}
	if !withinDeadline {
		Logger.Println("**ERROR** - problem - timed out waiting for the CA to come up. current wait time is seat at ", deadline)
		err := errors.New("timed out waiting for the CA to come up. current wait time is seat at \", deadline")
		return err
	}
	return nil
}

func removeIdentity(orgName string, enrollResp *lib.EnrollmentResponse) {
	removeRequest := &api.RemoveIdentityRequest{
		ID:    orgName,
		Force: true,
	}
	removeIdentityResponse, err := enrollResp.Identity.RemoveIdentity(removeRequest)
	if err != nil {
		Logger.Println("**ERROR** - problem removing identity for ", orgName)
	}
	Logger.Println("**SUCCESS** - the identity for "+orgName+"was deleted. Response: ", removeIdentityResponse)
}

func removeIdentityIfRegistered(name, errorAsString string, enrollResp *lib.EnrollmentResponse, retries *int) bool {
	wasRemoved := false
	if strings.Contains(errorAsString, "is already registered") && *retries < 3 { // already registered then remove the registration and try again
		Logger.Println("the identity " + name + " was already registered. trying again to remove it")
		*retries++
		removeIdentity(name, enrollResp)
		Logger.Println("**SUCCESS** - " + name + " identity was removed")
		wasRemoved = true
	}
	return wasRemoved
}
