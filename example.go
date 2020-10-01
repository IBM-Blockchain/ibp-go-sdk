
package main

import (
    "fmt"
    "github.com/IBM/go-sdk-core/v4/core"
    "github.com/IBM-Blockchain/ibp-go-sdk/blockchainv3"
)

func main() {
    fmt.Println("start")

    // Create an authenticator.
    authenticator := &core.IamAuthenticator{
        ApiKey: "my IAM api key",                                           // update field with your api key
    }
    /*authenticator := &core.BearerTokenAuthenticator{
        BearerToken: "my IAM access token",                                  // alternatively update field with access token
    }*/

    // Create an instance of the "BlockchainV3Options"  struct.
    myserviceURL := "https://my-ibp-console.uss01.blockchain.cloud.ibm.com"  // update field with service instance url
    options := &blockchainv3.BlockchainV3Options{
        Authenticator: authenticator,
        URL: myserviceURL,
    }

    // Create an instance of the "BlockchainV3" service client.
    service, err := blockchainv3.NewBlockchainV3(options)
    if err != nil {
        // handle error
    }

    // Example - List all components
    opts := service.NewListComponentsOptions()
    result, detailedResponse, err := service.ListComponents(opts)
    fmt.Println("result:", result)
    fmt.Println("response:", detailedResponse)
    fmt.Println("err:", err)

    fmt.Println("done")
    return
}

