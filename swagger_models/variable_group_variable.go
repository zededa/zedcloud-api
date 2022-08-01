// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

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
	"github.com/go-openapi/validate"
)

// VariableGroupVariable Edge Application detail
//
// Edge Application Artifact Details
// Example: {"format":"VARIABLE_FORMAT_TEXT","label":"Enter User Name","name":"username","required":"True"}
//
// swagger:model VariableGroupVariable
type VariableGroupVariable struct {

	// Default value of the variable. (Optional. Default: <Default value based on type>)
	Default string `json:"default,omitempty"`

	// Encoding of file content. Applicable if format is VARIABLE_FORMAT_FILE
	Encode *VariableFileEncoding `json:"encode,omitempty"`

	// Format of the user variable. (Required)
	// Required: true
	Format *VariableVariableFormat `json:"format"`

	// Label for the variable (Required)
	// Required: true
	Label *string `json:"label"`

	// Max length of the value of the variable(Optional. Default: 1024)
	MaxLength string `json:"maxLength,omitempty"`

	// Name of the Variable (Required)
	// Required: true
	Name *string `json:"name"`

	// Key-Value pair of options. Applicable if format is VARIABLE_FORMAT_DROPDOWN
	Options []*VariableOptionVal `json:"options"`

	// process input
	ProcessInput string `json:"processInput,omitempty"`

	// This variable MUST be specified when creating an App Instance. (Optional. Default: False)
	// Required: true
	Required *bool `json:"required"`

	// type
	Type string `json:"type,omitempty"`

	// User-specified value of the variable.(Required if required is true. Optional otherwise)
	Value string `json:"value,omitempty"`
}

// Validate validates this variable group variable
func (m *VariableGroupVariable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEncode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequired(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VariableGroupVariable) validateEncode(formats strfmt.Registry) error {
	if swag.IsZero(m.Encode) { // not required
		return nil
	}

	if m.Encode != nil {
		if err := m.Encode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("encode")
			}
			return err
		}
	}

	return nil
}

func (m *VariableGroupVariable) validateFormat(formats strfmt.Registry) error {

	if err := validate.Required("format", "body", m.Format); err != nil {
		return err
	}

	if err := validate.Required("format", "body", m.Format); err != nil {
		return err
	}

	if m.Format != nil {
		if err := m.Format.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("format")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("format")
			}
			return err
		}
	}

	return nil
}

func (m *VariableGroupVariable) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *VariableGroupVariable) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VariableGroupVariable) validateOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.Options) { // not required
		return nil
	}

	for i := 0; i < len(m.Options); i++ {
		if swag.IsZero(m.Options[i]) { // not required
			continue
		}

		if m.Options[i] != nil {
			if err := m.Options[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("options" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VariableGroupVariable) validateRequired(formats strfmt.Registry) error {

	if err := validate.Required("required", "body", m.Required); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this variable group variable based on the context it is used
func (m *VariableGroupVariable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEncode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFormat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VariableGroupVariable) contextValidateEncode(ctx context.Context, formats strfmt.Registry) error {

	if m.Encode != nil {
		if err := m.Encode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("encode")
			}
			return err
		}
	}

	return nil
}

func (m *VariableGroupVariable) contextValidateFormat(ctx context.Context, formats strfmt.Registry) error {

	if m.Format != nil {
		if err := m.Format.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("format")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("format")
			}
			return err
		}
	}

	return nil
}

func (m *VariableGroupVariable) contextValidateOptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Options); i++ {

		if m.Options[i] != nil {
			if err := m.Options[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("options" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VariableGroupVariable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VariableGroupVariable) UnmarshalBinary(b []byte) error {
	var res VariableGroupVariable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
