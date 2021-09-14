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

// QueryLatestImageVersionsReader is a Reader for the QueryLatestImageVersions structure.
type QueryLatestImageVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryLatestImageVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryLatestImageVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryLatestImageVersionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewQueryLatestImageVersionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryLatestImageVersionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryLatestImageVersionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewQueryLatestImageVersionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryLatestImageVersionsOK creates a QueryLatestImageVersionsOK with default headers values
func NewQueryLatestImageVersionsOK() *QueryLatestImageVersionsOK {
	return &QueryLatestImageVersionsOK{}
}

/* QueryLatestImageVersionsOK describes a response with status code 200, with default header values.

A successful response.
*/
type QueryLatestImageVersionsOK struct {
	Payload *swagger_models.Images
}

func (o *QueryLatestImageVersionsOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest][%d] queryLatestImageVersionsOK  %+v", 200, o.Payload)
}
func (o *QueryLatestImageVersionsOK) GetPayload() *swagger_models.Images {
	return o.Payload
}

func (o *QueryLatestImageVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.Images)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryLatestImageVersionsBadRequest creates a QueryLatestImageVersionsBadRequest with default headers values
func NewQueryLatestImageVersionsBadRequest() *QueryLatestImageVersionsBadRequest {
	return &QueryLatestImageVersionsBadRequest{}
}

/* QueryLatestImageVersionsBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type QueryLatestImageVersionsBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryLatestImageVersionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest][%d] queryLatestImageVersionsBadRequest  %+v", 400, o.Payload)
}
func (o *QueryLatestImageVersionsBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryLatestImageVersionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryLatestImageVersionsUnauthorized creates a QueryLatestImageVersionsUnauthorized with default headers values
func NewQueryLatestImageVersionsUnauthorized() *QueryLatestImageVersionsUnauthorized {
	return &QueryLatestImageVersionsUnauthorized{}
}

/* QueryLatestImageVersionsUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type QueryLatestImageVersionsUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryLatestImageVersionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest][%d] queryLatestImageVersionsUnauthorized  %+v", 401, o.Payload)
}
func (o *QueryLatestImageVersionsUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryLatestImageVersionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryLatestImageVersionsForbidden creates a QueryLatestImageVersionsForbidden with default headers values
func NewQueryLatestImageVersionsForbidden() *QueryLatestImageVersionsForbidden {
	return &QueryLatestImageVersionsForbidden{}
}

/* QueryLatestImageVersionsForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type QueryLatestImageVersionsForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryLatestImageVersionsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest][%d] queryLatestImageVersionsForbidden  %+v", 403, o.Payload)
}
func (o *QueryLatestImageVersionsForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryLatestImageVersionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryLatestImageVersionsInternalServerError creates a QueryLatestImageVersionsInternalServerError with default headers values
func NewQueryLatestImageVersionsInternalServerError() *QueryLatestImageVersionsInternalServerError {
	return &QueryLatestImageVersionsInternalServerError{}
}

/* QueryLatestImageVersionsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type QueryLatestImageVersionsInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryLatestImageVersionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest][%d] queryLatestImageVersionsInternalServerError  %+v", 500, o.Payload)
}
func (o *QueryLatestImageVersionsInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryLatestImageVersionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryLatestImageVersionsGatewayTimeout creates a QueryLatestImageVersionsGatewayTimeout with default headers values
func NewQueryLatestImageVersionsGatewayTimeout() *QueryLatestImageVersionsGatewayTimeout {
	return &QueryLatestImageVersionsGatewayTimeout{}
}

/* QueryLatestImageVersionsGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type QueryLatestImageVersionsGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryLatestImageVersionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest][%d] queryLatestImageVersionsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *QueryLatestImageVersionsGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryLatestImageVersionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
