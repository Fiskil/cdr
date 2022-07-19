package cdr

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"os"
)

// New creates a new http client with certificates attached to it. Most of the funcitons in this cdr package require a *http.Client as an argument. This is the way to build that client.
//
//	cert := []byte(`-----BEGIN CERTIFICATE-----
//	MIIBhTCCASugAwIBAgIQIRi6zePL6mKjOipn+dNuaTAKBggqhkjOPQQDAjASMRAw
//	DgYDVQQKEwdBY21lIENvMB4XDTE3MTAyMDE5NDMwNloXDTE4MTAyMDE5NDMwNlow
//	EjEQMA4GA1UEChMHQWNtZSBDbzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABD0d
//	7VNhbWvZLWPuj/RtHFjvtJBEwOkhbN/BnnE8rnZR8+sbwnc/KhCk3FhnpHZnQz7B
//	5aETbbIgmuvewdjvSBSjYzBhMA4GA1UdDwEB/wQEAwICpDATBgNVHSUEDDAKBggr
//	BgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MCkGA1UdEQQiMCCCDmxvY2FsaG9zdDo1
//	NDUzgg4xMjcuMC4wLjE6NTQ1MzAKBggqhkjOPQQDAgNIADBFAiEA2zpJEPQyz6/l
//	Wf86aX6PepsntZv2GYlA5UpabfT2EZICICpJ5h/iI+i341gBmLiAFQOyTDT+/wQc
//	6MF9+Yw1Yy0t
//	-----END CERTIFICATE-----`)
//	privKey := []byte(`-----BEGIN EC PRIVATE KEY-----
//	MHcCAQEEIIrYSSNQFaA2Hwf1duRSxKtLYX5CB04fSeQ6tF1aY/PuoAoGCCqGSM49
//	AwEHoUQDQgAEPR3tU2Fta9ktY+6P9G0cWO+0kETA6SFs38GecTyudlHz6xvCdz8q
//	EKTcWGekdmdDPsHloRNtsiCa697B2O9IFA==
//	-----END EC PRIVATE KEY-----`)
//
//	client, err := cdr.New(cert, privKey)
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
