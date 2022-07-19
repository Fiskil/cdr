package storage_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fiskil/cdr/storage"
	"github.com/jrapoport/chestnut/encryptor/crypto"
	"github.com/matryer/is"
)

func TestStorage(t *testing.T) {

	// Arrange
	is := is.New(t)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"access_token":"new-access","refresh_token":"a fresh start","cdr_arrangement_id":"a new deal"}`)
	}))
	a := storage.Arrangement{
		ID:               "my-token-id",
		RefreshToken:     "my-refresh-token",
		CDRArrangementID: "our-cdr-arrangement-id",
		TokenEndpoint:    ts.URL,
		ClientID:         "your-client-id-for-me",
	}
	store, err := storage.New("test", crypto.TextSecret("my-test-secret"))
	is.NoErr(err)
	e := store.NewArrangement(a)

	// Act
	access, err := store.AccessToken(context.Background(), "my-token-id")

	// Assert
	is.NoErr(e)
	is.Equal(access, "new-access")
	newArrangement, err := store.GetArrangement("my-token-id")
	is.Equal(newArrangement.CDRArrangementID, "a new deal")
	is.Equal(newArrangement.RefreshToken, "a fresh start")
}
