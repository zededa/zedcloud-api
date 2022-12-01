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

// AppInstSummaryReport app inst summary report
//
// swagger:model AppInstSummaryReport
type AppInstSummaryReport struct {

	// total app insts
	TotalAppInsts int64 `json:"totalAppInsts,omitempty"`
}

// Validate validates this app inst summary report
func (m *AppInstSummaryReport) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this app inst summary report based on context it is used
func (m *AppInstSummaryReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppInstSummaryReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppInstSummaryReport) UnmarshalBinary(b []byte) error {
	var res AppInstSummaryReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
