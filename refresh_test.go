package cdr_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/fiskil/cdr"
	"github.com/matryer/is"
)

func TestRefreshToken(t *testing.T) {

	// Arrange
	is := is.New(t)
	ctx := context.Background()
	clientID := "my-client-id"
	refreshToken := "my-refresh-token"
	assertions := "ey... some encrypted assertions"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		is.NoErr(err)
		params, err := url.ParseQuery(string(body))
		is.NoErr(err)
		is.Equal(params.Get("client_assertion_type"), "urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
		is.Equal(params.Get("client_assertion"), assertions)
		is.Equal(params.Get("client_id"), clientID)
		is.Equal(params.Get("grant_type"), "refresh_token")
		is.Equal(params.Get("refresh_token"), refreshToken)

		is.Equal(r.Header.Get("Content-Type"), "application/x-www-form-urlencoded")

		fmt.Fprintf(w, `{"refresh_token":"a new start","access_token":"access","id_token":"my id","scope":"all the scope","expires_in":123,"cdr_arrangement_id":"1"}`)
	}))

	// Act
	res, err := cdr.RefreshToken(ctx, http.DefaultClient, ts.URL, refreshToken, clientID, assertions)

	// Assert
	is.NoErr(err)
	is.Equal(res.RefreshToken, "a new start")
	is.Equal(res.IDToken, "my id")
	is.Equal(res.Scope, "all the scope")
	is.Equal(res.ExpiresIn, 123)
	is.Equal(res.AccessToken, "access")
	is.Equal(res.CDRArrangement, "1")

}
