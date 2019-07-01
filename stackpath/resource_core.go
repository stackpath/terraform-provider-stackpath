package stackpath

import "github.com/hashicorp/terraform/helper/schema"

func resourceComputeMatchExpressionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"operator": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"values": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}
