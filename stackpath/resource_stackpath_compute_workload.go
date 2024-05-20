package stackpath

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/workload/workload_client/instances"
	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/workload/workload_client/workloads"
	"github.com/stackpath/terraform-provider-stackpath/v2/stackpath/api/workload/workload_models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// annotation keys that should be ignored when diffing the state of a workload
var ignoredComputeWorkloadAnnotations = map[string]bool{
	"anycast.platform.stackpath.net/subnets": true,
}

func resourceComputeWorkload() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceComputeWorkloadCreate,
		ReadContext:   resourceComputeWorkloadRead,
		UpdateContext: resourceComputeWorkloadUpdate,
		DeleteContext: resourceComputeWorkloadDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceComputeWorkloadImportState,
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
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"annotations": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				DiffSuppressFunc: func(key, _, _ string, d *schema.ResourceData) bool {
					o, n := d.GetChange("annotations")
					oldData, newData := o.(map[string]interface{}), n.(map[string]interface{})
					for k, newVal := range newData {
						// check if it is an ignored annotation
						if ignoredComputeWorkloadAnnotations[k] {
							continue
						}
						// compare the previous value and see if it changed
						if oldVal, exists := oldData[k]; !exists || oldVal != newVal {
							return false
						}
					}

					for k, oldVal := range oldData {
						// check if it is an ignored annotation
						if ignoredComputeWorkloadAnnotations[k] {
							continue
						}
						// compare the previous value and see if it changed
						if newVal, exists := newData[k]; !exists || oldVal != newVal {
							return false
						}
					}

					return true
				},
			},
			"network_interface": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network": {
							Type:     schema.TypeString,
							Required: true,
						},
						"enable_one_to_one_nat": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"ip_families": {
							Type:     schema.TypeList,
							MaxItems: 2,
							Optional: true,
							// when ip_families is not provided, api creates workload with [IPv4] interface
							// terraform plugin sdk does not support a way to configure Default for TypeList
							// hence to avoid update in-place errors treating resource as computed
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"subnet": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipv6_subnet": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"image_pull_credentials": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"docker_registry": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"server": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"username": {
										Type:     schema.TypeString,
										Required: true,
									},
									"password": {
										Type:      schema.TypeString,
										Required:  true,
										Sensitive: true,
									},
									"email": {
										Type:      schema.TypeString,
										Optional:  true,
										Sensitive: true,
									},
								},
							},
						},
					},
				},
			},
			"virtual_machine": {
				Type:          schema.TypeList,
				ConflictsWith: []string{"container"},
				MaxItems:      1,
				Optional:      true,
				Elem:          resourceComputeWorkloadVirtualMachine(),
			},
			"container": {
				Type:          schema.TypeList,
				Optional:      true,
				ConflictsWith: []string{"virtual_machine"},
				Elem:          resourceComputeWorkloadContainer(),
			},
			"init_container": {
				Type:          schema.TypeList,
				Optional:      true,
				ConflictsWith: []string{"virtual_machine"},
				Elem:          resourceComputeWorkloadContainer(),
			},
			"volume_claim": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"slug": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"storage_class": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "stackpath-edge/network-standard",
						},
						"resources": resourceComputeWorkloadResourcesSchema(),
					},
				},
			},
			"target": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"min_replicas": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"max_replicas": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"scale_settings": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"metrics": {
										Type:     schema.TypeList,
										Required: true,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"metric": {
													Type:     schema.TypeString,
													Required: true,
												},
												"average_utilization": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"average_value": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"deployment_scope": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "cityCode",
						},
						"selector": {
							Type:     schema.TypeList,
							Required: true,
							MinItems: 1,
							Elem:     resourceComputeMatchExpressionSchema(),
						},
					},
				},
			},
			"container_runtime_environment": {
				Type:          schema.TypeList,
				Optional:      true,
				MaxItems:      1,
				ConflictsWith: []string{"virtual_machine_runtime_environment"},
				Elem:          resourceComputeContainerRuntimeEnvironment(),
			},
			"virtual_machine_runtime_environment": {
				Type:          schema.TypeList,
				Optional:      true,
				MaxItems:      1,
				ConflictsWith: []string{"container_runtime_environment"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
							Elem:     resourceComputeRuntimeEnvironmentDns(),
						},
					},
				},
			},
			"instances": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem:     resourceComputeWorkloadInstance(),
			},
		},
	}
}

func resourceComputeWorkloadVolumeMountSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"slug": {
					Type:     schema.TypeString,
					Required: true,
				},
				"mount_path": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	}
}

func resourceComputeWorkloadProbeSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"http_get": {
					Type:     schema.TypeList,
					MaxItems: 1,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"path": {
								Type:     schema.TypeString,
								Optional: true,
								Default:  "/",
							},
							"port": {
								Type:     schema.TypeInt,
								Required: true,
							},
							"scheme": {
								Type:     schema.TypeString,
								Optional: true,
								Default:  "http",
							},
							"http_headers": {
								Type:     schema.TypeMap,
								Optional: true,
								Elem: &schema.Schema{
									Type: schema.TypeString,
								},
							},
						},
					},
				},
				"tcp_socket": {
					Type:     schema.TypeList,
					MaxItems: 1,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"port": {
								Type:     schema.TypeInt,
								Required: true,
							},
						},
					},
				},
				"initial_delay_seconds": {
					Type:     schema.TypeInt,
					Optional: true,
					Default:  0,
				},
				"timeout_seconds": {
					Type:     schema.TypeInt,
					Optional: true,
					Default:  10,
				},
				"period_seconds": {
					Type:     schema.TypeInt,
					Optional: true,
					Default:  60,
				},
				"success_threshold": {
					Type:     schema.TypeInt,
					Required: true,
				},
				"failure_threshold": {
					Type:     schema.TypeInt,
					Required: true,
				},
			},
		},
	}
}

func resourceComputeWorkloadResourcesSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"requests": {
					Type:     schema.TypeMap,
					Required: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

func resourceComputeWorkloadEnvVars() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"key": {
					Type:     schema.TypeString,
					Required: true,
				},
				"value": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"secret_value": {
					Type:      schema.TypeString,
					Optional:  true,
					Sensitive: true,
				},
				"value_from": {
					Type:     schema.TypeList,
					MaxItems: 1,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"instance_field_ref": {
								Type:     schema.TypeList,
								MaxItems: 1,
								Optional: true,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"field_path": {
											Type:     schema.TypeString,
											Optional: true,
										},
										"optional": {
											Type:     schema.TypeBool,
											Optional: true,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceComputeWorkloadPortSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Required: true,
				},
				"enable_implicit_network_policy": {
					Type:     schema.TypeBool,
					Optional: true,
					Default:  false,
				},
				"port": {
					Type:     schema.TypeInt,
					Required: true,
				},
				"protocol": {
					Type:     schema.TypeString,
					Optional: true,
					Default:  "tcp",
				},
			},
		},
	}
}

func resourceComputeWorkloadSecurityContextSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"allow_privilege_escalation": {
					Type:     schema.TypeBool,
					Optional: true,
					Default:  false,
				},
				"read_only_root_filesystem": {
					Type:     schema.TypeBool,
					Optional: true,
					Default:  false,
				},
				"run_as_group": {
					Type:     schema.TypeString,
					Optional: true,
					Default:  "",
				},
				"run_as_user": {
					Type:     schema.TypeString,
					Optional: true,
					Default:  "",
				},
				"run_as_non_root": {
					Type:     schema.TypeBool,
					Optional: true,
					Default:  false,
				},
				"capabilities": {
					Type:     schema.TypeList,
					MaxItems: 1,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"add": {
								Type:     schema.TypeSet,
								Optional: true,
								Elem: &schema.Schema{
									Type: schema.TypeString,
								},
							},
							"drop": {
								Type:     schema.TypeSet,
								Optional: true,
								Elem: &schema.Schema{
									Type: schema.TypeString,
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceComputeContainerRuntimeEnvironment() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"termination_grace_period_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"share_process_namespace": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"security_context": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"run_as_group": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"run_as_user": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"run_as_non_root": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"supplemental_groups": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"sysctl": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"dns": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem:     resourceComputeRuntimeEnvironmentDns(),
			},
		},
	}
}

