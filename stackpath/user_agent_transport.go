package stackpath

import (
	"fmt"
	"net/http"

	"github.com/stackpath/terraform-provider-stackpath/version"
)

const userAgentFormat = "HashiCorp Terraform/%s (+https://www.terraform.io) terraform-provider-stackpath/%s (+https://registry.terraform.io/providers/stackpath/stackpath)"

// UserAgentTransport is an http RoundTripper that sets a descriptive User-Agent
//header for all StackPath API requests.
type UserAgentTransport struct {
	terraformVersion string
	http.RoundTripper
	parent http.RoundTripper
}

// NewUserAgentTransport builds a new UserAgentTransport around the underlying
// RoundTripper.
func NewUserAgentTransport(parent http.RoundTripper, terraformVersion string) *UserAgentTransport {
	return &UserAgentTransport{parent: parent, terraformVersion: terraformVersion}
}

// RoundTrip implements the http.RoundTripper interface, setting a User-Agent
// header on the HTTP request.
func (t *UserAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", fmt.Sprintf(userAgentFormat, t.terraformVersion, version.ProviderVersion))
	return t.parent.RoundTrip(req)
}
