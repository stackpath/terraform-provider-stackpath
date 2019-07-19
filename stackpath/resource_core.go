package stackpath

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceComputeMatchExpressionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"operator": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"values": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceComputeWorkloadContainer() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"image": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"command": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"env": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"secret_value": &schema.Schema{
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
					},
				},
			},
			"port":            resourceComputeWorkloadPortSchema(),
			"readiness_probe": resourceComputeWorkloadProbeSchema(),
			"liveness_probe":  resourceComputeWorkloadProbeSchema(),
			"resources":       resourceComputeWorkloadResourcesSchema(),
			"volume_mount":    resourceComputeWorkloadVolumeMountSchema(),
		},
	}
}

func resourceComputeWorkloadVirtualMachine() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"image": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"port":            resourceComputeWorkloadPortSchema(),
			"liveness_probe":  resourceComputeWorkloadProbeSchema(),
			"readiness_probe": resourceComputeWorkloadProbeSchema(),
			"resources":       resourceComputeWorkloadResourcesSchema(),
			"volume_mount":    resourceComputeWorkloadVolumeMountSchema(),
			"user_data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceComputeMetadata() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"labels": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"annotations": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}

func resourceComputeLocation() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"city": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"city_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subdivision": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subdivision_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"country": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"country_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"region_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"continent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"continent_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"latitude": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"longitude": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}
