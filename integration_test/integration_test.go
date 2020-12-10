package main_test

import (
	"encoding/json"
	"net/http"

	"crypto/tls"
	//"encoding/json"
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
	org1CAName             = "Org1 CA"
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
	ItTest                 = "[IT_TEST] "
	//pemCertFilePath        = "./env/tmpCert.pem"
)

var (
	l                        func(...interface{}) // store log.Println here for easier coding
	e                        func(...interface{}) // store log.Fatalln here for easier coding
	file                     *os.File
	service                  *blockchainv3.BlockchainV3
	encodedTlsCert, caApiUrl string
)

type setupInformation struct {
	APIKey       string `json:"api_key"`
	IdentityURL  string `json:"identity_url"`
	MyServiceURL string `json:"my_service_url"` // service instance url
}

var _ = Describe("GOLANG SDK Integration Test", func() {
	//const (
	//	org1CAName             = "Org1 CA"
	//	//org1AdminName          = "org1admin"
	//	//org1AdminPassword      = "org1adminpw"
	//	//peerType               = "peer"
	//	//peer1AdminName         = "peer1"
	//	//peer1AdminPassword     = "peer1pw"
	//	//org1MSPDisplayName     = "Org1 MSP"
	//	//org1MSPID              = "org1msp"
	//	//osCAName               = "Ordering Service CA"
	//	//osAdminName            = "OSadmin"
	//	//osAdminPassword        = "OSadminpw"
	//	//ordererType            = "orderer"
	//	//orderer1Name           = "OS1"
	//	//orderer1Password       = "OS1pw"
	//	//orderer1MSPDisplayName = "Ordering Service MSP"
	//	//orderer1MSPID          = "osmsp"
	//	ItTest                 = "[IT_TEST] "
	//	//pemCertFilePath        = "./env/tmpCert.pem"
	//)
	//
	//var (
	//	l                        func(...interface{}) // store log.Println here for easier coding
	//	//e                        func(...interface{}) // store log.Fatalln here for easier coding
	//	file                     *os.File
	//	service                  *blockchainv3.BlockchainV3
	//	encodedTlsCert, caApiUrl string
	//)
	var _ = BeforeSuite(func() {
		// setup a logger
		file = createLogFile() // we need a file to write the logs to
		defer file.Close()
		setupLogger(file)

		// get global setup information from a file
		s := setupInformation{}
		getSetupInfo(&s)
		l(ItTest + "start")

		// create a blockchain service to work with
		service = createAService(s)

		// make sure everything is cleaned up
		l(ItTest + "delete any existing components in the cluster")
		deleteAllComponents(service)

		// if the CA is not given a moment the new CA might fail to come up
		l(ItTest + "wait 10 seconds to make sure that everything was deleted")
		time.Sleep(15 * time.Second)
	})

	var _ = AfterSuite(func() {
		//----------------------------------------------------------------------------------------------
		// Cleanup
		//----------------------------------------------------------------------------------------------
		l(ItTest + "finally, delete any existing components in the cluster")
		deleteAllComponents(service)
		l(ItTest + "**SUCCESS** - test completed")
	})

	Describe("Creating Org1 CA", func() {
		It("should successfully create a CA and return the api url and the tls cert", func() {
			// we'll create our first certificate authority
			encodedTlsCert, caApiUrl = createCA(service, org1CAName)
			Expect(encodedTlsCert).NotTo(Equal(""))
			Expect(caApiUrl).NotTo(Equal(""))
		})
	})
})

