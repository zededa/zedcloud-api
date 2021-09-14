// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

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

// ConfigZNetworkOpaqueConfigType config z network opaque config type
//
// swagger:model configZNetworkOpaqueConfigType
type ConfigZNetworkOpaqueConfigType string

func NewConfigZNetworkOpaqueConfigType(value ConfigZNetworkOpaqueConfigType) *ConfigZNetworkOpaqueConfigType {
	v := value
	return &v
}

const (

	// ConfigZNetworkOpaqueConfigTypeZNetOConfigVPN captures enum value "ZNetOConfigVPN"
	ConfigZNetworkOpaqueConfigTypeZNetOConfigVPN ConfigZNetworkOpaqueConfigType = "ZNetOConfigVPN"

	// ConfigZNetworkOpaqueConfigTypeZNetOConfigLisp captures enum value "ZNetOConfigLisp"
	ConfigZNetworkOpaqueConfigTypeZNetOConfigLisp ConfigZNetworkOpaqueConfigType = "ZNetOConfigLisp"
)

// for schema
var configZNetworkOpaqueConfigTypeEnum []interface{}

func init() {
	var res []ConfigZNetworkOpaqueConfigType
	if err := json.Unmarshal([]byte(`["ZNetOConfigVPN","ZNetOConfigLisp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configZNetworkOpaqueConfigTypeEnum = append(configZNetworkOpaqueConfigTypeEnum, v)
	}
}

func (m ConfigZNetworkOpaqueConfigType) validateConfigZNetworkOpaqueConfigTypeEnum(path, location string, value ConfigZNetworkOpaqueConfigType) error {
	if err := validate.EnumCase(path, location, value, configZNetworkOpaqueConfigTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this config z network opaque config type
func (m ConfigZNetworkOpaqueConfigType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateConfigZNetworkOpaqueConfigTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this config z network opaque config type based on context it is used
func (m ConfigZNetworkOpaqueConfigType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
