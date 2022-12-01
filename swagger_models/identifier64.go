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

// Identifier64 Generic identifier, base64 encoding
//
// swagger:model Identifier64
type Identifier64 struct {

	// base64
	Base64 string `json:"base64,omitempty"`
}

// Validate validates this identifier64
func (m *Identifier64) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this identifier64 based on context it is used
func (m *Identifier64) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Identifier64) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Identifier64) UnmarshalBinary(b []byte) error {
	var res Identifier64
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
