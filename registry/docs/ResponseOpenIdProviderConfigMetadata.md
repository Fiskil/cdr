# ResponseOpenIdProviderConfigMetadata

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | **string** | URL using the https scheme with no query or fragment component that the CDR Register asserts as its Issuer Identifier | [default to null]
**JwksUri** | **string** | URL of the CDR Register&#x27;s JSON Web Key Set **[[JWK]](#nref-JWK)** document. This contains the signing key(s) used to validate access tokens issued from the CDR Register. Note that this differs from the JWKS endpoint used to validate SSAs and CDR Register client authentication | [default to null]
**TokenEndpoint** | **string** | URL of the CDR Register&#x27;s OAuth 2.0 Token Endpoint | [default to null]
**ClaimsSupported** | **[]string** | JSON array containing a list of the Claim Names of the Claims that the CDR Register supplies values for | [default to null]
**IdTokenSigningAlgValuesSupported** | **[]string** | JSON array containing a list of the JWS signing algorithms (alg values) supported by the CDR Register for the ID Token to encode the Claims in a JWT. Given the CDR Register does not issue ID tokens, this field can be safely ignored | [default to null]
**SubjectTypesSupported** | **[]string** | JSON array containing a list of the Subject Identifier types that the CDR Register supports. Given the CDR Register does not issue ID tokens, this field can be safely ignored | [default to null]
**CodeChallengeMethodsSupported** | **[]string** | JSON array containing a list of Proof Key for Code Exchange (PKCE) **[[RFC7636]](#nref-RFC7636)** code challenge methods supported by this authorization server. Given the CDR Register does not support PKCE, this field can be safely ignored | [default to null]
**ScopesSupported** | **[]string** | JSON array containing a list of the OAuth 2.0 **[[RFC6749]](#nref-RFC6749)** scope values that the CDR Register supports | [default to null]
**ResponseTypesSupported** | **[]string** | JSON array containing a list of the OAuth 2.0 response_type values that the CDR Registrer supports | [default to null]
**GrantTypesSupported** | **[]string** | JSON array containing a list of the OAuth 2.0 Grant Type values that the CDR Register supports | [default to null]
**TokenEndpointAuthMethodsSupported** | **[]string** | JSON array containing a list of Client Authentication methods supported by this Token Endpoint | [default to null]
**TlsClientCertificateBoundAccessTokens** | **bool** | Boolean value indicating server support for mutual TLS client certificate bound access tokens | [default to null]
**TokenEndpointAuthSigningAlgValuesSupported** | **[]string** | JSON array containing a list of the JWS signing algorithms (alg values) supported by the token endpoint for the signature on the JWT **[[JWT]](#nref-JWT)** used to authenticate the client at the token endpoint for the \\\&quot;private_key_jwt\\\&quot; authentication method | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

