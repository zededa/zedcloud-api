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

// IdentityAccessManagementGetUserSelfReader is a Reader for the IdentityAccessManagementGetUserSelf structure.
type IdentityAccessManagementGetUserSelfReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementGetUserSelfReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementGetUserSelfOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIdentityAccessManagementGetUserSelfUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementGetUserSelfForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIdentityAccessManagementGetUserSelfNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementGetUserSelfInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementGetUserSelfGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementGetUserSelfDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementGetUserSelfOK creates a IdentityAccessManagementGetUserSelfOK with default headers values
func NewIdentityAccessManagementGetUserSelfOK() *IdentityAccessManagementGetUserSelfOK {
	return &IdentityAccessManagementGetUserSelfOK{}
}

/*
IdentityAccessManagementGetUserSelfOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementGetUserSelfOK struct {
	Payload *swagger_models.CrudResponseRead
}

// IsSuccess returns true when this identity access management get user self o k response has a 2xx status code
func (o *IdentityAccessManagementGetUserSelfOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this identity access management get user self o k response has a 3xx status code
func (o *IdentityAccessManagementGetUserSelfOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get user self o k response has a 4xx status code
func (o *IdentityAccessManagementGetUserSelfOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management get user self o k response has a 5xx status code
func (o *IdentityAccessManagementGetUserSelfOK) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get user self o k response a status code equal to that given
func (o *IdentityAccessManagementGetUserSelfOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the identity access management get user self o k response
func (o *IdentityAccessManagementGetUserSelfOK) Code() int {
	return 200
}

func (o *IdentityAccessManagementGetUserSelfOK) Error() string {
	return fmt.Sprintf("[GET /v1/users/self][%d] identityAccessManagementGetUserSelfOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementGetUserSelfOK) String() string {
	return fmt.Sprintf("[GET /v1/users/self][%d] identityAccessManagementGetUserSelfOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementGetUserSelfOK) GetPayload() *swagger_models.CrudResponseRead {
	return o.Payload
}

func (o *IdentityAccessManagementGetUserSelfOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.CrudResponseRead)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetUserSelfUnauthorized creates a IdentityAccessManagementGetUserSelfUnauthorized with default headers values
func NewIdentityAccessManagementGetUserSelfUnauthorized() *IdentityAccessManagementGetUserSelfUnauthorized {
	return &IdentityAccessManagementGetUserSelfUnauthorized{}
}

/*
IdentityAccessManagementGetUserSelfUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementGetUserSelfUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get user self unauthorized response has a 2xx status code
func (o *IdentityAccessManagementGetUserSelfUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get user self unauthorized response has a 3xx status code
func (o *IdentityAccessManagementGetUserSelfUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get user self unauthorized response has a 4xx status code
func (o *IdentityAccessManagementGetUserSelfUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management get user self unauthorized response has a 5xx status code
func (o *IdentityAccessManagementGetUserSelfUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get user self unauthorized response a status code equal to that given
func (o *IdentityAccessManagementGetUserSelfUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the identity access management get user self unauthorized response
func (o *IdentityAccessManagementGetUserSelfUnauthorized) Code() int {
	return 401
}

func (o *IdentityAccessManagementGetUserSelfUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/users/self][%d] identityAccessManagementGetUserSelfUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementGetUserSelfUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/users/self][%d] identityAccessManagementGetUserSelfUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementGetUserSelfUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetUserSelfUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetUserSelfForbidden creates a IdentityAccessManagementGetUserSelfForbidden with default headers values
func NewIdentityAccessManagementGetUserSelfForbidden() *IdentityAccessManagementGetUserSelfForbidden {
	return &IdentityAccessManagementGetUserSelfForbidden{}
}

/*
IdentityAccessManagementGetUserSelfForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementGetUserSelfForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get user self forbidden response has a 2xx status code
func (o *IdentityAccessManagementGetUserSelfForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get user self forbidden response has a 3xx status code
func (o *IdentityAccessManagementGetUserSelfForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get user self forbidden response has a 4xx status code
func (o *IdentityAccessManagementGetUserSelfForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management get user self forbidden response has a 5xx status code
func (o *IdentityAccessManagementGetUserSelfForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get user self forbidden response a status code equal to that given
func (o *IdentityAccessManagementGetUserSelfForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the identity access management get user self forbidden response
func (o *IdentityAccessManagementGetUserSelfForbidden) Code() int {
	return 403
}

func (o *IdentityAccessManagementGetUserSelfForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/users/self][%d] identityAccessManagementGetUserSelfForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementGetUserSelfForbidden) String() string {
	return fmt.Sprintf("[GET /v1/users/self][%d] identityAccessManagementGetUserSelfForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementGetUserSelfForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetUserSelfForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetUserSelfNotFound creates a IdentityAccessManagementGetUserSelfNotFound with default headers values
func NewIdentityAccessManagementGetUserSelfNotFound() *IdentityAccessManagementGetUserSelfNotFound {
	return &IdentityAccessManagementGetUserSelfNotFound{}
}

/*
IdentityAccessManagementGetUserSelfNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type IdentityAccessManagementGetUserSelfNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get user self not found response has a 2xx status code
func (o *IdentityAccessManagementGetUserSelfNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get user self not found response has a 3xx status code
func (o *IdentityAccessManagementGetUserSelfNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get user self not found response has a 4xx status code
func (o *IdentityAccessManagementGetUserSelfNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management get user self not found response has a 5xx status code
func (o *IdentityAccessManagementGetUserSelfNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get user self not found response a status code equal to that given
func (o *IdentityAccessManagementGetUserSelfNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the identity access management get user self not found response
func (o *IdentityAccessManagementGetUserSelfNotFound) Code() int {
	return 404
}

func (o *IdentityAccessManagementGetUserSelfNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/users/self][%d] identityAccessManagementGetUserSelfNotFound  %+v", 404, o.Payload)
}

func (o *IdentityAccessManagementGetUserSelfNotFound) String() string {
	return fmt.Sprintf("[GET /v1/users/self][%d] identityAccessManagementGetUserSelfNotFound  %+v", 404, o.Payload)
}

func (o *IdentityAccessManagementGetUserSelfNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetUserSelfNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetUserSelfInternalServerError creates a IdentityAccessManagementGetUserSelfInternalServerError with default headers values
func NewIdentityAccessManagementGetUserSelfInternalServerError() *IdentityAccessManagementGetUserSelfInternalServerError {
	return &IdentityAccessManagementGetUserSelfInternalServerError{}
}

/*
IdentityAccessManagementGetUserSelfInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementGetUserSelfInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get user self internal server error response has a 2xx status code
func (o *IdentityAccessManagementGetUserSelfInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get user self internal server error response has a 3xx status code
func (o *IdentityAccessManagementGetUserSelfInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get user self internal server error response has a 4xx status code
func (o *IdentityAccessManagementGetUserSelfInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management get user self internal server error response has a 5xx status code
func (o *IdentityAccessManagementGetUserSelfInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management get user self internal server error response a status code equal to that given
func (o *IdentityAccessManagementGetUserSelfInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the identity access management get user self internal server error response
func (o *IdentityAccessManagementGetUserSelfInternalServerError) Code() int {
	return 500
}

func (o *IdentityAccessManagementGetUserSelfInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/users/self][%d] identityAccessManagementGetUserSelfInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementGetUserSelfInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/users/self][%d] identityAccessManagementGetUserSelfInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementGetUserSelfInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetUserSelfInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetUserSelfGatewayTimeout creates a IdentityAccessManagementGetUserSelfGatewayTimeout with default headers values
func NewIdentityAccessManagementGetUserSelfGatewayTimeout() *IdentityAccessManagementGetUserSelfGatewayTimeout {
	return &IdentityAccessManagementGetUserSelfGatewayTimeout{}
}

/*
IdentityAccessManagementGetUserSelfGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementGetUserSelfGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get user self gateway timeout response has a 2xx status code
func (o *IdentityAccessManagementGetUserSelfGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get user self gateway timeout response has a 3xx status code
func (o *IdentityAccessManagementGetUserSelfGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get user self gateway timeout response has a 4xx status code
func (o *IdentityAccessManagementGetUserSelfGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management get user self gateway timeout response has a 5xx status code
func (o *IdentityAccessManagementGetUserSelfGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management get user self gateway timeout response a status code equal to that given
func (o *IdentityAccessManagementGetUserSelfGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the identity access management get user self gateway timeout response
func (o *IdentityAccessManagementGetUserSelfGatewayTimeout) Code() int {
	return 504
}

func (o *IdentityAccessManagementGetUserSelfGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/users/self][%d] identityAccessManagementGetUserSelfGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementGetUserSelfGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/users/self][%d] identityAccessManagementGetUserSelfGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementGetUserSelfGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetUserSelfGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetUserSelfDefault creates a IdentityAccessManagementGetUserSelfDefault with default headers values
func NewIdentityAccessManagementGetUserSelfDefault(code int) *IdentityAccessManagementGetUserSelfDefault {
	return &IdentityAccessManagementGetUserSelfDefault{
		_statusCode: code,
	}
}

/*
IdentityAccessManagementGetUserSelfDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementGetUserSelfDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// IsSuccess returns true when this identity access management get user self default response has a 2xx status code
func (o *IdentityAccessManagementGetUserSelfDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this identity access management get user self default response has a 3xx status code
func (o *IdentityAccessManagementGetUserSelfDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this identity access management get user self default response has a 4xx status code
func (o *IdentityAccessManagementGetUserSelfDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this identity access management get user self default response has a 5xx status code
func (o *IdentityAccessManagementGetUserSelfDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this identity access management get user self default response a status code equal to that given
func (o *IdentityAccessManagementGetUserSelfDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the identity access management get user self default response
func (o *IdentityAccessManagementGetUserSelfDefault) Code() int {
	return o._statusCode
}

func (o *IdentityAccessManagementGetUserSelfDefault) Error() string {
	return fmt.Sprintf("[GET /v1/users/self][%d] IdentityAccessManagement_GetUserSelf default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementGetUserSelfDefault) String() string {
	return fmt.Sprintf("[GET /v1/users/self][%d] IdentityAccessManagement_GetUserSelf default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementGetUserSelfDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementGetUserSelfDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
