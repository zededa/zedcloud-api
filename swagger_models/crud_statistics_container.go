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

// CrudStatisticsContainer Container for per-clazz statistics
//
// swagger:model CrudStatisticsContainer
type CrudStatisticsContainer struct {

	// clazz
	Clazz *ModelClazz `json:"clazz,omitempty"`

	// user
	User *CrudStatisticsUser `json:"user,omitempty"`
}

// Validate validates this crud statistics container
func (m *CrudStatisticsContainer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClazz(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CrudStatisticsContainer) validateClazz(formats strfmt.Registry) error {
	if swag.IsZero(m.Clazz) { // not required
		return nil
	}

	if m.Clazz != nil {
		if err := m.Clazz.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clazz")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clazz")
			}
			return err
		}
	}

	return nil
}

func (m *CrudStatisticsContainer) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this crud statistics container based on the context it is used
func (m *CrudStatisticsContainer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClazz(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CrudStatisticsContainer) contextValidateClazz(ctx context.Context, formats strfmt.Registry) error {

	if m.Clazz != nil {
		if err := m.Clazz.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clazz")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clazz")
			}
			return err
		}
	}

	return nil
}

func (m *CrudStatisticsContainer) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {
		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CrudStatisticsContainer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CrudStatisticsContainer) UnmarshalBinary(b []byte) error {
	var res CrudStatisticsContainer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
