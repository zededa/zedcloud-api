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

// StorageSummary Storage Summary
//
// # Storage Summary
//
// swagger:model StorageSummary
type StorageSummary struct {

	// Total reserved for running applications + temp. images etc
	AllocatedMB float64 `json:"AllocatedMB,omitempty"`

	// Total Storage for the device in MBs
	TotalMB float64 `json:"TotalMB,omitempty"`

	// How much is used within the allocated total storage
	UsedMB float64 `json:"UsedMB,omitempty"`
}

// Validate validates this storage summary
func (m *StorageSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storage summary based on context it is used
func (m *StorageSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageSummary) UnmarshalBinary(b []byte) error {
	var res StorageSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
