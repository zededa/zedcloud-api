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

// ConfigWifiConfig config wifi config
//
// swagger:model configWifiConfig
type ConfigWifiConfig struct {

	// cipher data
	CipherData *ConfigCipherBlock `json:"cipherData,omitempty"`

	// crypto
	Crypto *WifiConfigcryptoblock `json:"crypto,omitempty"`

	// to be deprecated, use cipherData instead
	Identity string `json:"identity,omitempty"`

	// key scheme
	KeyScheme *ConfigWiFiKeyScheme `json:"keyScheme,omitempty"`

	// to be deprecated, use cipherData instead
	Password string `json:"password,omitempty"`

	// priority
	Priority int32 `json:"priority,omitempty"`

	// wifi s s ID
	WifiSSID string `json:"wifiSSID,omitempty"`
}

// Validate validates this config wifi config
func (m *ConfigWifiConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCipherData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCrypto(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyScheme(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigWifiConfig) validateCipherData(formats strfmt.Registry) error {
	if swag.IsZero(m.CipherData) { // not required
		return nil
	}

	if m.CipherData != nil {
		if err := m.CipherData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cipherData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cipherData")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigWifiConfig) validateCrypto(formats strfmt.Registry) error {
	if swag.IsZero(m.Crypto) { // not required
		return nil
	}

	if m.Crypto != nil {
		if err := m.Crypto.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("crypto")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("crypto")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigWifiConfig) validateKeyScheme(formats strfmt.Registry) error {
	if swag.IsZero(m.KeyScheme) { // not required
		return nil
	}

	if m.KeyScheme != nil {
		if err := m.KeyScheme.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keyScheme")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keyScheme")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config wifi config based on the context it is used
func (m *ConfigWifiConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCipherData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCrypto(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKeyScheme(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigWifiConfig) contextValidateCipherData(ctx context.Context, formats strfmt.Registry) error {

	if m.CipherData != nil {
		if err := m.CipherData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cipherData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cipherData")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigWifiConfig) contextValidateCrypto(ctx context.Context, formats strfmt.Registry) error {

	if m.Crypto != nil {
		if err := m.Crypto.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("crypto")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("crypto")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigWifiConfig) contextValidateKeyScheme(ctx context.Context, formats strfmt.Registry) error {

	if m.KeyScheme != nil {
		if err := m.KeyScheme.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keyScheme")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keyScheme")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigWifiConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigWifiConfig) UnmarshalBinary(b []byte) error {
	var res ConfigWifiConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
