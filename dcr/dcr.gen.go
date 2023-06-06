// Package dcr provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package dcr

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// Defines values for RegistrationErrorError.
const (
	InvalidClientMetadata       RegistrationErrorError = "invalid_client_metadata"
	InvalidRedirectUri          RegistrationErrorError = "invalid_redirect_uri"
	InvalidSoftwareStatement    RegistrationErrorError = "invalid_software_statement"
	UnapprovedSoftwareStatement RegistrationErrorError = "unapproved_software_statement"
)

// Defines values for RegistrationPropertiesApplicationType.
const (
	Web RegistrationPropertiesApplicationType = "web"
)

// Defines values for RegistrationPropertiesAuthorizationEncryptedResponseAlg.
const (
	RSAOAEP    RegistrationPropertiesAuthorizationEncryptedResponseAlg = "RSA-OAEP"
	RSAOAEP256 RegistrationPropertiesAuthorizationEncryptedResponseAlg = "RSA-OAEP-256"
)

// Defines values for RegistrationPropertiesAuthorizationEncryptedResponseEnc.
const (
	A128CBCHS256 RegistrationPropertiesAuthorizationEncryptedResponseEnc = "A128CBC-HS256"
	A256GCM      RegistrationPropertiesAuthorizationEncryptedResponseEnc = "A256GCM"
)

// Defines values for RegistrationPropertiesAuthorizationSignedResponseAlg.
const (
	RegistrationPropertiesAuthorizationSignedResponseAlgES256 RegistrationPropertiesAuthorizationSignedResponseAlg = "ES256"
	RegistrationPropertiesAuthorizationSignedResponseAlgPS256 RegistrationPropertiesAuthorizationSignedResponseAlg = "PS256"
)

// Defines values for RegistrationPropertiesGrantTypes.
const (
	AuthorizationCode RegistrationPropertiesGrantTypes = "authorization_code"
	ClientCredentials RegistrationPropertiesGrantTypes = "client_credentials"
	RefreshToken      RegistrationPropertiesGrantTypes = "refresh_token"
)

// Defines values for RegistrationPropertiesIdTokenSignedResponseAlg.
const (
	RegistrationPropertiesIdTokenSignedResponseAlgES256 RegistrationPropertiesIdTokenSignedResponseAlg = "ES256"
	RegistrationPropertiesIdTokenSignedResponseAlgPS256 RegistrationPropertiesIdTokenSignedResponseAlg = "PS256"
)

// Defines values for RegistrationPropertiesRequestObjectSigningAlg.
const (
	RegistrationPropertiesRequestObjectSigningAlgES256 RegistrationPropertiesRequestObjectSigningAlg = "ES256"
	RegistrationPropertiesRequestObjectSigningAlgPS256 RegistrationPropertiesRequestObjectSigningAlg = "PS256"
)

// Defines values for RegistrationPropertiesResponseTypes.
const (
	Code        RegistrationPropertiesResponseTypes = "code"
	CodeIdToken RegistrationPropertiesResponseTypes = "code id_token"
)

// Defines values for RegistrationPropertiesSoftwareRoles.
const (
	DataRecipientSoftwareProduct RegistrationPropertiesSoftwareRoles = "data-recipient-software-product"
)

// Defines values for RegistrationPropertiesTokenEndpointAuthMethod.
const (
	PrivateKeyJwt RegistrationPropertiesTokenEndpointAuthMethod = "private_key_jwt"
)

// Defines values for RegistrationPropertiesTokenEndpointAuthSigningAlg.
const (
	RegistrationPropertiesTokenEndpointAuthSigningAlgES256 RegistrationPropertiesTokenEndpointAuthSigningAlg = "ES256"
	RegistrationPropertiesTokenEndpointAuthSigningAlgPS256 RegistrationPropertiesTokenEndpointAuthSigningAlg = "PS256"
)

// ClientRegistrationRequest The registration request JWT to be used to register with a Data Holder.
type ClientRegistrationRequest = string

// RegistrationError defines model for RegistrationError.
type RegistrationError struct {
	// Error Predefined error code as described in [section 3.3 OIDC Dynamic Client Registration](https://openid.net/specs/openid-connect-registration-1_0.html)
	Error RegistrationErrorError `json:"error"`

	// ErrorDescription Additional text description of the error for debugging.
	ErrorDescription *string `json:"error_description,omitempty"`
}

