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

// NetworkStatus NetworkStatus is used to store the Network Status
//
// # NetworkStatus is used to store the Network status
//
// swagger:model NetworkStatus
type NetworkStatus struct {

	// Default Routers
	DefaultRouters []string `json:"defaultRouters"`

	// DNS Configuration
	DNS *DNSInfo `json:"dns,omitempty"`

	// Network error details
	ErrInfo *DeviceError `json:"errInfo,omitempty"`

	// Location from GNSS receivers on WWAN type adapters
	GpsLocation *GPSLocation `json:"gpsLocation,omitempty"`

	// ifName
	IfName string `json:"ifName,omitempty"`

	// Array of IP addresses
	IPAddrs []string `json:"ipAddrs"`

	// Geo Location Details
	Location *GeoLocation `json:"location,omitempty"`

	// mac Address
	MacAddr string `json:"macAddr,omitempty"`

	// Network Proxy status
	Proxy *NetProxyStatus `json:"proxy,omitempty"`

	// Network Status flag
	Up bool `json:"up,omitempty"`

	// Uplink flag
	Uplink bool `json:"uplink,omitempty"`
}

// Validate validates this network status
func (m *NetworkStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDNS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpsLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkStatus) validateDNS(formats strfmt.Registry) error {
	if swag.IsZero(m.DNS) { // not required
		return nil
	}

	if m.DNS != nil {
		if err := m.DNS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dns")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dns")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkStatus) validateErrInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ErrInfo) { // not required
		return nil
	}

	if m.ErrInfo != nil {
		if err := m.ErrInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("errInfo")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkStatus) validateGpsLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.GpsLocation) { // not required
		return nil
	}

	if m.GpsLocation != nil {
		if err := m.GpsLocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gpsLocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gpsLocation")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkStatus) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkStatus) validateProxy(formats strfmt.Registry) error {
	if swag.IsZero(m.Proxy) { // not required
		return nil
	}

	if m.Proxy != nil {
		if err := m.Proxy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proxy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this network status based on the context it is used
func (m *NetworkStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDNS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateErrInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGpsLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProxy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkStatus) contextValidateDNS(ctx context.Context, formats strfmt.Registry) error {

	if m.DNS != nil {
		if err := m.DNS.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dns")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dns")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkStatus) contextValidateErrInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.ErrInfo != nil {
		if err := m.ErrInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("errInfo")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkStatus) contextValidateGpsLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.GpsLocation != nil {
		if err := m.GpsLocation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gpsLocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gpsLocation")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkStatus) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {
		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkStatus) contextValidateProxy(ctx context.Context, formats strfmt.Registry) error {

	if m.Proxy != nil {
		if err := m.Proxy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proxy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkStatus) UnmarshalBinary(b []byte) error {
	var res NetworkStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
