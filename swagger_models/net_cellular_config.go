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

// NetCellularConfig net cellular config
//
// swagger:model NetCellularConfig
type NetCellularConfig struct {

	// a p n
	APN string `json:"APN,omitempty"`

	// location tracking
	LocationTracking bool `json:"locationTracking,omitempty"`
}

// Validate validates this net cellular config
func (m *NetCellularConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this net cellular config based on context it is used
func (m *NetCellularConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetCellularConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetCellularConfig) UnmarshalBinary(b []byte) error {
	var res NetCellularConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