// RegistrationErrorError Predefined error code as described in [section 3.3 OIDC Dynamic Client Registration](https://openid.net/specs/openid-connect-registration-1_0.html)
type RegistrationErrorError string

// RegistrationProperties defines model for RegistrationProperties.
type RegistrationProperties struct {
	// ApplicationType Kind of the application. The only supported application type will be `web`
	ApplicationType *RegistrationPropertiesApplicationType `json:"application_type,omitempty"`

	// AuthorizationEncryptedResponseAlg The JWE `alg` algorithm required for encrypting authorization responses. If unspecified, the default is that no encryption is performed.<br><br>Required if “authorization_encrypted_response_enc” is included.
	AuthorizationEncryptedResponseAlg *RegistrationPropertiesAuthorizationEncryptedResponseAlg `json:"authorization_encrypted_response_alg,omitempty"`

	// AuthorizationEncryptedResponseEnc The JWE `enc` algorithm required for encrypting authorization responses. If “authorization_encrypted_response_alg” is specified, the default for this value is “A128CBC-HS256”.
	AuthorizationEncryptedResponseEnc *RegistrationPropertiesAuthorizationEncryptedResponseEnc `json:"authorization_encrypted_response_enc,omitempty"`

	// AuthorizationSignedResponseAlg The JWS `alg` algorithm required for signing authorization responses. If this is specified, the response will be signed using JWS and the configured algorithm. The algorithm “none” is not allowed.<br><br>Required if response_type of “code” is registered by the client.
	AuthorizationSignedResponseAlg *RegistrationPropertiesAuthorizationSignedResponseAlg `json:"authorization_signed_response_alg,omitempty"`

	// ClientDescription Human-readable string name of the software product description to be presented to the end user during authorization
	ClientDescription string `json:"client_description"`

	// ClientId Data Holder issued client identifier string
	ClientId string `json:"client_id"`

	// ClientIdIssuedAt Time at which the client identifier was issued expressed as seconds since 1970-01-01T00:00:00Z as measured in UTC
	ClientIdIssuedAt *int `json:"client_id_issued_at,omitempty"`

	// ClientName Human-readable string name of the software product to be presented to the end-user during authorization
	ClientName string `json:"client_name"`

	// ClientUri URL string of a web page providing information about the client
	ClientUri string `json:"client_uri"`

	// GrantTypes Array of OAuth 2.0 grant type strings that the client can use at the token endpoint
	GrantTypes []RegistrationPropertiesGrantTypes `json:"grant_types"`

	// IdTokenEncryptedResponseAlg JWE `alg` algorithm with which an id_token is to be encrypted.<br/><br/>Required if OIDC Hybrid Flow (response type `code id_token`) is registered.
	IdTokenEncryptedResponseAlg *string `json:"id_token_encrypted_response_alg,omitempty"`

	// IdTokenEncryptedResponseEnc JWE `enc` algorithm with which an id_token is to be encrypted.<br/><br/>Required if OIDC Hybrid Flow (response type `code id_token`) is registered.
	IdTokenEncryptedResponseEnc *string `json:"id_token_encrypted_response_enc,omitempty"`

	// IdTokenSignedResponseAlg Algorithm with which an id_token is to be signed
	IdTokenSignedResponseAlg RegistrationPropertiesIdTokenSignedResponseAlg `json:"id_token_signed_response_alg"`

	// JwksUri URL string referencing the client JSON Web Key (JWK) Set **[[RFC7517]](#nref-RFC7517)** document, which contains the client public keys
	JwksUri string `json:"jwks_uri"`

	// LegalEntityId A unique identifier string assigned by the CDR Register that identifies the Accredited Data Recipient Legal Entity
	LegalEntityId *string `json:"legal_entity_id,omitempty"`

	// LegalEntityName Human-readable string name of the Accredited Data Recipient Legal Entity
	LegalEntityName *string `json:"legal_entity_name,omitempty"`

	// LogoUri URL string that references a logo for the client. If present, the server SHOULD display this image to the end-user during approval
	LogoUri string `json:"logo_uri"`

	// OrgId A unique identifier string assigned by the CDR Register that identifies the Accredited Data Recipient Brand
	OrgId string `json:"org_id"`

	// OrgName Human-readable string name of the Accredited Data Recipient to be presented to the end user during authorization
	OrgName string `json:"org_name"`

	// PolicyUri URL string that points to a human-readable policy document for the Software Product
	PolicyUri *string `json:"policy_uri,omitempty"`

	// RecipientBaseUri Base URI for the Consumer Data Standard Data Recipient endpoints. This should be the base to provide reference to all other Data Recipient Endpoints
	RecipientBaseUri *string `json:"recipient_base_uri,omitempty"`

	// RedirectUris Array of redirection URI strings for use in redirect-based flows. If used, redirect_uris MUST match or be a subset of the redirect_uris as defined in the SSA
	RedirectUris []string `json:"redirect_uris"`

	// RequestObjectSigningAlg Algorithm which the ADR expects to sign the request object if a request object will be part of the authorization request sent to the Data Holder
	RequestObjectSigningAlg RegistrationPropertiesRequestObjectSigningAlg `json:"request_object_signing_alg"`

	// ResponseTypes Array of the OAuth 2.0 response type strings that the client can use at the authorization endpoint.<br><br>Response type value `code` is required for Authorization Code Flow. Response type value `code id_token` is required for OIDC Hybrid Flow.
	ResponseTypes []RegistrationPropertiesResponseTypes `json:"response_types"`

	// RevocationUri URI string that references the location of the Software Product consent revocation endpoint
	RevocationUri *string `json:"revocation_uri,omitempty"`

	// Scope String containing a space-separated list of scope values that the client can use when requesting access tokens.
	Scope string `json:"scope"`

	// SectorIdentifierUri URL string referencing the client sector identifier URI, used as an optional input to the Pairwise Identifier
	SectorIdentifierUri *string `json:"sector_identifier_uri,omitempty"`

	// SoftwareId String representing a unique identifier assigned by the Register and used by registration endpoints to identify the software product to be dynamically registered. </br></br>The "software_id" will remain the same for the lifetime of the product, across multiple updates and versions
	SoftwareId string `json:"software_id"`

	// SoftwareRoles String containing a role of the software in the CDR Regime. Initially the only value used with be `data-recipient-software-product`
	SoftwareRoles *RegistrationPropertiesSoftwareRoles `json:"software_roles,omitempty"`

	// SoftwareStatement The Software Statement Assertion, as defined in CDR standards
	SoftwareStatement string `json:"software_statement"`

	// TokenEndpointAuthMethod The requested authentication method for the token endpoint
	TokenEndpointAuthMethod RegistrationPropertiesTokenEndpointAuthMethod `json:"token_endpoint_auth_method"`

	// TokenEndpointAuthSigningAlg The algorithm used for signing the JWT
	TokenEndpointAuthSigningAlg RegistrationPropertiesTokenEndpointAuthSigningAlg `json:"token_endpoint_auth_signing_alg"`

	// TosUri URL string that points to a human-readable terms of service document for the Software Product
	TosUri *string `json:"tos_uri,omitempty"`
}

