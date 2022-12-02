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
	"github.com/go-openapi/validate"
)

// DocPolicySummary DocPolicy detail
//
// DocPolicy meta data
// Example: {"fileURL":"xxxxxxxx","policyName":"xxxxxxxxx","revision":{"createdAt":"2020-07-17T06:03:14Z","createdBy":"us.root@zededa.com","curr":"1","prev":"","updatedAt":"2020-07-17T06:03:14Z","updatedBy":"us.root@zededa.com"},"version":"xxxxxxx"}
//
// swagger:model DocPolicySummary
type DocPolicySummary struct {

	// Policy doc fileURL
	FileURL string `json:"fileURL,omitempty"`

	// User defined name of the docpolicy. Name cannot be changed once created
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	Policy string `json:"policy,omitempty"`

	// System defined info
	// Read Only: true
	Revision *ObjectRevision `json:"revision,omitempty"`

	// Policy doc version
	Version string `json:"version,omitempty"`
}

// Validate validates this doc policy summary
func (m *DocPolicySummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DocPolicySummary) validatePolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.Policy) { // not required
		return nil
	}

	if err := validate.MinLength("policy", "body", m.Policy, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("policy", "body", m.Policy, 256); err != nil {
		return err
	}

	if err := validate.Pattern("policy", "body", m.Policy, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *DocPolicySummary) validateRevision(formats strfmt.Registry) error {
	if swag.IsZero(m.Revision) { // not required
		return nil
	}

	if m.Revision != nil {
		if err := m.Revision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("revision")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this doc policy summary based on the context it is used
func (m *DocPolicySummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DocPolicySummary) contextValidateRevision(ctx context.Context, formats strfmt.Registry) error {

	if m.Revision != nil {
		if err := m.Revision.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("revision")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DocPolicySummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DocPolicySummary) UnmarshalBinary(b []byte) error {
	var res DocPolicySummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
