/*
 * CDR Participant Discovery API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.17.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Response containing the Open ID Provider Configuration Metadata
type ResponseOpenIdProviderConfigMetadata struct {
	// URL using the https scheme with no query or fragment component that the CDR Register asserts as its Issuer Identifier
	Issuer string `json:"issuer"`
	// URL of the CDR Register's JSON Web Key Set **[[JWK]](#nref-JWK)** document. This contains the signing key(s) used to validate access tokens issued from the CDR Register. Note that this differs from the JWKS endpoint used to validate SSAs and CDR Register client authentication
	JwksUri string `json:"jwks_uri"`
	// URL of the CDR Register's OAuth 2.0 Token Endpoint
	TokenEndpoint string `json:"token_endpoint"`
	// JSON array containing a list of the Claim Names of the Claims that the CDR Register supplies values for
	ClaimsSupported []string `json:"claims_supported"`
	// JSON array containing a list of the JWS signing algorithms (alg values) supported by the CDR Register for the ID Token to encode the Claims in a JWT. Given the CDR Register does not issue ID tokens, this field can be safely ignored
	IdTokenSigningAlgValuesSupported []string `json:"id_token_signing_alg_values_supported"`
	// JSON array containing a list of the Subject Identifier types that the CDR Register supports. Given the CDR Register does not issue ID tokens, this field can be safely ignored
	SubjectTypesSupported []string `json:"subject_types_supported"`
	// JSON array containing a list of Proof Key for Code Exchange (PKCE) **[[RFC7636]](#nref-RFC7636)** code challenge methods supported by this authorization server. Given the CDR Register does not support PKCE, this field can be safely ignored
	CodeChallengeMethodsSupported []string `json:"code_challenge_methods_supported"`
	// JSON array containing a list of the OAuth 2.0 **[[RFC6749]](#nref-RFC6749)** scope values that the CDR Register supports
	ScopesSupported []string `json:"scopes_supported"`
	// JSON array containing a list of the OAuth 2.0 response_type values that the CDR Registrer supports
	ResponseTypesSupported []string `json:"response_types_supported"`
	// JSON array containing a list of the OAuth 2.0 Grant Type values that the CDR Register supports
	GrantTypesSupported []string `json:"grant_types_supported"`
	// JSON array containing a list of Client Authentication methods supported by this Token Endpoint
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported"`
	// Boolean value indicating server support for mutual TLS client certificate bound access tokens
	TlsClientCertificateBoundAccessTokens bool `json:"tls_client_certificate_bound_access_tokens"`
	// JSON array containing a list of the JWS signing algorithms (alg values) supported by the token endpoint for the signature on the JWT **[[JWT]](#nref-JWT)** used to authenticate the client at the token endpoint for the \\\"private_key_jwt\\\" authentication method
	TokenEndpointAuthSigningAlgValuesSupported []string `json:"token_endpoint_auth_signing_alg_values_supported"`
}