func resourceComputeRuntimeEnvironmentDns() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host_aliases": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:     schema.TypeString,
							Required: true,
						},
						"hostnames": {
							Type:     schema.TypeSet,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"resolver_config": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{

					Schema: map[string]*schema.Schema{
						"nameservers": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"search": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"options": {
							Type:     schema.TypeMap,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceComputeWorkloadCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	// Create the workload
	resp, err := config.edgeCompute.Workloads.CreateWorkload(&workloads.CreateWorkloadParams{
		Context: ctx,
		StackID: config.StackID,
		Body: &workload_models.V1CreateWorkloadRequest{
			Workload: convertComputeWorkload(data),
		},
	}, nil)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create compute workload: %v", NewStackPathError(err)))
	}

	// Set the ID based on the workload created in the API
	data.SetId(resp.Payload.Workload.ID)

	return resourceComputeWorkloadRead(ctx, data, meta)
}

func resourceComputeWorkloadReplace(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	_, err := config.edgeCompute.Workloads.PutWorkload(&workloads.PutWorkloadParams{
		Context:    ctx,
		StackID:    config.StackID,
		WorkloadID: data.Id(),
		Body: &workload_models.V1PutWorkloadRequest{
			Workload: convertComputeWorkload(data),
		},
	}, nil)
	if c, ok := err.(interface{ Code() int }); ok && c.Code() == http.StatusNotFound {
		// Clear out the ID in terraform if the
		// resource no longer exists in the API
		data.SetId("")
		return diag.Diagnostics{}
	} else if err != nil {
		return diag.FromErr(fmt.Errorf("failed to update compute workload: %v", NewStackPathError(err)))
	}
	return resourceComputeWorkloadRead(ctx, data, meta)
}

func resourceComputeWorkloadUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)

	wk := convertComputeWorkload(data)
	if wk.Metadata.Version != "" {
		// Currently using PATCH in cases where we are making updates
		// can miss array values, so we prefer to use PUT. But that requires
		// the version to be present, which might not be there for
		// older resources
		return resourceComputeWorkloadReplace(ctx, data, meta)
	}

	_, err := config.edgeCompute.Workloads.UpdateWorkload(&workloads.UpdateWorkloadParams{
		Context:    ctx,
		StackID:    config.StackID,
		WorkloadID: data.Id(),
		Body: &workload_models.V1UpdateWorkloadRequest{
			Workload: convertComputeWorkload(data),
		},
	}, nil)
	if c, ok := err.(interface{ Code() int }); ok && c.Code() == http.StatusNotFound {
		// Clear out the ID in terraform if the
		// resource no longer exists in the API
		data.SetId("")
		return diag.Diagnostics{}
	} else if err != nil {
		return diag.FromErr(fmt.Errorf("failed to update compute workload: %v", NewStackPathError(err)))
	}
	return resourceComputeWorkloadRead(ctx, data, meta)
}

func resourceComputeWorkloadRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)

	resp, err := config.edgeCompute.Workloads.GetWorkload(&workloads.GetWorkloadParams{
		Context:    ctx,
		StackID:    config.StackID,
		WorkloadID: data.Id(),
	}, nil)
	if c, ok := err.(interface{ Code() int }); ok && c.Code() == http.StatusNotFound {
		// Clear out the ID in terraform if the
		// resource no longer exists in the API
		data.SetId("")
		return diag.Diagnostics{}
	} else if err != nil {
		return diag.FromErr(fmt.Errorf("failed to read compute workload: %v", NewStackPathError(err)))
	}

	if err := flattenComputeWorkload(data, resp.Payload.Workload); err != nil {
		return diag.FromErr(err)
	}

	return resourceComputeWorkloadReadInstances(ctx, data, meta)
}

func resourceComputeWorkloadReadInstances(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)

	pageSize := "50"
	// variable to keep track of our location through pagination
	var endCursor string
	var terraformInstances []interface{}
	for {
		params := &instances.GetWorkloadInstancesParams{
			StackID:          config.StackID,
			WorkloadID:       data.Id(),
			Context:          ctx,
			PageRequestFirst: &pageSize,
		}
		if endCursor != "" {
			params.PageRequestAfter = &endCursor
		}
		resp, err := config.edgeCompute.Instances.GetWorkloadInstances(params, nil)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to read compute workload instances: %v", NewStackPathError(err)))
		}
		for _, result := range resp.Payload.Results {
			terraformInstances = append(terraformInstances, flattenComputeWorkloadInstance(result))
		}
		// Continue paginating until we get all the results
		if !resp.Payload.PageInfo.HasNextPage {
			break
		}
		endCursor = resp.Payload.PageInfo.EndCursor
	}

	if err := data.Set("instances", terraformInstances); err != nil {
		return diag.FromErr(fmt.Errorf("error setting instances: %v", err))
	}

	return diag.Diagnostics{}
}

func resourceComputeWorkloadDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)

	_, err := config.edgeCompute.Workloads.DeleteWorkload(&workloads.DeleteWorkloadParams{
		Context:    ctx,
		StackID:    config.StackID,
		WorkloadID: data.Id(),
	}, nil)
	if c, ok := err.(interface{ Code() int }); ok && c.Code() == http.StatusNotFound {
		// Clear out the ID in terraform if the
		// resource no longer exists in the API
		data.SetId("")
		return diag.Diagnostics{}
	} else if err != nil {
		return diag.FromErr(fmt.Errorf("failed to delete compute workload: %v", NewStackPathError(err)))
	}
	return diag.Diagnostics{}
}

func resourceComputeWorkloadImportState(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// We expect that to import a resource, the user will pass in
	// the full UUID of the workload they're attempting to import.
	return []*schema.ResourceData{d}, nil
}
