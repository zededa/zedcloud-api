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

// AAAFailureResponseSessionDetailsCause a a a failure response session details cause
//
// swagger:model AAAFailureResponseSessionDetailsCause
type AAAFailureResponseSessionDetailsCause string

func NewAAAFailureResponseSessionDetailsCause(value AAAFailureResponseSessionDetailsCause) *AAAFailureResponseSessionDetailsCause {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AAAFailureResponseSessionDetailsCause.
func (m AAAFailureResponseSessionDetailsCause) Pointer() *AAAFailureResponseSessionDetailsCause {
	return &m
}

const (

	// AAAFailureResponseSessionDetailsCauseINVALID captures enum value "__INVALID__"
	AAAFailureResponseSessionDetailsCauseINVALID AAAFailureResponseSessionDetailsCause = "__INVALID__"

	// AAAFailureResponseSessionDetailsCauseAAAFailureResponseSessionDetailsCauseUnknown captures enum value "AAAFailureResponseSessionDetailsCauseUnknown"
	AAAFailureResponseSessionDetailsCauseAAAFailureResponseSessionDetailsCauseUnknown AAAFailureResponseSessionDetailsCause = "AAAFailureResponseSessionDetailsCauseUnknown"

	// AAAFailureResponseSessionDetailsCauseAAAFailureResponseSessionDetailsCauseNoSuchSession captures enum value "AAAFailureResponseSessionDetailsCauseNoSuchSession"
	AAAFailureResponseSessionDetailsCauseAAAFailureResponseSessionDetailsCauseNoSuchSession AAAFailureResponseSessionDetailsCause = "AAAFailureResponseSessionDetailsCauseNoSuchSession"
)

// for schema
var aAAFailureResponseSessionDetailsCauseEnum []interface{}

func init() {
	var res []AAAFailureResponseSessionDetailsCause
	if err := json.Unmarshal([]byte(`["__INVALID__","AAAFailureResponseSessionDetailsCauseUnknown","AAAFailureResponseSessionDetailsCauseNoSuchSession"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aAAFailureResponseSessionDetailsCauseEnum = append(aAAFailureResponseSessionDetailsCauseEnum, v)
	}
}

func (m AAAFailureResponseSessionDetailsCause) validateAAAFailureResponseSessionDetailsCauseEnum(path, location string, value AAAFailureResponseSessionDetailsCause) error {
	if err := validate.EnumCase(path, location, value, aAAFailureResponseSessionDetailsCauseEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this a a a failure response session details cause
func (m AAAFailureResponseSessionDetailsCause) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAAAFailureResponseSessionDetailsCauseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this a a a failure response session details cause based on context it is used
func (m AAAFailureResponseSessionDetailsCause) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
