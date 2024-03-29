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

// ConfigDrive XXX the Drive will be deprecated and we will use Volumes instead
//
// swagger:model configDrive
type ConfigDrive struct {

	// drvtype
	Drvtype *ConfigDriveType `json:"drvtype,omitempty"`

	// image
	Image *EveconfigImage `json:"image,omitempty"`

	// maxsizebytes indicates the maximum size of the volume.
	// Initial image size will be resized to the maxsizebytes
	// iff maxsizebytes is greater than the image size.
	Maxsizebytes string `json:"maxsizebytes,omitempty"`

	// preserve
	Preserve bool `json:"preserve,omitempty"`

	// readonly
	Readonly bool `json:"readonly,omitempty"`

	// target
	Target *ConfigTarget `json:"target,omitempty"`
}

// Validate validates this config drive
func (m *ConfigDrive) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDrvtype(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigDrive) validateDrvtype(formats strfmt.Registry) error {
	if swag.IsZero(m.Drvtype) { // not required
		return nil
	}

	if m.Drvtype != nil {
		if err := m.Drvtype.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("drvtype")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("drvtype")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigDrive) validateImage(formats strfmt.Registry) error {
	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if m.Image != nil {
		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigDrive) validateTarget(formats strfmt.Registry) error {
	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config drive based on the context it is used
func (m *ConfigDrive) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDrvtype(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigDrive) contextValidateDrvtype(ctx context.Context, formats strfmt.Registry) error {

	if m.Drvtype != nil {

		if swag.IsZero(m.Drvtype) { // not required
			return nil
		}

		if err := m.Drvtype.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("drvtype")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("drvtype")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigDrive) contextValidateImage(ctx context.Context, formats strfmt.Registry) error {

	if m.Image != nil {

		if swag.IsZero(m.Image) { // not required
			return nil
		}

		if err := m.Image.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigDrive) contextValidateTarget(ctx context.Context, formats strfmt.Registry) error {

	if m.Target != nil {

		if swag.IsZero(m.Target) { // not required
			return nil
		}

		if err := m.Target.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigDrive) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigDrive) UnmarshalBinary(b []byte) error {
	var res ConfigDrive
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
