// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserUsagePerEnterprise user usage per enterprise
//
// swagger:model UserUsagePerEnterprise
type UserUsagePerEnterprise struct {

	// Enterprise id
	// Read Only: true
	// Pattern: [0-9A-Za-z-]+
	EntpID string `json:"entpId,omitempty"`

	// User usage for that enterprise
	UserUsage string `json:"userUsage,omitempty"`
}

// Validate validates this user usage per enterprise
func (m *UserUsagePerEnterprise) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntpID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserUsagePerEnterprise) validateEntpID(formats strfmt.Registry) error {
	if swag.IsZero(m.EntpID) { // not required
		return nil
	}

	if err := validate.Pattern("entpId", "body", m.EntpID, `[0-9A-Za-z-]+`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this user usage per enterprise based on the context it is used
func (m *UserUsagePerEnterprise) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntpID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserUsagePerEnterprise) contextValidateEntpID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "entpId", "body", string(m.EntpID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserUsagePerEnterprise) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserUsagePerEnterprise) UnmarshalBinary(b []byte) error {
	var res UserUsagePerEnterprise
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
