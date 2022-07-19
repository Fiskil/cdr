package cdr

import (
	"errors"
	"fmt"
	"io"
	"net/url"
)

var (
	// ErrInvalidKeys is when an invalid mTLS certificate is supplied.
	ErrInvalidKeys = errors.New("cdr : invalid mTLS certificate")
)

// ErrNon2xxResponse is returned when a non 2xx status code is received.
type ErrNon2xxResponse struct {
	StatusCode int
	Response   io.ReadCloser
	URL        *url.URL
}

func (e *ErrNon2xxResponse) Error() string {
	return fmt.Sprintf("cdr : non 2xx status returned : got %d", e.StatusCode)
}
