package rancher2

import (
	"testing"

	managementClient "github.com/rancher/rancher/pkg/client/generated/management/v3"
	"github.com/stretchr/testify/assert"
)

var (
	testCloudCredentialS3Conf      *managementClient.S3CredentialConfig
	testCloudCredentialS3Interface []interface{}
)

func init() {
	testCloudCredentialS3Conf = &managementClient.S3CredentialConfig{
		AccessKey:            "access_key",
		SecretKey:            "secret_key",
		DefaultBucket:        "default_bucket",
		DefaultEndpoint:      "default_endpoint",
		DefaultEndpointCA:    "default_endpoint_ca",
		DefaultFolder:        "default_folder",
		DefaultRegion:        "default_region",
		DefaultSkipSSLVerify: "true",
	}
	testCloudCredentialS3Interface = []interface{}{
		map[string]interface{}{
			"access_key":              "access_key",
			"secret_key":              "secret_key",
			"default_bucket":          "default_bucket",
			"default_endpoint":        "default_endpoint",
			"default_endpoint_ca":     "default_endpoint_ca",
			"default_folder":          "default_folder",
			"default_region":          "default_region",
			"default_skip_ssl_verify": true,
		},
	}
}

func TestFlattenCloudCredentialS3(t *testing.T) {

	cases := []struct {
		Input          *managementClient.S3CredentialConfig
		ExpectedOutput []interface{}
	}{
		{
			testCloudCredentialS3Conf,
			testCloudCredentialS3Interface,
		},
	}

	for _, tc := range cases {
		output := flattenCloudCredentialS3(tc.Input, tc.ExpectedOutput)
		assert.Equal(t, tc.ExpectedOutput, output, "Unexpected output from flattener.")
	}
}

func TestExpandCloudCredentialS3(t *testing.T) {

	cases := []struct {
		Input          []interface{}
		ExpectedOutput *managementClient.S3CredentialConfig
	}{
		{
			testCloudCredentialS3Interface,
			testCloudCredentialS3Conf,
		},
	}

	for _, tc := range cases {
		output := expandCloudCredentialS3(tc.Input)
		assert.Equal(t, tc.ExpectedOutput, output, "Unexpected output from expander.")
	}
}
