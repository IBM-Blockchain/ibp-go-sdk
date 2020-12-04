package main

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"github.com/hyperledger/fabric-ca/api"
	"github.com/hyperledger/fabric-ca/lib"
	catls "github.com/hyperledger/fabric-ca/lib/tls"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/IBM-Blockchain/ibp-go-sdk/blockchainv3"
	"github.com/IBM/go-sdk-core/v4/core"
)

const IT_TEST = "[IT_TEST] "

type setupInformation struct {
	APIKey       string `json:"api_key"`
	IdentityURL  string `json:"identity_url"`
	MyServiceURL string `json:"my_service_url"` // service instance url
}

func getSetupInfo(setupInfo *setupInformation) {
	log.Println(IT_TEST + "reading in the setup information from dev.json")
	file, err := ioutil.ReadFile("./env/dev.json")
	if err != nil {
		log.Fatalln(IT_TEST+"**ERROR** - problem reading in the setup info: ", err)
	}

	err = json.Unmarshal(file, setupInfo)
	if err != nil {
		log.Fatalln(IT_TEST+"**ERROR** - problem unmarshalling the setup info: ", err)
	}
	log.Println(IT_TEST + "**Success** - setup information transferred to the test")
}

func main() {
	// get global setup information from a file
	s := setupInformation{}
	getSetupInfo(&s)

	log.Println(IT_TEST + "start")

	// create a blockchain service to work with
	service := createAService(s)

	// make sure everything is cleaned up
	log.Println(IT_TEST + "delete any existing components in the cluster")
	deleteAllComponents(service)

	// if the CA is not given a moment the new CA might fail to come up
	log.Println(IT_TEST + "wait 5 seconds to make sure that everything was deleted")
	time.Sleep(5 * time.Second)

	//----------------------------------------------------------------------------------------------
	// Create Org 1 and it's components
	//----------------------------------------------------------------------------------------------
	// we'll create our first certificate authority
	caResult := createCA(service, "Org1 CA")

	// Get TLS Cert
	tlsCert := getDecodedTlsCert(*caResult.Msp.Component.TlsCert)

	// filepath and name for the cert we're creating
	fp := "./env/tmpCert.pem"

	// write TLS Cert to a file
	writeFileToLocalDirectory(fp, tlsCert)

	// create a tls client to use to enroll the CA
	client := createClient(fp, *caResult.ApiURL)

	// enroll the CA using the client we just made
	enrollResp := enrollCA(client)

	// register the admins for org 1
	numRetries := 1
	//orgIdentity := registeringTheAdmins(enrollResp, "1", &numRetries)
	orgIdentity := registerAndEnrollAdmin(enrollResp, "org1", &numRetries)
	numRetries = 1
	_ = registerAdmin(enrollResp, "peer", &numRetries)

	// create/import the msp definition
	createOrImportMSP(tlsCert, orgIdentity, service, "Org1 MSP")

	// create peer org 1
	createPeer(*caResult.ApiURL, "1", tlsCert, orgIdentity, service)

	////----------------------------------------------------------------------------------------------
	//// Create Ordering Org and it's components
	////----------------------------------------------------------------------------------------------
	//// we'll create our first certificate authority
	//caResult = createCA(service,"Ordering Service CA")
	//
	//// Get TLS Cert
	//tlsCert = getDecodedTlsCert(*caResult.Msp.Component.TlsCert)
	//
	//// filepath and name for the cert we're creating
	//fp = "./env/tmpCert.pem"
	//
	//// write TLS Cert to a file
	//writeFileToLocalDirectory(fp, tlsCert)
	//
	//// create a tls client to use to enroll the CA
	//client = createClient(fp, *caResult.ApiURL)
	//
	//// enroll the CA using the client we just made
	//enrollResp = enrollCA(client)
	//
	//// register the admins for org 2
	//numRetries = 1
	//orgIdentity = registerAndEnrollAdmin(enrollResp, "OS", &numRetries)
	//registerAdmin(enrollResp, "OS")
	////orgIdentity = registeringTheAdmins(enrollResp, "2", &numRetries)
	//
	////// create/import the msp definition
	////createOrImportMSP(tlsCert, orgIdentity, service, "2")
	////
	////// create peer org 2
	////createPeer(*caResult.ApiURL, "2", tlsCert, orgIdentity, service)

	log.Println(IT_TEST + "finally, delete any existing components in the cluster")
	deleteAllComponents(service)
	log.Println(IT_TEST + "**SUCCESS** - test completed with no known errors")

}

