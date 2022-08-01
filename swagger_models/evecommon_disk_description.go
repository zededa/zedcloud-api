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

// EvecommonDiskDescription DiskDescription describes disk
// we can use different data to locate disk in the system
//
// swagger:model evecommonDiskDescription
type EvecommonDiskDescription struct {

	// logical name
	LogicalName string `json:"logicalName,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// serial
	Serial string `json:"serial,omitempty"`
}

// Validate validates this evecommon disk description
func (m *EvecommonDiskDescription) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this evecommon disk description based on context it is used
func (m *EvecommonDiskDescription) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EvecommonDiskDescription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EvecommonDiskDescription) UnmarshalBinary(b []byte) error {
	var res EvecommonDiskDescription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}