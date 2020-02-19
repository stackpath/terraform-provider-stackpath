package stackpath

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/api/ipam/ipam_client"
	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/api/workload/workload_client"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/hashicorp/terraform/helper/logging"
	"github.com/hashicorp/terraform/helper/pathorcontents"
	"github.com/hashicorp/terraform/httpclient"

	object_storage "github.com/terraform-providers/terraform-provider-stackpath/stackpath/api/object_storage/client"
	"github.com/terraform-providers/terraform-provider-stackpath/version"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	// default to the official StackPath API
	defaultBaseURL = "https://gateway.stackpath.com"
)

// Config is used to configure the StackPath provider.
type Config struct {
	// The Stack in which all resources should be created
	StackID string

	// The Client ID that should be used to retrieve an access token. This
	// option must not be used with the access token option.
	ClientID string

	// The Client Secret that should be used to retrieve an access
	// token. This options must not be used with the access token option.
	ClientSecret string

	// The AccessToken that should be used in the request. This must not
	// be used in combination with the client ID and client secret options.
	AccessToken string

	// The base URL to use for API requests. This value must not end
	// in a trailing slash. This value will default to the official
	// StackPath API.
	BaseURL string

	client      *http.Client
	userAgent   string
	tokenSource oauth2.TokenSource

	edgeCompute           *workload_client.EdgeCompute
	edgeComputeNetworking *ipam_client.EdgeComputeNetworking
	objectStorage         *object_storage.ObjectStorage
}

// LoadAndValidate will load the configuration and validate the configuration
// options. An error will be returned when the configuration is invalid.
func (c *Config) LoadAndValidate() error {
	if c.ClientID == "" && c.ClientSecret == "" && c.AccessToken == "" {
		// Require the user to provide at least one form of authentication
		return fmt.Errorf("must provide either an access_token or a client_id and client_secret")
	} else if (c.ClientID != "" || c.ClientSecret != "") && c.AccessToken != "" {
		// Do not allow the access token to be provided when
		// either the client ID or client secret are set.
		return fmt.Errorf("must not use the client_id or client_secret option with the access_token option")
	} else if c.AccessToken == "" && (c.ClientID == "" || c.ClientSecret == "") {
		// Require the user to provide both the client ID and
		// client secret when using OAuth for authentication
		return fmt.Errorf("must provide both a client_id and client_secret to authenticate")
	}

	// Validate the base URL
	if c.BaseURL != "" && strings.HasSuffix(c.BaseURL, "/") {
		return fmt.Errorf("base_url must not end in a trailing slash")
	} else if baseURL, err := url.Parse(c.BaseURL); err != nil {
		return fmt.Errorf("failed to parse base_url as valid URL: %v", err)
	} else if baseURL.Scheme != "http" && baseURL.Scheme != "https" {
		return fmt.Errorf("invalid base_url option, must be a valid URL that begins with http:// or https://")
	} else if baseURL.Host == "" {
		return fmt.Errorf("invalid base_url option, must be a valid URL")
	} else {
		c.BaseURL = baseURL.Host
	}

	if c.StackID == "" {
		// Require the user to provide a stack
		return fmt.Errorf("must provide a stack to create resources in")
	}

	tokenSource, err := c.getTokenSource()
	if err != nil {
		return err
	}

	// Create a new http Client that will pull authentication tokens
	// from the configured token source
	c.client = oauth2.NewClient(context.Background(), tokenSource)
	c.client.Transport = logging.NewTransport("StackPath", c.client.Transport)
	// Each individual request should return within 60s - timeouts will be retried.
	// This is a timeout for, e.g. a single GET request of an operation - not a
	// timeout for the maximum amount of time a logical request can take.
	c.client.Timeout = 60 * time.Second

	// Set the user agent that's used by the Terraform Provider
	c.userAgent = fmt.Sprintf(
		"%s terraform-provider-stackpath/%s (+https://www.terraform.io)",
		httpclient.UserAgentString(),
		version.ProviderVersion,
	)

	// Create a new openAPI runtime
	runtime := httptransport.NewWithClient(c.BaseURL, "/", []string{"https"}, c.client)

	c.edgeCompute = workload_client.New(runtime, nil)
	c.edgeComputeNetworking = ipam_client.New(runtime, nil)
	c.objectStorage = object_storage.New(runtime, nil)

	return nil
}

func (c *Config) getTokenSource() (oauth2.TokenSource, error) {
	if c.AccessToken != "" {
		contents, _, err := pathorcontents.Read(c.AccessToken)
		if err != nil {
			return nil, fmt.Errorf("error loading access_token: %v", err)
		}

		// Create a static token source based on the configured access token
		return oauth2.StaticTokenSource(&oauth2.Token{AccessToken: contents}), nil
	}

	clientID, _, err := pathorcontents.Read(c.ClientID)
	if err != nil {
		return nil, fmt.Errorf("error loading client_id: %v", err)
	}

	clientSecret, _, err := pathorcontents.Read(c.ClientSecret)
	if err != nil {
		return nil, fmt.Errorf("error loading client_secret: %v", err)
	}

	oauthConfig := clientcredentials.Config{
		AuthStyle:    oauth2.AuthStyleInParams,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     "https://" + c.BaseURL + "/identity/v1/oauth2/token",
	}

	return oauthConfig.TokenSource(context.Background()), nil
}
