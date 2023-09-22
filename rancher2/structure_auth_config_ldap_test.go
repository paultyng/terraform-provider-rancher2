package rancher2

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	managementClient "github.com/rancher/rancher/pkg/client/generated/management/v3"
	"github.com/stretchr/testify/assert"
)

var (
	testAuthConfigLdapConf      *managementClient.LdapConfig
	testAuthConfigLdapInterface map[string]interface{}
)

func init() {
	testAuthConfigLdapConf = &managementClient.LdapConfig{
		AccessMode:                      "access",
		AllowedPrincipalIDs:             []string{"allowed1", "allowed2"},
		Enabled:                         true,
		Servers:                         []string{"server1", "server2"},
		ServiceAccountDistinguishedName: "service_account_distinguished_name",
		UserSearchBase:                  "user_search_base",
		Certificate:                     "certificate",
		ConnectionTimeout:               10,
		GroupDNAttribute:                "group_dn_attribute",
		GroupMemberMappingAttribute:     "group_member_mapping_attribute",
		GroupMemberUserAttribute:        "group_member_user_attribute",
		GroupNameAttribute:              "group_name_attribute",
		GroupObjectClass:                "group_object_class",
		GroupSearchAttribute:            "group_search_attribute",
		GroupSearchBase:                 "group_search_base",
		GroupSearchFilter:               "(cn=$SEARCH_STRING)",
		NestedGroupMembershipEnabled:    true,
		Port:                            389,
		TLS:                             true,
		StartTLS:                        true,
		UserDisabledBitMask:             0,
		UserLoginAttribute:              "user_login_attribute",
		UserMemberAttribute:             "user_member_attribute",
		UserNameAttribute:               "user_name_attribute",
		UserObjectClass:                 "user_object_class",
		UserSearchAttribute:             "user_search_attribute",
		UserSearchFilter:                "(|(cn=$SEARCH_STRING)(sAMAccountName=$SEARCH_STRING))",
	}
	testAuthConfigLdapInterface = map[string]interface{}{
		"access_mode":                        "access",
		"allowed_principal_ids":              []interface{}{"allowed1", "allowed2"},
		"enabled":                            true,
		"servers":                            []interface{}{"server1", "server2"},
		"service_account_distinguished_name": "service_account_distinguished_name",
		"user_search_base":                   "user_search_base",
		"certificate":                        Base64Encode("certificate"),
		"connection_timeout":                 10,
		"group_dn_attribute":                 "group_dn_attribute",
		"group_member_mapping_attribute":     "group_member_mapping_attribute",
		"group_member_user_attribute":        "group_member_user_attribute",
		"group_name_attribute":               "group_name_attribute",
		"group_object_class":                 "group_object_class",
		"group_search_attribute":             "group_search_attribute",
		"group_search_base":                  "group_search_base",
		"group_search_filter":                "(cn=$SEARCH_STRING)",
		"nested_group_membership_enabled":    true,
		"port":                               389,
		"tls":                                true,
		"start_tls":                          true,
		"user_disabled_bit_mask":             0,
		"user_login_attribute":               "user_login_attribute",
		"user_member_attribute":              "user_member_attribute",
		"user_name_attribute":                "user_name_attribute",
		"user_object_class":                  "user_object_class",
		"user_search_attribute":              "user_search_attribute",
		"user_search_filter":                 "(|(cn=$SEARCH_STRING)(sAMAccountName=$SEARCH_STRING))",
	}
}

func TestFlattenAuthConfigLdap(t *testing.T) {

	cases := []struct {
		Input          *managementClient.LdapConfig
		ExpectedOutput map[string]interface{}
	}{
		{
			testAuthConfigLdapConf,
			testAuthConfigLdapInterface,
		},
	}

	for _, tc := range cases {
		output := schema.TestResourceDataRaw(t, authConfigLdapFields(), map[string]interface{}{})
		err := flattenAuthConfigLdap(output, tc.Input)
		if err != nil {
			assert.FailNow(t, "[ERROR] on flattener: %#v", err)
		}
		expectedOutput := map[string]interface{}{}
		for k := range tc.ExpectedOutput {
			expectedOutput[k] = output.Get(k)
		}
		if !reflect.DeepEqual(expectedOutput, tc.ExpectedOutput) {
			assert.FailNow(t, "Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, expectedOutput)
		}
	}
}

func TestExpandAuthConfigLdap(t *testing.T) {

	cases := []struct {
		Input          map[string]interface{}
		ExpectedOutput *managementClient.LdapConfig
	}{
		{
			testAuthConfigLdapInterface,
			testAuthConfigLdapConf,
		},
	}

	for _, tc := range cases {
		inputResourceData := schema.TestResourceDataRaw(t, authConfigLdapFields(), tc.Input)
		output, err := expandAuthConfigLdap(inputResourceData)
		if err != nil {
			assert.FailNow(t, "[ERROR] on expander: %#v", err)
		}
		assert.Equal(t, tc.ExpectedOutput, output, "Unexpected output from expander.")
	}
}