func waitForCaToComeUp(apiUrl string) {
	log.Println(IT_TEST + "waiting for the CA to come up")
	// first, set the tls config to allow unsafe responses - WARNING - DO NOT DO THIS IN PRODUCTION
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	start := time.Now()

	client := http.Client{Timeout: 5 * time.Second} // create a client to handle the Get requests - this allows us the timeout

	deadline := start.Add(10 * 60 * time.Second) // ten minutes

	withinDeadline := false // final check at the end

	for !time.Now().After(deadline) { // make sure we're not past our five minute deadline
		log.Println(IT_TEST + "CA's cainfo polled ")
		resp, err := client.Get(apiUrl + "/cainfo")
		if err != nil {
			if os.IsTimeout(err) {
				continue
			} else {
				log.Fatalln(IT_TEST+"**ERROR** - problem reaching CA - Not a timeout: ", err)
			}
		} else if resp.StatusCode != 200 {
			log.Fatalln(IT_TEST+"**ERROR** - problem recieved a status code other than 200 while polling the CA's cainfo: ", resp)
		} else {
			elapsedTime := time.Since(start)
			log.Printf(IT_TEST+"CA came up - elapsed time was %v: ", elapsedTime)
			withinDeadline = true
			break
		}
	}
	if !withinDeadline {
		log.Fatalln(IT_TEST+"**ERROR** - problem - timed out waiting for the CA to come up. current wait time is seat at ", deadline)
	}
}

func deleteAllComponents(service *blockchainv3.BlockchainV3) {
	log.Println(IT_TEST + "deleting all components")
	opts := service.NewDeleteAllComponentsOptions()
	_, _, err := service.DeleteAllComponents(opts)
	if err != nil {
		log.Fatalln(IT_TEST+"**ERROR** - problem deleting all components: ", err)
	}
	log.Println(IT_TEST + "**SUCCESS** - all components were deleted")
}

func getDecodedTlsCert(ec string) []byte {
	// decode the base64 string in the CA's MSP
	tlsCert, err := base64.StdEncoding.DecodeString(ec)
	if err != nil {
		log.Fatalln(IT_TEST+"error copying the cert", err)
	}
	return tlsCert
}

func writeFileToLocalDirectory(filename string, tlsCert []byte) {
	// convert the tls cert from the newly created CA into a format that can be used to create a PEM file
	log.Println(IT_TEST + "creating pem file locally from the tls cert passed in")

	f, err := os.Create(filename)
	if err != nil {
		log.Fatalln(IT_TEST+"**ERROR** - problem creating the tempCert.pem file: ", err)
	}

	defer f.Close()

	// write out the decoded PEM to the file
	_, err = f.Write(tlsCert)
	if err != nil {
		log.Fatalln(IT_TEST+"**ERROR** - problem writing out the decoded PEM file: ", err)
	}
	if err := f.Sync(); err != nil {
		log.Fatalln(IT_TEST+"**ERROR** - problem during file sync: ", err)
	}
}

func createCA(service *blockchainv3.BlockchainV3, displayName string) *blockchainv3.CaResponse {
	log.Println(IT_TEST + "creating a CA")
	var identities []blockchainv3.ConfigCARegistryIdentitiesItem
	svc, err := service.NewConfigCARegistryIdentitiesItem("admin", "adminpw", "client")
	if err != nil {
		log.Fatalln(IT_TEST+"**ERROR** - problem with NewConfigCARegistryIdentitiesItem: ", err)
	}
	roles := "*"
	svc.Attrs = &blockchainv3.IdentityAttrs{
		HfRegistrarRoles:      &roles,
		HfRegistrarAttributes: &roles,
	}
	identities = append(identities, *svc)

	registry, err := service.NewConfigCARegistry(-1, identities)
	if err != nil {
		log.Fatalln(IT_TEST+"**ERROR** - problem with NewConfigCARegistry: ", err)
	}
	caConfigCreate, err := service.NewConfigCACreate(registry)
	if err != nil {
		log.Fatalln(IT_TEST+"**ERROR** - problem with NewConfigCACreate: ", err)
	}
	configOverride, err := service.NewCreateCaBodyConfigOverride(caConfigCreate)
	if err != nil {
		log.Fatalln(IT_TEST+"**ERROR** - problem with NewCreateCaBodyConfigOverride: ", err)
	}
	opts := service.NewCreateCaOptions(displayName, configOverride)
	result, _, err := service.CreateCa(opts)
	if err != nil {
		log.Fatalln(IT_TEST+"**ERROR** - problem creating CA: ", err)
	}
	log.Println(IT_TEST + "**SUCCESS** - CA created")

	// as a last step, we'll wait on the CA to come up before allowing anything else to happen
	waitForCaToComeUp(*result.ApiURL)
	return result
}

