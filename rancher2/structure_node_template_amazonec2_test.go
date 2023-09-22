package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	managementClient "github.com/rancher/rancher/pkg/client/generated/management/v3"
	"github.com/stretchr/testify/assert"

	"testing"
)

var (
	testNodeTemplateNodeTaintsConf               []managementClient.Taint
	testNodeTemplateAmazonEc2Conf                amazonec2Config
	testNodeTemplateAmazonEc2Interface           map[string]interface{}
	testNodeTemplateConf                         *NodeTemplate
	testNodeTemplateSquashAmazonEc2ConfInterface map[string]interface{}
	testNodeTemplateExpandAmazonEc2ConfInterface map[string]interface{}
	testNodeTemplateNodeTaintsInterface          interface{}
)

func init() {
	testNodeTemplateNodeTaintsConf = []managementClient.Taint{
		{
			Key:       "key",
			Value:     "value",
			Effect:    "recipient",
			TimeAdded: "time_added",
		},
	}
	testNodeTemplateNodeTaintsInterface = []interface{}{
		map[string]interface{}{
			"key":        "key",
			"value":      "value",
			"effect":     "recipient",
			"time_added": "time_added",
		},
	}
	testNodeTemplateAmazonEc2Conf = amazonec2Config{
		Ami:                  "ubuntu",
		DeviceName:           "/dev/sda1",
		HTTPTokens:           "true",
		HTTPEndpoint:         "true",
		EncryptEbsVolume:     false,
		InstanceType:         "t2.micro",
		BlockDurationMinutes: "0",
		Region:               "us-east-1",
		Retries:              "3",
		RootSize:             "16",
		SpotPrice:            "0.50",
		VolumeType:           "gp2",
		SSHUser:              "ubuntu",
		Zone:                 "a",
	}
	testNodeTemplateAmazonEc2Interface = map[string]interface{}{
		"ami":                  "ubuntu",
		"deviceName":           "/dev/sda1",
		"http_tokens":          "true",
		"http_endpoint":        "true",
		"encryptEbsVolume":     "false",
		"instanceType":         "t2.micro",
		"blockDurationMinutes": "0",
		"region":               "us-east-1",
		"retries":              "3",
		"rootSize":             "16",
		"spotPrice":            "0.50",
		"volumeType":           "gp2",
		"sshUser":              "ubuntu",
		"zone":                 "a",
	}
	testNodeTemplateAnnotationsConf := map[string]string{
		"key": "value",
	}
	testNodeTemplateAnnotationsInterface := map[string]interface{}{
		"key": "value",
	}
	useInternalIP := false
	testNodeTemplateConf = &NodeTemplate{
		NodeTemplate: managementClient.NodeTemplate{
			Driver:               "amazonec2",
			Annotations:          testNodeTemplateAnnotationsConf,
			CloudCredentialID:    "abc-test-123",
			NodeTaints:           testNodeTemplateNodeTaintsConf,
			EngineInstallURL:     "http://fake.url",
			Name:                 "test-node-template",
			UseInternalIPAddress: &useInternalIP,
		},
		Amazonec2Config: &testNodeTemplateAmazonEc2Conf,
	}
	testNodeTemplateSquashAmazonEc2ConfInterface = map[string]interface{}{
		"annotations":             testNodeTemplateAnnotationsInterface,
		"driver":                  "amazonec2",
		"cloud_credential_id":     "abc-test-123",
		"use_internal_ip_address": useInternalIP,
		"engine_install_url":      "http://fake.url",
		"name":                    "test-node-template",
	}

	testNodeTemplateExpandAmazonEc2ConfInterface = map[string]interface{}{
		"annotations":             testNodeTemplateAnnotationsInterface,
		"node_taints":             testNodeTemplateNodeTaintsInterface,
		"driver":                  "amazonec2",
		"cloud_credential_id":     "abc-test-123",
		"use_internal_ip_address": useInternalIP,
		"engine_install_url":      "http://fake.url",
		"name":                    "test-node-template",
		"amazonec2_config":        []interface{}{testNodeTemplateAmazonEc2Interface},
	}

}

func TestFlattenNodeTemplate(t *testing.T) {
	cases := []struct {
		Input          *NodeTemplate
		ExpectedOutput map[string]interface{}
	}{
		{
			testNodeTemplateConf,
			testNodeTemplateSquashAmazonEc2ConfInterface,
		},
	}

	for _, tc := range cases {
		output := schema.TestResourceDataRaw(t, nodeTemplateFields(), map[string]interface{}{})
		err := flattenNodeTemplate(output, tc.Input)
		if err != nil {
			assert.FailNow(t, "[ERROR] on flattener: %#v", err)
		}
		expectedOutput := map[string]interface{}{}
		for k := range tc.ExpectedOutput {
			expectedOutput[k] = output.Get(k)
		}
		assert.Equal(t, tc.ExpectedOutput, expectedOutput, "Unexpected output from flattener.")
	}
}

func TestExpandNodeTemplate(t *testing.T) {
	cases := []struct {
		Input          map[string]interface{}
		ExpectedOutput *NodeTemplate
	}{
		{
			Input:          testNodeTemplateExpandAmazonEc2ConfInterface,
			ExpectedOutput: testNodeTemplateConf,
		},
	}

	for _, tc := range cases {
		inputData := schema.TestResourceDataRaw(t, nodeTemplateFields(), tc.Input)
		output := expandNodeTemplate(inputData)
		assert.Equal(t, tc.ExpectedOutput, output, "Unexpected output from expander.")
	}
}
