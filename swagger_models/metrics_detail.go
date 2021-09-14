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

// MetricsDetail metrics detail
//
// swagger:model MetricsDetail
type MetricsDetail struct {

	// queries
	Queries map[string]string `json:"queries,omitempty"`

	// results
	Results map[string]string `json:"results,omitempty"`
}

// Validate validates this metrics detail
func (m *MetricsDetail) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this metrics detail based on context it is used
func (m *MetricsDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MetricsDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricsDetail) UnmarshalBinary(b []byte) error {
	var res MetricsDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
