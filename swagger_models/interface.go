// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

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

// Interface interface
//
// swagger:model Interface
type Interface struct {

	// Traffic access control rules for this interface. Applicable only when "direct attach" flag is false.
	Acls []*ACL `json:"acls"`

	// If true, a physical adapter is assigned to the edge application directly. If false, a network instance is assigned to the edge application.
	Directattach bool `json:"directattach,omitempty"`

	// Interface name used by the edge application
	Name string `json:"name,omitempty"`

	// Indicates if the interface is optional for edge application.
	Optional bool `json:"optional,omitempty"`

	// If true, DHCP network can't be assigned and user needs to provide a static IP address.
	Privateip bool `json:"privateip,omitempty"`

	// Physical Adapter type for this interface. Applicable only when "direct attach" flag is true.
	Type string `json:"type,omitempty"`
}

// Validate validates this interface
func (m *Interface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcls(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Interface) validateAcls(formats strfmt.Registry) error {
	if swag.IsZero(m.Acls) { // not required
		return nil
	}

	for i := 0; i < len(m.Acls); i++ {
		if swag.IsZero(m.Acls[i]) { // not required
			continue
		}

		if m.Acls[i] != nil {
			if err := m.Acls[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("acls" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("acls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this interface based on the context it is used
func (m *Interface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAcls(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Interface) contextValidateAcls(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Acls); i++ {

		if m.Acls[i] != nil {
			if err := m.Acls[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("acls" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("acls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Interface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Interface) UnmarshalBinary(b []byte) error {
	var res Interface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
