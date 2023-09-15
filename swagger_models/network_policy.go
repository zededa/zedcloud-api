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
	"github.com/go-openapi/validate"
)

// NetworkPolicy Network policy body detail
//
// list of networks that will be created on all the devices of the project to which this policy is attached
//
// swagger:model NetworkPolicy
type NetworkPolicy struct {

	// list of network details that will be created on all the devices of the project to which this policy is attached
	// Required: true
	NetInstanceConfig []*NetInstConfig `json:"netInstanceConfig"`
}

// Validate validates this network policy
func (m *NetworkPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetInstanceConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkPolicy) validateNetInstanceConfig(formats strfmt.Registry) error {

	if err := validate.Required("netInstanceConfig", "body", m.NetInstanceConfig); err != nil {
		return err
	}

	for i := 0; i < len(m.NetInstanceConfig); i++ {
		if swag.IsZero(m.NetInstanceConfig[i]) { // not required
			continue
		}

		if m.NetInstanceConfig[i] != nil {
			if err := m.NetInstanceConfig[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("netInstanceConfig" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("netInstanceConfig" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this network policy based on the context it is used
func (m *NetworkPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetInstanceConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkPolicy) contextValidateNetInstanceConfig(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NetInstanceConfig); i++ {

		if m.NetInstanceConfig[i] != nil {

			if swag.IsZero(m.NetInstanceConfig[i]) { // not required
				return nil
			}

			if err := m.NetInstanceConfig[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("netInstanceConfig" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("netInstanceConfig" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkPolicy) UnmarshalBinary(b []byte) error {
	var res NetworkPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
