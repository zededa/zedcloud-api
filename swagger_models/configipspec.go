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

// Configipspec Common for IPv4 and IPv6
//
// swagger:model configipspec
type Configipspec struct {

	// dhcp
	Dhcp *ConfigDHCPType `json:"dhcp,omitempty"`

	// for IPAM management when dhcp is turned on.
	// If none provided, system will default pool.
	DhcpRange *ConfigipRange `json:"dhcpRange,omitempty"`

	// dns
	DNS []string `json:"dns"`

	// domain
	Domain string `json:"domain,omitempty"`

	// gateway
	Gateway string `json:"gateway,omitempty"`

	// ntp
	Ntp string `json:"ntp,omitempty"`

	// subnet is CIDR format...x.y.z.l/nn
	Subnet string `json:"subnet,omitempty"`
}

// Validate validates this configipspec
func (m *Configipspec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDhcp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDhcpRange(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configipspec) validateDhcp(formats strfmt.Registry) error {
	if swag.IsZero(m.Dhcp) { // not required
		return nil
	}

	if m.Dhcp != nil {
		if err := m.Dhcp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dhcp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dhcp")
			}
			return err
		}
	}

	return nil
}

func (m *Configipspec) validateDhcpRange(formats strfmt.Registry) error {
	if swag.IsZero(m.DhcpRange) { // not required
		return nil
	}

	if m.DhcpRange != nil {
		if err := m.DhcpRange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dhcpRange")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dhcpRange")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configipspec based on the context it is used
func (m *Configipspec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDhcp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDhcpRange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configipspec) contextValidateDhcp(ctx context.Context, formats strfmt.Registry) error {

	if m.Dhcp != nil {

		if swag.IsZero(m.Dhcp) { // not required
			return nil
		}

		if err := m.Dhcp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dhcp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dhcp")
			}
			return err
		}
	}

	return nil
}

func (m *Configipspec) contextValidateDhcpRange(ctx context.Context, formats strfmt.Registry) error {

	if m.DhcpRange != nil {

		if swag.IsZero(m.DhcpRange) { // not required
			return nil
		}

		if err := m.DhcpRange.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dhcpRange")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dhcpRange")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configipspec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configipspec) UnmarshalBinary(b []byte) error {
	var res Configipspec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
