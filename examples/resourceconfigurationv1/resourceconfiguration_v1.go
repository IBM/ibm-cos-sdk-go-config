package main

import (
	"fmt"

	rc "github.com/IBM/ibm-cos-sdk-go-config/resourceconfigurationv1"

	"github.com/IBM/ibm-cos-sdk-go/aws"
	"github.com/IBM/ibm-cos-sdk-go/aws/credentials/ibmiam"
	"github.com/IBM/ibm-cos-sdk-go/aws/session"
	"github.com/IBM/ibm-cos-sdk-go/service/s3"

	"github.com/IBM/go-sdk-core/v3/core"
)

// A Resource Configuration Service example
// First, ensure the credentials are correct, and then run this script:
//    go run -v resource_configuration_v1.go

const (
	apiKey            = "<api_key>"
	serviceInstanceID = "<service_instance_id>"
	authEndpoint      = "https://iam.cloud.ibm.com/identity/token"
	serviceEndpoint   = "s3.us.cloud-object-storage.appdomain.cloud"
)

func main() {

	// Bucket name
	bName := "testbucket"

	// Institate S3 SDK Client on IBM Cloud PROD
	conf := aws.NewConfig().
		WithEndpoint(serviceEndpoint).
		WithCredentials(ibmiam.NewStaticCredentials(aws.NewConfig(),
			authEndpoint, apiKey, serviceInstanceID)).
		WithS3ForcePathStyle(true)

	// Create S3 client
	sess := session.Must(session.NewSession())
	client := s3.New(sess, conf)

	// Create a bucket
	input := &s3.CreateBucketInput{
		Bucket: aws.String(bName),
	}
	d, e := client.CreateBucket(input)
	fmt.Println(d) // should print an empty bracket
	fmt.Println(e) // should print <nil>

	// Instantiate the IAM Authenticator with API KEY
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
	}

	// Build an RC Service with the IAM authenticator
	service, serviceErr := rc.NewResourceConfigurationV1(&rc.ResourceConfigurationV1Options{
		Authenticator: authenticator,
	})
	// Check successful instantiation
	if serviceErr != nil {
		fmt.Println(serviceErr)
		return
	}

	// Update Config Options
	uOptions := &rc.UpdateBucketConfigOptions{
		Bucket: core.StringPtr(bName),
		Firewall: &rc.Firewall{
			AllowedIp: []string{"192.168.1.95", "192.168.1.100"},
		},
	}

	// Update Bucket Config
	_, e = service.UpdateBucketConfig(uOptions)
	// Check successful call
	if e != nil {
		fmt.Println(e)
		return
	}

	// Bucket Config Options
	GetBucketConfigOptions := service.NewGetBucketConfigOptions(bName)

	// Get Bucket Config
	result, response, e := service.GetBucketConfig(GetBucketConfigOptions)
	// Check successful call
	if e != nil {
		fmt.Println(e)
		return
	}

	// Print response and result outputs
	fmt.Println(result)
	fmt.Println(response)
}
