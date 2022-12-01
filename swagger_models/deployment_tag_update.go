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

// DeploymentTagUpdate deployment tag required for the devices
//
// deployment tag will be required to target deployment version on the device
//
// swagger:model DeploymentTagUpdate
type DeploymentTagUpdate struct {

	// deployment tag to be updated on the device
	// Required: true
	// Read Only: true
	DeploymentTag string `json:"deploymentTag"`
}

// Validate validates this deployment tag update
func (m *DeploymentTagUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeploymentTag(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentTagUpdate) validateDeploymentTag(formats strfmt.Registry) error {

	if err := validate.RequiredString("deploymentTag", "body", m.DeploymentTag); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this deployment tag update based on the context it is used
func (m *DeploymentTagUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeploymentTag(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentTagUpdate) contextValidateDeploymentTag(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "deploymentTag", "body", string(m.DeploymentTag)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentTagUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentTagUpdate) UnmarshalBinary(b []byte) error {
	var res DeploymentTagUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
