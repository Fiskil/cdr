// Package storage provides encrypted data store on your local machine that can handle the process
// of getting refresh tokens for you.
package storage

import (
	"context"
	"encoding/json"
	"os"
	"path"

	"github.com/fiskil/cdr"
	"github.com/fiskil/cdr/assertions"
	"github.com/jrapoport/chestnut"
	"github.com/jrapoport/chestnut/encryptor/aes"
	"github.com/jrapoport/chestnut/encryptor/crypto"
	"github.com/jrapoport/chestnut/storage/bolt"
)

var arrangmentBucket = "arrangements"

// Store can track your access tokens and refresh tokens for you.
type Store struct {
	db *chestnut.Chestnut
}

// New opens a new connection to your data storage and uses the proviced secret for encryption/decription.
func New(name string, secret crypto.Secret) (*Store, error) {
	confDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}
	dir := path.Join(confDir, "fiskil")

	err = os.Mkdir(dir, 0750)
	if err != nil && !os.IsExist(err) {
		return nil, err
	}

	file := path.Join(dir, name+".db")
	store := bolt.NewStore(file)
	opt := chestnut.WithAES(crypto.Key256, aes.CFB, secret)

	cn := chestnut.NewChestnut(store, opt)
	err = cn.Open()

	return &Store{cn}, err
}

// Arrangement represents an CDR arrangement with a dataholder.
type Arrangement struct {
	// ID is how a given arrangement is found.
	ID string
	// RefreshToken is the refresh token for this arrangement.
	RefreshToken string
	// CDRArrangementID is the arrangement ID provided by the dataholder.
	CDRArrangementID string
	// ClientID is the dataholders client id for you.
	ClientID string
	// TokenEndpoint is the dataholders token endpoint
	TokenEndpoint string
}

// NewArrangement saves a new arrangement into the local data store. After it is stored you will be able to fetch new access tokens automatically without worrying about refresh token exchange.
func (s *Store) NewArrangement(a Arrangement) error {
	raw, err := json.Marshal(a)
	if err != nil {
		return err
	}

	return s.db.Put(arrangmentBucket, []byte(a.ID), raw)
}

// GetArrangement finds an arrangement for a given ID.
func (s *Store) GetArrangement(id string) (Arrangement, error) {
	raw, err := s.db.Get(arrangmentBucket, []byte(id))
	if err != nil {
		return Arrangement{}, err
	}

	var arrangement Arrangement
	if err := json.Unmarshal(raw, &arrangement); err != nil {
		return Arrangement{}, err
	}

	return arrangement, nil
}

// AccessToken gets a new access token and refreshes the token in the database.
// This is an oppinionated function and relies on environment variables being configured for getting an MTLS certificate and signing a client assertion token.
// For a more flexible method of getting an access token use GetArrangement and cdr.RefreshToken.
func (s *Store) AccessToken(ctx context.Context, id string) (string, error) {

	arrangement, err := s.GetArrangement(id)
	if err != nil {
		return "", err
	}

	cli, err := cdr.NewFromEnv()
	if err != nil {
		return "", err
	}

	signer, err := assertions.NewSignerFromEnv()
	if err != nil {
		return "", err
	}
	assertion, err := signer.ClientAssertions(arrangement.ClientID, arrangement.TokenEndpoint)
	if err != nil {
		return "", err
	}

	refresh, err := cdr.RefreshToken(ctx, cli, arrangement.TokenEndpoint, arrangement.RefreshToken, arrangement.ClientID, assertion)
	if err != nil {
		return "", err
	}

	arrangement.RefreshToken = refresh.RefreshToken
	arrangement.CDRArrangementID = refresh.CDRArrangement

	if err := s.NewArrangement(arrangement); err != nil {
		return "", err
	}

	return refresh.AccessToken, nil
}
