package stackpath

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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
			"stackpath_compute_vpc_network":    resourceComputeVPCNetwork(),
			"stackpath_compute_vpc_route":      resourceComputeVPCRoute(),
			"stackpath_compute_network_policy": resourceComputeNetworkPolicy(),
			"stackpath_object_storage_bucket":  resourceObjectStorageBucket(),
		},
	}

	provider.ConfigureContextFunc = func(ctx context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
		return configureProvider(ctx, data, provider.TerraformVersion)
	}

	return provider
}

func configureProvider(ctx context.Context, data *schema.ResourceData, terraformVersion string) (interface{}, diag.Diagnostics) {
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
	if err := config.LoadAndValidate(ctx, terraformVersion); err != nil {
		return nil, diag.FromErr(fmt.Errorf("unable to validate configuration: %w", err))
	}

	return config, diag.Diagnostics{}
}
