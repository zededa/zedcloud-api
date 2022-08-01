// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

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

// ConfigBondAdapter BondAdapter aggregates multiple PhysicalIOs into one adapter for redundancy
// and load-spreading purposes.
//
// swagger:model configBondAdapter
type ConfigBondAdapter struct {

	// The ARP monitor is based on the communication to the target
	// hosts designated by their IP addresses. Even if the link is beyond
	// the nearest connected switch, the APR monitor can detect it.
	Arp *ConfigArpMonitor `json:"arp,omitempty"`

	// A bonding mode specifies the policy indicating how bonding slaves are used
	// during network transmission
	BondMode *ConfigBondMode `json:"bondMode,omitempty"`

	// A physical name of the bond interface.
	// Note that the interface name is limited in Linux kernel to 15 characters.
	// If not defined, logicallabel will be used instead.
	InterfaceName string `json:"interfaceName,omitempty"`

	// Option specifying the rate in which EVE will ask LACP link partners
	// to transmit LACPDU packets in 802.3ad mode.
	LacpRate *ConfigLacpRate `json:"lacpRate,omitempty"`

	// Name of this bond adapter.
	Logicallabel string `json:"logicallabel,omitempty"`

	// Logical names of aggregated PhysicalIOs.
	// For all bonding modes but Active-Backup the order is irrelevant.
	// In the Active-Backup mode (BOND_MODE_ACTIVE_BACKUP), the first PhysicalIO
	// in the list will be considered as the primary port (i.e. only when
	// the primary is off-line will alternate ports be used).
	LowerLayerNames []string `json:"lowerLayerNames"`

	// The MII monitor is driver-dependent. It monitors the links from the device
	// to the nearest connected switch. If the failure occurs beyond the nearest
	// connected switch, it cannot be detected by MII monitor.
	Mii *ConfigMIIMonitor `json:"mii,omitempty"`
}

// Validate validates this config bond adapter
func (m *ConfigBondAdapter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBondMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLacpRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMii(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigBondAdapter) validateArp(formats strfmt.Registry) error {
	if swag.IsZero(m.Arp) { // not required
		return nil
	}

	if m.Arp != nil {
		if err := m.Arp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("arp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("arp")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigBondAdapter) validateBondMode(formats strfmt.Registry) error {
	if swag.IsZero(m.BondMode) { // not required
		return nil
	}

	if m.BondMode != nil {
		if err := m.BondMode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bondMode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bondMode")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigBondAdapter) validateLacpRate(formats strfmt.Registry) error {
	if swag.IsZero(m.LacpRate) { // not required
		return nil
	}

	if m.LacpRate != nil {
		if err := m.LacpRate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lacpRate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lacpRate")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigBondAdapter) validateMii(formats strfmt.Registry) error {
	if swag.IsZero(m.Mii) { // not required
		return nil
	}

	if m.Mii != nil {
		if err := m.Mii.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mii")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mii")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config bond adapter based on the context it is used
func (m *ConfigBondAdapter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBondMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLacpRate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMii(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigBondAdapter) contextValidateArp(ctx context.Context, formats strfmt.Registry) error {

	if m.Arp != nil {
		if err := m.Arp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("arp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("arp")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigBondAdapter) contextValidateBondMode(ctx context.Context, formats strfmt.Registry) error {

	if m.BondMode != nil {
		if err := m.BondMode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bondMode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bondMode")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigBondAdapter) contextValidateLacpRate(ctx context.Context, formats strfmt.Registry) error {

	if m.LacpRate != nil {
		if err := m.LacpRate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lacpRate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lacpRate")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigBondAdapter) contextValidateMii(ctx context.Context, formats strfmt.Registry) error {

	if m.Mii != nil {
		if err := m.Mii.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mii")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mii")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigBondAdapter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigBondAdapter) UnmarshalBinary(b []byte) error {
	var res ConfigBondAdapter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}