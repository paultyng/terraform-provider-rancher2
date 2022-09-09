package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const AuthConfigKeyCloakName = "keycloak"

//Schemas

func authConfigKeyCloakFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"display_name_field": {
			Type:     schema.TypeString,
			Required: true,
		},
		"groups_field": {
			Type:     schema.TypeString,
			Required: true,
		},
		"idp_metadata_content": {
			Type:      schema.TypeString,
			Required:  true,
			Sensitive: true,
			StateFunc: TrimSpace,
		},
		"rancher_api_host": {
			Type:     schema.TypeString,
			Required: true,
		},
		"sp_cert": {
			Type:      schema.TypeString,
			Required:  true,
			StateFunc: TrimSpace,
		},
		"sp_key": {
			Type:      schema.TypeString,
			Required:  true,
			Sensitive: true,
			StateFunc: TrimSpace,
		},
		"uid_field": {
			Type:     schema.TypeString,
			Required: true,
		},
		"user_name_field": {
			Type:     schema.TypeString,
			Required: true,
		},
		"entity_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}

	for k, v := range authConfigFields() {
		s[k] = v
	}

	return s
}
