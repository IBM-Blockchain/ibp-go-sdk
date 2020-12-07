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

const ItTest = "[IT_TEST] "

type setupInformation struct {
	APIKey       string `json:"api_key"`
	IdentityURL  string `json:"identity_url"`
	MyServiceURL string `json:"my_service_url"` // service instance url
}

func getSetupInfo(setupInfo *setupInformation) {
	log.Println(ItTest + "reading in the setup information from dev.json")
	file, err := ioutil.ReadFile("./env/dev.json")
	if err != nil {
		log.Fatalln(ItTest+"**ERROR** - problem reading in the setup info: ", err)
	}

	err = json.Unmarshal(file, setupInfo)
	if err != nil {
		log.Fatalln(ItTest+"**ERROR** - problem unmarshalling the setup info: ", err)
	}
	log.Println(ItTest + "**Success** - setup information transferred to the test")
}

func main() {
	// get global setup information from a file
	s := setupInformation{}
	getSetupInfo(&s)

	log.Println(ItTest + "start")

	// create a blockchain service to work with
	service := createAService(s)

	// make sure everything is cleaned up
	log.Println(ItTest + "delete any existing components in the cluster")
	deleteAllComponents(service)

	// if the CA is not given a moment the new CA might fail to come up
	log.Println(ItTest + "wait 10 seconds to make sure that everything was deleted")
	time.Sleep(10 * time.Second)

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
	org1EnrollResponse := enrollCA(client)

	// register the admins for org 1
	numRetries := 1
	orgIdentity := registerAndEnrollAdmin(org1EnrollResponse, "org1", &numRetries)
	numRetries = 1
	_ = registerAdmin(org1EnrollResponse, "peer", "peer", &numRetries)

	// create/import the msp definition
	createOrImportMSP(tlsCert, orgIdentity, service, "Org1 MSP", "org1msp")

	// create a crypto object
	cryptoObject := createCryptoObject(*caResult.ApiURL, "peer1", "peer1pw", tlsCert, orgIdentity, service)

	// create peer org 1
	createPeer(service, cryptoObject)

	//----------------------------------------------------------------------------------------------
	// Create Ordering Org and it's components
	//----------------------------------------------------------------------------------------------
	// we'll create our first certificate authority
	caResult = createCA(service,"Ordering Service CA")

	// Get TLS Cert
	tlsCert = getDecodedTlsCert(*caResult.Msp.Component.TlsCert)

	// filepath and name for the cert we're creating
	fp = "./env/tmpCert.pem"

	// write TLS Cert to a file
	writeFileToLocalDirectory(fp, tlsCert)

	// create a tls client to use to enroll the CA
	client = createClient(fp, *caResult.ApiURL)

	// enroll the CA using the client we just made
	OS1EnrollResponse := enrollCA(client)

	// register the admins for the ordering org
	numRetries = 1
	orgIdentity = registerAndEnrollAdmin(OS1EnrollResponse, "OS", &numRetries)
	numRetries = 1
	_ = registerAdmin(OS1EnrollResponse, "OS", "orderer", &numRetries)

	// create/import the msp definition
	createOrImportMSP(tlsCert, orgIdentity, service, "Ordering Service MSP", "osmsp")

	// create a crypto object
	cryptoObject = createCryptoObject(*caResult.ApiURL, "OS1", "OS1pw", tlsCert, orgIdentity, service)
	cryptoObjectSlice := []blockchainv3.CryptoObject{*cryptoObject}

	// create orderer
	createOrderer(service, cryptoObjectSlice)

	//----------------------------------------------------------------------------------------------
	// Cleanup
	//----------------------------------------------------------------------------------------------
	log.Println(ItTest + "finally, delete any existing components in the cluster")
	deleteAllComponents(service)
	//removeIdentity("org1admin", org1EnrollResponse)	// dsh - comment out these two calls to "removeIdentity" to run without removing the identities
	//removeIdentity("peer1", org1EnrollResponse)		// dsh - see below for one more spot
	//removeIdentity("OS1", OS1EnrollResponse)
	log.Println(ItTest + "**SUCCESS** - test completed")

}

func waitForCaToComeUp(apiUrl string) {
	log.Println(ItTest + "waiting for the CA to come up")
	// first, set the tls config to allow unsafe responses - WARNING - DO NOT DO THIS IN PRODUCTION
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	start := time.Now()

	client := http.Client{Timeout: 5 * time.Second} // create a client to handle the Get requests - this allows us the timeout

	deadline := start.Add(10 * 60 * time.Second) // ten minutes

	withinDeadline := false // final check at the end

	for !time.Now().After(deadline) { // make sure we're not past our five minute deadline
		log.Println(ItTest + "CA's cainfo polled ")
		resp, err := client.Get(apiUrl + "/cainfo")
		if err != nil {
			if os.IsTimeout(err) {
				continue
			} else {
				log.Fatalln(ItTest+"**ERROR** - problem reaching CA - Not a timeout: ", err)
			}
		} else if resp.StatusCode != 200 {
			log.Fatalln(ItTest+"**ERROR** - problem received a status code other than 200 while polling the CA's cainfo: ", resp)
		} else {
			elapsedTime := time.Since(start)
			log.Printf(ItTest+"CA came up - elapsed time was %v: ", elapsedTime)
			withinDeadline = true
			break
		}
	}
	if !withinDeadline {
		log.Fatalln(ItTest+"**ERROR** - problem - timed out waiting for the CA to come up. current wait time is seat at ", deadline)
	}
}

