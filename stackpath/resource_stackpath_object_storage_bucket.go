package stackpath

import (
	"context"
	"fmt"
	"net/http"

	"github.com/antihax/optional"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/api_client"
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
	resp, _, err := config.apiClient.BucketsApi.CreateBucket(context.Background(), config.StackID, api_client.InlineObject{
		Label:  data.Get("label").(string),
		Region: data.Get("region").(string),
	})
	// Handle error
	if err != nil {
		return fmt.Errorf("failed to create object storage bucket: %v", NewStackPathError(err))
	}
	// Assign ID from the response
	data.SetId(resp.Bucket.Id)
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
	resp, httpResponse, err := config.apiClient.BucketsApi.GetBucket(context.Background(), config.StackID, data.Id())
	// Handle error
	if err != nil {
		if httpResponse.StatusCode == http.StatusNotFound {
			// Clear out the ID in terraform if the
			// resource no longer exists in the API
			data.SetId("")
			return nil
		}
		return fmt.Errorf("failed to read object storage bucket: %v", NewStackPathError(err))
	}
	// Set properties
	data.Set("endpoint_url", resp.Bucket.EndpointUrl)
	data.Set("label", resp.Bucket.Label)
	data.Set("region", resp.Bucket.Region)
	data.Set("visibility", resp.Bucket.Visibility)
	return nil
}

func resourceObjectStorageBucketUpdate(data *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	// Update in API
	_, httpResponse, err := config.apiClient.BucketsApi.UpdateBucket(
		context.Background(),
		config.StackID,
		data.Id(),
		api_client.InlineObject1{
			Visibility: data.Get("visibility").(string),
		},
	)
	// Handle error
	if err != nil {
		if httpResponse.StatusCode == http.StatusNotFound {
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
	_, err := config.apiClient.BucketsApi.DeleteBucket(
		context.Background(),
		config.StackID,
		data.Id(),
		&api_client.DeleteBucketOpts{
			ForceDelete: optional.NewBool(true),
		},
	)
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
