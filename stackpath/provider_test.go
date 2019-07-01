package stackpath

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var (
	testAccProvider  = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"stackpath": testAccProvider,
	}
)

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}
