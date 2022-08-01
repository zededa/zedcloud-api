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

// IdentityAccessManagementLoginForgotPasswordReader is a Reader for the IdentityAccessManagementLoginForgotPassword structure.
type IdentityAccessManagementLoginForgotPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementLoginForgotPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementLoginForgotPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIdentityAccessManagementLoginForgotPasswordUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementLoginForgotPasswordForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementLoginForgotPasswordInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementLoginForgotPasswordGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementLoginForgotPasswordDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementLoginForgotPasswordOK creates a IdentityAccessManagementLoginForgotPasswordOK with default headers values
func NewIdentityAccessManagementLoginForgotPasswordOK() *IdentityAccessManagementLoginForgotPasswordOK {
	return &IdentityAccessManagementLoginForgotPasswordOK{}
}

/* IdentityAccessManagementLoginForgotPasswordOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementLoginForgotPasswordOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementLoginForgotPasswordOK) Error() string {
	return fmt.Sprintf("[POST /v1/login/forgot][%d] identityAccessManagementLoginForgotPasswordOK  %+v", 200, o.Payload)
}
func (o *IdentityAccessManagementLoginForgotPasswordOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLoginForgotPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLoginForgotPasswordUnauthorized creates a IdentityAccessManagementLoginForgotPasswordUnauthorized with default headers values
func NewIdentityAccessManagementLoginForgotPasswordUnauthorized() *IdentityAccessManagementLoginForgotPasswordUnauthorized {
	return &IdentityAccessManagementLoginForgotPasswordUnauthorized{}
}

/* IdentityAccessManagementLoginForgotPasswordUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementLoginForgotPasswordUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementLoginForgotPasswordUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/login/forgot][%d] identityAccessManagementLoginForgotPasswordUnauthorized  %+v", 401, o.Payload)
}
func (o *IdentityAccessManagementLoginForgotPasswordUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLoginForgotPasswordUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLoginForgotPasswordForbidden creates a IdentityAccessManagementLoginForgotPasswordForbidden with default headers values
func NewIdentityAccessManagementLoginForgotPasswordForbidden() *IdentityAccessManagementLoginForgotPasswordForbidden {
	return &IdentityAccessManagementLoginForgotPasswordForbidden{}
}

/* IdentityAccessManagementLoginForgotPasswordForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementLoginForgotPasswordForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementLoginForgotPasswordForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/login/forgot][%d] identityAccessManagementLoginForgotPasswordForbidden  %+v", 403, o.Payload)
}
func (o *IdentityAccessManagementLoginForgotPasswordForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLoginForgotPasswordForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLoginForgotPasswordInternalServerError creates a IdentityAccessManagementLoginForgotPasswordInternalServerError with default headers values
func NewIdentityAccessManagementLoginForgotPasswordInternalServerError() *IdentityAccessManagementLoginForgotPasswordInternalServerError {
	return &IdentityAccessManagementLoginForgotPasswordInternalServerError{}
}

/* IdentityAccessManagementLoginForgotPasswordInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementLoginForgotPasswordInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementLoginForgotPasswordInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/login/forgot][%d] identityAccessManagementLoginForgotPasswordInternalServerError  %+v", 500, o.Payload)
}
func (o *IdentityAccessManagementLoginForgotPasswordInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLoginForgotPasswordInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLoginForgotPasswordGatewayTimeout creates a IdentityAccessManagementLoginForgotPasswordGatewayTimeout with default headers values
func NewIdentityAccessManagementLoginForgotPasswordGatewayTimeout() *IdentityAccessManagementLoginForgotPasswordGatewayTimeout {
	return &IdentityAccessManagementLoginForgotPasswordGatewayTimeout{}
}

/* IdentityAccessManagementLoginForgotPasswordGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementLoginForgotPasswordGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementLoginForgotPasswordGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/login/forgot][%d] identityAccessManagementLoginForgotPasswordGatewayTimeout  %+v", 504, o.Payload)
}
func (o *IdentityAccessManagementLoginForgotPasswordGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLoginForgotPasswordGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLoginForgotPasswordDefault creates a IdentityAccessManagementLoginForgotPasswordDefault with default headers values
func NewIdentityAccessManagementLoginForgotPasswordDefault(code int) *IdentityAccessManagementLoginForgotPasswordDefault {
	return &IdentityAccessManagementLoginForgotPasswordDefault{
		_statusCode: code,
	}
}

/* IdentityAccessManagementLoginForgotPasswordDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementLoginForgotPasswordDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the identity access management login forgot password default response
func (o *IdentityAccessManagementLoginForgotPasswordDefault) Code() int {
	return o._statusCode
}

func (o *IdentityAccessManagementLoginForgotPasswordDefault) Error() string {
	return fmt.Sprintf("[POST /v1/login/forgot][%d] IdentityAccessManagement_LoginForgotPassword default  %+v", o._statusCode, o.Payload)
}
func (o *IdentityAccessManagementLoginForgotPasswordDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementLoginForgotPasswordDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
