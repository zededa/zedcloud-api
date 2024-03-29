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

// TagStatusListMsg tag status list msg
//
// swagger:model TagStatusListMsg
type TagStatusListMsg struct {

	// List of filtered resource group records
	List []*TagStatusMsg `json:"list"`

	// Responded page details of filtered records
	Next *Cursor `json:"next,omitempty"`

	// Summary of filtered resource group records
	SummaryByState *Summary `json:"summaryByState,omitempty"`

	// Summary of filtered resource group records
	SummaryByType *Summary `json:"summaryByType,omitempty"`
}

// Validate validates this tag status list msg
func (m *TagStatusListMsg) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummaryByState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummaryByType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TagStatusListMsg) validateList(formats strfmt.Registry) error {
	if swag.IsZero(m.List) { // not required
		return nil
	}

	for i := 0; i < len(m.List); i++ {
		if swag.IsZero(m.List[i]) { // not required
			continue
		}

		if m.List[i] != nil {
			if err := m.List[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("list" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TagStatusListMsg) validateNext(formats strfmt.Registry) error {
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

func (m *TagStatusListMsg) validateSummaryByState(formats strfmt.Registry) error {
	if swag.IsZero(m.SummaryByState) { // not required
		return nil
	}

	if m.SummaryByState != nil {
		if err := m.SummaryByState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summaryByState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summaryByState")
			}
			return err
		}
	}

	return nil
}

func (m *TagStatusListMsg) validateSummaryByType(formats strfmt.Registry) error {
	if swag.IsZero(m.SummaryByType) { // not required
		return nil
	}

	if m.SummaryByType != nil {
		if err := m.SummaryByType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summaryByType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summaryByType")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this tag status list msg based on the context it is used
func (m *TagStatusListMsg) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSummaryByState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSummaryByType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TagStatusListMsg) contextValidateList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.List); i++ {

		if m.List[i] != nil {

			if swag.IsZero(m.List[i]) { // not required
				return nil
			}

			if err := m.List[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("list" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TagStatusListMsg) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

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

func (m *TagStatusListMsg) contextValidateSummaryByState(ctx context.Context, formats strfmt.Registry) error {

	if m.SummaryByState != nil {

		if swag.IsZero(m.SummaryByState) { // not required
			return nil
		}

		if err := m.SummaryByState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summaryByState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summaryByState")
			}
			return err
		}
	}

	return nil
}

func (m *TagStatusListMsg) contextValidateSummaryByType(ctx context.Context, formats strfmt.Registry) error {

	if m.SummaryByType != nil {

		if swag.IsZero(m.SummaryByType) { // not required
			return nil
		}

		if err := m.SummaryByType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summaryByType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summaryByType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TagStatusListMsg) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TagStatusListMsg) UnmarshalBinary(b []byte) error {
	var res TagStatusListMsg
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
