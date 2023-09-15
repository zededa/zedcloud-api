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
	"github.com/go-openapi/validate"
)

// DeviceStatusConfigList Device status config composite api response detail
//
// # Device statu sna config composite api response details
//
// swagger:model DeviceStatusConfigList
type DeviceStatusConfigList struct {

	// List of device status config
	// Required: true
	List []*DeviceStatusConfig `json:"list"`

	// Page details of the filtered records
	Next *Cursor `json:"next,omitempty"`

	// Device status config summary by app instance count
	// Required: true
	SummaryByAppInstanceCount *Summary `json:"summaryByAppInstanceCount"`

	// Device status config summary by eve distribution
	// Required: true
	SummaryByEVEDistribution *Summary `json:"summaryByEVEDistribution"`

	// Device status config summary by state
	// Required: true
	SummaryByState *Summary `json:"summaryByState"`

	// total count of devices
	// Required: true
	TotalCount *int64 `json:"totalCount"`

	// total count of edgeview active of devices
	TotalEvActiveCount int64 `json:"totalEvActiveCount,omitempty"`
}

// Validate validates this device status config list
func (m *DeviceStatusConfigList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummaryByAppInstanceCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummaryByEVEDistribution(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummaryByState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceStatusConfigList) validateList(formats strfmt.Registry) error {

	if err := validate.Required("list", "body", m.List); err != nil {
		return err
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

func (m *DeviceStatusConfigList) validateNext(formats strfmt.Registry) error {
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

func (m *DeviceStatusConfigList) validateSummaryByAppInstanceCount(formats strfmt.Registry) error {

	if err := validate.Required("summaryByAppInstanceCount", "body", m.SummaryByAppInstanceCount); err != nil {
		return err
	}

	if m.SummaryByAppInstanceCount != nil {
		if err := m.SummaryByAppInstanceCount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summaryByAppInstanceCount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summaryByAppInstanceCount")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceStatusConfigList) validateSummaryByEVEDistribution(formats strfmt.Registry) error {

	if err := validate.Required("summaryByEVEDistribution", "body", m.SummaryByEVEDistribution); err != nil {
		return err
	}

	if m.SummaryByEVEDistribution != nil {
		if err := m.SummaryByEVEDistribution.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summaryByEVEDistribution")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summaryByEVEDistribution")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceStatusConfigList) validateSummaryByState(formats strfmt.Registry) error {

	if err := validate.Required("summaryByState", "body", m.SummaryByState); err != nil {
		return err
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

func (m *DeviceStatusConfigList) validateTotalCount(formats strfmt.Registry) error {

	if err := validate.Required("totalCount", "body", m.TotalCount); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this device status config list based on the context it is used
func (m *DeviceStatusConfigList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSummaryByAppInstanceCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSummaryByEVEDistribution(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSummaryByState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceStatusConfigList) contextValidateList(ctx context.Context, formats strfmt.Registry) error {

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

func (m *DeviceStatusConfigList) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

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

func (m *DeviceStatusConfigList) contextValidateSummaryByAppInstanceCount(ctx context.Context, formats strfmt.Registry) error {

	if m.SummaryByAppInstanceCount != nil {

		if err := m.SummaryByAppInstanceCount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summaryByAppInstanceCount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summaryByAppInstanceCount")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceStatusConfigList) contextValidateSummaryByEVEDistribution(ctx context.Context, formats strfmt.Registry) error {

	if m.SummaryByEVEDistribution != nil {

		if err := m.SummaryByEVEDistribution.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summaryByEVEDistribution")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summaryByEVEDistribution")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceStatusConfigList) contextValidateSummaryByState(ctx context.Context, formats strfmt.Registry) error {

	if m.SummaryByState != nil {

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

// MarshalBinary interface implementation
func (m *DeviceStatusConfigList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceStatusConfigList) UnmarshalBinary(b []byte) error {
	var res DeviceStatusConfigList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
