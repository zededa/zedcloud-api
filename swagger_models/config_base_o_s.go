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

// ConfigBaseOS config base o s
//
// swagger:model configBaseOS
type ConfigBaseOS struct {

	// if not set BaseOS will be installed,
	// but not activated
	Activate bool `json:"activate,omitempty"`

	// base os version
	BaseOsVersion string `json:"baseOsVersion,omitempty"`

	// UUID for ContentTree with BaseOS image
	ContentTreeUUID string `json:"contentTreeUuid,omitempty"`

	// retry_update
	// Retry the BaseOs update if the update failed previously.
	// 1) If this image is in FAILED state, retry the image update.
	// 2) If this image is already active and fully installed (PartitionState = UPDATED),
	//    Do nothing. Just update the baseos_update_counter in Info message.
	// 3) If this image is same as active image, but status is NOT yet UPDATED, or
	//    if the update to this image is in progress, wait till the update
	//    concludes (Success / Error+rollback) - then trigger the retry as needed.
	RetryUpdate *ConfigDeviceOpsCmd `json:"retryUpdate,omitempty"`
}

// Validate validates this config base o s
func (m *ConfigBaseOS) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRetryUpdate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigBaseOS) validateRetryUpdate(formats strfmt.Registry) error {
	if swag.IsZero(m.RetryUpdate) { // not required
		return nil
	}

	if m.RetryUpdate != nil {
		if err := m.RetryUpdate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("retryUpdate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("retryUpdate")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config base o s based on the context it is used
func (m *ConfigBaseOS) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRetryUpdate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigBaseOS) contextValidateRetryUpdate(ctx context.Context, formats strfmt.Registry) error {

	if m.RetryUpdate != nil {

		if swag.IsZero(m.RetryUpdate) { // not required
			return nil
		}

		if err := m.RetryUpdate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("retryUpdate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("retryUpdate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigBaseOS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigBaseOS) UnmarshalBinary(b []byte) error {
	var res ConfigBaseOS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
