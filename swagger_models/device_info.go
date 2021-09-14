// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeviceInfo device info
//
// swagger:model DeviceInfo
type DeviceInfo struct {

	// cpu arch
	CPUArch string `json:"cpuArch,omitempty"`

	// machine arch
	MachineArch string `json:"machineArch,omitempty"`

	// mem m b
	MemMB string `json:"memMB,omitempty"`

	// n Cpu
	NCPU int64 `json:"nCpu,omitempty"`

	// platform
	Platform string `json:"platform,omitempty"`

	// storage m b
	StorageMB string `json:"storageMB,omitempty"`
}

// Validate validates this device info
func (m *DeviceInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this device info based on context it is used
func (m *DeviceInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceInfo) UnmarshalBinary(b []byte) error {
	var res DeviceInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
