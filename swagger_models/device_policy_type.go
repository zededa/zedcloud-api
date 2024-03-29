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

// DevicePolicyType device policy type
//
// swagger:model DevicePolicyType
type DevicePolicyType string

func NewDevicePolicyType(value DevicePolicyType) *DevicePolicyType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated DevicePolicyType.
func (m DevicePolicyType) Pointer() *DevicePolicyType {
	return &m
}

const (

	// DevicePolicyTypeDEVICEPOLICYTYPEUNSPECIFIED captures enum value "DEVICE_POLICY_TYPE_UNSPECIFIED"
	DevicePolicyTypeDEVICEPOLICYTYPEUNSPECIFIED DevicePolicyType = "DEVICE_POLICY_TYPE_UNSPECIFIED"

	// DevicePolicyTypeDEVICEPOLICYTYPEATTESTATION captures enum value "DEVICE_POLICY_TYPE_ATTESTATION"
	DevicePolicyTypeDEVICEPOLICYTYPEATTESTATION DevicePolicyType = "DEVICE_POLICY_TYPE_ATTESTATION"
)

// for schema
var devicePolicyTypeEnum []interface{}

func init() {
	var res []DevicePolicyType
	if err := json.Unmarshal([]byte(`["DEVICE_POLICY_TYPE_UNSPECIFIED","DEVICE_POLICY_TYPE_ATTESTATION"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		devicePolicyTypeEnum = append(devicePolicyTypeEnum, v)
	}
}

func (m DevicePolicyType) validateDevicePolicyTypeEnum(path, location string, value DevicePolicyType) error {
	if err := validate.EnumCase(path, location, value, devicePolicyTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this device policy type
func (m DevicePolicyType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDevicePolicyTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this device policy type based on context it is used
func (m DevicePolicyType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
