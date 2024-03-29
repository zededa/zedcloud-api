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

// MemorySummary Memory Summary
//
// # Memory Summary
//
// swagger:model MemorySummary
type MemorySummary struct {

	// Total allocated memory in MBs
	AllocatedMB float64 `json:"AllocatedMB,omitempty"`

	// Total memory for the device in MBs
	TotalMB float64 `json:"TotalMB,omitempty"`

	// Total memory used by the device in MBs within allocated memory
	UsedMB float64 `json:"UsedMB,omitempty"`
}

// Validate validates this memory summary
func (m *MemorySummary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this memory summary based on context it is used
func (m *MemorySummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MemorySummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MemorySummary) UnmarshalBinary(b []byte) error {
	var res MemorySummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
