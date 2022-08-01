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

// ImageConfigurationGetImageByNameReader is a Reader for the ImageConfigurationGetImageByName structure.
type ImageConfigurationGetImageByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageConfigurationGetImageByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageConfigurationGetImageByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewImageConfigurationGetImageByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImageConfigurationGetImageByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImageConfigurationGetImageByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageConfigurationGetImageByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewImageConfigurationGetImageByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewImageConfigurationGetImageByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewImageConfigurationGetImageByNameOK creates a ImageConfigurationGetImageByNameOK with default headers values
func NewImageConfigurationGetImageByNameOK() *ImageConfigurationGetImageByNameOK {
	return &ImageConfigurationGetImageByNameOK{}
}

/* ImageConfigurationGetImageByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type ImageConfigurationGetImageByNameOK struct {
	Payload *swagger_models.ImageConfig
}

func (o *ImageConfigurationGetImageByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/name/{name}][%d] imageConfigurationGetImageByNameOK  %+v", 200, o.Payload)
}
func (o *ImageConfigurationGetImageByNameOK) GetPayload() *swagger_models.ImageConfig {
	return o.Payload
}

func (o *ImageConfigurationGetImageByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ImageConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetImageByNameUnauthorized creates a ImageConfigurationGetImageByNameUnauthorized with default headers values
func NewImageConfigurationGetImageByNameUnauthorized() *ImageConfigurationGetImageByNameUnauthorized {
	return &ImageConfigurationGetImageByNameUnauthorized{}
}

/* ImageConfigurationGetImageByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ImageConfigurationGetImageByNameUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ImageConfigurationGetImageByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/name/{name}][%d] imageConfigurationGetImageByNameUnauthorized  %+v", 401, o.Payload)
}
func (o *ImageConfigurationGetImageByNameUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationGetImageByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetImageByNameForbidden creates a ImageConfigurationGetImageByNameForbidden with default headers values
func NewImageConfigurationGetImageByNameForbidden() *ImageConfigurationGetImageByNameForbidden {
	return &ImageConfigurationGetImageByNameForbidden{}
}

/* ImageConfigurationGetImageByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ImageConfigurationGetImageByNameForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ImageConfigurationGetImageByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/name/{name}][%d] imageConfigurationGetImageByNameForbidden  %+v", 403, o.Payload)
}
func (o *ImageConfigurationGetImageByNameForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationGetImageByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetImageByNameNotFound creates a ImageConfigurationGetImageByNameNotFound with default headers values
func NewImageConfigurationGetImageByNameNotFound() *ImageConfigurationGetImageByNameNotFound {
	return &ImageConfigurationGetImageByNameNotFound{}
}

/* ImageConfigurationGetImageByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type ImageConfigurationGetImageByNameNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ImageConfigurationGetImageByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/name/{name}][%d] imageConfigurationGetImageByNameNotFound  %+v", 404, o.Payload)
}
func (o *ImageConfigurationGetImageByNameNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationGetImageByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetImageByNameInternalServerError creates a ImageConfigurationGetImageByNameInternalServerError with default headers values
func NewImageConfigurationGetImageByNameInternalServerError() *ImageConfigurationGetImageByNameInternalServerError {
	return &ImageConfigurationGetImageByNameInternalServerError{}
}

/* ImageConfigurationGetImageByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ImageConfigurationGetImageByNameInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ImageConfigurationGetImageByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/name/{name}][%d] imageConfigurationGetImageByNameInternalServerError  %+v", 500, o.Payload)
}
func (o *ImageConfigurationGetImageByNameInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationGetImageByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetImageByNameGatewayTimeout creates a ImageConfigurationGetImageByNameGatewayTimeout with default headers values
func NewImageConfigurationGetImageByNameGatewayTimeout() *ImageConfigurationGetImageByNameGatewayTimeout {
	return &ImageConfigurationGetImageByNameGatewayTimeout{}
}

/* ImageConfigurationGetImageByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ImageConfigurationGetImageByNameGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ImageConfigurationGetImageByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/name/{name}][%d] imageConfigurationGetImageByNameGatewayTimeout  %+v", 504, o.Payload)
}
func (o *ImageConfigurationGetImageByNameGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationGetImageByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationGetImageByNameDefault creates a ImageConfigurationGetImageByNameDefault with default headers values
func NewImageConfigurationGetImageByNameDefault(code int) *ImageConfigurationGetImageByNameDefault {
	return &ImageConfigurationGetImageByNameDefault{
		_statusCode: code,
	}
}

/* ImageConfigurationGetImageByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ImageConfigurationGetImageByNameDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the image configuration get image by name default response
func (o *ImageConfigurationGetImageByNameDefault) Code() int {
	return o._statusCode
}

func (o *ImageConfigurationGetImageByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/name/{name}][%d] ImageConfiguration_GetImageByName default  %+v", o._statusCode, o.Payload)
}
func (o *ImageConfigurationGetImageByNameDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *ImageConfigurationGetImageByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
