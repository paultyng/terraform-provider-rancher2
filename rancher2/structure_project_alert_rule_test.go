package rancher2

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	managementClient "github.com/rancher/rancher/pkg/client/generated/management/v3"
	"github.com/stretchr/testify/assert"
)

var (
	testProjectAlertMetricRuleConf        *managementClient.MetricRule
	testProjectAlertMetricRuleInterface   []interface{}
	testProjectAlertPodRuleConf           *managementClient.PodRule
	testProjectAlertPodRuleInterface      []interface{}
	testProjectAlertWorkloadRuleConf      *managementClient.WorkloadRule
	testProjectAlertWorkloadRuleInterface []interface{}
	testProjectAlertRuleConf              *managementClient.ProjectAlertRule
	testProjectAlertRuleInterface         map[string]interface{}
)

func init() {
	testProjectAlertMetricRuleConf = &managementClient.MetricRule{
		Comparison:     metricRuleComparisonEqual,
		Duration:       "30",
		Expression:     "expression",
		ThresholdValue: float64(3.5),
		Description:    "description",
	}
	testProjectAlertMetricRuleInterface = []interface{}{
		map[string]interface{}{
			"comparison":      metricRuleComparisonEqual,
			"duration":        "30",
			"expression":      "expression",
			"threshold_value": float64(3.5),
			"description":     "description",
		},
	}
	testProjectAlertPodRuleConf = &managementClient.PodRule{
		Condition:              podRuleConditionNotRunning,
		PodID:                  "pod_id",
		RestartIntervalSeconds: int64(70),
		RestartTimes:           int64(5),
	}
	testProjectAlertPodRuleInterface = []interface{}{
		map[string]interface{}{
			"condition":                podRuleConditionNotRunning,
			"pod_id":                   "pod_id",
			"restart_interval_seconds": 70,
			"restart_times":            5,
		},
	}
	testProjectAlertWorkloadRuleConf = &managementClient.WorkloadRule{
		AvailablePercentage: int64(70),
		Selector: map[string]string{
			"selector1": "selector1",
			"selector2": "selector2",
		},
		WorkloadID: "workload_id",
	}
	testProjectAlertWorkloadRuleInterface = []interface{}{
		map[string]interface{}{
			"available_percentage": 70,
			"selector": map[string]interface{}{
				"selector1": "selector1",
				"selector2": "selector2",
			},
			"workload_id": "workload_id",
		},
	}
	testProjectAlertRuleConf = &managementClient.ProjectAlertRule{
		Name:                  "name",
		ProjectID:             "project_id",
		GroupID:               "group_id",
		GroupIntervalSeconds:  300,
		GroupWaitSeconds:      300,
		Inherited:             newTrue(),
		MetricRule:            testProjectAlertMetricRuleConf,
		PodRule:               testProjectAlertPodRuleConf,
		RepeatIntervalSeconds: 6000,
		Severity:              alertRuleSeverityCritical,
		WorkloadRule:          testProjectAlertWorkloadRuleConf,
	}
	testProjectAlertRuleInterface = map[string]interface{}{
		"name":                    "name",
		"project_id":              "project_id",
		"group_id":                "group_id",
		"group_interval_seconds":  300,
		"group_wait_seconds":      300,
		"inherited":               true,
		"metric_rule":             testProjectAlertMetricRuleInterface,
		"pod_rule":                testProjectAlertPodRuleInterface,
		"repeat_interval_seconds": 6000,
		"severity":                alertRuleSeverityCritical,
		"workload_rule":           testProjectAlertWorkloadRuleInterface,
	}
}

func TestFlattenProjectAlertRule(t *testing.T) {

	cases := []struct {
		Input          *managementClient.ProjectAlertRule
		ExpectedOutput map[string]interface{}
	}{
		{
			testProjectAlertRuleConf,
			testProjectAlertRuleInterface,
		},
	}

	for _, tc := range cases {
		output := schema.TestResourceDataRaw(t, projectAlertRuleFields(), map[string]interface{}{})
		err := flattenProjectAlertRule(output, tc.Input)
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

func TestExpandProjectAlertRule(t *testing.T) {

	cases := []struct {
		Input          map[string]interface{}
		ExpectedOutput *managementClient.ProjectAlertRule
	}{
		{
			testProjectAlertRuleInterface,
			testProjectAlertRuleConf,
		},
	}

	for _, tc := range cases {
		inputResourceData := schema.TestResourceDataRaw(t, projectAlertRuleFields(), tc.Input)
		output := expandProjectAlertRule(inputResourceData)
		assert.Equal(t, tc.ExpectedOutput, output, "Unexpected output from expander.")
	}
}
