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
	"github.com/go-openapi/validate"
)

// ProjectReport Get project reports
//
// Get summary reports of project objects for self/child enterprise
//
// swagger:model ProjectReport
type ProjectReport struct {

	// Enterprise id for which we want to get summary report for all objects
	// Pattern: [0-9A-Za-z-]+
	EntpID string `json:"entpId,omitempty"`

	// Error while fetching report for enterprise, if any
	Error string `json:"error,omitempty"`

	// cursor next
	Next *Cursor `json:"next,omitempty"`

	// Enterprise project report
	ProjectSummaryReport *ProjectSummaryReport `json:"projectSummaryReport,omitempty"`
}

// Validate validates this project report
func (m *ProjectReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntpID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectSummaryReport(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectReport) validateEntpID(formats strfmt.Registry) error {
	if swag.IsZero(m.EntpID) { // not required
		return nil
	}

	if err := validate.Pattern("entpId", "body", m.EntpID, `[0-9A-Za-z-]+`); err != nil {
		return err
	}

	return nil
}

func (m *ProjectReport) validateNext(formats strfmt.Registry) error {
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

func (m *ProjectReport) validateProjectSummaryReport(formats strfmt.Registry) error {
	if swag.IsZero(m.ProjectSummaryReport) { // not required
		return nil
	}

	if m.ProjectSummaryReport != nil {
		if err := m.ProjectSummaryReport.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("projectSummaryReport")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("projectSummaryReport")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this project report based on the context it is used
func (m *ProjectReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProjectSummaryReport(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectReport) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if m.Next != nil {
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

func (m *ProjectReport) contextValidateProjectSummaryReport(ctx context.Context, formats strfmt.Registry) error {

	if m.ProjectSummaryReport != nil {
		if err := m.ProjectSummaryReport.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("projectSummaryReport")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("projectSummaryReport")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectReport) UnmarshalBinary(b []byte) error {
	var res ProjectReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
