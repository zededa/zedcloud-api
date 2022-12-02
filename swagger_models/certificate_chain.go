// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CertificateChain certificate chain
//
// swagger:model CertificateChain
type CertificateChain struct {

	// List of Certificate object holds the details of certificate like cert block, encryption type, validity, subject etc
	Certificates []*Certificate `json:"certificates"`
}

// Validate validates this certificate chain
func (m *CertificateChain) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificateChain) validateCertificates(formats strfmt.Registry) error {
	if swag.IsZero(m.Certificates) { // not required
		return nil
	}

	for i := 0; i < len(m.Certificates); i++ {
		if swag.IsZero(m.Certificates[i]) { // not required
			continue
		}

		if m.Certificates[i] != nil {
			if err := m.Certificates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("certificates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("certificates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this certificate chain based on the context it is used
func (m *CertificateChain) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCertificates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificateChain) contextValidateCertificates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Certificates); i++ {

		if m.Certificates[i] != nil {
			if err := m.Certificates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("certificates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("certificates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificateChain) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateChain) UnmarshalBinary(b []byte) error {
	var res CertificateChain
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
