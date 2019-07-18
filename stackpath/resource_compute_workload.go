package stackpath

import (
	"context"
	"net/http"

	"github.com/stackpath/terraform-provider-stackpath/stackpath/internal/client"
	workload "github.com/stackpath/terraform-provider-stackpath/stackpath/internal/client"
	"github.com/stackpath/terraform-provider-stackpath/stackpath/internal/models"

	"github.com/hashicorp/terraform/helper/schema"
)

// annotation keys that should be ignored when diffing the state of a workload
var ignoredComputeWorkloadAnnotations = map[string]bool{
	"annotations.anycast.platform.stackpath.net/subnets": true,
}

func resourceComputeWorkload() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeWorkloadCreate,
		Read:   resourceComputeWorkloadRead,
		Update: resourceComputeWorkloadUpdate,
		Delete: resourceComputeWorkloadDelete,
		Importer: &schema.ResourceImporter{
			State: resourceComputeWorkloadImportState,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"labels": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"annotations": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(key, _, _ string, d *schema.ResourceData) bool {
					o, n := d.GetChange("annotations")
					oldData, newData := o.(map[string]interface{}), n.(map[string]interface{})
					for k, newVal := range newData {
						// check if its an ignored annotation
						if ignoredComputeWorkloadAnnotations[k] {
							continue
						}
						// compare the its previous value and see if its changed
						if oldVal, exists := oldData[k]; !exists || oldVal != newVal {
							return true
						}
					}

					for k, oldVal := range oldData {
						// check if its an ignored annotation
						if ignoredComputeWorkloadAnnotations[k] {
							continue
						}
						// compare the its previous value and see if its changed
						if newVal, exists := newData[k]; !exists || oldVal != newVal {
							return true
						}
					}

					return false
				},
			},
			"network_interface": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"image_pull_credentials": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"docker_registry": &schema.Schema{
							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"server": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"username": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"password": &schema.Schema{
										Type:      schema.TypeString,
										Required:  true,
										Sensitive: true,
									},
									"email": &schema.Schema{
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
			"virtual_machine": &schema.Schema{
				Type:          schema.TypeList,
				ConflictsWith: []string{"container"},
				MaxItems:      1,
				Optional:      true,
				Elem:          resourceComputeWorkloadVirtualMachine(),
			},
			"container": &schema.Schema{
				Type:          schema.TypeList,
				Optional:      true,
				ConflictsWith: []string{"virtual_machine"},
				Elem:          resourceComputeWorkloadContainer(),
			},
			"volume_claim": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"slug": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"resources": resourceComputeWorkloadResourcesSchema(),
					},
				},
			},
			"target": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"min_replicas": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"max_replicas": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"scale_settings": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"metrics": &schema.Schema{
										Type:     schema.TypeList,
										Required: true,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"metric": &schema.Schema{
													Type:     schema.TypeString,
													Required: true,
												},
												"average_utilization": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"average_value": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"deployment_scope": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Default:  "cityCode",
						},
						"selector": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							MinItems: 1,
							Elem:     resourceComputeMatchExpressionSchema(),
						},
					},
				},
			},
			"instances": &schema.Schema{
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
				"slug": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"mount_path": &schema.Schema{
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
				"http_get": &schema.Schema{
					Type:     schema.TypeList,
					MaxItems: 1,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"path": &schema.Schema{
								Type:     schema.TypeString,
								Optional: true,
								Default:  "/",
							},
							"port": &schema.Schema{
								Type:     schema.TypeInt,
								Required: true,
							},
							"scheme": &schema.Schema{
								Type:     schema.TypeString,
								Optional: true,
								Default:  "http",
							},
							"http_headers": &schema.Schema{
								Type:     schema.TypeMap,
								Optional: true,
								Elem: &schema.Schema{
									Type: schema.TypeString,
								},
							},
						},
					},
				},
				"tcp_socket": &schema.Schema{
					Type:     schema.TypeList,
					MaxItems: 1,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"port": &schema.Schema{
								Type:     schema.TypeInt,
								Required: true,
							},
						},
					},
				},
				"initial_delay_seconds": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
					Default:  0,
				},
				"timeout_seconds": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
					Default:  10,
				},
				"period_seconds": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
					Default:  60,
				},
				"success_threshold": &schema.Schema{
					Type:     schema.TypeInt,
					Required: true,
				},
				"failure_threshold": &schema.Schema{
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
				"requests": &schema.Schema{
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

func resourceComputeWorkloadPortSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"port": &schema.Schema{
					Type:     schema.TypeInt,
					Required: true,
				},
				"protocol": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Default:  "tcp",
				},
			},
		},
	}
}

