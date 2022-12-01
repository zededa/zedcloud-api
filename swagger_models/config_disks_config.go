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

// ConfigDisksConfig DisksConfig is a configuration of disks
// We expect information about disks to be filled and will try to adjust disks states accordingly
// All disks defined in disks field expected to have array type defined in array_type
// To support nested topologies we can use children field
//
// For example to use stripe of two pairs of mirrored disks we should define
// DisksConfig without disks with array_type DISKS_ARRAY_TYPE_RAID0
// with two children with properly defined disks inside and with array_type DISKS_ARRAY_TYPE_RAID1
// and empty children
//
// swagger:model configDisksConfig
type ConfigDisksConfig struct {

	// array type
	ArrayType *ConfigDisksArrayType `json:"arrayType,omitempty"`

	// children
	Children []*ConfigDisksConfig `json:"children"`

	// disks
	Disks []*ConfigDiskConfig `json:"disks"`
}

// Validate validates this config disks config
func (m *ConfigDisksConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArrayType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChildren(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigDisksConfig) validateArrayType(formats strfmt.Registry) error {
	if swag.IsZero(m.ArrayType) { // not required
		return nil
	}

	if m.ArrayType != nil {
		if err := m.ArrayType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("arrayType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("arrayType")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigDisksConfig) validateChildren(formats strfmt.Registry) error {
	if swag.IsZero(m.Children) { // not required
		return nil
	}

	for i := 0; i < len(m.Children); i++ {
		if swag.IsZero(m.Children[i]) { // not required
			continue
		}

		if m.Children[i] != nil {
			if err := m.Children[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("children" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("children" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigDisksConfig) validateDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.Disks) { // not required
		return nil
	}

	for i := 0; i < len(m.Disks); i++ {
		if swag.IsZero(m.Disks[i]) { // not required
			continue
		}

		if m.Disks[i] != nil {
			if err := m.Disks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this config disks config based on the context it is used
func (m *ConfigDisksConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArrayType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChildren(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigDisksConfig) contextValidateArrayType(ctx context.Context, formats strfmt.Registry) error {

	if m.ArrayType != nil {
		if err := m.ArrayType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("arrayType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("arrayType")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigDisksConfig) contextValidateChildren(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Children); i++ {

		if m.Children[i] != nil {
			if err := m.Children[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("children" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("children" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigDisksConfig) contextValidateDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Disks); i++ {

		if m.Disks[i] != nil {
			if err := m.Disks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigDisksConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigDisksConfig) UnmarshalBinary(b []byte) error {
	var res ConfigDisksConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
