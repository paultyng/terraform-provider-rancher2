package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const (
	DefaultProcMount  string = "Default"
	UnmaskedProcMount string = "Unmasked"
)

var (
	ProcMountTypes = []string{
		DefaultProcMount,
		UnmaskedProcMount,
	}
)

func podSecurityPolicyProcMountTypeFields() *schema.Schema {
	s := &schema.Schema{
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice(ProcMountTypes, true),
	}

	return s
}
