// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

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
	"github.com/go-openapi/validate"
)

// AppConfig Edge application configuration for App policy
//
// Edge application configuration for an App policy defines configuration details of the Edge application to be installed to all Edge node(s) under the App policy.
//
// swagger:model AppConfig
type AppConfig struct {

	// User defined name of the edge app, unique across the enterprise. Once app name is created, name can’t be changed
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	AppID string `json:"appId,omitempty"`

	// Current version of the attached bundle
	AppVersion string `json:"appVersion,omitempty"`

	// user defined cpus for bundle
	// Required: true
	Cpus *int64 `json:"cpus"`

	// Detailed description of the edge application
	// Max Length: 256
	Description string `json:"description,omitempty"`

	// user defined drives
	// Required: true
	// Read Only: true
	Drives int64 `json:"drives"`

	// System defined universally unique Id of the edge application
	// Read Only: true
	// Pattern: [a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}
	ID string `json:"id,omitempty"`

	// application interfaces
	Interfaces []*AppInterface `json:"interfaces"`

	// Manifest data
	// Required: true
	ManifestJSON *VMManifest `json:"manifestJSON"`

	// user defined memory for bundle
	// Required: true
	Memory *int64 `json:"memory"`

	// User defined name of the edge application, unique across the enterprise. Once object is created, name can’t be changed
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	Name *string `json:"name"`

	// User provided name part  for the auto deployed app
	NameAppPart string `json:"nameAppPart,omitempty"`

	// User provided name part  for the auto deployed app
	NameProjectPart string `json:"nameProjectPart,omitempty"`

	// app naming scheme
	NamingScheme *AppNamingScheme `json:"namingScheme,omitempty"`

	// user defined network options
	// Required: true
	Networks *int64 `json:"networks"`

	// origin of object
	// Required: true
	OriginType *Origin `json:"originType"`

	// origin and parent related details
	ParentDetail *ObjectParentDetail `json:"parentDetail,omitempty"`

	// start delay is the time in seconds EVE should wait after boot before starting the application instance
	StartDelayInSeconds int64 `json:"startDelayInSeconds,omitempty"`

	// user defined storage for bundle
	Storage int64 `json:"storage,omitempty"`

	// Tags are name/value pairs that enable you to categorize resources. Tag names are case insensitive with max_length 512 and min_length 3. Tag values are case sensitive with max_length 256 and min_length 3.
	Tags map[string]string `json:"tags,omitempty"`

	// User defined title of the edge application. Title can be changed at any time
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: ^[a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+$
	Title *string `json:"title"`
}

// Validate validates this app config
func (m *AppConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCpus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDrives(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterfaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManifestJSON(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamingScheme(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentDetail(formats); err != nil {
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

func (m *AppConfig) validateAppID(formats strfmt.Registry) error {
	if swag.IsZero(m.AppID) { // not required
		return nil
	}

	if err := validate.MinLength("appId", "body", m.AppID, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("appId", "body", m.AppID, 256); err != nil {
		return err
	}

	if err := validate.Pattern("appId", "body", m.AppID, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *AppConfig) validateCpus(formats strfmt.Registry) error {

	if err := validate.Required("cpus", "body", m.Cpus); err != nil {
		return err
	}

	return nil
}

func (m *AppConfig) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 256); err != nil {
		return err
	}

	return nil
}

func (m *AppConfig) validateDrives(formats strfmt.Registry) error {

	if err := validate.Required("drives", "body", int64(m.Drives)); err != nil {
		return err
	}

	return nil
}

func (m *AppConfig) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.Pattern("id", "body", m.ID, `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}`); err != nil {
		return err
	}

	return nil
}

func (m *AppConfig) validateInterfaces(formats strfmt.Registry) error {
	if swag.IsZero(m.Interfaces) { // not required
		return nil
	}

	for i := 0; i < len(m.Interfaces); i++ {
		if swag.IsZero(m.Interfaces[i]) { // not required
			continue
		}

		if m.Interfaces[i] != nil {
			if err := m.Interfaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("interfaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppConfig) validateManifestJSON(formats strfmt.Registry) error {

	if err := validate.Required("manifestJSON", "body", m.ManifestJSON); err != nil {
		return err
	}

	if m.ManifestJSON != nil {
		if err := m.ManifestJSON.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("manifestJSON")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("manifestJSON")
			}
			return err
		}
	}

	return nil
}

func (m *AppConfig) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("memory", "body", m.Memory); err != nil {
		return err
	}

	return nil
}

func (m *AppConfig) validateName(formats strfmt.Registry) error {

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

func (m *AppConfig) validateNamingScheme(formats strfmt.Registry) error {
	if swag.IsZero(m.NamingScheme) { // not required
		return nil
	}

	if m.NamingScheme != nil {
		if err := m.NamingScheme.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namingScheme")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namingScheme")
			}
			return err
		}
	}

	return nil
}

func (m *AppConfig) validateNetworks(formats strfmt.Registry) error {

	if err := validate.Required("networks", "body", m.Networks); err != nil {
		return err
	}

	return nil
}

func (m *AppConfig) validateOriginType(formats strfmt.Registry) error {

	if err := validate.Required("originType", "body", m.OriginType); err != nil {
		return err
	}

	if err := validate.Required("originType", "body", m.OriginType); err != nil {
		return err
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

func (m *AppConfig) validateParentDetail(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentDetail) { // not required
		return nil
	}

	if m.ParentDetail != nil {
		if err := m.ParentDetail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parentDetail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parentDetail")
			}
			return err
		}
	}

	return nil
}

func (m *AppConfig) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	if err := validate.MinLength("title", "body", *m.Title, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", *m.Title, 256); err != nil {
		return err
	}

	if err := validate.Pattern("title", "body", *m.Title, `^[a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this app config based on the context it is used
func (m *AppConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDrives(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterfaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateManifestJSON(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNamingScheme(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOriginType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParentDetail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppConfig) contextValidateDrives(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "drives", "body", int64(m.Drives)); err != nil {
		return err
	}

	return nil
}

func (m *AppConfig) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *AppConfig) contextValidateInterfaces(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Interfaces); i++ {

		if m.Interfaces[i] != nil {
			if err := m.Interfaces[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("interfaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppConfig) contextValidateManifestJSON(ctx context.Context, formats strfmt.Registry) error {

	if m.ManifestJSON != nil {
		if err := m.ManifestJSON.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("manifestJSON")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("manifestJSON")
			}
			return err
		}
	}

	return nil
}

func (m *AppConfig) contextValidateNamingScheme(ctx context.Context, formats strfmt.Registry) error {

	if m.NamingScheme != nil {
		if err := m.NamingScheme.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namingScheme")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("namingScheme")
			}
			return err
		}
	}

	return nil
}

func (m *AppConfig) contextValidateOriginType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AppConfig) contextValidateParentDetail(ctx context.Context, formats strfmt.Registry) error {

	if m.ParentDetail != nil {
		if err := m.ParentDetail.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parentDetail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parentDetail")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppConfig) UnmarshalBinary(b []byte) error {
	var res AppConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
