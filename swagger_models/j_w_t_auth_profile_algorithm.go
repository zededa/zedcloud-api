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

// JWTAuthProfileAlgorithm j w t auth profile algorithm
//
// swagger:model JWTAuthProfileAlgorithm
type JWTAuthProfileAlgorithm string

func NewJWTAuthProfileAlgorithm(value JWTAuthProfileAlgorithm) *JWTAuthProfileAlgorithm {
	return &value
}

// Pointer returns a pointer to a freshly-allocated JWTAuthProfileAlgorithm.
func (m JWTAuthProfileAlgorithm) Pointer() *JWTAuthProfileAlgorithm {
	return &m
}

const (

	// JWTAuthProfileAlgorithmNone captures enum value "None"
	JWTAuthProfileAlgorithmNone JWTAuthProfileAlgorithm = "None"

	// JWTAuthProfileAlgorithmRS256 captures enum value "RS256"
	JWTAuthProfileAlgorithmRS256 JWTAuthProfileAlgorithm = "RS256"

	// JWTAuthProfileAlgorithmRS384 captures enum value "RS384"
	JWTAuthProfileAlgorithmRS384 JWTAuthProfileAlgorithm = "RS384"

	// JWTAuthProfileAlgorithmRS512 captures enum value "RS512"
	JWTAuthProfileAlgorithmRS512 JWTAuthProfileAlgorithm = "RS512"
)

// for schema
var jWTAuthProfileAlgorithmEnum []interface{}

func init() {
	var res []JWTAuthProfileAlgorithm
	if err := json.Unmarshal([]byte(`["None","RS256","RS384","RS512"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		jWTAuthProfileAlgorithmEnum = append(jWTAuthProfileAlgorithmEnum, v)
	}
}

func (m JWTAuthProfileAlgorithm) validateJWTAuthProfileAlgorithmEnum(path, location string, value JWTAuthProfileAlgorithm) error {
	if err := validate.EnumCase(path, location, value, jWTAuthProfileAlgorithmEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this j w t auth profile algorithm
func (m JWTAuthProfileAlgorithm) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateJWTAuthProfileAlgorithmEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this j w t auth profile algorithm based on context it is used
func (m JWTAuthProfileAlgorithm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
