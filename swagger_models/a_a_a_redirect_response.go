// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AAARedirectResponse a a a redirect response
//
// swagger:model AAARedirectResponse
type AAARedirectResponse struct {

	// code
	Code int64 `json:"code,omitempty"`

	// redirect Url
	RedirectURL string `json:"redirectUrl,omitempty"`
}

// Validate validates this a a a redirect response
func (m *AAARedirectResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this a a a redirect response based on context it is used
func (m *AAARedirectResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AAARedirectResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AAARedirectResponse) UnmarshalBinary(b []byte) error {
	var res AAARedirectResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
