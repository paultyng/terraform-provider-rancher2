package rancher2

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRancher2CatalogV2() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRancher2CatalogV2Read,

		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ca_bundle": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"git_branch": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"git_repo": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"insecure": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"resource_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secret_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secret_namespace": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_account": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_account_namespace": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"annotations": {
				Type:     schema.TypeMap,
				Computed: true,
			},
			"labels": {
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}

func dataSourceRancher2CatalogV2Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	clusterID := d.Get("cluster_id").(string)
	name := d.Get("name").(string)

	catalog, err := getCatalogV2ByID(meta.(*Config), clusterID, name)
	if err != nil {
		if IsNotFound(err) || IsForbidden(err) {
			log.Printf("[INFO] Catalog V2 %s not found at cluster %s", name, clusterID)
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	return diag.FromErr(flattenCatalogV2(d, catalog))
}
