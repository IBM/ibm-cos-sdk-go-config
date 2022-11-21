/**
 * (C) Copyright IBM Corp. 2022.
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
				"RESOURCE_CONFIGURATION_URL": "https://resourceconfigurationv1/api",
				"RESOURCE_CONFIGURATION_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1UsingExternalConfig(&resourceconfigurationv1.ResourceConfigurationV1Options{
				})
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
				resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1UsingExternalConfig(&resourceconfigurationv1.ResourceConfigurationV1Options{
				})
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
				"RESOURCE_CONFIGURATION_URL": "https://resourceconfigurationv1/api",
				"RESOURCE_CONFIGURATION_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			resourceConfigurationService, serviceErr := resourceconfigurationv1.NewResourceConfigurationV1UsingExternalConfig(&resourceconfigurationv1.ResourceConfigurationV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(resourceConfigurationService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"RESOURCE_CONFIGURATION_AUTH_TYPE":   "NOAuth",
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
					fmt.Fprintf(res, "%s", `{"name": "my-new-bucket", "crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/3bf0d9003abfb5d29761c3e97696b71c:xxxxxxx-6c4f-4a62-a165-696756d63903:bucket:my-new-bucket", "service_instance_id": "d6f04d83-6c4f-4a62-a165-696756d63903", "service_instance_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/3bf0d9003abfb5d29761c3e97696b71c:xxxxxxx-6c4f-4a62-a165-696756d63903::", "time_created": "2018-03-26T16:23:36.980Z", "time_updated": "2018-10-17T19:29:10.117Z", "object_count": 764265234, "bytes_used": 28198745752445145, "noncurrent_object_count": 764265234, "noncurrent_bytes_used": 844239234, "delete_marker_count": 827201, "firewall": {"allowed_ip": ["AllowedIp"], "denied_ip": ["DeniedIp"], "allowed_network_type": ["public"]}, "activity_tracking": {"read_data_events": false, "write_data_events": false, "activity_tracker_crn": "ActivityTrackerCrn"}, "metrics_monitoring": {"usage_metrics_enabled": false, "request_metrics_enabled": false, "metrics_monitoring_crn": "MetricsMonitoringCrn"}, "hard_quota": 28198745752445146}`)
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
					fmt.Fprintf(res, "%s", `{"name": "my-new-bucket", "crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/3bf0d9003abfb5d29761c3e97696b71c:xxxxxxx-6c4f-4a62-a165-696756d63903:bucket:my-new-bucket", "service_instance_id": "d6f04d83-6c4f-4a62-a165-696756d63903", "service_instance_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/3bf0d9003abfb5d29761c3e97696b71c:xxxxxxx-6c4f-4a62-a165-696756d63903::", "time_created": "2018-03-26T16:23:36.980Z", "time_updated": "2018-10-17T19:29:10.117Z", "object_count": 764265234, "bytes_used": 28198745752445145, "noncurrent_object_count": 764265234, "noncurrent_bytes_used": 844239234, "delete_marker_count": 827201, "firewall": {"allowed_ip": ["AllowedIp"], "denied_ip": ["DeniedIp"], "allowed_network_type": ["public"]}, "activity_tracking": {"read_data_events": false, "write_data_events": false, "activity_tracker_crn": "ActivityTrackerCrn"}, "metrics_monitoring": {"usage_metrics_enabled": false, "request_metrics_enabled": false, "metrics_monitoring_crn": "MetricsMonitoringCrn"}, "hard_quota": 28198745752445146}`)
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
				firewallModel.DeniedIp = []string{"testString"}
				firewallModel.AllowedNetworkType = []string{"public"}

				// Construct an instance of the ActivityTracking model
				activityTrackingModel := new(resourceconfigurationv1.ActivityTracking)
				activityTrackingModel.ReadDataEvents = core.BoolPtr(false)
				activityTrackingModel.WriteDataEvents = core.BoolPtr(false)
				activityTrackingModel.ActivityTrackerCrn = core.StringPtr("testString")

				// Construct an instance of the MetricsMonitoring model
				metricsMonitoringModel := new(resourceconfigurationv1.MetricsMonitoring)
				metricsMonitoringModel.UsageMetricsEnabled = core.BoolPtr(false)
				metricsMonitoringModel.RequestMetricsEnabled = core.BoolPtr(false)
				metricsMonitoringModel.MetricsMonitoringCrn = core.StringPtr("testString")

				// Construct an instance of the BucketPatch model
				bucketPatchModel := new(resourceconfigurationv1.BucketPatch)
				bucketPatchModel.Firewall = firewallModel
				bucketPatchModel.ActivityTracking = activityTrackingModel
				bucketPatchModel.MetricsMonitoring = metricsMonitoringModel
				bucketPatchModel.HardQuota = core.Int64Ptr(int64(28198745752445146))
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
				firewallModel.DeniedIp = []string{"testString"}
				firewallModel.AllowedNetworkType = []string{"public"}

				// Construct an instance of the ActivityTracking model
				activityTrackingModel := new(resourceconfigurationv1.ActivityTracking)
				activityTrackingModel.ReadDataEvents = core.BoolPtr(false)
				activityTrackingModel.WriteDataEvents = core.BoolPtr(false)
				activityTrackingModel.ActivityTrackerCrn = core.StringPtr("testString")

				// Construct an instance of the MetricsMonitoring model
				metricsMonitoringModel := new(resourceconfigurationv1.MetricsMonitoring)
				metricsMonitoringModel.UsageMetricsEnabled = core.BoolPtr(false)
				metricsMonitoringModel.RequestMetricsEnabled = core.BoolPtr(false)
				metricsMonitoringModel.MetricsMonitoringCrn = core.StringPtr("testString")

				// Construct an instance of the BucketPatch model
				bucketPatchModel := new(resourceconfigurationv1.BucketPatch)
				bucketPatchModel.Firewall = firewallModel
				bucketPatchModel.ActivityTracking = activityTrackingModel
				bucketPatchModel.MetricsMonitoring = metricsMonitoringModel
				bucketPatchModel.HardQuota = core.Int64Ptr(int64(28198745752445146))
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
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			resourceConfigurationService, _ := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
				URL:           "http://resourceconfigurationv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
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
