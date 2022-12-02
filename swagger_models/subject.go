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

// Subject subject
//
// swagger:model Subject
type Subject struct {

	// Certificate common name.
	CommonName string `json:"commonName,omitempty"`

	// List of countries.
	Country []string `json:"country"`

	// List of locallity.
	Locality []string `json:"locality"`

	// List of organization.
	Organization []string `json:"organization"`

	// List of Organizational Unit.
	OrganizationalUnit []string `json:"organizationalUnit"`

	// List of Postal codes.
	PostalCode []string `json:"postalCode"`

	// List of List of Prvince.
	Province []string `json:"province"`

	// Subject cerial number
	SerialNumber string `json:"serialNumber,omitempty"`
}

// Validate validates this subject
func (m *Subject) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this subject based on context it is used
func (m *Subject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Subject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Subject) UnmarshalBinary(b []byte) error {
	var res Subject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
