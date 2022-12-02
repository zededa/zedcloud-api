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

// CrudContentProtobuf crud content protobuf
//
// swagger:model CrudContentProtobuf
type CrudContentProtobuf struct {

	// credential
	Credential *Credential `json:"credential,omitempty"`

	// crypto key
	CryptoKey string `json:"cryptoKey,omitempty"`

	// detailed user
	DetailedUser *DetailedUser `json:"detailedUser,omitempty"`

	// doc policy
	DocPolicy *DocPolicy `json:"docPolicy,omitempty"`

	// encrypted secrets
	EncryptedSecrets map[string]string `json:"encryptedSecrets,omitempty"`

	// enterprise
	Enterprise *Enterprise `json:"enterprise,omitempty"`

	// policy
	Policy *Policy `json:"policy,omitempty"`

	// profile
	Profile *AuthorizationProfile `json:"profile,omitempty"`

	// realm
	Realm *Realm `json:"realm,omitempty"`

	// role
	Role *Role `json:"role,omitempty"`

	// simple user
	SimpleUser *SimpleUser `json:"simpleUser,omitempty"`

	// user
	User *DetailedUser `json:"user,omitempty"`
}

// Validate validates this crud content protobuf
func (m *CrudContentProtobuf) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredential(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDetailedUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDocPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnterprise(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRealm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSimpleUser(formats); err != nil {
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

func (m *CrudContentProtobuf) validateCredential(formats strfmt.Registry) error {
	if swag.IsZero(m.Credential) { // not required
		return nil
	}

	if m.Credential != nil {
		if err := m.Credential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credential")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credential")
			}
			return err
		}
	}

	return nil
}

func (m *CrudContentProtobuf) validateDetailedUser(formats strfmt.Registry) error {
	if swag.IsZero(m.DetailedUser) { // not required
		return nil
	}

	if m.DetailedUser != nil {
		if err := m.DetailedUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("detailedUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("detailedUser")
			}
			return err
		}
	}

	return nil
}

func (m *CrudContentProtobuf) validateDocPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.DocPolicy) { // not required
		return nil
	}

	if m.DocPolicy != nil {
		if err := m.DocPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("docPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("docPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *CrudContentProtobuf) validateEnterprise(formats strfmt.Registry) error {
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

func (m *CrudContentProtobuf) validatePolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.Policy) { // not required
		return nil
	}

	if m.Policy != nil {
		if err := m.Policy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policy")
			}
			return err
		}
	}

	return nil
}

func (m *CrudContentProtobuf) validateProfile(formats strfmt.Registry) error {
	if swag.IsZero(m.Profile) { // not required
		return nil
	}

	if m.Profile != nil {
		if err := m.Profile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("profile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("profile")
			}
			return err
		}
	}

	return nil
}

func (m *CrudContentProtobuf) validateRealm(formats strfmt.Registry) error {
	if swag.IsZero(m.Realm) { // not required
		return nil
	}

	if m.Realm != nil {
		if err := m.Realm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("realm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("realm")
			}
			return err
		}
	}

	return nil
}

func (m *CrudContentProtobuf) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if m.Role != nil {
		if err := m.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

func (m *CrudContentProtobuf) validateSimpleUser(formats strfmt.Registry) error {
	if swag.IsZero(m.SimpleUser) { // not required
		return nil
	}

	if m.SimpleUser != nil {
		if err := m.SimpleUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("simpleUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("simpleUser")
			}
			return err
		}
	}

	return nil
}

func (m *CrudContentProtobuf) validateUser(formats strfmt.Registry) error {
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

// ContextValidate validate this crud content protobuf based on the context it is used
func (m *CrudContentProtobuf) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCredential(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDetailedUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDocPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnterprise(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProfile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRealm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRole(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSimpleUser(ctx, formats); err != nil {
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

func (m *CrudContentProtobuf) contextValidateCredential(ctx context.Context, formats strfmt.Registry) error {

	if m.Credential != nil {
		if err := m.Credential.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credential")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credential")
			}
			return err
		}
	}

	return nil
}

func (m *CrudContentProtobuf) contextValidateDetailedUser(ctx context.Context, formats strfmt.Registry) error {

	if m.DetailedUser != nil {
		if err := m.DetailedUser.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("detailedUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("detailedUser")
			}
			return err
		}
	}

	return nil
}

func (m *CrudContentProtobuf) contextValidateDocPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.DocPolicy != nil {
		if err := m.DocPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("docPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("docPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *CrudContentProtobuf) contextValidateEnterprise(ctx context.Context, formats strfmt.Registry) error {

	if m.Enterprise != nil {
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

func (m *CrudContentProtobuf) contextValidatePolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.Policy != nil {
		if err := m.Policy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policy")
			}
			return err
		}
	}

	return nil
}

func (m *CrudContentProtobuf) contextValidateProfile(ctx context.Context, formats strfmt.Registry) error {

	if m.Profile != nil {
		if err := m.Profile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("profile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("profile")
			}
			return err
		}
	}

	return nil
}

func (m *CrudContentProtobuf) contextValidateRealm(ctx context.Context, formats strfmt.Registry) error {

	if m.Realm != nil {
		if err := m.Realm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("realm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("realm")
			}
			return err
		}
	}

	return nil
}

func (m *CrudContentProtobuf) contextValidateRole(ctx context.Context, formats strfmt.Registry) error {

	if m.Role != nil {
		if err := m.Role.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

func (m *CrudContentProtobuf) contextValidateSimpleUser(ctx context.Context, formats strfmt.Registry) error {

	if m.SimpleUser != nil {
		if err := m.SimpleUser.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("simpleUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("simpleUser")
			}
			return err
		}
	}

	return nil
}

func (m *CrudContentProtobuf) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {
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
func (m *CrudContentProtobuf) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CrudContentProtobuf) UnmarshalBinary(b []byte) error {
	var res CrudContentProtobuf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
