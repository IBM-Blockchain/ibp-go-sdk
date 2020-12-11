package main_test

import (
	"encoding/base64"
	"encoding/json"
	"github.com/hyperledger/fabric-ca/lib"
	catls "github.com/hyperledger/fabric-ca/lib/tls"
	"net/http"

	"crypto/tls"
	//"encoding/json"
	"errors"
	"github.com/IBM-Blockchain/ibp-go-sdk/blockchainv3"
	//it "github.com/IBM-Blockchain/ibp-go-sdk/integration_test"
	"github.com/IBM/go-sdk-core/v4/core"
	"io/ioutil"
	"log"

	//"github.com/hyperledger/fabric-ca/api"
	//"github.com/hyperledger/fabric-ca/lib"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	//"io/ioutil"
	//"log"
	//"net/http"
	"os"
	"time"
)

const (
	colorRed   = "\033[31m" // for error messages
	org1CAName = "Org1 CA"
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
	ItTest          = "[IT_TEST] "
	pemCertFilePath = "./env/tmpCert.pem"
)

var (
	l                        func(...interface{}) // store log.Println here for easier coding
	lp                       string
	file                     *os.File
	service                  *blockchainv3.BlockchainV3
	encodedTlsCert, caApiUrl string
	tlsCert                  []byte
	client                   *lib.Client
)

type setupInformation struct {
	APIKey       string `json:"api_key"`
	IdentityURL  string `json:"identity_url"`
	MyServiceURL string `json:"my_service_url"` // service instance url
}