// RegistrationPropertiesApplicationType Kind of the application. The only supported application type will be `web`
type RegistrationPropertiesApplicationType string

// RegistrationPropertiesAuthorizationEncryptedResponseAlg The JWE `alg` algorithm required for encrypting authorization responses. If unspecified, the default is that no encryption is performed.<br><br>Required if “authorization_encrypted_response_enc” is included.
type RegistrationPropertiesAuthorizationEncryptedResponseAlg string

// RegistrationPropertiesAuthorizationEncryptedResponseEnc The JWE `enc` algorithm required for encrypting authorization responses. If “authorization_encrypted_response_alg” is specified, the default for this value is “A128CBC-HS256”.
type RegistrationPropertiesAuthorizationEncryptedResponseEnc string

// RegistrationPropertiesAuthorizationSignedResponseAlg The JWS `alg` algorithm required for signing authorization responses. If this is specified, the response will be signed using JWS and the configured algorithm. The algorithm “none” is not allowed.<br><br>Required if response_type of “code” is registered by the client.
type RegistrationPropertiesAuthorizationSignedResponseAlg string

// RegistrationPropertiesGrantTypes defines model for RegistrationProperties.GrantTypes.
type RegistrationPropertiesGrantTypes string

// RegistrationPropertiesIdTokenSignedResponseAlg Algorithm with which an id_token is to be signed
type RegistrationPropertiesIdTokenSignedResponseAlg string

// RegistrationPropertiesRequestObjectSigningAlg Algorithm which the ADR expects to sign the request object if a request object will be part of the authorization request sent to the Data Holder
type RegistrationPropertiesRequestObjectSigningAlg string

