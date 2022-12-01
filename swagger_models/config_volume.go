// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
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

// ConfigVolume The Volume describes a storage volume which should exist on the device.
// This can currently either be blank or created from a ContentTree
// If maxSizeBytes is zero it means unlimited by the controller. In that
// case EVE needs to determine how much space it can assign and limit any
// downloaded ContentTree and the created volume based on that calculated size.
//
// swagger:model configVolume
type ConfigVolume struct {

	// clear text
	ClearText bool `json:"clearText,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// change in generationCount indicates that the mutated volume needs to be
	// purged and built from scratch. This is a generalization of the purge
	// command for an application instance
	GenerationCount string `json:"generationCount,omitempty"`

	// miscellaneous attributes of the Volume
	// maxSizeBytes Used for capping resource consumption in EVE.
	// maxSizeBytes indicates the maximum size of the volume.
	Maxsizebytes string `json:"maxsizebytes,omitempty"`

	// origin of the volume content.
	Origin *ConfigVolumeContentOrigin `json:"origin,omitempty"`

	// describes all the different ways how this Volume can
	// be offered to Tasks
	Protocols []*ConfigVolumeAccessProtocols `json:"protocols"`

	// readonly
	Readonly bool `json:"readonly,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this config volume
func (m *ConfigVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrigin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocols(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigVolume) validateOrigin(formats strfmt.Registry) error {
	if swag.IsZero(m.Origin) { // not required
		return nil
	}

	if m.Origin != nil {
		if err := m.Origin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("origin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("origin")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigVolume) validateProtocols(formats strfmt.Registry) error {
	if swag.IsZero(m.Protocols) { // not required
		return nil
	}

	for i := 0; i < len(m.Protocols); i++ {
		if swag.IsZero(m.Protocols[i]) { // not required
			continue
		}

		if m.Protocols[i] != nil {
			if err := m.Protocols[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("protocols" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("protocols" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this config volume based on the context it is used
func (m *ConfigVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrigin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtocols(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigVolume) contextValidateOrigin(ctx context.Context, formats strfmt.Registry) error {

	if m.Origin != nil {
		if err := m.Origin.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("origin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("origin")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigVolume) contextValidateProtocols(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Protocols); i++ {

		if m.Protocols[i] != nil {
			if err := m.Protocols[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("protocols" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("protocols" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigVolume) UnmarshalBinary(b []byte) error {
	var res ConfigVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
