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

// EveconfigImage XXX the Image will be deprecated and we will use ContentTree instead
//
// swagger:model eveconfigImage
type EveconfigImage struct {

	// ds Id
	DsID string `json:"dsId,omitempty"`

	// iformat
	Iformat *ConfigFormat `json:"iformat,omitempty"`

	// it could be relative path/name as well; appended to the datastore dpath
	Name string `json:"name,omitempty"`

	// sha256
	Sha256 string `json:"sha256,omitempty"`

	// if its signed image
	Siginfo *ConfigSignatureInfo `json:"siginfo,omitempty"`

	// sizeBytes indicates the maximum download size of an image.
	// A value of 0 will indicate the unlimited download.
	SizeBytes string `json:"sizeBytes,omitempty"`

	// uuidandversion
	Uuidandversion *ConfigUUIDandVersion `json:"uuidandversion,omitempty"`
}

// Validate validates this eveconfig image
func (m *EveconfigImage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIformat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSiginfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUuidandversion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EveconfigImage) validateIformat(formats strfmt.Registry) error {
	if swag.IsZero(m.Iformat) { // not required
		return nil
	}

	if m.Iformat != nil {
		if err := m.Iformat.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iformat")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iformat")
			}
			return err
		}
	}

	return nil
}

func (m *EveconfigImage) validateSiginfo(formats strfmt.Registry) error {
	if swag.IsZero(m.Siginfo) { // not required
		return nil
	}

	if m.Siginfo != nil {
		if err := m.Siginfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("siginfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("siginfo")
			}
			return err
		}
	}

	return nil
}

func (m *EveconfigImage) validateUuidandversion(formats strfmt.Registry) error {
	if swag.IsZero(m.Uuidandversion) { // not required
		return nil
	}

	if m.Uuidandversion != nil {
		if err := m.Uuidandversion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uuidandversion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uuidandversion")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this eveconfig image based on the context it is used
func (m *EveconfigImage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIformat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSiginfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUuidandversion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EveconfigImage) contextValidateIformat(ctx context.Context, formats strfmt.Registry) error {

	if m.Iformat != nil {
		if err := m.Iformat.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iformat")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iformat")
			}
			return err
		}
	}

	return nil
}

func (m *EveconfigImage) contextValidateSiginfo(ctx context.Context, formats strfmt.Registry) error {

	if m.Siginfo != nil {
		if err := m.Siginfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("siginfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("siginfo")
			}
			return err
		}
	}

	return nil
}

func (m *EveconfigImage) contextValidateUuidandversion(ctx context.Context, formats strfmt.Registry) error {

	if m.Uuidandversion != nil {
		if err := m.Uuidandversion.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uuidandversion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uuidandversion")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EveconfigImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EveconfigImage) UnmarshalBinary(b []byte) error {
	var res EveconfigImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
