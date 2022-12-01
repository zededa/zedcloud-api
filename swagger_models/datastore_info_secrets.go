// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DatastoreInfoSecrets datastore info secrets
//
// swagger:model DatastoreInfoSecrets
type DatastoreInfoSecrets struct {

	// Datastore access API key in plain-text
	APIKey string `json:"apiKey,omitempty"`

	// Datastore access API password in plain-text
	APIPasswd string `json:"apiPasswd,omitempty"`
}

// Validate validates this datastore info secrets
func (m *DatastoreInfoSecrets) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this datastore info secrets based on context it is used
func (m *DatastoreInfoSecrets) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatastoreInfoSecrets) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatastoreInfoSecrets) UnmarshalBinary(b []byte) error {
	var res DatastoreInfoSecrets
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