func deleteAllComponents(service *blockchainv3.BlockchainV3) {
	log.Println(ItTest + "deleting all components")
	opts := service.NewDeleteAllComponentsOptions()
	_, _, err := service.DeleteAllComponents(opts)
	if err != nil {
		log.Fatalln(ItTest+"**ERROR** - problem deleting all components: ", err)
	}
	log.Println(ItTest + "**SUCCESS** - all components were deleted")
}

func getDecodedTlsCert(ec string) []byte {
	// decode the base64 string in the CA's MSP
	tlsCert, err := base64.StdEncoding.DecodeString(ec)
	if err != nil {
		log.Fatalln(ItTest+"error copying the cert", err)
	}
	return tlsCert
}

func writeFileToLocalDirectory(filename string, tlsCert []byte) {
	// convert the tls cert from the newly created CA into a format that can be used to create a PEM file
	log.Println(ItTest + "creating pem file locally from the tls cert passed in")

	f, err := os.Create(filename)
	if err != nil {
		log.Fatalln(ItTest+"**ERROR** - problem creating the tempCert.pem file: ", err)
	}

	defer f.Close()

	// write out the decoded PEM to the file
	_, err = f.Write(tlsCert)
	if err != nil {
		log.Fatalln(ItTest+"**ERROR** - problem writing out the decoded PEM file: ", err)
	}
	if err := f.Sync(); err != nil {
		log.Fatalln(ItTest+"**ERROR** - problem during file sync: ", err)
	}
}

func createCA(service *blockchainv3.BlockchainV3, displayName string) *blockchainv3.CaResponse {
	log.Println(ItTest + "creating a CA")
	var identities []blockchainv3.ConfigCARegistryIdentitiesItem
	svc, err := service.NewConfigCARegistryIdentitiesItem("admin", "adminpw", "client")
	if err != nil {
		log.Fatalln(ItTest+"**ERROR** - problem with NewConfigCARegistryIdentitiesItem: ", err)
	}
	roles := "*"
	svc.Attrs = &blockchainv3.IdentityAttrs{
		HfRegistrarRoles:      &roles,
		HfRegistrarAttributes: &roles,
	}
	identities = append(identities, *svc)

	registry, err := service.NewConfigCARegistry(-1, identities)
	if err != nil {
		log.Fatalln(ItTest+"**ERROR** - problem with NewConfigCARegistry: ", err)
	}
	caConfigCreate, err := service.NewConfigCACreate(registry)
	if err != nil {
		log.Fatalln(ItTest+"**ERROR** - problem with NewConfigCACreate: ", err)
	}
	configOverride, err := service.NewCreateCaBodyConfigOverride(caConfigCreate)
	if err != nil {
		log.Fatalln(ItTest+"**ERROR** - problem with NewCreateCaBodyConfigOverride: ", err)
	}
	opts := service.NewCreateCaOptions(displayName, configOverride)
	result, _, err := service.CreateCa(opts)
	if err != nil {
		log.Fatalln(ItTest+"**ERROR** - problem creating CA: ", err)
	}
	log.Println(ItTest + "**SUCCESS** - CA created")

	// as a last step, we'll wait on the CA to come up before allowing anything else to happen
	waitForCaToComeUp(*result.ApiURL)
	return result
}

func createAService(s setupInformation) *blockchainv3.BlockchainV3 {
	log.Println(ItTest + "creating a service")
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
		log.Fatalln(ItTest + "**ERROR** - problem creating an instance of blockchainv3")
	}
	log.Println(ItTest + "**SUCCESS** - service created")
	return service
}

func createClient(tlsCertFilePath, apiURL string) *lib.Client {
	// create a client config and enable it with TLS (using the tls cert from the CA's MSP)
	log.Println(ItTest + "creating the config to enroll the CA")
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
	log.Println(ItTest + "enrolling the CA")

	// create CA Enrollment request
	req := &api.EnrollmentRequest{
		Type:   "x509",
		Name:   "admin",
		Secret: "adminpw",
	}
	resp, err := client.Enroll(req)
	if err != nil {
		log.Fatalln(ItTest+"failed to enroll with CA", err)
	}
	log.Println(ItTest + "**SUCCESS** - CA enrolled without error")
	return resp
}

func registerAndEnrollAdmin(enrollResp *lib.EnrollmentResponse, prefix string, retries *int) *lib.Identity {
	log.Println(ItTest + "registering and enrolling the organization admin")
	name := prefix + "admin"
	secret := prefix + "adminpw"
	req := &api.RegistrationRequest{Name: name, Secret: secret, Type: "admin"} // registers user with the name
	identity, err := enrollResp.Identity.RegisterAndEnroll(req)

	if err != nil {
		log.Fatalln(ItTest+"**ERROR** - problem registering and enrolling "+name, err)
	}
	log.Println(ItTest + "**SUCCESS** - " + name + " registered")
	return identity
}

