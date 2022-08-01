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

// IoMember ioMember  payload details
//
//  ioMember consists of list of various networking attributes like physical label, physical address etc  associated with SysModel
// Example: {"logicallabel":"eth0","phylabel":"eth0","ztype":1}
//
// swagger:model ioMember
type IoMember struct {

	// Assign Group
	// Required: true
	Assigngrp *string `json:"assigngrp"`

	// physical and logical attributes
	//
	// attributes
	Cbattr map[string]string `json:"cbattr,omitempty"`

	// cost of using this ioMember. Default is 0.
	// Required: true
	// Maximum: 255
	Cost *int64 `json:"cost"`

	// Logical Label
	// Required: true
	Logicallabel *string `json:"logicallabel"`

	// Map of Physical Addresses
	// Required: true
	Phyaddrs map[string]string `json:"phyaddrs"`

	// Physical Label
	// Required: true
	Phylabel *string `json:"phylabel"`

	// Adopter Usage
	Usage *AdapterUsage `json:"usage,omitempty"`

	// usagePolicy
	UsagePolicy map[string]bool `json:"usagePolicy,omitempty"`

	// Z Type
	// Required: true
	Ztype *IoType `json:"ztype"`
}

// Validate validates this io member
func (m *IoMember) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssigngrp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogicallabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhyaddrs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhylabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZtype(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoMember) validateAssigngrp(formats strfmt.Registry) error {

	if err := validate.Required("assigngrp", "body", m.Assigngrp); err != nil {
		return err
	}

	return nil
}

func (m *IoMember) validateCost(formats strfmt.Registry) error {

	if err := validate.Required("cost", "body", m.Cost); err != nil {
		return err
	}

	if err := validate.MaximumInt("cost", "body", *m.Cost, 255, false); err != nil {
		return err
	}

	return nil
}

func (m *IoMember) validateLogicallabel(formats strfmt.Registry) error {

	if err := validate.Required("logicallabel", "body", m.Logicallabel); err != nil {
		return err
	}

	return nil
}

func (m *IoMember) validatePhyaddrs(formats strfmt.Registry) error {

	if err := validate.Required("phyaddrs", "body", m.Phyaddrs); err != nil {
		return err
	}

	return nil
}

func (m *IoMember) validatePhylabel(formats strfmt.Registry) error {

	if err := validate.Required("phylabel", "body", m.Phylabel); err != nil {
		return err
	}

	return nil
}

func (m *IoMember) validateUsage(formats strfmt.Registry) error {
	if swag.IsZero(m.Usage) { // not required
		return nil
	}

	if m.Usage != nil {
		if err := m.Usage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usage")
			}
			return err
		}
	}

	return nil
}

func (m *IoMember) validateZtype(formats strfmt.Registry) error {

	if err := validate.Required("ztype", "body", m.Ztype); err != nil {
		return err
	}

	if err := validate.Required("ztype", "body", m.Ztype); err != nil {
		return err
	}

	if m.Ztype != nil {
		if err := m.Ztype.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ztype")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ztype")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this io member based on the context it is used
func (m *IoMember) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateZtype(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoMember) contextValidateUsage(ctx context.Context, formats strfmt.Registry) error {

	if m.Usage != nil {
		if err := m.Usage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usage")
			}
			return err
		}
	}

	return nil
}

func (m *IoMember) contextValidateZtype(ctx context.Context, formats strfmt.Registry) error {

	if m.Ztype != nil {
		if err := m.Ztype.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ztype")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ztype")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoMember) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoMember) UnmarshalBinary(b []byte) error {
	var res IoMember
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
