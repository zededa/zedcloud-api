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

// NetworkInstanceDhcpType NetworkInstanceDhcpType: Used in the network instance provide
//
//	dhcp server configuration
//
// swagger:model NetworkInstanceDhcpType
type NetworkInstanceDhcpType string

func NewNetworkInstanceDhcpType(value NetworkInstanceDhcpType) *NetworkInstanceDhcpType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated NetworkInstanceDhcpType.
func (m NetworkInstanceDhcpType) Pointer() *NetworkInstanceDhcpType {
	return &m
}

const (

	// NetworkInstanceDhcpTypeNETWORKINSTANCEDHCPTYPEUNSPECIFIED captures enum value "NETWORK_INSTANCE_DHCP_TYPE_UNSPECIFIED"
	NetworkInstanceDhcpTypeNETWORKINSTANCEDHCPTYPEUNSPECIFIED NetworkInstanceDhcpType = "NETWORK_INSTANCE_DHCP_TYPE_UNSPECIFIED"

	// NetworkInstanceDhcpTypeNETWORKINSTANCEDHCPTYPEV4 captures enum value "NETWORK_INSTANCE_DHCP_TYPE_V4"
	NetworkInstanceDhcpTypeNETWORKINSTANCEDHCPTYPEV4 NetworkInstanceDhcpType = "NETWORK_INSTANCE_DHCP_TYPE_V4"

	// NetworkInstanceDhcpTypeNETWORKINSTANCEDHCPTYPEV6 captures enum value "NETWORK_INSTANCE_DHCP_TYPE_V6"
	NetworkInstanceDhcpTypeNETWORKINSTANCEDHCPTYPEV6 NetworkInstanceDhcpType = "NETWORK_INSTANCE_DHCP_TYPE_V6"

	// NetworkInstanceDhcpTypeNETWORKINSTANCEDHCPTYPECRYPTOEID captures enum value "NETWORK_INSTANCE_DHCP_TYPE_CRYPTOEID"
	NetworkInstanceDhcpTypeNETWORKINSTANCEDHCPTYPECRYPTOEID NetworkInstanceDhcpType = "NETWORK_INSTANCE_DHCP_TYPE_CRYPTOEID"

	// NetworkInstanceDhcpTypeNETWORKINSTANCEDHCPTYPECRYPTOV4 captures enum value "NETWORK_INSTANCE_DHCP_TYPE_CRYPTOV4"
	NetworkInstanceDhcpTypeNETWORKINSTANCEDHCPTYPECRYPTOV4 NetworkInstanceDhcpType = "NETWORK_INSTANCE_DHCP_TYPE_CRYPTOV4"

	// NetworkInstanceDhcpTypeNETWORKINSTANCEDHCPTYPECRYPTOV6 captures enum value "NETWORK_INSTANCE_DHCP_TYPE_CRYPTOV6"
	NetworkInstanceDhcpTypeNETWORKINSTANCEDHCPTYPECRYPTOV6 NetworkInstanceDhcpType = "NETWORK_INSTANCE_DHCP_TYPE_CRYPTOV6"
)

// for schema
var networkInstanceDhcpTypeEnum []interface{}

func init() {
	var res []NetworkInstanceDhcpType
	if err := json.Unmarshal([]byte(`["NETWORK_INSTANCE_DHCP_TYPE_UNSPECIFIED","NETWORK_INSTANCE_DHCP_TYPE_V4","NETWORK_INSTANCE_DHCP_TYPE_V6","NETWORK_INSTANCE_DHCP_TYPE_CRYPTOEID","NETWORK_INSTANCE_DHCP_TYPE_CRYPTOV4","NETWORK_INSTANCE_DHCP_TYPE_CRYPTOV6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkInstanceDhcpTypeEnum = append(networkInstanceDhcpTypeEnum, v)
	}
}

func (m NetworkInstanceDhcpType) validateNetworkInstanceDhcpTypeEnum(path, location string, value NetworkInstanceDhcpType) error {
	if err := validate.EnumCase(path, location, value, networkInstanceDhcpTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this network instance dhcp type
func (m NetworkInstanceDhcpType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateNetworkInstanceDhcpTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this network instance dhcp type based on context it is used
func (m NetworkInstanceDhcpType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
