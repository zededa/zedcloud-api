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

// AAAFailureResponseGenerateTokenCause a a a failure response generate token cause
//
// swagger:model AAAFailureResponseGenerateTokenCause
type AAAFailureResponseGenerateTokenCause string

func NewAAAFailureResponseGenerateTokenCause(value AAAFailureResponseGenerateTokenCause) *AAAFailureResponseGenerateTokenCause {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AAAFailureResponseGenerateTokenCause.
func (m AAAFailureResponseGenerateTokenCause) Pointer() *AAAFailureResponseGenerateTokenCause {
	return &m
}

const (

	// AAAFailureResponseGenerateTokenCauseINVALID captures enum value "__INVALID__"
	AAAFailureResponseGenerateTokenCauseINVALID AAAFailureResponseGenerateTokenCause = "__INVALID__"

	// AAAFailureResponseGenerateTokenCauseAAAFailureResponseGenerateTokenCauseUserUnknown captures enum value "AAAFailureResponseGenerateTokenCauseUserUnknown"
	AAAFailureResponseGenerateTokenCauseAAAFailureResponseGenerateTokenCauseUserUnknown AAAFailureResponseGenerateTokenCause = "AAAFailureResponseGenerateTokenCauseUserUnknown"

	// AAAFailureResponseGenerateTokenCauseAAAFailureResponseGenerateTokenCauseException captures enum value "AAAFailureResponseGenerateTokenCauseException"
	AAAFailureResponseGenerateTokenCauseAAAFailureResponseGenerateTokenCauseException AAAFailureResponseGenerateTokenCause = "AAAFailureResponseGenerateTokenCauseException"

	// AAAFailureResponseGenerateTokenCauseAAAFailureResponseGenerateTokenCauseEnterpriseUnknown captures enum value "AAAFailureResponseGenerateTokenCauseEnterpriseUnknown"
	AAAFailureResponseGenerateTokenCauseAAAFailureResponseGenerateTokenCauseEnterpriseUnknown AAAFailureResponseGenerateTokenCause = "AAAFailureResponseGenerateTokenCauseEnterpriseUnknown"

	// AAAFailureResponseGenerateTokenCauseAAAFailureResponseGenerateTokenCauseForbidden captures enum value "AAAFailureResponseGenerateTokenCauseForbidden"
	AAAFailureResponseGenerateTokenCauseAAAFailureResponseGenerateTokenCauseForbidden AAAFailureResponseGenerateTokenCause = "AAAFailureResponseGenerateTokenCauseForbidden"

	// AAAFailureResponseGenerateTokenCauseAAAFailureResponseGenerateTokenCauseInternalError captures enum value "AAAFailureResponseGenerateTokenCauseInternalError"
	AAAFailureResponseGenerateTokenCauseAAAFailureResponseGenerateTokenCauseInternalError AAAFailureResponseGenerateTokenCause = "AAAFailureResponseGenerateTokenCauseInternalError"
)

// for schema
var aAAFailureResponseGenerateTokenCauseEnum []interface{}

func init() {
	var res []AAAFailureResponseGenerateTokenCause
	if err := json.Unmarshal([]byte(`["__INVALID__","AAAFailureResponseGenerateTokenCauseUserUnknown","AAAFailureResponseGenerateTokenCauseException","AAAFailureResponseGenerateTokenCauseEnterpriseUnknown","AAAFailureResponseGenerateTokenCauseForbidden","AAAFailureResponseGenerateTokenCauseInternalError"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aAAFailureResponseGenerateTokenCauseEnum = append(aAAFailureResponseGenerateTokenCauseEnum, v)
	}
}

func (m AAAFailureResponseGenerateTokenCause) validateAAAFailureResponseGenerateTokenCauseEnum(path, location string, value AAAFailureResponseGenerateTokenCause) error {
	if err := validate.EnumCase(path, location, value, aAAFailureResponseGenerateTokenCauseEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this a a a failure response generate token cause
func (m AAAFailureResponseGenerateTokenCause) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAAAFailureResponseGenerateTokenCauseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this a a a failure response generate token cause based on context it is used
func (m AAAFailureResponseGenerateTokenCause) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
