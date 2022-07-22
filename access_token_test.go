package cdr_test

import (
	"context"
	"testing"

	"github.com/fiskil/cdr"
	"github.com/fiskil/cdr/data"
	"github.com/matryer/is"
)

func TestGetAccessToken(t *testing.T) {

	// Arrange
	is := is.New(t)
	ctx := context.Background()
	cli, err := cdr.NewFromEnv()
	is.NoErr(err)
	baseURL := data.BankAustBaseURL

	softwareProductID := ""
	tokenEndpoint := data.CDRBaseSecureURL

}
