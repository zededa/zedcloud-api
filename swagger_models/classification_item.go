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

// ClassificationItem classification item
//
// swagger:model ClassificationItem
type ClassificationItem struct {

	// dest node
	DestNode string `json:"destNode,omitempty"`

	// level
	Level int64 `json:"level,omitempty"`

	// metric value
	MetricValue string `json:"metricValue,omitempty"`

	// source node
	SourceNode string `json:"sourceNode,omitempty"`
}

// Validate validates this classification item
func (m *ClassificationItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this classification item based on context it is used
func (m *ClassificationItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClassificationItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClassificationItem) UnmarshalBinary(b []byte) error {
	var res ClassificationItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
