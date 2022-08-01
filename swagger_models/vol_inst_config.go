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

// VolInstConfig vol inst config
//
// swagger:model VolInstConfig
type VolInstConfig struct {

	// Access mode
	Accessmode *VolumeInstanceAccessMode `json:"accessmode,omitempty"`

	// flag to keep the contents of the volume unencrypted (in clear text)
	Cleartext bool `json:"cleartext,omitempty"`

	// content tree ID
	ContentTreeID string `json:"contentTreeId,omitempty"`

	// Detailed description of the volume instance.
	// Max Length: 256
	Description string `json:"description,omitempty"`

	// id of the device on which volume instance is created
	DeviceID string `json:"deviceId,omitempty"`

	// System defined universally unique Id of the volume instance.
	// Read Only: true
	// Pattern: [a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}
	ID string `json:"id,omitempty"`

	// name of the image
	Image string `json:"image,omitempty"`

	// flag to create implicit volumes
	Implicit string `json:"implicit,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// flag to enable the volume to be attached to multiple app instances
	Multiattach bool `json:"multiattach,omitempty"`

	// User defined name of the volume instance, unique across the enterprise. Once object is created, name can’t be changed.
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	Name string `json:"name,omitempty"`

	// id of the project in which the volume instance is created
	ProjectID string `json:"projectId,omitempty"`

	// Purge Counter information
	Purge *ZedCloudOpsCmd `json:"purge,omitempty"`

	// system defined Revision info of the object
	// Read Only: true
	Revision *ObjectRevision `json:"revision,omitempty"`

	// size of volume
	SizeBytes uint64 `json:"sizeBytes,omitempty"`

	// User defined title of the volume instance. Title can be changed at any time.
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+
	Title string `json:"title,omitempty"`

	// type of Volume Instance
	Type *VolumeInstanceType `json:"type,omitempty"`
}

// Validate validates this vol inst config
func (m *VolInstConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessmode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolInstConfig) validateAccessmode(formats strfmt.Registry) error {
	if swag.IsZero(m.Accessmode) { // not required
		return nil
	}

	if m.Accessmode != nil {
		if err := m.Accessmode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessmode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessmode")
			}
			return err
		}
	}

	return nil
}

func (m *VolInstConfig) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 256); err != nil {
		return err
	}

	return nil
}

func (m *VolInstConfig) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.Pattern("id", "body", m.ID, `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}`); err != nil {
		return err
	}

	return nil
}

func (m *VolInstConfig) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", m.Name, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", m.Name, 256); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", m.Name, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *VolInstConfig) validatePurge(formats strfmt.Registry) error {
	if swag.IsZero(m.Purge) { // not required
		return nil
	}

	if m.Purge != nil {
		if err := m.Purge.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("purge")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("purge")
			}
			return err
		}
	}

	return nil
}

func (m *VolInstConfig) validateRevision(formats strfmt.Registry) error {
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

func (m *VolInstConfig) validateTitle(formats strfmt.Registry) error {
	if swag.IsZero(m.Title) { // not required
		return nil
	}

	if err := validate.MinLength("title", "body", m.Title, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", m.Title, 256); err != nil {
		return err
	}

	if err := validate.Pattern("title", "body", m.Title, `[a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+`); err != nil {
		return err
	}

	return nil
}

func (m *VolInstConfig) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this vol inst config based on the context it is used
func (m *VolInstConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessmode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePurge(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolInstConfig) contextValidateAccessmode(ctx context.Context, formats strfmt.Registry) error {

	if m.Accessmode != nil {
		if err := m.Accessmode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessmode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessmode")
			}
			return err
		}
	}

	return nil
}

func (m *VolInstConfig) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *VolInstConfig) contextValidatePurge(ctx context.Context, formats strfmt.Registry) error {

	if m.Purge != nil {
		if err := m.Purge.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("purge")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("purge")
			}
			return err
		}
	}

	return nil
}

func (m *VolInstConfig) contextValidateRevision(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VolInstConfig) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VolInstConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolInstConfig) UnmarshalBinary(b []byte) error {
	var res VolInstConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
