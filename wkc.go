package cdr

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// WellKnownConfigResponse is well known config for a registar.
type WellKnownConfigResponse struct {
	Issuer                                  string   `json:"issuer"`
	JwksURI                                 string   `json:"jwks_uri"`
	TokenEndpoint                           string   `json:"token_endpoint"`
	SupportedClaims                         []string `json:"claims_supported"`
	IDTokenSigningAlgorithms                []string `json:"id_token_signing_alg_values_supported"`
	SupportedSubjectTypes                   []string `json:"subject_types_supported"`
	SupportedCodeChallengeMethods           []string `json:"code_challenge_methods_supported"`
	SupportedScopes                         []string `json:"scopes_supported"`
	SupportedResponseTypes                  []string `json:"response_types_supported"`
	SupportedGrantTypes                     []string `json:"grant_types_supported"`
	SupportedTokenEndpointAuthMethods       []string `json:"token_endpoint_auth_methods_supported"`
	TLSClientCertificateBoundAccessToken    bool     `json:"tls_client_certificate_bound_access_tokens"`
	SupportedRequestObjectSigningAlgorithms []string `json:"token_endpoint_auth_signing_alg_values_supported"`
}

// GetWellKnownConfig fetches the well known config based on the OIDC spec.
// You should not use the secure endpoint for the path.
func GetWellKnownConfig(ctx context.Context, cli *http.Client, baseURL string) (WellKnownConfigResponse, error) {

	url := baseURL + "/.well-known/openid-configuration"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return WellKnownConfigResponse{}, fmt.Errorf("cdr : building wkc request : %w", err)
	}

	res, err := cli.Do(req)
	if err != nil {
		return WellKnownConfigResponse{}, err
	}

	if res.StatusCode > 299 {
		return WellKnownConfigResponse{}, &ErrNon2xxResponse{
			StatusCode: res.StatusCode,
			Response:   res.Body,
			URL:        res.Request.URL,
		}
	}

	result := WellKnownConfigResponse{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return WellKnownConfigResponse{}, fmt.Errorf("cdr: unable to decode response: %v", err)
	}

	return result, nil
}
