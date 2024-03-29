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

// DetailedUsers detailed users
//
// swagger:model DetailedUsers
type DetailedUsers struct {

	// List of users
	List []*DetailedUser `json:"list"`

	// Page details of the filtered records
	Next *Cursor `json:"next,omitempty"`

	// User distribution by role
	SummaryByRoleDistribution *Summary `json:"summaryByRoleDistribution,omitempty"`

	// Summary of filtered users
	SummaryByState *Summary `json:"summaryByState,omitempty"`

	// Count of online/offline users
	SummaryByUserActivity *Summary `json:"summaryByUserActivity,omitempty"`
}

// Validate validates this detailed users
func (m *DetailedUsers) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummaryByRoleDistribution(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummaryByState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummaryByUserActivity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DetailedUsers) validateList(formats strfmt.Registry) error {
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

func (m *DetailedUsers) validateNext(formats strfmt.Registry) error {
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

func (m *DetailedUsers) validateSummaryByRoleDistribution(formats strfmt.Registry) error {
	if swag.IsZero(m.SummaryByRoleDistribution) { // not required
		return nil
	}

	if m.SummaryByRoleDistribution != nil {
		if err := m.SummaryByRoleDistribution.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summaryByRoleDistribution")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summaryByRoleDistribution")
			}
			return err
		}
	}

	return nil
}

func (m *DetailedUsers) validateSummaryByState(formats strfmt.Registry) error {
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

func (m *DetailedUsers) validateSummaryByUserActivity(formats strfmt.Registry) error {
	if swag.IsZero(m.SummaryByUserActivity) { // not required
		return nil
	}

	if m.SummaryByUserActivity != nil {
		if err := m.SummaryByUserActivity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summaryByUserActivity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summaryByUserActivity")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this detailed users based on the context it is used
func (m *DetailedUsers) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSummaryByRoleDistribution(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSummaryByState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSummaryByUserActivity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DetailedUsers) contextValidateList(ctx context.Context, formats strfmt.Registry) error {

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

func (m *DetailedUsers) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

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

func (m *DetailedUsers) contextValidateSummaryByRoleDistribution(ctx context.Context, formats strfmt.Registry) error {

	if m.SummaryByRoleDistribution != nil {

		if swag.IsZero(m.SummaryByRoleDistribution) { // not required
			return nil
		}

		if err := m.SummaryByRoleDistribution.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summaryByRoleDistribution")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summaryByRoleDistribution")
			}
			return err
		}
	}

	return nil
}

func (m *DetailedUsers) contextValidateSummaryByState(ctx context.Context, formats strfmt.Registry) error {

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

func (m *DetailedUsers) contextValidateSummaryByUserActivity(ctx context.Context, formats strfmt.Registry) error {

	if m.SummaryByUserActivity != nil {

		if swag.IsZero(m.SummaryByUserActivity) { // not required
			return nil
		}

		if err := m.SummaryByUserActivity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summaryByUserActivity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summaryByUserActivity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DetailedUsers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DetailedUsers) UnmarshalBinary(b []byte) error {
	var res DetailedUsers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
