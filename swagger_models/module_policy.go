// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

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
	"github.com/go-openapi/validate"
)

// ModulePolicy Module policy body detail
//
// list of modules that will be provisioned on all the devices of the project to which this policy is attached
//
// swagger:model ModulePolicy
type ModulePolicy struct {

	// etag for deployment
	Etag string `json:"Etag,omitempty"`

	// list of app details that will be provisioned on all the devices of the project to which this policy is attached
	// Required: true
	Apps []*AppConfig `json:"apps"`

	// unique id for deployment
	// Read Only: true
	// Pattern: [0-9-a-z-]+
	ID string `json:"id,omitempty"`

	// Mapping of label variable keys and value
	Labels map[string]string `json:"labels,omitempty"`

	// custom metrics for deployment
	Metrics *MetricsDetail `json:"metrics,omitempty"`

	// deployment priority of module manifest
	// Required: true
	Priority *int64 `json:"priority"`

	// Mapping of routes variable keys and value
	Routes map[string]string `json:"routes,omitempty"`

	// target condition for deployment that matches single device or group of devices
	TargetCondition string `json:"targetCondition,omitempty"`

	// target condition for deployment that matches single device or group of devices
	TargetConditionNew map[string]string `json:"targetConditionNew,omitempty"`
}

// Validate validates this module policy
func (m *ModulePolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModulePolicy) validateApps(formats strfmt.Registry) error {

	if err := validate.Required("apps", "body", m.Apps); err != nil {
		return err
	}

	for i := 0; i < len(m.Apps); i++ {
		if swag.IsZero(m.Apps[i]) { // not required
			continue
		}

		if m.Apps[i] != nil {
			if err := m.Apps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("apps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModulePolicy) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.Pattern("id", "body", m.ID, `[0-9-a-z-]+`); err != nil {
		return err
	}

	return nil
}

func (m *ModulePolicy) validateMetrics(formats strfmt.Registry) error {
	if swag.IsZero(m.Metrics) { // not required
		return nil
	}

	if m.Metrics != nil {
		if err := m.Metrics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metrics")
			}
			return err
		}
	}

	return nil
}

func (m *ModulePolicy) validatePriority(formats strfmt.Registry) error {

	if err := validate.Required("priority", "body", m.Priority); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this module policy based on the context it is used
func (m *ModulePolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetrics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModulePolicy) contextValidateApps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Apps); i++ {

		if m.Apps[i] != nil {
			if err := m.Apps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("apps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModulePolicy) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ModulePolicy) contextValidateMetrics(ctx context.Context, formats strfmt.Registry) error {

	if m.Metrics != nil {
		if err := m.Metrics.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metrics")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModulePolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModulePolicy) UnmarshalBinary(b []byte) error {
	var res ModulePolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
