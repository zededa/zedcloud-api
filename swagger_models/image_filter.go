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

// ImageFilter image filter
//
// swagger:model ImageFilter
type ImageFilter struct {

	// Datastore id to be matched.
	// Pattern: [a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}
	DatastoreID string `json:"datastoreId,omitempty"`

	// Image architecture to be matched.
	ImageArch *ModelArchType `json:"imageArch,omitempty"`

	// Image status to be matched.
	ImageStatus *ImageStatus `json:"imageStatus,omitempty"`

	// Image type to ne matched.
	ImageType *ImageType `json:"imageType,omitempty"`

	// Image name pattern to be matched.
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9_.-]{3,256}
	NamePattern string `json:"namePattern,omitempty"`
}

// Validate validates this image filter
func (m *ImageFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatastoreID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageArch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamePattern(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageFilter) validateDatastoreID(formats strfmt.Registry) error {
	if swag.IsZero(m.DatastoreID) { // not required
		return nil
	}

	if err := validate.Pattern("datastoreId", "body", m.DatastoreID, `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}`); err != nil {
		return err
	}

	return nil
}

func (m *ImageFilter) validateImageArch(formats strfmt.Registry) error {
	if swag.IsZero(m.ImageArch) { // not required
		return nil
	}

	if m.ImageArch != nil {
		if err := m.ImageArch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("imageArch")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("imageArch")
			}
			return err
		}
	}

	return nil
}

func (m *ImageFilter) validateImageStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.ImageStatus) { // not required
		return nil
	}

	if m.ImageStatus != nil {
		if err := m.ImageStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("imageStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("imageStatus")
			}
			return err
		}
	}

	return nil
}

func (m *ImageFilter) validateImageType(formats strfmt.Registry) error {
	if swag.IsZero(m.ImageType) { // not required
		return nil
	}

	if m.ImageType != nil {
		if err := m.ImageType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("imageType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("imageType")
			}
			return err
		}
	}

	return nil
}

func (m *ImageFilter) validateNamePattern(formats strfmt.Registry) error {
	if swag.IsZero(m.NamePattern) { // not required
		return nil
	}

	if err := validate.MinLength("namePattern", "body", m.NamePattern, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("namePattern", "body", m.NamePattern, 256); err != nil {
		return err
	}

	if err := validate.Pattern("namePattern", "body", m.NamePattern, `[a-zA-Z0-9_.-]{3,256}`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this image filter based on the context it is used
func (m *ImageFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateImageArch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImageStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImageType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageFilter) contextValidateImageArch(ctx context.Context, formats strfmt.Registry) error {

	if m.ImageArch != nil {

		if swag.IsZero(m.ImageArch) { // not required
			return nil
		}

		if err := m.ImageArch.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("imageArch")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("imageArch")
			}
			return err
		}
	}

	return nil
}

func (m *ImageFilter) contextValidateImageStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.ImageStatus != nil {

		if swag.IsZero(m.ImageStatus) { // not required
			return nil
		}

		if err := m.ImageStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("imageStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("imageStatus")
			}
			return err
		}
	}

	return nil
}

func (m *ImageFilter) contextValidateImageType(ctx context.Context, formats strfmt.Registry) error {

	if m.ImageType != nil {

		if swag.IsZero(m.ImageType) { // not required
			return nil
		}

		if err := m.ImageType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("imageType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("imageType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageFilter) UnmarshalBinary(b []byte) error {
	var res ImageFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
