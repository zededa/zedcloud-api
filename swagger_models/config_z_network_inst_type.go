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

// ConfigZNetworkInstType config z network inst type
//
// swagger:model configZNetworkInstType
type ConfigZNetworkInstType string

func NewConfigZNetworkInstType(value ConfigZNetworkInstType) *ConfigZNetworkInstType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ConfigZNetworkInstType.
func (m ConfigZNetworkInstType) Pointer() *ConfigZNetworkInstType {
	return &m
}

const (

	// ConfigZNetworkInstTypeZNetInstFirst captures enum value "ZNetInstFirst"
	ConfigZNetworkInstTypeZNetInstFirst ConfigZNetworkInstType = "ZNetInstFirst"

	// ConfigZNetworkInstTypeZnetInstSwitch captures enum value "ZnetInstSwitch"
	ConfigZNetworkInstTypeZnetInstSwitch ConfigZNetworkInstType = "ZnetInstSwitch"

	// ConfigZNetworkInstTypeZnetInstLocal captures enum value "ZnetInstLocal"
	ConfigZNetworkInstTypeZnetInstLocal ConfigZNetworkInstType = "ZnetInstLocal"

	// ConfigZNetworkInstTypeZnetInstCloud captures enum value "ZnetInstCloud"
	ConfigZNetworkInstTypeZnetInstCloud ConfigZNetworkInstType = "ZnetInstCloud"

	// ConfigZNetworkInstTypeZnetInstMesh captures enum value "ZnetInstMesh"
	ConfigZNetworkInstTypeZnetInstMesh ConfigZNetworkInstType = "ZnetInstMesh"

	// ConfigZNetworkInstTypeZnetInstHoneyPot captures enum value "ZnetInstHoneyPot"
	ConfigZNetworkInstTypeZnetInstHoneyPot ConfigZNetworkInstType = "ZnetInstHoneyPot"

	// ConfigZNetworkInstTypeZnetInstTransparent captures enum value "ZnetInstTransparent"
	ConfigZNetworkInstTypeZnetInstTransparent ConfigZNetworkInstType = "ZnetInstTransparent"

	// ConfigZNetworkInstTypeZNetInstLast captures enum value "ZNetInstLast"
	ConfigZNetworkInstTypeZNetInstLast ConfigZNetworkInstType = "ZNetInstLast"
)

// for schema
var configZNetworkInstTypeEnum []interface{}

func init() {
	var res []ConfigZNetworkInstType
	if err := json.Unmarshal([]byte(`["ZNetInstFirst","ZnetInstSwitch","ZnetInstLocal","ZnetInstCloud","ZnetInstMesh","ZnetInstHoneyPot","ZnetInstTransparent","ZNetInstLast"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configZNetworkInstTypeEnum = append(configZNetworkInstTypeEnum, v)
	}
}

func (m ConfigZNetworkInstType) validateConfigZNetworkInstTypeEnum(path, location string, value ConfigZNetworkInstType) error {
	if err := validate.EnumCase(path, location, value, configZNetworkInstTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this config z network inst type
func (m ConfigZNetworkInstType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateConfigZNetworkInstTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this config z network inst type based on context it is used
func (m ConfigZNetworkInstType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
