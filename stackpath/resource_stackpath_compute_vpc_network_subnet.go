package stackpath

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/ipam/ipam_client/virtual_private_cloud"
	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/ipam/ipam_models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceComputeVPCNetworkSubnet() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceComputeVPCNetworkSubnetCreate,
		ReadContext:   resourceComputeVPCNetworkSubnetRead,
		UpdateContext: resourceComputeVPCNetworkSubnetUpdate,
		DeleteContext: resourceComputeVPCNetworkSubnetDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceComputeVPCNetworkSubnetImportState,
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
			"prefix": {
				Type:     schema.TypeString,
				Required: true,
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
			},
		},
	}
}

func resourceComputeVPCNetworkSubnetCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	computeNetworkSubnet := convertComputeNetworkSubnet(data)
	resp, err := config.edgeComputeNetworking.VirtualPrivateCloud.CreateNetworkSubnet(&virtual_private_cloud.CreateNetworkSubnetParams{
		Context:   ctx,
		StackID:   config.StackID,
		NetworkID: computeNetworkSubnet.NetworkID,
		Body: &ipam_models.NetworkCreateNetworkSubnetRequest{
			Subnet: computeNetworkSubnet,
		},
	}, nil)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create network subnet: %w", NewStackPathError(err)))
	}

	data.SetId(formatNetworkSubnetID(resp.Payload.Subnet.NetworkID, resp.Payload.Subnet.ID))
	return resourceComputeVPCNetworkSubnetRead(ctx, data, meta)
}

func resourceComputeVPCNetworkSubnetRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	params := virtual_private_cloud.GetNetworkSubnetParams{
		Context: ctx,
		StackID: config.StackID,
	}
	var err error
	id := data.Id()
	if params.NetworkID, params.SubnetID, err = parseNetworkSubnetID(id); err != nil {
		return diag.FromErr(fmt.Errorf("failed to parse subnet ID (%s): %w", id, NewStackPathError(err)))
	}

	resp, err := config.edgeComputeNetworking.VirtualPrivateCloud.GetNetworkSubnet(&params, nil)
	if c, ok := err.(interface{ Code() int }); ok && c.Code() == http.StatusNotFound {
		// Clear out the ID in terraform if the
		// resource no longer exists in the API
		data.SetId("")
		return diag.FromErr(fmt.Errorf("failed to get subnet (%s)%s/%s: %w", data.Id(), params.NetworkID, params.SubnetID, NewStackPathError(err)))
	} else if err != nil {
		return diag.FromErr(fmt.Errorf("failed to read subnet: %w", NewStackPathError(err)))
	}

	if err := data.Set("name", resp.Payload.Subnet.Name); err != nil {
		return diag.FromErr(fmt.Errorf("error setting name: %w", NewStackPathError(err)))
	}

	if err := data.Set("slug", resp.Payload.Subnet.Slug); err != nil {
		return diag.FromErr(fmt.Errorf("error setting slug: %w", NewStackPathError(err)))
	}

	if err := data.Set("prefix", resp.Payload.Subnet.Prefix); err != nil {
		return diag.FromErr(fmt.Errorf("error setting prefix: %w", NewStackPathError(err)))
	}

	if resp.Payload.Subnet.Metadata != nil {
		if err := data.Set("labels", flattenStringMap(convertIPAMToWorkloadStringMapEntry(resp.Payload.Subnet.Metadata.Labels))); err != nil {
			return diag.FromErr(fmt.Errorf("error setting labels: %w", NewStackPathError(err)))
		}

		if err := data.Set("annotations", flattenStringMap(convertIPAMToWorkloadStringMapEntry(resp.Payload.Subnet.Metadata.Annotations))); err != nil {
			return diag.FromErr(fmt.Errorf("error setting annotations: %w", NewStackPathError(err)))
		}

		if err := data.Set("version", resp.Payload.Subnet.Metadata.Version); err != nil {
			return diag.FromErr(fmt.Errorf("error setting version: %w", NewStackPathError(err)))
		}
	}
	return diag.Diagnostics{}
}

func resourceComputeVPCNetworkSubnetUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	networkSubnet := convertComputeNetworkSubnet(data)
	var err error
	id := data.Id()
	if networkSubnet.NetworkID, networkSubnet.ID, err = parseNetworkSubnetID(id); err != nil {
		return diag.FromErr(fmt.Errorf("failed to parse subnet ID (%s): %w", id, NewStackPathError(err)))
	}

	_, err = config.edgeComputeNetworking.VirtualPrivateCloud.UpdateNetworkSubnet(&virtual_private_cloud.UpdateNetworkSubnetParams{
		Context:   ctx,
		StackID:   config.StackID,
		NetworkID: networkSubnet.NetworkID,
		SubnetID:  networkSubnet.ID,
		Body: &ipam_models.NetworkUpdateNetworkSubnetRequest{
			Subnet: networkSubnet,
		},
	}, nil)
	if c, ok := err.(interface{ Code() int }); ok && c.Code() == http.StatusNotFound {
		// Clear out the ID in terraform if the
		// resource no longer exists in the API
		data.SetId("")
		return diag.Diagnostics{}
	} else if err != nil {
		return diag.FromErr(fmt.Errorf("failed to update subnet: %w", NewStackPathError(err)))
	}

	return resourceComputeVPCNetworkSubnetRead(ctx, data, meta)
}

func resourceComputeVPCNetworkSubnetDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	params := virtual_private_cloud.DeleteNetworkSubnetParams{
		Context: ctx,
		StackID: config.StackID,
	}
	var err error
	id := data.Id()
	if params.NetworkID, params.SubnetID, err = parseNetworkSubnetID(id); err != nil {
		return diag.FromErr(fmt.Errorf("failed to parse subnet ID (%s): %w", id, err))
	}
	_, err = config.edgeComputeNetworking.VirtualPrivateCloud.DeleteNetworkSubnet(&params, nil)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to delete subnet: %w", NewStackPathError(err)))
	}

	data.SetId("")
	return diag.Diagnostics{}
}

func resourceComputeVPCNetworkSubnetImportState(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// We expect that to import a resource, the user will pass in the
	// full UUID of the network subnet they're attempting to import.
	return []*schema.ResourceData{d}, nil
}
