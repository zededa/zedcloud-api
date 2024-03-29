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

// AAAFrontendSessionDetailsResponseCause a a a frontend session details response cause
//
// swagger:model AAA_Frontend_SessionDetailsResponseCause
type AAAFrontendSessionDetailsResponseCause string

func NewAAAFrontendSessionDetailsResponseCause(value AAAFrontendSessionDetailsResponseCause) *AAAFrontendSessionDetailsResponseCause {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AAAFrontendSessionDetailsResponseCause.
func (m AAAFrontendSessionDetailsResponseCause) Pointer() *AAAFrontendSessionDetailsResponseCause {
	return &m
}

const (

	// AAAFrontendSessionDetailsResponseCauseUNSPECIFIED captures enum value "UNSPECIFIED"
	AAAFrontendSessionDetailsResponseCauseUNSPECIFIED AAAFrontendSessionDetailsResponseCause = "UNSPECIFIED"

	// AAAFrontendSessionDetailsResponseCauseOK captures enum value "OK"
	AAAFrontendSessionDetailsResponseCauseOK AAAFrontendSessionDetailsResponseCause = "OK"

	// AAAFrontendSessionDetailsResponseCauseFAILED captures enum value "FAILED"
	AAAFrontendSessionDetailsResponseCauseFAILED AAAFrontendSessionDetailsResponseCause = "FAILED"
)

// for schema
var aAAFrontendSessionDetailsResponseCauseEnum []interface{}

func init() {
	var res []AAAFrontendSessionDetailsResponseCause
	if err := json.Unmarshal([]byte(`["UNSPECIFIED","OK","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aAAFrontendSessionDetailsResponseCauseEnum = append(aAAFrontendSessionDetailsResponseCauseEnum, v)
	}
}

func (m AAAFrontendSessionDetailsResponseCause) validateAAAFrontendSessionDetailsResponseCauseEnum(path, location string, value AAAFrontendSessionDetailsResponseCause) error {
	if err := validate.EnumCase(path, location, value, aAAFrontendSessionDetailsResponseCauseEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this a a a frontend session details response cause
func (m AAAFrontendSessionDetailsResponseCause) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAAAFrontendSessionDetailsResponseCauseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this a a a frontend session details response cause based on context it is used
func (m AAAFrontendSessionDetailsResponseCause) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
