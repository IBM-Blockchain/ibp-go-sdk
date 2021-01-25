package integration_test

import (
	"encoding/json"
	"errors"
	"github.com/IBM-Blockchain/ibp-go-sdk/blockchainv3"
	it "github.com/IBM-Blockchain/ibp-go-sdk/integration_test"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/hyperledger/fabric-ca/lib"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"log"
	"os"

	"time"
)

const (
	adminName     = "admin"
	adminPassword = "adminpw"
	org1CAName    = "Org1 CA"
	//org1CAId               = "org1ca"
	org1AdminName      = "org1admin"
	org1AdminPassword  = "org1adminpw"
	peerType           = "peer"
	peer1AdminName     = "peer1"
	peer1AdminPassword = "peer1pw"
	//peerOrg1Id             = "peerorg1"
	peerOrg1DisplayName    = "Peer Org1"
	org1MSPDisplayName     = "Org1 MSP"
	org1MSPID              = "org1msp"
	osCAName               = "Ordering Service CA"
	osCAId                 = "orderingserviceca"
	osCAImportedName       = "orderingserviceca_0"
	osAdminName            = "OSadmin"
	osAdminPassword        = "OSadminpw"
	ordererType            = "orderer"
	orderer1Name           = "OS1"
	orderer1Password       = "OS1pw"
	orderer1MSPDisplayName = "Ordering Service MSP"
	orderer1MSPID          = "osmsp"
	OsId                   = "orderingservicemsp"
	pemCertFilePath        = ".tlsca.pem"
	mspDirectory           = "./msp/"
	genericGrpcwpUrl       = "https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084"
	clusterName            = "paidcluster"
)

var (
	service                  *blockchainv3.BlockchainV3
	encodedTlsCert, caApiUrl string
	tlsCert                  []byte
	client                   *lib.Client
	org1EnrollResponse       *lib.EnrollmentResponse
	OS1EnrollResponse        *lib.EnrollmentResponse
	orgIdentity              *lib.Identity
	cryptoObject             *blockchainv3.CryptoObject
	cryptoObjectSlice        []blockchainv3.CryptoObject
	setupInfo                setupInformation
	orderer1Id               string
	org1CAId                 string
	org1PeerId               string
)

type setupInformation struct {
	APIKey       string `json:"IAM_API_KEY"`
	IdentityURL  string `json:"IAM_IDENTITY_URL"`
	MyServiceURL string `json:"IBP_SERVICE_INSTANCE_URL"` // service instance url
}

var _ = BeforeSuite(func() {
	// initialize integration functions lib variables
	//it.LogPrefix = "[IT_TEST] "
	it.Logger = log.New(GinkgoWriter, "", 0)
	it.Logger.SetPrefix(time.Now().Format("2006-01-02 15:04:05.000 MST " + it.LogPrefix))
	it.Logger.Println("\n\n***********************************STARTING INTEGRATION TEST***********************************")

	// setup the test environment
	err := setupTestEnv()
	Expect(err).NotTo(HaveOccurred())

	// create a blockchain service instance to work with
	service, err = createAService(setupInfo)
	Expect(err).NotTo(HaveOccurred())

	// make sure everything is cleaned up
	it.Logger.Println("delete any existing components in the cluster")
	err = deleteAllComponents(service) // maybe a bit redundant here since we'll use the same code in the library and test below - todo sanity check this lcs
	Expect(err).NotTo(HaveOccurred())
	// if the CA is not given a moment the new CA might fail to come up
	it.Logger.Println("wait 10 seconds to make sure that everything was deleted")
	time.Sleep(15 * time.Second)
})

var _ = AfterSuite(func() {
	//----------------------------------------------------------------------------------------------
	// Cleanup
	//----------------------------------------------------------------------------------------------
	it.Logger.Println("finally, delete any existing components in the cluster")
	err := deleteAllComponents(service)
	Expect(err).NotTo(HaveOccurred())
	err = deleteLocallyCreatedFiles()
	Expect(err).NotTo(HaveOccurred())
	if err == nil {
		it.Logger.Println("**SUCCESS** - test completed")
	} else {
		it.Logger.Println("***UNSUCCESSFUL*** one or more errors occurred")
	}
})

