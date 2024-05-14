package stackpath

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/ipam/ipam_client/allocations"
	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/ipam/ipam_models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceComputeNetworkAllocation() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceComputeNetworkAllocationCreate,
		ReadContext:   resourceComputeNetworkAllocationRead,
		UpdateContext: resourceComputeNetworkAllocationUpdate,
		DeleteContext: resourceComputeNetworkAllocationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceComputeNetworkAllocationImportState,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"annotations": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"allocation_class": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ip_family": {
				Type:     schema.TypeString,
				Required: true,
			},
			"prefix_length": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"reclaim_policy": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
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
				},
			},
			"selectors": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceComputeMatchExpressionSchema(),
			},
			"status": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"prefix": {
							Type:     schema.TypeString,
							Required: true,
						},
						"parent_allocation": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"conditions": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     resourceComputeCondition(),
						},
					},
				},
			},
		},
	}
}

func resourceComputeNetworkAllocationCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	resp, err := config.edgeComputeNetworking.Allocations.CreateAllocation(&allocations.CreateAllocationParams{
		Context: ctx,
		StackID: config.StackID,
		Body: &ipam_models.V1CreateAllocationRequest{
			Allocation: convertComputeNetworkAllocation(data),
		},
	}, nil)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create network allocation: %v", NewStackPathError(err)))
	}

	// wait for operation to complete
	if operation, err := waitForIPAMOperationToBeDone(ctx, resp.Payload.Name, config); err != nil {
		return diag.FromErr(err)
	} else if operation.Error != nil {
		// (TODO)- print *ipam_models.GooglerpcStatus in format aligning with NewStackPathError
		return diag.FromErr(fmt.Errorf("network allocation operation failed: %v", operation.Error))
	} else if allocation, err := unmarshalProtobufAny(operation.Response); err != nil {
		return diag.FromErr(err)
	} else if allocation != nil {
		data.SetId(allocation["id"].(string))
	}

	return resourceComputeNetworkAllocationRead(ctx, data, meta)
}

func resourceComputeNetworkAllocationRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)

	resp, err := config.edgeComputeNetworking.Allocations.GetAllocation(&allocations.GetAllocationParams{
		StackID:        config.StackID,
		AllocationSlug: data.Get("slug").(string),
		Context:        ctx,
	}, nil)
	if c, ok := err.(interface{ Code() int }); ok && c.Code() == http.StatusNotFound {
		// Clear out the ID in terraform if the
		// resource no longer exists in the API
		data.SetId("")
		return nil
	} else if err != nil {
		return diag.FromErr(fmt.Errorf("failed to read network allocation: %v", NewStackPathError(err)))
	}

	if err := data.Set("name", resp.Payload.Allocation.Name); err != nil {
		return diag.FromErr(fmt.Errorf("error setting name: %v", err))
	}

	if err := data.Set("slug", resp.Payload.Allocation.Slug); err != nil {
		return diag.FromErr(fmt.Errorf("error setting slug: %v", err))
	}

	if err := data.Set("version", resp.Payload.Allocation.Metadata.Version); err != nil {
		return diag.FromErr(fmt.Errorf("error setting version: %v", err))
	}

	if err := data.Set("resource_name", resp.Payload.Allocation.Metadata.ResourceName); err != nil {
		return diag.FromErr(fmt.Errorf("error setting resource name: %v", err))
	}

	if err := data.Set("labels", flattenMetaV1StringMap(resp.Payload.Allocation.Metadata.Labels)); err != nil {
		return diag.FromErr(fmt.Errorf("error setting labels: %v", err))
	}

	if err := data.Set("annotations", flattenMetaV1StringMap(resp.Payload.Allocation.Metadata.Annotations)); err != nil {
		return diag.FromErr(fmt.Errorf("error setting annotations: %v", err))
	}

	if err := data.Set("allocation_class", resp.Payload.Allocation.Spec.AllocationClass); err != nil {
		return diag.FromErr(fmt.Errorf("error setting allocation class: %v", err))
	}

	if err := data.Set("prefix_length", resp.Payload.Allocation.Spec.PrefixLength); err != nil {
		return diag.FromErr(fmt.Errorf("error setting prefix length: %v", err))
	}

	if err := data.Set("ip_family", resp.Payload.Allocation.Spec.IPFamily); err != nil {
		return diag.FromErr(fmt.Errorf("error setting ip family: %v", err))
	}

	if err := data.Set("version", flattenComputeNetworkAllocationReclaimPolicy(resp.Payload.Allocation.Spec.ReclaimPolicy)); err != nil {
		return diag.FromErr(fmt.Errorf("error setting reclaim policy: %v", err))
	}

	if err := data.Set("selectors", flattenComputeMetaV1MatchExpressionsOrdered("selectors", data, (resp.Payload.Allocation.Spec.Selectors))); err != nil {
		return diag.FromErr(fmt.Errorf("error setting selectors: %v", err))
	}

	if resp.Payload.Allocation.Status != nil {
		if err := data.Set("status", flattenComputeNetworkAllocationStatus(resp.Payload.Allocation.Status)); err != nil {
			return diag.FromErr(fmt.Errorf("error setting allocation status: %v", err))
		}
	}

	return diag.Diagnostics{}
}

func resourceComputeNetworkAllocationUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	allocation := convertComputeNetworkAllocation(data)
	allocation.ID = data.Id()

	resp, err := config.edgeComputeNetworking.Allocations.UpdateAllocation(&allocations.UpdateAllocationParams{
		Context:        ctx,
		StackID:        config.StackID,
		AllocationSlug: data.Get("slug").(string),
		Body: &ipam_models.V1UpdateAllocationRequest{
			Allocation: allocation,
		},
	}, nil)
	if c, ok := err.(interface{ Code() int }); ok && c.Code() == http.StatusNotFound {
		// Clear out the ID in terraform if the
		// resource no longer exists in the API
		data.SetId("")
		return diag.Diagnostics{}
	} else if err != nil {
		return diag.FromErr(fmt.Errorf("failed to update network allocation: %v", NewStackPathError(err)))
	}

	if _, err := waitForIPAMOperationToBeDone(ctx, resp.Payload.Name, config); err != nil {
		return diag.FromErr(err)
	} else if resp.Payload.Error != nil {
		// (TODO)- print *ipam_models.GooglerpcStatus in format aligning with NewStackPathError
		return diag.FromErr(fmt.Errorf("network allocation update operation failed: %v", resp.Payload.Error))
	}

	return resourceComputeNetworkAllocationRead(ctx, data, meta)
}

func resourceComputeNetworkAllocationDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	resp, err := config.edgeComputeNetworking.Allocations.DeleteAllocation(&allocations.DeleteAllocationParams{
		Context:        ctx,
		StackID:        config.StackID,
		AllocationSlug: data.Get("slug").(string),
	}, nil)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to delete network allocation: %v", NewStackPathError(err)))
	}

	if _, err := waitForIPAMOperationToBeDone(ctx, resp.Payload.Name, config); err != nil {
		return diag.FromErr(err)
	}

	data.SetId("")
	return diag.Diagnostics{}
}

func resourceComputeNetworkAllocationImportState(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// We expect that to import a resource, the user will pass in the
	// full UUID of the network policy they're attempting to import.
	return []*schema.ResourceData{d}, nil
}
