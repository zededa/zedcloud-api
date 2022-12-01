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

// DeviceSWStatus device s w status
//
// swagger:model DeviceSWStatus
type DeviceSWStatus string

func NewDeviceSWStatus(value DeviceSWStatus) *DeviceSWStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated DeviceSWStatus.
func (m DeviceSWStatus) Pointer() *DeviceSWStatus {
	return &m
}

const (

	// DeviceSWStatusDEVICESWSTATUSUNSPECIFIED captures enum value "DEVICE_SW_STATUS_UNSPECIFIED"
	DeviceSWStatusDEVICESWSTATUSUNSPECIFIED DeviceSWStatus = "DEVICE_SW_STATUS_UNSPECIFIED"

	// DeviceSWStatusDEVICESWSTATUSDOWNLOADING captures enum value "DEVICE_SW_STATUS_DOWNLOADING"
	DeviceSWStatusDEVICESWSTATUSDOWNLOADING DeviceSWStatus = "DEVICE_SW_STATUS_DOWNLOADING"

	// DeviceSWStatusDEVICESWSTATUSDOWNLOADDONE captures enum value "DEVICE_SW_STATUS_DOWNLOADDONE"
	DeviceSWStatusDEVICESWSTATUSDOWNLOADDONE DeviceSWStatus = "DEVICE_SW_STATUS_DOWNLOADDONE"

	// DeviceSWStatusDEVICESWSTATUSUPDATING captures enum value "DEVICE_SW_STATUS_UPDATING"
	DeviceSWStatusDEVICESWSTATUSUPDATING DeviceSWStatus = "DEVICE_SW_STATUS_UPDATING"

	// DeviceSWStatusDEVICESWSTATUSUPDATED captures enum value "DEVICE_SW_STATUS_UPDATED"
	DeviceSWStatusDEVICESWSTATUSUPDATED DeviceSWStatus = "DEVICE_SW_STATUS_UPDATED"

	// DeviceSWStatusDEVICESWSTATUSFALLBACK captures enum value "DEVICE_SW_STATUS_FALLBACK"
	DeviceSWStatusDEVICESWSTATUSFALLBACK DeviceSWStatus = "DEVICE_SW_STATUS_FALLBACK"

	// DeviceSWStatusDEVICESWSTATUSFAILED captures enum value "DEVICE_SW_STATUS_FAILED"
	DeviceSWStatusDEVICESWSTATUSFAILED DeviceSWStatus = "DEVICE_SW_STATUS_FAILED"
)

// for schema
var deviceSWStatusEnum []interface{}

func init() {
	var res []DeviceSWStatus
	if err := json.Unmarshal([]byte(`["DEVICE_SW_STATUS_UNSPECIFIED","DEVICE_SW_STATUS_DOWNLOADING","DEVICE_SW_STATUS_DOWNLOADDONE","DEVICE_SW_STATUS_UPDATING","DEVICE_SW_STATUS_UPDATED","DEVICE_SW_STATUS_FALLBACK","DEVICE_SW_STATUS_FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceSWStatusEnum = append(deviceSWStatusEnum, v)
	}
}

func (m DeviceSWStatus) validateDeviceSWStatusEnum(path, location string, value DeviceSWStatus) error {
	if err := validate.EnumCase(path, location, value, deviceSWStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this device s w status
func (m DeviceSWStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDeviceSWStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this device s w status based on context it is used
func (m DeviceSWStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
