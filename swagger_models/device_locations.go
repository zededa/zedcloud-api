// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeviceLocations device locations
//
// swagger:model DeviceLocations
type DeviceLocations struct {

	// device locations
	DeviceLocations []*DeviceLocation `json:"deviceLocations"`

	// next
	Next *Cursor `json:"next,omitempty"`
}

// Validate validates this device locations
func (m *DeviceLocations) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceLocations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceLocations) validateDeviceLocations(formats strfmt.Registry) error {
	if swag.IsZero(m.DeviceLocations) { // not required
		return nil
	}

	for i := 0; i < len(m.DeviceLocations); i++ {
		if swag.IsZero(m.DeviceLocations[i]) { // not required
			continue
		}

		if m.DeviceLocations[i] != nil {
			if err := m.DeviceLocations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deviceLocations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("deviceLocations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceLocations) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this device locations based on the context it is used
func (m *DeviceLocations) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeviceLocations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceLocations) contextValidateDeviceLocations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DeviceLocations); i++ {

		if m.DeviceLocations[i] != nil {

			if swag.IsZero(m.DeviceLocations[i]) { // not required
				return nil
			}

			if err := m.DeviceLocations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deviceLocations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("deviceLocations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceLocations) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if m.Next != nil {

		if swag.IsZero(m.Next) { // not required
			return nil
		}

		if err := m.Next.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceLocations) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceLocations) UnmarshalBinary(b []byte) error {
	var res DeviceLocations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
