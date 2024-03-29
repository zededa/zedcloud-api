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

// ConfigEncryptionScheme Encryption Scheme for Cipher Payload
//
// swagger:model configEncryptionScheme
type ConfigEncryptionScheme string

func NewConfigEncryptionScheme(value ConfigEncryptionScheme) *ConfigEncryptionScheme {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ConfigEncryptionScheme.
func (m ConfigEncryptionScheme) Pointer() *ConfigEncryptionScheme {
	return &m
}

const (

	// ConfigEncryptionSchemeSANONE captures enum value "SA_NONE"
	ConfigEncryptionSchemeSANONE ConfigEncryptionScheme = "SA_NONE"

	// ConfigEncryptionSchemeSAAES256CFB captures enum value "SA_AES_256_CFB"
	ConfigEncryptionSchemeSAAES256CFB ConfigEncryptionScheme = "SA_AES_256_CFB"
)

// for schema
var configEncryptionSchemeEnum []interface{}

func init() {
	var res []ConfigEncryptionScheme
	if err := json.Unmarshal([]byte(`["SA_NONE","SA_AES_256_CFB"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configEncryptionSchemeEnum = append(configEncryptionSchemeEnum, v)
	}
}

func (m ConfigEncryptionScheme) validateConfigEncryptionSchemeEnum(path, location string, value ConfigEncryptionScheme) error {
	if err := validate.EnumCase(path, location, value, configEncryptionSchemeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this config encryption scheme
func (m ConfigEncryptionScheme) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateConfigEncryptionSchemeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this config encryption scheme based on context it is used
func (m ConfigEncryptionScheme) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
