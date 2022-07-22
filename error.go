package cdr

import (
	"fmt"
	"io"
	"net/url"
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
