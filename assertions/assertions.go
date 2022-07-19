// Package assertions is used for signing assertions that are used when authenticating with a data holder.
package assertions

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/segmentio/ksuid"
)

// SingleKeySigner can sign assertions with a single pub private key.
type SingleKeySigner struct {
	kid    string
	pk     *rsa.PrivateKey
	method jwt.SigningMethod
}

// NewSingleKeySigner creates a new single key signer.
//
//	kid := "fiskil_kid_sig"
//	privKey := []byte(`-----BEGIN EC PRIVATE KEY-----
//	MHcCAQEEIIrYSSNQFaA2Hwf1duRSxKtLYX5CB04fSeQ6tF1aY/PuoAoGCCqGSM49
//	AwEHoUQDQgAEPR3tU2Fta9ktY+6P9G0cWO+0kETA6SFs38GecTyudlHz6xvCdz8q
//	EKTcWGekdmdDPsHloRNtsiCa697B2O9IFA==
//	-----END EC PRIVATE KEY-----`)
//
//	signer, err := assertions.NewSingleKeySigner(kid, privKey)
func NewSingleKeySigner(kid string, pk []byte) (*SingleKeySigner, error) {

	privBlock, _ := pem.Decode(pk)
	if privBlock == nil {
		return nil, errors.New("failed decoding pem key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(privBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed decoding pem key: %w", err)
	}

	method := jwt.GetSigningMethod("PS256")

	return &SingleKeySigner{
		pk:     privateKey,
		method: method,
		kid:    kid,
	}, nil
}

// NewSignerFromEnv creates a new signer from environment variables.
//	CDR_SIGNER_KID
//	CDR_SIGNER_PRIVATE_KEY
func NewSignerFromEnv() (*SingleKeySigner, error) {
	kid := os.Getenv("CDR_SIGNER_KID")
	pk := os.Getenv("CDR_SIGNER_PRIVATE_KEY")

	return NewSingleKeySigner(kid, []byte(pk))
}

// GenerateToken generates a signed JWT token string representing the user Claims.
func (s *SingleKeySigner) GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(s.method, claims)
	token.Header["kid"] = s.kid

	str, err := token.SignedString(s.pk)
	if err != nil {
		return "", fmt.Errorf("signing token: %w", err)
	}

	return str, nil
}

// ClientAssertions generates and signs client assertions used to authenticate against a bank.
//
// The sub field is (normally) your data holder specific client id. While the audience will be the endpoint your are requesting from.
//
//	signer := cdr.NewSingleKeySignerFromEnv()
//	token, err := signer.ClientAssertions("my-client-id-with-bank-australia", "https://identity-mtls.cdr-api.bankaust.com.au/par")
func (s *SingleKeySigner) ClientAssertions(sub string, aud string) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(10 * time.Minute)),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ID:        ksuid.New().String(),
		Issuer:    sub,
		Subject:   sub,
		Audience:  []string{aud},
	}

	return s.GenerateToken(claims)
}
