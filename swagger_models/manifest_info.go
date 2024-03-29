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
)

// ManifestInfo Transition action information
//
// # Details about transition parameters and indication for cause of the transition and recommendation for the transition action
//
// swagger:model ManifestInfo
type ManifestInfo struct {

	// Current version of edge application being used
	BundleVersion string `json:"bundleVersion,omitempty"`

	// Details for recommended transition action
	Details *TransDetails `json:"details,omitempty"`

	// Next version of edge application available
	NextBundleVersion string `json:"nextBundleVersion,omitempty"`

	// Common and Custom parameters for deciding on transition actions
	Params map[string]string `json:"params,omitempty"`

	// Recommended transition action
	TransitionAction *InstanceTransitionAction `json:"transitionAction,omitempty"`
}

// Validate validates this manifest info
func (m *ManifestInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransitionAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManifestInfo) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.Details) { // not required
		return nil
	}

	if m.Details != nil {
		if err := m.Details.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("details")
			}
			return err
		}
	}

	return nil
}

func (m *ManifestInfo) validateTransitionAction(formats strfmt.Registry) error {
	if swag.IsZero(m.TransitionAction) { // not required
		return nil
	}

	if m.TransitionAction != nil {
		if err := m.TransitionAction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transitionAction")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("transitionAction")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this manifest info based on the context it is used
func (m *ManifestInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransitionAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManifestInfo) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.Details != nil {

		if swag.IsZero(m.Details) { // not required
			return nil
		}

		if err := m.Details.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("details")
			}
			return err
		}
	}

	return nil
}

func (m *ManifestInfo) contextValidateTransitionAction(ctx context.Context, formats strfmt.Registry) error {

	if m.TransitionAction != nil {

		if swag.IsZero(m.TransitionAction) { // not required
			return nil
		}

		if err := m.TransitionAction.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transitionAction")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("transitionAction")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ManifestInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManifestInfo) UnmarshalBinary(b []byte) error {
	var res ManifestInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
