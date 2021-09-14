// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WifiConfigcryptoblock wifi configcryptoblock
//
// swagger:model WifiConfigcryptoblock
type WifiConfigcryptoblock struct {

	// identity
	Identity string `json:"identity,omitempty"`

	// password
	Password string `json:"password,omitempty"`
}

// Validate validates this wifi configcryptoblock
func (m *WifiConfigcryptoblock) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this wifi configcryptoblock based on context it is used
func (m *WifiConfigcryptoblock) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WifiConfigcryptoblock) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WifiConfigcryptoblock) UnmarshalBinary(b []byte) error {
	var res WifiConfigcryptoblock
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
