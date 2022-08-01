// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package image_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// ImageConfigurationGetLatestImageVersionReader is a Reader for the ImageConfigurationGetLatestImageVersion structure.
type ImageConfigurationGetLatestImageVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageConfigurationGetLatestImageVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageConfigurationGetLatestImageVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewImageConfigurationGetLatestImageVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImageConfigurationGetLatestImageVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImageConfigurationGetLatestImageVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageConfigurationGetLatestImageVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewImageConfigurationGetLatestImageVersionGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewImageConfigurationGetLatestImageVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewImageConfigurationGetLatestImageVersionOK creates a ImageConfigurationGetLatestImageVersionOK with default headers values
func NewImageConfigurationGetLatestImageVersionOK() *ImageConfigurationGetLatestImageVersionOK {
	return &ImageConfigurationGetLatestImageVersionOK{}
}

/* ImageConfigurationGetLatestImageVersionOK describes a response with status code 200, with default header values.

A successful response.
*/
type ImageConfigurationGetLatestImageVersionOK struct {
	Payload *swagger_models.ImageConfig
}

func (o *ImageConfigurationGetLatestImageVersionOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationGetLatestImageVersionOK  %+v", 200, o.Payload)
}
func (o *ImageConfigurationGetLatestImageVersionOK) GetPayload() *swagger_models.ImageConfig {
	return o.Payload
}

func (o *ImageConfigurationGetLatestImageVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ImageConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetLatestImageVersionUnauthorized creates a ImageConfigurationGetLatestImageVersionUnauthorized with default headers values
func NewImageConfigurationGetLatestImageVersionUnauthorized() *ImageConfigurationGetLatestImageVersionUnauthorized {
	return &ImageConfigurationGetLatestImageVersionUnauthorized{}
}

/* ImageConfigurationGetLatestImageVersionUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ImageConfigurationGetLatestImageVersionUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ImageConfigurationGetLatestImageVersionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationGetLatestImageVersionUnauthorized  %+v", 401, o.Payload)
}
func (o *ImageConfigurationGetLatestImageVersionUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationGetLatestImageVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetLatestImageVersionForbidden creates a ImageConfigurationGetLatestImageVersionForbidden with default headers values
func NewImageConfigurationGetLatestImageVersionForbidden() *ImageConfigurationGetLatestImageVersionForbidden {
	return &ImageConfigurationGetLatestImageVersionForbidden{}
}

/* ImageConfigurationGetLatestImageVersionForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ImageConfigurationGetLatestImageVersionForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ImageConfigurationGetLatestImageVersionForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationGetLatestImageVersionForbidden  %+v", 403, o.Payload)
}
func (o *ImageConfigurationGetLatestImageVersionForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationGetLatestImageVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetLatestImageVersionNotFound creates a ImageConfigurationGetLatestImageVersionNotFound with default headers values
func NewImageConfigurationGetLatestImageVersionNotFound() *ImageConfigurationGetLatestImageVersionNotFound {
	return &ImageConfigurationGetLatestImageVersionNotFound{}
}

/* ImageConfigurationGetLatestImageVersionNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type ImageConfigurationGetLatestImageVersionNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ImageConfigurationGetLatestImageVersionNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationGetLatestImageVersionNotFound  %+v", 404, o.Payload)
}
func (o *ImageConfigurationGetLatestImageVersionNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationGetLatestImageVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetLatestImageVersionInternalServerError creates a ImageConfigurationGetLatestImageVersionInternalServerError with default headers values
func NewImageConfigurationGetLatestImageVersionInternalServerError() *ImageConfigurationGetLatestImageVersionInternalServerError {
	return &ImageConfigurationGetLatestImageVersionInternalServerError{}
}

/* ImageConfigurationGetLatestImageVersionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ImageConfigurationGetLatestImageVersionInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ImageConfigurationGetLatestImageVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationGetLatestImageVersionInternalServerError  %+v", 500, o.Payload)
}
func (o *ImageConfigurationGetLatestImageVersionInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationGetLatestImageVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetLatestImageVersionGatewayTimeout creates a ImageConfigurationGetLatestImageVersionGatewayTimeout with default headers values
func NewImageConfigurationGetLatestImageVersionGatewayTimeout() *ImageConfigurationGetLatestImageVersionGatewayTimeout {
	return &ImageConfigurationGetLatestImageVersionGatewayTimeout{}
}

/* ImageConfigurationGetLatestImageVersionGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ImageConfigurationGetLatestImageVersionGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ImageConfigurationGetLatestImageVersionGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] imageConfigurationGetLatestImageVersionGatewayTimeout  %+v", 504, o.Payload)
}
func (o *ImageConfigurationGetLatestImageVersionGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationGetLatestImageVersionGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetLatestImageVersionDefault creates a ImageConfigurationGetLatestImageVersionDefault with default headers values
func NewImageConfigurationGetLatestImageVersionDefault(code int) *ImageConfigurationGetLatestImageVersionDefault {
	return &ImageConfigurationGetLatestImageVersionDefault{
		_statusCode: code,
	}
}

/* ImageConfigurationGetLatestImageVersionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ImageConfigurationGetLatestImageVersionDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the image configuration get latest image version default response
func (o *ImageConfigurationGetLatestImageVersionDefault) Code() int {
	return o._statusCode
}

func (o *ImageConfigurationGetLatestImageVersionDefault) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest/hwclass/{imageArch}][%d] ImageConfiguration_GetLatestImageVersion default  %+v", o._statusCode, o.Payload)
}
func (o *ImageConfigurationGetLatestImageVersionDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *ImageConfigurationGetLatestImageVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}