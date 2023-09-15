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

// AAARequestEnterpriseSignup a a a request enterprise signup
//
// swagger:model AAARequestEnterpriseSignup
type AAARequestEnterpriseSignup struct {

	// admin user
	AdminUser *AdminUserSignup `json:"adminUser,omitempty"`

	// enterprise
	Enterprise *Enterprise `json:"enterprise,omitempty"`

	// profile type
	ProfileType *AuthProfileType `json:"profileType,omitempty"`

	// realm list
	RealmList []string `json:"realmList"`

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this a a a request enterprise signup
func (m *AAARequestEnterpriseSignup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdminUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnterprise(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfileType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AAARequestEnterpriseSignup) validateAdminUser(formats strfmt.Registry) error {
	if swag.IsZero(m.AdminUser) { // not required
		return nil
	}

	if m.AdminUser != nil {
		if err := m.AdminUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("adminUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("adminUser")
			}
			return err
		}
	}

	return nil
}

func (m *AAARequestEnterpriseSignup) validateEnterprise(formats strfmt.Registry) error {
	if swag.IsZero(m.Enterprise) { // not required
		return nil
	}

	if m.Enterprise != nil {
		if err := m.Enterprise.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("enterprise")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("enterprise")
			}
			return err
		}
	}

	return nil
}

func (m *AAARequestEnterpriseSignup) validateProfileType(formats strfmt.Registry) error {
	if swag.IsZero(m.ProfileType) { // not required
		return nil
	}

	if m.ProfileType != nil {
		if err := m.ProfileType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("profileType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("profileType")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this a a a request enterprise signup based on the context it is used
func (m *AAARequestEnterpriseSignup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdminUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnterprise(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProfileType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AAARequestEnterpriseSignup) contextValidateAdminUser(ctx context.Context, formats strfmt.Registry) error {

	if m.AdminUser != nil {

		if swag.IsZero(m.AdminUser) { // not required
			return nil
		}

		if err := m.AdminUser.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("adminUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("adminUser")
			}
			return err
		}
	}

	return nil
}

func (m *AAARequestEnterpriseSignup) contextValidateEnterprise(ctx context.Context, formats strfmt.Registry) error {

	if m.Enterprise != nil {

		if swag.IsZero(m.Enterprise) { // not required
			return nil
		}

		if err := m.Enterprise.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("enterprise")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("enterprise")
			}
			return err
		}
	}

	return nil
}

func (m *AAARequestEnterpriseSignup) contextValidateProfileType(ctx context.Context, formats strfmt.Registry) error {

	if m.ProfileType != nil {

		if swag.IsZero(m.ProfileType) { // not required
			return nil
		}

		if err := m.ProfileType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("profileType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("profileType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AAARequestEnterpriseSignup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AAARequestEnterpriseSignup) UnmarshalBinary(b []byte) error {
	var res AAARequestEnterpriseSignup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
