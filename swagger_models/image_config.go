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

// ImageConfig Image metadata detail
//
// Image metadata for edge gateway Base OS or for eedge applications.
// Example: {"description":"My test image in QCOW2 format for Edge computing","dsId":"7927f6e3-484d-4105-a98e-868b21c1cb61","id":"d1125b0f-633d-459c-99c6-f47e7467cebc","imageArch":"AMD64","imageError":"Image uplinked successfully...","imageFormat":3,"imageLocal":"","imageRelUrl":"edge/computing/AMD64","imageSha256":"ADC5BB9DD39F83DD74C276B0BA119FB27925A5CDEA343FE1F2C8433F28AB345B","imageSizeBytes":142016512,"imageStatus":4,"imageType":2,"imageVersion":"","name":"my-test-image","originType":2,"revision":{"createdAt":{"seconds":1592068270},"createdBy":"admin@my-company.com","curr":"1","updatedAt":{"seconds":1592068271},"updatedBy":"admin@my-company.com"},"title":"My Test Image for Edge Computing"}
//
// swagger:model ImageConfig
type ImageConfig struct {

	// Datastore Id where image binary is located.
	// Required: true
	// Pattern: [a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}
	DatastoreID *string `json:"datastoreId"`

	// Detailed description of the image.
	// Max Length: 256
	Description string `json:"description,omitempty"`

	// System defined universally unique Id of the image.
	// Read Only: true
	// Pattern: [a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}
	ID string `json:"id,omitempty"`

	// Image Architecture.
	// Required: true
	ImageArch *ModelArchType `json:"imageArch"`

	// Image upload/uplink detailed error/status message
	// Read Only: true
	ImageError string `json:"imageError,omitempty"`

	// Image binary format.
	// Required: true
	ImageFormat *ConfigFormat `json:"imageFormat"`

	// Internal image location.
	// Read Only: true
	ImageLocal string `json:"imageLocal,omitempty"`

	// Image relative path w.r.t. Datastore
	ImageRelURL string `json:"imageRelUrl,omitempty"`

	// Image checksum in SHA256 format
	ImageSha256 string `json:"imageSha256,omitempty"`

	// Image size in KBytes.
	ImageSizeBytes string `json:"imageSizeBytes,omitempty"`

	// Image status
	// Read Only: true
	ImageStatus *ImageStatus `json:"imageStatus,omitempty"`

	// Image type
	// Required: true
	ImageType *ImageType `json:"imageType"`

	// system defined info
	ImageVersion string `json:"imageVersion,omitempty"`

	// User defined name of the image, unique across the enterprise. Once image is created, name can’t be changed.
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	Name *string `json:"name"`

	// Origin type of image.
	// Read Only: true
	OriginType *Origin `json:"originType,omitempty"`

	// project access list of the image
	ProjectAccessList []string `json:"projectAccessList"`

	// system defined info
	// Read Only: true
	Revision *ObjectRevision `json:"revision,omitempty"`

	// User defined title of the image. Title can be changed at any time.
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+
	Title *string `json:"title"`
}

// Validate validates this image config
func (m *ImageConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatastoreID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageArch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageConfig) validateDatastoreID(formats strfmt.Registry) error {

	if err := validate.Required("datastoreId", "body", m.DatastoreID); err != nil {
		return err
	}

	if err := validate.Pattern("datastoreId", "body", *m.DatastoreID, `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}`); err != nil {
		return err
	}

	return nil
}

func (m *ImageConfig) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 256); err != nil {
		return err
	}

	return nil
}

func (m *ImageConfig) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.Pattern("id", "body", m.ID, `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}`); err != nil {
		return err
	}

	return nil
}

func (m *ImageConfig) validateImageArch(formats strfmt.Registry) error {

	if err := validate.Required("imageArch", "body", m.ImageArch); err != nil {
		return err
	}

	if err := validate.Required("imageArch", "body", m.ImageArch); err != nil {
		return err
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

func (m *ImageConfig) validateImageFormat(formats strfmt.Registry) error {

	if err := validate.Required("imageFormat", "body", m.ImageFormat); err != nil {
		return err
	}

	if err := validate.Required("imageFormat", "body", m.ImageFormat); err != nil {
		return err
	}

	if m.ImageFormat != nil {
		if err := m.ImageFormat.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("imageFormat")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("imageFormat")
			}
			return err
		}
	}

	return nil
}

func (m *ImageConfig) validateImageStatus(formats strfmt.Registry) error {
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

func (m *ImageConfig) validateImageType(formats strfmt.Registry) error {

	if err := validate.Required("imageType", "body", m.ImageType); err != nil {
		return err
	}

	if err := validate.Required("imageType", "body", m.ImageType); err != nil {
		return err
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

func (m *ImageConfig) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 256); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", *m.Name, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *ImageConfig) validateOriginType(formats strfmt.Registry) error {
	if swag.IsZero(m.OriginType) { // not required
		return nil
	}

	if m.OriginType != nil {
		if err := m.OriginType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("originType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("originType")
			}
			return err
		}
	}

	return nil
}

func (m *ImageConfig) validateRevision(formats strfmt.Registry) error {
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

func (m *ImageConfig) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	if err := validate.MinLength("title", "body", *m.Title, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", *m.Title, 256); err != nil {
		return err
	}

	if err := validate.Pattern("title", "body", *m.Title, `[a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this image config based on the context it is used
func (m *ImageConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImageArch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImageError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImageFormat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImageLocal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImageStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImageType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOriginType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageConfig) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ImageConfig) contextValidateImageArch(ctx context.Context, formats strfmt.Registry) error {

	if m.ImageArch != nil {
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

func (m *ImageConfig) contextValidateImageError(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "imageError", "body", string(m.ImageError)); err != nil {
		return err
	}

	return nil
}

func (m *ImageConfig) contextValidateImageFormat(ctx context.Context, formats strfmt.Registry) error {

	if m.ImageFormat != nil {
		if err := m.ImageFormat.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("imageFormat")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("imageFormat")
			}
			return err
		}
	}

	return nil
}

func (m *ImageConfig) contextValidateImageLocal(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "imageLocal", "body", string(m.ImageLocal)); err != nil {
		return err
	}

	return nil
}

func (m *ImageConfig) contextValidateImageStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.ImageStatus != nil {
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

func (m *ImageConfig) contextValidateImageType(ctx context.Context, formats strfmt.Registry) error {

	if m.ImageType != nil {
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

func (m *ImageConfig) contextValidateOriginType(ctx context.Context, formats strfmt.Registry) error {

	if m.OriginType != nil {
		if err := m.OriginType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("originType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("originType")
			}
			return err
		}
	}

	return nil
}

func (m *ImageConfig) contextValidateRevision(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ImageConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageConfig) UnmarshalBinary(b []byte) error {
	var res ImageConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
