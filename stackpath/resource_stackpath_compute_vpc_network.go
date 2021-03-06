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

func resourceComputeVPCNetwork() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceComputeVPCNetworkCreate,
		ReadContext:   resourceComputeVPCNetworkRead,
		UpdateContext: resourceComputeVPCNetworkUpdate,
		DeleteContext: resourceComputeVPCNetworkDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceComputeVPCNetworkImportState,
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
			"root_subnet": {
				Type:     schema.TypeString,
				Required: true,
			},
			"virtual_network_identifier": {
				Type:     schema.TypeInt,
				Computed: true,
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

func resourceComputeVPCNetworkCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	computeNetwork := convertComputeNetwork(data)
	resp, err := config.edgeComputeNetworking.VirtualPrivateCloud.CreateNetwork(&virtual_private_cloud.CreateNetworkParams{
		Context: ctx,
		StackID: config.StackID,
		Body: &ipam_models.NetworkCreateNetworkRequest{
			Name:       computeNetwork.Name,
			Slug:       computeNetwork.Slug,
			RootSubnet: computeNetwork.RootSubnet,
			Metadata:   computeNetwork.Metadata,
		},
	}, nil)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create network: %v", NewStackPathError(err)))
	}

	data.SetId(resp.Payload.Network.ID)
	return resourceComputeVPCNetworkRead(ctx, data, meta)
}

func resourceComputeVPCNetworkRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)

	resp, err := config.edgeComputeNetworking.VirtualPrivateCloud.GetNetwork(&virtual_private_cloud.GetNetworkParams{
		StackID:   config.StackID,
		NetworkID: data.Id(),
		Context:   ctx,
	}, nil)
	if c, ok := err.(interface{ Code() int }); ok && c.Code() == http.StatusNotFound {
		// Clear out the ID in terraform if the
		// resource no longer exists in the API
		data.SetId("")
		return diag.Diagnostics{}
	} else if err != nil {
		return diag.FromErr(fmt.Errorf("failed to read network: %v", NewStackPathError(err)))
	}

	if err := data.Set("name", resp.Payload.Network.Name); err != nil {
		return diag.FromErr(fmt.Errorf("error setting name: %v", err))
	}

	if err := data.Set("slug", resp.Payload.Network.Slug); err != nil {
		return diag.FromErr(fmt.Errorf("error setting slug: %v", err))
	}

	if err := data.Set("root_subnet", resp.Payload.Network.RootSubnet); err != nil {
		return diag.FromErr(fmt.Errorf("error setting root_subnet: %v", err))
	}

	if err := data.Set("virtual_network_identifier", resp.Payload.Network.VirtualNetworkIdentifier); err != nil {
		return diag.FromErr(fmt.Errorf("error setting virtual_network_identifier: %v", err))
	}

	if resp.Payload.Network.Metadata != nil {
		if err := data.Set("labels", flattenStringMap(convertIPAMToWorkloadStringMapEntry(resp.Payload.Network.Metadata.Labels))); err != nil {
			return diag.FromErr(fmt.Errorf("error setting labels: %v", err))
		}

		if err := data.Set("annotations", flattenStringMap(convertIPAMToWorkloadStringMapEntry(resp.Payload.Network.Metadata.Annotations))); err != nil {
			return diag.FromErr(fmt.Errorf("error setting annotations: %v", err))
		}

		if err := data.Set("version", resp.Payload.Network.Metadata.Version); err != nil {
			return diag.FromErr(fmt.Errorf("error setting version: %v", err))
		}
	}
	return diag.Diagnostics{}
}

func resourceComputeVPCNetworkUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	network := convertComputeNetwork(data)
	network.ID = data.Id()

	_, err := config.edgeComputeNetworking.VirtualPrivateCloud.UpdateNetwork(&virtual_private_cloud.UpdateNetworkParams{
		Context:   ctx,
		StackID:   config.StackID,
		NetworkID: data.Id(),
		Body: &ipam_models.NetworkUpdateNetworkRequest{
			Name:     network.Name,
			Metadata: network.Metadata,
		},
	}, nil)
	if c, ok := err.(interface{ Code() int }); ok && c.Code() == http.StatusNotFound {
		// Clear out the ID in terraform if the
		// resource no longer exists in the API
		data.SetId("")
		return diag.Diagnostics{}
	} else if err != nil {
		return diag.FromErr(fmt.Errorf("failed to update network: %v", NewStackPathError(err)))
	}

	return resourceComputeVPCNetworkRead(ctx, data, meta)
}

func resourceComputeVPCNetworkDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	_, err := config.edgeComputeNetworking.VirtualPrivateCloud.DeleteNetwork(&virtual_private_cloud.DeleteNetworkParams{
		Context:   ctx,
		StackID:   config.StackID,
		NetworkID: data.Id(),
	}, nil)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to delete network: %v", NewStackPathError(err)))
	}

	data.SetId("")
	return diag.Diagnostics{}
}

func resourceComputeVPCNetworkImportState(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// We expect that to import a resource, the user will pass in the
	// full UUID of the network they're attempting to import.
	return []*schema.ResourceData{d}, nil
}
