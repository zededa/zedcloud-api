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

// AppAccessPolicy app access policy
//
// swagger:model AppAccessPolicy
type AppAccessPolicy struct {

	// app side of edge-view access is allowed or not
	AllowApp bool `json:"allowApp,omitempty"`
}

// Validate validates this app access policy
func (m *AppAccessPolicy) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this app access policy based on context it is used
func (m *AppAccessPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppAccessPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppAccessPolicy) UnmarshalBinary(b []byte) error {
	var res AppAccessPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
