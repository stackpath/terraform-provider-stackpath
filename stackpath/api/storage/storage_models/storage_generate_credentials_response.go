// Code generated by go-swagger; DO NOT EDIT.

package storage_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageGenerateCredentialsResponse A response with new credentials
//
// swagger:model storageGenerateCredentialsResponse
type StorageGenerateCredentialsResponse struct {

	// The ID for the access key
	AccessKey string `json:"accessKey,omitempty"`

	// The secret key used to sign requests
	SecretKey string `json:"secretKey,omitempty"`
}

// Validate validates this storage generate credentials response
func (m *StorageGenerateCredentialsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storage generate credentials response based on context it is used
func (m *StorageGenerateCredentialsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageGenerateCredentialsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageGenerateCredentialsResponse) UnmarshalBinary(b []byte) error {
	var res StorageGenerateCredentialsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
