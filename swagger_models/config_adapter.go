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

// ConfigAdapter Adapter bundles corresponding to a subset of what is in ZioBundle
// When used by a NetworkInstanceConfig the name is the logicallabel
// for the network adapter.
//
// swagger:model configAdapter
type ConfigAdapter struct {

	// eth vf
	EthVf *ConfigEthVF `json:"ethVf,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// type
	Type *CommonPhyIoType `json:"type,omitempty"`
}

// Validate validates this config adapter
func (m *ConfigAdapter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEthVf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigAdapter) validateEthVf(formats strfmt.Registry) error {
	if swag.IsZero(m.EthVf) { // not required
		return nil
	}

	if m.EthVf != nil {
		if err := m.EthVf.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ethVf")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ethVf")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigAdapter) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config adapter based on the context it is used
func (m *ConfigAdapter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEthVf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigAdapter) contextValidateEthVf(ctx context.Context, formats strfmt.Registry) error {

	if m.EthVf != nil {

		if swag.IsZero(m.EthVf) { // not required
			return nil
		}

		if err := m.EthVf.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ethVf")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ethVf")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigAdapter) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {

		if swag.IsZero(m.Type) { // not required
			return nil
		}

		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigAdapter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigAdapter) UnmarshalBinary(b []byte) error {
	var res ConfigAdapter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
