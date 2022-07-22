package cdr_test

import (
	"context"
	"testing"

	"github.com/fiskil/cdr"
	"github.com/fiskil/cdr/data"
	"github.com/matryer/is"
)

func TestGetWellKnownConfigFromCDR(t *testing.T) {

	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	// Arrange
	is := is.New(t)
	ctx := context.Background()
	cli, err := cdr.NewFromEnv()
	is.NoErr(err)
	path := data.CDRBaseURL + "/idp"

	// Act
	res, err := cdr.GetWellKnownConfig(ctx, cli, path)

	// Assert
	is.NoErr(err)
	is.Equal(res.Issuer, "https://secure.api.cdr.gov.au/idp")
}
