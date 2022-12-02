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

// VifInfo vif info
//
// swagger:model VifInfo
type VifInfo struct {

	// app name
	AppName string `json:"appName,omitempty"`

	// mac address
	MacAddress string `json:"macAddress,omitempty"`

	// vif name
	VifName string `json:"vifName,omitempty"`
}

// Validate validates this vif info
func (m *VifInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this vif info based on context it is used
func (m *VifInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VifInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VifInfo) UnmarshalBinary(b []byte) error {
	var res VifInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
