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

// DeviceDistribution device distribution
//
// swagger:model DeviceDistribution
type DeviceDistribution struct {

	// device count
	DeviceCount int64 `json:"deviceCount,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`
}

// Validate validates this device distribution
func (m *DeviceDistribution) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this device distribution based on context it is used
func (m *DeviceDistribution) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceDistribution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceDistribution) UnmarshalBinary(b []byte) error {
	var res DeviceDistribution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
