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

// SANValues s a n values
//
// swagger:model SANValues
type SANValues struct {

	// List of permitted DNS names.
	DNS []string `json:"dns"`

	// List of permitted email addresses.
	EmaildIds []string `json:"emaildIds"`

	// List of permitted hosts.
	Hosts []string `json:"hosts"`

	// List of permitted IP addresses.
	Ips []string `json:"ips"`

	// List of permitted User principal names.
	Upns []string `json:"upns"`

	// List of permitted URIs.
	Uris []string `json:"uris"`
}

// Validate validates this s a n values
func (m *SANValues) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this s a n values based on context it is used
func (m *SANValues) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SANValues) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SANValues) UnmarshalBinary(b []byte) error {
	var res SANValues
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
