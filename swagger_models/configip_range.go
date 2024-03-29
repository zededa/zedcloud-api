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

// ConfigipRange configip range
//
// swagger:model configipRange
type ConfigipRange struct {

	// end
	End string `json:"end,omitempty"`

	// start
	Start string `json:"start,omitempty"`
}

// Validate validates this configip range
func (m *ConfigipRange) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this configip range based on context it is used
func (m *ConfigipRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConfigipRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigipRange) UnmarshalBinary(b []byte) error {
	var res ConfigipRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
