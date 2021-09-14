// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AAAFailureResponseGenerateToken a a a failure response generate token
//
// swagger:model AAAFailureResponseGenerateToken
type AAAFailureResponseGenerateToken struct {

	// cause
	Cause *AAAFailureResponseGenerateTokenCause `json:"cause,omitempty"`

	// error
	Error string `json:"error,omitempty"`
}

// Validate validates this a a a failure response generate token
func (m *AAAFailureResponseGenerateToken) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCause(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AAAFailureResponseGenerateToken) validateCause(formats strfmt.Registry) error {
	if swag.IsZero(m.Cause) { // not required
		return nil
	}

	if m.Cause != nil {
		if err := m.Cause.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cause")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this a a a failure response generate token based on the context it is used
func (m *AAAFailureResponseGenerateToken) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCause(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AAAFailureResponseGenerateToken) contextValidateCause(ctx context.Context, formats strfmt.Registry) error {

	if m.Cause != nil {
		if err := m.Cause.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cause")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AAAFailureResponseGenerateToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AAAFailureResponseGenerateToken) UnmarshalBinary(b []byte) error {
	var res AAAFailureResponseGenerateToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
