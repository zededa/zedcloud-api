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

// DeviceSize device size
//
// swagger:model DeviceSize
type DeviceSize string

func NewDeviceSize(value DeviceSize) *DeviceSize {
	return &value
}

// Pointer returns a pointer to a freshly-allocated DeviceSize.
func (m DeviceSize) Pointer() *DeviceSize {
	return &m
}

const (

	// DeviceSizeDEFAULTSIZE captures enum value "DEFAULT_SIZE"
	DeviceSizeDEFAULTSIZE DeviceSize = "DEFAULT_SIZE"

	// DeviceSizeSMALLSIZE captures enum value "SMALL_SIZE"
	DeviceSizeSMALLSIZE DeviceSize = "SMALL_SIZE"

	// DeviceSizeMEDIUMSIZE captures enum value "MEDIUM_SIZE"
	DeviceSizeMEDIUMSIZE DeviceSize = "MEDIUM_SIZE"

	// DeviceSizeLARGESIZE captures enum value "LARGE_SIZE"
	DeviceSizeLARGESIZE DeviceSize = "LARGE_SIZE"

	// DeviceSizeXLARGESIZE captures enum value "XLARGE_SIZE"
	DeviceSizeXLARGESIZE DeviceSize = "XLARGE_SIZE"
)

// for schema
var deviceSizeEnum []interface{}

func init() {
	var res []DeviceSize
	if err := json.Unmarshal([]byte(`["DEFAULT_SIZE","SMALL_SIZE","MEDIUM_SIZE","LARGE_SIZE","XLARGE_SIZE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceSizeEnum = append(deviceSizeEnum, v)
	}
}

func (m DeviceSize) validateDeviceSizeEnum(path, location string, value DeviceSize) error {
	if err := validate.EnumCase(path, location, value, deviceSizeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this device size
func (m DeviceSize) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDeviceSizeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this device size based on context it is used
func (m DeviceSize) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
