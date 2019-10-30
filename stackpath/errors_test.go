package stackpath

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"golang.org/x/oauth2"
)

func TestBuildAnInvalidClientIDError(t *testing.T) {
	err := NewStackPathError(&url.Error{
		Err: &oauth2.RetrieveError{
			Response: &http.Response{
				StatusCode: 404,
			},
		},
	})

	switch err.(type) {
	case *InvalidClientIDError:
		if err.Error() != "invalid or unknown StackPath client ID" {
			t.Errorf("An invalid client ID error has the unexpected error message \"%s\"", err.Error())
		}
	default:
		t.Errorf("NewStackPathError() built an incorrect error type. Expected InvalidClientIDError but got %t", err)
	}
}

func TestBuildAnInvalidClientSecretError(t *testing.T) {
	err := NewStackPathError(&url.Error{
		Err: &oauth2.RetrieveError{
			Response: &http.Response{
				StatusCode: 401,
			},
		},
	})

	switch err.(type) {
	case *InvalidClientSecretError:
		if err.Error() != "invalid StackPath client secret" {
			t.Errorf("An invalid client ID error has the unexpected error message \"%s\"", err.Error())
		}
	default:
		t.Errorf("NewStackPathError() built an incorrect error type. Expected InvalidClientSecretError but got %t", err)
	}
}

func TestBuildANonStackPathError(t *testing.T) {
	type TestError struct {
		error
	}

	err := NewStackPathError(&TestError{errors.New("foo")})

	switch err.(type) {
	case *TestError:
	default:
		t.Errorf("NewStackPathError built an incorrect error type. Expected TestError but got %t", err)
	}
}
