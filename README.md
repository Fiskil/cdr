# cdr

[![GoDoc](https://godoc.org/github.com/fiskil/cdr?status.svg)](https://godoc.org/github.com/fiskil/cdr)
[![Go Report Card](https://goreportcard.com/badge/github.com/fiskil/cdr)](https://goreportcard.com/report/github.com/fiskil/cdr)


This is a library and set of tools to ease some of the issues we at Fiskil have had with configuring and testing interactions with other CDR entities.

This is still a work in progress and does not have complete coverage over the CDR specification.

## Getting started with the library

### mTLS certs

Almost all of the endpoints specified by the spec require an mTLS certificate to verify clients. The functions in this library assume that mTLS certificates have already been attached to the provided http client, but also provide a way of buliding an http client with certificates:

```go
cert := []byte(`-----BEGIN CERTIFICATE-----
MIIBhTCCASugAwIBAgIQIRi6zePL6mKjOipn+dNuaTAKBggqhkjOPQQDAjASMRAw
DgYDVQQKEwdBY21lIENvMB4XDTE3MTAyMDE5NDMwNloXDTE4MTAyMDE5NDMwNlow
EjEQMA4GA1UEChMHQWNtZSBDbzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABD0d
7VNhbWvZLWPuj/RtHFjvtJBEwOkhbN/BnnE8rnZR8+sbwnc/KhCk3FhnpHZnQz7B
5aETbbIgmuvewdjvSBSjYzBhMA4GA1UdDwEB/wQEAwICpDATBgNVHSUEDDAKBggr
BgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MCkGA1UdEQQiMCCCDmxvY2FsaG9zdDo1
NDUzgg4xMjcuMC4wLjE6NTQ1MzAKBggqhkjOPQQDAgNIADBFAiEA2zpJEPQyz6/l
Wf86aX6PepsntZv2GYlA5UpabfT2EZICICpJ5h/iI+i341gBmLiAFQOyTDT+/wQc
6MF9+Yw1Yy0t
-----END CERTIFICATE-----`)

privKey := []byte(`-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIIrYSSNQFaA2Hwf1duRSxKtLYX5CB04fSeQ6tF1aY/PuoAoGCCqGSM49
AwEHoUQDQgAEPR3tU2Fta9ktY+6P9G0cWO+0kETA6SFs38GecTyudlHz6xvCdz8q
EKTcWGekdmdDPsHloRNtsiCa697B2O9IFA==
-----END EC PRIVATE KEY-----`)

cli, err := cdr.New(cert, privKey)
```

The `cdr.New` returns an `*http.Client` with certificates attached. This http client can now be used as an argument to other functions in the libary.

```go 
transactions, err := cdr.GetTransactionsByAccount(ctx, cli, other args ...)
```

### Access Tokens

We have found that getting an access token for a specific data holder quickly can be challenging. To deal with this issue there is a tool attached to this library for managing refresh tokens and cdr arrangements. 

Refresh tokens are stored locally with encryption at rest and associated access tokens can be easily accessed.

```bash

# To create a new token named "aus_bank" run:
go run ./cmd access set aus_bank

# You will be prompted for an encryption secret, refresh token, cdr arrangment id, and various information about the data holder.

# Once a token has been set you can access it:
go run ./cmd access fetch aus_bank

# You will again be prompted for the decryption secret.
```
