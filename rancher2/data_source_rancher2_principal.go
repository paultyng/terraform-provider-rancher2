package rancher2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	managementClient "github.com/rancher/rancher/pkg/client/generated/management/v3"
)

const (
	principalTypeGroup = "group"
	principalTypeUser  = "user"
)

var (
	principalTypes = []string{principalTypeGroup, principalTypeUser}
)

func dataSourceRancher2Principal() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRancher2PrincipalRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      principalTypeUser,
				ValidateFunc: validation.StringInSlice(principalTypes, true),
			},
		},
	}
}

func dataSourceRancher2PrincipalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, err := meta.(*Config).ManagementClient()
	if err != nil {
		return diag.FromErr(err)
	}

	name := d.Get("name").(string)
	principalType := d.Get("type").(string)

	collection, err := client.Principal.List(nil)
	if err != nil {
		return diag.FromErr(err)
	}

	principals, err := client.Principal.CollectionActionSearch(collection, &managementClient.SearchPrincipalsInput{
		Name:          name,
		PrincipalType: principalType,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	count := len(principals.Data)
	if count <= 0 {
		return diag.Errorf("[ERROR] principal \"%s\" of type \"%s\" not found", name, principalType)
	}

	return diag.FromErr(flattenDataSourcePrincipal(d, &principals.Data[0]))
}

func flattenDataSourcePrincipal(d *schema.ResourceData, in *managementClient.Principal) error {
	if in == nil {
		return nil
	}

	d.SetId(in.ID)
	d.Set("name", in.Name)
	d.Set("type", in.PrincipalType)

	return nil
}
