package main

import (
	"fmt"

	"github.com/IBM-Blockchain/ibp-go-sdk/blockchainv2"
	"github.com/IBM/go-sdk-core/v4/core"
)

func main() {
	fmt.Println("start")

	// globals
	ApiKey := "api-key"
	IdentityUrl := "https://iam.test.cloud.ibm.com/identity/token"
	myserviceURL := "http://localhost:3000"

	// Create an authenticator.
	authenticator := &core.IamAuthenticator{
		ApiKey: ApiKey, // update field with your api key
		URL:    IdentityUrl,
	}
	/*authenticator := &core.BearerTokenAuthenticator{
	    BearerToken: "my IAM access token",                                  // alternatively update field with access token
	}*/

	// Create an instance of the "BlockchainV2Options"  struct.
	//  myserviceURL := "http://localhost:3000"  // update field with service instance url
	options := &blockchainv2.BlockchainV2Options{
		Authenticator: authenticator,
		URL:           myserviceURL,
	}

	// Create an instance of the "BlockchainV2" service client.
	service, err := blockchainv2.NewBlockchainV2(options)
	if err != nil {
		// handle error
	}

	// Example - List all components
	opts := service.NewListComponentsOptions()
	result, detailedResponse, err := service.ListComponents(opts)
	fmt.Println("result:", result)
	fmt.Println("response:", detailedResponse)

	// Create a CA
	// Create an authenticator
	authenticator = &core.IamAuthenticator{
		ApiKey: ApiKey,
		URL:    IdentityUrl,
	}

	// Create an instance of the "BlockchainV2Options" struct
	options = &blockchainv2.BlockchainV2Options{
		Authenticator: authenticator,
		URL:           myserviceURL,
	}

	// Create an instance of the "BlockchainV2" service client.
	service, err = blockchainv2.NewBlockchainV2(options)
	if err != nil {
		return
	}

	// Create CA
	var identities []blockchainv2.ConfigCARegistryIdentitiesItem
	svc, err := service.NewConfigCARegistryIdentitiesItem("admin", "password", "client")
	if err != nil {
		return //err
	}
	identities = append(identities, *svc)
	registry, registery_err := service.NewConfigCARegistry(-1, identities)
	if registery_err != nil {
		return
	}
	ca_config_create, ca_config_create_err := service.NewConfigCACreate(registry)
	if ca_config_create_err != nil {
		return
	}
	config_override, config_override_err := service.NewCreateCaBodyConfigOverride(ca_config_create)
	if config_override_err != nil {
		return
	}
	ca_opts := service.NewCreateCaOptions("My CA", config_override)
	ca_result, detailedResponse, err := service.CreateCa(ca_opts)
	fmt.Println("result:", ca_result)
	fmt.Println("response:", detailedResponse)

	// // REMOVE COMPONENTS BY Tag
	// // Create an authenticator
	// authenticator = &core.IamAuthenticator{
	//     ApiKey: ApiKey,
	//     URL: IdentityUrl,
	// }

	// // Create an instance of the "BlockchainV2Options" struct
	// options = &blockchainv2.BlockchainV2Options{
	//     Authenticator: authenticator,
	//     URL: myserviceURL,
	// }

	// // Create an instance of the "BlockchainV2" service client.
	// service, err = blockchainv2.NewBlockchainV2(options)
	// if err != nil {
	//     return
	// }

	// // Remove components by tag
	// opts2 := service.NewRemoveComponentsByTagOptions("msp")
	// result2, detailedResponse, err := service.RemoveComponentsByTag(opts2)
	// fmt.Println("api key - lcsharp", ApiKey)
	// fmt.Println("result:", result2)
	// fmt.Println("response:", detailedResponse)

	// // DELETE COMPONENTS BY Tag
	// // Create an authenticator
	// authenticator = &core.IamAuthenticator{
	//     ApiKey: ApiKey,
	//     URL: IdentityUrl,
	// }

	// // Create an instance of the "BlockchainV2Options" struct
	// options = &blockchainv2.BlockchainV2Options{
	//     Authenticator: authenticator,
	//     URL: myserviceURL,
	// }

	// // Create an instance of the "BlockchainV2" service client.
	// service, err = blockchainv2.NewBlockchainV2(options)
	// if err != nil {
	//     return
	// }

	// // Delete components by tag
	// opts3 := service.NewDeleteComponentsByTagOptions("fabric-ca")
	// result3, detailedResponse, err := service.DeleteComponentsByTag(opts3)
	// fmt.Println("result:", result3)
	// fmt.Println("response:", detailedResponse)

	// // DELETE ALL COMPONENTS
	// // Create an authenticatorr
	// authenticator = &core.IamAuthenticator{
	//     ApiKey: ApiKey,
	//     URL: IdentityUrl,
	// }

	// // Create an instance of the "BlockchainV2Options" struct
	// options = &blockchainv2.BlockchainV2Options{
	//     Authenticator: authenticator,
	//     URL: myserviceURL,
	// }

	// // Create an instance of the "BlockchainV2" service client.
	// service, err = blockchainv2.NewBlockchainV2(options)
	// if err != nil {
	//     return
	// }

	// // Delete all components
	// opts4 := service.NewDeleteAllComponentsOptions()
	// result4, detailedResponse, err := service.DeleteAllComponents(opts4)
	// fmt.Println("result:", result4)
	// fmt.Println("response:", detailedResponse)

	// Create an authenticator
	authenticator = &core.IamAuthenticator{
		ApiKey: ApiKey,
		URL:    IdentityUrl,
	}

	// Create an instance of the "BlockchainV2Options" struct
	options = &blockchainv2.BlockchainV2Options{
		Authenticator: authenticator,
		URL:           myserviceURL,
	}

	// Create an instance of the "BlockchainV2" service client.
	service, err = blockchainv2.NewBlockchainV2(options)
	if err != nil {
		return
	}

	boolean_false := false
	boolean_true := true
	type_float64 := float64(89999)
	debug_level := "debug"
	info_level := "info"
	clientLoggingSettings := &blockchainv2.LoggingSettingsClient{Enabled: &boolean_true, Level: &debug_level, UniqueName: &boolean_false}
	serverLoggingSettings := &blockchainv2.LoggingSettingsServer{Enabled: &boolean_true, Level: &info_level, UniqueName: &boolean_false}
	opts5 := service.NewEditSettingsOptions()
	opts5.SetMaxReqPerMin(float64(50))
	opts5.SetInactivityTimeouts(&blockchainv2.EditSettingsBodyInactivityTimeouts{Enabled: &boolean_false, MaxIdleTime: &type_float64})
	opts5.SetFileLogging(&blockchainv2.EditLogSettingsBody{Client: clientLoggingSettings, Server: serverLoggingSettings})
	opts5.SetFabricLcGetCcTimeoutMs(float64(350000))
	result5, detailedResponse, err := service.EditSettings(opts5)
	fmt.Println("result:", result5)
	fmt.Println("response:", detailedResponse)

	fmt.Println("err:", err)
	fmt.Println("done")
	return
}
