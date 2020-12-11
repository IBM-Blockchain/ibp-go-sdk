package integration

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	//"fmt"
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

const (
	//org1CAName             = "Org1 CA"
	//org1AdminName          = "org1admin"
	//org1AdminPassword      = "org1adminpw"
	//peerType               = "peer"
	//peer1AdminName         = "peer1"
	//peer1AdminPassword     = "peer1pw"
	//org1MSPDisplayName     = "Org1 MSP"
	//org1MSPID              = "org1msp"
	//osCAName               = "Ordering Service CA"
	//osAdminName            = "OSadmin"
	//osAdminPassword        = "OSadminpw"
	//ordererType            = "orderer"
	//orderer1Name           = "OS1"
	//orderer1Password       = "OS1pw"
	//orderer1MSPDisplayName = "Ordering Service MSP"
	//orderer1MSPID          = "osmsp"
	//LogPrefix                 = "[IT_TEST] "
	//pemCertFilePath        = "./env/tmpCert.pem"
)

var (
	LogPrefix string
	Logger *log.Logger
	l = Logger.Println
	e = Logger.Println
	//l func(...interface{}) // store log.Println here for easier coding
	//e func(...interface{}) // store log.Fatalln here for easier coding
	//file *os.File
	//service *blockchainv3.BlockchainV3
)

//type setupInformation struct {
//	APIKey       string `json:"api_key"`
//	IdentityURL  string `json:"identity_url"`
//	MyServiceURL string `json:"my_service_url"` // service instance url
//}

//func createLogFile() *os.File {
//	// get the timestamp to add to the log name
//	t := getCurrentTimeFormatted()
//
//	lp := "./logs/it_testing_" + t + ".log"
//	file, err := os.OpenFile(lp, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	return file
//}
//
//func setupLogger(file *os.File) {
//	log.SetOutput(file)
//	//l.Println()
//	//e = log.Fatalln
//}
//
//func getSetupInfo(setupInfo *setupInformation, l func(...interface{})) {
//	Logger.Println(LogPrefix + "\n\n***********************************STARTING INTEGRATION TEST***********************************")
//	Logger.Println(LogPrefix + "reading in the setup information from dev.json")
//	file, err := ioutil.ReadFile("./env/dev.json")
//	if err != nil {
//		Logger.Println(LogPrefix+"**ERROR** - problem reading in the setup info: ", err)
//	}
//
//	err = json.Unmarshal(file, setupInfo)
//	if err != nil {
//		Logger.Println(LogPrefix+"**ERROR** - problem unmarshalling the setup info: ", err)
//	}
//	Logger.Println(LogPrefix + "**Success** - setup information transferred to the test")
//}
//
//func getCurrentTimeFormatted() string {
//	t := time.Now()
//	z, _ := t.Zone()
//	return t.Format("2006 Jan _2 15:04:05") + " " + z
//}
//
//func createAService(s setupInformation, l, e func(...interface{})) *blockchainv3.BlockchainV3 {
//	Logger.Println(LogPrefix + "creating a service")
//	// Create an authenticator
//	authenticator := &core.IamAuthenticator{
//		ApiKey: s.APIKey,
//		URL:    s.IdentityURL,
//	}
//
//	// Create an instance of the "BlockchainV3Options" struct
//	options := &blockchainv3.BlockchainV3Options{
//		Authenticator: authenticator,
//		URL:           s.MyServiceURL,
//	}
//
//	// Create an instance of the "BlockchainV3" service client.
//	service, err := blockchainv3.NewBlockchainV3(options)
//	if err != nil {
//		Logger.Println(LogPrefix + "**ERROR** - problem creating an instance of blockchainv3")
//	}
//	Logger.Println(LogPrefix + "**SUCCESS** - service created")
//	return service
//}
//
//func deleteAllComponents(service *blockchainv3.BlockchainV3, l, e func(...interface{})) {
//	Logger.Println(LogPrefix + "deleting all components")
//	opts := service.NewDeleteAllComponentsOptions()
//	_, _, err := service.DeleteAllComponents(opts)
//	if err != nil {
//		Logger.Println(LogPrefix+"**ERROR** - problem deleting all components: ", err)
//	}
//	Logger.Println(LogPrefix + "**SUCCESS** - all components were deleted")
//}

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

