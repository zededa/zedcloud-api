// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// VariableVariableFormat Custom variable format
//
// - VARIABLE_FORMAT_UNSPECIFIED: Invalid format
//  - VARIABLE_FORMAT_TEXT: Value in string format
//  - VARIABLE_FORMAT_NUMBER: Value in integer format
//  - VARIABLE_FORMAT_FILE: Value to be read from a file
//  - VARIABLE_FORMAT_DROPDOWN: Value to be selected from dropdown of options
//  - VARIABLE_FORMAT_BOOLEAN: Value in boolean format
//  - VARIABLE_FORMAT_PASSWORD: Value in string to be masked in User-Agent
//
// swagger:model VariableVariableFormat
type VariableVariableFormat string

func NewVariableVariableFormat(value VariableVariableFormat) *VariableVariableFormat {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VariableVariableFormat.
func (m VariableVariableFormat) Pointer() *VariableVariableFormat {
	return &m
}

const (

	// VariableVariableFormatVARIABLEFORMATUNSPECIFIED captures enum value "VARIABLE_FORMAT_UNSPECIFIED"
	VariableVariableFormatVARIABLEFORMATUNSPECIFIED VariableVariableFormat = "VARIABLE_FORMAT_UNSPECIFIED"

	// VariableVariableFormatVARIABLEFORMATTEXT captures enum value "VARIABLE_FORMAT_TEXT"
	VariableVariableFormatVARIABLEFORMATTEXT VariableVariableFormat = "VARIABLE_FORMAT_TEXT"

	// VariableVariableFormatVARIABLEFORMATNUMBER captures enum value "VARIABLE_FORMAT_NUMBER"
	VariableVariableFormatVARIABLEFORMATNUMBER VariableVariableFormat = "VARIABLE_FORMAT_NUMBER"

	// VariableVariableFormatVARIABLEFORMATFILE captures enum value "VARIABLE_FORMAT_FILE"
	VariableVariableFormatVARIABLEFORMATFILE VariableVariableFormat = "VARIABLE_FORMAT_FILE"

	// VariableVariableFormatVARIABLEFORMATDROPDOWN captures enum value "VARIABLE_FORMAT_DROPDOWN"
	VariableVariableFormatVARIABLEFORMATDROPDOWN VariableVariableFormat = "VARIABLE_FORMAT_DROPDOWN"

	// VariableVariableFormatVARIABLEFORMATBOOLEAN captures enum value "VARIABLE_FORMAT_BOOLEAN"
	VariableVariableFormatVARIABLEFORMATBOOLEAN VariableVariableFormat = "VARIABLE_FORMAT_BOOLEAN"

	// VariableVariableFormatVARIABLEFORMATPASSWORD captures enum value "VARIABLE_FORMAT_PASSWORD"
	VariableVariableFormatVARIABLEFORMATPASSWORD VariableVariableFormat = "VARIABLE_FORMAT_PASSWORD"
)

// for schema
var variableVariableFormatEnum []interface{}

func init() {
	var res []VariableVariableFormat
	if err := json.Unmarshal([]byte(`["VARIABLE_FORMAT_UNSPECIFIED","VARIABLE_FORMAT_TEXT","VARIABLE_FORMAT_NUMBER","VARIABLE_FORMAT_FILE","VARIABLE_FORMAT_DROPDOWN","VARIABLE_FORMAT_BOOLEAN","VARIABLE_FORMAT_PASSWORD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		variableVariableFormatEnum = append(variableVariableFormatEnum, v)
	}
}

func (m VariableVariableFormat) validateVariableVariableFormatEnum(path, location string, value VariableVariableFormat) error {
	if err := validate.EnumCase(path, location, value, variableVariableFormatEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this variable variable format
func (m VariableVariableFormat) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVariableVariableFormatEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this variable variable format based on context it is used
func (m VariableVariableFormat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
