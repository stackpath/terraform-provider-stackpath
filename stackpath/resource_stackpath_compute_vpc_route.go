package stackpath

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/stackpath/terraform-provider-stackpath/stackpath/api/ipam/ipam_client/virtual_private_cloud"
	"github.com/stackpath/terraform-provider-stackpath/stackpath/api/ipam/ipam_models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceComputeVPCRoute() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceComputeVPCRouteCreate,
		ReadContext:   resourceComputeVPCRouteRead,
		UpdateContext: resourceComputeVPCRouteUpdate,
		DeleteContext: resourceComputeVPCRouteDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceComputeVPCRouteImportState,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"network_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"destination_prefixes": {
				Required: true,
				MinItems: 1,
				Type:     schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"gateway_selectors": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_selectors": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
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
										Required: true,
										MinItems: 1,
										Type:     schema.TypeList,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
			"version": {
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
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// Ignore the default network-slug annotation
					return k == "annotations.ipam.platform.stackpath.net/network-slug" || k == "annotations.%"
				},
			},
		},
	}
}

func resourceComputeVPCRouteCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	computeRoute := convertComputeRoute(data)
	resp, err := config.edgeComputeNetworking.VirtualPrivateCloud.CreateRoute(&virtual_private_cloud.CreateRouteParams{
		Context:   ctx,
		StackID:   config.StackID,
		NetworkID: computeRoute.NetworkID,
		Body: &ipam_models.NetworkCreateRouteRequest{
			Route: computeRoute,
		},
	}, nil)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create route: %w", NewStackPathError(err)))
	}

	data.SetId(formatRouteID(resp.Payload.Route.NetworkID, resp.Payload.Route.ID))
	return resourceComputeVPCRouteRead(ctx, data, meta)
}

func resourceComputeVPCRouteRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	params := virtual_private_cloud.GetRouteParams{
		Context: ctx,
		StackID: config.StackID,
	}
	var err error
	id := data.Id()
	if params.NetworkID, params.RouteID, err = parseRouteID(id); err != nil {
		return diag.FromErr(fmt.Errorf("failed to parse route ID (%s): %w", id, NewStackPathError(err)))
	}

	resp, err := config.edgeComputeNetworking.VirtualPrivateCloud.GetRoute(&params, nil)
	if c, ok := err.(interface{ Code() int }); ok && c.Code() == http.StatusNotFound {
		// Clear out the ID in terraform if the
		// resource no longer exists in the API
		data.SetId("")
		return diag.FromErr(fmt.Errorf("failed to get route (%s)%s/%s: %w", data.Id(), params.NetworkID, params.RouteID, NewStackPathError(err)))
	} else if err != nil {
		return diag.FromErr(fmt.Errorf("failed to read route: %w", NewStackPathError(err)))
	}

	if err := data.Set("name", resp.Payload.Route.Name); err != nil {
		return diag.FromErr(fmt.Errorf("error setting name: %w", NewStackPathError(err)))
	}

	if err := data.Set("slug", resp.Payload.Route.Slug); err != nil {
		return diag.FromErr(fmt.Errorf("error setting slug: %w", NewStackPathError(err)))
	}

	if err := data.Set("destination_prefixes", flattenStringArray(resp.Payload.Route.DestinationPrefixes)); err != nil {
		return diag.FromErr(fmt.Errorf("error setting destination_prefixes: %w", NewStackPathError(err)))
	}
	if err := data.Set("gateway_selectors", flattenGatewaySelectors(resp.Payload.Route.GatewaySelectors)); err != nil {
		return diag.FromErr(fmt.Errorf("error setting gateway_selectors: %w", NewStackPathError(err)))
	}

	if resp.Payload.Route.Metadata != nil {
		if err := data.Set("labels", flattenStringMap(convertIPAMToWorkloadStringMapEntry(resp.Payload.Route.Metadata.Labels))); err != nil {
			return diag.FromErr(fmt.Errorf("error setting labels: %w", NewStackPathError(err)))
		}

		if err := data.Set("annotations", flattenStringMap(convertIPAMToWorkloadStringMapEntry(resp.Payload.Route.Metadata.Annotations))); err != nil {
			return diag.FromErr(fmt.Errorf("error setting annotations: %w", NewStackPathError(err)))
		}

		if err := data.Set("version", resp.Payload.Route.Metadata.Version); err != nil {
			return diag.FromErr(fmt.Errorf("error setting version: %w", NewStackPathError(err)))
		}
	}
	return diag.Diagnostics{}
}

func resourceComputeVPCRouteUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	route := convertComputeRoute(data)
	var err error
	id := data.Id()
	if route.NetworkID, route.ID, err = parseRouteID(id); err != nil {
		return diag.FromErr(fmt.Errorf("failed to parse route ID (%s): %w", id, NewStackPathError(err)))
	}

	_, err = config.edgeComputeNetworking.VirtualPrivateCloud.UpdateRoute(&virtual_private_cloud.UpdateRouteParams{
		Context:   ctx,
		StackID:   config.StackID,
		NetworkID: route.NetworkID,
		RouteID:   route.ID,
		Body: &ipam_models.NetworkUpdateRouteRequest{
			Route: route,
		},
	}, nil)
	if c, ok := err.(interface{ Code() int }); ok && c.Code() == http.StatusNotFound {
		// Clear out the ID in terraform if the
		// resource no longer exists in the API
		data.SetId("")
		return diag.Diagnostics{}
	} else if err != nil {
		return diag.FromErr(fmt.Errorf("failed to update route: %w", NewStackPathError(err)))
	}

	return resourceComputeVPCRouteRead(ctx, data, meta)
}

func resourceComputeVPCRouteDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	params := virtual_private_cloud.DeleteRouteParams{
		Context: ctx,
		StackID: config.StackID,
	}
	var err error
	id := data.Id()
	if params.NetworkID, params.RouteID, err = parseRouteID(id); err != nil {
		return diag.FromErr(fmt.Errorf("failed to parse route ID (%s): %w", id, err))
	}
	_, err = config.edgeComputeNetworking.VirtualPrivateCloud.DeleteRoute(&params, nil)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to delete route: %w", NewStackPathError(err)))
	}

	data.SetId("")
	return diag.Diagnostics{}
}

func resourceComputeVPCRouteImportState(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// We expect that to import a resource, the user will pass in the
	// full UUID of the network they're attempting to import.
	return []*schema.ResourceData{d}, nil
}
