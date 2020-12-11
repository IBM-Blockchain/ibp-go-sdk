package integration_test

import (
	"encoding/json"
	"github.com/hyperledger/fabric-ca/lib"

	"github.com/IBM-Blockchain/ibp-go-sdk/blockchainv3"
	it "github.com/IBM-Blockchain/ibp-go-sdk/integration_test"
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
	//l = log.Println //GinkgoWriter
	colorRed           = "\033[31m" // for error messages
	org1CAName         = "Org1 CA"
	org1AdminName      = "org1admin"
	org1AdminPassword  = "org1adminpw"
	peerType           = "peer"
	peer1AdminName     = "peer1"
	peer1AdminPassword = "peer1pw"
	org1MSPDisplayName = "Org1 MSP"
	org1MSPID          = "org1msp"
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
	//l = fmt.Println
	l                        func(...interface{}) // store log.Println here for easier coding
	lp                       string
	file                     *os.File
	service                  *blockchainv3.BlockchainV3
	encodedTlsCert, caApiUrl string
	tlsCert                  []byte
	client                   *lib.Client
	org1EnrollResponse       *lib.EnrollmentResponse
	orgIdentity              *lib.Identity
	cryptoObject             *blockchainv3.CryptoObject
)

type setupInformation struct {
	APIKey       string `json:"api_key"`
	IdentityURL  string `json:"identity_url"`
	MyServiceURL string `json:"my_service_url"` // service instance url
}

var _ = BeforeSuite(func() {
	// initialize integration functions lib variables
	it.Logger = log.New(GinkgoWriter, "IT", log.Ldate)
	it.LogPrefix = "[IT_TEST] "

	// get global setup information from a file
	s := setupInformation{}
	err := getSetupInfo(&s)
	Expect(err).NotTo(HaveOccurred())
	it.Logger.Println(ItTest + "\n\n***********************************STARTING INTEGRATION TEST***********************************")

	// create a blockchain service to work with
	service, err = createAService(s)
	Expect(err).NotTo(HaveOccurred())

	// make sure everything is cleaned up
	it.Logger.Println(ItTest + "delete any existing components in the cluster")
	err = deleteAllComponents(service) // maybe a bit redundant here since we'll use the same code in the library and test below - todo sanity check this lcs
	Expect(err).NotTo(HaveOccurred())
	// if the CA is not given a moment the new CA might fail to come up
	it.Logger.Println(ItTest + "wait 10 seconds to make sure that everything was deleted")
	time.Sleep(15 * time.Second)
})

var _ = AfterSuite(func() {
	//----------------------------------------------------------------------------------------------
	// Cleanup
	//----------------------------------------------------------------------------------------------
	it.Logger.Println(ItTest + "finally, delete any existing components in the cluster")
	err := deleteAllComponents(service)
	Expect(err).NotTo(HaveOccurred())
	if err == nil {
		it.Logger.Println(ItTest + "**SUCCESS** - test completed")
	}
})

