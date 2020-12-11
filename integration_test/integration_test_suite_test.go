package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestIntegrationTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IntegrationTest Suite")
}

//const (
//	org1CAName             = "Org1 CA"
//	org1AdminName          = "org1admin"
//	org1AdminPassword      = "org1adminpw"
//	peerType               = "peer"
//	peer1AdminName         = "peer1"
//	peer1AdminPassword     = "peer1pw"
//	org1MSPDisplayName     = "Org1 MSP"
//	org1MSPID              = "org1msp"
//	osCAName               = "Ordering Service CA"
//	osAdminName            = "OSadmin"
//	osAdminPassword        = "OSadminpw"
//	ordererType            = "orderer"
//	orderer1Name           = "OS1"
//	orderer1Password       = "OS1pw"
//	orderer1MSPDisplayName = "Ordering Service MSP"
//	orderer1MSPID          = "osmsp"
//	ItTest                 = "[IT_TEST] "
//	pemCertFilePath        = "./env/tmpCert.pem"
//)
//
//var (
//	l func(...interface{}) // store log.Println here for easier coding
//	e func(...interface{}) // store log.Fatalln here for easier coding
//	file *os.File
//	service *blockchainv3.BlockchainV3
//	encodedTlsCert, caApiUrl string
//)
//
//var _ = BeforeSuite(func() {
//	// setup a logger
//	file = createLogFile() // we need a file to write the logs to
//	defer file.Close()
//	setupLogger(file)
//
//	// get global setup information from a file
//	s := SetupInformation{}
//	getSetupInfo(&s)
//	l(ItTest + "start")
//
//	// create a blockchain service to work with
//	service = createAService(s)
//
//	// make sure everything is cleaned up
//	l(ItTest + "delete any existing components in the cluster")
//	deleteAllComponents(service)
//
//	// if the CA is not given a moment the new CA might fail to come up
//	l(ItTest + "wait 10 seconds to make sure that everything was deleted")
//	time.Sleep(15 * time.Second)
//
//	//----------------------------------------------------------------------------------------------
//	// Create Org 1 and it's components
//	//----------------------------------------------------------------------------------------------
//	By("creating Org1's CA", func() {
//		// we'll create our first certificate authority
//		encodedTlsCert, caApiUrl = createCA(service, org1CAName)
//		Expect(encodedTlsCert).NotTo(Equal(""))
//		Expect(caApiUrl).NotTo(Equal(""))
//	})
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
//})
//
//var _ = AfterSuite(func() {
//	//----------------------------------------------------------------------------------------------
//	// Cleanup
//	//----------------------------------------------------------------------------------------------
//	l(ItTest + "finally, delete any existing components in the cluster")
//	deleteAllComponents(service)
//	l(ItTest + "**SUCCESS** - test completed")
//})
//
//func setupLogger(file *os.File) {
//	log.SetOutput(file)
//	l = log.Println
//	e = log.Fatalln
//}
//
//func createLogFile() *os.File {
//	// get the timestamp to add to the log name
//	t := getCurrentTimeFormatted()
//
//	lp := "./logs/it_testing_" + t + ".log"
//	file, err := os.OpenFile(lp, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
//	Expect(err).NotTo(HaveOccurred())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	return file
//}
//
//func getCurrentTimeFormatted() string {
//	t := time.Now()
//	z, _ := t.Zone()
//	return t.Format("2006 Jan _2 15:04:05") + " " + z
//}
//
//type SetupInformation struct {
//	APIKey       string `json:"api_key"`
//	IdentityURL  string `json:"identity_url"`
//	MyServiceURL string `json:"my_service_url"` // service instance url
//}
//
//func getSetupInfo(setupInfo *SetupInformation) {
//	l(ItTest + "\n\n***********************************STARTING INTEGRATION TEST***********************************")
//	l(ItTest + "reading in the setup information from dev.json")
//	file, err := ioutil.ReadFile("./env/dev.json")
//	if err != nil {
//		e(ItTest+"**ERROR** - problem reading in the setup info: ", err)
//	}
//
//	err = json.Unmarshal(file, setupInfo)
//	if err != nil {
//		e(ItTest+"**ERROR** - problem unmarshalling the setup info: ", err)
//	}
//	l(ItTest + "**Success** - setup information transferred to the test")
//}
//
//func waitForCaToComeUp(apiUrl string) {
//	l(ItTest + "waiting for the CA to come up")
//	// first, set the tls config to allow unsafe responses - WARNING - DO NOT DO THIS IN PRODUCTION
//	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
//
//	start := time.Now()
//
//	client := http.Client{Timeout: 5 * time.Second} // create a client to handle the Get requests - this allows us the timeout
//
//	deadline := start.Add(10 * 60 * time.Second) // ten minutes
//
//	withinDeadline := false // final check at the end
//
//	for !time.Now().After(deadline) { // make sure we're not past our five minute deadline
//		l(ItTest + "CA's cainfo polled ")
//		resp, err := client.Get(apiUrl + "/cainfo")
//		if err != nil {
//			if os.IsTimeout(err) {
//				continue
//			} else {
//				e(ItTest+"**ERROR** - problem reaching CA - Not a timeout: ", err)
//			}
//		} else if resp.StatusCode != 200 {
//			e(ItTest+"**ERROR** - problem received a status code other than 200 while polling the CA's cainfo: ", resp)
//		} else {
//			elapsedTime := time.Since(start)
//			log.Printf(ItTest+"CA came up - elapsed time was %v: ", elapsedTime)
//			withinDeadline = true
//			break
//		}
//	}
//	if !withinDeadline {
//		e(ItTest+"**ERROR** - problem - timed out waiting for the CA to come up. current wait time is seat at ", deadline)
//	}
//}
//
//func deleteAllComponents(service *blockchainv3.BlockchainV3) {
//	l(ItTest + "deleting all components")
//	opts := service.NewDeleteAllComponentsOptions()
//	_, _, err := service.DeleteAllComponents(opts)
//	if err != nil {
//		e(ItTest+"**ERROR** - problem deleting all components: ", err)
//	}
//	l(ItTest + "**SUCCESS** - all components were deleted")
//}
//
//func getDecodedTlsCert(ec string) []byte {
//	// decode the base64 string in the CA's MSP
//	tlsCert, err := base64.StdEncoding.DecodeString(ec)
//	if err != nil {
//		e(ItTest+"error copying the cert", err)
//	}
//	return tlsCert
//}
//
//func writeFileToLocalDirectory(filename string, tlsCert []byte) {
//	// convert the tls cert from the newly created CA into a format that can be used to create a PEM file
//	l(ItTest + "creating pem file locally from the tls cert passed in")
//
//	f, err := os.Create(filename)
//	if err != nil {
//		e(ItTest+"**ERROR** - problem creating "+filename, err)
//	}
//
//	defer f.Close()
//
//	// write out the decoded PEM to the file
//	_, err = f.Write(tlsCert)
//	if err != nil {
//		e(ItTest+"**ERROR** - problem writing out the decoded PEM file: ", err)
//	}
//	if err := f.Sync(); err != nil {
//		e(ItTest+"**ERROR** - problem during file sync: ", err)
//	}
//}

