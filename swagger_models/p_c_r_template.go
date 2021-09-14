// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

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

// PCRTemplate p c r template
//
// swagger:model PCRTemplate
type PCRTemplate struct {

	// p c r values
	PCRValues []*PCRValue `json:"PCRValues"`

	// eve version
	EveVersion string `json:"eveVersion,omitempty"`

	// firmware version
	FirmwareVersion string `json:"firmwareVersion,omitempty"`
}

// Validate validates this p c r template
func (m *PCRTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePCRValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PCRTemplate) validatePCRValues(formats strfmt.Registry) error {
	if swag.IsZero(m.PCRValues) { // not required
		return nil
	}

	for i := 0; i < len(m.PCRValues); i++ {
		if swag.IsZero(m.PCRValues[i]) { // not required
			continue
		}

		if m.PCRValues[i] != nil {
			if err := m.PCRValues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PCRValues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this p c r template based on the context it is used
func (m *PCRTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePCRValues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PCRTemplate) contextValidatePCRValues(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PCRValues); i++ {

		if m.PCRValues[i] != nil {
			if err := m.PCRValues[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PCRValues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PCRTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PCRTemplate) UnmarshalBinary(b []byte) error {
	var res PCRTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
