package diaspora

import (
	"net/url"

	json "github.com/pquerna/ffjson/ffjson"
)

type (
	// Client represents a pod client.
	Client struct {
		Pod *url.URL
	}

	// ProviderConfiguration contains information about current pod andpoints, scopes and other authentication stuff.
	ProviderConfiguration struct {
		AuthorizationEndpoint                  string   `json:"authorization_endpoint"`
		ClaimsParameterSupported               bool     `json:"claims_parameter_supported"`
		ClaimsSupported                        []string `json:"claims_supported"`
		IDTokenSigningAlgValuesSupported       []string `json:"id_token_signing_alg_values_supported"`
		Issuer                                 string   `json:"issuer"`
		JwksURI                                string   `json:"jwks_uri"`
		RegistrationEndpoint                   string   `json:"registration_endpoint"`
		RequestObjectSigningAlgValuesSupported []string `json:"request_object_signing_alg_values_supported"`
		RequestParameterSupported              bool     `json:"request_parameter_supported"`
		RequestURIParameterSupported           bool     `json:"request_uri_parameter_supported"`
		ResponseTypesSupported                 []string `json:"response_types_supported"`
		ScopesSupported                        []string `json:"scopes_supported"`
		SubjectTypesSupported                  []string `json:"subject_types_supported"`
		TokenEndpoint                          string   `json:"token_endpoint"`
		TokenEndpointAuthMethodsSupported      []string `json:"token_endpoint_auth_methods_supported"`
		UserinfoEndpoint                       string   `json:"userinfo_endpoint"`
		UserinfoSigningAlgValuesSupported      []string `json:"userinfo_signing_alg_values_supported"`
	}
)

// NewClient create a new client for making all other requests.
func NewClient(podEndpoint string) (*Client, error) {
	pod, err := url.Parse(podEndpoint)
	if err != nil {
		return nil, err
	}

	return &Client{Pod: pod}, nil
}

// DiscoveryEndpoint return ProviderConfiguration of client pod.
//
// OpenID endpoints can be discovered using OpenID Connect Discovery 1.0
// (http://openid.net/specs/openid-connect-discovery-1_0.html). In addition to all required endpoints,
// the response will also include available scopes and other information.
func (c *Client) DiscoveryEndpoint() (*ProviderConfiguration, error) {
	body, err := c.get("/.well-known/openid-configuration", nil)
	if err != nil {
		return nil, err
	}

	var resp ProviderConfiguration
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
