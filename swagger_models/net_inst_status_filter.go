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

// NetInstStatusFilter net inst status filter
//
// swagger:model NetInstStatusFilter
type NetInstStatusFilter struct {

	// device name
	DeviceName string `json:"deviceName,omitempty"`

	// device name pattern
	DeviceNamePattern string `json:"deviceNamePattern,omitempty"`

	// kind
	Kind *NetworkInstanceKind `json:"kind,omitempty"`

	// name pattern
	NamePattern string `json:"namePattern,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`

	// project name pattern
	ProjectNamePattern string `json:"projectNamePattern,omitempty"`

	// run state
	RunState *RunState `json:"runState,omitempty"`

	// tags
	Tags map[string]string `json:"tags,omitempty"`

	// type
	Type *NetworkInstanceDhcpType `json:"type,omitempty"`
}

// Validate validates this net inst status filter
func (m *NetInstStatusFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunState(formats); err != nil {
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

func (m *NetInstStatusFilter) validateKind(formats strfmt.Registry) error {
	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	if m.Kind != nil {
		if err := m.Kind.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kind")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kind")
			}
			return err
		}
	}

	return nil
}

func (m *NetInstStatusFilter) validateRunState(formats strfmt.Registry) error {
	if swag.IsZero(m.RunState) { // not required
		return nil
	}

	if m.RunState != nil {
		if err := m.RunState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("runState")
			}
			return err
		}
	}

	return nil
}

func (m *NetInstStatusFilter) validateType(formats strfmt.Registry) error {
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

// ContextValidate validate this net inst status filter based on the context it is used
func (m *NetInstStatusFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKind(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRunState(ctx, formats); err != nil {
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

func (m *NetInstStatusFilter) contextValidateKind(ctx context.Context, formats strfmt.Registry) error {

	if m.Kind != nil {
		if err := m.Kind.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kind")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kind")
			}
			return err
		}
	}

	return nil
}

func (m *NetInstStatusFilter) contextValidateRunState(ctx context.Context, formats strfmt.Registry) error {

	if m.RunState != nil {
		if err := m.RunState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("runState")
			}
			return err
		}
	}

	return nil
}

func (m *NetInstStatusFilter) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
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
func (m *NetInstStatusFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetInstStatusFilter) UnmarshalBinary(b []byte) error {
	var res NetInstStatusFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