var _ = Describe("GOLANG SDK Integration Test", func() {
	Describe("Creating Org1 CA Components", func() {
		// we'll create our first certificate authority
		It("should successfully create a CA and return the api url and the tls cert", func() {
			cert, url, err := it.CreateCA(service, org1CAName)
			encodedTlsCert = cert	// initialize global variables
			caApiUrl = url
			it.Logger.Println("encodedTlsCert from CreateCA call: ", encodedTlsCert)
			Expect(err).NotTo(HaveOccurred())
			Expect(encodedTlsCert).NotTo(Equal(""))
			Expect(caApiUrl).To(ContainSubstring("https://"))
		})
		It("should decode and return the TLS cert passed in from the CA", func() {
			it.Logger.Println("encodedTlsCert in the GetDecodedTlsCert test: ", encodedTlsCert)
			resp, err := it.GetDecodedTlsCert(encodedTlsCert)
			tlsCert = resp
			Expect(err).NotTo(HaveOccurred())
			Expect(tlsCert).NotTo(Equal(nil))
		})
		It("should write the TLS cert to a pem file", func() {
			err := it.WriteFileToLocalDirectory(pemCertFilePath, tlsCert)
			Expect(err).NotTo(HaveOccurred())
		})
		It("should create a tls client to use to enroll the CA", func() {
			resp := it.CreateClient(pemCertFilePath, caApiUrl)
			client = resp
			Expect(client).NotTo(Equal(nil))
		})
		It("should enroll the CA using the client we just made", func() {
			resp, err := it.EnrollCA(client)
			org1EnrollResponse = resp
			Expect(err).NotTo(HaveOccurred())
			Expect(org1EnrollResponse).NotTo(BeNil()) // todo do better. this is a weak assertion - lcs
		})
		It("should register the admins for org 1", func() {
			retries := 1
			resp, err := it.RegisterAndEnrollAdmin(org1EnrollResponse, org1AdminName, org1AdminPassword, &retries)
			orgIdentity = resp
			Expect(err).NotTo(HaveOccurred())
			Expect(orgIdentity).NotTo(BeNil()) // todo fix this weak assumption - lcs
		})
		It("should register the peer1 admin", func() {
			retries := 1
			err := it.RegisterAdmin(org1EnrollResponse, peerType, peer1AdminName, peer1AdminPassword, &retries)
			Expect(err).NotTo(HaveOccurred())
		})
		It("should create/import the msp definition", func() {
			err := it.CreateOrImportMSP(tlsCert, orgIdentity, service, org1MSPDisplayName, org1MSPID)
			Expect(err).NotTo(HaveOccurred())
		})
		It("should create a crypto object", func() {
			resp, err := it.CreateCryptoObject(caApiUrl, peer1AdminName, peer1AdminPassword, tlsCert, orgIdentity, service)
			cryptoObject = resp
			Expect(err).NotTo(HaveOccurred())
			Expect(cryptoObject).NotTo(BeNil()) // todo make better assertions - lcs
		})
		It("should create peer org 1", func() {
			err := it.CreatePeer(service, cryptoObject)
			Expect(err).NotTo(HaveOccurred())
		})
	})

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

//----------------------------------------------------------------------------------------------
// Helper/Aux functions
//----------------------------------------------------------------------------------------------

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

//func setupLogger(file *os.File) {
//	log.SetOutput(file)
//	l = log.Println
//}

func getSetupInfo(setupInfo *setupInformation) error {
	it.Logger.Println(ItTest + "\n\n***********************************STARTING INTEGRATION TEST***********************************")
	it.Logger.Println(ItTest + "reading in the setup information from dev.json")
	file, err := ioutil.ReadFile("./env/dev.json")
	if err != nil {
		it.Logger.Println(ItTest+"**ERROR** - problem reading in the setup info: ", err)
		return err
	}

	err = json.Unmarshal(file, setupInfo)
	if err != nil {
		it.Logger.Println(ItTest+"**ERROR** - problem unmarshalling the setup info: ", err)
		return err
	}
	it.Logger.Println(ItTest + "**Success** - setup information transferred to the test")
	return nil
}

func createAService(s setupInformation) (*blockchainv3.BlockchainV3, error) {
	it.Logger.Println(ItTest + "creating a service")
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
		it.Logger.Println(colorRed, ItTest+"**ERROR** - problem creating an instance of blockchainv3")
		return nil, err
	}
	it.Logger.Println(ItTest + "**SUCCESS** - service created")
	return service, nil
}

func deleteAllComponents(service *blockchainv3.BlockchainV3) error {
	it.Logger.Println(ItTest + "deleting all components")
	opts := service.NewDeleteAllComponentsOptions()
	_, _, err := service.DeleteAllComponents(opts)
	if err != nil {
		it.Logger.Println(ItTest+"**ERROR** - problem deleting all components: ", err)
		return err
	}
	it.Logger.Println(ItTest + "**SUCCESS** - all components were deleted")
	return nil
}