func createAService(s setupInformation) *blockchainv3.BlockchainV3 {
	log.Println(IT_TEST + "creating a service")
	// Create an authenticator
	authenticator := &core.IamAuthenticator{
		ApiKey: s.APIKey,
		URL:    s.IdentityURL,
	}

	// Create an instance of the "BlockchainV3Options" struct
	options := &blockchainv3.BlockchainV3Options{
		Authenticator: authenticator,
		URL:           s.MyServiceURL,
	}

	// Create an instance of the "BlockchainV3" service client.
	service, err := blockchainv3.NewBlockchainV3(options)
	if err != nil {
		log.Fatalln(IT_TEST + "**ERROR** - problem creating an instance of blockchainv3")
	}
	log.Println(IT_TEST + "**SUCCESS** - service created")
	return service
}

func createClient(tlsCertFilePath, apiURL string) *lib.Client {
	// create a client config and enable it with TLS (using the tls cert from the CA's MSP)
	log.Println(IT_TEST + "creating the config to enroll the CA")
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
	log.Println("**SUCCESS** - client created")
	return client
}

func enrollCA(client *lib.Client) *lib.EnrollmentResponse {
	// use the client to enroll the CA
	log.Println(IT_TEST + "enrolling the CA")

	// create CA Enrollment request
	req := &api.EnrollmentRequest{
		Type:   "x509",
		Name:   "admin",
		Secret: "adminpw",
	}
	resp, err := client.Enroll(req)
	if err != nil {
		log.Fatalln(IT_TEST+"failed to enroll with CA", err)
	}
	log.Println(IT_TEST + "**SUCCESS** - CA enrolled without error")
	return resp
}

func registerAndEnrollAdmin(enrollResp *lib.EnrollmentResponse, prefix string, retries *int) *lib.Identity {
	log.Println(IT_TEST + "registering and enrolling the organization admin")
	name := prefix + "admin"
	secret := prefix + "adminpw"
	req := &api.RegistrationRequest{Name: name, Secret: secret, Type: "admin"} // registers user with the name
	identity, err := enrollResp.Identity.RegisterAndEnroll(req)

	if err != nil {
		errorAsString := err.Error()
		if strings.Contains(errorAsString, "is already registered") && *retries < 3 { // already registered then remove the registration and try again
			*retries++
			removeIdentity(name, enrollResp)
			log.Println(IT_TEST + "the identity " + name + " was already registered. trying again to remove it")
			return registerAndEnrollAdmin(enrollResp, prefix, retries)
		}
		log.Fatalln(IT_TEST+"**ERROR** - problem registering and enrolling "+name, err)
	}
	log.Println(IT_TEST + "**SUCCESS** - " + name + " registered")
	return identity
}

func registerAdmin(enrollResp *lib.EnrollmentResponse, prefix string, retries *int) error {
	log.Println(IT_TEST + "registering peer admin")
	name := prefix + "1"
	secret := prefix + "1pw"
	regReq := &api.RegistrationRequest{Name: name, Secret: secret, Type: "peer"} // registers user with the name
	_, err := enrollResp.Identity.Register(regReq)
	if err != nil {
		errorAsString := err.Error()
		if strings.Contains(errorAsString, "is already registered") && *retries < 3 { // already registered then remove the registration and try again
			*retries++
			removeIdentity(name, enrollResp)
			log.Println(IT_TEST + "the identity " + name + " was already registered. trying again to remove it")
			return registerAdmin(enrollResp, prefix, retries)
		}
		log.Fatalln(IT_TEST+"**ERROR** - problem registering "+name, err)
	}
	log.Println(IT_TEST + "**SUCCESS** - " + name + " admin was registered")
	return nil
}

