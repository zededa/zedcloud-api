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

// DeviceHWSecurityModuleStatus device h w security module status
//
// swagger:model DeviceHWSecurityModuleStatus
type DeviceHWSecurityModuleStatus string

func NewDeviceHWSecurityModuleStatus(value DeviceHWSecurityModuleStatus) *DeviceHWSecurityModuleStatus {
	v := value
	return &v
}

const (

	// DeviceHWSecurityModuleStatusDEVICEHWSECURITYMODULESTATUSUNSPECIFIED captures enum value "DEVICE_HW_SECURITY_MODULE_STATUS_UNSPECIFIED"
	DeviceHWSecurityModuleStatusDEVICEHWSECURITYMODULESTATUSUNSPECIFIED DeviceHWSecurityModuleStatus = "DEVICE_HW_SECURITY_MODULE_STATUS_UNSPECIFIED"

	// DeviceHWSecurityModuleStatusDEVICEHWSECURITYMODULESTATUSNOTFOUND captures enum value "DEVICE_HW_SECURITY_MODULE_STATUS_NOT_FOUND"
	DeviceHWSecurityModuleStatusDEVICEHWSECURITYMODULESTATUSNOTFOUND DeviceHWSecurityModuleStatus = "DEVICE_HW_SECURITY_MODULE_STATUS_NOT_FOUND"

	// DeviceHWSecurityModuleStatusDEVICEHWSECURITYMODULESTATUSDISABLED captures enum value "DEVICE_HW_SECURITY_MODULE_STATUS_DISABLED"
	DeviceHWSecurityModuleStatusDEVICEHWSECURITYMODULESTATUSDISABLED DeviceHWSecurityModuleStatus = "DEVICE_HW_SECURITY_MODULE_STATUS_DISABLED"

	// DeviceHWSecurityModuleStatusDEVICEHWSECURITYMODULESTATUSENABLED captures enum value "DEVICE_HW_SECURITY_MODULE_STATUS_ENABLED"
	DeviceHWSecurityModuleStatusDEVICEHWSECURITYMODULESTATUSENABLED DeviceHWSecurityModuleStatus = "DEVICE_HW_SECURITY_MODULE_STATUS_ENABLED"
)

// for schema
var deviceHWSecurityModuleStatusEnum []interface{}

func init() {
	var res []DeviceHWSecurityModuleStatus
	if err := json.Unmarshal([]byte(`["DEVICE_HW_SECURITY_MODULE_STATUS_UNSPECIFIED","DEVICE_HW_SECURITY_MODULE_STATUS_NOT_FOUND","DEVICE_HW_SECURITY_MODULE_STATUS_DISABLED","DEVICE_HW_SECURITY_MODULE_STATUS_ENABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceHWSecurityModuleStatusEnum = append(deviceHWSecurityModuleStatusEnum, v)
	}
}

func (m DeviceHWSecurityModuleStatus) validateDeviceHWSecurityModuleStatusEnum(path, location string, value DeviceHWSecurityModuleStatus) error {
	if err := validate.EnumCase(path, location, value, deviceHWSecurityModuleStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this device h w security module status
func (m DeviceHWSecurityModuleStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDeviceHWSecurityModuleStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this device h w security module status based on context it is used
func (m DeviceHWSecurityModuleStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
