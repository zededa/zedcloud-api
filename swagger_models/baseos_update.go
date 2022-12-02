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

// BaseosUpdate baseos update
//
// swagger:model BaseosUpdate
type BaseosUpdate struct {

	// baseimage
	Baseimage *BaseOSImage `json:"baseimage,omitempty"`
}

// Validate validates this baseos update
func (m *BaseosUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseimage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BaseosUpdate) validateBaseimage(formats strfmt.Registry) error {
	if swag.IsZero(m.Baseimage) { // not required
		return nil
	}

	if m.Baseimage != nil {
		if err := m.Baseimage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baseimage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("baseimage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this baseos update based on the context it is used
func (m *BaseosUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBaseimage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BaseosUpdate) contextValidateBaseimage(ctx context.Context, formats strfmt.Registry) error {

	if m.Baseimage != nil {
		if err := m.Baseimage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baseimage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("baseimage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BaseosUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaseosUpdate) UnmarshalBinary(b []byte) error {
	var res BaseosUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
