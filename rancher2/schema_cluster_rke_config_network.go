package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

const (
	networkPluginAciName     = "aci"
	networkPluginCalicoName  = "calico"
	networkPluginCanalName   = "canal"
	networkPluginFlannelName = "flannel"
	networkPluginNonelName   = "none"
	networkPluginWeaveName   = "weave"
)

var (
	networkPluginDefault = networkPluginCanalName
	networkPluginList    = []string{
		networkPluginAciName,
		networkPluginCalicoName,
		networkPluginCanalName,
		networkPluginFlannelName,
		networkPluginNonelName,
		networkPluginWeaveName,
	}
)

//Schemas

func clusterRKEConfigNetworkAciFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"aep": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"apic_hosts": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"apic_refresh_time": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"apic_user_crt": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"apic_user_key": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"apic_user_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"capic": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"controller_log_level": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"drop_log_enable": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"extern_dynamic": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"enable_endpoint_slice": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"encap_type": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"ep_registry": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"gbp_pod_subnet": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"host_agent_log_level": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"image_pull_policy": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"image_pull_secret": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"infra_vlan": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"install_istio": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"istio_profile": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"kafka_brokers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"kafka_client_crt": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"kafka_client_key": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"kube_api_vlan": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"l3out": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"l3out_external_networks": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"max_nodes_svc_graph": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"mcast_range_end": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"mcast_range_start": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"no_priority_class": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"node_subnet": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"ovs_memory_limit": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"opflex_log_level": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"opflex_client_ssl": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"opflex_mode": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"opflex_server_port": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"overlay_vrf_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"pbr_tracking_non_snat": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"pod_subnet_chunk_size": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"run_gbp_container": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"run_opflex_server_container": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"node_svc_subnet": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"service_monitor_interval": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"service_vlan": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"snat_contract_scope": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"snat_namespace": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"snat_port_range_end": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"snat_port_range_start": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"snat_ports_per_node": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"extern_static": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"subnet_domain_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"system_id": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"tenant": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"token": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"use_aci_anywhere_crd": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"use_aci_cni_priority_class": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"use_host_netns_volume": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"use_opflex_server_volume": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"use_privileged_container": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"vrf_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"vrf_tenant": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"vmm_controller": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"vmm_domain": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
	return s
}

func clusterRKEConfigNetworkCalicoFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"cloud_provider": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
	return s
}

func clusterRKEConfigNetworkCanalFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"iface": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
	return s
}

func clusterRKEConfigNetworkFlannelFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"iface": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
	return s
}

func clusterRKEConfigNetworkWeaveFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"password": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
	return s
}

func clusterRKEConfigNetworkFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"aci_network_provider": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: clusterRKEConfigNetworkAciFields(),
			},
		},
		"calico_network_provider": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: clusterRKEConfigNetworkCalicoFields(),
			},
		},
		"canal_network_provider": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: clusterRKEConfigNetworkCanalFields(),
			},
		},
		"flannel_network_provider": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: clusterRKEConfigNetworkFlannelFields(),
			},
		},
		"weave_network_provider": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: clusterRKEConfigNetworkWeaveFields(),
			},
		},
		"mtu": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      0,
			ValidateFunc: validation.IntBetween(0, 9000),
		},
		"options": {
			Type:     schema.TypeMap,
			Optional: true,
			Computed: true,
		},
		"plugin": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validation.StringInSlice(networkPluginList, true),
		},
		"tolerations": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Network add-on tolerations",
			Elem: &schema.Resource{
				Schema: tolerationFields(),
			},
		},
	}
	return s
}
