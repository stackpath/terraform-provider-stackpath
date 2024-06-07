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

func resourceComputeNetworkAllocationClaim() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceComputeNetworkAllocationClaimCreate,
		ReadContext:   resourceComputeNetworkAllocationClaimRead,
		UpdateContext: resourceComputeNetworkAllocationClaimUpdate,
		DeleteContext: resourceComputeNetworkAllocationClaimDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceComputeNetworkAllocationClaimImportState,
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
				Elem:     resourceComputeNetworkAllocationReclaimPolicy(),
			},
			"resource_binding": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     resourceComputeNetworkAllocationResourceBinding(),
			},
			"allocation": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"allocation.0.selector", "allocation.0.template"},
						},
						"selector": {
							Type:          schema.TypeList,
							MaxItems:      1,
							Optional:      true,
							ConflictsWith: []string{"allocation.0.name", "allocation.0.template"},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allocation_class": {
										Type:     schema.TypeString,
										Required: true,
									},
									"match_expressions": {
										Type:     schema.TypeList,
										Required: true,
										Elem:     resourceComputeMatchExpressionSchema(),
									},
								},
							},
						},
						"template": {
							Type:          schema.TypeList,
							MaxItems:      1,
							Optional:      true,
							ConflictsWith: []string{"allocation.0.selector", "allocation.0.name"},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
										Required: true,
										Elem:     resourceComputeNetworkAllocationReclaimPolicy(),
									},
									"selectors": {
										Type:     schema.TypeList,
										Optional: true,
										Elem:     resourceComputeMatchExpressionSchema(),
									},
								},
							},
						},
					},
				},
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
						"allocation": {
							Type:     schema.TypeString,
							Required: true,
						},
						"resource_binding": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem:     resourceComputeNetworkAllocationResourceBinding(),
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

func resourceComputeNetworkAllocationClaimCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	resp, err := config.edgeComputeNetworking.Allocations.CreateAllocationClaim(&allocations.CreateAllocationClaimParams{
		Context: ctx,
		StackID: config.StackID,
		Body: &ipam_models.V1CreateAllocationClaimRequest{
			AllocationClaim: convertComputeNetworkAllocationClaim(data),
		},
	}, nil)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create network allocation claim: %v", NewStackPathError(err)))
	}

	// wait for operation to complete
	if operation, err := waitForIPAMOperationToBeDone(ctx, resp.Payload.Name, config); err != nil {
		return diag.FromErr(err)
	} else if operation.Error != nil {
		// (TODO)- print *ipam_models.GooglerpcStatus in format aligning with NewStackPathError
		return diag.FromErr(fmt.Errorf("network allocation claim operation failed: %v", operation.Error))
	} else {
		// (TODO)- Currently there is an issue in the GetOperation client API generated through
		// swagger which leads to ProtobufAny typed fields like .Response or .Metadata in
		// ipam_models.V1Operation to be always nil even if the REST API response has that data.
		// Hence currently there is no way to get allocation ID during create context.
		// Until we get that working, we are going to use stackID/allocationSlug named string
		// as ID since it is expected that the allocation slug be unique in a stack.
		data.SetId(formatAllocationID(config.StackID, data.Get("slug").(string)))
	}

	return resourceComputeNetworkAllocationClaimRead(ctx, data, meta)
}

func resourceComputeNetworkAllocationClaimRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)

	params := allocations.GetAllocationClaimParams{
		Context: ctx,
	}
	var err error
	id := data.Id()
	if params.StackID, params.AllocationClaimSlug, err = parseAllocationID(id); err != nil {
		return diag.FromErr(fmt.Errorf("failed to parse allocation ID (%s): %w", id, NewStackPathError(err)))
	}

	resp, err := config.edgeComputeNetworking.Allocations.GetAllocationClaim(&params, nil)
	if c, ok := err.(interface{ Code() int }); ok && c.Code() == http.StatusNotFound {
		// Clear out the ID in terraform if the
		// resource no longer exists in the API
		data.SetId("")
		return nil
	} else if err != nil {
		return diag.FromErr(fmt.Errorf("failed to read network allocation claim: %v", NewStackPathError(err)))
	}

	if err := data.Set("name", resp.Payload.AllocationClaim.Name); err != nil {
		return diag.FromErr(fmt.Errorf("error setting name: %v", err))
	}

	if err := data.Set("slug", resp.Payload.AllocationClaim.Slug); err != nil {
		return diag.FromErr(fmt.Errorf("error setting slug: %v", err))
	}

	if err := data.Set("version", resp.Payload.AllocationClaim.Metadata.Version); err != nil {
		return diag.FromErr(fmt.Errorf("error setting version: %v", err))
	}

	if err := data.Set("resource_name", resp.Payload.AllocationClaim.Metadata.ResourceName); err != nil {
		return diag.FromErr(fmt.Errorf("error setting resource name: %v", err))
	}

	if err := data.Set("labels", flattenMetaV1StringMap(resp.Payload.AllocationClaim.Metadata.Labels)); err != nil {
		return diag.FromErr(fmt.Errorf("error setting labels: %v", err))
	}

	if err := data.Set("annotations", flattenMetaV1StringMap(resp.Payload.AllocationClaim.Metadata.Annotations)); err != nil {
		return diag.FromErr(fmt.Errorf("error setting annotations: %v", err))
	}

	if err := data.Set("prefix_length", int(resp.Payload.AllocationClaim.Spec.PrefixLength)); err != nil {
		return diag.FromErr(fmt.Errorf("error setting prefix length: %v", err))
	}

	if err := data.Set("ip_family", string(*resp.Payload.AllocationClaim.Spec.IPFamily)); err != nil {
		return diag.FromErr(fmt.Errorf("error setting ip family: %v", err))
	}

	if err := data.Set("reclaim_policy", flattenComputeNetworkAllocationReclaimPolicy(resp.Payload.AllocationClaim.Spec.ReclaimPolicy)); err != nil {
		return diag.FromErr(fmt.Errorf("error setting reclaim policy: %v", err))
	}

	if err := data.Set("resource_binding", flattenComputeNetworkAllocationResourceBinding(resp.Payload.AllocationClaim.Spec.ResourceBinding)); err != nil {
		return diag.FromErr(fmt.Errorf("error setting resource binding: %v", err))
	}

	if err := data.Set("allocation", flattenComputeNetworkAllocationClaimSpecAllocation(resp.Payload.AllocationClaim.Spec.Allocation)); err != nil {
		return diag.FromErr(fmt.Errorf("error setting allocation claim allocation spec: %v", err))
	}

	if resp.Payload.AllocationClaim.Status != nil {
		if err := data.Set("status", flattenComputeNetworkAllocationClaimStatus(resp.Payload.AllocationClaim.Status)); err != nil {
			return diag.FromErr(fmt.Errorf("error setting status: %v", err))
		}
	}

	return diag.Diagnostics{}
}

func resourceComputeNetworkAllocationClaimUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	allocationClaim := convertComputeNetworkAllocationClaimUpdate(data)

	resp, err := config.edgeComputeNetworking.Allocations.UpdateAllocationClaim(&allocations.UpdateAllocationClaimParams{
		Context:             ctx,
		StackID:             config.StackID,
		AllocationClaimSlug: data.Get("slug").(string),
		Body: &ipam_models.V1UpdateAllocationClaimRequest{
			AllocationClaim: allocationClaim,
		},
	}, nil)
	if c, ok := err.(interface{ Code() int }); ok && c.Code() == http.StatusNotFound {
		// Clear out the ID in terraform if the
		// resource no longer exists in the API
		data.SetId("")
		return diag.Diagnostics{}
	} else if err != nil {
		return diag.FromErr(fmt.Errorf("failed to update network allocation claim: %v", NewStackPathError(err)))
	}

	if _, err := waitForIPAMOperationToBeDone(ctx, resp.Payload.Name, config); err != nil {
		return diag.FromErr(err)
	} else if resp.Payload.Error != nil {
		// (TODO)- print *ipam_models.GooglerpcStatus in format aligning with NewStackPathError
		return diag.FromErr(fmt.Errorf("network allocation claim update operation failed: %v", resp.Payload.Error))
	}

	return resourceComputeNetworkAllocationClaimRead(ctx, data, meta)
}

func resourceComputeNetworkAllocationClaimDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	resp, err := config.edgeComputeNetworking.Allocations.DeleteAllocationClaim(&allocations.DeleteAllocationClaimParams{
		Context:             ctx,
		StackID:             config.StackID,
		AllocationClaimSlug: data.Get("slug").(string),
	}, nil)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to delete network allocation claim: %v", NewStackPathError(err)))
	}

	if _, err := waitForIPAMOperationToBeDone(ctx, resp.Payload.Name, config); err != nil {
		return diag.FromErr(err)
	}

	data.SetId("")
	return diag.Diagnostics{}
}

func resourceComputeNetworkAllocationClaimImportState(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// We expect that to import a resource, the user will pass in the
	// stackID/allocationSlug formatted name of allocation they're attempting to import.
	return []*schema.ResourceData{d}, nil
}
