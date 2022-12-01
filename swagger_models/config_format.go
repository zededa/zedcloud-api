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

// ConfigFormat config format
//
// swagger:model configFormat
type ConfigFormat string

func NewConfigFormat(value ConfigFormat) *ConfigFormat {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ConfigFormat.
func (m ConfigFormat) Pointer() *ConfigFormat {
	return &m
}

const (

	// ConfigFormatFmtUnknown captures enum value "FmtUnknown"
	ConfigFormatFmtUnknown ConfigFormat = "FmtUnknown"

	// ConfigFormatRAW captures enum value "RAW"
	ConfigFormatRAW ConfigFormat = "RAW"

	// ConfigFormatQCOW captures enum value "QCOW"
	ConfigFormatQCOW ConfigFormat = "QCOW"

	// ConfigFormatQCOW2 captures enum value "QCOW2"
	ConfigFormatQCOW2 ConfigFormat = "QCOW2"

	// ConfigFormatVHD captures enum value "VHD"
	ConfigFormatVHD ConfigFormat = "VHD"

	// ConfigFormatVMDK captures enum value "VMDK"
	ConfigFormatVMDK ConfigFormat = "VMDK"

	// ConfigFormatOVA captures enum value "OVA"
	ConfigFormatOVA ConfigFormat = "OVA"

	// ConfigFormatVHDX captures enum value "VHDX"
	ConfigFormatVHDX ConfigFormat = "VHDX"

	// ConfigFormatCONTAINER captures enum value "CONTAINER"
	ConfigFormatCONTAINER ConfigFormat = "CONTAINER"

	// ConfigFormatISO captures enum value "ISO"
	ConfigFormatISO ConfigFormat = "ISO"
)

// for schema
var configFormatEnum []interface{}

func init() {
	var res []ConfigFormat
	if err := json.Unmarshal([]byte(`["FmtUnknown","RAW","QCOW","QCOW2","VHD","VMDK","OVA","VHDX","CONTAINER","ISO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configFormatEnum = append(configFormatEnum, v)
	}
}

func (m ConfigFormat) validateConfigFormatEnum(path, location string, value ConfigFormat) error {
	if err := validate.EnumCase(path, location, value, configFormatEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this config format
func (m ConfigFormat) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateConfigFormatEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this config format based on context it is used
func (m ConfigFormat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