func createCA(service *blockchainv3.BlockchainV3, displayName string) (string, string) {
	l(ItTest + "creating a CA")
	var identities []blockchainv3.ConfigCARegistryIdentitiesItem
	svc, err := service.NewConfigCARegistryIdentitiesItem("admin", "adminpw", "client")
	if err != nil {
		e(ItTest+"**ERROR** - problem with NewConfigCARegistryIdentitiesItem: ", err)
	}
	roles := "*"
	svc.Attrs = &blockchainv3.IdentityAttrs{
		HfRegistrarRoles:      &roles,
		HfRegistrarAttributes: &roles,
	}
	identities = append(identities, *svc)

	registry, err := service.NewConfigCARegistry(-1, identities)
	if err != nil {
		e(ItTest+"**ERROR** - problem with NewConfigCARegistry: ", err)
	}
	caConfigCreate, err := service.NewConfigCACreate(registry)
	if err != nil {
		e(ItTest+"**ERROR** - problem with NewConfigCACreate: ", err)
	}
	configOverride, err := service.NewCreateCaBodyConfigOverride(caConfigCreate)
	if err != nil {
		e(ItTest+"**ERROR** - problem with NewCreateCaBodyConfigOverride: ", err)
	}
	opts := service.NewCreateCaOptions(displayName, configOverride)
	result, _, err := service.CreateCa(opts)
	if err != nil {
		e(ItTest+"**ERROR** - problem creating CA: ", err)
	}
	l(ItTest + "**SUCCESS** - CA created")
	l(ItTest+"[DEBUG] CA's api url: ", *result.ApiURL)
	l(ItTest+"[DEBUG] CA's ID: ", *result.ID)
	l(ItTest+"[DEBUG] CA's DepComponentID: ", *result.DepComponentID)
	// as a last step, we'll wait on the CA to come up before allowing anything else to happen
	waitForCaToComeUp(*result.ApiURL)
	return *result.Msp.Component.TlsCert, *result.ApiURL
}

func waitForCaToComeUp(apiUrl string) {
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
				e(ItTest+"**ERROR** - problem reaching CA - Not a timeout: ", err)
			}
		} else if resp.StatusCode != 200 {
			e(ItTest+"**ERROR** - problem received a status code other than 200 while polling the CA's cainfo: ", resp)
		} else {
			elapsedTime := time.Since(start)
			log.Printf(ItTest+"CA came up - elapsed time was %v: ", elapsedTime)
			withinDeadline = true
			break
		}
	}
	if !withinDeadline {
		e(ItTest+"**ERROR** - problem - timed out waiting for the CA to come up. current wait time is seat at ", deadline)
	}
}

//----------------------------------------------------------------------------------------------
// Setup and teardown functions
//----------------------------------------------------------------------------------------------

func createLogFile() *os.File {
	// get the timestamp to add to the log name
	t := getCurrentTimeFormatted()

	lp := "./logs/it_testing_" + t + ".log"
	file, err := os.OpenFile(lp, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func getCurrentTimeFormatted() string {
	t := time.Now()
	z, _ := t.Zone()
	return t.Format("2006 Jan _2 15:04:05") + " " + z
}

func setupLogger(file *os.File) {
	log.SetOutput(file)
	l = log.Println
	e = log.Fatalln
}

func getSetupInfo(setupInfo *setupInformation) {
	l(ItTest + "\n\n***********************************STARTING INTEGRATION TEST***********************************")
	l(ItTest + "reading in the setup information from dev.json")
	file, err := ioutil.ReadFile("./env/dev.json")
	if err != nil {
		e(ItTest+"**ERROR** - problem reading in the setup info: ", err)
	}

	err = json.Unmarshal(file, setupInfo)
	if err != nil {
		e(ItTest+"**ERROR** - problem unmarshalling the setup info: ", err)
	}
	l(ItTest + "**Success** - setup information transferred to the test")
}

func createAService(s setupInformation) *blockchainv3.BlockchainV3 {
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
		e(ItTest + "**ERROR** - problem creating an instance of blockchainv3")
	}
	l(ItTest + "**SUCCESS** - service created")
	return service
}

func deleteAllComponents(service *blockchainv3.BlockchainV3) {
	l(ItTest + "deleting all components")
	opts := service.NewDeleteAllComponentsOptions()
	_, _, err := service.DeleteAllComponents(opts)
	if err != nil {
		e(ItTest+"**ERROR** - problem deleting all components: ", err)
	}
	l(ItTest + "**SUCCESS** - all components were deleted")
}