package cdr

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"os"
)

// New creates a new http client with certificates attached to it. Most of the funcitons in this cdr package require a *http.Client as an argument. This is the way to build that client.
func New(cert []byte, privKey []byte) (*http.Client, error) {

	certificate, err := tls.X509KeyPair(cert, privKey)
	if err != nil {
		return nil, ErrInvalidKeys
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(cert)

	/* #nosec */
	rt := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:            caCertPool,
			Certificates:       []tls.Certificate{certificate},
			InsecureSkipVerify: true,
			Renegotiation:      tls.RenegotiateOnceAsClient,
		},
	}

	client := &http.Client{
		Transport: rt,
	}

	return client, nil
}

// NewFromEnv creates a new http client taking its certificates from environment variables. This ishelpful if the certificates are const.
// The environment variables are:
//	CDR_MTLS_CERTIFICATE
//	CDR_MTLS_PRIVATE_KEY
func NewFromEnv() (*http.Client, error) {
	cert := os.Getenv("CDR_MTLS_CERTIFICATE")
	pk := os.Getenv("CDR_MTLS_PRIVATE_KEY")

	return New([]byte(cert), []byte(pk))
}
