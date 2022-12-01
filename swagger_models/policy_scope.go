// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// PolicyScope Scope of the permission, similar to Unix file system
//
// swagger:model PolicyScope
type PolicyScope string

func NewPolicyScope(value PolicyScope) *PolicyScope {
	return &value
}

// Pointer returns a pointer to a freshly-allocated PolicyScope.
func (m PolicyScope) Pointer() *PolicyScope {
	return &m
}

const (

	// PolicyScopePermissionScopeINVALID captures enum value "PermissionScope__INVALID__"
	PolicyScopePermissionScopeINVALID PolicyScope = "PermissionScope__INVALID__"

	// PolicyScopePermissionScopeOther captures enum value "PermissionScopeOther"
	PolicyScopePermissionScopeOther PolicyScope = "PermissionScopeOther"

	// PolicyScopePermissionScopeGroup captures enum value "PermissionScopeGroup"
	PolicyScopePermissionScopeGroup PolicyScope = "PermissionScopeGroup"

	// PolicyScopePermissionScopeOtherGroup captures enum value "PermissionScopeOtherGroup"
	PolicyScopePermissionScopeOtherGroup PolicyScope = "PermissionScopeOtherGroup"

	// PolicyScopePermissionScopeOwner captures enum value "PermissionScopeOwner"
	PolicyScopePermissionScopeOwner PolicyScope = "PermissionScopeOwner"

	// PolicyScopePermissionScopeOtherOwner captures enum value "PermissionScopeOtherOwner"
	PolicyScopePermissionScopeOtherOwner PolicyScope = "PermissionScopeOtherOwner"

	// PolicyScopePermissionScopeGroupOwner captures enum value "PermissionScopeGroupOwner"
	PolicyScopePermissionScopeGroupOwner PolicyScope = "PermissionScopeGroupOwner"

	// PolicyScopePermissionScopeOtherGroupOwner captures enum value "PermissionScopeOtherGroupOwner"
	PolicyScopePermissionScopeOtherGroupOwner PolicyScope = "PermissionScopeOtherGroupOwner"

	// PolicyScopePermissionScopeEnterprise captures enum value "PermissionScopeEnterprise"
	PolicyScopePermissionScopeEnterprise PolicyScope = "PermissionScopeEnterprise"

	// PolicyScopePermissionScopeOtherEnterprise captures enum value "PermissionScopeOtherEnterprise"
	PolicyScopePermissionScopeOtherEnterprise PolicyScope = "PermissionScopeOtherEnterprise"

	// PolicyScopePermissionScopeGroupEnterprise captures enum value "PermissionScopeGroupEnterprise"
	PolicyScopePermissionScopeGroupEnterprise PolicyScope = "PermissionScopeGroupEnterprise"

	// PolicyScopePermissionScopeOtherGroupEnterprise captures enum value "PermissionScopeOtherGroupEnterprise"
	PolicyScopePermissionScopeOtherGroupEnterprise PolicyScope = "PermissionScopeOtherGroupEnterprise"

	// PolicyScopePermissionScopeOwnerEnterprise captures enum value "PermissionScopeOwnerEnterprise"
	PolicyScopePermissionScopeOwnerEnterprise PolicyScope = "PermissionScopeOwnerEnterprise"

	// PolicyScopePermissionScopeOtherOwnerEnterprise captures enum value "PermissionScopeOtherOwnerEnterprise"
	PolicyScopePermissionScopeOtherOwnerEnterprise PolicyScope = "PermissionScopeOtherOwnerEnterprise"

	// PolicyScopePermissionScopeGroupOwnerEnterprise captures enum value "PermissionScopeGroupOwnerEnterprise"
	PolicyScopePermissionScopeGroupOwnerEnterprise PolicyScope = "PermissionScopeGroupOwnerEnterprise"

	// PolicyScopePermissionScopeOtherGroupOwnerEnterprise captures enum value "PermissionScopeOtherGroupOwnerEnterprise"
	PolicyScopePermissionScopeOtherGroupOwnerEnterprise PolicyScope = "PermissionScopeOtherGroupOwnerEnterprise"
)

// for schema
var policyScopeEnum []interface{}

func init() {
	var res []PolicyScope
	if err := json.Unmarshal([]byte(`["PermissionScope__INVALID__","PermissionScopeOther","PermissionScopeGroup","PermissionScopeOtherGroup","PermissionScopeOwner","PermissionScopeOtherOwner","PermissionScopeGroupOwner","PermissionScopeOtherGroupOwner","PermissionScopeEnterprise","PermissionScopeOtherEnterprise","PermissionScopeGroupEnterprise","PermissionScopeOtherGroupEnterprise","PermissionScopeOwnerEnterprise","PermissionScopeOtherOwnerEnterprise","PermissionScopeGroupOwnerEnterprise","PermissionScopeOtherGroupOwnerEnterprise"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		policyScopeEnum = append(policyScopeEnum, v)
	}
}

func (m PolicyScope) validatePolicyScopeEnum(path, location string, value PolicyScope) error {
	if err := validate.EnumCase(path, location, value, policyScopeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this policy scope
func (m PolicyScope) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePolicyScopeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this policy scope based on context it is used
func (m PolicyScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
