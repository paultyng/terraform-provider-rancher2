package rancher2

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRancher2ConfigMapV2() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRancher2ConfigMapV2Read,

		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},
			"immutable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"data": {
				Type:     schema.TypeMap,
				Computed: true,
			},
			"resource_version": {
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

func dataSourceRancher2ConfigMapV2Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	clusterID := d.Get("cluster_id").(string)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)
	rancherID := namespace + "/" + name
	d.SetId(clusterID + configMapV2ClusterIDsep + rancherID)

	configMap, err := getConfigMapV2ByID(meta.(*Config), clusterID, rancherID)
	if err != nil {
		if IsNotFound(err) || IsForbidden(err) {
			log.Printf("[INFO] ConfigMap V2 %s not found at cluster %s", rancherID, clusterID)
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	return diag.FromErr(flattenConfigMapV2(d, configMap))
}
