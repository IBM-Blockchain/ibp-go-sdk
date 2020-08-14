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

	// // Create a CA
	// // Create an authenticator
	// authenticator = &core.IamAuthenticator{
	// 	ApiKey: ApiKey,
	// 	URL:    IdentityUrl,
	// }

	// // Create an instance of the "BlockchainV2Options" struct
	// options = &blockchainv2.BlockchainV2Options{
	// 	Authenticator: authenticator,
	// 	URL:           myserviceURL,
	// }

	// // Create an instance of the "BlockchainV2" service client.
	// service, err = blockchainv2.NewBlockchainV2(options)
	// if err != nil {
	// 	return
	// }

	// // Create CA
	// var identities []blockchainv2.ConfigCARegistryIdentitiesItem
	// svc, err := service.NewConfigCARegistryIdentitiesItem("admin", "password", "client")
	// if err != nil {
	// 	return //err
	// }
	// identities = append(identities, *svc)
	// registry, err := service.NewConfigCARegistry(-1, identities)
	// if err != nil {
	// 	return
	// }
	// caConfigCreate, err := service.NewConfigCACreate(registry)
	// if err != nil {
	// 	return
	// }
	// configOverride, err := service.NewCreateCaBodyConfigOverride(caConfigCreate)
	// if err != nil {
	// 	return
	// }
	// caOpts := service.NewCreateCaOptions("My CA", configOverride)
	// caResult, detailedResponse, err := service.CreateCa(caOpts)
	// fmt.Println("result:", caResult)
	// fmt.Println("response:", detailedResponse)

	// // Import a CA
	// // Create an authenticator
	// authenticator = &core.IamAuthenticator{
	// 	ApiKey: ApiKey,
	// 	URL:    IdentityUrl,
	// }

	// // Create an instance of the "BlockchainV2Options" struct
	// options = &blockchainv2.BlockchainV2Options{
	// 	Authenticator: authenticator,
	// 	URL:           myserviceURL,
	// }

	// // Create an instance of the "BlockchainV2" service client.
	// service, err = blockchainv2.NewBlockchainV2(options)
	// if err != nil {
	// 	return
	// }

	// // Import CA
	// importCaOpts := service.NewImportCaOptions("My Imported CA", "http://localhost:3000", "myca", "tlsca", "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2ekNDQVpHZ0F3SUJBZ0lVV1lVRVNMV1FXMGZTRWhtVkxWLzNuWTcvYmFZd0NnWUlLb1pJemowRUF3SXcKYURFTE1Ba0dBMVVFQmhNQ1ZWTXhGekFWQmdOVkJBZ1REazV2Y25Sb0lFTmhjbTlzYVc1aE1SUXdFZ1lEVlFRSwpFd3RJZVhCbGNteGxaR2RsY2pFUE1BMEdBMVVFQ3hNR1JtRmljbWxqTVJrd0Z3WURWUVFERXhCbVlXSnlhV010ClkyRXRjMlZ5ZG1WeU1CNFhEVEl3TURFd09URTBORFV3TUZvWERUSXhNREV3T0RFME5UQXdNRm93SVRFUE1BMEcKQTFVRUN4TUdZMnhwWlc1ME1RNHdEQVlEVlFRREV3VmhaRzFwYmpCWk1CTUdCeXFHU000OUFnRUdDQ3FHU000OQpBd0VIQTBJQUJMc29Md1VDMGRCSkJlZEcwOXRyN2xuNm84T2JMWVgyZVJEZVByWlRzWm8yVjhPZjFBSkl1SEk0CmhEZHFSV0tITXRuamowUUMwK09WNEpYay9LSExtbytqWURCZU1BNEdBMVVkRHdFQi93UUVBd0lIZ0RBTUJnTlYKSFJNQkFmOEVBakFBTUIwR0ExVWREZ1FXQkJUYjB0RGs2WU5LRG9EMWxCWExocWtjTzc0Z1pUQWZCZ05WSFNNRQpHREFXZ0JUOENHNFdINCtJQXdYVWNFMFNVaDJhakowRVV6QUtCZ2dxaGtqT1BRUURBZ05JQURCRkFpRUE5cWl4CktmdC9hL1FqRXZMTXJXOVpyaDRnakVML01tVnpmSjlHZ0pwNUN5a0NJRVFUdmZBbkNHR1BHcVpJa0FNSUFpc1gKbVVkVTFreVZiai96bGhJa0JDSW0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo")
	// importCaResult, detailedResponse, err := service.ImportCa(importCaOpts)
	// fmt.Println("result:", importCaResult)
	// fmt.Println("response:", detailedResponse)

	// // Update a CA
	// // Create an authenticator
	// authenticator = &core.IamAuthenticator{
	// 	ApiKey: ApiKey,
	// 	URL:    IdentityUrl,
	// }

	// // Create an instance of the "BlockchainV2Options" struct
	// options = &blockchainv2.BlockchainV2Options{
	// 	Authenticator: authenticator,
	// 	URL:           myserviceURL,
	// }

	// // Create an instance of the "BlockchainV2" service client.
	// service, err = blockchainv2.NewBlockchainV2(options)
	// if err != nil {
	// 	return
	// }

	// // Update CA
	// cpu := "200m"
	// memory := "256Mi"
	// resourceRequests := blockchainv2.ResourceRequests{Cpu: &cpu, Memory: &memory}
	// resourceObject, err := service.NewResourceObject(&resourceRequests)
	// if err != nil {
	// 	return
	// }

	// caBodyResources, err := service.NewUpdateCaBodyResources(resourceObject)
	// if err != nil {
	// 	return
	// }

	// updateCaOpts := service.NewUpdateCaOptions("myca")
	// updateCaOpts.SetResources(caBodyResources)
	// updateCaResult, detailedResponse, err := service.UpdateCa(updateCaOpts)
	// fmt.Println("result:", updateCaResult)
	// fmt.Println("response:", detailedResponse)

	// // Edit Data About a CA
	// // Create an authenticator
	// authenticator = &core.IamAuthenticator{
	// 	ApiKey: ApiKey,
	// 	URL:    IdentityUrl,
	// }

	// // Create an instance of the "BlockchainV2Options" struct
	// options = &blockchainv2.BlockchainV2Options{
	// 	Authenticator: authenticator,
	// 	URL:           myserviceURL,
	// }

	// // Create an instance of the "BlockchainV2" service client.
	// service, err = blockchainv2.NewBlockchainV2(options)
	// if err != nil {
	// 	return
	// }

	// // Edit CA Data
	// tags := [4]string{"fabric-ca", "ibm_sass", "blue_team", "dev"}
	// editCaDataOpts := service.NewEditCaOptions("myca")
	// editCaDataOpts.SetCaName("My Ca Edited")
	// editCaDataOpts.SetTags(tags[:])
	// editCaDataResult, detailedResponse, err := service.EditCa(editCaDataOpts)
	// fmt.Println("result:", editCaDataResult)
	// fmt.Println("response:", detailedResponse)

	// // Create a Peer
	// // Create an authenticator
	// authenticator = &core.IamAuthenticator{
	// 	ApiKey: ApiKey,
	// 	URL:    IdentityUrl,
	// }

	// // Create an instance of the "BlockchainV2Options" struct
	// options = &blockchainv2.BlockchainV2Options{
	// 	Authenticator: authenticator,
	// 	URL:           myserviceURL,
	// }

	// // Create an instance of the "BlockchainV2" service client.
	// service, err = blockchainv2.NewBlockchainV2(options)
	// if err != nil {
	// 	return
	// }

	// // Create Peer
	// cacerts := [1]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2ekNDQVpHZ0F3SUJBZ0lVV1lVRVNMV1FXMGZTRWhtVkxWLzNuWTcvYmFZd0NnWUlLb1pJemowRUF3SXcKYURFTE1Ba0dBMVVFQmhNQ1ZWTXhGekFWQmdOVkJBZ1REazV2Y25Sb0lFTmhjbTlzYVc1aE1SUXdFZ1lEVlFRSwpFd3RJZVhCbGNteGxaR2RsY2pFUE1BMEdBMVVFQ3hNR1JtRmljbWxqTVJrd0Z3WURWUVFERXhCbVlXSnlhV010ClkyRXRjMlZ5ZG1WeU1CNFhEVEl3TURFd09URTBORFV3TUZvWERUSXhNREV3T0RFME5UQXdNRm93SVRFUE1BMEcKQTFVRUN4TUdZMnhwWlc1ME1RNHdEQVlEVlFRREV3VmhaRzFwYmpCWk1CTUdCeXFHU000OUFnRUdDQ3FHU000OQpBd0VIQTBJQUJMc29Md1VDMGRCSkJlZEcwOXRyN2xuNm84T2JMWVgyZVJEZVByWlRzWm8yVjhPZjFBSkl1SEk0CmhEZHFSV0tITXRuamowUUMwK09WNEpYay9LSExtbytqWURCZU1BNEdBMVVkRHdFQi93UUVBd0lIZ0RBTUJnTlYKSFJNQkFmOEVBakFBTUIwR0ExVWREZ1FXQkJUYjB0RGs2WU5LRG9EMWxCWExocWtjTzc0Z1pUQWZCZ05WSFNNRQpHREFXZ0JUOENHNFdINCtJQXdYVWNFMFNVaDJhakowRVV6QUtCZ2dxaGtqT1BRUURBZ05JQURCRkFpRUE5cWl4CktmdC9hL1FqRXZMTXJXOVpyaDRnakVML01tVnpmSjlHZ0pwNUN5a0NJRVFUdmZBbkNHR1BHcVpJa0FNSUFpc1gKbVVkVTFreVZiai96bGhJa0JDSW0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo"}
	// component, err := service.NewMspConfigData(
	// 	"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2ekNDQVpHZ0F3SUJBZ0lVV1lVRVNMV1FXMGZTRWhtVkxWLzNuWTcvYmFZd0NnWUlLb1pJemowRUF3SXcKYURFTE1Ba0dBMVVFQmhNQ1ZWTXhGekFWQmdOVkJBZ1REazV2Y25Sb0lFTmhjbTlzYVc1aE1SUXdFZ1lEVlFRSwpFd3RJZVhCbGNteGxaR2RsY2pFUE1BMEdBMVVFQ3hNR1JtRmljbWxqTVJrd0Z3WURWUVFERXhCbVlXSnlhV010ClkyRXRjMlZ5ZG1WeU1CNFhEVEl3TURFd09URTBORFV3TUZvWERUSXhNREV3T0RFME5UQXdNRm93SVRFUE1BMEcKQTFVRUN4TUdZMnhwWlc1ME1RNHdEQVlEVlFRREV3VmhaRzFwYmpCWk1CTUdCeXFHU000OUFnRUdDQ3FHU000OQpBd0VIQTBJQUJMc29Md1VDMGRCSkJlZEcwOXRyN2xuNm84T2JMWVgyZVJEZVByWlRzWm8yVjhPZjFBSkl1SEk0CmhEZHFSV0tITXRuamowUUMwK09WNEpYay9LSExtbytqWURCZU1BNEdBMVVkRHdFQi93UUVBd0lIZ0RBTUJnTlYKSFJNQkFmOEVBakFBTUIwR0ExVWREZ1FXQkJUYjB0RGs2WU5LRG9EMWxCWExocWtjTzc0Z1pUQWZCZ05WSFNNRQpHREFXZ0JUOENHNFdINCtJQXdYVWNFMFNVaDJhakowRVV6QUtCZ2dxaGtqT1BRUURBZ05JQURCRkFpRUE5cWl4CktmdC9hL1FqRXZMTXJXOVpyaDRnakVML01tVnpmSjlHZ0pwNUN5a0NJRVFUdmZBbkNHR1BHcVpJa0FNSUFpc1gKbVVkVTFreVZiai96bGhJa0JDSW0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo",
	// 	"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2ekNDQVpHZ0F3SUJBZ0lVV1lVRVNMV1FXMGZTRWhtVkxWLzNuWTcvYmFZd0NnWUlLb1pJemowRUF3SXcKYURFTE1Ba0dBMVVFQmhNQ1ZWTXhGekFWQmdOVkJBZ1REazV2Y25Sb0lFTmhjbTlzYVc1aE1SUXdFZ1lEVlFRSwpFd3RJZVhCbGNteGxaR2RsY2pFUE1BMEdBMVVFQ3hNR1JtRmljbWxqTVJrd0Z3WURWUVFERXhCbVlXSnlhV010ClkyRXRjMlZ5ZG1WeU1CNFhEVEl3TURFd09URTBORFV3TUZvWERUSXhNREV3T0RFME5UQXdNRm93SVRFUE1BMEcKQTFVRUN4TUdZMnhwWlc1ME1RNHdEQVlEVlFRREV3VmhaRzFwYmpCWk1CTUdCeXFHU000OUFnRUdDQ3FHU000OQpBd0VIQTBJQUJMc29Md1VDMGRCSkJlZEcwOXRyN2xuNm84T2JMWVgyZVJEZVByWlRzWm8yVjhPZjFBSkl1SEk0CmhEZHFSV0tITXRuamowUUMwK09WNEpYay9LSExtbytqWURCZU1BNEdBMVVkRHdFQi93UUVBd0lIZ0RBTUJnTlYKSFJNQkFmOEVBakFBTUIwR0ExVWREZ1FXQkJUYjB0RGs2WU5LRG9EMWxCWExocWtjTzc0Z1pUQWZCZ05WSFNNRQpHREFXZ0JUOENHNFdINCtJQXdYVWNFMFNVaDJhakowRVV6QUtCZ2dxaGtqT1BRUURBZ05JQURCRkFpRUE5cWl4CktmdC9hL1FqRXZMTXJXOVpyaDRnakVML01tVnpmSjlHZ0pwNUN5a0NJRVFUdmZBbkNHR1BHcVpJa0FNSUFpc1gKbVVkVTFreVZiai96bGhJa0JDSW0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo",
	// 	cacerts[:])
	// if err != nil {
	// 	return
	// }

	// tls, err := service.NewMspConfigData(
	// 	"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2ekNDQVpHZ0F3SUJBZ0lVV1lVRVNMV1FXMGZTRWhtVkxWLzNuWTcvYmFZd0NnWUlLb1pJemowRUF3SXcKYURFTE1Ba0dBMVVFQmhNQ1ZWTXhGekFWQmdOVkJBZ1REazV2Y25Sb0lFTmhjbTlzYVc1aE1SUXdFZ1lEVlFRSwpFd3RJZVhCbGNteGxaR2RsY2pFUE1BMEdBMVVFQ3hNR1JtRmljbWxqTVJrd0Z3WURWUVFERXhCbVlXSnlhV010ClkyRXRjMlZ5ZG1WeU1CNFhEVEl3TURFd09URTBORFV3TUZvWERUSXhNREV3T0RFME5UQXdNRm93SVRFUE1BMEcKQTFVRUN4TUdZMnhwWlc1ME1RNHdEQVlEVlFRREV3VmhaRzFwYmpCWk1CTUdCeXFHU000OUFnRUdDQ3FHU000OQpBd0VIQTBJQUJMc29Md1VDMGRCSkJlZEcwOXRyN2xuNm84T2JMWVgyZVJEZVByWlRzWm8yVjhPZjFBSkl1SEk0CmhEZHFSV0tITXRuamowUUMwK09WNEpYay9LSExtbytqWURCZU1BNEdBMVVkRHdFQi93UUVBd0lIZ0RBTUJnTlYKSFJNQkFmOEVBakFBTUIwR0ExVWREZ1FXQkJUYjB0RGs2WU5LRG9EMWxCWExocWtjTzc0Z1pUQWZCZ05WSFNNRQpHREFXZ0JUOENHNFdINCtJQXdYVWNFMFNVaDJhakowRVV6QUtCZ2dxaGtqT1BRUURBZ05JQURCRkFpRUE5cWl4CktmdC9hL1FqRXZMTXJXOVpyaDRnakVML01tVnpmSjlHZ0pwNUN5a0NJRVFUdmZBbkNHR1BHcVpJa0FNSUFpc1gKbVVkVTFreVZiai96bGhJa0JDSW0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo",
	// 	"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2ekNDQVpHZ0F3SUJBZ0lVV1lVRVNMV1FXMGZTRWhtVkxWLzNuWTcvYmFZd0NnWUlLb1pJemowRUF3SXcKYURFTE1Ba0dBMVVFQmhNQ1ZWTXhGekFWQmdOVkJBZ1REazV2Y25Sb0lFTmhjbTlzYVc1aE1SUXdFZ1lEVlFRSwpFd3RJZVhCbGNteGxaR2RsY2pFUE1BMEdBMVVFQ3hNR1JtRmljbWxqTVJrd0Z3WURWUVFERXhCbVlXSnlhV010ClkyRXRjMlZ5ZG1WeU1CNFhEVEl3TURFd09URTBORFV3TUZvWERUSXhNREV3T0RFME5UQXdNRm93SVRFUE1BMEcKQTFVRUN4TUdZMnhwWlc1ME1RNHdEQVlEVlFRREV3VmhaRzFwYmpCWk1CTUdCeXFHU000OUFnRUdDQ3FHU000OQpBd0VIQTBJQUJMc29Md1VDMGRCSkJlZEcwOXRyN2xuNm84T2JMWVgyZVJEZVByWlRzWm8yVjhPZjFBSkl1SEk0CmhEZHFSV0tITXRuamowUUMwK09WNEpYay9LSExtbytqWURCZU1BNEdBMVVkRHdFQi93UUVBd0lIZ0RBTUJnTlYKSFJNQkFmOEVBakFBTUIwR0ExVWREZ1FXQkJUYjB0RGs2WU5LRG9EMWxCWExocWtjTzc0Z1pUQWZCZ05WSFNNRQpHREFXZ0JUOENHNFdINCtJQXdYVWNFMFNVaDJhakowRVV6QUtCZ2dxaGtqT1BRUURBZ05JQURCRkFpRUE5cWl4CktmdC9hL1FqRXZMTXJXOVpyaDRnakVML01tVnpmSjlHZ0pwNUN5a0NJRVFUdmZBbkNHR1BHcVpJa0FNSUFpc1gKbVVkVTFreVZiai96bGhJa0JDSW0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo",
	// 	cacerts[:])
	// if err != nil {
	// 	return
	// }

	// configObjectMsp, err := service.NewConfigObjectMsp(component, tls)
	// if err != nil {
	// 	return
	// }
	// configObject := blockchainv2.ConfigObject{Msp: configObjectMsp}
	// peerOpts := service.NewCreatePeerOptions("org1msp", "Peer1", &configObject)
	// peerResult, detailedResponse, err := service.CreatePeer(peerOpts)
	// fmt.Println("result:", peerResult)
	// fmt.Println("response:", detailedResponse)

	// // Import a Peer
	// // Create an authenticator
	// authenticator = &core.IamAuthenticator{
	// 	ApiKey: ApiKey,
	// 	URL:    IdentityUrl,
	// }

	// // Create an instance of the "BlockchainV2Options" struct
	// options = &blockchainv2.BlockchainV2Options{
	// 	Authenticator: authenticator,
	// 	URL:           myserviceURL,
	// }

	// // Create an instance of the "BlockchainV2" service client.
	// service, err = blockchainv2.NewBlockchainV2(options)
	// if err != nil {
	// 	return
	// }

	// // Import Peer
	// importPeerOpts := service.NewImportPeerOptions("peer1", "ImportedMSP", "http://localhost:3000", "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2ekNDQVpHZ0F3SUJBZ0lVV1lVRVNMV1FXMGZTRWhtVkxWLzNuWTcvYmFZd0NnWUlLb1pJemowRUF3SXcKYURFTE1Ba0dBMVVFQmhNQ1ZWTXhGekFWQmdOVkJBZ1REazV2Y25Sb0lFTmhjbTlzYVc1aE1SUXdFZ1lEVlFRSwpFd3RJZVhCbGNteGxaR2RsY2pFUE1BMEdBMVVFQ3hNR1JtRmljbWxqTVJrd0Z3WURWUVFERXhCbVlXSnlhV010ClkyRXRjMlZ5ZG1WeU1CNFhEVEl3TURFd09URTBORFV3TUZvWERUSXhNREV3T0RFME5UQXdNRm93SVRFUE1BMEcKQTFVRUN4TUdZMnhwWlc1ME1RNHdEQVlEVlFRREV3VmhaRzFwYmpCWk1CTUdCeXFHU000OUFnRUdDQ3FHU000OQpBd0VIQTBJQUJMc29Md1VDMGRCSkJlZEcwOXRyN2xuNm84T2JMWVgyZVJEZVByWlRzWm8yVjhPZjFBSkl1SEk0CmhEZHFSV0tITXRuamowUUMwK09WNEpYay9LSExtbytqWURCZU1BNEdBMVVkRHdFQi93UUVBd0lIZ0RBTUJnTlYKSFJNQkFmOEVBakFBTUIwR0ExVWREZ1FXQkJUYjB0RGs2WU5LRG9EMWxCWExocWtjTzc0Z1pUQWZCZ05WSFNNRQpHREFXZ0JUOENHNFdINCtJQXdYVWNFMFNVaDJhakowRVV6QUtCZ2dxaGtqT1BRUURBZ05JQURCRkFpRUE5cWl4CktmdC9hL1FqRXZMTXJXOVpyaDRnakVML01tVnpmSjlHZ0pwNUN5a0NJRVFUdmZBbkNHR1BHcVpJa0FNSUFpc1gKbVVkVTFreVZiai96bGhJa0JDSW0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo")
	// importPeerResult, detailedResponse, err := service.ImportPeer(importPeerOpts)
	// fmt.Println("result:", importPeerResult)
	// fmt.Println("response:", detailedResponse)

	// // Import Data About a Peer
	// // Create an authenticator
	// authenticator = &core.IamAuthenticator{
	// 	ApiKey: ApiKey,
	// 	URL:    IdentityUrl,
	// }

	// // Create an instance of the "BlockchainV2Options" struct
	// options = &blockchainv2.BlockchainV2Options{
	// 	Authenticator: authenticator,
	// 	URL:           myserviceURL,
	// }

	// // Create an instance of the "BlockchainV2" service client.
	// service, err = blockchainv2.NewBlockchainV2(options)
	// if err != nil {
	// 	return
	// }

	// // Edit Peer Data
	// tags := [4]string{"fabric-peer", "ibm_sass", "red_team", "dev"}
	// editPeerDataOpts := service.NewEditPeerOptions("peer1")
	// editPeerDataOpts.SetDisplayName("My Other Peer")
	// editPeerDataOpts.SetMspID("peermsp")
	// editPeerDataOpts.SetTags(tags[:])
	// editCaDataResult, detailedResponse, err := service.EditPeer(editPeerDataOpts)
	// fmt.Println("result:", editCaDataResult)
	// fmt.Println("response:", detailedResponse)

	// // Update a Peer
	// // Create an authenticator
	// authenticator = &core.IamAuthenticator{
	// 	ApiKey: ApiKey,
	// 	URL:    IdentityUrl,
	// }

	// // Create an instance of the "BlockchainV2Options" struct
	// options = &blockchainv2.BlockchainV2Options{
	// 	Authenticator: authenticator,
	// 	URL:           myserviceURL,
	// }

	// // Create an instance of the "BlockchainV2" service client.
	// service, err = blockchainv2.NewBlockchainV2(options)
	// if err != nil {
	// 	return
	// }

	// // Update Peer
	// cpu := "400m"
	// memory := "800Mi"
	// resourceRequests := blockchainv2.ResourceRequests{Cpu: &cpu, Memory: &memory}
	// resourceObject, err := service.NewResourceObject(&resourceRequests)
	// if err != nil {
	// 	return
	// }

	// peerBodyResources := &blockchainv2.PeerResources{Peer: resourceObject}

	// updatePeerOpts := service.NewUpdatePeerOptions("peer1")
	// updatePeerOpts.SetResources(peerBodyResources)
	// updateCaResult, detailedResponse, err := service.UpdatePeer(updatePeerOpts)
	// fmt.Println("result:", updateCaResult)
	// fmt.Println("response:", detailedResponse)

	// // Create an Ordering Service
	// authenticator = &core.IamAuthenticator{
	// 	ApiKey: ApiKey,
	// 	URL:    IdentityUrl,
	// }

	// // Create an instance of the "BlockchainV2Options" struct
	// options = &blockchainv2.BlockchainV2Options{
	// 	Authenticator: authenticator,
	// 	URL:           myserviceURL,
	// }

	// // Create an instance of the "BlockchainV2" service client.
	// service, err = blockchainv2.NewBlockchainV2(options)
	// if err != nil {
	// 	return
	// }

	// // Create Ordering Service
	// cacert, err := service.NewConfigObjectEnrollmentComponentCatls("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2ekNDQVpHZ0F3SUJBZ0lVV1lVRVNMV1FXMGZTRWhtVkxWLzNuWTcvYmFZd0NnWUlLb1pJemowRUF3SXcKYURFTE1Ba0dBMVVFQmhNQ1ZWTXhGekFWQmdOVkJBZ1REazV2Y25Sb0lFTmhjbTlzYVc1aE1SUXdFZ1lEVlFRSwpFd3RJZVhCbGNteGxaR2RsY2pFUE1BMEdBMVVFQ3hNR1JtRmljbWxqTVJrd0Z3WURWUVFERXhCbVlXSnlhV010ClkyRXRjMlZ5ZG1WeU1CNFhEVEl3TURFd09URTBORFV3TUZvWERUSXhNREV3T0RFME5UQXdNRm93SVRFUE1BMEcKQTFVRUN4TUdZMnhwWlc1ME1RNHdEQVlEVlFRREV3VmhaRzFwYmpCWk1CTUdCeXFHU000OUFnRUdDQ3FHU000OQpBd0VIQTBJQUJMc29Md1VDMGRCSkJlZEcwOXRyN2xuNm84T2JMWVgyZVJEZVByWlRzWm8yVjhPZjFBSkl1SEk0CmhEZHFSV0tITXRuamowUUMwK09WNEpYay9LSExtbytqWURCZU1BNEdBMVVkRHdFQi93UUVBd0lIZ0RBTUJnTlYKSFJNQkFmOEVBakFBTUIwR0ExVWREZ1FXQkJUYjB0RGs2WU5LRG9EMWxCWExocWtjTzc0Z1pUQWZCZ05WSFNNRQpHREFXZ0JUOENHNFdINCtJQXdYVWNFMFNVaDJhakowRVV6QUtCZ2dxaGtqT1BRUURBZ05JQURCRkFpRUE5cWl4CktmdC9hL1FqRXZMTXJXOVpyaDRnakVML01tVnpmSjlHZ0pwNUN5a0NJRVFUdmZBbkNHR1BHcVpJa0FNSUFpc1gKbVVkVTFreVZiai96bGhJa0JDSW0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo")
	// if err != nil {
	// 	return
	// }

	// caTlsCert, err := service.NewConfigObjectEnrollmentTlsCatls("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2ekNDQVpHZ0F3SUJBZ0lVV1lVRVNMV1FXMGZTRWhtVkxWLzNuWTcvYmFZd0NnWUlLb1pJemowRUF3SXcKYURFTE1Ba0dBMVVFQmhNQ1ZWTXhGekFWQmdOVkJBZ1REazV2Y25Sb0lFTmhjbTlzYVc1aE1SUXdFZ1lEVlFRSwpFd3RJZVhCbGNteGxaR2RsY2pFUE1BMEdBMVVFQ3hNR1JtRmljbWxqTVJrd0Z3WURWUVFERXhCbVlXSnlhV010ClkyRXRjMlZ5ZG1WeU1CNFhEVEl3TURFd09URTBORFV3TUZvWERUSXhNREV3T0RFME5UQXdNRm93SVRFUE1BMEcKQTFVRUN4TUdZMnhwWlc1ME1RNHdEQVlEVlFRREV3VmhaRzFwYmpCWk1CTUdCeXFHU000OUFnRUdDQ3FHU000OQpBd0VIQTBJQUJMc29Md1VDMGRCSkJlZEcwOXRyN2xuNm84T2JMWVgyZVJEZVByWlRzWm8yVjhPZjFBSkl1SEk0CmhEZHFSV0tITXRuamowUUMwK09WNEpYay9LSExtbytqWURCZU1BNEdBMVVkRHdFQi93UUVBd0lIZ0RBTUJnTlYKSFJNQkFmOEVBakFBTUIwR0ExVWREZ1FXQkJUYjB0RGs2WU5LRG9EMWxCWExocWtjTzc0Z1pUQWZCZ05WSFNNRQpHREFXZ0JUOENHNFdINCtJQXdYVWNFMFNVaDJhakowRVV6QUtCZ2dxaGtqT1BRUURBZ05JQURCRkFpRUE5cWl4CktmdC9hL1FqRXZMTXJXOVpyaDRnakVML01tVnpmSjlHZ0pwNUN5a0NJRVFUdmZBbkNHR1BHcVpJa0FNSUFpc1gKbVVkVTFreVZiai96bGhJa0JDSW0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo")
	// if err != nil {
	// 	return
	// }

	// enrollmentComponent, err := service.NewConfigObjectEnrollmentComponent("localhost", float64(3000), "myca", cacert, "admin", "password")
	// if err != nil {
	// 	return
	// }

	// tlsEnrollment, err := service.NewConfigObjectEnrollmentTls("localhost", float64(3000), "myca", caTlsCert)
	// if err != nil {
	// 	return
	// }

	// objectEnrollment, err := service.NewConfigObjectEnrollment(enrollmentComponent, tlsEnrollment)
	// if err != nil {
	// 	return
	// }

	// configObject := blockchainv2.ConfigObject{Enrollment: objectEnrollment}
	// configObjectArray := [1]blockchainv2.ConfigObject{configObject}
	// ordererOpts := service.NewCreateOrdererOptions("raft", "orderermsp", "ordering service node", configObjectArray[:])
	// ordererOpts.SetOrdererType("raft")
	// ordererOpts.SetMspID("orderermsp")
	// ordererOpts.SetConfig(configObjectArray[:])
	// ordererOpts.SetClusterName("My three Node Raft")
	// ordererOpts.SetDisplayName("ordering service node")
	// ordererResult, detailedResponse, err := service.CreateOrderer(ordererOpts)
	// fmt.Println("result:", ordererResult)
	// fmt.Println("response:", detailedResponse)

	// // Import an Ordering Service
	// // Create an authenticator
	// authenticator = &core.IamAuthenticator{
	// 	ApiKey: ApiKey,
	// 	URL:    IdentityUrl,
	// }

	// // Create an instance of the "BlockchainV2Options" struct
	// options = &blockchainv2.BlockchainV2Options{
	// 	Authenticator: authenticator,
	// 	URL:           myserviceURL,
	// }

	// // Create an instance of the "BlockchainV2" service client.
	// service, err = blockchainv2.NewBlockchainV2(options)
	// if err != nil {
	// 	return
	// }

	// // Import Ordering Service
	// importOrdererOpts := service.NewImportOrdererOptions("My Raft OS", "orderer node", "org1", "https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443", "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2ekNDQVpHZ0F3SUJBZ0lVV1lVRVNMV1FXMGZTRWhtVkxWLzNuWTcvYmFZd0NnWUlLb1pJemowRUF3SXcKYURFTE1Ba0dBMVVFQmhNQ1ZWTXhGekFWQmdOVkJBZ1REazV2Y25Sb0lFTmhjbTlzYVc1aE1SUXdFZ1lEVlFRSwpFd3RJZVhCbGNteGxaR2RsY2pFUE1BMEdBMVVFQ3hNR1JtRmljbWxqTVJrd0Z3WURWUVFERXhCbVlXSnlhV010ClkyRXRjMlZ5ZG1WeU1CNFhEVEl3TURFd09URTBORFV3TUZvWERUSXhNREV3T0RFME5UQXdNRm93SVRFUE1BMEcKQTFVRUN4TUdZMnhwWlc1ME1RNHdEQVlEVlFRREV3VmhaRzFwYmpCWk1CTUdCeXFHU000OUFnRUdDQ3FHU000OQpBd0VIQTBJQUJMc29Md1VDMGRCSkJlZEcwOXRyN2xuNm84T2JMWVgyZVJEZVByWlRzWm8yVjhPZjFBSkl1SEk0CmhEZHFSV0tITXRuamowUUMwK09WNEpYay9LSExtbytqWURCZU1BNEdBMVVkRHdFQi93UUVBd0lIZ0RBTUJnTlYKSFJNQkFmOEVBakFBTUIwR0ExVWREZ1FXQkJUYjB0RGs2WU5LRG9EMWxCWExocWtjTzc0Z1pUQWZCZ05WSFNNRQpHREFXZ0JUOENHNFdINCtJQXdYVWNFMFNVaDJhakowRVV6QUtCZ2dxaGtqT1BRUURBZ05JQURCRkFpRUE5cWl4CktmdC9hL1FqRXZMTXJXOVpyaDRnakVML01tVnpmSjlHZ0pwNUN5a0NJRVFUdmZBbkNHR1BHcVpJa0FNSUFpc1gKbVVkVTFreVZiai96bGhJa0JDSW0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo")
	// importOrdererOpts.SetClusterID("abcde")
	// importOrdererOpts.SetApiURL("grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
	// importOrdererOpts.SetOperationsURL("https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443")
	// importOrdererOpts.SetSystemChannelID("testchainid")
	// importOrdererOpts.SetTlsCert("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2ekNDQVpHZ0F3SUJBZ0lVV1lVRVNMV1FXMGZTRWhtVkxWLzNuWTcvYmFZd0NnWUlLb1pJemowRUF3SXcKYURFTE1Ba0dBMVVFQmhNQ1ZWTXhGekFWQmdOVkJBZ1REazV2Y25Sb0lFTmhjbTlzYVc1aE1SUXdFZ1lEVlFRSwpFd3RJZVhCbGNteGxaR2RsY2pFUE1BMEdBMVVFQ3hNR1JtRmljbWxqTVJrd0Z3WURWUVFERXhCbVlXSnlhV010ClkyRXRjMlZ5ZG1WeU1CNFhEVEl3TURFd09URTBORFV3TUZvWERUSXhNREV3T0RFME5UQXdNRm93SVRFUE1BMEcKQTFVRUN4TUdZMnhwWlc1ME1RNHdEQVlEVlFRREV3VmhaRzFwYmpCWk1CTUdCeXFHU000OUFnRUdDQ3FHU000OQpBd0VIQTBJQUJMc29Md1VDMGRCSkJlZEcwOXRyN2xuNm84T2JMWVgyZVJEZVByWlRzWm8yVjhPZjFBSkl1SEk0CmhEZHFSV0tITXRuamowUUMwK09WNEpYay9LSExtbytqWURCZU1BNEdBMVVkRHdFQi93UUVBd0lIZ0RBTUJnTlYKSFJNQkFmOEVBakFBTUIwR0ExVWREZ1FXQkJUYjB0RGs2WU5LRG9EMWxCWExocWtjTzc0Z1pUQWZCZ05WSFNNRQpHREFXZ0JUOENHNFdINCtJQXdYVWNFMFNVaDJhakowRVV6QUtCZ2dxaGtqT1BRUURBZ05JQURCRkFpRUE5cWl4CktmdC9hL1FqRXZMTXJXOVpyaDRnakVML01tVnpmSjlHZ0pwNUN5a0NJRVFUdmZBbkNHR1BHcVpJa0FNSUFpc1gKbVVkVTFreVZiai96bGhJa0JDSW0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo")
	// importPeerResult, detailedResponse, err := service.ImportOrderer(importOrdererOpts)
	// fmt.Println("result:", importPeerResult)
	// fmt.Println("response:", detailedResponse)

	// // Edit Data About an Orderer
	// // Create an authenticator
	// authenticator = &core.IamAuthenticator{
	// 	ApiKey: ApiKey,
	// 	URL:    IdentityUrl,
	// }

	// // Create an instance of the "BlockchainV2Options" struct
	// options = &blockchainv2.BlockchainV2Options{
	// 	Authenticator: authenticator,
	// 	URL:           myserviceURL,
	// }

	// // Create an instance of the "BlockchainV2" service client.
	// service, err = blockchainv2.NewBlockchainV2(options)
	// if err != nil {
	// 	return
	// }

	// // Edit Orderer Data
	// editOrdererDataOpts := service.NewEditOrdererOptions("orderernode")
	// editOrdererDataOpts.SetClusterName("My Other OS")
	// editOrdererDataOpts.SetDisplayName("ordering node")
	// editOrdererDataOpts.SetMspID("orderermsp")
	// editCaDataResult, detailedResponse, err := service.EditOrderer(editOrdererDataOpts)
	// fmt.Println("result:", editCaDataResult)
	// fmt.Println("response:", detailedResponse)

	// // Update an Orderer Node
	// // Create an authenticator
	// authenticator = &core.IamAuthenticator{
	// 	ApiKey: ApiKey,
	// 	URL:    IdentityUrl,
	// }

	// // Create an instance of the "BlockchainV2Options" struct
	// options = &blockchainv2.BlockchainV2Options{
	// 	Authenticator: authenticator,
	// 	URL:           myserviceURL,
	// }

	// // Create an instance of the "BlockchainV2" service client.
	// service, err = blockchainv2.NewBlockchainV2(options)
	// if err != nil {
	// 	return
	// }

	// // Update Orderer
	// cpu := "500m"
	// memory := "1024Mi"
	// requests := &blockchainv2.ResourceRequests{Cpu: &cpu, Memory: &memory}

	// resourceObject, err := service.NewResourceObject(requests)
	// if err != nil {
	// 	return
	// }

	// ordererResources := &blockchainv2.UpdateOrdererBodyResources{Orderer: resourceObject}
	// updateOrdererOpts := service.NewUpdateOrdererOptions("orderernode")
	// updateOrdererOpts.SetResources(ordererResources)
	// updateCaResult, detailedResponse, err := service.UpdateOrderer(updateOrdererOpts)
	// fmt.Println("result:", updateCaResult)
	// fmt.Println("response:", detailedResponse)

	// // Submit Config Block to Orderer
	// // Create an authenticator
	// authenticator = &core.IamAuthenticator{
	// 	ApiKey: ApiKey,
	// 	URL:    IdentityUrl,
	// }

	// // Create an instance of the "BlockchainV2Options" struct
	// options = &blockchainv2.BlockchainV2Options{
	// 	Authenticator: authenticator,
	// 	URL:           myserviceURL,
	// }

	// // Create an instance of the "BlockchainV2" service client.
	// service, err = blockchainv2.NewBlockchainV2(options)
	// if err != nil {
	// 	return
	// }

	// // Submit Config Block
	// submitBlockOpts := service.NewSubmitBlockOptions("orderernode")
	// submitBlockOpts.SetB64Block("bWFzc2l2ZSBiaW5hcnkgb2YgYSBjb25maWcgYmxvY2sgd291bGQgYmUgaGVyZSBpZiB0aGlzIHdhcyByZWFs")
	// updateCaResult, detailedResponse, err := service.SubmitBlock(submitBlockOpts)
	// fmt.Println("result:", updateCaResult)
	// fmt.Println("response:", detailedResponse)

	// // Import an MSP Definition
	// // Create an authenticator
	// authenticator = &core.IamAuthenticator{
	// 	ApiKey: ApiKey,
	// 	URL:    IdentityUrl,
	// }

	// // Create an instance of the "BlockchainV2Options" struct
	// options = &blockchainv2.BlockchainV2Options{
	// 	Authenticator: authenticator,
	// 	URL:           myserviceURL,
	// }

	// // Create an instance of the "BlockchainV2" service client.
	// service, err = blockchainv2.NewBlockchainV2(options)
	// if err != nil {
	// 	return
	// }

	// // Import MSP Definition
	// rootCerts := [1]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2ekNDQVpHZ0F3SUJBZ0lVV1lVRVNMV1FXMGZTRWhtVkxWLzNuWTcvYmFZd0NnWUlLb1pJemowRUF3SXcKYURFTE1Ba0dBMVVFQmhNQ1ZWTXhGekFWQmdOVkJBZ1REazV2Y25Sb0lFTmhjbTlzYVc1aE1SUXdFZ1lEVlFRSwpFd3RJZVhCbGNteGxaR2RsY2pFUE1BMEdBMVVFQ3hNR1JtRmljbWxqTVJrd0Z3WURWUVFERXhCbVlXSnlhV010ClkyRXRjMlZ5ZG1WeU1CNFhEVEl3TURFd09URTBORFV3TUZvWERUSXhNREV3T0RFME5UQXdNRm93SVRFUE1BMEcKQTFVRUN4TUdZMnhwWlc1ME1RNHdEQVlEVlFRREV3VmhaRzFwYmpCWk1CTUdCeXFHU000OUFnRUdDQ3FHU000OQpBd0VIQTBJQUJMc29Md1VDMGRCSkJlZEcwOXRyN2xuNm84T2JMWVgyZVJEZVByWlRzWm8yVjhPZjFBSkl1SEk0CmhEZHFSV0tITXRuamowUUMwK09WNEpYay9LSExtbytqWURCZU1BNEdBMVVkRHdFQi93UUVBd0lIZ0RBTUJnTlYKSFJNQkFmOEVBakFBTUIwR0ExVWREZ1FXQkJUYjB0RGs2WU5LRG9EMWxCWExocWtjTzc0Z1pUQWZCZ05WSFNNRQpHREFXZ0JUOENHNFdINCtJQXdYVWNFMFNVaDJhakowRVV6QUtCZ2dxaGtqT1BRUURBZ05JQURCRkFpRUE5cWl4CktmdC9hL1FqRXZMTXJXOVpyaDRnakVML01tVnpmSjlHZ0pwNUN5a0NJRVFUdmZBbkNHR1BHcVpJa0FNSUFpc1gKbVVkVTFreVZiai96bGhJa0JDSW0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo"}
	// admins := [1]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2ekNDQVpHZ0F3SUJBZ0lVV1lVRVNMV1FXMGZTRWhtVkxWLzNuWTcvYmFZd0NnWUlLb1pJemowRUF3SXcKYURFTE1Ba0dBMVVFQmhNQ1ZWTXhGekFWQmdOVkJBZ1REazV2Y25Sb0lFTmhjbTlzYVc1aE1SUXdFZ1lEVlFRSwpFd3RJZVhCbGNteGxaR2RsY2pFUE1BMEdBMVVFQ3hNR1JtRmljbWxqTVJrd0Z3WURWUVFERXhCbVlXSnlhV010ClkyRXRjMlZ5ZG1WeU1CNFhEVEl3TURFd09URTBORFV3TUZvWERUSXhNREV3T0RFME5UQXdNRm93SVRFUE1BMEcKQTFVRUN4TUdZMnhwWlc1ME1RNHdEQVlEVlFRREV3VmhaRzFwYmpCWk1CTUdCeXFHU000OUFnRUdDQ3FHU000OQpBd0VIQTBJQUJMc29Md1VDMGRCSkJlZEcwOXRyN2xuNm84T2JMWVgyZVJEZVByWlRzWm8yVjhPZjFBSkl1SEk0CmhEZHFSV0tITXRuamowUUMwK09WNEpYay9LSExtbytqWURCZU1BNEdBMVVkRHdFQi93UUVBd0lIZ0RBTUJnTlYKSFJNQkFmOEVBakFBTUIwR0ExVWREZ1FXQkJUYjB0RGs2WU5LRG9EMWxCWExocWtjTzc0Z1pUQWZCZ05WSFNNRQpHREFXZ0JUOENHNFdINCtJQXdYVWNFMFNVaDJhakowRVV6QUtCZ2dxaGtqT1BRUURBZ05JQURCRkFpRUE5cWl4CktmdC9hL1FqRXZMTXJXOVpyaDRnakVML01tVnpmSjlHZ0pwNUN5a0NJRVFUdmZBbkNHR1BHcVpJa0FNSUFpc1gKbVVkVTFreVZiai96bGhJa0JDSW0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo"}
	// tlsRootCerts := [1]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2ekNDQVpHZ0F3SUJBZ0lVV1lVRVNMV1FXMGZTRWhtVkxWLzNuWTcvYmFZd0NnWUlLb1pJemowRUF3SXcKYURFTE1Ba0dBMVVFQmhNQ1ZWTXhGekFWQmdOVkJBZ1REazV2Y25Sb0lFTmhjbTlzYVc1aE1SUXdFZ1lEVlFRSwpFd3RJZVhCbGNteGxaR2RsY2pFUE1BMEdBMVVFQ3hNR1JtRmljbWxqTVJrd0Z3WURWUVFERXhCbVlXSnlhV010ClkyRXRjMlZ5ZG1WeU1CNFhEVEl3TURFd09URTBORFV3TUZvWERUSXhNREV3T0RFME5UQXdNRm93SVRFUE1BMEcKQTFVRUN4TUdZMnhwWlc1ME1RNHdEQVlEVlFRREV3VmhaRzFwYmpCWk1CTUdCeXFHU000OUFnRUdDQ3FHU000OQpBd0VIQTBJQUJMc29Md1VDMGRCSkJlZEcwOXRyN2xuNm84T2JMWVgyZVJEZVByWlRzWm8yVjhPZjFBSkl1SEk0CmhEZHFSV0tITXRuamowUUMwK09WNEpYay9LSExtbytqWURCZU1BNEdBMVVkRHdFQi93UUVBd0lIZ0RBTUJnTlYKSFJNQkFmOEVBakFBTUIwR0ExVWREZ1FXQkJUYjB0RGs2WU5LRG9EMWxCWExocWtjTzc0Z1pUQWZCZ05WSFNNRQpHREFXZ0JUOENHNFdINCtJQXdYVWNFMFNVaDJhakowRVV6QUtCZ2dxaGtqT1BRUURBZ05JQURCRkFpRUE5cWl4CktmdC9hL1FqRXZMTXJXOVpyaDRnakVML01tVnpmSjlHZ0pwNUN5a0NJRVFUdmZBbkNHR1BHcVpJa0FNSUFpc1gKbVVkVTFreVZiai96bGhJa0JDSW0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo"}
	// importMspOpts := service.NewImportMspOptions("org1", "My First Org", rootCerts[:])
	// importMspOpts.SetAdmins(admins[:])
	// importMspOpts.SetTlsRootCerts(tlsRootCerts[:])
	// importMspResult, detailedResponse, err := service.ImportMsp(importMspOpts)
	// fmt.Println("result:", importMspResult)
	// fmt.Println("response:", detailedResponse)

	// // Edit an MSP
	// // Create an authenticator
	// authenticator = &core.IamAuthenticator{
	// 	ApiKey: ApiKey,
	// 	URL:    IdentityUrl,
	// }

	// // Create an instance of the "BlockchainV2Options" struct
	// options = &blockchainv2.BlockchainV2Options{
	// 	Authenticator: authenticator,
	// 	URL:           myserviceURL,
	// }

	// // Create an instance of the "BlockchainV2" service client.
	// service, err = blockchainv2.NewBlockchainV2(options)
	// if err != nil {
	// 	return
	// }

	// // Edit MSP
	// editMSPOpts := service.NewEditMspOptions("org1")
	// editMSPOpts.SetDisplayName("My Other MSP")
	// editMspResult, detailedResponse, err := service.EditMsp(editMSPOpts)
	// fmt.Println("result:", editMspResult)
	// fmt.Println("response:", detailedResponse)

	// Edit Admin Certs on a Component
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

	// Edit Admin Certs
	appendAdminCerts := [1]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2ekNDQVpHZ0F3SUJBZ0lVV1lVRVNMV1FXMGZTRWhtVkxWLzNuWTcvYmFZd0NnWUlLb1pJemowRUF3SXcKYURFTE1Ba0dBMVVFQmhNQ1ZWTXhGekFWQmdOVkJBZ1REazV2Y25Sb0lFTmhjbTlzYVc1aE1SUXdFZ1lEVlFRSwpFd3RJZVhCbGNteGxaR2RsY2pFUE1BMEdBMVVFQ3hNR1JtRmljbWxqTVJrd0Z3WURWUVFERXhCbVlXSnlhV010ClkyRXRjMlZ5ZG1WeU1CNFhEVEl3TURFd09URTBORFV3TUZvWERUSXhNREV3T0RFME5UQXdNRm93SVRFUE1BMEcKQTFVRUN4TUdZMnhwWlc1ME1RNHdEQVlEVlFRREV3VmhaRzFwYmpCWk1CTUdCeXFHU000OUFnRUdDQ3FHU000OQpBd0VIQTBJQUJMc29Md1VDMGRCSkJlZEcwOXRyN2xuNm84T2JMWVgyZVJEZVByWlRzWm8yVjhPZjFBSkl1SEk0CmhEZHFSV0tITXRuamowUUMwK09WNEpYay9LSExtbytqWURCZU1BNEdBMVVkRHdFQi93UUVBd0lIZ0RBTUJnTlYKSFJNQkFmOEVBakFBTUIwR0ExVWREZ1FXQkJUYjB0RGs2WU5LRG9EMWxCWExocWtjTzc0Z1pUQWZCZ05WSFNNRQpHREFXZ0JUOENHNFdINCtJQXdYVWNFMFNVaDJhakowRVV6QUtCZ2dxaGtqT1BRUURBZ05JQURCRkFpRUE5cWl4CktmdC9hL1FqRXZMTXJXOVpyaDRnakVML01tVnpmSjlHZ0pwNUN5a0NJRVFUdmZBbkNHR1BHcVpJa0FNSUFpc1gKbVVkVTFreVZiai96bGhJa0JDSW0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo"}
	removeAdminCerts := []string{}
	editAdminCertsOpts := service.NewEditAdminCertsOptions("orderernode")
	editAdminCertsOpts.SetAppendAdminCerts(appendAdminCerts[:])
	editAdminCertsOpts.SetRemoveAdminCerts(removeAdminCerts)
	editAdminCertsResult, detailedResponse, err := service.EditAdminCerts(editAdminCertsOpts)
	fmt.Println("result:", editAdminCertsResult)
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

	// boolean_false := false
	// boolean_true := true
	// type_float64 := float64(89999)
	// debug_level := "debug"
	// info_level := "info"
	// clientLoggingSettings := &blockchainv2.LoggingSettingsClient{Enabled: &boolean_true, Level: &debug_level, UniqueName: &boolean_false}
	// serverLoggingSettings := &blockchainv2.LoggingSettingsServer{Enabled: &boolean_true, Level: &info_level, UniqueName: &boolean_false}
	// opts5 := service.NewEditSettingsOptions()
	// opts5.SetMaxReqPerMin(float64(50))
	// opts5.SetInactivityTimeouts(&blockchainv2.EditSettingsBodyInactivityTimeouts{Enabled: &boolean_false, MaxIdleTime: &type_float64})
	// opts5.SetFileLogging(&blockchainv2.EditLogSettingsBody{Client: clientLoggingSettings, Server: serverLoggingSettings})
	// opts5.SetFabricLcGetCcTimeoutMs(float64(350000))
	// result5, detailedResponse, err := service.EditSettings(opts5)
	// fmt.Println("result:", result5)
	// fmt.Println("response:", detailedResponse)

	fmt.Println("err:", err)
	fmt.Println("done")
	return
}

