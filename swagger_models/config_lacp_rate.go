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

// ConfigLacpRate Option specifying the rate in which EVE will ask LACP link partners
// to transmit LACPDU packets in 802.3ad mode.
//
//   - LACP_RATE_UNSPECIFIED: Default LACP rate is Slow.
//   - LACP_RATE_SLOW: Request LACP partner to transmit LACPDUs every 30 seconds.
//   - LACP_RATE_FAST: Request LACP partner to transmit LACPDUs every 1 second.
//
// swagger:model configLacpRate
type ConfigLacpRate string

func NewConfigLacpRate(value ConfigLacpRate) *ConfigLacpRate {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ConfigLacpRate.
func (m ConfigLacpRate) Pointer() *ConfigLacpRate {
	return &m
}

const (

	// ConfigLacpRateLACPRATEUNSPECIFIED captures enum value "LACP_RATE_UNSPECIFIED"
	ConfigLacpRateLACPRATEUNSPECIFIED ConfigLacpRate = "LACP_RATE_UNSPECIFIED"

	// ConfigLacpRateLACPRATESLOW captures enum value "LACP_RATE_SLOW"
	ConfigLacpRateLACPRATESLOW ConfigLacpRate = "LACP_RATE_SLOW"

	// ConfigLacpRateLACPRATEFAST captures enum value "LACP_RATE_FAST"
	ConfigLacpRateLACPRATEFAST ConfigLacpRate = "LACP_RATE_FAST"
)

// for schema
var configLacpRateEnum []interface{}

func init() {
	var res []ConfigLacpRate
	if err := json.Unmarshal([]byte(`["LACP_RATE_UNSPECIFIED","LACP_RATE_SLOW","LACP_RATE_FAST"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configLacpRateEnum = append(configLacpRateEnum, v)
	}
}

func (m ConfigLacpRate) validateConfigLacpRateEnum(path, location string, value ConfigLacpRate) error {
	if err := validate.EnumCase(path, location, value, configLacpRateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this config lacp rate
func (m ConfigLacpRate) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateConfigLacpRateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this config lacp rate based on context it is used
func (m ConfigLacpRate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
