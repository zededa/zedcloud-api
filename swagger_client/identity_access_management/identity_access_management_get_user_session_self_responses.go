// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
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

// IdentityAccessManagementGetUserSessionSelfReader is a Reader for the IdentityAccessManagementGetUserSessionSelf structure.
type IdentityAccessManagementGetUserSessionSelfReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementGetUserSessionSelfReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementGetUserSessionSelfOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIdentityAccessManagementGetUserSessionSelfUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementGetUserSessionSelfForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementGetUserSessionSelfInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementGetUserSessionSelfGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementGetUserSessionSelfDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementGetUserSessionSelfOK creates a IdentityAccessManagementGetUserSessionSelfOK with default headers values
func NewIdentityAccessManagementGetUserSessionSelfOK() *IdentityAccessManagementGetUserSessionSelfOK {
	return &IdentityAccessManagementGetUserSessionSelfOK{}
}

/*
IdentityAccessManagementGetUserSessionSelfOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementGetUserSessionSelfOK struct {
	Payload *swagger_models.AAAFrontendSessionDetailsResponse
}

// IsSuccess returns true when this identity access management get user session self o k response has a 2xx status code
func (o *IdentityAccessManagementGetUserSessionSelfOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this identity access management get user session self o k response has a 3xx status code
func (o *IdentityAccessManagementGetUserSessionSelfOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get user session self o k response has a 4xx status code
func (o *IdentityAccessManagementGetUserSessionSelfOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management get user session self o k response has a 5xx status code
func (o *IdentityAccessManagementGetUserSessionSelfOK) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get user session self o k response a status code equal to that given
func (o *IdentityAccessManagementGetUserSessionSelfOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the identity access management get user session self o k response
func (o *IdentityAccessManagementGetUserSessionSelfOK) Code() int {
	return 200
}

func (o *IdentityAccessManagementGetUserSessionSelfOK) Error() string {
	return fmt.Sprintf("[GET /v1/sessions/self][%d] identityAccessManagementGetUserSessionSelfOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementGetUserSessionSelfOK) String() string {
	return fmt.Sprintf("[GET /v1/sessions/self][%d] identityAccessManagementGetUserSessionSelfOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementGetUserSessionSelfOK) GetPayload() *swagger_models.AAAFrontendSessionDetailsResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetUserSessionSelfOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.AAAFrontendSessionDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetUserSessionSelfUnauthorized creates a IdentityAccessManagementGetUserSessionSelfUnauthorized with default headers values
func NewIdentityAccessManagementGetUserSessionSelfUnauthorized() *IdentityAccessManagementGetUserSessionSelfUnauthorized {
	return &IdentityAccessManagementGetUserSessionSelfUnauthorized{}
}

/*
IdentityAccessManagementGetUserSessionSelfUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementGetUserSessionSelfUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get user session self unauthorized response has a 2xx status code
func (o *IdentityAccessManagementGetUserSessionSelfUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get user session self unauthorized response has a 3xx status code
func (o *IdentityAccessManagementGetUserSessionSelfUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get user session self unauthorized response has a 4xx status code
func (o *IdentityAccessManagementGetUserSessionSelfUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management get user session self unauthorized response has a 5xx status code
func (o *IdentityAccessManagementGetUserSessionSelfUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get user session self unauthorized response a status code equal to that given
func (o *IdentityAccessManagementGetUserSessionSelfUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the identity access management get user session self unauthorized response
func (o *IdentityAccessManagementGetUserSessionSelfUnauthorized) Code() int {
	return 401
}

func (o *IdentityAccessManagementGetUserSessionSelfUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/sessions/self][%d] identityAccessManagementGetUserSessionSelfUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementGetUserSessionSelfUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/sessions/self][%d] identityAccessManagementGetUserSessionSelfUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementGetUserSessionSelfUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetUserSessionSelfUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetUserSessionSelfForbidden creates a IdentityAccessManagementGetUserSessionSelfForbidden with default headers values
func NewIdentityAccessManagementGetUserSessionSelfForbidden() *IdentityAccessManagementGetUserSessionSelfForbidden {
	return &IdentityAccessManagementGetUserSessionSelfForbidden{}
}

/*
IdentityAccessManagementGetUserSessionSelfForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementGetUserSessionSelfForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get user session self forbidden response has a 2xx status code
func (o *IdentityAccessManagementGetUserSessionSelfForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get user session self forbidden response has a 3xx status code
func (o *IdentityAccessManagementGetUserSessionSelfForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get user session self forbidden response has a 4xx status code
func (o *IdentityAccessManagementGetUserSessionSelfForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management get user session self forbidden response has a 5xx status code
func (o *IdentityAccessManagementGetUserSessionSelfForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get user session self forbidden response a status code equal to that given
func (o *IdentityAccessManagementGetUserSessionSelfForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the identity access management get user session self forbidden response
func (o *IdentityAccessManagementGetUserSessionSelfForbidden) Code() int {
	return 403
}

func (o *IdentityAccessManagementGetUserSessionSelfForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/sessions/self][%d] identityAccessManagementGetUserSessionSelfForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementGetUserSessionSelfForbidden) String() string {
	return fmt.Sprintf("[GET /v1/sessions/self][%d] identityAccessManagementGetUserSessionSelfForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementGetUserSessionSelfForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetUserSessionSelfForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetUserSessionSelfInternalServerError creates a IdentityAccessManagementGetUserSessionSelfInternalServerError with default headers values
func NewIdentityAccessManagementGetUserSessionSelfInternalServerError() *IdentityAccessManagementGetUserSessionSelfInternalServerError {
	return &IdentityAccessManagementGetUserSessionSelfInternalServerError{}
}

/*
IdentityAccessManagementGetUserSessionSelfInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementGetUserSessionSelfInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get user session self internal server error response has a 2xx status code
func (o *IdentityAccessManagementGetUserSessionSelfInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get user session self internal server error response has a 3xx status code
func (o *IdentityAccessManagementGetUserSessionSelfInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get user session self internal server error response has a 4xx status code
func (o *IdentityAccessManagementGetUserSessionSelfInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management get user session self internal server error response has a 5xx status code
func (o *IdentityAccessManagementGetUserSessionSelfInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management get user session self internal server error response a status code equal to that given
func (o *IdentityAccessManagementGetUserSessionSelfInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the identity access management get user session self internal server error response
func (o *IdentityAccessManagementGetUserSessionSelfInternalServerError) Code() int {
	return 500
}

func (o *IdentityAccessManagementGetUserSessionSelfInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/sessions/self][%d] identityAccessManagementGetUserSessionSelfInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementGetUserSessionSelfInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/sessions/self][%d] identityAccessManagementGetUserSessionSelfInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementGetUserSessionSelfInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetUserSessionSelfInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetUserSessionSelfGatewayTimeout creates a IdentityAccessManagementGetUserSessionSelfGatewayTimeout with default headers values
func NewIdentityAccessManagementGetUserSessionSelfGatewayTimeout() *IdentityAccessManagementGetUserSessionSelfGatewayTimeout {
	return &IdentityAccessManagementGetUserSessionSelfGatewayTimeout{}
}

/*
IdentityAccessManagementGetUserSessionSelfGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementGetUserSessionSelfGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get user session self gateway timeout response has a 2xx status code
func (o *IdentityAccessManagementGetUserSessionSelfGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get user session self gateway timeout response has a 3xx status code
func (o *IdentityAccessManagementGetUserSessionSelfGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get user session self gateway timeout response has a 4xx status code
func (o *IdentityAccessManagementGetUserSessionSelfGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management get user session self gateway timeout response has a 5xx status code
func (o *IdentityAccessManagementGetUserSessionSelfGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management get user session self gateway timeout response a status code equal to that given
func (o *IdentityAccessManagementGetUserSessionSelfGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the identity access management get user session self gateway timeout response
func (o *IdentityAccessManagementGetUserSessionSelfGatewayTimeout) Code() int {
	return 504
}

func (o *IdentityAccessManagementGetUserSessionSelfGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/sessions/self][%d] identityAccessManagementGetUserSessionSelfGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementGetUserSessionSelfGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/sessions/self][%d] identityAccessManagementGetUserSessionSelfGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementGetUserSessionSelfGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetUserSessionSelfGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetUserSessionSelfDefault creates a IdentityAccessManagementGetUserSessionSelfDefault with default headers values
func NewIdentityAccessManagementGetUserSessionSelfDefault(code int) *IdentityAccessManagementGetUserSessionSelfDefault {
	return &IdentityAccessManagementGetUserSessionSelfDefault{
		_statusCode: code,
	}
}

/*
IdentityAccessManagementGetUserSessionSelfDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementGetUserSessionSelfDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// IsSuccess returns true when this identity access management get user session self default response has a 2xx status code
func (o *IdentityAccessManagementGetUserSessionSelfDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this identity access management get user session self default response has a 3xx status code
func (o *IdentityAccessManagementGetUserSessionSelfDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this identity access management get user session self default response has a 4xx status code
func (o *IdentityAccessManagementGetUserSessionSelfDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this identity access management get user session self default response has a 5xx status code
func (o *IdentityAccessManagementGetUserSessionSelfDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this identity access management get user session self default response a status code equal to that given
func (o *IdentityAccessManagementGetUserSessionSelfDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the identity access management get user session self default response
func (o *IdentityAccessManagementGetUserSessionSelfDefault) Code() int {
	return o._statusCode
}

func (o *IdentityAccessManagementGetUserSessionSelfDefault) Error() string {
	return fmt.Sprintf("[GET /v1/sessions/self][%d] IdentityAccessManagement_GetUserSessionSelf default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementGetUserSessionSelfDefault) String() string {
	return fmt.Sprintf("[GET /v1/sessions/self][%d] IdentityAccessManagement_GetUserSessionSelf default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementGetUserSessionSelfDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementGetUserSessionSelfDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