var _ = Describe("GOLANG SDK Integration Test", func() {
	var _ = BeforeSuite(func() {
		// setup a logger
		var err error
		file, err = createLogFile() // we need a file to write the logs to
		Expect(err).NotTo(HaveOccurred())
		defer file.Close()
		setupLogger(file)

		// get global setup information from a file
		s := setupInformation{}
		err = getSetupInfo(&s)
		Expect(err).NotTo(HaveOccurred())
		l(ItTest + "start")

		// create a blockchain service to work with
		service, err = createAService(s)
		Expect(err).NotTo(HaveOccurred())

		// make sure everything is cleaned up
		l(ItTest + "delete any existing components in the cluster")
		err = deleteAllComponents(service)
		Expect(err).NotTo(HaveOccurred())
		// if the CA is not given a moment the new CA might fail to come up
		l(ItTest + "wait 10 seconds to make sure that everything was deleted")
		time.Sleep(15 * time.Second)
	})

	var _ = AfterSuite(func() {
		//----------------------------------------------------------------------------------------------
		// Cleanup
		//----------------------------------------------------------------------------------------------
		l(ItTest + "finally, delete any existing components in the cluster")
		err := deleteAllComponents(service)
		Expect(err).NotTo(HaveOccurred())
		if err == nil {
			l(ItTest + "**SUCCESS** - test completed")
		}
	})

	//----------------------------------------------------------------------------------------------
	// Create Org 1 and it's components
	//----------------------------------------------------------------------------------------------

	Describe("Creating Org1 CA Components", func() {
		// reopen the file for logging
		//var err error
		file, _ = os.OpenFile(lp, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		setupLogger(file)
		//Expect(err).NotTo(HaveOccurred())
		l(ItTest + "Inside the test now")
		// we'll create our first certificate authority
		It("should successfully create a CA and return the api url and the tls cert", func() {
			encodedTlsCert, caApiUrl, err := createCA(service, org1CAName)
			Expect(err).NotTo(HaveOccurred())
			Expect(encodedTlsCert).NotTo(Equal(""))
			Expect(caApiUrl).To(ContainSubstring("https://"))
		})
		It("should decode and return the TLS cert passed in from the CA", func() {
			tlsCert, err := getDecodedTlsCert(encodedTlsCert)
			Expect(err).NotTo(HaveOccurred())
			Expect(tlsCert).NotTo(Equal(nil))
		})
		It("should write the TLS cert to a pem file", func() {
			err := writeFileToLocalDirectory(pemCertFilePath, tlsCert)
			Expect(err).NotTo(HaveOccurred())
		})
		It("should create a tls client to use to enroll the CA", func() {
			client = createClient(pemCertFilePath, caApiUrl)
			Expect(client).NotTo(Equal(nil))
		})
		file.Close()
	})

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
})

func createCA(service *blockchainv3.BlockchainV3, displayName string) (string, string, error) {
	l(ItTest + "creating a CA")
	var identities []blockchainv3.ConfigCARegistryIdentitiesItem
	svc, err := service.NewConfigCARegistryIdentitiesItem("admin", "adminpw", "client")
	if err != nil {
		l(colorRed, ItTest+"**ERROR** - problem with NewConfigCARegistryIdentitiesItem: ", err)
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
		l(colorRed, ItTest+"**ERROR** - problem with NewConfigCARegistry: ", err)
		return "", "", err
	}
	caConfigCreate, err := service.NewConfigCACreate(registry)
	if err != nil {
		l(colorRed, ItTest+"**ERROR** - problem with NewConfigCACreate: ", err)
		return "", "", err
	}
	configOverride, err := service.NewCreateCaBodyConfigOverride(caConfigCreate)
	if err != nil {
		l(colorRed, ItTest+"**ERROR** - problem with NewCreateCaBodyConfigOverride: ", err)
		return "", "", err
	}
	opts := service.NewCreateCaOptions(displayName, configOverride)
	result, _, err := service.CreateCa(opts)
	if err != nil {
		l(colorRed, ItTest+"**ERROR** - problem creating CA: ", err)
		return "", "", err
	}
	l(ItTest + "**SUCCESS** - CA created")
	l(ItTest+"[DEBUG] CA's api url: ", *result.ApiURL)
	l(ItTest+"[DEBUG] CA's ID: ", *result.ID)
	l(ItTest+"[DEBUG] CA's DepComponentID: ", *result.DepComponentID)
	// as a last step, we'll wait on the CA to come up before allowing anything else to happen
	err = waitForCaToComeUp(*result.ApiURL)
	return *result.Msp.Component.TlsCert, *result.ApiURL, err
}

func getDecodedTlsCert(ec string) ([]byte, error) {
	// decode the base64 string in the CA's MSP
	tlsCert, err := base64.StdEncoding.DecodeString(ec)
	if err != nil {
		l(colorRed, ItTest+"error copying the cert", err)
		return nil, err
	}
	return tlsCert, nil
}

func writeFileToLocalDirectory(filename string, tlsCert []byte) error {
	// convert the tls cert from the newly created CA into a format that can be used to create a PEM file
	l(ItTest + "creating pem file locally from the tls cert passed in")

	f, err := os.Create(filename)
	if err != nil {
		l(colorRed, ItTest+"**ERROR** - problem creating "+filename, err)
		return err
	}

	defer f.Close()

	// write out the decoded PEM to the file
	_, err = f.Write(tlsCert)
	if err != nil {
		l(colorRed, ItTest+"**ERROR** - problem writing out the decoded PEM file: ", err)
		return err
	}
	if err := f.Sync(); err != nil {
		l(colorRed, ItTest+"**ERROR** - problem during file sync: ", err)
		return err
	}
	return nil
}

func createClient(tlsCertFilePath, apiURL string) *lib.Client {
	// create a client config and enable it with TLS (using the tls cert from the CA's MSP)
	l(ItTest + "creating the config to enroll the CA")
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
	l("**SUCCESS** - client created")
	return client
}

//----------------------------------------------------------------------------------------------
// Helper/Aux functions
//----------------------------------------------------------------------------------------------

func waitForCaToComeUp(apiUrl string) error {
	l(ItTest + "waiting for the CA to come up")
	// first, set the tls config to allow unsafe responses - WARNING - DO NOT DO THIS IN PRODUCTION
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	start := time.Now()

	client := http.Client{Timeout: 5 * time.Second} // create a client to handle the Get requests - this allows us the timeout

	deadline := start.Add(10 * 60 * time.Second) // ten minutes

	withinDeadline := false // final check at the end

	for !time.Now().After(deadline) { // make sure we're not past our five minute deadline
		l(ItTest + "CA's cainfo polled ")
		resp, err := client.Get(apiUrl + "/cainfo")
		if err != nil {
			if os.IsTimeout(err) {
				continue
			} else {
				l(colorRed, ItTest+"**ERROR** - problem reaching CA - Not a timeout: ", err)
				return err
			}
		} else if resp.StatusCode != 200 {
			l(colorRed, ItTest+"**ERROR** - problem received a status code other than 200 while polling the CA's cainfo: ", resp)
			return err
		} else {
			elapsedTime := time.Since(start)
			log.Printf(ItTest+"CA came up - elapsed time was %v: ", elapsedTime)
			withinDeadline = true
			break
		}
	}
	if !withinDeadline {
		l(colorRed, ItTest+"**ERROR** - problem - timed out waiting for the CA to come up. current wait time is seat at ", deadline)
		err := errors.New("timed out waiting for the CA to come up. current wait time is seat at \", deadline")
		return err
	}
	return nil
}

//----------------------------------------------------------------------------------------------
// Setup and teardown functions
//----------------------------------------------------------------------------------------------

func createLogFile() (*os.File, error) {
	// get the timestamp to add to the log name
	t := getCurrentTimeFormatted()

	lp = "./logs/it_testing_" + t + ".log"
	file, err := os.OpenFile(lp, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(colorRed, ItTest, err)
		return nil, err
	}

	return file, nil
}

func getCurrentTimeFormatted() string {
	t := time.Now()
	z, _ := t.Zone()
	return t.Format("2006 Jan _2 15:04:05") + " " + z
}

func setupLogger(file *os.File) {
	log.SetOutput(file)
	l = log.Println
}

func getSetupInfo(setupInfo *setupInformation) error {
	l(ItTest + "\n\n***********************************STARTING INTEGRATION TEST***********************************")
	l(ItTest + "reading in the setup information from dev.json")
	file, err := ioutil.ReadFile("./env/dev.json")
	if err != nil {
		l(colorRed, ItTest+"**ERROR** - problem reading in the setup info: ", err)
		return err
	}

	err = json.Unmarshal(file, setupInfo)
	if err != nil {
		l(colorRed, ItTest+"**ERROR** - problem unmarshalling the setup info: ", err)
		return err
	}
	l(ItTest + "**Success** - setup information transferred to the test")
	return nil
}

func createAService(s setupInformation) (*blockchainv3.BlockchainV3, error) {
	l(ItTest + "creating a service")
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
		l(colorRed, ItTest+"**ERROR** - problem creating an instance of blockchainv3")
		return nil, err
	}
	l(ItTest + "**SUCCESS** - service created")
	return service, nil
}

func deleteAllComponents(service *blockchainv3.BlockchainV3) error {
	l(ItTest + "deleting all components")
	opts := service.NewDeleteAllComponentsOptions()
	_, _, err := service.DeleteAllComponents(opts)
	if err != nil {
		l(colorRed, ItTest+"**ERROR** - problem deleting all components: ", err)
		return err
	}
	l(ItTest + "**SUCCESS** - all components were deleted")
	return nil
}
