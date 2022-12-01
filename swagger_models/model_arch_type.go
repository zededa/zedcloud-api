// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
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

// ModelArchType model arch type
//
// swagger:model ModelArchType
type ModelArchType string

func NewModelArchType(value ModelArchType) *ModelArchType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ModelArchType.
func (m ModelArchType) Pointer() *ModelArchType {
	return &m
}

const (

	// ModelArchTypeUNSPECIFIED captures enum value "UNSPECIFIED"
	ModelArchTypeUNSPECIFIED ModelArchType = "UNSPECIFIED"

	// ModelArchTypeUNDEFINED captures enum value "UNDEFINED"
	ModelArchTypeUNDEFINED ModelArchType = "UNDEFINED"

	// ModelArchTypeAMD64 captures enum value "AMD64"
	ModelArchTypeAMD64 ModelArchType = "AMD64"

	// ModelArchTypeARM64 captures enum value "ARM64"
	ModelArchTypeARM64 ModelArchType = "ARM64"
)

// for schema
var modelArchTypeEnum []interface{}

func init() {
	var res []ModelArchType
	if err := json.Unmarshal([]byte(`["UNSPECIFIED","UNDEFINED","AMD64","ARM64"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		modelArchTypeEnum = append(modelArchTypeEnum, v)
	}
}

func (m ModelArchType) validateModelArchTypeEnum(path, location string, value ModelArchType) error {
	if err := validate.EnumCase(path, location, value, modelArchTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this model arch type
func (m ModelArchType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateModelArchTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this model arch type based on context it is used
func (m ModelArchType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