// cacerts := [1]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2ekNDQVpHZ0F3SUJBZ0lVV1lVRVNMV1FXMGZTRWhtVkxWLzNuWTcvYmFZd0NnWUlLb1pJemowRUF3SXcKYURFTE1Ba0dBMVVFQmhNQ1ZWTXhGekFWQmdOVkJBZ1REazV2Y25Sb0lFTmhjbTlzYVc1aE1SUXdFZ1lEVlFRSwpFd3RJZVhCbGNteGxaR2RsY2pFUE1BMEdBMVVFQ3hNR1JtRmljbWxqTVJrd0Z3WURWUVFERXhCbVlXSnlhV010ClkyRXRjMlZ5ZG1WeU1CNFhEVEl3TURFd09URTBORFV3TUZvWERUSXhNREV3T0RFME5UQXdNRm93SVRFUE1BMEcKQTFVRUN4TUdZMnhwWlc1ME1RNHdEQVlEVlFRREV3VmhaRzFwYmpCWk1CTUdCeXFHU000OUFnRUdDQ3FHU000OQpBd0VIQTBJQUJMc29Md1VDMGRCSkJlZEcwOXRyN2xuNm84T2JMWVgyZVJEZVByWlRzWm8yVjhPZjFBSkl1SEk0CmhEZHFSV0tITXRuamowUUMwK09WNEpYay9LSExtbytqWURCZU1BNEdBMVVkRHdFQi93UUVBd0lIZ0RBTUJnTlYKSFJNQkFmOEVBakFBTUIwR0ExVWREZ1FXQkJUYjB0RGs2WU5LRG9EMWxCWExocWtjTzc0Z1pUQWZCZ05WSFNNRQpHREFXZ0JUOENHNFdINCtJQXdYVWNFMFNVaDJhakowRVV6QUtCZ2dxaGtqT1BRUURBZ05JQURCRkFpRUE5cWl4CktmdC9hL1FqRXZMTXJXOVpyaDRnakVML01tVnpmSjlHZ0pwNUN5a0NJRVFUdmZBbkNHR1BHcVpJa0FNSUFpc1gKbVVkVTFreVZiai96bGhJa0JDSW0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo"}
// component, err := service.NewMspConfigData(
// 	"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2ekNDQVpHZ0F3SUJBZ0lVV1lVRVNMV1FXMGZTRWhtVkxWLzNuWTcvYmFZd0NnWUlLb1pJemowRUF3SXcKYURFTE1Ba0dBMVVFQmhNQ1ZWTXhGekFWQmdOVkJBZ1REazV2Y25Sb0lFTmhjbTlzYVc1aE1SUXdFZ1lEVlFRSwpFd3RJZVhCbGNteGxaR2RsY2pFUE1BMEdBMVVFQ3hNR1JtRmljbWxqTVJrd0Z3WURWUVFERXhCbVlXSnlhV010ClkyRXRjMlZ5ZG1WeU1CNFhEVEl3TURFd09URTBORFV3TUZvWERUSXhNREV3T0RFME5UQXdNRm93SVRFUE1BMEcKQTFVRUN4TUdZMnhwWlc1ME1RNHdEQVlEVlFRREV3VmhaRzFwYmpCWk1CTUdCeXFHU000OUFnRUdDQ3FHU000OQpBd0VIQTBJQUJMc29Md1VDMGRCSkJlZEcwOXRyN2xuNm84T2JMWVgyZVJEZVByWlRzWm8yVjhPZjFBSkl1SEk0CmhEZHFSV0tITXRuamowUUMwK09WNEpYay9LSExtbytqWURCZU1BNEdBMVVkRHdFQi93UUVBd0lIZ0RBTUJnTlYKSFJNQkFmOEVBakFBTUIwR0ExVWREZ1FXQkJUYjB0RGs2WU5LRG9EMWxCWExocWtjTzc0Z1pUQWZCZ05WSFNNRQpHREFXZ0JUOENHNFdINCtJQXdYVWNFMFNVaDJhakowRVV6QUtCZ2dxaGtqT1BRUURBZ05JQURCRkFpRUE5cWl4CktmdC9hL1FqRXZMTXJXOVpyaDRnakVML01tVnpmSjlHZ0pwNUN5a0NJRVFUdmZBbkNHR1BHcVpJa0FNSUFpc1gKbVVkVTFreVZiai96bGhJa0JDSW0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo",
// 	"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2ekNDQVpHZ0F3SUJBZ0lVV1lVRVNMV1FXMGZTRWhtVkxWLzNuWTcvYmFZd0NnWUlLb1pJemowRUF3SXcKYURFTE1Ba0dBMVVFQmhNQ1ZWTXhGekFWQmdOVkJBZ1REazV2Y25Sb0lFTmhjbTlzYVc1aE1SUXdFZ1lEVlFRSwpFd3RJZVhCbGNteGxaR2RsY2pFUE1BMEdBMVVFQ3hNR1JtRmljbWxqTVJrd0Z3WURWUVFERXhCbVlXSnlhV010ClkyRXRjMlZ5ZG1WeU1CNFhEVEl3TURFd09URTBORFV3TUZvWERUSXhNREV3T0RFME5UQXdNRm93SVRFUE1BMEcKQTFVRUN4TUdZMnhwWlc1ME1RNHdEQVlEVlFRREV3VmhaRzFwYmpCWk1CTUdCeXFHU000OUFnRUdDQ3FHU000OQpBd0VIQTBJQUJMc29Md1VDMGRCSkJlZEcwOXRyN2xuNm84T2JMWVgyZVJEZVByWlRzWm8yVjhPZjFBSkl1SEk0CmhEZHFSV0tITXRuamowUUMwK09WNEpYay9LSExtbytqWURCZU1BNEdBMVVkRHdFQi93UUVBd0lIZ0RBTUJnTlYKSFJNQkFmOEVBakFBTUIwR0ExVWREZ1FXQkJUYjB0RGs2WU5LRG9EMWxCWExocWtjTzc0Z1pUQWZCZ05WSFNNRQpHREFXZ0JUOENHNFdINCtJQXdYVWNFMFNVaDJhakowRVV6QUtCZ2dxaGtqT1BRUURBZ05JQURCRkFpRUE5cWl4CktmdC9hL1FqRXZMTXJXOVpyaDRnakVML01tVnpmSjlHZ0pwNUN5a0NJRVFUdmZBbkNHR1BHcVpJa0FNSUFpc1gKbVVkVTFreVZiai96bGhJa0JDSW0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo",
// 	cacerts[:])
// if err != nil {
// 	return
// }

