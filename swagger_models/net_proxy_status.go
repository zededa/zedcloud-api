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
)

// NetProxyStatus NetProxyStatus is used to store the proxy configurations
//
// NetProxyStatus is used to store the proxy configurations
//
// swagger:model NetProxyStatus
type NetProxyStatus struct {

	// exceptions
	Exceptions string `json:"exceptions,omitempty"`

	// Enable, the proxy configuration coming from network
	//
	// Network Proxy flag
	NetworkProxy bool `json:"networkProxy,omitempty"`

	// Direct URL for wpad.dat download
	//
	// Network Proxy URL
	NetworkProxyURL string `json:"networkProxyURL,omitempty"`

	// proxy configuration in a pacfile
	//
	// proxy configuration in a pacfile
	Pacfile string `json:"pacfile,omitempty"`

	// protocol level proxies
	//
	// protocol level proxies
	Proxies []*NetProxyServer `json:"proxies"`

	// WPAD Proxy URL
	WpadProxyURL string `json:"wpadProxyURL,omitempty"`
}

// Validate validates this net proxy status
func (m *NetProxyStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProxies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetProxyStatus) validateProxies(formats strfmt.Registry) error {
	if swag.IsZero(m.Proxies) { // not required
		return nil
	}

	for i := 0; i < len(m.Proxies); i++ {
		if swag.IsZero(m.Proxies[i]) { // not required
			continue
		}

		if m.Proxies[i] != nil {
			if err := m.Proxies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("proxies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this net proxy status based on the context it is used
func (m *NetProxyStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProxies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetProxyStatus) contextValidateProxies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Proxies); i++ {

		if m.Proxies[i] != nil {
			if err := m.Proxies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("proxies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetProxyStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetProxyStatus) UnmarshalBinary(b []byte) error {
	var res NetProxyStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
