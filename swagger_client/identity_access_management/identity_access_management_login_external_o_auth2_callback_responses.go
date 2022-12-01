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

// IdentityAccessManagementLoginExternalOAuth2CallbackReader is a Reader for the IdentityAccessManagementLoginExternalOAuth2Callback structure.
type IdentityAccessManagementLoginExternalOAuth2CallbackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementLoginExternalOAuth2CallbackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIdentityAccessManagementLoginExternalOAuth2CallbackUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementLoginExternalOAuth2CallbackForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementLoginExternalOAuth2CallbackInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementLoginExternalOAuth2CallbackGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementLoginExternalOAuth2CallbackDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementLoginExternalOAuth2CallbackOK creates a IdentityAccessManagementLoginExternalOAuth2CallbackOK with default headers values
func NewIdentityAccessManagementLoginExternalOAuth2CallbackOK() *IdentityAccessManagementLoginExternalOAuth2CallbackOK {
	return &IdentityAccessManagementLoginExternalOAuth2CallbackOK{}
}

/*
IdentityAccessManagementLoginExternalOAuth2CallbackOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementLoginExternalOAuth2CallbackOK struct {
	Payload *swagger_models.AAAFrontendLoginResponse
}

// IsSuccess returns true when this identity access management login external o auth2 callback o k response has a 2xx status code
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this identity access management login external o auth2 callback o k response has a 3xx status code
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management login external o auth2 callback o k response has a 4xx status code
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management login external o auth2 callback o k response has a 5xx status code
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackOK) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management login external o auth2 callback o k response a status code equal to that given
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackOK) IsCode(code int) bool {
	return code == 200
}

func (o *IdentityAccessManagementLoginExternalOAuth2CallbackOK) Error() string {
	return fmt.Sprintf("[POST /v1/login/oauth/callback][%d] identityAccessManagementLoginExternalOAuth2CallbackOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementLoginExternalOAuth2CallbackOK) String() string {
	return fmt.Sprintf("[POST /v1/login/oauth/callback][%d] identityAccessManagementLoginExternalOAuth2CallbackOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementLoginExternalOAuth2CallbackOK) GetPayload() *swagger_models.AAAFrontendLoginResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLoginExternalOAuth2CallbackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.AAAFrontendLoginResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLoginExternalOAuth2CallbackUnauthorized creates a IdentityAccessManagementLoginExternalOAuth2CallbackUnauthorized with default headers values
func NewIdentityAccessManagementLoginExternalOAuth2CallbackUnauthorized() *IdentityAccessManagementLoginExternalOAuth2CallbackUnauthorized {
	return &IdentityAccessManagementLoginExternalOAuth2CallbackUnauthorized{}
}

/*
IdentityAccessManagementLoginExternalOAuth2CallbackUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementLoginExternalOAuth2CallbackUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management login external o auth2 callback unauthorized response has a 2xx status code
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management login external o auth2 callback unauthorized response has a 3xx status code
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management login external o auth2 callback unauthorized response has a 4xx status code
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management login external o auth2 callback unauthorized response has a 5xx status code
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management login external o auth2 callback unauthorized response a status code equal to that given
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *IdentityAccessManagementLoginExternalOAuth2CallbackUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/login/oauth/callback][%d] identityAccessManagementLoginExternalOAuth2CallbackUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementLoginExternalOAuth2CallbackUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/login/oauth/callback][%d] identityAccessManagementLoginExternalOAuth2CallbackUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementLoginExternalOAuth2CallbackUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLoginExternalOAuth2CallbackUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLoginExternalOAuth2CallbackForbidden creates a IdentityAccessManagementLoginExternalOAuth2CallbackForbidden with default headers values
func NewIdentityAccessManagementLoginExternalOAuth2CallbackForbidden() *IdentityAccessManagementLoginExternalOAuth2CallbackForbidden {
	return &IdentityAccessManagementLoginExternalOAuth2CallbackForbidden{}
}

/*
IdentityAccessManagementLoginExternalOAuth2CallbackForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementLoginExternalOAuth2CallbackForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management login external o auth2 callback forbidden response has a 2xx status code
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management login external o auth2 callback forbidden response has a 3xx status code
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management login external o auth2 callback forbidden response has a 4xx status code
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management login external o auth2 callback forbidden response has a 5xx status code
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management login external o auth2 callback forbidden response a status code equal to that given
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *IdentityAccessManagementLoginExternalOAuth2CallbackForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/login/oauth/callback][%d] identityAccessManagementLoginExternalOAuth2CallbackForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementLoginExternalOAuth2CallbackForbidden) String() string {
	return fmt.Sprintf("[POST /v1/login/oauth/callback][%d] identityAccessManagementLoginExternalOAuth2CallbackForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementLoginExternalOAuth2CallbackForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLoginExternalOAuth2CallbackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLoginExternalOAuth2CallbackInternalServerError creates a IdentityAccessManagementLoginExternalOAuth2CallbackInternalServerError with default headers values
func NewIdentityAccessManagementLoginExternalOAuth2CallbackInternalServerError() *IdentityAccessManagementLoginExternalOAuth2CallbackInternalServerError {
	return &IdentityAccessManagementLoginExternalOAuth2CallbackInternalServerError{}
}

/*
IdentityAccessManagementLoginExternalOAuth2CallbackInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementLoginExternalOAuth2CallbackInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management login external o auth2 callback internal server error response has a 2xx status code
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management login external o auth2 callback internal server error response has a 3xx status code
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management login external o auth2 callback internal server error response has a 4xx status code
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management login external o auth2 callback internal server error response has a 5xx status code
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management login external o auth2 callback internal server error response a status code equal to that given
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *IdentityAccessManagementLoginExternalOAuth2CallbackInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/login/oauth/callback][%d] identityAccessManagementLoginExternalOAuth2CallbackInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementLoginExternalOAuth2CallbackInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/login/oauth/callback][%d] identityAccessManagementLoginExternalOAuth2CallbackInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementLoginExternalOAuth2CallbackInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLoginExternalOAuth2CallbackInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLoginExternalOAuth2CallbackGatewayTimeout creates a IdentityAccessManagementLoginExternalOAuth2CallbackGatewayTimeout with default headers values
func NewIdentityAccessManagementLoginExternalOAuth2CallbackGatewayTimeout() *IdentityAccessManagementLoginExternalOAuth2CallbackGatewayTimeout {
	return &IdentityAccessManagementLoginExternalOAuth2CallbackGatewayTimeout{}
}

/*
IdentityAccessManagementLoginExternalOAuth2CallbackGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementLoginExternalOAuth2CallbackGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management login external o auth2 callback gateway timeout response has a 2xx status code
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management login external o auth2 callback gateway timeout response has a 3xx status code
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management login external o auth2 callback gateway timeout response has a 4xx status code
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management login external o auth2 callback gateway timeout response has a 5xx status code
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management login external o auth2 callback gateway timeout response a status code equal to that given
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *IdentityAccessManagementLoginExternalOAuth2CallbackGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/login/oauth/callback][%d] identityAccessManagementLoginExternalOAuth2CallbackGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementLoginExternalOAuth2CallbackGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /v1/login/oauth/callback][%d] identityAccessManagementLoginExternalOAuth2CallbackGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementLoginExternalOAuth2CallbackGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementLoginExternalOAuth2CallbackGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementLoginExternalOAuth2CallbackDefault creates a IdentityAccessManagementLoginExternalOAuth2CallbackDefault with default headers values
func NewIdentityAccessManagementLoginExternalOAuth2CallbackDefault(code int) *IdentityAccessManagementLoginExternalOAuth2CallbackDefault {
	return &IdentityAccessManagementLoginExternalOAuth2CallbackDefault{
		_statusCode: code,
	}
}

/*
IdentityAccessManagementLoginExternalOAuth2CallbackDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementLoginExternalOAuth2CallbackDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the identity access management login external o auth2 callback default response
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this identity access management login external o auth2 callback default response has a 2xx status code
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this identity access management login external o auth2 callback default response has a 3xx status code
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this identity access management login external o auth2 callback default response has a 4xx status code
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this identity access management login external o auth2 callback default response has a 5xx status code
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this identity access management login external o auth2 callback default response a status code equal to that given
func (o *IdentityAccessManagementLoginExternalOAuth2CallbackDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IdentityAccessManagementLoginExternalOAuth2CallbackDefault) Error() string {
	return fmt.Sprintf("[POST /v1/login/oauth/callback][%d] IdentityAccessManagement_LoginExternalOAuth2Callback default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementLoginExternalOAuth2CallbackDefault) String() string {
	return fmt.Sprintf("[POST /v1/login/oauth/callback][%d] IdentityAccessManagement_LoginExternalOAuth2Callback default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementLoginExternalOAuth2CallbackDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementLoginExternalOAuth2CallbackDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
