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

// DeviceLisp DeviceLisp payload detail
//
// # DeviceLisp request paylod
//
// swagger:model DeviceLisp
type DeviceLisp struct {

	// EID
	// Required: true
	EID *string `json:"EID"`

	// EID hash length
	// Required: true
	EIDHashLen *int64 `json:"EIDHashLen"`

	// Client Address
	// Required: true
	ClientAddr *string `json:"clientAddr"`

	// EID allocation prefix
	// Required: true
	// Format: byte
	EidAllocationPrefix *strfmt.Base64 `json:"eidAllocationPrefix"`

	// EID allocation prefix length
	// Required: true
	EidAllocationPrefixLen *int64 `json:"eidAllocationPrefixLen"`

	// LISP instance
	// Required: true
	LispInstance *int64 `json:"lispInstance"`

	// List of Lisp servers
	// Required: true
	LispMapServers []*LispServer `json:"lispMapServers"`

	// TEMP : flag to indicate which version of LISP data plane should be running on the device
	// Required: true
	Mode *string `json:"mode"`

	// Zed development servers
	// Required: true
	ZedServers []*DevZedServer `json:"zedServers"`
}

// Validate validates this device lisp
func (m *DeviceLisp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEIDHashLen(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientAddr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEidAllocationPrefix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEidAllocationPrefixLen(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLispInstance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLispMapServers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZedServers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceLisp) validateEID(formats strfmt.Registry) error {

	if err := validate.Required("EID", "body", m.EID); err != nil {
		return err
	}

	return nil
}

func (m *DeviceLisp) validateEIDHashLen(formats strfmt.Registry) error {

	if err := validate.Required("EIDHashLen", "body", m.EIDHashLen); err != nil {
		return err
	}

	return nil
}

func (m *DeviceLisp) validateClientAddr(formats strfmt.Registry) error {

	if err := validate.Required("clientAddr", "body", m.ClientAddr); err != nil {
		return err
	}

	return nil
}

func (m *DeviceLisp) validateEidAllocationPrefix(formats strfmt.Registry) error {

	if err := validate.Required("eidAllocationPrefix", "body", m.EidAllocationPrefix); err != nil {
		return err
	}

	return nil
}

func (m *DeviceLisp) validateEidAllocationPrefixLen(formats strfmt.Registry) error {

	if err := validate.Required("eidAllocationPrefixLen", "body", m.EidAllocationPrefixLen); err != nil {
		return err
	}

	return nil
}

func (m *DeviceLisp) validateLispInstance(formats strfmt.Registry) error {

	if err := validate.Required("lispInstance", "body", m.LispInstance); err != nil {
		return err
	}

	return nil
}

func (m *DeviceLisp) validateLispMapServers(formats strfmt.Registry) error {

	if err := validate.Required("lispMapServers", "body", m.LispMapServers); err != nil {
		return err
	}

	for i := 0; i < len(m.LispMapServers); i++ {
		if swag.IsZero(m.LispMapServers[i]) { // not required
			continue
		}

		if m.LispMapServers[i] != nil {
			if err := m.LispMapServers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lispMapServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("lispMapServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceLisp) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *DeviceLisp) validateZedServers(formats strfmt.Registry) error {

	if err := validate.Required("zedServers", "body", m.ZedServers); err != nil {
		return err
	}

	for i := 0; i < len(m.ZedServers); i++ {
		if swag.IsZero(m.ZedServers[i]) { // not required
			continue
		}

		if m.ZedServers[i] != nil {
			if err := m.ZedServers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("zedServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("zedServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this device lisp based on the context it is used
func (m *DeviceLisp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLispMapServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateZedServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceLisp) contextValidateLispMapServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LispMapServers); i++ {

		if m.LispMapServers[i] != nil {

			if swag.IsZero(m.LispMapServers[i]) { // not required
				return nil
			}

			if err := m.LispMapServers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lispMapServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("lispMapServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceLisp) contextValidateZedServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ZedServers); i++ {

		if m.ZedServers[i] != nil {

			if swag.IsZero(m.ZedServers[i]) { // not required
				return nil
			}

			if err := m.ZedServers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("zedServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("zedServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceLisp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceLisp) UnmarshalBinary(b []byte) error {
	var res DeviceLisp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
