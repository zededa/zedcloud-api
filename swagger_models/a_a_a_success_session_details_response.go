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

// AAASuccessSessionDetailsResponse a a a success session details response
//
// swagger:model AAASuccessSessionDetailsResponse
type AAASuccessSessionDetailsResponse struct {

	// original
	Original *OpaqueToken64 `json:"original,omitempty"`

	// policies
	Policies []*Policy `json:"policies"`

	// user
	User *DetailedUser `json:"user,omitempty"`
}

// Validate validates this a a a success session details response
func (m *AAASuccessSessionDetailsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOriginal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AAASuccessSessionDetailsResponse) validateOriginal(formats strfmt.Registry) error {
	if swag.IsZero(m.Original) { // not required
		return nil
	}

	if m.Original != nil {
		if err := m.Original.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("original")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("original")
			}
			return err
		}
	}

	return nil
}

func (m *AAASuccessSessionDetailsResponse) validatePolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.Policies) { // not required
		return nil
	}

	for i := 0; i < len(m.Policies); i++ {
		if swag.IsZero(m.Policies[i]) { // not required
			continue
		}

		if m.Policies[i] != nil {
			if err := m.Policies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AAASuccessSessionDetailsResponse) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this a a a success session details response based on the context it is used
func (m *AAASuccessSessionDetailsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOriginal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AAASuccessSessionDetailsResponse) contextValidateOriginal(ctx context.Context, formats strfmt.Registry) error {

	if m.Original != nil {

		if swag.IsZero(m.Original) { // not required
			return nil
		}

		if err := m.Original.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("original")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("original")
			}
			return err
		}
	}

	return nil
}

func (m *AAASuccessSessionDetailsResponse) contextValidatePolicies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Policies); i++ {

		if m.Policies[i] != nil {

			if swag.IsZero(m.Policies[i]) { // not required
				return nil
			}

			if err := m.Policies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AAASuccessSessionDetailsResponse) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {

		if swag.IsZero(m.User) { // not required
			return nil
		}

		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AAASuccessSessionDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AAASuccessSessionDetailsResponse) UnmarshalBinary(b []byte) error {
	var res AAASuccessSessionDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
