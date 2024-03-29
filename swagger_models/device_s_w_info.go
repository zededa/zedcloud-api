// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
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

// DeviceSWInfo device s w info
//
// swagger:model DeviceSWInfo
type DeviceSWInfo struct {

	// activated
	Activated bool `json:"activated,omitempty"`

	// download progress
	DownloadProgress int64 `json:"downloadProgress,omitempty"`

	// long version
	LongVersion string `json:"longVersion,omitempty"`

	// partition device
	PartitionDevice string `json:"partitionDevice,omitempty"`

	// partition label
	PartitionLabel string `json:"partitionLabel,omitempty"`

	// partition state
	PartitionState string `json:"partitionState,omitempty"`

	// short version
	ShortVersion string `json:"shortVersion,omitempty"`

	// status
	Status *SWState `json:"status,omitempty"`

	// sub status progress
	SubStatusProgress int64 `json:"subStatusProgress,omitempty"`

	// sw error
	SwError *DeviceError `json:"swError,omitempty"`

	// sw status
	SwStatus *DeviceSWStatus `json:"swStatus,omitempty"`

	// sw sub status
	SwSubStatus *DeviceSWSubStatus `json:"swSubStatus,omitempty"`

	// sw sub status str
	SwSubStatusStr string `json:"swSubStatusStr,omitempty"`
}

// Validate validates this device s w info
func (m *DeviceSWInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwSubStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceSWInfo) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceSWInfo) validateSwError(formats strfmt.Registry) error {
	if swag.IsZero(m.SwError) { // not required
		return nil
	}

	if m.SwError != nil {
		if err := m.SwError.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("swError")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("swError")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceSWInfo) validateSwStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.SwStatus) { // not required
		return nil
	}

	if m.SwStatus != nil {
		if err := m.SwStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("swStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("swStatus")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceSWInfo) validateSwSubStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.SwSubStatus) { // not required
		return nil
	}

	if m.SwSubStatus != nil {
		if err := m.SwSubStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("swSubStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("swSubStatus")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this device s w info based on the context it is used
func (m *DeviceSWInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSwError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSwStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSwSubStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceSWInfo) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {

		if swag.IsZero(m.Status) { // not required
			return nil
		}

		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceSWInfo) contextValidateSwError(ctx context.Context, formats strfmt.Registry) error {

	if m.SwError != nil {

		if swag.IsZero(m.SwError) { // not required
			return nil
		}

		if err := m.SwError.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("swError")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("swError")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceSWInfo) contextValidateSwStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.SwStatus != nil {

		if swag.IsZero(m.SwStatus) { // not required
			return nil
		}

		if err := m.SwStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("swStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("swStatus")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceSWInfo) contextValidateSwSubStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.SwSubStatus != nil {

		if swag.IsZero(m.SwSubStatus) { // not required
			return nil
		}

		if err := m.SwSubStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("swSubStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("swSubStatus")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceSWInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceSWInfo) UnmarshalBinary(b []byte) error {
	var res DeviceSWInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
