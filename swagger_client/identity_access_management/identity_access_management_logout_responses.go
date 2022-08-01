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

// IdentityAccessManagementLogoutReader is a Reader for the IdentityAccessManagementLogout structure.
type IdentityAccessManagementLogoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementLogoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementLogoutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIdentityAccessManagementLogoutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementLogoutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementLogoutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementLogoutGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementLogoutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementLogoutOK creates a IdentityAccessManagementLogoutOK with default headers values
func NewIdentityAccessManagementLogoutOK() *IdentityAccessManagementLogoutOK {
	return &IdentityAccessManagementLogoutOK{}
}

/* IdentityAccessManagementLogoutOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementLogoutOK struct {
	Payload *swagger_models.AAAFrontendLogoutResponse
}

func (o *IdentityAccessManagementLogoutOK) Error() string {
	return fmt.Sprintf("[POST /v1/logout][%d] identityAccessManagementLogoutOK  %+v", 200, o.Payload)
}
func (o *IdentityAccessManagementLogoutOK) GetPayload() *swagger_models.AAAFrontendLogoutResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLogoutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.AAAFrontendLogoutResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLogoutUnauthorized creates a IdentityAccessManagementLogoutUnauthorized with default headers values
func NewIdentityAccessManagementLogoutUnauthorized() *IdentityAccessManagementLogoutUnauthorized {
	return &IdentityAccessManagementLogoutUnauthorized{}
}

/* IdentityAccessManagementLogoutUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementLogoutUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementLogoutUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/logout][%d] identityAccessManagementLogoutUnauthorized  %+v", 401, o.Payload)
}
func (o *IdentityAccessManagementLogoutUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLogoutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLogoutForbidden creates a IdentityAccessManagementLogoutForbidden with default headers values
func NewIdentityAccessManagementLogoutForbidden() *IdentityAccessManagementLogoutForbidden {
	return &IdentityAccessManagementLogoutForbidden{}
}

/* IdentityAccessManagementLogoutForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementLogoutForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementLogoutForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/logout][%d] identityAccessManagementLogoutForbidden  %+v", 403, o.Payload)
}
func (o *IdentityAccessManagementLogoutForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLogoutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLogoutInternalServerError creates a IdentityAccessManagementLogoutInternalServerError with default headers values
func NewIdentityAccessManagementLogoutInternalServerError() *IdentityAccessManagementLogoutInternalServerError {
	return &IdentityAccessManagementLogoutInternalServerError{}
}

/* IdentityAccessManagementLogoutInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementLogoutInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementLogoutInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/logout][%d] identityAccessManagementLogoutInternalServerError  %+v", 500, o.Payload)
}
func (o *IdentityAccessManagementLogoutInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLogoutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLogoutGatewayTimeout creates a IdentityAccessManagementLogoutGatewayTimeout with default headers values
func NewIdentityAccessManagementLogoutGatewayTimeout() *IdentityAccessManagementLogoutGatewayTimeout {
	return &IdentityAccessManagementLogoutGatewayTimeout{}
}

/* IdentityAccessManagementLogoutGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementLogoutGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementLogoutGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/logout][%d] identityAccessManagementLogoutGatewayTimeout  %+v", 504, o.Payload)
}
func (o *IdentityAccessManagementLogoutGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLogoutGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLogoutDefault creates a IdentityAccessManagementLogoutDefault with default headers values
func NewIdentityAccessManagementLogoutDefault(code int) *IdentityAccessManagementLogoutDefault {
	return &IdentityAccessManagementLogoutDefault{
		_statusCode: code,
	}
}

/* IdentityAccessManagementLogoutDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementLogoutDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the identity access management logout default response
func (o *IdentityAccessManagementLogoutDefault) Code() int {
	return o._statusCode
}

func (o *IdentityAccessManagementLogoutDefault) Error() string {
	return fmt.Sprintf("[POST /v1/logout][%d] IdentityAccessManagement_Logout default  %+v", o._statusCode, o.Payload)
}
func (o *IdentityAccessManagementLogoutDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementLogoutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}