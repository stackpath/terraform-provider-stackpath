package stackpath

import (
	"net/url"

	"golang.org/x/oauth2"
)

// InvalidClientSecretError models when a StackPath API OAuth 2 client ID is
// invalid, either due to an invalid format or because the client ID does not
// exist at StackPath.
type InvalidClientIDError struct {
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
	switch rootErr := err.(type) {
	// Look for errors performing underlying OAuth 2 authentication.
	case *url.Error:
		switch typedErr := rootErr.Err.(type) {
		case *oauth2.RetrieveError:
			switch typedErr.Response.StatusCode {
			// A 401 Unauthorized error means the client ID was valid, but the
			// corresponding secret wasn't.
			case 401:
				return NewInvalidClientSecretError(err)

			// A 404 Not Found error means the client ID didn't exist.
			case 404:
				return NewInvalidClientIDError(err)
			}
		}
	}

	return err
}
