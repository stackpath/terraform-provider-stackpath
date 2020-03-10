package stackpath

import (
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var clientID, clientSecret, stackID string
var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	clientID = os.Getenv("STACKPATH_CLIENT_ID")
	clientSecret = os.Getenv("STACKPATH_CLIENT_SECRET")
	stackID = os.Getenv("STACKPATH_STACK_ID")
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"stackpath": testAccProvider,
	}
}

func testAccPreCheck(t *testing.T) {
	missingVars := make([]string, 0)

	if clientID == "" {
		missingVars = append(missingVars, "STACKPATH_CLIENT_ID")
	}

	if clientSecret == "" {
		missingVars = append(missingVars, "STACKPATH_CLIENT_SECRET")
	}

	if stackID == "" {
		missingVars = append(missingVars, "STACKPATH_STACK_ID")
	}

	if len(missingVars) > 0 {
		t.Fatalf("%s must be set for acceptance tests", strings.Join(missingVars, ","))
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}