// RegistrationPropertiesResponseTypes defines model for RegistrationProperties.ResponseTypes.
type RegistrationPropertiesResponseTypes string

// RegistrationPropertiesSoftwareRoles String containing a role of the software in the CDR Regime. Initially the only value used with be `data-recipient-software-product`
type RegistrationPropertiesSoftwareRoles string

// RegistrationPropertiesTokenEndpointAuthMethod The requested authentication method for the token endpoint
type RegistrationPropertiesTokenEndpointAuthMethod string

// RegistrationPropertiesTokenEndpointAuthSigningAlg The algorithm used for signing the JWT
type RegistrationPropertiesTokenEndpointAuthSigningAlg string

// DeleteDataRecipientRegistrationParams defines parameters for DeleteDataRecipientRegistration.
type DeleteDataRecipientRegistrationParams struct {
	// Authorization An Authorisation Token as per **[[RFC6750]](#nref-RFC6750)**
	Authorization string `json:"Authorization"`
}

// GetClientRegistrationParams defines parameters for GetClientRegistration.
type GetClientRegistrationParams struct {
	// Authorization An Authorisation Token as per **[[RFC6750]](#nref-RFC6750)**
	Authorization string `json:"Authorization"`
}

// PutDataRecipientRegistrationParams defines parameters for PutDataRecipientRegistration.
type PutDataRecipientRegistrationParams struct {
	// Authorization An Authorisation Token as per **[[RFC6750]](#nref-RFC6750)**
	Authorization string `json:"Authorization"`
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// PostDataRecipientRegistration request with any body
	PostDataRecipientRegistrationWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteDataRecipientRegistration request
	DeleteDataRecipientRegistration(ctx context.Context, clientId string, params *DeleteDataRecipientRegistrationParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetClientRegistration request
	GetClientRegistration(ctx context.Context, clientId string, params *GetClientRegistrationParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutDataRecipientRegistration request with any body
	PutDataRecipientRegistrationWithBody(ctx context.Context, clientId string, params *PutDataRecipientRegistrationParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) PostDataRecipientRegistrationWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostDataRecipientRegistrationRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteDataRecipientRegistration(ctx context.Context, clientId string, params *DeleteDataRecipientRegistrationParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteDataRecipientRegistrationRequest(c.Server, clientId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetClientRegistration(ctx context.Context, clientId string, params *GetClientRegistrationParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetClientRegistrationRequest(c.Server, clientId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PutDataRecipientRegistrationWithBody(ctx context.Context, clientId string, params *PutDataRecipientRegistrationParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutDataRecipientRegistrationRequestWithBody(c.Server, clientId, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewPostDataRecipientRegistrationRequestWithBody generates requests for PostDataRecipientRegistration with any type of body
func NewPostDataRecipientRegistrationRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/register")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeleteDataRecipientRegistrationRequest generates requests for DeleteDataRecipientRegistration
func NewDeleteDataRecipientRegistrationRequest(server string, clientId string, params *DeleteDataRecipientRegistrationParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "ClientId", runtime.ParamLocationPath, clientId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/register/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var headerParam0 string

	headerParam0, err = runtime.StyleParamWithLocation("simple", false, "Authorization", runtime.ParamLocationHeader, params.Authorization)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", headerParam0)

	return req, nil
}

// NewGetClientRegistrationRequest generates requests for GetClientRegistration
func NewGetClientRegistrationRequest(server string, clientId string, params *GetClientRegistrationParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "ClientId", runtime.ParamLocationPath, clientId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/register/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var headerParam0 string

	headerParam0, err = runtime.StyleParamWithLocation("simple", false, "Authorization", runtime.ParamLocationHeader, params.Authorization)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", headerParam0)

	return req, nil
}

// NewPutDataRecipientRegistrationRequestWithBody generates requests for PutDataRecipientRegistration with any type of body
func NewPutDataRecipientRegistrationRequestWithBody(server string, clientId string, params *PutDataRecipientRegistrationParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "ClientId", runtime.ParamLocationPath, clientId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/register/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	var headerParam0 string

	headerParam0, err = runtime.StyleParamWithLocation("simple", false, "Authorization", runtime.ParamLocationHeader, params.Authorization)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", headerParam0)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// PostDataRecipientRegistration request with any body
	PostDataRecipientRegistrationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostDataRecipientRegistrationResponse, error)

	// DeleteDataRecipientRegistration request
	DeleteDataRecipientRegistrationWithResponse(ctx context.Context, clientId string, params *DeleteDataRecipientRegistrationParams, reqEditors ...RequestEditorFn) (*DeleteDataRecipientRegistrationResponse, error)

	// GetClientRegistration request
	GetClientRegistrationWithResponse(ctx context.Context, clientId string, params *GetClientRegistrationParams, reqEditors ...RequestEditorFn) (*GetClientRegistrationResponse, error)

	// PutDataRecipientRegistration request with any body
	PutDataRecipientRegistrationWithBodyWithResponse(ctx context.Context, clientId string, params *PutDataRecipientRegistrationParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutDataRecipientRegistrationResponse, error)
}

type PostDataRecipientRegistrationResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *RegistrationProperties
	JSON400      *RegistrationError
}

// Status returns HTTPResponse.Status
func (r PostDataRecipientRegistrationResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostDataRecipientRegistrationResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteDataRecipientRegistrationResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r DeleteDataRecipientRegistrationResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteDataRecipientRegistrationResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetClientRegistrationResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *RegistrationProperties
}

// Status returns HTTPResponse.Status
func (r GetClientRegistrationResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetClientRegistrationResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PutDataRecipientRegistrationResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *RegistrationProperties
	JSON400      *RegistrationError
}

// Status returns HTTPResponse.Status
func (r PutDataRecipientRegistrationResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PutDataRecipientRegistrationResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// PostDataRecipientRegistrationWithBodyWithResponse request with arbitrary body returning *PostDataRecipientRegistrationResponse
func (c *ClientWithResponses) PostDataRecipientRegistrationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostDataRecipientRegistrationResponse, error) {
	rsp, err := c.PostDataRecipientRegistrationWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostDataRecipientRegistrationResponse(rsp)
}

// DeleteDataRecipientRegistrationWithResponse request returning *DeleteDataRecipientRegistrationResponse
func (c *ClientWithResponses) DeleteDataRecipientRegistrationWithResponse(ctx context.Context, clientId string, params *DeleteDataRecipientRegistrationParams, reqEditors ...RequestEditorFn) (*DeleteDataRecipientRegistrationResponse, error) {
	rsp, err := c.DeleteDataRecipientRegistration(ctx, clientId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteDataRecipientRegistrationResponse(rsp)
}

// GetClientRegistrationWithResponse request returning *GetClientRegistrationResponse
func (c *ClientWithResponses) GetClientRegistrationWithResponse(ctx context.Context, clientId string, params *GetClientRegistrationParams, reqEditors ...RequestEditorFn) (*GetClientRegistrationResponse, error) {
	rsp, err := c.GetClientRegistration(ctx, clientId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetClientRegistrationResponse(rsp)
}

// PutDataRecipientRegistrationWithBodyWithResponse request with arbitrary body returning *PutDataRecipientRegistrationResponse
func (c *ClientWithResponses) PutDataRecipientRegistrationWithBodyWithResponse(ctx context.Context, clientId string, params *PutDataRecipientRegistrationParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutDataRecipientRegistrationResponse, error) {
	rsp, err := c.PutDataRecipientRegistrationWithBody(ctx, clientId, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutDataRecipientRegistrationResponse(rsp)
}

// ParsePostDataRecipientRegistrationResponse parses an HTTP response from a PostDataRecipientRegistrationWithResponse call
func ParsePostDataRecipientRegistrationResponse(rsp *http.Response) (*PostDataRecipientRegistrationResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostDataRecipientRegistrationResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest RegistrationProperties
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest RegistrationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParseDeleteDataRecipientRegistrationResponse parses an HTTP response from a DeleteDataRecipientRegistrationWithResponse call
func ParseDeleteDataRecipientRegistrationResponse(rsp *http.Response) (*DeleteDataRecipientRegistrationResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteDataRecipientRegistrationResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseGetClientRegistrationResponse parses an HTTP response from a GetClientRegistrationWithResponse call
func ParseGetClientRegistrationResponse(rsp *http.Response) (*GetClientRegistrationResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetClientRegistrationResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest RegistrationProperties
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParsePutDataRecipientRegistrationResponse parses an HTTP response from a PutDataRecipientRegistrationWithResponse call
func ParsePutDataRecipientRegistrationResponse(rsp *http.Response) (*PutDataRecipientRegistrationResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutDataRecipientRegistrationResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest RegistrationProperties
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest RegistrationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}
