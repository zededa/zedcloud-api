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

// ConfigTarget config target
//
// swagger:model configTarget
type ConfigTarget string

func NewConfigTarget(value ConfigTarget) *ConfigTarget {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ConfigTarget.
func (m ConfigTarget) Pointer() *ConfigTarget {
	return &m
}

const (

	// ConfigTargetTgtUnknown captures enum value "TgtUnknown"
	ConfigTargetTgtUnknown ConfigTarget = "TgtUnknown"

	// ConfigTargetDisk captures enum value "Disk"
	ConfigTargetDisk ConfigTarget = "Disk"

	// ConfigTargetKernel captures enum value "Kernel"
	ConfigTargetKernel ConfigTarget = "Kernel"

	// ConfigTargetInitrd captures enum value "Initrd"
	ConfigTargetInitrd ConfigTarget = "Initrd"

	// ConfigTargetRAMDisk captures enum value "RamDisk"
	ConfigTargetRAMDisk ConfigTarget = "RamDisk"
)

// for schema
var configTargetEnum []interface{}

func init() {
	var res []ConfigTarget
	if err := json.Unmarshal([]byte(`["TgtUnknown","Disk","Kernel","Initrd","RamDisk"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configTargetEnum = append(configTargetEnum, v)
	}
}

func (m ConfigTarget) validateConfigTargetEnum(path, location string, value ConfigTarget) error {
	if err := validate.EnumCase(path, location, value, configTargetEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this config target
func (m ConfigTarget) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateConfigTargetEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this config target based on context it is used
func (m ConfigTarget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
