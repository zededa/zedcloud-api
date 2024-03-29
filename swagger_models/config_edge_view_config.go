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

// ConfigEdgeViewConfig config edge view config
//
// swagger:model configEdgeViewConfig
type ConfigEdgeViewConfig struct {

	// policy access for apps through edge-view
	AppPolicy *ConfigAppDebugAccessPolicy `json:"appPolicy,omitempty"`

	// policy for device access through edge-view
	DevPolicy *ConfigDevDebugAccessPolicy `json:"devPolicy,omitempty"`

	// dispatcher certificate(s) if it's not well-known CA signed
	DispCertPem []strfmt.Base64 `json:"dispCertPem"`

	// policy access for external endpoint through edge-view
	ExtPolicy *ConfigExternalEndPointPolicy `json:"extPolicy,omitempty"`

	// Generation ID for re-start edgeview without parameter changes
	GenerationID int64 `json:"generationId,omitempty"`

	// JWT token for signed info, it contains the dispatcher
	// endpoint IP:Port, device UUID, nonce and expiration time
	Token string `json:"token,omitempty"`
}

// Validate validates this config edge view config
func (m *ConfigEdgeViewConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigEdgeViewConfig) validateAppPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.AppPolicy) { // not required
		return nil
	}

	if m.AppPolicy != nil {
		if err := m.AppPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigEdgeViewConfig) validateDevPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.DevPolicy) { // not required
		return nil
	}

	if m.DevPolicy != nil {
		if err := m.DevPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("devPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("devPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigEdgeViewConfig) validateExtPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.ExtPolicy) { // not required
		return nil
	}

	if m.ExtPolicy != nil {
		if err := m.ExtPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("extPolicy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config edge view config based on the context it is used
func (m *ConfigEdgeViewConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDevPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExtPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigEdgeViewConfig) contextValidateAppPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.AppPolicy != nil {

		if swag.IsZero(m.AppPolicy) { // not required
			return nil
		}

		if err := m.AppPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigEdgeViewConfig) contextValidateDevPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.DevPolicy != nil {

		if swag.IsZero(m.DevPolicy) { // not required
			return nil
		}

		if err := m.DevPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("devPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("devPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigEdgeViewConfig) contextValidateExtPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.ExtPolicy != nil {

		if swag.IsZero(m.ExtPolicy) { // not required
			return nil
		}

		if err := m.ExtPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("extPolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigEdgeViewConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigEdgeViewConfig) UnmarshalBinary(b []byte) error {
	var res ConfigEdgeViewConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