func createOrImportMSP(tlsCert []byte, identity *lib.Identity, service *blockchainv3.BlockchainV3, displayName string) {
	log.Println(IT_TEST + "creating/importing the msp definition for " + identity.GetName())
	tlsRootCerts := []string{string(tlsCert)}
	admins := []string{string(identity.GetECert().Cert())} // registers using the identity
	mspID := strings.ToLower(strings.Join(strings.Fields(displayName), ""))
	log.Println(IT_TEST+"The MSP ID is: ", mspID)
	importMspOpts := service.NewImportMspOptions(mspID, displayName, tlsRootCerts)
	importMspOpts.SetAdmins(admins[:])
	_, _, err := service.ImportMsp(importMspOpts)
	if err != nil {
		log.Fatalln(IT_TEST+"**ERROR** - problem importing MSP: ", err)
	}
	log.Println(IT_TEST + "**SUCCESS** - created/imported MSP definition")
}

func createPeer(apiUrl, orgNumber string, tlsCert []byte, identity *lib.Identity, service *blockchainv3.BlockchainV3) {
	caName := "ca"
	tlsName := "tlsca"
	caTlsCert := base64.StdEncoding.EncodeToString(tlsCert)
	displayName := "Peer Org" + orgNumber
	enrollID := "peer" + orgNumber
	enrollSecret := "peer" + orgNumber + "pw"
	mspID := "org" + orgNumber + "msp"

	// create the 2nd argument
	// - gather the properties needed for the argument
	parsedUrl, err := url.Parse(apiUrl)
	if err != nil {
		log.Fatalln(IT_TEST+"**ERROR** - problem parsing the api url while creating an instance of cryptoObjectEnrollmentsCa: ", err)
	}
	hostname := strings.Split(parsedUrl.Host, ":")[0]
	port, err := strconv.ParseFloat(parsedUrl.Port(), 64)
	if err != nil {
		log.Fatalln(IT_TEST + "**ERROR** - problem getting the port from the url")
	}

	// - create the argument
	cryptoObjectEnrollmentCa := &blockchainv3.CryptoObjectEnrollmentCa{
		Host:         &hostname,
		Port:         &port,
		Name:         &caName,
		TlsCert:      &caTlsCert,
		EnrollID:     &enrollID,
		EnrollSecret: &enrollSecret,
	}

	// create the 3rd argument
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
		log.Fatalln(IT_TEST+"**ERROR** - problem enrolling the crypto object: ", err)
	}

	cryptoObject := &blockchainv3.CryptoObject{Enrollment: cryptoObjectEnrollment}
	createPeerOptions := service.NewCreatePeerOptions(mspID, displayName, cryptoObject)
	_, _, err = service.CreatePeer(createPeerOptions)
	if err != nil {
		log.Fatalln(IT_TEST+"**ERROR** - problem creating the peer: ", err)
	}
	log.Println(IT_TEST + "**SUCCESS** - peer " + displayName + " created")
}

func removeIdentity(orgName string, enrollResp *lib.EnrollmentResponse) {
	rr := &api.RemoveIdentityRequest{
		ID:    orgName,
		Force: true,
	}
	_, err2 := enrollResp.Identity.RemoveIdentity(rr)
	if err2 != nil {
		log.Fatalln(IT_TEST + "Failed to remove the identity")
	}
}

////---------------------------------------------------------------------------------------------------------------------
//// Create CA workaround
////---------------------------------------------------------------------------------------------------------------------
//// workaround if you don't want to keep creating CAs - useful for debugging this test by allowing a single run of the "Create CA" code above - otherwise comment it out
//caComponentID := "org1ca"
//getComponentOptions := &blockchainv3.GetComponentOptions{ID: &caComponentID}
//caResult, detailedResponse, err := service.GetComponent(getComponentOptions)
//log.Println(IT_TEST + "The returned CA: ", *caResult)
//// END OF WORKAROUND
////---------------------------------------------------------------------------------------------------------------------
