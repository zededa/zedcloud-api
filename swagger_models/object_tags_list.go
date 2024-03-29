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

// ObjectTagsList Object tag list records
//
// # Returned list of Object tags matched by filter criteria
//
// swagger:model ObjectTagsList
type ObjectTagsList struct {

	// Returned record page.
	Next *Cursor `json:"next,omitempty"`

	// Details of Object tags list records.
	ObjectTags []*ObjectTag `json:"objectTags"`
}

// Validate validates this object tags list
func (m *ObjectTagsList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectTagsList) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

func (m *ObjectTagsList) validateObjectTags(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectTags) { // not required
		return nil
	}

	for i := 0; i < len(m.ObjectTags); i++ {
		if swag.IsZero(m.ObjectTags[i]) { // not required
			continue
		}

		if m.ObjectTags[i] != nil {
			if err := m.ObjectTags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objectTags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objectTags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this object tags list based on the context it is used
func (m *ObjectTagsList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateObjectTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectTagsList) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if m.Next != nil {

		if swag.IsZero(m.Next) { // not required
			return nil
		}

		if err := m.Next.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

func (m *ObjectTagsList) contextValidateObjectTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ObjectTags); i++ {

		if m.ObjectTags[i] != nil {

			if swag.IsZero(m.ObjectTags[i]) { // not required
				return nil
			}

			if err := m.ObjectTags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objectTags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objectTags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ObjectTagsList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectTagsList) UnmarshalBinary(b []byte) error {
	var res ObjectTagsList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
