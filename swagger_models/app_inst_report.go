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

// AppInstReport Get appInst reports
//
// Get summary reports of appInst objects for self/child enterprise
//
// swagger:model AppInstReport
type AppInstReport struct {

	// Enterprise appInst report
	AppInstSummaryReport *AppInstSummaryReport `json:"appInstSummaryReport,omitempty"`

	// Enterprise id for which we want to get summary report for all objects
	// Pattern: [0-9A-Za-z-]+
	EntpID string `json:"entpId,omitempty"`

	// Error while fetching report for enterprise, if any
	Error string `json:"error,omitempty"`
}

// Validate validates this app inst report
func (m *AppInstReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppInstSummaryReport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntpID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstReport) validateAppInstSummaryReport(formats strfmt.Registry) error {
	if swag.IsZero(m.AppInstSummaryReport) { // not required
		return nil
	}

	if m.AppInstSummaryReport != nil {
		if err := m.AppInstSummaryReport.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appInstSummaryReport")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appInstSummaryReport")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstReport) validateEntpID(formats strfmt.Registry) error {
	if swag.IsZero(m.EntpID) { // not required
		return nil
	}

	if err := validate.Pattern("entpId", "body", m.EntpID, `[0-9A-Za-z-]+`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this app inst report based on the context it is used
func (m *AppInstReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppInstSummaryReport(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstReport) contextValidateAppInstSummaryReport(ctx context.Context, formats strfmt.Registry) error {

	if m.AppInstSummaryReport != nil {
		if err := m.AppInstSummaryReport.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appInstSummaryReport")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appInstSummaryReport")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppInstReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppInstReport) UnmarshalBinary(b []byte) error {
	var res AppInstReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
