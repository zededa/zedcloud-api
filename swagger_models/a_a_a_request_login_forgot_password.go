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

// AAARequestLoginForgotPassword a a a request login forgot password
//
// swagger:model AAARequestLoginForgotPassword
type AAARequestLoginForgotPassword struct {

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this a a a request login forgot password
func (m *AAARequestLoginForgotPassword) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this a a a request login forgot password based on context it is used
func (m *AAARequestLoginForgotPassword) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AAARequestLoginForgotPassword) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AAARequestLoginForgotPassword) UnmarshalBinary(b []byte) error {
	var res AAARequestLoginForgotPassword
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
