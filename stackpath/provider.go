package stackpath

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns the configured provider for managing StackPath resources.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"client_id": {
				Type:        schema.TypeString,
				Sensitive:   true,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("STACKPATH_CLIENT_ID", ""),
			},
			"client_secret": {
				Type:        schema.TypeString,
				Sensitive:   true,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("STACKPATH_CLIENT_SECRET", ""),
			},
			"access_token": {
				Type:      schema.TypeString,
				Sensitive: true,
				Optional:  true,
			},
			"stack_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("STACKPATH_STACK_ID", ""),
			},
			"base_url": {
				Type:     schema.TypeString,
				Optional: true,
				// Default to the official StackPath API
				DefaultFunc: schema.EnvDefaultFunc("STACKPATH_BASE_URL", defaultBaseURL),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"stackpath_compute_workload":       resourceComputeWorkload(),
			"stackpath_compute_network_policy": resourceComputeNetworkPolicy(),
			"stackpath_object_storage_bucket":  resourceObjectStorageBucket(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(data *schema.ResourceData) (interface{}, error) {
	config := &Config{
		StackID: data.Get("stack_id").(string),
		BaseURL: data.Get("base_url").(string),
	}

	if v, ok := data.GetOk("access_token"); ok {
		config.AccessToken = v.(string)
	}
	if v, ok := data.GetOk("client_id"); ok {
		config.ClientID = v.(string)
	}
	if v, ok := data.GetOk("client_secret"); ok {
		config.ClientSecret = v.(string)
	}

	log.Printf("[INFO] configuring stackpath provider")
	if err := config.LoadAndValidate(); err != nil {
		return nil, fmt.Errorf("unable to validate configuration: %v", err)
	}

	return config, nil
}
