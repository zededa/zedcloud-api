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

// IAMFilter i a m filter
//
// swagger:model IAMFilter
type IAMFilter struct {

	// all
	All bool `json:"all,omitempty"`

	// entpstate
	Entpstate string `json:"entpstate,omitempty"`

	// hubspotid
	Hubspotid string `json:"hubspotid,omitempty"`

	// name pattern
	NamePattern string `json:"namePattern,omitempty"`

	// projectid
	Projectid []string `json:"projectid"`

	// role name
	RoleName string `json:"roleName,omitempty"`

	// sfdcid
	Sfdcid string `json:"sfdcid,omitempty"`

	// userstate
	Userstate string `json:"userstate,omitempty"`
}

// Validate validates this i a m filter
func (m *IAMFilter) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this i a m filter based on context it is used
func (m *IAMFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IAMFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IAMFilter) UnmarshalBinary(b []byte) error {
	var res IAMFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
