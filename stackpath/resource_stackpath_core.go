package stackpath

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceComputeMatchExpressionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"operator": {
				Type:     schema.TypeString,
				Required: true,
			},
			"values": {
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
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"image": {
				Type:     schema.TypeString,
				Required: true,
			},
			"command": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"args": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"env":              resourceComputeWorkloadEnvVars(),
			"port":             resourceComputeWorkloadPortSchema(),
			"readiness_probe":  resourceComputeWorkloadProbeSchema(),
			"liveness_probe":   resourceComputeWorkloadProbeSchema(),
			"resources":        resourceComputeWorkloadResourcesSchema(),
			"volume_mount":     resourceComputeWorkloadVolumeMountSchema(),
			"security_context": resourceComputeWorkloadSecurityContextSchema(),
		},
	}
}

func resourceComputeWorkloadVirtualMachine() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"image": {
				Type:     schema.TypeString,
				Required: true,
			},
			"port":            resourceComputeWorkloadPortSchema(),
			"liveness_probe":  resourceComputeWorkloadProbeSchema(),
			"readiness_probe": resourceComputeWorkloadProbeSchema(),
			"resources":       resourceComputeWorkloadResourcesSchema(),
			"volume_mount":    resourceComputeWorkloadVolumeMountSchema(),
			"user_data": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceComputeMetadata() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
			},
			"annotations": {
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}

func resourceComputeLocation() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"city": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"city_code": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subdivision": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subdivision_code": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"country": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"country_code": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"region_code": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"continent": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"continent_code": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"latitude": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"longitude": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}

func resourceComputeCondition() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
				Required: true,
			},
			"observed_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_transition_time": {
				Type:     schema.TypeString,
				Required: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Required: true,
			},
			"message": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceComputeNetworkAllocationReclaimPolicy() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"idle_retention_period": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceComputeNetworkAllocationResourceBinding() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
