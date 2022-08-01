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

// ImageConfigurationMarkEveImageLatestReader is a Reader for the ImageConfigurationMarkEveImageLatest structure.
type ImageConfigurationMarkEveImageLatestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageConfigurationMarkEveImageLatestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageConfigurationMarkEveImageLatestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImageConfigurationMarkEveImageLatestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImageConfigurationMarkEveImageLatestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImageConfigurationMarkEveImageLatestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImageConfigurationMarkEveImageLatestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageConfigurationMarkEveImageLatestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewImageConfigurationMarkEveImageLatestGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewImageConfigurationMarkEveImageLatestDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewImageConfigurationMarkEveImageLatestOK creates a ImageConfigurationMarkEveImageLatestOK with default headers values
func NewImageConfigurationMarkEveImageLatestOK() *ImageConfigurationMarkEveImageLatestOK {
	return &ImageConfigurationMarkEveImageLatestOK{}
}

/* ImageConfigurationMarkEveImageLatestOK describes a response with status code 200, with default header values.

A successful response.
*/
type ImageConfigurationMarkEveImageLatestOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ImageConfigurationMarkEveImageLatestOK) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/baseos/latest][%d] imageConfigurationMarkEveImageLatestOK  %+v", 200, o.Payload)
}
func (o *ImageConfigurationMarkEveImageLatestOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationMarkEveImageLatestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationMarkEveImageLatestBadRequest creates a ImageConfigurationMarkEveImageLatestBadRequest with default headers values
func NewImageConfigurationMarkEveImageLatestBadRequest() *ImageConfigurationMarkEveImageLatestBadRequest {
	return &ImageConfigurationMarkEveImageLatestBadRequest{}
}

/* ImageConfigurationMarkEveImageLatestBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type ImageConfigurationMarkEveImageLatestBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ImageConfigurationMarkEveImageLatestBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/baseos/latest][%d] imageConfigurationMarkEveImageLatestBadRequest  %+v", 400, o.Payload)
}
func (o *ImageConfigurationMarkEveImageLatestBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationMarkEveImageLatestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationMarkEveImageLatestUnauthorized creates a ImageConfigurationMarkEveImageLatestUnauthorized with default headers values
func NewImageConfigurationMarkEveImageLatestUnauthorized() *ImageConfigurationMarkEveImageLatestUnauthorized {
	return &ImageConfigurationMarkEveImageLatestUnauthorized{}
}

/* ImageConfigurationMarkEveImageLatestUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ImageConfigurationMarkEveImageLatestUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ImageConfigurationMarkEveImageLatestUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/baseos/latest][%d] imageConfigurationMarkEveImageLatestUnauthorized  %+v", 401, o.Payload)
}
func (o *ImageConfigurationMarkEveImageLatestUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationMarkEveImageLatestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationMarkEveImageLatestForbidden creates a ImageConfigurationMarkEveImageLatestForbidden with default headers values
func NewImageConfigurationMarkEveImageLatestForbidden() *ImageConfigurationMarkEveImageLatestForbidden {
	return &ImageConfigurationMarkEveImageLatestForbidden{}
}

/* ImageConfigurationMarkEveImageLatestForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ImageConfigurationMarkEveImageLatestForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ImageConfigurationMarkEveImageLatestForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/baseos/latest][%d] imageConfigurationMarkEveImageLatestForbidden  %+v", 403, o.Payload)
}
func (o *ImageConfigurationMarkEveImageLatestForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationMarkEveImageLatestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationMarkEveImageLatestNotFound creates a ImageConfigurationMarkEveImageLatestNotFound with default headers values
func NewImageConfigurationMarkEveImageLatestNotFound() *ImageConfigurationMarkEveImageLatestNotFound {
	return &ImageConfigurationMarkEveImageLatestNotFound{}
}

/* ImageConfigurationMarkEveImageLatestNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type ImageConfigurationMarkEveImageLatestNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ImageConfigurationMarkEveImageLatestNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/baseos/latest][%d] imageConfigurationMarkEveImageLatestNotFound  %+v", 404, o.Payload)
}
func (o *ImageConfigurationMarkEveImageLatestNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationMarkEveImageLatestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationMarkEveImageLatestInternalServerError creates a ImageConfigurationMarkEveImageLatestInternalServerError with default headers values
func NewImageConfigurationMarkEveImageLatestInternalServerError() *ImageConfigurationMarkEveImageLatestInternalServerError {
	return &ImageConfigurationMarkEveImageLatestInternalServerError{}
}

/* ImageConfigurationMarkEveImageLatestInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ImageConfigurationMarkEveImageLatestInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ImageConfigurationMarkEveImageLatestInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/baseos/latest][%d] imageConfigurationMarkEveImageLatestInternalServerError  %+v", 500, o.Payload)
}
func (o *ImageConfigurationMarkEveImageLatestInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationMarkEveImageLatestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationMarkEveImageLatestGatewayTimeout creates a ImageConfigurationMarkEveImageLatestGatewayTimeout with default headers values
func NewImageConfigurationMarkEveImageLatestGatewayTimeout() *ImageConfigurationMarkEveImageLatestGatewayTimeout {
	return &ImageConfigurationMarkEveImageLatestGatewayTimeout{}
}

/* ImageConfigurationMarkEveImageLatestGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ImageConfigurationMarkEveImageLatestGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ImageConfigurationMarkEveImageLatestGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/baseos/latest][%d] imageConfigurationMarkEveImageLatestGatewayTimeout  %+v", 504, o.Payload)
}
func (o *ImageConfigurationMarkEveImageLatestGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationMarkEveImageLatestGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationMarkEveImageLatestDefault creates a ImageConfigurationMarkEveImageLatestDefault with default headers values
func NewImageConfigurationMarkEveImageLatestDefault(code int) *ImageConfigurationMarkEveImageLatestDefault {
	return &ImageConfigurationMarkEveImageLatestDefault{
		_statusCode: code,
	}
}

/* ImageConfigurationMarkEveImageLatestDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ImageConfigurationMarkEveImageLatestDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the image configuration mark eve image latest default response
func (o *ImageConfigurationMarkEveImageLatestDefault) Code() int {
	return o._statusCode
}

func (o *ImageConfigurationMarkEveImageLatestDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/baseos/latest][%d] ImageConfiguration_MarkEveImageLatest default  %+v", o._statusCode, o.Payload)
}
func (o *ImageConfigurationMarkEveImageLatestDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *ImageConfigurationMarkEveImageLatestDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
