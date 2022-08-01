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

// VariableOptionVal variable option val
//
// swagger:model VariableOptionVal
type VariableOptionVal struct {

	// Display label of the key in User-Agent
	Label string `json:"label,omitempty"`

	// Value of the key to be used
	Value string `json:"value,omitempty"`
}

// Validate validates this variable option val
func (m *VariableOptionVal) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this variable option val based on context it is used
func (m *VariableOptionVal) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VariableOptionVal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VariableOptionVal) UnmarshalBinary(b []byte) error {
	var res VariableOptionVal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