func registerAdmin(enrollResp *lib.EnrollmentResponse, prefix, identityType string, retries *int) error {
	log.Println(ItTest + "registering peer admin")
	name := prefix + "1"                        // todo - abstract this part
	secret := prefix + "1pw"
	regReq := &api.RegistrationRequest{Name: name, Secret: secret, Type: identityType} // registers user with the name
	_, err := enrollResp.Identity.Register(regReq)
	if err != nil {
		log.Fatalln(ItTest+"**ERROR** - problem registering "+name, err)
	}
	log.Println(ItTest + "**SUCCESS** - " + name + " admin was registered")
	return nil
}

func createOrImportMSP(tlsCert []byte, identity *lib.Identity, service *blockchainv3.BlockchainV3, displayName, mspID string) {
	log.Println(ItTest + "creating/importing the msp definition for " + identity.GetName())
	tlsRootCerts := []string{string(tlsCert)}
	admins := []string{string(identity.GetECert().Cert())} // registers using the identity
	//mspID := strings.ToLower(strings.Join(strings.Fields(displayName), ""))
	log.Println(ItTest+"The MSP ID is: ", mspID)
	importMspOpts := service.NewImportMspOptions(mspID, displayName, tlsRootCerts)
	importMspOpts.SetAdmins(admins[:])
	_, _, err := service.ImportMsp(importMspOpts)
	if err != nil {
		log.Fatalln(ItTest+"**ERROR** - problem importing MSP: ", err)
	}
	log.Println(ItTest + "**SUCCESS** - created/imported MSP definition")
}

func createPeer(service *blockchainv3.BlockchainV3, cryptoObject *blockchainv3.CryptoObject) {
	opts := service.NewCreatePeerOptions("org1msp", "Peer Org1", cryptoObject)
	_, _, err := service.CreatePeer(opts)
	if err != nil {
		log.Fatalln(ItTest+"**ERROR** - problem creating the peer: ", err)
	}
	log.Println(ItTest + "**SUCCESS** - Peer Org1 created")
}

func createOrderer(service *blockchainv3.BlockchainV3, cryptoObjectSlice []blockchainv3.CryptoObject) {
	opts := service.NewCreateOrdererOptions("raft", "osmsp", "Ordering Service MSP", cryptoObjectSlice)
	_, _, err := service.CreateOrderer(opts)
	if err != nil {
		log.Fatalln(ItTest + "**ERROR** - problem creating the orderer", err)
	}
	log.Println(ItTest + "**SUCCESS** - Ordering Org1 created")
}

func createCryptoObject(apiUrl, enrollID, enrollSecret string, tlsCert []byte, identity *lib.Identity,
	service *blockchainv3.BlockchainV3) *blockchainv3.CryptoObject {
	caName := "ca"
	tlsName := "tlsca"
	caTlsCert := base64.StdEncoding.EncodeToString(tlsCert)

	parsedUrl, err := url.Parse(apiUrl)
	if err != nil {
		log.Fatalln(ItTest+"**ERROR** - problem parsing the api url while creating an instance of cryptoObjectEnrollmentsCa: ", err)
	}
	hostname := strings.Split(parsedUrl.Host, ":")[0]
	port, err := strconv.ParseFloat(parsedUrl.Port(), 64)
	if err != nil {
		log.Fatalln(ItTest + "**ERROR** - problem getting the port from the url")
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
		log.Fatalln(ItTest+"**ERROR** - problem enrolling the crypto object: ", err)
	}

	cryptoObject := &blockchainv3.CryptoObject{Enrollment: cryptoObjectEnrollment}
	return cryptoObject
}

//func removeIdentity(orgName string, enrollResp *lib.EnrollmentResponse) {
//	rr := &api.RemoveIdentityRequest{
//		ID:    orgName,
//		Force: true,
//	}
//	ir, err := enrollResp.Identity.RemoveIdentity(rr)
//	if err != nil {
//		log.Println(ItTest + "**ERROR** - problem removing identity for ", orgName) // use log.Println here so it won't stop the script during cleanup
//	}
//	log.Println(ItTest + "**SUCCESS** - the identity for " + orgName + "was deleted. Response: ", ir)
//}

//func removeIdentityIfRegistered(name, errorAsString string, enrollResp *lib.EnrollmentResponse, retries *int) bool {
//	wasRemoved := false
//	if strings.Contains(errorAsString, "is already registered") && *retries < 3 { // already registered then remove the registration and try again
//		log.Println(ItTest + "the identity " + name + " was already registered. trying again to remove it")
//		*retries++
//		//removeIdentity(name, enrollResp)											// dsh comment out this line to run it without it ever removing anything
//		log.Println(ItTest + "**SUCCESS** - " + name + " identity was removed")
//		wasRemoved = true
//	}
//	return wasRemoved
//}

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
