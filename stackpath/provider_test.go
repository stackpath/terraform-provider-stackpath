package stackpath

import (
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var clientID, clientSecret, stackID string
var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider
var testAccProviderFactories = map[string]func() (*schema.Provider, error){
	"stackpath": func() (*schema.Provider, error) {
		return Provider(), nil
	},
}

func init() {
	clientID = os.Getenv("STACKPATH_CLIENT_ID")
	clientSecret = os.Getenv("STACKPATH_CLIENT_SECRET")
	stackID = os.Getenv("STACKPATH_STACK_ID")
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"stackpath": testAccProvider,
	}
	testAccProviderFactories = map[string]func() (*schema.Provider, error){
		"stackpath": func() (*schema.Provider, error) {
			return testAccProvider, nil
		},
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
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}
