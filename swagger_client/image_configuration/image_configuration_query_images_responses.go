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

// ImageConfigurationQueryImagesReader is a Reader for the ImageConfigurationQueryImages structure.
type ImageConfigurationQueryImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageConfigurationQueryImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageConfigurationQueryImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImageConfigurationQueryImagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImageConfigurationQueryImagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImageConfigurationQueryImagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageConfigurationQueryImagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewImageConfigurationQueryImagesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewImageConfigurationQueryImagesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewImageConfigurationQueryImagesOK creates a ImageConfigurationQueryImagesOK with default headers values
func NewImageConfigurationQueryImagesOK() *ImageConfigurationQueryImagesOK {
	return &ImageConfigurationQueryImagesOK{}
}

/* ImageConfigurationQueryImagesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ImageConfigurationQueryImagesOK struct {
	Payload *swagger_models.Images
}

func (o *ImageConfigurationQueryImagesOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images][%d] imageConfigurationQueryImagesOK  %+v", 200, o.Payload)
}
func (o *ImageConfigurationQueryImagesOK) GetPayload() *swagger_models.Images {
	return o.Payload
}

func (o *ImageConfigurationQueryImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.Images)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationQueryImagesBadRequest creates a ImageConfigurationQueryImagesBadRequest with default headers values
func NewImageConfigurationQueryImagesBadRequest() *ImageConfigurationQueryImagesBadRequest {
	return &ImageConfigurationQueryImagesBadRequest{}
}

/* ImageConfigurationQueryImagesBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type ImageConfigurationQueryImagesBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ImageConfigurationQueryImagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images][%d] imageConfigurationQueryImagesBadRequest  %+v", 400, o.Payload)
}
func (o *ImageConfigurationQueryImagesBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationQueryImagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationQueryImagesUnauthorized creates a ImageConfigurationQueryImagesUnauthorized with default headers values
func NewImageConfigurationQueryImagesUnauthorized() *ImageConfigurationQueryImagesUnauthorized {
	return &ImageConfigurationQueryImagesUnauthorized{}
}

/* ImageConfigurationQueryImagesUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ImageConfigurationQueryImagesUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ImageConfigurationQueryImagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images][%d] imageConfigurationQueryImagesUnauthorized  %+v", 401, o.Payload)
}
func (o *ImageConfigurationQueryImagesUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationQueryImagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationQueryImagesForbidden creates a ImageConfigurationQueryImagesForbidden with default headers values
func NewImageConfigurationQueryImagesForbidden() *ImageConfigurationQueryImagesForbidden {
	return &ImageConfigurationQueryImagesForbidden{}
}

/* ImageConfigurationQueryImagesForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ImageConfigurationQueryImagesForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ImageConfigurationQueryImagesForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images][%d] imageConfigurationQueryImagesForbidden  %+v", 403, o.Payload)
}
func (o *ImageConfigurationQueryImagesForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationQueryImagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationQueryImagesInternalServerError creates a ImageConfigurationQueryImagesInternalServerError with default headers values
func NewImageConfigurationQueryImagesInternalServerError() *ImageConfigurationQueryImagesInternalServerError {
	return &ImageConfigurationQueryImagesInternalServerError{}
}

/* ImageConfigurationQueryImagesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ImageConfigurationQueryImagesInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ImageConfigurationQueryImagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images][%d] imageConfigurationQueryImagesInternalServerError  %+v", 500, o.Payload)
}
func (o *ImageConfigurationQueryImagesInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationQueryImagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationQueryImagesGatewayTimeout creates a ImageConfigurationQueryImagesGatewayTimeout with default headers values
func NewImageConfigurationQueryImagesGatewayTimeout() *ImageConfigurationQueryImagesGatewayTimeout {
	return &ImageConfigurationQueryImagesGatewayTimeout{}
}

/* ImageConfigurationQueryImagesGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ImageConfigurationQueryImagesGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ImageConfigurationQueryImagesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images][%d] imageConfigurationQueryImagesGatewayTimeout  %+v", 504, o.Payload)
}
func (o *ImageConfigurationQueryImagesGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationQueryImagesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationQueryImagesDefault creates a ImageConfigurationQueryImagesDefault with default headers values
func NewImageConfigurationQueryImagesDefault(code int) *ImageConfigurationQueryImagesDefault {
	return &ImageConfigurationQueryImagesDefault{
		_statusCode: code,
	}
}

/* ImageConfigurationQueryImagesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ImageConfigurationQueryImagesDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the image configuration query images default response
func (o *ImageConfigurationQueryImagesDefault) Code() int {
	return o._statusCode
}

func (o *ImageConfigurationQueryImagesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images][%d] ImageConfiguration_QueryImages default  %+v", o._statusCode, o.Payload)
}
func (o *ImageConfigurationQueryImagesDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *ImageConfigurationQueryImagesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}