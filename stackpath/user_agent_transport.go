package stackpath

import (
	"fmt"
	"net/http"

	terraform_version "github.com/hashicorp/terraform/version"
	"github.com/terraform-providers/terraform-provider-stackpath/version"
)

const userAgentFormat = "HashiCorp Terraform/%s (+https://www.terraform.io) terraform-provider-stackpath/%s (+https://www.terraform.io/docs/providers/stackpath)"

// UserAgentTransport is an http RoundTripper that sets a descriptive User-Agent
//header for all StackPath API requests.
type UserAgentTransport struct {
	http.RoundTripper
	parent http.RoundTripper
}

// NewUserAgentTransport builds a new UserAgentTransport around the underlying
// RoundTripper.
func NewUserAgentTransport(parent http.RoundTripper) *UserAgentTransport {
	return &UserAgentTransport{parent: parent}
}

// RoundTrip implements the http.RoundTripper interface, setting a User-Agent
// header on the HTTP request.
func (t *UserAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", fmt.Sprintf(userAgentFormat, terraform_version.Version, version.ProviderVersion))
	return t.parent.RoundTrip(req)
}