var _ = Describe("GOLANG SDK Integration Test", func() {
	Describe("Creating Org1 CA Components", func() {
		// we'll create our first certificate authority
		It("should successfully create Org1 CA and return the api url and the tls cert", func() {
			id, cert, url, err := it.CreateCA(service, org1CAName, adminName, adminPassword)
			org1CAId = id
			encodedTlsCert = cert // initialize global variables
			caApiUrl = url
			Expect(err).NotTo(HaveOccurred())
			Expect(encodedTlsCert).NotTo(Equal(""))
			Expect(caApiUrl).To(ContainSubstring("https://"))
		})
		It("should decode and return the TLS cert passed in from Org1 CA", func() {
			resp, err := it.GetDecodedTlsCert(encodedTlsCert)
			tlsCert = resp
			Expect(err).NotTo(HaveOccurred())
			Expect(tlsCert).NotTo(Equal(nil))
		})
		It("should write the TLS cert for Org1 to a pem file", func() {
			err := it.WriteFileToLocalDirectory(pemCertFilePath, tlsCert)
			Expect(err).NotTo(HaveOccurred())
		})
		It("should create a tls client to use to enroll the CA", func() {
			resp := it.CreateClient(pemCertFilePath, caApiUrl, org1CAName)
			client = resp
			Expect(client).NotTo(Equal(nil))
		})
		It("should enroll Org1 CA using the client we just made", func() {
			resp, err := it.EnrollCA(client, adminName, adminPassword)
			org1EnrollResponse = resp
			Expect(err).NotTo(HaveOccurred())
			Expect(org1EnrollResponse).NotTo(BeNil()) // todo do better. this is a weak assertion - lcs
		})
		It("should register the admins for Org1", func() {
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
		It("should create/import the msp definition for Org1 flow", func() {
			err := it.CreateOrImportMSP(tlsCert, orgIdentity, service, org1MSPDisplayName, org1MSPID)
			Expect(err).NotTo(HaveOccurred())
		})
		It("should create a crypto object for Org1 flow", func() {
			resp, err := it.CreateCryptoObject(caApiUrl, peer1AdminName, peer1AdminPassword, tlsCert, orgIdentity, service)
			cryptoObject = resp
			Expect(err).NotTo(HaveOccurred())
			Expect(cryptoObject).NotTo(BeNil()) // todo make better assertions - lcs
		})
		It("should create Peer Org1", func() {
			id, err := it.CreatePeer(service, cryptoObject, org1MSPID, peerOrg1DisplayName)
			org1PeerId = id
			Expect(err).NotTo(HaveOccurred())
		})
	})
	Describe("Creating Ordering Org Components", func() {
		// we'll create our first certificate authority
		It("should successfully create the Ordering Org CA and return the api url and the tls cert", func() {
			_, cert, url, err := it.CreateCA(service, osCAName, adminName, adminPassword)
			encodedTlsCert = cert // initialize global variables
			caApiUrl = url
			Expect(err).NotTo(HaveOccurred())
			Expect(encodedTlsCert).NotTo(Equal(""))
			Expect(caApiUrl).To(ContainSubstring("https://"))
		})
		It("should decode and return the TLS cert passed in from the Ordering Org CA", func() {
			resp, err := it.GetDecodedTlsCert(encodedTlsCert)
			tlsCert = resp
			Expect(err).NotTo(HaveOccurred())
			Expect(tlsCert).NotTo(Equal(nil))
		})
		It("should write the TLS cert for Ordering Org1 CA to a pem file", func() {
			err := it.WriteFileToLocalDirectory(pemCertFilePath, tlsCert)
			Expect(err).NotTo(HaveOccurred())
		})
		It("should create a tls client to use to enroll the Ordering Org CA", func() {
			resp := it.CreateClient(pemCertFilePath, caApiUrl, osCAName)
			client = resp
			Expect(client).NotTo(Equal(nil))
		})
		It("should enroll the Ordering Org CA using the client we just made", func() {
			resp, err := it.EnrollCA(client, adminName, adminPassword)
			OS1EnrollResponse = resp
			Expect(err).NotTo(HaveOccurred())
			Expect(org1EnrollResponse).NotTo(BeNil()) // todo do better. this is a weak assertion - lcs
		})
		It("should register the admins for the Ordering Org", func() {
			retries := 1
			resp, err := it.RegisterAndEnrollAdmin(OS1EnrollResponse, osAdminName, osAdminPassword, &retries)
			orgIdentity = resp
			Expect(err).NotTo(HaveOccurred())
			Expect(orgIdentity).NotTo(BeNil()) // todo fix this weak assumption - lcs
		})
		It("should register the orderer admin", func() {
			retries := 1
			err := it.RegisterAdmin(OS1EnrollResponse, ordererType, orderer1Name, orderer1Password, &retries)
			Expect(err).NotTo(HaveOccurred())
		})
		It("should create/import the msp definition for the orderer flow", func() {
			err := it.CreateOrImportMSP(tlsCert, orgIdentity, service, orderer1MSPDisplayName, orderer1MSPID)
			Expect(err).NotTo(HaveOccurred())
		})
		It("should create a crypto object for the orderer flow", func() {
			resp, err := it.CreateCryptoObject(caApiUrl, orderer1Name, orderer1Password, tlsCert, orgIdentity, service)
			Expect(resp).NotTo(BeNil())
			if resp != nil {
				cryptoObjectSlice = []blockchainv3.CryptoObject{*resp}
				Expect(err).NotTo(HaveOccurred())
				Expect(cryptoObject).NotTo(BeNil()) // todo make better assertions - lcs
			} else {
				err = errors.New("**ERROR** - problem creating crypto object for the orderer flow")
			}
			Expect(err).NotTo(HaveOccurred())
		})
		It("should create Orderer 1", func() { // todo - uncomment this test when the unmarshalling issue is resolved lcs - 12/16/2020
			id, err := it.CreateOrderer(service, cryptoObjectSlice, orderer1MSPID, orderer1MSPDisplayName)
			Expect(id).NotTo(Equal(""))
			Expect(err).NotTo(HaveOccurred())
			orderer1Id = id
		})
	})
	Describe("Get component data", func() {
		It("should successfully get component data for the Org1 CA", func() {
			statusCode, err := it.GetComponentData(service, org1CAId)
			Expect(statusCode).To(Equal(200))
			Expect(err).NotTo(HaveOccurred())
		})
	})
	Describe("Import a CA", func() {
		It("should successfully import the Ordering Service CA", func() {
			statusCode, err := it.ImportCA(service, osCAName, caApiUrl, tlsCert)
			Expect(statusCode).To(Equal(200))
			Expect(err).NotTo(HaveOccurred())
		})
	})
	Describe("Remove Imported Component", func() {
		It("should successfully remove the imported Ordering Service CA", func() {
			statusCode, err := it.RemoveImportedComponent(service, osCAImportedName)
			Expect(statusCode).To(Equal(200))
			Expect(err).NotTo(HaveOccurred())
		})
	})
	Describe("Delete Component", func() {
		It("should successfully delete the Org1 CA", func() {
			statusCode, err := it.DeleteComponent(service, org1CAId)
			Expect(statusCode).To(Equal(200))
			Expect(err).NotTo(HaveOccurred())
		})
	})
	Describe("Update a CA", func() {
		It("should successfully update the Ordering Service CA", func() {
			statusCode, err := it.UpdateCA(service, osCAId, tlsCert)
			Expect(statusCode).To(Equal(200))
			Expect(err).NotTo(HaveOccurred())
		})
	})
	Describe("Edit Data about a CA", func() {
		It("should successfully edit data about the Ordering Service CA", func() {
			statusCode, err := it.EditDataAboutCA(service, osCAId)
			Expect(statusCode).To(Equal(200))
			Expect(err).NotTo(HaveOccurred())
		})
	})
	Describe("Submit Action to a CA", func() {
		It("should successfully restart the Ordering Service CA", func() {
			statusCode, err := it.SubmitActionToCa(service, osCAId)
			Expect(statusCode).To(Equal(202))
			Expect(err).NotTo(HaveOccurred())
		})
	})
	Describe("Import a Peer", func() {
		It("should successfully import PeerOrg 1", func() {
			statusCode, err := it.ImportAPeer(service, peerOrg1DisplayName, genericGrpcwpUrl, org1MSPID, tlsCert)
			Expect(statusCode).To(Equal(200))
			Expect(err).NotTo(HaveOccurred())
		})
	})
	Describe("Edit Data about a Peer", func() {
		It("should successfully edit data about the Org 1 Peer", func() {
			statusCode, err := it.EditDataAboutPeer(service, org1PeerId)
			Expect(statusCode).To(Equal(200))
			Expect(err).NotTo(HaveOccurred())
		})
	})
	Describe("Submit Action to a Peer", func() {
		It("should successfully restart the Org 1 Peer", func() {
			statusCode, err := it.SubmitActionToPeer(service, org1PeerId)
			Expect(statusCode).To(Equal(202))
			Expect(err).NotTo(HaveOccurred())
		})
	})
	/* not available on free clusters
	Describe("Update a Peer", func() {
		It("should successfully update the Org 1 Peer", func() {
			statusCode, err := it.UpdatePeer(service, org1PeerId)
			Expect(statusCode).To(Equal(200))
			Expect(err).NotTo(HaveOccurred())
		})
	})
	*/
	Describe("Import an Orderer", func() {
		It("should successfully import OS1", func() {
			statusCode, err := it.ImportAnOrderer(service, orderer1Name, genericGrpcwpUrl, orderer1MSPID, clusterName, tlsCert)
			Expect(statusCode).To(Equal(200))
			Expect(err).NotTo(HaveOccurred())
		})
	})
	Describe("Edit Data about an Orderer", func() {
		It("should successfully edit data about the OS1", func() {
			statusCode, err := it.EditDataAboutOrderer(service, orderer1Id)
			Expect(statusCode).To(Equal(200))
			Expect(err).NotTo(HaveOccurred())
		})
	})
	Describe("Submit Action to an Orderer", func() {
		It("should successfully restart Ordering Service MSP", func() {
			statusCode, err := it.SubmitActionToOrderer(service, orderer1Id)
			Expect(statusCode).To(Equal(202))
			Expect(err).NotTo(HaveOccurred())
		})
	})
	/* not available on free clusters
	Describe("Update an Orderer", func() {
		It("should successfully update the Ordering Service MSP", func() {
			statusCode, err := it.UpdateOrderer(service, orderer1Id)
			Expect(statusCode).To(Equal(200))
			Expect(err).NotTo(HaveOccurred())
		})
	})
	*/
})

//----------------------------------------------------------------------------------------------
// Setup and teardown functions
//----------------------------------------------------------------------------------------------

func setupTestEnv() error {
	it.Logger.Println("setting up the test env")
	setupInfo = setupInformation{}

	// setup private information the test needs from the env
	it.Logger.Println("read in the variables from the env")
	getSetupInfoFromEnv()

	// verify that we got the setup info
	err := verifySetupInfo()
	if err != nil { // failed to read the info from the env. let's try a file
		it.Logger.Println("failed to initialize the setup variables from the env. let's try a file")
		err = getSetupInfoFromFile(&setupInfo)
	}
	return err
}

func getSetupInfoFromEnv() {
	setupInfo.APIKey = os.Getenv("IAM_API_KEY")
	setupInfo.IdentityURL = os.Getenv("IAM_IDENTITY_URL")
	setupInfo.MyServiceURL = os.Getenv("IBP_SERVICE_INSTANCE_URL")
}

func verifySetupInfo() error {
	if setupInfo.APIKey == "" || setupInfo.IdentityURL == "" || setupInfo.MyServiceURL == "" {
		return errors.New("failed to initialize the setup variables from the env")
	}
	return nil
}

func getSetupInfoFromFile(setupInfo *setupInformation) error {
	it.Logger.Println("reading in the setup information from dev.json")
	file, err := ioutil.ReadFile("./env/dev.json")
	if err != nil {
		it.Logger.Println("**ERROR** - problem reading setup variables file: ", err)
		return err
	}

	err = json.Unmarshal(file, setupInfo)
	if err != nil {
		it.Logger.Println("**ERROR** - problem unmarshalling the setup variables obtained from the file: ", err)
		return err
	}
	it.Logger.Println("**Success** - setup information transferred to the test")
	return nil
}

func createAService(s setupInformation) (*blockchainv3.BlockchainV3, error) {
	it.Logger.Println("creating a service")
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
		it.Logger.Println("**ERROR** - problem creating an instance of blockchainv3")
		return nil, err
	}
	it.Logger.Println("**SUCCESS** - service created")
	return service, nil
}

func deleteAllComponents(service *blockchainv3.BlockchainV3) error {
	it.Logger.Println("deleting all components")
	opts := service.NewDeleteAllComponentsOptions()
	_, _, err := service.DeleteAllComponents(opts)
	if err != nil {
		it.Logger.Println("**ERROR** - problem deleting all components: ", err)
		return err
	}
	it.Logger.Println("**SUCCESS** - all components were deleted")
	return nil
}

func deleteLocallyCreatedFiles() error {
	it.Logger.Println("deleting locally created files (cert stores, etc)")
	err := os.Remove(pemCertFilePath)
	if err != nil {
		it.Logger.Println("**ERROR** - problem removing the pemCert created during the test", err)
	}
	err = os.RemoveAll(mspDirectory)
	if err != nil {
		it.Logger.Println("**ERROR** - problem removing the '/msp/' directory created during the test", err)
	}
	return nil
}
