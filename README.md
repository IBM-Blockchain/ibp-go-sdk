Blockchain Go SDK/Module

GoLang client library to use the IBM Cloud Blockchain **Service**.

**This module will allow you to use native golang functions to leverage the same functionality seen in the [IBP APIs](https://cloud.ibm.com/apidocs/blockchain)**

<details>
<summary>Table of Contents</summary>

* [Overview](#overview)
* [Prerequisites](#prerequisites)
* [Installation](#installation)
* [Using the SDK](#using-the-sdk)
  * [Constructing service clients](#constructing-service-clients)
  * [Authentication](#authentication)
  * [Passing operation parameters via an options struct](#passing-operation-parameters-via-an-options-struct)
  * [Receiving operation responses](#receiving-operation-responses)
  * [Error Handling](#error-handling)
  * [Default headers](#default-headers)
  * [Sending request headers](#sending-request-headers)
* [Explore the SDK](#explore-the-sdk)
* [Generation](#generation)
* [License](#license)

</details>

## Overview

The IBM Cloud Blockchain Go SDK allows developers to programmatically interact with the
IBM Cloud Blockchain service.

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration?target=%2Fdeveloper%2Fwatson&

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An [IBM Blockchain Platform Service instance](https://cloud.ibm.com/catalog/services/blockchain-platform)
* An IAM API key to allow the SDK to access your service instance. Create an account level api key [here](https://cloud.ibm.com/iam/apikeys) (alternatively you can create a service instance level api key from the IBM cloud UI).
* An installation of Go (version 1.12 or above) on your local machine.


## Installation
There are a few different ways to download and install the Blockchain Go SDK project for use by your
Go application:
##### 1. `go get` command
Use this command to download and install the Blockchain Go SDK project to allow your Go application to
use it:
```
go get -u github.com/IBM-Blockchain/ibp-go-sdk
```
##### 2. Go modules
If your application is using Go modules, you can add a suitable import to your
Go application, like this:
```go
import (
    "github.com/IBM-Blockchain/ibp-go-sdk/blockchainv2"
)
```
then run `go mod tidy` to download and install the new dependency and update your Go application's
`go.mod` file.
##### 3. `dep` dependency manager
If your application is using the `dep` dependency management tool, you can add a dependency
to your `Gopkg.toml` file.  Here is an example:
```
[[constraint]]
  name = "github.com/IBM-Blockchain/ibp-go-sdk/blockchainv2"
  version = "0.0.1"

```
then run `dep ensure`.

## Using the SDK
This section provides general information on how to use the services contained in this SDK.
### Constructing service clients
Each service is implemented in its own package (e.g. `blockchainv2`).
The package will contain a "service client"
struct (a client-side representation of the service), as well as an "options" struct that is used to
construct instances of the service client.  
Here's an example of how to construct an instance of "My Service":
```go
import (
    "github.com/IBM/go-sdk-core/core"
    "github.com/IBM-Blockchain/ibp-go-sdk/blockchainv2"
)

// Create an authenticator.
authenticator := /* create an authenticator - see examples below */

// Create an instance of the "BlockchainV2Options"  struct.
myserviceURL := "https://myservice.cloud.ibm.com/api"
options := &blockchainv2.BlockchainV2Options{
    Authenticator: authenticator,
    URL: myserviceURL,
}

// Create an instance of the "BlockchainV2" service client.
service, err := NewBlockchainV2(options)
if err != nil {
    // handle error
}

// Service operations can now be called using the "service" variable.

```

### Authentication
Blockchain services use token-based Identity and Access Management (IAM) authentication.

IAM authentication uses an API key to obtain an access token, which is then used to authenticate
each API request.  Access tokens are valid for a limited amount of time and must be regenerated.

To provide credentials to the SDK, you supply either an IAM service **API key** or an **access token**:

- Specify the IAM API key to have the SDK manage the lifecycle of the access token.
The SDK requests an access token, ensures that the access token is valid, and refreshes it when
necessary.
- Specify the access token if you want to manage the lifecycle yourself.
For details, see [Authenticating with IAM tokens](https://cloud.ibm.com/docs/services/watson/getting-started-iam.html).

##### Examples:
* Supplying the IAM API key and letting the SDK manage the access token for you:

```go
// letting the SDK manage the IAM access token
import (
    "github.com/IBM/go-sdk-core/core"
    "github.com/IBM-Blockchain/ibp-go-sdk/blockchainv2"
)
...
// Create the IAM authenticator.
authenticator := &core.IamAuthenticator{
    ApiKey: "myapikey",
}

// Create the service options struct.
options := &blockchainv2.BlockchainV2Options{
    Authenticator: authenticator,
}

// Construct the service instance.
service, err := blockchainv2.NewBlockchainV2(options)

```

* Supplying the access token (a bearer token) and managing it yourself:

```go
import (
    "github.com/IBM/go-sdk-core/core"
    "github.com/IBM-Blockchain/ibp-go-sdk/blockchainv2"
)
...
// Create the BearerToken authenticator.
authenticator := &core.BearerTokenAuthenticator{
    BearerToken: "my IAM access token",
}

// Create the service options struct.
options := &blockchainv2.BlockchainV2Options{
    Authenticator: authenticator,
}

// Construct the service instance.
service, err := blockchainv2.NewBlockchainV2(options)

...
// Later when the access token expires, the application must refresh the access token,
// then set the new access token on the authenticator.
// Subsequent request invocations will include the new access token.
authenticator.BearerToken = /* new access token */
```
For more information on authentication, including the full set of authentication schemes supported by
the underlying Go Core library, see
[this page](https://github.com/IBM/go-sdk-core/blob/master/Authentication.md)

### Passing operation parameters via an options struct
For each operation belonging to a service, an "options" struct is defined as a container for
the parameters associated with the operation.
The name of the struct will be `<operation-name>Options` and it will contain a field for each
operation parameter.  
Here's an example of an options struct for the `GetComponent` operation:
```go
// GetComponentOptions : The GetComponent options.
type GetComponentOptions struct {

    // The id of the resource to retrieve.
    ID *string `json:"resource_id" validate:"required"`

    ...
}
```
When invoking this operation, the application first creates an instance of the `GetComponentOptions`
struct and then sets the parameter values within it.  Along with the "options" struct, a constructor
function is also provided.  
Here's an example:
```go
options := service.NewGetComponentOptions("resource-id-1")
```
Then the operation can be called like this:
```go
result, detailedResponse, err := service.GetComponent(options)
```
This use of the "options" struct pattern (instead of listing each operation parameter within the
argument list of the service method) allows for future expansion of the API (within certain
guidelines) without impacting applications.

### Receiving operation responses

Each service method (operation) will return the following values:
1. `result` - An operation-specific result (if the operation is defined as returning a result).
2. `detailedResponse` - An instance of the `core.DetailedResponse` struct.
This will contain the following fields:
* `StatusCode` - the HTTP status code returned in the response message
* `Headers` - the HTTP headers returned in the response message
* `Result` - the operation result (if available). This is the same value returned in the `result` return value
mentioned above.
3. `err` - An error object.  This return value will be nil if the operation was successful, or non-nil
if unsuccessful.

##### Example:
1. Here's an example of calling the `GetComponent` operation which returns an instance of the `Resource`
struct as its result:
```go
// Construct the service instance.
service, err := blockchainv2.NewBlockchainV2(
    &blockchainv2.BlockchainV2Options{
        Authenticator: authenticator,
    })

// Call the GetComponent operation and receive the returned Resource.
options := service.NewGetComponentOptions("resource-id-1")
result, detailedResponse, err := service.GetComponent(options)

// Now use 'result' which should be an instance of 'Resource'.
```
2. Here's an example of calling the `DeleteResource` operation which does not return a response object:
```
// Construct the service instance.
service, err := blockchainv2.NewBlockchainV2(
    &blockchainv2.BlockchainV2Options{
        Authenticator: authenticator,
    })

// Call the DeleteResource operation and receive the returned Resource.
options := service.NewDeleteResourceOptions("resource-id-1")
detailedResponse, err := service.DeleteResource(options)
```

### Error Handling

In the case of an error response from the server endpoint, the Blockchain Go SDK will do the following:
1. The service method (operation) will return a non-nil `error` object.  This `error` object will
contain the error message retrieved from the HTTP response if possible, or a generic error message
otherwise.
2. The `detailedResponse.Result` field will contain the unmarshalled response (in the form of a
`map[string]interface{}`) if the operation returned a JSON response.  
This allows the application to examine all of the error information returned in the HTTP
response message.
3. The `detailedResponse.RawResult` field will contain the raw response body as a `[]byte` if the
operation returned a non-JSON response.

##### Example:
Here's an example of checking the `error` object after invoking the `GetComponent` operation:
```go
// Call the GetComponent operation and receive the returned Resource.
options := service.NewGetComponentOptions("bad-resource-id", "bad-resource-type")
result, detailedResponse, err := service.GetComponent(options)
if err != nil {
    fmt.Println("Error retrieving the resource: ", err.Error())
    fmt.Println("   full error response: ", detailedResponse.Result)
}
```

### Default headers
Default HTTP headers can be specified by using the `SetDefaultHeaders(http.Header)`
method of the client instance.  Once set on the service client, default headers are sent with
every outbound request.  
##### Example:
The example below sets the header `Custom-Header` with the value "custom_value" as a default
header:
```go
// Construct the service instance.
service, err := blockchainv2.NewBlockchainV2(
    &blockchainv2.BlockchainV2Options{
        Authenticator: authenticator,
    })

customHeaders := http.Header{}
customHeaders.Add("Custom-Header", "custom_value")
service.Service.SetDefaultHeaders(customHeaders)

// "Custom-Header" will now be included with all subsequent requests invoked from "service".
```

### Sending request headers
Custom HTTP headers can also be passed with any individual request.
To do so, add the headers to the "options" struct passed to the service method.
##### Example:
Here's an example that sets "Custom-Header" on the `GetComponentOptions` instance and then
invokes the `GetComponent` operation:
```go

// Call the GetComponent operation, passing our Custom-Header.
options := service.NewGetComponentOptions("resource-id-1")
customHeaders := make(map[string]interface{})
customHeaders["Custom-Header"] = "custom_value"
options.SetHeaders(customHeaders)
result, detailedResponse, err := service.GetComponent(options)
// "Custom-Header" will be sent along with the "GetComponent" request.
```

## Explore the SDK
This module is generated from the OpenAPI (swagger) file that populated the [IBP APIs documentation](https://cloud.ibm.com/apidocs/blockchain).
It is recommended to explore the IBP APIs documentation to find the desired functionality.
Then find the corresponding go example to the right of the api documentation.

Alternatively you could manually browse the SDK's main file [blockchain_v2.go](./blockchainv2/blockchain_v2.go).

## Generation
This is a note for developers of this repository on how to rebuild the SDK.
- this module was generated/built via the [IBM Cloud OpenAPI SDK generator](https://github.ibm.com/CloudEngineering/openapi-sdkgen)
    - [SDK generator overview](https://github.ibm.com/CloudEngineering/openapi-sdkgen/wiki/SDK-Gen-Overview)
    - [Configuration option code](https://github.ibm.com/CloudEngineering/openapi-sdkgen/blob/ab7d50a1dcdc707faad8cbe4f86de2d2ca510d24/src/main/java/com/ibm/sdk/codegen/IBMDefaultCodegen.java)
    - [IBP's OpenAPI source](https://github.ibm.com/cloud-api-docs/ibp/blob/master/ibp.yaml)
1. download the  latest sdk generator **release** (should see the java file `lib/openapi-sdkgen.jar`)
1. clone/download the IBP OpenAPI file
1. build command w/o shell: 
```
cd code/openapi-sdkgen
java -jar ./lib/openapi-sdkgen.jar generate -g ibm-go -i C:/code/cloud-api-docs/ibp.yaml -o C:/code/openapi-sdkgen/build --apiref C:/code/openapi-sdkgen/go-apiref.json
// inspect the files in C:/code/openapi-sdkgen/build and copy to this repo if they look okay
// copy file C:/code/openapi-sdkgen/go-apiref.json to the `cloud-api-docs` repo (this is IBP's IBM Cloud ApiDocs source)
```

## License

The IBM Cloud Blockchain Go SDK is released under the Apache 2.0 license. The license's full text can be found in [LICENSE](LICENSE).
