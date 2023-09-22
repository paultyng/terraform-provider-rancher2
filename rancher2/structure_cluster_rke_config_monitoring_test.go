package rancher2

import (
	"testing"

	managementClient "github.com/rancher/rancher/pkg/client/generated/management/v3"
	"github.com/stretchr/testify/assert"
)

var (
	testClusterRKEConfigMonitoringTolerationsConf      []managementClient.Toleration
	testClusterRKEConfigMonitoringTolerationsInterface []interface{}
	testClusterRKEConfigMonitoringConf                 *managementClient.MonitoringConfig
	testClusterRKEConfigMonitoringInterface            []interface{}
	testClusterRKEConfigMonitoringReplicas             int64
)

func init() {
	seconds := int64(10)
	testClusterRKEConfigMonitoringTolerationsConf = []managementClient.Toleration{
		{
			Key:               "key",
			Value:             "value",
			Effect:            "recipient",
			Operator:          "operator",
			TolerationSeconds: &seconds,
		},
	}
	testClusterRKEConfigMonitoringTolerationsInterface = []interface{}{
		map[string]interface{}{
			"key":      "key",
			"value":    "value",
			"effect":   "recipient",
			"operator": "operator",
			"seconds":  10,
		},
	}
	testClusterRKEConfigMonitoringReplicas = int64(2)
	testClusterRKEConfigMonitoringConf = &managementClient.MonitoringConfig{
		NodeSelector: map[string]string{
			"selector1": "value1",
			"selector2": "value2",
		},
		Options: map[string]string{
			"option1": "value1",
			"option2": "value2",
		},
		Provider:    "test",
		Replicas:    &testClusterRKEConfigMonitoringReplicas,
		Tolerations: testClusterRKEConfigMonitoringTolerationsConf,
	}
	testClusterRKEConfigMonitoringInterface = []interface{}{
		map[string]interface{}{
			"node_selector": map[string]interface{}{
				"selector1": "value1",
				"selector2": "value2",
			},
			"options": map[string]interface{}{
				"option1": "value1",
				"option2": "value2",
			},
			"provider":    "test",
			"replicas":    2,
			"tolerations": testClusterRKEConfigMonitoringTolerationsInterface,
		},
	}
}

func TestFlattenClusterRKEConfigMonitoring(t *testing.T) {

	cases := []struct {
		Input          *managementClient.MonitoringConfig
		ExpectedOutput []interface{}
	}{
		{
			testClusterRKEConfigMonitoringConf,
			testClusterRKEConfigMonitoringInterface,
		},
	}

	for _, tc := range cases {
		output, err := flattenClusterRKEConfigMonitoring(tc.Input)
		if err != nil {
			assert.FailNow(t, "[ERROR] on flattener: %#v", err)
		}
		assert.Equal(t, tc.ExpectedOutput, output, "Unexpected output from flattener.")
	}
}

func TestExpandClusterRKEConfigMonitoring(t *testing.T) {

	cases := []struct {
		Input          []interface{}
		ExpectedOutput *managementClient.MonitoringConfig
	}{
		{
			testClusterRKEConfigMonitoringInterface,
			testClusterRKEConfigMonitoringConf,
		},
	}

	for _, tc := range cases {
		output, err := expandClusterRKEConfigMonitoring(tc.Input)
		if err != nil {
			assert.FailNow(t, "[ERROR] on expander: %#v", err)
		}
		assert.Equal(t, tc.ExpectedOutput, output, "Unexpected output from expander.")
	}
}
