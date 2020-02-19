package stackpath

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/api/object_storage/client/buckets"
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
			"label": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"visibility": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "PRIVATE",
				ValidateFunc: validateObjectStorageBucketVisibility,
			},
			"endpoint_url": &schema.Schema{
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
		Body: buckets.CreateBucketBody{
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
		resourceObjectStorageBucketUpdate(data, meta)
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
	data.Set("endpoint_url", resp.GetPayload().Bucket.EndpointURL)
	data.Set("label", resp.GetPayload().Bucket.Label)
	data.Set("region", resp.GetPayload().Bucket.Region)
	data.Set("visibility", resp.GetPayload().Bucket.Visibility)
	return nil
}

func resourceObjectStorageBucketUpdate(data *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	visibility := data.Get("visibility").(string)
	// Update in API
	_, err := config.objectStorage.Buckets.UpdateBucket(&buckets.UpdateBucketParams{
		BucketID: data.Id(),
		Context:  context.Background(),
		StackID:  config.StackID,
		Body: buckets.UpdateBucketBody{
			Visibility: &visibility,
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
	resourceObjectStorageBucketRead(data, meta)
	return []*schema.ResourceData{data}, nil
}
