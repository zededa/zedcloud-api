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

// VolInstStatusListMsg vol inst status list msg
//
// swagger:model VolInstStatusListMsg
type VolInstStatusListMsg struct {

	// list
	List []*VolInstStatusSummaryMsg `json:"list"`

	// next
	Next *Cursor `json:"next,omitempty"`

	// summary by state
	SummaryByState *Summary `json:"summaryByState,omitempty"`

	// summary by type
	SummaryByType *Summary `json:"summaryByType,omitempty"`

	// total count
	TotalCount int64 `json:"totalCount,omitempty"`
}

// Validate validates this vol inst status list msg
func (m *VolInstStatusListMsg) Validate(formats strfmt.Registry) error {
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

func (m *VolInstStatusListMsg) validateList(formats strfmt.Registry) error {
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

func (m *VolInstStatusListMsg) validateNext(formats strfmt.Registry) error {
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

func (m *VolInstStatusListMsg) validateSummaryByState(formats strfmt.Registry) error {
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

func (m *VolInstStatusListMsg) validateSummaryByType(formats strfmt.Registry) error {
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

// ContextValidate validate this vol inst status list msg based on the context it is used
func (m *VolInstStatusListMsg) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *VolInstStatusListMsg) contextValidateList(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VolInstStatusListMsg) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VolInstStatusListMsg) contextValidateSummaryByState(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VolInstStatusListMsg) contextValidateSummaryByType(ctx context.Context, formats strfmt.Registry) error {

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
func (m *VolInstStatusListMsg) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolInstStatusListMsg) UnmarshalBinary(b []byte) error {
	var res VolInstStatusListMsg
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
