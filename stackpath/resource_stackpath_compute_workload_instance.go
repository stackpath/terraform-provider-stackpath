package stackpath

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceComputeWorkloadInstance() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"metadata": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     resourceComputeMetadata(),
			},
			"location": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     resourceComputeLocation(),
			},
			"external_ip_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_interface": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network": {
							Type:     schema.TypeString,
							Required: true,
						},
						"ip_address": {
							Type:     schema.TypeString,
							Required: true,
						},
						"ip_address_aliases": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"gateway": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"virtual_machine": {
				Type:          schema.TypeList,
				ConflictsWith: []string{"container"},
				MaxItems:      1,
				Optional:      true,
				Elem:          resourceComputeWorkloadVirtualMachine(),
			},
			"container": {
				Type:          schema.TypeList,
				Optional:      true,
				ConflictsWith: []string{"virtual_machine"},
				Elem:          resourceComputeWorkloadContainer(),
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"message": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"phase": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceComputeWorkloadVirtualMachineStatus() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"phase": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"message": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceComputeWorkloadContainerStatus() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"phase": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"started_at": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"finished_at": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"waiting": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"reason": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"message": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"running": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"started_at": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"terminated": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"reason": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"message": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"started_at": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"finished_at": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"exit_code": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"ready": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"restart_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"container_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
