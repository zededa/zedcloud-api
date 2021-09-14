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
	"github.com/go-openapi/validate"
)

// AppInstFilter app inst filter
//
// swagger:model AppInstFilter
type AppInstFilter struct {

	// User defined name of the app instance, unique across the enterprise. Once app instance is created, name can’t be changed
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	AppName string `json:"appName,omitempty"`

	// type of bundle
	AppType *AppType `json:"appType,omitempty"`

	// type of deployment for the app, eg: azure, k3s, standalone
	DeploymentType *DeploymentType `json:"deploymentType,omitempty"`

	// User defined name of the device, unique across the enterprise. Once device is created, name can’t be changed
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	DeviceName string `json:"deviceName,omitempty"`

	// name pattern
	NamePattern string `json:"namePattern,omitempty"`

	// User defined name of the project, unique across the enterprise. Once project is created, name can’t be changed
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	ProjectName string `json:"projectName,omitempty"`
}

// Validate validates this app inst filter
func (m *AppInstFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeploymentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstFilter) validateAppName(formats strfmt.Registry) error {
	if swag.IsZero(m.AppName) { // not required
		return nil
	}

	if err := validate.MinLength("appName", "body", m.AppName, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("appName", "body", m.AppName, 256); err != nil {
		return err
	}

	if err := validate.Pattern("appName", "body", m.AppName, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *AppInstFilter) validateAppType(formats strfmt.Registry) error {
	if swag.IsZero(m.AppType) { // not required
		return nil
	}

	if m.AppType != nil {
		if err := m.AppType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appType")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstFilter) validateDeploymentType(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentType) { // not required
		return nil
	}

	if m.DeploymentType != nil {
		if err := m.DeploymentType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deploymentType")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstFilter) validateDeviceName(formats strfmt.Registry) error {
	if swag.IsZero(m.DeviceName) { // not required
		return nil
	}

	if err := validate.MinLength("deviceName", "body", m.DeviceName, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("deviceName", "body", m.DeviceName, 256); err != nil {
		return err
	}

	if err := validate.Pattern("deviceName", "body", m.DeviceName, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *AppInstFilter) validateProjectName(formats strfmt.Registry) error {
	if swag.IsZero(m.ProjectName) { // not required
		return nil
	}

	if err := validate.MinLength("projectName", "body", m.ProjectName, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("projectName", "body", m.ProjectName, 256); err != nil {
		return err
	}

	if err := validate.Pattern("projectName", "body", m.ProjectName, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this app inst filter based on the context it is used
func (m *AppInstFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeploymentType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstFilter) contextValidateAppType(ctx context.Context, formats strfmt.Registry) error {

	if m.AppType != nil {
		if err := m.AppType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appType")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstFilter) contextValidateDeploymentType(ctx context.Context, formats strfmt.Registry) error {

	if m.DeploymentType != nil {
		if err := m.DeploymentType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deploymentType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppInstFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppInstFilter) UnmarshalBinary(b []byte) error {
	var res AppInstFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
