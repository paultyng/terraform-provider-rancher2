package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

//Schemas

func clusterDriverFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"active": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"builtin": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"url": {
			Type:     schema.TypeString,
			Required: true,
		},
		"actual_url": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"checksum": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"ui_url": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"whitelist_domains": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}

	for k, v := range commonAnnotationLabelFields() {
		s[k] = v
	}

	return s
}
