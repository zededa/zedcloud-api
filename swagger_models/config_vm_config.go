// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConfigVMConfig config Vm config
//
// swagger:model configVmConfig
type ConfigVMConfig struct {

	// bootloader
	Bootloader string `json:"bootloader,omitempty"`

	// cpus
	Cpus string `json:"cpus,omitempty"`

	// devicetree
	Devicetree string `json:"devicetree,omitempty"`

	// disable logs
	DisableLogs bool `json:"disableLogs,omitempty"`

	// dtdev
	Dtdev []string `json:"dtdev"`

	// enable vnc
	EnableVnc bool `json:"enableVnc,omitempty"`

	// extraargs
	Extraargs string `json:"extraargs,omitempty"`

	// iomem
	Iomem []string `json:"iomem"`

	// irqs
	Irqs []int64 `json:"irqs"`

	// kernel
	Kernel string `json:"kernel,omitempty"`

	// maxcpus
	Maxcpus int64 `json:"maxcpus,omitempty"`

	// maxmem
	Maxmem int64 `json:"maxmem,omitempty"`

	// memory
	Memory int64 `json:"memory,omitempty"`

	// ramdisk
	Ramdisk string `json:"ramdisk,omitempty"`

	// rootdev
	Rootdev string `json:"rootdev,omitempty"`

	// vcpus
	Vcpus int64 `json:"vcpus,omitempty"`

	// virtualization mode
	VirtualizationMode *ConfigVMMode `json:"virtualizationMode,omitempty"`

	// vnc display
	VncDisplay int64 `json:"vncDisplay,omitempty"`

	// vnc passwd
	VncPasswd string `json:"vncPasswd,omitempty"`
}

// Validate validates this config Vm config
func (m *ConfigVMConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVirtualizationMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigVMConfig) validateVirtualizationMode(formats strfmt.Registry) error {
	if swag.IsZero(m.VirtualizationMode) { // not required
		return nil
	}

	if m.VirtualizationMode != nil {
		if err := m.VirtualizationMode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualizationMode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("virtualizationMode")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config Vm config based on the context it is used
func (m *ConfigVMConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVirtualizationMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigVMConfig) contextValidateVirtualizationMode(ctx context.Context, formats strfmt.Registry) error {

	if m.VirtualizationMode != nil {
		if err := m.VirtualizationMode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualizationMode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("virtualizationMode")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigVMConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigVMConfig) UnmarshalBinary(b []byte) error {
	var res ConfigVMConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
