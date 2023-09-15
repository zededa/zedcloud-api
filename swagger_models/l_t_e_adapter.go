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

// LTEAdapter LTE Adapter
//
// # Details of LTE Adapter
//
// swagger:model LTEAdapter
type LTEAdapter struct {

	// Name of Cell Module
	CellModuleName string `json:"cellModuleName,omitempty"`

	// Firmware Version of Cell Radio.
	FirmwareVersion string `json:"firmwareVersion,omitempty"`

	// iccid of the SIM
	Iccid string `json:"iccid,omitempty"`

	// IMEI of Cell Radio.
	Imei string `json:"imei,omitempty"`

	// imsi of the SIM
	Imsi string `json:"imsi,omitempty"`

	// Name of SIM card.
	SimName string `json:"simName,omitempty"`

	// State of SimCard
	SimcardState *SimcardState `json:"simcardState,omitempty"`
}

// Validate validates this l t e adapter
func (m *LTEAdapter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSimcardState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LTEAdapter) validateSimcardState(formats strfmt.Registry) error {
	if swag.IsZero(m.SimcardState) { // not required
		return nil
	}

	if m.SimcardState != nil {
		if err := m.SimcardState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("simcardState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("simcardState")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this l t e adapter based on the context it is used
func (m *LTEAdapter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSimcardState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LTEAdapter) contextValidateSimcardState(ctx context.Context, formats strfmt.Registry) error {

	if m.SimcardState != nil {

		if swag.IsZero(m.SimcardState) { // not required
			return nil
		}

		if err := m.SimcardState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("simcardState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("simcardState")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LTEAdapter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LTEAdapter) UnmarshalBinary(b []byte) error {
	var res LTEAdapter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
