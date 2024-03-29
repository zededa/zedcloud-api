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

// BulkInstanceConfig bulk instance config
//
// swagger:model BulkInstanceConfig
type BulkInstanceConfig struct {

	// app instance config
	AppInstanceConfig *AppInstance `json:"appInstanceConfig,omitempty"`

	// device Id
	DeviceID string `json:"deviceId,omitempty"`

	// net instance config
	NetInstanceConfig []*NetInstConfig `json:"netInstanceConfig"`

	// net instance detail
	NetInstanceDetail *NetworkConfigOrDefault `json:"netInstanceDetail,omitempty"`
}

// Validate validates this bulk instance config
func (m *BulkInstanceConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppInstanceConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetInstanceConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetInstanceDetail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulkInstanceConfig) validateAppInstanceConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.AppInstanceConfig) { // not required
		return nil
	}

	if m.AppInstanceConfig != nil {
		if err := m.AppInstanceConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appInstanceConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appInstanceConfig")
			}
			return err
		}
	}

	return nil
}

func (m *BulkInstanceConfig) validateNetInstanceConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.NetInstanceConfig) { // not required
		return nil
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

func (m *BulkInstanceConfig) validateNetInstanceDetail(formats strfmt.Registry) error {
	if swag.IsZero(m.NetInstanceDetail) { // not required
		return nil
	}

	if m.NetInstanceDetail != nil {
		if err := m.NetInstanceDetail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netInstanceDetail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("netInstanceDetail")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this bulk instance config based on the context it is used
func (m *BulkInstanceConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppInstanceConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetInstanceConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetInstanceDetail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulkInstanceConfig) contextValidateAppInstanceConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.AppInstanceConfig != nil {

		if swag.IsZero(m.AppInstanceConfig) { // not required
			return nil
		}

		if err := m.AppInstanceConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appInstanceConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appInstanceConfig")
			}
			return err
		}
	}

	return nil
}

func (m *BulkInstanceConfig) contextValidateNetInstanceConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BulkInstanceConfig) contextValidateNetInstanceDetail(ctx context.Context, formats strfmt.Registry) error {

	if m.NetInstanceDetail != nil {

		if swag.IsZero(m.NetInstanceDetail) { // not required
			return nil
		}

		if err := m.NetInstanceDetail.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netInstanceDetail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("netInstanceDetail")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BulkInstanceConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BulkInstanceConfig) UnmarshalBinary(b []byte) error {
	var res BulkInstanceConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