func resourceComputeWorkloadCreate(data *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	// Create the workload
	resp, err := config.compute.CreateWorkload(&workload.CreateWorkloadParams{
		Context: context.Background(),
		StackID: config.Stack,
		Body: &models.V1CreateWorkloadRequest{
			Workload: convertComputeWorkload(data),
		},
	}, nil)
	if err != nil {
		return err
	}

	// Set the ID based on the workload created in the API
	data.SetId(resp.Payload.Workload.ID)

	return resourceComputeWorkloadRead(data, meta)
}

func resourceComputeWorkloadUpdate(data *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	_, err := config.compute.UpdateWorkload(&workload.UpdateWorkloadParams{
		Context:    context.Background(),
		StackID:    config.Stack,
		WorkloadID: data.Id(),
		Body: &models.V1UpdateWorkloadRequest{
			Workload: convertComputeWorkload(data),
		},
	}, nil)
	if c, ok := err.(interface{ Code() int }); ok && c.Code() == http.StatusNotFound {
		// Clear out the ID in terraform if the
		// resource no longer exists in the API
		data.SetId("")
		return nil
	} else if err != nil {
		return err
	}
	return resourceComputeWorkloadRead(data, meta)
}

func resourceComputeWorkloadRead(data *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	resp, err := config.compute.GetWorkload(&workload.GetWorkloadParams{
		Context:    context.Background(),
		StackID:    config.Stack,
		WorkloadID: data.Id(),
	}, nil)
	if c, ok := err.(interface{ Code() int }); ok && c.Code() == http.StatusNotFound {
		// Clear out the ID in terraform if the
		// resource no longer exists in the API
		data.SetId("")
		return nil
	} else if err != nil {
		return err
	}

	flattenComputeWorkload(data, resp.Payload.Workload)
	return resourceComputeWorkloadReadInstances(data, meta)
}

func resourceComputeWorkloadReadInstances(data *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	pageSize := "50"
	// variable to keep track of our location through pagination
	var endCursor string
	var instances []interface{}
	for {
		params := &client.GetWorkloadInstancesParams{
			StackID:          config.Stack,
			WorkloadID:       data.Id(),
			Context:          context.Background(),
			PageRequestFirst: &pageSize,
		}
		if endCursor != "" {
			params.PageRequestAfter = &endCursor
		}
		resp, err := config.compute.GetWorkloadInstances(params, nil)
		if err != nil {
			return err
		}
		for _, result := range resp.Payload.Results {
			instances = append(instances, flattenComputeWorkloadInstance(result))
		}
		// Continue paginating until we get all the results
		if !resp.Payload.PageInfo.HasNextPage {
			break
		}
		endCursor = resp.Payload.PageInfo.EndCursor
	}

	data.Set("instances", instances)
	return nil
}

func resourceComputeWorkloadDelete(data *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	_, err := config.compute.DeleteWorkload(&workload.DeleteWorkloadParams{
		Context:    context.Background(),
		StackID:    config.Stack,
		WorkloadID: data.Id(),
	}, nil)
	if c, ok := err.(interface{ Code() int }); ok && c.Code() == http.StatusNotFound {
		// Clear out the ID in terraform if the
		// resource no longer exists in the API
		data.SetId("")
		return nil
	} else if err != nil {
		return err
	}
	return nil
}

func resourceComputeWorkloadImportState(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// We expect that to import a resource, the user will pass in
	// the full UUID of the workload they're attempting to import.
	return []*schema.ResourceData{d}, nil
}
