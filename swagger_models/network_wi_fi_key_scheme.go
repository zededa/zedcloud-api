// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NetworkWiFiKeyScheme network wi fi key scheme
//
// swagger:model NetworkWiFiKeyScheme
type NetworkWiFiKeyScheme string

func NewNetworkWiFiKeyScheme(value NetworkWiFiKeyScheme) *NetworkWiFiKeyScheme {
	return &value
}

// Pointer returns a pointer to a freshly-allocated NetworkWiFiKeyScheme.
func (m NetworkWiFiKeyScheme) Pointer() *NetworkWiFiKeyScheme {
	return &m
}

const (

	// NetworkWiFiKeySchemeNETWORKWIFIKEYSCHEMEUNSPECIFIED captures enum value "NETWORK_WIFIKEY_SCHEME_UNSPECIFIED"
	NetworkWiFiKeySchemeNETWORKWIFIKEYSCHEMEUNSPECIFIED NetworkWiFiKeyScheme = "NETWORK_WIFIKEY_SCHEME_UNSPECIFIED"

	// NetworkWiFiKeySchemeNETWORKWIFIKEYSCHEMEWPAPSK captures enum value "NETWORK_WIFIKEY_SCHEME_WPAPSK"
	NetworkWiFiKeySchemeNETWORKWIFIKEYSCHEMEWPAPSK NetworkWiFiKeyScheme = "NETWORK_WIFIKEY_SCHEME_WPAPSK"

	// NetworkWiFiKeySchemeNETWORKWIFIKEYSCHEMEWPAEAP captures enum value "NETWORK_WIFIKEY_SCHEME_WPAEAP"
	NetworkWiFiKeySchemeNETWORKWIFIKEYSCHEMEWPAEAP NetworkWiFiKeyScheme = "NETWORK_WIFIKEY_SCHEME_WPAEAP"
)

// for schema
var networkWiFiKeySchemeEnum []interface{}

func init() {
	var res []NetworkWiFiKeyScheme
	if err := json.Unmarshal([]byte(`["NETWORK_WIFIKEY_SCHEME_UNSPECIFIED","NETWORK_WIFIKEY_SCHEME_WPAPSK","NETWORK_WIFIKEY_SCHEME_WPAEAP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkWiFiKeySchemeEnum = append(networkWiFiKeySchemeEnum, v)
	}
}

func (m NetworkWiFiKeyScheme) validateNetworkWiFiKeySchemeEnum(path, location string, value NetworkWiFiKeyScheme) error {
	if err := validate.EnumCase(path, location, value, networkWiFiKeySchemeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this network wi fi key scheme
func (m NetworkWiFiKeyScheme) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateNetworkWiFiKeySchemeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this network wi fi key scheme based on context it is used
func (m NetworkWiFiKeyScheme) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
