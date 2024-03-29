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
)

// AAAFailureResponseLogout a a a failure response logout
//
// swagger:model AAAFailureResponseLogout
type AAAFailureResponseLogout struct {

	// cause
	Cause *AAAFailureResponseLogoutCause `json:"cause,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// original
	Original *OpaqueToken64 `json:"original,omitempty"`
}

// Validate validates this a a a failure response logout
func (m *AAAFailureResponseLogout) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCause(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AAAFailureResponseLogout) validateCause(formats strfmt.Registry) error {
	if swag.IsZero(m.Cause) { // not required
		return nil
	}

	if m.Cause != nil {
		if err := m.Cause.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cause")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cause")
			}
			return err
		}
	}

	return nil
}

func (m *AAAFailureResponseLogout) validateOriginal(formats strfmt.Registry) error {
	if swag.IsZero(m.Original) { // not required
		return nil
	}

	if m.Original != nil {
		if err := m.Original.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("original")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("original")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this a a a failure response logout based on the context it is used
func (m *AAAFailureResponseLogout) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCause(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOriginal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AAAFailureResponseLogout) contextValidateCause(ctx context.Context, formats strfmt.Registry) error {

	if m.Cause != nil {

		if swag.IsZero(m.Cause) { // not required
			return nil
		}

		if err := m.Cause.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cause")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cause")
			}
			return err
		}
	}

	return nil
}

func (m *AAAFailureResponseLogout) contextValidateOriginal(ctx context.Context, formats strfmt.Registry) error {

	if m.Original != nil {

		if swag.IsZero(m.Original) { // not required
			return nil
		}

		if err := m.Original.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("original")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("original")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AAAFailureResponseLogout) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AAAFailureResponseLogout) UnmarshalBinary(b []byte) error {
	var res AAAFailureResponseLogout
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
