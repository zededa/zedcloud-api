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

// NetworkFilter network filter
//
// swagger:model NetworkFilter
type NetworkFilter struct {

	// dist
	Dist *NetworkWirelessType `json:"dist,omitempty"`

	// kind
	Kind *NetworkKind `json:"kind,omitempty"`

	// name pattern
	NamePattern string `json:"namePattern,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`

	// project name pattern
	ProjectNamePattern string `json:"projectNamePattern,omitempty"`
}

// Validate validates this network filter
func (m *NetworkFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDist(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkFilter) validateDist(formats strfmt.Registry) error {
	if swag.IsZero(m.Dist) { // not required
		return nil
	}

	if m.Dist != nil {
		if err := m.Dist.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dist")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dist")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkFilter) validateKind(formats strfmt.Registry) error {
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

// ContextValidate validate this network filter based on the context it is used
func (m *NetworkFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDist(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKind(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkFilter) contextValidateDist(ctx context.Context, formats strfmt.Registry) error {

	if m.Dist != nil {

		if swag.IsZero(m.Dist) { // not required
			return nil
		}

		if err := m.Dist.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dist")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dist")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkFilter) contextValidateKind(ctx context.Context, formats strfmt.Registry) error {

	if m.Kind != nil {

		if swag.IsZero(m.Kind) { // not required
			return nil
		}

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

// MarshalBinary interface implementation
func (m *NetworkFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkFilter) UnmarshalBinary(b []byte) error {
	var res NetworkFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
