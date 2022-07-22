package cdr_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fiskil/cdr"
	"github.com/fiskil/cdr/data"
	"github.com/matryer/is"
)

func TestGetWellKnownConfig(t *testing.T) {

	// Arrange
	is := is.New(t)
	ctx := context.Background()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"issuer":"https://secure.api.cdr.gov.au/idp","jwks_uri":"https://api.cdr.gov.au/idp/.well-known/openid-configuration/jwks","token_endpoint":"https://secure.api.cdr.gov.au/idp/connect/token","claims_supported":["sub"],"id_token_signing_alg_values_supported":["PS256"],"subject_types_supported":["public"],"code_challenge_methods_supported":["plain","S256"],"scopes_supported":["cdr-register:bank:read"],"response_types_supported":["token"],"grant_types_supported":["client_credentials"],"token_endpoint_auth_methods_supported":["private_key_jwt"],"tls_client_certificate_bound_access_tokens":true,"token_endpoint_auth_signing_alg_values_supported":["PS256","ES256"]}`)
	}))
	cli, err := cdr.NewFromEnv()
	is.NoErr(err)

	// Act
	res, err := cdr.GetWellKnownConfig(ctx, cli, ts.URL)

	// Assert
	is.NoErr(err)
	is.Equal(res.Issuer, "https://secure.api.cdr.gov.au/idp")

}

func ExampleGetWellKnownConfig() {

	// This gets the well known config for the CDR, but this endpoint can also be used for other data holders.
	ctx := context.Background()
	path := data.CDRBaseURL + "/idp" // Notice the use of the non secure URL.
	cli := http.DefaultClient

	res, _ := cdr.GetWellKnownConfig(ctx, cli, path)

	fmt.Println(res.Issuer)
	// Output: https://secure.api.cdr.gov.au/idp
}
