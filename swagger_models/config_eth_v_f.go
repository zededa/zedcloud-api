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

// ConfigEthVF Information regarding Virtual Function (VF) customisation
//
// swagger:model configEthVF
type ConfigEthVF struct {

	// index
	Index int64 `json:"index,omitempty"`

	// mac
	Mac string `json:"mac,omitempty"`

	// vlan Id
	VlanID int64 `json:"vlanId,omitempty"`
}

// Validate validates this config eth v f
func (m *ConfigEthVF) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this config eth v f based on context it is used
func (m *ConfigEthVF) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConfigEthVF) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigEthVF) UnmarshalBinary(b []byte) error {
	var res ConfigEthVF
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}