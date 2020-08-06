
package main

import (
    "fmt"
    "github.com/IBM/go-sdk-core/v4/core"
    "github.com/IBM-Blockchain/ibp-go-sdk/blockchainv2"
)

func main() {
    fmt.Println("start")

    // globals
    ApiKey:= "api-key"
    IdentityUrl:= "https://iam.test.cloud.ibm.com/identity/token"
    myserviceURL:= "http://localhost:3000"

    // Create an authenticator.
    authenticator := &core.IamAuthenticator{
        ApiKey: ApiKey,                                           // update field with your api key
        URL: IdentityUrl,
    }
    /*authenticator := &core.BearerTokenAuthenticator{
        BearerToken: "my IAM access token",                                  // alternatively update field with access token
    }*/

    // Create an instance of the "BlockchainV2Options"  struct.
    //  myserviceURL := "http://localhost:3000"  // update field with service instance url
    options := &blockchainv2.BlockchainV2Options{
        Authenticator: authenticator,
        URL: myserviceURL,
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

   
   
    // REMOVE COMPONENTS BY Tag
    // Create an authenticator
    authenticator = &core.IamAuthenticator{
        ApiKey: ApiKey,
        URL: IdentityUrl,
    }

    // Create an instance of the "BlockchainV2Options" struct
    options = &blockchainv2.BlockchainV2Options{
        Authenticator: authenticator,
        URL: myserviceURL,
    }

    // Create an instance of the "BlockchainV2" service client.
    service, err = blockchainv2.NewBlockchainV2(options)
    if err != nil {
        return
    }

    // Get all component data
    opts2 := service.NewRemoveComponentsByTagOptions("msp")
    result2, detailedResponse, err := service.RemoveComponentsByTag(opts2)
    fmt.Println("api key - lcsharp", ApiKey)
    fmt.Println("result:", result2)
    fmt.Println("response:", detailedResponse)
    fmt.Println("err:", err)

    fmt.Println("done")

    return
}

