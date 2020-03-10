package stackpath

import (
	"context"
	"fmt"
	"net/http"

	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/api/ipam/ipam_client/network_policies"
	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/api/ipam/ipam_models"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceComputeNetworkPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeNetworkPolicyCreate,
		Read:   resourceComputeNetworkPolicyRead,
		Update: resourceComputeNetworkPolicyUpdate,
		Delete: resourceComputeNetworkPolicyDelete,
		Importer: &schema.ResourceImporter{
			State: resourceComputeNetworkPolicyImportState,
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
			"description": {
				Type:     schema.TypeString,
				Optional: true,
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
			"instance_selector": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceComputeMatchExpressionSchema(),
			},
			"network_selector": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceComputeMatchExpressionSchema(),
			},
			"policy_types": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"egress": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:     schema.TypeString,
							Required: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"protocol": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ah": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem:     &schema.Resource{},
									},
									"esp": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem:     &schema.Resource{},
									},
									"gre": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem:     &schema.Resource{},
									},
									"icmp": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem:     &schema.Resource{},
									},
									"tcp": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"destination_ports": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"source_ports": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"tcp_udp": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"destination_ports": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"source_ports": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"udp": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"destination_ports": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"source_ports": {
													Type:     schema.TypeList,
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
						},
						"to": resourceComputeNetworkPolicyHostRuleSchema(),
					},
				},
			},
			"ingress": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:     schema.TypeString,
							Required: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"protocol": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ah": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem:     &schema.Resource{},
									},
									"esp": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem:     &schema.Resource{},
									},
									"gre": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem:     &schema.Resource{},
									},
									"icmp": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem:     &schema.Resource{},
									},
									"tcp": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"destination_ports": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"source_ports": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"tcp_udp": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"destination_ports": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"source_ports": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"udp": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"destination_ports": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"source_ports": {
													Type:     schema.TypeList,
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
						},
						"from": resourceComputeNetworkPolicyHostRuleSchema(),
					},
				},
			},
		},
	}
}

func resourceComputeNetworkPolicyHostRuleSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"instance_selector": {
					Type:     schema.TypeList,
					Optional: true,
					Elem:     resourceComputeMatchExpressionSchema(),
				},
				"network_selector": {
					Type:     schema.TypeList,
					Optional: true,
					Elem:     resourceComputeMatchExpressionSchema(),
				},
				"ip_block": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"cidr": {
								Type:         schema.TypeString,
								Required:     true,
								ValidateFunc: validateSubnet,
							},
							"except": {
								Type:     schema.TypeList,
								Optional: true,
								Elem: &schema.Schema{
									Type:         schema.TypeString,
									ValidateFunc: validateSubnet,
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceComputeNetworkPolicyCreate(data *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	resp, err := config.edgeComputeNetworking.NetworkPolicies.CreateNetworkPolicy(&network_policies.CreateNetworkPolicyParams{
		Context: context.Background(),
		StackID: config.StackID,
		Body: &ipam_models.V1CreateNetworkPolicyRequest{
			NetworkPolicy: convertComputeNetworkPolicy(data),
		},
	}, nil)
	if err != nil {
		return fmt.Errorf("failed to create network policy: %v", NewStackPathError(err))
	}

	data.SetId(resp.Payload.NetworkPolicy.ID)
	return resourceComputeNetworkPolicyRead(data, meta)
}

func resourceComputeNetworkPolicyRead(data *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	resp, err := config.edgeComputeNetworking.NetworkPolicies.GetNetworkPolicy(&network_policies.GetNetworkPolicyParams{
		StackID:         config.StackID,
		NetworkPolicyID: data.Id(),
		Context:         context.Background(),
	}, nil)
	if c, ok := err.(interface{ Code() int }); ok && c.Code() == http.StatusNotFound {
		// Clear out the ID in terraform if the
		// resource no longer exists in the API
		data.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to read network policy: %v", NewStackPathError(err))
	}

	if err := data.Set("name", resp.Payload.NetworkPolicy.Name); err != nil {
		return fmt.Errorf("error setting name: %v", err)
	}

	if err := data.Set("slug", resp.Payload.NetworkPolicy.Slug); err != nil {
		return fmt.Errorf("error setting slug: %v", err)
	}

	if err := data.Set("description", resp.Payload.NetworkPolicy.Description); err != nil {
		return fmt.Errorf("error setting description: %v", err)
	}

	if err := data.Set("labels", flattenStringMap(convertIPAMToWorkloadStringMapEntry(resp.Payload.NetworkPolicy.Metadata.Labels))); err != nil {
		return fmt.Errorf("error setting labels: %v", err)
	}

	if err := data.Set("annotations", flattenStringMap(convertIPAMToWorkloadStringMapEntry(resp.Payload.NetworkPolicy.Metadata.Annotations))); err != nil {
		return fmt.Errorf("error setting annotations: %v", err)
	}

	if err := data.Set("instance_selector", flattenComputeMatchExpressionsOrdered("instance_selector", data, convertIPAMToWorkloadMatchExpression(resp.Payload.NetworkPolicy.Spec.InstanceSelectors))); err != nil {
		return fmt.Errorf("error setting instance_selector: %v", err)
	}

	if err := data.Set("network_selector", flattenComputeMatchExpressionsOrdered("network_selector", data, convertIPAMToWorkloadMatchExpression(resp.Payload.NetworkPolicy.Spec.NetworkSelectors))); err != nil {
		return fmt.Errorf("error setting network_selector: %v", err)
	}

	if err := data.Set("policy_types", flattenComputeNetworkPolicyTypes(resp.Payload.NetworkPolicy.Spec.PolicyTypes)); err != nil {
		return fmt.Errorf("error setting policy_types: %v", err)
	}

	if err := data.Set("priority", resp.Payload.NetworkPolicy.Spec.Priority); err != nil {
		return fmt.Errorf("error setting priority: %v", err)
	}

	if err := data.Set("ingress", flattenComputeNetworkPolicyIngress(resp.Payload.NetworkPolicy.Spec.Ingress)); err != nil {
		return fmt.Errorf("error setting ingress: %v", err)
	}
	if err := data.Set("egress", flattenComputeNetworkPolicyEgress(resp.Payload.NetworkPolicy.Spec.Egress)); err != nil {
		return fmt.Errorf("error setting egress: %v", err)
	}

	return nil
}

func resourceComputeNetworkPolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkPolicy := convertComputeNetworkPolicy(data)
	networkPolicy.ID = data.Id()

	_, err := config.edgeComputeNetworking.NetworkPolicies.UpdateNetworkPolicy(&network_policies.UpdateNetworkPolicyParams{
		Context: context.Background(),
		StackID: config.StackID,
		Body: &ipam_models.V1UpdateNetworkPolicyRequest{
			NetworkPolicy: networkPolicy,
		},
	}, nil)
	if c, ok := err.(interface{ Code() int }); ok && c.Code() == http.StatusNotFound {
		// Clear out the ID in terraform if the
		// resource no longer exists in the API
		data.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to update network policy: %v", NewStackPathError(err))
	}

	return resourceComputeNetworkPolicyRead(data, meta)
}

func resourceComputeNetworkPolicyDelete(data *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	_, err := config.edgeComputeNetworking.NetworkPolicies.DeleteNetworkPolicy(&network_policies.DeleteNetworkPolicyParams{
		Context:         context.Background(),
		StackID:         config.StackID,
		NetworkPolicyID: data.Id(),
	}, nil)
	if err != nil {
		return fmt.Errorf("failed to delete network policy: %v", NewStackPathError(err))
	}

	data.SetId("")
	return nil
}

func resourceComputeNetworkPolicyImportState(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// We expect that to import a resource, the user will pass in the
	// full UUID of the network policy they're attempting to import.
	return []*schema.ResourceData{d}, nil
}
