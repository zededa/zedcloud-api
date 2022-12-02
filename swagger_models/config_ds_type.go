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

// ConfigDsType config ds type
//
// swagger:model configDsType
type ConfigDsType string

func NewConfigDsType(value ConfigDsType) *ConfigDsType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ConfigDsType.
func (m ConfigDsType) Pointer() *ConfigDsType {
	return &m
}

const (

	// ConfigDsTypeDsUnknown captures enum value "DsUnknown"
	ConfigDsTypeDsUnknown ConfigDsType = "DsUnknown"

	// ConfigDsTypeDsHTTP captures enum value "DsHttp"
	ConfigDsTypeDsHTTP ConfigDsType = "DsHttp"

	// ConfigDsTypeDsHTTPS captures enum value "DsHttps"
	ConfigDsTypeDsHTTPS ConfigDsType = "DsHttps"

	// ConfigDsTypeDsS3 captures enum value "DsS3"
	ConfigDsTypeDsS3 ConfigDsType = "DsS3"

	// ConfigDsTypeDsSFTP captures enum value "DsSFTP"
	ConfigDsTypeDsSFTP ConfigDsType = "DsSFTP"

	// ConfigDsTypeDsContainerRegistry captures enum value "DsContainerRegistry"
	ConfigDsTypeDsContainerRegistry ConfigDsType = "DsContainerRegistry"

	// ConfigDsTypeDsAzureBlob captures enum value "DsAzureBlob"
	ConfigDsTypeDsAzureBlob ConfigDsType = "DsAzureBlob"

	// ConfigDsTypeDsGoogleStorage captures enum value "DsGoogleStorage"
	ConfigDsTypeDsGoogleStorage ConfigDsType = "DsGoogleStorage"
)

// for schema
var configDsTypeEnum []interface{}

func init() {
	var res []ConfigDsType
	if err := json.Unmarshal([]byte(`["DsUnknown","DsHttp","DsHttps","DsS3","DsSFTP","DsContainerRegistry","DsAzureBlob","DsGoogleStorage"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configDsTypeEnum = append(configDsTypeEnum, v)
	}
}

func (m ConfigDsType) validateConfigDsTypeEnum(path, location string, value ConfigDsType) error {
	if err := validate.EnumCase(path, location, value, configDsTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this config ds type
func (m ConfigDsType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateConfigDsTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this config ds type based on context it is used
func (m ConfigDsType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