//func createCA(service *blockchainv3.BlockchainV3, displayName string) (string, string) {
//	l(ItTest + "creating a CA")
//	var identities []blockchainv3.ConfigCARegistryIdentitiesItem
//	svc, err := service.NewConfigCARegistryIdentitiesItem("admin", "adminpw", "client")
//	if err != nil {
//		e(ItTest+"**ERROR** - problem with NewConfigCARegistryIdentitiesItem: ", err)
//	}
//	roles := "*"
//	svc.Attrs = &blockchainv3.IdentityAttrs{
//		HfRegistrarRoles:      &roles,
//		HfRegistrarAttributes: &roles,
//	}
//	identities = append(identities, *svc)
//
//	registry, err := service.NewConfigCARegistry(-1, identities)
//	if err != nil {
//		e(ItTest+"**ERROR** - problem with NewConfigCARegistry: ", err)
//	}
//	caConfigCreate, err := service.NewConfigCACreate(registry)
//	if err != nil {
//		e(ItTest+"**ERROR** - problem with NewConfigCACreate: ", err)
//	}
//	configOverride, err := service.NewCreateCaBodyConfigOverride(caConfigCreate)
//	if err != nil {
//		e(ItTest+"**ERROR** - problem with NewCreateCaBodyConfigOverride: ", err)
//	}
//	opts := service.NewCreateCaOptions(displayName, configOverride)
//	result, _, err := service.CreateCa(opts)
//	if err != nil {
//		e(ItTest+"**ERROR** - problem creating CA: ", err)
//	}
//	l(ItTest + "**SUCCESS** - CA created")
//	l(ItTest+"[DEBUG] CA's api url: ", *result.ApiURL)
//	l(ItTest+"[DEBUG] CA's ID: ", *result.ID)
//	l(ItTest+"[DEBUG] CA's DepComponentID: ", *result.DepComponentID)
//	// as a last step, we'll wait on the CA to come up before allowing anything else to happen
//	waitForCaToComeUp(*result.ApiURL)
//	return *result.Msp.Component.TlsCert, *result.ApiURL
//}
//
//func createAService(s SetupInformation) *blockchainv3.BlockchainV3 {
//	l(ItTest + "creating a service")
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
//		e(ItTest + "**ERROR** - problem creating an instance of blockchainv3")
//	}
//	l(ItTest + "**SUCCESS** - service created")
//	return service
//}
//
//func createClient(tlsCertFilePath, apiURL string) *lib.Client {
//	// create a client config and enable it with TLS (using the tls cert from the CA's MSP)
//	l(ItTest + "creating the config to enroll the CA")
//	cfg := &lib.ClientConfig{
//		TLS: catls.ClientTLSConfig{
//			Enabled:   true,
//			CertFiles: []string{tlsCertFilePath},
//		},
//		CAName: "Org1 CA",
//		URL:    apiURL,
//	}
//
//	// use the config to create the client
//	client := &lib.Client{
//		HomeDir: ".",
//		Config:  cfg,
//	}
//	l("**SUCCESS** - client created")
//	return client
//}
//
//func enrollCA(client *lib.Client) *lib.EnrollmentResponse {
//	// use the client to enroll the CA
//	l(ItTest + "enrolling the CA admin")
//
//	// create CA Enrollment request and enroll the CA
//	req := &api.EnrollmentRequest{
//		Type:   "x509",
//		Name:   "admin",
//		Secret: "adminpw",
//	}
//	resp, err := client.Enroll(req)
//	if resp == nil {
//		e(ItTest+"failed to enroll with CA", err)
//	}
//	l(ItTest + "**SUCCESS** - CA enrolled without error")
//	return resp
//}
//
//func registerAndEnrollAdmin(enrollResp *lib.EnrollmentResponse, name, secret string, retries *int) *lib.Identity {
//	l(ItTest+"registering and enrolling ", name)
//	req := &api.RegistrationRequest{Name: name, Secret: secret, Type: "admin"} // registers user with the name
//	identity, err := enrollResp.Identity.RegisterAndEnroll(req)
//
//	if err != nil {
//		errorAsString := err.Error()
//		if removeIdentityIfRegistered(name, errorAsString, enrollResp, retries) {
//			return registerAndEnrollAdmin(enrollResp, name, secret, retries)
//		}
//		e(ItTest+"**ERROR** - problem registering and enrolling "+name, err)
//	}
//	l(ItTest + "**SUCCESS** - " + name + " registered")
//	return identity
//}
//
//func registerAdmin(enrollResp *lib.EnrollmentResponse, identityType, name, secret string, retries *int) error {
//	l(ItTest + "registering peer admin")
//	regReq := &api.RegistrationRequest{Name: name, Secret: secret, Type: identityType} // registers user with the name
//	_, err := enrollResp.Identity.Register(regReq)
//	if err != nil {
//		errorAsString := err.Error()
//		if removeIdentityIfRegistered(name, errorAsString, enrollResp, retries) {
//			return registerAdmin(enrollResp, identityType, name, secret, retries)
//		}
//		e(ItTest+"**ERROR** - problem registering "+name, err)
//	}
//	l(ItTest + "**SUCCESS** - " + name + " admin was registered")
//	return nil
//}
//
//func createOrImportMSP(tlsCert []byte, identity *lib.Identity, service *blockchainv3.BlockchainV3, displayName, mspID string) {
//	l(ItTest + "creating/importing the msp definition for " + identity.GetName())
//	tlsRootCerts := []string{string(tlsCert)}
//	admins := []string{string(identity.GetECert().Cert())} // registers using the identity
//	//mspID := strings.ToLower(strings.Join(strings.Fields(displayName), ""))
//	l(ItTest+"The MSP ID is: ", mspID)
//	importMspOpts := service.NewImportMspOptions(mspID, displayName, tlsRootCerts)
//	importMspOpts.SetAdmins(admins[:])
//	_, _, err := service.ImportMsp(importMspOpts)
//	if err != nil {
//		e(ItTest+"**ERROR** - problem importing MSP: ", err)
//	}
//	l(ItTest + "**SUCCESS** - created/imported MSP definition")
//}
//
//func createPeer(service *blockchainv3.BlockchainV3, cryptoObject *blockchainv3.CryptoObject) {
//	opts := service.NewCreatePeerOptions("org1msp", "Peer Org1", cryptoObject)
//	_, _, err := service.CreatePeer(opts)
//	if err != nil {
//		e(ItTest+"**ERROR** - problem creating the peer: ", err)
//	}
//	l(ItTest + "**SUCCESS** - Peer Org1 created")
//}
//
//func createOrderer(service *blockchainv3.BlockchainV3, cryptoObjectSlice []blockchainv3.CryptoObject) {
//	opts := service.NewCreateOrdererOptions("raft", "osmsp", "Ordering Service MSP", cryptoObjectSlice)
//	_, _, err := service.CreateOrderer(opts)
//	if err != nil {
//		e(ItTest+"**ERROR** - problem creating the orderer", err)
//	}
//	l(ItTest + "**SUCCESS** - Ordering Org1 created")
//}
//
//func createCryptoObject(apiUrl, enrollID, enrollSecret string, tlsCert []byte, identity *lib.Identity,
//	service *blockchainv3.BlockchainV3) *blockchainv3.CryptoObject {
//	l(ItTest+"[DEBUG] - inside createCryptoObject - api url: ", apiUrl)
//	caName := "ca"
//	tlsName := "tlsca"
//	caTlsCert := base64.StdEncoding.EncodeToString(tlsCert)
//
//	parsedUrl, err := url.Parse(apiUrl)
//	if err != nil {
//		e(ItTest+"**ERROR** - problem parsing the api url while creating an instance of cryptoObjectEnrollmentsCa: ", err)
//	}
//	hostname := strings.Split(parsedUrl.Host, ":")[0]
//	port, err := strconv.ParseFloat(parsedUrl.Port(), 64)
//	if err != nil {
//		e(ItTest + "**ERROR** - problem getting the port from the url")
//	}
//
//	// - create the arguments
//	cryptoObjectEnrollmentCa := &blockchainv3.CryptoObjectEnrollmentCa{
//		Host:         &hostname,
//		Port:         &port,
//		Name:         &caName,
//		TlsCert:      &caTlsCert,
//		EnrollID:     &enrollID,
//		EnrollSecret: &enrollSecret,
//	}
//
//	cryptoObjectEnrollmentTlsca := &blockchainv3.CryptoObjectEnrollmentTlsca{
//		Host:         &hostname,
//		Port:         &port,
//		Name:         &tlsName,
//		TlsCert:      &caTlsCert,
//		EnrollID:     &enrollID,
//		EnrollSecret: &enrollSecret,
//		CsrHosts:     nil,
//	}
//
//	cryptoEnrollmentComponent := &blockchainv3.CryptoEnrollmentComponent{
//		Admincerts: []string{base64.StdEncoding.EncodeToString(identity.GetECert().Cert())}}
//
//	// enroll the crypto object
//	cryptoObjectEnrollment, err := service.NewCryptoObjectEnrollment(cryptoEnrollmentComponent, cryptoObjectEnrollmentCa, cryptoObjectEnrollmentTlsca)
//	if err != nil {
//		e(ItTest+"**ERROR** - problem enrolling the crypto object: ", err)
//	}
//
//	cryptoObject := &blockchainv3.CryptoObject{Enrollment: cryptoObjectEnrollment}
//	return cryptoObject
//}
//
//func removeIdentity(orgName string, enrollResp *lib.EnrollmentResponse) {
//	rr := &api.RemoveIdentityRequest{
//		ID:    orgName,
//		Force: true,
//	}
//	ir, err := enrollResp.Identity.RemoveIdentity(rr)
//	if err != nil {
//		l(ItTest+"**ERROR** - problem removing identity for ", orgName) // use log.Println here so it won't stop the script during cleanup
//	}
//	l(ItTest+"**SUCCESS** - the identity for "+orgName+"was deleted. Response: ", ir)
//}
//
//func removeIdentityIfRegistered(name, errorAsString string, enrollResp *lib.EnrollmentResponse, retries *int) bool {
//	wasRemoved := false
//	if strings.Contains(errorAsString, "is already registered") && *retries < 3 { // already registered then remove the registration and try again
//		l(ItTest + "the identity " + name + " was already registered. trying again to remove it")
//		*retries++
//		removeIdentity(name, enrollResp) // dsh comment out this line to run it without it ever removing anything
//		l(ItTest + "**SUCCESS** - " + name + " identity was removed")
//		wasRemoved = true
//	}
//	return wasRemoved
//}
//
//////---------------------------------------------------------------------------------------------------------------------
////// Create CA workaround
//////---------------------------------------------------------------------------------------------------------------------
////// workaround if you don't want to keep creating CAs - useful for debugging this test by allowing a single run of the "Create CA" code above - otherwise comment it out
////caComponentID := "org1ca"
////getComponentOptions := &blockchainv3.GetComponentOptions{ID: &caComponentID}
////caResult, detailedResponse, err := service.GetComponent(getComponentOptions)
////l(IT_TEST + "The returned CA: ", *caResult)
////// END OF WORKAROUND
//////---------------------------------------------------------------------------------------------------------------------