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

// VMManifestSummary VMManifestSummary - App summary for list views
//
// swagger:model VMManifestSummary
type VMManifestSummary struct {

	// Edge Application type
	//
	// UI map: N/A - not exposed to users
	AcKind *string `json:"acKind,omitempty"`

	// Manifest version
	//
	// UI map: N/A - not exposed to users
	AcVersion *string `json:"acVersion,omitempty"`

	// bundle type, eg: vm, container, module
	AppType *AppType `json:"appType,omitempty"`

	// type of deployment for the app, eg: azure, k3s, standalone
	DeploymentType *DeploymentType `json:"deploymentType,omitempty"`

	// Details of the Edge App
	Desc *Details `json:"desc,omitempty"`

	// Description of the Edge application
	//
	// UI map: AppDetailsPage:IdentityPane:DescriptionField, AppMarketplacePage:AppCard:DescriptionField
	// Pattern: [0-9A-Za-z-]+
	Description string `json:"description,omitempty"`

	// Display name or title of app manifest
	//
	// UI map: AppEditPage:IdentityPane:Title_Field, AppDetailsPage:IdentityPane:Title_Field
	DisplayName string `json:"displayName,omitempty"`

	// Module specific details
	//
	// Azure module specific details like module twin, environment variable, routes
	Module *ModuleSummary `json:"module,omitempty"`

	// Unique id of app manifest, should match object name
	//
	// UI map: AppEditPage:IdentityPane:Name_Field, AppDetailsPage:IdentityPane:Name_Field
	Name string `json:"name,omitempty"`
}

// Validate validates this VM manifest summary
func (m *VMManifestSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeploymentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDesc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMManifestSummary) validateAppType(formats strfmt.Registry) error {
	if swag.IsZero(m.AppType) { // not required
		return nil
	}

	if m.AppType != nil {
		if err := m.AppType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appType")
			}
			return err
		}
	}

	return nil
}

func (m *VMManifestSummary) validateDeploymentType(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentType) { // not required
		return nil
	}

	if m.DeploymentType != nil {
		if err := m.DeploymentType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deploymentType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deploymentType")
			}
			return err
		}
	}

	return nil
}

func (m *VMManifestSummary) validateDesc(formats strfmt.Registry) error {
	if swag.IsZero(m.Desc) { // not required
		return nil
	}

	if m.Desc != nil {
		if err := m.Desc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("desc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("desc")
			}
			return err
		}
	}

	return nil
}

func (m *VMManifestSummary) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.Pattern("description", "body", m.Description, `[0-9A-Za-z-]+`); err != nil {
		return err
	}

	return nil
}

func (m *VMManifestSummary) validateModule(formats strfmt.Registry) error {
	if swag.IsZero(m.Module) { // not required
		return nil
	}

	if m.Module != nil {
		if err := m.Module.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("module")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("module")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this VM manifest summary based on the context it is used
func (m *VMManifestSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeploymentType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDesc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMManifestSummary) contextValidateAppType(ctx context.Context, formats strfmt.Registry) error {

	if m.AppType != nil {

		if swag.IsZero(m.AppType) { // not required
			return nil
		}

		if err := m.AppType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appType")
			}
			return err
		}
	}

	return nil
}

func (m *VMManifestSummary) contextValidateDeploymentType(ctx context.Context, formats strfmt.Registry) error {

	if m.DeploymentType != nil {

		if swag.IsZero(m.DeploymentType) { // not required
			return nil
		}

		if err := m.DeploymentType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deploymentType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deploymentType")
			}
			return err
		}
	}

	return nil
}

func (m *VMManifestSummary) contextValidateDesc(ctx context.Context, formats strfmt.Registry) error {

	if m.Desc != nil {

		if swag.IsZero(m.Desc) { // not required
			return nil
		}

		if err := m.Desc.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("desc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("desc")
			}
			return err
		}
	}

	return nil
}

func (m *VMManifestSummary) contextValidateModule(ctx context.Context, formats strfmt.Registry) error {

	if m.Module != nil {

		if swag.IsZero(m.Module) { // not required
			return nil
		}

		if err := m.Module.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("module")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("module")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMManifestSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMManifestSummary) UnmarshalBinary(b []byte) error {
	var res VMManifestSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
