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

// Details Edge Application detail
//
// Edge Application Artifact Details
// Example: {"appCategory":"APP_CATEGORY_CLOUD_APPLICATION","os":"Linux"}
//
// swagger:model Details
type Details struct {

	// agreement list
	AgreementList map[string]string `json:"agreementList,omitempty"`

	// Edge application category
	// Required: true
	AppCategory *AppCategory `json:"appCategory"`

	// Type of the Edge application
	//
	// UI map: AppMarketplacePage:AppCard:DescriptionField, AppEditPage:IdentityPane:CategoryField, AppDetailsPage:IdentityPane:CategoryField
	Category *string `json:"category,omitempty"`

	// Schema: {<license_name>:<url>}
	LicenseList map[string]string `json:"licenseList,omitempty"`

	// logo
	Logo map[string]string `json:"logo,omitempty"`

	// Edge application's Operating System
	Os string `json:"os,omitempty"`

	// screenshot list
	ScreenshotList map[string]string `json:"screenshotList,omitempty"`

	// Support Description
	//
	// UI map: AppEditPage:DeveloperPane:Support_Description_Field, AppDetailsPage:DeveloperPane:Support_Description_Field
	Support string `json:"support,omitempty"`
}

// Validate validates this details
func (m *Details) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppCategory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Details) validateAppCategory(formats strfmt.Registry) error {

	if err := validate.Required("appCategory", "body", m.AppCategory); err != nil {
		return err
	}

	if err := validate.Required("appCategory", "body", m.AppCategory); err != nil {
		return err
	}

	if m.AppCategory != nil {
		if err := m.AppCategory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appCategory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appCategory")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this details based on the context it is used
func (m *Details) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppCategory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Details) contextValidateAppCategory(ctx context.Context, formats strfmt.Registry) error {

	if m.AppCategory != nil {
		if err := m.AppCategory.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appCategory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appCategory")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Details) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Details) UnmarshalBinary(b []byte) error {
	var res Details
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
