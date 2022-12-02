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

// Policy policy
//
// swagger:model Policy
type Policy struct {

	// access
	Access *PolicyAccess `json:"access,omitempty"`

	// external
	External bool `json:"external,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// owner Id
	OwnerID string `json:"ownerId,omitempty"`

	// scope
	Scope *PolicyScope `json:"scope,omitempty"`

	// zone
	Zone string `json:"zone,omitempty"`
}

// Validate validates this policy
func (m *Policy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Policy) validateAccess(formats strfmt.Registry) error {
	if swag.IsZero(m.Access) { // not required
		return nil
	}

	if m.Access != nil {
		if err := m.Access.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("access")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("access")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	if m.Scope != nil {
		if err := m.Scope.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scope")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scope")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this policy based on the context it is used
func (m *Policy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Policy) contextValidateAccess(ctx context.Context, formats strfmt.Registry) error {

	if m.Access != nil {
		if err := m.Access.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("access")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("access")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	if m.Scope != nil {
		if err := m.Scope.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scope")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scope")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Policy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Policy) UnmarshalBinary(b []byte) error {
	var res Policy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
