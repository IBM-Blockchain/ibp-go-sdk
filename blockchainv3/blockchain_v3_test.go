/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package blockchainv3_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/IBM-Blockchain/ibp-go-sdk/blockchainv3"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`BlockchainV3`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(blockchainService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(blockchainService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
				URL: "https://blockchainv3/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(blockchainService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BLOCKCHAIN_URL": "https://blockchainv3/api",
				"BLOCKCHAIN_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3UsingExternalConfig(&blockchainv3.BlockchainV3Options{
				})
				Expect(blockchainService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := blockchainService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != blockchainService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(blockchainService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(blockchainService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3UsingExternalConfig(&blockchainv3.BlockchainV3Options{
					URL: "https://testService/api",
				})
				Expect(blockchainService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := blockchainService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != blockchainService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(blockchainService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(blockchainService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3UsingExternalConfig(&blockchainv3.BlockchainV3Options{
				})
				err := blockchainService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := blockchainService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != blockchainService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(blockchainService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(blockchainService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BLOCKCHAIN_URL": "https://blockchainv3/api",
				"BLOCKCHAIN_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			blockchainService, serviceErr := blockchainv3.NewBlockchainV3UsingExternalConfig(&blockchainv3.BlockchainV3Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(blockchainService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BLOCKCHAIN_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			blockchainService, serviceErr := blockchainv3.NewBlockchainV3UsingExternalConfig(&blockchainv3.BlockchainV3Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(blockchainService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = blockchainv3.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetComponent(getComponentOptions *GetComponentOptions) - Operation response error`, func() {
		getComponentPath := "/ak/api/v3/components/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getComponentPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["deployment_attrs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["parsed_certs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["cache"]).To(Equal([]string{"skip"}))

					Expect(req.URL.Query()["ca_attrs"]).To(Equal([]string{"included"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetComponent with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the GetComponentOptions model
				getComponentOptionsModel := new(blockchainv3.GetComponentOptions)
				getComponentOptionsModel.ID = core.StringPtr("testString")
				getComponentOptionsModel.DeploymentAttrs = core.StringPtr("included")
				getComponentOptionsModel.ParsedCerts = core.StringPtr("included")
				getComponentOptionsModel.Cache = core.StringPtr("skip")
				getComponentOptionsModel.CaAttrs = core.StringPtr("included")
				getComponentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.GetComponent(getComponentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.GetComponent(getComponentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetComponent(getComponentOptions *GetComponentOptions)`, func() {
		getComponentPath := "/ak/api/v3/components/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getComponentPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["deployment_attrs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["parsed_certs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["cache"]).To(Equal([]string{"skip"}))

					Expect(req.URL.Query()["ca_attrs"]).To(Equal([]string{"included"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "myca-2", "type": "fabric-ca", "display_name": "Example CA", "cluster_id": "mzdqhdifnl", "cluster_name": "ordering service 1", "grpcwp_url": "https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084", "api_url": "grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051", "operations_url": "https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443", "msp": {"ca": {"name": "org1CA", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "tlsca": {"name": "org1tlsCA", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "component": {"tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "ecert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "admin_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}}, "msp_id": "Org1", "location": "ibmcloud", "node_ou": {"enabled": true}, "resources": {"ca": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "peer": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "orderer": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "proxy": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "statedb": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "state_db": "couchdb", "storage": {"ca": {"size": "4GiB", "class": "default"}, "peer": {"size": "4GiB", "class": "default"}, "orderer": {"size": "4GiB", "class": "default"}, "statedb": {"size": "4GiB", "class": "default"}}, "timestamp": 1537262855753, "tags": ["fabric-ca"], "version": "1.4.6-1", "zone": "-"}`)
				}))
			})
			It(`Invoke GetComponent successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.GetComponent(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetComponentOptions model
				getComponentOptionsModel := new(blockchainv3.GetComponentOptions)
				getComponentOptionsModel.ID = core.StringPtr("testString")
				getComponentOptionsModel.DeploymentAttrs = core.StringPtr("included")
				getComponentOptionsModel.ParsedCerts = core.StringPtr("included")
				getComponentOptionsModel.Cache = core.StringPtr("skip")
				getComponentOptionsModel.CaAttrs = core.StringPtr("included")
				getComponentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.GetComponent(getComponentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.GetComponentWithContext(ctx, getComponentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.GetComponent(getComponentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.GetComponentWithContext(ctx, getComponentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetComponent with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the GetComponentOptions model
				getComponentOptionsModel := new(blockchainv3.GetComponentOptions)
				getComponentOptionsModel.ID = core.StringPtr("testString")
				getComponentOptionsModel.DeploymentAttrs = core.StringPtr("included")
				getComponentOptionsModel.ParsedCerts = core.StringPtr("included")
				getComponentOptionsModel.Cache = core.StringPtr("skip")
				getComponentOptionsModel.CaAttrs = core.StringPtr("included")
				getComponentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.GetComponent(getComponentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetComponentOptions model with no property values
				getComponentOptionsModelNew := new(blockchainv3.GetComponentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.GetComponent(getComponentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RemoveComponent(removeComponentOptions *RemoveComponentOptions) - Operation response error`, func() {
		removeComponentPath := "/ak/api/v3/components/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(removeComponentPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RemoveComponent with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the RemoveComponentOptions model
				removeComponentOptionsModel := new(blockchainv3.RemoveComponentOptions)
				removeComponentOptionsModel.ID = core.StringPtr("testString")
				removeComponentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.RemoveComponent(removeComponentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.RemoveComponent(removeComponentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`RemoveComponent(removeComponentOptions *RemoveComponentOptions)`, func() {
		removeComponentPath := "/ak/api/v3/components/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(removeComponentPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"message": "deleted", "type": "fabric-peer", "id": "component1", "display_name": "My Peer"}`)
				}))
			})
			It(`Invoke RemoveComponent successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.RemoveComponent(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RemoveComponentOptions model
				removeComponentOptionsModel := new(blockchainv3.RemoveComponentOptions)
				removeComponentOptionsModel.ID = core.StringPtr("testString")
				removeComponentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.RemoveComponent(removeComponentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.RemoveComponentWithContext(ctx, removeComponentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.RemoveComponent(removeComponentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.RemoveComponentWithContext(ctx, removeComponentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke RemoveComponent with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the RemoveComponentOptions model
				removeComponentOptionsModel := new(blockchainv3.RemoveComponentOptions)
				removeComponentOptionsModel.ID = core.StringPtr("testString")
				removeComponentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.RemoveComponent(removeComponentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RemoveComponentOptions model with no property values
				removeComponentOptionsModelNew := new(blockchainv3.RemoveComponentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.RemoveComponent(removeComponentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteComponent(deleteComponentOptions *DeleteComponentOptions) - Operation response error`, func() {
		deleteComponentPath := "/ak/api/v3/kubernetes/components/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteComponentPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteComponent with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the DeleteComponentOptions model
				deleteComponentOptionsModel := new(blockchainv3.DeleteComponentOptions)
				deleteComponentOptionsModel.ID = core.StringPtr("testString")
				deleteComponentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.DeleteComponent(deleteComponentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.DeleteComponent(deleteComponentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteComponent(deleteComponentOptions *DeleteComponentOptions)`, func() {
		deleteComponentPath := "/ak/api/v3/kubernetes/components/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteComponentPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"message": "deleted", "type": "fabric-peer", "id": "component1", "display_name": "My Peer"}`)
				}))
			})
			It(`Invoke DeleteComponent successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.DeleteComponent(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteComponentOptions model
				deleteComponentOptionsModel := new(blockchainv3.DeleteComponentOptions)
				deleteComponentOptionsModel.ID = core.StringPtr("testString")
				deleteComponentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.DeleteComponent(deleteComponentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.DeleteComponentWithContext(ctx, deleteComponentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.DeleteComponent(deleteComponentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.DeleteComponentWithContext(ctx, deleteComponentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke DeleteComponent with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the DeleteComponentOptions model
				deleteComponentOptionsModel := new(blockchainv3.DeleteComponentOptions)
				deleteComponentOptionsModel.ID = core.StringPtr("testString")
				deleteComponentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.DeleteComponent(deleteComponentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteComponentOptions model with no property values
				deleteComponentOptionsModelNew := new(blockchainv3.DeleteComponentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.DeleteComponent(deleteComponentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCa(createCaOptions *CreateCaOptions) - Operation response error`, func() {
		createCaPath := "/ak/api/v3/kubernetes/components/fabric-ca"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCaPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCa with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the ConfigCACors model
				configCaCorsModel := new(blockchainv3.ConfigCACors)
				configCaCorsModel.Enabled = core.BoolPtr(true)
				configCaCorsModel.Origins = []string{"*"}

				// Construct an instance of the ConfigCATlsClientauth model
				configCaTlsClientauthModel := new(blockchainv3.ConfigCATlsClientauth)
				configCaTlsClientauthModel.Type = core.StringPtr("noclientcert")
				configCaTlsClientauthModel.Certfiles = []string{"testString"}

				// Construct an instance of the ConfigCATls model
				configCaTlsModel := new(blockchainv3.ConfigCATls)
				configCaTlsModel.Keyfile = core.StringPtr("testString")
				configCaTlsModel.Certfile = core.StringPtr("testString")
				configCaTlsModel.Clientauth = configCaTlsClientauthModel

				// Construct an instance of the ConfigCACa model
				configCaCaModel := new(blockchainv3.ConfigCACa)
				configCaCaModel.Keyfile = core.StringPtr("testString")
				configCaCaModel.Certfile = core.StringPtr("testString")
				configCaCaModel.Chainfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCACrl model
				configCaCrlModel := new(blockchainv3.ConfigCACrl)
				configCaCrlModel.Expiry = core.StringPtr("24h")

				// Construct an instance of the IdentityAttrs model
				identityAttrsModel := new(blockchainv3.IdentityAttrs)
				identityAttrsModel.HfRegistrarRoles = core.StringPtr("*")
				identityAttrsModel.HfRegistrarDelegateRoles = core.StringPtr("*")
				identityAttrsModel.HfRevoker = core.BoolPtr(true)
				identityAttrsModel.HfIntermediateCA = core.BoolPtr(true)
				identityAttrsModel.HfGenCRL = core.BoolPtr(true)
				identityAttrsModel.HfRegistrarAttributes = core.StringPtr("*")
				identityAttrsModel.HfAffiliationMgr = core.BoolPtr(true)

				// Construct an instance of the ConfigCARegistryIdentitiesItem model
				configCaRegistryIdentitiesItemModel := new(blockchainv3.ConfigCARegistryIdentitiesItem)
				configCaRegistryIdentitiesItemModel.Name = core.StringPtr("admin")
				configCaRegistryIdentitiesItemModel.Pass = core.StringPtr("password")
				configCaRegistryIdentitiesItemModel.Type = core.StringPtr("client")
				configCaRegistryIdentitiesItemModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryIdentitiesItemModel.Affiliation = core.StringPtr("testString")
				configCaRegistryIdentitiesItemModel.Attrs = identityAttrsModel

				// Construct an instance of the ConfigCARegistry model
				configCaRegistryModel := new(blockchainv3.ConfigCARegistry)
				configCaRegistryModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryModel.Identities = []blockchainv3.ConfigCARegistryIdentitiesItem{*configCaRegistryIdentitiesItemModel}

				// Construct an instance of the ConfigCADbTlsClient model
				configCaDbTlsClientModel := new(blockchainv3.ConfigCADbTlsClient)
				configCaDbTlsClientModel.Certfile = core.StringPtr("testString")
				configCaDbTlsClientModel.Keyfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCADbTls model
				configCaDbTlsModel := new(blockchainv3.ConfigCADbTls)
				configCaDbTlsModel.Certfiles = []string{"testString"}
				configCaDbTlsModel.Client = configCaDbTlsClientModel
				configCaDbTlsModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the ConfigCADb model
				configCaDbModel := new(blockchainv3.ConfigCADb)
				configCaDbModel.Type = core.StringPtr("postgres")
				configCaDbModel.Datasource = core.StringPtr("host=fake.databases.appdomain.cloud port=31941 user=ibm_cloud password=password dbname=ibmclouddb sslmode=verify-full")
				configCaDbModel.Tls = configCaDbTlsModel

				// Construct an instance of the ConfigCAAffiliations model
				configCaAffiliationsModel := new(blockchainv3.ConfigCAAffiliations)
				configCaAffiliationsModel.Org1 = []string{"department1"}
				configCaAffiliationsModel.Org2 = []string{"department1"}
				configCaAffiliationsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ConfigCACsrKeyrequest model
				configCaCsrKeyrequestModel := new(blockchainv3.ConfigCACsrKeyrequest)
				configCaCsrKeyrequestModel.Algo = core.StringPtr("ecdsa")
				configCaCsrKeyrequestModel.Size = core.Float64Ptr(float64(256))

				// Construct an instance of the ConfigCACsrNamesItem model
				configCaCsrNamesItemModel := new(blockchainv3.ConfigCACsrNamesItem)
				configCaCsrNamesItemModel.C = core.StringPtr("US")
				configCaCsrNamesItemModel.ST = core.StringPtr("North Carolina")
				configCaCsrNamesItemModel.L = core.StringPtr("Raleigh")
				configCaCsrNamesItemModel.O = core.StringPtr("Hyperledger")
				configCaCsrNamesItemModel.OU = core.StringPtr("Fabric")

				// Construct an instance of the ConfigCACsrCa model
				configCaCsrCaModel := new(blockchainv3.ConfigCACsrCa)
				configCaCsrCaModel.Expiry = core.StringPtr("131400h")
				configCaCsrCaModel.Pathlength = core.Float64Ptr(float64(0))

				// Construct an instance of the ConfigCACsr model
				configCaCsrModel := new(blockchainv3.ConfigCACsr)
				configCaCsrModel.Cn = core.StringPtr("ca")
				configCaCsrModel.Keyrequest = configCaCsrKeyrequestModel
				configCaCsrModel.Names = []blockchainv3.ConfigCACsrNamesItem{*configCaCsrNamesItemModel}
				configCaCsrModel.Hosts = []string{"localhost"}
				configCaCsrModel.Ca = configCaCsrCaModel

				// Construct an instance of the ConfigCAIdemix model
				configCaIdemixModel := new(blockchainv3.ConfigCAIdemix)
				configCaIdemixModel.Rhpoolsize = core.Float64Ptr(float64(100))
				configCaIdemixModel.Nonceexpiration = core.StringPtr("15s")
				configCaIdemixModel.Noncesweepinterval = core.StringPtr("15m")

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv3.BccspSW)
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv3.BccspPKCS11)
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv3.Bccsp)
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model

				// Construct an instance of the ConfigCAIntermediateParentserver model
				configCaIntermediateParentserverModel := new(blockchainv3.ConfigCAIntermediateParentserver)
				configCaIntermediateParentserverModel.URL = core.StringPtr("testString")
				configCaIntermediateParentserverModel.Caname = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateEnrollment model
				configCaIntermediateEnrollmentModel := new(blockchainv3.ConfigCAIntermediateEnrollment)
				configCaIntermediateEnrollmentModel.Hosts = core.StringPtr("localhost")
				configCaIntermediateEnrollmentModel.Profile = core.StringPtr("testString")
				configCaIntermediateEnrollmentModel.Label = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateTlsClient model
				configCaIntermediateTlsClientModel := new(blockchainv3.ConfigCAIntermediateTlsClient)
				configCaIntermediateTlsClientModel.Certfile = core.StringPtr("testString")
				configCaIntermediateTlsClientModel.Keyfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateTls model
				configCaIntermediateTlsModel := new(blockchainv3.ConfigCAIntermediateTls)
				configCaIntermediateTlsModel.Certfiles = []string{"testString"}
				configCaIntermediateTlsModel.Client = configCaIntermediateTlsClientModel

				// Construct an instance of the ConfigCAIntermediate model
				configCaIntermediateModel := new(blockchainv3.ConfigCAIntermediate)
				configCaIntermediateModel.Parentserver = configCaIntermediateParentserverModel
				configCaIntermediateModel.Enrollment = configCaIntermediateEnrollmentModel
				configCaIntermediateModel.Tls = configCaIntermediateTlsModel

				// Construct an instance of the ConfigCACfgIdentities model
				configCaCfgIdentitiesModel := new(blockchainv3.ConfigCACfgIdentities)
				configCaCfgIdentitiesModel.Passwordattempts = core.Float64Ptr(float64(10))
				configCaCfgIdentitiesModel.Allowremove = core.BoolPtr(false)

				// Construct an instance of the ConfigCACfg model
				configCaCfgModel := new(blockchainv3.ConfigCACfg)
				configCaCfgModel.Identities = configCaCfgIdentitiesModel

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv3.MetricsStatsd)
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv3.Metrics)
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel

				// Construct an instance of the ConfigCASigningDefault model
				configCaSigningDefaultModel := new(blockchainv3.ConfigCASigningDefault)
				configCaSigningDefaultModel.Usage = []string{"cert sign"}
				configCaSigningDefaultModel.Expiry = core.StringPtr("8760h")

				// Construct an instance of the ConfigCASigningProfilesCaCaconstraint model
				configCaSigningProfilesCaCaconstraintModel := new(blockchainv3.ConfigCASigningProfilesCaCaconstraint)
				configCaSigningProfilesCaCaconstraintModel.Isca = core.BoolPtr(true)
				configCaSigningProfilesCaCaconstraintModel.Maxpathlen = core.Float64Ptr(float64(0))
				configCaSigningProfilesCaCaconstraintModel.Maxpathlenzero = core.BoolPtr(true)

				// Construct an instance of the ConfigCASigningProfilesCa model
				configCaSigningProfilesCaModel := new(blockchainv3.ConfigCASigningProfilesCa)
				configCaSigningProfilesCaModel.Usage = []string{"cert sign"}
				configCaSigningProfilesCaModel.Expiry = core.StringPtr("43800h")
				configCaSigningProfilesCaModel.Caconstraint = configCaSigningProfilesCaCaconstraintModel

				// Construct an instance of the ConfigCASigningProfilesTls model
				configCaSigningProfilesTlsModel := new(blockchainv3.ConfigCASigningProfilesTls)
				configCaSigningProfilesTlsModel.Usage = []string{"cert sign"}
				configCaSigningProfilesTlsModel.Expiry = core.StringPtr("43800h")

				// Construct an instance of the ConfigCASigningProfiles model
				configCaSigningProfilesModel := new(blockchainv3.ConfigCASigningProfiles)
				configCaSigningProfilesModel.Ca = configCaSigningProfilesCaModel
				configCaSigningProfilesModel.Tls = configCaSigningProfilesTlsModel

				// Construct an instance of the ConfigCASigning model
				configCaSigningModel := new(blockchainv3.ConfigCASigning)
				configCaSigningModel.Default = configCaSigningDefaultModel
				configCaSigningModel.Profiles = configCaSigningProfilesModel

				// Construct an instance of the ConfigCACreate model
				configCaCreateModel := new(blockchainv3.ConfigCACreate)
				configCaCreateModel.Cors = configCaCorsModel
				configCaCreateModel.Debug = core.BoolPtr(false)
				configCaCreateModel.Crlsizelimit = core.Float64Ptr(float64(512000))
				configCaCreateModel.Tls = configCaTlsModel
				configCaCreateModel.Ca = configCaCaModel
				configCaCreateModel.Crl = configCaCrlModel
				configCaCreateModel.Registry = configCaRegistryModel
				configCaCreateModel.Db = configCaDbModel
				configCaCreateModel.Affiliations = configCaAffiliationsModel
				configCaCreateModel.Csr = configCaCsrModel
				configCaCreateModel.Idemix = configCaIdemixModel
				configCaCreateModel.BCCSP = bccspModel
				configCaCreateModel.Intermediate = configCaIntermediateModel
				configCaCreateModel.Cfg = configCaCfgModel
				configCaCreateModel.Metrics = metricsModel
				configCaCreateModel.Signing = configCaSigningModel

				// Construct an instance of the CreateCaBodyConfigOverride model
				createCaBodyConfigOverrideModel := new(blockchainv3.CreateCaBodyConfigOverride)
				createCaBodyConfigOverrideModel.Ca = configCaCreateModel
				createCaBodyConfigOverrideModel.Tlsca = configCaCreateModel

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv3.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv3.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv3.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the CreateCaBodyResources model
				createCaBodyResourcesModel := new(blockchainv3.CreateCaBodyResources)
				createCaBodyResourcesModel.Ca = resourceObjectModel

				// Construct an instance of the StorageObject model
				storageObjectModel := new(blockchainv3.StorageObject)
				storageObjectModel.Size = core.StringPtr("4GiB")
				storageObjectModel.Class = core.StringPtr("default")

				// Construct an instance of the CreateCaBodyStorage model
				createCaBodyStorageModel := new(blockchainv3.CreateCaBodyStorage)
				createCaBodyStorageModel.Ca = storageObjectModel

				// Construct an instance of the Hsm model
				hsmModel := new(blockchainv3.Hsm)
				hsmModel.Pkcs11endpoint = core.StringPtr("tcp://example.com:666")

				// Construct an instance of the CreateCaOptions model
				createCaOptionsModel := new(blockchainv3.CreateCaOptions)
				createCaOptionsModel.DisplayName = core.StringPtr("My CA")
				createCaOptionsModel.ConfigOverride = createCaBodyConfigOverrideModel
				createCaOptionsModel.ID = core.StringPtr("component1")
				createCaOptionsModel.Resources = createCaBodyResourcesModel
				createCaOptionsModel.Storage = createCaBodyStorageModel
				createCaOptionsModel.Zone = core.StringPtr("-")
				createCaOptionsModel.Replicas = core.Float64Ptr(float64(1))
				createCaOptionsModel.Tags = []string{"fabric-ca"}
				createCaOptionsModel.Hsm = hsmModel
				createCaOptionsModel.Region = core.StringPtr("-")
				createCaOptionsModel.Version = core.StringPtr("1.4.6-1")
				createCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.CreateCa(createCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.CreateCa(createCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateCa(createCaOptions *CreateCaOptions)`, func() {
		createCaPath := "/ak/api/v3/kubernetes/components/fabric-ca"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCaPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "component1", "dep_component_id": "admin", "display_name": "My CA", "api_url": "grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051", "operations_url": "https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443", "config_override": {"anyKey": "anyValue"}, "location": "ibmcloud", "msp": {"ca": {"name": "ca", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "tlsca": {"name": "tlsca", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "component": {"tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "ecert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "admin_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}}, "resources": {"ca": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "storage": {"ca": {"size": "4GiB", "class": "default"}}, "tags": ["fabric-ca"], "timestamp": 1537262855753, "version": "1.4.6-1", "zone": "-"}`)
				}))
			})
			It(`Invoke CreateCa successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.CreateCa(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ConfigCACors model
				configCaCorsModel := new(blockchainv3.ConfigCACors)
				configCaCorsModel.Enabled = core.BoolPtr(true)
				configCaCorsModel.Origins = []string{"*"}

				// Construct an instance of the ConfigCATlsClientauth model
				configCaTlsClientauthModel := new(blockchainv3.ConfigCATlsClientauth)
				configCaTlsClientauthModel.Type = core.StringPtr("noclientcert")
				configCaTlsClientauthModel.Certfiles = []string{"testString"}

				// Construct an instance of the ConfigCATls model
				configCaTlsModel := new(blockchainv3.ConfigCATls)
				configCaTlsModel.Keyfile = core.StringPtr("testString")
				configCaTlsModel.Certfile = core.StringPtr("testString")
				configCaTlsModel.Clientauth = configCaTlsClientauthModel

				// Construct an instance of the ConfigCACa model
				configCaCaModel := new(blockchainv3.ConfigCACa)
				configCaCaModel.Keyfile = core.StringPtr("testString")
				configCaCaModel.Certfile = core.StringPtr("testString")
				configCaCaModel.Chainfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCACrl model
				configCaCrlModel := new(blockchainv3.ConfigCACrl)
				configCaCrlModel.Expiry = core.StringPtr("24h")

				// Construct an instance of the IdentityAttrs model
				identityAttrsModel := new(blockchainv3.IdentityAttrs)
				identityAttrsModel.HfRegistrarRoles = core.StringPtr("*")
				identityAttrsModel.HfRegistrarDelegateRoles = core.StringPtr("*")
				identityAttrsModel.HfRevoker = core.BoolPtr(true)
				identityAttrsModel.HfIntermediateCA = core.BoolPtr(true)
				identityAttrsModel.HfGenCRL = core.BoolPtr(true)
				identityAttrsModel.HfRegistrarAttributes = core.StringPtr("*")
				identityAttrsModel.HfAffiliationMgr = core.BoolPtr(true)

				// Construct an instance of the ConfigCARegistryIdentitiesItem model
				configCaRegistryIdentitiesItemModel := new(blockchainv3.ConfigCARegistryIdentitiesItem)
				configCaRegistryIdentitiesItemModel.Name = core.StringPtr("admin")
				configCaRegistryIdentitiesItemModel.Pass = core.StringPtr("password")
				configCaRegistryIdentitiesItemModel.Type = core.StringPtr("client")
				configCaRegistryIdentitiesItemModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryIdentitiesItemModel.Affiliation = core.StringPtr("testString")
				configCaRegistryIdentitiesItemModel.Attrs = identityAttrsModel

				// Construct an instance of the ConfigCARegistry model
				configCaRegistryModel := new(blockchainv3.ConfigCARegistry)
				configCaRegistryModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryModel.Identities = []blockchainv3.ConfigCARegistryIdentitiesItem{*configCaRegistryIdentitiesItemModel}

				// Construct an instance of the ConfigCADbTlsClient model
				configCaDbTlsClientModel := new(blockchainv3.ConfigCADbTlsClient)
				configCaDbTlsClientModel.Certfile = core.StringPtr("testString")
				configCaDbTlsClientModel.Keyfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCADbTls model
				configCaDbTlsModel := new(blockchainv3.ConfigCADbTls)
				configCaDbTlsModel.Certfiles = []string{"testString"}
				configCaDbTlsModel.Client = configCaDbTlsClientModel
				configCaDbTlsModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the ConfigCADb model
				configCaDbModel := new(blockchainv3.ConfigCADb)
				configCaDbModel.Type = core.StringPtr("postgres")
				configCaDbModel.Datasource = core.StringPtr("host=fake.databases.appdomain.cloud port=31941 user=ibm_cloud password=password dbname=ibmclouddb sslmode=verify-full")
				configCaDbModel.Tls = configCaDbTlsModel

				// Construct an instance of the ConfigCAAffiliations model
				configCaAffiliationsModel := new(blockchainv3.ConfigCAAffiliations)
				configCaAffiliationsModel.Org1 = []string{"department1"}
				configCaAffiliationsModel.Org2 = []string{"department1"}
				configCaAffiliationsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ConfigCACsrKeyrequest model
				configCaCsrKeyrequestModel := new(blockchainv3.ConfigCACsrKeyrequest)
				configCaCsrKeyrequestModel.Algo = core.StringPtr("ecdsa")
				configCaCsrKeyrequestModel.Size = core.Float64Ptr(float64(256))

				// Construct an instance of the ConfigCACsrNamesItem model
				configCaCsrNamesItemModel := new(blockchainv3.ConfigCACsrNamesItem)
				configCaCsrNamesItemModel.C = core.StringPtr("US")
				configCaCsrNamesItemModel.ST = core.StringPtr("North Carolina")
				configCaCsrNamesItemModel.L = core.StringPtr("Raleigh")
				configCaCsrNamesItemModel.O = core.StringPtr("Hyperledger")
				configCaCsrNamesItemModel.OU = core.StringPtr("Fabric")

				// Construct an instance of the ConfigCACsrCa model
				configCaCsrCaModel := new(blockchainv3.ConfigCACsrCa)
				configCaCsrCaModel.Expiry = core.StringPtr("131400h")
				configCaCsrCaModel.Pathlength = core.Float64Ptr(float64(0))

				// Construct an instance of the ConfigCACsr model
				configCaCsrModel := new(blockchainv3.ConfigCACsr)
				configCaCsrModel.Cn = core.StringPtr("ca")
				configCaCsrModel.Keyrequest = configCaCsrKeyrequestModel
				configCaCsrModel.Names = []blockchainv3.ConfigCACsrNamesItem{*configCaCsrNamesItemModel}
				configCaCsrModel.Hosts = []string{"localhost"}
				configCaCsrModel.Ca = configCaCsrCaModel

				// Construct an instance of the ConfigCAIdemix model
				configCaIdemixModel := new(blockchainv3.ConfigCAIdemix)
				configCaIdemixModel.Rhpoolsize = core.Float64Ptr(float64(100))
				configCaIdemixModel.Nonceexpiration = core.StringPtr("15s")
				configCaIdemixModel.Noncesweepinterval = core.StringPtr("15m")

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv3.BccspSW)
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv3.BccspPKCS11)
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv3.Bccsp)
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model

				// Construct an instance of the ConfigCAIntermediateParentserver model
				configCaIntermediateParentserverModel := new(blockchainv3.ConfigCAIntermediateParentserver)
				configCaIntermediateParentserverModel.URL = core.StringPtr("testString")
				configCaIntermediateParentserverModel.Caname = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateEnrollment model
				configCaIntermediateEnrollmentModel := new(blockchainv3.ConfigCAIntermediateEnrollment)
				configCaIntermediateEnrollmentModel.Hosts = core.StringPtr("localhost")
				configCaIntermediateEnrollmentModel.Profile = core.StringPtr("testString")
				configCaIntermediateEnrollmentModel.Label = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateTlsClient model
				configCaIntermediateTlsClientModel := new(blockchainv3.ConfigCAIntermediateTlsClient)
				configCaIntermediateTlsClientModel.Certfile = core.StringPtr("testString")
				configCaIntermediateTlsClientModel.Keyfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateTls model
				configCaIntermediateTlsModel := new(blockchainv3.ConfigCAIntermediateTls)
				configCaIntermediateTlsModel.Certfiles = []string{"testString"}
				configCaIntermediateTlsModel.Client = configCaIntermediateTlsClientModel

				// Construct an instance of the ConfigCAIntermediate model
				configCaIntermediateModel := new(blockchainv3.ConfigCAIntermediate)
				configCaIntermediateModel.Parentserver = configCaIntermediateParentserverModel
				configCaIntermediateModel.Enrollment = configCaIntermediateEnrollmentModel
				configCaIntermediateModel.Tls = configCaIntermediateTlsModel

				// Construct an instance of the ConfigCACfgIdentities model
				configCaCfgIdentitiesModel := new(blockchainv3.ConfigCACfgIdentities)
				configCaCfgIdentitiesModel.Passwordattempts = core.Float64Ptr(float64(10))
				configCaCfgIdentitiesModel.Allowremove = core.BoolPtr(false)

				// Construct an instance of the ConfigCACfg model
				configCaCfgModel := new(blockchainv3.ConfigCACfg)
				configCaCfgModel.Identities = configCaCfgIdentitiesModel

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv3.MetricsStatsd)
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv3.Metrics)
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel

				// Construct an instance of the ConfigCASigningDefault model
				configCaSigningDefaultModel := new(blockchainv3.ConfigCASigningDefault)
				configCaSigningDefaultModel.Usage = []string{"cert sign"}
				configCaSigningDefaultModel.Expiry = core.StringPtr("8760h")

				// Construct an instance of the ConfigCASigningProfilesCaCaconstraint model
				configCaSigningProfilesCaCaconstraintModel := new(blockchainv3.ConfigCASigningProfilesCaCaconstraint)
				configCaSigningProfilesCaCaconstraintModel.Isca = core.BoolPtr(true)
				configCaSigningProfilesCaCaconstraintModel.Maxpathlen = core.Float64Ptr(float64(0))
				configCaSigningProfilesCaCaconstraintModel.Maxpathlenzero = core.BoolPtr(true)

				// Construct an instance of the ConfigCASigningProfilesCa model
				configCaSigningProfilesCaModel := new(blockchainv3.ConfigCASigningProfilesCa)
				configCaSigningProfilesCaModel.Usage = []string{"cert sign"}
				configCaSigningProfilesCaModel.Expiry = core.StringPtr("43800h")
				configCaSigningProfilesCaModel.Caconstraint = configCaSigningProfilesCaCaconstraintModel

				// Construct an instance of the ConfigCASigningProfilesTls model
				configCaSigningProfilesTlsModel := new(blockchainv3.ConfigCASigningProfilesTls)
				configCaSigningProfilesTlsModel.Usage = []string{"cert sign"}
				configCaSigningProfilesTlsModel.Expiry = core.StringPtr("43800h")

				// Construct an instance of the ConfigCASigningProfiles model
				configCaSigningProfilesModel := new(blockchainv3.ConfigCASigningProfiles)
				configCaSigningProfilesModel.Ca = configCaSigningProfilesCaModel
				configCaSigningProfilesModel.Tls = configCaSigningProfilesTlsModel

				// Construct an instance of the ConfigCASigning model
				configCaSigningModel := new(blockchainv3.ConfigCASigning)
				configCaSigningModel.Default = configCaSigningDefaultModel
				configCaSigningModel.Profiles = configCaSigningProfilesModel

				// Construct an instance of the ConfigCACreate model
				configCaCreateModel := new(blockchainv3.ConfigCACreate)
				configCaCreateModel.Cors = configCaCorsModel
				configCaCreateModel.Debug = core.BoolPtr(false)
				configCaCreateModel.Crlsizelimit = core.Float64Ptr(float64(512000))
				configCaCreateModel.Tls = configCaTlsModel
				configCaCreateModel.Ca = configCaCaModel
				configCaCreateModel.Crl = configCaCrlModel
				configCaCreateModel.Registry = configCaRegistryModel
				configCaCreateModel.Db = configCaDbModel
				configCaCreateModel.Affiliations = configCaAffiliationsModel
				configCaCreateModel.Csr = configCaCsrModel
				configCaCreateModel.Idemix = configCaIdemixModel
				configCaCreateModel.BCCSP = bccspModel
				configCaCreateModel.Intermediate = configCaIntermediateModel
				configCaCreateModel.Cfg = configCaCfgModel
				configCaCreateModel.Metrics = metricsModel
				configCaCreateModel.Signing = configCaSigningModel

				// Construct an instance of the CreateCaBodyConfigOverride model
				createCaBodyConfigOverrideModel := new(blockchainv3.CreateCaBodyConfigOverride)
				createCaBodyConfigOverrideModel.Ca = configCaCreateModel
				createCaBodyConfigOverrideModel.Tlsca = configCaCreateModel

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv3.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv3.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv3.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the CreateCaBodyResources model
				createCaBodyResourcesModel := new(blockchainv3.CreateCaBodyResources)
				createCaBodyResourcesModel.Ca = resourceObjectModel

				// Construct an instance of the StorageObject model
				storageObjectModel := new(blockchainv3.StorageObject)
				storageObjectModel.Size = core.StringPtr("4GiB")
				storageObjectModel.Class = core.StringPtr("default")

				// Construct an instance of the CreateCaBodyStorage model
				createCaBodyStorageModel := new(blockchainv3.CreateCaBodyStorage)
				createCaBodyStorageModel.Ca = storageObjectModel

				// Construct an instance of the Hsm model
				hsmModel := new(blockchainv3.Hsm)
				hsmModel.Pkcs11endpoint = core.StringPtr("tcp://example.com:666")

				// Construct an instance of the CreateCaOptions model
				createCaOptionsModel := new(blockchainv3.CreateCaOptions)
				createCaOptionsModel.DisplayName = core.StringPtr("My CA")
				createCaOptionsModel.ConfigOverride = createCaBodyConfigOverrideModel
				createCaOptionsModel.ID = core.StringPtr("component1")
				createCaOptionsModel.Resources = createCaBodyResourcesModel
				createCaOptionsModel.Storage = createCaBodyStorageModel
				createCaOptionsModel.Zone = core.StringPtr("-")
				createCaOptionsModel.Replicas = core.Float64Ptr(float64(1))
				createCaOptionsModel.Tags = []string{"fabric-ca"}
				createCaOptionsModel.Hsm = hsmModel
				createCaOptionsModel.Region = core.StringPtr("-")
				createCaOptionsModel.Version = core.StringPtr("1.4.6-1")
				createCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.CreateCa(createCaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.CreateCaWithContext(ctx, createCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.CreateCa(createCaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.CreateCaWithContext(ctx, createCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateCa with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the ConfigCACors model
				configCaCorsModel := new(blockchainv3.ConfigCACors)
				configCaCorsModel.Enabled = core.BoolPtr(true)
				configCaCorsModel.Origins = []string{"*"}

				// Construct an instance of the ConfigCATlsClientauth model
				configCaTlsClientauthModel := new(blockchainv3.ConfigCATlsClientauth)
				configCaTlsClientauthModel.Type = core.StringPtr("noclientcert")
				configCaTlsClientauthModel.Certfiles = []string{"testString"}

				// Construct an instance of the ConfigCATls model
				configCaTlsModel := new(blockchainv3.ConfigCATls)
				configCaTlsModel.Keyfile = core.StringPtr("testString")
				configCaTlsModel.Certfile = core.StringPtr("testString")
				configCaTlsModel.Clientauth = configCaTlsClientauthModel

				// Construct an instance of the ConfigCACa model
				configCaCaModel := new(blockchainv3.ConfigCACa)
				configCaCaModel.Keyfile = core.StringPtr("testString")
				configCaCaModel.Certfile = core.StringPtr("testString")
				configCaCaModel.Chainfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCACrl model
				configCaCrlModel := new(blockchainv3.ConfigCACrl)
				configCaCrlModel.Expiry = core.StringPtr("24h")

				// Construct an instance of the IdentityAttrs model
				identityAttrsModel := new(blockchainv3.IdentityAttrs)
				identityAttrsModel.HfRegistrarRoles = core.StringPtr("*")
				identityAttrsModel.HfRegistrarDelegateRoles = core.StringPtr("*")
				identityAttrsModel.HfRevoker = core.BoolPtr(true)
				identityAttrsModel.HfIntermediateCA = core.BoolPtr(true)
				identityAttrsModel.HfGenCRL = core.BoolPtr(true)
				identityAttrsModel.HfRegistrarAttributes = core.StringPtr("*")
				identityAttrsModel.HfAffiliationMgr = core.BoolPtr(true)

				// Construct an instance of the ConfigCARegistryIdentitiesItem model
				configCaRegistryIdentitiesItemModel := new(blockchainv3.ConfigCARegistryIdentitiesItem)
				configCaRegistryIdentitiesItemModel.Name = core.StringPtr("admin")
				configCaRegistryIdentitiesItemModel.Pass = core.StringPtr("password")
				configCaRegistryIdentitiesItemModel.Type = core.StringPtr("client")
				configCaRegistryIdentitiesItemModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryIdentitiesItemModel.Affiliation = core.StringPtr("testString")
				configCaRegistryIdentitiesItemModel.Attrs = identityAttrsModel

				// Construct an instance of the ConfigCARegistry model
				configCaRegistryModel := new(blockchainv3.ConfigCARegistry)
				configCaRegistryModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryModel.Identities = []blockchainv3.ConfigCARegistryIdentitiesItem{*configCaRegistryIdentitiesItemModel}

				// Construct an instance of the ConfigCADbTlsClient model
				configCaDbTlsClientModel := new(blockchainv3.ConfigCADbTlsClient)
				configCaDbTlsClientModel.Certfile = core.StringPtr("testString")
				configCaDbTlsClientModel.Keyfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCADbTls model
				configCaDbTlsModel := new(blockchainv3.ConfigCADbTls)
				configCaDbTlsModel.Certfiles = []string{"testString"}
				configCaDbTlsModel.Client = configCaDbTlsClientModel
				configCaDbTlsModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the ConfigCADb model
				configCaDbModel := new(blockchainv3.ConfigCADb)
				configCaDbModel.Type = core.StringPtr("postgres")
				configCaDbModel.Datasource = core.StringPtr("host=fake.databases.appdomain.cloud port=31941 user=ibm_cloud password=password dbname=ibmclouddb sslmode=verify-full")
				configCaDbModel.Tls = configCaDbTlsModel

				// Construct an instance of the ConfigCAAffiliations model
				configCaAffiliationsModel := new(blockchainv3.ConfigCAAffiliations)
				configCaAffiliationsModel.Org1 = []string{"department1"}
				configCaAffiliationsModel.Org2 = []string{"department1"}
				configCaAffiliationsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ConfigCACsrKeyrequest model
				configCaCsrKeyrequestModel := new(blockchainv3.ConfigCACsrKeyrequest)
				configCaCsrKeyrequestModel.Algo = core.StringPtr("ecdsa")
				configCaCsrKeyrequestModel.Size = core.Float64Ptr(float64(256))

				// Construct an instance of the ConfigCACsrNamesItem model
				configCaCsrNamesItemModel := new(blockchainv3.ConfigCACsrNamesItem)
				configCaCsrNamesItemModel.C = core.StringPtr("US")
				configCaCsrNamesItemModel.ST = core.StringPtr("North Carolina")
				configCaCsrNamesItemModel.L = core.StringPtr("Raleigh")
				configCaCsrNamesItemModel.O = core.StringPtr("Hyperledger")
				configCaCsrNamesItemModel.OU = core.StringPtr("Fabric")

				// Construct an instance of the ConfigCACsrCa model
				configCaCsrCaModel := new(blockchainv3.ConfigCACsrCa)
				configCaCsrCaModel.Expiry = core.StringPtr("131400h")
				configCaCsrCaModel.Pathlength = core.Float64Ptr(float64(0))

				// Construct an instance of the ConfigCACsr model
				configCaCsrModel := new(blockchainv3.ConfigCACsr)
				configCaCsrModel.Cn = core.StringPtr("ca")
				configCaCsrModel.Keyrequest = configCaCsrKeyrequestModel
				configCaCsrModel.Names = []blockchainv3.ConfigCACsrNamesItem{*configCaCsrNamesItemModel}
				configCaCsrModel.Hosts = []string{"localhost"}
				configCaCsrModel.Ca = configCaCsrCaModel

				// Construct an instance of the ConfigCAIdemix model
				configCaIdemixModel := new(blockchainv3.ConfigCAIdemix)
				configCaIdemixModel.Rhpoolsize = core.Float64Ptr(float64(100))
				configCaIdemixModel.Nonceexpiration = core.StringPtr("15s")
				configCaIdemixModel.Noncesweepinterval = core.StringPtr("15m")

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv3.BccspSW)
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv3.BccspPKCS11)
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv3.Bccsp)
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model

				// Construct an instance of the ConfigCAIntermediateParentserver model
				configCaIntermediateParentserverModel := new(blockchainv3.ConfigCAIntermediateParentserver)
				configCaIntermediateParentserverModel.URL = core.StringPtr("testString")
				configCaIntermediateParentserverModel.Caname = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateEnrollment model
				configCaIntermediateEnrollmentModel := new(blockchainv3.ConfigCAIntermediateEnrollment)
				configCaIntermediateEnrollmentModel.Hosts = core.StringPtr("localhost")
				configCaIntermediateEnrollmentModel.Profile = core.StringPtr("testString")
				configCaIntermediateEnrollmentModel.Label = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateTlsClient model
				configCaIntermediateTlsClientModel := new(blockchainv3.ConfigCAIntermediateTlsClient)
				configCaIntermediateTlsClientModel.Certfile = core.StringPtr("testString")
				configCaIntermediateTlsClientModel.Keyfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateTls model
				configCaIntermediateTlsModel := new(blockchainv3.ConfigCAIntermediateTls)
				configCaIntermediateTlsModel.Certfiles = []string{"testString"}
				configCaIntermediateTlsModel.Client = configCaIntermediateTlsClientModel

				// Construct an instance of the ConfigCAIntermediate model
				configCaIntermediateModel := new(blockchainv3.ConfigCAIntermediate)
				configCaIntermediateModel.Parentserver = configCaIntermediateParentserverModel
				configCaIntermediateModel.Enrollment = configCaIntermediateEnrollmentModel
				configCaIntermediateModel.Tls = configCaIntermediateTlsModel

				// Construct an instance of the ConfigCACfgIdentities model
				configCaCfgIdentitiesModel := new(blockchainv3.ConfigCACfgIdentities)
				configCaCfgIdentitiesModel.Passwordattempts = core.Float64Ptr(float64(10))
				configCaCfgIdentitiesModel.Allowremove = core.BoolPtr(false)

				// Construct an instance of the ConfigCACfg model
				configCaCfgModel := new(blockchainv3.ConfigCACfg)
				configCaCfgModel.Identities = configCaCfgIdentitiesModel

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv3.MetricsStatsd)
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv3.Metrics)
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel

				// Construct an instance of the ConfigCASigningDefault model
				configCaSigningDefaultModel := new(blockchainv3.ConfigCASigningDefault)
				configCaSigningDefaultModel.Usage = []string{"cert sign"}
				configCaSigningDefaultModel.Expiry = core.StringPtr("8760h")

				// Construct an instance of the ConfigCASigningProfilesCaCaconstraint model
				configCaSigningProfilesCaCaconstraintModel := new(blockchainv3.ConfigCASigningProfilesCaCaconstraint)
				configCaSigningProfilesCaCaconstraintModel.Isca = core.BoolPtr(true)
				configCaSigningProfilesCaCaconstraintModel.Maxpathlen = core.Float64Ptr(float64(0))
				configCaSigningProfilesCaCaconstraintModel.Maxpathlenzero = core.BoolPtr(true)

				// Construct an instance of the ConfigCASigningProfilesCa model
				configCaSigningProfilesCaModel := new(blockchainv3.ConfigCASigningProfilesCa)
				configCaSigningProfilesCaModel.Usage = []string{"cert sign"}
				configCaSigningProfilesCaModel.Expiry = core.StringPtr("43800h")
				configCaSigningProfilesCaModel.Caconstraint = configCaSigningProfilesCaCaconstraintModel

				// Construct an instance of the ConfigCASigningProfilesTls model
				configCaSigningProfilesTlsModel := new(blockchainv3.ConfigCASigningProfilesTls)
				configCaSigningProfilesTlsModel.Usage = []string{"cert sign"}
				configCaSigningProfilesTlsModel.Expiry = core.StringPtr("43800h")

				// Construct an instance of the ConfigCASigningProfiles model
				configCaSigningProfilesModel := new(blockchainv3.ConfigCASigningProfiles)
				configCaSigningProfilesModel.Ca = configCaSigningProfilesCaModel
				configCaSigningProfilesModel.Tls = configCaSigningProfilesTlsModel

				// Construct an instance of the ConfigCASigning model
				configCaSigningModel := new(blockchainv3.ConfigCASigning)
				configCaSigningModel.Default = configCaSigningDefaultModel
				configCaSigningModel.Profiles = configCaSigningProfilesModel

				// Construct an instance of the ConfigCACreate model
				configCaCreateModel := new(blockchainv3.ConfigCACreate)
				configCaCreateModel.Cors = configCaCorsModel
				configCaCreateModel.Debug = core.BoolPtr(false)
				configCaCreateModel.Crlsizelimit = core.Float64Ptr(float64(512000))
				configCaCreateModel.Tls = configCaTlsModel
				configCaCreateModel.Ca = configCaCaModel
				configCaCreateModel.Crl = configCaCrlModel
				configCaCreateModel.Registry = configCaRegistryModel
				configCaCreateModel.Db = configCaDbModel
				configCaCreateModel.Affiliations = configCaAffiliationsModel
				configCaCreateModel.Csr = configCaCsrModel
				configCaCreateModel.Idemix = configCaIdemixModel
				configCaCreateModel.BCCSP = bccspModel
				configCaCreateModel.Intermediate = configCaIntermediateModel
				configCaCreateModel.Cfg = configCaCfgModel
				configCaCreateModel.Metrics = metricsModel
				configCaCreateModel.Signing = configCaSigningModel

				// Construct an instance of the CreateCaBodyConfigOverride model
				createCaBodyConfigOverrideModel := new(blockchainv3.CreateCaBodyConfigOverride)
				createCaBodyConfigOverrideModel.Ca = configCaCreateModel
				createCaBodyConfigOverrideModel.Tlsca = configCaCreateModel

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv3.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv3.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv3.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the CreateCaBodyResources model
				createCaBodyResourcesModel := new(blockchainv3.CreateCaBodyResources)
				createCaBodyResourcesModel.Ca = resourceObjectModel

				// Construct an instance of the StorageObject model
				storageObjectModel := new(blockchainv3.StorageObject)
				storageObjectModel.Size = core.StringPtr("4GiB")
				storageObjectModel.Class = core.StringPtr("default")

				// Construct an instance of the CreateCaBodyStorage model
				createCaBodyStorageModel := new(blockchainv3.CreateCaBodyStorage)
				createCaBodyStorageModel.Ca = storageObjectModel

				// Construct an instance of the Hsm model
				hsmModel := new(blockchainv3.Hsm)
				hsmModel.Pkcs11endpoint = core.StringPtr("tcp://example.com:666")

				// Construct an instance of the CreateCaOptions model
				createCaOptionsModel := new(blockchainv3.CreateCaOptions)
				createCaOptionsModel.DisplayName = core.StringPtr("My CA")
				createCaOptionsModel.ConfigOverride = createCaBodyConfigOverrideModel
				createCaOptionsModel.ID = core.StringPtr("component1")
				createCaOptionsModel.Resources = createCaBodyResourcesModel
				createCaOptionsModel.Storage = createCaBodyStorageModel
				createCaOptionsModel.Zone = core.StringPtr("-")
				createCaOptionsModel.Replicas = core.Float64Ptr(float64(1))
				createCaOptionsModel.Tags = []string{"fabric-ca"}
				createCaOptionsModel.Hsm = hsmModel
				createCaOptionsModel.Region = core.StringPtr("-")
				createCaOptionsModel.Version = core.StringPtr("1.4.6-1")
				createCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.CreateCa(createCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateCaOptions model with no property values
				createCaOptionsModelNew := new(blockchainv3.CreateCaOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.CreateCa(createCaOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ImportCa(importCaOptions *ImportCaOptions) - Operation response error`, func() {
		importCaPath := "/ak/api/v3/components/fabric-ca"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(importCaPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ImportCa with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the ImportCaBodyMspCa model
				importCaBodyMspCaModel := new(blockchainv3.ImportCaBodyMspCa)
				importCaBodyMspCaModel.Name = core.StringPtr("org1CA")
				importCaBodyMspCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the ImportCaBodyMspTlsca model
				importCaBodyMspTlscaModel := new(blockchainv3.ImportCaBodyMspTlsca)
				importCaBodyMspTlscaModel.Name = core.StringPtr("org1tlsCA")
				importCaBodyMspTlscaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the ImportCaBodyMspComponent model
				importCaBodyMspComponentModel := new(blockchainv3.ImportCaBodyMspComponent)
				importCaBodyMspComponentModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")

				// Construct an instance of the ImportCaBodyMsp model
				importCaBodyMspModel := new(blockchainv3.ImportCaBodyMsp)
				importCaBodyMspModel.Ca = importCaBodyMspCaModel
				importCaBodyMspModel.Tlsca = importCaBodyMspTlscaModel
				importCaBodyMspModel.Component = importCaBodyMspComponentModel

				// Construct an instance of the ImportCaOptions model
				importCaOptionsModel := new(blockchainv3.ImportCaOptions)
				importCaOptionsModel.DisplayName = core.StringPtr("Sample CA")
				importCaOptionsModel.ApiURL = core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:7054")
				importCaOptionsModel.Msp = importCaBodyMspModel
				importCaOptionsModel.ID = core.StringPtr("component1")
				importCaOptionsModel.Location = core.StringPtr("ibmcloud")
				importCaOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443")
				importCaOptionsModel.Tags = []string{"fabric-ca"}
				importCaOptionsModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				importCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.ImportCa(importCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.ImportCa(importCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ImportCa(importCaOptions *ImportCaOptions)`, func() {
		importCaPath := "/ak/api/v3/components/fabric-ca"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(importCaPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "component1", "dep_component_id": "admin", "display_name": "My CA", "api_url": "grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051", "operations_url": "https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443", "config_override": {"anyKey": "anyValue"}, "location": "ibmcloud", "msp": {"ca": {"name": "ca", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "tlsca": {"name": "tlsca", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "component": {"tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "ecert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "admin_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}}, "resources": {"ca": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "storage": {"ca": {"size": "4GiB", "class": "default"}}, "tags": ["fabric-ca"], "timestamp": 1537262855753, "version": "1.4.6-1", "zone": "-"}`)
				}))
			})
			It(`Invoke ImportCa successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.ImportCa(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ImportCaBodyMspCa model
				importCaBodyMspCaModel := new(blockchainv3.ImportCaBodyMspCa)
				importCaBodyMspCaModel.Name = core.StringPtr("org1CA")
				importCaBodyMspCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the ImportCaBodyMspTlsca model
				importCaBodyMspTlscaModel := new(blockchainv3.ImportCaBodyMspTlsca)
				importCaBodyMspTlscaModel.Name = core.StringPtr("org1tlsCA")
				importCaBodyMspTlscaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the ImportCaBodyMspComponent model
				importCaBodyMspComponentModel := new(blockchainv3.ImportCaBodyMspComponent)
				importCaBodyMspComponentModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")

				// Construct an instance of the ImportCaBodyMsp model
				importCaBodyMspModel := new(blockchainv3.ImportCaBodyMsp)
				importCaBodyMspModel.Ca = importCaBodyMspCaModel
				importCaBodyMspModel.Tlsca = importCaBodyMspTlscaModel
				importCaBodyMspModel.Component = importCaBodyMspComponentModel

				// Construct an instance of the ImportCaOptions model
				importCaOptionsModel := new(blockchainv3.ImportCaOptions)
				importCaOptionsModel.DisplayName = core.StringPtr("Sample CA")
				importCaOptionsModel.ApiURL = core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:7054")
				importCaOptionsModel.Msp = importCaBodyMspModel
				importCaOptionsModel.ID = core.StringPtr("component1")
				importCaOptionsModel.Location = core.StringPtr("ibmcloud")
				importCaOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443")
				importCaOptionsModel.Tags = []string{"fabric-ca"}
				importCaOptionsModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				importCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.ImportCa(importCaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.ImportCaWithContext(ctx, importCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.ImportCa(importCaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.ImportCaWithContext(ctx, importCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ImportCa with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the ImportCaBodyMspCa model
				importCaBodyMspCaModel := new(blockchainv3.ImportCaBodyMspCa)
				importCaBodyMspCaModel.Name = core.StringPtr("org1CA")
				importCaBodyMspCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the ImportCaBodyMspTlsca model
				importCaBodyMspTlscaModel := new(blockchainv3.ImportCaBodyMspTlsca)
				importCaBodyMspTlscaModel.Name = core.StringPtr("org1tlsCA")
				importCaBodyMspTlscaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the ImportCaBodyMspComponent model
				importCaBodyMspComponentModel := new(blockchainv3.ImportCaBodyMspComponent)
				importCaBodyMspComponentModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")

				// Construct an instance of the ImportCaBodyMsp model
				importCaBodyMspModel := new(blockchainv3.ImportCaBodyMsp)
				importCaBodyMspModel.Ca = importCaBodyMspCaModel
				importCaBodyMspModel.Tlsca = importCaBodyMspTlscaModel
				importCaBodyMspModel.Component = importCaBodyMspComponentModel

				// Construct an instance of the ImportCaOptions model
				importCaOptionsModel := new(blockchainv3.ImportCaOptions)
				importCaOptionsModel.DisplayName = core.StringPtr("Sample CA")
				importCaOptionsModel.ApiURL = core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:7054")
				importCaOptionsModel.Msp = importCaBodyMspModel
				importCaOptionsModel.ID = core.StringPtr("component1")
				importCaOptionsModel.Location = core.StringPtr("ibmcloud")
				importCaOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443")
				importCaOptionsModel.Tags = []string{"fabric-ca"}
				importCaOptionsModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				importCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.ImportCa(importCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ImportCaOptions model with no property values
				importCaOptionsModelNew := new(blockchainv3.ImportCaOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.ImportCa(importCaOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCa(updateCaOptions *UpdateCaOptions) - Operation response error`, func() {
		updateCaPath := "/ak/api/v3/kubernetes/components/fabric-ca/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCaPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateCa with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the ConfigCACors model
				configCaCorsModel := new(blockchainv3.ConfigCACors)
				configCaCorsModel.Enabled = core.BoolPtr(true)
				configCaCorsModel.Origins = []string{"*"}

				// Construct an instance of the ConfigCATlsClientauth model
				configCaTlsClientauthModel := new(blockchainv3.ConfigCATlsClientauth)
				configCaTlsClientauthModel.Type = core.StringPtr("noclientcert")
				configCaTlsClientauthModel.Certfiles = []string{"testString"}

				// Construct an instance of the ConfigCATls model
				configCaTlsModel := new(blockchainv3.ConfigCATls)
				configCaTlsModel.Keyfile = core.StringPtr("testString")
				configCaTlsModel.Certfile = core.StringPtr("testString")
				configCaTlsModel.Clientauth = configCaTlsClientauthModel

				// Construct an instance of the ConfigCACa model
				configCaCaModel := new(blockchainv3.ConfigCACa)
				configCaCaModel.Keyfile = core.StringPtr("testString")
				configCaCaModel.Certfile = core.StringPtr("testString")
				configCaCaModel.Chainfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCACrl model
				configCaCrlModel := new(blockchainv3.ConfigCACrl)
				configCaCrlModel.Expiry = core.StringPtr("24h")

				// Construct an instance of the IdentityAttrs model
				identityAttrsModel := new(blockchainv3.IdentityAttrs)
				identityAttrsModel.HfRegistrarRoles = core.StringPtr("*")
				identityAttrsModel.HfRegistrarDelegateRoles = core.StringPtr("*")
				identityAttrsModel.HfRevoker = core.BoolPtr(true)
				identityAttrsModel.HfIntermediateCA = core.BoolPtr(true)
				identityAttrsModel.HfGenCRL = core.BoolPtr(true)
				identityAttrsModel.HfRegistrarAttributes = core.StringPtr("*")
				identityAttrsModel.HfAffiliationMgr = core.BoolPtr(true)

				// Construct an instance of the ConfigCARegistryIdentitiesItem model
				configCaRegistryIdentitiesItemModel := new(blockchainv3.ConfigCARegistryIdentitiesItem)
				configCaRegistryIdentitiesItemModel.Name = core.StringPtr("admin")
				configCaRegistryIdentitiesItemModel.Pass = core.StringPtr("password")
				configCaRegistryIdentitiesItemModel.Type = core.StringPtr("client")
				configCaRegistryIdentitiesItemModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryIdentitiesItemModel.Affiliation = core.StringPtr("testString")
				configCaRegistryIdentitiesItemModel.Attrs = identityAttrsModel

				// Construct an instance of the ConfigCARegistry model
				configCaRegistryModel := new(blockchainv3.ConfigCARegistry)
				configCaRegistryModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryModel.Identities = []blockchainv3.ConfigCARegistryIdentitiesItem{*configCaRegistryIdentitiesItemModel}

				// Construct an instance of the ConfigCADbTlsClient model
				configCaDbTlsClientModel := new(blockchainv3.ConfigCADbTlsClient)
				configCaDbTlsClientModel.Certfile = core.StringPtr("testString")
				configCaDbTlsClientModel.Keyfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCADbTls model
				configCaDbTlsModel := new(blockchainv3.ConfigCADbTls)
				configCaDbTlsModel.Certfiles = []string{"testString"}
				configCaDbTlsModel.Client = configCaDbTlsClientModel
				configCaDbTlsModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the ConfigCADb model
				configCaDbModel := new(blockchainv3.ConfigCADb)
				configCaDbModel.Type = core.StringPtr("postgres")
				configCaDbModel.Datasource = core.StringPtr("host=fake.databases.appdomain.cloud port=31941 user=ibm_cloud password=password dbname=ibmclouddb sslmode=verify-full")
				configCaDbModel.Tls = configCaDbTlsModel

				// Construct an instance of the ConfigCAAffiliations model
				configCaAffiliationsModel := new(blockchainv3.ConfigCAAffiliations)
				configCaAffiliationsModel.Org1 = []string{"department1"}
				configCaAffiliationsModel.Org2 = []string{"department1"}
				configCaAffiliationsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ConfigCACsrKeyrequest model
				configCaCsrKeyrequestModel := new(blockchainv3.ConfigCACsrKeyrequest)
				configCaCsrKeyrequestModel.Algo = core.StringPtr("ecdsa")
				configCaCsrKeyrequestModel.Size = core.Float64Ptr(float64(256))

				// Construct an instance of the ConfigCACsrNamesItem model
				configCaCsrNamesItemModel := new(blockchainv3.ConfigCACsrNamesItem)
				configCaCsrNamesItemModel.C = core.StringPtr("US")
				configCaCsrNamesItemModel.ST = core.StringPtr("North Carolina")
				configCaCsrNamesItemModel.L = core.StringPtr("Raleigh")
				configCaCsrNamesItemModel.O = core.StringPtr("Hyperledger")
				configCaCsrNamesItemModel.OU = core.StringPtr("Fabric")

				// Construct an instance of the ConfigCACsrCa model
				configCaCsrCaModel := new(blockchainv3.ConfigCACsrCa)
				configCaCsrCaModel.Expiry = core.StringPtr("131400h")
				configCaCsrCaModel.Pathlength = core.Float64Ptr(float64(0))

				// Construct an instance of the ConfigCACsr model
				configCaCsrModel := new(blockchainv3.ConfigCACsr)
				configCaCsrModel.Cn = core.StringPtr("ca")
				configCaCsrModel.Keyrequest = configCaCsrKeyrequestModel
				configCaCsrModel.Names = []blockchainv3.ConfigCACsrNamesItem{*configCaCsrNamesItemModel}
				configCaCsrModel.Hosts = []string{"localhost"}
				configCaCsrModel.Ca = configCaCsrCaModel

				// Construct an instance of the ConfigCAIdemix model
				configCaIdemixModel := new(blockchainv3.ConfigCAIdemix)
				configCaIdemixModel.Rhpoolsize = core.Float64Ptr(float64(100))
				configCaIdemixModel.Nonceexpiration = core.StringPtr("15s")
				configCaIdemixModel.Noncesweepinterval = core.StringPtr("15m")

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv3.BccspSW)
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv3.BccspPKCS11)
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv3.Bccsp)
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model

				// Construct an instance of the ConfigCAIntermediateParentserver model
				configCaIntermediateParentserverModel := new(blockchainv3.ConfigCAIntermediateParentserver)
				configCaIntermediateParentserverModel.URL = core.StringPtr("testString")
				configCaIntermediateParentserverModel.Caname = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateEnrollment model
				configCaIntermediateEnrollmentModel := new(blockchainv3.ConfigCAIntermediateEnrollment)
				configCaIntermediateEnrollmentModel.Hosts = core.StringPtr("localhost")
				configCaIntermediateEnrollmentModel.Profile = core.StringPtr("testString")
				configCaIntermediateEnrollmentModel.Label = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateTlsClient model
				configCaIntermediateTlsClientModel := new(blockchainv3.ConfigCAIntermediateTlsClient)
				configCaIntermediateTlsClientModel.Certfile = core.StringPtr("testString")
				configCaIntermediateTlsClientModel.Keyfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateTls model
				configCaIntermediateTlsModel := new(blockchainv3.ConfigCAIntermediateTls)
				configCaIntermediateTlsModel.Certfiles = []string{"testString"}
				configCaIntermediateTlsModel.Client = configCaIntermediateTlsClientModel

				// Construct an instance of the ConfigCAIntermediate model
				configCaIntermediateModel := new(blockchainv3.ConfigCAIntermediate)
				configCaIntermediateModel.Parentserver = configCaIntermediateParentserverModel
				configCaIntermediateModel.Enrollment = configCaIntermediateEnrollmentModel
				configCaIntermediateModel.Tls = configCaIntermediateTlsModel

				// Construct an instance of the ConfigCACfgIdentities model
				configCaCfgIdentitiesModel := new(blockchainv3.ConfigCACfgIdentities)
				configCaCfgIdentitiesModel.Passwordattempts = core.Float64Ptr(float64(10))
				configCaCfgIdentitiesModel.Allowremove = core.BoolPtr(false)

				// Construct an instance of the ConfigCACfg model
				configCaCfgModel := new(blockchainv3.ConfigCACfg)
				configCaCfgModel.Identities = configCaCfgIdentitiesModel

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv3.MetricsStatsd)
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv3.Metrics)
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel

				// Construct an instance of the ConfigCAUpdate model
				configCaUpdateModel := new(blockchainv3.ConfigCAUpdate)
				configCaUpdateModel.Cors = configCaCorsModel
				configCaUpdateModel.Debug = core.BoolPtr(false)
				configCaUpdateModel.Crlsizelimit = core.Float64Ptr(float64(512000))
				configCaUpdateModel.Tls = configCaTlsModel
				configCaUpdateModel.Ca = configCaCaModel
				configCaUpdateModel.Crl = configCaCrlModel
				configCaUpdateModel.Registry = configCaRegistryModel
				configCaUpdateModel.Db = configCaDbModel
				configCaUpdateModel.Affiliations = configCaAffiliationsModel
				configCaUpdateModel.Csr = configCaCsrModel
				configCaUpdateModel.Idemix = configCaIdemixModel
				configCaUpdateModel.BCCSP = bccspModel
				configCaUpdateModel.Intermediate = configCaIntermediateModel
				configCaUpdateModel.Cfg = configCaCfgModel
				configCaUpdateModel.Metrics = metricsModel

				// Construct an instance of the UpdateCaBodyConfigOverride model
				updateCaBodyConfigOverrideModel := new(blockchainv3.UpdateCaBodyConfigOverride)
				updateCaBodyConfigOverrideModel.Ca = configCaUpdateModel

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv3.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv3.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv3.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the UpdateCaBodyResources model
				updateCaBodyResourcesModel := new(blockchainv3.UpdateCaBodyResources)
				updateCaBodyResourcesModel.Ca = resourceObjectModel

				// Construct an instance of the UpdateCaOptions model
				updateCaOptionsModel := new(blockchainv3.UpdateCaOptions)
				updateCaOptionsModel.ID = core.StringPtr("testString")
				updateCaOptionsModel.ConfigOverride = updateCaBodyConfigOverrideModel
				updateCaOptionsModel.Replicas = core.Float64Ptr(float64(1))
				updateCaOptionsModel.Resources = updateCaBodyResourcesModel
				updateCaOptionsModel.Version = core.StringPtr("1.4.6-1")
				updateCaOptionsModel.Zone = core.StringPtr("-")
				updateCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.UpdateCa(updateCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.UpdateCa(updateCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateCa(updateCaOptions *UpdateCaOptions)`, func() {
		updateCaPath := "/ak/api/v3/kubernetes/components/fabric-ca/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCaPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "component1", "dep_component_id": "admin", "display_name": "My CA", "api_url": "grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051", "operations_url": "https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443", "config_override": {"anyKey": "anyValue"}, "location": "ibmcloud", "msp": {"ca": {"name": "ca", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "tlsca": {"name": "tlsca", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "component": {"tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "ecert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "admin_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}}, "resources": {"ca": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "storage": {"ca": {"size": "4GiB", "class": "default"}}, "tags": ["fabric-ca"], "timestamp": 1537262855753, "version": "1.4.6-1", "zone": "-"}`)
				}))
			})
			It(`Invoke UpdateCa successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.UpdateCa(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ConfigCACors model
				configCaCorsModel := new(blockchainv3.ConfigCACors)
				configCaCorsModel.Enabled = core.BoolPtr(true)
				configCaCorsModel.Origins = []string{"*"}

				// Construct an instance of the ConfigCATlsClientauth model
				configCaTlsClientauthModel := new(blockchainv3.ConfigCATlsClientauth)
				configCaTlsClientauthModel.Type = core.StringPtr("noclientcert")
				configCaTlsClientauthModel.Certfiles = []string{"testString"}

				// Construct an instance of the ConfigCATls model
				configCaTlsModel := new(blockchainv3.ConfigCATls)
				configCaTlsModel.Keyfile = core.StringPtr("testString")
				configCaTlsModel.Certfile = core.StringPtr("testString")
				configCaTlsModel.Clientauth = configCaTlsClientauthModel

				// Construct an instance of the ConfigCACa model
				configCaCaModel := new(blockchainv3.ConfigCACa)
				configCaCaModel.Keyfile = core.StringPtr("testString")
				configCaCaModel.Certfile = core.StringPtr("testString")
				configCaCaModel.Chainfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCACrl model
				configCaCrlModel := new(blockchainv3.ConfigCACrl)
				configCaCrlModel.Expiry = core.StringPtr("24h")

				// Construct an instance of the IdentityAttrs model
				identityAttrsModel := new(blockchainv3.IdentityAttrs)
				identityAttrsModel.HfRegistrarRoles = core.StringPtr("*")
				identityAttrsModel.HfRegistrarDelegateRoles = core.StringPtr("*")
				identityAttrsModel.HfRevoker = core.BoolPtr(true)
				identityAttrsModel.HfIntermediateCA = core.BoolPtr(true)
				identityAttrsModel.HfGenCRL = core.BoolPtr(true)
				identityAttrsModel.HfRegistrarAttributes = core.StringPtr("*")
				identityAttrsModel.HfAffiliationMgr = core.BoolPtr(true)

				// Construct an instance of the ConfigCARegistryIdentitiesItem model
				configCaRegistryIdentitiesItemModel := new(blockchainv3.ConfigCARegistryIdentitiesItem)
				configCaRegistryIdentitiesItemModel.Name = core.StringPtr("admin")
				configCaRegistryIdentitiesItemModel.Pass = core.StringPtr("password")
				configCaRegistryIdentitiesItemModel.Type = core.StringPtr("client")
				configCaRegistryIdentitiesItemModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryIdentitiesItemModel.Affiliation = core.StringPtr("testString")
				configCaRegistryIdentitiesItemModel.Attrs = identityAttrsModel

				// Construct an instance of the ConfigCARegistry model
				configCaRegistryModel := new(blockchainv3.ConfigCARegistry)
				configCaRegistryModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryModel.Identities = []blockchainv3.ConfigCARegistryIdentitiesItem{*configCaRegistryIdentitiesItemModel}

				// Construct an instance of the ConfigCADbTlsClient model
				configCaDbTlsClientModel := new(blockchainv3.ConfigCADbTlsClient)
				configCaDbTlsClientModel.Certfile = core.StringPtr("testString")
				configCaDbTlsClientModel.Keyfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCADbTls model
				configCaDbTlsModel := new(blockchainv3.ConfigCADbTls)
				configCaDbTlsModel.Certfiles = []string{"testString"}
				configCaDbTlsModel.Client = configCaDbTlsClientModel
				configCaDbTlsModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the ConfigCADb model
				configCaDbModel := new(blockchainv3.ConfigCADb)
				configCaDbModel.Type = core.StringPtr("postgres")
				configCaDbModel.Datasource = core.StringPtr("host=fake.databases.appdomain.cloud port=31941 user=ibm_cloud password=password dbname=ibmclouddb sslmode=verify-full")
				configCaDbModel.Tls = configCaDbTlsModel

				// Construct an instance of the ConfigCAAffiliations model
				configCaAffiliationsModel := new(blockchainv3.ConfigCAAffiliations)
				configCaAffiliationsModel.Org1 = []string{"department1"}
				configCaAffiliationsModel.Org2 = []string{"department1"}
				configCaAffiliationsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ConfigCACsrKeyrequest model
				configCaCsrKeyrequestModel := new(blockchainv3.ConfigCACsrKeyrequest)
				configCaCsrKeyrequestModel.Algo = core.StringPtr("ecdsa")
				configCaCsrKeyrequestModel.Size = core.Float64Ptr(float64(256))

				// Construct an instance of the ConfigCACsrNamesItem model
				configCaCsrNamesItemModel := new(blockchainv3.ConfigCACsrNamesItem)
				configCaCsrNamesItemModel.C = core.StringPtr("US")
				configCaCsrNamesItemModel.ST = core.StringPtr("North Carolina")
				configCaCsrNamesItemModel.L = core.StringPtr("Raleigh")
				configCaCsrNamesItemModel.O = core.StringPtr("Hyperledger")
				configCaCsrNamesItemModel.OU = core.StringPtr("Fabric")

				// Construct an instance of the ConfigCACsrCa model
				configCaCsrCaModel := new(blockchainv3.ConfigCACsrCa)
				configCaCsrCaModel.Expiry = core.StringPtr("131400h")
				configCaCsrCaModel.Pathlength = core.Float64Ptr(float64(0))

				// Construct an instance of the ConfigCACsr model
				configCaCsrModel := new(blockchainv3.ConfigCACsr)
				configCaCsrModel.Cn = core.StringPtr("ca")
				configCaCsrModel.Keyrequest = configCaCsrKeyrequestModel
				configCaCsrModel.Names = []blockchainv3.ConfigCACsrNamesItem{*configCaCsrNamesItemModel}
				configCaCsrModel.Hosts = []string{"localhost"}
				configCaCsrModel.Ca = configCaCsrCaModel

				// Construct an instance of the ConfigCAIdemix model
				configCaIdemixModel := new(blockchainv3.ConfigCAIdemix)
				configCaIdemixModel.Rhpoolsize = core.Float64Ptr(float64(100))
				configCaIdemixModel.Nonceexpiration = core.StringPtr("15s")
				configCaIdemixModel.Noncesweepinterval = core.StringPtr("15m")

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv3.BccspSW)
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv3.BccspPKCS11)
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv3.Bccsp)
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model

				// Construct an instance of the ConfigCAIntermediateParentserver model
				configCaIntermediateParentserverModel := new(blockchainv3.ConfigCAIntermediateParentserver)
				configCaIntermediateParentserverModel.URL = core.StringPtr("testString")
				configCaIntermediateParentserverModel.Caname = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateEnrollment model
				configCaIntermediateEnrollmentModel := new(blockchainv3.ConfigCAIntermediateEnrollment)
				configCaIntermediateEnrollmentModel.Hosts = core.StringPtr("localhost")
				configCaIntermediateEnrollmentModel.Profile = core.StringPtr("testString")
				configCaIntermediateEnrollmentModel.Label = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateTlsClient model
				configCaIntermediateTlsClientModel := new(blockchainv3.ConfigCAIntermediateTlsClient)
				configCaIntermediateTlsClientModel.Certfile = core.StringPtr("testString")
				configCaIntermediateTlsClientModel.Keyfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateTls model
				configCaIntermediateTlsModel := new(blockchainv3.ConfigCAIntermediateTls)
				configCaIntermediateTlsModel.Certfiles = []string{"testString"}
				configCaIntermediateTlsModel.Client = configCaIntermediateTlsClientModel

				// Construct an instance of the ConfigCAIntermediate model
				configCaIntermediateModel := new(blockchainv3.ConfigCAIntermediate)
				configCaIntermediateModel.Parentserver = configCaIntermediateParentserverModel
				configCaIntermediateModel.Enrollment = configCaIntermediateEnrollmentModel
				configCaIntermediateModel.Tls = configCaIntermediateTlsModel

				// Construct an instance of the ConfigCACfgIdentities model
				configCaCfgIdentitiesModel := new(blockchainv3.ConfigCACfgIdentities)
				configCaCfgIdentitiesModel.Passwordattempts = core.Float64Ptr(float64(10))
				configCaCfgIdentitiesModel.Allowremove = core.BoolPtr(false)

				// Construct an instance of the ConfigCACfg model
				configCaCfgModel := new(blockchainv3.ConfigCACfg)
				configCaCfgModel.Identities = configCaCfgIdentitiesModel

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv3.MetricsStatsd)
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv3.Metrics)
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel

				// Construct an instance of the ConfigCAUpdate model
				configCaUpdateModel := new(blockchainv3.ConfigCAUpdate)
				configCaUpdateModel.Cors = configCaCorsModel
				configCaUpdateModel.Debug = core.BoolPtr(false)
				configCaUpdateModel.Crlsizelimit = core.Float64Ptr(float64(512000))
				configCaUpdateModel.Tls = configCaTlsModel
				configCaUpdateModel.Ca = configCaCaModel
				configCaUpdateModel.Crl = configCaCrlModel
				configCaUpdateModel.Registry = configCaRegistryModel
				configCaUpdateModel.Db = configCaDbModel
				configCaUpdateModel.Affiliations = configCaAffiliationsModel
				configCaUpdateModel.Csr = configCaCsrModel
				configCaUpdateModel.Idemix = configCaIdemixModel
				configCaUpdateModel.BCCSP = bccspModel
				configCaUpdateModel.Intermediate = configCaIntermediateModel
				configCaUpdateModel.Cfg = configCaCfgModel
				configCaUpdateModel.Metrics = metricsModel

				// Construct an instance of the UpdateCaBodyConfigOverride model
				updateCaBodyConfigOverrideModel := new(blockchainv3.UpdateCaBodyConfigOverride)
				updateCaBodyConfigOverrideModel.Ca = configCaUpdateModel

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv3.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv3.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv3.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the UpdateCaBodyResources model
				updateCaBodyResourcesModel := new(blockchainv3.UpdateCaBodyResources)
				updateCaBodyResourcesModel.Ca = resourceObjectModel

				// Construct an instance of the UpdateCaOptions model
				updateCaOptionsModel := new(blockchainv3.UpdateCaOptions)
				updateCaOptionsModel.ID = core.StringPtr("testString")
				updateCaOptionsModel.ConfigOverride = updateCaBodyConfigOverrideModel
				updateCaOptionsModel.Replicas = core.Float64Ptr(float64(1))
				updateCaOptionsModel.Resources = updateCaBodyResourcesModel
				updateCaOptionsModel.Version = core.StringPtr("1.4.6-1")
				updateCaOptionsModel.Zone = core.StringPtr("-")
				updateCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.UpdateCa(updateCaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.UpdateCaWithContext(ctx, updateCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.UpdateCa(updateCaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.UpdateCaWithContext(ctx, updateCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateCa with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the ConfigCACors model
				configCaCorsModel := new(blockchainv3.ConfigCACors)
				configCaCorsModel.Enabled = core.BoolPtr(true)
				configCaCorsModel.Origins = []string{"*"}

				// Construct an instance of the ConfigCATlsClientauth model
				configCaTlsClientauthModel := new(blockchainv3.ConfigCATlsClientauth)
				configCaTlsClientauthModel.Type = core.StringPtr("noclientcert")
				configCaTlsClientauthModel.Certfiles = []string{"testString"}

				// Construct an instance of the ConfigCATls model
				configCaTlsModel := new(blockchainv3.ConfigCATls)
				configCaTlsModel.Keyfile = core.StringPtr("testString")
				configCaTlsModel.Certfile = core.StringPtr("testString")
				configCaTlsModel.Clientauth = configCaTlsClientauthModel

				// Construct an instance of the ConfigCACa model
				configCaCaModel := new(blockchainv3.ConfigCACa)
				configCaCaModel.Keyfile = core.StringPtr("testString")
				configCaCaModel.Certfile = core.StringPtr("testString")
				configCaCaModel.Chainfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCACrl model
				configCaCrlModel := new(blockchainv3.ConfigCACrl)
				configCaCrlModel.Expiry = core.StringPtr("24h")

				// Construct an instance of the IdentityAttrs model
				identityAttrsModel := new(blockchainv3.IdentityAttrs)
				identityAttrsModel.HfRegistrarRoles = core.StringPtr("*")
				identityAttrsModel.HfRegistrarDelegateRoles = core.StringPtr("*")
				identityAttrsModel.HfRevoker = core.BoolPtr(true)
				identityAttrsModel.HfIntermediateCA = core.BoolPtr(true)
				identityAttrsModel.HfGenCRL = core.BoolPtr(true)
				identityAttrsModel.HfRegistrarAttributes = core.StringPtr("*")
				identityAttrsModel.HfAffiliationMgr = core.BoolPtr(true)

				// Construct an instance of the ConfigCARegistryIdentitiesItem model
				configCaRegistryIdentitiesItemModel := new(blockchainv3.ConfigCARegistryIdentitiesItem)
				configCaRegistryIdentitiesItemModel.Name = core.StringPtr("admin")
				configCaRegistryIdentitiesItemModel.Pass = core.StringPtr("password")
				configCaRegistryIdentitiesItemModel.Type = core.StringPtr("client")
				configCaRegistryIdentitiesItemModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryIdentitiesItemModel.Affiliation = core.StringPtr("testString")
				configCaRegistryIdentitiesItemModel.Attrs = identityAttrsModel

				// Construct an instance of the ConfigCARegistry model
				configCaRegistryModel := new(blockchainv3.ConfigCARegistry)
				configCaRegistryModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryModel.Identities = []blockchainv3.ConfigCARegistryIdentitiesItem{*configCaRegistryIdentitiesItemModel}

				// Construct an instance of the ConfigCADbTlsClient model
				configCaDbTlsClientModel := new(blockchainv3.ConfigCADbTlsClient)
				configCaDbTlsClientModel.Certfile = core.StringPtr("testString")
				configCaDbTlsClientModel.Keyfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCADbTls model
				configCaDbTlsModel := new(blockchainv3.ConfigCADbTls)
				configCaDbTlsModel.Certfiles = []string{"testString"}
				configCaDbTlsModel.Client = configCaDbTlsClientModel
				configCaDbTlsModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the ConfigCADb model
				configCaDbModel := new(blockchainv3.ConfigCADb)
				configCaDbModel.Type = core.StringPtr("postgres")
				configCaDbModel.Datasource = core.StringPtr("host=fake.databases.appdomain.cloud port=31941 user=ibm_cloud password=password dbname=ibmclouddb sslmode=verify-full")
				configCaDbModel.Tls = configCaDbTlsModel

				// Construct an instance of the ConfigCAAffiliations model
				configCaAffiliationsModel := new(blockchainv3.ConfigCAAffiliations)
				configCaAffiliationsModel.Org1 = []string{"department1"}
				configCaAffiliationsModel.Org2 = []string{"department1"}
				configCaAffiliationsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ConfigCACsrKeyrequest model
				configCaCsrKeyrequestModel := new(blockchainv3.ConfigCACsrKeyrequest)
				configCaCsrKeyrequestModel.Algo = core.StringPtr("ecdsa")
				configCaCsrKeyrequestModel.Size = core.Float64Ptr(float64(256))

				// Construct an instance of the ConfigCACsrNamesItem model
				configCaCsrNamesItemModel := new(blockchainv3.ConfigCACsrNamesItem)
				configCaCsrNamesItemModel.C = core.StringPtr("US")
				configCaCsrNamesItemModel.ST = core.StringPtr("North Carolina")
				configCaCsrNamesItemModel.L = core.StringPtr("Raleigh")
				configCaCsrNamesItemModel.O = core.StringPtr("Hyperledger")
				configCaCsrNamesItemModel.OU = core.StringPtr("Fabric")

				// Construct an instance of the ConfigCACsrCa model
				configCaCsrCaModel := new(blockchainv3.ConfigCACsrCa)
				configCaCsrCaModel.Expiry = core.StringPtr("131400h")
				configCaCsrCaModel.Pathlength = core.Float64Ptr(float64(0))

				// Construct an instance of the ConfigCACsr model
				configCaCsrModel := new(blockchainv3.ConfigCACsr)
				configCaCsrModel.Cn = core.StringPtr("ca")
				configCaCsrModel.Keyrequest = configCaCsrKeyrequestModel
				configCaCsrModel.Names = []blockchainv3.ConfigCACsrNamesItem{*configCaCsrNamesItemModel}
				configCaCsrModel.Hosts = []string{"localhost"}
				configCaCsrModel.Ca = configCaCsrCaModel

				// Construct an instance of the ConfigCAIdemix model
				configCaIdemixModel := new(blockchainv3.ConfigCAIdemix)
				configCaIdemixModel.Rhpoolsize = core.Float64Ptr(float64(100))
				configCaIdemixModel.Nonceexpiration = core.StringPtr("15s")
				configCaIdemixModel.Noncesweepinterval = core.StringPtr("15m")

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv3.BccspSW)
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv3.BccspPKCS11)
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv3.Bccsp)
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model

				// Construct an instance of the ConfigCAIntermediateParentserver model
				configCaIntermediateParentserverModel := new(blockchainv3.ConfigCAIntermediateParentserver)
				configCaIntermediateParentserverModel.URL = core.StringPtr("testString")
				configCaIntermediateParentserverModel.Caname = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateEnrollment model
				configCaIntermediateEnrollmentModel := new(blockchainv3.ConfigCAIntermediateEnrollment)
				configCaIntermediateEnrollmentModel.Hosts = core.StringPtr("localhost")
				configCaIntermediateEnrollmentModel.Profile = core.StringPtr("testString")
				configCaIntermediateEnrollmentModel.Label = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateTlsClient model
				configCaIntermediateTlsClientModel := new(blockchainv3.ConfigCAIntermediateTlsClient)
				configCaIntermediateTlsClientModel.Certfile = core.StringPtr("testString")
				configCaIntermediateTlsClientModel.Keyfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateTls model
				configCaIntermediateTlsModel := new(blockchainv3.ConfigCAIntermediateTls)
				configCaIntermediateTlsModel.Certfiles = []string{"testString"}
				configCaIntermediateTlsModel.Client = configCaIntermediateTlsClientModel

				// Construct an instance of the ConfigCAIntermediate model
				configCaIntermediateModel := new(blockchainv3.ConfigCAIntermediate)
				configCaIntermediateModel.Parentserver = configCaIntermediateParentserverModel
				configCaIntermediateModel.Enrollment = configCaIntermediateEnrollmentModel
				configCaIntermediateModel.Tls = configCaIntermediateTlsModel

				// Construct an instance of the ConfigCACfgIdentities model
				configCaCfgIdentitiesModel := new(blockchainv3.ConfigCACfgIdentities)
				configCaCfgIdentitiesModel.Passwordattempts = core.Float64Ptr(float64(10))
				configCaCfgIdentitiesModel.Allowremove = core.BoolPtr(false)

				// Construct an instance of the ConfigCACfg model
				configCaCfgModel := new(blockchainv3.ConfigCACfg)
				configCaCfgModel.Identities = configCaCfgIdentitiesModel

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv3.MetricsStatsd)
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv3.Metrics)
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel

				// Construct an instance of the ConfigCAUpdate model
				configCaUpdateModel := new(blockchainv3.ConfigCAUpdate)
				configCaUpdateModel.Cors = configCaCorsModel
				configCaUpdateModel.Debug = core.BoolPtr(false)
				configCaUpdateModel.Crlsizelimit = core.Float64Ptr(float64(512000))
				configCaUpdateModel.Tls = configCaTlsModel
				configCaUpdateModel.Ca = configCaCaModel
				configCaUpdateModel.Crl = configCaCrlModel
				configCaUpdateModel.Registry = configCaRegistryModel
				configCaUpdateModel.Db = configCaDbModel
				configCaUpdateModel.Affiliations = configCaAffiliationsModel
				configCaUpdateModel.Csr = configCaCsrModel
				configCaUpdateModel.Idemix = configCaIdemixModel
				configCaUpdateModel.BCCSP = bccspModel
				configCaUpdateModel.Intermediate = configCaIntermediateModel
				configCaUpdateModel.Cfg = configCaCfgModel
				configCaUpdateModel.Metrics = metricsModel

				// Construct an instance of the UpdateCaBodyConfigOverride model
				updateCaBodyConfigOverrideModel := new(blockchainv3.UpdateCaBodyConfigOverride)
				updateCaBodyConfigOverrideModel.Ca = configCaUpdateModel

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv3.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv3.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv3.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the UpdateCaBodyResources model
				updateCaBodyResourcesModel := new(blockchainv3.UpdateCaBodyResources)
				updateCaBodyResourcesModel.Ca = resourceObjectModel

				// Construct an instance of the UpdateCaOptions model
				updateCaOptionsModel := new(blockchainv3.UpdateCaOptions)
				updateCaOptionsModel.ID = core.StringPtr("testString")
				updateCaOptionsModel.ConfigOverride = updateCaBodyConfigOverrideModel
				updateCaOptionsModel.Replicas = core.Float64Ptr(float64(1))
				updateCaOptionsModel.Resources = updateCaBodyResourcesModel
				updateCaOptionsModel.Version = core.StringPtr("1.4.6-1")
				updateCaOptionsModel.Zone = core.StringPtr("-")
				updateCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.UpdateCa(updateCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateCaOptions model with no property values
				updateCaOptionsModelNew := new(blockchainv3.UpdateCaOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.UpdateCa(updateCaOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`EditCa(editCaOptions *EditCaOptions) - Operation response error`, func() {
		editCaPath := "/ak/api/v3/components/fabric-ca/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(editCaPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke EditCa with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the EditCaOptions model
				editCaOptionsModel := new(blockchainv3.EditCaOptions)
				editCaOptionsModel.ID = core.StringPtr("testString")
				editCaOptionsModel.DisplayName = core.StringPtr("My CA")
				editCaOptionsModel.ApiURL = core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:7054")
				editCaOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443")
				editCaOptionsModel.CaName = core.StringPtr("ca")
				editCaOptionsModel.Location = core.StringPtr("ibmcloud")
				editCaOptionsModel.Tags = []string{"fabric-ca"}
				editCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.EditCa(editCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.EditCa(editCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`EditCa(editCaOptions *EditCaOptions)`, func() {
		editCaPath := "/ak/api/v3/components/fabric-ca/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(editCaPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "component1", "dep_component_id": "admin", "display_name": "My CA", "api_url": "grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051", "operations_url": "https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443", "config_override": {"anyKey": "anyValue"}, "location": "ibmcloud", "msp": {"ca": {"name": "ca", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "tlsca": {"name": "tlsca", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "component": {"tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "ecert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "admin_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}}, "resources": {"ca": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "storage": {"ca": {"size": "4GiB", "class": "default"}}, "tags": ["fabric-ca"], "timestamp": 1537262855753, "version": "1.4.6-1", "zone": "-"}`)
				}))
			})
			It(`Invoke EditCa successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.EditCa(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EditCaOptions model
				editCaOptionsModel := new(blockchainv3.EditCaOptions)
				editCaOptionsModel.ID = core.StringPtr("testString")
				editCaOptionsModel.DisplayName = core.StringPtr("My CA")
				editCaOptionsModel.ApiURL = core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:7054")
				editCaOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443")
				editCaOptionsModel.CaName = core.StringPtr("ca")
				editCaOptionsModel.Location = core.StringPtr("ibmcloud")
				editCaOptionsModel.Tags = []string{"fabric-ca"}
				editCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.EditCa(editCaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.EditCaWithContext(ctx, editCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.EditCa(editCaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.EditCaWithContext(ctx, editCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke EditCa with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the EditCaOptions model
				editCaOptionsModel := new(blockchainv3.EditCaOptions)
				editCaOptionsModel.ID = core.StringPtr("testString")
				editCaOptionsModel.DisplayName = core.StringPtr("My CA")
				editCaOptionsModel.ApiURL = core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:7054")
				editCaOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443")
				editCaOptionsModel.CaName = core.StringPtr("ca")
				editCaOptionsModel.Location = core.StringPtr("ibmcloud")
				editCaOptionsModel.Tags = []string{"fabric-ca"}
				editCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.EditCa(editCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the EditCaOptions model with no property values
				editCaOptionsModelNew := new(blockchainv3.EditCaOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.EditCa(editCaOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CaAction(caActionOptions *CaActionOptions) - Operation response error`, func() {
		caActionPath := "/ak/api/v3/kubernetes/components/fabric-ca/testString/actions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(caActionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CaAction with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the ActionRenew model
				actionRenewModel := new(blockchainv3.ActionRenew)
				actionRenewModel.TlsCert = core.BoolPtr(true)

				// Construct an instance of the CaActionOptions model
				caActionOptionsModel := new(blockchainv3.CaActionOptions)
				caActionOptionsModel.ID = core.StringPtr("testString")
				caActionOptionsModel.Restart = core.BoolPtr(true)
				caActionOptionsModel.Renew = actionRenewModel
				caActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.CaAction(caActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.CaAction(caActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CaAction(caActionOptions *CaActionOptions)`, func() {
		caActionPath := "/ak/api/v3/kubernetes/components/fabric-ca/testString/actions"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(caActionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"message": "accepted", "id": "myca", "actions": ["restart"]}`)
				}))
			})
			It(`Invoke CaAction successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.CaAction(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ActionRenew model
				actionRenewModel := new(blockchainv3.ActionRenew)
				actionRenewModel.TlsCert = core.BoolPtr(true)

				// Construct an instance of the CaActionOptions model
				caActionOptionsModel := new(blockchainv3.CaActionOptions)
				caActionOptionsModel.ID = core.StringPtr("testString")
				caActionOptionsModel.Restart = core.BoolPtr(true)
				caActionOptionsModel.Renew = actionRenewModel
				caActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.CaAction(caActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.CaActionWithContext(ctx, caActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.CaAction(caActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.CaActionWithContext(ctx, caActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CaAction with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the ActionRenew model
				actionRenewModel := new(blockchainv3.ActionRenew)
				actionRenewModel.TlsCert = core.BoolPtr(true)

				// Construct an instance of the CaActionOptions model
				caActionOptionsModel := new(blockchainv3.CaActionOptions)
				caActionOptionsModel.ID = core.StringPtr("testString")
				caActionOptionsModel.Restart = core.BoolPtr(true)
				caActionOptionsModel.Renew = actionRenewModel
				caActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.CaAction(caActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CaActionOptions model with no property values
				caActionOptionsModelNew := new(blockchainv3.CaActionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.CaAction(caActionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreatePeer(createPeerOptions *CreatePeerOptions) - Operation response error`, func() {
		createPeerPath := "/ak/api/v3/kubernetes/components/fabric-peer"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createPeerPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreatePeer with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the CryptoEnrollmentComponent model
				cryptoEnrollmentComponentModel := new(blockchainv3.CryptoEnrollmentComponent)
				cryptoEnrollmentComponentModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the CryptoObjectEnrollmentCa model
				cryptoObjectEnrollmentCaModel := new(blockchainv3.CryptoObjectEnrollmentCa)
				cryptoObjectEnrollmentCaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				cryptoObjectEnrollmentCaModel.Port = core.Float64Ptr(float64(7054))
				cryptoObjectEnrollmentCaModel.Name = core.StringPtr("ca")
				cryptoObjectEnrollmentCaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				cryptoObjectEnrollmentCaModel.EnrollID = core.StringPtr("admin")
				cryptoObjectEnrollmentCaModel.EnrollSecret = core.StringPtr("password")

				// Construct an instance of the CryptoObjectEnrollmentTlsca model
				cryptoObjectEnrollmentTlscaModel := new(blockchainv3.CryptoObjectEnrollmentTlsca)
				cryptoObjectEnrollmentTlscaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				cryptoObjectEnrollmentTlscaModel.Port = core.Float64Ptr(float64(7054))
				cryptoObjectEnrollmentTlscaModel.Name = core.StringPtr("tlsca")
				cryptoObjectEnrollmentTlscaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				cryptoObjectEnrollmentTlscaModel.EnrollID = core.StringPtr("admin")
				cryptoObjectEnrollmentTlscaModel.EnrollSecret = core.StringPtr("password")
				cryptoObjectEnrollmentTlscaModel.CsrHosts = []string{"testString"}

				// Construct an instance of the CryptoObjectEnrollment model
				cryptoObjectEnrollmentModel := new(blockchainv3.CryptoObjectEnrollment)
				cryptoObjectEnrollmentModel.Component = cryptoEnrollmentComponentModel
				cryptoObjectEnrollmentModel.Ca = cryptoObjectEnrollmentCaModel
				cryptoObjectEnrollmentModel.Tlsca = cryptoObjectEnrollmentTlscaModel

				// Construct an instance of the ClientAuth model
				clientAuthModel := new(blockchainv3.ClientAuth)
				clientAuthModel.Type = core.StringPtr("noclientcert")
				clientAuthModel.TlsCerts = []string{"testString"}

				// Construct an instance of the MspCryptoComp model
				mspCryptoCompModel := new(blockchainv3.MspCryptoComp)
				mspCryptoCompModel.Ekey = core.StringPtr("testString")
				mspCryptoCompModel.Ecert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoCompModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				mspCryptoCompModel.TlsKey = core.StringPtr("testString")
				mspCryptoCompModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoCompModel.ClientAuth = clientAuthModel

				// Construct an instance of the MspCryptoCa model
				mspCryptoCaModel := new(blockchainv3.MspCryptoCa)
				mspCryptoCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				mspCryptoCaModel.CaIntermediateCerts = []string{"testString"}

				// Construct an instance of the CryptoObjectMsp model
				cryptoObjectMspModel := new(blockchainv3.CryptoObjectMsp)
				cryptoObjectMspModel.Component = mspCryptoCompModel
				cryptoObjectMspModel.Ca = mspCryptoCaModel
				cryptoObjectMspModel.Tlsca = mspCryptoCaModel

				// Construct an instance of the CryptoObject model
				cryptoObjectModel := new(blockchainv3.CryptoObject)
				cryptoObjectModel.Enrollment = cryptoObjectEnrollmentModel
				cryptoObjectModel.Msp = cryptoObjectMspModel

				// Construct an instance of the ConfigPeerKeepaliveClient model
				configPeerKeepaliveClientModel := new(blockchainv3.ConfigPeerKeepaliveClient)
				configPeerKeepaliveClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveClientModel.Timeout = core.StringPtr("20s")

				// Construct an instance of the ConfigPeerKeepaliveDeliveryClient model
				configPeerKeepaliveDeliveryClientModel := new(blockchainv3.ConfigPeerKeepaliveDeliveryClient)
				configPeerKeepaliveDeliveryClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveDeliveryClientModel.Timeout = core.StringPtr("20s")

				// Construct an instance of the ConfigPeerKeepalive model
				configPeerKeepaliveModel := new(blockchainv3.ConfigPeerKeepalive)
				configPeerKeepaliveModel.MinInterval = core.StringPtr("60s")
				configPeerKeepaliveModel.Client = configPeerKeepaliveClientModel
				configPeerKeepaliveModel.DeliveryClient = configPeerKeepaliveDeliveryClientModel

				// Construct an instance of the ConfigPeerGossipElection model
				configPeerGossipElectionModel := new(blockchainv3.ConfigPeerGossipElection)
				configPeerGossipElectionModel.StartupGracePeriod = core.StringPtr("15s")
				configPeerGossipElectionModel.MembershipSampleInterval = core.StringPtr("1s")
				configPeerGossipElectionModel.LeaderAliveThreshold = core.StringPtr("10s")
				configPeerGossipElectionModel.LeaderElectionDuration = core.StringPtr("5s")

				// Construct an instance of the ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy model
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel := new(blockchainv3.ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy)
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.RequiredPeerCount = core.Float64Ptr(float64(0))
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.MaxPeerCount = core.Float64Ptr(float64(1))

				// Construct an instance of the ConfigPeerGossipPvtData model
				configPeerGossipPvtDataModel := new(blockchainv3.ConfigPeerGossipPvtData)
				configPeerGossipPvtDataModel.PullRetryThreshold = core.StringPtr("60s")
				configPeerGossipPvtDataModel.TransientstoreMaxBlockRetention = core.Float64Ptr(float64(1000))
				configPeerGossipPvtDataModel.PushAckTimeout = core.StringPtr("3s")
				configPeerGossipPvtDataModel.BtlPullMargin = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileBatchSize = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileSleepInterval = core.StringPtr("1m")
				configPeerGossipPvtDataModel.ReconciliationEnabled = core.BoolPtr(true)
				configPeerGossipPvtDataModel.SkipPullingInvalidTransactionsDuringCommit = core.BoolPtr(false)
				configPeerGossipPvtDataModel.ImplicitCollectionDisseminationPolicy = configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel

				// Construct an instance of the ConfigPeerGossipState model
				configPeerGossipStateModel := new(blockchainv3.ConfigPeerGossipState)
				configPeerGossipStateModel.Enabled = core.BoolPtr(true)
				configPeerGossipStateModel.CheckInterval = core.StringPtr("10s")
				configPeerGossipStateModel.ResponseTimeout = core.StringPtr("3s")
				configPeerGossipStateModel.BatchSize = core.Float64Ptr(float64(10))
				configPeerGossipStateModel.BlockBufferSize = core.Float64Ptr(float64(100))
				configPeerGossipStateModel.MaxRetries = core.Float64Ptr(float64(3))

				// Construct an instance of the ConfigPeerGossip model
				configPeerGossipModel := new(blockchainv3.ConfigPeerGossip)
				configPeerGossipModel.UseLeaderElection = core.BoolPtr(true)
				configPeerGossipModel.OrgLeader = core.BoolPtr(false)
				configPeerGossipModel.MembershipTrackerInterval = core.StringPtr("5s")
				configPeerGossipModel.MaxBlockCountToStore = core.Float64Ptr(float64(100))
				configPeerGossipModel.MaxPropagationBurstLatency = core.StringPtr("10ms")
				configPeerGossipModel.MaxPropagationBurstSize = core.Float64Ptr(float64(10))
				configPeerGossipModel.PropagateIterations = core.Float64Ptr(float64(3))
				configPeerGossipModel.PullInterval = core.StringPtr("4s")
				configPeerGossipModel.PullPeerNum = core.Float64Ptr(float64(3))
				configPeerGossipModel.RequestStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.PublishStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.StateInfoRetentionInterval = core.StringPtr("0s")
				configPeerGossipModel.PublishCertPeriod = core.StringPtr("10s")
				configPeerGossipModel.SkipBlockVerification = core.BoolPtr(false)
				configPeerGossipModel.DialTimeout = core.StringPtr("3s")
				configPeerGossipModel.ConnTimeout = core.StringPtr("2s")
				configPeerGossipModel.RecvBuffSize = core.Float64Ptr(float64(20))
				configPeerGossipModel.SendBuffSize = core.Float64Ptr(float64(200))
				configPeerGossipModel.DigestWaitTime = core.StringPtr("1s")
				configPeerGossipModel.RequestWaitTime = core.StringPtr("1500ms")
				configPeerGossipModel.ResponseWaitTime = core.StringPtr("2s")
				configPeerGossipModel.AliveTimeInterval = core.StringPtr("5s")
				configPeerGossipModel.AliveExpirationTimeout = core.StringPtr("25s")
				configPeerGossipModel.ReconnectInterval = core.StringPtr("25s")
				configPeerGossipModel.Election = configPeerGossipElectionModel
				configPeerGossipModel.PvtData = configPeerGossipPvtDataModel
				configPeerGossipModel.State = configPeerGossipStateModel

				// Construct an instance of the ConfigPeerAuthentication model
				configPeerAuthenticationModel := new(blockchainv3.ConfigPeerAuthentication)
				configPeerAuthenticationModel.Timewindow = core.StringPtr("15m")

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv3.BccspSW)
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv3.BccspPKCS11)
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv3.Bccsp)
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model

				// Construct an instance of the ConfigPeerClient model
				configPeerClientModel := new(blockchainv3.ConfigPeerClient)
				configPeerClientModel.ConnTimeout = core.StringPtr("2s")

				// Construct an instance of the ConfigPeerDeliveryclientAddressOverridesItem model
				configPeerDeliveryclientAddressOverridesItemModel := new(blockchainv3.ConfigPeerDeliveryclientAddressOverridesItem)
				configPeerDeliveryclientAddressOverridesItemModel.From = core.StringPtr("n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.To = core.StringPtr("n3a3ec3-myorderer2.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.CaCertsFile = core.StringPtr("my-data/cert.pem")

				// Construct an instance of the ConfigPeerDeliveryclient model
				configPeerDeliveryclientModel := new(blockchainv3.ConfigPeerDeliveryclient)
				configPeerDeliveryclientModel.ReconnectTotalTimeThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.ConnTimeout = core.StringPtr("2s")
				configPeerDeliveryclientModel.ReConnectBackoffThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.AddressOverrides = []blockchainv3.ConfigPeerDeliveryclientAddressOverridesItem{*configPeerDeliveryclientAddressOverridesItemModel}

				// Construct an instance of the ConfigPeerAdminService model
				configPeerAdminServiceModel := new(blockchainv3.ConfigPeerAdminService)
				configPeerAdminServiceModel.ListenAddress = core.StringPtr("0.0.0.0:7051")

				// Construct an instance of the ConfigPeerDiscovery model
				configPeerDiscoveryModel := new(blockchainv3.ConfigPeerDiscovery)
				configPeerDiscoveryModel.Enabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheEnabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheMaxSize = core.Float64Ptr(float64(1000))
				configPeerDiscoveryModel.AuthCachePurgeRetentionRatio = core.Float64Ptr(float64(0.75))
				configPeerDiscoveryModel.OrgMembersAllowedAccess = core.BoolPtr(false)

				// Construct an instance of the ConfigPeerLimitsConcurrency model
				configPeerLimitsConcurrencyModel := new(blockchainv3.ConfigPeerLimitsConcurrency)
				configPeerLimitsConcurrencyModel.EndorserService = core.Float64Ptr(float64(2500))
				configPeerLimitsConcurrencyModel.DeliverService = core.Float64Ptr(float64(2500))

				// Construct an instance of the ConfigPeerLimits model
				configPeerLimitsModel := new(blockchainv3.ConfigPeerLimits)
				configPeerLimitsModel.Concurrency = configPeerLimitsConcurrencyModel

				// Construct an instance of the ConfigPeerGateway model
				configPeerGatewayModel := new(blockchainv3.ConfigPeerGateway)
				configPeerGatewayModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the ConfigPeerCreatePeer model
				configPeerCreatePeerModel := new(blockchainv3.ConfigPeerCreatePeer)
				configPeerCreatePeerModel.ID = core.StringPtr("john-doe")
				configPeerCreatePeerModel.NetworkID = core.StringPtr("dev")
				configPeerCreatePeerModel.Keepalive = configPeerKeepaliveModel
				configPeerCreatePeerModel.Gossip = configPeerGossipModel
				configPeerCreatePeerModel.Authentication = configPeerAuthenticationModel
				configPeerCreatePeerModel.BCCSP = bccspModel
				configPeerCreatePeerModel.Client = configPeerClientModel
				configPeerCreatePeerModel.Deliveryclient = configPeerDeliveryclientModel
				configPeerCreatePeerModel.AdminService = configPeerAdminServiceModel
				configPeerCreatePeerModel.ValidatorPoolSize = core.Float64Ptr(float64(8))
				configPeerCreatePeerModel.Discovery = configPeerDiscoveryModel
				configPeerCreatePeerModel.Limits = configPeerLimitsModel
				configPeerCreatePeerModel.Gateway = configPeerGatewayModel

				// Construct an instance of the ConfigPeerChaincodeGolang model
				configPeerChaincodeGolangModel := new(blockchainv3.ConfigPeerChaincodeGolang)
				configPeerChaincodeGolangModel.DynamicLink = core.BoolPtr(false)

				// Construct an instance of the ConfigPeerChaincodeExternalBuildersItem model
				configPeerChaincodeExternalBuildersItemModel := new(blockchainv3.ConfigPeerChaincodeExternalBuildersItem)
				configPeerChaincodeExternalBuildersItemModel.Path = core.StringPtr("/path/to/directory")
				configPeerChaincodeExternalBuildersItemModel.Name = core.StringPtr("descriptive-build-name")
				configPeerChaincodeExternalBuildersItemModel.EnvironmentWhitelist = []string{"GOPROXY"}

				// Construct an instance of the ConfigPeerChaincodeSystem model
				configPeerChaincodeSystemModel := new(blockchainv3.ConfigPeerChaincodeSystem)
				configPeerChaincodeSystemModel.Cscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Lscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Escc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Vscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Qscc = core.BoolPtr(true)

				// Construct an instance of the ConfigPeerChaincodeLogging model
				configPeerChaincodeLoggingModel := new(blockchainv3.ConfigPeerChaincodeLogging)
				configPeerChaincodeLoggingModel.Level = core.StringPtr("info")
				configPeerChaincodeLoggingModel.Shim = core.StringPtr("warning")
				configPeerChaincodeLoggingModel.Format = core.StringPtr("%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}")

				// Construct an instance of the ConfigPeerChaincode model
				configPeerChaincodeModel := new(blockchainv3.ConfigPeerChaincode)
				configPeerChaincodeModel.Golang = configPeerChaincodeGolangModel
				configPeerChaincodeModel.ExternalBuilders = []blockchainv3.ConfigPeerChaincodeExternalBuildersItem{*configPeerChaincodeExternalBuildersItemModel}
				configPeerChaincodeModel.InstallTimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Startuptimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Executetimeout = core.StringPtr("30s")
				configPeerChaincodeModel.System = configPeerChaincodeSystemModel
				configPeerChaincodeModel.Logging = configPeerChaincodeLoggingModel

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv3.MetricsStatsd)
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv3.Metrics)
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel

				// Construct an instance of the ConfigPeerCreate model
				configPeerCreateModel := new(blockchainv3.ConfigPeerCreate)
				configPeerCreateModel.Peer = configPeerCreatePeerModel
				configPeerCreateModel.Chaincode = configPeerChaincodeModel
				configPeerCreateModel.Metrics = metricsModel

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv3.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv3.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceObjectFabV2 model
				resourceObjectFabV2Model := new(blockchainv3.ResourceObjectFabV2)
				resourceObjectFabV2Model.Requests = resourceRequestsModel
				resourceObjectFabV2Model.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectCouchDb model
				resourceObjectCouchDbModel := new(blockchainv3.ResourceObjectCouchDb)
				resourceObjectCouchDbModel.Requests = resourceRequestsModel
				resourceObjectCouchDbModel.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv3.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectFabV1 model
				resourceObjectFabV1Model := new(blockchainv3.ResourceObjectFabV1)
				resourceObjectFabV1Model.Requests = resourceRequestsModel
				resourceObjectFabV1Model.Limits = resourceLimitsModel

				// Construct an instance of the PeerResources model
				peerResourcesModel := new(blockchainv3.PeerResources)
				peerResourcesModel.Chaincodelauncher = resourceObjectFabV2Model
				peerResourcesModel.Couchdb = resourceObjectCouchDbModel
				peerResourcesModel.Statedb = resourceObjectModel
				peerResourcesModel.Dind = resourceObjectFabV1Model
				peerResourcesModel.Fluentd = resourceObjectFabV1Model
				peerResourcesModel.Peer = resourceObjectModel
				peerResourcesModel.Proxy = resourceObjectModel

				// Construct an instance of the StorageObject model
				storageObjectModel := new(blockchainv3.StorageObject)
				storageObjectModel.Size = core.StringPtr("4GiB")
				storageObjectModel.Class = core.StringPtr("default")

				// Construct an instance of the CreatePeerBodyStorage model
				createPeerBodyStorageModel := new(blockchainv3.CreatePeerBodyStorage)
				createPeerBodyStorageModel.Peer = storageObjectModel
				createPeerBodyStorageModel.Statedb = storageObjectModel

				// Construct an instance of the Hsm model
				hsmModel := new(blockchainv3.Hsm)
				hsmModel.Pkcs11endpoint = core.StringPtr("tcp://example.com:666")

				// Construct an instance of the CreatePeerOptions model
				createPeerOptionsModel := new(blockchainv3.CreatePeerOptions)
				createPeerOptionsModel.MspID = core.StringPtr("Org1")
				createPeerOptionsModel.DisplayName = core.StringPtr("My Peer")
				createPeerOptionsModel.Crypto = cryptoObjectModel
				createPeerOptionsModel.ID = core.StringPtr("component1")
				createPeerOptionsModel.ConfigOverride = configPeerCreateModel
				createPeerOptionsModel.Resources = peerResourcesModel
				createPeerOptionsModel.Storage = createPeerBodyStorageModel
				createPeerOptionsModel.Zone = core.StringPtr("-")
				createPeerOptionsModel.StateDb = core.StringPtr("couchdb")
				createPeerOptionsModel.Tags = []string{"fabric-ca"}
				createPeerOptionsModel.Hsm = hsmModel
				createPeerOptionsModel.Region = core.StringPtr("-")
				createPeerOptionsModel.Version = core.StringPtr("1.4.6-1")
				createPeerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.CreatePeer(createPeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.CreatePeer(createPeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreatePeer(createPeerOptions *CreatePeerOptions)`, func() {
		createPeerPath := "/ak/api/v3/kubernetes/components/fabric-peer"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createPeerPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "component1", "dep_component_id": "admin", "api_url": "grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051", "display_name": "My Peer", "grpcwp_url": "https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084", "location": "ibmcloud", "operations_url": "https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443", "config_override": {"anyKey": "anyValue"}, "node_ou": {"enabled": true}, "msp": {"ca": {"name": "ca", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "tlsca": {"name": "tlsca", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "component": {"tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "ecert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "admin_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}}, "msp_id": "Org1", "resources": {"peer": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "proxy": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "statedb": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "state_db": "couchdb", "storage": {"peer": {"size": "4GiB", "class": "default"}, "statedb": {"size": "4GiB", "class": "default"}}, "tags": ["fabric-ca"], "timestamp": 1537262855753, "type": "fabric-peer", "version": "1.4.6-1", "zone": "-"}`)
				}))
			})
			It(`Invoke CreatePeer successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.CreatePeer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CryptoEnrollmentComponent model
				cryptoEnrollmentComponentModel := new(blockchainv3.CryptoEnrollmentComponent)
				cryptoEnrollmentComponentModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the CryptoObjectEnrollmentCa model
				cryptoObjectEnrollmentCaModel := new(blockchainv3.CryptoObjectEnrollmentCa)
				cryptoObjectEnrollmentCaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				cryptoObjectEnrollmentCaModel.Port = core.Float64Ptr(float64(7054))
				cryptoObjectEnrollmentCaModel.Name = core.StringPtr("ca")
				cryptoObjectEnrollmentCaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				cryptoObjectEnrollmentCaModel.EnrollID = core.StringPtr("admin")
				cryptoObjectEnrollmentCaModel.EnrollSecret = core.StringPtr("password")

				// Construct an instance of the CryptoObjectEnrollmentTlsca model
				cryptoObjectEnrollmentTlscaModel := new(blockchainv3.CryptoObjectEnrollmentTlsca)
				cryptoObjectEnrollmentTlscaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				cryptoObjectEnrollmentTlscaModel.Port = core.Float64Ptr(float64(7054))
				cryptoObjectEnrollmentTlscaModel.Name = core.StringPtr("tlsca")
				cryptoObjectEnrollmentTlscaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				cryptoObjectEnrollmentTlscaModel.EnrollID = core.StringPtr("admin")
				cryptoObjectEnrollmentTlscaModel.EnrollSecret = core.StringPtr("password")
				cryptoObjectEnrollmentTlscaModel.CsrHosts = []string{"testString"}

				// Construct an instance of the CryptoObjectEnrollment model
				cryptoObjectEnrollmentModel := new(blockchainv3.CryptoObjectEnrollment)
				cryptoObjectEnrollmentModel.Component = cryptoEnrollmentComponentModel
				cryptoObjectEnrollmentModel.Ca = cryptoObjectEnrollmentCaModel
				cryptoObjectEnrollmentModel.Tlsca = cryptoObjectEnrollmentTlscaModel

				// Construct an instance of the ClientAuth model
				clientAuthModel := new(blockchainv3.ClientAuth)
				clientAuthModel.Type = core.StringPtr("noclientcert")
				clientAuthModel.TlsCerts = []string{"testString"}

				// Construct an instance of the MspCryptoComp model
				mspCryptoCompModel := new(blockchainv3.MspCryptoComp)
				mspCryptoCompModel.Ekey = core.StringPtr("testString")
				mspCryptoCompModel.Ecert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoCompModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				mspCryptoCompModel.TlsKey = core.StringPtr("testString")
				mspCryptoCompModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoCompModel.ClientAuth = clientAuthModel

				// Construct an instance of the MspCryptoCa model
				mspCryptoCaModel := new(blockchainv3.MspCryptoCa)
				mspCryptoCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				mspCryptoCaModel.CaIntermediateCerts = []string{"testString"}

				// Construct an instance of the CryptoObjectMsp model
				cryptoObjectMspModel := new(blockchainv3.CryptoObjectMsp)
				cryptoObjectMspModel.Component = mspCryptoCompModel
				cryptoObjectMspModel.Ca = mspCryptoCaModel
				cryptoObjectMspModel.Tlsca = mspCryptoCaModel

				// Construct an instance of the CryptoObject model
				cryptoObjectModel := new(blockchainv3.CryptoObject)
				cryptoObjectModel.Enrollment = cryptoObjectEnrollmentModel
				cryptoObjectModel.Msp = cryptoObjectMspModel

				// Construct an instance of the ConfigPeerKeepaliveClient model
				configPeerKeepaliveClientModel := new(blockchainv3.ConfigPeerKeepaliveClient)
				configPeerKeepaliveClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveClientModel.Timeout = core.StringPtr("20s")

				// Construct an instance of the ConfigPeerKeepaliveDeliveryClient model
				configPeerKeepaliveDeliveryClientModel := new(blockchainv3.ConfigPeerKeepaliveDeliveryClient)
				configPeerKeepaliveDeliveryClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveDeliveryClientModel.Timeout = core.StringPtr("20s")

				// Construct an instance of the ConfigPeerKeepalive model
				configPeerKeepaliveModel := new(blockchainv3.ConfigPeerKeepalive)
				configPeerKeepaliveModel.MinInterval = core.StringPtr("60s")
				configPeerKeepaliveModel.Client = configPeerKeepaliveClientModel
				configPeerKeepaliveModel.DeliveryClient = configPeerKeepaliveDeliveryClientModel

				// Construct an instance of the ConfigPeerGossipElection model
				configPeerGossipElectionModel := new(blockchainv3.ConfigPeerGossipElection)
				configPeerGossipElectionModel.StartupGracePeriod = core.StringPtr("15s")
				configPeerGossipElectionModel.MembershipSampleInterval = core.StringPtr("1s")
				configPeerGossipElectionModel.LeaderAliveThreshold = core.StringPtr("10s")
				configPeerGossipElectionModel.LeaderElectionDuration = core.StringPtr("5s")

				// Construct an instance of the ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy model
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel := new(blockchainv3.ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy)
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.RequiredPeerCount = core.Float64Ptr(float64(0))
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.MaxPeerCount = core.Float64Ptr(float64(1))

				// Construct an instance of the ConfigPeerGossipPvtData model
				configPeerGossipPvtDataModel := new(blockchainv3.ConfigPeerGossipPvtData)
				configPeerGossipPvtDataModel.PullRetryThreshold = core.StringPtr("60s")
				configPeerGossipPvtDataModel.TransientstoreMaxBlockRetention = core.Float64Ptr(float64(1000))
				configPeerGossipPvtDataModel.PushAckTimeout = core.StringPtr("3s")
				configPeerGossipPvtDataModel.BtlPullMargin = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileBatchSize = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileSleepInterval = core.StringPtr("1m")
				configPeerGossipPvtDataModel.ReconciliationEnabled = core.BoolPtr(true)
				configPeerGossipPvtDataModel.SkipPullingInvalidTransactionsDuringCommit = core.BoolPtr(false)
				configPeerGossipPvtDataModel.ImplicitCollectionDisseminationPolicy = configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel

				// Construct an instance of the ConfigPeerGossipState model
				configPeerGossipStateModel := new(blockchainv3.ConfigPeerGossipState)
				configPeerGossipStateModel.Enabled = core.BoolPtr(true)
				configPeerGossipStateModel.CheckInterval = core.StringPtr("10s")
				configPeerGossipStateModel.ResponseTimeout = core.StringPtr("3s")
				configPeerGossipStateModel.BatchSize = core.Float64Ptr(float64(10))
				configPeerGossipStateModel.BlockBufferSize = core.Float64Ptr(float64(100))
				configPeerGossipStateModel.MaxRetries = core.Float64Ptr(float64(3))

				// Construct an instance of the ConfigPeerGossip model
				configPeerGossipModel := new(blockchainv3.ConfigPeerGossip)
				configPeerGossipModel.UseLeaderElection = core.BoolPtr(true)
				configPeerGossipModel.OrgLeader = core.BoolPtr(false)
				configPeerGossipModel.MembershipTrackerInterval = core.StringPtr("5s")
				configPeerGossipModel.MaxBlockCountToStore = core.Float64Ptr(float64(100))
				configPeerGossipModel.MaxPropagationBurstLatency = core.StringPtr("10ms")
				configPeerGossipModel.MaxPropagationBurstSize = core.Float64Ptr(float64(10))
				configPeerGossipModel.PropagateIterations = core.Float64Ptr(float64(3))
				configPeerGossipModel.PullInterval = core.StringPtr("4s")
				configPeerGossipModel.PullPeerNum = core.Float64Ptr(float64(3))
				configPeerGossipModel.RequestStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.PublishStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.StateInfoRetentionInterval = core.StringPtr("0s")
				configPeerGossipModel.PublishCertPeriod = core.StringPtr("10s")
				configPeerGossipModel.SkipBlockVerification = core.BoolPtr(false)
				configPeerGossipModel.DialTimeout = core.StringPtr("3s")
				configPeerGossipModel.ConnTimeout = core.StringPtr("2s")
				configPeerGossipModel.RecvBuffSize = core.Float64Ptr(float64(20))
				configPeerGossipModel.SendBuffSize = core.Float64Ptr(float64(200))
				configPeerGossipModel.DigestWaitTime = core.StringPtr("1s")
				configPeerGossipModel.RequestWaitTime = core.StringPtr("1500ms")
				configPeerGossipModel.ResponseWaitTime = core.StringPtr("2s")
				configPeerGossipModel.AliveTimeInterval = core.StringPtr("5s")
				configPeerGossipModel.AliveExpirationTimeout = core.StringPtr("25s")
				configPeerGossipModel.ReconnectInterval = core.StringPtr("25s")
				configPeerGossipModel.Election = configPeerGossipElectionModel
				configPeerGossipModel.PvtData = configPeerGossipPvtDataModel
				configPeerGossipModel.State = configPeerGossipStateModel

				// Construct an instance of the ConfigPeerAuthentication model
				configPeerAuthenticationModel := new(blockchainv3.ConfigPeerAuthentication)
				configPeerAuthenticationModel.Timewindow = core.StringPtr("15m")

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv3.BccspSW)
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv3.BccspPKCS11)
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv3.Bccsp)
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model

				// Construct an instance of the ConfigPeerClient model
				configPeerClientModel := new(blockchainv3.ConfigPeerClient)
				configPeerClientModel.ConnTimeout = core.StringPtr("2s")

				// Construct an instance of the ConfigPeerDeliveryclientAddressOverridesItem model
				configPeerDeliveryclientAddressOverridesItemModel := new(blockchainv3.ConfigPeerDeliveryclientAddressOverridesItem)
				configPeerDeliveryclientAddressOverridesItemModel.From = core.StringPtr("n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.To = core.StringPtr("n3a3ec3-myorderer2.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.CaCertsFile = core.StringPtr("my-data/cert.pem")

				// Construct an instance of the ConfigPeerDeliveryclient model
				configPeerDeliveryclientModel := new(blockchainv3.ConfigPeerDeliveryclient)
				configPeerDeliveryclientModel.ReconnectTotalTimeThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.ConnTimeout = core.StringPtr("2s")
				configPeerDeliveryclientModel.ReConnectBackoffThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.AddressOverrides = []blockchainv3.ConfigPeerDeliveryclientAddressOverridesItem{*configPeerDeliveryclientAddressOverridesItemModel}

				// Construct an instance of the ConfigPeerAdminService model
				configPeerAdminServiceModel := new(blockchainv3.ConfigPeerAdminService)
				configPeerAdminServiceModel.ListenAddress = core.StringPtr("0.0.0.0:7051")

				// Construct an instance of the ConfigPeerDiscovery model
				configPeerDiscoveryModel := new(blockchainv3.ConfigPeerDiscovery)
				configPeerDiscoveryModel.Enabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheEnabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheMaxSize = core.Float64Ptr(float64(1000))
				configPeerDiscoveryModel.AuthCachePurgeRetentionRatio = core.Float64Ptr(float64(0.75))
				configPeerDiscoveryModel.OrgMembersAllowedAccess = core.BoolPtr(false)

				// Construct an instance of the ConfigPeerLimitsConcurrency model
				configPeerLimitsConcurrencyModel := new(blockchainv3.ConfigPeerLimitsConcurrency)
				configPeerLimitsConcurrencyModel.EndorserService = core.Float64Ptr(float64(2500))
				configPeerLimitsConcurrencyModel.DeliverService = core.Float64Ptr(float64(2500))

				// Construct an instance of the ConfigPeerLimits model
				configPeerLimitsModel := new(blockchainv3.ConfigPeerLimits)
				configPeerLimitsModel.Concurrency = configPeerLimitsConcurrencyModel

				// Construct an instance of the ConfigPeerGateway model
				configPeerGatewayModel := new(blockchainv3.ConfigPeerGateway)
				configPeerGatewayModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the ConfigPeerCreatePeer model
				configPeerCreatePeerModel := new(blockchainv3.ConfigPeerCreatePeer)
				configPeerCreatePeerModel.ID = core.StringPtr("john-doe")
				configPeerCreatePeerModel.NetworkID = core.StringPtr("dev")
				configPeerCreatePeerModel.Keepalive = configPeerKeepaliveModel
				configPeerCreatePeerModel.Gossip = configPeerGossipModel
				configPeerCreatePeerModel.Authentication = configPeerAuthenticationModel
				configPeerCreatePeerModel.BCCSP = bccspModel
				configPeerCreatePeerModel.Client = configPeerClientModel
				configPeerCreatePeerModel.Deliveryclient = configPeerDeliveryclientModel
				configPeerCreatePeerModel.AdminService = configPeerAdminServiceModel
				configPeerCreatePeerModel.ValidatorPoolSize = core.Float64Ptr(float64(8))
				configPeerCreatePeerModel.Discovery = configPeerDiscoveryModel
				configPeerCreatePeerModel.Limits = configPeerLimitsModel
				configPeerCreatePeerModel.Gateway = configPeerGatewayModel

				// Construct an instance of the ConfigPeerChaincodeGolang model
				configPeerChaincodeGolangModel := new(blockchainv3.ConfigPeerChaincodeGolang)
				configPeerChaincodeGolangModel.DynamicLink = core.BoolPtr(false)

				// Construct an instance of the ConfigPeerChaincodeExternalBuildersItem model
				configPeerChaincodeExternalBuildersItemModel := new(blockchainv3.ConfigPeerChaincodeExternalBuildersItem)
				configPeerChaincodeExternalBuildersItemModel.Path = core.StringPtr("/path/to/directory")
				configPeerChaincodeExternalBuildersItemModel.Name = core.StringPtr("descriptive-build-name")
				configPeerChaincodeExternalBuildersItemModel.EnvironmentWhitelist = []string{"GOPROXY"}

				// Construct an instance of the ConfigPeerChaincodeSystem model
				configPeerChaincodeSystemModel := new(blockchainv3.ConfigPeerChaincodeSystem)
				configPeerChaincodeSystemModel.Cscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Lscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Escc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Vscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Qscc = core.BoolPtr(true)

				// Construct an instance of the ConfigPeerChaincodeLogging model
				configPeerChaincodeLoggingModel := new(blockchainv3.ConfigPeerChaincodeLogging)
				configPeerChaincodeLoggingModel.Level = core.StringPtr("info")
				configPeerChaincodeLoggingModel.Shim = core.StringPtr("warning")
				configPeerChaincodeLoggingModel.Format = core.StringPtr("%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}")

				// Construct an instance of the ConfigPeerChaincode model
				configPeerChaincodeModel := new(blockchainv3.ConfigPeerChaincode)
				configPeerChaincodeModel.Golang = configPeerChaincodeGolangModel
				configPeerChaincodeModel.ExternalBuilders = []blockchainv3.ConfigPeerChaincodeExternalBuildersItem{*configPeerChaincodeExternalBuildersItemModel}
				configPeerChaincodeModel.InstallTimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Startuptimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Executetimeout = core.StringPtr("30s")
				configPeerChaincodeModel.System = configPeerChaincodeSystemModel
				configPeerChaincodeModel.Logging = configPeerChaincodeLoggingModel

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv3.MetricsStatsd)
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv3.Metrics)
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel

				// Construct an instance of the ConfigPeerCreate model
				configPeerCreateModel := new(blockchainv3.ConfigPeerCreate)
				configPeerCreateModel.Peer = configPeerCreatePeerModel
				configPeerCreateModel.Chaincode = configPeerChaincodeModel
				configPeerCreateModel.Metrics = metricsModel

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv3.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv3.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceObjectFabV2 model
				resourceObjectFabV2Model := new(blockchainv3.ResourceObjectFabV2)
				resourceObjectFabV2Model.Requests = resourceRequestsModel
				resourceObjectFabV2Model.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectCouchDb model
				resourceObjectCouchDbModel := new(blockchainv3.ResourceObjectCouchDb)
				resourceObjectCouchDbModel.Requests = resourceRequestsModel
				resourceObjectCouchDbModel.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv3.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectFabV1 model
				resourceObjectFabV1Model := new(blockchainv3.ResourceObjectFabV1)
				resourceObjectFabV1Model.Requests = resourceRequestsModel
				resourceObjectFabV1Model.Limits = resourceLimitsModel

				// Construct an instance of the PeerResources model
				peerResourcesModel := new(blockchainv3.PeerResources)
				peerResourcesModel.Chaincodelauncher = resourceObjectFabV2Model
				peerResourcesModel.Couchdb = resourceObjectCouchDbModel
				peerResourcesModel.Statedb = resourceObjectModel
				peerResourcesModel.Dind = resourceObjectFabV1Model
				peerResourcesModel.Fluentd = resourceObjectFabV1Model
				peerResourcesModel.Peer = resourceObjectModel
				peerResourcesModel.Proxy = resourceObjectModel

				// Construct an instance of the StorageObject model
				storageObjectModel := new(blockchainv3.StorageObject)
				storageObjectModel.Size = core.StringPtr("4GiB")
				storageObjectModel.Class = core.StringPtr("default")

				// Construct an instance of the CreatePeerBodyStorage model
				createPeerBodyStorageModel := new(blockchainv3.CreatePeerBodyStorage)
				createPeerBodyStorageModel.Peer = storageObjectModel
				createPeerBodyStorageModel.Statedb = storageObjectModel

				// Construct an instance of the Hsm model
				hsmModel := new(blockchainv3.Hsm)
				hsmModel.Pkcs11endpoint = core.StringPtr("tcp://example.com:666")

				// Construct an instance of the CreatePeerOptions model
				createPeerOptionsModel := new(blockchainv3.CreatePeerOptions)
				createPeerOptionsModel.MspID = core.StringPtr("Org1")
				createPeerOptionsModel.DisplayName = core.StringPtr("My Peer")
				createPeerOptionsModel.Crypto = cryptoObjectModel
				createPeerOptionsModel.ID = core.StringPtr("component1")
				createPeerOptionsModel.ConfigOverride = configPeerCreateModel
				createPeerOptionsModel.Resources = peerResourcesModel
				createPeerOptionsModel.Storage = createPeerBodyStorageModel
				createPeerOptionsModel.Zone = core.StringPtr("-")
				createPeerOptionsModel.StateDb = core.StringPtr("couchdb")
				createPeerOptionsModel.Tags = []string{"fabric-ca"}
				createPeerOptionsModel.Hsm = hsmModel
				createPeerOptionsModel.Region = core.StringPtr("-")
				createPeerOptionsModel.Version = core.StringPtr("1.4.6-1")
				createPeerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.CreatePeer(createPeerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.CreatePeerWithContext(ctx, createPeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.CreatePeer(createPeerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.CreatePeerWithContext(ctx, createPeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreatePeer with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the CryptoEnrollmentComponent model
				cryptoEnrollmentComponentModel := new(blockchainv3.CryptoEnrollmentComponent)
				cryptoEnrollmentComponentModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the CryptoObjectEnrollmentCa model
				cryptoObjectEnrollmentCaModel := new(blockchainv3.CryptoObjectEnrollmentCa)
				cryptoObjectEnrollmentCaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				cryptoObjectEnrollmentCaModel.Port = core.Float64Ptr(float64(7054))
				cryptoObjectEnrollmentCaModel.Name = core.StringPtr("ca")
				cryptoObjectEnrollmentCaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				cryptoObjectEnrollmentCaModel.EnrollID = core.StringPtr("admin")
				cryptoObjectEnrollmentCaModel.EnrollSecret = core.StringPtr("password")

				// Construct an instance of the CryptoObjectEnrollmentTlsca model
				cryptoObjectEnrollmentTlscaModel := new(blockchainv3.CryptoObjectEnrollmentTlsca)
				cryptoObjectEnrollmentTlscaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				cryptoObjectEnrollmentTlscaModel.Port = core.Float64Ptr(float64(7054))
				cryptoObjectEnrollmentTlscaModel.Name = core.StringPtr("tlsca")
				cryptoObjectEnrollmentTlscaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				cryptoObjectEnrollmentTlscaModel.EnrollID = core.StringPtr("admin")
				cryptoObjectEnrollmentTlscaModel.EnrollSecret = core.StringPtr("password")
				cryptoObjectEnrollmentTlscaModel.CsrHosts = []string{"testString"}

				// Construct an instance of the CryptoObjectEnrollment model
				cryptoObjectEnrollmentModel := new(blockchainv3.CryptoObjectEnrollment)
				cryptoObjectEnrollmentModel.Component = cryptoEnrollmentComponentModel
				cryptoObjectEnrollmentModel.Ca = cryptoObjectEnrollmentCaModel
				cryptoObjectEnrollmentModel.Tlsca = cryptoObjectEnrollmentTlscaModel

				// Construct an instance of the ClientAuth model
				clientAuthModel := new(blockchainv3.ClientAuth)
				clientAuthModel.Type = core.StringPtr("noclientcert")
				clientAuthModel.TlsCerts = []string{"testString"}

				// Construct an instance of the MspCryptoComp model
				mspCryptoCompModel := new(blockchainv3.MspCryptoComp)
				mspCryptoCompModel.Ekey = core.StringPtr("testString")
				mspCryptoCompModel.Ecert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoCompModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				mspCryptoCompModel.TlsKey = core.StringPtr("testString")
				mspCryptoCompModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoCompModel.ClientAuth = clientAuthModel

				// Construct an instance of the MspCryptoCa model
				mspCryptoCaModel := new(blockchainv3.MspCryptoCa)
				mspCryptoCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				mspCryptoCaModel.CaIntermediateCerts = []string{"testString"}

				// Construct an instance of the CryptoObjectMsp model
				cryptoObjectMspModel := new(blockchainv3.CryptoObjectMsp)
				cryptoObjectMspModel.Component = mspCryptoCompModel
				cryptoObjectMspModel.Ca = mspCryptoCaModel
				cryptoObjectMspModel.Tlsca = mspCryptoCaModel

				// Construct an instance of the CryptoObject model
				cryptoObjectModel := new(blockchainv3.CryptoObject)
				cryptoObjectModel.Enrollment = cryptoObjectEnrollmentModel
				cryptoObjectModel.Msp = cryptoObjectMspModel

				// Construct an instance of the ConfigPeerKeepaliveClient model
				configPeerKeepaliveClientModel := new(blockchainv3.ConfigPeerKeepaliveClient)
				configPeerKeepaliveClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveClientModel.Timeout = core.StringPtr("20s")

				// Construct an instance of the ConfigPeerKeepaliveDeliveryClient model
				configPeerKeepaliveDeliveryClientModel := new(blockchainv3.ConfigPeerKeepaliveDeliveryClient)
				configPeerKeepaliveDeliveryClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveDeliveryClientModel.Timeout = core.StringPtr("20s")

				// Construct an instance of the ConfigPeerKeepalive model
				configPeerKeepaliveModel := new(blockchainv3.ConfigPeerKeepalive)
				configPeerKeepaliveModel.MinInterval = core.StringPtr("60s")
				configPeerKeepaliveModel.Client = configPeerKeepaliveClientModel
				configPeerKeepaliveModel.DeliveryClient = configPeerKeepaliveDeliveryClientModel

				// Construct an instance of the ConfigPeerGossipElection model
				configPeerGossipElectionModel := new(blockchainv3.ConfigPeerGossipElection)
				configPeerGossipElectionModel.StartupGracePeriod = core.StringPtr("15s")
				configPeerGossipElectionModel.MembershipSampleInterval = core.StringPtr("1s")
				configPeerGossipElectionModel.LeaderAliveThreshold = core.StringPtr("10s")
				configPeerGossipElectionModel.LeaderElectionDuration = core.StringPtr("5s")

				// Construct an instance of the ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy model
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel := new(blockchainv3.ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy)
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.RequiredPeerCount = core.Float64Ptr(float64(0))
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.MaxPeerCount = core.Float64Ptr(float64(1))

				// Construct an instance of the ConfigPeerGossipPvtData model
				configPeerGossipPvtDataModel := new(blockchainv3.ConfigPeerGossipPvtData)
				configPeerGossipPvtDataModel.PullRetryThreshold = core.StringPtr("60s")
				configPeerGossipPvtDataModel.TransientstoreMaxBlockRetention = core.Float64Ptr(float64(1000))
				configPeerGossipPvtDataModel.PushAckTimeout = core.StringPtr("3s")
				configPeerGossipPvtDataModel.BtlPullMargin = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileBatchSize = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileSleepInterval = core.StringPtr("1m")
				configPeerGossipPvtDataModel.ReconciliationEnabled = core.BoolPtr(true)
				configPeerGossipPvtDataModel.SkipPullingInvalidTransactionsDuringCommit = core.BoolPtr(false)
				configPeerGossipPvtDataModel.ImplicitCollectionDisseminationPolicy = configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel

				// Construct an instance of the ConfigPeerGossipState model
				configPeerGossipStateModel := new(blockchainv3.ConfigPeerGossipState)
				configPeerGossipStateModel.Enabled = core.BoolPtr(true)
				configPeerGossipStateModel.CheckInterval = core.StringPtr("10s")
				configPeerGossipStateModel.ResponseTimeout = core.StringPtr("3s")
				configPeerGossipStateModel.BatchSize = core.Float64Ptr(float64(10))
				configPeerGossipStateModel.BlockBufferSize = core.Float64Ptr(float64(100))
				configPeerGossipStateModel.MaxRetries = core.Float64Ptr(float64(3))

				// Construct an instance of the ConfigPeerGossip model
				configPeerGossipModel := new(blockchainv3.ConfigPeerGossip)
				configPeerGossipModel.UseLeaderElection = core.BoolPtr(true)
				configPeerGossipModel.OrgLeader = core.BoolPtr(false)
				configPeerGossipModel.MembershipTrackerInterval = core.StringPtr("5s")
				configPeerGossipModel.MaxBlockCountToStore = core.Float64Ptr(float64(100))
				configPeerGossipModel.MaxPropagationBurstLatency = core.StringPtr("10ms")
				configPeerGossipModel.MaxPropagationBurstSize = core.Float64Ptr(float64(10))
				configPeerGossipModel.PropagateIterations = core.Float64Ptr(float64(3))
				configPeerGossipModel.PullInterval = core.StringPtr("4s")
				configPeerGossipModel.PullPeerNum = core.Float64Ptr(float64(3))
				configPeerGossipModel.RequestStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.PublishStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.StateInfoRetentionInterval = core.StringPtr("0s")
				configPeerGossipModel.PublishCertPeriod = core.StringPtr("10s")
				configPeerGossipModel.SkipBlockVerification = core.BoolPtr(false)
				configPeerGossipModel.DialTimeout = core.StringPtr("3s")
				configPeerGossipModel.ConnTimeout = core.StringPtr("2s")
				configPeerGossipModel.RecvBuffSize = core.Float64Ptr(float64(20))
				configPeerGossipModel.SendBuffSize = core.Float64Ptr(float64(200))
				configPeerGossipModel.DigestWaitTime = core.StringPtr("1s")
				configPeerGossipModel.RequestWaitTime = core.StringPtr("1500ms")
				configPeerGossipModel.ResponseWaitTime = core.StringPtr("2s")
				configPeerGossipModel.AliveTimeInterval = core.StringPtr("5s")
				configPeerGossipModel.AliveExpirationTimeout = core.StringPtr("25s")
				configPeerGossipModel.ReconnectInterval = core.StringPtr("25s")
				configPeerGossipModel.Election = configPeerGossipElectionModel
				configPeerGossipModel.PvtData = configPeerGossipPvtDataModel
				configPeerGossipModel.State = configPeerGossipStateModel

				// Construct an instance of the ConfigPeerAuthentication model
				configPeerAuthenticationModel := new(blockchainv3.ConfigPeerAuthentication)
				configPeerAuthenticationModel.Timewindow = core.StringPtr("15m")

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv3.BccspSW)
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv3.BccspPKCS11)
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv3.Bccsp)
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model

				// Construct an instance of the ConfigPeerClient model
				configPeerClientModel := new(blockchainv3.ConfigPeerClient)
				configPeerClientModel.ConnTimeout = core.StringPtr("2s")

				// Construct an instance of the ConfigPeerDeliveryclientAddressOverridesItem model
				configPeerDeliveryclientAddressOverridesItemModel := new(blockchainv3.ConfigPeerDeliveryclientAddressOverridesItem)
				configPeerDeliveryclientAddressOverridesItemModel.From = core.StringPtr("n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.To = core.StringPtr("n3a3ec3-myorderer2.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.CaCertsFile = core.StringPtr("my-data/cert.pem")

				// Construct an instance of the ConfigPeerDeliveryclient model
				configPeerDeliveryclientModel := new(blockchainv3.ConfigPeerDeliveryclient)
				configPeerDeliveryclientModel.ReconnectTotalTimeThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.ConnTimeout = core.StringPtr("2s")
				configPeerDeliveryclientModel.ReConnectBackoffThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.AddressOverrides = []blockchainv3.ConfigPeerDeliveryclientAddressOverridesItem{*configPeerDeliveryclientAddressOverridesItemModel}

				// Construct an instance of the ConfigPeerAdminService model
				configPeerAdminServiceModel := new(blockchainv3.ConfigPeerAdminService)
				configPeerAdminServiceModel.ListenAddress = core.StringPtr("0.0.0.0:7051")

				// Construct an instance of the ConfigPeerDiscovery model
				configPeerDiscoveryModel := new(blockchainv3.ConfigPeerDiscovery)
				configPeerDiscoveryModel.Enabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheEnabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheMaxSize = core.Float64Ptr(float64(1000))
				configPeerDiscoveryModel.AuthCachePurgeRetentionRatio = core.Float64Ptr(float64(0.75))
				configPeerDiscoveryModel.OrgMembersAllowedAccess = core.BoolPtr(false)

				// Construct an instance of the ConfigPeerLimitsConcurrency model
				configPeerLimitsConcurrencyModel := new(blockchainv3.ConfigPeerLimitsConcurrency)
				configPeerLimitsConcurrencyModel.EndorserService = core.Float64Ptr(float64(2500))
				configPeerLimitsConcurrencyModel.DeliverService = core.Float64Ptr(float64(2500))

				// Construct an instance of the ConfigPeerLimits model
				configPeerLimitsModel := new(blockchainv3.ConfigPeerLimits)
				configPeerLimitsModel.Concurrency = configPeerLimitsConcurrencyModel

				// Construct an instance of the ConfigPeerGateway model
				configPeerGatewayModel := new(blockchainv3.ConfigPeerGateway)
				configPeerGatewayModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the ConfigPeerCreatePeer model
				configPeerCreatePeerModel := new(blockchainv3.ConfigPeerCreatePeer)
				configPeerCreatePeerModel.ID = core.StringPtr("john-doe")
				configPeerCreatePeerModel.NetworkID = core.StringPtr("dev")
				configPeerCreatePeerModel.Keepalive = configPeerKeepaliveModel
				configPeerCreatePeerModel.Gossip = configPeerGossipModel
				configPeerCreatePeerModel.Authentication = configPeerAuthenticationModel
				configPeerCreatePeerModel.BCCSP = bccspModel
				configPeerCreatePeerModel.Client = configPeerClientModel
				configPeerCreatePeerModel.Deliveryclient = configPeerDeliveryclientModel
				configPeerCreatePeerModel.AdminService = configPeerAdminServiceModel
				configPeerCreatePeerModel.ValidatorPoolSize = core.Float64Ptr(float64(8))
				configPeerCreatePeerModel.Discovery = configPeerDiscoveryModel
				configPeerCreatePeerModel.Limits = configPeerLimitsModel
				configPeerCreatePeerModel.Gateway = configPeerGatewayModel

				// Construct an instance of the ConfigPeerChaincodeGolang model
				configPeerChaincodeGolangModel := new(blockchainv3.ConfigPeerChaincodeGolang)
				configPeerChaincodeGolangModel.DynamicLink = core.BoolPtr(false)

				// Construct an instance of the ConfigPeerChaincodeExternalBuildersItem model
				configPeerChaincodeExternalBuildersItemModel := new(blockchainv3.ConfigPeerChaincodeExternalBuildersItem)
				configPeerChaincodeExternalBuildersItemModel.Path = core.StringPtr("/path/to/directory")
				configPeerChaincodeExternalBuildersItemModel.Name = core.StringPtr("descriptive-build-name")
				configPeerChaincodeExternalBuildersItemModel.EnvironmentWhitelist = []string{"GOPROXY"}

				// Construct an instance of the ConfigPeerChaincodeSystem model
				configPeerChaincodeSystemModel := new(blockchainv3.ConfigPeerChaincodeSystem)
				configPeerChaincodeSystemModel.Cscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Lscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Escc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Vscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Qscc = core.BoolPtr(true)

				// Construct an instance of the ConfigPeerChaincodeLogging model
				configPeerChaincodeLoggingModel := new(blockchainv3.ConfigPeerChaincodeLogging)
				configPeerChaincodeLoggingModel.Level = core.StringPtr("info")
				configPeerChaincodeLoggingModel.Shim = core.StringPtr("warning")
				configPeerChaincodeLoggingModel.Format = core.StringPtr("%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}")

				// Construct an instance of the ConfigPeerChaincode model
				configPeerChaincodeModel := new(blockchainv3.ConfigPeerChaincode)
				configPeerChaincodeModel.Golang = configPeerChaincodeGolangModel
				configPeerChaincodeModel.ExternalBuilders = []blockchainv3.ConfigPeerChaincodeExternalBuildersItem{*configPeerChaincodeExternalBuildersItemModel}
				configPeerChaincodeModel.InstallTimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Startuptimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Executetimeout = core.StringPtr("30s")
				configPeerChaincodeModel.System = configPeerChaincodeSystemModel
				configPeerChaincodeModel.Logging = configPeerChaincodeLoggingModel

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv3.MetricsStatsd)
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv3.Metrics)
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel

				// Construct an instance of the ConfigPeerCreate model
				configPeerCreateModel := new(blockchainv3.ConfigPeerCreate)
				configPeerCreateModel.Peer = configPeerCreatePeerModel
				configPeerCreateModel.Chaincode = configPeerChaincodeModel
				configPeerCreateModel.Metrics = metricsModel

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv3.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv3.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceObjectFabV2 model
				resourceObjectFabV2Model := new(blockchainv3.ResourceObjectFabV2)
				resourceObjectFabV2Model.Requests = resourceRequestsModel
				resourceObjectFabV2Model.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectCouchDb model
				resourceObjectCouchDbModel := new(blockchainv3.ResourceObjectCouchDb)
				resourceObjectCouchDbModel.Requests = resourceRequestsModel
				resourceObjectCouchDbModel.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv3.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectFabV1 model
				resourceObjectFabV1Model := new(blockchainv3.ResourceObjectFabV1)
				resourceObjectFabV1Model.Requests = resourceRequestsModel
				resourceObjectFabV1Model.Limits = resourceLimitsModel

				// Construct an instance of the PeerResources model
				peerResourcesModel := new(blockchainv3.PeerResources)
				peerResourcesModel.Chaincodelauncher = resourceObjectFabV2Model
				peerResourcesModel.Couchdb = resourceObjectCouchDbModel
				peerResourcesModel.Statedb = resourceObjectModel
				peerResourcesModel.Dind = resourceObjectFabV1Model
				peerResourcesModel.Fluentd = resourceObjectFabV1Model
				peerResourcesModel.Peer = resourceObjectModel
				peerResourcesModel.Proxy = resourceObjectModel

				// Construct an instance of the StorageObject model
				storageObjectModel := new(blockchainv3.StorageObject)
				storageObjectModel.Size = core.StringPtr("4GiB")
				storageObjectModel.Class = core.StringPtr("default")

				// Construct an instance of the CreatePeerBodyStorage model
				createPeerBodyStorageModel := new(blockchainv3.CreatePeerBodyStorage)
				createPeerBodyStorageModel.Peer = storageObjectModel
				createPeerBodyStorageModel.Statedb = storageObjectModel

				// Construct an instance of the Hsm model
				hsmModel := new(blockchainv3.Hsm)
				hsmModel.Pkcs11endpoint = core.StringPtr("tcp://example.com:666")

				// Construct an instance of the CreatePeerOptions model
				createPeerOptionsModel := new(blockchainv3.CreatePeerOptions)
				createPeerOptionsModel.MspID = core.StringPtr("Org1")
				createPeerOptionsModel.DisplayName = core.StringPtr("My Peer")
				createPeerOptionsModel.Crypto = cryptoObjectModel
				createPeerOptionsModel.ID = core.StringPtr("component1")
				createPeerOptionsModel.ConfigOverride = configPeerCreateModel
				createPeerOptionsModel.Resources = peerResourcesModel
				createPeerOptionsModel.Storage = createPeerBodyStorageModel
				createPeerOptionsModel.Zone = core.StringPtr("-")
				createPeerOptionsModel.StateDb = core.StringPtr("couchdb")
				createPeerOptionsModel.Tags = []string{"fabric-ca"}
				createPeerOptionsModel.Hsm = hsmModel
				createPeerOptionsModel.Region = core.StringPtr("-")
				createPeerOptionsModel.Version = core.StringPtr("1.4.6-1")
				createPeerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.CreatePeer(createPeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreatePeerOptions model with no property values
				createPeerOptionsModelNew := new(blockchainv3.CreatePeerOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.CreatePeer(createPeerOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ImportPeer(importPeerOptions *ImportPeerOptions) - Operation response error`, func() {
		importPeerPath := "/ak/api/v3/components/fabric-peer"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(importPeerPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ImportPeer with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the MspCryptoFieldCa model
				mspCryptoFieldCaModel := new(blockchainv3.MspCryptoFieldCa)
				mspCryptoFieldCaModel.Name = core.StringPtr("ca")
				mspCryptoFieldCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the MspCryptoFieldTlsca model
				mspCryptoFieldTlscaModel := new(blockchainv3.MspCryptoFieldTlsca)
				mspCryptoFieldTlscaModel.Name = core.StringPtr("tlsca")
				mspCryptoFieldTlscaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the MspCryptoFieldComponent model
				mspCryptoFieldComponentModel := new(blockchainv3.MspCryptoFieldComponent)
				mspCryptoFieldComponentModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoFieldComponentModel.Ecert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoFieldComponentModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the MspCryptoField model
				mspCryptoFieldModel := new(blockchainv3.MspCryptoField)
				mspCryptoFieldModel.Ca = mspCryptoFieldCaModel
				mspCryptoFieldModel.Tlsca = mspCryptoFieldTlscaModel
				mspCryptoFieldModel.Component = mspCryptoFieldComponentModel

				// Construct an instance of the ImportPeerOptions model
				importPeerOptionsModel := new(blockchainv3.ImportPeerOptions)
				importPeerOptionsModel.DisplayName = core.StringPtr("My Peer")
				importPeerOptionsModel.GrpcwpURL = core.StringPtr("https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084")
				importPeerOptionsModel.Msp = mspCryptoFieldModel
				importPeerOptionsModel.MspID = core.StringPtr("Org1")
				importPeerOptionsModel.ID = core.StringPtr("component1")
				importPeerOptionsModel.ApiURL = core.StringPtr("grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051")
				importPeerOptionsModel.Location = core.StringPtr("ibmcloud")
				importPeerOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443")
				importPeerOptionsModel.Tags = []string{"fabric-ca"}
				importPeerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.ImportPeer(importPeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.ImportPeer(importPeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ImportPeer(importPeerOptions *ImportPeerOptions)`, func() {
		importPeerPath := "/ak/api/v3/components/fabric-peer"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(importPeerPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "component1", "dep_component_id": "admin", "api_url": "grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051", "display_name": "My Peer", "grpcwp_url": "https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084", "location": "ibmcloud", "operations_url": "https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443", "config_override": {"anyKey": "anyValue"}, "node_ou": {"enabled": true}, "msp": {"ca": {"name": "ca", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "tlsca": {"name": "tlsca", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "component": {"tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "ecert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "admin_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}}, "msp_id": "Org1", "resources": {"peer": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "proxy": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "statedb": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "state_db": "couchdb", "storage": {"peer": {"size": "4GiB", "class": "default"}, "statedb": {"size": "4GiB", "class": "default"}}, "tags": ["fabric-ca"], "timestamp": 1537262855753, "type": "fabric-peer", "version": "1.4.6-1", "zone": "-"}`)
				}))
			})
			It(`Invoke ImportPeer successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.ImportPeer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the MspCryptoFieldCa model
				mspCryptoFieldCaModel := new(blockchainv3.MspCryptoFieldCa)
				mspCryptoFieldCaModel.Name = core.StringPtr("ca")
				mspCryptoFieldCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the MspCryptoFieldTlsca model
				mspCryptoFieldTlscaModel := new(blockchainv3.MspCryptoFieldTlsca)
				mspCryptoFieldTlscaModel.Name = core.StringPtr("tlsca")
				mspCryptoFieldTlscaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the MspCryptoFieldComponent model
				mspCryptoFieldComponentModel := new(blockchainv3.MspCryptoFieldComponent)
				mspCryptoFieldComponentModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoFieldComponentModel.Ecert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoFieldComponentModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the MspCryptoField model
				mspCryptoFieldModel := new(blockchainv3.MspCryptoField)
				mspCryptoFieldModel.Ca = mspCryptoFieldCaModel
				mspCryptoFieldModel.Tlsca = mspCryptoFieldTlscaModel
				mspCryptoFieldModel.Component = mspCryptoFieldComponentModel

				// Construct an instance of the ImportPeerOptions model
				importPeerOptionsModel := new(blockchainv3.ImportPeerOptions)
				importPeerOptionsModel.DisplayName = core.StringPtr("My Peer")
				importPeerOptionsModel.GrpcwpURL = core.StringPtr("https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084")
				importPeerOptionsModel.Msp = mspCryptoFieldModel
				importPeerOptionsModel.MspID = core.StringPtr("Org1")
				importPeerOptionsModel.ID = core.StringPtr("component1")
				importPeerOptionsModel.ApiURL = core.StringPtr("grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051")
				importPeerOptionsModel.Location = core.StringPtr("ibmcloud")
				importPeerOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443")
				importPeerOptionsModel.Tags = []string{"fabric-ca"}
				importPeerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.ImportPeer(importPeerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.ImportPeerWithContext(ctx, importPeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.ImportPeer(importPeerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.ImportPeerWithContext(ctx, importPeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ImportPeer with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the MspCryptoFieldCa model
				mspCryptoFieldCaModel := new(blockchainv3.MspCryptoFieldCa)
				mspCryptoFieldCaModel.Name = core.StringPtr("ca")
				mspCryptoFieldCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the MspCryptoFieldTlsca model
				mspCryptoFieldTlscaModel := new(blockchainv3.MspCryptoFieldTlsca)
				mspCryptoFieldTlscaModel.Name = core.StringPtr("tlsca")
				mspCryptoFieldTlscaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the MspCryptoFieldComponent model
				mspCryptoFieldComponentModel := new(blockchainv3.MspCryptoFieldComponent)
				mspCryptoFieldComponentModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoFieldComponentModel.Ecert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoFieldComponentModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the MspCryptoField model
				mspCryptoFieldModel := new(blockchainv3.MspCryptoField)
				mspCryptoFieldModel.Ca = mspCryptoFieldCaModel
				mspCryptoFieldModel.Tlsca = mspCryptoFieldTlscaModel
				mspCryptoFieldModel.Component = mspCryptoFieldComponentModel

				// Construct an instance of the ImportPeerOptions model
				importPeerOptionsModel := new(blockchainv3.ImportPeerOptions)
				importPeerOptionsModel.DisplayName = core.StringPtr("My Peer")
				importPeerOptionsModel.GrpcwpURL = core.StringPtr("https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084")
				importPeerOptionsModel.Msp = mspCryptoFieldModel
				importPeerOptionsModel.MspID = core.StringPtr("Org1")
				importPeerOptionsModel.ID = core.StringPtr("component1")
				importPeerOptionsModel.ApiURL = core.StringPtr("grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051")
				importPeerOptionsModel.Location = core.StringPtr("ibmcloud")
				importPeerOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443")
				importPeerOptionsModel.Tags = []string{"fabric-ca"}
				importPeerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.ImportPeer(importPeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ImportPeerOptions model with no property values
				importPeerOptionsModelNew := new(blockchainv3.ImportPeerOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.ImportPeer(importPeerOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`EditPeer(editPeerOptions *EditPeerOptions) - Operation response error`, func() {
		editPeerPath := "/ak/api/v3/components/fabric-peer/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(editPeerPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke EditPeer with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the EditPeerOptions model
				editPeerOptionsModel := new(blockchainv3.EditPeerOptions)
				editPeerOptionsModel.ID = core.StringPtr("testString")
				editPeerOptionsModel.DisplayName = core.StringPtr("My Peer")
				editPeerOptionsModel.ApiURL = core.StringPtr("grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051")
				editPeerOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443")
				editPeerOptionsModel.GrpcwpURL = core.StringPtr("https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084")
				editPeerOptionsModel.MspID = core.StringPtr("Org1")
				editPeerOptionsModel.Location = core.StringPtr("ibmcloud")
				editPeerOptionsModel.Tags = []string{"fabric-ca"}
				editPeerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.EditPeer(editPeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.EditPeer(editPeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`EditPeer(editPeerOptions *EditPeerOptions)`, func() {
		editPeerPath := "/ak/api/v3/components/fabric-peer/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(editPeerPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "component1", "dep_component_id": "admin", "api_url": "grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051", "display_name": "My Peer", "grpcwp_url": "https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084", "location": "ibmcloud", "operations_url": "https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443", "config_override": {"anyKey": "anyValue"}, "node_ou": {"enabled": true}, "msp": {"ca": {"name": "ca", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "tlsca": {"name": "tlsca", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "component": {"tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "ecert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "admin_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}}, "msp_id": "Org1", "resources": {"peer": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "proxy": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "statedb": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "state_db": "couchdb", "storage": {"peer": {"size": "4GiB", "class": "default"}, "statedb": {"size": "4GiB", "class": "default"}}, "tags": ["fabric-ca"], "timestamp": 1537262855753, "type": "fabric-peer", "version": "1.4.6-1", "zone": "-"}`)
				}))
			})
			It(`Invoke EditPeer successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.EditPeer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EditPeerOptions model
				editPeerOptionsModel := new(blockchainv3.EditPeerOptions)
				editPeerOptionsModel.ID = core.StringPtr("testString")
				editPeerOptionsModel.DisplayName = core.StringPtr("My Peer")
				editPeerOptionsModel.ApiURL = core.StringPtr("grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051")
				editPeerOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443")
				editPeerOptionsModel.GrpcwpURL = core.StringPtr("https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084")
				editPeerOptionsModel.MspID = core.StringPtr("Org1")
				editPeerOptionsModel.Location = core.StringPtr("ibmcloud")
				editPeerOptionsModel.Tags = []string{"fabric-ca"}
				editPeerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.EditPeer(editPeerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.EditPeerWithContext(ctx, editPeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.EditPeer(editPeerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.EditPeerWithContext(ctx, editPeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke EditPeer with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the EditPeerOptions model
				editPeerOptionsModel := new(blockchainv3.EditPeerOptions)
				editPeerOptionsModel.ID = core.StringPtr("testString")
				editPeerOptionsModel.DisplayName = core.StringPtr("My Peer")
				editPeerOptionsModel.ApiURL = core.StringPtr("grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051")
				editPeerOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443")
				editPeerOptionsModel.GrpcwpURL = core.StringPtr("https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084")
				editPeerOptionsModel.MspID = core.StringPtr("Org1")
				editPeerOptionsModel.Location = core.StringPtr("ibmcloud")
				editPeerOptionsModel.Tags = []string{"fabric-ca"}
				editPeerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.EditPeer(editPeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the EditPeerOptions model with no property values
				editPeerOptionsModelNew := new(blockchainv3.EditPeerOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.EditPeer(editPeerOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PeerAction(peerActionOptions *PeerActionOptions) - Operation response error`, func() {
		peerActionPath := "/ak/api/v3/kubernetes/components/fabric-peer/testString/actions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(peerActionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PeerAction with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the ActionReenroll model
				actionReenrollModel := new(blockchainv3.ActionReenroll)
				actionReenrollModel.TlsCert = core.BoolPtr(true)
				actionReenrollModel.Ecert = core.BoolPtr(true)

				// Construct an instance of the ActionEnroll model
				actionEnrollModel := new(blockchainv3.ActionEnroll)
				actionEnrollModel.TlsCert = core.BoolPtr(true)
				actionEnrollModel.Ecert = core.BoolPtr(true)

				// Construct an instance of the PeerActionOptions model
				peerActionOptionsModel := new(blockchainv3.PeerActionOptions)
				peerActionOptionsModel.ID = core.StringPtr("testString")
				peerActionOptionsModel.Restart = core.BoolPtr(true)
				peerActionOptionsModel.Reenroll = actionReenrollModel
				peerActionOptionsModel.Enroll = actionEnrollModel
				peerActionOptionsModel.UpgradeDbs = core.BoolPtr(true)
				peerActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.PeerAction(peerActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.PeerAction(peerActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PeerAction(peerActionOptions *PeerActionOptions)`, func() {
		peerActionPath := "/ak/api/v3/kubernetes/components/fabric-peer/testString/actions"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(peerActionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"message": "accepted", "id": "myca", "actions": ["restart"]}`)
				}))
			})
			It(`Invoke PeerAction successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.PeerAction(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ActionReenroll model
				actionReenrollModel := new(blockchainv3.ActionReenroll)
				actionReenrollModel.TlsCert = core.BoolPtr(true)
				actionReenrollModel.Ecert = core.BoolPtr(true)

				// Construct an instance of the ActionEnroll model
				actionEnrollModel := new(blockchainv3.ActionEnroll)
				actionEnrollModel.TlsCert = core.BoolPtr(true)
				actionEnrollModel.Ecert = core.BoolPtr(true)

				// Construct an instance of the PeerActionOptions model
				peerActionOptionsModel := new(blockchainv3.PeerActionOptions)
				peerActionOptionsModel.ID = core.StringPtr("testString")
				peerActionOptionsModel.Restart = core.BoolPtr(true)
				peerActionOptionsModel.Reenroll = actionReenrollModel
				peerActionOptionsModel.Enroll = actionEnrollModel
				peerActionOptionsModel.UpgradeDbs = core.BoolPtr(true)
				peerActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.PeerAction(peerActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.PeerActionWithContext(ctx, peerActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.PeerAction(peerActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.PeerActionWithContext(ctx, peerActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke PeerAction with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the ActionReenroll model
				actionReenrollModel := new(blockchainv3.ActionReenroll)
				actionReenrollModel.TlsCert = core.BoolPtr(true)
				actionReenrollModel.Ecert = core.BoolPtr(true)

				// Construct an instance of the ActionEnroll model
				actionEnrollModel := new(blockchainv3.ActionEnroll)
				actionEnrollModel.TlsCert = core.BoolPtr(true)
				actionEnrollModel.Ecert = core.BoolPtr(true)

				// Construct an instance of the PeerActionOptions model
				peerActionOptionsModel := new(blockchainv3.PeerActionOptions)
				peerActionOptionsModel.ID = core.StringPtr("testString")
				peerActionOptionsModel.Restart = core.BoolPtr(true)
				peerActionOptionsModel.Reenroll = actionReenrollModel
				peerActionOptionsModel.Enroll = actionEnrollModel
				peerActionOptionsModel.UpgradeDbs = core.BoolPtr(true)
				peerActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.PeerAction(peerActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PeerActionOptions model with no property values
				peerActionOptionsModelNew := new(blockchainv3.PeerActionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.PeerAction(peerActionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdatePeer(updatePeerOptions *UpdatePeerOptions) - Operation response error`, func() {
		updatePeerPath := "/ak/api/v3/kubernetes/components/fabric-peer/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePeerPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdatePeer with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the ConfigPeerKeepaliveClient model
				configPeerKeepaliveClientModel := new(blockchainv3.ConfigPeerKeepaliveClient)
				configPeerKeepaliveClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveClientModel.Timeout = core.StringPtr("20s")

				// Construct an instance of the ConfigPeerKeepaliveDeliveryClient model
				configPeerKeepaliveDeliveryClientModel := new(blockchainv3.ConfigPeerKeepaliveDeliveryClient)
				configPeerKeepaliveDeliveryClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveDeliveryClientModel.Timeout = core.StringPtr("20s")

				// Construct an instance of the ConfigPeerKeepalive model
				configPeerKeepaliveModel := new(blockchainv3.ConfigPeerKeepalive)
				configPeerKeepaliveModel.MinInterval = core.StringPtr("60s")
				configPeerKeepaliveModel.Client = configPeerKeepaliveClientModel
				configPeerKeepaliveModel.DeliveryClient = configPeerKeepaliveDeliveryClientModel

				// Construct an instance of the ConfigPeerGossipElection model
				configPeerGossipElectionModel := new(blockchainv3.ConfigPeerGossipElection)
				configPeerGossipElectionModel.StartupGracePeriod = core.StringPtr("15s")
				configPeerGossipElectionModel.MembershipSampleInterval = core.StringPtr("1s")
				configPeerGossipElectionModel.LeaderAliveThreshold = core.StringPtr("10s")
				configPeerGossipElectionModel.LeaderElectionDuration = core.StringPtr("5s")

				// Construct an instance of the ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy model
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel := new(blockchainv3.ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy)
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.RequiredPeerCount = core.Float64Ptr(float64(0))
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.MaxPeerCount = core.Float64Ptr(float64(1))

				// Construct an instance of the ConfigPeerGossipPvtData model
				configPeerGossipPvtDataModel := new(blockchainv3.ConfigPeerGossipPvtData)
				configPeerGossipPvtDataModel.PullRetryThreshold = core.StringPtr("60s")
				configPeerGossipPvtDataModel.TransientstoreMaxBlockRetention = core.Float64Ptr(float64(1000))
				configPeerGossipPvtDataModel.PushAckTimeout = core.StringPtr("3s")
				configPeerGossipPvtDataModel.BtlPullMargin = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileBatchSize = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileSleepInterval = core.StringPtr("1m")
				configPeerGossipPvtDataModel.ReconciliationEnabled = core.BoolPtr(true)
				configPeerGossipPvtDataModel.SkipPullingInvalidTransactionsDuringCommit = core.BoolPtr(false)
				configPeerGossipPvtDataModel.ImplicitCollectionDisseminationPolicy = configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel

				// Construct an instance of the ConfigPeerGossipState model
				configPeerGossipStateModel := new(blockchainv3.ConfigPeerGossipState)
				configPeerGossipStateModel.Enabled = core.BoolPtr(true)
				configPeerGossipStateModel.CheckInterval = core.StringPtr("10s")
				configPeerGossipStateModel.ResponseTimeout = core.StringPtr("3s")
				configPeerGossipStateModel.BatchSize = core.Float64Ptr(float64(10))
				configPeerGossipStateModel.BlockBufferSize = core.Float64Ptr(float64(100))
				configPeerGossipStateModel.MaxRetries = core.Float64Ptr(float64(3))

				// Construct an instance of the ConfigPeerGossip model
				configPeerGossipModel := new(blockchainv3.ConfigPeerGossip)
				configPeerGossipModel.UseLeaderElection = core.BoolPtr(true)
				configPeerGossipModel.OrgLeader = core.BoolPtr(false)
				configPeerGossipModel.MembershipTrackerInterval = core.StringPtr("5s")
				configPeerGossipModel.MaxBlockCountToStore = core.Float64Ptr(float64(100))
				configPeerGossipModel.MaxPropagationBurstLatency = core.StringPtr("10ms")
				configPeerGossipModel.MaxPropagationBurstSize = core.Float64Ptr(float64(10))
				configPeerGossipModel.PropagateIterations = core.Float64Ptr(float64(3))
				configPeerGossipModel.PullInterval = core.StringPtr("4s")
				configPeerGossipModel.PullPeerNum = core.Float64Ptr(float64(3))
				configPeerGossipModel.RequestStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.PublishStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.StateInfoRetentionInterval = core.StringPtr("0s")
				configPeerGossipModel.PublishCertPeriod = core.StringPtr("10s")
				configPeerGossipModel.SkipBlockVerification = core.BoolPtr(false)
				configPeerGossipModel.DialTimeout = core.StringPtr("3s")
				configPeerGossipModel.ConnTimeout = core.StringPtr("2s")
				configPeerGossipModel.RecvBuffSize = core.Float64Ptr(float64(20))
				configPeerGossipModel.SendBuffSize = core.Float64Ptr(float64(200))
				configPeerGossipModel.DigestWaitTime = core.StringPtr("1s")
				configPeerGossipModel.RequestWaitTime = core.StringPtr("1500ms")
				configPeerGossipModel.ResponseWaitTime = core.StringPtr("2s")
				configPeerGossipModel.AliveTimeInterval = core.StringPtr("5s")
				configPeerGossipModel.AliveExpirationTimeout = core.StringPtr("25s")
				configPeerGossipModel.ReconnectInterval = core.StringPtr("25s")
				configPeerGossipModel.Election = configPeerGossipElectionModel
				configPeerGossipModel.PvtData = configPeerGossipPvtDataModel
				configPeerGossipModel.State = configPeerGossipStateModel

				// Construct an instance of the ConfigPeerAuthentication model
				configPeerAuthenticationModel := new(blockchainv3.ConfigPeerAuthentication)
				configPeerAuthenticationModel.Timewindow = core.StringPtr("15m")

				// Construct an instance of the ConfigPeerClient model
				configPeerClientModel := new(blockchainv3.ConfigPeerClient)
				configPeerClientModel.ConnTimeout = core.StringPtr("2s")

				// Construct an instance of the ConfigPeerDeliveryclientAddressOverridesItem model
				configPeerDeliveryclientAddressOverridesItemModel := new(blockchainv3.ConfigPeerDeliveryclientAddressOverridesItem)
				configPeerDeliveryclientAddressOverridesItemModel.From = core.StringPtr("n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.To = core.StringPtr("n3a3ec3-myorderer2.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.CaCertsFile = core.StringPtr("my-data/cert.pem")

				// Construct an instance of the ConfigPeerDeliveryclient model
				configPeerDeliveryclientModel := new(blockchainv3.ConfigPeerDeliveryclient)
				configPeerDeliveryclientModel.ReconnectTotalTimeThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.ConnTimeout = core.StringPtr("2s")
				configPeerDeliveryclientModel.ReConnectBackoffThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.AddressOverrides = []blockchainv3.ConfigPeerDeliveryclientAddressOverridesItem{*configPeerDeliveryclientAddressOverridesItemModel}

				// Construct an instance of the ConfigPeerAdminService model
				configPeerAdminServiceModel := new(blockchainv3.ConfigPeerAdminService)
				configPeerAdminServiceModel.ListenAddress = core.StringPtr("0.0.0.0:7051")

				// Construct an instance of the ConfigPeerDiscovery model
				configPeerDiscoveryModel := new(blockchainv3.ConfigPeerDiscovery)
				configPeerDiscoveryModel.Enabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheEnabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheMaxSize = core.Float64Ptr(float64(1000))
				configPeerDiscoveryModel.AuthCachePurgeRetentionRatio = core.Float64Ptr(float64(0.75))
				configPeerDiscoveryModel.OrgMembersAllowedAccess = core.BoolPtr(false)

				// Construct an instance of the ConfigPeerLimitsConcurrency model
				configPeerLimitsConcurrencyModel := new(blockchainv3.ConfigPeerLimitsConcurrency)
				configPeerLimitsConcurrencyModel.EndorserService = core.Float64Ptr(float64(2500))
				configPeerLimitsConcurrencyModel.DeliverService = core.Float64Ptr(float64(2500))

				// Construct an instance of the ConfigPeerLimits model
				configPeerLimitsModel := new(blockchainv3.ConfigPeerLimits)
				configPeerLimitsModel.Concurrency = configPeerLimitsConcurrencyModel

				// Construct an instance of the ConfigPeerGateway model
				configPeerGatewayModel := new(blockchainv3.ConfigPeerGateway)
				configPeerGatewayModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the ConfigPeerUpdatePeer model
				configPeerUpdatePeerModel := new(blockchainv3.ConfigPeerUpdatePeer)
				configPeerUpdatePeerModel.ID = core.StringPtr("john-doe")
				configPeerUpdatePeerModel.NetworkID = core.StringPtr("dev")
				configPeerUpdatePeerModel.Keepalive = configPeerKeepaliveModel
				configPeerUpdatePeerModel.Gossip = configPeerGossipModel
				configPeerUpdatePeerModel.Authentication = configPeerAuthenticationModel
				configPeerUpdatePeerModel.Client = configPeerClientModel
				configPeerUpdatePeerModel.Deliveryclient = configPeerDeliveryclientModel
				configPeerUpdatePeerModel.AdminService = configPeerAdminServiceModel
				configPeerUpdatePeerModel.ValidatorPoolSize = core.Float64Ptr(float64(8))
				configPeerUpdatePeerModel.Discovery = configPeerDiscoveryModel
				configPeerUpdatePeerModel.Limits = configPeerLimitsModel
				configPeerUpdatePeerModel.Gateway = configPeerGatewayModel

				// Construct an instance of the ConfigPeerChaincodeGolang model
				configPeerChaincodeGolangModel := new(blockchainv3.ConfigPeerChaincodeGolang)
				configPeerChaincodeGolangModel.DynamicLink = core.BoolPtr(false)

				// Construct an instance of the ConfigPeerChaincodeExternalBuildersItem model
				configPeerChaincodeExternalBuildersItemModel := new(blockchainv3.ConfigPeerChaincodeExternalBuildersItem)
				configPeerChaincodeExternalBuildersItemModel.Path = core.StringPtr("/path/to/directory")
				configPeerChaincodeExternalBuildersItemModel.Name = core.StringPtr("descriptive-build-name")
				configPeerChaincodeExternalBuildersItemModel.EnvironmentWhitelist = []string{"GOPROXY"}

				// Construct an instance of the ConfigPeerChaincodeSystem model
				configPeerChaincodeSystemModel := new(blockchainv3.ConfigPeerChaincodeSystem)
				configPeerChaincodeSystemModel.Cscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Lscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Escc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Vscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Qscc = core.BoolPtr(true)

				// Construct an instance of the ConfigPeerChaincodeLogging model
				configPeerChaincodeLoggingModel := new(blockchainv3.ConfigPeerChaincodeLogging)
				configPeerChaincodeLoggingModel.Level = core.StringPtr("info")
				configPeerChaincodeLoggingModel.Shim = core.StringPtr("warning")
				configPeerChaincodeLoggingModel.Format = core.StringPtr("%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}")

				// Construct an instance of the ConfigPeerChaincode model
				configPeerChaincodeModel := new(blockchainv3.ConfigPeerChaincode)
				configPeerChaincodeModel.Golang = configPeerChaincodeGolangModel
				configPeerChaincodeModel.ExternalBuilders = []blockchainv3.ConfigPeerChaincodeExternalBuildersItem{*configPeerChaincodeExternalBuildersItemModel}
				configPeerChaincodeModel.InstallTimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Startuptimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Executetimeout = core.StringPtr("30s")
				configPeerChaincodeModel.System = configPeerChaincodeSystemModel
				configPeerChaincodeModel.Logging = configPeerChaincodeLoggingModel

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv3.MetricsStatsd)
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv3.Metrics)
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel

				// Construct an instance of the ConfigPeerUpdate model
				configPeerUpdateModel := new(blockchainv3.ConfigPeerUpdate)
				configPeerUpdateModel.Peer = configPeerUpdatePeerModel
				configPeerUpdateModel.Chaincode = configPeerChaincodeModel
				configPeerUpdateModel.Metrics = metricsModel

				// Construct an instance of the CryptoEnrollmentComponent model
				cryptoEnrollmentComponentModel := new(blockchainv3.CryptoEnrollmentComponent)
				cryptoEnrollmentComponentModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the UpdateEnrollmentCryptoFieldCa model
				updateEnrollmentCryptoFieldCaModel := new(blockchainv3.UpdateEnrollmentCryptoFieldCa)
				updateEnrollmentCryptoFieldCaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				updateEnrollmentCryptoFieldCaModel.Port = core.Float64Ptr(float64(7054))
				updateEnrollmentCryptoFieldCaModel.Name = core.StringPtr("ca")
				updateEnrollmentCryptoFieldCaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateEnrollmentCryptoFieldCaModel.EnrollID = core.StringPtr("admin")
				updateEnrollmentCryptoFieldCaModel.EnrollSecret = core.StringPtr("password")

				// Construct an instance of the UpdateEnrollmentCryptoFieldTlsca model
				updateEnrollmentCryptoFieldTlscaModel := new(blockchainv3.UpdateEnrollmentCryptoFieldTlsca)
				updateEnrollmentCryptoFieldTlscaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				updateEnrollmentCryptoFieldTlscaModel.Port = core.Float64Ptr(float64(7054))
				updateEnrollmentCryptoFieldTlscaModel.Name = core.StringPtr("tlsca")
				updateEnrollmentCryptoFieldTlscaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateEnrollmentCryptoFieldTlscaModel.EnrollID = core.StringPtr("admin")
				updateEnrollmentCryptoFieldTlscaModel.EnrollSecret = core.StringPtr("password")
				updateEnrollmentCryptoFieldTlscaModel.CsrHosts = []string{"testString"}

				// Construct an instance of the UpdateEnrollmentCryptoField model
				updateEnrollmentCryptoFieldModel := new(blockchainv3.UpdateEnrollmentCryptoField)
				updateEnrollmentCryptoFieldModel.Component = cryptoEnrollmentComponentModel
				updateEnrollmentCryptoFieldModel.Ca = updateEnrollmentCryptoFieldCaModel
				updateEnrollmentCryptoFieldModel.Tlsca = updateEnrollmentCryptoFieldTlscaModel

				// Construct an instance of the UpdateMspCryptoFieldCa model
				updateMspCryptoFieldCaModel := new(blockchainv3.UpdateMspCryptoFieldCa)
				updateMspCryptoFieldCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateMspCryptoFieldCaModel.CaIntermediateCerts = []string{"testString"}

				// Construct an instance of the UpdateMspCryptoFieldTlsca model
				updateMspCryptoFieldTlscaModel := new(blockchainv3.UpdateMspCryptoFieldTlsca)
				updateMspCryptoFieldTlscaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateMspCryptoFieldTlscaModel.CaIntermediateCerts = []string{"testString"}

				// Construct an instance of the ClientAuth model
				clientAuthModel := new(blockchainv3.ClientAuth)
				clientAuthModel.Type = core.StringPtr("noclientcert")
				clientAuthModel.TlsCerts = []string{"testString"}

				// Construct an instance of the UpdateMspCryptoFieldComponent model
				updateMspCryptoFieldComponentModel := new(blockchainv3.UpdateMspCryptoFieldComponent)
				updateMspCryptoFieldComponentModel.Ekey = core.StringPtr("testString")
				updateMspCryptoFieldComponentModel.Ecert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateMspCryptoFieldComponentModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateMspCryptoFieldComponentModel.TlsKey = core.StringPtr("testString")
				updateMspCryptoFieldComponentModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateMspCryptoFieldComponentModel.ClientAuth = clientAuthModel

				// Construct an instance of the UpdateMspCryptoField model
				updateMspCryptoFieldModel := new(blockchainv3.UpdateMspCryptoField)
				updateMspCryptoFieldModel.Ca = updateMspCryptoFieldCaModel
				updateMspCryptoFieldModel.Tlsca = updateMspCryptoFieldTlscaModel
				updateMspCryptoFieldModel.Component = updateMspCryptoFieldComponentModel

				// Construct an instance of the UpdatePeerBodyCrypto model
				updatePeerBodyCryptoModel := new(blockchainv3.UpdatePeerBodyCrypto)
				updatePeerBodyCryptoModel.Enrollment = updateEnrollmentCryptoFieldModel
				updatePeerBodyCryptoModel.Msp = updateMspCryptoFieldModel

				// Construct an instance of the NodeOu model
				nodeOuModel := new(blockchainv3.NodeOu)
				nodeOuModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv3.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv3.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceObjectFabV2 model
				resourceObjectFabV2Model := new(blockchainv3.ResourceObjectFabV2)
				resourceObjectFabV2Model.Requests = resourceRequestsModel
				resourceObjectFabV2Model.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectCouchDb model
				resourceObjectCouchDbModel := new(blockchainv3.ResourceObjectCouchDb)
				resourceObjectCouchDbModel.Requests = resourceRequestsModel
				resourceObjectCouchDbModel.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv3.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectFabV1 model
				resourceObjectFabV1Model := new(blockchainv3.ResourceObjectFabV1)
				resourceObjectFabV1Model.Requests = resourceRequestsModel
				resourceObjectFabV1Model.Limits = resourceLimitsModel

				// Construct an instance of the PeerResources model
				peerResourcesModel := new(blockchainv3.PeerResources)
				peerResourcesModel.Chaincodelauncher = resourceObjectFabV2Model
				peerResourcesModel.Couchdb = resourceObjectCouchDbModel
				peerResourcesModel.Statedb = resourceObjectModel
				peerResourcesModel.Dind = resourceObjectFabV1Model
				peerResourcesModel.Fluentd = resourceObjectFabV1Model
				peerResourcesModel.Peer = resourceObjectModel
				peerResourcesModel.Proxy = resourceObjectModel

				// Construct an instance of the UpdatePeerOptions model
				updatePeerOptionsModel := new(blockchainv3.UpdatePeerOptions)
				updatePeerOptionsModel.ID = core.StringPtr("testString")
				updatePeerOptionsModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updatePeerOptionsModel.ConfigOverride = configPeerUpdateModel
				updatePeerOptionsModel.Crypto = updatePeerBodyCryptoModel
				updatePeerOptionsModel.NodeOu = nodeOuModel
				updatePeerOptionsModel.Replicas = core.Float64Ptr(float64(1))
				updatePeerOptionsModel.Resources = peerResourcesModel
				updatePeerOptionsModel.Version = core.StringPtr("1.4.6-1")
				updatePeerOptionsModel.Zone = core.StringPtr("-")
				updatePeerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.UpdatePeer(updatePeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.UpdatePeer(updatePeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdatePeer(updatePeerOptions *UpdatePeerOptions)`, func() {
		updatePeerPath := "/ak/api/v3/kubernetes/components/fabric-peer/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePeerPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "component1", "dep_component_id": "admin", "api_url": "grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051", "display_name": "My Peer", "grpcwp_url": "https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084", "location": "ibmcloud", "operations_url": "https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443", "config_override": {"anyKey": "anyValue"}, "node_ou": {"enabled": true}, "msp": {"ca": {"name": "ca", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "tlsca": {"name": "tlsca", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "component": {"tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "ecert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "admin_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}}, "msp_id": "Org1", "resources": {"peer": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "proxy": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "statedb": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "state_db": "couchdb", "storage": {"peer": {"size": "4GiB", "class": "default"}, "statedb": {"size": "4GiB", "class": "default"}}, "tags": ["fabric-ca"], "timestamp": 1537262855753, "type": "fabric-peer", "version": "1.4.6-1", "zone": "-"}`)
				}))
			})
			It(`Invoke UpdatePeer successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.UpdatePeer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ConfigPeerKeepaliveClient model
				configPeerKeepaliveClientModel := new(blockchainv3.ConfigPeerKeepaliveClient)
				configPeerKeepaliveClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveClientModel.Timeout = core.StringPtr("20s")

				// Construct an instance of the ConfigPeerKeepaliveDeliveryClient model
				configPeerKeepaliveDeliveryClientModel := new(blockchainv3.ConfigPeerKeepaliveDeliveryClient)
				configPeerKeepaliveDeliveryClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveDeliveryClientModel.Timeout = core.StringPtr("20s")

				// Construct an instance of the ConfigPeerKeepalive model
				configPeerKeepaliveModel := new(blockchainv3.ConfigPeerKeepalive)
				configPeerKeepaliveModel.MinInterval = core.StringPtr("60s")
				configPeerKeepaliveModel.Client = configPeerKeepaliveClientModel
				configPeerKeepaliveModel.DeliveryClient = configPeerKeepaliveDeliveryClientModel

				// Construct an instance of the ConfigPeerGossipElection model
				configPeerGossipElectionModel := new(blockchainv3.ConfigPeerGossipElection)
				configPeerGossipElectionModel.StartupGracePeriod = core.StringPtr("15s")
				configPeerGossipElectionModel.MembershipSampleInterval = core.StringPtr("1s")
				configPeerGossipElectionModel.LeaderAliveThreshold = core.StringPtr("10s")
				configPeerGossipElectionModel.LeaderElectionDuration = core.StringPtr("5s")

				// Construct an instance of the ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy model
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel := new(blockchainv3.ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy)
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.RequiredPeerCount = core.Float64Ptr(float64(0))
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.MaxPeerCount = core.Float64Ptr(float64(1))

				// Construct an instance of the ConfigPeerGossipPvtData model
				configPeerGossipPvtDataModel := new(blockchainv3.ConfigPeerGossipPvtData)
				configPeerGossipPvtDataModel.PullRetryThreshold = core.StringPtr("60s")
				configPeerGossipPvtDataModel.TransientstoreMaxBlockRetention = core.Float64Ptr(float64(1000))
				configPeerGossipPvtDataModel.PushAckTimeout = core.StringPtr("3s")
				configPeerGossipPvtDataModel.BtlPullMargin = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileBatchSize = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileSleepInterval = core.StringPtr("1m")
				configPeerGossipPvtDataModel.ReconciliationEnabled = core.BoolPtr(true)
				configPeerGossipPvtDataModel.SkipPullingInvalidTransactionsDuringCommit = core.BoolPtr(false)
				configPeerGossipPvtDataModel.ImplicitCollectionDisseminationPolicy = configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel

				// Construct an instance of the ConfigPeerGossipState model
				configPeerGossipStateModel := new(blockchainv3.ConfigPeerGossipState)
				configPeerGossipStateModel.Enabled = core.BoolPtr(true)
				configPeerGossipStateModel.CheckInterval = core.StringPtr("10s")
				configPeerGossipStateModel.ResponseTimeout = core.StringPtr("3s")
				configPeerGossipStateModel.BatchSize = core.Float64Ptr(float64(10))
				configPeerGossipStateModel.BlockBufferSize = core.Float64Ptr(float64(100))
				configPeerGossipStateModel.MaxRetries = core.Float64Ptr(float64(3))

				// Construct an instance of the ConfigPeerGossip model
				configPeerGossipModel := new(blockchainv3.ConfigPeerGossip)
				configPeerGossipModel.UseLeaderElection = core.BoolPtr(true)
				configPeerGossipModel.OrgLeader = core.BoolPtr(false)
				configPeerGossipModel.MembershipTrackerInterval = core.StringPtr("5s")
				configPeerGossipModel.MaxBlockCountToStore = core.Float64Ptr(float64(100))
				configPeerGossipModel.MaxPropagationBurstLatency = core.StringPtr("10ms")
				configPeerGossipModel.MaxPropagationBurstSize = core.Float64Ptr(float64(10))
				configPeerGossipModel.PropagateIterations = core.Float64Ptr(float64(3))
				configPeerGossipModel.PullInterval = core.StringPtr("4s")
				configPeerGossipModel.PullPeerNum = core.Float64Ptr(float64(3))
				configPeerGossipModel.RequestStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.PublishStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.StateInfoRetentionInterval = core.StringPtr("0s")
				configPeerGossipModel.PublishCertPeriod = core.StringPtr("10s")
				configPeerGossipModel.SkipBlockVerification = core.BoolPtr(false)
				configPeerGossipModel.DialTimeout = core.StringPtr("3s")
				configPeerGossipModel.ConnTimeout = core.StringPtr("2s")
				configPeerGossipModel.RecvBuffSize = core.Float64Ptr(float64(20))
				configPeerGossipModel.SendBuffSize = core.Float64Ptr(float64(200))
				configPeerGossipModel.DigestWaitTime = core.StringPtr("1s")
				configPeerGossipModel.RequestWaitTime = core.StringPtr("1500ms")
				configPeerGossipModel.ResponseWaitTime = core.StringPtr("2s")
				configPeerGossipModel.AliveTimeInterval = core.StringPtr("5s")
				configPeerGossipModel.AliveExpirationTimeout = core.StringPtr("25s")
				configPeerGossipModel.ReconnectInterval = core.StringPtr("25s")
				configPeerGossipModel.Election = configPeerGossipElectionModel
				configPeerGossipModel.PvtData = configPeerGossipPvtDataModel
				configPeerGossipModel.State = configPeerGossipStateModel

				// Construct an instance of the ConfigPeerAuthentication model
				configPeerAuthenticationModel := new(blockchainv3.ConfigPeerAuthentication)
				configPeerAuthenticationModel.Timewindow = core.StringPtr("15m")

				// Construct an instance of the ConfigPeerClient model
				configPeerClientModel := new(blockchainv3.ConfigPeerClient)
				configPeerClientModel.ConnTimeout = core.StringPtr("2s")

				// Construct an instance of the ConfigPeerDeliveryclientAddressOverridesItem model
				configPeerDeliveryclientAddressOverridesItemModel := new(blockchainv3.ConfigPeerDeliveryclientAddressOverridesItem)
				configPeerDeliveryclientAddressOverridesItemModel.From = core.StringPtr("n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.To = core.StringPtr("n3a3ec3-myorderer2.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.CaCertsFile = core.StringPtr("my-data/cert.pem")

				// Construct an instance of the ConfigPeerDeliveryclient model
				configPeerDeliveryclientModel := new(blockchainv3.ConfigPeerDeliveryclient)
				configPeerDeliveryclientModel.ReconnectTotalTimeThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.ConnTimeout = core.StringPtr("2s")
				configPeerDeliveryclientModel.ReConnectBackoffThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.AddressOverrides = []blockchainv3.ConfigPeerDeliveryclientAddressOverridesItem{*configPeerDeliveryclientAddressOverridesItemModel}

				// Construct an instance of the ConfigPeerAdminService model
				configPeerAdminServiceModel := new(blockchainv3.ConfigPeerAdminService)
				configPeerAdminServiceModel.ListenAddress = core.StringPtr("0.0.0.0:7051")

				// Construct an instance of the ConfigPeerDiscovery model
				configPeerDiscoveryModel := new(blockchainv3.ConfigPeerDiscovery)
				configPeerDiscoveryModel.Enabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheEnabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheMaxSize = core.Float64Ptr(float64(1000))
				configPeerDiscoveryModel.AuthCachePurgeRetentionRatio = core.Float64Ptr(float64(0.75))
				configPeerDiscoveryModel.OrgMembersAllowedAccess = core.BoolPtr(false)

				// Construct an instance of the ConfigPeerLimitsConcurrency model
				configPeerLimitsConcurrencyModel := new(blockchainv3.ConfigPeerLimitsConcurrency)
				configPeerLimitsConcurrencyModel.EndorserService = core.Float64Ptr(float64(2500))
				configPeerLimitsConcurrencyModel.DeliverService = core.Float64Ptr(float64(2500))

				// Construct an instance of the ConfigPeerLimits model
				configPeerLimitsModel := new(blockchainv3.ConfigPeerLimits)
				configPeerLimitsModel.Concurrency = configPeerLimitsConcurrencyModel

				// Construct an instance of the ConfigPeerGateway model
				configPeerGatewayModel := new(blockchainv3.ConfigPeerGateway)
				configPeerGatewayModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the ConfigPeerUpdatePeer model
				configPeerUpdatePeerModel := new(blockchainv3.ConfigPeerUpdatePeer)
				configPeerUpdatePeerModel.ID = core.StringPtr("john-doe")
				configPeerUpdatePeerModel.NetworkID = core.StringPtr("dev")
				configPeerUpdatePeerModel.Keepalive = configPeerKeepaliveModel
				configPeerUpdatePeerModel.Gossip = configPeerGossipModel
				configPeerUpdatePeerModel.Authentication = configPeerAuthenticationModel
				configPeerUpdatePeerModel.Client = configPeerClientModel
				configPeerUpdatePeerModel.Deliveryclient = configPeerDeliveryclientModel
				configPeerUpdatePeerModel.AdminService = configPeerAdminServiceModel
				configPeerUpdatePeerModel.ValidatorPoolSize = core.Float64Ptr(float64(8))
				configPeerUpdatePeerModel.Discovery = configPeerDiscoveryModel
				configPeerUpdatePeerModel.Limits = configPeerLimitsModel
				configPeerUpdatePeerModel.Gateway = configPeerGatewayModel

				// Construct an instance of the ConfigPeerChaincodeGolang model
				configPeerChaincodeGolangModel := new(blockchainv3.ConfigPeerChaincodeGolang)
				configPeerChaincodeGolangModel.DynamicLink = core.BoolPtr(false)

				// Construct an instance of the ConfigPeerChaincodeExternalBuildersItem model
				configPeerChaincodeExternalBuildersItemModel := new(blockchainv3.ConfigPeerChaincodeExternalBuildersItem)
				configPeerChaincodeExternalBuildersItemModel.Path = core.StringPtr("/path/to/directory")
				configPeerChaincodeExternalBuildersItemModel.Name = core.StringPtr("descriptive-build-name")
				configPeerChaincodeExternalBuildersItemModel.EnvironmentWhitelist = []string{"GOPROXY"}

				// Construct an instance of the ConfigPeerChaincodeSystem model
				configPeerChaincodeSystemModel := new(blockchainv3.ConfigPeerChaincodeSystem)
				configPeerChaincodeSystemModel.Cscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Lscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Escc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Vscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Qscc = core.BoolPtr(true)

				// Construct an instance of the ConfigPeerChaincodeLogging model
				configPeerChaincodeLoggingModel := new(blockchainv3.ConfigPeerChaincodeLogging)
				configPeerChaincodeLoggingModel.Level = core.StringPtr("info")
				configPeerChaincodeLoggingModel.Shim = core.StringPtr("warning")
				configPeerChaincodeLoggingModel.Format = core.StringPtr("%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}")

				// Construct an instance of the ConfigPeerChaincode model
				configPeerChaincodeModel := new(blockchainv3.ConfigPeerChaincode)
				configPeerChaincodeModel.Golang = configPeerChaincodeGolangModel
				configPeerChaincodeModel.ExternalBuilders = []blockchainv3.ConfigPeerChaincodeExternalBuildersItem{*configPeerChaincodeExternalBuildersItemModel}
				configPeerChaincodeModel.InstallTimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Startuptimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Executetimeout = core.StringPtr("30s")
				configPeerChaincodeModel.System = configPeerChaincodeSystemModel
				configPeerChaincodeModel.Logging = configPeerChaincodeLoggingModel

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv3.MetricsStatsd)
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv3.Metrics)
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel

				// Construct an instance of the ConfigPeerUpdate model
				configPeerUpdateModel := new(blockchainv3.ConfigPeerUpdate)
				configPeerUpdateModel.Peer = configPeerUpdatePeerModel
				configPeerUpdateModel.Chaincode = configPeerChaincodeModel
				configPeerUpdateModel.Metrics = metricsModel

				// Construct an instance of the CryptoEnrollmentComponent model
				cryptoEnrollmentComponentModel := new(blockchainv3.CryptoEnrollmentComponent)
				cryptoEnrollmentComponentModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the UpdateEnrollmentCryptoFieldCa model
				updateEnrollmentCryptoFieldCaModel := new(blockchainv3.UpdateEnrollmentCryptoFieldCa)
				updateEnrollmentCryptoFieldCaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				updateEnrollmentCryptoFieldCaModel.Port = core.Float64Ptr(float64(7054))
				updateEnrollmentCryptoFieldCaModel.Name = core.StringPtr("ca")
				updateEnrollmentCryptoFieldCaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateEnrollmentCryptoFieldCaModel.EnrollID = core.StringPtr("admin")
				updateEnrollmentCryptoFieldCaModel.EnrollSecret = core.StringPtr("password")

				// Construct an instance of the UpdateEnrollmentCryptoFieldTlsca model
				updateEnrollmentCryptoFieldTlscaModel := new(blockchainv3.UpdateEnrollmentCryptoFieldTlsca)
				updateEnrollmentCryptoFieldTlscaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				updateEnrollmentCryptoFieldTlscaModel.Port = core.Float64Ptr(float64(7054))
				updateEnrollmentCryptoFieldTlscaModel.Name = core.StringPtr("tlsca")
				updateEnrollmentCryptoFieldTlscaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateEnrollmentCryptoFieldTlscaModel.EnrollID = core.StringPtr("admin")
				updateEnrollmentCryptoFieldTlscaModel.EnrollSecret = core.StringPtr("password")
				updateEnrollmentCryptoFieldTlscaModel.CsrHosts = []string{"testString"}

				// Construct an instance of the UpdateEnrollmentCryptoField model
				updateEnrollmentCryptoFieldModel := new(blockchainv3.UpdateEnrollmentCryptoField)
				updateEnrollmentCryptoFieldModel.Component = cryptoEnrollmentComponentModel
				updateEnrollmentCryptoFieldModel.Ca = updateEnrollmentCryptoFieldCaModel
				updateEnrollmentCryptoFieldModel.Tlsca = updateEnrollmentCryptoFieldTlscaModel

				// Construct an instance of the UpdateMspCryptoFieldCa model
				updateMspCryptoFieldCaModel := new(blockchainv3.UpdateMspCryptoFieldCa)
				updateMspCryptoFieldCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateMspCryptoFieldCaModel.CaIntermediateCerts = []string{"testString"}

				// Construct an instance of the UpdateMspCryptoFieldTlsca model
				updateMspCryptoFieldTlscaModel := new(blockchainv3.UpdateMspCryptoFieldTlsca)
				updateMspCryptoFieldTlscaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateMspCryptoFieldTlscaModel.CaIntermediateCerts = []string{"testString"}

				// Construct an instance of the ClientAuth model
				clientAuthModel := new(blockchainv3.ClientAuth)
				clientAuthModel.Type = core.StringPtr("noclientcert")
				clientAuthModel.TlsCerts = []string{"testString"}

				// Construct an instance of the UpdateMspCryptoFieldComponent model
				updateMspCryptoFieldComponentModel := new(blockchainv3.UpdateMspCryptoFieldComponent)
				updateMspCryptoFieldComponentModel.Ekey = core.StringPtr("testString")
				updateMspCryptoFieldComponentModel.Ecert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateMspCryptoFieldComponentModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateMspCryptoFieldComponentModel.TlsKey = core.StringPtr("testString")
				updateMspCryptoFieldComponentModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateMspCryptoFieldComponentModel.ClientAuth = clientAuthModel

				// Construct an instance of the UpdateMspCryptoField model
				updateMspCryptoFieldModel := new(blockchainv3.UpdateMspCryptoField)
				updateMspCryptoFieldModel.Ca = updateMspCryptoFieldCaModel
				updateMspCryptoFieldModel.Tlsca = updateMspCryptoFieldTlscaModel
				updateMspCryptoFieldModel.Component = updateMspCryptoFieldComponentModel

				// Construct an instance of the UpdatePeerBodyCrypto model
				updatePeerBodyCryptoModel := new(blockchainv3.UpdatePeerBodyCrypto)
				updatePeerBodyCryptoModel.Enrollment = updateEnrollmentCryptoFieldModel
				updatePeerBodyCryptoModel.Msp = updateMspCryptoFieldModel

				// Construct an instance of the NodeOu model
				nodeOuModel := new(blockchainv3.NodeOu)
				nodeOuModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv3.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv3.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceObjectFabV2 model
				resourceObjectFabV2Model := new(blockchainv3.ResourceObjectFabV2)
				resourceObjectFabV2Model.Requests = resourceRequestsModel
				resourceObjectFabV2Model.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectCouchDb model
				resourceObjectCouchDbModel := new(blockchainv3.ResourceObjectCouchDb)
				resourceObjectCouchDbModel.Requests = resourceRequestsModel
				resourceObjectCouchDbModel.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv3.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectFabV1 model
				resourceObjectFabV1Model := new(blockchainv3.ResourceObjectFabV1)
				resourceObjectFabV1Model.Requests = resourceRequestsModel
				resourceObjectFabV1Model.Limits = resourceLimitsModel

				// Construct an instance of the PeerResources model
				peerResourcesModel := new(blockchainv3.PeerResources)
				peerResourcesModel.Chaincodelauncher = resourceObjectFabV2Model
				peerResourcesModel.Couchdb = resourceObjectCouchDbModel
				peerResourcesModel.Statedb = resourceObjectModel
				peerResourcesModel.Dind = resourceObjectFabV1Model
				peerResourcesModel.Fluentd = resourceObjectFabV1Model
				peerResourcesModel.Peer = resourceObjectModel
				peerResourcesModel.Proxy = resourceObjectModel

				// Construct an instance of the UpdatePeerOptions model
				updatePeerOptionsModel := new(blockchainv3.UpdatePeerOptions)
				updatePeerOptionsModel.ID = core.StringPtr("testString")
				updatePeerOptionsModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updatePeerOptionsModel.ConfigOverride = configPeerUpdateModel
				updatePeerOptionsModel.Crypto = updatePeerBodyCryptoModel
				updatePeerOptionsModel.NodeOu = nodeOuModel
				updatePeerOptionsModel.Replicas = core.Float64Ptr(float64(1))
				updatePeerOptionsModel.Resources = peerResourcesModel
				updatePeerOptionsModel.Version = core.StringPtr("1.4.6-1")
				updatePeerOptionsModel.Zone = core.StringPtr("-")
				updatePeerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.UpdatePeer(updatePeerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.UpdatePeerWithContext(ctx, updatePeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.UpdatePeer(updatePeerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.UpdatePeerWithContext(ctx, updatePeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdatePeer with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the ConfigPeerKeepaliveClient model
				configPeerKeepaliveClientModel := new(blockchainv3.ConfigPeerKeepaliveClient)
				configPeerKeepaliveClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveClientModel.Timeout = core.StringPtr("20s")

				// Construct an instance of the ConfigPeerKeepaliveDeliveryClient model
				configPeerKeepaliveDeliveryClientModel := new(blockchainv3.ConfigPeerKeepaliveDeliveryClient)
				configPeerKeepaliveDeliveryClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveDeliveryClientModel.Timeout = core.StringPtr("20s")

				// Construct an instance of the ConfigPeerKeepalive model
				configPeerKeepaliveModel := new(blockchainv3.ConfigPeerKeepalive)
				configPeerKeepaliveModel.MinInterval = core.StringPtr("60s")
				configPeerKeepaliveModel.Client = configPeerKeepaliveClientModel
				configPeerKeepaliveModel.DeliveryClient = configPeerKeepaliveDeliveryClientModel

				// Construct an instance of the ConfigPeerGossipElection model
				configPeerGossipElectionModel := new(blockchainv3.ConfigPeerGossipElection)
				configPeerGossipElectionModel.StartupGracePeriod = core.StringPtr("15s")
				configPeerGossipElectionModel.MembershipSampleInterval = core.StringPtr("1s")
				configPeerGossipElectionModel.LeaderAliveThreshold = core.StringPtr("10s")
				configPeerGossipElectionModel.LeaderElectionDuration = core.StringPtr("5s")

				// Construct an instance of the ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy model
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel := new(blockchainv3.ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy)
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.RequiredPeerCount = core.Float64Ptr(float64(0))
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.MaxPeerCount = core.Float64Ptr(float64(1))

				// Construct an instance of the ConfigPeerGossipPvtData model
				configPeerGossipPvtDataModel := new(blockchainv3.ConfigPeerGossipPvtData)
				configPeerGossipPvtDataModel.PullRetryThreshold = core.StringPtr("60s")
				configPeerGossipPvtDataModel.TransientstoreMaxBlockRetention = core.Float64Ptr(float64(1000))
				configPeerGossipPvtDataModel.PushAckTimeout = core.StringPtr("3s")
				configPeerGossipPvtDataModel.BtlPullMargin = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileBatchSize = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileSleepInterval = core.StringPtr("1m")
				configPeerGossipPvtDataModel.ReconciliationEnabled = core.BoolPtr(true)
				configPeerGossipPvtDataModel.SkipPullingInvalidTransactionsDuringCommit = core.BoolPtr(false)
				configPeerGossipPvtDataModel.ImplicitCollectionDisseminationPolicy = configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel

				// Construct an instance of the ConfigPeerGossipState model
				configPeerGossipStateModel := new(blockchainv3.ConfigPeerGossipState)
				configPeerGossipStateModel.Enabled = core.BoolPtr(true)
				configPeerGossipStateModel.CheckInterval = core.StringPtr("10s")
				configPeerGossipStateModel.ResponseTimeout = core.StringPtr("3s")
				configPeerGossipStateModel.BatchSize = core.Float64Ptr(float64(10))
				configPeerGossipStateModel.BlockBufferSize = core.Float64Ptr(float64(100))
				configPeerGossipStateModel.MaxRetries = core.Float64Ptr(float64(3))

				// Construct an instance of the ConfigPeerGossip model
				configPeerGossipModel := new(blockchainv3.ConfigPeerGossip)
				configPeerGossipModel.UseLeaderElection = core.BoolPtr(true)
				configPeerGossipModel.OrgLeader = core.BoolPtr(false)
				configPeerGossipModel.MembershipTrackerInterval = core.StringPtr("5s")
				configPeerGossipModel.MaxBlockCountToStore = core.Float64Ptr(float64(100))
				configPeerGossipModel.MaxPropagationBurstLatency = core.StringPtr("10ms")
				configPeerGossipModel.MaxPropagationBurstSize = core.Float64Ptr(float64(10))
				configPeerGossipModel.PropagateIterations = core.Float64Ptr(float64(3))
				configPeerGossipModel.PullInterval = core.StringPtr("4s")
				configPeerGossipModel.PullPeerNum = core.Float64Ptr(float64(3))
				configPeerGossipModel.RequestStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.PublishStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.StateInfoRetentionInterval = core.StringPtr("0s")
				configPeerGossipModel.PublishCertPeriod = core.StringPtr("10s")
				configPeerGossipModel.SkipBlockVerification = core.BoolPtr(false)
				configPeerGossipModel.DialTimeout = core.StringPtr("3s")
				configPeerGossipModel.ConnTimeout = core.StringPtr("2s")
				configPeerGossipModel.RecvBuffSize = core.Float64Ptr(float64(20))
				configPeerGossipModel.SendBuffSize = core.Float64Ptr(float64(200))
				configPeerGossipModel.DigestWaitTime = core.StringPtr("1s")
				configPeerGossipModel.RequestWaitTime = core.StringPtr("1500ms")
				configPeerGossipModel.ResponseWaitTime = core.StringPtr("2s")
				configPeerGossipModel.AliveTimeInterval = core.StringPtr("5s")
				configPeerGossipModel.AliveExpirationTimeout = core.StringPtr("25s")
				configPeerGossipModel.ReconnectInterval = core.StringPtr("25s")
				configPeerGossipModel.Election = configPeerGossipElectionModel
				configPeerGossipModel.PvtData = configPeerGossipPvtDataModel
				configPeerGossipModel.State = configPeerGossipStateModel

				// Construct an instance of the ConfigPeerAuthentication model
				configPeerAuthenticationModel := new(blockchainv3.ConfigPeerAuthentication)
				configPeerAuthenticationModel.Timewindow = core.StringPtr("15m")

				// Construct an instance of the ConfigPeerClient model
				configPeerClientModel := new(blockchainv3.ConfigPeerClient)
				configPeerClientModel.ConnTimeout = core.StringPtr("2s")

				// Construct an instance of the ConfigPeerDeliveryclientAddressOverridesItem model
				configPeerDeliveryclientAddressOverridesItemModel := new(blockchainv3.ConfigPeerDeliveryclientAddressOverridesItem)
				configPeerDeliveryclientAddressOverridesItemModel.From = core.StringPtr("n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.To = core.StringPtr("n3a3ec3-myorderer2.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.CaCertsFile = core.StringPtr("my-data/cert.pem")

				// Construct an instance of the ConfigPeerDeliveryclient model
				configPeerDeliveryclientModel := new(blockchainv3.ConfigPeerDeliveryclient)
				configPeerDeliveryclientModel.ReconnectTotalTimeThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.ConnTimeout = core.StringPtr("2s")
				configPeerDeliveryclientModel.ReConnectBackoffThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.AddressOverrides = []blockchainv3.ConfigPeerDeliveryclientAddressOverridesItem{*configPeerDeliveryclientAddressOverridesItemModel}

				// Construct an instance of the ConfigPeerAdminService model
				configPeerAdminServiceModel := new(blockchainv3.ConfigPeerAdminService)
				configPeerAdminServiceModel.ListenAddress = core.StringPtr("0.0.0.0:7051")

				// Construct an instance of the ConfigPeerDiscovery model
				configPeerDiscoveryModel := new(blockchainv3.ConfigPeerDiscovery)
				configPeerDiscoveryModel.Enabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheEnabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheMaxSize = core.Float64Ptr(float64(1000))
				configPeerDiscoveryModel.AuthCachePurgeRetentionRatio = core.Float64Ptr(float64(0.75))
				configPeerDiscoveryModel.OrgMembersAllowedAccess = core.BoolPtr(false)

				// Construct an instance of the ConfigPeerLimitsConcurrency model
				configPeerLimitsConcurrencyModel := new(blockchainv3.ConfigPeerLimitsConcurrency)
				configPeerLimitsConcurrencyModel.EndorserService = core.Float64Ptr(float64(2500))
				configPeerLimitsConcurrencyModel.DeliverService = core.Float64Ptr(float64(2500))

				// Construct an instance of the ConfigPeerLimits model
				configPeerLimitsModel := new(blockchainv3.ConfigPeerLimits)
				configPeerLimitsModel.Concurrency = configPeerLimitsConcurrencyModel

				// Construct an instance of the ConfigPeerGateway model
				configPeerGatewayModel := new(blockchainv3.ConfigPeerGateway)
				configPeerGatewayModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the ConfigPeerUpdatePeer model
				configPeerUpdatePeerModel := new(blockchainv3.ConfigPeerUpdatePeer)
				configPeerUpdatePeerModel.ID = core.StringPtr("john-doe")
				configPeerUpdatePeerModel.NetworkID = core.StringPtr("dev")
				configPeerUpdatePeerModel.Keepalive = configPeerKeepaliveModel
				configPeerUpdatePeerModel.Gossip = configPeerGossipModel
				configPeerUpdatePeerModel.Authentication = configPeerAuthenticationModel
				configPeerUpdatePeerModel.Client = configPeerClientModel
				configPeerUpdatePeerModel.Deliveryclient = configPeerDeliveryclientModel
				configPeerUpdatePeerModel.AdminService = configPeerAdminServiceModel
				configPeerUpdatePeerModel.ValidatorPoolSize = core.Float64Ptr(float64(8))
				configPeerUpdatePeerModel.Discovery = configPeerDiscoveryModel
				configPeerUpdatePeerModel.Limits = configPeerLimitsModel
				configPeerUpdatePeerModel.Gateway = configPeerGatewayModel

				// Construct an instance of the ConfigPeerChaincodeGolang model
				configPeerChaincodeGolangModel := new(blockchainv3.ConfigPeerChaincodeGolang)
				configPeerChaincodeGolangModel.DynamicLink = core.BoolPtr(false)

				// Construct an instance of the ConfigPeerChaincodeExternalBuildersItem model
				configPeerChaincodeExternalBuildersItemModel := new(blockchainv3.ConfigPeerChaincodeExternalBuildersItem)
				configPeerChaincodeExternalBuildersItemModel.Path = core.StringPtr("/path/to/directory")
				configPeerChaincodeExternalBuildersItemModel.Name = core.StringPtr("descriptive-build-name")
				configPeerChaincodeExternalBuildersItemModel.EnvironmentWhitelist = []string{"GOPROXY"}

				// Construct an instance of the ConfigPeerChaincodeSystem model
				configPeerChaincodeSystemModel := new(blockchainv3.ConfigPeerChaincodeSystem)
				configPeerChaincodeSystemModel.Cscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Lscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Escc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Vscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Qscc = core.BoolPtr(true)

				// Construct an instance of the ConfigPeerChaincodeLogging model
				configPeerChaincodeLoggingModel := new(blockchainv3.ConfigPeerChaincodeLogging)
				configPeerChaincodeLoggingModel.Level = core.StringPtr("info")
				configPeerChaincodeLoggingModel.Shim = core.StringPtr("warning")
				configPeerChaincodeLoggingModel.Format = core.StringPtr("%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}")

				// Construct an instance of the ConfigPeerChaincode model
				configPeerChaincodeModel := new(blockchainv3.ConfigPeerChaincode)
				configPeerChaincodeModel.Golang = configPeerChaincodeGolangModel
				configPeerChaincodeModel.ExternalBuilders = []blockchainv3.ConfigPeerChaincodeExternalBuildersItem{*configPeerChaincodeExternalBuildersItemModel}
				configPeerChaincodeModel.InstallTimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Startuptimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Executetimeout = core.StringPtr("30s")
				configPeerChaincodeModel.System = configPeerChaincodeSystemModel
				configPeerChaincodeModel.Logging = configPeerChaincodeLoggingModel

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv3.MetricsStatsd)
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv3.Metrics)
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel

				// Construct an instance of the ConfigPeerUpdate model
				configPeerUpdateModel := new(blockchainv3.ConfigPeerUpdate)
				configPeerUpdateModel.Peer = configPeerUpdatePeerModel
				configPeerUpdateModel.Chaincode = configPeerChaincodeModel
				configPeerUpdateModel.Metrics = metricsModel

				// Construct an instance of the CryptoEnrollmentComponent model
				cryptoEnrollmentComponentModel := new(blockchainv3.CryptoEnrollmentComponent)
				cryptoEnrollmentComponentModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the UpdateEnrollmentCryptoFieldCa model
				updateEnrollmentCryptoFieldCaModel := new(blockchainv3.UpdateEnrollmentCryptoFieldCa)
				updateEnrollmentCryptoFieldCaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				updateEnrollmentCryptoFieldCaModel.Port = core.Float64Ptr(float64(7054))
				updateEnrollmentCryptoFieldCaModel.Name = core.StringPtr("ca")
				updateEnrollmentCryptoFieldCaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateEnrollmentCryptoFieldCaModel.EnrollID = core.StringPtr("admin")
				updateEnrollmentCryptoFieldCaModel.EnrollSecret = core.StringPtr("password")

				// Construct an instance of the UpdateEnrollmentCryptoFieldTlsca model
				updateEnrollmentCryptoFieldTlscaModel := new(blockchainv3.UpdateEnrollmentCryptoFieldTlsca)
				updateEnrollmentCryptoFieldTlscaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				updateEnrollmentCryptoFieldTlscaModel.Port = core.Float64Ptr(float64(7054))
				updateEnrollmentCryptoFieldTlscaModel.Name = core.StringPtr("tlsca")
				updateEnrollmentCryptoFieldTlscaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateEnrollmentCryptoFieldTlscaModel.EnrollID = core.StringPtr("admin")
				updateEnrollmentCryptoFieldTlscaModel.EnrollSecret = core.StringPtr("password")
				updateEnrollmentCryptoFieldTlscaModel.CsrHosts = []string{"testString"}

				// Construct an instance of the UpdateEnrollmentCryptoField model
				updateEnrollmentCryptoFieldModel := new(blockchainv3.UpdateEnrollmentCryptoField)
				updateEnrollmentCryptoFieldModel.Component = cryptoEnrollmentComponentModel
				updateEnrollmentCryptoFieldModel.Ca = updateEnrollmentCryptoFieldCaModel
				updateEnrollmentCryptoFieldModel.Tlsca = updateEnrollmentCryptoFieldTlscaModel

				// Construct an instance of the UpdateMspCryptoFieldCa model
				updateMspCryptoFieldCaModel := new(blockchainv3.UpdateMspCryptoFieldCa)
				updateMspCryptoFieldCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateMspCryptoFieldCaModel.CaIntermediateCerts = []string{"testString"}

				// Construct an instance of the UpdateMspCryptoFieldTlsca model
				updateMspCryptoFieldTlscaModel := new(blockchainv3.UpdateMspCryptoFieldTlsca)
				updateMspCryptoFieldTlscaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateMspCryptoFieldTlscaModel.CaIntermediateCerts = []string{"testString"}

				// Construct an instance of the ClientAuth model
				clientAuthModel := new(blockchainv3.ClientAuth)
				clientAuthModel.Type = core.StringPtr("noclientcert")
				clientAuthModel.TlsCerts = []string{"testString"}

				// Construct an instance of the UpdateMspCryptoFieldComponent model
				updateMspCryptoFieldComponentModel := new(blockchainv3.UpdateMspCryptoFieldComponent)
				updateMspCryptoFieldComponentModel.Ekey = core.StringPtr("testString")
				updateMspCryptoFieldComponentModel.Ecert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateMspCryptoFieldComponentModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateMspCryptoFieldComponentModel.TlsKey = core.StringPtr("testString")
				updateMspCryptoFieldComponentModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateMspCryptoFieldComponentModel.ClientAuth = clientAuthModel

				// Construct an instance of the UpdateMspCryptoField model
				updateMspCryptoFieldModel := new(blockchainv3.UpdateMspCryptoField)
				updateMspCryptoFieldModel.Ca = updateMspCryptoFieldCaModel
				updateMspCryptoFieldModel.Tlsca = updateMspCryptoFieldTlscaModel
				updateMspCryptoFieldModel.Component = updateMspCryptoFieldComponentModel

				// Construct an instance of the UpdatePeerBodyCrypto model
				updatePeerBodyCryptoModel := new(blockchainv3.UpdatePeerBodyCrypto)
				updatePeerBodyCryptoModel.Enrollment = updateEnrollmentCryptoFieldModel
				updatePeerBodyCryptoModel.Msp = updateMspCryptoFieldModel

				// Construct an instance of the NodeOu model
				nodeOuModel := new(blockchainv3.NodeOu)
				nodeOuModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv3.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv3.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceObjectFabV2 model
				resourceObjectFabV2Model := new(blockchainv3.ResourceObjectFabV2)
				resourceObjectFabV2Model.Requests = resourceRequestsModel
				resourceObjectFabV2Model.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectCouchDb model
				resourceObjectCouchDbModel := new(blockchainv3.ResourceObjectCouchDb)
				resourceObjectCouchDbModel.Requests = resourceRequestsModel
				resourceObjectCouchDbModel.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv3.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectFabV1 model
				resourceObjectFabV1Model := new(blockchainv3.ResourceObjectFabV1)
				resourceObjectFabV1Model.Requests = resourceRequestsModel
				resourceObjectFabV1Model.Limits = resourceLimitsModel

				// Construct an instance of the PeerResources model
				peerResourcesModel := new(blockchainv3.PeerResources)
				peerResourcesModel.Chaincodelauncher = resourceObjectFabV2Model
				peerResourcesModel.Couchdb = resourceObjectCouchDbModel
				peerResourcesModel.Statedb = resourceObjectModel
				peerResourcesModel.Dind = resourceObjectFabV1Model
				peerResourcesModel.Fluentd = resourceObjectFabV1Model
				peerResourcesModel.Peer = resourceObjectModel
				peerResourcesModel.Proxy = resourceObjectModel

				// Construct an instance of the UpdatePeerOptions model
				updatePeerOptionsModel := new(blockchainv3.UpdatePeerOptions)
				updatePeerOptionsModel.ID = core.StringPtr("testString")
				updatePeerOptionsModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updatePeerOptionsModel.ConfigOverride = configPeerUpdateModel
				updatePeerOptionsModel.Crypto = updatePeerBodyCryptoModel
				updatePeerOptionsModel.NodeOu = nodeOuModel
				updatePeerOptionsModel.Replicas = core.Float64Ptr(float64(1))
				updatePeerOptionsModel.Resources = peerResourcesModel
				updatePeerOptionsModel.Version = core.StringPtr("1.4.6-1")
				updatePeerOptionsModel.Zone = core.StringPtr("-")
				updatePeerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.UpdatePeer(updatePeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdatePeerOptions model with no property values
				updatePeerOptionsModelNew := new(blockchainv3.UpdatePeerOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.UpdatePeer(updatePeerOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateOrderer(createOrdererOptions *CreateOrdererOptions) - Operation response error`, func() {
		createOrdererPath := "/ak/api/v3/kubernetes/components/fabric-orderer"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createOrdererPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateOrderer with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the CryptoEnrollmentComponent model
				cryptoEnrollmentComponentModel := new(blockchainv3.CryptoEnrollmentComponent)
				cryptoEnrollmentComponentModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the CryptoObjectEnrollmentCa model
				cryptoObjectEnrollmentCaModel := new(blockchainv3.CryptoObjectEnrollmentCa)
				cryptoObjectEnrollmentCaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				cryptoObjectEnrollmentCaModel.Port = core.Float64Ptr(float64(7054))
				cryptoObjectEnrollmentCaModel.Name = core.StringPtr("ca")
				cryptoObjectEnrollmentCaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				cryptoObjectEnrollmentCaModel.EnrollID = core.StringPtr("admin")
				cryptoObjectEnrollmentCaModel.EnrollSecret = core.StringPtr("password")

				// Construct an instance of the CryptoObjectEnrollmentTlsca model
				cryptoObjectEnrollmentTlscaModel := new(blockchainv3.CryptoObjectEnrollmentTlsca)
				cryptoObjectEnrollmentTlscaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				cryptoObjectEnrollmentTlscaModel.Port = core.Float64Ptr(float64(7054))
				cryptoObjectEnrollmentTlscaModel.Name = core.StringPtr("tlsca")
				cryptoObjectEnrollmentTlscaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				cryptoObjectEnrollmentTlscaModel.EnrollID = core.StringPtr("admin")
				cryptoObjectEnrollmentTlscaModel.EnrollSecret = core.StringPtr("password")
				cryptoObjectEnrollmentTlscaModel.CsrHosts = []string{"testString"}

				// Construct an instance of the CryptoObjectEnrollment model
				cryptoObjectEnrollmentModel := new(blockchainv3.CryptoObjectEnrollment)
				cryptoObjectEnrollmentModel.Component = cryptoEnrollmentComponentModel
				cryptoObjectEnrollmentModel.Ca = cryptoObjectEnrollmentCaModel
				cryptoObjectEnrollmentModel.Tlsca = cryptoObjectEnrollmentTlscaModel

				// Construct an instance of the ClientAuth model
				clientAuthModel := new(blockchainv3.ClientAuth)
				clientAuthModel.Type = core.StringPtr("noclientcert")
				clientAuthModel.TlsCerts = []string{"testString"}

				// Construct an instance of the MspCryptoComp model
				mspCryptoCompModel := new(blockchainv3.MspCryptoComp)
				mspCryptoCompModel.Ekey = core.StringPtr("testString")
				mspCryptoCompModel.Ecert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoCompModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				mspCryptoCompModel.TlsKey = core.StringPtr("testString")
				mspCryptoCompModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoCompModel.ClientAuth = clientAuthModel

				// Construct an instance of the MspCryptoCa model
				mspCryptoCaModel := new(blockchainv3.MspCryptoCa)
				mspCryptoCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				mspCryptoCaModel.CaIntermediateCerts = []string{"testString"}

				// Construct an instance of the CryptoObjectMsp model
				cryptoObjectMspModel := new(blockchainv3.CryptoObjectMsp)
				cryptoObjectMspModel.Component = mspCryptoCompModel
				cryptoObjectMspModel.Ca = mspCryptoCaModel
				cryptoObjectMspModel.Tlsca = mspCryptoCaModel

				// Construct an instance of the CryptoObject model
				cryptoObjectModel := new(blockchainv3.CryptoObject)
				cryptoObjectModel.Enrollment = cryptoObjectEnrollmentModel
				cryptoObjectModel.Msp = cryptoObjectMspModel

				// Construct an instance of the ConfigOrdererKeepalive model
				configOrdererKeepaliveModel := new(blockchainv3.ConfigOrdererKeepalive)
				configOrdererKeepaliveModel.ServerMinInterval = core.StringPtr("60s")
				configOrdererKeepaliveModel.ServerInterval = core.StringPtr("2h")
				configOrdererKeepaliveModel.ServerTimeout = core.StringPtr("20s")

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv3.BccspSW)
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv3.BccspPKCS11)
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv3.Bccsp)
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model

				// Construct an instance of the ConfigOrdererAuthentication model
				configOrdererAuthenticationModel := new(blockchainv3.ConfigOrdererAuthentication)
				configOrdererAuthenticationModel.TimeWindow = core.StringPtr("15m")
				configOrdererAuthenticationModel.NoExpirationChecks = core.BoolPtr(false)

				// Construct an instance of the ConfigOrdererGeneral model
				configOrdererGeneralModel := new(blockchainv3.ConfigOrdererGeneral)
				configOrdererGeneralModel.Keepalive = configOrdererKeepaliveModel
				configOrdererGeneralModel.BCCSP = bccspModel
				configOrdererGeneralModel.Authentication = configOrdererAuthenticationModel

				// Construct an instance of the ConfigOrdererDebug model
				configOrdererDebugModel := new(blockchainv3.ConfigOrdererDebug)
				configOrdererDebugModel.BroadcastTraceDir = core.StringPtr("testString")
				configOrdererDebugModel.DeliverTraceDir = core.StringPtr("testString")

				// Construct an instance of the ConfigOrdererMetricsStatsd model
				configOrdererMetricsStatsdModel := new(blockchainv3.ConfigOrdererMetricsStatsd)
				configOrdererMetricsStatsdModel.Network = core.StringPtr("udp")
				configOrdererMetricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				configOrdererMetricsStatsdModel.WriteInterval = core.StringPtr("10s")
				configOrdererMetricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the ConfigOrdererMetrics model
				configOrdererMetricsModel := new(blockchainv3.ConfigOrdererMetrics)
				configOrdererMetricsModel.Provider = core.StringPtr("disabled")
				configOrdererMetricsModel.Statsd = configOrdererMetricsStatsdModel

				// Construct an instance of the ConfigOrdererCreate model
				configOrdererCreateModel := new(blockchainv3.ConfigOrdererCreate)
				configOrdererCreateModel.General = configOrdererGeneralModel
				configOrdererCreateModel.Debug = configOrdererDebugModel
				configOrdererCreateModel.Metrics = configOrdererMetricsModel

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv3.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv3.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv3.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the CreateOrdererRaftBodyResources model
				createOrdererRaftBodyResourcesModel := new(blockchainv3.CreateOrdererRaftBodyResources)
				createOrdererRaftBodyResourcesModel.Orderer = resourceObjectModel
				createOrdererRaftBodyResourcesModel.Proxy = resourceObjectModel

				// Construct an instance of the StorageObject model
				storageObjectModel := new(blockchainv3.StorageObject)
				storageObjectModel.Size = core.StringPtr("4GiB")
				storageObjectModel.Class = core.StringPtr("default")

				// Construct an instance of the CreateOrdererRaftBodyStorage model
				createOrdererRaftBodyStorageModel := new(blockchainv3.CreateOrdererRaftBodyStorage)
				createOrdererRaftBodyStorageModel.Orderer = storageObjectModel

				// Construct an instance of the Hsm model
				hsmModel := new(blockchainv3.Hsm)
				hsmModel.Pkcs11endpoint = core.StringPtr("tcp://example.com:666")

				// Construct an instance of the CreateOrdererOptions model
				createOrdererOptionsModel := new(blockchainv3.CreateOrdererOptions)
				createOrdererOptionsModel.OrdererType = core.StringPtr("raft")
				createOrdererOptionsModel.MspID = core.StringPtr("Org1")
				createOrdererOptionsModel.DisplayName = core.StringPtr("orderer")
				createOrdererOptionsModel.Crypto = []blockchainv3.CryptoObject{*cryptoObjectModel}
				createOrdererOptionsModel.ClusterName = core.StringPtr("ordering service 1")
				createOrdererOptionsModel.ID = core.StringPtr("component1")
				createOrdererOptionsModel.ClusterID = core.StringPtr("abcde")
				createOrdererOptionsModel.ExternalAppend = core.BoolPtr(false)
				createOrdererOptionsModel.ConfigOverride = []blockchainv3.ConfigOrdererCreate{*configOrdererCreateModel}
				createOrdererOptionsModel.Resources = createOrdererRaftBodyResourcesModel
				createOrdererOptionsModel.Storage = createOrdererRaftBodyStorageModel
				createOrdererOptionsModel.SystemChannelID = core.StringPtr("testchainid")
				createOrdererOptionsModel.Zone = []string{"-"}
				createOrdererOptionsModel.Tags = []string{"fabric-ca"}
				createOrdererOptionsModel.Region = []string{"-"}
				createOrdererOptionsModel.Hsm = hsmModel
				createOrdererOptionsModel.Version = core.StringPtr("1.4.6-1")
				createOrdererOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.CreateOrderer(createOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.CreateOrderer(createOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateOrderer(createOrdererOptions *CreateOrdererOptions)`, func() {
		createOrdererPath := "/ak/api/v3/kubernetes/components/fabric-orderer"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createOrdererPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"created": [{"id": "component1", "dep_component_id": "admin", "api_url": "grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050", "display_name": "orderer", "cluster_id": "mzdqhdifnl", "cluster_name": "ordering service 1", "grpcwp_url": "https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443", "location": "ibmcloud", "operations_url": "https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443", "orderer_type": "raft", "config_override": {"anyKey": "anyValue"}, "consenter_proposal_fin": true, "node_ou": {"enabled": true}, "msp": {"ca": {"name": "ca", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "tlsca": {"name": "tlsca", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "component": {"tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "ecert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "admin_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}}, "msp_id": "Org1", "resources": {"orderer": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "proxy": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "storage": {"orderer": {"size": "4GiB", "class": "default"}}, "system_channel_id": "testchainid", "tags": ["fabric-ca"], "timestamp": 1537262855753, "type": "fabric-peer", "version": "1.4.6-1", "zone": "-"}]}`)
				}))
			})
			It(`Invoke CreateOrderer successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.CreateOrderer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CryptoEnrollmentComponent model
				cryptoEnrollmentComponentModel := new(blockchainv3.CryptoEnrollmentComponent)
				cryptoEnrollmentComponentModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the CryptoObjectEnrollmentCa model
				cryptoObjectEnrollmentCaModel := new(blockchainv3.CryptoObjectEnrollmentCa)
				cryptoObjectEnrollmentCaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				cryptoObjectEnrollmentCaModel.Port = core.Float64Ptr(float64(7054))
				cryptoObjectEnrollmentCaModel.Name = core.StringPtr("ca")
				cryptoObjectEnrollmentCaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				cryptoObjectEnrollmentCaModel.EnrollID = core.StringPtr("admin")
				cryptoObjectEnrollmentCaModel.EnrollSecret = core.StringPtr("password")

				// Construct an instance of the CryptoObjectEnrollmentTlsca model
				cryptoObjectEnrollmentTlscaModel := new(blockchainv3.CryptoObjectEnrollmentTlsca)
				cryptoObjectEnrollmentTlscaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				cryptoObjectEnrollmentTlscaModel.Port = core.Float64Ptr(float64(7054))
				cryptoObjectEnrollmentTlscaModel.Name = core.StringPtr("tlsca")
				cryptoObjectEnrollmentTlscaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				cryptoObjectEnrollmentTlscaModel.EnrollID = core.StringPtr("admin")
				cryptoObjectEnrollmentTlscaModel.EnrollSecret = core.StringPtr("password")
				cryptoObjectEnrollmentTlscaModel.CsrHosts = []string{"testString"}

				// Construct an instance of the CryptoObjectEnrollment model
				cryptoObjectEnrollmentModel := new(blockchainv3.CryptoObjectEnrollment)
				cryptoObjectEnrollmentModel.Component = cryptoEnrollmentComponentModel
				cryptoObjectEnrollmentModel.Ca = cryptoObjectEnrollmentCaModel
				cryptoObjectEnrollmentModel.Tlsca = cryptoObjectEnrollmentTlscaModel

				// Construct an instance of the ClientAuth model
				clientAuthModel := new(blockchainv3.ClientAuth)
				clientAuthModel.Type = core.StringPtr("noclientcert")
				clientAuthModel.TlsCerts = []string{"testString"}

				// Construct an instance of the MspCryptoComp model
				mspCryptoCompModel := new(blockchainv3.MspCryptoComp)
				mspCryptoCompModel.Ekey = core.StringPtr("testString")
				mspCryptoCompModel.Ecert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoCompModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				mspCryptoCompModel.TlsKey = core.StringPtr("testString")
				mspCryptoCompModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoCompModel.ClientAuth = clientAuthModel

				// Construct an instance of the MspCryptoCa model
				mspCryptoCaModel := new(blockchainv3.MspCryptoCa)
				mspCryptoCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				mspCryptoCaModel.CaIntermediateCerts = []string{"testString"}

				// Construct an instance of the CryptoObjectMsp model
				cryptoObjectMspModel := new(blockchainv3.CryptoObjectMsp)
				cryptoObjectMspModel.Component = mspCryptoCompModel
				cryptoObjectMspModel.Ca = mspCryptoCaModel
				cryptoObjectMspModel.Tlsca = mspCryptoCaModel

				// Construct an instance of the CryptoObject model
				cryptoObjectModel := new(blockchainv3.CryptoObject)
				cryptoObjectModel.Enrollment = cryptoObjectEnrollmentModel
				cryptoObjectModel.Msp = cryptoObjectMspModel

				// Construct an instance of the ConfigOrdererKeepalive model
				configOrdererKeepaliveModel := new(blockchainv3.ConfigOrdererKeepalive)
				configOrdererKeepaliveModel.ServerMinInterval = core.StringPtr("60s")
				configOrdererKeepaliveModel.ServerInterval = core.StringPtr("2h")
				configOrdererKeepaliveModel.ServerTimeout = core.StringPtr("20s")

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv3.BccspSW)
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv3.BccspPKCS11)
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv3.Bccsp)
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model

				// Construct an instance of the ConfigOrdererAuthentication model
				configOrdererAuthenticationModel := new(blockchainv3.ConfigOrdererAuthentication)
				configOrdererAuthenticationModel.TimeWindow = core.StringPtr("15m")
				configOrdererAuthenticationModel.NoExpirationChecks = core.BoolPtr(false)

				// Construct an instance of the ConfigOrdererGeneral model
				configOrdererGeneralModel := new(blockchainv3.ConfigOrdererGeneral)
				configOrdererGeneralModel.Keepalive = configOrdererKeepaliveModel
				configOrdererGeneralModel.BCCSP = bccspModel
				configOrdererGeneralModel.Authentication = configOrdererAuthenticationModel

				// Construct an instance of the ConfigOrdererDebug model
				configOrdererDebugModel := new(blockchainv3.ConfigOrdererDebug)
				configOrdererDebugModel.BroadcastTraceDir = core.StringPtr("testString")
				configOrdererDebugModel.DeliverTraceDir = core.StringPtr("testString")

				// Construct an instance of the ConfigOrdererMetricsStatsd model
				configOrdererMetricsStatsdModel := new(blockchainv3.ConfigOrdererMetricsStatsd)
				configOrdererMetricsStatsdModel.Network = core.StringPtr("udp")
				configOrdererMetricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				configOrdererMetricsStatsdModel.WriteInterval = core.StringPtr("10s")
				configOrdererMetricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the ConfigOrdererMetrics model
				configOrdererMetricsModel := new(blockchainv3.ConfigOrdererMetrics)
				configOrdererMetricsModel.Provider = core.StringPtr("disabled")
				configOrdererMetricsModel.Statsd = configOrdererMetricsStatsdModel

				// Construct an instance of the ConfigOrdererCreate model
				configOrdererCreateModel := new(blockchainv3.ConfigOrdererCreate)
				configOrdererCreateModel.General = configOrdererGeneralModel
				configOrdererCreateModel.Debug = configOrdererDebugModel
				configOrdererCreateModel.Metrics = configOrdererMetricsModel

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv3.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv3.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv3.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the CreateOrdererRaftBodyResources model
				createOrdererRaftBodyResourcesModel := new(blockchainv3.CreateOrdererRaftBodyResources)
				createOrdererRaftBodyResourcesModel.Orderer = resourceObjectModel
				createOrdererRaftBodyResourcesModel.Proxy = resourceObjectModel

				// Construct an instance of the StorageObject model
				storageObjectModel := new(blockchainv3.StorageObject)
				storageObjectModel.Size = core.StringPtr("4GiB")
				storageObjectModel.Class = core.StringPtr("default")

				// Construct an instance of the CreateOrdererRaftBodyStorage model
				createOrdererRaftBodyStorageModel := new(blockchainv3.CreateOrdererRaftBodyStorage)
				createOrdererRaftBodyStorageModel.Orderer = storageObjectModel

				// Construct an instance of the Hsm model
				hsmModel := new(blockchainv3.Hsm)
				hsmModel.Pkcs11endpoint = core.StringPtr("tcp://example.com:666")

				// Construct an instance of the CreateOrdererOptions model
				createOrdererOptionsModel := new(blockchainv3.CreateOrdererOptions)
				createOrdererOptionsModel.OrdererType = core.StringPtr("raft")
				createOrdererOptionsModel.MspID = core.StringPtr("Org1")
				createOrdererOptionsModel.DisplayName = core.StringPtr("orderer")
				createOrdererOptionsModel.Crypto = []blockchainv3.CryptoObject{*cryptoObjectModel}
				createOrdererOptionsModel.ClusterName = core.StringPtr("ordering service 1")
				createOrdererOptionsModel.ID = core.StringPtr("component1")
				createOrdererOptionsModel.ClusterID = core.StringPtr("abcde")
				createOrdererOptionsModel.ExternalAppend = core.BoolPtr(false)
				createOrdererOptionsModel.ConfigOverride = []blockchainv3.ConfigOrdererCreate{*configOrdererCreateModel}
				createOrdererOptionsModel.Resources = createOrdererRaftBodyResourcesModel
				createOrdererOptionsModel.Storage = createOrdererRaftBodyStorageModel
				createOrdererOptionsModel.SystemChannelID = core.StringPtr("testchainid")
				createOrdererOptionsModel.Zone = []string{"-"}
				createOrdererOptionsModel.Tags = []string{"fabric-ca"}
				createOrdererOptionsModel.Region = []string{"-"}
				createOrdererOptionsModel.Hsm = hsmModel
				createOrdererOptionsModel.Version = core.StringPtr("1.4.6-1")
				createOrdererOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.CreateOrderer(createOrdererOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.CreateOrdererWithContext(ctx, createOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.CreateOrderer(createOrdererOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.CreateOrdererWithContext(ctx, createOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateOrderer with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the CryptoEnrollmentComponent model
				cryptoEnrollmentComponentModel := new(blockchainv3.CryptoEnrollmentComponent)
				cryptoEnrollmentComponentModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the CryptoObjectEnrollmentCa model
				cryptoObjectEnrollmentCaModel := new(blockchainv3.CryptoObjectEnrollmentCa)
				cryptoObjectEnrollmentCaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				cryptoObjectEnrollmentCaModel.Port = core.Float64Ptr(float64(7054))
				cryptoObjectEnrollmentCaModel.Name = core.StringPtr("ca")
				cryptoObjectEnrollmentCaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				cryptoObjectEnrollmentCaModel.EnrollID = core.StringPtr("admin")
				cryptoObjectEnrollmentCaModel.EnrollSecret = core.StringPtr("password")

				// Construct an instance of the CryptoObjectEnrollmentTlsca model
				cryptoObjectEnrollmentTlscaModel := new(blockchainv3.CryptoObjectEnrollmentTlsca)
				cryptoObjectEnrollmentTlscaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				cryptoObjectEnrollmentTlscaModel.Port = core.Float64Ptr(float64(7054))
				cryptoObjectEnrollmentTlscaModel.Name = core.StringPtr("tlsca")
				cryptoObjectEnrollmentTlscaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				cryptoObjectEnrollmentTlscaModel.EnrollID = core.StringPtr("admin")
				cryptoObjectEnrollmentTlscaModel.EnrollSecret = core.StringPtr("password")
				cryptoObjectEnrollmentTlscaModel.CsrHosts = []string{"testString"}

				// Construct an instance of the CryptoObjectEnrollment model
				cryptoObjectEnrollmentModel := new(blockchainv3.CryptoObjectEnrollment)
				cryptoObjectEnrollmentModel.Component = cryptoEnrollmentComponentModel
				cryptoObjectEnrollmentModel.Ca = cryptoObjectEnrollmentCaModel
				cryptoObjectEnrollmentModel.Tlsca = cryptoObjectEnrollmentTlscaModel

				// Construct an instance of the ClientAuth model
				clientAuthModel := new(blockchainv3.ClientAuth)
				clientAuthModel.Type = core.StringPtr("noclientcert")
				clientAuthModel.TlsCerts = []string{"testString"}

				// Construct an instance of the MspCryptoComp model
				mspCryptoCompModel := new(blockchainv3.MspCryptoComp)
				mspCryptoCompModel.Ekey = core.StringPtr("testString")
				mspCryptoCompModel.Ecert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoCompModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				mspCryptoCompModel.TlsKey = core.StringPtr("testString")
				mspCryptoCompModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoCompModel.ClientAuth = clientAuthModel

				// Construct an instance of the MspCryptoCa model
				mspCryptoCaModel := new(blockchainv3.MspCryptoCa)
				mspCryptoCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				mspCryptoCaModel.CaIntermediateCerts = []string{"testString"}

				// Construct an instance of the CryptoObjectMsp model
				cryptoObjectMspModel := new(blockchainv3.CryptoObjectMsp)
				cryptoObjectMspModel.Component = mspCryptoCompModel
				cryptoObjectMspModel.Ca = mspCryptoCaModel
				cryptoObjectMspModel.Tlsca = mspCryptoCaModel

				// Construct an instance of the CryptoObject model
				cryptoObjectModel := new(blockchainv3.CryptoObject)
				cryptoObjectModel.Enrollment = cryptoObjectEnrollmentModel
				cryptoObjectModel.Msp = cryptoObjectMspModel

				// Construct an instance of the ConfigOrdererKeepalive model
				configOrdererKeepaliveModel := new(blockchainv3.ConfigOrdererKeepalive)
				configOrdererKeepaliveModel.ServerMinInterval = core.StringPtr("60s")
				configOrdererKeepaliveModel.ServerInterval = core.StringPtr("2h")
				configOrdererKeepaliveModel.ServerTimeout = core.StringPtr("20s")

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv3.BccspSW)
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv3.BccspPKCS11)
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv3.Bccsp)
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model

				// Construct an instance of the ConfigOrdererAuthentication model
				configOrdererAuthenticationModel := new(blockchainv3.ConfigOrdererAuthentication)
				configOrdererAuthenticationModel.TimeWindow = core.StringPtr("15m")
				configOrdererAuthenticationModel.NoExpirationChecks = core.BoolPtr(false)

				// Construct an instance of the ConfigOrdererGeneral model
				configOrdererGeneralModel := new(blockchainv3.ConfigOrdererGeneral)
				configOrdererGeneralModel.Keepalive = configOrdererKeepaliveModel
				configOrdererGeneralModel.BCCSP = bccspModel
				configOrdererGeneralModel.Authentication = configOrdererAuthenticationModel

				// Construct an instance of the ConfigOrdererDebug model
				configOrdererDebugModel := new(blockchainv3.ConfigOrdererDebug)
				configOrdererDebugModel.BroadcastTraceDir = core.StringPtr("testString")
				configOrdererDebugModel.DeliverTraceDir = core.StringPtr("testString")

				// Construct an instance of the ConfigOrdererMetricsStatsd model
				configOrdererMetricsStatsdModel := new(blockchainv3.ConfigOrdererMetricsStatsd)
				configOrdererMetricsStatsdModel.Network = core.StringPtr("udp")
				configOrdererMetricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				configOrdererMetricsStatsdModel.WriteInterval = core.StringPtr("10s")
				configOrdererMetricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the ConfigOrdererMetrics model
				configOrdererMetricsModel := new(blockchainv3.ConfigOrdererMetrics)
				configOrdererMetricsModel.Provider = core.StringPtr("disabled")
				configOrdererMetricsModel.Statsd = configOrdererMetricsStatsdModel

				// Construct an instance of the ConfigOrdererCreate model
				configOrdererCreateModel := new(blockchainv3.ConfigOrdererCreate)
				configOrdererCreateModel.General = configOrdererGeneralModel
				configOrdererCreateModel.Debug = configOrdererDebugModel
				configOrdererCreateModel.Metrics = configOrdererMetricsModel

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv3.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv3.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv3.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the CreateOrdererRaftBodyResources model
				createOrdererRaftBodyResourcesModel := new(blockchainv3.CreateOrdererRaftBodyResources)
				createOrdererRaftBodyResourcesModel.Orderer = resourceObjectModel
				createOrdererRaftBodyResourcesModel.Proxy = resourceObjectModel

				// Construct an instance of the StorageObject model
				storageObjectModel := new(blockchainv3.StorageObject)
				storageObjectModel.Size = core.StringPtr("4GiB")
				storageObjectModel.Class = core.StringPtr("default")

				// Construct an instance of the CreateOrdererRaftBodyStorage model
				createOrdererRaftBodyStorageModel := new(blockchainv3.CreateOrdererRaftBodyStorage)
				createOrdererRaftBodyStorageModel.Orderer = storageObjectModel

				// Construct an instance of the Hsm model
				hsmModel := new(blockchainv3.Hsm)
				hsmModel.Pkcs11endpoint = core.StringPtr("tcp://example.com:666")

				// Construct an instance of the CreateOrdererOptions model
				createOrdererOptionsModel := new(blockchainv3.CreateOrdererOptions)
				createOrdererOptionsModel.OrdererType = core.StringPtr("raft")
				createOrdererOptionsModel.MspID = core.StringPtr("Org1")
				createOrdererOptionsModel.DisplayName = core.StringPtr("orderer")
				createOrdererOptionsModel.Crypto = []blockchainv3.CryptoObject{*cryptoObjectModel}
				createOrdererOptionsModel.ClusterName = core.StringPtr("ordering service 1")
				createOrdererOptionsModel.ID = core.StringPtr("component1")
				createOrdererOptionsModel.ClusterID = core.StringPtr("abcde")
				createOrdererOptionsModel.ExternalAppend = core.BoolPtr(false)
				createOrdererOptionsModel.ConfigOverride = []blockchainv3.ConfigOrdererCreate{*configOrdererCreateModel}
				createOrdererOptionsModel.Resources = createOrdererRaftBodyResourcesModel
				createOrdererOptionsModel.Storage = createOrdererRaftBodyStorageModel
				createOrdererOptionsModel.SystemChannelID = core.StringPtr("testchainid")
				createOrdererOptionsModel.Zone = []string{"-"}
				createOrdererOptionsModel.Tags = []string{"fabric-ca"}
				createOrdererOptionsModel.Region = []string{"-"}
				createOrdererOptionsModel.Hsm = hsmModel
				createOrdererOptionsModel.Version = core.StringPtr("1.4.6-1")
				createOrdererOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.CreateOrderer(createOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateOrdererOptions model with no property values
				createOrdererOptionsModelNew := new(blockchainv3.CreateOrdererOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.CreateOrderer(createOrdererOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ImportOrderer(importOrdererOptions *ImportOrdererOptions) - Operation response error`, func() {
		importOrdererPath := "/ak/api/v3/components/fabric-orderer"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(importOrdererPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ImportOrderer with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the MspCryptoFieldCa model
				mspCryptoFieldCaModel := new(blockchainv3.MspCryptoFieldCa)
				mspCryptoFieldCaModel.Name = core.StringPtr("ca")
				mspCryptoFieldCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the MspCryptoFieldTlsca model
				mspCryptoFieldTlscaModel := new(blockchainv3.MspCryptoFieldTlsca)
				mspCryptoFieldTlscaModel.Name = core.StringPtr("tlsca")
				mspCryptoFieldTlscaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the MspCryptoFieldComponent model
				mspCryptoFieldComponentModel := new(blockchainv3.MspCryptoFieldComponent)
				mspCryptoFieldComponentModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoFieldComponentModel.Ecert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoFieldComponentModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the MspCryptoField model
				mspCryptoFieldModel := new(blockchainv3.MspCryptoField)
				mspCryptoFieldModel.Ca = mspCryptoFieldCaModel
				mspCryptoFieldModel.Tlsca = mspCryptoFieldTlscaModel
				mspCryptoFieldModel.Component = mspCryptoFieldComponentModel

				// Construct an instance of the ImportOrdererOptions model
				importOrdererOptionsModel := new(blockchainv3.ImportOrdererOptions)
				importOrdererOptionsModel.ClusterName = core.StringPtr("ordering service 1")
				importOrdererOptionsModel.DisplayName = core.StringPtr("orderer")
				importOrdererOptionsModel.GrpcwpURL = core.StringPtr("https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443")
				importOrdererOptionsModel.Msp = mspCryptoFieldModel
				importOrdererOptionsModel.MspID = core.StringPtr("Org1")
				importOrdererOptionsModel.ApiURL = core.StringPtr("grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				importOrdererOptionsModel.ClusterID = core.StringPtr("mzdqhdifnl")
				importOrdererOptionsModel.ID = core.StringPtr("component1")
				importOrdererOptionsModel.Location = core.StringPtr("ibmcloud")
				importOrdererOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443")
				importOrdererOptionsModel.SystemChannelID = core.StringPtr("testchainid")
				importOrdererOptionsModel.Tags = []string{"fabric-ca"}
				importOrdererOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.ImportOrderer(importOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.ImportOrderer(importOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ImportOrderer(importOrdererOptions *ImportOrdererOptions)`, func() {
		importOrdererPath := "/ak/api/v3/components/fabric-orderer"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(importOrdererPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "component1", "dep_component_id": "admin", "api_url": "grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050", "display_name": "orderer", "cluster_id": "mzdqhdifnl", "cluster_name": "ordering service 1", "grpcwp_url": "https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443", "location": "ibmcloud", "operations_url": "https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443", "orderer_type": "raft", "config_override": {"anyKey": "anyValue"}, "consenter_proposal_fin": true, "node_ou": {"enabled": true}, "msp": {"ca": {"name": "ca", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "tlsca": {"name": "tlsca", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "component": {"tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "ecert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "admin_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}}, "msp_id": "Org1", "resources": {"orderer": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "proxy": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "storage": {"orderer": {"size": "4GiB", "class": "default"}}, "system_channel_id": "testchainid", "tags": ["fabric-ca"], "timestamp": 1537262855753, "type": "fabric-peer", "version": "1.4.6-1", "zone": "-"}`)
				}))
			})
			It(`Invoke ImportOrderer successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.ImportOrderer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the MspCryptoFieldCa model
				mspCryptoFieldCaModel := new(blockchainv3.MspCryptoFieldCa)
				mspCryptoFieldCaModel.Name = core.StringPtr("ca")
				mspCryptoFieldCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the MspCryptoFieldTlsca model
				mspCryptoFieldTlscaModel := new(blockchainv3.MspCryptoFieldTlsca)
				mspCryptoFieldTlscaModel.Name = core.StringPtr("tlsca")
				mspCryptoFieldTlscaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the MspCryptoFieldComponent model
				mspCryptoFieldComponentModel := new(blockchainv3.MspCryptoFieldComponent)
				mspCryptoFieldComponentModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoFieldComponentModel.Ecert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoFieldComponentModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the MspCryptoField model
				mspCryptoFieldModel := new(blockchainv3.MspCryptoField)
				mspCryptoFieldModel.Ca = mspCryptoFieldCaModel
				mspCryptoFieldModel.Tlsca = mspCryptoFieldTlscaModel
				mspCryptoFieldModel.Component = mspCryptoFieldComponentModel

				// Construct an instance of the ImportOrdererOptions model
				importOrdererOptionsModel := new(blockchainv3.ImportOrdererOptions)
				importOrdererOptionsModel.ClusterName = core.StringPtr("ordering service 1")
				importOrdererOptionsModel.DisplayName = core.StringPtr("orderer")
				importOrdererOptionsModel.GrpcwpURL = core.StringPtr("https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443")
				importOrdererOptionsModel.Msp = mspCryptoFieldModel
				importOrdererOptionsModel.MspID = core.StringPtr("Org1")
				importOrdererOptionsModel.ApiURL = core.StringPtr("grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				importOrdererOptionsModel.ClusterID = core.StringPtr("mzdqhdifnl")
				importOrdererOptionsModel.ID = core.StringPtr("component1")
				importOrdererOptionsModel.Location = core.StringPtr("ibmcloud")
				importOrdererOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443")
				importOrdererOptionsModel.SystemChannelID = core.StringPtr("testchainid")
				importOrdererOptionsModel.Tags = []string{"fabric-ca"}
				importOrdererOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.ImportOrderer(importOrdererOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.ImportOrdererWithContext(ctx, importOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.ImportOrderer(importOrdererOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.ImportOrdererWithContext(ctx, importOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ImportOrderer with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the MspCryptoFieldCa model
				mspCryptoFieldCaModel := new(blockchainv3.MspCryptoFieldCa)
				mspCryptoFieldCaModel.Name = core.StringPtr("ca")
				mspCryptoFieldCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the MspCryptoFieldTlsca model
				mspCryptoFieldTlscaModel := new(blockchainv3.MspCryptoFieldTlsca)
				mspCryptoFieldTlscaModel.Name = core.StringPtr("tlsca")
				mspCryptoFieldTlscaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the MspCryptoFieldComponent model
				mspCryptoFieldComponentModel := new(blockchainv3.MspCryptoFieldComponent)
				mspCryptoFieldComponentModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoFieldComponentModel.Ecert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoFieldComponentModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the MspCryptoField model
				mspCryptoFieldModel := new(blockchainv3.MspCryptoField)
				mspCryptoFieldModel.Ca = mspCryptoFieldCaModel
				mspCryptoFieldModel.Tlsca = mspCryptoFieldTlscaModel
				mspCryptoFieldModel.Component = mspCryptoFieldComponentModel

				// Construct an instance of the ImportOrdererOptions model
				importOrdererOptionsModel := new(blockchainv3.ImportOrdererOptions)
				importOrdererOptionsModel.ClusterName = core.StringPtr("ordering service 1")
				importOrdererOptionsModel.DisplayName = core.StringPtr("orderer")
				importOrdererOptionsModel.GrpcwpURL = core.StringPtr("https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443")
				importOrdererOptionsModel.Msp = mspCryptoFieldModel
				importOrdererOptionsModel.MspID = core.StringPtr("Org1")
				importOrdererOptionsModel.ApiURL = core.StringPtr("grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				importOrdererOptionsModel.ClusterID = core.StringPtr("mzdqhdifnl")
				importOrdererOptionsModel.ID = core.StringPtr("component1")
				importOrdererOptionsModel.Location = core.StringPtr("ibmcloud")
				importOrdererOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443")
				importOrdererOptionsModel.SystemChannelID = core.StringPtr("testchainid")
				importOrdererOptionsModel.Tags = []string{"fabric-ca"}
				importOrdererOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.ImportOrderer(importOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ImportOrdererOptions model with no property values
				importOrdererOptionsModelNew := new(blockchainv3.ImportOrdererOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.ImportOrderer(importOrdererOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`EditOrderer(editOrdererOptions *EditOrdererOptions) - Operation response error`, func() {
		editOrdererPath := "/ak/api/v3/components/fabric-orderer/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(editOrdererPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke EditOrderer with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the EditOrdererOptions model
				editOrdererOptionsModel := new(blockchainv3.EditOrdererOptions)
				editOrdererOptionsModel.ID = core.StringPtr("testString")
				editOrdererOptionsModel.ClusterName = core.StringPtr("ordering service 1")
				editOrdererOptionsModel.DisplayName = core.StringPtr("orderer")
				editOrdererOptionsModel.ApiURL = core.StringPtr("grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				editOrdererOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443")
				editOrdererOptionsModel.GrpcwpURL = core.StringPtr("https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443")
				editOrdererOptionsModel.MspID = core.StringPtr("Org1")
				editOrdererOptionsModel.ConsenterProposalFin = core.BoolPtr(true)
				editOrdererOptionsModel.Location = core.StringPtr("ibmcloud")
				editOrdererOptionsModel.SystemChannelID = core.StringPtr("testchainid")
				editOrdererOptionsModel.Tags = []string{"fabric-ca"}
				editOrdererOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.EditOrderer(editOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.EditOrderer(editOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`EditOrderer(editOrdererOptions *EditOrdererOptions)`, func() {
		editOrdererPath := "/ak/api/v3/components/fabric-orderer/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(editOrdererPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "component1", "dep_component_id": "admin", "api_url": "grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050", "display_name": "orderer", "cluster_id": "mzdqhdifnl", "cluster_name": "ordering service 1", "grpcwp_url": "https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443", "location": "ibmcloud", "operations_url": "https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443", "orderer_type": "raft", "config_override": {"anyKey": "anyValue"}, "consenter_proposal_fin": true, "node_ou": {"enabled": true}, "msp": {"ca": {"name": "ca", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "tlsca": {"name": "tlsca", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "component": {"tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "ecert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "admin_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}}, "msp_id": "Org1", "resources": {"orderer": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "proxy": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "storage": {"orderer": {"size": "4GiB", "class": "default"}}, "system_channel_id": "testchainid", "tags": ["fabric-ca"], "timestamp": 1537262855753, "type": "fabric-peer", "version": "1.4.6-1", "zone": "-"}`)
				}))
			})
			It(`Invoke EditOrderer successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.EditOrderer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EditOrdererOptions model
				editOrdererOptionsModel := new(blockchainv3.EditOrdererOptions)
				editOrdererOptionsModel.ID = core.StringPtr("testString")
				editOrdererOptionsModel.ClusterName = core.StringPtr("ordering service 1")
				editOrdererOptionsModel.DisplayName = core.StringPtr("orderer")
				editOrdererOptionsModel.ApiURL = core.StringPtr("grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				editOrdererOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443")
				editOrdererOptionsModel.GrpcwpURL = core.StringPtr("https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443")
				editOrdererOptionsModel.MspID = core.StringPtr("Org1")
				editOrdererOptionsModel.ConsenterProposalFin = core.BoolPtr(true)
				editOrdererOptionsModel.Location = core.StringPtr("ibmcloud")
				editOrdererOptionsModel.SystemChannelID = core.StringPtr("testchainid")
				editOrdererOptionsModel.Tags = []string{"fabric-ca"}
				editOrdererOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.EditOrderer(editOrdererOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.EditOrdererWithContext(ctx, editOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.EditOrderer(editOrdererOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.EditOrdererWithContext(ctx, editOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke EditOrderer with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the EditOrdererOptions model
				editOrdererOptionsModel := new(blockchainv3.EditOrdererOptions)
				editOrdererOptionsModel.ID = core.StringPtr("testString")
				editOrdererOptionsModel.ClusterName = core.StringPtr("ordering service 1")
				editOrdererOptionsModel.DisplayName = core.StringPtr("orderer")
				editOrdererOptionsModel.ApiURL = core.StringPtr("grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				editOrdererOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443")
				editOrdererOptionsModel.GrpcwpURL = core.StringPtr("https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443")
				editOrdererOptionsModel.MspID = core.StringPtr("Org1")
				editOrdererOptionsModel.ConsenterProposalFin = core.BoolPtr(true)
				editOrdererOptionsModel.Location = core.StringPtr("ibmcloud")
				editOrdererOptionsModel.SystemChannelID = core.StringPtr("testchainid")
				editOrdererOptionsModel.Tags = []string{"fabric-ca"}
				editOrdererOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.EditOrderer(editOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the EditOrdererOptions model with no property values
				editOrdererOptionsModelNew := new(blockchainv3.EditOrdererOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.EditOrderer(editOrdererOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`OrdererAction(ordererActionOptions *OrdererActionOptions) - Operation response error`, func() {
		ordererActionPath := "/ak/api/v3/kubernetes/components/fabric-orderer/testString/actions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(ordererActionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke OrdererAction with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the ActionReenroll model
				actionReenrollModel := new(blockchainv3.ActionReenroll)
				actionReenrollModel.TlsCert = core.BoolPtr(true)
				actionReenrollModel.Ecert = core.BoolPtr(true)

				// Construct an instance of the ActionEnroll model
				actionEnrollModel := new(blockchainv3.ActionEnroll)
				actionEnrollModel.TlsCert = core.BoolPtr(true)
				actionEnrollModel.Ecert = core.BoolPtr(true)

				// Construct an instance of the OrdererActionOptions model
				ordererActionOptionsModel := new(blockchainv3.OrdererActionOptions)
				ordererActionOptionsModel.ID = core.StringPtr("testString")
				ordererActionOptionsModel.Restart = core.BoolPtr(true)
				ordererActionOptionsModel.Reenroll = actionReenrollModel
				ordererActionOptionsModel.Enroll = actionEnrollModel
				ordererActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.OrdererAction(ordererActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.OrdererAction(ordererActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`OrdererAction(ordererActionOptions *OrdererActionOptions)`, func() {
		ordererActionPath := "/ak/api/v3/kubernetes/components/fabric-orderer/testString/actions"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(ordererActionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"message": "accepted", "id": "myca", "actions": ["restart"]}`)
				}))
			})
			It(`Invoke OrdererAction successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.OrdererAction(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ActionReenroll model
				actionReenrollModel := new(blockchainv3.ActionReenroll)
				actionReenrollModel.TlsCert = core.BoolPtr(true)
				actionReenrollModel.Ecert = core.BoolPtr(true)

				// Construct an instance of the ActionEnroll model
				actionEnrollModel := new(blockchainv3.ActionEnroll)
				actionEnrollModel.TlsCert = core.BoolPtr(true)
				actionEnrollModel.Ecert = core.BoolPtr(true)

				// Construct an instance of the OrdererActionOptions model
				ordererActionOptionsModel := new(blockchainv3.OrdererActionOptions)
				ordererActionOptionsModel.ID = core.StringPtr("testString")
				ordererActionOptionsModel.Restart = core.BoolPtr(true)
				ordererActionOptionsModel.Reenroll = actionReenrollModel
				ordererActionOptionsModel.Enroll = actionEnrollModel
				ordererActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.OrdererAction(ordererActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.OrdererActionWithContext(ctx, ordererActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.OrdererAction(ordererActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.OrdererActionWithContext(ctx, ordererActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke OrdererAction with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the ActionReenroll model
				actionReenrollModel := new(blockchainv3.ActionReenroll)
				actionReenrollModel.TlsCert = core.BoolPtr(true)
				actionReenrollModel.Ecert = core.BoolPtr(true)

				// Construct an instance of the ActionEnroll model
				actionEnrollModel := new(blockchainv3.ActionEnroll)
				actionEnrollModel.TlsCert = core.BoolPtr(true)
				actionEnrollModel.Ecert = core.BoolPtr(true)

				// Construct an instance of the OrdererActionOptions model
				ordererActionOptionsModel := new(blockchainv3.OrdererActionOptions)
				ordererActionOptionsModel.ID = core.StringPtr("testString")
				ordererActionOptionsModel.Restart = core.BoolPtr(true)
				ordererActionOptionsModel.Reenroll = actionReenrollModel
				ordererActionOptionsModel.Enroll = actionEnrollModel
				ordererActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.OrdererAction(ordererActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the OrdererActionOptions model with no property values
				ordererActionOptionsModelNew := new(blockchainv3.OrdererActionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.OrdererAction(ordererActionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateOrderer(updateOrdererOptions *UpdateOrdererOptions) - Operation response error`, func() {
		updateOrdererPath := "/ak/api/v3/kubernetes/components/fabric-orderer/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateOrdererPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateOrderer with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the ConfigOrdererKeepalive model
				configOrdererKeepaliveModel := new(blockchainv3.ConfigOrdererKeepalive)
				configOrdererKeepaliveModel.ServerMinInterval = core.StringPtr("60s")
				configOrdererKeepaliveModel.ServerInterval = core.StringPtr("2h")
				configOrdererKeepaliveModel.ServerTimeout = core.StringPtr("20s")

				// Construct an instance of the ConfigOrdererAuthentication model
				configOrdererAuthenticationModel := new(blockchainv3.ConfigOrdererAuthentication)
				configOrdererAuthenticationModel.TimeWindow = core.StringPtr("15m")
				configOrdererAuthenticationModel.NoExpirationChecks = core.BoolPtr(false)

				// Construct an instance of the ConfigOrdererGeneralUpdate model
				configOrdererGeneralUpdateModel := new(blockchainv3.ConfigOrdererGeneralUpdate)
				configOrdererGeneralUpdateModel.Keepalive = configOrdererKeepaliveModel
				configOrdererGeneralUpdateModel.Authentication = configOrdererAuthenticationModel

				// Construct an instance of the ConfigOrdererDebug model
				configOrdererDebugModel := new(blockchainv3.ConfigOrdererDebug)
				configOrdererDebugModel.BroadcastTraceDir = core.StringPtr("testString")
				configOrdererDebugModel.DeliverTraceDir = core.StringPtr("testString")

				// Construct an instance of the ConfigOrdererMetricsStatsd model
				configOrdererMetricsStatsdModel := new(blockchainv3.ConfigOrdererMetricsStatsd)
				configOrdererMetricsStatsdModel.Network = core.StringPtr("udp")
				configOrdererMetricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				configOrdererMetricsStatsdModel.WriteInterval = core.StringPtr("10s")
				configOrdererMetricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the ConfigOrdererMetrics model
				configOrdererMetricsModel := new(blockchainv3.ConfigOrdererMetrics)
				configOrdererMetricsModel.Provider = core.StringPtr("disabled")
				configOrdererMetricsModel.Statsd = configOrdererMetricsStatsdModel

				// Construct an instance of the ConfigOrdererUpdate model
				configOrdererUpdateModel := new(blockchainv3.ConfigOrdererUpdate)
				configOrdererUpdateModel.General = configOrdererGeneralUpdateModel
				configOrdererUpdateModel.Debug = configOrdererDebugModel
				configOrdererUpdateModel.Metrics = configOrdererMetricsModel

				// Construct an instance of the CryptoEnrollmentComponent model
				cryptoEnrollmentComponentModel := new(blockchainv3.CryptoEnrollmentComponent)
				cryptoEnrollmentComponentModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the UpdateEnrollmentCryptoFieldCa model
				updateEnrollmentCryptoFieldCaModel := new(blockchainv3.UpdateEnrollmentCryptoFieldCa)
				updateEnrollmentCryptoFieldCaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				updateEnrollmentCryptoFieldCaModel.Port = core.Float64Ptr(float64(7054))
				updateEnrollmentCryptoFieldCaModel.Name = core.StringPtr("ca")
				updateEnrollmentCryptoFieldCaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateEnrollmentCryptoFieldCaModel.EnrollID = core.StringPtr("admin")
				updateEnrollmentCryptoFieldCaModel.EnrollSecret = core.StringPtr("password")

				// Construct an instance of the UpdateEnrollmentCryptoFieldTlsca model
				updateEnrollmentCryptoFieldTlscaModel := new(blockchainv3.UpdateEnrollmentCryptoFieldTlsca)
				updateEnrollmentCryptoFieldTlscaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				updateEnrollmentCryptoFieldTlscaModel.Port = core.Float64Ptr(float64(7054))
				updateEnrollmentCryptoFieldTlscaModel.Name = core.StringPtr("tlsca")
				updateEnrollmentCryptoFieldTlscaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateEnrollmentCryptoFieldTlscaModel.EnrollID = core.StringPtr("admin")
				updateEnrollmentCryptoFieldTlscaModel.EnrollSecret = core.StringPtr("password")
				updateEnrollmentCryptoFieldTlscaModel.CsrHosts = []string{"testString"}

				// Construct an instance of the UpdateEnrollmentCryptoField model
				updateEnrollmentCryptoFieldModel := new(blockchainv3.UpdateEnrollmentCryptoField)
				updateEnrollmentCryptoFieldModel.Component = cryptoEnrollmentComponentModel
				updateEnrollmentCryptoFieldModel.Ca = updateEnrollmentCryptoFieldCaModel
				updateEnrollmentCryptoFieldModel.Tlsca = updateEnrollmentCryptoFieldTlscaModel

				// Construct an instance of the UpdateMspCryptoFieldCa model
				updateMspCryptoFieldCaModel := new(blockchainv3.UpdateMspCryptoFieldCa)
				updateMspCryptoFieldCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateMspCryptoFieldCaModel.CaIntermediateCerts = []string{"testString"}

				// Construct an instance of the UpdateMspCryptoFieldTlsca model
				updateMspCryptoFieldTlscaModel := new(blockchainv3.UpdateMspCryptoFieldTlsca)
				updateMspCryptoFieldTlscaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateMspCryptoFieldTlscaModel.CaIntermediateCerts = []string{"testString"}

				// Construct an instance of the ClientAuth model
				clientAuthModel := new(blockchainv3.ClientAuth)
				clientAuthModel.Type = core.StringPtr("noclientcert")
				clientAuthModel.TlsCerts = []string{"testString"}

				// Construct an instance of the UpdateMspCryptoFieldComponent model
				updateMspCryptoFieldComponentModel := new(blockchainv3.UpdateMspCryptoFieldComponent)
				updateMspCryptoFieldComponentModel.Ekey = core.StringPtr("testString")
				updateMspCryptoFieldComponentModel.Ecert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateMspCryptoFieldComponentModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateMspCryptoFieldComponentModel.TlsKey = core.StringPtr("testString")
				updateMspCryptoFieldComponentModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateMspCryptoFieldComponentModel.ClientAuth = clientAuthModel

				// Construct an instance of the UpdateMspCryptoField model
				updateMspCryptoFieldModel := new(blockchainv3.UpdateMspCryptoField)
				updateMspCryptoFieldModel.Ca = updateMspCryptoFieldCaModel
				updateMspCryptoFieldModel.Tlsca = updateMspCryptoFieldTlscaModel
				updateMspCryptoFieldModel.Component = updateMspCryptoFieldComponentModel

				// Construct an instance of the UpdateOrdererBodyCrypto model
				updateOrdererBodyCryptoModel := new(blockchainv3.UpdateOrdererBodyCrypto)
				updateOrdererBodyCryptoModel.Enrollment = updateEnrollmentCryptoFieldModel
				updateOrdererBodyCryptoModel.Msp = updateMspCryptoFieldModel

				// Construct an instance of the NodeOu model
				nodeOuModel := new(blockchainv3.NodeOu)
				nodeOuModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv3.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv3.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv3.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the UpdateOrdererBodyResources model
				updateOrdererBodyResourcesModel := new(blockchainv3.UpdateOrdererBodyResources)
				updateOrdererBodyResourcesModel.Orderer = resourceObjectModel
				updateOrdererBodyResourcesModel.Proxy = resourceObjectModel

				// Construct an instance of the UpdateOrdererOptions model
				updateOrdererOptionsModel := new(blockchainv3.UpdateOrdererOptions)
				updateOrdererOptionsModel.ID = core.StringPtr("testString")
				updateOrdererOptionsModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateOrdererOptionsModel.ConfigOverride = configOrdererUpdateModel
				updateOrdererOptionsModel.Crypto = updateOrdererBodyCryptoModel
				updateOrdererOptionsModel.NodeOu = nodeOuModel
				updateOrdererOptionsModel.Replicas = core.Float64Ptr(float64(1))
				updateOrdererOptionsModel.Resources = updateOrdererBodyResourcesModel
				updateOrdererOptionsModel.Version = core.StringPtr("1.4.6-1")
				updateOrdererOptionsModel.Zone = core.StringPtr("-")
				updateOrdererOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.UpdateOrderer(updateOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.UpdateOrderer(updateOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateOrderer(updateOrdererOptions *UpdateOrdererOptions)`, func() {
		updateOrdererPath := "/ak/api/v3/kubernetes/components/fabric-orderer/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateOrdererPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "component1", "dep_component_id": "admin", "api_url": "grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050", "display_name": "orderer", "cluster_id": "mzdqhdifnl", "cluster_name": "ordering service 1", "grpcwp_url": "https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443", "location": "ibmcloud", "operations_url": "https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443", "orderer_type": "raft", "config_override": {"anyKey": "anyValue"}, "consenter_proposal_fin": true, "node_ou": {"enabled": true}, "msp": {"ca": {"name": "ca", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "tlsca": {"name": "tlsca", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "component": {"tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "ecert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "admin_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}}, "msp_id": "Org1", "resources": {"orderer": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "proxy": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "storage": {"orderer": {"size": "4GiB", "class": "default"}}, "system_channel_id": "testchainid", "tags": ["fabric-ca"], "timestamp": 1537262855753, "type": "fabric-peer", "version": "1.4.6-1", "zone": "-"}`)
				}))
			})
			It(`Invoke UpdateOrderer successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.UpdateOrderer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ConfigOrdererKeepalive model
				configOrdererKeepaliveModel := new(blockchainv3.ConfigOrdererKeepalive)
				configOrdererKeepaliveModel.ServerMinInterval = core.StringPtr("60s")
				configOrdererKeepaliveModel.ServerInterval = core.StringPtr("2h")
				configOrdererKeepaliveModel.ServerTimeout = core.StringPtr("20s")

				// Construct an instance of the ConfigOrdererAuthentication model
				configOrdererAuthenticationModel := new(blockchainv3.ConfigOrdererAuthentication)
				configOrdererAuthenticationModel.TimeWindow = core.StringPtr("15m")
				configOrdererAuthenticationModel.NoExpirationChecks = core.BoolPtr(false)

				// Construct an instance of the ConfigOrdererGeneralUpdate model
				configOrdererGeneralUpdateModel := new(blockchainv3.ConfigOrdererGeneralUpdate)
				configOrdererGeneralUpdateModel.Keepalive = configOrdererKeepaliveModel
				configOrdererGeneralUpdateModel.Authentication = configOrdererAuthenticationModel

				// Construct an instance of the ConfigOrdererDebug model
				configOrdererDebugModel := new(blockchainv3.ConfigOrdererDebug)
				configOrdererDebugModel.BroadcastTraceDir = core.StringPtr("testString")
				configOrdererDebugModel.DeliverTraceDir = core.StringPtr("testString")

				// Construct an instance of the ConfigOrdererMetricsStatsd model
				configOrdererMetricsStatsdModel := new(blockchainv3.ConfigOrdererMetricsStatsd)
				configOrdererMetricsStatsdModel.Network = core.StringPtr("udp")
				configOrdererMetricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				configOrdererMetricsStatsdModel.WriteInterval = core.StringPtr("10s")
				configOrdererMetricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the ConfigOrdererMetrics model
				configOrdererMetricsModel := new(blockchainv3.ConfigOrdererMetrics)
				configOrdererMetricsModel.Provider = core.StringPtr("disabled")
				configOrdererMetricsModel.Statsd = configOrdererMetricsStatsdModel

				// Construct an instance of the ConfigOrdererUpdate model
				configOrdererUpdateModel := new(blockchainv3.ConfigOrdererUpdate)
				configOrdererUpdateModel.General = configOrdererGeneralUpdateModel
				configOrdererUpdateModel.Debug = configOrdererDebugModel
				configOrdererUpdateModel.Metrics = configOrdererMetricsModel

				// Construct an instance of the CryptoEnrollmentComponent model
				cryptoEnrollmentComponentModel := new(blockchainv3.CryptoEnrollmentComponent)
				cryptoEnrollmentComponentModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the UpdateEnrollmentCryptoFieldCa model
				updateEnrollmentCryptoFieldCaModel := new(blockchainv3.UpdateEnrollmentCryptoFieldCa)
				updateEnrollmentCryptoFieldCaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				updateEnrollmentCryptoFieldCaModel.Port = core.Float64Ptr(float64(7054))
				updateEnrollmentCryptoFieldCaModel.Name = core.StringPtr("ca")
				updateEnrollmentCryptoFieldCaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateEnrollmentCryptoFieldCaModel.EnrollID = core.StringPtr("admin")
				updateEnrollmentCryptoFieldCaModel.EnrollSecret = core.StringPtr("password")

				// Construct an instance of the UpdateEnrollmentCryptoFieldTlsca model
				updateEnrollmentCryptoFieldTlscaModel := new(blockchainv3.UpdateEnrollmentCryptoFieldTlsca)
				updateEnrollmentCryptoFieldTlscaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				updateEnrollmentCryptoFieldTlscaModel.Port = core.Float64Ptr(float64(7054))
				updateEnrollmentCryptoFieldTlscaModel.Name = core.StringPtr("tlsca")
				updateEnrollmentCryptoFieldTlscaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateEnrollmentCryptoFieldTlscaModel.EnrollID = core.StringPtr("admin")
				updateEnrollmentCryptoFieldTlscaModel.EnrollSecret = core.StringPtr("password")
				updateEnrollmentCryptoFieldTlscaModel.CsrHosts = []string{"testString"}

				// Construct an instance of the UpdateEnrollmentCryptoField model
				updateEnrollmentCryptoFieldModel := new(blockchainv3.UpdateEnrollmentCryptoField)
				updateEnrollmentCryptoFieldModel.Component = cryptoEnrollmentComponentModel
				updateEnrollmentCryptoFieldModel.Ca = updateEnrollmentCryptoFieldCaModel
				updateEnrollmentCryptoFieldModel.Tlsca = updateEnrollmentCryptoFieldTlscaModel

				// Construct an instance of the UpdateMspCryptoFieldCa model
				updateMspCryptoFieldCaModel := new(blockchainv3.UpdateMspCryptoFieldCa)
				updateMspCryptoFieldCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateMspCryptoFieldCaModel.CaIntermediateCerts = []string{"testString"}

				// Construct an instance of the UpdateMspCryptoFieldTlsca model
				updateMspCryptoFieldTlscaModel := new(blockchainv3.UpdateMspCryptoFieldTlsca)
				updateMspCryptoFieldTlscaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateMspCryptoFieldTlscaModel.CaIntermediateCerts = []string{"testString"}

				// Construct an instance of the ClientAuth model
				clientAuthModel := new(blockchainv3.ClientAuth)
				clientAuthModel.Type = core.StringPtr("noclientcert")
				clientAuthModel.TlsCerts = []string{"testString"}

				// Construct an instance of the UpdateMspCryptoFieldComponent model
				updateMspCryptoFieldComponentModel := new(blockchainv3.UpdateMspCryptoFieldComponent)
				updateMspCryptoFieldComponentModel.Ekey = core.StringPtr("testString")
				updateMspCryptoFieldComponentModel.Ecert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateMspCryptoFieldComponentModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateMspCryptoFieldComponentModel.TlsKey = core.StringPtr("testString")
				updateMspCryptoFieldComponentModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateMspCryptoFieldComponentModel.ClientAuth = clientAuthModel

				// Construct an instance of the UpdateMspCryptoField model
				updateMspCryptoFieldModel := new(blockchainv3.UpdateMspCryptoField)
				updateMspCryptoFieldModel.Ca = updateMspCryptoFieldCaModel
				updateMspCryptoFieldModel.Tlsca = updateMspCryptoFieldTlscaModel
				updateMspCryptoFieldModel.Component = updateMspCryptoFieldComponentModel

				// Construct an instance of the UpdateOrdererBodyCrypto model
				updateOrdererBodyCryptoModel := new(blockchainv3.UpdateOrdererBodyCrypto)
				updateOrdererBodyCryptoModel.Enrollment = updateEnrollmentCryptoFieldModel
				updateOrdererBodyCryptoModel.Msp = updateMspCryptoFieldModel

				// Construct an instance of the NodeOu model
				nodeOuModel := new(blockchainv3.NodeOu)
				nodeOuModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv3.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv3.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv3.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the UpdateOrdererBodyResources model
				updateOrdererBodyResourcesModel := new(blockchainv3.UpdateOrdererBodyResources)
				updateOrdererBodyResourcesModel.Orderer = resourceObjectModel
				updateOrdererBodyResourcesModel.Proxy = resourceObjectModel

				// Construct an instance of the UpdateOrdererOptions model
				updateOrdererOptionsModel := new(blockchainv3.UpdateOrdererOptions)
				updateOrdererOptionsModel.ID = core.StringPtr("testString")
				updateOrdererOptionsModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateOrdererOptionsModel.ConfigOverride = configOrdererUpdateModel
				updateOrdererOptionsModel.Crypto = updateOrdererBodyCryptoModel
				updateOrdererOptionsModel.NodeOu = nodeOuModel
				updateOrdererOptionsModel.Replicas = core.Float64Ptr(float64(1))
				updateOrdererOptionsModel.Resources = updateOrdererBodyResourcesModel
				updateOrdererOptionsModel.Version = core.StringPtr("1.4.6-1")
				updateOrdererOptionsModel.Zone = core.StringPtr("-")
				updateOrdererOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.UpdateOrderer(updateOrdererOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.UpdateOrdererWithContext(ctx, updateOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.UpdateOrderer(updateOrdererOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.UpdateOrdererWithContext(ctx, updateOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateOrderer with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the ConfigOrdererKeepalive model
				configOrdererKeepaliveModel := new(blockchainv3.ConfigOrdererKeepalive)
				configOrdererKeepaliveModel.ServerMinInterval = core.StringPtr("60s")
				configOrdererKeepaliveModel.ServerInterval = core.StringPtr("2h")
				configOrdererKeepaliveModel.ServerTimeout = core.StringPtr("20s")

				// Construct an instance of the ConfigOrdererAuthentication model
				configOrdererAuthenticationModel := new(blockchainv3.ConfigOrdererAuthentication)
				configOrdererAuthenticationModel.TimeWindow = core.StringPtr("15m")
				configOrdererAuthenticationModel.NoExpirationChecks = core.BoolPtr(false)

				// Construct an instance of the ConfigOrdererGeneralUpdate model
				configOrdererGeneralUpdateModel := new(blockchainv3.ConfigOrdererGeneralUpdate)
				configOrdererGeneralUpdateModel.Keepalive = configOrdererKeepaliveModel
				configOrdererGeneralUpdateModel.Authentication = configOrdererAuthenticationModel

				// Construct an instance of the ConfigOrdererDebug model
				configOrdererDebugModel := new(blockchainv3.ConfigOrdererDebug)
				configOrdererDebugModel.BroadcastTraceDir = core.StringPtr("testString")
				configOrdererDebugModel.DeliverTraceDir = core.StringPtr("testString")

				// Construct an instance of the ConfigOrdererMetricsStatsd model
				configOrdererMetricsStatsdModel := new(blockchainv3.ConfigOrdererMetricsStatsd)
				configOrdererMetricsStatsdModel.Network = core.StringPtr("udp")
				configOrdererMetricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				configOrdererMetricsStatsdModel.WriteInterval = core.StringPtr("10s")
				configOrdererMetricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the ConfigOrdererMetrics model
				configOrdererMetricsModel := new(blockchainv3.ConfigOrdererMetrics)
				configOrdererMetricsModel.Provider = core.StringPtr("disabled")
				configOrdererMetricsModel.Statsd = configOrdererMetricsStatsdModel

				// Construct an instance of the ConfigOrdererUpdate model
				configOrdererUpdateModel := new(blockchainv3.ConfigOrdererUpdate)
				configOrdererUpdateModel.General = configOrdererGeneralUpdateModel
				configOrdererUpdateModel.Debug = configOrdererDebugModel
				configOrdererUpdateModel.Metrics = configOrdererMetricsModel

				// Construct an instance of the CryptoEnrollmentComponent model
				cryptoEnrollmentComponentModel := new(blockchainv3.CryptoEnrollmentComponent)
				cryptoEnrollmentComponentModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the UpdateEnrollmentCryptoFieldCa model
				updateEnrollmentCryptoFieldCaModel := new(blockchainv3.UpdateEnrollmentCryptoFieldCa)
				updateEnrollmentCryptoFieldCaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				updateEnrollmentCryptoFieldCaModel.Port = core.Float64Ptr(float64(7054))
				updateEnrollmentCryptoFieldCaModel.Name = core.StringPtr("ca")
				updateEnrollmentCryptoFieldCaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateEnrollmentCryptoFieldCaModel.EnrollID = core.StringPtr("admin")
				updateEnrollmentCryptoFieldCaModel.EnrollSecret = core.StringPtr("password")

				// Construct an instance of the UpdateEnrollmentCryptoFieldTlsca model
				updateEnrollmentCryptoFieldTlscaModel := new(blockchainv3.UpdateEnrollmentCryptoFieldTlsca)
				updateEnrollmentCryptoFieldTlscaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				updateEnrollmentCryptoFieldTlscaModel.Port = core.Float64Ptr(float64(7054))
				updateEnrollmentCryptoFieldTlscaModel.Name = core.StringPtr("tlsca")
				updateEnrollmentCryptoFieldTlscaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateEnrollmentCryptoFieldTlscaModel.EnrollID = core.StringPtr("admin")
				updateEnrollmentCryptoFieldTlscaModel.EnrollSecret = core.StringPtr("password")
				updateEnrollmentCryptoFieldTlscaModel.CsrHosts = []string{"testString"}

				// Construct an instance of the UpdateEnrollmentCryptoField model
				updateEnrollmentCryptoFieldModel := new(blockchainv3.UpdateEnrollmentCryptoField)
				updateEnrollmentCryptoFieldModel.Component = cryptoEnrollmentComponentModel
				updateEnrollmentCryptoFieldModel.Ca = updateEnrollmentCryptoFieldCaModel
				updateEnrollmentCryptoFieldModel.Tlsca = updateEnrollmentCryptoFieldTlscaModel

				// Construct an instance of the UpdateMspCryptoFieldCa model
				updateMspCryptoFieldCaModel := new(blockchainv3.UpdateMspCryptoFieldCa)
				updateMspCryptoFieldCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateMspCryptoFieldCaModel.CaIntermediateCerts = []string{"testString"}

				// Construct an instance of the UpdateMspCryptoFieldTlsca model
				updateMspCryptoFieldTlscaModel := new(blockchainv3.UpdateMspCryptoFieldTlsca)
				updateMspCryptoFieldTlscaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateMspCryptoFieldTlscaModel.CaIntermediateCerts = []string{"testString"}

				// Construct an instance of the ClientAuth model
				clientAuthModel := new(blockchainv3.ClientAuth)
				clientAuthModel.Type = core.StringPtr("noclientcert")
				clientAuthModel.TlsCerts = []string{"testString"}

				// Construct an instance of the UpdateMspCryptoFieldComponent model
				updateMspCryptoFieldComponentModel := new(blockchainv3.UpdateMspCryptoFieldComponent)
				updateMspCryptoFieldComponentModel.Ekey = core.StringPtr("testString")
				updateMspCryptoFieldComponentModel.Ecert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateMspCryptoFieldComponentModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateMspCryptoFieldComponentModel.TlsKey = core.StringPtr("testString")
				updateMspCryptoFieldComponentModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateMspCryptoFieldComponentModel.ClientAuth = clientAuthModel

				// Construct an instance of the UpdateMspCryptoField model
				updateMspCryptoFieldModel := new(blockchainv3.UpdateMspCryptoField)
				updateMspCryptoFieldModel.Ca = updateMspCryptoFieldCaModel
				updateMspCryptoFieldModel.Tlsca = updateMspCryptoFieldTlscaModel
				updateMspCryptoFieldModel.Component = updateMspCryptoFieldComponentModel

				// Construct an instance of the UpdateOrdererBodyCrypto model
				updateOrdererBodyCryptoModel := new(blockchainv3.UpdateOrdererBodyCrypto)
				updateOrdererBodyCryptoModel.Enrollment = updateEnrollmentCryptoFieldModel
				updateOrdererBodyCryptoModel.Msp = updateMspCryptoFieldModel

				// Construct an instance of the NodeOu model
				nodeOuModel := new(blockchainv3.NodeOu)
				nodeOuModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv3.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv3.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv3.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the UpdateOrdererBodyResources model
				updateOrdererBodyResourcesModel := new(blockchainv3.UpdateOrdererBodyResources)
				updateOrdererBodyResourcesModel.Orderer = resourceObjectModel
				updateOrdererBodyResourcesModel.Proxy = resourceObjectModel

				// Construct an instance of the UpdateOrdererOptions model
				updateOrdererOptionsModel := new(blockchainv3.UpdateOrdererOptions)
				updateOrdererOptionsModel.ID = core.StringPtr("testString")
				updateOrdererOptionsModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateOrdererOptionsModel.ConfigOverride = configOrdererUpdateModel
				updateOrdererOptionsModel.Crypto = updateOrdererBodyCryptoModel
				updateOrdererOptionsModel.NodeOu = nodeOuModel
				updateOrdererOptionsModel.Replicas = core.Float64Ptr(float64(1))
				updateOrdererOptionsModel.Resources = updateOrdererBodyResourcesModel
				updateOrdererOptionsModel.Version = core.StringPtr("1.4.6-1")
				updateOrdererOptionsModel.Zone = core.StringPtr("-")
				updateOrdererOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.UpdateOrderer(updateOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateOrdererOptions model with no property values
				updateOrdererOptionsModelNew := new(blockchainv3.UpdateOrdererOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.UpdateOrderer(updateOrdererOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SubmitBlock(submitBlockOptions *SubmitBlockOptions) - Operation response error`, func() {
		submitBlockPath := "/ak/api/v3/kubernetes/components/testString/config"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(submitBlockPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SubmitBlock with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the SubmitBlockOptions model
				submitBlockOptionsModel := new(blockchainv3.SubmitBlockOptions)
				submitBlockOptionsModel.ID = core.StringPtr("testString")
				submitBlockOptionsModel.B64Block = core.StringPtr("bWFzc2l2ZSBiaW5hcnkgb2YgYSBjb25maWcgYmxvY2sgd291bGQgYmUgaGVyZSBpZiB0aGlzIHdhcyByZWFsLCBwbGVhc2UgZG9udCBzZW5kIHRoaXM=")
				submitBlockOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.SubmitBlock(submitBlockOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.SubmitBlock(submitBlockOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SubmitBlock(submitBlockOptions *SubmitBlockOptions)`, func() {
		submitBlockPath := "/ak/api/v3/kubernetes/components/testString/config"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(submitBlockPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "myca-2", "type": "fabric-ca", "display_name": "Example CA", "cluster_id": "mzdqhdifnl", "cluster_name": "ordering service 1", "grpcwp_url": "https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084", "api_url": "grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051", "operations_url": "https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443", "msp": {"ca": {"name": "org1CA", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "tlsca": {"name": "org1tlsCA", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "component": {"tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "ecert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "admin_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}}, "msp_id": "Org1", "location": "ibmcloud", "node_ou": {"enabled": true}, "resources": {"ca": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "peer": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "orderer": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "proxy": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "statedb": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "state_db": "couchdb", "storage": {"ca": {"size": "4GiB", "class": "default"}, "peer": {"size": "4GiB", "class": "default"}, "orderer": {"size": "4GiB", "class": "default"}, "statedb": {"size": "4GiB", "class": "default"}}, "timestamp": 1537262855753, "tags": ["fabric-ca"], "version": "1.4.6-1", "zone": "-"}`)
				}))
			})
			It(`Invoke SubmitBlock successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.SubmitBlock(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SubmitBlockOptions model
				submitBlockOptionsModel := new(blockchainv3.SubmitBlockOptions)
				submitBlockOptionsModel.ID = core.StringPtr("testString")
				submitBlockOptionsModel.B64Block = core.StringPtr("bWFzc2l2ZSBiaW5hcnkgb2YgYSBjb25maWcgYmxvY2sgd291bGQgYmUgaGVyZSBpZiB0aGlzIHdhcyByZWFsLCBwbGVhc2UgZG9udCBzZW5kIHRoaXM=")
				submitBlockOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.SubmitBlock(submitBlockOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.SubmitBlockWithContext(ctx, submitBlockOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.SubmitBlock(submitBlockOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.SubmitBlockWithContext(ctx, submitBlockOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke SubmitBlock with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the SubmitBlockOptions model
				submitBlockOptionsModel := new(blockchainv3.SubmitBlockOptions)
				submitBlockOptionsModel.ID = core.StringPtr("testString")
				submitBlockOptionsModel.B64Block = core.StringPtr("bWFzc2l2ZSBiaW5hcnkgb2YgYSBjb25maWcgYmxvY2sgd291bGQgYmUgaGVyZSBpZiB0aGlzIHdhcyByZWFsLCBwbGVhc2UgZG9udCBzZW5kIHRoaXM=")
				submitBlockOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.SubmitBlock(submitBlockOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SubmitBlockOptions model with no property values
				submitBlockOptionsModelNew := new(blockchainv3.SubmitBlockOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.SubmitBlock(submitBlockOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ImportMsp(importMspOptions *ImportMspOptions) - Operation response error`, func() {
		importMspPath := "/ak/api/v3/components/msp"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(importMspPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ImportMsp with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the ImportMspOptions model
				importMspOptionsModel := new(blockchainv3.ImportMspOptions)
				importMspOptionsModel.MspID = core.StringPtr("Org1")
				importMspOptionsModel.DisplayName = core.StringPtr("My Peer")
				importMspOptionsModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				importMspOptionsModel.IntermediateCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkRhdGEgaGVyZSBpZiB0aGlzIHdhcyByZWFsCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"}
				importMspOptionsModel.Admins = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				importMspOptionsModel.TlsRootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				importMspOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.ImportMsp(importMspOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.ImportMsp(importMspOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ImportMsp(importMspOptions *ImportMspOptions)`, func() {
		importMspPath := "/ak/api/v3/components/msp"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(importMspPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "component1", "type": "fabric-peer", "display_name": "My Peer", "msp_id": "Org1", "timestamp": 1537262855753, "tags": ["fabric-ca"], "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="], "intermediate_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkRhdGEgaGVyZSBpZiB0aGlzIHdhcyByZWFsCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"], "admins": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="], "scheme_version": "v1", "tls_root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}`)
				}))
			})
			It(`Invoke ImportMsp successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.ImportMsp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ImportMspOptions model
				importMspOptionsModel := new(blockchainv3.ImportMspOptions)
				importMspOptionsModel.MspID = core.StringPtr("Org1")
				importMspOptionsModel.DisplayName = core.StringPtr("My Peer")
				importMspOptionsModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				importMspOptionsModel.IntermediateCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkRhdGEgaGVyZSBpZiB0aGlzIHdhcyByZWFsCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"}
				importMspOptionsModel.Admins = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				importMspOptionsModel.TlsRootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				importMspOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.ImportMsp(importMspOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.ImportMspWithContext(ctx, importMspOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.ImportMsp(importMspOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.ImportMspWithContext(ctx, importMspOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ImportMsp with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the ImportMspOptions model
				importMspOptionsModel := new(blockchainv3.ImportMspOptions)
				importMspOptionsModel.MspID = core.StringPtr("Org1")
				importMspOptionsModel.DisplayName = core.StringPtr("My Peer")
				importMspOptionsModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				importMspOptionsModel.IntermediateCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkRhdGEgaGVyZSBpZiB0aGlzIHdhcyByZWFsCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"}
				importMspOptionsModel.Admins = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				importMspOptionsModel.TlsRootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				importMspOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.ImportMsp(importMspOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ImportMspOptions model with no property values
				importMspOptionsModelNew := new(blockchainv3.ImportMspOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.ImportMsp(importMspOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`EditMsp(editMspOptions *EditMspOptions) - Operation response error`, func() {
		editMspPath := "/ak/api/v3/components/msp/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(editMspPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke EditMsp with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the EditMspOptions model
				editMspOptionsModel := new(blockchainv3.EditMspOptions)
				editMspOptionsModel.ID = core.StringPtr("testString")
				editMspOptionsModel.MspID = core.StringPtr("Org1")
				editMspOptionsModel.DisplayName = core.StringPtr("My Peer")
				editMspOptionsModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editMspOptionsModel.IntermediateCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkRhdGEgaGVyZSBpZiB0aGlzIHdhcyByZWFsCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"}
				editMspOptionsModel.Admins = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editMspOptionsModel.TlsRootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editMspOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.EditMsp(editMspOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.EditMsp(editMspOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`EditMsp(editMspOptions *EditMspOptions)`, func() {
		editMspPath := "/ak/api/v3/components/msp/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(editMspPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "component1", "type": "fabric-peer", "display_name": "My Peer", "msp_id": "Org1", "timestamp": 1537262855753, "tags": ["fabric-ca"], "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="], "intermediate_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkRhdGEgaGVyZSBpZiB0aGlzIHdhcyByZWFsCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"], "admins": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="], "scheme_version": "v1", "tls_root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}`)
				}))
			})
			It(`Invoke EditMsp successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.EditMsp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EditMspOptions model
				editMspOptionsModel := new(blockchainv3.EditMspOptions)
				editMspOptionsModel.ID = core.StringPtr("testString")
				editMspOptionsModel.MspID = core.StringPtr("Org1")
				editMspOptionsModel.DisplayName = core.StringPtr("My Peer")
				editMspOptionsModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editMspOptionsModel.IntermediateCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkRhdGEgaGVyZSBpZiB0aGlzIHdhcyByZWFsCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"}
				editMspOptionsModel.Admins = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editMspOptionsModel.TlsRootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editMspOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.EditMsp(editMspOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.EditMspWithContext(ctx, editMspOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.EditMsp(editMspOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.EditMspWithContext(ctx, editMspOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke EditMsp with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the EditMspOptions model
				editMspOptionsModel := new(blockchainv3.EditMspOptions)
				editMspOptionsModel.ID = core.StringPtr("testString")
				editMspOptionsModel.MspID = core.StringPtr("Org1")
				editMspOptionsModel.DisplayName = core.StringPtr("My Peer")
				editMspOptionsModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editMspOptionsModel.IntermediateCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkRhdGEgaGVyZSBpZiB0aGlzIHdhcyByZWFsCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"}
				editMspOptionsModel.Admins = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editMspOptionsModel.TlsRootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editMspOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.EditMsp(editMspOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the EditMspOptions model with no property values
				editMspOptionsModelNew := new(blockchainv3.EditMspOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.EditMsp(editMspOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMspCertificate(getMspCertificateOptions *GetMspCertificateOptions) - Operation response error`, func() {
		getMspCertificatePath := "/ak/api/v3/components/msps/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMspCertificatePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["cache"]).To(Equal([]string{"skip"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetMspCertificate with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the GetMspCertificateOptions model
				getMspCertificateOptionsModel := new(blockchainv3.GetMspCertificateOptions)
				getMspCertificateOptionsModel.MspID = core.StringPtr("testString")
				getMspCertificateOptionsModel.Cache = core.StringPtr("skip")
				getMspCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.GetMspCertificate(getMspCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.GetMspCertificate(getMspCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetMspCertificate(getMspCertificateOptions *GetMspCertificateOptions)`, func() {
		getMspCertificatePath := "/ak/api/v3/components/msps/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMspCertificatePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["cache"]).To(Equal([]string{"skip"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"msps": [{"msp_id": "Org1", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="], "admins": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="], "tls_root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}]}`)
				}))
			})
			It(`Invoke GetMspCertificate successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.GetMspCertificate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetMspCertificateOptions model
				getMspCertificateOptionsModel := new(blockchainv3.GetMspCertificateOptions)
				getMspCertificateOptionsModel.MspID = core.StringPtr("testString")
				getMspCertificateOptionsModel.Cache = core.StringPtr("skip")
				getMspCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.GetMspCertificate(getMspCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.GetMspCertificateWithContext(ctx, getMspCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.GetMspCertificate(getMspCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.GetMspCertificateWithContext(ctx, getMspCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetMspCertificate with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the GetMspCertificateOptions model
				getMspCertificateOptionsModel := new(blockchainv3.GetMspCertificateOptions)
				getMspCertificateOptionsModel.MspID = core.StringPtr("testString")
				getMspCertificateOptionsModel.Cache = core.StringPtr("skip")
				getMspCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.GetMspCertificate(getMspCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetMspCertificateOptions model with no property values
				getMspCertificateOptionsModelNew := new(blockchainv3.GetMspCertificateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.GetMspCertificate(getMspCertificateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`EditAdminCerts(editAdminCertsOptions *EditAdminCertsOptions) - Operation response error`, func() {
		editAdminCertsPath := "/ak/api/v3/kubernetes/components/testString/certs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(editAdminCertsPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke EditAdminCerts with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the EditAdminCertsOptions model
				editAdminCertsOptionsModel := new(blockchainv3.EditAdminCertsOptions)
				editAdminCertsOptionsModel.ID = core.StringPtr("testString")
				editAdminCertsOptionsModel.AppendAdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editAdminCertsOptionsModel.RemoveAdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editAdminCertsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.EditAdminCerts(editAdminCertsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.EditAdminCerts(editAdminCertsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`EditAdminCerts(editAdminCertsOptions *EditAdminCertsOptions)`, func() {
		editAdminCertsPath := "/ak/api/v3/kubernetes/components/testString/certs"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(editAdminCertsPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"changes_made": 1, "set_admin_certs": [{"base_64_pem": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "issuer": "/C=US/ST=North Carolina/O=Hyperledger/OU=Fabric/CN=fabric-ca-server", "not_after_ts": 1597770420000, "not_before_ts": 1566234120000, "serial_number_hex": "649a1206fd0bc8be994886dd715cecb0a7a21276", "signature_algorithm": "SHA256withECDSA", "subject": "/OU=client/CN=admin", "X509_version": 3, "time_left": "10 hrs"}]}`)
				}))
			})
			It(`Invoke EditAdminCerts successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.EditAdminCerts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EditAdminCertsOptions model
				editAdminCertsOptionsModel := new(blockchainv3.EditAdminCertsOptions)
				editAdminCertsOptionsModel.ID = core.StringPtr("testString")
				editAdminCertsOptionsModel.AppendAdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editAdminCertsOptionsModel.RemoveAdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editAdminCertsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.EditAdminCerts(editAdminCertsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.EditAdminCertsWithContext(ctx, editAdminCertsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.EditAdminCerts(editAdminCertsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.EditAdminCertsWithContext(ctx, editAdminCertsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke EditAdminCerts with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the EditAdminCertsOptions model
				editAdminCertsOptionsModel := new(blockchainv3.EditAdminCertsOptions)
				editAdminCertsOptionsModel.ID = core.StringPtr("testString")
				editAdminCertsOptionsModel.AppendAdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editAdminCertsOptionsModel.RemoveAdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editAdminCertsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.EditAdminCerts(editAdminCertsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the EditAdminCertsOptions model with no property values
				editAdminCertsOptionsModelNew := new(blockchainv3.EditAdminCertsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.EditAdminCerts(editAdminCertsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(blockchainService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(blockchainService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
				URL: "https://blockchainv3/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(blockchainService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BLOCKCHAIN_URL": "https://blockchainv3/api",
				"BLOCKCHAIN_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3UsingExternalConfig(&blockchainv3.BlockchainV3Options{
				})
				Expect(blockchainService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := blockchainService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != blockchainService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(blockchainService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(blockchainService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3UsingExternalConfig(&blockchainv3.BlockchainV3Options{
					URL: "https://testService/api",
				})
				Expect(blockchainService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := blockchainService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != blockchainService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(blockchainService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(blockchainService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3UsingExternalConfig(&blockchainv3.BlockchainV3Options{
				})
				err := blockchainService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := blockchainService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != blockchainService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(blockchainService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(blockchainService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BLOCKCHAIN_URL": "https://blockchainv3/api",
				"BLOCKCHAIN_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			blockchainService, serviceErr := blockchainv3.NewBlockchainV3UsingExternalConfig(&blockchainv3.BlockchainV3Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(blockchainService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BLOCKCHAIN_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			blockchainService, serviceErr := blockchainv3.NewBlockchainV3UsingExternalConfig(&blockchainv3.BlockchainV3Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(blockchainService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = blockchainv3.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListComponents(listComponentsOptions *ListComponentsOptions) - Operation response error`, func() {
		listComponentsPath := "/ak/api/v3/components"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listComponentsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["deployment_attrs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["parsed_certs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["cache"]).To(Equal([]string{"skip"}))

					Expect(req.URL.Query()["ca_attrs"]).To(Equal([]string{"included"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListComponents with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the ListComponentsOptions model
				listComponentsOptionsModel := new(blockchainv3.ListComponentsOptions)
				listComponentsOptionsModel.DeploymentAttrs = core.StringPtr("included")
				listComponentsOptionsModel.ParsedCerts = core.StringPtr("included")
				listComponentsOptionsModel.Cache = core.StringPtr("skip")
				listComponentsOptionsModel.CaAttrs = core.StringPtr("included")
				listComponentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.ListComponents(listComponentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.ListComponents(listComponentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListComponents(listComponentsOptions *ListComponentsOptions)`, func() {
		listComponentsPath := "/ak/api/v3/components"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listComponentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["deployment_attrs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["parsed_certs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["cache"]).To(Equal([]string{"skip"}))

					Expect(req.URL.Query()["ca_attrs"]).To(Equal([]string{"included"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"components": [{"id": "myca-2", "type": "fabric-ca", "display_name": "Example CA", "cluster_id": "mzdqhdifnl", "cluster_name": "ordering service 1", "grpcwp_url": "https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084", "api_url": "grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051", "operations_url": "https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443", "msp": {"ca": {"name": "org1CA", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "tlsca": {"name": "org1tlsCA", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "component": {"tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "ecert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "admin_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}}, "msp_id": "Org1", "location": "ibmcloud", "node_ou": {"enabled": true}, "resources": {"ca": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "peer": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "orderer": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "proxy": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "statedb": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "state_db": "couchdb", "storage": {"ca": {"size": "4GiB", "class": "default"}, "peer": {"size": "4GiB", "class": "default"}, "orderer": {"size": "4GiB", "class": "default"}, "statedb": {"size": "4GiB", "class": "default"}}, "timestamp": 1537262855753, "tags": ["fabric-ca"], "version": "1.4.6-1", "zone": "-"}]}`)
				}))
			})
			It(`Invoke ListComponents successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.ListComponents(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListComponentsOptions model
				listComponentsOptionsModel := new(blockchainv3.ListComponentsOptions)
				listComponentsOptionsModel.DeploymentAttrs = core.StringPtr("included")
				listComponentsOptionsModel.ParsedCerts = core.StringPtr("included")
				listComponentsOptionsModel.Cache = core.StringPtr("skip")
				listComponentsOptionsModel.CaAttrs = core.StringPtr("included")
				listComponentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.ListComponents(listComponentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.ListComponentsWithContext(ctx, listComponentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.ListComponents(listComponentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.ListComponentsWithContext(ctx, listComponentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListComponents with error: Operation request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the ListComponentsOptions model
				listComponentsOptionsModel := new(blockchainv3.ListComponentsOptions)
				listComponentsOptionsModel.DeploymentAttrs = core.StringPtr("included")
				listComponentsOptionsModel.ParsedCerts = core.StringPtr("included")
				listComponentsOptionsModel.Cache = core.StringPtr("skip")
				listComponentsOptionsModel.CaAttrs = core.StringPtr("included")
				listComponentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.ListComponents(listComponentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetComponentsByType(getComponentsByTypeOptions *GetComponentsByTypeOptions) - Operation response error`, func() {
		getComponentsByTypePath := "/ak/api/v3/components/types/fabric-peer"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getComponentsByTypePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["deployment_attrs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["parsed_certs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["cache"]).To(Equal([]string{"skip"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetComponentsByType with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the GetComponentsByTypeOptions model
				getComponentsByTypeOptionsModel := new(blockchainv3.GetComponentsByTypeOptions)
				getComponentsByTypeOptionsModel.Type = core.StringPtr("fabric-peer")
				getComponentsByTypeOptionsModel.DeploymentAttrs = core.StringPtr("included")
				getComponentsByTypeOptionsModel.ParsedCerts = core.StringPtr("included")
				getComponentsByTypeOptionsModel.Cache = core.StringPtr("skip")
				getComponentsByTypeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.GetComponentsByType(getComponentsByTypeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.GetComponentsByType(getComponentsByTypeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetComponentsByType(getComponentsByTypeOptions *GetComponentsByTypeOptions)`, func() {
		getComponentsByTypePath := "/ak/api/v3/components/types/fabric-peer"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getComponentsByTypePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["deployment_attrs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["parsed_certs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["cache"]).To(Equal([]string{"skip"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"components": [{"id": "myca-2", "type": "fabric-ca", "display_name": "Example CA", "cluster_id": "mzdqhdifnl", "cluster_name": "ordering service 1", "grpcwp_url": "https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084", "api_url": "grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051", "operations_url": "https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443", "msp": {"ca": {"name": "org1CA", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "tlsca": {"name": "org1tlsCA", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "component": {"tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "ecert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "admin_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}}, "msp_id": "Org1", "location": "ibmcloud", "node_ou": {"enabled": true}, "resources": {"ca": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "peer": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "orderer": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "proxy": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "statedb": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "state_db": "couchdb", "storage": {"ca": {"size": "4GiB", "class": "default"}, "peer": {"size": "4GiB", "class": "default"}, "orderer": {"size": "4GiB", "class": "default"}, "statedb": {"size": "4GiB", "class": "default"}}, "timestamp": 1537262855753, "tags": ["fabric-ca"], "version": "1.4.6-1", "zone": "-"}]}`)
				}))
			})
			It(`Invoke GetComponentsByType successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.GetComponentsByType(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetComponentsByTypeOptions model
				getComponentsByTypeOptionsModel := new(blockchainv3.GetComponentsByTypeOptions)
				getComponentsByTypeOptionsModel.Type = core.StringPtr("fabric-peer")
				getComponentsByTypeOptionsModel.DeploymentAttrs = core.StringPtr("included")
				getComponentsByTypeOptionsModel.ParsedCerts = core.StringPtr("included")
				getComponentsByTypeOptionsModel.Cache = core.StringPtr("skip")
				getComponentsByTypeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.GetComponentsByType(getComponentsByTypeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.GetComponentsByTypeWithContext(ctx, getComponentsByTypeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.GetComponentsByType(getComponentsByTypeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.GetComponentsByTypeWithContext(ctx, getComponentsByTypeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetComponentsByType with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the GetComponentsByTypeOptions model
				getComponentsByTypeOptionsModel := new(blockchainv3.GetComponentsByTypeOptions)
				getComponentsByTypeOptionsModel.Type = core.StringPtr("fabric-peer")
				getComponentsByTypeOptionsModel.DeploymentAttrs = core.StringPtr("included")
				getComponentsByTypeOptionsModel.ParsedCerts = core.StringPtr("included")
				getComponentsByTypeOptionsModel.Cache = core.StringPtr("skip")
				getComponentsByTypeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.GetComponentsByType(getComponentsByTypeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetComponentsByTypeOptions model with no property values
				getComponentsByTypeOptionsModelNew := new(blockchainv3.GetComponentsByTypeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.GetComponentsByType(getComponentsByTypeOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetComponentsByTag(getComponentsByTagOptions *GetComponentsByTagOptions) - Operation response error`, func() {
		getComponentsByTagPath := "/ak/api/v3/components/tags/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getComponentsByTagPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["deployment_attrs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["parsed_certs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["cache"]).To(Equal([]string{"skip"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetComponentsByTag with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the GetComponentsByTagOptions model
				getComponentsByTagOptionsModel := new(blockchainv3.GetComponentsByTagOptions)
				getComponentsByTagOptionsModel.Tag = core.StringPtr("testString")
				getComponentsByTagOptionsModel.DeploymentAttrs = core.StringPtr("included")
				getComponentsByTagOptionsModel.ParsedCerts = core.StringPtr("included")
				getComponentsByTagOptionsModel.Cache = core.StringPtr("skip")
				getComponentsByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.GetComponentsByTag(getComponentsByTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.GetComponentsByTag(getComponentsByTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetComponentsByTag(getComponentsByTagOptions *GetComponentsByTagOptions)`, func() {
		getComponentsByTagPath := "/ak/api/v3/components/tags/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getComponentsByTagPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["deployment_attrs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["parsed_certs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["cache"]).To(Equal([]string{"skip"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"components": [{"id": "myca-2", "type": "fabric-ca", "display_name": "Example CA", "cluster_id": "mzdqhdifnl", "cluster_name": "ordering service 1", "grpcwp_url": "https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084", "api_url": "grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051", "operations_url": "https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443", "msp": {"ca": {"name": "org1CA", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "tlsca": {"name": "org1tlsCA", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}, "component": {"tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "ecert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "admin_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}}, "msp_id": "Org1", "location": "ibmcloud", "node_ou": {"enabled": true}, "resources": {"ca": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "peer": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "orderer": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "proxy": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "statedb": {"requests": {"cpu": "100m", "memory": "256M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "state_db": "couchdb", "storage": {"ca": {"size": "4GiB", "class": "default"}, "peer": {"size": "4GiB", "class": "default"}, "orderer": {"size": "4GiB", "class": "default"}, "statedb": {"size": "4GiB", "class": "default"}}, "timestamp": 1537262855753, "tags": ["fabric-ca"], "version": "1.4.6-1", "zone": "-"}]}`)
				}))
			})
			It(`Invoke GetComponentsByTag successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.GetComponentsByTag(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetComponentsByTagOptions model
				getComponentsByTagOptionsModel := new(blockchainv3.GetComponentsByTagOptions)
				getComponentsByTagOptionsModel.Tag = core.StringPtr("testString")
				getComponentsByTagOptionsModel.DeploymentAttrs = core.StringPtr("included")
				getComponentsByTagOptionsModel.ParsedCerts = core.StringPtr("included")
				getComponentsByTagOptionsModel.Cache = core.StringPtr("skip")
				getComponentsByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.GetComponentsByTag(getComponentsByTagOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.GetComponentsByTagWithContext(ctx, getComponentsByTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.GetComponentsByTag(getComponentsByTagOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.GetComponentsByTagWithContext(ctx, getComponentsByTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetComponentsByTag with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the GetComponentsByTagOptions model
				getComponentsByTagOptionsModel := new(blockchainv3.GetComponentsByTagOptions)
				getComponentsByTagOptionsModel.Tag = core.StringPtr("testString")
				getComponentsByTagOptionsModel.DeploymentAttrs = core.StringPtr("included")
				getComponentsByTagOptionsModel.ParsedCerts = core.StringPtr("included")
				getComponentsByTagOptionsModel.Cache = core.StringPtr("skip")
				getComponentsByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.GetComponentsByTag(getComponentsByTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetComponentsByTagOptions model with no property values
				getComponentsByTagOptionsModelNew := new(blockchainv3.GetComponentsByTagOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.GetComponentsByTag(getComponentsByTagOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RemoveComponentsByTag(removeComponentsByTagOptions *RemoveComponentsByTagOptions) - Operation response error`, func() {
		removeComponentsByTagPath := "/ak/api/v3/components/tags/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(removeComponentsByTagPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RemoveComponentsByTag with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the RemoveComponentsByTagOptions model
				removeComponentsByTagOptionsModel := new(blockchainv3.RemoveComponentsByTagOptions)
				removeComponentsByTagOptionsModel.Tag = core.StringPtr("testString")
				removeComponentsByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.RemoveComponentsByTag(removeComponentsByTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.RemoveComponentsByTag(removeComponentsByTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`RemoveComponentsByTag(removeComponentsByTagOptions *RemoveComponentsByTagOptions)`, func() {
		removeComponentsByTagPath := "/ak/api/v3/components/tags/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(removeComponentsByTagPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"removed": [{"message": "deleted", "type": "fabric-peer", "id": "component1", "display_name": "My Peer"}]}`)
				}))
			})
			It(`Invoke RemoveComponentsByTag successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.RemoveComponentsByTag(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RemoveComponentsByTagOptions model
				removeComponentsByTagOptionsModel := new(blockchainv3.RemoveComponentsByTagOptions)
				removeComponentsByTagOptionsModel.Tag = core.StringPtr("testString")
				removeComponentsByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.RemoveComponentsByTag(removeComponentsByTagOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.RemoveComponentsByTagWithContext(ctx, removeComponentsByTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.RemoveComponentsByTag(removeComponentsByTagOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.RemoveComponentsByTagWithContext(ctx, removeComponentsByTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke RemoveComponentsByTag with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the RemoveComponentsByTagOptions model
				removeComponentsByTagOptionsModel := new(blockchainv3.RemoveComponentsByTagOptions)
				removeComponentsByTagOptionsModel.Tag = core.StringPtr("testString")
				removeComponentsByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.RemoveComponentsByTag(removeComponentsByTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RemoveComponentsByTagOptions model with no property values
				removeComponentsByTagOptionsModelNew := new(blockchainv3.RemoveComponentsByTagOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.RemoveComponentsByTag(removeComponentsByTagOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteComponentsByTag(deleteComponentsByTagOptions *DeleteComponentsByTagOptions) - Operation response error`, func() {
		deleteComponentsByTagPath := "/ak/api/v3/kubernetes/components/tags/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteComponentsByTagPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteComponentsByTag with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the DeleteComponentsByTagOptions model
				deleteComponentsByTagOptionsModel := new(blockchainv3.DeleteComponentsByTagOptions)
				deleteComponentsByTagOptionsModel.Tag = core.StringPtr("testString")
				deleteComponentsByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.DeleteComponentsByTag(deleteComponentsByTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.DeleteComponentsByTag(deleteComponentsByTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteComponentsByTag(deleteComponentsByTagOptions *DeleteComponentsByTagOptions)`, func() {
		deleteComponentsByTagPath := "/ak/api/v3/kubernetes/components/tags/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteComponentsByTagPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"deleted": [{"message": "deleted", "type": "fabric-peer", "id": "component1", "display_name": "My Peer"}]}`)
				}))
			})
			It(`Invoke DeleteComponentsByTag successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.DeleteComponentsByTag(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteComponentsByTagOptions model
				deleteComponentsByTagOptionsModel := new(blockchainv3.DeleteComponentsByTagOptions)
				deleteComponentsByTagOptionsModel.Tag = core.StringPtr("testString")
				deleteComponentsByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.DeleteComponentsByTag(deleteComponentsByTagOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.DeleteComponentsByTagWithContext(ctx, deleteComponentsByTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.DeleteComponentsByTag(deleteComponentsByTagOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.DeleteComponentsByTagWithContext(ctx, deleteComponentsByTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke DeleteComponentsByTag with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the DeleteComponentsByTagOptions model
				deleteComponentsByTagOptionsModel := new(blockchainv3.DeleteComponentsByTagOptions)
				deleteComponentsByTagOptionsModel.Tag = core.StringPtr("testString")
				deleteComponentsByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.DeleteComponentsByTag(deleteComponentsByTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteComponentsByTagOptions model with no property values
				deleteComponentsByTagOptionsModelNew := new(blockchainv3.DeleteComponentsByTagOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.DeleteComponentsByTag(deleteComponentsByTagOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteAllComponents(deleteAllComponentsOptions *DeleteAllComponentsOptions) - Operation response error`, func() {
		deleteAllComponentsPath := "/ak/api/v3/kubernetes/components/purge"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAllComponentsPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteAllComponents with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the DeleteAllComponentsOptions model
				deleteAllComponentsOptionsModel := new(blockchainv3.DeleteAllComponentsOptions)
				deleteAllComponentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.DeleteAllComponents(deleteAllComponentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.DeleteAllComponents(deleteAllComponentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteAllComponents(deleteAllComponentsOptions *DeleteAllComponentsOptions)`, func() {
		deleteAllComponentsPath := "/ak/api/v3/kubernetes/components/purge"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAllComponentsPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"deleted": [{"message": "deleted", "type": "fabric-peer", "id": "component1", "display_name": "My Peer"}]}`)
				}))
			})
			It(`Invoke DeleteAllComponents successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.DeleteAllComponents(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteAllComponentsOptions model
				deleteAllComponentsOptionsModel := new(blockchainv3.DeleteAllComponentsOptions)
				deleteAllComponentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.DeleteAllComponents(deleteAllComponentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.DeleteAllComponentsWithContext(ctx, deleteAllComponentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.DeleteAllComponents(deleteAllComponentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.DeleteAllComponentsWithContext(ctx, deleteAllComponentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke DeleteAllComponents with error: Operation request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the DeleteAllComponentsOptions model
				deleteAllComponentsOptionsModel := new(blockchainv3.DeleteAllComponentsOptions)
				deleteAllComponentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.DeleteAllComponents(deleteAllComponentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(blockchainService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(blockchainService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
				URL: "https://blockchainv3/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(blockchainService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BLOCKCHAIN_URL": "https://blockchainv3/api",
				"BLOCKCHAIN_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3UsingExternalConfig(&blockchainv3.BlockchainV3Options{
				})
				Expect(blockchainService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := blockchainService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != blockchainService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(blockchainService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(blockchainService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3UsingExternalConfig(&blockchainv3.BlockchainV3Options{
					URL: "https://testService/api",
				})
				Expect(blockchainService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := blockchainService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != blockchainService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(blockchainService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(blockchainService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3UsingExternalConfig(&blockchainv3.BlockchainV3Options{
				})
				err := blockchainService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := blockchainService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != blockchainService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(blockchainService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(blockchainService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BLOCKCHAIN_URL": "https://blockchainv3/api",
				"BLOCKCHAIN_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			blockchainService, serviceErr := blockchainv3.NewBlockchainV3UsingExternalConfig(&blockchainv3.BlockchainV3Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(blockchainService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BLOCKCHAIN_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			blockchainService, serviceErr := blockchainv3.NewBlockchainV3UsingExternalConfig(&blockchainv3.BlockchainV3Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(blockchainService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = blockchainv3.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetSettings(getSettingsOptions *GetSettingsOptions) - Operation response error`, func() {
		getSettingsPath := "/ak/api/v3/settings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSettings with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(blockchainv3.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSettings(getSettingsOptions *GetSettingsOptions)`, func() {
		getSettingsPath := "/ak/api/v3/settings"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ACTIVITY_TRACKER_PATH": "/logs", "ATHENA_ID": "17v7e", "AUTH_SCHEME": "iam", "CALLBACK_URI": "/auth/cb", "CLUSTER_DATA": {"type": "paid"}, "CONFIGTXLATOR_URL": "https://n3a3ec3-configtxlator.ibp.us-south.containers.appdomain.cloud", "CRN": {"account_id": "a/abcd", "c_name": "staging", "c_type": "public", "instance_id": "abc123", "location": "us-south", "resource_id": "-", "resource_type": "-", "service_name": "blockchain", "version": "v1"}, "CRN_STRING": "crn:v1:staging:public:blockchain:us-south:a/abcd:abc123::", "CSP_HEADER_VALUES": ["-"], "DB_SYSTEM": "system", "DEPLOYER_URL": "https://api.dev.blockchain.cloud.ibm.com", "DOMAIN": "localhost", "ENVIRONMENT": "prod", "FABRIC_CAPABILITIES": {"application": ["V1_1"], "channel": ["V1_1"], "orderer": ["V1_1"]}, "FEATURE_FLAGS": {"anyKey": "anyValue"}, "FILE_LOGGING": {"server": {"client": {"enabled": true, "level": "silly", "unique_name": false}, "server": {"enabled": true, "level": "silly", "unique_name": false}}, "client": {"client": {"enabled": true, "level": "silly", "unique_name": false}, "server": {"enabled": true, "level": "silly", "unique_name": false}}}, "HOST_URL": "http://localhost:3000", "IAM_CACHE_ENABLED": true, "IAM_URL": "-", "IBM_ID_CALLBACK_URL": "http://localhost:3000/auth/login", "IGNORE_CONFIG_FILE": true, "INACTIVITY_TIMEOUTS": {"enabled": true, "max_idle_time": 60000}, "INFRASTRUCTURE": "ibmcloud", "LANDING_URL": "http://localhost:3000", "LOGIN_URI": "/auth/login", "LOGOUT_URI": "/auth/logout", "MAX_REQ_PER_MIN": 25, "MAX_REQ_PER_MIN_AK": 25, "MEMORY_CACHE_ENABLED": true, "PORT": 3000, "PROXY_CACHE_ENABLED": true, "PROXY_TLS_FABRIC_REQS": "always", "PROXY_TLS_HTTP_URL": "http://localhost:3000", "PROXY_TLS_WS_URL": "http://localhost:3000", "REGION": "us_south", "SESSION_CACHE_ENABLED": true, "TIMEOUTS": {"anyKey": "anyValue"}, "TIMESTAMPS": {"now": 1542746836056, "born": 1542746836056, "next_settings_update": "1.2 mins", "up_time": "30 days"}, "TRANSACTION_VISIBILITY": {"anyKey": "anyValue"}, "TRUST_PROXY": "loopback", "TRUST_UNKNOWN_CERTS": true, "VERSIONS": {"apollo": "65f3cbfd", "athena": "1198f94", "stitch": "0f1a0c6", "tag": "v0.4.31"}}`)
				}))
			})
			It(`Invoke GetSettings successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.GetSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(blockchainv3.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.GetSettingsWithContext(ctx, getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.GetSettingsWithContext(ctx, getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetSettings with error: Operation request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(blockchainv3.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`EditSettings(editSettingsOptions *EditSettingsOptions) - Operation response error`, func() {
		editSettingsPath := "/ak/api/v3/settings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(editSettingsPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke EditSettings with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the EditSettingsBodyInactivityTimeouts model
				editSettingsBodyInactivityTimeoutsModel := new(blockchainv3.EditSettingsBodyInactivityTimeouts)
				editSettingsBodyInactivityTimeoutsModel.Enabled = core.BoolPtr(false)
				editSettingsBodyInactivityTimeoutsModel.MaxIdleTime = core.Float64Ptr(float64(90000))

				// Construct an instance of the LoggingSettingsClient model
				loggingSettingsClientModel := new(blockchainv3.LoggingSettingsClient)
				loggingSettingsClientModel.Enabled = core.BoolPtr(true)
				loggingSettingsClientModel.Level = core.StringPtr("silly")
				loggingSettingsClientModel.UniqueName = core.BoolPtr(false)

				// Construct an instance of the LoggingSettingsServer model
				loggingSettingsServerModel := new(blockchainv3.LoggingSettingsServer)
				loggingSettingsServerModel.Enabled = core.BoolPtr(true)
				loggingSettingsServerModel.Level = core.StringPtr("silly")
				loggingSettingsServerModel.UniqueName = core.BoolPtr(false)

				// Construct an instance of the EditLogSettingsBody model
				editLogSettingsBodyModel := new(blockchainv3.EditLogSettingsBody)
				editLogSettingsBodyModel.Client = loggingSettingsClientModel
				editLogSettingsBodyModel.Server = loggingSettingsServerModel

				// Construct an instance of the EditSettingsOptions model
				editSettingsOptionsModel := new(blockchainv3.EditSettingsOptions)
				editSettingsOptionsModel.InactivityTimeouts = editSettingsBodyInactivityTimeoutsModel
				editSettingsOptionsModel.FileLogging = editLogSettingsBodyModel
				editSettingsOptionsModel.MaxReqPerMin = core.Float64Ptr(float64(25))
				editSettingsOptionsModel.MaxReqPerMinAk = core.Float64Ptr(float64(25))
				editSettingsOptionsModel.FabricGetBlockTimeoutMs = core.Float64Ptr(float64(10000))
				editSettingsOptionsModel.FabricInstantiateTimeoutMs = core.Float64Ptr(float64(300000))
				editSettingsOptionsModel.FabricJoinChannelTimeoutMs = core.Float64Ptr(float64(25000))
				editSettingsOptionsModel.FabricInstallCcTimeoutMs = core.Float64Ptr(float64(300000))
				editSettingsOptionsModel.FabricLcInstallCcTimeoutMs = core.Float64Ptr(float64(300000))
				editSettingsOptionsModel.FabricLcGetCcTimeoutMs = core.Float64Ptr(float64(180000))
				editSettingsOptionsModel.FabricGeneralTimeoutMs = core.Float64Ptr(float64(10000))
				editSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.EditSettings(editSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.EditSettings(editSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`EditSettings(editSettingsOptions *EditSettingsOptions)`, func() {
		editSettingsPath := "/ak/api/v3/settings"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(editSettingsPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ACTIVITY_TRACKER_PATH": "/logs", "ATHENA_ID": "17v7e", "AUTH_SCHEME": "iam", "CALLBACK_URI": "/auth/cb", "CLUSTER_DATA": {"type": "paid"}, "CONFIGTXLATOR_URL": "https://n3a3ec3-configtxlator.ibp.us-south.containers.appdomain.cloud", "CRN": {"account_id": "a/abcd", "c_name": "staging", "c_type": "public", "instance_id": "abc123", "location": "us-south", "resource_id": "-", "resource_type": "-", "service_name": "blockchain", "version": "v1"}, "CRN_STRING": "crn:v1:staging:public:blockchain:us-south:a/abcd:abc123::", "CSP_HEADER_VALUES": ["-"], "DB_SYSTEM": "system", "DEPLOYER_URL": "https://api.dev.blockchain.cloud.ibm.com", "DOMAIN": "localhost", "ENVIRONMENT": "prod", "FABRIC_CAPABILITIES": {"application": ["V1_1"], "channel": ["V1_1"], "orderer": ["V1_1"]}, "FEATURE_FLAGS": {"anyKey": "anyValue"}, "FILE_LOGGING": {"server": {"client": {"enabled": true, "level": "silly", "unique_name": false}, "server": {"enabled": true, "level": "silly", "unique_name": false}}, "client": {"client": {"enabled": true, "level": "silly", "unique_name": false}, "server": {"enabled": true, "level": "silly", "unique_name": false}}}, "HOST_URL": "http://localhost:3000", "IAM_CACHE_ENABLED": true, "IAM_URL": "-", "IBM_ID_CALLBACK_URL": "http://localhost:3000/auth/login", "IGNORE_CONFIG_FILE": true, "INACTIVITY_TIMEOUTS": {"enabled": true, "max_idle_time": 60000}, "INFRASTRUCTURE": "ibmcloud", "LANDING_URL": "http://localhost:3000", "LOGIN_URI": "/auth/login", "LOGOUT_URI": "/auth/logout", "MAX_REQ_PER_MIN": 25, "MAX_REQ_PER_MIN_AK": 25, "MEMORY_CACHE_ENABLED": true, "PORT": 3000, "PROXY_CACHE_ENABLED": true, "PROXY_TLS_FABRIC_REQS": "always", "PROXY_TLS_HTTP_URL": "http://localhost:3000", "PROXY_TLS_WS_URL": "http://localhost:3000", "REGION": "us_south", "SESSION_CACHE_ENABLED": true, "TIMEOUTS": {"anyKey": "anyValue"}, "TIMESTAMPS": {"now": 1542746836056, "born": 1542746836056, "next_settings_update": "1.2 mins", "up_time": "30 days"}, "TRANSACTION_VISIBILITY": {"anyKey": "anyValue"}, "TRUST_PROXY": "loopback", "TRUST_UNKNOWN_CERTS": true, "VERSIONS": {"apollo": "65f3cbfd", "athena": "1198f94", "stitch": "0f1a0c6", "tag": "v0.4.31"}}`)
				}))
			})
			It(`Invoke EditSettings successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.EditSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EditSettingsBodyInactivityTimeouts model
				editSettingsBodyInactivityTimeoutsModel := new(blockchainv3.EditSettingsBodyInactivityTimeouts)
				editSettingsBodyInactivityTimeoutsModel.Enabled = core.BoolPtr(false)
				editSettingsBodyInactivityTimeoutsModel.MaxIdleTime = core.Float64Ptr(float64(90000))

				// Construct an instance of the LoggingSettingsClient model
				loggingSettingsClientModel := new(blockchainv3.LoggingSettingsClient)
				loggingSettingsClientModel.Enabled = core.BoolPtr(true)
				loggingSettingsClientModel.Level = core.StringPtr("silly")
				loggingSettingsClientModel.UniqueName = core.BoolPtr(false)

				// Construct an instance of the LoggingSettingsServer model
				loggingSettingsServerModel := new(blockchainv3.LoggingSettingsServer)
				loggingSettingsServerModel.Enabled = core.BoolPtr(true)
				loggingSettingsServerModel.Level = core.StringPtr("silly")
				loggingSettingsServerModel.UniqueName = core.BoolPtr(false)

				// Construct an instance of the EditLogSettingsBody model
				editLogSettingsBodyModel := new(blockchainv3.EditLogSettingsBody)
				editLogSettingsBodyModel.Client = loggingSettingsClientModel
				editLogSettingsBodyModel.Server = loggingSettingsServerModel

				// Construct an instance of the EditSettingsOptions model
				editSettingsOptionsModel := new(blockchainv3.EditSettingsOptions)
				editSettingsOptionsModel.InactivityTimeouts = editSettingsBodyInactivityTimeoutsModel
				editSettingsOptionsModel.FileLogging = editLogSettingsBodyModel
				editSettingsOptionsModel.MaxReqPerMin = core.Float64Ptr(float64(25))
				editSettingsOptionsModel.MaxReqPerMinAk = core.Float64Ptr(float64(25))
				editSettingsOptionsModel.FabricGetBlockTimeoutMs = core.Float64Ptr(float64(10000))
				editSettingsOptionsModel.FabricInstantiateTimeoutMs = core.Float64Ptr(float64(300000))
				editSettingsOptionsModel.FabricJoinChannelTimeoutMs = core.Float64Ptr(float64(25000))
				editSettingsOptionsModel.FabricInstallCcTimeoutMs = core.Float64Ptr(float64(300000))
				editSettingsOptionsModel.FabricLcInstallCcTimeoutMs = core.Float64Ptr(float64(300000))
				editSettingsOptionsModel.FabricLcGetCcTimeoutMs = core.Float64Ptr(float64(180000))
				editSettingsOptionsModel.FabricGeneralTimeoutMs = core.Float64Ptr(float64(10000))
				editSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.EditSettings(editSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.EditSettingsWithContext(ctx, editSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.EditSettings(editSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.EditSettingsWithContext(ctx, editSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke EditSettings with error: Operation request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the EditSettingsBodyInactivityTimeouts model
				editSettingsBodyInactivityTimeoutsModel := new(blockchainv3.EditSettingsBodyInactivityTimeouts)
				editSettingsBodyInactivityTimeoutsModel.Enabled = core.BoolPtr(false)
				editSettingsBodyInactivityTimeoutsModel.MaxIdleTime = core.Float64Ptr(float64(90000))

				// Construct an instance of the LoggingSettingsClient model
				loggingSettingsClientModel := new(blockchainv3.LoggingSettingsClient)
				loggingSettingsClientModel.Enabled = core.BoolPtr(true)
				loggingSettingsClientModel.Level = core.StringPtr("silly")
				loggingSettingsClientModel.UniqueName = core.BoolPtr(false)

				// Construct an instance of the LoggingSettingsServer model
				loggingSettingsServerModel := new(blockchainv3.LoggingSettingsServer)
				loggingSettingsServerModel.Enabled = core.BoolPtr(true)
				loggingSettingsServerModel.Level = core.StringPtr("silly")
				loggingSettingsServerModel.UniqueName = core.BoolPtr(false)

				// Construct an instance of the EditLogSettingsBody model
				editLogSettingsBodyModel := new(blockchainv3.EditLogSettingsBody)
				editLogSettingsBodyModel.Client = loggingSettingsClientModel
				editLogSettingsBodyModel.Server = loggingSettingsServerModel

				// Construct an instance of the EditSettingsOptions model
				editSettingsOptionsModel := new(blockchainv3.EditSettingsOptions)
				editSettingsOptionsModel.InactivityTimeouts = editSettingsBodyInactivityTimeoutsModel
				editSettingsOptionsModel.FileLogging = editLogSettingsBodyModel
				editSettingsOptionsModel.MaxReqPerMin = core.Float64Ptr(float64(25))
				editSettingsOptionsModel.MaxReqPerMinAk = core.Float64Ptr(float64(25))
				editSettingsOptionsModel.FabricGetBlockTimeoutMs = core.Float64Ptr(float64(10000))
				editSettingsOptionsModel.FabricInstantiateTimeoutMs = core.Float64Ptr(float64(300000))
				editSettingsOptionsModel.FabricJoinChannelTimeoutMs = core.Float64Ptr(float64(25000))
				editSettingsOptionsModel.FabricInstallCcTimeoutMs = core.Float64Ptr(float64(300000))
				editSettingsOptionsModel.FabricLcInstallCcTimeoutMs = core.Float64Ptr(float64(300000))
				editSettingsOptionsModel.FabricLcGetCcTimeoutMs = core.Float64Ptr(float64(180000))
				editSettingsOptionsModel.FabricGeneralTimeoutMs = core.Float64Ptr(float64(10000))
				editSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.EditSettings(editSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetFabVersions(getFabVersionsOptions *GetFabVersionsOptions) - Operation response error`, func() {
		getFabVersionsPath := "/ak/api/v3/kubernetes/fabric/versions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getFabVersionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["cache"]).To(Equal([]string{"skip"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetFabVersions with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the GetFabVersionsOptions model
				getFabVersionsOptionsModel := new(blockchainv3.GetFabVersionsOptions)
				getFabVersionsOptionsModel.Cache = core.StringPtr("skip")
				getFabVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.GetFabVersions(getFabVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.GetFabVersions(getFabVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetFabVersions(getFabVersionsOptions *GetFabVersionsOptions)`, func() {
		getFabVersionsPath := "/ak/api/v3/kubernetes/fabric/versions"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getFabVersionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["cache"]).To(Equal([]string{"skip"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"versions": {"ca": {"1.4.6-2": {"default": true, "version": "1.4.6-2", "image": {"anyKey": "anyValue"}}, "2.1.0-0": {"default": true, "version": "1.4.6-2", "image": {"anyKey": "anyValue"}}}, "peer": {"1.4.6-2": {"default": true, "version": "1.4.6-2", "image": {"anyKey": "anyValue"}}, "2.1.0-0": {"default": true, "version": "1.4.6-2", "image": {"anyKey": "anyValue"}}}, "orderer": {"1.4.6-2": {"default": true, "version": "1.4.6-2", "image": {"anyKey": "anyValue"}}, "2.1.0-0": {"default": true, "version": "1.4.6-2", "image": {"anyKey": "anyValue"}}}}}`)
				}))
			})
			It(`Invoke GetFabVersions successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.GetFabVersions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetFabVersionsOptions model
				getFabVersionsOptionsModel := new(blockchainv3.GetFabVersionsOptions)
				getFabVersionsOptionsModel.Cache = core.StringPtr("skip")
				getFabVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.GetFabVersions(getFabVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.GetFabVersionsWithContext(ctx, getFabVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.GetFabVersions(getFabVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.GetFabVersionsWithContext(ctx, getFabVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetFabVersions with error: Operation request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the GetFabVersionsOptions model
				getFabVersionsOptionsModel := new(blockchainv3.GetFabVersionsOptions)
				getFabVersionsOptionsModel.Cache = core.StringPtr("skip")
				getFabVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.GetFabVersions(getFabVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetHealth(getHealthOptions *GetHealthOptions) - Operation response error`, func() {
		getHealthPath := "/ak/api/v3/health"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getHealthPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetHealth with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the GetHealthOptions model
				getHealthOptionsModel := new(blockchainv3.GetHealthOptions)
				getHealthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.GetHealth(getHealthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.GetHealth(getHealthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetHealth(getHealthOptions *GetHealthOptions)`, func() {
		getHealthPath := "/ak/api/v3/health"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getHealthPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"OPTOOLS": {"instance_id": "p59ta", "now": 1542746836056, "born": 1542746836056, "up_time": "30 days", "memory_usage": {"rss": "56.1 MB", "heapTotal": "34.4 MB", "heapUsed": "28.4 MB", "external": "369.3 KB"}, "session_cache_stats": {"hits": 42, "misses": 11, "keys": 100, "cache_size": "4.19 KiB"}, "couch_cache_stats": {"hits": 42, "misses": 11, "keys": 100, "cache_size": "4.19 KiB"}, "iam_cache_stats": {"hits": 42, "misses": 11, "keys": 100, "cache_size": "4.19 KiB"}, "proxy_cache": {"hits": 42, "misses": 11, "keys": 100, "cache_size": "4.19 KiB"}}, "OS": {"arch": "x64", "type": "Windows_NT", "endian": "LE", "loadavg": [0], "cpus": [{"model": "Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz", "speed": 2592, "times": {"idle": 131397203, "irq": 6068640, "nice": 0, "sys": 9652328, "user": 4152187}}], "total_memory": "31.7 GB", "free_memory": "21.9 GB", "up_time": "4.9 days"}}`)
				}))
			})
			It(`Invoke GetHealth successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.GetHealth(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetHealthOptions model
				getHealthOptionsModel := new(blockchainv3.GetHealthOptions)
				getHealthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.GetHealth(getHealthOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.GetHealthWithContext(ctx, getHealthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.GetHealth(getHealthOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.GetHealthWithContext(ctx, getHealthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetHealth with error: Operation request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the GetHealthOptions model
				getHealthOptionsModel := new(blockchainv3.GetHealthOptions)
				getHealthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.GetHealth(getHealthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListNotifications(listNotificationsOptions *ListNotificationsOptions) - Operation response error`, func() {
		listNotificationsPath := "/ak/api/v3/notifications"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listNotificationsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for limit query parameter


					// TODO: Add check for skip query parameter

					Expect(req.URL.Query()["component_id"]).To(Equal([]string{"MyPeer"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListNotifications with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the ListNotificationsOptions model
				listNotificationsOptionsModel := new(blockchainv3.ListNotificationsOptions)
				listNotificationsOptionsModel.Limit = core.Float64Ptr(float64(1))
				listNotificationsOptionsModel.Skip = core.Float64Ptr(float64(1))
				listNotificationsOptionsModel.ComponentID = core.StringPtr("MyPeer")
				listNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.ListNotifications(listNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.ListNotifications(listNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListNotifications(listNotificationsOptions *ListNotificationsOptions)`, func() {
		listNotificationsPath := "/ak/api/v3/notifications"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listNotificationsPath))
					Expect(req.Method).To(Equal("GET"))


					// TODO: Add check for limit query parameter


					// TODO: Add check for skip query parameter

					Expect(req.URL.Query()["component_id"]).To(Equal([]string{"MyPeer"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total": 10, "returning": 3, "notifications": [{"id": "60d84819bfa17adb4174ff3a1c52b5d6", "type": "notification", "status": "pending", "by": "d******a@us.ibm.com", "message": "Restarting application", "ts_display": 1537262855753}]}`)
				}))
			})
			It(`Invoke ListNotifications successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.ListNotifications(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListNotificationsOptions model
				listNotificationsOptionsModel := new(blockchainv3.ListNotificationsOptions)
				listNotificationsOptionsModel.Limit = core.Float64Ptr(float64(1))
				listNotificationsOptionsModel.Skip = core.Float64Ptr(float64(1))
				listNotificationsOptionsModel.ComponentID = core.StringPtr("MyPeer")
				listNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.ListNotifications(listNotificationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.ListNotificationsWithContext(ctx, listNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.ListNotifications(listNotificationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.ListNotificationsWithContext(ctx, listNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListNotifications with error: Operation request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the ListNotificationsOptions model
				listNotificationsOptionsModel := new(blockchainv3.ListNotificationsOptions)
				listNotificationsOptionsModel.Limit = core.Float64Ptr(float64(1))
				listNotificationsOptionsModel.Skip = core.Float64Ptr(float64(1))
				listNotificationsOptionsModel.ComponentID = core.StringPtr("MyPeer")
				listNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.ListNotifications(listNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteSigTx(deleteSigTxOptions *DeleteSigTxOptions) - Operation response error`, func() {
		deleteSigTxPath := "/ak/api/v3/signature_collections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSigTxPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteSigTx with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the DeleteSigTxOptions model
				deleteSigTxOptionsModel := new(blockchainv3.DeleteSigTxOptions)
				deleteSigTxOptionsModel.ID = core.StringPtr("testString")
				deleteSigTxOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.DeleteSigTx(deleteSigTxOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.DeleteSigTx(deleteSigTxOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteSigTx(deleteSigTxOptions *DeleteSigTxOptions)`, func() {
		deleteSigTxPath := "/ak/api/v3/signature_collections/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSigTxPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"message": "ok", "tx_id": "abcde"}`)
				}))
			})
			It(`Invoke DeleteSigTx successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.DeleteSigTx(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteSigTxOptions model
				deleteSigTxOptionsModel := new(blockchainv3.DeleteSigTxOptions)
				deleteSigTxOptionsModel.ID = core.StringPtr("testString")
				deleteSigTxOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.DeleteSigTx(deleteSigTxOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.DeleteSigTxWithContext(ctx, deleteSigTxOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.DeleteSigTx(deleteSigTxOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.DeleteSigTxWithContext(ctx, deleteSigTxOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke DeleteSigTx with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the DeleteSigTxOptions model
				deleteSigTxOptionsModel := new(blockchainv3.DeleteSigTxOptions)
				deleteSigTxOptionsModel.ID = core.StringPtr("testString")
				deleteSigTxOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.DeleteSigTx(deleteSigTxOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteSigTxOptions model with no property values
				deleteSigTxOptionsModelNew := new(blockchainv3.DeleteSigTxOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.DeleteSigTx(deleteSigTxOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ArchiveNotifications(archiveNotificationsOptions *ArchiveNotificationsOptions) - Operation response error`, func() {
		archiveNotificationsPath := "/ak/api/v3/notifications/bulk"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(archiveNotificationsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ArchiveNotifications with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the ArchiveNotificationsOptions model
				archiveNotificationsOptionsModel := new(blockchainv3.ArchiveNotificationsOptions)
				archiveNotificationsOptionsModel.NotificationIds = []string{"c9d00ebf849051e4f102008dc0be2488"}
				archiveNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.ArchiveNotifications(archiveNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.ArchiveNotifications(archiveNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ArchiveNotifications(archiveNotificationsOptions *ArchiveNotificationsOptions)`, func() {
		archiveNotificationsPath := "/ak/api/v3/notifications/bulk"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(archiveNotificationsPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"message": "ok", "details": "archived 3 notification(s)"}`)
				}))
			})
			It(`Invoke ArchiveNotifications successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.ArchiveNotifications(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ArchiveNotificationsOptions model
				archiveNotificationsOptionsModel := new(blockchainv3.ArchiveNotificationsOptions)
				archiveNotificationsOptionsModel.NotificationIds = []string{"c9d00ebf849051e4f102008dc0be2488"}
				archiveNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.ArchiveNotifications(archiveNotificationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.ArchiveNotificationsWithContext(ctx, archiveNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.ArchiveNotifications(archiveNotificationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.ArchiveNotificationsWithContext(ctx, archiveNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ArchiveNotifications with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the ArchiveNotificationsOptions model
				archiveNotificationsOptionsModel := new(blockchainv3.ArchiveNotificationsOptions)
				archiveNotificationsOptionsModel.NotificationIds = []string{"c9d00ebf849051e4f102008dc0be2488"}
				archiveNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.ArchiveNotifications(archiveNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ArchiveNotificationsOptions model with no property values
				archiveNotificationsOptionsModelNew := new(blockchainv3.ArchiveNotificationsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = blockchainService.ArchiveNotifications(archiveNotificationsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Restart(restartOptions *RestartOptions) - Operation response error`, func() {
		restartPath := "/ak/api/v3/restart"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(restartPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Restart with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the RestartOptions model
				restartOptionsModel := new(blockchainv3.RestartOptions)
				restartOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.Restart(restartOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.Restart(restartOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`Restart(restartOptions *RestartOptions)`, func() {
		restartPath := "/ak/api/v3/restart"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(restartPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"message": "restarting - give me 5-30 seconds"}`)
				}))
			})
			It(`Invoke Restart successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.Restart(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RestartOptions model
				restartOptionsModel := new(blockchainv3.RestartOptions)
				restartOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.Restart(restartOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.RestartWithContext(ctx, restartOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.Restart(restartOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.RestartWithContext(ctx, restartOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke Restart with error: Operation request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the RestartOptions model
				restartOptionsModel := new(blockchainv3.RestartOptions)
				restartOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.Restart(restartOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteAllSessions(deleteAllSessionsOptions *DeleteAllSessionsOptions) - Operation response error`, func() {
		deleteAllSessionsPath := "/ak/api/v3/sessions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAllSessionsPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteAllSessions with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the DeleteAllSessionsOptions model
				deleteAllSessionsOptionsModel := new(blockchainv3.DeleteAllSessionsOptions)
				deleteAllSessionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.DeleteAllSessions(deleteAllSessionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.DeleteAllSessions(deleteAllSessionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteAllSessions(deleteAllSessionsOptions *DeleteAllSessionsOptions)`, func() {
		deleteAllSessionsPath := "/ak/api/v3/sessions"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAllSessionsPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"message": "delete submitted"}`)
				}))
			})
			It(`Invoke DeleteAllSessions successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.DeleteAllSessions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteAllSessionsOptions model
				deleteAllSessionsOptionsModel := new(blockchainv3.DeleteAllSessionsOptions)
				deleteAllSessionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.DeleteAllSessions(deleteAllSessionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.DeleteAllSessionsWithContext(ctx, deleteAllSessionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.DeleteAllSessions(deleteAllSessionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.DeleteAllSessionsWithContext(ctx, deleteAllSessionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke DeleteAllSessions with error: Operation request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the DeleteAllSessionsOptions model
				deleteAllSessionsOptionsModel := new(blockchainv3.DeleteAllSessionsOptions)
				deleteAllSessionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.DeleteAllSessions(deleteAllSessionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteAllNotifications(deleteAllNotificationsOptions *DeleteAllNotificationsOptions) - Operation response error`, func() {
		deleteAllNotificationsPath := "/ak/api/v3/notifications/purge"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAllNotificationsPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteAllNotifications with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the DeleteAllNotificationsOptions model
				deleteAllNotificationsOptionsModel := new(blockchainv3.DeleteAllNotificationsOptions)
				deleteAllNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.DeleteAllNotifications(deleteAllNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.DeleteAllNotifications(deleteAllNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteAllNotifications(deleteAllNotificationsOptions *DeleteAllNotificationsOptions)`, func() {
		deleteAllNotificationsPath := "/ak/api/v3/notifications/purge"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAllNotificationsPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"message": "ok", "details": "deleted 101 notification(s)"}`)
				}))
			})
			It(`Invoke DeleteAllNotifications successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.DeleteAllNotifications(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteAllNotificationsOptions model
				deleteAllNotificationsOptionsModel := new(blockchainv3.DeleteAllNotificationsOptions)
				deleteAllNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.DeleteAllNotifications(deleteAllNotificationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.DeleteAllNotificationsWithContext(ctx, deleteAllNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.DeleteAllNotifications(deleteAllNotificationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.DeleteAllNotificationsWithContext(ctx, deleteAllNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke DeleteAllNotifications with error: Operation request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the DeleteAllNotificationsOptions model
				deleteAllNotificationsOptionsModel := new(blockchainv3.DeleteAllNotificationsOptions)
				deleteAllNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.DeleteAllNotifications(deleteAllNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ClearCaches(clearCachesOptions *ClearCachesOptions) - Operation response error`, func() {
		clearCachesPath := "/ak/api/v3/cache"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(clearCachesPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ClearCaches with error: Operation response processing error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the ClearCachesOptions model
				clearCachesOptionsModel := new(blockchainv3.ClearCachesOptions)
				clearCachesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := blockchainService.ClearCaches(clearCachesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				blockchainService.EnableRetries(0, 0)
				result, response, operationErr = blockchainService.ClearCaches(clearCachesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ClearCaches(clearCachesOptions *ClearCachesOptions)`, func() {
		clearCachesPath := "/ak/api/v3/cache"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(clearCachesPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"message": "ok", "flushed": ["iam_cache"]}`)
				}))
			})
			It(`Invoke ClearCaches successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.ClearCaches(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ClearCachesOptions model
				clearCachesOptionsModel := new(blockchainv3.ClearCachesOptions)
				clearCachesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.ClearCaches(clearCachesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.ClearCachesWithContext(ctx, clearCachesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.ClearCaches(clearCachesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.ClearCachesWithContext(ctx, clearCachesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ClearCaches with error: Operation request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the ClearCachesOptions model
				clearCachesOptionsModel := new(blockchainv3.ClearCachesOptions)
				clearCachesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.ClearCaches(clearCachesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(blockchainService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(blockchainService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
				URL: "https://blockchainv3/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(blockchainService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BLOCKCHAIN_URL": "https://blockchainv3/api",
				"BLOCKCHAIN_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3UsingExternalConfig(&blockchainv3.BlockchainV3Options{
				})
				Expect(blockchainService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := blockchainService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != blockchainService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(blockchainService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(blockchainService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3UsingExternalConfig(&blockchainv3.BlockchainV3Options{
					URL: "https://testService/api",
				})
				Expect(blockchainService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := blockchainService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != blockchainService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(blockchainService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(blockchainService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3UsingExternalConfig(&blockchainv3.BlockchainV3Options{
				})
				err := blockchainService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := blockchainService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != blockchainService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(blockchainService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(blockchainService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BLOCKCHAIN_URL": "https://blockchainv3/api",
				"BLOCKCHAIN_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			blockchainService, serviceErr := blockchainv3.NewBlockchainV3UsingExternalConfig(&blockchainv3.BlockchainV3Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(blockchainService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BLOCKCHAIN_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			blockchainService, serviceErr := blockchainv3.NewBlockchainV3UsingExternalConfig(&blockchainv3.BlockchainV3Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(blockchainService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = blockchainv3.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})

	Describe(`GetPostman(getPostmanOptions *GetPostmanOptions)`, func() {
		getPostmanPath := "/ak/api/v3/postman"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPostmanPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["auth_type"]).To(Equal([]string{"bearer"}))

					Expect(req.URL.Query()["token"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["api_key"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["username"]).To(Equal([]string{"admin"}))

					Expect(req.URL.Query()["password"]).To(Equal([]string{"password"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetPostman successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := blockchainService.GetPostman(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the GetPostmanOptions model
				getPostmanOptionsModel := new(blockchainv3.GetPostmanOptions)
				getPostmanOptionsModel.AuthType = core.StringPtr("bearer")
				getPostmanOptionsModel.Token = core.StringPtr("testString")
				getPostmanOptionsModel.ApiKey = core.StringPtr("testString")
				getPostmanOptionsModel.Username = core.StringPtr("admin")
				getPostmanOptionsModel.Password = core.StringPtr("password")
				getPostmanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = blockchainService.GetPostman(getPostmanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				blockchainService.DisableRetries()
				response, operationErr = blockchainService.GetPostman(getPostmanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke GetPostman with error: Operation validation and request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the GetPostmanOptions model
				getPostmanOptionsModel := new(blockchainv3.GetPostmanOptions)
				getPostmanOptionsModel.AuthType = core.StringPtr("bearer")
				getPostmanOptionsModel.Token = core.StringPtr("testString")
				getPostmanOptionsModel.ApiKey = core.StringPtr("testString")
				getPostmanOptionsModel.Username = core.StringPtr("admin")
				getPostmanOptionsModel.Password = core.StringPtr("password")
				getPostmanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := blockchainService.GetPostman(getPostmanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the GetPostmanOptions model with no property values
				getPostmanOptionsModelNew := new(blockchainv3.GetPostmanOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = blockchainService.GetPostman(getPostmanOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSwagger(getSwaggerOptions *GetSwaggerOptions)`, func() {
		getSwaggerPath := "/ak/api/v3/openapi"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSwaggerPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "text/plain")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `"OperationResponse"`)
				}))
			})
			It(`Invoke GetSwagger successfully`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())
				blockchainService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := blockchainService.GetSwagger(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSwaggerOptions model
				getSwaggerOptionsModel := new(blockchainv3.GetSwaggerOptions)
				getSwaggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = blockchainService.GetSwagger(getSwaggerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.GetSwaggerWithContext(ctx, getSwaggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				blockchainService.DisableRetries()
				result, response, operationErr = blockchainService.GetSwagger(getSwaggerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = blockchainService.GetSwaggerWithContext(ctx, getSwaggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetSwagger with error: Operation request error`, func() {
				blockchainService, serviceErr := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(blockchainService).ToNot(BeNil())

				// Construct an instance of the GetSwaggerOptions model
				getSwaggerOptionsModel := new(blockchainv3.GetSwaggerOptions)
				getSwaggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := blockchainService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := blockchainService.GetSwagger(getSwaggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			blockchainService, _ := blockchainv3.NewBlockchainV3(&blockchainv3.BlockchainV3Options{
				URL:           "http://blockchainv3modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewArchiveNotificationsOptions successfully`, func() {
				// Construct an instance of the ArchiveNotificationsOptions model
				archiveNotificationsOptionsNotificationIds := []string{"c9d00ebf849051e4f102008dc0be2488"}
				archiveNotificationsOptionsModel := blockchainService.NewArchiveNotificationsOptions(archiveNotificationsOptionsNotificationIds)
				archiveNotificationsOptionsModel.SetNotificationIds([]string{"c9d00ebf849051e4f102008dc0be2488"})
				archiveNotificationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(archiveNotificationsOptionsModel).ToNot(BeNil())
				Expect(archiveNotificationsOptionsModel.NotificationIds).To(Equal([]string{"c9d00ebf849051e4f102008dc0be2488"}))
				Expect(archiveNotificationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewBccspPKCS11 successfully`, func() {
				label := "testString"
				pin := "testString"
				model, err := blockchainService.NewBccspPKCS11(label, pin)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewBccspSW successfully`, func() {
				hash := "SHA2"
				security := float64(256)
				model, err := blockchainService.NewBccspSW(hash, security)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCaActionOptions successfully`, func() {
				// Construct an instance of the ActionRenew model
				actionRenewModel := new(blockchainv3.ActionRenew)
				Expect(actionRenewModel).ToNot(BeNil())
				actionRenewModel.TlsCert = core.BoolPtr(true)
				Expect(actionRenewModel.TlsCert).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the CaActionOptions model
				id := "testString"
				caActionOptionsModel := blockchainService.NewCaActionOptions(id)
				caActionOptionsModel.SetID("testString")
				caActionOptionsModel.SetRestart(true)
				caActionOptionsModel.SetRenew(actionRenewModel)
				caActionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(caActionOptionsModel).ToNot(BeNil())
				Expect(caActionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(caActionOptionsModel.Restart).To(Equal(core.BoolPtr(true)))
				Expect(caActionOptionsModel.Renew).To(Equal(actionRenewModel))
				Expect(caActionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewClearCachesOptions successfully`, func() {
				// Construct an instance of the ClearCachesOptions model
				clearCachesOptionsModel := blockchainService.NewClearCachesOptions()
				clearCachesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(clearCachesOptionsModel).ToNot(BeNil())
				Expect(clearCachesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewConfigCACfgIdentities successfully`, func() {
				passwordattempts := float64(10)
				model, err := blockchainService.NewConfigCACfgIdentities(passwordattempts)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCACreate successfully`, func() {
				var registry *blockchainv3.ConfigCARegistry = nil
				_, err := blockchainService.NewConfigCACreate(registry)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewConfigCACsrKeyrequest successfully`, func() {
				algo := "ecdsa"
				size := float64(256)
				model, err := blockchainService.NewConfigCACsrKeyrequest(algo, size)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCACsrNamesItem successfully`, func() {
				c := "US"
				st := "North Carolina"
				o := "Hyperledger"
				model, err := blockchainService.NewConfigCACsrNamesItem(c, st, o)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCADbTlsClient successfully`, func() {
				certfile := "testString"
				keyfile := "testString"
				model, err := blockchainService.NewConfigCADbTlsClient(certfile, keyfile)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCAIntermediateEnrollment successfully`, func() {
				hosts := "localhost"
				profile := "testString"
				label := "testString"
				model, err := blockchainService.NewConfigCAIntermediateEnrollment(hosts, profile, label)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCAIntermediateParentserver successfully`, func() {
				url := "testString"
				caname := "testString"
				model, err := blockchainService.NewConfigCAIntermediateParentserver(url, caname)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCAIntermediateTls successfully`, func() {
				certfiles := []string{"testString"}
				model, err := blockchainService.NewConfigCAIntermediateTls(certfiles)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCAIntermediateTlsClient successfully`, func() {
				certfile := "testString"
				keyfile := "testString"
				model, err := blockchainService.NewConfigCAIntermediateTlsClient(certfile, keyfile)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCARegistryIdentitiesItem successfully`, func() {
				name := "admin"
				pass := "password"
				typeVar := "client"
				model, err := blockchainService.NewConfigCARegistryIdentitiesItem(name, pass, typeVar)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCATlsClientauth successfully`, func() {
				typeVar := "noclientcert"
				certfiles := []string{"testString"}
				model, err := blockchainService.NewConfigCATlsClientauth(typeVar, certfiles)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCACfg successfully`, func() {
				var identities *blockchainv3.ConfigCACfgIdentities = nil
				_, err := blockchainService.NewConfigCACfg(identities)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewConfigCACors successfully`, func() {
				enabled := true
				origins := []string{"*"}
				model, err := blockchainService.NewConfigCACors(enabled, origins)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCACrl successfully`, func() {
				expiry := "24h"
				model, err := blockchainService.NewConfigCACrl(expiry)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCACsr successfully`, func() {
				cn := "ca"
				names := []blockchainv3.ConfigCACsrNamesItem{}
				var ca *blockchainv3.ConfigCACsrCa = nil
				_, err := blockchainService.NewConfigCACsr(cn, names, ca)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewConfigCADb successfully`, func() {
				typeVar := "postgres"
				datasource := "host=fake.databases.appdomain.cloud port=31941 user=ibm_cloud password=password dbname=ibmclouddb sslmode=verify-full"
				model, err := blockchainService.NewConfigCADb(typeVar, datasource)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCAIdemix successfully`, func() {
				rhpoolsize := float64(100)
				nonceexpiration := "15s"
				noncesweepinterval := "15m"
				model, err := blockchainService.NewConfigCAIdemix(rhpoolsize, nonceexpiration, noncesweepinterval)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCAIntermediate successfully`, func() {
				var parentserver *blockchainv3.ConfigCAIntermediateParentserver = nil
				_, err := blockchainService.NewConfigCAIntermediate(parentserver)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewConfigCARegistry successfully`, func() {
				maxenrollments := float64(-1)
				identities := []blockchainv3.ConfigCARegistryIdentitiesItem{}
				model, err := blockchainService.NewConfigCARegistry(maxenrollments, identities)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCATls successfully`, func() {
				keyfile := "testString"
				certfile := "testString"
				model, err := blockchainService.NewConfigCATls(keyfile, certfile)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigPeerAdminService successfully`, func() {
				listenAddress := "0.0.0.0:7051"
				model, err := blockchainService.NewConfigPeerAdminService(listenAddress)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigPeerAuthentication successfully`, func() {
				timewindow := "15m"
				model, err := blockchainService.NewConfigPeerAuthentication(timewindow)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigPeerClient successfully`, func() {
				connTimeout := "2s"
				model, err := blockchainService.NewConfigPeerClient(connTimeout)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateCaBodyConfigOverride successfully`, func() {
				var ca *blockchainv3.ConfigCACreate = nil
				_, err := blockchainService.NewCreateCaBodyConfigOverride(ca)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCreateCaBodyResources successfully`, func() {
				var ca *blockchainv3.ResourceObject = nil
				_, err := blockchainService.NewCreateCaBodyResources(ca)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCreateCaBodyStorage successfully`, func() {
				var ca *blockchainv3.StorageObject = nil
				_, err := blockchainService.NewCreateCaBodyStorage(ca)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCreateCaOptions successfully`, func() {
				// Construct an instance of the ConfigCACors model
				configCaCorsModel := new(blockchainv3.ConfigCACors)
				Expect(configCaCorsModel).ToNot(BeNil())
				configCaCorsModel.Enabled = core.BoolPtr(true)
				configCaCorsModel.Origins = []string{"*"}
				Expect(configCaCorsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(configCaCorsModel.Origins).To(Equal([]string{"*"}))

				// Construct an instance of the ConfigCATlsClientauth model
				configCaTlsClientauthModel := new(blockchainv3.ConfigCATlsClientauth)
				Expect(configCaTlsClientauthModel).ToNot(BeNil())
				configCaTlsClientauthModel.Type = core.StringPtr("noclientcert")
				configCaTlsClientauthModel.Certfiles = []string{"testString"}
				Expect(configCaTlsClientauthModel.Type).To(Equal(core.StringPtr("noclientcert")))
				Expect(configCaTlsClientauthModel.Certfiles).To(Equal([]string{"testString"}))

				// Construct an instance of the ConfigCATls model
				configCaTlsModel := new(blockchainv3.ConfigCATls)
				Expect(configCaTlsModel).ToNot(BeNil())
				configCaTlsModel.Keyfile = core.StringPtr("testString")
				configCaTlsModel.Certfile = core.StringPtr("testString")
				configCaTlsModel.Clientauth = configCaTlsClientauthModel
				Expect(configCaTlsModel.Keyfile).To(Equal(core.StringPtr("testString")))
				Expect(configCaTlsModel.Certfile).To(Equal(core.StringPtr("testString")))
				Expect(configCaTlsModel.Clientauth).To(Equal(configCaTlsClientauthModel))

				// Construct an instance of the ConfigCACa model
				configCaCaModel := new(blockchainv3.ConfigCACa)
				Expect(configCaCaModel).ToNot(BeNil())
				configCaCaModel.Keyfile = core.StringPtr("testString")
				configCaCaModel.Certfile = core.StringPtr("testString")
				configCaCaModel.Chainfile = core.StringPtr("testString")
				Expect(configCaCaModel.Keyfile).To(Equal(core.StringPtr("testString")))
				Expect(configCaCaModel.Certfile).To(Equal(core.StringPtr("testString")))
				Expect(configCaCaModel.Chainfile).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ConfigCACrl model
				configCaCrlModel := new(blockchainv3.ConfigCACrl)
				Expect(configCaCrlModel).ToNot(BeNil())
				configCaCrlModel.Expiry = core.StringPtr("24h")
				Expect(configCaCrlModel.Expiry).To(Equal(core.StringPtr("24h")))

				// Construct an instance of the IdentityAttrs model
				identityAttrsModel := new(blockchainv3.IdentityAttrs)
				Expect(identityAttrsModel).ToNot(BeNil())
				identityAttrsModel.HfRegistrarRoles = core.StringPtr("*")
				identityAttrsModel.HfRegistrarDelegateRoles = core.StringPtr("*")
				identityAttrsModel.HfRevoker = core.BoolPtr(true)
				identityAttrsModel.HfIntermediateCA = core.BoolPtr(true)
				identityAttrsModel.HfGenCRL = core.BoolPtr(true)
				identityAttrsModel.HfRegistrarAttributes = core.StringPtr("*")
				identityAttrsModel.HfAffiliationMgr = core.BoolPtr(true)
				Expect(identityAttrsModel.HfRegistrarRoles).To(Equal(core.StringPtr("*")))
				Expect(identityAttrsModel.HfRegistrarDelegateRoles).To(Equal(core.StringPtr("*")))
				Expect(identityAttrsModel.HfRevoker).To(Equal(core.BoolPtr(true)))
				Expect(identityAttrsModel.HfIntermediateCA).To(Equal(core.BoolPtr(true)))
				Expect(identityAttrsModel.HfGenCRL).To(Equal(core.BoolPtr(true)))
				Expect(identityAttrsModel.HfRegistrarAttributes).To(Equal(core.StringPtr("*")))
				Expect(identityAttrsModel.HfAffiliationMgr).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the ConfigCARegistryIdentitiesItem model
				configCaRegistryIdentitiesItemModel := new(blockchainv3.ConfigCARegistryIdentitiesItem)
				Expect(configCaRegistryIdentitiesItemModel).ToNot(BeNil())
				configCaRegistryIdentitiesItemModel.Name = core.StringPtr("admin")
				configCaRegistryIdentitiesItemModel.Pass = core.StringPtr("password")
				configCaRegistryIdentitiesItemModel.Type = core.StringPtr("client")
				configCaRegistryIdentitiesItemModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryIdentitiesItemModel.Affiliation = core.StringPtr("testString")
				configCaRegistryIdentitiesItemModel.Attrs = identityAttrsModel
				Expect(configCaRegistryIdentitiesItemModel.Name).To(Equal(core.StringPtr("admin")))
				Expect(configCaRegistryIdentitiesItemModel.Pass).To(Equal(core.StringPtr("password")))
				Expect(configCaRegistryIdentitiesItemModel.Type).To(Equal(core.StringPtr("client")))
				Expect(configCaRegistryIdentitiesItemModel.Maxenrollments).To(Equal(core.Float64Ptr(float64(-1))))
				Expect(configCaRegistryIdentitiesItemModel.Affiliation).To(Equal(core.StringPtr("testString")))
				Expect(configCaRegistryIdentitiesItemModel.Attrs).To(Equal(identityAttrsModel))

				// Construct an instance of the ConfigCARegistry model
				configCaRegistryModel := new(blockchainv3.ConfigCARegistry)
				Expect(configCaRegistryModel).ToNot(BeNil())
				configCaRegistryModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryModel.Identities = []blockchainv3.ConfigCARegistryIdentitiesItem{*configCaRegistryIdentitiesItemModel}
				Expect(configCaRegistryModel.Maxenrollments).To(Equal(core.Float64Ptr(float64(-1))))
				Expect(configCaRegistryModel.Identities).To(Equal([]blockchainv3.ConfigCARegistryIdentitiesItem{*configCaRegistryIdentitiesItemModel}))

				// Construct an instance of the ConfigCADbTlsClient model
				configCaDbTlsClientModel := new(blockchainv3.ConfigCADbTlsClient)
				Expect(configCaDbTlsClientModel).ToNot(BeNil())
				configCaDbTlsClientModel.Certfile = core.StringPtr("testString")
				configCaDbTlsClientModel.Keyfile = core.StringPtr("testString")
				Expect(configCaDbTlsClientModel.Certfile).To(Equal(core.StringPtr("testString")))
				Expect(configCaDbTlsClientModel.Keyfile).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ConfigCADbTls model
				configCaDbTlsModel := new(blockchainv3.ConfigCADbTls)
				Expect(configCaDbTlsModel).ToNot(BeNil())
				configCaDbTlsModel.Certfiles = []string{"testString"}
				configCaDbTlsModel.Client = configCaDbTlsClientModel
				configCaDbTlsModel.Enabled = core.BoolPtr(false)
				Expect(configCaDbTlsModel.Certfiles).To(Equal([]string{"testString"}))
				Expect(configCaDbTlsModel.Client).To(Equal(configCaDbTlsClientModel))
				Expect(configCaDbTlsModel.Enabled).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the ConfigCADb model
				configCaDbModel := new(blockchainv3.ConfigCADb)
				Expect(configCaDbModel).ToNot(BeNil())
				configCaDbModel.Type = core.StringPtr("postgres")
				configCaDbModel.Datasource = core.StringPtr("host=fake.databases.appdomain.cloud port=31941 user=ibm_cloud password=password dbname=ibmclouddb sslmode=verify-full")
				configCaDbModel.Tls = configCaDbTlsModel
				Expect(configCaDbModel.Type).To(Equal(core.StringPtr("postgres")))
				Expect(configCaDbModel.Datasource).To(Equal(core.StringPtr("host=fake.databases.appdomain.cloud port=31941 user=ibm_cloud password=password dbname=ibmclouddb sslmode=verify-full")))
				Expect(configCaDbModel.Tls).To(Equal(configCaDbTlsModel))

				// Construct an instance of the ConfigCAAffiliations model
				configCaAffiliationsModel := new(blockchainv3.ConfigCAAffiliations)
				Expect(configCaAffiliationsModel).ToNot(BeNil())
				configCaAffiliationsModel.Org1 = []string{"department1"}
				configCaAffiliationsModel.Org2 = []string{"department1"}
				configCaAffiliationsModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(configCaAffiliationsModel.Org1).To(Equal([]string{"department1"}))
				Expect(configCaAffiliationsModel.Org2).To(Equal([]string{"department1"}))
				Expect(configCaAffiliationsModel.GetProperties()).ToNot(BeEmpty())
				Expect(configCaAffiliationsModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ConfigCACsrKeyrequest model
				configCaCsrKeyrequestModel := new(blockchainv3.ConfigCACsrKeyrequest)
				Expect(configCaCsrKeyrequestModel).ToNot(BeNil())
				configCaCsrKeyrequestModel.Algo = core.StringPtr("ecdsa")
				configCaCsrKeyrequestModel.Size = core.Float64Ptr(float64(256))
				Expect(configCaCsrKeyrequestModel.Algo).To(Equal(core.StringPtr("ecdsa")))
				Expect(configCaCsrKeyrequestModel.Size).To(Equal(core.Float64Ptr(float64(256))))

				// Construct an instance of the ConfigCACsrNamesItem model
				configCaCsrNamesItemModel := new(blockchainv3.ConfigCACsrNamesItem)
				Expect(configCaCsrNamesItemModel).ToNot(BeNil())
				configCaCsrNamesItemModel.C = core.StringPtr("US")
				configCaCsrNamesItemModel.ST = core.StringPtr("North Carolina")
				configCaCsrNamesItemModel.L = core.StringPtr("Raleigh")
				configCaCsrNamesItemModel.O = core.StringPtr("Hyperledger")
				configCaCsrNamesItemModel.OU = core.StringPtr("Fabric")
				Expect(configCaCsrNamesItemModel.C).To(Equal(core.StringPtr("US")))
				Expect(configCaCsrNamesItemModel.ST).To(Equal(core.StringPtr("North Carolina")))
				Expect(configCaCsrNamesItemModel.L).To(Equal(core.StringPtr("Raleigh")))
				Expect(configCaCsrNamesItemModel.O).To(Equal(core.StringPtr("Hyperledger")))
				Expect(configCaCsrNamesItemModel.OU).To(Equal(core.StringPtr("Fabric")))

				// Construct an instance of the ConfigCACsrCa model
				configCaCsrCaModel := new(blockchainv3.ConfigCACsrCa)
				Expect(configCaCsrCaModel).ToNot(BeNil())
				configCaCsrCaModel.Expiry = core.StringPtr("131400h")
				configCaCsrCaModel.Pathlength = core.Float64Ptr(float64(0))
				Expect(configCaCsrCaModel.Expiry).To(Equal(core.StringPtr("131400h")))
				Expect(configCaCsrCaModel.Pathlength).To(Equal(core.Float64Ptr(float64(0))))

				// Construct an instance of the ConfigCACsr model
				configCaCsrModel := new(blockchainv3.ConfigCACsr)
				Expect(configCaCsrModel).ToNot(BeNil())
				configCaCsrModel.Cn = core.StringPtr("ca")
				configCaCsrModel.Keyrequest = configCaCsrKeyrequestModel
				configCaCsrModel.Names = []blockchainv3.ConfigCACsrNamesItem{*configCaCsrNamesItemModel}
				configCaCsrModel.Hosts = []string{"localhost"}
				configCaCsrModel.Ca = configCaCsrCaModel
				Expect(configCaCsrModel.Cn).To(Equal(core.StringPtr("ca")))
				Expect(configCaCsrModel.Keyrequest).To(Equal(configCaCsrKeyrequestModel))
				Expect(configCaCsrModel.Names).To(Equal([]blockchainv3.ConfigCACsrNamesItem{*configCaCsrNamesItemModel}))
				Expect(configCaCsrModel.Hosts).To(Equal([]string{"localhost"}))
				Expect(configCaCsrModel.Ca).To(Equal(configCaCsrCaModel))

				// Construct an instance of the ConfigCAIdemix model
				configCaIdemixModel := new(blockchainv3.ConfigCAIdemix)
				Expect(configCaIdemixModel).ToNot(BeNil())
				configCaIdemixModel.Rhpoolsize = core.Float64Ptr(float64(100))
				configCaIdemixModel.Nonceexpiration = core.StringPtr("15s")
				configCaIdemixModel.Noncesweepinterval = core.StringPtr("15m")
				Expect(configCaIdemixModel.Rhpoolsize).To(Equal(core.Float64Ptr(float64(100))))
				Expect(configCaIdemixModel.Nonceexpiration).To(Equal(core.StringPtr("15s")))
				Expect(configCaIdemixModel.Noncesweepinterval).To(Equal(core.StringPtr("15m")))

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv3.BccspSW)
				Expect(bccspSwModel).ToNot(BeNil())
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))
				Expect(bccspSwModel.Hash).To(Equal(core.StringPtr("SHA2")))
				Expect(bccspSwModel.Security).To(Equal(core.Float64Ptr(float64(256))))

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv3.BccspPKCS11)
				Expect(bccspPkcS11Model).ToNot(BeNil())
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))
				Expect(bccspPkcS11Model.Label).To(Equal(core.StringPtr("testString")))
				Expect(bccspPkcS11Model.Pin).To(Equal(core.StringPtr("testString")))
				Expect(bccspPkcS11Model.Hash).To(Equal(core.StringPtr("SHA2")))
				Expect(bccspPkcS11Model.Security).To(Equal(core.Float64Ptr(float64(256))))

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv3.Bccsp)
				Expect(bccspModel).ToNot(BeNil())
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model
				Expect(bccspModel.Default).To(Equal(core.StringPtr("SW")))
				Expect(bccspModel.SW).To(Equal(bccspSwModel))
				Expect(bccspModel.PKCS11).To(Equal(bccspPkcS11Model))

				// Construct an instance of the ConfigCAIntermediateParentserver model
				configCaIntermediateParentserverModel := new(blockchainv3.ConfigCAIntermediateParentserver)
				Expect(configCaIntermediateParentserverModel).ToNot(BeNil())
				configCaIntermediateParentserverModel.URL = core.StringPtr("testString")
				configCaIntermediateParentserverModel.Caname = core.StringPtr("testString")
				Expect(configCaIntermediateParentserverModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(configCaIntermediateParentserverModel.Caname).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ConfigCAIntermediateEnrollment model
				configCaIntermediateEnrollmentModel := new(blockchainv3.ConfigCAIntermediateEnrollment)
				Expect(configCaIntermediateEnrollmentModel).ToNot(BeNil())
				configCaIntermediateEnrollmentModel.Hosts = core.StringPtr("localhost")
				configCaIntermediateEnrollmentModel.Profile = core.StringPtr("testString")
				configCaIntermediateEnrollmentModel.Label = core.StringPtr("testString")
				Expect(configCaIntermediateEnrollmentModel.Hosts).To(Equal(core.StringPtr("localhost")))
				Expect(configCaIntermediateEnrollmentModel.Profile).To(Equal(core.StringPtr("testString")))
				Expect(configCaIntermediateEnrollmentModel.Label).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ConfigCAIntermediateTlsClient model
				configCaIntermediateTlsClientModel := new(blockchainv3.ConfigCAIntermediateTlsClient)
				Expect(configCaIntermediateTlsClientModel).ToNot(BeNil())
				configCaIntermediateTlsClientModel.Certfile = core.StringPtr("testString")
				configCaIntermediateTlsClientModel.Keyfile = core.StringPtr("testString")
				Expect(configCaIntermediateTlsClientModel.Certfile).To(Equal(core.StringPtr("testString")))
				Expect(configCaIntermediateTlsClientModel.Keyfile).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ConfigCAIntermediateTls model
				configCaIntermediateTlsModel := new(blockchainv3.ConfigCAIntermediateTls)
				Expect(configCaIntermediateTlsModel).ToNot(BeNil())
				configCaIntermediateTlsModel.Certfiles = []string{"testString"}
				configCaIntermediateTlsModel.Client = configCaIntermediateTlsClientModel
				Expect(configCaIntermediateTlsModel.Certfiles).To(Equal([]string{"testString"}))
				Expect(configCaIntermediateTlsModel.Client).To(Equal(configCaIntermediateTlsClientModel))

				// Construct an instance of the ConfigCAIntermediate model
				configCaIntermediateModel := new(blockchainv3.ConfigCAIntermediate)
				Expect(configCaIntermediateModel).ToNot(BeNil())
				configCaIntermediateModel.Parentserver = configCaIntermediateParentserverModel
				configCaIntermediateModel.Enrollment = configCaIntermediateEnrollmentModel
				configCaIntermediateModel.Tls = configCaIntermediateTlsModel
				Expect(configCaIntermediateModel.Parentserver).To(Equal(configCaIntermediateParentserverModel))
				Expect(configCaIntermediateModel.Enrollment).To(Equal(configCaIntermediateEnrollmentModel))
				Expect(configCaIntermediateModel.Tls).To(Equal(configCaIntermediateTlsModel))

				// Construct an instance of the ConfigCACfgIdentities model
				configCaCfgIdentitiesModel := new(blockchainv3.ConfigCACfgIdentities)
				Expect(configCaCfgIdentitiesModel).ToNot(BeNil())
				configCaCfgIdentitiesModel.Passwordattempts = core.Float64Ptr(float64(10))
				configCaCfgIdentitiesModel.Allowremove = core.BoolPtr(false)
				Expect(configCaCfgIdentitiesModel.Passwordattempts).To(Equal(core.Float64Ptr(float64(10))))
				Expect(configCaCfgIdentitiesModel.Allowremove).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the ConfigCACfg model
				configCaCfgModel := new(blockchainv3.ConfigCACfg)
				Expect(configCaCfgModel).ToNot(BeNil())
				configCaCfgModel.Identities = configCaCfgIdentitiesModel
				Expect(configCaCfgModel.Identities).To(Equal(configCaCfgIdentitiesModel))

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv3.MetricsStatsd)
				Expect(metricsStatsdModel).ToNot(BeNil())
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")
				Expect(metricsStatsdModel.Network).To(Equal(core.StringPtr("udp")))
				Expect(metricsStatsdModel.Address).To(Equal(core.StringPtr("127.0.0.1:8125")))
				Expect(metricsStatsdModel.WriteInterval).To(Equal(core.StringPtr("10s")))
				Expect(metricsStatsdModel.Prefix).To(Equal(core.StringPtr("server")))

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv3.Metrics)
				Expect(metricsModel).ToNot(BeNil())
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel
				Expect(metricsModel.Provider).To(Equal(core.StringPtr("prometheus")))
				Expect(metricsModel.Statsd).To(Equal(metricsStatsdModel))

				// Construct an instance of the ConfigCASigningDefault model
				configCaSigningDefaultModel := new(blockchainv3.ConfigCASigningDefault)
				Expect(configCaSigningDefaultModel).ToNot(BeNil())
				configCaSigningDefaultModel.Usage = []string{"cert sign"}
				configCaSigningDefaultModel.Expiry = core.StringPtr("8760h")
				Expect(configCaSigningDefaultModel.Usage).To(Equal([]string{"cert sign"}))
				Expect(configCaSigningDefaultModel.Expiry).To(Equal(core.StringPtr("8760h")))

				// Construct an instance of the ConfigCASigningProfilesCaCaconstraint model
				configCaSigningProfilesCaCaconstraintModel := new(blockchainv3.ConfigCASigningProfilesCaCaconstraint)
				Expect(configCaSigningProfilesCaCaconstraintModel).ToNot(BeNil())
				configCaSigningProfilesCaCaconstraintModel.Isca = core.BoolPtr(true)
				configCaSigningProfilesCaCaconstraintModel.Maxpathlen = core.Float64Ptr(float64(0))
				configCaSigningProfilesCaCaconstraintModel.Maxpathlenzero = core.BoolPtr(true)
				Expect(configCaSigningProfilesCaCaconstraintModel.Isca).To(Equal(core.BoolPtr(true)))
				Expect(configCaSigningProfilesCaCaconstraintModel.Maxpathlen).To(Equal(core.Float64Ptr(float64(0))))
				Expect(configCaSigningProfilesCaCaconstraintModel.Maxpathlenzero).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the ConfigCASigningProfilesCa model
				configCaSigningProfilesCaModel := new(blockchainv3.ConfigCASigningProfilesCa)
				Expect(configCaSigningProfilesCaModel).ToNot(BeNil())
				configCaSigningProfilesCaModel.Usage = []string{"cert sign"}
				configCaSigningProfilesCaModel.Expiry = core.StringPtr("43800h")
				configCaSigningProfilesCaModel.Caconstraint = configCaSigningProfilesCaCaconstraintModel
				Expect(configCaSigningProfilesCaModel.Usage).To(Equal([]string{"cert sign"}))
				Expect(configCaSigningProfilesCaModel.Expiry).To(Equal(core.StringPtr("43800h")))
				Expect(configCaSigningProfilesCaModel.Caconstraint).To(Equal(configCaSigningProfilesCaCaconstraintModel))

				// Construct an instance of the ConfigCASigningProfilesTls model
				configCaSigningProfilesTlsModel := new(blockchainv3.ConfigCASigningProfilesTls)
				Expect(configCaSigningProfilesTlsModel).ToNot(BeNil())
				configCaSigningProfilesTlsModel.Usage = []string{"cert sign"}
				configCaSigningProfilesTlsModel.Expiry = core.StringPtr("43800h")
				Expect(configCaSigningProfilesTlsModel.Usage).To(Equal([]string{"cert sign"}))
				Expect(configCaSigningProfilesTlsModel.Expiry).To(Equal(core.StringPtr("43800h")))

				// Construct an instance of the ConfigCASigningProfiles model
				configCaSigningProfilesModel := new(blockchainv3.ConfigCASigningProfiles)
				Expect(configCaSigningProfilesModel).ToNot(BeNil())
				configCaSigningProfilesModel.Ca = configCaSigningProfilesCaModel
				configCaSigningProfilesModel.Tls = configCaSigningProfilesTlsModel
				Expect(configCaSigningProfilesModel.Ca).To(Equal(configCaSigningProfilesCaModel))
				Expect(configCaSigningProfilesModel.Tls).To(Equal(configCaSigningProfilesTlsModel))

				// Construct an instance of the ConfigCASigning model
				configCaSigningModel := new(blockchainv3.ConfigCASigning)
				Expect(configCaSigningModel).ToNot(BeNil())
				configCaSigningModel.Default = configCaSigningDefaultModel
				configCaSigningModel.Profiles = configCaSigningProfilesModel
				Expect(configCaSigningModel.Default).To(Equal(configCaSigningDefaultModel))
				Expect(configCaSigningModel.Profiles).To(Equal(configCaSigningProfilesModel))

				// Construct an instance of the ConfigCACreate model
				configCaCreateModel := new(blockchainv3.ConfigCACreate)
				Expect(configCaCreateModel).ToNot(BeNil())
				configCaCreateModel.Cors = configCaCorsModel
				configCaCreateModel.Debug = core.BoolPtr(false)
				configCaCreateModel.Crlsizelimit = core.Float64Ptr(float64(512000))
				configCaCreateModel.Tls = configCaTlsModel
				configCaCreateModel.Ca = configCaCaModel
				configCaCreateModel.Crl = configCaCrlModel
				configCaCreateModel.Registry = configCaRegistryModel
				configCaCreateModel.Db = configCaDbModel
				configCaCreateModel.Affiliations = configCaAffiliationsModel
				configCaCreateModel.Csr = configCaCsrModel
				configCaCreateModel.Idemix = configCaIdemixModel
				configCaCreateModel.BCCSP = bccspModel
				configCaCreateModel.Intermediate = configCaIntermediateModel
				configCaCreateModel.Cfg = configCaCfgModel
				configCaCreateModel.Metrics = metricsModel
				configCaCreateModel.Signing = configCaSigningModel
				Expect(configCaCreateModel.Cors).To(Equal(configCaCorsModel))
				Expect(configCaCreateModel.Debug).To(Equal(core.BoolPtr(false)))
				Expect(configCaCreateModel.Crlsizelimit).To(Equal(core.Float64Ptr(float64(512000))))
				Expect(configCaCreateModel.Tls).To(Equal(configCaTlsModel))
				Expect(configCaCreateModel.Ca).To(Equal(configCaCaModel))
				Expect(configCaCreateModel.Crl).To(Equal(configCaCrlModel))
				Expect(configCaCreateModel.Registry).To(Equal(configCaRegistryModel))
				Expect(configCaCreateModel.Db).To(Equal(configCaDbModel))
				Expect(configCaCreateModel.Affiliations).To(Equal(configCaAffiliationsModel))
				Expect(configCaCreateModel.Csr).To(Equal(configCaCsrModel))
				Expect(configCaCreateModel.Idemix).To(Equal(configCaIdemixModel))
				Expect(configCaCreateModel.BCCSP).To(Equal(bccspModel))
				Expect(configCaCreateModel.Intermediate).To(Equal(configCaIntermediateModel))
				Expect(configCaCreateModel.Cfg).To(Equal(configCaCfgModel))
				Expect(configCaCreateModel.Metrics).To(Equal(metricsModel))
				Expect(configCaCreateModel.Signing).To(Equal(configCaSigningModel))

				// Construct an instance of the CreateCaBodyConfigOverride model
				createCaBodyConfigOverrideModel := new(blockchainv3.CreateCaBodyConfigOverride)
				Expect(createCaBodyConfigOverrideModel).ToNot(BeNil())
				createCaBodyConfigOverrideModel.Ca = configCaCreateModel
				createCaBodyConfigOverrideModel.Tlsca = configCaCreateModel
				Expect(createCaBodyConfigOverrideModel.Ca).To(Equal(configCaCreateModel))
				Expect(createCaBodyConfigOverrideModel.Tlsca).To(Equal(configCaCreateModel))

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv3.ResourceRequests)
				Expect(resourceRequestsModel).ToNot(BeNil())
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")
				Expect(resourceRequestsModel.Cpu).To(Equal(core.StringPtr("100m")))
				Expect(resourceRequestsModel.Memory).To(Equal(core.StringPtr("256MiB")))

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv3.ResourceLimits)
				Expect(resourceLimitsModel).ToNot(BeNil())
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")
				Expect(resourceLimitsModel.Cpu).To(Equal(core.StringPtr("100m")))
				Expect(resourceLimitsModel.Memory).To(Equal(core.StringPtr("256MiB")))

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv3.ResourceObject)
				Expect(resourceObjectModel).ToNot(BeNil())
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel
				Expect(resourceObjectModel.Requests).To(Equal(resourceRequestsModel))
				Expect(resourceObjectModel.Limits).To(Equal(resourceLimitsModel))

				// Construct an instance of the CreateCaBodyResources model
				createCaBodyResourcesModel := new(blockchainv3.CreateCaBodyResources)
				Expect(createCaBodyResourcesModel).ToNot(BeNil())
				createCaBodyResourcesModel.Ca = resourceObjectModel
				Expect(createCaBodyResourcesModel.Ca).To(Equal(resourceObjectModel))

				// Construct an instance of the StorageObject model
				storageObjectModel := new(blockchainv3.StorageObject)
				Expect(storageObjectModel).ToNot(BeNil())
				storageObjectModel.Size = core.StringPtr("4GiB")
				storageObjectModel.Class = core.StringPtr("default")
				Expect(storageObjectModel.Size).To(Equal(core.StringPtr("4GiB")))
				Expect(storageObjectModel.Class).To(Equal(core.StringPtr("default")))

				// Construct an instance of the CreateCaBodyStorage model
				createCaBodyStorageModel := new(blockchainv3.CreateCaBodyStorage)
				Expect(createCaBodyStorageModel).ToNot(BeNil())
				createCaBodyStorageModel.Ca = storageObjectModel
				Expect(createCaBodyStorageModel.Ca).To(Equal(storageObjectModel))

				// Construct an instance of the Hsm model
				hsmModel := new(blockchainv3.Hsm)
				Expect(hsmModel).ToNot(BeNil())
				hsmModel.Pkcs11endpoint = core.StringPtr("tcp://example.com:666")
				Expect(hsmModel.Pkcs11endpoint).To(Equal(core.StringPtr("tcp://example.com:666")))

				// Construct an instance of the CreateCaOptions model
				createCaOptionsDisplayName := "My CA"
				var createCaOptionsConfigOverride *blockchainv3.CreateCaBodyConfigOverride = nil
				createCaOptionsModel := blockchainService.NewCreateCaOptions(createCaOptionsDisplayName, createCaOptionsConfigOverride)
				createCaOptionsModel.SetDisplayName("My CA")
				createCaOptionsModel.SetConfigOverride(createCaBodyConfigOverrideModel)
				createCaOptionsModel.SetID("component1")
				createCaOptionsModel.SetResources(createCaBodyResourcesModel)
				createCaOptionsModel.SetStorage(createCaBodyStorageModel)
				createCaOptionsModel.SetZone("-")
				createCaOptionsModel.SetReplicas(float64(1))
				createCaOptionsModel.SetTags([]string{"fabric-ca"})
				createCaOptionsModel.SetHsm(hsmModel)
				createCaOptionsModel.SetRegion("-")
				createCaOptionsModel.SetVersion("1.4.6-1")
				createCaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCaOptionsModel).ToNot(BeNil())
				Expect(createCaOptionsModel.DisplayName).To(Equal(core.StringPtr("My CA")))
				Expect(createCaOptionsModel.ConfigOverride).To(Equal(createCaBodyConfigOverrideModel))
				Expect(createCaOptionsModel.ID).To(Equal(core.StringPtr("component1")))
				Expect(createCaOptionsModel.Resources).To(Equal(createCaBodyResourcesModel))
				Expect(createCaOptionsModel.Storage).To(Equal(createCaBodyStorageModel))
				Expect(createCaOptionsModel.Zone).To(Equal(core.StringPtr("-")))
				Expect(createCaOptionsModel.Replicas).To(Equal(core.Float64Ptr(float64(1))))
				Expect(createCaOptionsModel.Tags).To(Equal([]string{"fabric-ca"}))
				Expect(createCaOptionsModel.Hsm).To(Equal(hsmModel))
				Expect(createCaOptionsModel.Region).To(Equal(core.StringPtr("-")))
				Expect(createCaOptionsModel.Version).To(Equal(core.StringPtr("1.4.6-1")))
				Expect(createCaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateOrdererOptions successfully`, func() {
				// Construct an instance of the CryptoEnrollmentComponent model
				cryptoEnrollmentComponentModel := new(blockchainv3.CryptoEnrollmentComponent)
				Expect(cryptoEnrollmentComponentModel).ToNot(BeNil())
				cryptoEnrollmentComponentModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				Expect(cryptoEnrollmentComponentModel.Admincerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))

				// Construct an instance of the CryptoObjectEnrollmentCa model
				cryptoObjectEnrollmentCaModel := new(blockchainv3.CryptoObjectEnrollmentCa)
				Expect(cryptoObjectEnrollmentCaModel).ToNot(BeNil())
				cryptoObjectEnrollmentCaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				cryptoObjectEnrollmentCaModel.Port = core.Float64Ptr(float64(7054))
				cryptoObjectEnrollmentCaModel.Name = core.StringPtr("ca")
				cryptoObjectEnrollmentCaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				cryptoObjectEnrollmentCaModel.EnrollID = core.StringPtr("admin")
				cryptoObjectEnrollmentCaModel.EnrollSecret = core.StringPtr("password")
				Expect(cryptoObjectEnrollmentCaModel.Host).To(Equal(core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")))
				Expect(cryptoObjectEnrollmentCaModel.Port).To(Equal(core.Float64Ptr(float64(7054))))
				Expect(cryptoObjectEnrollmentCaModel.Name).To(Equal(core.StringPtr("ca")))
				Expect(cryptoObjectEnrollmentCaModel.TlsCert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(cryptoObjectEnrollmentCaModel.EnrollID).To(Equal(core.StringPtr("admin")))
				Expect(cryptoObjectEnrollmentCaModel.EnrollSecret).To(Equal(core.StringPtr("password")))

				// Construct an instance of the CryptoObjectEnrollmentTlsca model
				cryptoObjectEnrollmentTlscaModel := new(blockchainv3.CryptoObjectEnrollmentTlsca)
				Expect(cryptoObjectEnrollmentTlscaModel).ToNot(BeNil())
				cryptoObjectEnrollmentTlscaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				cryptoObjectEnrollmentTlscaModel.Port = core.Float64Ptr(float64(7054))
				cryptoObjectEnrollmentTlscaModel.Name = core.StringPtr("tlsca")
				cryptoObjectEnrollmentTlscaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				cryptoObjectEnrollmentTlscaModel.EnrollID = core.StringPtr("admin")
				cryptoObjectEnrollmentTlscaModel.EnrollSecret = core.StringPtr("password")
				cryptoObjectEnrollmentTlscaModel.CsrHosts = []string{"testString"}
				Expect(cryptoObjectEnrollmentTlscaModel.Host).To(Equal(core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")))
				Expect(cryptoObjectEnrollmentTlscaModel.Port).To(Equal(core.Float64Ptr(float64(7054))))
				Expect(cryptoObjectEnrollmentTlscaModel.Name).To(Equal(core.StringPtr("tlsca")))
				Expect(cryptoObjectEnrollmentTlscaModel.TlsCert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(cryptoObjectEnrollmentTlscaModel.EnrollID).To(Equal(core.StringPtr("admin")))
				Expect(cryptoObjectEnrollmentTlscaModel.EnrollSecret).To(Equal(core.StringPtr("password")))
				Expect(cryptoObjectEnrollmentTlscaModel.CsrHosts).To(Equal([]string{"testString"}))

				// Construct an instance of the CryptoObjectEnrollment model
				cryptoObjectEnrollmentModel := new(blockchainv3.CryptoObjectEnrollment)
				Expect(cryptoObjectEnrollmentModel).ToNot(BeNil())
				cryptoObjectEnrollmentModel.Component = cryptoEnrollmentComponentModel
				cryptoObjectEnrollmentModel.Ca = cryptoObjectEnrollmentCaModel
				cryptoObjectEnrollmentModel.Tlsca = cryptoObjectEnrollmentTlscaModel
				Expect(cryptoObjectEnrollmentModel.Component).To(Equal(cryptoEnrollmentComponentModel))
				Expect(cryptoObjectEnrollmentModel.Ca).To(Equal(cryptoObjectEnrollmentCaModel))
				Expect(cryptoObjectEnrollmentModel.Tlsca).To(Equal(cryptoObjectEnrollmentTlscaModel))

				// Construct an instance of the ClientAuth model
				clientAuthModel := new(blockchainv3.ClientAuth)
				Expect(clientAuthModel).ToNot(BeNil())
				clientAuthModel.Type = core.StringPtr("noclientcert")
				clientAuthModel.TlsCerts = []string{"testString"}
				Expect(clientAuthModel.Type).To(Equal(core.StringPtr("noclientcert")))
				Expect(clientAuthModel.TlsCerts).To(Equal([]string{"testString"}))

				// Construct an instance of the MspCryptoComp model
				mspCryptoCompModel := new(blockchainv3.MspCryptoComp)
				Expect(mspCryptoCompModel).ToNot(BeNil())
				mspCryptoCompModel.Ekey = core.StringPtr("testString")
				mspCryptoCompModel.Ecert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoCompModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				mspCryptoCompModel.TlsKey = core.StringPtr("testString")
				mspCryptoCompModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoCompModel.ClientAuth = clientAuthModel
				Expect(mspCryptoCompModel.Ekey).To(Equal(core.StringPtr("testString")))
				Expect(mspCryptoCompModel.Ecert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(mspCryptoCompModel.AdminCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(mspCryptoCompModel.TlsKey).To(Equal(core.StringPtr("testString")))
				Expect(mspCryptoCompModel.TlsCert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(mspCryptoCompModel.ClientAuth).To(Equal(clientAuthModel))

				// Construct an instance of the MspCryptoCa model
				mspCryptoCaModel := new(blockchainv3.MspCryptoCa)
				Expect(mspCryptoCaModel).ToNot(BeNil())
				mspCryptoCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				mspCryptoCaModel.CaIntermediateCerts = []string{"testString"}
				Expect(mspCryptoCaModel.RootCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(mspCryptoCaModel.CaIntermediateCerts).To(Equal([]string{"testString"}))

				// Construct an instance of the CryptoObjectMsp model
				cryptoObjectMspModel := new(blockchainv3.CryptoObjectMsp)
				Expect(cryptoObjectMspModel).ToNot(BeNil())
				cryptoObjectMspModel.Component = mspCryptoCompModel
				cryptoObjectMspModel.Ca = mspCryptoCaModel
				cryptoObjectMspModel.Tlsca = mspCryptoCaModel
				Expect(cryptoObjectMspModel.Component).To(Equal(mspCryptoCompModel))
				Expect(cryptoObjectMspModel.Ca).To(Equal(mspCryptoCaModel))
				Expect(cryptoObjectMspModel.Tlsca).To(Equal(mspCryptoCaModel))

				// Construct an instance of the CryptoObject model
				cryptoObjectModel := new(blockchainv3.CryptoObject)
				Expect(cryptoObjectModel).ToNot(BeNil())
				cryptoObjectModel.Enrollment = cryptoObjectEnrollmentModel
				cryptoObjectModel.Msp = cryptoObjectMspModel
				Expect(cryptoObjectModel.Enrollment).To(Equal(cryptoObjectEnrollmentModel))
				Expect(cryptoObjectModel.Msp).To(Equal(cryptoObjectMspModel))

				// Construct an instance of the ConfigOrdererKeepalive model
				configOrdererKeepaliveModel := new(blockchainv3.ConfigOrdererKeepalive)
				Expect(configOrdererKeepaliveModel).ToNot(BeNil())
				configOrdererKeepaliveModel.ServerMinInterval = core.StringPtr("60s")
				configOrdererKeepaliveModel.ServerInterval = core.StringPtr("2h")
				configOrdererKeepaliveModel.ServerTimeout = core.StringPtr("20s")
				Expect(configOrdererKeepaliveModel.ServerMinInterval).To(Equal(core.StringPtr("60s")))
				Expect(configOrdererKeepaliveModel.ServerInterval).To(Equal(core.StringPtr("2h")))
				Expect(configOrdererKeepaliveModel.ServerTimeout).To(Equal(core.StringPtr("20s")))

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv3.BccspSW)
				Expect(bccspSwModel).ToNot(BeNil())
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))
				Expect(bccspSwModel.Hash).To(Equal(core.StringPtr("SHA2")))
				Expect(bccspSwModel.Security).To(Equal(core.Float64Ptr(float64(256))))

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv3.BccspPKCS11)
				Expect(bccspPkcS11Model).ToNot(BeNil())
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))
				Expect(bccspPkcS11Model.Label).To(Equal(core.StringPtr("testString")))
				Expect(bccspPkcS11Model.Pin).To(Equal(core.StringPtr("testString")))
				Expect(bccspPkcS11Model.Hash).To(Equal(core.StringPtr("SHA2")))
				Expect(bccspPkcS11Model.Security).To(Equal(core.Float64Ptr(float64(256))))

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv3.Bccsp)
				Expect(bccspModel).ToNot(BeNil())
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model
				Expect(bccspModel.Default).To(Equal(core.StringPtr("SW")))
				Expect(bccspModel.SW).To(Equal(bccspSwModel))
				Expect(bccspModel.PKCS11).To(Equal(bccspPkcS11Model))

				// Construct an instance of the ConfigOrdererAuthentication model
				configOrdererAuthenticationModel := new(blockchainv3.ConfigOrdererAuthentication)
				Expect(configOrdererAuthenticationModel).ToNot(BeNil())
				configOrdererAuthenticationModel.TimeWindow = core.StringPtr("15m")
				configOrdererAuthenticationModel.NoExpirationChecks = core.BoolPtr(false)
				Expect(configOrdererAuthenticationModel.TimeWindow).To(Equal(core.StringPtr("15m")))
				Expect(configOrdererAuthenticationModel.NoExpirationChecks).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the ConfigOrdererGeneral model
				configOrdererGeneralModel := new(blockchainv3.ConfigOrdererGeneral)
				Expect(configOrdererGeneralModel).ToNot(BeNil())
				configOrdererGeneralModel.Keepalive = configOrdererKeepaliveModel
				configOrdererGeneralModel.BCCSP = bccspModel
				configOrdererGeneralModel.Authentication = configOrdererAuthenticationModel
				Expect(configOrdererGeneralModel.Keepalive).To(Equal(configOrdererKeepaliveModel))
				Expect(configOrdererGeneralModel.BCCSP).To(Equal(bccspModel))
				Expect(configOrdererGeneralModel.Authentication).To(Equal(configOrdererAuthenticationModel))

				// Construct an instance of the ConfigOrdererDebug model
				configOrdererDebugModel := new(blockchainv3.ConfigOrdererDebug)
				Expect(configOrdererDebugModel).ToNot(BeNil())
				configOrdererDebugModel.BroadcastTraceDir = core.StringPtr("testString")
				configOrdererDebugModel.DeliverTraceDir = core.StringPtr("testString")
				Expect(configOrdererDebugModel.BroadcastTraceDir).To(Equal(core.StringPtr("testString")))
				Expect(configOrdererDebugModel.DeliverTraceDir).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ConfigOrdererMetricsStatsd model
				configOrdererMetricsStatsdModel := new(blockchainv3.ConfigOrdererMetricsStatsd)
				Expect(configOrdererMetricsStatsdModel).ToNot(BeNil())
				configOrdererMetricsStatsdModel.Network = core.StringPtr("udp")
				configOrdererMetricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				configOrdererMetricsStatsdModel.WriteInterval = core.StringPtr("10s")
				configOrdererMetricsStatsdModel.Prefix = core.StringPtr("server")
				Expect(configOrdererMetricsStatsdModel.Network).To(Equal(core.StringPtr("udp")))
				Expect(configOrdererMetricsStatsdModel.Address).To(Equal(core.StringPtr("127.0.0.1:8125")))
				Expect(configOrdererMetricsStatsdModel.WriteInterval).To(Equal(core.StringPtr("10s")))
				Expect(configOrdererMetricsStatsdModel.Prefix).To(Equal(core.StringPtr("server")))

				// Construct an instance of the ConfigOrdererMetrics model
				configOrdererMetricsModel := new(blockchainv3.ConfigOrdererMetrics)
				Expect(configOrdererMetricsModel).ToNot(BeNil())
				configOrdererMetricsModel.Provider = core.StringPtr("disabled")
				configOrdererMetricsModel.Statsd = configOrdererMetricsStatsdModel
				Expect(configOrdererMetricsModel.Provider).To(Equal(core.StringPtr("disabled")))
				Expect(configOrdererMetricsModel.Statsd).To(Equal(configOrdererMetricsStatsdModel))

				// Construct an instance of the ConfigOrdererCreate model
				configOrdererCreateModel := new(blockchainv3.ConfigOrdererCreate)
				Expect(configOrdererCreateModel).ToNot(BeNil())
				configOrdererCreateModel.General = configOrdererGeneralModel
				configOrdererCreateModel.Debug = configOrdererDebugModel
				configOrdererCreateModel.Metrics = configOrdererMetricsModel
				Expect(configOrdererCreateModel.General).To(Equal(configOrdererGeneralModel))
				Expect(configOrdererCreateModel.Debug).To(Equal(configOrdererDebugModel))
				Expect(configOrdererCreateModel.Metrics).To(Equal(configOrdererMetricsModel))

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv3.ResourceRequests)
				Expect(resourceRequestsModel).ToNot(BeNil())
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")
				Expect(resourceRequestsModel.Cpu).To(Equal(core.StringPtr("100m")))
				Expect(resourceRequestsModel.Memory).To(Equal(core.StringPtr("256MiB")))

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv3.ResourceLimits)
				Expect(resourceLimitsModel).ToNot(BeNil())
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")
				Expect(resourceLimitsModel.Cpu).To(Equal(core.StringPtr("100m")))
				Expect(resourceLimitsModel.Memory).To(Equal(core.StringPtr("256MiB")))

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv3.ResourceObject)
				Expect(resourceObjectModel).ToNot(BeNil())
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel
				Expect(resourceObjectModel.Requests).To(Equal(resourceRequestsModel))
				Expect(resourceObjectModel.Limits).To(Equal(resourceLimitsModel))

				// Construct an instance of the CreateOrdererRaftBodyResources model
				createOrdererRaftBodyResourcesModel := new(blockchainv3.CreateOrdererRaftBodyResources)
				Expect(createOrdererRaftBodyResourcesModel).ToNot(BeNil())
				createOrdererRaftBodyResourcesModel.Orderer = resourceObjectModel
				createOrdererRaftBodyResourcesModel.Proxy = resourceObjectModel
				Expect(createOrdererRaftBodyResourcesModel.Orderer).To(Equal(resourceObjectModel))
				Expect(createOrdererRaftBodyResourcesModel.Proxy).To(Equal(resourceObjectModel))

				// Construct an instance of the StorageObject model
				storageObjectModel := new(blockchainv3.StorageObject)
				Expect(storageObjectModel).ToNot(BeNil())
				storageObjectModel.Size = core.StringPtr("4GiB")
				storageObjectModel.Class = core.StringPtr("default")
				Expect(storageObjectModel.Size).To(Equal(core.StringPtr("4GiB")))
				Expect(storageObjectModel.Class).To(Equal(core.StringPtr("default")))

				// Construct an instance of the CreateOrdererRaftBodyStorage model
				createOrdererRaftBodyStorageModel := new(blockchainv3.CreateOrdererRaftBodyStorage)
				Expect(createOrdererRaftBodyStorageModel).ToNot(BeNil())
				createOrdererRaftBodyStorageModel.Orderer = storageObjectModel
				Expect(createOrdererRaftBodyStorageModel.Orderer).To(Equal(storageObjectModel))

				// Construct an instance of the Hsm model
				hsmModel := new(blockchainv3.Hsm)
				Expect(hsmModel).ToNot(BeNil())
				hsmModel.Pkcs11endpoint = core.StringPtr("tcp://example.com:666")
				Expect(hsmModel.Pkcs11endpoint).To(Equal(core.StringPtr("tcp://example.com:666")))

				// Construct an instance of the CreateOrdererOptions model
				createOrdererOptionsOrdererType := "raft"
				createOrdererOptionsMspID := "Org1"
				createOrdererOptionsDisplayName := "orderer"
				createOrdererOptionsCrypto := []blockchainv3.CryptoObject{}
				createOrdererOptionsModel := blockchainService.NewCreateOrdererOptions(createOrdererOptionsOrdererType, createOrdererOptionsMspID, createOrdererOptionsDisplayName, createOrdererOptionsCrypto)
				createOrdererOptionsModel.SetOrdererType("raft")
				createOrdererOptionsModel.SetMspID("Org1")
				createOrdererOptionsModel.SetDisplayName("orderer")
				createOrdererOptionsModel.SetCrypto([]blockchainv3.CryptoObject{*cryptoObjectModel})
				createOrdererOptionsModel.SetClusterName("ordering service 1")
				createOrdererOptionsModel.SetID("component1")
				createOrdererOptionsModel.SetClusterID("abcde")
				createOrdererOptionsModel.SetExternalAppend(false)
				createOrdererOptionsModel.SetConfigOverride([]blockchainv3.ConfigOrdererCreate{*configOrdererCreateModel})
				createOrdererOptionsModel.SetResources(createOrdererRaftBodyResourcesModel)
				createOrdererOptionsModel.SetStorage(createOrdererRaftBodyStorageModel)
				createOrdererOptionsModel.SetSystemChannelID("testchainid")
				createOrdererOptionsModel.SetZone([]string{"-"})
				createOrdererOptionsModel.SetTags([]string{"fabric-ca"})
				createOrdererOptionsModel.SetRegion([]string{"-"})
				createOrdererOptionsModel.SetHsm(hsmModel)
				createOrdererOptionsModel.SetVersion("1.4.6-1")
				createOrdererOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createOrdererOptionsModel).ToNot(BeNil())
				Expect(createOrdererOptionsModel.OrdererType).To(Equal(core.StringPtr("raft")))
				Expect(createOrdererOptionsModel.MspID).To(Equal(core.StringPtr("Org1")))
				Expect(createOrdererOptionsModel.DisplayName).To(Equal(core.StringPtr("orderer")))
				Expect(createOrdererOptionsModel.Crypto).To(Equal([]blockchainv3.CryptoObject{*cryptoObjectModel}))
				Expect(createOrdererOptionsModel.ClusterName).To(Equal(core.StringPtr("ordering service 1")))
				Expect(createOrdererOptionsModel.ID).To(Equal(core.StringPtr("component1")))
				Expect(createOrdererOptionsModel.ClusterID).To(Equal(core.StringPtr("abcde")))
				Expect(createOrdererOptionsModel.ExternalAppend).To(Equal(core.BoolPtr(false)))
				Expect(createOrdererOptionsModel.ConfigOverride).To(Equal([]blockchainv3.ConfigOrdererCreate{*configOrdererCreateModel}))
				Expect(createOrdererOptionsModel.Resources).To(Equal(createOrdererRaftBodyResourcesModel))
				Expect(createOrdererOptionsModel.Storage).To(Equal(createOrdererRaftBodyStorageModel))
				Expect(createOrdererOptionsModel.SystemChannelID).To(Equal(core.StringPtr("testchainid")))
				Expect(createOrdererOptionsModel.Zone).To(Equal([]string{"-"}))
				Expect(createOrdererOptionsModel.Tags).To(Equal([]string{"fabric-ca"}))
				Expect(createOrdererOptionsModel.Region).To(Equal([]string{"-"}))
				Expect(createOrdererOptionsModel.Hsm).To(Equal(hsmModel))
				Expect(createOrdererOptionsModel.Version).To(Equal(core.StringPtr("1.4.6-1")))
				Expect(createOrdererOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateOrdererRaftBodyResources successfully`, func() {
				var orderer *blockchainv3.ResourceObject = nil
				_, err := blockchainService.NewCreateOrdererRaftBodyResources(orderer)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCreateOrdererRaftBodyStorage successfully`, func() {
				var orderer *blockchainv3.StorageObject = nil
				_, err := blockchainService.NewCreateOrdererRaftBodyStorage(orderer)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCreatePeerBodyStorage successfully`, func() {
				var peer *blockchainv3.StorageObject = nil
				_, err := blockchainService.NewCreatePeerBodyStorage(peer)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCreatePeerOptions successfully`, func() {
				// Construct an instance of the CryptoEnrollmentComponent model
				cryptoEnrollmentComponentModel := new(blockchainv3.CryptoEnrollmentComponent)
				Expect(cryptoEnrollmentComponentModel).ToNot(BeNil())
				cryptoEnrollmentComponentModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				Expect(cryptoEnrollmentComponentModel.Admincerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))

				// Construct an instance of the CryptoObjectEnrollmentCa model
				cryptoObjectEnrollmentCaModel := new(blockchainv3.CryptoObjectEnrollmentCa)
				Expect(cryptoObjectEnrollmentCaModel).ToNot(BeNil())
				cryptoObjectEnrollmentCaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				cryptoObjectEnrollmentCaModel.Port = core.Float64Ptr(float64(7054))
				cryptoObjectEnrollmentCaModel.Name = core.StringPtr("ca")
				cryptoObjectEnrollmentCaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				cryptoObjectEnrollmentCaModel.EnrollID = core.StringPtr("admin")
				cryptoObjectEnrollmentCaModel.EnrollSecret = core.StringPtr("password")
				Expect(cryptoObjectEnrollmentCaModel.Host).To(Equal(core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")))
				Expect(cryptoObjectEnrollmentCaModel.Port).To(Equal(core.Float64Ptr(float64(7054))))
				Expect(cryptoObjectEnrollmentCaModel.Name).To(Equal(core.StringPtr("ca")))
				Expect(cryptoObjectEnrollmentCaModel.TlsCert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(cryptoObjectEnrollmentCaModel.EnrollID).To(Equal(core.StringPtr("admin")))
				Expect(cryptoObjectEnrollmentCaModel.EnrollSecret).To(Equal(core.StringPtr("password")))

				// Construct an instance of the CryptoObjectEnrollmentTlsca model
				cryptoObjectEnrollmentTlscaModel := new(blockchainv3.CryptoObjectEnrollmentTlsca)
				Expect(cryptoObjectEnrollmentTlscaModel).ToNot(BeNil())
				cryptoObjectEnrollmentTlscaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				cryptoObjectEnrollmentTlscaModel.Port = core.Float64Ptr(float64(7054))
				cryptoObjectEnrollmentTlscaModel.Name = core.StringPtr("tlsca")
				cryptoObjectEnrollmentTlscaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				cryptoObjectEnrollmentTlscaModel.EnrollID = core.StringPtr("admin")
				cryptoObjectEnrollmentTlscaModel.EnrollSecret = core.StringPtr("password")
				cryptoObjectEnrollmentTlscaModel.CsrHosts = []string{"testString"}
				Expect(cryptoObjectEnrollmentTlscaModel.Host).To(Equal(core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")))
				Expect(cryptoObjectEnrollmentTlscaModel.Port).To(Equal(core.Float64Ptr(float64(7054))))
				Expect(cryptoObjectEnrollmentTlscaModel.Name).To(Equal(core.StringPtr("tlsca")))
				Expect(cryptoObjectEnrollmentTlscaModel.TlsCert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(cryptoObjectEnrollmentTlscaModel.EnrollID).To(Equal(core.StringPtr("admin")))
				Expect(cryptoObjectEnrollmentTlscaModel.EnrollSecret).To(Equal(core.StringPtr("password")))
				Expect(cryptoObjectEnrollmentTlscaModel.CsrHosts).To(Equal([]string{"testString"}))

				// Construct an instance of the CryptoObjectEnrollment model
				cryptoObjectEnrollmentModel := new(blockchainv3.CryptoObjectEnrollment)
				Expect(cryptoObjectEnrollmentModel).ToNot(BeNil())
				cryptoObjectEnrollmentModel.Component = cryptoEnrollmentComponentModel
				cryptoObjectEnrollmentModel.Ca = cryptoObjectEnrollmentCaModel
				cryptoObjectEnrollmentModel.Tlsca = cryptoObjectEnrollmentTlscaModel
				Expect(cryptoObjectEnrollmentModel.Component).To(Equal(cryptoEnrollmentComponentModel))
				Expect(cryptoObjectEnrollmentModel.Ca).To(Equal(cryptoObjectEnrollmentCaModel))
				Expect(cryptoObjectEnrollmentModel.Tlsca).To(Equal(cryptoObjectEnrollmentTlscaModel))

				// Construct an instance of the ClientAuth model
				clientAuthModel := new(blockchainv3.ClientAuth)
				Expect(clientAuthModel).ToNot(BeNil())
				clientAuthModel.Type = core.StringPtr("noclientcert")
				clientAuthModel.TlsCerts = []string{"testString"}
				Expect(clientAuthModel.Type).To(Equal(core.StringPtr("noclientcert")))
				Expect(clientAuthModel.TlsCerts).To(Equal([]string{"testString"}))

				// Construct an instance of the MspCryptoComp model
				mspCryptoCompModel := new(blockchainv3.MspCryptoComp)
				Expect(mspCryptoCompModel).ToNot(BeNil())
				mspCryptoCompModel.Ekey = core.StringPtr("testString")
				mspCryptoCompModel.Ecert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoCompModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				mspCryptoCompModel.TlsKey = core.StringPtr("testString")
				mspCryptoCompModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoCompModel.ClientAuth = clientAuthModel
				Expect(mspCryptoCompModel.Ekey).To(Equal(core.StringPtr("testString")))
				Expect(mspCryptoCompModel.Ecert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(mspCryptoCompModel.AdminCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(mspCryptoCompModel.TlsKey).To(Equal(core.StringPtr("testString")))
				Expect(mspCryptoCompModel.TlsCert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(mspCryptoCompModel.ClientAuth).To(Equal(clientAuthModel))

				// Construct an instance of the MspCryptoCa model
				mspCryptoCaModel := new(blockchainv3.MspCryptoCa)
				Expect(mspCryptoCaModel).ToNot(BeNil())
				mspCryptoCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				mspCryptoCaModel.CaIntermediateCerts = []string{"testString"}
				Expect(mspCryptoCaModel.RootCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(mspCryptoCaModel.CaIntermediateCerts).To(Equal([]string{"testString"}))

				// Construct an instance of the CryptoObjectMsp model
				cryptoObjectMspModel := new(blockchainv3.CryptoObjectMsp)
				Expect(cryptoObjectMspModel).ToNot(BeNil())
				cryptoObjectMspModel.Component = mspCryptoCompModel
				cryptoObjectMspModel.Ca = mspCryptoCaModel
				cryptoObjectMspModel.Tlsca = mspCryptoCaModel
				Expect(cryptoObjectMspModel.Component).To(Equal(mspCryptoCompModel))
				Expect(cryptoObjectMspModel.Ca).To(Equal(mspCryptoCaModel))
				Expect(cryptoObjectMspModel.Tlsca).To(Equal(mspCryptoCaModel))

				// Construct an instance of the CryptoObject model
				cryptoObjectModel := new(blockchainv3.CryptoObject)
				Expect(cryptoObjectModel).ToNot(BeNil())
				cryptoObjectModel.Enrollment = cryptoObjectEnrollmentModel
				cryptoObjectModel.Msp = cryptoObjectMspModel
				Expect(cryptoObjectModel.Enrollment).To(Equal(cryptoObjectEnrollmentModel))
				Expect(cryptoObjectModel.Msp).To(Equal(cryptoObjectMspModel))

				// Construct an instance of the ConfigPeerKeepaliveClient model
				configPeerKeepaliveClientModel := new(blockchainv3.ConfigPeerKeepaliveClient)
				Expect(configPeerKeepaliveClientModel).ToNot(BeNil())
				configPeerKeepaliveClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveClientModel.Timeout = core.StringPtr("20s")
				Expect(configPeerKeepaliveClientModel.Interval).To(Equal(core.StringPtr("60s")))
				Expect(configPeerKeepaliveClientModel.Timeout).To(Equal(core.StringPtr("20s")))

				// Construct an instance of the ConfigPeerKeepaliveDeliveryClient model
				configPeerKeepaliveDeliveryClientModel := new(blockchainv3.ConfigPeerKeepaliveDeliveryClient)
				Expect(configPeerKeepaliveDeliveryClientModel).ToNot(BeNil())
				configPeerKeepaliveDeliveryClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveDeliveryClientModel.Timeout = core.StringPtr("20s")
				Expect(configPeerKeepaliveDeliveryClientModel.Interval).To(Equal(core.StringPtr("60s")))
				Expect(configPeerKeepaliveDeliveryClientModel.Timeout).To(Equal(core.StringPtr("20s")))

				// Construct an instance of the ConfigPeerKeepalive model
				configPeerKeepaliveModel := new(blockchainv3.ConfigPeerKeepalive)
				Expect(configPeerKeepaliveModel).ToNot(BeNil())
				configPeerKeepaliveModel.MinInterval = core.StringPtr("60s")
				configPeerKeepaliveModel.Client = configPeerKeepaliveClientModel
				configPeerKeepaliveModel.DeliveryClient = configPeerKeepaliveDeliveryClientModel
				Expect(configPeerKeepaliveModel.MinInterval).To(Equal(core.StringPtr("60s")))
				Expect(configPeerKeepaliveModel.Client).To(Equal(configPeerKeepaliveClientModel))
				Expect(configPeerKeepaliveModel.DeliveryClient).To(Equal(configPeerKeepaliveDeliveryClientModel))

				// Construct an instance of the ConfigPeerGossipElection model
				configPeerGossipElectionModel := new(blockchainv3.ConfigPeerGossipElection)
				Expect(configPeerGossipElectionModel).ToNot(BeNil())
				configPeerGossipElectionModel.StartupGracePeriod = core.StringPtr("15s")
				configPeerGossipElectionModel.MembershipSampleInterval = core.StringPtr("1s")
				configPeerGossipElectionModel.LeaderAliveThreshold = core.StringPtr("10s")
				configPeerGossipElectionModel.LeaderElectionDuration = core.StringPtr("5s")
				Expect(configPeerGossipElectionModel.StartupGracePeriod).To(Equal(core.StringPtr("15s")))
				Expect(configPeerGossipElectionModel.MembershipSampleInterval).To(Equal(core.StringPtr("1s")))
				Expect(configPeerGossipElectionModel.LeaderAliveThreshold).To(Equal(core.StringPtr("10s")))
				Expect(configPeerGossipElectionModel.LeaderElectionDuration).To(Equal(core.StringPtr("5s")))

				// Construct an instance of the ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy model
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel := new(blockchainv3.ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy)
				Expect(configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel).ToNot(BeNil())
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.RequiredPeerCount = core.Float64Ptr(float64(0))
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.MaxPeerCount = core.Float64Ptr(float64(1))
				Expect(configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.RequiredPeerCount).To(Equal(core.Float64Ptr(float64(0))))
				Expect(configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.MaxPeerCount).To(Equal(core.Float64Ptr(float64(1))))

				// Construct an instance of the ConfigPeerGossipPvtData model
				configPeerGossipPvtDataModel := new(blockchainv3.ConfigPeerGossipPvtData)
				Expect(configPeerGossipPvtDataModel).ToNot(BeNil())
				configPeerGossipPvtDataModel.PullRetryThreshold = core.StringPtr("60s")
				configPeerGossipPvtDataModel.TransientstoreMaxBlockRetention = core.Float64Ptr(float64(1000))
				configPeerGossipPvtDataModel.PushAckTimeout = core.StringPtr("3s")
				configPeerGossipPvtDataModel.BtlPullMargin = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileBatchSize = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileSleepInterval = core.StringPtr("1m")
				configPeerGossipPvtDataModel.ReconciliationEnabled = core.BoolPtr(true)
				configPeerGossipPvtDataModel.SkipPullingInvalidTransactionsDuringCommit = core.BoolPtr(false)
				configPeerGossipPvtDataModel.ImplicitCollectionDisseminationPolicy = configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel
				Expect(configPeerGossipPvtDataModel.PullRetryThreshold).To(Equal(core.StringPtr("60s")))
				Expect(configPeerGossipPvtDataModel.TransientstoreMaxBlockRetention).To(Equal(core.Float64Ptr(float64(1000))))
				Expect(configPeerGossipPvtDataModel.PushAckTimeout).To(Equal(core.StringPtr("3s")))
				Expect(configPeerGossipPvtDataModel.BtlPullMargin).To(Equal(core.Float64Ptr(float64(10))))
				Expect(configPeerGossipPvtDataModel.ReconcileBatchSize).To(Equal(core.Float64Ptr(float64(10))))
				Expect(configPeerGossipPvtDataModel.ReconcileSleepInterval).To(Equal(core.StringPtr("1m")))
				Expect(configPeerGossipPvtDataModel.ReconciliationEnabled).To(Equal(core.BoolPtr(true)))
				Expect(configPeerGossipPvtDataModel.SkipPullingInvalidTransactionsDuringCommit).To(Equal(core.BoolPtr(false)))
				Expect(configPeerGossipPvtDataModel.ImplicitCollectionDisseminationPolicy).To(Equal(configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel))

				// Construct an instance of the ConfigPeerGossipState model
				configPeerGossipStateModel := new(blockchainv3.ConfigPeerGossipState)
				Expect(configPeerGossipStateModel).ToNot(BeNil())
				configPeerGossipStateModel.Enabled = core.BoolPtr(true)
				configPeerGossipStateModel.CheckInterval = core.StringPtr("10s")
				configPeerGossipStateModel.ResponseTimeout = core.StringPtr("3s")
				configPeerGossipStateModel.BatchSize = core.Float64Ptr(float64(10))
				configPeerGossipStateModel.BlockBufferSize = core.Float64Ptr(float64(100))
				configPeerGossipStateModel.MaxRetries = core.Float64Ptr(float64(3))
				Expect(configPeerGossipStateModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(configPeerGossipStateModel.CheckInterval).To(Equal(core.StringPtr("10s")))
				Expect(configPeerGossipStateModel.ResponseTimeout).To(Equal(core.StringPtr("3s")))
				Expect(configPeerGossipStateModel.BatchSize).To(Equal(core.Float64Ptr(float64(10))))
				Expect(configPeerGossipStateModel.BlockBufferSize).To(Equal(core.Float64Ptr(float64(100))))
				Expect(configPeerGossipStateModel.MaxRetries).To(Equal(core.Float64Ptr(float64(3))))

				// Construct an instance of the ConfigPeerGossip model
				configPeerGossipModel := new(blockchainv3.ConfigPeerGossip)
				Expect(configPeerGossipModel).ToNot(BeNil())
				configPeerGossipModel.UseLeaderElection = core.BoolPtr(true)
				configPeerGossipModel.OrgLeader = core.BoolPtr(false)
				configPeerGossipModel.MembershipTrackerInterval = core.StringPtr("5s")
				configPeerGossipModel.MaxBlockCountToStore = core.Float64Ptr(float64(100))
				configPeerGossipModel.MaxPropagationBurstLatency = core.StringPtr("10ms")
				configPeerGossipModel.MaxPropagationBurstSize = core.Float64Ptr(float64(10))
				configPeerGossipModel.PropagateIterations = core.Float64Ptr(float64(3))
				configPeerGossipModel.PullInterval = core.StringPtr("4s")
				configPeerGossipModel.PullPeerNum = core.Float64Ptr(float64(3))
				configPeerGossipModel.RequestStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.PublishStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.StateInfoRetentionInterval = core.StringPtr("0s")
				configPeerGossipModel.PublishCertPeriod = core.StringPtr("10s")
				configPeerGossipModel.SkipBlockVerification = core.BoolPtr(false)
				configPeerGossipModel.DialTimeout = core.StringPtr("3s")
				configPeerGossipModel.ConnTimeout = core.StringPtr("2s")
				configPeerGossipModel.RecvBuffSize = core.Float64Ptr(float64(20))
				configPeerGossipModel.SendBuffSize = core.Float64Ptr(float64(200))
				configPeerGossipModel.DigestWaitTime = core.StringPtr("1s")
				configPeerGossipModel.RequestWaitTime = core.StringPtr("1500ms")
				configPeerGossipModel.ResponseWaitTime = core.StringPtr("2s")
				configPeerGossipModel.AliveTimeInterval = core.StringPtr("5s")
				configPeerGossipModel.AliveExpirationTimeout = core.StringPtr("25s")
				configPeerGossipModel.ReconnectInterval = core.StringPtr("25s")
				configPeerGossipModel.Election = configPeerGossipElectionModel
				configPeerGossipModel.PvtData = configPeerGossipPvtDataModel
				configPeerGossipModel.State = configPeerGossipStateModel
				Expect(configPeerGossipModel.UseLeaderElection).To(Equal(core.BoolPtr(true)))
				Expect(configPeerGossipModel.OrgLeader).To(Equal(core.BoolPtr(false)))
				Expect(configPeerGossipModel.MembershipTrackerInterval).To(Equal(core.StringPtr("5s")))
				Expect(configPeerGossipModel.MaxBlockCountToStore).To(Equal(core.Float64Ptr(float64(100))))
				Expect(configPeerGossipModel.MaxPropagationBurstLatency).To(Equal(core.StringPtr("10ms")))
				Expect(configPeerGossipModel.MaxPropagationBurstSize).To(Equal(core.Float64Ptr(float64(10))))
				Expect(configPeerGossipModel.PropagateIterations).To(Equal(core.Float64Ptr(float64(3))))
				Expect(configPeerGossipModel.PullInterval).To(Equal(core.StringPtr("4s")))
				Expect(configPeerGossipModel.PullPeerNum).To(Equal(core.Float64Ptr(float64(3))))
				Expect(configPeerGossipModel.RequestStateInfoInterval).To(Equal(core.StringPtr("4s")))
				Expect(configPeerGossipModel.PublishStateInfoInterval).To(Equal(core.StringPtr("4s")))
				Expect(configPeerGossipModel.StateInfoRetentionInterval).To(Equal(core.StringPtr("0s")))
				Expect(configPeerGossipModel.PublishCertPeriod).To(Equal(core.StringPtr("10s")))
				Expect(configPeerGossipModel.SkipBlockVerification).To(Equal(core.BoolPtr(false)))
				Expect(configPeerGossipModel.DialTimeout).To(Equal(core.StringPtr("3s")))
				Expect(configPeerGossipModel.ConnTimeout).To(Equal(core.StringPtr("2s")))
				Expect(configPeerGossipModel.RecvBuffSize).To(Equal(core.Float64Ptr(float64(20))))
				Expect(configPeerGossipModel.SendBuffSize).To(Equal(core.Float64Ptr(float64(200))))
				Expect(configPeerGossipModel.DigestWaitTime).To(Equal(core.StringPtr("1s")))
				Expect(configPeerGossipModel.RequestWaitTime).To(Equal(core.StringPtr("1500ms")))
				Expect(configPeerGossipModel.ResponseWaitTime).To(Equal(core.StringPtr("2s")))
				Expect(configPeerGossipModel.AliveTimeInterval).To(Equal(core.StringPtr("5s")))
				Expect(configPeerGossipModel.AliveExpirationTimeout).To(Equal(core.StringPtr("25s")))
				Expect(configPeerGossipModel.ReconnectInterval).To(Equal(core.StringPtr("25s")))
				Expect(configPeerGossipModel.Election).To(Equal(configPeerGossipElectionModel))
				Expect(configPeerGossipModel.PvtData).To(Equal(configPeerGossipPvtDataModel))
				Expect(configPeerGossipModel.State).To(Equal(configPeerGossipStateModel))

				// Construct an instance of the ConfigPeerAuthentication model
				configPeerAuthenticationModel := new(blockchainv3.ConfigPeerAuthentication)
				Expect(configPeerAuthenticationModel).ToNot(BeNil())
				configPeerAuthenticationModel.Timewindow = core.StringPtr("15m")
				Expect(configPeerAuthenticationModel.Timewindow).To(Equal(core.StringPtr("15m")))

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv3.BccspSW)
				Expect(bccspSwModel).ToNot(BeNil())
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))
				Expect(bccspSwModel.Hash).To(Equal(core.StringPtr("SHA2")))
				Expect(bccspSwModel.Security).To(Equal(core.Float64Ptr(float64(256))))

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv3.BccspPKCS11)
				Expect(bccspPkcS11Model).ToNot(BeNil())
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))
				Expect(bccspPkcS11Model.Label).To(Equal(core.StringPtr("testString")))
				Expect(bccspPkcS11Model.Pin).To(Equal(core.StringPtr("testString")))
				Expect(bccspPkcS11Model.Hash).To(Equal(core.StringPtr("SHA2")))
				Expect(bccspPkcS11Model.Security).To(Equal(core.Float64Ptr(float64(256))))

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv3.Bccsp)
				Expect(bccspModel).ToNot(BeNil())
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model
				Expect(bccspModel.Default).To(Equal(core.StringPtr("SW")))
				Expect(bccspModel.SW).To(Equal(bccspSwModel))
				Expect(bccspModel.PKCS11).To(Equal(bccspPkcS11Model))

				// Construct an instance of the ConfigPeerClient model
				configPeerClientModel := new(blockchainv3.ConfigPeerClient)
				Expect(configPeerClientModel).ToNot(BeNil())
				configPeerClientModel.ConnTimeout = core.StringPtr("2s")
				Expect(configPeerClientModel.ConnTimeout).To(Equal(core.StringPtr("2s")))

				// Construct an instance of the ConfigPeerDeliveryclientAddressOverridesItem model
				configPeerDeliveryclientAddressOverridesItemModel := new(blockchainv3.ConfigPeerDeliveryclientAddressOverridesItem)
				Expect(configPeerDeliveryclientAddressOverridesItemModel).ToNot(BeNil())
				configPeerDeliveryclientAddressOverridesItemModel.From = core.StringPtr("n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.To = core.StringPtr("n3a3ec3-myorderer2.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.CaCertsFile = core.StringPtr("my-data/cert.pem")
				Expect(configPeerDeliveryclientAddressOverridesItemModel.From).To(Equal(core.StringPtr("n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")))
				Expect(configPeerDeliveryclientAddressOverridesItemModel.To).To(Equal(core.StringPtr("n3a3ec3-myorderer2.ibp.us-south.containers.appdomain.cloud:7050")))
				Expect(configPeerDeliveryclientAddressOverridesItemModel.CaCertsFile).To(Equal(core.StringPtr("my-data/cert.pem")))

				// Construct an instance of the ConfigPeerDeliveryclient model
				configPeerDeliveryclientModel := new(blockchainv3.ConfigPeerDeliveryclient)
				Expect(configPeerDeliveryclientModel).ToNot(BeNil())
				configPeerDeliveryclientModel.ReconnectTotalTimeThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.ConnTimeout = core.StringPtr("2s")
				configPeerDeliveryclientModel.ReConnectBackoffThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.AddressOverrides = []blockchainv3.ConfigPeerDeliveryclientAddressOverridesItem{*configPeerDeliveryclientAddressOverridesItemModel}
				Expect(configPeerDeliveryclientModel.ReconnectTotalTimeThreshold).To(Equal(core.StringPtr("60m")))
				Expect(configPeerDeliveryclientModel.ConnTimeout).To(Equal(core.StringPtr("2s")))
				Expect(configPeerDeliveryclientModel.ReConnectBackoffThreshold).To(Equal(core.StringPtr("60m")))
				Expect(configPeerDeliveryclientModel.AddressOverrides).To(Equal([]blockchainv3.ConfigPeerDeliveryclientAddressOverridesItem{*configPeerDeliveryclientAddressOverridesItemModel}))

				// Construct an instance of the ConfigPeerAdminService model
				configPeerAdminServiceModel := new(blockchainv3.ConfigPeerAdminService)
				Expect(configPeerAdminServiceModel).ToNot(BeNil())
				configPeerAdminServiceModel.ListenAddress = core.StringPtr("0.0.0.0:7051")
				Expect(configPeerAdminServiceModel.ListenAddress).To(Equal(core.StringPtr("0.0.0.0:7051")))

				// Construct an instance of the ConfigPeerDiscovery model
				configPeerDiscoveryModel := new(blockchainv3.ConfigPeerDiscovery)
				Expect(configPeerDiscoveryModel).ToNot(BeNil())
				configPeerDiscoveryModel.Enabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheEnabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheMaxSize = core.Float64Ptr(float64(1000))
				configPeerDiscoveryModel.AuthCachePurgeRetentionRatio = core.Float64Ptr(float64(0.75))
				configPeerDiscoveryModel.OrgMembersAllowedAccess = core.BoolPtr(false)
				Expect(configPeerDiscoveryModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(configPeerDiscoveryModel.AuthCacheEnabled).To(Equal(core.BoolPtr(true)))
				Expect(configPeerDiscoveryModel.AuthCacheMaxSize).To(Equal(core.Float64Ptr(float64(1000))))
				Expect(configPeerDiscoveryModel.AuthCachePurgeRetentionRatio).To(Equal(core.Float64Ptr(float64(0.75))))
				Expect(configPeerDiscoveryModel.OrgMembersAllowedAccess).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the ConfigPeerLimitsConcurrency model
				configPeerLimitsConcurrencyModel := new(blockchainv3.ConfigPeerLimitsConcurrency)
				Expect(configPeerLimitsConcurrencyModel).ToNot(BeNil())
				configPeerLimitsConcurrencyModel.EndorserService = core.Float64Ptr(float64(2500))
				configPeerLimitsConcurrencyModel.DeliverService = core.Float64Ptr(float64(2500))
				Expect(configPeerLimitsConcurrencyModel.EndorserService).To(Equal(core.Float64Ptr(float64(2500))))
				Expect(configPeerLimitsConcurrencyModel.DeliverService).To(Equal(core.Float64Ptr(float64(2500))))

				// Construct an instance of the ConfigPeerLimits model
				configPeerLimitsModel := new(blockchainv3.ConfigPeerLimits)
				Expect(configPeerLimitsModel).ToNot(BeNil())
				configPeerLimitsModel.Concurrency = configPeerLimitsConcurrencyModel
				Expect(configPeerLimitsModel.Concurrency).To(Equal(configPeerLimitsConcurrencyModel))

				// Construct an instance of the ConfigPeerGateway model
				configPeerGatewayModel := new(blockchainv3.ConfigPeerGateway)
				Expect(configPeerGatewayModel).ToNot(BeNil())
				configPeerGatewayModel.Enabled = core.BoolPtr(true)
				Expect(configPeerGatewayModel.Enabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the ConfigPeerCreatePeer model
				configPeerCreatePeerModel := new(blockchainv3.ConfigPeerCreatePeer)
				Expect(configPeerCreatePeerModel).ToNot(BeNil())
				configPeerCreatePeerModel.ID = core.StringPtr("john-doe")
				configPeerCreatePeerModel.NetworkID = core.StringPtr("dev")
				configPeerCreatePeerModel.Keepalive = configPeerKeepaliveModel
				configPeerCreatePeerModel.Gossip = configPeerGossipModel
				configPeerCreatePeerModel.Authentication = configPeerAuthenticationModel
				configPeerCreatePeerModel.BCCSP = bccspModel
				configPeerCreatePeerModel.Client = configPeerClientModel
				configPeerCreatePeerModel.Deliveryclient = configPeerDeliveryclientModel
				configPeerCreatePeerModel.AdminService = configPeerAdminServiceModel
				configPeerCreatePeerModel.ValidatorPoolSize = core.Float64Ptr(float64(8))
				configPeerCreatePeerModel.Discovery = configPeerDiscoveryModel
				configPeerCreatePeerModel.Limits = configPeerLimitsModel
				configPeerCreatePeerModel.Gateway = configPeerGatewayModel
				Expect(configPeerCreatePeerModel.ID).To(Equal(core.StringPtr("john-doe")))
				Expect(configPeerCreatePeerModel.NetworkID).To(Equal(core.StringPtr("dev")))
				Expect(configPeerCreatePeerModel.Keepalive).To(Equal(configPeerKeepaliveModel))
				Expect(configPeerCreatePeerModel.Gossip).To(Equal(configPeerGossipModel))
				Expect(configPeerCreatePeerModel.Authentication).To(Equal(configPeerAuthenticationModel))
				Expect(configPeerCreatePeerModel.BCCSP).To(Equal(bccspModel))
				Expect(configPeerCreatePeerModel.Client).To(Equal(configPeerClientModel))
				Expect(configPeerCreatePeerModel.Deliveryclient).To(Equal(configPeerDeliveryclientModel))
				Expect(configPeerCreatePeerModel.AdminService).To(Equal(configPeerAdminServiceModel))
				Expect(configPeerCreatePeerModel.ValidatorPoolSize).To(Equal(core.Float64Ptr(float64(8))))
				Expect(configPeerCreatePeerModel.Discovery).To(Equal(configPeerDiscoveryModel))
				Expect(configPeerCreatePeerModel.Limits).To(Equal(configPeerLimitsModel))
				Expect(configPeerCreatePeerModel.Gateway).To(Equal(configPeerGatewayModel))

				// Construct an instance of the ConfigPeerChaincodeGolang model
				configPeerChaincodeGolangModel := new(blockchainv3.ConfigPeerChaincodeGolang)
				Expect(configPeerChaincodeGolangModel).ToNot(BeNil())
				configPeerChaincodeGolangModel.DynamicLink = core.BoolPtr(false)
				Expect(configPeerChaincodeGolangModel.DynamicLink).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the ConfigPeerChaincodeExternalBuildersItem model
				configPeerChaincodeExternalBuildersItemModel := new(blockchainv3.ConfigPeerChaincodeExternalBuildersItem)
				Expect(configPeerChaincodeExternalBuildersItemModel).ToNot(BeNil())
				configPeerChaincodeExternalBuildersItemModel.Path = core.StringPtr("/path/to/directory")
				configPeerChaincodeExternalBuildersItemModel.Name = core.StringPtr("descriptive-build-name")
				configPeerChaincodeExternalBuildersItemModel.EnvironmentWhitelist = []string{"GOPROXY"}
				Expect(configPeerChaincodeExternalBuildersItemModel.Path).To(Equal(core.StringPtr("/path/to/directory")))
				Expect(configPeerChaincodeExternalBuildersItemModel.Name).To(Equal(core.StringPtr("descriptive-build-name")))
				Expect(configPeerChaincodeExternalBuildersItemModel.EnvironmentWhitelist).To(Equal([]string{"GOPROXY"}))

				// Construct an instance of the ConfigPeerChaincodeSystem model
				configPeerChaincodeSystemModel := new(blockchainv3.ConfigPeerChaincodeSystem)
				Expect(configPeerChaincodeSystemModel).ToNot(BeNil())
				configPeerChaincodeSystemModel.Cscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Lscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Escc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Vscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Qscc = core.BoolPtr(true)
				Expect(configPeerChaincodeSystemModel.Cscc).To(Equal(core.BoolPtr(true)))
				Expect(configPeerChaincodeSystemModel.Lscc).To(Equal(core.BoolPtr(true)))
				Expect(configPeerChaincodeSystemModel.Escc).To(Equal(core.BoolPtr(true)))
				Expect(configPeerChaincodeSystemModel.Vscc).To(Equal(core.BoolPtr(true)))
				Expect(configPeerChaincodeSystemModel.Qscc).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the ConfigPeerChaincodeLogging model
				configPeerChaincodeLoggingModel := new(blockchainv3.ConfigPeerChaincodeLogging)
				Expect(configPeerChaincodeLoggingModel).ToNot(BeNil())
				configPeerChaincodeLoggingModel.Level = core.StringPtr("info")
				configPeerChaincodeLoggingModel.Shim = core.StringPtr("warning")
				configPeerChaincodeLoggingModel.Format = core.StringPtr("%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}")
				Expect(configPeerChaincodeLoggingModel.Level).To(Equal(core.StringPtr("info")))
				Expect(configPeerChaincodeLoggingModel.Shim).To(Equal(core.StringPtr("warning")))
				Expect(configPeerChaincodeLoggingModel.Format).To(Equal(core.StringPtr("%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}")))

				// Construct an instance of the ConfigPeerChaincode model
				configPeerChaincodeModel := new(blockchainv3.ConfigPeerChaincode)
				Expect(configPeerChaincodeModel).ToNot(BeNil())
				configPeerChaincodeModel.Golang = configPeerChaincodeGolangModel
				configPeerChaincodeModel.ExternalBuilders = []blockchainv3.ConfigPeerChaincodeExternalBuildersItem{*configPeerChaincodeExternalBuildersItemModel}
				configPeerChaincodeModel.InstallTimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Startuptimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Executetimeout = core.StringPtr("30s")
				configPeerChaincodeModel.System = configPeerChaincodeSystemModel
				configPeerChaincodeModel.Logging = configPeerChaincodeLoggingModel
				Expect(configPeerChaincodeModel.Golang).To(Equal(configPeerChaincodeGolangModel))
				Expect(configPeerChaincodeModel.ExternalBuilders).To(Equal([]blockchainv3.ConfigPeerChaincodeExternalBuildersItem{*configPeerChaincodeExternalBuildersItemModel}))
				Expect(configPeerChaincodeModel.InstallTimeout).To(Equal(core.StringPtr("300s")))
				Expect(configPeerChaincodeModel.Startuptimeout).To(Equal(core.StringPtr("300s")))
				Expect(configPeerChaincodeModel.Executetimeout).To(Equal(core.StringPtr("30s")))
				Expect(configPeerChaincodeModel.System).To(Equal(configPeerChaincodeSystemModel))
				Expect(configPeerChaincodeModel.Logging).To(Equal(configPeerChaincodeLoggingModel))

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv3.MetricsStatsd)
				Expect(metricsStatsdModel).ToNot(BeNil())
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")
				Expect(metricsStatsdModel.Network).To(Equal(core.StringPtr("udp")))
				Expect(metricsStatsdModel.Address).To(Equal(core.StringPtr("127.0.0.1:8125")))
				Expect(metricsStatsdModel.WriteInterval).To(Equal(core.StringPtr("10s")))
				Expect(metricsStatsdModel.Prefix).To(Equal(core.StringPtr("server")))

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv3.Metrics)
				Expect(metricsModel).ToNot(BeNil())
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel
				Expect(metricsModel.Provider).To(Equal(core.StringPtr("prometheus")))
				Expect(metricsModel.Statsd).To(Equal(metricsStatsdModel))

				// Construct an instance of the ConfigPeerCreate model
				configPeerCreateModel := new(blockchainv3.ConfigPeerCreate)
				Expect(configPeerCreateModel).ToNot(BeNil())
				configPeerCreateModel.Peer = configPeerCreatePeerModel
				configPeerCreateModel.Chaincode = configPeerChaincodeModel
				configPeerCreateModel.Metrics = metricsModel
				Expect(configPeerCreateModel.Peer).To(Equal(configPeerCreatePeerModel))
				Expect(configPeerCreateModel.Chaincode).To(Equal(configPeerChaincodeModel))
				Expect(configPeerCreateModel.Metrics).To(Equal(metricsModel))

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv3.ResourceRequests)
				Expect(resourceRequestsModel).ToNot(BeNil())
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")
				Expect(resourceRequestsModel.Cpu).To(Equal(core.StringPtr("100m")))
				Expect(resourceRequestsModel.Memory).To(Equal(core.StringPtr("256MiB")))

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv3.ResourceLimits)
				Expect(resourceLimitsModel).ToNot(BeNil())
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")
				Expect(resourceLimitsModel.Cpu).To(Equal(core.StringPtr("100m")))
				Expect(resourceLimitsModel.Memory).To(Equal(core.StringPtr("256MiB")))

				// Construct an instance of the ResourceObjectFabV2 model
				resourceObjectFabV2Model := new(blockchainv3.ResourceObjectFabV2)
				Expect(resourceObjectFabV2Model).ToNot(BeNil())
				resourceObjectFabV2Model.Requests = resourceRequestsModel
				resourceObjectFabV2Model.Limits = resourceLimitsModel
				Expect(resourceObjectFabV2Model.Requests).To(Equal(resourceRequestsModel))
				Expect(resourceObjectFabV2Model.Limits).To(Equal(resourceLimitsModel))

				// Construct an instance of the ResourceObjectCouchDb model
				resourceObjectCouchDbModel := new(blockchainv3.ResourceObjectCouchDb)
				Expect(resourceObjectCouchDbModel).ToNot(BeNil())
				resourceObjectCouchDbModel.Requests = resourceRequestsModel
				resourceObjectCouchDbModel.Limits = resourceLimitsModel
				Expect(resourceObjectCouchDbModel.Requests).To(Equal(resourceRequestsModel))
				Expect(resourceObjectCouchDbModel.Limits).To(Equal(resourceLimitsModel))

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv3.ResourceObject)
				Expect(resourceObjectModel).ToNot(BeNil())
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel
				Expect(resourceObjectModel.Requests).To(Equal(resourceRequestsModel))
				Expect(resourceObjectModel.Limits).To(Equal(resourceLimitsModel))

				// Construct an instance of the ResourceObjectFabV1 model
				resourceObjectFabV1Model := new(blockchainv3.ResourceObjectFabV1)
				Expect(resourceObjectFabV1Model).ToNot(BeNil())
				resourceObjectFabV1Model.Requests = resourceRequestsModel
				resourceObjectFabV1Model.Limits = resourceLimitsModel
				Expect(resourceObjectFabV1Model.Requests).To(Equal(resourceRequestsModel))
				Expect(resourceObjectFabV1Model.Limits).To(Equal(resourceLimitsModel))

				// Construct an instance of the PeerResources model
				peerResourcesModel := new(blockchainv3.PeerResources)
				Expect(peerResourcesModel).ToNot(BeNil())
				peerResourcesModel.Chaincodelauncher = resourceObjectFabV2Model
				peerResourcesModel.Couchdb = resourceObjectCouchDbModel
				peerResourcesModel.Statedb = resourceObjectModel
				peerResourcesModel.Dind = resourceObjectFabV1Model
				peerResourcesModel.Fluentd = resourceObjectFabV1Model
				peerResourcesModel.Peer = resourceObjectModel
				peerResourcesModel.Proxy = resourceObjectModel
				Expect(peerResourcesModel.Chaincodelauncher).To(Equal(resourceObjectFabV2Model))
				Expect(peerResourcesModel.Couchdb).To(Equal(resourceObjectCouchDbModel))
				Expect(peerResourcesModel.Statedb).To(Equal(resourceObjectModel))
				Expect(peerResourcesModel.Dind).To(Equal(resourceObjectFabV1Model))
				Expect(peerResourcesModel.Fluentd).To(Equal(resourceObjectFabV1Model))
				Expect(peerResourcesModel.Peer).To(Equal(resourceObjectModel))
				Expect(peerResourcesModel.Proxy).To(Equal(resourceObjectModel))

				// Construct an instance of the StorageObject model
				storageObjectModel := new(blockchainv3.StorageObject)
				Expect(storageObjectModel).ToNot(BeNil())
				storageObjectModel.Size = core.StringPtr("4GiB")
				storageObjectModel.Class = core.StringPtr("default")
				Expect(storageObjectModel.Size).To(Equal(core.StringPtr("4GiB")))
				Expect(storageObjectModel.Class).To(Equal(core.StringPtr("default")))

				// Construct an instance of the CreatePeerBodyStorage model
				createPeerBodyStorageModel := new(blockchainv3.CreatePeerBodyStorage)
				Expect(createPeerBodyStorageModel).ToNot(BeNil())
				createPeerBodyStorageModel.Peer = storageObjectModel
				createPeerBodyStorageModel.Statedb = storageObjectModel
				Expect(createPeerBodyStorageModel.Peer).To(Equal(storageObjectModel))
				Expect(createPeerBodyStorageModel.Statedb).To(Equal(storageObjectModel))

				// Construct an instance of the Hsm model
				hsmModel := new(blockchainv3.Hsm)
				Expect(hsmModel).ToNot(BeNil())
				hsmModel.Pkcs11endpoint = core.StringPtr("tcp://example.com:666")
				Expect(hsmModel.Pkcs11endpoint).To(Equal(core.StringPtr("tcp://example.com:666")))

				// Construct an instance of the CreatePeerOptions model
				createPeerOptionsMspID := "Org1"
				createPeerOptionsDisplayName := "My Peer"
				var createPeerOptionsCrypto *blockchainv3.CryptoObject = nil
				createPeerOptionsModel := blockchainService.NewCreatePeerOptions(createPeerOptionsMspID, createPeerOptionsDisplayName, createPeerOptionsCrypto)
				createPeerOptionsModel.SetMspID("Org1")
				createPeerOptionsModel.SetDisplayName("My Peer")
				createPeerOptionsModel.SetCrypto(cryptoObjectModel)
				createPeerOptionsModel.SetID("component1")
				createPeerOptionsModel.SetConfigOverride(configPeerCreateModel)
				createPeerOptionsModel.SetResources(peerResourcesModel)
				createPeerOptionsModel.SetStorage(createPeerBodyStorageModel)
				createPeerOptionsModel.SetZone("-")
				createPeerOptionsModel.SetStateDb("couchdb")
				createPeerOptionsModel.SetTags([]string{"fabric-ca"})
				createPeerOptionsModel.SetHsm(hsmModel)
				createPeerOptionsModel.SetRegion("-")
				createPeerOptionsModel.SetVersion("1.4.6-1")
				createPeerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createPeerOptionsModel).ToNot(BeNil())
				Expect(createPeerOptionsModel.MspID).To(Equal(core.StringPtr("Org1")))
				Expect(createPeerOptionsModel.DisplayName).To(Equal(core.StringPtr("My Peer")))
				Expect(createPeerOptionsModel.Crypto).To(Equal(cryptoObjectModel))
				Expect(createPeerOptionsModel.ID).To(Equal(core.StringPtr("component1")))
				Expect(createPeerOptionsModel.ConfigOverride).To(Equal(configPeerCreateModel))
				Expect(createPeerOptionsModel.Resources).To(Equal(peerResourcesModel))
				Expect(createPeerOptionsModel.Storage).To(Equal(createPeerBodyStorageModel))
				Expect(createPeerOptionsModel.Zone).To(Equal(core.StringPtr("-")))
				Expect(createPeerOptionsModel.StateDb).To(Equal(core.StringPtr("couchdb")))
				Expect(createPeerOptionsModel.Tags).To(Equal([]string{"fabric-ca"}))
				Expect(createPeerOptionsModel.Hsm).To(Equal(hsmModel))
				Expect(createPeerOptionsModel.Region).To(Equal(core.StringPtr("-")))
				Expect(createPeerOptionsModel.Version).To(Equal(core.StringPtr("1.4.6-1")))
				Expect(createPeerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCryptoObjectEnrollment successfully`, func() {
				var component *blockchainv3.CryptoEnrollmentComponent = nil
				var ca *blockchainv3.CryptoObjectEnrollmentCa = nil
				var tlsca *blockchainv3.CryptoObjectEnrollmentTlsca = nil
				_, err := blockchainService.NewCryptoObjectEnrollment(component, ca, tlsca)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCryptoObjectEnrollmentCa successfully`, func() {
				host := "n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud"
				port := float64(7054)
				name := "ca"
				tlsCert := "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
				enrollID := "admin"
				enrollSecret := "password"
				model, err := blockchainService.NewCryptoObjectEnrollmentCa(host, port, name, tlsCert, enrollID, enrollSecret)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCryptoObjectEnrollmentTlsca successfully`, func() {
				host := "n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud"
				port := float64(7054)
				name := "tlsca"
				tlsCert := "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
				enrollID := "admin"
				enrollSecret := "password"
				model, err := blockchainService.NewCryptoObjectEnrollmentTlsca(host, port, name, tlsCert, enrollID, enrollSecret)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCryptoObjectMsp successfully`, func() {
				var component *blockchainv3.MspCryptoComp = nil
				var ca *blockchainv3.MspCryptoCa = nil
				var tlsca *blockchainv3.MspCryptoCa = nil
				_, err := blockchainService.NewCryptoObjectMsp(component, ca, tlsca)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewDeleteAllComponentsOptions successfully`, func() {
				// Construct an instance of the DeleteAllComponentsOptions model
				deleteAllComponentsOptionsModel := blockchainService.NewDeleteAllComponentsOptions()
				deleteAllComponentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAllComponentsOptionsModel).ToNot(BeNil())
				Expect(deleteAllComponentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAllNotificationsOptions successfully`, func() {
				// Construct an instance of the DeleteAllNotificationsOptions model
				deleteAllNotificationsOptionsModel := blockchainService.NewDeleteAllNotificationsOptions()
				deleteAllNotificationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAllNotificationsOptionsModel).ToNot(BeNil())
				Expect(deleteAllNotificationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAllSessionsOptions successfully`, func() {
				// Construct an instance of the DeleteAllSessionsOptions model
				deleteAllSessionsOptionsModel := blockchainService.NewDeleteAllSessionsOptions()
				deleteAllSessionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAllSessionsOptionsModel).ToNot(BeNil())
				Expect(deleteAllSessionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteComponentOptions successfully`, func() {
				// Construct an instance of the DeleteComponentOptions model
				id := "testString"
				deleteComponentOptionsModel := blockchainService.NewDeleteComponentOptions(id)
				deleteComponentOptionsModel.SetID("testString")
				deleteComponentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteComponentOptionsModel).ToNot(BeNil())
				Expect(deleteComponentOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteComponentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteComponentsByTagOptions successfully`, func() {
				// Construct an instance of the DeleteComponentsByTagOptions model
				tag := "testString"
				deleteComponentsByTagOptionsModel := blockchainService.NewDeleteComponentsByTagOptions(tag)
				deleteComponentsByTagOptionsModel.SetTag("testString")
				deleteComponentsByTagOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteComponentsByTagOptionsModel).ToNot(BeNil())
				Expect(deleteComponentsByTagOptionsModel.Tag).To(Equal(core.StringPtr("testString")))
				Expect(deleteComponentsByTagOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteSigTxOptions successfully`, func() {
				// Construct an instance of the DeleteSigTxOptions model
				id := "testString"
				deleteSigTxOptionsModel := blockchainService.NewDeleteSigTxOptions(id)
				deleteSigTxOptionsModel.SetID("testString")
				deleteSigTxOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteSigTxOptionsModel).ToNot(BeNil())
				Expect(deleteSigTxOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSigTxOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEditAdminCertsOptions successfully`, func() {
				// Construct an instance of the EditAdminCertsOptions model
				id := "testString"
				editAdminCertsOptionsModel := blockchainService.NewEditAdminCertsOptions(id)
				editAdminCertsOptionsModel.SetID("testString")
				editAdminCertsOptionsModel.SetAppendAdminCerts([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="})
				editAdminCertsOptionsModel.SetRemoveAdminCerts([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="})
				editAdminCertsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(editAdminCertsOptionsModel).ToNot(BeNil())
				Expect(editAdminCertsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(editAdminCertsOptionsModel.AppendAdminCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(editAdminCertsOptionsModel.RemoveAdminCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(editAdminCertsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEditCaOptions successfully`, func() {
				// Construct an instance of the EditCaOptions model
				id := "testString"
				editCaOptionsModel := blockchainService.NewEditCaOptions(id)
				editCaOptionsModel.SetID("testString")
				editCaOptionsModel.SetDisplayName("My CA")
				editCaOptionsModel.SetApiURL("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:7054")
				editCaOptionsModel.SetOperationsURL("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443")
				editCaOptionsModel.SetCaName("ca")
				editCaOptionsModel.SetLocation("ibmcloud")
				editCaOptionsModel.SetTags([]string{"fabric-ca"})
				editCaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(editCaOptionsModel).ToNot(BeNil())
				Expect(editCaOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(editCaOptionsModel.DisplayName).To(Equal(core.StringPtr("My CA")))
				Expect(editCaOptionsModel.ApiURL).To(Equal(core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:7054")))
				Expect(editCaOptionsModel.OperationsURL).To(Equal(core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443")))
				Expect(editCaOptionsModel.CaName).To(Equal(core.StringPtr("ca")))
				Expect(editCaOptionsModel.Location).To(Equal(core.StringPtr("ibmcloud")))
				Expect(editCaOptionsModel.Tags).To(Equal([]string{"fabric-ca"}))
				Expect(editCaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEditMspOptions successfully`, func() {
				// Construct an instance of the EditMspOptions model
				id := "testString"
				editMspOptionsModel := blockchainService.NewEditMspOptions(id)
				editMspOptionsModel.SetID("testString")
				editMspOptionsModel.SetMspID("Org1")
				editMspOptionsModel.SetDisplayName("My Peer")
				editMspOptionsModel.SetRootCerts([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="})
				editMspOptionsModel.SetIntermediateCerts([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkRhdGEgaGVyZSBpZiB0aGlzIHdhcyByZWFsCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"})
				editMspOptionsModel.SetAdmins([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="})
				editMspOptionsModel.SetTlsRootCerts([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="})
				editMspOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(editMspOptionsModel).ToNot(BeNil())
				Expect(editMspOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(editMspOptionsModel.MspID).To(Equal(core.StringPtr("Org1")))
				Expect(editMspOptionsModel.DisplayName).To(Equal(core.StringPtr("My Peer")))
				Expect(editMspOptionsModel.RootCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(editMspOptionsModel.IntermediateCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkRhdGEgaGVyZSBpZiB0aGlzIHdhcyByZWFsCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"}))
				Expect(editMspOptionsModel.Admins).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(editMspOptionsModel.TlsRootCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(editMspOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEditOrdererOptions successfully`, func() {
				// Construct an instance of the EditOrdererOptions model
				id := "testString"
				editOrdererOptionsModel := blockchainService.NewEditOrdererOptions(id)
				editOrdererOptionsModel.SetID("testString")
				editOrdererOptionsModel.SetClusterName("ordering service 1")
				editOrdererOptionsModel.SetDisplayName("orderer")
				editOrdererOptionsModel.SetApiURL("grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				editOrdererOptionsModel.SetOperationsURL("https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443")
				editOrdererOptionsModel.SetGrpcwpURL("https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443")
				editOrdererOptionsModel.SetMspID("Org1")
				editOrdererOptionsModel.SetConsenterProposalFin(true)
				editOrdererOptionsModel.SetLocation("ibmcloud")
				editOrdererOptionsModel.SetSystemChannelID("testchainid")
				editOrdererOptionsModel.SetTags([]string{"fabric-ca"})
				editOrdererOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(editOrdererOptionsModel).ToNot(BeNil())
				Expect(editOrdererOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(editOrdererOptionsModel.ClusterName).To(Equal(core.StringPtr("ordering service 1")))
				Expect(editOrdererOptionsModel.DisplayName).To(Equal(core.StringPtr("orderer")))
				Expect(editOrdererOptionsModel.ApiURL).To(Equal(core.StringPtr("grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")))
				Expect(editOrdererOptionsModel.OperationsURL).To(Equal(core.StringPtr("https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443")))
				Expect(editOrdererOptionsModel.GrpcwpURL).To(Equal(core.StringPtr("https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443")))
				Expect(editOrdererOptionsModel.MspID).To(Equal(core.StringPtr("Org1")))
				Expect(editOrdererOptionsModel.ConsenterProposalFin).To(Equal(core.BoolPtr(true)))
				Expect(editOrdererOptionsModel.Location).To(Equal(core.StringPtr("ibmcloud")))
				Expect(editOrdererOptionsModel.SystemChannelID).To(Equal(core.StringPtr("testchainid")))
				Expect(editOrdererOptionsModel.Tags).To(Equal([]string{"fabric-ca"}))
				Expect(editOrdererOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEditPeerOptions successfully`, func() {
				// Construct an instance of the EditPeerOptions model
				id := "testString"
				editPeerOptionsModel := blockchainService.NewEditPeerOptions(id)
				editPeerOptionsModel.SetID("testString")
				editPeerOptionsModel.SetDisplayName("My Peer")
				editPeerOptionsModel.SetApiURL("grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051")
				editPeerOptionsModel.SetOperationsURL("https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443")
				editPeerOptionsModel.SetGrpcwpURL("https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084")
				editPeerOptionsModel.SetMspID("Org1")
				editPeerOptionsModel.SetLocation("ibmcloud")
				editPeerOptionsModel.SetTags([]string{"fabric-ca"})
				editPeerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(editPeerOptionsModel).ToNot(BeNil())
				Expect(editPeerOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(editPeerOptionsModel.DisplayName).To(Equal(core.StringPtr("My Peer")))
				Expect(editPeerOptionsModel.ApiURL).To(Equal(core.StringPtr("grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051")))
				Expect(editPeerOptionsModel.OperationsURL).To(Equal(core.StringPtr("https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443")))
				Expect(editPeerOptionsModel.GrpcwpURL).To(Equal(core.StringPtr("https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084")))
				Expect(editPeerOptionsModel.MspID).To(Equal(core.StringPtr("Org1")))
				Expect(editPeerOptionsModel.Location).To(Equal(core.StringPtr("ibmcloud")))
				Expect(editPeerOptionsModel.Tags).To(Equal([]string{"fabric-ca"}))
				Expect(editPeerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEditSettingsOptions successfully`, func() {
				// Construct an instance of the EditSettingsBodyInactivityTimeouts model
				editSettingsBodyInactivityTimeoutsModel := new(blockchainv3.EditSettingsBodyInactivityTimeouts)
				Expect(editSettingsBodyInactivityTimeoutsModel).ToNot(BeNil())
				editSettingsBodyInactivityTimeoutsModel.Enabled = core.BoolPtr(false)
				editSettingsBodyInactivityTimeoutsModel.MaxIdleTime = core.Float64Ptr(float64(90000))
				Expect(editSettingsBodyInactivityTimeoutsModel.Enabled).To(Equal(core.BoolPtr(false)))
				Expect(editSettingsBodyInactivityTimeoutsModel.MaxIdleTime).To(Equal(core.Float64Ptr(float64(90000))))

				// Construct an instance of the LoggingSettingsClient model
				loggingSettingsClientModel := new(blockchainv3.LoggingSettingsClient)
				Expect(loggingSettingsClientModel).ToNot(BeNil())
				loggingSettingsClientModel.Enabled = core.BoolPtr(true)
				loggingSettingsClientModel.Level = core.StringPtr("silly")
				loggingSettingsClientModel.UniqueName = core.BoolPtr(false)
				Expect(loggingSettingsClientModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(loggingSettingsClientModel.Level).To(Equal(core.StringPtr("silly")))
				Expect(loggingSettingsClientModel.UniqueName).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the LoggingSettingsServer model
				loggingSettingsServerModel := new(blockchainv3.LoggingSettingsServer)
				Expect(loggingSettingsServerModel).ToNot(BeNil())
				loggingSettingsServerModel.Enabled = core.BoolPtr(true)
				loggingSettingsServerModel.Level = core.StringPtr("silly")
				loggingSettingsServerModel.UniqueName = core.BoolPtr(false)
				Expect(loggingSettingsServerModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(loggingSettingsServerModel.Level).To(Equal(core.StringPtr("silly")))
				Expect(loggingSettingsServerModel.UniqueName).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the EditLogSettingsBody model
				editLogSettingsBodyModel := new(blockchainv3.EditLogSettingsBody)
				Expect(editLogSettingsBodyModel).ToNot(BeNil())
				editLogSettingsBodyModel.Client = loggingSettingsClientModel
				editLogSettingsBodyModel.Server = loggingSettingsServerModel
				Expect(editLogSettingsBodyModel.Client).To(Equal(loggingSettingsClientModel))
				Expect(editLogSettingsBodyModel.Server).To(Equal(loggingSettingsServerModel))

				// Construct an instance of the EditSettingsOptions model
				editSettingsOptionsModel := blockchainService.NewEditSettingsOptions()
				editSettingsOptionsModel.SetInactivityTimeouts(editSettingsBodyInactivityTimeoutsModel)
				editSettingsOptionsModel.SetFileLogging(editLogSettingsBodyModel)
				editSettingsOptionsModel.SetMaxReqPerMin(float64(25))
				editSettingsOptionsModel.SetMaxReqPerMinAk(float64(25))
				editSettingsOptionsModel.SetFabricGetBlockTimeoutMs(float64(10000))
				editSettingsOptionsModel.SetFabricInstantiateTimeoutMs(float64(300000))
				editSettingsOptionsModel.SetFabricJoinChannelTimeoutMs(float64(25000))
				editSettingsOptionsModel.SetFabricInstallCcTimeoutMs(float64(300000))
				editSettingsOptionsModel.SetFabricLcInstallCcTimeoutMs(float64(300000))
				editSettingsOptionsModel.SetFabricLcGetCcTimeoutMs(float64(180000))
				editSettingsOptionsModel.SetFabricGeneralTimeoutMs(float64(10000))
				editSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(editSettingsOptionsModel).ToNot(BeNil())
				Expect(editSettingsOptionsModel.InactivityTimeouts).To(Equal(editSettingsBodyInactivityTimeoutsModel))
				Expect(editSettingsOptionsModel.FileLogging).To(Equal(editLogSettingsBodyModel))
				Expect(editSettingsOptionsModel.MaxReqPerMin).To(Equal(core.Float64Ptr(float64(25))))
				Expect(editSettingsOptionsModel.MaxReqPerMinAk).To(Equal(core.Float64Ptr(float64(25))))
				Expect(editSettingsOptionsModel.FabricGetBlockTimeoutMs).To(Equal(core.Float64Ptr(float64(10000))))
				Expect(editSettingsOptionsModel.FabricInstantiateTimeoutMs).To(Equal(core.Float64Ptr(float64(300000))))
				Expect(editSettingsOptionsModel.FabricJoinChannelTimeoutMs).To(Equal(core.Float64Ptr(float64(25000))))
				Expect(editSettingsOptionsModel.FabricInstallCcTimeoutMs).To(Equal(core.Float64Ptr(float64(300000))))
				Expect(editSettingsOptionsModel.FabricLcInstallCcTimeoutMs).To(Equal(core.Float64Ptr(float64(300000))))
				Expect(editSettingsOptionsModel.FabricLcGetCcTimeoutMs).To(Equal(core.Float64Ptr(float64(180000))))
				Expect(editSettingsOptionsModel.FabricGeneralTimeoutMs).To(Equal(core.Float64Ptr(float64(10000))))
				Expect(editSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetComponentOptions successfully`, func() {
				// Construct an instance of the GetComponentOptions model
				id := "testString"
				getComponentOptionsModel := blockchainService.NewGetComponentOptions(id)
				getComponentOptionsModel.SetID("testString")
				getComponentOptionsModel.SetDeploymentAttrs("included")
				getComponentOptionsModel.SetParsedCerts("included")
				getComponentOptionsModel.SetCache("skip")
				getComponentOptionsModel.SetCaAttrs("included")
				getComponentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getComponentOptionsModel).ToNot(BeNil())
				Expect(getComponentOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getComponentOptionsModel.DeploymentAttrs).To(Equal(core.StringPtr("included")))
				Expect(getComponentOptionsModel.ParsedCerts).To(Equal(core.StringPtr("included")))
				Expect(getComponentOptionsModel.Cache).To(Equal(core.StringPtr("skip")))
				Expect(getComponentOptionsModel.CaAttrs).To(Equal(core.StringPtr("included")))
				Expect(getComponentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetComponentsByTagOptions successfully`, func() {
				// Construct an instance of the GetComponentsByTagOptions model
				tag := "testString"
				getComponentsByTagOptionsModel := blockchainService.NewGetComponentsByTagOptions(tag)
				getComponentsByTagOptionsModel.SetTag("testString")
				getComponentsByTagOptionsModel.SetDeploymentAttrs("included")
				getComponentsByTagOptionsModel.SetParsedCerts("included")
				getComponentsByTagOptionsModel.SetCache("skip")
				getComponentsByTagOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getComponentsByTagOptionsModel).ToNot(BeNil())
				Expect(getComponentsByTagOptionsModel.Tag).To(Equal(core.StringPtr("testString")))
				Expect(getComponentsByTagOptionsModel.DeploymentAttrs).To(Equal(core.StringPtr("included")))
				Expect(getComponentsByTagOptionsModel.ParsedCerts).To(Equal(core.StringPtr("included")))
				Expect(getComponentsByTagOptionsModel.Cache).To(Equal(core.StringPtr("skip")))
				Expect(getComponentsByTagOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetComponentsByTypeOptions successfully`, func() {
				// Construct an instance of the GetComponentsByTypeOptions model
				typeVar := "fabric-peer"
				getComponentsByTypeOptionsModel := blockchainService.NewGetComponentsByTypeOptions(typeVar)
				getComponentsByTypeOptionsModel.SetType("fabric-peer")
				getComponentsByTypeOptionsModel.SetDeploymentAttrs("included")
				getComponentsByTypeOptionsModel.SetParsedCerts("included")
				getComponentsByTypeOptionsModel.SetCache("skip")
				getComponentsByTypeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getComponentsByTypeOptionsModel).ToNot(BeNil())
				Expect(getComponentsByTypeOptionsModel.Type).To(Equal(core.StringPtr("fabric-peer")))
				Expect(getComponentsByTypeOptionsModel.DeploymentAttrs).To(Equal(core.StringPtr("included")))
				Expect(getComponentsByTypeOptionsModel.ParsedCerts).To(Equal(core.StringPtr("included")))
				Expect(getComponentsByTypeOptionsModel.Cache).To(Equal(core.StringPtr("skip")))
				Expect(getComponentsByTypeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetFabVersionsOptions successfully`, func() {
				// Construct an instance of the GetFabVersionsOptions model
				getFabVersionsOptionsModel := blockchainService.NewGetFabVersionsOptions()
				getFabVersionsOptionsModel.SetCache("skip")
				getFabVersionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getFabVersionsOptionsModel).ToNot(BeNil())
				Expect(getFabVersionsOptionsModel.Cache).To(Equal(core.StringPtr("skip")))
				Expect(getFabVersionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetHealthOptions successfully`, func() {
				// Construct an instance of the GetHealthOptions model
				getHealthOptionsModel := blockchainService.NewGetHealthOptions()
				getHealthOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getHealthOptionsModel).ToNot(BeNil())
				Expect(getHealthOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetMspCertificateOptions successfully`, func() {
				// Construct an instance of the GetMspCertificateOptions model
				mspID := "testString"
				getMspCertificateOptionsModel := blockchainService.NewGetMspCertificateOptions(mspID)
				getMspCertificateOptionsModel.SetMspID("testString")
				getMspCertificateOptionsModel.SetCache("skip")
				getMspCertificateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getMspCertificateOptionsModel).ToNot(BeNil())
				Expect(getMspCertificateOptionsModel.MspID).To(Equal(core.StringPtr("testString")))
				Expect(getMspCertificateOptionsModel.Cache).To(Equal(core.StringPtr("skip")))
				Expect(getMspCertificateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPostmanOptions successfully`, func() {
				// Construct an instance of the GetPostmanOptions model
				authType := "bearer"
				getPostmanOptionsModel := blockchainService.NewGetPostmanOptions(authType)
				getPostmanOptionsModel.SetAuthType("bearer")
				getPostmanOptionsModel.SetToken("testString")
				getPostmanOptionsModel.SetApiKey("testString")
				getPostmanOptionsModel.SetUsername("admin")
				getPostmanOptionsModel.SetPassword("password")
				getPostmanOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPostmanOptionsModel).ToNot(BeNil())
				Expect(getPostmanOptionsModel.AuthType).To(Equal(core.StringPtr("bearer")))
				Expect(getPostmanOptionsModel.Token).To(Equal(core.StringPtr("testString")))
				Expect(getPostmanOptionsModel.ApiKey).To(Equal(core.StringPtr("testString")))
				Expect(getPostmanOptionsModel.Username).To(Equal(core.StringPtr("admin")))
				Expect(getPostmanOptionsModel.Password).To(Equal(core.StringPtr("password")))
				Expect(getPostmanOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSettingsOptions successfully`, func() {
				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := blockchainService.NewGetSettingsOptions()
				getSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSettingsOptionsModel).ToNot(BeNil())
				Expect(getSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSwaggerOptions successfully`, func() {
				// Construct an instance of the GetSwaggerOptions model
				getSwaggerOptionsModel := blockchainService.NewGetSwaggerOptions()
				getSwaggerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSwaggerOptionsModel).ToNot(BeNil())
				Expect(getSwaggerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewImportCaBodyMsp successfully`, func() {
				var ca *blockchainv3.ImportCaBodyMspCa = nil
				var tlsca *blockchainv3.ImportCaBodyMspTlsca = nil
				var component *blockchainv3.ImportCaBodyMspComponent = nil
				_, err := blockchainService.NewImportCaBodyMsp(ca, tlsca, component)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewImportCaBodyMspCa successfully`, func() {
				name := "org1CA"
				model, err := blockchainService.NewImportCaBodyMspCa(name)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewImportCaBodyMspComponent successfully`, func() {
				tlsCert := "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
				model, err := blockchainService.NewImportCaBodyMspComponent(tlsCert)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewImportCaBodyMspTlsca successfully`, func() {
				name := "org1tlsCA"
				model, err := blockchainService.NewImportCaBodyMspTlsca(name)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewImportCaOptions successfully`, func() {
				// Construct an instance of the ImportCaBodyMspCa model
				importCaBodyMspCaModel := new(blockchainv3.ImportCaBodyMspCa)
				Expect(importCaBodyMspCaModel).ToNot(BeNil())
				importCaBodyMspCaModel.Name = core.StringPtr("org1CA")
				importCaBodyMspCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				Expect(importCaBodyMspCaModel.Name).To(Equal(core.StringPtr("org1CA")))
				Expect(importCaBodyMspCaModel.RootCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))

				// Construct an instance of the ImportCaBodyMspTlsca model
				importCaBodyMspTlscaModel := new(blockchainv3.ImportCaBodyMspTlsca)
				Expect(importCaBodyMspTlscaModel).ToNot(BeNil())
				importCaBodyMspTlscaModel.Name = core.StringPtr("org1tlsCA")
				importCaBodyMspTlscaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				Expect(importCaBodyMspTlscaModel.Name).To(Equal(core.StringPtr("org1tlsCA")))
				Expect(importCaBodyMspTlscaModel.RootCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))

				// Construct an instance of the ImportCaBodyMspComponent model
				importCaBodyMspComponentModel := new(blockchainv3.ImportCaBodyMspComponent)
				Expect(importCaBodyMspComponentModel).ToNot(BeNil())
				importCaBodyMspComponentModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				Expect(importCaBodyMspComponentModel.TlsCert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))

				// Construct an instance of the ImportCaBodyMsp model
				importCaBodyMspModel := new(blockchainv3.ImportCaBodyMsp)
				Expect(importCaBodyMspModel).ToNot(BeNil())
				importCaBodyMspModel.Ca = importCaBodyMspCaModel
				importCaBodyMspModel.Tlsca = importCaBodyMspTlscaModel
				importCaBodyMspModel.Component = importCaBodyMspComponentModel
				Expect(importCaBodyMspModel.Ca).To(Equal(importCaBodyMspCaModel))
				Expect(importCaBodyMspModel.Tlsca).To(Equal(importCaBodyMspTlscaModel))
				Expect(importCaBodyMspModel.Component).To(Equal(importCaBodyMspComponentModel))

				// Construct an instance of the ImportCaOptions model
				importCaOptionsDisplayName := "Sample CA"
				importCaOptionsApiURL := "https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:7054"
				var importCaOptionsMsp *blockchainv3.ImportCaBodyMsp = nil
				importCaOptionsModel := blockchainService.NewImportCaOptions(importCaOptionsDisplayName, importCaOptionsApiURL, importCaOptionsMsp)
				importCaOptionsModel.SetDisplayName("Sample CA")
				importCaOptionsModel.SetApiURL("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:7054")
				importCaOptionsModel.SetMsp(importCaBodyMspModel)
				importCaOptionsModel.SetID("component1")
				importCaOptionsModel.SetLocation("ibmcloud")
				importCaOptionsModel.SetOperationsURL("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443")
				importCaOptionsModel.SetTags([]string{"fabric-ca"})
				importCaOptionsModel.SetTlsCert("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				importCaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(importCaOptionsModel).ToNot(BeNil())
				Expect(importCaOptionsModel.DisplayName).To(Equal(core.StringPtr("Sample CA")))
				Expect(importCaOptionsModel.ApiURL).To(Equal(core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:7054")))
				Expect(importCaOptionsModel.Msp).To(Equal(importCaBodyMspModel))
				Expect(importCaOptionsModel.ID).To(Equal(core.StringPtr("component1")))
				Expect(importCaOptionsModel.Location).To(Equal(core.StringPtr("ibmcloud")))
				Expect(importCaOptionsModel.OperationsURL).To(Equal(core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443")))
				Expect(importCaOptionsModel.Tags).To(Equal([]string{"fabric-ca"}))
				Expect(importCaOptionsModel.TlsCert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(importCaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewImportMspOptions successfully`, func() {
				// Construct an instance of the ImportMspOptions model
				importMspOptionsMspID := "Org1"
				importMspOptionsDisplayName := "My Peer"
				importMspOptionsRootCerts := []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				importMspOptionsModel := blockchainService.NewImportMspOptions(importMspOptionsMspID, importMspOptionsDisplayName, importMspOptionsRootCerts)
				importMspOptionsModel.SetMspID("Org1")
				importMspOptionsModel.SetDisplayName("My Peer")
				importMspOptionsModel.SetRootCerts([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="})
				importMspOptionsModel.SetIntermediateCerts([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkRhdGEgaGVyZSBpZiB0aGlzIHdhcyByZWFsCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"})
				importMspOptionsModel.SetAdmins([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="})
				importMspOptionsModel.SetTlsRootCerts([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="})
				importMspOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(importMspOptionsModel).ToNot(BeNil())
				Expect(importMspOptionsModel.MspID).To(Equal(core.StringPtr("Org1")))
				Expect(importMspOptionsModel.DisplayName).To(Equal(core.StringPtr("My Peer")))
				Expect(importMspOptionsModel.RootCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(importMspOptionsModel.IntermediateCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkRhdGEgaGVyZSBpZiB0aGlzIHdhcyByZWFsCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"}))
				Expect(importMspOptionsModel.Admins).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(importMspOptionsModel.TlsRootCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(importMspOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewImportOrdererOptions successfully`, func() {
				// Construct an instance of the MspCryptoFieldCa model
				mspCryptoFieldCaModel := new(blockchainv3.MspCryptoFieldCa)
				Expect(mspCryptoFieldCaModel).ToNot(BeNil())
				mspCryptoFieldCaModel.Name = core.StringPtr("ca")
				mspCryptoFieldCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				Expect(mspCryptoFieldCaModel.Name).To(Equal(core.StringPtr("ca")))
				Expect(mspCryptoFieldCaModel.RootCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))

				// Construct an instance of the MspCryptoFieldTlsca model
				mspCryptoFieldTlscaModel := new(blockchainv3.MspCryptoFieldTlsca)
				Expect(mspCryptoFieldTlscaModel).ToNot(BeNil())
				mspCryptoFieldTlscaModel.Name = core.StringPtr("tlsca")
				mspCryptoFieldTlscaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				Expect(mspCryptoFieldTlscaModel.Name).To(Equal(core.StringPtr("tlsca")))
				Expect(mspCryptoFieldTlscaModel.RootCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))

				// Construct an instance of the MspCryptoFieldComponent model
				mspCryptoFieldComponentModel := new(blockchainv3.MspCryptoFieldComponent)
				Expect(mspCryptoFieldComponentModel).ToNot(BeNil())
				mspCryptoFieldComponentModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoFieldComponentModel.Ecert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoFieldComponentModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				Expect(mspCryptoFieldComponentModel.TlsCert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(mspCryptoFieldComponentModel.Ecert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(mspCryptoFieldComponentModel.AdminCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))

				// Construct an instance of the MspCryptoField model
				mspCryptoFieldModel := new(blockchainv3.MspCryptoField)
				Expect(mspCryptoFieldModel).ToNot(BeNil())
				mspCryptoFieldModel.Ca = mspCryptoFieldCaModel
				mspCryptoFieldModel.Tlsca = mspCryptoFieldTlscaModel
				mspCryptoFieldModel.Component = mspCryptoFieldComponentModel
				Expect(mspCryptoFieldModel.Ca).To(Equal(mspCryptoFieldCaModel))
				Expect(mspCryptoFieldModel.Tlsca).To(Equal(mspCryptoFieldTlscaModel))
				Expect(mspCryptoFieldModel.Component).To(Equal(mspCryptoFieldComponentModel))

				// Construct an instance of the ImportOrdererOptions model
				importOrdererOptionsClusterName := "ordering service 1"
				importOrdererOptionsDisplayName := "orderer"
				importOrdererOptionsGrpcwpURL := "https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443"
				var importOrdererOptionsMsp *blockchainv3.MspCryptoField = nil
				importOrdererOptionsMspID := "Org1"
				importOrdererOptionsModel := blockchainService.NewImportOrdererOptions(importOrdererOptionsClusterName, importOrdererOptionsDisplayName, importOrdererOptionsGrpcwpURL, importOrdererOptionsMsp, importOrdererOptionsMspID)
				importOrdererOptionsModel.SetClusterName("ordering service 1")
				importOrdererOptionsModel.SetDisplayName("orderer")
				importOrdererOptionsModel.SetGrpcwpURL("https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443")
				importOrdererOptionsModel.SetMsp(mspCryptoFieldModel)
				importOrdererOptionsModel.SetMspID("Org1")
				importOrdererOptionsModel.SetApiURL("grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				importOrdererOptionsModel.SetClusterID("mzdqhdifnl")
				importOrdererOptionsModel.SetID("component1")
				importOrdererOptionsModel.SetLocation("ibmcloud")
				importOrdererOptionsModel.SetOperationsURL("https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443")
				importOrdererOptionsModel.SetSystemChannelID("testchainid")
				importOrdererOptionsModel.SetTags([]string{"fabric-ca"})
				importOrdererOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(importOrdererOptionsModel).ToNot(BeNil())
				Expect(importOrdererOptionsModel.ClusterName).To(Equal(core.StringPtr("ordering service 1")))
				Expect(importOrdererOptionsModel.DisplayName).To(Equal(core.StringPtr("orderer")))
				Expect(importOrdererOptionsModel.GrpcwpURL).To(Equal(core.StringPtr("https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443")))
				Expect(importOrdererOptionsModel.Msp).To(Equal(mspCryptoFieldModel))
				Expect(importOrdererOptionsModel.MspID).To(Equal(core.StringPtr("Org1")))
				Expect(importOrdererOptionsModel.ApiURL).To(Equal(core.StringPtr("grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")))
				Expect(importOrdererOptionsModel.ClusterID).To(Equal(core.StringPtr("mzdqhdifnl")))
				Expect(importOrdererOptionsModel.ID).To(Equal(core.StringPtr("component1")))
				Expect(importOrdererOptionsModel.Location).To(Equal(core.StringPtr("ibmcloud")))
				Expect(importOrdererOptionsModel.OperationsURL).To(Equal(core.StringPtr("https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443")))
				Expect(importOrdererOptionsModel.SystemChannelID).To(Equal(core.StringPtr("testchainid")))
				Expect(importOrdererOptionsModel.Tags).To(Equal([]string{"fabric-ca"}))
				Expect(importOrdererOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewImportPeerOptions successfully`, func() {
				// Construct an instance of the MspCryptoFieldCa model
				mspCryptoFieldCaModel := new(blockchainv3.MspCryptoFieldCa)
				Expect(mspCryptoFieldCaModel).ToNot(BeNil())
				mspCryptoFieldCaModel.Name = core.StringPtr("ca")
				mspCryptoFieldCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				Expect(mspCryptoFieldCaModel.Name).To(Equal(core.StringPtr("ca")))
				Expect(mspCryptoFieldCaModel.RootCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))

				// Construct an instance of the MspCryptoFieldTlsca model
				mspCryptoFieldTlscaModel := new(blockchainv3.MspCryptoFieldTlsca)
				Expect(mspCryptoFieldTlscaModel).ToNot(BeNil())
				mspCryptoFieldTlscaModel.Name = core.StringPtr("tlsca")
				mspCryptoFieldTlscaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				Expect(mspCryptoFieldTlscaModel.Name).To(Equal(core.StringPtr("tlsca")))
				Expect(mspCryptoFieldTlscaModel.RootCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))

				// Construct an instance of the MspCryptoFieldComponent model
				mspCryptoFieldComponentModel := new(blockchainv3.MspCryptoFieldComponent)
				Expect(mspCryptoFieldComponentModel).ToNot(BeNil())
				mspCryptoFieldComponentModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoFieldComponentModel.Ecert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspCryptoFieldComponentModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				Expect(mspCryptoFieldComponentModel.TlsCert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(mspCryptoFieldComponentModel.Ecert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(mspCryptoFieldComponentModel.AdminCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))

				// Construct an instance of the MspCryptoField model
				mspCryptoFieldModel := new(blockchainv3.MspCryptoField)
				Expect(mspCryptoFieldModel).ToNot(BeNil())
				mspCryptoFieldModel.Ca = mspCryptoFieldCaModel
				mspCryptoFieldModel.Tlsca = mspCryptoFieldTlscaModel
				mspCryptoFieldModel.Component = mspCryptoFieldComponentModel
				Expect(mspCryptoFieldModel.Ca).To(Equal(mspCryptoFieldCaModel))
				Expect(mspCryptoFieldModel.Tlsca).To(Equal(mspCryptoFieldTlscaModel))
				Expect(mspCryptoFieldModel.Component).To(Equal(mspCryptoFieldComponentModel))

				// Construct an instance of the ImportPeerOptions model
				importPeerOptionsDisplayName := "My Peer"
				importPeerOptionsGrpcwpURL := "https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084"
				var importPeerOptionsMsp *blockchainv3.MspCryptoField = nil
				importPeerOptionsMspID := "Org1"
				importPeerOptionsModel := blockchainService.NewImportPeerOptions(importPeerOptionsDisplayName, importPeerOptionsGrpcwpURL, importPeerOptionsMsp, importPeerOptionsMspID)
				importPeerOptionsModel.SetDisplayName("My Peer")
				importPeerOptionsModel.SetGrpcwpURL("https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084")
				importPeerOptionsModel.SetMsp(mspCryptoFieldModel)
				importPeerOptionsModel.SetMspID("Org1")
				importPeerOptionsModel.SetID("component1")
				importPeerOptionsModel.SetApiURL("grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051")
				importPeerOptionsModel.SetLocation("ibmcloud")
				importPeerOptionsModel.SetOperationsURL("https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443")
				importPeerOptionsModel.SetTags([]string{"fabric-ca"})
				importPeerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(importPeerOptionsModel).ToNot(BeNil())
				Expect(importPeerOptionsModel.DisplayName).To(Equal(core.StringPtr("My Peer")))
				Expect(importPeerOptionsModel.GrpcwpURL).To(Equal(core.StringPtr("https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084")))
				Expect(importPeerOptionsModel.Msp).To(Equal(mspCryptoFieldModel))
				Expect(importPeerOptionsModel.MspID).To(Equal(core.StringPtr("Org1")))
				Expect(importPeerOptionsModel.ID).To(Equal(core.StringPtr("component1")))
				Expect(importPeerOptionsModel.ApiURL).To(Equal(core.StringPtr("grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051")))
				Expect(importPeerOptionsModel.Location).To(Equal(core.StringPtr("ibmcloud")))
				Expect(importPeerOptionsModel.OperationsURL).To(Equal(core.StringPtr("https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443")))
				Expect(importPeerOptionsModel.Tags).To(Equal([]string{"fabric-ca"}))
				Expect(importPeerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListComponentsOptions successfully`, func() {
				// Construct an instance of the ListComponentsOptions model
				listComponentsOptionsModel := blockchainService.NewListComponentsOptions()
				listComponentsOptionsModel.SetDeploymentAttrs("included")
				listComponentsOptionsModel.SetParsedCerts("included")
				listComponentsOptionsModel.SetCache("skip")
				listComponentsOptionsModel.SetCaAttrs("included")
				listComponentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listComponentsOptionsModel).ToNot(BeNil())
				Expect(listComponentsOptionsModel.DeploymentAttrs).To(Equal(core.StringPtr("included")))
				Expect(listComponentsOptionsModel.ParsedCerts).To(Equal(core.StringPtr("included")))
				Expect(listComponentsOptionsModel.Cache).To(Equal(core.StringPtr("skip")))
				Expect(listComponentsOptionsModel.CaAttrs).To(Equal(core.StringPtr("included")))
				Expect(listComponentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListNotificationsOptions successfully`, func() {
				// Construct an instance of the ListNotificationsOptions model
				listNotificationsOptionsModel := blockchainService.NewListNotificationsOptions()
				listNotificationsOptionsModel.SetLimit(float64(1))
				listNotificationsOptionsModel.SetSkip(float64(1))
				listNotificationsOptionsModel.SetComponentID("MyPeer")
				listNotificationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listNotificationsOptionsModel).ToNot(BeNil())
				Expect(listNotificationsOptionsModel.Limit).To(Equal(core.Float64Ptr(float64(1))))
				Expect(listNotificationsOptionsModel.Skip).To(Equal(core.Float64Ptr(float64(1))))
				Expect(listNotificationsOptionsModel.ComponentID).To(Equal(core.StringPtr("MyPeer")))
				Expect(listNotificationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewMetrics successfully`, func() {
				provider := "prometheus"
				model, err := blockchainService.NewMetrics(provider)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewMetricsStatsd successfully`, func() {
				network := "udp"
				address := "127.0.0.1:8125"
				writeInterval := "10s"
				prefix := "server"
				model, err := blockchainService.NewMetricsStatsd(network, address, writeInterval, prefix)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewMspCryptoCa successfully`, func() {
				rootCerts := []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				model, err := blockchainService.NewMspCryptoCa(rootCerts)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewMspCryptoComp successfully`, func() {
				ekey := "testString"
				ecert := "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
				tlsKey := "testString"
				tlsCert := "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
				model, err := blockchainService.NewMspCryptoComp(ekey, ecert, tlsKey, tlsCert)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewMspCryptoFieldComponent successfully`, func() {
				tlsCert := "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
				model, err := blockchainService.NewMspCryptoFieldComponent(tlsCert)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewMspCryptoFieldTlsca successfully`, func() {
				rootCerts := []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				model, err := blockchainService.NewMspCryptoFieldTlsca(rootCerts)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewOrdererActionOptions successfully`, func() {
				// Construct an instance of the ActionReenroll model
				actionReenrollModel := new(blockchainv3.ActionReenroll)
				Expect(actionReenrollModel).ToNot(BeNil())
				actionReenrollModel.TlsCert = core.BoolPtr(true)
				actionReenrollModel.Ecert = core.BoolPtr(true)
				Expect(actionReenrollModel.TlsCert).To(Equal(core.BoolPtr(true)))
				Expect(actionReenrollModel.Ecert).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the ActionEnroll model
				actionEnrollModel := new(blockchainv3.ActionEnroll)
				Expect(actionEnrollModel).ToNot(BeNil())
				actionEnrollModel.TlsCert = core.BoolPtr(true)
				actionEnrollModel.Ecert = core.BoolPtr(true)
				Expect(actionEnrollModel.TlsCert).To(Equal(core.BoolPtr(true)))
				Expect(actionEnrollModel.Ecert).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the OrdererActionOptions model
				id := "testString"
				ordererActionOptionsModel := blockchainService.NewOrdererActionOptions(id)
				ordererActionOptionsModel.SetID("testString")
				ordererActionOptionsModel.SetRestart(true)
				ordererActionOptionsModel.SetReenroll(actionReenrollModel)
				ordererActionOptionsModel.SetEnroll(actionEnrollModel)
				ordererActionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(ordererActionOptionsModel).ToNot(BeNil())
				Expect(ordererActionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(ordererActionOptionsModel.Restart).To(Equal(core.BoolPtr(true)))
				Expect(ordererActionOptionsModel.Reenroll).To(Equal(actionReenrollModel))
				Expect(ordererActionOptionsModel.Enroll).To(Equal(actionEnrollModel))
				Expect(ordererActionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPeerActionOptions successfully`, func() {
				// Construct an instance of the ActionReenroll model
				actionReenrollModel := new(blockchainv3.ActionReenroll)
				Expect(actionReenrollModel).ToNot(BeNil())
				actionReenrollModel.TlsCert = core.BoolPtr(true)
				actionReenrollModel.Ecert = core.BoolPtr(true)
				Expect(actionReenrollModel.TlsCert).To(Equal(core.BoolPtr(true)))
				Expect(actionReenrollModel.Ecert).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the ActionEnroll model
				actionEnrollModel := new(blockchainv3.ActionEnroll)
				Expect(actionEnrollModel).ToNot(BeNil())
				actionEnrollModel.TlsCert = core.BoolPtr(true)
				actionEnrollModel.Ecert = core.BoolPtr(true)
				Expect(actionEnrollModel.TlsCert).To(Equal(core.BoolPtr(true)))
				Expect(actionEnrollModel.Ecert).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the PeerActionOptions model
				id := "testString"
				peerActionOptionsModel := blockchainService.NewPeerActionOptions(id)
				peerActionOptionsModel.SetID("testString")
				peerActionOptionsModel.SetRestart(true)
				peerActionOptionsModel.SetReenroll(actionReenrollModel)
				peerActionOptionsModel.SetEnroll(actionEnrollModel)
				peerActionOptionsModel.SetUpgradeDbs(true)
				peerActionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(peerActionOptionsModel).ToNot(BeNil())
				Expect(peerActionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(peerActionOptionsModel.Restart).To(Equal(core.BoolPtr(true)))
				Expect(peerActionOptionsModel.Reenroll).To(Equal(actionReenrollModel))
				Expect(peerActionOptionsModel.Enroll).To(Equal(actionEnrollModel))
				Expect(peerActionOptionsModel.UpgradeDbs).To(Equal(core.BoolPtr(true)))
				Expect(peerActionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRemoveComponentOptions successfully`, func() {
				// Construct an instance of the RemoveComponentOptions model
				id := "testString"
				removeComponentOptionsModel := blockchainService.NewRemoveComponentOptions(id)
				removeComponentOptionsModel.SetID("testString")
				removeComponentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(removeComponentOptionsModel).ToNot(BeNil())
				Expect(removeComponentOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(removeComponentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRemoveComponentsByTagOptions successfully`, func() {
				// Construct an instance of the RemoveComponentsByTagOptions model
				tag := "testString"
				removeComponentsByTagOptionsModel := blockchainService.NewRemoveComponentsByTagOptions(tag)
				removeComponentsByTagOptionsModel.SetTag("testString")
				removeComponentsByTagOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(removeComponentsByTagOptionsModel).ToNot(BeNil())
				Expect(removeComponentsByTagOptionsModel.Tag).To(Equal(core.StringPtr("testString")))
				Expect(removeComponentsByTagOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewResourceObject successfully`, func() {
				var requests *blockchainv3.ResourceRequests = nil
				_, err := blockchainService.NewResourceObject(requests)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewResourceObjectCouchDb successfully`, func() {
				var requests *blockchainv3.ResourceRequests = nil
				_, err := blockchainService.NewResourceObjectCouchDb(requests)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewResourceObjectFabV1 successfully`, func() {
				var requests *blockchainv3.ResourceRequests = nil
				_, err := blockchainService.NewResourceObjectFabV1(requests)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewResourceObjectFabV2 successfully`, func() {
				var requests *blockchainv3.ResourceRequests = nil
				_, err := blockchainService.NewResourceObjectFabV2(requests)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewRestartOptions successfully`, func() {
				// Construct an instance of the RestartOptions model
				restartOptionsModel := blockchainService.NewRestartOptions()
				restartOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(restartOptionsModel).ToNot(BeNil())
				Expect(restartOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSubmitBlockOptions successfully`, func() {
				// Construct an instance of the SubmitBlockOptions model
				id := "testString"
				submitBlockOptionsModel := blockchainService.NewSubmitBlockOptions(id)
				submitBlockOptionsModel.SetID("testString")
				submitBlockOptionsModel.SetB64Block("bWFzc2l2ZSBiaW5hcnkgb2YgYSBjb25maWcgYmxvY2sgd291bGQgYmUgaGVyZSBpZiB0aGlzIHdhcyByZWFsLCBwbGVhc2UgZG9udCBzZW5kIHRoaXM=")
				submitBlockOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(submitBlockOptionsModel).ToNot(BeNil())
				Expect(submitBlockOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(submitBlockOptionsModel.B64Block).To(Equal(core.StringPtr("bWFzc2l2ZSBiaW5hcnkgb2YgYSBjb25maWcgYmxvY2sgd291bGQgYmUgaGVyZSBpZiB0aGlzIHdhcyByZWFsLCBwbGVhc2UgZG9udCBzZW5kIHRoaXM=")))
				Expect(submitBlockOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateCaBodyConfigOverride successfully`, func() {
				var ca *blockchainv3.ConfigCAUpdate = nil
				_, err := blockchainService.NewUpdateCaBodyConfigOverride(ca)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewUpdateCaBodyResources successfully`, func() {
				var ca *blockchainv3.ResourceObject = nil
				_, err := blockchainService.NewUpdateCaBodyResources(ca)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewUpdateCaOptions successfully`, func() {
				// Construct an instance of the ConfigCACors model
				configCaCorsModel := new(blockchainv3.ConfigCACors)
				Expect(configCaCorsModel).ToNot(BeNil())
				configCaCorsModel.Enabled = core.BoolPtr(true)
				configCaCorsModel.Origins = []string{"*"}
				Expect(configCaCorsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(configCaCorsModel.Origins).To(Equal([]string{"*"}))

				// Construct an instance of the ConfigCATlsClientauth model
				configCaTlsClientauthModel := new(blockchainv3.ConfigCATlsClientauth)
				Expect(configCaTlsClientauthModel).ToNot(BeNil())
				configCaTlsClientauthModel.Type = core.StringPtr("noclientcert")
				configCaTlsClientauthModel.Certfiles = []string{"testString"}
				Expect(configCaTlsClientauthModel.Type).To(Equal(core.StringPtr("noclientcert")))
				Expect(configCaTlsClientauthModel.Certfiles).To(Equal([]string{"testString"}))

				// Construct an instance of the ConfigCATls model
				configCaTlsModel := new(blockchainv3.ConfigCATls)
				Expect(configCaTlsModel).ToNot(BeNil())
				configCaTlsModel.Keyfile = core.StringPtr("testString")
				configCaTlsModel.Certfile = core.StringPtr("testString")
				configCaTlsModel.Clientauth = configCaTlsClientauthModel
				Expect(configCaTlsModel.Keyfile).To(Equal(core.StringPtr("testString")))
				Expect(configCaTlsModel.Certfile).To(Equal(core.StringPtr("testString")))
				Expect(configCaTlsModel.Clientauth).To(Equal(configCaTlsClientauthModel))

				// Construct an instance of the ConfigCACa model
				configCaCaModel := new(blockchainv3.ConfigCACa)
				Expect(configCaCaModel).ToNot(BeNil())
				configCaCaModel.Keyfile = core.StringPtr("testString")
				configCaCaModel.Certfile = core.StringPtr("testString")
				configCaCaModel.Chainfile = core.StringPtr("testString")
				Expect(configCaCaModel.Keyfile).To(Equal(core.StringPtr("testString")))
				Expect(configCaCaModel.Certfile).To(Equal(core.StringPtr("testString")))
				Expect(configCaCaModel.Chainfile).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ConfigCACrl model
				configCaCrlModel := new(blockchainv3.ConfigCACrl)
				Expect(configCaCrlModel).ToNot(BeNil())
				configCaCrlModel.Expiry = core.StringPtr("24h")
				Expect(configCaCrlModel.Expiry).To(Equal(core.StringPtr("24h")))

				// Construct an instance of the IdentityAttrs model
				identityAttrsModel := new(blockchainv3.IdentityAttrs)
				Expect(identityAttrsModel).ToNot(BeNil())
				identityAttrsModel.HfRegistrarRoles = core.StringPtr("*")
				identityAttrsModel.HfRegistrarDelegateRoles = core.StringPtr("*")
				identityAttrsModel.HfRevoker = core.BoolPtr(true)
				identityAttrsModel.HfIntermediateCA = core.BoolPtr(true)
				identityAttrsModel.HfGenCRL = core.BoolPtr(true)
				identityAttrsModel.HfRegistrarAttributes = core.StringPtr("*")
				identityAttrsModel.HfAffiliationMgr = core.BoolPtr(true)
				Expect(identityAttrsModel.HfRegistrarRoles).To(Equal(core.StringPtr("*")))
				Expect(identityAttrsModel.HfRegistrarDelegateRoles).To(Equal(core.StringPtr("*")))
				Expect(identityAttrsModel.HfRevoker).To(Equal(core.BoolPtr(true)))
				Expect(identityAttrsModel.HfIntermediateCA).To(Equal(core.BoolPtr(true)))
				Expect(identityAttrsModel.HfGenCRL).To(Equal(core.BoolPtr(true)))
				Expect(identityAttrsModel.HfRegistrarAttributes).To(Equal(core.StringPtr("*")))
				Expect(identityAttrsModel.HfAffiliationMgr).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the ConfigCARegistryIdentitiesItem model
				configCaRegistryIdentitiesItemModel := new(blockchainv3.ConfigCARegistryIdentitiesItem)
				Expect(configCaRegistryIdentitiesItemModel).ToNot(BeNil())
				configCaRegistryIdentitiesItemModel.Name = core.StringPtr("admin")
				configCaRegistryIdentitiesItemModel.Pass = core.StringPtr("password")
				configCaRegistryIdentitiesItemModel.Type = core.StringPtr("client")
				configCaRegistryIdentitiesItemModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryIdentitiesItemModel.Affiliation = core.StringPtr("testString")
				configCaRegistryIdentitiesItemModel.Attrs = identityAttrsModel
				Expect(configCaRegistryIdentitiesItemModel.Name).To(Equal(core.StringPtr("admin")))
				Expect(configCaRegistryIdentitiesItemModel.Pass).To(Equal(core.StringPtr("password")))
				Expect(configCaRegistryIdentitiesItemModel.Type).To(Equal(core.StringPtr("client")))
				Expect(configCaRegistryIdentitiesItemModel.Maxenrollments).To(Equal(core.Float64Ptr(float64(-1))))
				Expect(configCaRegistryIdentitiesItemModel.Affiliation).To(Equal(core.StringPtr("testString")))
				Expect(configCaRegistryIdentitiesItemModel.Attrs).To(Equal(identityAttrsModel))

				// Construct an instance of the ConfigCARegistry model
				configCaRegistryModel := new(blockchainv3.ConfigCARegistry)
				Expect(configCaRegistryModel).ToNot(BeNil())
				configCaRegistryModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryModel.Identities = []blockchainv3.ConfigCARegistryIdentitiesItem{*configCaRegistryIdentitiesItemModel}
				Expect(configCaRegistryModel.Maxenrollments).To(Equal(core.Float64Ptr(float64(-1))))
				Expect(configCaRegistryModel.Identities).To(Equal([]blockchainv3.ConfigCARegistryIdentitiesItem{*configCaRegistryIdentitiesItemModel}))

				// Construct an instance of the ConfigCADbTlsClient model
				configCaDbTlsClientModel := new(blockchainv3.ConfigCADbTlsClient)
				Expect(configCaDbTlsClientModel).ToNot(BeNil())
				configCaDbTlsClientModel.Certfile = core.StringPtr("testString")
				configCaDbTlsClientModel.Keyfile = core.StringPtr("testString")
				Expect(configCaDbTlsClientModel.Certfile).To(Equal(core.StringPtr("testString")))
				Expect(configCaDbTlsClientModel.Keyfile).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ConfigCADbTls model
				configCaDbTlsModel := new(blockchainv3.ConfigCADbTls)
				Expect(configCaDbTlsModel).ToNot(BeNil())
				configCaDbTlsModel.Certfiles = []string{"testString"}
				configCaDbTlsModel.Client = configCaDbTlsClientModel
				configCaDbTlsModel.Enabled = core.BoolPtr(false)
				Expect(configCaDbTlsModel.Certfiles).To(Equal([]string{"testString"}))
				Expect(configCaDbTlsModel.Client).To(Equal(configCaDbTlsClientModel))
				Expect(configCaDbTlsModel.Enabled).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the ConfigCADb model
				configCaDbModel := new(blockchainv3.ConfigCADb)
				Expect(configCaDbModel).ToNot(BeNil())
				configCaDbModel.Type = core.StringPtr("postgres")
				configCaDbModel.Datasource = core.StringPtr("host=fake.databases.appdomain.cloud port=31941 user=ibm_cloud password=password dbname=ibmclouddb sslmode=verify-full")
				configCaDbModel.Tls = configCaDbTlsModel
				Expect(configCaDbModel.Type).To(Equal(core.StringPtr("postgres")))
				Expect(configCaDbModel.Datasource).To(Equal(core.StringPtr("host=fake.databases.appdomain.cloud port=31941 user=ibm_cloud password=password dbname=ibmclouddb sslmode=verify-full")))
				Expect(configCaDbModel.Tls).To(Equal(configCaDbTlsModel))

				// Construct an instance of the ConfigCAAffiliations model
				configCaAffiliationsModel := new(blockchainv3.ConfigCAAffiliations)
				Expect(configCaAffiliationsModel).ToNot(BeNil())
				configCaAffiliationsModel.Org1 = []string{"department1"}
				configCaAffiliationsModel.Org2 = []string{"department1"}
				configCaAffiliationsModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(configCaAffiliationsModel.Org1).To(Equal([]string{"department1"}))
				Expect(configCaAffiliationsModel.Org2).To(Equal([]string{"department1"}))
				Expect(configCaAffiliationsModel.GetProperties()).ToNot(BeEmpty())
				Expect(configCaAffiliationsModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ConfigCACsrKeyrequest model
				configCaCsrKeyrequestModel := new(blockchainv3.ConfigCACsrKeyrequest)
				Expect(configCaCsrKeyrequestModel).ToNot(BeNil())
				configCaCsrKeyrequestModel.Algo = core.StringPtr("ecdsa")
				configCaCsrKeyrequestModel.Size = core.Float64Ptr(float64(256))
				Expect(configCaCsrKeyrequestModel.Algo).To(Equal(core.StringPtr("ecdsa")))
				Expect(configCaCsrKeyrequestModel.Size).To(Equal(core.Float64Ptr(float64(256))))

				// Construct an instance of the ConfigCACsrNamesItem model
				configCaCsrNamesItemModel := new(blockchainv3.ConfigCACsrNamesItem)
				Expect(configCaCsrNamesItemModel).ToNot(BeNil())
				configCaCsrNamesItemModel.C = core.StringPtr("US")
				configCaCsrNamesItemModel.ST = core.StringPtr("North Carolina")
				configCaCsrNamesItemModel.L = core.StringPtr("Raleigh")
				configCaCsrNamesItemModel.O = core.StringPtr("Hyperledger")
				configCaCsrNamesItemModel.OU = core.StringPtr("Fabric")
				Expect(configCaCsrNamesItemModel.C).To(Equal(core.StringPtr("US")))
				Expect(configCaCsrNamesItemModel.ST).To(Equal(core.StringPtr("North Carolina")))
				Expect(configCaCsrNamesItemModel.L).To(Equal(core.StringPtr("Raleigh")))
				Expect(configCaCsrNamesItemModel.O).To(Equal(core.StringPtr("Hyperledger")))
				Expect(configCaCsrNamesItemModel.OU).To(Equal(core.StringPtr("Fabric")))

				// Construct an instance of the ConfigCACsrCa model
				configCaCsrCaModel := new(blockchainv3.ConfigCACsrCa)
				Expect(configCaCsrCaModel).ToNot(BeNil())
				configCaCsrCaModel.Expiry = core.StringPtr("131400h")
				configCaCsrCaModel.Pathlength = core.Float64Ptr(float64(0))
				Expect(configCaCsrCaModel.Expiry).To(Equal(core.StringPtr("131400h")))
				Expect(configCaCsrCaModel.Pathlength).To(Equal(core.Float64Ptr(float64(0))))

				// Construct an instance of the ConfigCACsr model
				configCaCsrModel := new(blockchainv3.ConfigCACsr)
				Expect(configCaCsrModel).ToNot(BeNil())
				configCaCsrModel.Cn = core.StringPtr("ca")
				configCaCsrModel.Keyrequest = configCaCsrKeyrequestModel
				configCaCsrModel.Names = []blockchainv3.ConfigCACsrNamesItem{*configCaCsrNamesItemModel}
				configCaCsrModel.Hosts = []string{"localhost"}
				configCaCsrModel.Ca = configCaCsrCaModel
				Expect(configCaCsrModel.Cn).To(Equal(core.StringPtr("ca")))
				Expect(configCaCsrModel.Keyrequest).To(Equal(configCaCsrKeyrequestModel))
				Expect(configCaCsrModel.Names).To(Equal([]blockchainv3.ConfigCACsrNamesItem{*configCaCsrNamesItemModel}))
				Expect(configCaCsrModel.Hosts).To(Equal([]string{"localhost"}))
				Expect(configCaCsrModel.Ca).To(Equal(configCaCsrCaModel))

				// Construct an instance of the ConfigCAIdemix model
				configCaIdemixModel := new(blockchainv3.ConfigCAIdemix)
				Expect(configCaIdemixModel).ToNot(BeNil())
				configCaIdemixModel.Rhpoolsize = core.Float64Ptr(float64(100))
				configCaIdemixModel.Nonceexpiration = core.StringPtr("15s")
				configCaIdemixModel.Noncesweepinterval = core.StringPtr("15m")
				Expect(configCaIdemixModel.Rhpoolsize).To(Equal(core.Float64Ptr(float64(100))))
				Expect(configCaIdemixModel.Nonceexpiration).To(Equal(core.StringPtr("15s")))
				Expect(configCaIdemixModel.Noncesweepinterval).To(Equal(core.StringPtr("15m")))

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv3.BccspSW)
				Expect(bccspSwModel).ToNot(BeNil())
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))
				Expect(bccspSwModel.Hash).To(Equal(core.StringPtr("SHA2")))
				Expect(bccspSwModel.Security).To(Equal(core.Float64Ptr(float64(256))))

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv3.BccspPKCS11)
				Expect(bccspPkcS11Model).ToNot(BeNil())
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))
				Expect(bccspPkcS11Model.Label).To(Equal(core.StringPtr("testString")))
				Expect(bccspPkcS11Model.Pin).To(Equal(core.StringPtr("testString")))
				Expect(bccspPkcS11Model.Hash).To(Equal(core.StringPtr("SHA2")))
				Expect(bccspPkcS11Model.Security).To(Equal(core.Float64Ptr(float64(256))))

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv3.Bccsp)
				Expect(bccspModel).ToNot(BeNil())
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model
				Expect(bccspModel.Default).To(Equal(core.StringPtr("SW")))
				Expect(bccspModel.SW).To(Equal(bccspSwModel))
				Expect(bccspModel.PKCS11).To(Equal(bccspPkcS11Model))

				// Construct an instance of the ConfigCAIntermediateParentserver model
				configCaIntermediateParentserverModel := new(blockchainv3.ConfigCAIntermediateParentserver)
				Expect(configCaIntermediateParentserverModel).ToNot(BeNil())
				configCaIntermediateParentserverModel.URL = core.StringPtr("testString")
				configCaIntermediateParentserverModel.Caname = core.StringPtr("testString")
				Expect(configCaIntermediateParentserverModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(configCaIntermediateParentserverModel.Caname).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ConfigCAIntermediateEnrollment model
				configCaIntermediateEnrollmentModel := new(blockchainv3.ConfigCAIntermediateEnrollment)
				Expect(configCaIntermediateEnrollmentModel).ToNot(BeNil())
				configCaIntermediateEnrollmentModel.Hosts = core.StringPtr("localhost")
				configCaIntermediateEnrollmentModel.Profile = core.StringPtr("testString")
				configCaIntermediateEnrollmentModel.Label = core.StringPtr("testString")
				Expect(configCaIntermediateEnrollmentModel.Hosts).To(Equal(core.StringPtr("localhost")))
				Expect(configCaIntermediateEnrollmentModel.Profile).To(Equal(core.StringPtr("testString")))
				Expect(configCaIntermediateEnrollmentModel.Label).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ConfigCAIntermediateTlsClient model
				configCaIntermediateTlsClientModel := new(blockchainv3.ConfigCAIntermediateTlsClient)
				Expect(configCaIntermediateTlsClientModel).ToNot(BeNil())
				configCaIntermediateTlsClientModel.Certfile = core.StringPtr("testString")
				configCaIntermediateTlsClientModel.Keyfile = core.StringPtr("testString")
				Expect(configCaIntermediateTlsClientModel.Certfile).To(Equal(core.StringPtr("testString")))
				Expect(configCaIntermediateTlsClientModel.Keyfile).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ConfigCAIntermediateTls model
				configCaIntermediateTlsModel := new(blockchainv3.ConfigCAIntermediateTls)
				Expect(configCaIntermediateTlsModel).ToNot(BeNil())
				configCaIntermediateTlsModel.Certfiles = []string{"testString"}
				configCaIntermediateTlsModel.Client = configCaIntermediateTlsClientModel
				Expect(configCaIntermediateTlsModel.Certfiles).To(Equal([]string{"testString"}))
				Expect(configCaIntermediateTlsModel.Client).To(Equal(configCaIntermediateTlsClientModel))

				// Construct an instance of the ConfigCAIntermediate model
				configCaIntermediateModel := new(blockchainv3.ConfigCAIntermediate)
				Expect(configCaIntermediateModel).ToNot(BeNil())
				configCaIntermediateModel.Parentserver = configCaIntermediateParentserverModel
				configCaIntermediateModel.Enrollment = configCaIntermediateEnrollmentModel
				configCaIntermediateModel.Tls = configCaIntermediateTlsModel
				Expect(configCaIntermediateModel.Parentserver).To(Equal(configCaIntermediateParentserverModel))
				Expect(configCaIntermediateModel.Enrollment).To(Equal(configCaIntermediateEnrollmentModel))
				Expect(configCaIntermediateModel.Tls).To(Equal(configCaIntermediateTlsModel))

				// Construct an instance of the ConfigCACfgIdentities model
				configCaCfgIdentitiesModel := new(blockchainv3.ConfigCACfgIdentities)
				Expect(configCaCfgIdentitiesModel).ToNot(BeNil())
				configCaCfgIdentitiesModel.Passwordattempts = core.Float64Ptr(float64(10))
				configCaCfgIdentitiesModel.Allowremove = core.BoolPtr(false)
				Expect(configCaCfgIdentitiesModel.Passwordattempts).To(Equal(core.Float64Ptr(float64(10))))
				Expect(configCaCfgIdentitiesModel.Allowremove).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the ConfigCACfg model
				configCaCfgModel := new(blockchainv3.ConfigCACfg)
				Expect(configCaCfgModel).ToNot(BeNil())
				configCaCfgModel.Identities = configCaCfgIdentitiesModel
				Expect(configCaCfgModel.Identities).To(Equal(configCaCfgIdentitiesModel))

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv3.MetricsStatsd)
				Expect(metricsStatsdModel).ToNot(BeNil())
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")
				Expect(metricsStatsdModel.Network).To(Equal(core.StringPtr("udp")))
				Expect(metricsStatsdModel.Address).To(Equal(core.StringPtr("127.0.0.1:8125")))
				Expect(metricsStatsdModel.WriteInterval).To(Equal(core.StringPtr("10s")))
				Expect(metricsStatsdModel.Prefix).To(Equal(core.StringPtr("server")))

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv3.Metrics)
				Expect(metricsModel).ToNot(BeNil())
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel
				Expect(metricsModel.Provider).To(Equal(core.StringPtr("prometheus")))
				Expect(metricsModel.Statsd).To(Equal(metricsStatsdModel))

				// Construct an instance of the ConfigCAUpdate model
				configCaUpdateModel := new(blockchainv3.ConfigCAUpdate)
				Expect(configCaUpdateModel).ToNot(BeNil())
				configCaUpdateModel.Cors = configCaCorsModel
				configCaUpdateModel.Debug = core.BoolPtr(false)
				configCaUpdateModel.Crlsizelimit = core.Float64Ptr(float64(512000))
				configCaUpdateModel.Tls = configCaTlsModel
				configCaUpdateModel.Ca = configCaCaModel
				configCaUpdateModel.Crl = configCaCrlModel
				configCaUpdateModel.Registry = configCaRegistryModel
				configCaUpdateModel.Db = configCaDbModel
				configCaUpdateModel.Affiliations = configCaAffiliationsModel
				configCaUpdateModel.Csr = configCaCsrModel
				configCaUpdateModel.Idemix = configCaIdemixModel
				configCaUpdateModel.BCCSP = bccspModel
				configCaUpdateModel.Intermediate = configCaIntermediateModel
				configCaUpdateModel.Cfg = configCaCfgModel
				configCaUpdateModel.Metrics = metricsModel
				Expect(configCaUpdateModel.Cors).To(Equal(configCaCorsModel))
				Expect(configCaUpdateModel.Debug).To(Equal(core.BoolPtr(false)))
				Expect(configCaUpdateModel.Crlsizelimit).To(Equal(core.Float64Ptr(float64(512000))))
				Expect(configCaUpdateModel.Tls).To(Equal(configCaTlsModel))
				Expect(configCaUpdateModel.Ca).To(Equal(configCaCaModel))
				Expect(configCaUpdateModel.Crl).To(Equal(configCaCrlModel))
				Expect(configCaUpdateModel.Registry).To(Equal(configCaRegistryModel))
				Expect(configCaUpdateModel.Db).To(Equal(configCaDbModel))
				Expect(configCaUpdateModel.Affiliations).To(Equal(configCaAffiliationsModel))
				Expect(configCaUpdateModel.Csr).To(Equal(configCaCsrModel))
				Expect(configCaUpdateModel.Idemix).To(Equal(configCaIdemixModel))
				Expect(configCaUpdateModel.BCCSP).To(Equal(bccspModel))
				Expect(configCaUpdateModel.Intermediate).To(Equal(configCaIntermediateModel))
				Expect(configCaUpdateModel.Cfg).To(Equal(configCaCfgModel))
				Expect(configCaUpdateModel.Metrics).To(Equal(metricsModel))

				// Construct an instance of the UpdateCaBodyConfigOverride model
				updateCaBodyConfigOverrideModel := new(blockchainv3.UpdateCaBodyConfigOverride)
				Expect(updateCaBodyConfigOverrideModel).ToNot(BeNil())
				updateCaBodyConfigOverrideModel.Ca = configCaUpdateModel
				Expect(updateCaBodyConfigOverrideModel.Ca).To(Equal(configCaUpdateModel))

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv3.ResourceRequests)
				Expect(resourceRequestsModel).ToNot(BeNil())
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")
				Expect(resourceRequestsModel.Cpu).To(Equal(core.StringPtr("100m")))
				Expect(resourceRequestsModel.Memory).To(Equal(core.StringPtr("256MiB")))

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv3.ResourceLimits)
				Expect(resourceLimitsModel).ToNot(BeNil())
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")
				Expect(resourceLimitsModel.Cpu).To(Equal(core.StringPtr("100m")))
				Expect(resourceLimitsModel.Memory).To(Equal(core.StringPtr("256MiB")))

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv3.ResourceObject)
				Expect(resourceObjectModel).ToNot(BeNil())
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel
				Expect(resourceObjectModel.Requests).To(Equal(resourceRequestsModel))
				Expect(resourceObjectModel.Limits).To(Equal(resourceLimitsModel))

				// Construct an instance of the UpdateCaBodyResources model
				updateCaBodyResourcesModel := new(blockchainv3.UpdateCaBodyResources)
				Expect(updateCaBodyResourcesModel).ToNot(BeNil())
				updateCaBodyResourcesModel.Ca = resourceObjectModel
				Expect(updateCaBodyResourcesModel.Ca).To(Equal(resourceObjectModel))

				// Construct an instance of the UpdateCaOptions model
				id := "testString"
				updateCaOptionsModel := blockchainService.NewUpdateCaOptions(id)
				updateCaOptionsModel.SetID("testString")
				updateCaOptionsModel.SetConfigOverride(updateCaBodyConfigOverrideModel)
				updateCaOptionsModel.SetReplicas(float64(1))
				updateCaOptionsModel.SetResources(updateCaBodyResourcesModel)
				updateCaOptionsModel.SetVersion("1.4.6-1")
				updateCaOptionsModel.SetZone("-")
				updateCaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCaOptionsModel).ToNot(BeNil())
				Expect(updateCaOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateCaOptionsModel.ConfigOverride).To(Equal(updateCaBodyConfigOverrideModel))
				Expect(updateCaOptionsModel.Replicas).To(Equal(core.Float64Ptr(float64(1))))
				Expect(updateCaOptionsModel.Resources).To(Equal(updateCaBodyResourcesModel))
				Expect(updateCaOptionsModel.Version).To(Equal(core.StringPtr("1.4.6-1")))
				Expect(updateCaOptionsModel.Zone).To(Equal(core.StringPtr("-")))
				Expect(updateCaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateOrdererOptions successfully`, func() {
				// Construct an instance of the ConfigOrdererKeepalive model
				configOrdererKeepaliveModel := new(blockchainv3.ConfigOrdererKeepalive)
				Expect(configOrdererKeepaliveModel).ToNot(BeNil())
				configOrdererKeepaliveModel.ServerMinInterval = core.StringPtr("60s")
				configOrdererKeepaliveModel.ServerInterval = core.StringPtr("2h")
				configOrdererKeepaliveModel.ServerTimeout = core.StringPtr("20s")
				Expect(configOrdererKeepaliveModel.ServerMinInterval).To(Equal(core.StringPtr("60s")))
				Expect(configOrdererKeepaliveModel.ServerInterval).To(Equal(core.StringPtr("2h")))
				Expect(configOrdererKeepaliveModel.ServerTimeout).To(Equal(core.StringPtr("20s")))

				// Construct an instance of the ConfigOrdererAuthentication model
				configOrdererAuthenticationModel := new(blockchainv3.ConfigOrdererAuthentication)
				Expect(configOrdererAuthenticationModel).ToNot(BeNil())
				configOrdererAuthenticationModel.TimeWindow = core.StringPtr("15m")
				configOrdererAuthenticationModel.NoExpirationChecks = core.BoolPtr(false)
				Expect(configOrdererAuthenticationModel.TimeWindow).To(Equal(core.StringPtr("15m")))
				Expect(configOrdererAuthenticationModel.NoExpirationChecks).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the ConfigOrdererGeneralUpdate model
				configOrdererGeneralUpdateModel := new(blockchainv3.ConfigOrdererGeneralUpdate)
				Expect(configOrdererGeneralUpdateModel).ToNot(BeNil())
				configOrdererGeneralUpdateModel.Keepalive = configOrdererKeepaliveModel
				configOrdererGeneralUpdateModel.Authentication = configOrdererAuthenticationModel
				Expect(configOrdererGeneralUpdateModel.Keepalive).To(Equal(configOrdererKeepaliveModel))
				Expect(configOrdererGeneralUpdateModel.Authentication).To(Equal(configOrdererAuthenticationModel))

				// Construct an instance of the ConfigOrdererDebug model
				configOrdererDebugModel := new(blockchainv3.ConfigOrdererDebug)
				Expect(configOrdererDebugModel).ToNot(BeNil())
				configOrdererDebugModel.BroadcastTraceDir = core.StringPtr("testString")
				configOrdererDebugModel.DeliverTraceDir = core.StringPtr("testString")
				Expect(configOrdererDebugModel.BroadcastTraceDir).To(Equal(core.StringPtr("testString")))
				Expect(configOrdererDebugModel.DeliverTraceDir).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ConfigOrdererMetricsStatsd model
				configOrdererMetricsStatsdModel := new(blockchainv3.ConfigOrdererMetricsStatsd)
				Expect(configOrdererMetricsStatsdModel).ToNot(BeNil())
				configOrdererMetricsStatsdModel.Network = core.StringPtr("udp")
				configOrdererMetricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				configOrdererMetricsStatsdModel.WriteInterval = core.StringPtr("10s")
				configOrdererMetricsStatsdModel.Prefix = core.StringPtr("server")
				Expect(configOrdererMetricsStatsdModel.Network).To(Equal(core.StringPtr("udp")))
				Expect(configOrdererMetricsStatsdModel.Address).To(Equal(core.StringPtr("127.0.0.1:8125")))
				Expect(configOrdererMetricsStatsdModel.WriteInterval).To(Equal(core.StringPtr("10s")))
				Expect(configOrdererMetricsStatsdModel.Prefix).To(Equal(core.StringPtr("server")))

				// Construct an instance of the ConfigOrdererMetrics model
				configOrdererMetricsModel := new(blockchainv3.ConfigOrdererMetrics)
				Expect(configOrdererMetricsModel).ToNot(BeNil())
				configOrdererMetricsModel.Provider = core.StringPtr("disabled")
				configOrdererMetricsModel.Statsd = configOrdererMetricsStatsdModel
				Expect(configOrdererMetricsModel.Provider).To(Equal(core.StringPtr("disabled")))
				Expect(configOrdererMetricsModel.Statsd).To(Equal(configOrdererMetricsStatsdModel))

				// Construct an instance of the ConfigOrdererUpdate model
				configOrdererUpdateModel := new(blockchainv3.ConfigOrdererUpdate)
				Expect(configOrdererUpdateModel).ToNot(BeNil())
				configOrdererUpdateModel.General = configOrdererGeneralUpdateModel
				configOrdererUpdateModel.Debug = configOrdererDebugModel
				configOrdererUpdateModel.Metrics = configOrdererMetricsModel
				Expect(configOrdererUpdateModel.General).To(Equal(configOrdererGeneralUpdateModel))
				Expect(configOrdererUpdateModel.Debug).To(Equal(configOrdererDebugModel))
				Expect(configOrdererUpdateModel.Metrics).To(Equal(configOrdererMetricsModel))

				// Construct an instance of the CryptoEnrollmentComponent model
				cryptoEnrollmentComponentModel := new(blockchainv3.CryptoEnrollmentComponent)
				Expect(cryptoEnrollmentComponentModel).ToNot(BeNil())
				cryptoEnrollmentComponentModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				Expect(cryptoEnrollmentComponentModel.Admincerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))

				// Construct an instance of the UpdateEnrollmentCryptoFieldCa model
				updateEnrollmentCryptoFieldCaModel := new(blockchainv3.UpdateEnrollmentCryptoFieldCa)
				Expect(updateEnrollmentCryptoFieldCaModel).ToNot(BeNil())
				updateEnrollmentCryptoFieldCaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				updateEnrollmentCryptoFieldCaModel.Port = core.Float64Ptr(float64(7054))
				updateEnrollmentCryptoFieldCaModel.Name = core.StringPtr("ca")
				updateEnrollmentCryptoFieldCaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateEnrollmentCryptoFieldCaModel.EnrollID = core.StringPtr("admin")
				updateEnrollmentCryptoFieldCaModel.EnrollSecret = core.StringPtr("password")
				Expect(updateEnrollmentCryptoFieldCaModel.Host).To(Equal(core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")))
				Expect(updateEnrollmentCryptoFieldCaModel.Port).To(Equal(core.Float64Ptr(float64(7054))))
				Expect(updateEnrollmentCryptoFieldCaModel.Name).To(Equal(core.StringPtr("ca")))
				Expect(updateEnrollmentCryptoFieldCaModel.TlsCert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(updateEnrollmentCryptoFieldCaModel.EnrollID).To(Equal(core.StringPtr("admin")))
				Expect(updateEnrollmentCryptoFieldCaModel.EnrollSecret).To(Equal(core.StringPtr("password")))

				// Construct an instance of the UpdateEnrollmentCryptoFieldTlsca model
				updateEnrollmentCryptoFieldTlscaModel := new(blockchainv3.UpdateEnrollmentCryptoFieldTlsca)
				Expect(updateEnrollmentCryptoFieldTlscaModel).ToNot(BeNil())
				updateEnrollmentCryptoFieldTlscaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				updateEnrollmentCryptoFieldTlscaModel.Port = core.Float64Ptr(float64(7054))
				updateEnrollmentCryptoFieldTlscaModel.Name = core.StringPtr("tlsca")
				updateEnrollmentCryptoFieldTlscaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateEnrollmentCryptoFieldTlscaModel.EnrollID = core.StringPtr("admin")
				updateEnrollmentCryptoFieldTlscaModel.EnrollSecret = core.StringPtr("password")
				updateEnrollmentCryptoFieldTlscaModel.CsrHosts = []string{"testString"}
				Expect(updateEnrollmentCryptoFieldTlscaModel.Host).To(Equal(core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")))
				Expect(updateEnrollmentCryptoFieldTlscaModel.Port).To(Equal(core.Float64Ptr(float64(7054))))
				Expect(updateEnrollmentCryptoFieldTlscaModel.Name).To(Equal(core.StringPtr("tlsca")))
				Expect(updateEnrollmentCryptoFieldTlscaModel.TlsCert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(updateEnrollmentCryptoFieldTlscaModel.EnrollID).To(Equal(core.StringPtr("admin")))
				Expect(updateEnrollmentCryptoFieldTlscaModel.EnrollSecret).To(Equal(core.StringPtr("password")))
				Expect(updateEnrollmentCryptoFieldTlscaModel.CsrHosts).To(Equal([]string{"testString"}))

				// Construct an instance of the UpdateEnrollmentCryptoField model
				updateEnrollmentCryptoFieldModel := new(blockchainv3.UpdateEnrollmentCryptoField)
				Expect(updateEnrollmentCryptoFieldModel).ToNot(BeNil())
				updateEnrollmentCryptoFieldModel.Component = cryptoEnrollmentComponentModel
				updateEnrollmentCryptoFieldModel.Ca = updateEnrollmentCryptoFieldCaModel
				updateEnrollmentCryptoFieldModel.Tlsca = updateEnrollmentCryptoFieldTlscaModel
				Expect(updateEnrollmentCryptoFieldModel.Component).To(Equal(cryptoEnrollmentComponentModel))
				Expect(updateEnrollmentCryptoFieldModel.Ca).To(Equal(updateEnrollmentCryptoFieldCaModel))
				Expect(updateEnrollmentCryptoFieldModel.Tlsca).To(Equal(updateEnrollmentCryptoFieldTlscaModel))

				// Construct an instance of the UpdateMspCryptoFieldCa model
				updateMspCryptoFieldCaModel := new(blockchainv3.UpdateMspCryptoFieldCa)
				Expect(updateMspCryptoFieldCaModel).ToNot(BeNil())
				updateMspCryptoFieldCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateMspCryptoFieldCaModel.CaIntermediateCerts = []string{"testString"}
				Expect(updateMspCryptoFieldCaModel.RootCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(updateMspCryptoFieldCaModel.CaIntermediateCerts).To(Equal([]string{"testString"}))

				// Construct an instance of the UpdateMspCryptoFieldTlsca model
				updateMspCryptoFieldTlscaModel := new(blockchainv3.UpdateMspCryptoFieldTlsca)
				Expect(updateMspCryptoFieldTlscaModel).ToNot(BeNil())
				updateMspCryptoFieldTlscaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateMspCryptoFieldTlscaModel.CaIntermediateCerts = []string{"testString"}
				Expect(updateMspCryptoFieldTlscaModel.RootCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(updateMspCryptoFieldTlscaModel.CaIntermediateCerts).To(Equal([]string{"testString"}))

				// Construct an instance of the ClientAuth model
				clientAuthModel := new(blockchainv3.ClientAuth)
				Expect(clientAuthModel).ToNot(BeNil())
				clientAuthModel.Type = core.StringPtr("noclientcert")
				clientAuthModel.TlsCerts = []string{"testString"}
				Expect(clientAuthModel.Type).To(Equal(core.StringPtr("noclientcert")))
				Expect(clientAuthModel.TlsCerts).To(Equal([]string{"testString"}))

				// Construct an instance of the UpdateMspCryptoFieldComponent model
				updateMspCryptoFieldComponentModel := new(blockchainv3.UpdateMspCryptoFieldComponent)
				Expect(updateMspCryptoFieldComponentModel).ToNot(BeNil())
				updateMspCryptoFieldComponentModel.Ekey = core.StringPtr("testString")
				updateMspCryptoFieldComponentModel.Ecert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateMspCryptoFieldComponentModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateMspCryptoFieldComponentModel.TlsKey = core.StringPtr("testString")
				updateMspCryptoFieldComponentModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateMspCryptoFieldComponentModel.ClientAuth = clientAuthModel
				Expect(updateMspCryptoFieldComponentModel.Ekey).To(Equal(core.StringPtr("testString")))
				Expect(updateMspCryptoFieldComponentModel.Ecert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(updateMspCryptoFieldComponentModel.AdminCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(updateMspCryptoFieldComponentModel.TlsKey).To(Equal(core.StringPtr("testString")))
				Expect(updateMspCryptoFieldComponentModel.TlsCert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(updateMspCryptoFieldComponentModel.ClientAuth).To(Equal(clientAuthModel))

				// Construct an instance of the UpdateMspCryptoField model
				updateMspCryptoFieldModel := new(blockchainv3.UpdateMspCryptoField)
				Expect(updateMspCryptoFieldModel).ToNot(BeNil())
				updateMspCryptoFieldModel.Ca = updateMspCryptoFieldCaModel
				updateMspCryptoFieldModel.Tlsca = updateMspCryptoFieldTlscaModel
				updateMspCryptoFieldModel.Component = updateMspCryptoFieldComponentModel
				Expect(updateMspCryptoFieldModel.Ca).To(Equal(updateMspCryptoFieldCaModel))
				Expect(updateMspCryptoFieldModel.Tlsca).To(Equal(updateMspCryptoFieldTlscaModel))
				Expect(updateMspCryptoFieldModel.Component).To(Equal(updateMspCryptoFieldComponentModel))

				// Construct an instance of the UpdateOrdererBodyCrypto model
				updateOrdererBodyCryptoModel := new(blockchainv3.UpdateOrdererBodyCrypto)
				Expect(updateOrdererBodyCryptoModel).ToNot(BeNil())
				updateOrdererBodyCryptoModel.Enrollment = updateEnrollmentCryptoFieldModel
				updateOrdererBodyCryptoModel.Msp = updateMspCryptoFieldModel
				Expect(updateOrdererBodyCryptoModel.Enrollment).To(Equal(updateEnrollmentCryptoFieldModel))
				Expect(updateOrdererBodyCryptoModel.Msp).To(Equal(updateMspCryptoFieldModel))

				// Construct an instance of the NodeOu model
				nodeOuModel := new(blockchainv3.NodeOu)
				Expect(nodeOuModel).ToNot(BeNil())
				nodeOuModel.Enabled = core.BoolPtr(true)
				Expect(nodeOuModel.Enabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv3.ResourceRequests)
				Expect(resourceRequestsModel).ToNot(BeNil())
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")
				Expect(resourceRequestsModel.Cpu).To(Equal(core.StringPtr("100m")))
				Expect(resourceRequestsModel.Memory).To(Equal(core.StringPtr("256MiB")))

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv3.ResourceLimits)
				Expect(resourceLimitsModel).ToNot(BeNil())
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")
				Expect(resourceLimitsModel.Cpu).To(Equal(core.StringPtr("100m")))
				Expect(resourceLimitsModel.Memory).To(Equal(core.StringPtr("256MiB")))

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv3.ResourceObject)
				Expect(resourceObjectModel).ToNot(BeNil())
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel
				Expect(resourceObjectModel.Requests).To(Equal(resourceRequestsModel))
				Expect(resourceObjectModel.Limits).To(Equal(resourceLimitsModel))

				// Construct an instance of the UpdateOrdererBodyResources model
				updateOrdererBodyResourcesModel := new(blockchainv3.UpdateOrdererBodyResources)
				Expect(updateOrdererBodyResourcesModel).ToNot(BeNil())
				updateOrdererBodyResourcesModel.Orderer = resourceObjectModel
				updateOrdererBodyResourcesModel.Proxy = resourceObjectModel
				Expect(updateOrdererBodyResourcesModel.Orderer).To(Equal(resourceObjectModel))
				Expect(updateOrdererBodyResourcesModel.Proxy).To(Equal(resourceObjectModel))

				// Construct an instance of the UpdateOrdererOptions model
				id := "testString"
				updateOrdererOptionsModel := blockchainService.NewUpdateOrdererOptions(id)
				updateOrdererOptionsModel.SetID("testString")
				updateOrdererOptionsModel.SetAdminCerts([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="})
				updateOrdererOptionsModel.SetConfigOverride(configOrdererUpdateModel)
				updateOrdererOptionsModel.SetCrypto(updateOrdererBodyCryptoModel)
				updateOrdererOptionsModel.SetNodeOu(nodeOuModel)
				updateOrdererOptionsModel.SetReplicas(float64(1))
				updateOrdererOptionsModel.SetResources(updateOrdererBodyResourcesModel)
				updateOrdererOptionsModel.SetVersion("1.4.6-1")
				updateOrdererOptionsModel.SetZone("-")
				updateOrdererOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateOrdererOptionsModel).ToNot(BeNil())
				Expect(updateOrdererOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateOrdererOptionsModel.AdminCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(updateOrdererOptionsModel.ConfigOverride).To(Equal(configOrdererUpdateModel))
				Expect(updateOrdererOptionsModel.Crypto).To(Equal(updateOrdererBodyCryptoModel))
				Expect(updateOrdererOptionsModel.NodeOu).To(Equal(nodeOuModel))
				Expect(updateOrdererOptionsModel.Replicas).To(Equal(core.Float64Ptr(float64(1))))
				Expect(updateOrdererOptionsModel.Resources).To(Equal(updateOrdererBodyResourcesModel))
				Expect(updateOrdererOptionsModel.Version).To(Equal(core.StringPtr("1.4.6-1")))
				Expect(updateOrdererOptionsModel.Zone).To(Equal(core.StringPtr("-")))
				Expect(updateOrdererOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdatePeerOptions successfully`, func() {
				// Construct an instance of the ConfigPeerKeepaliveClient model
				configPeerKeepaliveClientModel := new(blockchainv3.ConfigPeerKeepaliveClient)
				Expect(configPeerKeepaliveClientModel).ToNot(BeNil())
				configPeerKeepaliveClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveClientModel.Timeout = core.StringPtr("20s")
				Expect(configPeerKeepaliveClientModel.Interval).To(Equal(core.StringPtr("60s")))
				Expect(configPeerKeepaliveClientModel.Timeout).To(Equal(core.StringPtr("20s")))

				// Construct an instance of the ConfigPeerKeepaliveDeliveryClient model
				configPeerKeepaliveDeliveryClientModel := new(blockchainv3.ConfigPeerKeepaliveDeliveryClient)
				Expect(configPeerKeepaliveDeliveryClientModel).ToNot(BeNil())
				configPeerKeepaliveDeliveryClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveDeliveryClientModel.Timeout = core.StringPtr("20s")
				Expect(configPeerKeepaliveDeliveryClientModel.Interval).To(Equal(core.StringPtr("60s")))
				Expect(configPeerKeepaliveDeliveryClientModel.Timeout).To(Equal(core.StringPtr("20s")))

				// Construct an instance of the ConfigPeerKeepalive model
				configPeerKeepaliveModel := new(blockchainv3.ConfigPeerKeepalive)
				Expect(configPeerKeepaliveModel).ToNot(BeNil())
				configPeerKeepaliveModel.MinInterval = core.StringPtr("60s")
				configPeerKeepaliveModel.Client = configPeerKeepaliveClientModel
				configPeerKeepaliveModel.DeliveryClient = configPeerKeepaliveDeliveryClientModel
				Expect(configPeerKeepaliveModel.MinInterval).To(Equal(core.StringPtr("60s")))
				Expect(configPeerKeepaliveModel.Client).To(Equal(configPeerKeepaliveClientModel))
				Expect(configPeerKeepaliveModel.DeliveryClient).To(Equal(configPeerKeepaliveDeliveryClientModel))

				// Construct an instance of the ConfigPeerGossipElection model
				configPeerGossipElectionModel := new(blockchainv3.ConfigPeerGossipElection)
				Expect(configPeerGossipElectionModel).ToNot(BeNil())
				configPeerGossipElectionModel.StartupGracePeriod = core.StringPtr("15s")
				configPeerGossipElectionModel.MembershipSampleInterval = core.StringPtr("1s")
				configPeerGossipElectionModel.LeaderAliveThreshold = core.StringPtr("10s")
				configPeerGossipElectionModel.LeaderElectionDuration = core.StringPtr("5s")
				Expect(configPeerGossipElectionModel.StartupGracePeriod).To(Equal(core.StringPtr("15s")))
				Expect(configPeerGossipElectionModel.MembershipSampleInterval).To(Equal(core.StringPtr("1s")))
				Expect(configPeerGossipElectionModel.LeaderAliveThreshold).To(Equal(core.StringPtr("10s")))
				Expect(configPeerGossipElectionModel.LeaderElectionDuration).To(Equal(core.StringPtr("5s")))

				// Construct an instance of the ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy model
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel := new(blockchainv3.ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy)
				Expect(configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel).ToNot(BeNil())
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.RequiredPeerCount = core.Float64Ptr(float64(0))
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.MaxPeerCount = core.Float64Ptr(float64(1))
				Expect(configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.RequiredPeerCount).To(Equal(core.Float64Ptr(float64(0))))
				Expect(configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.MaxPeerCount).To(Equal(core.Float64Ptr(float64(1))))

				// Construct an instance of the ConfigPeerGossipPvtData model
				configPeerGossipPvtDataModel := new(blockchainv3.ConfigPeerGossipPvtData)
				Expect(configPeerGossipPvtDataModel).ToNot(BeNil())
				configPeerGossipPvtDataModel.PullRetryThreshold = core.StringPtr("60s")
				configPeerGossipPvtDataModel.TransientstoreMaxBlockRetention = core.Float64Ptr(float64(1000))
				configPeerGossipPvtDataModel.PushAckTimeout = core.StringPtr("3s")
				configPeerGossipPvtDataModel.BtlPullMargin = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileBatchSize = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileSleepInterval = core.StringPtr("1m")
				configPeerGossipPvtDataModel.ReconciliationEnabled = core.BoolPtr(true)
				configPeerGossipPvtDataModel.SkipPullingInvalidTransactionsDuringCommit = core.BoolPtr(false)
				configPeerGossipPvtDataModel.ImplicitCollectionDisseminationPolicy = configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel
				Expect(configPeerGossipPvtDataModel.PullRetryThreshold).To(Equal(core.StringPtr("60s")))
				Expect(configPeerGossipPvtDataModel.TransientstoreMaxBlockRetention).To(Equal(core.Float64Ptr(float64(1000))))
				Expect(configPeerGossipPvtDataModel.PushAckTimeout).To(Equal(core.StringPtr("3s")))
				Expect(configPeerGossipPvtDataModel.BtlPullMargin).To(Equal(core.Float64Ptr(float64(10))))
				Expect(configPeerGossipPvtDataModel.ReconcileBatchSize).To(Equal(core.Float64Ptr(float64(10))))
				Expect(configPeerGossipPvtDataModel.ReconcileSleepInterval).To(Equal(core.StringPtr("1m")))
				Expect(configPeerGossipPvtDataModel.ReconciliationEnabled).To(Equal(core.BoolPtr(true)))
				Expect(configPeerGossipPvtDataModel.SkipPullingInvalidTransactionsDuringCommit).To(Equal(core.BoolPtr(false)))
				Expect(configPeerGossipPvtDataModel.ImplicitCollectionDisseminationPolicy).To(Equal(configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel))

				// Construct an instance of the ConfigPeerGossipState model
				configPeerGossipStateModel := new(blockchainv3.ConfigPeerGossipState)
				Expect(configPeerGossipStateModel).ToNot(BeNil())
				configPeerGossipStateModel.Enabled = core.BoolPtr(true)
				configPeerGossipStateModel.CheckInterval = core.StringPtr("10s")
				configPeerGossipStateModel.ResponseTimeout = core.StringPtr("3s")
				configPeerGossipStateModel.BatchSize = core.Float64Ptr(float64(10))
				configPeerGossipStateModel.BlockBufferSize = core.Float64Ptr(float64(100))
				configPeerGossipStateModel.MaxRetries = core.Float64Ptr(float64(3))
				Expect(configPeerGossipStateModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(configPeerGossipStateModel.CheckInterval).To(Equal(core.StringPtr("10s")))
				Expect(configPeerGossipStateModel.ResponseTimeout).To(Equal(core.StringPtr("3s")))
				Expect(configPeerGossipStateModel.BatchSize).To(Equal(core.Float64Ptr(float64(10))))
				Expect(configPeerGossipStateModel.BlockBufferSize).To(Equal(core.Float64Ptr(float64(100))))
				Expect(configPeerGossipStateModel.MaxRetries).To(Equal(core.Float64Ptr(float64(3))))

				// Construct an instance of the ConfigPeerGossip model
				configPeerGossipModel := new(blockchainv3.ConfigPeerGossip)
				Expect(configPeerGossipModel).ToNot(BeNil())
				configPeerGossipModel.UseLeaderElection = core.BoolPtr(true)
				configPeerGossipModel.OrgLeader = core.BoolPtr(false)
				configPeerGossipModel.MembershipTrackerInterval = core.StringPtr("5s")
				configPeerGossipModel.MaxBlockCountToStore = core.Float64Ptr(float64(100))
				configPeerGossipModel.MaxPropagationBurstLatency = core.StringPtr("10ms")
				configPeerGossipModel.MaxPropagationBurstSize = core.Float64Ptr(float64(10))
				configPeerGossipModel.PropagateIterations = core.Float64Ptr(float64(3))
				configPeerGossipModel.PullInterval = core.StringPtr("4s")
				configPeerGossipModel.PullPeerNum = core.Float64Ptr(float64(3))
				configPeerGossipModel.RequestStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.PublishStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.StateInfoRetentionInterval = core.StringPtr("0s")
				configPeerGossipModel.PublishCertPeriod = core.StringPtr("10s")
				configPeerGossipModel.SkipBlockVerification = core.BoolPtr(false)
				configPeerGossipModel.DialTimeout = core.StringPtr("3s")
				configPeerGossipModel.ConnTimeout = core.StringPtr("2s")
				configPeerGossipModel.RecvBuffSize = core.Float64Ptr(float64(20))
				configPeerGossipModel.SendBuffSize = core.Float64Ptr(float64(200))
				configPeerGossipModel.DigestWaitTime = core.StringPtr("1s")
				configPeerGossipModel.RequestWaitTime = core.StringPtr("1500ms")
				configPeerGossipModel.ResponseWaitTime = core.StringPtr("2s")
				configPeerGossipModel.AliveTimeInterval = core.StringPtr("5s")
				configPeerGossipModel.AliveExpirationTimeout = core.StringPtr("25s")
				configPeerGossipModel.ReconnectInterval = core.StringPtr("25s")
				configPeerGossipModel.Election = configPeerGossipElectionModel
				configPeerGossipModel.PvtData = configPeerGossipPvtDataModel
				configPeerGossipModel.State = configPeerGossipStateModel
				Expect(configPeerGossipModel.UseLeaderElection).To(Equal(core.BoolPtr(true)))
				Expect(configPeerGossipModel.OrgLeader).To(Equal(core.BoolPtr(false)))
				Expect(configPeerGossipModel.MembershipTrackerInterval).To(Equal(core.StringPtr("5s")))
				Expect(configPeerGossipModel.MaxBlockCountToStore).To(Equal(core.Float64Ptr(float64(100))))
				Expect(configPeerGossipModel.MaxPropagationBurstLatency).To(Equal(core.StringPtr("10ms")))
				Expect(configPeerGossipModel.MaxPropagationBurstSize).To(Equal(core.Float64Ptr(float64(10))))
				Expect(configPeerGossipModel.PropagateIterations).To(Equal(core.Float64Ptr(float64(3))))
				Expect(configPeerGossipModel.PullInterval).To(Equal(core.StringPtr("4s")))
				Expect(configPeerGossipModel.PullPeerNum).To(Equal(core.Float64Ptr(float64(3))))
				Expect(configPeerGossipModel.RequestStateInfoInterval).To(Equal(core.StringPtr("4s")))
				Expect(configPeerGossipModel.PublishStateInfoInterval).To(Equal(core.StringPtr("4s")))
				Expect(configPeerGossipModel.StateInfoRetentionInterval).To(Equal(core.StringPtr("0s")))
				Expect(configPeerGossipModel.PublishCertPeriod).To(Equal(core.StringPtr("10s")))
				Expect(configPeerGossipModel.SkipBlockVerification).To(Equal(core.BoolPtr(false)))
				Expect(configPeerGossipModel.DialTimeout).To(Equal(core.StringPtr("3s")))
				Expect(configPeerGossipModel.ConnTimeout).To(Equal(core.StringPtr("2s")))
				Expect(configPeerGossipModel.RecvBuffSize).To(Equal(core.Float64Ptr(float64(20))))
				Expect(configPeerGossipModel.SendBuffSize).To(Equal(core.Float64Ptr(float64(200))))
				Expect(configPeerGossipModel.DigestWaitTime).To(Equal(core.StringPtr("1s")))
				Expect(configPeerGossipModel.RequestWaitTime).To(Equal(core.StringPtr("1500ms")))
				Expect(configPeerGossipModel.ResponseWaitTime).To(Equal(core.StringPtr("2s")))
				Expect(configPeerGossipModel.AliveTimeInterval).To(Equal(core.StringPtr("5s")))
				Expect(configPeerGossipModel.AliveExpirationTimeout).To(Equal(core.StringPtr("25s")))
				Expect(configPeerGossipModel.ReconnectInterval).To(Equal(core.StringPtr("25s")))
				Expect(configPeerGossipModel.Election).To(Equal(configPeerGossipElectionModel))
				Expect(configPeerGossipModel.PvtData).To(Equal(configPeerGossipPvtDataModel))
				Expect(configPeerGossipModel.State).To(Equal(configPeerGossipStateModel))

				// Construct an instance of the ConfigPeerAuthentication model
				configPeerAuthenticationModel := new(blockchainv3.ConfigPeerAuthentication)
				Expect(configPeerAuthenticationModel).ToNot(BeNil())
				configPeerAuthenticationModel.Timewindow = core.StringPtr("15m")
				Expect(configPeerAuthenticationModel.Timewindow).To(Equal(core.StringPtr("15m")))

				// Construct an instance of the ConfigPeerClient model
				configPeerClientModel := new(blockchainv3.ConfigPeerClient)
				Expect(configPeerClientModel).ToNot(BeNil())
				configPeerClientModel.ConnTimeout = core.StringPtr("2s")
				Expect(configPeerClientModel.ConnTimeout).To(Equal(core.StringPtr("2s")))

				// Construct an instance of the ConfigPeerDeliveryclientAddressOverridesItem model
				configPeerDeliveryclientAddressOverridesItemModel := new(blockchainv3.ConfigPeerDeliveryclientAddressOverridesItem)
				Expect(configPeerDeliveryclientAddressOverridesItemModel).ToNot(BeNil())
				configPeerDeliveryclientAddressOverridesItemModel.From = core.StringPtr("n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.To = core.StringPtr("n3a3ec3-myorderer2.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.CaCertsFile = core.StringPtr("my-data/cert.pem")
				Expect(configPeerDeliveryclientAddressOverridesItemModel.From).To(Equal(core.StringPtr("n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")))
				Expect(configPeerDeliveryclientAddressOverridesItemModel.To).To(Equal(core.StringPtr("n3a3ec3-myorderer2.ibp.us-south.containers.appdomain.cloud:7050")))
				Expect(configPeerDeliveryclientAddressOverridesItemModel.CaCertsFile).To(Equal(core.StringPtr("my-data/cert.pem")))

				// Construct an instance of the ConfigPeerDeliveryclient model
				configPeerDeliveryclientModel := new(blockchainv3.ConfigPeerDeliveryclient)
				Expect(configPeerDeliveryclientModel).ToNot(BeNil())
				configPeerDeliveryclientModel.ReconnectTotalTimeThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.ConnTimeout = core.StringPtr("2s")
				configPeerDeliveryclientModel.ReConnectBackoffThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.AddressOverrides = []blockchainv3.ConfigPeerDeliveryclientAddressOverridesItem{*configPeerDeliveryclientAddressOverridesItemModel}
				Expect(configPeerDeliveryclientModel.ReconnectTotalTimeThreshold).To(Equal(core.StringPtr("60m")))
				Expect(configPeerDeliveryclientModel.ConnTimeout).To(Equal(core.StringPtr("2s")))
				Expect(configPeerDeliveryclientModel.ReConnectBackoffThreshold).To(Equal(core.StringPtr("60m")))
				Expect(configPeerDeliveryclientModel.AddressOverrides).To(Equal([]blockchainv3.ConfigPeerDeliveryclientAddressOverridesItem{*configPeerDeliveryclientAddressOverridesItemModel}))

				// Construct an instance of the ConfigPeerAdminService model
				configPeerAdminServiceModel := new(blockchainv3.ConfigPeerAdminService)
				Expect(configPeerAdminServiceModel).ToNot(BeNil())
				configPeerAdminServiceModel.ListenAddress = core.StringPtr("0.0.0.0:7051")
				Expect(configPeerAdminServiceModel.ListenAddress).To(Equal(core.StringPtr("0.0.0.0:7051")))

				// Construct an instance of the ConfigPeerDiscovery model
				configPeerDiscoveryModel := new(blockchainv3.ConfigPeerDiscovery)
				Expect(configPeerDiscoveryModel).ToNot(BeNil())
				configPeerDiscoveryModel.Enabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheEnabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheMaxSize = core.Float64Ptr(float64(1000))
				configPeerDiscoveryModel.AuthCachePurgeRetentionRatio = core.Float64Ptr(float64(0.75))
				configPeerDiscoveryModel.OrgMembersAllowedAccess = core.BoolPtr(false)
				Expect(configPeerDiscoveryModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(configPeerDiscoveryModel.AuthCacheEnabled).To(Equal(core.BoolPtr(true)))
				Expect(configPeerDiscoveryModel.AuthCacheMaxSize).To(Equal(core.Float64Ptr(float64(1000))))
				Expect(configPeerDiscoveryModel.AuthCachePurgeRetentionRatio).To(Equal(core.Float64Ptr(float64(0.75))))
				Expect(configPeerDiscoveryModel.OrgMembersAllowedAccess).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the ConfigPeerLimitsConcurrency model
				configPeerLimitsConcurrencyModel := new(blockchainv3.ConfigPeerLimitsConcurrency)
				Expect(configPeerLimitsConcurrencyModel).ToNot(BeNil())
				configPeerLimitsConcurrencyModel.EndorserService = core.Float64Ptr(float64(2500))
				configPeerLimitsConcurrencyModel.DeliverService = core.Float64Ptr(float64(2500))
				Expect(configPeerLimitsConcurrencyModel.EndorserService).To(Equal(core.Float64Ptr(float64(2500))))
				Expect(configPeerLimitsConcurrencyModel.DeliverService).To(Equal(core.Float64Ptr(float64(2500))))

				// Construct an instance of the ConfigPeerLimits model
				configPeerLimitsModel := new(blockchainv3.ConfigPeerLimits)
				Expect(configPeerLimitsModel).ToNot(BeNil())
				configPeerLimitsModel.Concurrency = configPeerLimitsConcurrencyModel
				Expect(configPeerLimitsModel.Concurrency).To(Equal(configPeerLimitsConcurrencyModel))

				// Construct an instance of the ConfigPeerGateway model
				configPeerGatewayModel := new(blockchainv3.ConfigPeerGateway)
				Expect(configPeerGatewayModel).ToNot(BeNil())
				configPeerGatewayModel.Enabled = core.BoolPtr(true)
				Expect(configPeerGatewayModel.Enabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the ConfigPeerUpdatePeer model
				configPeerUpdatePeerModel := new(blockchainv3.ConfigPeerUpdatePeer)
				Expect(configPeerUpdatePeerModel).ToNot(BeNil())
				configPeerUpdatePeerModel.ID = core.StringPtr("john-doe")
				configPeerUpdatePeerModel.NetworkID = core.StringPtr("dev")
				configPeerUpdatePeerModel.Keepalive = configPeerKeepaliveModel
				configPeerUpdatePeerModel.Gossip = configPeerGossipModel
				configPeerUpdatePeerModel.Authentication = configPeerAuthenticationModel
				configPeerUpdatePeerModel.Client = configPeerClientModel
				configPeerUpdatePeerModel.Deliveryclient = configPeerDeliveryclientModel
				configPeerUpdatePeerModel.AdminService = configPeerAdminServiceModel
				configPeerUpdatePeerModel.ValidatorPoolSize = core.Float64Ptr(float64(8))
				configPeerUpdatePeerModel.Discovery = configPeerDiscoveryModel
				configPeerUpdatePeerModel.Limits = configPeerLimitsModel
				configPeerUpdatePeerModel.Gateway = configPeerGatewayModel
				Expect(configPeerUpdatePeerModel.ID).To(Equal(core.StringPtr("john-doe")))
				Expect(configPeerUpdatePeerModel.NetworkID).To(Equal(core.StringPtr("dev")))
				Expect(configPeerUpdatePeerModel.Keepalive).To(Equal(configPeerKeepaliveModel))
				Expect(configPeerUpdatePeerModel.Gossip).To(Equal(configPeerGossipModel))
				Expect(configPeerUpdatePeerModel.Authentication).To(Equal(configPeerAuthenticationModel))
				Expect(configPeerUpdatePeerModel.Client).To(Equal(configPeerClientModel))
				Expect(configPeerUpdatePeerModel.Deliveryclient).To(Equal(configPeerDeliveryclientModel))
				Expect(configPeerUpdatePeerModel.AdminService).To(Equal(configPeerAdminServiceModel))
				Expect(configPeerUpdatePeerModel.ValidatorPoolSize).To(Equal(core.Float64Ptr(float64(8))))
				Expect(configPeerUpdatePeerModel.Discovery).To(Equal(configPeerDiscoveryModel))
				Expect(configPeerUpdatePeerModel.Limits).To(Equal(configPeerLimitsModel))
				Expect(configPeerUpdatePeerModel.Gateway).To(Equal(configPeerGatewayModel))

				// Construct an instance of the ConfigPeerChaincodeGolang model
				configPeerChaincodeGolangModel := new(blockchainv3.ConfigPeerChaincodeGolang)
				Expect(configPeerChaincodeGolangModel).ToNot(BeNil())
				configPeerChaincodeGolangModel.DynamicLink = core.BoolPtr(false)
				Expect(configPeerChaincodeGolangModel.DynamicLink).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the ConfigPeerChaincodeExternalBuildersItem model
				configPeerChaincodeExternalBuildersItemModel := new(blockchainv3.ConfigPeerChaincodeExternalBuildersItem)
				Expect(configPeerChaincodeExternalBuildersItemModel).ToNot(BeNil())
				configPeerChaincodeExternalBuildersItemModel.Path = core.StringPtr("/path/to/directory")
				configPeerChaincodeExternalBuildersItemModel.Name = core.StringPtr("descriptive-build-name")
				configPeerChaincodeExternalBuildersItemModel.EnvironmentWhitelist = []string{"GOPROXY"}
				Expect(configPeerChaincodeExternalBuildersItemModel.Path).To(Equal(core.StringPtr("/path/to/directory")))
				Expect(configPeerChaincodeExternalBuildersItemModel.Name).To(Equal(core.StringPtr("descriptive-build-name")))
				Expect(configPeerChaincodeExternalBuildersItemModel.EnvironmentWhitelist).To(Equal([]string{"GOPROXY"}))

				// Construct an instance of the ConfigPeerChaincodeSystem model
				configPeerChaincodeSystemModel := new(blockchainv3.ConfigPeerChaincodeSystem)
				Expect(configPeerChaincodeSystemModel).ToNot(BeNil())
				configPeerChaincodeSystemModel.Cscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Lscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Escc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Vscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Qscc = core.BoolPtr(true)
				Expect(configPeerChaincodeSystemModel.Cscc).To(Equal(core.BoolPtr(true)))
				Expect(configPeerChaincodeSystemModel.Lscc).To(Equal(core.BoolPtr(true)))
				Expect(configPeerChaincodeSystemModel.Escc).To(Equal(core.BoolPtr(true)))
				Expect(configPeerChaincodeSystemModel.Vscc).To(Equal(core.BoolPtr(true)))
				Expect(configPeerChaincodeSystemModel.Qscc).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the ConfigPeerChaincodeLogging model
				configPeerChaincodeLoggingModel := new(blockchainv3.ConfigPeerChaincodeLogging)
				Expect(configPeerChaincodeLoggingModel).ToNot(BeNil())
				configPeerChaincodeLoggingModel.Level = core.StringPtr("info")
				configPeerChaincodeLoggingModel.Shim = core.StringPtr("warning")
				configPeerChaincodeLoggingModel.Format = core.StringPtr("%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}")
				Expect(configPeerChaincodeLoggingModel.Level).To(Equal(core.StringPtr("info")))
				Expect(configPeerChaincodeLoggingModel.Shim).To(Equal(core.StringPtr("warning")))
				Expect(configPeerChaincodeLoggingModel.Format).To(Equal(core.StringPtr("%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}")))

				// Construct an instance of the ConfigPeerChaincode model
				configPeerChaincodeModel := new(blockchainv3.ConfigPeerChaincode)
				Expect(configPeerChaincodeModel).ToNot(BeNil())
				configPeerChaincodeModel.Golang = configPeerChaincodeGolangModel
				configPeerChaincodeModel.ExternalBuilders = []blockchainv3.ConfigPeerChaincodeExternalBuildersItem{*configPeerChaincodeExternalBuildersItemModel}
				configPeerChaincodeModel.InstallTimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Startuptimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Executetimeout = core.StringPtr("30s")
				configPeerChaincodeModel.System = configPeerChaincodeSystemModel
				configPeerChaincodeModel.Logging = configPeerChaincodeLoggingModel
				Expect(configPeerChaincodeModel.Golang).To(Equal(configPeerChaincodeGolangModel))
				Expect(configPeerChaincodeModel.ExternalBuilders).To(Equal([]blockchainv3.ConfigPeerChaincodeExternalBuildersItem{*configPeerChaincodeExternalBuildersItemModel}))
				Expect(configPeerChaincodeModel.InstallTimeout).To(Equal(core.StringPtr("300s")))
				Expect(configPeerChaincodeModel.Startuptimeout).To(Equal(core.StringPtr("300s")))
				Expect(configPeerChaincodeModel.Executetimeout).To(Equal(core.StringPtr("30s")))
				Expect(configPeerChaincodeModel.System).To(Equal(configPeerChaincodeSystemModel))
				Expect(configPeerChaincodeModel.Logging).To(Equal(configPeerChaincodeLoggingModel))

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv3.MetricsStatsd)
				Expect(metricsStatsdModel).ToNot(BeNil())
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")
				Expect(metricsStatsdModel.Network).To(Equal(core.StringPtr("udp")))
				Expect(metricsStatsdModel.Address).To(Equal(core.StringPtr("127.0.0.1:8125")))
				Expect(metricsStatsdModel.WriteInterval).To(Equal(core.StringPtr("10s")))
				Expect(metricsStatsdModel.Prefix).To(Equal(core.StringPtr("server")))

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv3.Metrics)
				Expect(metricsModel).ToNot(BeNil())
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel
				Expect(metricsModel.Provider).To(Equal(core.StringPtr("prometheus")))
				Expect(metricsModel.Statsd).To(Equal(metricsStatsdModel))

				// Construct an instance of the ConfigPeerUpdate model
				configPeerUpdateModel := new(blockchainv3.ConfigPeerUpdate)
				Expect(configPeerUpdateModel).ToNot(BeNil())
				configPeerUpdateModel.Peer = configPeerUpdatePeerModel
				configPeerUpdateModel.Chaincode = configPeerChaincodeModel
				configPeerUpdateModel.Metrics = metricsModel
				Expect(configPeerUpdateModel.Peer).To(Equal(configPeerUpdatePeerModel))
				Expect(configPeerUpdateModel.Chaincode).To(Equal(configPeerChaincodeModel))
				Expect(configPeerUpdateModel.Metrics).To(Equal(metricsModel))

				// Construct an instance of the CryptoEnrollmentComponent model
				cryptoEnrollmentComponentModel := new(blockchainv3.CryptoEnrollmentComponent)
				Expect(cryptoEnrollmentComponentModel).ToNot(BeNil())
				cryptoEnrollmentComponentModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				Expect(cryptoEnrollmentComponentModel.Admincerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))

				// Construct an instance of the UpdateEnrollmentCryptoFieldCa model
				updateEnrollmentCryptoFieldCaModel := new(blockchainv3.UpdateEnrollmentCryptoFieldCa)
				Expect(updateEnrollmentCryptoFieldCaModel).ToNot(BeNil())
				updateEnrollmentCryptoFieldCaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				updateEnrollmentCryptoFieldCaModel.Port = core.Float64Ptr(float64(7054))
				updateEnrollmentCryptoFieldCaModel.Name = core.StringPtr("ca")
				updateEnrollmentCryptoFieldCaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateEnrollmentCryptoFieldCaModel.EnrollID = core.StringPtr("admin")
				updateEnrollmentCryptoFieldCaModel.EnrollSecret = core.StringPtr("password")
				Expect(updateEnrollmentCryptoFieldCaModel.Host).To(Equal(core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")))
				Expect(updateEnrollmentCryptoFieldCaModel.Port).To(Equal(core.Float64Ptr(float64(7054))))
				Expect(updateEnrollmentCryptoFieldCaModel.Name).To(Equal(core.StringPtr("ca")))
				Expect(updateEnrollmentCryptoFieldCaModel.TlsCert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(updateEnrollmentCryptoFieldCaModel.EnrollID).To(Equal(core.StringPtr("admin")))
				Expect(updateEnrollmentCryptoFieldCaModel.EnrollSecret).To(Equal(core.StringPtr("password")))

				// Construct an instance of the UpdateEnrollmentCryptoFieldTlsca model
				updateEnrollmentCryptoFieldTlscaModel := new(blockchainv3.UpdateEnrollmentCryptoFieldTlsca)
				Expect(updateEnrollmentCryptoFieldTlscaModel).ToNot(BeNil())
				updateEnrollmentCryptoFieldTlscaModel.Host = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				updateEnrollmentCryptoFieldTlscaModel.Port = core.Float64Ptr(float64(7054))
				updateEnrollmentCryptoFieldTlscaModel.Name = core.StringPtr("tlsca")
				updateEnrollmentCryptoFieldTlscaModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateEnrollmentCryptoFieldTlscaModel.EnrollID = core.StringPtr("admin")
				updateEnrollmentCryptoFieldTlscaModel.EnrollSecret = core.StringPtr("password")
				updateEnrollmentCryptoFieldTlscaModel.CsrHosts = []string{"testString"}
				Expect(updateEnrollmentCryptoFieldTlscaModel.Host).To(Equal(core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")))
				Expect(updateEnrollmentCryptoFieldTlscaModel.Port).To(Equal(core.Float64Ptr(float64(7054))))
				Expect(updateEnrollmentCryptoFieldTlscaModel.Name).To(Equal(core.StringPtr("tlsca")))
				Expect(updateEnrollmentCryptoFieldTlscaModel.TlsCert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(updateEnrollmentCryptoFieldTlscaModel.EnrollID).To(Equal(core.StringPtr("admin")))
				Expect(updateEnrollmentCryptoFieldTlscaModel.EnrollSecret).To(Equal(core.StringPtr("password")))
				Expect(updateEnrollmentCryptoFieldTlscaModel.CsrHosts).To(Equal([]string{"testString"}))

				// Construct an instance of the UpdateEnrollmentCryptoField model
				updateEnrollmentCryptoFieldModel := new(blockchainv3.UpdateEnrollmentCryptoField)
				Expect(updateEnrollmentCryptoFieldModel).ToNot(BeNil())
				updateEnrollmentCryptoFieldModel.Component = cryptoEnrollmentComponentModel
				updateEnrollmentCryptoFieldModel.Ca = updateEnrollmentCryptoFieldCaModel
				updateEnrollmentCryptoFieldModel.Tlsca = updateEnrollmentCryptoFieldTlscaModel
				Expect(updateEnrollmentCryptoFieldModel.Component).To(Equal(cryptoEnrollmentComponentModel))
				Expect(updateEnrollmentCryptoFieldModel.Ca).To(Equal(updateEnrollmentCryptoFieldCaModel))
				Expect(updateEnrollmentCryptoFieldModel.Tlsca).To(Equal(updateEnrollmentCryptoFieldTlscaModel))

				// Construct an instance of the UpdateMspCryptoFieldCa model
				updateMspCryptoFieldCaModel := new(blockchainv3.UpdateMspCryptoFieldCa)
				Expect(updateMspCryptoFieldCaModel).ToNot(BeNil())
				updateMspCryptoFieldCaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateMspCryptoFieldCaModel.CaIntermediateCerts = []string{"testString"}
				Expect(updateMspCryptoFieldCaModel.RootCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(updateMspCryptoFieldCaModel.CaIntermediateCerts).To(Equal([]string{"testString"}))

				// Construct an instance of the UpdateMspCryptoFieldTlsca model
				updateMspCryptoFieldTlscaModel := new(blockchainv3.UpdateMspCryptoFieldTlsca)
				Expect(updateMspCryptoFieldTlscaModel).ToNot(BeNil())
				updateMspCryptoFieldTlscaModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateMspCryptoFieldTlscaModel.CaIntermediateCerts = []string{"testString"}
				Expect(updateMspCryptoFieldTlscaModel.RootCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(updateMspCryptoFieldTlscaModel.CaIntermediateCerts).To(Equal([]string{"testString"}))

				// Construct an instance of the ClientAuth model
				clientAuthModel := new(blockchainv3.ClientAuth)
				Expect(clientAuthModel).ToNot(BeNil())
				clientAuthModel.Type = core.StringPtr("noclientcert")
				clientAuthModel.TlsCerts = []string{"testString"}
				Expect(clientAuthModel.Type).To(Equal(core.StringPtr("noclientcert")))
				Expect(clientAuthModel.TlsCerts).To(Equal([]string{"testString"}))

				// Construct an instance of the UpdateMspCryptoFieldComponent model
				updateMspCryptoFieldComponentModel := new(blockchainv3.UpdateMspCryptoFieldComponent)
				Expect(updateMspCryptoFieldComponentModel).ToNot(BeNil())
				updateMspCryptoFieldComponentModel.Ekey = core.StringPtr("testString")
				updateMspCryptoFieldComponentModel.Ecert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateMspCryptoFieldComponentModel.AdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				updateMspCryptoFieldComponentModel.TlsKey = core.StringPtr("testString")
				updateMspCryptoFieldComponentModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				updateMspCryptoFieldComponentModel.ClientAuth = clientAuthModel
				Expect(updateMspCryptoFieldComponentModel.Ekey).To(Equal(core.StringPtr("testString")))
				Expect(updateMspCryptoFieldComponentModel.Ecert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(updateMspCryptoFieldComponentModel.AdminCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(updateMspCryptoFieldComponentModel.TlsKey).To(Equal(core.StringPtr("testString")))
				Expect(updateMspCryptoFieldComponentModel.TlsCert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(updateMspCryptoFieldComponentModel.ClientAuth).To(Equal(clientAuthModel))

				// Construct an instance of the UpdateMspCryptoField model
				updateMspCryptoFieldModel := new(blockchainv3.UpdateMspCryptoField)
				Expect(updateMspCryptoFieldModel).ToNot(BeNil())
				updateMspCryptoFieldModel.Ca = updateMspCryptoFieldCaModel
				updateMspCryptoFieldModel.Tlsca = updateMspCryptoFieldTlscaModel
				updateMspCryptoFieldModel.Component = updateMspCryptoFieldComponentModel
				Expect(updateMspCryptoFieldModel.Ca).To(Equal(updateMspCryptoFieldCaModel))
				Expect(updateMspCryptoFieldModel.Tlsca).To(Equal(updateMspCryptoFieldTlscaModel))
				Expect(updateMspCryptoFieldModel.Component).To(Equal(updateMspCryptoFieldComponentModel))

				// Construct an instance of the UpdatePeerBodyCrypto model
				updatePeerBodyCryptoModel := new(blockchainv3.UpdatePeerBodyCrypto)
				Expect(updatePeerBodyCryptoModel).ToNot(BeNil())
				updatePeerBodyCryptoModel.Enrollment = updateEnrollmentCryptoFieldModel
				updatePeerBodyCryptoModel.Msp = updateMspCryptoFieldModel
				Expect(updatePeerBodyCryptoModel.Enrollment).To(Equal(updateEnrollmentCryptoFieldModel))
				Expect(updatePeerBodyCryptoModel.Msp).To(Equal(updateMspCryptoFieldModel))

				// Construct an instance of the NodeOu model
				nodeOuModel := new(blockchainv3.NodeOu)
				Expect(nodeOuModel).ToNot(BeNil())
				nodeOuModel.Enabled = core.BoolPtr(true)
				Expect(nodeOuModel.Enabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv3.ResourceRequests)
				Expect(resourceRequestsModel).ToNot(BeNil())
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")
				Expect(resourceRequestsModel.Cpu).To(Equal(core.StringPtr("100m")))
				Expect(resourceRequestsModel.Memory).To(Equal(core.StringPtr("256MiB")))

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv3.ResourceLimits)
				Expect(resourceLimitsModel).ToNot(BeNil())
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")
				Expect(resourceLimitsModel.Cpu).To(Equal(core.StringPtr("100m")))
				Expect(resourceLimitsModel.Memory).To(Equal(core.StringPtr("256MiB")))

				// Construct an instance of the ResourceObjectFabV2 model
				resourceObjectFabV2Model := new(blockchainv3.ResourceObjectFabV2)
				Expect(resourceObjectFabV2Model).ToNot(BeNil())
				resourceObjectFabV2Model.Requests = resourceRequestsModel
				resourceObjectFabV2Model.Limits = resourceLimitsModel
				Expect(resourceObjectFabV2Model.Requests).To(Equal(resourceRequestsModel))
				Expect(resourceObjectFabV2Model.Limits).To(Equal(resourceLimitsModel))

				// Construct an instance of the ResourceObjectCouchDb model
				resourceObjectCouchDbModel := new(blockchainv3.ResourceObjectCouchDb)
				Expect(resourceObjectCouchDbModel).ToNot(BeNil())
				resourceObjectCouchDbModel.Requests = resourceRequestsModel
				resourceObjectCouchDbModel.Limits = resourceLimitsModel
				Expect(resourceObjectCouchDbModel.Requests).To(Equal(resourceRequestsModel))
				Expect(resourceObjectCouchDbModel.Limits).To(Equal(resourceLimitsModel))

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv3.ResourceObject)
				Expect(resourceObjectModel).ToNot(BeNil())
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel
				Expect(resourceObjectModel.Requests).To(Equal(resourceRequestsModel))
				Expect(resourceObjectModel.Limits).To(Equal(resourceLimitsModel))

				// Construct an instance of the ResourceObjectFabV1 model
				resourceObjectFabV1Model := new(blockchainv3.ResourceObjectFabV1)
				Expect(resourceObjectFabV1Model).ToNot(BeNil())
				resourceObjectFabV1Model.Requests = resourceRequestsModel
				resourceObjectFabV1Model.Limits = resourceLimitsModel
				Expect(resourceObjectFabV1Model.Requests).To(Equal(resourceRequestsModel))
				Expect(resourceObjectFabV1Model.Limits).To(Equal(resourceLimitsModel))

				// Construct an instance of the PeerResources model
				peerResourcesModel := new(blockchainv3.PeerResources)
				Expect(peerResourcesModel).ToNot(BeNil())
				peerResourcesModel.Chaincodelauncher = resourceObjectFabV2Model
				peerResourcesModel.Couchdb = resourceObjectCouchDbModel
				peerResourcesModel.Statedb = resourceObjectModel
				peerResourcesModel.Dind = resourceObjectFabV1Model
				peerResourcesModel.Fluentd = resourceObjectFabV1Model
				peerResourcesModel.Peer = resourceObjectModel
				peerResourcesModel.Proxy = resourceObjectModel
				Expect(peerResourcesModel.Chaincodelauncher).To(Equal(resourceObjectFabV2Model))
				Expect(peerResourcesModel.Couchdb).To(Equal(resourceObjectCouchDbModel))
				Expect(peerResourcesModel.Statedb).To(Equal(resourceObjectModel))
				Expect(peerResourcesModel.Dind).To(Equal(resourceObjectFabV1Model))
				Expect(peerResourcesModel.Fluentd).To(Equal(resourceObjectFabV1Model))
				Expect(peerResourcesModel.Peer).To(Equal(resourceObjectModel))
				Expect(peerResourcesModel.Proxy).To(Equal(resourceObjectModel))

				// Construct an instance of the UpdatePeerOptions model
				id := "testString"
				updatePeerOptionsModel := blockchainService.NewUpdatePeerOptions(id)
				updatePeerOptionsModel.SetID("testString")
				updatePeerOptionsModel.SetAdminCerts([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="})
				updatePeerOptionsModel.SetConfigOverride(configPeerUpdateModel)
				updatePeerOptionsModel.SetCrypto(updatePeerBodyCryptoModel)
				updatePeerOptionsModel.SetNodeOu(nodeOuModel)
				updatePeerOptionsModel.SetReplicas(float64(1))
				updatePeerOptionsModel.SetResources(peerResourcesModel)
				updatePeerOptionsModel.SetVersion("1.4.6-1")
				updatePeerOptionsModel.SetZone("-")
				updatePeerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updatePeerOptionsModel).ToNot(BeNil())
				Expect(updatePeerOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updatePeerOptionsModel.AdminCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(updatePeerOptionsModel.ConfigOverride).To(Equal(configPeerUpdateModel))
				Expect(updatePeerOptionsModel.Crypto).To(Equal(updatePeerBodyCryptoModel))
				Expect(updatePeerOptionsModel.NodeOu).To(Equal(nodeOuModel))
				Expect(updatePeerOptionsModel.Replicas).To(Equal(core.Float64Ptr(float64(1))))
				Expect(updatePeerOptionsModel.Resources).To(Equal(peerResourcesModel))
				Expect(updatePeerOptionsModel.Version).To(Equal(core.StringPtr("1.4.6-1")))
				Expect(updatePeerOptionsModel.Zone).To(Equal(core.StringPtr("-")))
				Expect(updatePeerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHsm successfully`, func() {
				pkcs11endpoint := "tcp://example.com:666"
				model, err := blockchainService.NewHsm(pkcs11endpoint)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewMspCryptoField successfully`, func() {
				var tlsca *blockchainv3.MspCryptoFieldTlsca = nil
				var component *blockchainv3.MspCryptoFieldComponent = nil
				_, err := blockchainService.NewMspCryptoField(tlsca, component)
				Expect(err).ToNot(BeNil())
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
