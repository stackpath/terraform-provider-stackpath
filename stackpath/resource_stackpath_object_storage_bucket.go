package stackpath

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/api/storage/storage_client/buckets"
	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/api/storage/storage_models"
)

func resourceObjectStorageBucket() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectStorageBucketCreate,
		Read:   resourceObjectStorageBucketRead,
		Update: resourceObjectStorageBucketUpdate,
		Delete: resourceObjectStorageBucketDelete,
		Importer: &schema.ResourceImporter{
			State: resourceObjectStorageBucketImportState,
		},
		Schema: map[string]*schema.Schema{
			"label": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"region": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"visibility": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "PRIVATE",
				ValidateFunc: validateObjectStorageBucketVisibility,
			},
			"endpoint_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceObjectStorageBucketCreate(data *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	// Create in API
	resp, err := config.objectStorage.Buckets.CreateBucket(&buckets.CreateBucketParams{
		Body: &storage_models.StorageCreateBucketRequest{
			Label:  data.Get("label").(string),
			Region: data.Get("region").(string),
		},
		StackID: config.StackID,
		Context: context.Background(),
	}, nil)
	// Handle error
	if err != nil {
		return fmt.Errorf("failed to create object storage bucket: %v", NewStackPathError(err))
	}
	// Assign ID from the response
	data.SetId(resp.Payload.Bucket.ID)
	// Run update if visibility is set to PUBLIC
	if data.Get("visibility").(string) != "PRIVATE" {
		if err := resourceObjectStorageBucketUpdate(data, meta); err != nil {
			return fmt.Errorf("failed to set object storage bucket visibility: %v", NewStackPathError(err))
		}
	}
	// Return read
	return resourceObjectStorageBucketRead(data, meta)
}

func resourceObjectStorageBucketRead(data *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	// Read from API
	resp, err := config.objectStorage.Buckets.GetBucket(&buckets.GetBucketParams{
		BucketID: data.Id(),
		StackID:  config.StackID,
		Context:  context.Background(),
	}, nil)
	// Handle error
	if err != nil {
		if c, ok := err.(interface{ Code() int }); ok && c.Code() == http.StatusNotFound {
			// Clear out the ID in terraform if the
			// resource no longer exists in the API
			data.SetId("")
			return nil
		}
		return fmt.Errorf("failed to read object storage bucket: %v", NewStackPathError(err))
	}
	// Set properties
	if err := data.Set("endpoint_url", resp.GetPayload().Bucket.EndpointURL); err != nil {
		return fmt.Errorf("failed to set object storage bucket endpoint_url: %v", err)
	}

	if err := data.Set("label", resp.GetPayload().Bucket.Label); err != nil {
		return fmt.Errorf("failed to set object storage bucket label: %v", err)
	}

	if err := data.Set("region", resp.GetPayload().Bucket.Region); err != nil {
		return fmt.Errorf("failed to set object storage bucket region: %v", err)
	}

	if err := data.Set("visibility", resp.GetPayload().Bucket.Visibility); err != nil {
		return fmt.Errorf("failed to set object storage bucket visibility: %v", err)
	}

	return nil
}

func resourceObjectStorageBucketUpdate(data *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	visibility := storage_models.StorageBucketVisibilityPRIVATE
	if strings.ToUpper(data.Get("visibility").(string)) == "PUBLIC" {
		visibility = storage_models.StorageBucketVisibilityPUBLIC
	}

	// Update in API
	_, err := config.objectStorage.Buckets.UpdateBucket(&buckets.UpdateBucketParams{
		BucketID: data.Id(),
		Context:  context.Background(),
		StackID:  config.StackID,
		Body: &storage_models.StorageUpdateBucketRequest{
			Visibility: visibility,
		},
	}, nil)
	// Handle error
	if err != nil {
		if c, ok := err.(interface{ Code() int }); ok && c.Code() == http.StatusNotFound {
			// Clear out the ID in terraform if the
			// resource no longer exists in the API
			data.SetId("")
			return nil
		}
		return fmt.Errorf("failed to update object storage bucket: %v", NewStackPathError(err))
	}
	// Return read
	return resourceObjectStorageBucketRead(data, meta)
}

func resourceObjectStorageBucketDelete(data *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	// Delete in API
	_, err := config.objectStorage.Buckets.DeleteBucket(&buckets.DeleteBucketParams{
		BucketID: data.Id(),
		Context:  context.Background(),
		StackID:  config.StackID,
	}, nil)
	// Handle error
	if err != nil {
		return fmt.Errorf("failed to delete object storage bucket: %v", NewStackPathError(err))
	}
	// Unset ID
	data.SetId("")
	return nil
}

func resourceObjectStorageBucketImportState(data *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// We expect that to import a resource, the user will pass in the
	// full UUID of the bucket they're attempting to import.
	// Update data from the read method and return
	if err := resourceObjectStorageBucketRead(data, meta); err != nil {
		return nil, fmt.Errorf("failed to read storage bucket: %v", err)
	}
	return []*schema.ResourceData{data}, nil
}
