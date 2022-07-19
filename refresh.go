package cdr

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// RefreshTokenResponse is a response you get from exchanging a refresh token.
type RefreshTokenResponse struct {
	AccessToken    string `json:"access_token"`
	IDToken        string `json:"id_token"`
	ExpiresIn      int    `json:"expires_in"`
	Scope          string `json:"scope"`
	RefreshToken   string `json:"refresh_token"`
	CDRArrangement string `json:"cdr_arrangement_id"`
}

// RefreshToken gets a new access token from a refresh token.
func RefreshToken(ctx context.Context, cli *http.Client, tokenEndpoint string, refreshToken string, clientID string, clientAssertion string) (RefreshTokenResponse, error) {

	requestParams := url.Values{}
	requestParams.Set("client_assertion_type", "urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
	requestParams.Set("client_assertion", clientAssertion)
	requestParams.Set("client_id", clientID)
	requestParams.Set("grant_type", "refresh_token")
	requestParams.Set("refresh_token", refreshToken)
	body := requestParams.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, tokenEndpoint, strings.NewReader(body))
	if err != nil {
		return RefreshTokenResponse{}, fmt.Errorf("cdr : building refresh token request : %w", err)

	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := cli.Do(req)
	if err != nil {
		return RefreshTokenResponse{}, fmt.Errorf("cdr : sending refresh token request : %w", err)
	}
	if res.StatusCode > 299 {
		return RefreshTokenResponse{}, &ErrNon2xxResponse{
			StatusCode: res.StatusCode,
			Response:   res.Body,
			URL:        res.Request.URL,
		}
	}

	var result RefreshTokenResponse
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return RefreshTokenResponse{}, fmt.Errorf("cdr : parsing refresh token response : %w", err)
	}

	return result, nil
}
