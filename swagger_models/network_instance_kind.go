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

// NetworkInstanceKind network instance kind
//
// swagger:model NetworkInstanceKind
type NetworkInstanceKind string

func NewNetworkInstanceKind(value NetworkInstanceKind) *NetworkInstanceKind {
	return &value
}

// Pointer returns a pointer to a freshly-allocated NetworkInstanceKind.
func (m NetworkInstanceKind) Pointer() *NetworkInstanceKind {
	return &m
}

const (

	// NetworkInstanceKindNETWORKINSTANCEKINDUNSPECIFIED captures enum value "NETWORK_INSTANCE_KIND_UNSPECIFIED"
	NetworkInstanceKindNETWORKINSTANCEKINDUNSPECIFIED NetworkInstanceKind = "NETWORK_INSTANCE_KIND_UNSPECIFIED"

	// NetworkInstanceKindNETWORKINSTANCEKINDTRANSPARENT captures enum value "NETWORK_INSTANCE_KIND_TRANSPARENT"
	NetworkInstanceKindNETWORKINSTANCEKINDTRANSPARENT NetworkInstanceKind = "NETWORK_INSTANCE_KIND_TRANSPARENT"

	// NetworkInstanceKindNETWORKINSTANCEKINDSWITCH captures enum value "NETWORK_INSTANCE_KIND_SWITCH"
	NetworkInstanceKindNETWORKINSTANCEKINDSWITCH NetworkInstanceKind = "NETWORK_INSTANCE_KIND_SWITCH"

	// NetworkInstanceKindNETWORKINSTANCEKINDLOCAL captures enum value "NETWORK_INSTANCE_KIND_LOCAL"
	NetworkInstanceKindNETWORKINSTANCEKINDLOCAL NetworkInstanceKind = "NETWORK_INSTANCE_KIND_LOCAL"

	// NetworkInstanceKindNETWORKINSTANCEKINDCLOUD captures enum value "NETWORK_INSTANCE_KIND_CLOUD"
	NetworkInstanceKindNETWORKINSTANCEKINDCLOUD NetworkInstanceKind = "NETWORK_INSTANCE_KIND_CLOUD"

	// NetworkInstanceKindNETWORKINSTANCEKINDMESH captures enum value "NETWORK_INSTANCE_KIND_MESH"
	NetworkInstanceKindNETWORKINSTANCEKINDMESH NetworkInstanceKind = "NETWORK_INSTANCE_KIND_MESH"

	// NetworkInstanceKindNETWORKINSTANCEKINDHONEYPOT captures enum value "NETWORK_INSTANCE_KIND_HONEYPOT"
	NetworkInstanceKindNETWORKINSTANCEKINDHONEYPOT NetworkInstanceKind = "NETWORK_INSTANCE_KIND_HONEYPOT"
)

// for schema
var networkInstanceKindEnum []interface{}

func init() {
	var res []NetworkInstanceKind
	if err := json.Unmarshal([]byte(`["NETWORK_INSTANCE_KIND_UNSPECIFIED","NETWORK_INSTANCE_KIND_TRANSPARENT","NETWORK_INSTANCE_KIND_SWITCH","NETWORK_INSTANCE_KIND_LOCAL","NETWORK_INSTANCE_KIND_CLOUD","NETWORK_INSTANCE_KIND_MESH","NETWORK_INSTANCE_KIND_HONEYPOT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkInstanceKindEnum = append(networkInstanceKindEnum, v)
	}
}

func (m NetworkInstanceKind) validateNetworkInstanceKindEnum(path, location string, value NetworkInstanceKind) error {
	if err := validate.EnumCase(path, location, value, networkInstanceKindEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this network instance kind
func (m NetworkInstanceKind) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateNetworkInstanceKindEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this network instance kind based on context it is used
func (m NetworkInstanceKind) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
