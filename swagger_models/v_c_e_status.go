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

// VCEStatus status of object on velocloud.
//
// swagger:model VCEStatus
type VCEStatus struct {

	// vce status detail
	VceStatusDetail *VCEStatusDetail `json:"vceStatusDetail,omitempty"`
}

// Validate validates this v c e status
func (m *VCEStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVceStatusDetail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VCEStatus) validateVceStatusDetail(formats strfmt.Registry) error {
	if swag.IsZero(m.VceStatusDetail) { // not required
		return nil
	}

	if m.VceStatusDetail != nil {
		if err := m.VceStatusDetail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vceStatusDetail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vceStatusDetail")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v c e status based on the context it is used
func (m *VCEStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVceStatusDetail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VCEStatus) contextValidateVceStatusDetail(ctx context.Context, formats strfmt.Registry) error {

	if m.VceStatusDetail != nil {
		if err := m.VceStatusDetail.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vceStatusDetail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vceStatusDetail")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VCEStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VCEStatus) UnmarshalBinary(b []byte) error {
	var res VCEStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