func createOrderer(service *blockchainv3.BlockchainV3, cryptoObjectSlice []blockchainv3.CryptoObject) {
	opts := service.NewCreateOrdererOptions("raft", "osmsp", "Ordering Service MSP", cryptoObjectSlice)
	_, _, err := service.CreateOrderer(opts)
	if err != nil {
		Logger.Println(LogPrefix+"**ERROR** - problem creating the orderer", err)
	}
	Logger.Println(LogPrefix + "**SUCCESS** - Ordering Org1 created")
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

////---------------------------------------------------------------------------------------------------------------------
//// Create CA workaround
////---------------------------------------------------------------------------------------------------------------------
//// workaround if you don't want to keep creating CAs - useful for debugging this test by allowing a single run of the "Create CA" code above - otherwise comment it out
//caComponentID := "org1ca"
//getComponentOptions := &blockchainv3.GetComponentOptions{ID: &caComponentID}
//caResult, detailedResponse, err := service.GetComponent(getComponentOptions)
//Logger.Println(IT_TEST + "The returned CA: ", *caResult)
//// END OF WORKAROUND
////---------------------------------------------------------------------------------------------------------------------


//func main() {
//	// setup a logger
//	file = createLogFile() // we need a file to write the logs to
//	defer file.Close()
//	setupLogger(file)
//
//	// get global setup information from a file
//	s := setupInformation{}
//	getSetupInfo(&s)
//	Logger.Println(LogPrefix + "start")
//
//	// create a blockchain service to work with
//	service := createAService(s)
//
//	// make sure everything is cleaned up
//	Logger.Println(LogPrefix + "delete any existing components in the cluster")
//	deleteAllComponents(service)
//
//	// if the CA is not given a moment the new CA might fail to come up
//	Logger.Println(LogPrefix + "wait 10 seconds to make sure that everything was deleted")
//	time.Sleep(15 * time.Second)
//
//	//----------------------------------------------------------------------------------------------
//	// Create Org 1 and it's components
//	//----------------------------------------------------------------------------------------------
//	// we'll create our first certificate authority
//	encodedTlsCert, caApiUrl := createCA(service, org1CAName)
//
//	// Get TLS Cert
//	tlsCert := getDecodedTlsCert(encodedTlsCert)
//
//	// filepath and name for the cert we're creating
//	//pemCertFilePath := "./env/"
//
//	// write TLS Cert to a file
//	writeFileToLocalDirectory(pemCertFilePath, tlsCert)
//
//	// create a tls client to use to enroll the CA
//	client := createClient(pemCertFilePath, caApiUrl)
//
//	// enroll the CA using the client we just made
//	org1EnrollResponse := enrollCA(client)
//
//	// register the admins for org 1
//	retries := 1
//	orgIdentity := registerAndEnrollAdmin(org1EnrollResponse, org1AdminName, org1AdminPassword, &retries)
//	retries = 1
//	_ = registerAdmin(org1EnrollResponse, peerType, peer1AdminName, peer1AdminPassword, &retries)
//
//	// create/import the msp definition
//	createOrImportMSP(tlsCert, orgIdentity, service, org1MSPDisplayName, org1MSPID)
//
//	// create a crypto object
//	cryptoObject := createCryptoObject(caApiUrl, peer1AdminName, peer1AdminPassword, tlsCert, orgIdentity, service)
//
//	// create peer org 1
//	createPeer(service, cryptoObject)
//
//	//----------------------------------------------------------------------------------------------
//	// Create Ordering Org and it's components
//	//----------------------------------------------------------------------------------------------
//	// we'll create our first certificate authority
//	encodedTlsCert, caApiUrl = createCA(service, osCAName)
//
//	// Get TLS Cert
//	tlsCert = getDecodedTlsCert(encodedTlsCert)
//
//	// write TLS Cert to a file
//	writeFileToLocalDirectory(pemCertFilePath, tlsCert)
//
//	// create a tls client to use to enroll the CA
//	client = createClient(pemCertFilePath, caApiUrl)
//
//	// enroll the CA using the client we just made
//	OS1EnrollResponse := enrollCA(client)
//
//	// register the admins for the ordering org
//	retries = 1
//	orgIdentity = registerAndEnrollAdmin(OS1EnrollResponse, osAdminName, osAdminPassword, &retries)
//	retries = 1
//	_ = registerAdmin(OS1EnrollResponse, ordererType, orderer1Name, orderer1Password, &retries)
//
//	// create/import the msp definition
//	createOrImportMSP(tlsCert, orgIdentity, service, orderer1MSPDisplayName, orderer1MSPID)
//
//	// create a crypto object
//	cryptoObject = createCryptoObject(caApiUrl, orderer1Name, orderer1Password, tlsCert, orgIdentity, service)
//	cryptoObjectSlice := []blockchainv3.CryptoObject{*cryptoObject}
//
//	// create orderer
//	createOrderer(service, cryptoObjectSlice)
//
//	//----------------------------------------------------------------------------------------------
//	// Cleanup
//	//----------------------------------------------------------------------------------------------
//	Logger.Println(LogPrefix + "finally, delete any existing components in the cluster")
//	deleteAllComponents(service)
//	Logger.Println(LogPrefix + "**SUCCESS** - test completed")
//
//}