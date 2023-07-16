package stackpath

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stackpath/terraform-provider-stackpath/stackpath/api/dns"
)

func resourceDNSRecord() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDNSRecordCreate,
		ReadContext:   resourceDNSRecordRead,
		UpdateContext: resourceDNSRecordUpdate,
		DeleteContext: resourceDNSRecordDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceDNSRecordImportState,
		},
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"class": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"data": {
				Type: schema.TypeString,
			},
			"weight": {
				Type: schema.TypeInt,
			},
			"labels": {
				Type: schema.TypeMap,
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func strPtr(s string) *string {
	return &s
}

func int32Ptr(i int32) *int32 {
	return &i
}

func mapStringPtr(m map[string]interface{}) *map[string]string {
	if m == nil {
		return nil
	}
	res := make(map[string]string)
	for k, v := range m {
		res[k] = v.(string)
	}
	return &res
}

func zoneRecTypePtr(s string) *dns.ZoneRecordType {
	t := dns.ZoneRecordType(s)
	return &t
}

func resourceDNSRecordCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	// Create in API
	msg := dns.ZoneUpdateZoneRecordMessage{
		Name:   strPtr(data.Get("name").(string)),
		Type:   zoneRecTypePtr(data.Get("type").(string)),
		Ttl:    int32Ptr(data.Get("ttl").(int32)),
		Data:   strPtr(data.Get("data").(string)),
		Weight: int32Ptr(data.Get("weight").(int32)),
		Labels: mapStringPtr(data.Get("labels").(map[string]interface{})),
	}

	resp, _, err := config.dns.ResourceRecordsAPI.CreateZoneRecord(
		ctx, config.StackID, data.Get("zone_id").(string)).ZoneUpdateZoneRecordMessage(msg).Execute()
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create zone record: %v", NewStackPathError(err)))
	}

	// Assign ID from the response
	record := resp.GetRecord()
	data.SetId(record.GetId())

	return resourceDNSRecordRead(ctx, data, meta)
}

func resourceDNSRecordRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	resp, _, err := config.dns.ResourceRecordsAPI.GetZoneRecord(
		ctx,
		config.StackID,
		data.Get("zone_id").(string),
		data.Id()).Execute()
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to get zone record: %v", NewStackPathError(err)))
	}

	record := resp.GetRecord()

	if err := data.Set("name", record.Name); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set record name: %v", err))
	}

	// Set properties
	if err := data.Set("type", record.Type); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set record type: %v", err))
	}

	if err := data.Set("class", record.Class); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set record class: %v", err))
	}

	if err := data.Set("ttl", record.Ttl); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set record ttl: %v", err))
	}

	if err := data.Set("data", record.Data); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set record data: %v", err))
	}

	if err := data.Set("weight", record.Weight); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set record weight: %v", err))
	}

	if err := data.Set("labels", record.Labels); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set record labels: %v", err))
	}

	if err := data.Set("created", record.Created); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set record created: %v", err))
	}

	if err := data.Set("updated", record.Updated); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set record updated: %v", err))
	}

	return diag.Diagnostics{}
}

func resourceDNSRecordUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	// Create in API
	msg := dns.ZoneUpdateZoneRecordMessage{
		Name:   strPtr(data.Get("name").(string)),
		Type:   zoneRecTypePtr(data.Get("type").(string)),
		Ttl:    int32Ptr(data.Get("ttl").(int32)),
		Data:   strPtr(data.Get("data").(string)),
		Weight: int32Ptr(data.Get("weight").(int32)),
		Labels: mapStringPtr(data.Get("labels").(map[string]interface{})),
	}

	_, _, err := config.dns.ResourceRecordsAPI.UpdateZoneRecord(
		ctx,
		config.StackID,
		data.Get("zone_id").(string),
		data.Id()).ZoneUpdateZoneRecordMessage(msg).Execute()
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to update zone record: %v", NewStackPathError(err)))
	}

	return resourceDNSRecordRead(ctx, data, meta)
}

func resourceDNSRecordDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	_, err := config.dns.ResourceRecordsAPI.DeleteZoneRecord(
		ctx,
		config.StackID,
		data.Get("zone_id").(string),
		data.Id()).Execute()
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to delete zone record: %v", NewStackPathError(err)))
	}
	return diag.Diagnostics{}
}

func resourceDNSRecordImportState(ctx context.Context, data *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// We expect that to import a resource, the user will pass in the
	// full UUID of the bucket they're attempting to import.
	// Update data from the read method and return
	if err := resourceDNSRecordRead(ctx, data, meta); err != nil {
		return nil, fmt.Errorf("failed to read dns record: %v", err)
	}
	return []*schema.ResourceData{data}, nil
}
