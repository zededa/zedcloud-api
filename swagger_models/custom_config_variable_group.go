// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustomConfigVariableGroup custom config variable group
//
// swagger:model CustomConfigVariableGroup
type CustomConfigVariableGroup struct {

	// Condition to apply the variable group. (Optional. Default: None)
	Condition *VariableGroupCondition `json:"condition,omitempty"`

	// Name of the Variable Group(Required)
	Name string `json:"name,omitempty"`

	// Indicates if the variable group is required to be specified for the App Instance. (Optional. Default:False)
	Required bool `json:"required,omitempty"`

	// List of variables(Required)
	Variables []*VariableGroupVariable `json:"variables"`
}

// Validate validates this custom config variable group
func (m *CustomConfigVariableGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCondition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomConfigVariableGroup) validateCondition(formats strfmt.Registry) error {
	if swag.IsZero(m.Condition) { // not required
		return nil
	}

	if m.Condition != nil {
		if err := m.Condition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("condition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("condition")
			}
			return err
		}
	}

	return nil
}

func (m *CustomConfigVariableGroup) validateVariables(formats strfmt.Registry) error {
	if swag.IsZero(m.Variables) { // not required
		return nil
	}

	for i := 0; i < len(m.Variables); i++ {
		if swag.IsZero(m.Variables[i]) { // not required
			continue
		}

		if m.Variables[i] != nil {
			if err := m.Variables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("variables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this custom config variable group based on the context it is used
func (m *CustomConfigVariableGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCondition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVariables(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomConfigVariableGroup) contextValidateCondition(ctx context.Context, formats strfmt.Registry) error {

	if m.Condition != nil {

		if swag.IsZero(m.Condition) { // not required
			return nil
		}

		if err := m.Condition.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("condition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("condition")
			}
			return err
		}
	}

	return nil
}

func (m *CustomConfigVariableGroup) contextValidateVariables(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Variables); i++ {

		if m.Variables[i] != nil {

			if swag.IsZero(m.Variables[i]) { // not required
				return nil
			}

			if err := m.Variables[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("variables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomConfigVariableGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomConfigVariableGroup) UnmarshalBinary(b []byte) error {
	var res CustomConfigVariableGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