// tls, err := service.NewMspConfigData(
// 	"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2ekNDQVpHZ0F3SUJBZ0lVV1lVRVNMV1FXMGZTRWhtVkxWLzNuWTcvYmFZd0NnWUlLb1pJemowRUF3SXcKYURFTE1Ba0dBMVVFQmhNQ1ZWTXhGekFWQmdOVkJBZ1REazV2Y25Sb0lFTmhjbTlzYVc1aE1SUXdFZ1lEVlFRSwpFd3RJZVhCbGNteGxaR2RsY2pFUE1BMEdBMVVFQ3hNR1JtRmljbWxqTVJrd0Z3WURWUVFERXhCbVlXSnlhV010ClkyRXRjMlZ5ZG1WeU1CNFhEVEl3TURFd09URTBORFV3TUZvWERUSXhNREV3T0RFME5UQXdNRm93SVRFUE1BMEcKQTFVRUN4TUdZMnhwWlc1ME1RNHdEQVlEVlFRREV3VmhaRzFwYmpCWk1CTUdCeXFHU000OUFnRUdDQ3FHU000OQpBd0VIQTBJQUJMc29Md1VDMGRCSkJlZEcwOXRyN2xuNm84T2JMWVgyZVJEZVByWlRzWm8yVjhPZjFBSkl1SEk0CmhEZHFSV0tITXRuamowUUMwK09WNEpYay9LSExtbytqWURCZU1BNEdBMVVkRHdFQi93UUVBd0lIZ0RBTUJnTlYKSFJNQkFmOEVBakFBTUIwR0ExVWREZ1FXQkJUYjB0RGs2WU5LRG9EMWxCWExocWtjTzc0Z1pUQWZCZ05WSFNNRQpHREFXZ0JUOENHNFdINCtJQXdYVWNFMFNVaDJhakowRVV6QUtCZ2dxaGtqT1BRUURBZ05JQURCRkFpRUE5cWl4CktmdC9hL1FqRXZMTXJXOVpyaDRnakVML01tVnpmSjlHZ0pwNUN5a0NJRVFUdmZBbkNHR1BHcVpJa0FNSUFpc1gKbVVkVTFreVZiai96bGhJa0JDSW0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo",
// 	"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUI2ekNDQVpHZ0F3SUJBZ0lVV1lVRVNMV1FXMGZTRWhtVkxWLzNuWTcvYmFZd0NnWUlLb1pJemowRUF3SXcKYURFTE1Ba0dBMVVFQmhNQ1ZWTXhGekFWQmdOVkJBZ1REazV2Y25Sb0lFTmhjbTlzYVc1aE1SUXdFZ1lEVlFRSwpFd3RJZVhCbGNteGxaR2RsY2pFUE1BMEdBMVVFQ3hNR1JtRmljbWxqTVJrd0Z3WURWUVFERXhCbVlXSnlhV010ClkyRXRjMlZ5ZG1WeU1CNFhEVEl3TURFd09URTBORFV3TUZvWERUSXhNREV3T0RFME5UQXdNRm93SVRFUE1BMEcKQTFVRUN4TUdZMnhwWlc1ME1RNHdEQVlEVlFRREV3VmhaRzFwYmpCWk1CTUdCeXFHU000OUFnRUdDQ3FHU000OQpBd0VIQTBJQUJMc29Md1VDMGRCSkJlZEcwOXRyN2xuNm84T2JMWVgyZVJEZVByWlRzWm8yVjhPZjFBSkl1SEk0CmhEZHFSV0tITXRuamowUUMwK09WNEpYay9LSExtbytqWURCZU1BNEdBMVVkRHdFQi93UUVBd0lIZ0RBTUJnTlYKSFJNQkFmOEVBakFBTUIwR0ExVWREZ1FXQkJUYjB0RGs2WU5LRG9EMWxCWExocWtjTzc0Z1pUQWZCZ05WSFNNRQpHREFXZ0JUOENHNFdINCtJQXdYVWNFMFNVaDJhakowRVV6QUtCZ2dxaGtqT1BRUURBZ05JQURCRkFpRUE5cWl4CktmdC9hL1FqRXZMTXJXOVpyaDRnakVML01tVnpmSjlHZ0pwNUN5a0NJRVFUdmZBbkNHR1BHcVpJa0FNSUFpc1gKbVVkVTFreVZiai96bGhJa0JDSW0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo",
// 	cacerts[:])
// if err != nil {
// 	return
// }
// configObjectMsp, err := service.NewConfigObjectMsp(component, tls)
// if err != nil {
// 	return
// }
