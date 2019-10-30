package stackpath

import (
	"net/url"
	"strings"

	"golang.org/x/oauth2"
)

// InvalidClientSecretError models when a StackPath API OAuth 2 client ID is
// invalid, either due to an invalid format or because the client ID does not
// exist at StackPath.
type InvalidClientIDError struct {
	error
	Err error
}

// NewInvalidClientIDError wraps an existing error as an invalid client ID error.
func NewInvalidClientIDError(err error) *InvalidClientIDError {
	return &InvalidClientIDError{Err: err}
}

// Error returns a human-readable invalid client ID error message.
func (e *InvalidClientIDError) Error() string {
	return "invalid or unknown StackPath client ID"
}

// InvalidClientSecretError models when a StackPath API OAuth 2 client ID is
// correct, but the client secret is incorrect.
type InvalidClientSecretError struct {
	error
	Err error
}

// NewInvalidClientSecretError wraps an existing error as an invalid client
// secret error.
func NewInvalidClientSecretError(err error) *InvalidClientSecretError {
	return &InvalidClientSecretError{Err: err}
}

// Error returns a human-readable invalid client secret error message.
func (e *InvalidClientSecretError) Error() string {
	return "invalid StackPath client secret"
}

// NewStackPathError factories common StackPath error scenarios into their own
// error types, or returns the original error.
func NewStackPathError(err error) error {
	switch err.(type) {
	// Look for errors performing underlying OAuth 2 authentication.
	case *url.Error:
		switch err.(*url.Error).Err.(type) {
		case *oauth2.RetrieveError:
			if strings.HasPrefix(err.(*url.Error).Err.(*oauth2.RetrieveError).Error(), "oauth2: cannot fetch token: 404 Not Found") {
				return NewInvalidClientIDError(err)
			}

			if strings.HasPrefix(err.(*url.Error).Err.(*oauth2.RetrieveError).Error(), "oauth2: cannot fetch token: 401 Unauthorized") {
				return NewInvalidClientSecretError(err)
			}
		}
	}

	return err
}
