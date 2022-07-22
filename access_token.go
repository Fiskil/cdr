package cdr

import (
	"context"
	"net/http"
)

// GetAccessToken uses the client credentials flow to get an access token.
func GetAccessToken(ctx context.Context, cli *http.Client, baseURL string, clientID string, clientAssertion string, scope string) (string, error) {

	// GrantType:           "client_credentials",
	// 	ClientID:            h.SoftwareProductID,
	// 	ClientAssertionType: "urn:ietf:params:oauth:client-assertion-type:jwt-bearer",
	// 	ClientAssertion:     cat,
	// 	Scope:               "cdr-register:bank:read",

	return "", nil
}
