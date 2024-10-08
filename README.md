# cdr

[![GoDoc](https://godoc.org/github.com/fiskil/cdr?status.svg)](https://godoc.org/github.com/fiskil/cdr)
[![Go Report Card](https://goreportcard.com/badge/github.com/fiskil/cdr)](https://goreportcard.com/report/github.com/fiskil/cdr)

This is a library and set of tools to ease some of the issues we at Fiskil have had with configuring and testing interactions with other CDR entities.

## Who is this library for?

This library is for accredited data recipients. This client uses the CDR specification which is restricted to data recipients registered with CDR. To become an ADR follow the instructions [here](https://www.cdr.gov.au/for-providers/become-accredited-data-recipient); alternatively, [Fiskil](https://fiskil.com.au/) provides similar APIs without the requirement of becoming a data recipient. If Fiskil isn't for you there is a list of other data recipients [here](https://www.cdr.gov.au/find-a-provider?providerType=Data%2520Recipient).

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

## Testing

An easy way to test interacting with the CDR is to setup a temporary test. The file `my_test.go` and directory `my_tests` are added to gitignore to encourage you to write your own stockpile of examples and references.

```go
func TestTransactions(t *testing.T) {

	// Arrange
	is := is.New(t)
	ctx := context.Background()
	cli, err := cdr.NewFromEnv()
	is.NoErr(err)
	cdrCli, err := banking.NewClientWithResponses("https://resource.cdr-api.bankaust.com.au/cds-au/v1", banking.WithHTTPClient(cli))
	is.NoErr(err)
	tok := "insert-your-token-here"

	// Act
	res, err := cdrCli.ListAccountsWithResponse(ctx, &banking.ListAccountsParams{
		XV: "1",
	}, func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tok))
		return nil
	})

	// Assert
	is.NoErr(err)
	if res.StatusCode() > 299 {
		is.NoErr(fmt.Errorf("non 2xx response"))
	}

	bytes, err := json.Marshal(res.JSON200.Data)
	fmt.Println(string(bytes), err)

}

// Output: {"accounts":[{"accountId":"1","creationDate":"2022-05-19","displayName":"Everyday Access","isOwned":true,"maskedNumber":"xxxx7889","openStatus":"OPEN","productCategory":"TRANS_AND_SAVINGS_ACCOUNTS","productName":"Everyday Access"},{"accountId":"2","creationDate":"2022-06-04","displayName":"Bonus Saver","isOwned":true,"maskedNumber":"xxxx0241","nickname":"Bonus Saver","openStatus":"OPEN","productCategory":"TRANS_AND_SAVINGS_ACCOUNTS","productName":"Bonus Saver"}]}
```

## Generating Go

Note that the cdr openapi specifications are obtained from [here](https://github.com/ConsumerDataStandardsAustralia/standards/tree/master/slate/source/includes/swagger). 

```bash
resource=banking
cd ${resource}
oapi-codegen -config config.yaml cdr_${resource}.swagger.json > ${resource}.gen.go
cd -
```

However, for energy additional patching is required to be compatible with AER energy API endpoints. 
```bash
resource=energy
cd ${resource}
cat cdr_energy.swagger.json | json-patch -p cdr_energy.swagger.json.patch > patched_cdr_energy.swagger.json
oapi-codegen -config config.yaml patched_cdr_${resource}.swagger.json > ${resource}.gen.go
cd -
```

When new cdr resource type definitions are published they can be appended to the current package using the `diff-gen` subcommand of the `cdr` cli.
`diff-gen` will generate go models from the provided openAPI/Swagger definition but omit those already declared in the package. 
```bash
./cdr diff-gen ./energy ./energy/cdr_energy.swagger.1.2.4.json 
```