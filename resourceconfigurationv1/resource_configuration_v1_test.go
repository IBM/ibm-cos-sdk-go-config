/**
 * (C) Copyright IBM Corp. 2019.
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
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/IBM/ibm-cos-sdk-go-config/resourceconfigurationv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

var _ = Describe(`ResourceConfigurationV1`, func() {
	Describe(`GetBucketConfig(getBucketConfigOptions *GetBucketConfigOptions)`, func() {
		bearerToken := "0ui9876453"
		getBucketConfigPath := "/b/{bucket}"
		getBucketConfigPath = strings.Replace(getBucketConfigPath, "{bucket}", "testString", 1)
		Context(`Successfully - Returns metadata for the specified bucket`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getBucketConfigPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"name": "fake_Name", "crn": "fake_Crn", "service_instance_id": "fake_ServiceInstanceID", "service_instance_crn": "fake_ServiceInstanceCrn", "time_created": "2017-05-16T13:56:54.957Z", "time_updated": "2017-05-16T13:56:54.957Z", "object_count": 11, "bytes_used": 9, "firewall": {"allowed_ip": []}, "activity_tracking": {"read_data_events": true, "write_data_events": false, "activity_tracker_crn": "fake_ActivityTrackerCrn"}, "metrics_monitoring": {"usage_metrics_enabled": false, "metrics_monitoring_crn": "fake_MetricsMonitoringCrn"}}`)
			}))
			It(`Succeed to call GetBucketConfig`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetBucketConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBucketConfigOptions model
				getBucketConfigOptionsModel := new(resourceconfigurationv1.GetBucketConfigOptions)
				getBucketConfigOptionsModel.Bucket = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetBucketConfig(getBucketConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateBucketConfig(updateBucketConfigOptions *UpdateBucketConfigOptions)`, func() {
		bearerToken := "0ui9876453"
		updateBucketConfigPath := "/b/{bucket}"
		updateBucketConfigPath = strings.Replace(updateBucketConfigPath, "{bucket}", "testString", 1)
		Context(`Successfully - Make changes to a bucket's configuration`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateBucketConfigPath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.Header["If-Match"]).ToNot(BeNil())
				Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.WriteHeader(200)
			}))
			It(`Succeed to call UpdateBucketConfig`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourceconfigurationv1.NewResourceConfigurationV1(&resourceconfigurationv1.ResourceConfigurationV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.UpdateBucketConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the MetricsMonitoring model
				metricsMonitoringModel := new(resourceconfigurationv1.MetricsMonitoring)
				metricsMonitoringModel.UsageMetricsEnabled = core.BoolPtr(true)
				metricsMonitoringModel.MetricsMonitoringCrn = core.StringPtr("testString")

				// Construct an instance of the ActivityTracking model
				activityTrackingModel := new(resourceconfigurationv1.ActivityTracking)
				activityTrackingModel.ReadDataEvents = core.BoolPtr(true)
				activityTrackingModel.WriteDataEvents = core.BoolPtr(true)
				activityTrackingModel.ActivityTrackerCrn = core.StringPtr("testString")

				// Construct an instance of the Firewall model
				firewallModel := new(resourceconfigurationv1.Firewall)
				firewallModel.AllowedIp = []string{"testString"}

				// Construct an instance of the UpdateBucketConfigOptions model
				updateBucketConfigOptionsModel := new(resourceconfigurationv1.UpdateBucketConfigOptions)
				updateBucketConfigOptionsModel.Bucket = core.StringPtr("testString")
				updateBucketConfigOptionsModel.Firewall = firewallModel
				updateBucketConfigOptionsModel.ActivityTracking = activityTrackingModel
				updateBucketConfigOptionsModel.MetricsMonitoring = metricsMonitoringModel
				updateBucketConfigOptionsModel.IfMatch = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.UpdateBucketConfig(updateBucketConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockMap() map[string]interface{} {
	m := make(map[string]interface{})
	return m
}

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, len(mockData))
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
	d := strfmt.Date(time.Now())
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Now())
	return &d
}
