// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package identity_access_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// GetUserSelfReader is a Reader for the GetUserSelf structure.
type GetUserSelfReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserSelfReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserSelfOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUserSelfUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserSelfForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserSelfNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserSelfInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUserSelfGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserSelfOK creates a GetUserSelfOK with default headers values
func NewGetUserSelfOK() *GetUserSelfOK {
	return &GetUserSelfOK{}
}

/* GetUserSelfOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetUserSelfOK struct {
	Payload *swagger_models.CrudResponseRead
}

func (o *GetUserSelfOK) Error() string {
	return fmt.Sprintf("[GET /v1/users/self][%d] getUserSelfOK  %+v", 200, o.Payload)
}
func (o *GetUserSelfOK) GetPayload() *swagger_models.CrudResponseRead {
	return o.Payload
}

func (o *GetUserSelfOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.CrudResponseRead)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserSelfUnauthorized creates a GetUserSelfUnauthorized with default headers values
func NewGetUserSelfUnauthorized() *GetUserSelfUnauthorized {
	return &GetUserSelfUnauthorized{}
}

/* GetUserSelfUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetUserSelfUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetUserSelfUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/users/self][%d] getUserSelfUnauthorized  %+v", 401, o.Payload)
}
func (o *GetUserSelfUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetUserSelfUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserSelfForbidden creates a GetUserSelfForbidden with default headers values
func NewGetUserSelfForbidden() *GetUserSelfForbidden {
	return &GetUserSelfForbidden{}
}

/* GetUserSelfForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type GetUserSelfForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetUserSelfForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/users/self][%d] getUserSelfForbidden  %+v", 403, o.Payload)
}
func (o *GetUserSelfForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetUserSelfForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserSelfNotFound creates a GetUserSelfNotFound with default headers values
func NewGetUserSelfNotFound() *GetUserSelfNotFound {
	return &GetUserSelfNotFound{}
}

/* GetUserSelfNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetUserSelfNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetUserSelfNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/users/self][%d] getUserSelfNotFound  %+v", 404, o.Payload)
}
func (o *GetUserSelfNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetUserSelfNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserSelfInternalServerError creates a GetUserSelfInternalServerError with default headers values
func NewGetUserSelfInternalServerError() *GetUserSelfInternalServerError {
	return &GetUserSelfInternalServerError{}
}

/* GetUserSelfInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetUserSelfInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetUserSelfInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/users/self][%d] getUserSelfInternalServerError  %+v", 500, o.Payload)
}
func (o *GetUserSelfInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetUserSelfInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserSelfGatewayTimeout creates a GetUserSelfGatewayTimeout with default headers values
func NewGetUserSelfGatewayTimeout() *GetUserSelfGatewayTimeout {
	return &GetUserSelfGatewayTimeout{}
}

/* GetUserSelfGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetUserSelfGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetUserSelfGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/users/self][%d] getUserSelfGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetUserSelfGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetUserSelfGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
