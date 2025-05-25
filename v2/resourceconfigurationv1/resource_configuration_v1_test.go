/**
 * (C) Copyright IBM Corp. 2025.
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

package resourceconfigurationv1_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-cos-sdk-go-config/v2/resourceconfigurationv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`ResourceConfigurationV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(resourceConfigurationService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(resourceConfigurationService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
				URL: "https://resourceconfigurationv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(resourceConfigurationService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"RESOURCE_CONFIGURATION_URL":       "https://resourceconfigurationv1/api",
				"RESOURCE_CONFIGURATION_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1UsingExternalConfig(&resourceconfigurationv1.ResourceConfigurationV1Options{})
				Expect(resourceConfigurationService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := resourceConfigurationService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != resourceConfigurationService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(resourceConfigurationService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(resourceConfigurationService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1UsingExternalConfig(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL: "https://testService/api",
				})
				Expect(resourceConfigurationService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := resourceConfigurationService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != resourceConfigurationService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(resourceConfigurationService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(resourceConfigurationService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1UsingExternalConfig(&resourceconfigurationv1.ResourceConfigurationV1Options{})
				err := resourceConfigurationService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := resourceConfigurationService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != resourceConfigurationService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(resourceConfigurationService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(resourceConfigurationService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"RESOURCE_CONFIGURATION_URL":       "https://resourceconfigurationv1/api",
				"RESOURCE_CONFIGURATION_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1UsingExternalConfig(&resourceconfigurationv1.ResourceConfigurationV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(resourceConfigurationService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"RESOURCE_CONFIGURATION_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1UsingExternalConfig(&resourceconfigurationv1.ResourceConfigurationV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(resourceConfigurationService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = resourceconfigurationv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateBackupPolicy(createBackupPolicyOptions *CreateBackupPolicyOptions) - Operation response error`, func() {
		createBackupPolicyPath := "/buckets/testString/backup_policies"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createBackupPolicyPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Md5"]).ToNot(BeNil())
					Expect(req.Header["Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateBackupPolicy with error: Operation response processing error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the DeleteAfterDays model
				deleteAfterDaysModel := new(resourceconfigurationv1.DeleteAfterDays)
				deleteAfterDaysModel.DeleteAfterDays = core.Int64Ptr(int64(10))

				// Construct an instance of the CreateBackupPolicyOptions model
				createBackupPolicyOptionsModel := new(resourceconfigurationv1.CreateBackupPolicyOptions)
				createBackupPolicyOptionsModel.Bucket = core.StringPtr("testString")
				createBackupPolicyOptionsModel.InitialRetention = deleteAfterDaysModel
				createBackupPolicyOptionsModel.PolicyName = core.StringPtr("myBackupPolicy")
				createBackupPolicyOptionsModel.TargetBackupVaultCrn = core.StringPtr("testString")
				createBackupPolicyOptionsModel.BackupType = core.StringPtr("continuous")
				createBackupPolicyOptionsModel.MD5 = core.StringPtr("testString")
				createBackupPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceConfigurationService.CreateBackupPolicy(createBackupPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = resourceConfigurationService.CreateBackupPolicy(createBackupPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateBackupPolicy(createBackupPolicyOptions *CreateBackupPolicyOptions)`, func() {
		createBackupPolicyPath := "/buckets/testString/backup_policies"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createBackupPolicyPath))
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

					Expect(req.Header["Md5"]).ToNot(BeNil())
					Expect(req.Header["Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"initial_retention": {"delete_after_days": 10}, "policy_name": "myBackupPolicy", "target_backup_vault_crn": "TargetBackupVaultCrn", "backup_type": "continuous", "policy_id": "PolicyID", "policy_status": "pending", "initial_sync_progress": 50.0, "error_cause": "Some error has occurred that prevents backup coverage from being created."}`)
				}))
			})
			It(`Invoke CreateBackupPolicy successfully with retries`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())
				resourceConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the DeleteAfterDays model
				deleteAfterDaysModel := new(resourceconfigurationv1.DeleteAfterDays)
				deleteAfterDaysModel.DeleteAfterDays = core.Int64Ptr(int64(10))

				// Construct an instance of the CreateBackupPolicyOptions model
				createBackupPolicyOptionsModel := new(resourceconfigurationv1.CreateBackupPolicyOptions)
				createBackupPolicyOptionsModel.Bucket = core.StringPtr("testString")
				createBackupPolicyOptionsModel.InitialRetention = deleteAfterDaysModel
				createBackupPolicyOptionsModel.PolicyName = core.StringPtr("myBackupPolicy")
				createBackupPolicyOptionsModel.TargetBackupVaultCrn = core.StringPtr("testString")
				createBackupPolicyOptionsModel.BackupType = core.StringPtr("continuous")
				createBackupPolicyOptionsModel.MD5 = core.StringPtr("testString")
				createBackupPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceConfigurationService.CreateBackupPolicyWithContext(ctx, createBackupPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceConfigurationService.DisableRetries()
				result, response, operationErr := resourceConfigurationService.CreateBackupPolicy(createBackupPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceConfigurationService.CreateBackupPolicyWithContext(ctx, createBackupPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createBackupPolicyPath))
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

					Expect(req.Header["Md5"]).ToNot(BeNil())
					Expect(req.Header["Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"initial_retention": {"delete_after_days": 10}, "policy_name": "myBackupPolicy", "target_backup_vault_crn": "TargetBackupVaultCrn", "backup_type": "continuous", "policy_id": "PolicyID", "policy_status": "pending", "initial_sync_progress": 50.0, "error_cause": "Some error has occurred that prevents backup coverage from being created."}`)
				}))
			})
			It(`Invoke CreateBackupPolicy successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceConfigurationService.CreateBackupPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteAfterDays model
				deleteAfterDaysModel := new(resourceconfigurationv1.DeleteAfterDays)
				deleteAfterDaysModel.DeleteAfterDays = core.Int64Ptr(int64(10))

				// Construct an instance of the CreateBackupPolicyOptions model
				createBackupPolicyOptionsModel := new(resourceconfigurationv1.CreateBackupPolicyOptions)
				createBackupPolicyOptionsModel.Bucket = core.StringPtr("testString")
				createBackupPolicyOptionsModel.InitialRetention = deleteAfterDaysModel
				createBackupPolicyOptionsModel.PolicyName = core.StringPtr("myBackupPolicy")
				createBackupPolicyOptionsModel.TargetBackupVaultCrn = core.StringPtr("testString")
				createBackupPolicyOptionsModel.BackupType = core.StringPtr("continuous")
				createBackupPolicyOptionsModel.MD5 = core.StringPtr("testString")
				createBackupPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceConfigurationService.CreateBackupPolicy(createBackupPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateBackupPolicy with error: Operation validation and request error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the DeleteAfterDays model
				deleteAfterDaysModel := new(resourceconfigurationv1.DeleteAfterDays)
				deleteAfterDaysModel.DeleteAfterDays = core.Int64Ptr(int64(10))

				// Construct an instance of the CreateBackupPolicyOptions model
				createBackupPolicyOptionsModel := new(resourceconfigurationv1.CreateBackupPolicyOptions)
				createBackupPolicyOptionsModel.Bucket = core.StringPtr("testString")
				createBackupPolicyOptionsModel.InitialRetention = deleteAfterDaysModel
				createBackupPolicyOptionsModel.PolicyName = core.StringPtr("myBackupPolicy")
				createBackupPolicyOptionsModel.TargetBackupVaultCrn = core.StringPtr("testString")
				createBackupPolicyOptionsModel.BackupType = core.StringPtr("continuous")
				createBackupPolicyOptionsModel.MD5 = core.StringPtr("testString")
				createBackupPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceConfigurationService.CreateBackupPolicy(createBackupPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateBackupPolicyOptions model with no property values
				createBackupPolicyOptionsModelNew := new(resourceconfigurationv1.CreateBackupPolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceConfigurationService.CreateBackupPolicy(createBackupPolicyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateBackupPolicy successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the DeleteAfterDays model
				deleteAfterDaysModel := new(resourceconfigurationv1.DeleteAfterDays)
				deleteAfterDaysModel.DeleteAfterDays = core.Int64Ptr(int64(10))

				// Construct an instance of the CreateBackupPolicyOptions model
				createBackupPolicyOptionsModel := new(resourceconfigurationv1.CreateBackupPolicyOptions)
				createBackupPolicyOptionsModel.Bucket = core.StringPtr("testString")
				createBackupPolicyOptionsModel.InitialRetention = deleteAfterDaysModel
				createBackupPolicyOptionsModel.PolicyName = core.StringPtr("myBackupPolicy")
				createBackupPolicyOptionsModel.TargetBackupVaultCrn = core.StringPtr("testString")
				createBackupPolicyOptionsModel.BackupType = core.StringPtr("continuous")
				createBackupPolicyOptionsModel.MD5 = core.StringPtr("testString")
				createBackupPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceConfigurationService.CreateBackupPolicy(createBackupPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListBackupPolicies(listBackupPoliciesOptions *ListBackupPoliciesOptions) - Operation response error`, func() {
		listBackupPoliciesPath := "/buckets/testString/backup_policies"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBackupPoliciesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListBackupPolicies with error: Operation response processing error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListBackupPoliciesOptions model
				listBackupPoliciesOptionsModel := new(resourceconfigurationv1.ListBackupPoliciesOptions)
				listBackupPoliciesOptionsModel.Bucket = core.StringPtr("testString")
				listBackupPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceConfigurationService.ListBackupPolicies(listBackupPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = resourceConfigurationService.ListBackupPolicies(listBackupPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListBackupPolicies(listBackupPoliciesOptions *ListBackupPoliciesOptions)`, func() {
		listBackupPoliciesPath := "/buckets/testString/backup_policies"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBackupPoliciesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"backup_policies": [{"initial_retention": {"delete_after_days": 10}, "policy_name": "myBackupPolicy", "target_backup_vault_crn": "TargetBackupVaultCrn", "backup_type": "continuous", "policy_id": "PolicyID", "policy_status": "pending", "initial_sync_progress": 50.0, "error_cause": "Some error has occurred that prevents backup coverage from being created."}]}`)
				}))
			})
			It(`Invoke ListBackupPolicies successfully with retries`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())
				resourceConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the ListBackupPoliciesOptions model
				listBackupPoliciesOptionsModel := new(resourceconfigurationv1.ListBackupPoliciesOptions)
				listBackupPoliciesOptionsModel.Bucket = core.StringPtr("testString")
				listBackupPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceConfigurationService.ListBackupPoliciesWithContext(ctx, listBackupPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceConfigurationService.DisableRetries()
				result, response, operationErr := resourceConfigurationService.ListBackupPolicies(listBackupPoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceConfigurationService.ListBackupPoliciesWithContext(ctx, listBackupPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBackupPoliciesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"backup_policies": [{"initial_retention": {"delete_after_days": 10}, "policy_name": "myBackupPolicy", "target_backup_vault_crn": "TargetBackupVaultCrn", "backup_type": "continuous", "policy_id": "PolicyID", "policy_status": "pending", "initial_sync_progress": 50.0, "error_cause": "Some error has occurred that prevents backup coverage from being created."}]}`)
				}))
			})
			It(`Invoke ListBackupPolicies successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceConfigurationService.ListBackupPolicies(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListBackupPoliciesOptions model
				listBackupPoliciesOptionsModel := new(resourceconfigurationv1.ListBackupPoliciesOptions)
				listBackupPoliciesOptionsModel.Bucket = core.StringPtr("testString")
				listBackupPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceConfigurationService.ListBackupPolicies(listBackupPoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListBackupPolicies with error: Operation validation and request error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListBackupPoliciesOptions model
				listBackupPoliciesOptionsModel := new(resourceconfigurationv1.ListBackupPoliciesOptions)
				listBackupPoliciesOptionsModel.Bucket = core.StringPtr("testString")
				listBackupPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceConfigurationService.ListBackupPolicies(listBackupPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListBackupPoliciesOptions model with no property values
				listBackupPoliciesOptionsModelNew := new(resourceconfigurationv1.ListBackupPoliciesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceConfigurationService.ListBackupPolicies(listBackupPoliciesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListBackupPolicies successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListBackupPoliciesOptions model
				listBackupPoliciesOptionsModel := new(resourceconfigurationv1.ListBackupPoliciesOptions)
				listBackupPoliciesOptionsModel.Bucket = core.StringPtr("testString")
				listBackupPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceConfigurationService.ListBackupPolicies(listBackupPoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBackupPolicy(getBackupPolicyOptions *GetBackupPolicyOptions) - Operation response error`, func() {
		getBackupPolicyPath := "/buckets/testString/backup_policies/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBackupPolicyPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBackupPolicy with error: Operation response processing error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetBackupPolicyOptions model
				getBackupPolicyOptionsModel := new(resourceconfigurationv1.GetBackupPolicyOptions)
				getBackupPolicyOptionsModel.Bucket = core.StringPtr("testString")
				getBackupPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				getBackupPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceConfigurationService.GetBackupPolicy(getBackupPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = resourceConfigurationService.GetBackupPolicy(getBackupPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBackupPolicy(getBackupPolicyOptions *GetBackupPolicyOptions)`, func() {
		getBackupPolicyPath := "/buckets/testString/backup_policies/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBackupPolicyPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"initial_retention": {"delete_after_days": 10}, "policy_name": "myBackupPolicy", "target_backup_vault_crn": "TargetBackupVaultCrn", "backup_type": "continuous", "policy_id": "PolicyID", "policy_status": "pending", "initial_sync_progress": 50.0, "error_cause": "Some error has occurred that prevents backup coverage from being created."}`)
				}))
			})
			It(`Invoke GetBackupPolicy successfully with retries`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())
				resourceConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the GetBackupPolicyOptions model
				getBackupPolicyOptionsModel := new(resourceconfigurationv1.GetBackupPolicyOptions)
				getBackupPolicyOptionsModel.Bucket = core.StringPtr("testString")
				getBackupPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				getBackupPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceConfigurationService.GetBackupPolicyWithContext(ctx, getBackupPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceConfigurationService.DisableRetries()
				result, response, operationErr := resourceConfigurationService.GetBackupPolicy(getBackupPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceConfigurationService.GetBackupPolicyWithContext(ctx, getBackupPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBackupPolicyPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"initial_retention": {"delete_after_days": 10}, "policy_name": "myBackupPolicy", "target_backup_vault_crn": "TargetBackupVaultCrn", "backup_type": "continuous", "policy_id": "PolicyID", "policy_status": "pending", "initial_sync_progress": 50.0, "error_cause": "Some error has occurred that prevents backup coverage from being created."}`)
				}))
			})
			It(`Invoke GetBackupPolicy successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceConfigurationService.GetBackupPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBackupPolicyOptions model
				getBackupPolicyOptionsModel := new(resourceconfigurationv1.GetBackupPolicyOptions)
				getBackupPolicyOptionsModel.Bucket = core.StringPtr("testString")
				getBackupPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				getBackupPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceConfigurationService.GetBackupPolicy(getBackupPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetBackupPolicy with error: Operation validation and request error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetBackupPolicyOptions model
				getBackupPolicyOptionsModel := new(resourceconfigurationv1.GetBackupPolicyOptions)
				getBackupPolicyOptionsModel.Bucket = core.StringPtr("testString")
				getBackupPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				getBackupPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceConfigurationService.GetBackupPolicy(getBackupPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetBackupPolicyOptions model with no property values
				getBackupPolicyOptionsModelNew := new(resourceconfigurationv1.GetBackupPolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceConfigurationService.GetBackupPolicy(getBackupPolicyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetBackupPolicy successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetBackupPolicyOptions model
				getBackupPolicyOptionsModel := new(resourceconfigurationv1.GetBackupPolicyOptions)
				getBackupPolicyOptionsModel.Bucket = core.StringPtr("testString")
				getBackupPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				getBackupPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceConfigurationService.GetBackupPolicy(getBackupPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteBackupPolicy(deleteBackupPolicyOptions *DeleteBackupPolicyOptions)`, func() {
		deleteBackupPolicyPath := "/buckets/testString/backup_policies/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteBackupPolicyPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteBackupPolicy successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := resourceConfigurationService.DeleteBackupPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteBackupPolicyOptions model
				deleteBackupPolicyOptionsModel := new(resourceconfigurationv1.DeleteBackupPolicyOptions)
				deleteBackupPolicyOptionsModel.Bucket = core.StringPtr("testString")
				deleteBackupPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				deleteBackupPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = resourceConfigurationService.DeleteBackupPolicy(deleteBackupPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteBackupPolicy with error: Operation validation and request error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the DeleteBackupPolicyOptions model
				deleteBackupPolicyOptionsModel := new(resourceconfigurationv1.DeleteBackupPolicyOptions)
				deleteBackupPolicyOptionsModel.Bucket = core.StringPtr("testString")
				deleteBackupPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				deleteBackupPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := resourceConfigurationService.DeleteBackupPolicy(deleteBackupPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteBackupPolicyOptions model with no property values
				deleteBackupPolicyOptionsModelNew := new(resourceconfigurationv1.DeleteBackupPolicyOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = resourceConfigurationService.DeleteBackupPolicy(deleteBackupPolicyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListBackupVaults(listBackupVaultsOptions *ListBackupVaultsOptions) - Operation response error`, func() {
		listBackupVaultsPath := "/backup_vaults"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBackupVaultsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["service_instance_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["token"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListBackupVaults with error: Operation response processing error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListBackupVaultsOptions model
				listBackupVaultsOptionsModel := new(resourceconfigurationv1.ListBackupVaultsOptions)
				listBackupVaultsOptionsModel.ServiceInstanceID = core.StringPtr("testString")
				listBackupVaultsOptionsModel.Token = core.StringPtr("testString")
				listBackupVaultsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceConfigurationService.ListBackupVaults(listBackupVaultsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = resourceConfigurationService.ListBackupVaults(listBackupVaultsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListBackupVaults(listBackupVaultsOptions *ListBackupVaultsOptions)`, func() {
		listBackupVaultsPath := "/backup_vaults"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBackupVaultsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["service_instance_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["token"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"next": {"href": "Href", "token": "Token"}, "backup_vaults": ["BackupVaults"]}`)
				}))
			})
			It(`Invoke ListBackupVaults successfully with retries`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())
				resourceConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the ListBackupVaultsOptions model
				listBackupVaultsOptionsModel := new(resourceconfigurationv1.ListBackupVaultsOptions)
				listBackupVaultsOptionsModel.ServiceInstanceID = core.StringPtr("testString")
				listBackupVaultsOptionsModel.Token = core.StringPtr("testString")
				listBackupVaultsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceConfigurationService.ListBackupVaultsWithContext(ctx, listBackupVaultsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceConfigurationService.DisableRetries()
				result, response, operationErr := resourceConfigurationService.ListBackupVaults(listBackupVaultsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceConfigurationService.ListBackupVaultsWithContext(ctx, listBackupVaultsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBackupVaultsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["service_instance_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["token"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"next": {"href": "Href", "token": "Token"}, "backup_vaults": ["BackupVaults"]}`)
				}))
			})
			It(`Invoke ListBackupVaults successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceConfigurationService.ListBackupVaults(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListBackupVaultsOptions model
				listBackupVaultsOptionsModel := new(resourceconfigurationv1.ListBackupVaultsOptions)
				listBackupVaultsOptionsModel.ServiceInstanceID = core.StringPtr("testString")
				listBackupVaultsOptionsModel.Token = core.StringPtr("testString")
				listBackupVaultsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceConfigurationService.ListBackupVaults(listBackupVaultsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListBackupVaults with error: Operation validation and request error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListBackupVaultsOptions model
				listBackupVaultsOptionsModel := new(resourceconfigurationv1.ListBackupVaultsOptions)
				listBackupVaultsOptionsModel.ServiceInstanceID = core.StringPtr("testString")
				listBackupVaultsOptionsModel.Token = core.StringPtr("testString")
				listBackupVaultsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceConfigurationService.ListBackupVaults(listBackupVaultsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListBackupVaultsOptions model with no property values
				listBackupVaultsOptionsModelNew := new(resourceconfigurationv1.ListBackupVaultsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceConfigurationService.ListBackupVaults(listBackupVaultsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListBackupVaults successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListBackupVaultsOptions model
				listBackupVaultsOptionsModel := new(resourceconfigurationv1.ListBackupVaultsOptions)
				listBackupVaultsOptionsModel.ServiceInstanceID = core.StringPtr("testString")
				listBackupVaultsOptionsModel.Token = core.StringPtr("testString")
				listBackupVaultsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceConfigurationService.ListBackupVaults(listBackupVaultsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextToken successfully`, func() {
				responseObject := new(resourceconfigurationv1.BackupVaultCollection)
				nextObject := new(resourceconfigurationv1.NextPagination)
				nextObject.Token = core.StringPtr("abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextToken()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextToken without a "Next" property in the response`, func() {
				responseObject := new(resourceconfigurationv1.BackupVaultCollection)

				value, err := responseObject.GetNextToken()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBackupVaultsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"token":"1"},"backup_vaults":["BackupVaults"],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"backup_vaults":["BackupVaults"],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use BackupVaultsPager.GetNext successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				listBackupVaultsOptionsModel := &resourceconfigurationv1.ListBackupVaultsOptions{
					ServiceInstanceID: core.StringPtr("testString"),
				}

				pager, err := resourceConfigurationService.NewBackupVaultsPager(listBackupVaultsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []string
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use BackupVaultsPager.GetAll successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				listBackupVaultsOptionsModel := &resourceconfigurationv1.ListBackupVaultsOptions{
					ServiceInstanceID: core.StringPtr("testString"),
				}

				pager, err := resourceConfigurationService.NewBackupVaultsPager(listBackupVaultsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateBackupVault(createBackupVaultOptions *CreateBackupVaultOptions) - Operation response error`, func() {
		createBackupVaultPath := "/backup_vaults"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createBackupVaultPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["service_instance_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateBackupVault with error: Operation response processing error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the BackupVaultActivityTracking model
				backupVaultActivityTrackingModel := new(resourceconfigurationv1.BackupVaultActivityTracking)
				backupVaultActivityTrackingModel.ManagementEvents = core.BoolPtr(true)

				// Construct an instance of the BackupVaultMetricsMonitoring model
				backupVaultMetricsMonitoringModel := new(resourceconfigurationv1.BackupVaultMetricsMonitoring)
				backupVaultMetricsMonitoringModel.UsageMetricsEnabled = core.BoolPtr(true)

				// Construct an instance of the CreateBackupVaultOptions model
				createBackupVaultOptionsModel := new(resourceconfigurationv1.CreateBackupVaultOptions)
				createBackupVaultOptionsModel.ServiceInstanceID = core.StringPtr("testString")
				createBackupVaultOptionsModel.BackupVaultName = core.StringPtr("testString")
				createBackupVaultOptionsModel.Region = core.StringPtr("testString")
				createBackupVaultOptionsModel.ActivityTracking = backupVaultActivityTrackingModel
				createBackupVaultOptionsModel.MetricsMonitoring = backupVaultMetricsMonitoringModel
				createBackupVaultOptionsModel.SseKpCustomerRootKeyCrn = core.StringPtr("testString")
				createBackupVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceConfigurationService.CreateBackupVault(createBackupVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = resourceConfigurationService.CreateBackupVault(createBackupVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateBackupVault(createBackupVaultOptions *CreateBackupVaultOptions)`, func() {
		createBackupVaultPath := "/backup_vaults"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createBackupVaultPath))
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

					Expect(req.URL.Query()["service_instance_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"activity_tracking": {"management_events": true}, "metrics_monitoring": {"usage_metrics_enabled": false}, "backup_vault_name": "BackupVaultName", "region": "Region", "sse_kp_customer_root_key_crn": "SseKpCustomerRootKeyCrn", "crn": "Crn", "service_instance_crn": "ServiceInstanceCrn", "time_created": "2019-01-01T12:00:00.000Z", "time_updated": "2019-01-01T12:00:00.000Z", "bytes_used": 0}`)
				}))
			})
			It(`Invoke CreateBackupVault successfully with retries`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())
				resourceConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the BackupVaultActivityTracking model
				backupVaultActivityTrackingModel := new(resourceconfigurationv1.BackupVaultActivityTracking)
				backupVaultActivityTrackingModel.ManagementEvents = core.BoolPtr(true)

				// Construct an instance of the BackupVaultMetricsMonitoring model
				backupVaultMetricsMonitoringModel := new(resourceconfigurationv1.BackupVaultMetricsMonitoring)
				backupVaultMetricsMonitoringModel.UsageMetricsEnabled = core.BoolPtr(true)

				// Construct an instance of the CreateBackupVaultOptions model
				createBackupVaultOptionsModel := new(resourceconfigurationv1.CreateBackupVaultOptions)
				createBackupVaultOptionsModel.ServiceInstanceID = core.StringPtr("testString")
				createBackupVaultOptionsModel.BackupVaultName = core.StringPtr("testString")
				createBackupVaultOptionsModel.Region = core.StringPtr("testString")
				createBackupVaultOptionsModel.ActivityTracking = backupVaultActivityTrackingModel
				createBackupVaultOptionsModel.MetricsMonitoring = backupVaultMetricsMonitoringModel
				createBackupVaultOptionsModel.SseKpCustomerRootKeyCrn = core.StringPtr("testString")
				createBackupVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceConfigurationService.CreateBackupVaultWithContext(ctx, createBackupVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceConfigurationService.DisableRetries()
				result, response, operationErr := resourceConfigurationService.CreateBackupVault(createBackupVaultOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceConfigurationService.CreateBackupVaultWithContext(ctx, createBackupVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createBackupVaultPath))
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

					Expect(req.URL.Query()["service_instance_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"activity_tracking": {"management_events": true}, "metrics_monitoring": {"usage_metrics_enabled": false}, "backup_vault_name": "BackupVaultName", "region": "Region", "sse_kp_customer_root_key_crn": "SseKpCustomerRootKeyCrn", "crn": "Crn", "service_instance_crn": "ServiceInstanceCrn", "time_created": "2019-01-01T12:00:00.000Z", "time_updated": "2019-01-01T12:00:00.000Z", "bytes_used": 0}`)
				}))
			})
			It(`Invoke CreateBackupVault successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceConfigurationService.CreateBackupVault(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the BackupVaultActivityTracking model
				backupVaultActivityTrackingModel := new(resourceconfigurationv1.BackupVaultActivityTracking)
				backupVaultActivityTrackingModel.ManagementEvents = core.BoolPtr(true)

				// Construct an instance of the BackupVaultMetricsMonitoring model
				backupVaultMetricsMonitoringModel := new(resourceconfigurationv1.BackupVaultMetricsMonitoring)
				backupVaultMetricsMonitoringModel.UsageMetricsEnabled = core.BoolPtr(true)

				// Construct an instance of the CreateBackupVaultOptions model
				createBackupVaultOptionsModel := new(resourceconfigurationv1.CreateBackupVaultOptions)
				createBackupVaultOptionsModel.ServiceInstanceID = core.StringPtr("testString")
				createBackupVaultOptionsModel.BackupVaultName = core.StringPtr("testString")
				createBackupVaultOptionsModel.Region = core.StringPtr("testString")
				createBackupVaultOptionsModel.ActivityTracking = backupVaultActivityTrackingModel
				createBackupVaultOptionsModel.MetricsMonitoring = backupVaultMetricsMonitoringModel
				createBackupVaultOptionsModel.SseKpCustomerRootKeyCrn = core.StringPtr("testString")
				createBackupVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceConfigurationService.CreateBackupVault(createBackupVaultOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateBackupVault with error: Operation validation and request error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the BackupVaultActivityTracking model
				backupVaultActivityTrackingModel := new(resourceconfigurationv1.BackupVaultActivityTracking)
				backupVaultActivityTrackingModel.ManagementEvents = core.BoolPtr(true)

				// Construct an instance of the BackupVaultMetricsMonitoring model
				backupVaultMetricsMonitoringModel := new(resourceconfigurationv1.BackupVaultMetricsMonitoring)
				backupVaultMetricsMonitoringModel.UsageMetricsEnabled = core.BoolPtr(true)

				// Construct an instance of the CreateBackupVaultOptions model
				createBackupVaultOptionsModel := new(resourceconfigurationv1.CreateBackupVaultOptions)
				createBackupVaultOptionsModel.ServiceInstanceID = core.StringPtr("testString")
				createBackupVaultOptionsModel.BackupVaultName = core.StringPtr("testString")
				createBackupVaultOptionsModel.Region = core.StringPtr("testString")
				createBackupVaultOptionsModel.ActivityTracking = backupVaultActivityTrackingModel
				createBackupVaultOptionsModel.MetricsMonitoring = backupVaultMetricsMonitoringModel
				createBackupVaultOptionsModel.SseKpCustomerRootKeyCrn = core.StringPtr("testString")
				createBackupVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceConfigurationService.CreateBackupVault(createBackupVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateBackupVaultOptions model with no property values
				createBackupVaultOptionsModelNew := new(resourceconfigurationv1.CreateBackupVaultOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceConfigurationService.CreateBackupVault(createBackupVaultOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateBackupVault successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the BackupVaultActivityTracking model
				backupVaultActivityTrackingModel := new(resourceconfigurationv1.BackupVaultActivityTracking)
				backupVaultActivityTrackingModel.ManagementEvents = core.BoolPtr(true)

				// Construct an instance of the BackupVaultMetricsMonitoring model
				backupVaultMetricsMonitoringModel := new(resourceconfigurationv1.BackupVaultMetricsMonitoring)
				backupVaultMetricsMonitoringModel.UsageMetricsEnabled = core.BoolPtr(true)

				// Construct an instance of the CreateBackupVaultOptions model
				createBackupVaultOptionsModel := new(resourceconfigurationv1.CreateBackupVaultOptions)
				createBackupVaultOptionsModel.ServiceInstanceID = core.StringPtr("testString")
				createBackupVaultOptionsModel.BackupVaultName = core.StringPtr("testString")
				createBackupVaultOptionsModel.Region = core.StringPtr("testString")
				createBackupVaultOptionsModel.ActivityTracking = backupVaultActivityTrackingModel
				createBackupVaultOptionsModel.MetricsMonitoring = backupVaultMetricsMonitoringModel
				createBackupVaultOptionsModel.SseKpCustomerRootKeyCrn = core.StringPtr("testString")
				createBackupVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceConfigurationService.CreateBackupVault(createBackupVaultOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBackupVault(getBackupVaultOptions *GetBackupVaultOptions) - Operation response error`, func() {
		getBackupVaultPath := "/backup_vaults/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBackupVaultPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBackupVault with error: Operation response processing error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetBackupVaultOptions model
				getBackupVaultOptionsModel := new(resourceconfigurationv1.GetBackupVaultOptions)
				getBackupVaultOptionsModel.BackupVaultName = core.StringPtr("testString")
				getBackupVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceConfigurationService.GetBackupVault(getBackupVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = resourceConfigurationService.GetBackupVault(getBackupVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBackupVault(getBackupVaultOptions *GetBackupVaultOptions)`, func() {
		getBackupVaultPath := "/backup_vaults/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBackupVaultPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"activity_tracking": {"management_events": true}, "metrics_monitoring": {"usage_metrics_enabled": false}, "backup_vault_name": "BackupVaultName", "region": "Region", "sse_kp_customer_root_key_crn": "SseKpCustomerRootKeyCrn", "crn": "Crn", "service_instance_crn": "ServiceInstanceCrn", "time_created": "2019-01-01T12:00:00.000Z", "time_updated": "2019-01-01T12:00:00.000Z", "bytes_used": 0}`)
				}))
			})
			It(`Invoke GetBackupVault successfully with retries`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())
				resourceConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the GetBackupVaultOptions model
				getBackupVaultOptionsModel := new(resourceconfigurationv1.GetBackupVaultOptions)
				getBackupVaultOptionsModel.BackupVaultName = core.StringPtr("testString")
				getBackupVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceConfigurationService.GetBackupVaultWithContext(ctx, getBackupVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceConfigurationService.DisableRetries()
				result, response, operationErr := resourceConfigurationService.GetBackupVault(getBackupVaultOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceConfigurationService.GetBackupVaultWithContext(ctx, getBackupVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBackupVaultPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"activity_tracking": {"management_events": true}, "metrics_monitoring": {"usage_metrics_enabled": false}, "backup_vault_name": "BackupVaultName", "region": "Region", "sse_kp_customer_root_key_crn": "SseKpCustomerRootKeyCrn", "crn": "Crn", "service_instance_crn": "ServiceInstanceCrn", "time_created": "2019-01-01T12:00:00.000Z", "time_updated": "2019-01-01T12:00:00.000Z", "bytes_used": 0}`)
				}))
			})
			It(`Invoke GetBackupVault successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceConfigurationService.GetBackupVault(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBackupVaultOptions model
				getBackupVaultOptionsModel := new(resourceconfigurationv1.GetBackupVaultOptions)
				getBackupVaultOptionsModel.BackupVaultName = core.StringPtr("testString")
				getBackupVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceConfigurationService.GetBackupVault(getBackupVaultOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetBackupVault with error: Operation validation and request error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetBackupVaultOptions model
				getBackupVaultOptionsModel := new(resourceconfigurationv1.GetBackupVaultOptions)
				getBackupVaultOptionsModel.BackupVaultName = core.StringPtr("testString")
				getBackupVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceConfigurationService.GetBackupVault(getBackupVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetBackupVaultOptions model with no property values
				getBackupVaultOptionsModelNew := new(resourceconfigurationv1.GetBackupVaultOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceConfigurationService.GetBackupVault(getBackupVaultOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetBackupVault successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetBackupVaultOptions model
				getBackupVaultOptionsModel := new(resourceconfigurationv1.GetBackupVaultOptions)
				getBackupVaultOptionsModel.BackupVaultName = core.StringPtr("testString")
				getBackupVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceConfigurationService.GetBackupVault(getBackupVaultOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateBackupVault(updateBackupVaultOptions *UpdateBackupVaultOptions) - Operation response error`, func() {
		updateBackupVaultPath := "/backup_vaults/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateBackupVaultPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateBackupVault with error: Operation response processing error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the BackupVaultActivityTracking model
				backupVaultActivityTrackingModel := new(resourceconfigurationv1.BackupVaultActivityTracking)
				backupVaultActivityTrackingModel.ManagementEvents = core.BoolPtr(true)

				// Construct an instance of the BackupVaultMetricsMonitoring model
				backupVaultMetricsMonitoringModel := new(resourceconfigurationv1.BackupVaultMetricsMonitoring)
				backupVaultMetricsMonitoringModel.UsageMetricsEnabled = core.BoolPtr(true)

				// Construct an instance of the BackupVaultPatch model
				backupVaultPatchModel := new(resourceconfigurationv1.BackupVaultPatch)
				backupVaultPatchModel.ActivityTracking = backupVaultActivityTrackingModel
				backupVaultPatchModel.MetricsMonitoring = backupVaultMetricsMonitoringModel
				backupVaultPatchModelAsPatch, asPatchErr := backupVaultPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateBackupVaultOptions model
				updateBackupVaultOptionsModel := new(resourceconfigurationv1.UpdateBackupVaultOptions)
				updateBackupVaultOptionsModel.BackupVaultName = core.StringPtr("testString")
				updateBackupVaultOptionsModel.BackupVaultPatch = backupVaultPatchModelAsPatch
				updateBackupVaultOptionsModel.IfMatch = core.StringPtr("testString")
				updateBackupVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceConfigurationService.UpdateBackupVault(updateBackupVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = resourceConfigurationService.UpdateBackupVault(updateBackupVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateBackupVault(updateBackupVaultOptions *UpdateBackupVaultOptions)`, func() {
		updateBackupVaultPath := "/backup_vaults/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateBackupVaultPath))
					Expect(req.Method).To(Equal("PATCH"))

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

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"activity_tracking": {"management_events": true}, "metrics_monitoring": {"usage_metrics_enabled": false}, "backup_vault_name": "BackupVaultName", "region": "Region", "sse_kp_customer_root_key_crn": "SseKpCustomerRootKeyCrn", "crn": "Crn", "service_instance_crn": "ServiceInstanceCrn", "time_created": "2019-01-01T12:00:00.000Z", "time_updated": "2019-01-01T12:00:00.000Z", "bytes_used": 0}`)
				}))
			})
			It(`Invoke UpdateBackupVault successfully with retries`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())
				resourceConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the BackupVaultActivityTracking model
				backupVaultActivityTrackingModel := new(resourceconfigurationv1.BackupVaultActivityTracking)
				backupVaultActivityTrackingModel.ManagementEvents = core.BoolPtr(true)

				// Construct an instance of the BackupVaultMetricsMonitoring model
				backupVaultMetricsMonitoringModel := new(resourceconfigurationv1.BackupVaultMetricsMonitoring)
				backupVaultMetricsMonitoringModel.UsageMetricsEnabled = core.BoolPtr(true)

				// Construct an instance of the BackupVaultPatch model
				backupVaultPatchModel := new(resourceconfigurationv1.BackupVaultPatch)
				backupVaultPatchModel.ActivityTracking = backupVaultActivityTrackingModel
				backupVaultPatchModel.MetricsMonitoring = backupVaultMetricsMonitoringModel
				backupVaultPatchModelAsPatch, asPatchErr := backupVaultPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateBackupVaultOptions model
				updateBackupVaultOptionsModel := new(resourceconfigurationv1.UpdateBackupVaultOptions)
				updateBackupVaultOptionsModel.BackupVaultName = core.StringPtr("testString")
				updateBackupVaultOptionsModel.BackupVaultPatch = backupVaultPatchModelAsPatch
				updateBackupVaultOptionsModel.IfMatch = core.StringPtr("testString")
				updateBackupVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceConfigurationService.UpdateBackupVaultWithContext(ctx, updateBackupVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceConfigurationService.DisableRetries()
				result, response, operationErr := resourceConfigurationService.UpdateBackupVault(updateBackupVaultOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceConfigurationService.UpdateBackupVaultWithContext(ctx, updateBackupVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateBackupVaultPath))
					Expect(req.Method).To(Equal("PATCH"))

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

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"activity_tracking": {"management_events": true}, "metrics_monitoring": {"usage_metrics_enabled": false}, "backup_vault_name": "BackupVaultName", "region": "Region", "sse_kp_customer_root_key_crn": "SseKpCustomerRootKeyCrn", "crn": "Crn", "service_instance_crn": "ServiceInstanceCrn", "time_created": "2019-01-01T12:00:00.000Z", "time_updated": "2019-01-01T12:00:00.000Z", "bytes_used": 0}`)
				}))
			})
			It(`Invoke UpdateBackupVault successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceConfigurationService.UpdateBackupVault(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the BackupVaultActivityTracking model
				backupVaultActivityTrackingModel := new(resourceconfigurationv1.BackupVaultActivityTracking)
				backupVaultActivityTrackingModel.ManagementEvents = core.BoolPtr(true)

				// Construct an instance of the BackupVaultMetricsMonitoring model
				backupVaultMetricsMonitoringModel := new(resourceconfigurationv1.BackupVaultMetricsMonitoring)
				backupVaultMetricsMonitoringModel.UsageMetricsEnabled = core.BoolPtr(true)

				// Construct an instance of the BackupVaultPatch model
				backupVaultPatchModel := new(resourceconfigurationv1.BackupVaultPatch)
				backupVaultPatchModel.ActivityTracking = backupVaultActivityTrackingModel
				backupVaultPatchModel.MetricsMonitoring = backupVaultMetricsMonitoringModel
				backupVaultPatchModelAsPatch, asPatchErr := backupVaultPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateBackupVaultOptions model
				updateBackupVaultOptionsModel := new(resourceconfigurationv1.UpdateBackupVaultOptions)
				updateBackupVaultOptionsModel.BackupVaultName = core.StringPtr("testString")
				updateBackupVaultOptionsModel.BackupVaultPatch = backupVaultPatchModelAsPatch
				updateBackupVaultOptionsModel.IfMatch = core.StringPtr("testString")
				updateBackupVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceConfigurationService.UpdateBackupVault(updateBackupVaultOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateBackupVault with error: Operation validation and request error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the BackupVaultActivityTracking model
				backupVaultActivityTrackingModel := new(resourceconfigurationv1.BackupVaultActivityTracking)
				backupVaultActivityTrackingModel.ManagementEvents = core.BoolPtr(true)

				// Construct an instance of the BackupVaultMetricsMonitoring model
				backupVaultMetricsMonitoringModel := new(resourceconfigurationv1.BackupVaultMetricsMonitoring)
				backupVaultMetricsMonitoringModel.UsageMetricsEnabled = core.BoolPtr(true)

				// Construct an instance of the BackupVaultPatch model
				backupVaultPatchModel := new(resourceconfigurationv1.BackupVaultPatch)
				backupVaultPatchModel.ActivityTracking = backupVaultActivityTrackingModel
				backupVaultPatchModel.MetricsMonitoring = backupVaultMetricsMonitoringModel
				backupVaultPatchModelAsPatch, asPatchErr := backupVaultPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateBackupVaultOptions model
				updateBackupVaultOptionsModel := new(resourceconfigurationv1.UpdateBackupVaultOptions)
				updateBackupVaultOptionsModel.BackupVaultName = core.StringPtr("testString")
				updateBackupVaultOptionsModel.BackupVaultPatch = backupVaultPatchModelAsPatch
				updateBackupVaultOptionsModel.IfMatch = core.StringPtr("testString")
				updateBackupVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceConfigurationService.UpdateBackupVault(updateBackupVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateBackupVaultOptions model with no property values
				updateBackupVaultOptionsModelNew := new(resourceconfigurationv1.UpdateBackupVaultOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceConfigurationService.UpdateBackupVault(updateBackupVaultOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateBackupVault successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the BackupVaultActivityTracking model
				backupVaultActivityTrackingModel := new(resourceconfigurationv1.BackupVaultActivityTracking)
				backupVaultActivityTrackingModel.ManagementEvents = core.BoolPtr(true)

				// Construct an instance of the BackupVaultMetricsMonitoring model
				backupVaultMetricsMonitoringModel := new(resourceconfigurationv1.BackupVaultMetricsMonitoring)
				backupVaultMetricsMonitoringModel.UsageMetricsEnabled = core.BoolPtr(true)

				// Construct an instance of the BackupVaultPatch model
				backupVaultPatchModel := new(resourceconfigurationv1.BackupVaultPatch)
				backupVaultPatchModel.ActivityTracking = backupVaultActivityTrackingModel
				backupVaultPatchModel.MetricsMonitoring = backupVaultMetricsMonitoringModel
				backupVaultPatchModelAsPatch, asPatchErr := backupVaultPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateBackupVaultOptions model
				updateBackupVaultOptionsModel := new(resourceconfigurationv1.UpdateBackupVaultOptions)
				updateBackupVaultOptionsModel.BackupVaultName = core.StringPtr("testString")
				updateBackupVaultOptionsModel.BackupVaultPatch = backupVaultPatchModelAsPatch
				updateBackupVaultOptionsModel.IfMatch = core.StringPtr("testString")
				updateBackupVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceConfigurationService.UpdateBackupVault(updateBackupVaultOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteBackupVault(deleteBackupVaultOptions *DeleteBackupVaultOptions)`, func() {
		deleteBackupVaultPath := "/backup_vaults/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteBackupVaultPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteBackupVault successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := resourceConfigurationService.DeleteBackupVault(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteBackupVaultOptions model
				deleteBackupVaultOptionsModel := new(resourceconfigurationv1.DeleteBackupVaultOptions)
				deleteBackupVaultOptionsModel.BackupVaultName = core.StringPtr("testString")
				deleteBackupVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = resourceConfigurationService.DeleteBackupVault(deleteBackupVaultOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteBackupVault with error: Operation validation and request error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the DeleteBackupVaultOptions model
				deleteBackupVaultOptionsModel := new(resourceconfigurationv1.DeleteBackupVaultOptions)
				deleteBackupVaultOptionsModel.BackupVaultName = core.StringPtr("testString")
				deleteBackupVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := resourceConfigurationService.DeleteBackupVault(deleteBackupVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteBackupVaultOptions model with no property values
				deleteBackupVaultOptionsModelNew := new(resourceconfigurationv1.DeleteBackupVaultOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = resourceConfigurationService.DeleteBackupVault(deleteBackupVaultOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBucketConfig(getBucketConfigOptions *GetBucketConfigOptions) - Operation response error`, func() {
		getBucketConfigPath := "/b/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBucketConfigPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBucketConfig with error: Operation response processing error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetBucketConfigOptions model
				getBucketConfigOptionsModel := new(resourceconfigurationv1.GetBucketConfigOptions)
				getBucketConfigOptionsModel.Bucket = core.StringPtr("testString")
				getBucketConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceConfigurationService.GetBucketConfig(getBucketConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = resourceConfigurationService.GetBucketConfig(getBucketConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBucketConfig(getBucketConfigOptions *GetBucketConfigOptions)`, func() {
		getBucketConfigPath := "/b/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBucketConfigPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "my-new-bucket", "crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/3bf0d9003abfb5d29761c3e97696b71c:xxxxxxx-6c4f-4a62-a165-696756d63903:bucket:my-new-bucket", "service_instance_id": "d6f04d83-6c4f-4a62-a165-696756d63903", "service_instance_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/3bf0d9003abfb5d29761c3e97696b71c:xxxxxxx-6c4f-4a62-a165-696756d63903::", "time_created": "2018-03-26T16:23:36.980Z", "time_updated": "2018-10-17T19:29:10.117Z", "object_count": 764265234, "bytes_used": 28198745752445145, "noncurrent_object_count": 764265234, "noncurrent_bytes_used": 844239234, "delete_marker_count": 827201, "firewall": {"allowed_ip": ["AllowedIp"], "denied_ip": ["DeniedIp"], "allowed_network_type": ["public"]}, "activity_tracking": {"read_data_events": false, "write_data_events": false, "activity_tracker_crn": "ActivityTrackerCrn", "management_events": false}, "metrics_monitoring": {"usage_metrics_enabled": false, "request_metrics_enabled": false, "metrics_monitoring_crn": "MetricsMonitoringCrn"}, "hard_quota": 28198745752445146, "protection_management": {"token_applied_counter": "TokenAppliedCounter", "token_entries": [{"token_id": "TokenID", "token_expiration_time": "TokenExpirationTime", "token_reference_id": "TokenReferenceID", "applied_time": "AppliedTime", "invalidated_time": "InvalidatedTime", "expiration_time": "ExpirationTime", "shorten_retention_flag": true}]}}`)
				}))
			})
			It(`Invoke GetBucketConfig successfully with retries`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())
				resourceConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the GetBucketConfigOptions model
				getBucketConfigOptionsModel := new(resourceconfigurationv1.GetBucketConfigOptions)
				getBucketConfigOptionsModel.Bucket = core.StringPtr("testString")
				getBucketConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceConfigurationService.GetBucketConfigWithContext(ctx, getBucketConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceConfigurationService.DisableRetries()
				result, response, operationErr := resourceConfigurationService.GetBucketConfig(getBucketConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceConfigurationService.GetBucketConfigWithContext(ctx, getBucketConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBucketConfigPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "my-new-bucket", "crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/3bf0d9003abfb5d29761c3e97696b71c:xxxxxxx-6c4f-4a62-a165-696756d63903:bucket:my-new-bucket", "service_instance_id": "d6f04d83-6c4f-4a62-a165-696756d63903", "service_instance_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/3bf0d9003abfb5d29761c3e97696b71c:xxxxxxx-6c4f-4a62-a165-696756d63903::", "time_created": "2018-03-26T16:23:36.980Z", "time_updated": "2018-10-17T19:29:10.117Z", "object_count": 764265234, "bytes_used": 28198745752445145, "noncurrent_object_count": 764265234, "noncurrent_bytes_used": 844239234, "delete_marker_count": 827201, "firewall": {"allowed_ip": ["AllowedIp"], "denied_ip": ["DeniedIp"], "allowed_network_type": ["public"]}, "activity_tracking": {"read_data_events": false, "write_data_events": false, "activity_tracker_crn": "ActivityTrackerCrn", "management_events": false}, "metrics_monitoring": {"usage_metrics_enabled": false, "request_metrics_enabled": false, "metrics_monitoring_crn": "MetricsMonitoringCrn"}, "hard_quota": 28198745752445146, "protection_management": {"token_applied_counter": "TokenAppliedCounter", "token_entries": [{"token_id": "TokenID", "token_expiration_time": "TokenExpirationTime", "token_reference_id": "TokenReferenceID", "applied_time": "AppliedTime", "invalidated_time": "InvalidatedTime", "expiration_time": "ExpirationTime", "shorten_retention_flag": true}]}}`)
				}))
			})
			It(`Invoke GetBucketConfig successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceConfigurationService.GetBucketConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBucketConfigOptions model
				getBucketConfigOptionsModel := new(resourceconfigurationv1.GetBucketConfigOptions)
				getBucketConfigOptionsModel.Bucket = core.StringPtr("testString")
				getBucketConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceConfigurationService.GetBucketConfig(getBucketConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetBucketConfig with error: Operation validation and request error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetBucketConfigOptions model
				getBucketConfigOptionsModel := new(resourceconfigurationv1.GetBucketConfigOptions)
				getBucketConfigOptionsModel.Bucket = core.StringPtr("testString")
				getBucketConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceConfigurationService.GetBucketConfig(getBucketConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetBucketConfigOptions model with no property values
				getBucketConfigOptionsModelNew := new(resourceconfigurationv1.GetBucketConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceConfigurationService.GetBucketConfig(getBucketConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetBucketConfig successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetBucketConfigOptions model
				getBucketConfigOptionsModel := new(resourceconfigurationv1.GetBucketConfigOptions)
				getBucketConfigOptionsModel.Bucket = core.StringPtr("testString")
				getBucketConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceConfigurationService.GetBucketConfig(getBucketConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateBucketConfig(updateBucketConfigOptions *UpdateBucketConfigOptions)`, func() {
		updateBucketConfigPath := "/b/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateBucketConfigPath))
					Expect(req.Method).To(Equal("PATCH"))

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

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateBucketConfig successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := resourceConfigurationService.UpdateBucketConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the Firewall model
				firewallModel := new(resourceconfigurationv1.Firewall)
				firewallModel.AllowedIp = []string{"10.142.175.0/22", "10.198.243.79"}

				// Construct an instance of the ActivityTracking model
				activityTrackingModel := new(resourceconfigurationv1.ActivityTracking)
				activityTrackingModel.ReadDataEvents = core.BoolPtr(false)
				activityTrackingModel.WriteDataEvents = core.BoolPtr(false)
				activityTrackingModel.ActivityTrackerCrn = core.StringPtr("testString")
				activityTrackingModel.ManagementEvents = core.BoolPtr(false)

				// Construct an instance of the MetricsMonitoring model
				metricsMonitoringModel := new(resourceconfigurationv1.MetricsMonitoring)
				metricsMonitoringModel.UsageMetricsEnabled = core.BoolPtr(false)
				metricsMonitoringModel.RequestMetricsEnabled = core.BoolPtr(false)
				metricsMonitoringModel.MetricsMonitoringCrn = core.StringPtr("testString")

				// Construct an instance of the ProtectionManagement model
				protectionManagementModel := new(resourceconfigurationv1.ProtectionManagement)
				protectionManagementModel.RequestedState = core.StringPtr("activate")
				protectionManagementModel.ProtectionManagementToken = core.StringPtr("testString")

				// Construct an instance of the BucketPatch model
				bucketPatchModel := new(resourceconfigurationv1.BucketPatch)
				bucketPatchModel.Firewall = firewallModel
				bucketPatchModel.ActivityTracking = activityTrackingModel
				bucketPatchModel.MetricsMonitoring = metricsMonitoringModel
				bucketPatchModel.HardQuota = core.Int64Ptr(int64(28198745752445146))
				bucketPatchModel.ProtectionManagement = protectionManagementModel
				bucketPatchModelAsPatch, asPatchErr := bucketPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateBucketConfigOptions model
				updateBucketConfigOptionsModel := new(resourceconfigurationv1.UpdateBucketConfigOptions)
				updateBucketConfigOptionsModel.Bucket = core.StringPtr("testString")
				updateBucketConfigOptionsModel.BucketPatch = bucketPatchModelAsPatch
				updateBucketConfigOptionsModel.IfMatch = core.StringPtr("testString")
				updateBucketConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = resourceConfigurationService.UpdateBucketConfig(updateBucketConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateBucketConfig with error: Operation validation and request error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the Firewall model
				firewallModel := new(resourceconfigurationv1.Firewall)
				firewallModel.AllowedIp = []string{"10.142.175.0/22", "10.198.243.79"}

				// Construct an instance of the ActivityTracking model
				activityTrackingModel := new(resourceconfigurationv1.ActivityTracking)
				activityTrackingModel.ReadDataEvents = core.BoolPtr(false)
				activityTrackingModel.WriteDataEvents = core.BoolPtr(false)
				activityTrackingModel.ActivityTrackerCrn = core.StringPtr("testString")
				activityTrackingModel.ManagementEvents = core.BoolPtr(false)

				// Construct an instance of the MetricsMonitoring model
				metricsMonitoringModel := new(resourceconfigurationv1.MetricsMonitoring)
				metricsMonitoringModel.UsageMetricsEnabled = core.BoolPtr(false)
				metricsMonitoringModel.RequestMetricsEnabled = core.BoolPtr(false)
				metricsMonitoringModel.MetricsMonitoringCrn = core.StringPtr("testString")

				// Construct an instance of the ProtectionManagement model
				protectionManagementModel := new(resourceconfigurationv1.ProtectionManagement)
				protectionManagementModel.RequestedState = core.StringPtr("activate")
				protectionManagementModel.ProtectionManagementToken = core.StringPtr("testString")

				// Construct an instance of the BucketPatch model
				bucketPatchModel := new(resourceconfigurationv1.BucketPatch)
				bucketPatchModel.Firewall = firewallModel
				bucketPatchModel.ActivityTracking = activityTrackingModel
				bucketPatchModel.MetricsMonitoring = metricsMonitoringModel
				bucketPatchModel.HardQuota = core.Int64Ptr(int64(28198745752445146))
				bucketPatchModel.ProtectionManagement = protectionManagementModel
				bucketPatchModelAsPatch, asPatchErr := bucketPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateBucketConfigOptions model
				updateBucketConfigOptionsModel := new(resourceconfigurationv1.UpdateBucketConfigOptions)
				updateBucketConfigOptionsModel.Bucket = core.StringPtr("testString")
				updateBucketConfigOptionsModel.BucketPatch = bucketPatchModelAsPatch
				updateBucketConfigOptionsModel.IfMatch = core.StringPtr("testString")
				updateBucketConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := resourceConfigurationService.UpdateBucketConfig(updateBucketConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateBucketConfigOptions model with no property values
				updateBucketConfigOptionsModelNew := new(resourceconfigurationv1.UpdateBucketConfigOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = resourceConfigurationService.UpdateBucketConfig(updateBucketConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListRecoveryRanges(listRecoveryRangesOptions *ListRecoveryRangesOptions) - Operation response error`, func() {
		listRecoveryRangesPath := "/backup_vaults/testString/recovery_ranges"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRecoveryRangesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["source_resource_crn"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["latest"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["token"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListRecoveryRanges with error: Operation response processing error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListRecoveryRangesOptions model
				listRecoveryRangesOptionsModel := new(resourceconfigurationv1.ListRecoveryRangesOptions)
				listRecoveryRangesOptionsModel.BackupVaultName = core.StringPtr("testString")
				listRecoveryRangesOptionsModel.SourceResourceCrn = core.StringPtr("testString")
				listRecoveryRangesOptionsModel.Latest = core.StringPtr("testString")
				listRecoveryRangesOptionsModel.Token = core.StringPtr("testString")
				listRecoveryRangesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceConfigurationService.ListRecoveryRanges(listRecoveryRangesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = resourceConfigurationService.ListRecoveryRanges(listRecoveryRangesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListRecoveryRanges(listRecoveryRangesOptions *ListRecoveryRangesOptions)`, func() {
		listRecoveryRangesPath := "/backup_vaults/testString/recovery_ranges"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRecoveryRangesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["source_resource_crn"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["latest"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["token"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"next": {"href": "Href", "token": "Token"}, "recovery_ranges": [{"source_resource_crn": "SourceResourceCrn", "backup_policy_name": "myBackupPolicy", "range_start_time": "2019-01-01T12:00:00.000Z", "range_end_time": "2019-01-01T12:00:00.000Z", "range_create_time": "2019-01-01T12:00:00.000Z", "retention": {"delete_after_days": 10}, "recovery_range_id": "RecoveryRangeID"}]}`)
				}))
			})
			It(`Invoke ListRecoveryRanges successfully with retries`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())
				resourceConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the ListRecoveryRangesOptions model
				listRecoveryRangesOptionsModel := new(resourceconfigurationv1.ListRecoveryRangesOptions)
				listRecoveryRangesOptionsModel.BackupVaultName = core.StringPtr("testString")
				listRecoveryRangesOptionsModel.SourceResourceCrn = core.StringPtr("testString")
				listRecoveryRangesOptionsModel.Latest = core.StringPtr("testString")
				listRecoveryRangesOptionsModel.Token = core.StringPtr("testString")
				listRecoveryRangesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceConfigurationService.ListRecoveryRangesWithContext(ctx, listRecoveryRangesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceConfigurationService.DisableRetries()
				result, response, operationErr := resourceConfigurationService.ListRecoveryRanges(listRecoveryRangesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceConfigurationService.ListRecoveryRangesWithContext(ctx, listRecoveryRangesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRecoveryRangesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["source_resource_crn"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["latest"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["token"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"next": {"href": "Href", "token": "Token"}, "recovery_ranges": [{"source_resource_crn": "SourceResourceCrn", "backup_policy_name": "myBackupPolicy", "range_start_time": "2019-01-01T12:00:00.000Z", "range_end_time": "2019-01-01T12:00:00.000Z", "range_create_time": "2019-01-01T12:00:00.000Z", "retention": {"delete_after_days": 10}, "recovery_range_id": "RecoveryRangeID"}]}`)
				}))
			})
			It(`Invoke ListRecoveryRanges successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceConfigurationService.ListRecoveryRanges(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListRecoveryRangesOptions model
				listRecoveryRangesOptionsModel := new(resourceconfigurationv1.ListRecoveryRangesOptions)
				listRecoveryRangesOptionsModel.BackupVaultName = core.StringPtr("testString")
				listRecoveryRangesOptionsModel.SourceResourceCrn = core.StringPtr("testString")
				listRecoveryRangesOptionsModel.Latest = core.StringPtr("testString")
				listRecoveryRangesOptionsModel.Token = core.StringPtr("testString")
				listRecoveryRangesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceConfigurationService.ListRecoveryRanges(listRecoveryRangesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListRecoveryRanges with error: Operation validation and request error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListRecoveryRangesOptions model
				listRecoveryRangesOptionsModel := new(resourceconfigurationv1.ListRecoveryRangesOptions)
				listRecoveryRangesOptionsModel.BackupVaultName = core.StringPtr("testString")
				listRecoveryRangesOptionsModel.SourceResourceCrn = core.StringPtr("testString")
				listRecoveryRangesOptionsModel.Latest = core.StringPtr("testString")
				listRecoveryRangesOptionsModel.Token = core.StringPtr("testString")
				listRecoveryRangesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceConfigurationService.ListRecoveryRanges(listRecoveryRangesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListRecoveryRangesOptions model with no property values
				listRecoveryRangesOptionsModelNew := new(resourceconfigurationv1.ListRecoveryRangesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceConfigurationService.ListRecoveryRanges(listRecoveryRangesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListRecoveryRanges successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListRecoveryRangesOptions model
				listRecoveryRangesOptionsModel := new(resourceconfigurationv1.ListRecoveryRangesOptions)
				listRecoveryRangesOptionsModel.BackupVaultName = core.StringPtr("testString")
				listRecoveryRangesOptionsModel.SourceResourceCrn = core.StringPtr("testString")
				listRecoveryRangesOptionsModel.Latest = core.StringPtr("testString")
				listRecoveryRangesOptionsModel.Token = core.StringPtr("testString")
				listRecoveryRangesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceConfigurationService.ListRecoveryRanges(listRecoveryRangesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextToken successfully`, func() {
				responseObject := new(resourceconfigurationv1.RecoveryRangeCollection)
				nextObject := new(resourceconfigurationv1.NextPagination)
				nextObject.Token = core.StringPtr("abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextToken()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextToken without a "Next" property in the response`, func() {
				responseObject := new(resourceconfigurationv1.RecoveryRangeCollection)

				value, err := responseObject.GetNextToken()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRecoveryRangesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"token":"1"},"total_count":2,"recovery_ranges":[{"source_resource_crn":"SourceResourceCrn","backup_policy_name":"myBackupPolicy","range_start_time":"2019-01-01T12:00:00.000Z","range_end_time":"2019-01-01T12:00:00.000Z","range_create_time":"2019-01-01T12:00:00.000Z","retention":{"delete_after_days":10},"recovery_range_id":"RecoveryRangeID"}],"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"recovery_ranges":[{"source_resource_crn":"SourceResourceCrn","backup_policy_name":"myBackupPolicy","range_start_time":"2019-01-01T12:00:00.000Z","range_end_time":"2019-01-01T12:00:00.000Z","range_create_time":"2019-01-01T12:00:00.000Z","retention":{"delete_after_days":10},"recovery_range_id":"RecoveryRangeID"}],"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use RecoveryRangesPager.GetNext successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				listRecoveryRangesOptionsModel := &resourceconfigurationv1.ListRecoveryRangesOptions{
					BackupVaultName:   core.StringPtr("testString"),
					SourceResourceCrn: core.StringPtr("testString"),
					Latest:            core.StringPtr("testString"),
				}

				pager, err := resourceConfigurationService.NewRecoveryRangesPager(listRecoveryRangesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []resourceconfigurationv1.RecoveryRange
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use RecoveryRangesPager.GetAll successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				listRecoveryRangesOptionsModel := &resourceconfigurationv1.ListRecoveryRangesOptions{
					BackupVaultName:   core.StringPtr("testString"),
					SourceResourceCrn: core.StringPtr("testString"),
					Latest:            core.StringPtr("testString"),
				}

				pager, err := resourceConfigurationService.NewRecoveryRangesPager(listRecoveryRangesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`GetSourceResourceRecoveryRange(getSourceResourceRecoveryRangeOptions *GetSourceResourceRecoveryRangeOptions) - Operation response error`, func() {
		getSourceResourceRecoveryRangePath := "/backup_vaults/testString/recovery_ranges/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSourceResourceRecoveryRangePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSourceResourceRecoveryRange with error: Operation response processing error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetSourceResourceRecoveryRangeOptions model
				getSourceResourceRecoveryRangeOptionsModel := new(resourceconfigurationv1.GetSourceResourceRecoveryRangeOptions)
				getSourceResourceRecoveryRangeOptionsModel.BackupVaultName = core.StringPtr("testString")
				getSourceResourceRecoveryRangeOptionsModel.RecoveryRangeID = core.StringPtr("testString")
				getSourceResourceRecoveryRangeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceConfigurationService.GetSourceResourceRecoveryRange(getSourceResourceRecoveryRangeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = resourceConfigurationService.GetSourceResourceRecoveryRange(getSourceResourceRecoveryRangeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSourceResourceRecoveryRange(getSourceResourceRecoveryRangeOptions *GetSourceResourceRecoveryRangeOptions)`, func() {
		getSourceResourceRecoveryRangePath := "/backup_vaults/testString/recovery_ranges/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSourceResourceRecoveryRangePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"source_resource_crn": "SourceResourceCrn", "backup_policy_name": "myBackupPolicy", "range_start_time": "2019-01-01T12:00:00.000Z", "range_end_time": "2019-01-01T12:00:00.000Z", "range_create_time": "2019-01-01T12:00:00.000Z", "retention": {"delete_after_days": 10}, "recovery_range_id": "RecoveryRangeID"}`)
				}))
			})
			It(`Invoke GetSourceResourceRecoveryRange successfully with retries`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())
				resourceConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the GetSourceResourceRecoveryRangeOptions model
				getSourceResourceRecoveryRangeOptionsModel := new(resourceconfigurationv1.GetSourceResourceRecoveryRangeOptions)
				getSourceResourceRecoveryRangeOptionsModel.BackupVaultName = core.StringPtr("testString")
				getSourceResourceRecoveryRangeOptionsModel.RecoveryRangeID = core.StringPtr("testString")
				getSourceResourceRecoveryRangeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceConfigurationService.GetSourceResourceRecoveryRangeWithContext(ctx, getSourceResourceRecoveryRangeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceConfigurationService.DisableRetries()
				result, response, operationErr := resourceConfigurationService.GetSourceResourceRecoveryRange(getSourceResourceRecoveryRangeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceConfigurationService.GetSourceResourceRecoveryRangeWithContext(ctx, getSourceResourceRecoveryRangeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSourceResourceRecoveryRangePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"source_resource_crn": "SourceResourceCrn", "backup_policy_name": "myBackupPolicy", "range_start_time": "2019-01-01T12:00:00.000Z", "range_end_time": "2019-01-01T12:00:00.000Z", "range_create_time": "2019-01-01T12:00:00.000Z", "retention": {"delete_after_days": 10}, "recovery_range_id": "RecoveryRangeID"}`)
				}))
			})
			It(`Invoke GetSourceResourceRecoveryRange successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceConfigurationService.GetSourceResourceRecoveryRange(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSourceResourceRecoveryRangeOptions model
				getSourceResourceRecoveryRangeOptionsModel := new(resourceconfigurationv1.GetSourceResourceRecoveryRangeOptions)
				getSourceResourceRecoveryRangeOptionsModel.BackupVaultName = core.StringPtr("testString")
				getSourceResourceRecoveryRangeOptionsModel.RecoveryRangeID = core.StringPtr("testString")
				getSourceResourceRecoveryRangeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceConfigurationService.GetSourceResourceRecoveryRange(getSourceResourceRecoveryRangeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSourceResourceRecoveryRange with error: Operation validation and request error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetSourceResourceRecoveryRangeOptions model
				getSourceResourceRecoveryRangeOptionsModel := new(resourceconfigurationv1.GetSourceResourceRecoveryRangeOptions)
				getSourceResourceRecoveryRangeOptionsModel.BackupVaultName = core.StringPtr("testString")
				getSourceResourceRecoveryRangeOptionsModel.RecoveryRangeID = core.StringPtr("testString")
				getSourceResourceRecoveryRangeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceConfigurationService.GetSourceResourceRecoveryRange(getSourceResourceRecoveryRangeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSourceResourceRecoveryRangeOptions model with no property values
				getSourceResourceRecoveryRangeOptionsModelNew := new(resourceconfigurationv1.GetSourceResourceRecoveryRangeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceConfigurationService.GetSourceResourceRecoveryRange(getSourceResourceRecoveryRangeOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetSourceResourceRecoveryRange successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetSourceResourceRecoveryRangeOptions model
				getSourceResourceRecoveryRangeOptionsModel := new(resourceconfigurationv1.GetSourceResourceRecoveryRangeOptions)
				getSourceResourceRecoveryRangeOptionsModel.BackupVaultName = core.StringPtr("testString")
				getSourceResourceRecoveryRangeOptionsModel.RecoveryRangeID = core.StringPtr("testString")
				getSourceResourceRecoveryRangeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceConfigurationService.GetSourceResourceRecoveryRange(getSourceResourceRecoveryRangeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PatchSourceResourceRecoveryRange(patchSourceResourceRecoveryRangeOptions *PatchSourceResourceRecoveryRangeOptions) - Operation response error`, func() {
		patchSourceResourceRecoveryRangePath := "/backup_vaults/testString/recovery_ranges/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(patchSourceResourceRecoveryRangePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PatchSourceResourceRecoveryRange with error: Operation response processing error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the DeleteAfterDays model
				deleteAfterDaysModel := new(resourceconfigurationv1.DeleteAfterDays)
				deleteAfterDaysModel.DeleteAfterDays = core.Int64Ptr(int64(10))

				// Construct an instance of the RecoveryRangePatch model
				recoveryRangePatchModel := new(resourceconfigurationv1.RecoveryRangePatch)
				recoveryRangePatchModel.Retention = deleteAfterDaysModel
				recoveryRangePatchModelAsPatch, asPatchErr := recoveryRangePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the PatchSourceResourceRecoveryRangeOptions model
				patchSourceResourceRecoveryRangeOptionsModel := new(resourceconfigurationv1.PatchSourceResourceRecoveryRangeOptions)
				patchSourceResourceRecoveryRangeOptionsModel.BackupVaultName = core.StringPtr("testString")
				patchSourceResourceRecoveryRangeOptionsModel.RecoveryRangeID = core.StringPtr("testString")
				patchSourceResourceRecoveryRangeOptionsModel.RecoveryRangePatch = recoveryRangePatchModelAsPatch
				patchSourceResourceRecoveryRangeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceConfigurationService.PatchSourceResourceRecoveryRange(patchSourceResourceRecoveryRangeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = resourceConfigurationService.PatchSourceResourceRecoveryRange(patchSourceResourceRecoveryRangeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PatchSourceResourceRecoveryRange(patchSourceResourceRecoveryRangeOptions *PatchSourceResourceRecoveryRangeOptions)`, func() {
		patchSourceResourceRecoveryRangePath := "/backup_vaults/testString/recovery_ranges/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(patchSourceResourceRecoveryRangePath))
					Expect(req.Method).To(Equal("PATCH"))

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
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"source_resource_crn": "SourceResourceCrn", "backup_policy_name": "myBackupPolicy", "range_start_time": "2019-01-01T12:00:00.000Z", "range_end_time": "2019-01-01T12:00:00.000Z", "range_create_time": "2019-01-01T12:00:00.000Z", "retention": {"delete_after_days": 10}, "recovery_range_id": "RecoveryRangeID"}`)
				}))
			})
			It(`Invoke PatchSourceResourceRecoveryRange successfully with retries`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())
				resourceConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the DeleteAfterDays model
				deleteAfterDaysModel := new(resourceconfigurationv1.DeleteAfterDays)
				deleteAfterDaysModel.DeleteAfterDays = core.Int64Ptr(int64(10))

				// Construct an instance of the RecoveryRangePatch model
				recoveryRangePatchModel := new(resourceconfigurationv1.RecoveryRangePatch)
				recoveryRangePatchModel.Retention = deleteAfterDaysModel
				recoveryRangePatchModelAsPatch, asPatchErr := recoveryRangePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the PatchSourceResourceRecoveryRangeOptions model
				patchSourceResourceRecoveryRangeOptionsModel := new(resourceconfigurationv1.PatchSourceResourceRecoveryRangeOptions)
				patchSourceResourceRecoveryRangeOptionsModel.BackupVaultName = core.StringPtr("testString")
				patchSourceResourceRecoveryRangeOptionsModel.RecoveryRangeID = core.StringPtr("testString")
				patchSourceResourceRecoveryRangeOptionsModel.RecoveryRangePatch = recoveryRangePatchModelAsPatch
				patchSourceResourceRecoveryRangeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceConfigurationService.PatchSourceResourceRecoveryRangeWithContext(ctx, patchSourceResourceRecoveryRangeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceConfigurationService.DisableRetries()
				result, response, operationErr := resourceConfigurationService.PatchSourceResourceRecoveryRange(patchSourceResourceRecoveryRangeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceConfigurationService.PatchSourceResourceRecoveryRangeWithContext(ctx, patchSourceResourceRecoveryRangeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(patchSourceResourceRecoveryRangePath))
					Expect(req.Method).To(Equal("PATCH"))

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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"source_resource_crn": "SourceResourceCrn", "backup_policy_name": "myBackupPolicy", "range_start_time": "2019-01-01T12:00:00.000Z", "range_end_time": "2019-01-01T12:00:00.000Z", "range_create_time": "2019-01-01T12:00:00.000Z", "retention": {"delete_after_days": 10}, "recovery_range_id": "RecoveryRangeID"}`)
				}))
			})
			It(`Invoke PatchSourceResourceRecoveryRange successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceConfigurationService.PatchSourceResourceRecoveryRange(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteAfterDays model
				deleteAfterDaysModel := new(resourceconfigurationv1.DeleteAfterDays)
				deleteAfterDaysModel.DeleteAfterDays = core.Int64Ptr(int64(10))

				// Construct an instance of the RecoveryRangePatch model
				recoveryRangePatchModel := new(resourceconfigurationv1.RecoveryRangePatch)
				recoveryRangePatchModel.Retention = deleteAfterDaysModel
				recoveryRangePatchModelAsPatch, asPatchErr := recoveryRangePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the PatchSourceResourceRecoveryRangeOptions model
				patchSourceResourceRecoveryRangeOptionsModel := new(resourceconfigurationv1.PatchSourceResourceRecoveryRangeOptions)
				patchSourceResourceRecoveryRangeOptionsModel.BackupVaultName = core.StringPtr("testString")
				patchSourceResourceRecoveryRangeOptionsModel.RecoveryRangeID = core.StringPtr("testString")
				patchSourceResourceRecoveryRangeOptionsModel.RecoveryRangePatch = recoveryRangePatchModelAsPatch
				patchSourceResourceRecoveryRangeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceConfigurationService.PatchSourceResourceRecoveryRange(patchSourceResourceRecoveryRangeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PatchSourceResourceRecoveryRange with error: Operation validation and request error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the DeleteAfterDays model
				deleteAfterDaysModel := new(resourceconfigurationv1.DeleteAfterDays)
				deleteAfterDaysModel.DeleteAfterDays = core.Int64Ptr(int64(10))

				// Construct an instance of the RecoveryRangePatch model
				recoveryRangePatchModel := new(resourceconfigurationv1.RecoveryRangePatch)
				recoveryRangePatchModel.Retention = deleteAfterDaysModel
				recoveryRangePatchModelAsPatch, asPatchErr := recoveryRangePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the PatchSourceResourceRecoveryRangeOptions model
				patchSourceResourceRecoveryRangeOptionsModel := new(resourceconfigurationv1.PatchSourceResourceRecoveryRangeOptions)
				patchSourceResourceRecoveryRangeOptionsModel.BackupVaultName = core.StringPtr("testString")
				patchSourceResourceRecoveryRangeOptionsModel.RecoveryRangeID = core.StringPtr("testString")
				patchSourceResourceRecoveryRangeOptionsModel.RecoveryRangePatch = recoveryRangePatchModelAsPatch
				patchSourceResourceRecoveryRangeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceConfigurationService.PatchSourceResourceRecoveryRange(patchSourceResourceRecoveryRangeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PatchSourceResourceRecoveryRangeOptions model with no property values
				patchSourceResourceRecoveryRangeOptionsModelNew := new(resourceconfigurationv1.PatchSourceResourceRecoveryRangeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceConfigurationService.PatchSourceResourceRecoveryRange(patchSourceResourceRecoveryRangeOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke PatchSourceResourceRecoveryRange successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the DeleteAfterDays model
				deleteAfterDaysModel := new(resourceconfigurationv1.DeleteAfterDays)
				deleteAfterDaysModel.DeleteAfterDays = core.Int64Ptr(int64(10))

				// Construct an instance of the RecoveryRangePatch model
				recoveryRangePatchModel := new(resourceconfigurationv1.RecoveryRangePatch)
				recoveryRangePatchModel.Retention = deleteAfterDaysModel
				recoveryRangePatchModelAsPatch, asPatchErr := recoveryRangePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the PatchSourceResourceRecoveryRangeOptions model
				patchSourceResourceRecoveryRangeOptionsModel := new(resourceconfigurationv1.PatchSourceResourceRecoveryRangeOptions)
				patchSourceResourceRecoveryRangeOptionsModel.BackupVaultName = core.StringPtr("testString")
				patchSourceResourceRecoveryRangeOptionsModel.RecoveryRangeID = core.StringPtr("testString")
				patchSourceResourceRecoveryRangeOptionsModel.RecoveryRangePatch = recoveryRangePatchModelAsPatch
				patchSourceResourceRecoveryRangeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceConfigurationService.PatchSourceResourceRecoveryRange(patchSourceResourceRecoveryRangeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateRestore(createRestoreOptions *CreateRestoreOptions) - Operation response error`, func() {
		createRestorePath := "/backup_vaults/testString/restores"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRestorePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateRestore with error: Operation response processing error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the CreateRestoreOptions model
				createRestoreOptionsModel := new(resourceconfigurationv1.CreateRestoreOptions)
				createRestoreOptionsModel.BackupVaultName = core.StringPtr("testString")
				createRestoreOptionsModel.RecoveryRangeID = core.StringPtr("6ff0d31c-7583-4463-8ae5-208752f5769c")
				createRestoreOptionsModel.RestoreType = core.StringPtr("in_place")
				createRestoreOptionsModel.RestorePointInTime = CreateMockDateTime("2024-06-04T12:12:00.000Z")
				createRestoreOptionsModel.TargetResourceCrn = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a1229395:8dfbcba4e6a740e3866020847e525436:bucket:targetbucket")
				createRestoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceConfigurationService.CreateRestore(createRestoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = resourceConfigurationService.CreateRestore(createRestoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateRestore(createRestoreOptions *CreateRestoreOptions)`, func() {
		createRestorePath := "/backup_vaults/testString/restores"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRestorePath))
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
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"recovery_range_id": "RecoveryRangeID", "restore_type": "in_place", "restore_point_in_time": "2019-01-01T12:00:00.000Z", "target_resource_crn": "TargetResourceCrn", "source_resource_crn": "SourceResourceCrn", "restore_id": "RestoreID", "restore_status": "initializing", "init_time": "2019-01-01T12:00:00.000Z", "complete_time": "2019-01-01T12:00:00.000Z", "restore_percent_progress": 1, "error_cause": "ErrorCause"}`)
				}))
			})
			It(`Invoke CreateRestore successfully with retries`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())
				resourceConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the CreateRestoreOptions model
				createRestoreOptionsModel := new(resourceconfigurationv1.CreateRestoreOptions)
				createRestoreOptionsModel.BackupVaultName = core.StringPtr("testString")
				createRestoreOptionsModel.RecoveryRangeID = core.StringPtr("6ff0d31c-7583-4463-8ae5-208752f5769c")
				createRestoreOptionsModel.RestoreType = core.StringPtr("in_place")
				createRestoreOptionsModel.RestorePointInTime = CreateMockDateTime("2024-06-04T12:12:00.000Z")
				createRestoreOptionsModel.TargetResourceCrn = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a1229395:8dfbcba4e6a740e3866020847e525436:bucket:targetbucket")
				createRestoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceConfigurationService.CreateRestoreWithContext(ctx, createRestoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceConfigurationService.DisableRetries()
				result, response, operationErr := resourceConfigurationService.CreateRestore(createRestoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceConfigurationService.CreateRestoreWithContext(ctx, createRestoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRestorePath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"recovery_range_id": "RecoveryRangeID", "restore_type": "in_place", "restore_point_in_time": "2019-01-01T12:00:00.000Z", "target_resource_crn": "TargetResourceCrn", "source_resource_crn": "SourceResourceCrn", "restore_id": "RestoreID", "restore_status": "initializing", "init_time": "2019-01-01T12:00:00.000Z", "complete_time": "2019-01-01T12:00:00.000Z", "restore_percent_progress": 1, "error_cause": "ErrorCause"}`)
				}))
			})
			It(`Invoke CreateRestore successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceConfigurationService.CreateRestore(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateRestoreOptions model
				createRestoreOptionsModel := new(resourceconfigurationv1.CreateRestoreOptions)
				createRestoreOptionsModel.BackupVaultName = core.StringPtr("testString")
				createRestoreOptionsModel.RecoveryRangeID = core.StringPtr("6ff0d31c-7583-4463-8ae5-208752f5769c")
				createRestoreOptionsModel.RestoreType = core.StringPtr("in_place")
				createRestoreOptionsModel.RestorePointInTime = CreateMockDateTime("2024-06-04T12:12:00.000Z")
				createRestoreOptionsModel.TargetResourceCrn = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a1229395:8dfbcba4e6a740e3866020847e525436:bucket:targetbucket")
				createRestoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceConfigurationService.CreateRestore(createRestoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateRestore with error: Operation validation and request error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the CreateRestoreOptions model
				createRestoreOptionsModel := new(resourceconfigurationv1.CreateRestoreOptions)
				createRestoreOptionsModel.BackupVaultName = core.StringPtr("testString")
				createRestoreOptionsModel.RecoveryRangeID = core.StringPtr("6ff0d31c-7583-4463-8ae5-208752f5769c")
				createRestoreOptionsModel.RestoreType = core.StringPtr("in_place")
				createRestoreOptionsModel.RestorePointInTime = CreateMockDateTime("2024-06-04T12:12:00.000Z")
				createRestoreOptionsModel.TargetResourceCrn = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a1229395:8dfbcba4e6a740e3866020847e525436:bucket:targetbucket")
				createRestoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceConfigurationService.CreateRestore(createRestoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateRestoreOptions model with no property values
				createRestoreOptionsModelNew := new(resourceconfigurationv1.CreateRestoreOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceConfigurationService.CreateRestore(createRestoreOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateRestore successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the CreateRestoreOptions model
				createRestoreOptionsModel := new(resourceconfigurationv1.CreateRestoreOptions)
				createRestoreOptionsModel.BackupVaultName = core.StringPtr("testString")
				createRestoreOptionsModel.RecoveryRangeID = core.StringPtr("6ff0d31c-7583-4463-8ae5-208752f5769c")
				createRestoreOptionsModel.RestoreType = core.StringPtr("in_place")
				createRestoreOptionsModel.RestorePointInTime = CreateMockDateTime("2024-06-04T12:12:00.000Z")
				createRestoreOptionsModel.TargetResourceCrn = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a1229395:8dfbcba4e6a740e3866020847e525436:bucket:targetbucket")
				createRestoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceConfigurationService.CreateRestore(createRestoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListRestores(listRestoresOptions *ListRestoresOptions) - Operation response error`, func() {
		listRestoresPath := "/backup_vaults/testString/restores"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRestoresPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["token"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListRestores with error: Operation response processing error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListRestoresOptions model
				listRestoresOptionsModel := new(resourceconfigurationv1.ListRestoresOptions)
				listRestoresOptionsModel.BackupVaultName = core.StringPtr("testString")
				listRestoresOptionsModel.Token = core.StringPtr("testString")
				listRestoresOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceConfigurationService.ListRestores(listRestoresOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = resourceConfigurationService.ListRestores(listRestoresOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListRestores(listRestoresOptions *ListRestoresOptions)`, func() {
		listRestoresPath := "/backup_vaults/testString/restores"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRestoresPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["token"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"next": {"href": "Href", "token": "Token"}, "restores": [{"recovery_range_id": "RecoveryRangeID", "restore_type": "in_place", "restore_point_in_time": "2019-01-01T12:00:00.000Z", "target_resource_crn": "TargetResourceCrn", "source_resource_crn": "SourceResourceCrn", "restore_id": "RestoreID", "restore_status": "initializing", "init_time": "2019-01-01T12:00:00.000Z", "complete_time": "2019-01-01T12:00:00.000Z", "restore_percent_progress": 1, "error_cause": "ErrorCause"}]}`)
				}))
			})
			It(`Invoke ListRestores successfully with retries`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())
				resourceConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the ListRestoresOptions model
				listRestoresOptionsModel := new(resourceconfigurationv1.ListRestoresOptions)
				listRestoresOptionsModel.BackupVaultName = core.StringPtr("testString")
				listRestoresOptionsModel.Token = core.StringPtr("testString")
				listRestoresOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceConfigurationService.ListRestoresWithContext(ctx, listRestoresOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceConfigurationService.DisableRetries()
				result, response, operationErr := resourceConfigurationService.ListRestores(listRestoresOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceConfigurationService.ListRestoresWithContext(ctx, listRestoresOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRestoresPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["token"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"next": {"href": "Href", "token": "Token"}, "restores": [{"recovery_range_id": "RecoveryRangeID", "restore_type": "in_place", "restore_point_in_time": "2019-01-01T12:00:00.000Z", "target_resource_crn": "TargetResourceCrn", "source_resource_crn": "SourceResourceCrn", "restore_id": "RestoreID", "restore_status": "initializing", "init_time": "2019-01-01T12:00:00.000Z", "complete_time": "2019-01-01T12:00:00.000Z", "restore_percent_progress": 1, "error_cause": "ErrorCause"}]}`)
				}))
			})
			It(`Invoke ListRestores successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceConfigurationService.ListRestores(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListRestoresOptions model
				listRestoresOptionsModel := new(resourceconfigurationv1.ListRestoresOptions)
				listRestoresOptionsModel.BackupVaultName = core.StringPtr("testString")
				listRestoresOptionsModel.Token = core.StringPtr("testString")
				listRestoresOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceConfigurationService.ListRestores(listRestoresOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListRestores with error: Operation validation and request error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListRestoresOptions model
				listRestoresOptionsModel := new(resourceconfigurationv1.ListRestoresOptions)
				listRestoresOptionsModel.BackupVaultName = core.StringPtr("testString")
				listRestoresOptionsModel.Token = core.StringPtr("testString")
				listRestoresOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceConfigurationService.ListRestores(listRestoresOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListRestoresOptions model with no property values
				listRestoresOptionsModelNew := new(resourceconfigurationv1.ListRestoresOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceConfigurationService.ListRestores(listRestoresOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListRestores successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the ListRestoresOptions model
				listRestoresOptionsModel := new(resourceconfigurationv1.ListRestoresOptions)
				listRestoresOptionsModel.BackupVaultName = core.StringPtr("testString")
				listRestoresOptionsModel.Token = core.StringPtr("testString")
				listRestoresOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceConfigurationService.ListRestores(listRestoresOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextToken successfully`, func() {
				responseObject := new(resourceconfigurationv1.RestoreCollection)
				nextObject := new(resourceconfigurationv1.NextPagination)
				nextObject.Token = core.StringPtr("abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextToken()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextToken without a "Next" property in the response`, func() {
				responseObject := new(resourceconfigurationv1.RestoreCollection)

				value, err := responseObject.GetNextToken()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRestoresPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"token":"1"},"total_count":2,"limit":1,"restores":[{"recovery_range_id":"RecoveryRangeID","restore_type":"in_place","restore_point_in_time":"2019-01-01T12:00:00.000Z","target_resource_crn":"TargetResourceCrn","source_resource_crn":"SourceResourceCrn","restore_id":"RestoreID","restore_status":"initializing","init_time":"2019-01-01T12:00:00.000Z","complete_time":"2019-01-01T12:00:00.000Z","restore_percent_progress":1,"error_cause":"ErrorCause"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"restores":[{"recovery_range_id":"RecoveryRangeID","restore_type":"in_place","restore_point_in_time":"2019-01-01T12:00:00.000Z","target_resource_crn":"TargetResourceCrn","source_resource_crn":"SourceResourceCrn","restore_id":"RestoreID","restore_status":"initializing","init_time":"2019-01-01T12:00:00.000Z","complete_time":"2019-01-01T12:00:00.000Z","restore_percent_progress":1,"error_cause":"ErrorCause"}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use RestoresPager.GetNext successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				listRestoresOptionsModel := &resourceconfigurationv1.ListRestoresOptions{
					BackupVaultName: core.StringPtr("testString"),
				}

				pager, err := resourceConfigurationService.NewRestoresPager(listRestoresOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []resourceconfigurationv1.Restore
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use RestoresPager.GetAll successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				listRestoresOptionsModel := &resourceconfigurationv1.ListRestoresOptions{
					BackupVaultName: core.StringPtr("testString"),
				}

				pager, err := resourceConfigurationService.NewRestoresPager(listRestoresOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`GetRestore(getRestoreOptions *GetRestoreOptions) - Operation response error`, func() {
		getRestorePath := "/backup_vaults/testString/restores/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRestorePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetRestore with error: Operation response processing error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetRestoreOptions model
				getRestoreOptionsModel := new(resourceconfigurationv1.GetRestoreOptions)
				getRestoreOptionsModel.BackupVaultName = core.StringPtr("testString")
				getRestoreOptionsModel.RestoreID = core.StringPtr("testString")
				getRestoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceConfigurationService.GetRestore(getRestoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceConfigurationService.EnableRetries(0, 0)
				result, response, operationErr = resourceConfigurationService.GetRestore(getRestoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetRestore(getRestoreOptions *GetRestoreOptions)`, func() {
		getRestorePath := "/backup_vaults/testString/restores/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRestorePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"recovery_range_id": "RecoveryRangeID", "restore_type": "in_place", "restore_point_in_time": "2019-01-01T12:00:00.000Z", "target_resource_crn": "TargetResourceCrn", "source_resource_crn": "SourceResourceCrn", "restore_id": "RestoreID", "restore_status": "initializing", "init_time": "2019-01-01T12:00:00.000Z", "complete_time": "2019-01-01T12:00:00.000Z", "restore_percent_progress": 1, "error_cause": "ErrorCause"}`)
				}))
			})
			It(`Invoke GetRestore successfully with retries`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())
				resourceConfigurationService.EnableRetries(0, 0)

				// Construct an instance of the GetRestoreOptions model
				getRestoreOptionsModel := new(resourceconfigurationv1.GetRestoreOptions)
				getRestoreOptionsModel.BackupVaultName = core.StringPtr("testString")
				getRestoreOptionsModel.RestoreID = core.StringPtr("testString")
				getRestoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := resourceConfigurationService.GetRestoreWithContext(ctx, getRestoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				resourceConfigurationService.DisableRetries()
				result, response, operationErr := resourceConfigurationService.GetRestore(getRestoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = resourceConfigurationService.GetRestoreWithContext(ctx, getRestoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRestorePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"recovery_range_id": "RecoveryRangeID", "restore_type": "in_place", "restore_point_in_time": "2019-01-01T12:00:00.000Z", "target_resource_crn": "TargetResourceCrn", "source_resource_crn": "SourceResourceCrn", "restore_id": "RestoreID", "restore_status": "initializing", "init_time": "2019-01-01T12:00:00.000Z", "complete_time": "2019-01-01T12:00:00.000Z", "restore_percent_progress": 1, "error_cause": "ErrorCause"}`)
				}))
			})
			It(`Invoke GetRestore successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceConfigurationService.GetRestore(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRestoreOptions model
				getRestoreOptionsModel := new(resourceconfigurationv1.GetRestoreOptions)
				getRestoreOptionsModel.BackupVaultName = core.StringPtr("testString")
				getRestoreOptionsModel.RestoreID = core.StringPtr("testString")
				getRestoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceConfigurationService.GetRestore(getRestoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetRestore with error: Operation validation and request error`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetRestoreOptions model
				getRestoreOptionsModel := new(resourceconfigurationv1.GetRestoreOptions)
				getRestoreOptionsModel.BackupVaultName = core.StringPtr("testString")
				getRestoreOptionsModel.RestoreID = core.StringPtr("testString")
				getRestoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceConfigurationService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceConfigurationService.GetRestore(getRestoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetRestoreOptions model with no property values
				getRestoreOptionsModelNew := new(resourceconfigurationv1.GetRestoreOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceConfigurationService.GetRestore(getRestoreOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetRestore successfully`, func() {
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceConfigurationService).ToNot(BeNil())

				// Construct an instance of the GetRestoreOptions model
				getRestoreOptionsModel := new(resourceconfigurationv1.GetRestoreOptions)
				getRestoreOptionsModel.BackupVaultName = core.StringPtr("testString")
				getRestoreOptionsModel.RestoreID = core.StringPtr("testString")
				getRestoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := resourceConfigurationService.GetRestore(getRestoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			resourceConfigurationService, _ := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
				URL:           "http://resourceconfigurationv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateBackupPolicyOptions successfully`, func() {
				// Construct an instance of the DeleteAfterDays model
				deleteAfterDaysModel := new(resourceconfigurationv1.DeleteAfterDays)
				Expect(deleteAfterDaysModel).ToNot(BeNil())
				deleteAfterDaysModel.DeleteAfterDays = core.Int64Ptr(int64(10))
				Expect(deleteAfterDaysModel.DeleteAfterDays).To(Equal(core.Int64Ptr(int64(10))))

				// Construct an instance of the CreateBackupPolicyOptions model
				bucket := "testString"
				var createBackupPolicyOptionsInitialRetention *resourceconfigurationv1.DeleteAfterDays = nil
				createBackupPolicyOptionsPolicyName := "myBackupPolicy"
				createBackupPolicyOptionsTargetBackupVaultCrn := "testString"
				createBackupPolicyOptionsBackupType := "continuous"
				createBackupPolicyOptionsModel := resourceConfigurationService.NewCreateBackupPolicyOptions(bucket, createBackupPolicyOptionsInitialRetention, createBackupPolicyOptionsPolicyName, createBackupPolicyOptionsTargetBackupVaultCrn, createBackupPolicyOptionsBackupType)
				createBackupPolicyOptionsModel.SetBucket("testString")
				createBackupPolicyOptionsModel.SetInitialRetention(deleteAfterDaysModel)
				createBackupPolicyOptionsModel.SetPolicyName("myBackupPolicy")
				createBackupPolicyOptionsModel.SetTargetBackupVaultCrn("testString")
				createBackupPolicyOptionsModel.SetBackupType("continuous")
				createBackupPolicyOptionsModel.SetMD5("testString")
				createBackupPolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createBackupPolicyOptionsModel).ToNot(BeNil())
				Expect(createBackupPolicyOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(createBackupPolicyOptionsModel.InitialRetention).To(Equal(deleteAfterDaysModel))
				Expect(createBackupPolicyOptionsModel.PolicyName).To(Equal(core.StringPtr("myBackupPolicy")))
				Expect(createBackupPolicyOptionsModel.TargetBackupVaultCrn).To(Equal(core.StringPtr("testString")))
				Expect(createBackupPolicyOptionsModel.BackupType).To(Equal(core.StringPtr("continuous")))
				Expect(createBackupPolicyOptionsModel.MD5).To(Equal(core.StringPtr("testString")))
				Expect(createBackupPolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateBackupVaultOptions successfully`, func() {
				// Construct an instance of the BackupVaultActivityTracking model
				backupVaultActivityTrackingModel := new(resourceconfigurationv1.BackupVaultActivityTracking)
				Expect(backupVaultActivityTrackingModel).ToNot(BeNil())
				backupVaultActivityTrackingModel.ManagementEvents = core.BoolPtr(true)
				Expect(backupVaultActivityTrackingModel.ManagementEvents).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the BackupVaultMetricsMonitoring model
				backupVaultMetricsMonitoringModel := new(resourceconfigurationv1.BackupVaultMetricsMonitoring)
				Expect(backupVaultMetricsMonitoringModel).ToNot(BeNil())
				backupVaultMetricsMonitoringModel.UsageMetricsEnabled = core.BoolPtr(true)
				Expect(backupVaultMetricsMonitoringModel.UsageMetricsEnabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the CreateBackupVaultOptions model
				serviceInstanceID := "testString"
				createBackupVaultOptionsBackupVaultName := "testString"
				createBackupVaultOptionsRegion := "testString"
				createBackupVaultOptionsModel := resourceConfigurationService.NewCreateBackupVaultOptions(serviceInstanceID, createBackupVaultOptionsBackupVaultName, createBackupVaultOptionsRegion)
				createBackupVaultOptionsModel.SetServiceInstanceID("testString")
				createBackupVaultOptionsModel.SetBackupVaultName("testString")
				createBackupVaultOptionsModel.SetRegion("testString")
				createBackupVaultOptionsModel.SetActivityTracking(backupVaultActivityTrackingModel)
				createBackupVaultOptionsModel.SetMetricsMonitoring(backupVaultMetricsMonitoringModel)
				createBackupVaultOptionsModel.SetSseKpCustomerRootKeyCrn("testString")
				createBackupVaultOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createBackupVaultOptionsModel).ToNot(BeNil())
				Expect(createBackupVaultOptionsModel.ServiceInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createBackupVaultOptionsModel.BackupVaultName).To(Equal(core.StringPtr("testString")))
				Expect(createBackupVaultOptionsModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(createBackupVaultOptionsModel.ActivityTracking).To(Equal(backupVaultActivityTrackingModel))
				Expect(createBackupVaultOptionsModel.MetricsMonitoring).To(Equal(backupVaultMetricsMonitoringModel))
				Expect(createBackupVaultOptionsModel.SseKpCustomerRootKeyCrn).To(Equal(core.StringPtr("testString")))
				Expect(createBackupVaultOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateRestoreOptions successfully`, func() {
				// Construct an instance of the CreateRestoreOptions model
				backupVaultName := "testString"
				createRestoreOptionsRecoveryRangeID := "6ff0d31c-7583-4463-8ae5-208752f5769c"
				createRestoreOptionsRestoreType := "in_place"
				createRestoreOptionsRestorePointInTime := CreateMockDateTime("2024-06-04T12:12:00.000Z")
				createRestoreOptionsTargetResourceCrn := "crn:v1:bluemix:public:cloud-object-storage:global:a1229395:8dfbcba4e6a740e3866020847e525436:bucket:targetbucket"
				createRestoreOptionsModel := resourceConfigurationService.NewCreateRestoreOptions(backupVaultName, createRestoreOptionsRecoveryRangeID, createRestoreOptionsRestoreType, createRestoreOptionsRestorePointInTime, createRestoreOptionsTargetResourceCrn)
				createRestoreOptionsModel.SetBackupVaultName("testString")
				createRestoreOptionsModel.SetRecoveryRangeID("6ff0d31c-7583-4463-8ae5-208752f5769c")
				createRestoreOptionsModel.SetRestoreType("in_place")
				createRestoreOptionsModel.SetRestorePointInTime(CreateMockDateTime("2024-06-04T12:12:00.000Z"))
				createRestoreOptionsModel.SetTargetResourceCrn("crn:v1:bluemix:public:cloud-object-storage:global:a1229395:8dfbcba4e6a740e3866020847e525436:bucket:targetbucket")
				createRestoreOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createRestoreOptionsModel).ToNot(BeNil())
				Expect(createRestoreOptionsModel.BackupVaultName).To(Equal(core.StringPtr("testString")))
				Expect(createRestoreOptionsModel.RecoveryRangeID).To(Equal(core.StringPtr("6ff0d31c-7583-4463-8ae5-208752f5769c")))
				Expect(createRestoreOptionsModel.RestoreType).To(Equal(core.StringPtr("in_place")))
				Expect(createRestoreOptionsModel.RestorePointInTime).To(Equal(CreateMockDateTime("2024-06-04T12:12:00.000Z")))
				Expect(createRestoreOptionsModel.TargetResourceCrn).To(Equal(core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a1229395:8dfbcba4e6a740e3866020847e525436:bucket:targetbucket")))
				Expect(createRestoreOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteBackupPolicyOptions successfully`, func() {
				// Construct an instance of the DeleteBackupPolicyOptions model
				bucket := "testString"
				policyID := "testString"
				deleteBackupPolicyOptionsModel := resourceConfigurationService.NewDeleteBackupPolicyOptions(bucket, policyID)
				deleteBackupPolicyOptionsModel.SetBucket("testString")
				deleteBackupPolicyOptionsModel.SetPolicyID("testString")
				deleteBackupPolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteBackupPolicyOptionsModel).ToNot(BeNil())
				Expect(deleteBackupPolicyOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(deleteBackupPolicyOptionsModel.PolicyID).To(Equal(core.StringPtr("testString")))
				Expect(deleteBackupPolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteBackupVaultOptions successfully`, func() {
				// Construct an instance of the DeleteBackupVaultOptions model
				backupVaultName := "testString"
				deleteBackupVaultOptionsModel := resourceConfigurationService.NewDeleteBackupVaultOptions(backupVaultName)
				deleteBackupVaultOptionsModel.SetBackupVaultName("testString")
				deleteBackupVaultOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteBackupVaultOptionsModel).ToNot(BeNil())
				Expect(deleteBackupVaultOptionsModel.BackupVaultName).To(Equal(core.StringPtr("testString")))
				Expect(deleteBackupVaultOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetBackupPolicyOptions successfully`, func() {
				// Construct an instance of the GetBackupPolicyOptions model
				bucket := "testString"
				policyID := "testString"
				getBackupPolicyOptionsModel := resourceConfigurationService.NewGetBackupPolicyOptions(bucket, policyID)
				getBackupPolicyOptionsModel.SetBucket("testString")
				getBackupPolicyOptionsModel.SetPolicyID("testString")
				getBackupPolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBackupPolicyOptionsModel).ToNot(BeNil())
				Expect(getBackupPolicyOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(getBackupPolicyOptionsModel.PolicyID).To(Equal(core.StringPtr("testString")))
				Expect(getBackupPolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetBackupVaultOptions successfully`, func() {
				// Construct an instance of the GetBackupVaultOptions model
				backupVaultName := "testString"
				getBackupVaultOptionsModel := resourceConfigurationService.NewGetBackupVaultOptions(backupVaultName)
				getBackupVaultOptionsModel.SetBackupVaultName("testString")
				getBackupVaultOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBackupVaultOptionsModel).ToNot(BeNil())
				Expect(getBackupVaultOptionsModel.BackupVaultName).To(Equal(core.StringPtr("testString")))
				Expect(getBackupVaultOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetBucketConfigOptions successfully`, func() {
				// Construct an instance of the GetBucketConfigOptions model
				bucket := "testString"
				getBucketConfigOptionsModel := resourceConfigurationService.NewGetBucketConfigOptions(bucket)
				getBucketConfigOptionsModel.SetBucket("testString")
				getBucketConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBucketConfigOptionsModel).ToNot(BeNil())
				Expect(getBucketConfigOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(getBucketConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetRestoreOptions successfully`, func() {
				// Construct an instance of the GetRestoreOptions model
				backupVaultName := "testString"
				restoreID := "testString"
				getRestoreOptionsModel := resourceConfigurationService.NewGetRestoreOptions(backupVaultName, restoreID)
				getRestoreOptionsModel.SetBackupVaultName("testString")
				getRestoreOptionsModel.SetRestoreID("testString")
				getRestoreOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRestoreOptionsModel).ToNot(BeNil())
				Expect(getRestoreOptionsModel.BackupVaultName).To(Equal(core.StringPtr("testString")))
				Expect(getRestoreOptionsModel.RestoreID).To(Equal(core.StringPtr("testString")))
				Expect(getRestoreOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSourceResourceRecoveryRangeOptions successfully`, func() {
				// Construct an instance of the GetSourceResourceRecoveryRangeOptions model
				backupVaultName := "testString"
				recoveryRangeID := "testString"
				getSourceResourceRecoveryRangeOptionsModel := resourceConfigurationService.NewGetSourceResourceRecoveryRangeOptions(backupVaultName, recoveryRangeID)
				getSourceResourceRecoveryRangeOptionsModel.SetBackupVaultName("testString")
				getSourceResourceRecoveryRangeOptionsModel.SetRecoveryRangeID("testString")
				getSourceResourceRecoveryRangeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSourceResourceRecoveryRangeOptionsModel).ToNot(BeNil())
				Expect(getSourceResourceRecoveryRangeOptionsModel.BackupVaultName).To(Equal(core.StringPtr("testString")))
				Expect(getSourceResourceRecoveryRangeOptionsModel.RecoveryRangeID).To(Equal(core.StringPtr("testString")))
				Expect(getSourceResourceRecoveryRangeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListBackupPoliciesOptions successfully`, func() {
				// Construct an instance of the ListBackupPoliciesOptions model
				bucket := "testString"
				listBackupPoliciesOptionsModel := resourceConfigurationService.NewListBackupPoliciesOptions(bucket)
				listBackupPoliciesOptionsModel.SetBucket("testString")
				listBackupPoliciesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listBackupPoliciesOptionsModel).ToNot(BeNil())
				Expect(listBackupPoliciesOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(listBackupPoliciesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListBackupVaultsOptions successfully`, func() {
				// Construct an instance of the ListBackupVaultsOptions model
				serviceInstanceID := "testString"
				listBackupVaultsOptionsModel := resourceConfigurationService.NewListBackupVaultsOptions(serviceInstanceID)
				listBackupVaultsOptionsModel.SetServiceInstanceID("testString")
				listBackupVaultsOptionsModel.SetToken("testString")
				listBackupVaultsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listBackupVaultsOptionsModel).ToNot(BeNil())
				Expect(listBackupVaultsOptionsModel.ServiceInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listBackupVaultsOptionsModel.Token).To(Equal(core.StringPtr("testString")))
				Expect(listBackupVaultsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListRecoveryRangesOptions successfully`, func() {
				// Construct an instance of the ListRecoveryRangesOptions model
				backupVaultName := "testString"
				listRecoveryRangesOptionsModel := resourceConfigurationService.NewListRecoveryRangesOptions(backupVaultName)
				listRecoveryRangesOptionsModel.SetBackupVaultName("testString")
				listRecoveryRangesOptionsModel.SetSourceResourceCrn("testString")
				listRecoveryRangesOptionsModel.SetLatest("testString")
				listRecoveryRangesOptionsModel.SetToken("testString")
				listRecoveryRangesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listRecoveryRangesOptionsModel).ToNot(BeNil())
				Expect(listRecoveryRangesOptionsModel.BackupVaultName).To(Equal(core.StringPtr("testString")))
				Expect(listRecoveryRangesOptionsModel.SourceResourceCrn).To(Equal(core.StringPtr("testString")))
				Expect(listRecoveryRangesOptionsModel.Latest).To(Equal(core.StringPtr("testString")))
				Expect(listRecoveryRangesOptionsModel.Token).To(Equal(core.StringPtr("testString")))
				Expect(listRecoveryRangesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListRestoresOptions successfully`, func() {
				// Construct an instance of the ListRestoresOptions model
				backupVaultName := "testString"
				listRestoresOptionsModel := resourceConfigurationService.NewListRestoresOptions(backupVaultName)
				listRestoresOptionsModel.SetBackupVaultName("testString")
				listRestoresOptionsModel.SetToken("testString")
				listRestoresOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listRestoresOptionsModel).ToNot(BeNil())
				Expect(listRestoresOptionsModel.BackupVaultName).To(Equal(core.StringPtr("testString")))
				Expect(listRestoresOptionsModel.Token).To(Equal(core.StringPtr("testString")))
				Expect(listRestoresOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPatchSourceResourceRecoveryRangeOptions successfully`, func() {
				// Construct an instance of the PatchSourceResourceRecoveryRangeOptions model
				backupVaultName := "testString"
				recoveryRangeID := "testString"
				recoveryRangePatch := map[string]interface{}{"anyKey": "anyValue"}
				patchSourceResourceRecoveryRangeOptionsModel := resourceConfigurationService.NewPatchSourceResourceRecoveryRangeOptions(backupVaultName, recoveryRangeID, recoveryRangePatch)
				patchSourceResourceRecoveryRangeOptionsModel.SetBackupVaultName("testString")
				patchSourceResourceRecoveryRangeOptionsModel.SetRecoveryRangeID("testString")
				patchSourceResourceRecoveryRangeOptionsModel.SetRecoveryRangePatch(map[string]interface{}{"anyKey": "anyValue"})
				patchSourceResourceRecoveryRangeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(patchSourceResourceRecoveryRangeOptionsModel).ToNot(BeNil())
				Expect(patchSourceResourceRecoveryRangeOptionsModel.BackupVaultName).To(Equal(core.StringPtr("testString")))
				Expect(patchSourceResourceRecoveryRangeOptionsModel.RecoveryRangeID).To(Equal(core.StringPtr("testString")))
				Expect(patchSourceResourceRecoveryRangeOptionsModel.RecoveryRangePatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(patchSourceResourceRecoveryRangeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateBackupVaultOptions successfully`, func() {
				// Construct an instance of the UpdateBackupVaultOptions model
				backupVaultName := "testString"
				backupVaultPatch := map[string]interface{}{"anyKey": "anyValue"}
				updateBackupVaultOptionsModel := resourceConfigurationService.NewUpdateBackupVaultOptions(backupVaultName, backupVaultPatch)
				updateBackupVaultOptionsModel.SetBackupVaultName("testString")
				updateBackupVaultOptionsModel.SetBackupVaultPatch(map[string]interface{}{"anyKey": "anyValue"})
				updateBackupVaultOptionsModel.SetIfMatch("testString")
				updateBackupVaultOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateBackupVaultOptionsModel).ToNot(BeNil())
				Expect(updateBackupVaultOptionsModel.BackupVaultName).To(Equal(core.StringPtr("testString")))
				Expect(updateBackupVaultOptionsModel.BackupVaultPatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateBackupVaultOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateBackupVaultOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateBucketConfigOptions successfully`, func() {
				// Construct an instance of the UpdateBucketConfigOptions model
				bucket := "testString"
				updateBucketConfigOptionsModel := resourceConfigurationService.NewUpdateBucketConfigOptions(bucket)
				updateBucketConfigOptionsModel.SetBucket("testString")
				updateBucketConfigOptionsModel.SetBucketPatch(map[string]interface{}{"anyKey": "anyValue"})
				updateBucketConfigOptionsModel.SetIfMatch("testString")
				updateBucketConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateBucketConfigOptionsModel).ToNot(BeNil())
				Expect(updateBucketConfigOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(updateBucketConfigOptionsModel.BucketPatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateBucketConfigOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateBucketConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalActivityTracking successfully`, func() {
			// Construct an instance of the model.
			model := new(resourceconfigurationv1.ActivityTracking)
			model.ReadDataEvents = core.BoolPtr(false)
			model.WriteDataEvents = core.BoolPtr(false)
			model.ActivityTrackerCrn = core.StringPtr("testString")
			model.ManagementEvents = core.BoolPtr(false)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *resourceconfigurationv1.ActivityTracking
			err = resourceconfigurationv1.UnmarshalActivityTracking(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalBackupVaultActivityTracking successfully`, func() {
			// Construct an instance of the model.
			model := new(resourceconfigurationv1.BackupVaultActivityTracking)
			model.ManagementEvents = core.BoolPtr(true)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *resourceconfigurationv1.BackupVaultActivityTracking
			err = resourceconfigurationv1.UnmarshalBackupVaultActivityTracking(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalBackupVaultMetricsMonitoring successfully`, func() {
			// Construct an instance of the model.
			model := new(resourceconfigurationv1.BackupVaultMetricsMonitoring)
			model.UsageMetricsEnabled = core.BoolPtr(true)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *resourceconfigurationv1.BackupVaultMetricsMonitoring
			err = resourceconfigurationv1.UnmarshalBackupVaultMetricsMonitoring(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalBackupVaultPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(resourceconfigurationv1.BackupVaultPatch)
			model.ActivityTracking = nil
			model.MetricsMonitoring = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *resourceconfigurationv1.BackupVaultPatch
			err = resourceconfigurationv1.UnmarshalBackupVaultPatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalBucketPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(resourceconfigurationv1.BucketPatch)
			model.Firewall = nil
			model.ActivityTracking = nil
			model.MetricsMonitoring = nil
			model.HardQuota = core.Int64Ptr(int64(28198745752445146))
			model.ProtectionManagement = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *resourceconfigurationv1.BucketPatch
			err = resourceconfigurationv1.UnmarshalBucketPatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalDeleteAfterDays successfully`, func() {
			// Construct an instance of the model.
			model := new(resourceconfigurationv1.DeleteAfterDays)
			model.DeleteAfterDays = core.Int64Ptr(int64(10))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *resourceconfigurationv1.DeleteAfterDays
			err = resourceconfigurationv1.UnmarshalDeleteAfterDays(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalFirewall successfully`, func() {
			// Construct an instance of the model.
			model := new(resourceconfigurationv1.Firewall)
			model.AllowedIp = []string{"testString"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *resourceconfigurationv1.Firewall
			err = resourceconfigurationv1.UnmarshalFirewall(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalMetricsMonitoring successfully`, func() {
			// Construct an instance of the model.
			model := new(resourceconfigurationv1.MetricsMonitoring)
			model.UsageMetricsEnabled = core.BoolPtr(false)
			model.RequestMetricsEnabled = core.BoolPtr(false)
			model.MetricsMonitoringCrn = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *resourceconfigurationv1.MetricsMonitoring
			err = resourceconfigurationv1.UnmarshalMetricsMonitoring(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalProtectionManagement successfully`, func() {
			// Construct an instance of the model.
			model := new(resourceconfigurationv1.ProtectionManagement)
			model.RequestedState = core.StringPtr("activate")
			model.ProtectionManagementToken = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *resourceconfigurationv1.ProtectionManagement
			err = resourceconfigurationv1.UnmarshalProtectionManagement(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalRecoveryRangePatch successfully`, func() {
			// Construct an instance of the model.
			model := new(resourceconfigurationv1.RecoveryRangePatch)
			model.Retention = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *resourceconfigurationv1.RecoveryRangePatch
			err = resourceconfigurationv1.UnmarshalRecoveryRangePatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("VGhpcyBpcyBhIHRlc3Qgb2YgdGhlIGVtZXJnZW5jeSBicm9hZGNhc3Qgc3lzdGVt")
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
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(encodedString string) *[]byte {
	ba, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		panic(err)
	}
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
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
