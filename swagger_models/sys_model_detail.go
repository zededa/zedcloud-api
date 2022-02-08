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

// SysModelDetail sys model detail
//
// swagger:model SysModelDetail
type SysModelDetail struct {

	// custom model fields
	CustomModelFields *CustomUpdateModelFields `json:"customModelFields,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// parent model Id
	ParentModelID string `json:"parentModelId,omitempty"`
}

// Validate validates this sys model detail
func (m *SysModelDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomModelFields(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SysModelDetail) validateCustomModelFields(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomModelFields) { // not required
		return nil
	}

	if m.CustomModelFields != nil {
		if err := m.CustomModelFields.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customModelFields")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customModelFields")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sys model detail based on the context it is used
func (m *SysModelDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomModelFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SysModelDetail) contextValidateCustomModelFields(ctx context.Context, formats strfmt.Registry) error {

	if m.CustomModelFields != nil {
		if err := m.CustomModelFields.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customModelFields")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customModelFields")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SysModelDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SysModelDetail) UnmarshalBinary(b []byte) error {
	var res SysModelDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
