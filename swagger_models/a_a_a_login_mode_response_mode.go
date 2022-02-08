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

// AAALoginModeResponseMode a a a login mode response mode
//
// swagger:model AAALoginModeResponseMode
type AAALoginModeResponseMode string

func NewAAALoginModeResponseMode(value AAALoginModeResponseMode) *AAALoginModeResponseMode {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AAALoginModeResponseMode.
func (m AAALoginModeResponseMode) Pointer() *AAALoginModeResponseMode {
	return &m
}

const (

	// AAALoginModeResponseModeAAALoginModeLocal captures enum value "AAALoginModeLocal"
	AAALoginModeResponseModeAAALoginModeLocal AAALoginModeResponseMode = "AAALoginModeLocal"

	// AAALoginModeResponseModeAAALoginModeExternal captures enum value "AAALoginModeExternal"
	AAALoginModeResponseModeAAALoginModeExternal AAALoginModeResponseMode = "AAALoginModeExternal"
)

// for schema
var aAALoginModeResponseModeEnum []interface{}

func init() {
	var res []AAALoginModeResponseMode
	if err := json.Unmarshal([]byte(`["AAALoginModeLocal","AAALoginModeExternal"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aAALoginModeResponseModeEnum = append(aAALoginModeResponseModeEnum, v)
	}
}

func (m AAALoginModeResponseMode) validateAAALoginModeResponseModeEnum(path, location string, value AAALoginModeResponseMode) error {
	if err := validate.EnumCase(path, location, value, aAALoginModeResponseModeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this a a a login mode response mode
func (m AAALoginModeResponseMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAAALoginModeResponseModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this a a a login mode response mode based on context it is used
func (m AAALoginModeResponseMode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
