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

// IdentityAccessManagementLoginModeReader is a Reader for the IdentityAccessManagementLoginMode structure.
type IdentityAccessManagementLoginModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementLoginModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementLoginModeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIdentityAccessManagementLoginModeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementLoginModeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementLoginModeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementLoginModeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementLoginModeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementLoginModeOK creates a IdentityAccessManagementLoginModeOK with default headers values
func NewIdentityAccessManagementLoginModeOK() *IdentityAccessManagementLoginModeOK {
	return &IdentityAccessManagementLoginModeOK{}
}

/*
IdentityAccessManagementLoginModeOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementLoginModeOK struct {
	Payload *swagger_models.AAAFrontendLoginModeResponse
}

// IsSuccess returns true when this identity access management login mode o k response has a 2xx status code
func (o *IdentityAccessManagementLoginModeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this identity access management login mode o k response has a 3xx status code
func (o *IdentityAccessManagementLoginModeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management login mode o k response has a 4xx status code
func (o *IdentityAccessManagementLoginModeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management login mode o k response has a 5xx status code
func (o *IdentityAccessManagementLoginModeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management login mode o k response a status code equal to that given
func (o *IdentityAccessManagementLoginModeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the identity access management login mode o k response
func (o *IdentityAccessManagementLoginModeOK) Code() int {
	return 200
}

func (o *IdentityAccessManagementLoginModeOK) Error() string {
	return fmt.Sprintf("[POST /v1/login/mode][%d] identityAccessManagementLoginModeOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementLoginModeOK) String() string {
	return fmt.Sprintf("[POST /v1/login/mode][%d] identityAccessManagementLoginModeOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementLoginModeOK) GetPayload() *swagger_models.AAAFrontendLoginModeResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLoginModeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.AAAFrontendLoginModeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLoginModeUnauthorized creates a IdentityAccessManagementLoginModeUnauthorized with default headers values
func NewIdentityAccessManagementLoginModeUnauthorized() *IdentityAccessManagementLoginModeUnauthorized {
	return &IdentityAccessManagementLoginModeUnauthorized{}
}

/*
IdentityAccessManagementLoginModeUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementLoginModeUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management login mode unauthorized response has a 2xx status code
func (o *IdentityAccessManagementLoginModeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management login mode unauthorized response has a 3xx status code
func (o *IdentityAccessManagementLoginModeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management login mode unauthorized response has a 4xx status code
func (o *IdentityAccessManagementLoginModeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management login mode unauthorized response has a 5xx status code
func (o *IdentityAccessManagementLoginModeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management login mode unauthorized response a status code equal to that given
func (o *IdentityAccessManagementLoginModeUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the identity access management login mode unauthorized response
func (o *IdentityAccessManagementLoginModeUnauthorized) Code() int {
	return 401
}

func (o *IdentityAccessManagementLoginModeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/login/mode][%d] identityAccessManagementLoginModeUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementLoginModeUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/login/mode][%d] identityAccessManagementLoginModeUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementLoginModeUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLoginModeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLoginModeForbidden creates a IdentityAccessManagementLoginModeForbidden with default headers values
func NewIdentityAccessManagementLoginModeForbidden() *IdentityAccessManagementLoginModeForbidden {
	return &IdentityAccessManagementLoginModeForbidden{}
}

/*
IdentityAccessManagementLoginModeForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementLoginModeForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management login mode forbidden response has a 2xx status code
func (o *IdentityAccessManagementLoginModeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management login mode forbidden response has a 3xx status code
func (o *IdentityAccessManagementLoginModeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management login mode forbidden response has a 4xx status code
func (o *IdentityAccessManagementLoginModeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management login mode forbidden response has a 5xx status code
func (o *IdentityAccessManagementLoginModeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management login mode forbidden response a status code equal to that given
func (o *IdentityAccessManagementLoginModeForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the identity access management login mode forbidden response
func (o *IdentityAccessManagementLoginModeForbidden) Code() int {
	return 403
}

func (o *IdentityAccessManagementLoginModeForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/login/mode][%d] identityAccessManagementLoginModeForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementLoginModeForbidden) String() string {
	return fmt.Sprintf("[POST /v1/login/mode][%d] identityAccessManagementLoginModeForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementLoginModeForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLoginModeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLoginModeInternalServerError creates a IdentityAccessManagementLoginModeInternalServerError with default headers values
func NewIdentityAccessManagementLoginModeInternalServerError() *IdentityAccessManagementLoginModeInternalServerError {
	return &IdentityAccessManagementLoginModeInternalServerError{}
}

/*
IdentityAccessManagementLoginModeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementLoginModeInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management login mode internal server error response has a 2xx status code
func (o *IdentityAccessManagementLoginModeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management login mode internal server error response has a 3xx status code
func (o *IdentityAccessManagementLoginModeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management login mode internal server error response has a 4xx status code
func (o *IdentityAccessManagementLoginModeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management login mode internal server error response has a 5xx status code
func (o *IdentityAccessManagementLoginModeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management login mode internal server error response a status code equal to that given
func (o *IdentityAccessManagementLoginModeInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the identity access management login mode internal server error response
func (o *IdentityAccessManagementLoginModeInternalServerError) Code() int {
	return 500
}

func (o *IdentityAccessManagementLoginModeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/login/mode][%d] identityAccessManagementLoginModeInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementLoginModeInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/login/mode][%d] identityAccessManagementLoginModeInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementLoginModeInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLoginModeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLoginModeGatewayTimeout creates a IdentityAccessManagementLoginModeGatewayTimeout with default headers values
func NewIdentityAccessManagementLoginModeGatewayTimeout() *IdentityAccessManagementLoginModeGatewayTimeout {
	return &IdentityAccessManagementLoginModeGatewayTimeout{}
}

/*
IdentityAccessManagementLoginModeGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementLoginModeGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management login mode gateway timeout response has a 2xx status code
func (o *IdentityAccessManagementLoginModeGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management login mode gateway timeout response has a 3xx status code
func (o *IdentityAccessManagementLoginModeGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management login mode gateway timeout response has a 4xx status code
func (o *IdentityAccessManagementLoginModeGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management login mode gateway timeout response has a 5xx status code
func (o *IdentityAccessManagementLoginModeGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management login mode gateway timeout response a status code equal to that given
func (o *IdentityAccessManagementLoginModeGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the identity access management login mode gateway timeout response
func (o *IdentityAccessManagementLoginModeGatewayTimeout) Code() int {
	return 504
}

func (o *IdentityAccessManagementLoginModeGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/login/mode][%d] identityAccessManagementLoginModeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementLoginModeGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /v1/login/mode][%d] identityAccessManagementLoginModeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementLoginModeGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLoginModeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLoginModeDefault creates a IdentityAccessManagementLoginModeDefault with default headers values
func NewIdentityAccessManagementLoginModeDefault(code int) *IdentityAccessManagementLoginModeDefault {
	return &IdentityAccessManagementLoginModeDefault{
		_statusCode: code,
	}
}

/*
IdentityAccessManagementLoginModeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementLoginModeDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// IsSuccess returns true when this identity access management login mode default response has a 2xx status code
func (o *IdentityAccessManagementLoginModeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this identity access management login mode default response has a 3xx status code
func (o *IdentityAccessManagementLoginModeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this identity access management login mode default response has a 4xx status code
func (o *IdentityAccessManagementLoginModeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this identity access management login mode default response has a 5xx status code
func (o *IdentityAccessManagementLoginModeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this identity access management login mode default response a status code equal to that given
func (o *IdentityAccessManagementLoginModeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the identity access management login mode default response
func (o *IdentityAccessManagementLoginModeDefault) Code() int {
	return o._statusCode
}

func (o *IdentityAccessManagementLoginModeDefault) Error() string {
	return fmt.Sprintf("[POST /v1/login/mode][%d] IdentityAccessManagement_LoginMode default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementLoginModeDefault) String() string {
	return fmt.Sprintf("[POST /v1/login/mode][%d] IdentityAccessManagement_LoginMode default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementLoginModeDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementLoginModeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
