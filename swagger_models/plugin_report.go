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

// PluginReport Get plugin reports
//
// Get summary reports of plugin objects for self/child enterprise
//
// swagger:model PluginReport
type PluginReport struct {

	// Enterprise id for which we want to get summary report for all objects
	// Pattern: [0-9A-Za-z-]+
	EntpID string `json:"entpId,omitempty"`

	// Error while fetching report for enterprise, if any
	Error string `json:"error,omitempty"`

	// Enterprise plugin report
	PluginSummaryReport *PluginSummaryReport `json:"pluginSummaryReport,omitempty"`
}

// Validate validates this plugin report
func (m *PluginReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntpID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePluginSummaryReport(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginReport) validateEntpID(formats strfmt.Registry) error {
	if swag.IsZero(m.EntpID) { // not required
		return nil
	}

	if err := validate.Pattern("entpId", "body", m.EntpID, `[0-9A-Za-z-]+`); err != nil {
		return err
	}

	return nil
}

func (m *PluginReport) validatePluginSummaryReport(formats strfmt.Registry) error {
	if swag.IsZero(m.PluginSummaryReport) { // not required
		return nil
	}

	if m.PluginSummaryReport != nil {
		if err := m.PluginSummaryReport.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pluginSummaryReport")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pluginSummaryReport")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this plugin report based on the context it is used
func (m *PluginReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePluginSummaryReport(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginReport) contextValidatePluginSummaryReport(ctx context.Context, formats strfmt.Registry) error {

	if m.PluginSummaryReport != nil {
		if err := m.PluginSummaryReport.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pluginSummaryReport")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pluginSummaryReport")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PluginReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginReport) UnmarshalBinary(b []byte) error {
	var res PluginReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}