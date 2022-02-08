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

// ModuleType module type.
//
// swagger:model ModuleType
type ModuleType string

func NewModuleType(value ModuleType) *ModuleType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ModuleType.
func (m ModuleType) Pointer() *ModuleType {
	return &m
}

const (

	// ModuleTypeMODULETYPEUNSPECIFIED captures enum value "MODULE_TYPE_UNSPECIFIED"
	ModuleTypeMODULETYPEUNSPECIFIED ModuleType = "MODULE_TYPE_UNSPECIFIED"

	// ModuleTypeMODULETYPESYSTEMDEFINED captures enum value "MODULE_TYPE_SYSTEM_DEFINED"
	ModuleTypeMODULETYPESYSTEMDEFINED ModuleType = "MODULE_TYPE_SYSTEM_DEFINED"

	// ModuleTypeMODULETYPECUSTOM captures enum value "MODULE_TYPE_CUSTOM"
	ModuleTypeMODULETYPECUSTOM ModuleType = "MODULE_TYPE_CUSTOM"
)

// for schema
var moduleTypeEnum []interface{}

func init() {
	var res []ModuleType
	if err := json.Unmarshal([]byte(`["MODULE_TYPE_UNSPECIFIED","MODULE_TYPE_SYSTEM_DEFINED","MODULE_TYPE_CUSTOM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		moduleTypeEnum = append(moduleTypeEnum, v)
	}
}

func (m ModuleType) validateModuleTypeEnum(path, location string, value ModuleType) error {
	if err := validate.EnumCase(path, location, value, moduleTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this module type
func (m ModuleType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateModuleTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this module type based on context it is used
func (m ModuleType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
