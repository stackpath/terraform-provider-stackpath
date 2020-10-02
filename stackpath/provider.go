package stackpath

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns the configured provider for managing StackPath resources.
func Provider() *schema.Provider {
	provider := &schema.Provider{
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
	}

	provider.ConfigureFunc = func(data *schema.ResourceData) (interface{}, error) {
		// Taken from https://github.com/terraform-providers/terraform-provider-kubernetes/pull/620/files#diff-da3a5957d1adf1d97d4dec9f43b36ec1R171
		// as an example for how to get the Terraform version into ConfigureFunc.
		terraformVersion := provider.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return configureProvider(data, terraformVersion)
	}

	return provider
}

func configureProvider(data *schema.ResourceData, terraformVersion string) (interface{}, error) {
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
	if err := config.LoadAndValidate(terraformVersion); err != nil {
		return nil, fmt.Errorf("unable to validate configuration: %v", err)
	}

	return config, nil
}
