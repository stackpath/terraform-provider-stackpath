package stackpath

import "github.com/hashicorp/terraform/helper/schema"

func resourceComputeWorkloadInstance() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"metadata": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     resourceComputeMetadata(),
			},
			"location": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     resourceComputeLocation(),
			},
			"external_ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_interface": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"ip_address_aliases": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"gateway": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"virtual_machine": &schema.Schema{
				Type:          schema.TypeList,
				ConflictsWith: []string{"container"},
				MaxItems:      1,
				Optional:      true,
				Elem:          resourceComputeWorkloadVirtualMachine(),
			},
			"container": &schema.Schema{
				Type:          schema.TypeList,
				Optional:      true,
				ConflictsWith: []string{"virtual_machine"},
				Elem:          resourceComputeWorkloadContainer(),
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"phase": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceComputeWorkloadVirtualMachineStatus() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"phase": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceComputeWorkloadContainerStatus() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"phase": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"started_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"finished_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"waiting": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"reason": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"message": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"running": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"started_at": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"terminated": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"reason": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"message": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"started_at": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"finished_at": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"exit_code": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"ready": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"restart_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"container_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
