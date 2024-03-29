// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConfigZnetStaticDNSEntry These are list of static mapping that can be added to network
//
// swagger:model configZnetStaticDNSEntry
type ConfigZnetStaticDNSEntry struct {

	// address
	Address []string `json:"Address"`

	// host name
	HostName string `json:"HostName,omitempty"`
}

// Validate validates this config znet static DNS entry
func (m *ConfigZnetStaticDNSEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this config znet static DNS entry based on context it is used
func (m *ConfigZnetStaticDNSEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConfigZnetStaticDNSEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigZnetStaticDNSEntry) UnmarshalBinary(b []byte) error {
	var res ConfigZnetStaticDNSEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
