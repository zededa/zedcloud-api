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

// VolumeInstanceAccessMode volume instance access mode
//
// swagger:model VolumeInstanceAccessMode
type VolumeInstanceAccessMode string

func NewVolumeInstanceAccessMode(value VolumeInstanceAccessMode) *VolumeInstanceAccessMode {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VolumeInstanceAccessMode.
func (m VolumeInstanceAccessMode) Pointer() *VolumeInstanceAccessMode {
	return &m
}

const (

	// VolumeInstanceAccessModeVOLUMEINSTANCEACCESSMODEINVALID captures enum value "VOLUME_INSTANCE_ACCESS_MODE_INVALID"
	VolumeInstanceAccessModeVOLUMEINSTANCEACCESSMODEINVALID VolumeInstanceAccessMode = "VOLUME_INSTANCE_ACCESS_MODE_INVALID"

	// VolumeInstanceAccessModeVOLUMEINSTANCEACCESSMODEREADWRITE captures enum value "VOLUME_INSTANCE_ACCESS_MODE_READWRITE"
	VolumeInstanceAccessModeVOLUMEINSTANCEACCESSMODEREADWRITE VolumeInstanceAccessMode = "VOLUME_INSTANCE_ACCESS_MODE_READWRITE"

	// VolumeInstanceAccessModeVOLUMEINSTANCEACCESSMODEREADONLY captures enum value "VOLUME_INSTANCE_ACCESS_MODE_READONLY"
	VolumeInstanceAccessModeVOLUMEINSTANCEACCESSMODEREADONLY VolumeInstanceAccessMode = "VOLUME_INSTANCE_ACCESS_MODE_READONLY"

	// VolumeInstanceAccessModeVOLUMEINSTANCEACCESSMODEMULTIREADSINGLEWRITE captures enum value "VOLUME_INSTANCE_ACCESS_MODE_MULTIREAD_SINGLEWRITE"
	VolumeInstanceAccessModeVOLUMEINSTANCEACCESSMODEMULTIREADSINGLEWRITE VolumeInstanceAccessMode = "VOLUME_INSTANCE_ACCESS_MODE_MULTIREAD_SINGLEWRITE"
)

// for schema
var volumeInstanceAccessModeEnum []interface{}

func init() {
	var res []VolumeInstanceAccessMode
	if err := json.Unmarshal([]byte(`["VOLUME_INSTANCE_ACCESS_MODE_INVALID","VOLUME_INSTANCE_ACCESS_MODE_READWRITE","VOLUME_INSTANCE_ACCESS_MODE_READONLY","VOLUME_INSTANCE_ACCESS_MODE_MULTIREAD_SINGLEWRITE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		volumeInstanceAccessModeEnum = append(volumeInstanceAccessModeEnum, v)
	}
}

func (m VolumeInstanceAccessMode) validateVolumeInstanceAccessModeEnum(path, location string, value VolumeInstanceAccessMode) error {
	if err := validate.EnumCase(path, location, value, volumeInstanceAccessModeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this volume instance access mode
func (m VolumeInstanceAccessMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVolumeInstanceAccessModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this volume instance access mode based on context it is used
func (m VolumeInstanceAccessMode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
