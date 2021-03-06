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

// GetUserSessionReader is a Reader for the GetUserSession structure.
type GetUserSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserSessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUserSessionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserSessionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserSessionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUserSessionGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserSessionOK creates a GetUserSessionOK with default headers values
func NewGetUserSessionOK() *GetUserSessionOK {
	return &GetUserSessionOK{}
}

/* GetUserSessionOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetUserSessionOK struct {
	Payload *swagger_models.AAAFrontendSessionDetailsResponse
}

func (o *GetUserSessionOK) Error() string {
	return fmt.Sprintf("[GET /v1/sessions/token/{sessionToken.base64}][%d] getUserSessionOK  %+v", 200, o.Payload)
}
func (o *GetUserSessionOK) GetPayload() *swagger_models.AAAFrontendSessionDetailsResponse {
	return o.Payload
}

func (o *GetUserSessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.AAAFrontendSessionDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserSessionUnauthorized creates a GetUserSessionUnauthorized with default headers values
func NewGetUserSessionUnauthorized() *GetUserSessionUnauthorized {
	return &GetUserSessionUnauthorized{}
}

/* GetUserSessionUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetUserSessionUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetUserSessionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/sessions/token/{sessionToken.base64}][%d] getUserSessionUnauthorized  %+v", 401, o.Payload)
}
func (o *GetUserSessionUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetUserSessionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserSessionForbidden creates a GetUserSessionForbidden with default headers values
func NewGetUserSessionForbidden() *GetUserSessionForbidden {
	return &GetUserSessionForbidden{}
}

/* GetUserSessionForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type GetUserSessionForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetUserSessionForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/sessions/token/{sessionToken.base64}][%d] getUserSessionForbidden  %+v", 403, o.Payload)
}
func (o *GetUserSessionForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetUserSessionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserSessionInternalServerError creates a GetUserSessionInternalServerError with default headers values
func NewGetUserSessionInternalServerError() *GetUserSessionInternalServerError {
	return &GetUserSessionInternalServerError{}
}

/* GetUserSessionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetUserSessionInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetUserSessionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/sessions/token/{sessionToken.base64}][%d] getUserSessionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetUserSessionInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetUserSessionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserSessionGatewayTimeout creates a GetUserSessionGatewayTimeout with default headers values
func NewGetUserSessionGatewayTimeout() *GetUserSessionGatewayTimeout {
	return &GetUserSessionGatewayTimeout{}
}

/* GetUserSessionGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetUserSessionGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetUserSessionGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/sessions/token/{sessionToken.base64}][%d] getUserSessionGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetUserSessionGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetUserSessionGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
