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

// ImageStatus Image status
//
// - IMAGE_STATUS_CREATED: Image metadata is created
//   - IMAGE_STATUS_UPLOADING: Image binary is being uploaded to Datstore
//   - IMAGE_STATUS_READY: Image is ready for download
//   - IMAGE_STATUS_INUSE: Image is being used by edge applications
//   - IMAGE_STATUS_FAILED: Image binary upload has failed
//   - IMAGE_STATUS_UPLINKING: Image metadata is being uplinked with Datstore binary
//
// swagger:model ImageStatus
type ImageStatus string

func NewImageStatus(value ImageStatus) *ImageStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ImageStatus.
func (m ImageStatus) Pointer() *ImageStatus {
	return &m
}

const (

	// ImageStatusIMAGESTATUSUNSPECIFIED captures enum value "IMAGE_STATUS_UNSPECIFIED"
	ImageStatusIMAGESTATUSUNSPECIFIED ImageStatus = "IMAGE_STATUS_UNSPECIFIED"

	// ImageStatusIMAGESTATUSCREATED captures enum value "IMAGE_STATUS_CREATED"
	ImageStatusIMAGESTATUSCREATED ImageStatus = "IMAGE_STATUS_CREATED"

	// ImageStatusIMAGESTATUSPENDING captures enum value "IMAGE_STATUS_PENDING"
	ImageStatusIMAGESTATUSPENDING ImageStatus = "IMAGE_STATUS_PENDING"

	// ImageStatusIMAGESTATUSUPLOADING captures enum value "IMAGE_STATUS_UPLOADING"
	ImageStatusIMAGESTATUSUPLOADING ImageStatus = "IMAGE_STATUS_UPLOADING"

	// ImageStatusIMAGESTATUSREADY captures enum value "IMAGE_STATUS_READY"
	ImageStatusIMAGESTATUSREADY ImageStatus = "IMAGE_STATUS_READY"

	// ImageStatusIMAGESTATUSINUSE captures enum value "IMAGE_STATUS_INUSE"
	ImageStatusIMAGESTATUSINUSE ImageStatus = "IMAGE_STATUS_INUSE"

	// ImageStatusIMAGESTATUSARCHIVED captures enum value "IMAGE_STATUS_ARCHIVED"
	ImageStatusIMAGESTATUSARCHIVED ImageStatus = "IMAGE_STATUS_ARCHIVED"

	// ImageStatusIMAGESTATUSMAX captures enum value "IMAGE_STATUS_MAX"
	ImageStatusIMAGESTATUSMAX ImageStatus = "IMAGE_STATUS_MAX"

	// ImageStatusIMAGESTATUSFAILED captures enum value "IMAGE_STATUS_FAILED"
	ImageStatusIMAGESTATUSFAILED ImageStatus = "IMAGE_STATUS_FAILED"

	// ImageStatusIMAGESTATUSUPLINKING captures enum value "IMAGE_STATUS_UPLINKING"
	ImageStatusIMAGESTATUSUPLINKING ImageStatus = "IMAGE_STATUS_UPLINKING"
)

// for schema
var imageStatusEnum []interface{}

func init() {
	var res []ImageStatus
	if err := json.Unmarshal([]byte(`["IMAGE_STATUS_UNSPECIFIED","IMAGE_STATUS_CREATED","IMAGE_STATUS_PENDING","IMAGE_STATUS_UPLOADING","IMAGE_STATUS_READY","IMAGE_STATUS_INUSE","IMAGE_STATUS_ARCHIVED","IMAGE_STATUS_MAX","IMAGE_STATUS_FAILED","IMAGE_STATUS_UPLINKING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		imageStatusEnum = append(imageStatusEnum, v)
	}
}

func (m ImageStatus) validateImageStatusEnum(path, location string, value ImageStatus) error {
	if err := validate.EnumCase(path, location, value, imageStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this image status
func (m ImageStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateImageStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this image status based on context it is used
func (m ImageStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
