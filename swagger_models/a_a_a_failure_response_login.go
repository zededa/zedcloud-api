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

// AAAFailureResponseLogin a a a failure response login
//
// swagger:model AAAFailureResponseLogin
type AAAFailureResponseLogin struct {

	// cause
	Cause *AAAFailureResponseLoginCause `json:"cause,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// no of login attempts left
	NoOfLoginAttemptsLeft int64 `json:"noOfLoginAttemptsLeft,omitempty"`

	// Sessions depend heavily on AAASuccessResponseLogin. In case of password expired,
	// we need temporary token. We can not generate a temporary token for password reset with
	// AAAFailureResponseLogin itself. Therefore, adding this tempSuccessResponse, to be used to create new session.
	TempSuccessResponse *AAASuccessResponseLogin `json:"tempSuccessResponse,omitempty"`

	// temp token
	TempToken *Token64 `json:"tempToken,omitempty"`
}

// Validate validates this a a a failure response login
func (m *AAAFailureResponseLogin) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCause(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTempSuccessResponse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTempToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AAAFailureResponseLogin) validateCause(formats strfmt.Registry) error {
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

func (m *AAAFailureResponseLogin) validateTempSuccessResponse(formats strfmt.Registry) error {
	if swag.IsZero(m.TempSuccessResponse) { // not required
		return nil
	}

	if m.TempSuccessResponse != nil {
		if err := m.TempSuccessResponse.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tempSuccessResponse")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tempSuccessResponse")
			}
			return err
		}
	}

	return nil
}

func (m *AAAFailureResponseLogin) validateTempToken(formats strfmt.Registry) error {
	if swag.IsZero(m.TempToken) { // not required
		return nil
	}

	if m.TempToken != nil {
		if err := m.TempToken.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tempToken")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tempToken")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this a a a failure response login based on the context it is used
func (m *AAAFailureResponseLogin) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCause(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTempSuccessResponse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTempToken(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AAAFailureResponseLogin) contextValidateCause(ctx context.Context, formats strfmt.Registry) error {

	if m.Cause != nil {
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

func (m *AAAFailureResponseLogin) contextValidateTempSuccessResponse(ctx context.Context, formats strfmt.Registry) error {

	if m.TempSuccessResponse != nil {
		if err := m.TempSuccessResponse.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tempSuccessResponse")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tempSuccessResponse")
			}
			return err
		}
	}

	return nil
}

func (m *AAAFailureResponseLogin) contextValidateTempToken(ctx context.Context, formats strfmt.Registry) error {

	if m.TempToken != nil {
		if err := m.TempToken.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tempToken")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tempToken")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AAAFailureResponseLogin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AAAFailureResponseLogin) UnmarshalBinary(b []byte) error {
	var res AAAFailureResponseLogin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
