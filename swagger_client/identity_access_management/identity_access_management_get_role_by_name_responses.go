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

// IdentityAccessManagementGetRoleByNameReader is a Reader for the IdentityAccessManagementGetRoleByName structure.
type IdentityAccessManagementGetRoleByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementGetRoleByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementGetRoleByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIdentityAccessManagementGetRoleByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementGetRoleByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIdentityAccessManagementGetRoleByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementGetRoleByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementGetRoleByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementGetRoleByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementGetRoleByNameOK creates a IdentityAccessManagementGetRoleByNameOK with default headers values
func NewIdentityAccessManagementGetRoleByNameOK() *IdentityAccessManagementGetRoleByNameOK {
	return &IdentityAccessManagementGetRoleByNameOK{}
}

/*
IdentityAccessManagementGetRoleByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementGetRoleByNameOK struct {
	Payload *swagger_models.CrudResponseRead
}

// IsSuccess returns true when this identity access management get role by name o k response has a 2xx status code
func (o *IdentityAccessManagementGetRoleByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this identity access management get role by name o k response has a 3xx status code
func (o *IdentityAccessManagementGetRoleByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get role by name o k response has a 4xx status code
func (o *IdentityAccessManagementGetRoleByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management get role by name o k response has a 5xx status code
func (o *IdentityAccessManagementGetRoleByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get role by name o k response a status code equal to that given
func (o *IdentityAccessManagementGetRoleByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the identity access management get role by name o k response
func (o *IdentityAccessManagementGetRoleByNameOK) Code() int {
	return 200
}

func (o *IdentityAccessManagementGetRoleByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/roles/name/{name}][%d] identityAccessManagementGetRoleByNameOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementGetRoleByNameOK) String() string {
	return fmt.Sprintf("[GET /v1/roles/name/{name}][%d] identityAccessManagementGetRoleByNameOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementGetRoleByNameOK) GetPayload() *swagger_models.CrudResponseRead {
	return o.Payload
}

func (o *IdentityAccessManagementGetRoleByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.CrudResponseRead)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetRoleByNameUnauthorized creates a IdentityAccessManagementGetRoleByNameUnauthorized with default headers values
func NewIdentityAccessManagementGetRoleByNameUnauthorized() *IdentityAccessManagementGetRoleByNameUnauthorized {
	return &IdentityAccessManagementGetRoleByNameUnauthorized{}
}

/*
IdentityAccessManagementGetRoleByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementGetRoleByNameUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get role by name unauthorized response has a 2xx status code
func (o *IdentityAccessManagementGetRoleByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get role by name unauthorized response has a 3xx status code
func (o *IdentityAccessManagementGetRoleByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get role by name unauthorized response has a 4xx status code
func (o *IdentityAccessManagementGetRoleByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management get role by name unauthorized response has a 5xx status code
func (o *IdentityAccessManagementGetRoleByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get role by name unauthorized response a status code equal to that given
func (o *IdentityAccessManagementGetRoleByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the identity access management get role by name unauthorized response
func (o *IdentityAccessManagementGetRoleByNameUnauthorized) Code() int {
	return 401
}

func (o *IdentityAccessManagementGetRoleByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/roles/name/{name}][%d] identityAccessManagementGetRoleByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementGetRoleByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/roles/name/{name}][%d] identityAccessManagementGetRoleByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementGetRoleByNameUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetRoleByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetRoleByNameForbidden creates a IdentityAccessManagementGetRoleByNameForbidden with default headers values
func NewIdentityAccessManagementGetRoleByNameForbidden() *IdentityAccessManagementGetRoleByNameForbidden {
	return &IdentityAccessManagementGetRoleByNameForbidden{}
}

/*
IdentityAccessManagementGetRoleByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementGetRoleByNameForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get role by name forbidden response has a 2xx status code
func (o *IdentityAccessManagementGetRoleByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get role by name forbidden response has a 3xx status code
func (o *IdentityAccessManagementGetRoleByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get role by name forbidden response has a 4xx status code
func (o *IdentityAccessManagementGetRoleByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management get role by name forbidden response has a 5xx status code
func (o *IdentityAccessManagementGetRoleByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get role by name forbidden response a status code equal to that given
func (o *IdentityAccessManagementGetRoleByNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the identity access management get role by name forbidden response
func (o *IdentityAccessManagementGetRoleByNameForbidden) Code() int {
	return 403
}

func (o *IdentityAccessManagementGetRoleByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/roles/name/{name}][%d] identityAccessManagementGetRoleByNameForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementGetRoleByNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/roles/name/{name}][%d] identityAccessManagementGetRoleByNameForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementGetRoleByNameForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetRoleByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetRoleByNameNotFound creates a IdentityAccessManagementGetRoleByNameNotFound with default headers values
func NewIdentityAccessManagementGetRoleByNameNotFound() *IdentityAccessManagementGetRoleByNameNotFound {
	return &IdentityAccessManagementGetRoleByNameNotFound{}
}

/*
IdentityAccessManagementGetRoleByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type IdentityAccessManagementGetRoleByNameNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get role by name not found response has a 2xx status code
func (o *IdentityAccessManagementGetRoleByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get role by name not found response has a 3xx status code
func (o *IdentityAccessManagementGetRoleByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get role by name not found response has a 4xx status code
func (o *IdentityAccessManagementGetRoleByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management get role by name not found response has a 5xx status code
func (o *IdentityAccessManagementGetRoleByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get role by name not found response a status code equal to that given
func (o *IdentityAccessManagementGetRoleByNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the identity access management get role by name not found response
func (o *IdentityAccessManagementGetRoleByNameNotFound) Code() int {
	return 404
}

func (o *IdentityAccessManagementGetRoleByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/roles/name/{name}][%d] identityAccessManagementGetRoleByNameNotFound  %+v", 404, o.Payload)
}

func (o *IdentityAccessManagementGetRoleByNameNotFound) String() string {
	return fmt.Sprintf("[GET /v1/roles/name/{name}][%d] identityAccessManagementGetRoleByNameNotFound  %+v", 404, o.Payload)
}

func (o *IdentityAccessManagementGetRoleByNameNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetRoleByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetRoleByNameInternalServerError creates a IdentityAccessManagementGetRoleByNameInternalServerError with default headers values
func NewIdentityAccessManagementGetRoleByNameInternalServerError() *IdentityAccessManagementGetRoleByNameInternalServerError {
	return &IdentityAccessManagementGetRoleByNameInternalServerError{}
}

/*
IdentityAccessManagementGetRoleByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementGetRoleByNameInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get role by name internal server error response has a 2xx status code
func (o *IdentityAccessManagementGetRoleByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get role by name internal server error response has a 3xx status code
func (o *IdentityAccessManagementGetRoleByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get role by name internal server error response has a 4xx status code
func (o *IdentityAccessManagementGetRoleByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management get role by name internal server error response has a 5xx status code
func (o *IdentityAccessManagementGetRoleByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management get role by name internal server error response a status code equal to that given
func (o *IdentityAccessManagementGetRoleByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the identity access management get role by name internal server error response
func (o *IdentityAccessManagementGetRoleByNameInternalServerError) Code() int {
	return 500
}

func (o *IdentityAccessManagementGetRoleByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/roles/name/{name}][%d] identityAccessManagementGetRoleByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementGetRoleByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/roles/name/{name}][%d] identityAccessManagementGetRoleByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementGetRoleByNameInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetRoleByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetRoleByNameGatewayTimeout creates a IdentityAccessManagementGetRoleByNameGatewayTimeout with default headers values
func NewIdentityAccessManagementGetRoleByNameGatewayTimeout() *IdentityAccessManagementGetRoleByNameGatewayTimeout {
	return &IdentityAccessManagementGetRoleByNameGatewayTimeout{}
}

/*
IdentityAccessManagementGetRoleByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementGetRoleByNameGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get role by name gateway timeout response has a 2xx status code
func (o *IdentityAccessManagementGetRoleByNameGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get role by name gateway timeout response has a 3xx status code
func (o *IdentityAccessManagementGetRoleByNameGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get role by name gateway timeout response has a 4xx status code
func (o *IdentityAccessManagementGetRoleByNameGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management get role by name gateway timeout response has a 5xx status code
func (o *IdentityAccessManagementGetRoleByNameGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management get role by name gateway timeout response a status code equal to that given
func (o *IdentityAccessManagementGetRoleByNameGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the identity access management get role by name gateway timeout response
func (o *IdentityAccessManagementGetRoleByNameGatewayTimeout) Code() int {
	return 504
}

func (o *IdentityAccessManagementGetRoleByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/roles/name/{name}][%d] identityAccessManagementGetRoleByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementGetRoleByNameGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/roles/name/{name}][%d] identityAccessManagementGetRoleByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementGetRoleByNameGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetRoleByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetRoleByNameDefault creates a IdentityAccessManagementGetRoleByNameDefault with default headers values
func NewIdentityAccessManagementGetRoleByNameDefault(code int) *IdentityAccessManagementGetRoleByNameDefault {
	return &IdentityAccessManagementGetRoleByNameDefault{
		_statusCode: code,
	}
}

/*
IdentityAccessManagementGetRoleByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementGetRoleByNameDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// IsSuccess returns true when this identity access management get role by name default response has a 2xx status code
func (o *IdentityAccessManagementGetRoleByNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this identity access management get role by name default response has a 3xx status code
func (o *IdentityAccessManagementGetRoleByNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this identity access management get role by name default response has a 4xx status code
func (o *IdentityAccessManagementGetRoleByNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this identity access management get role by name default response has a 5xx status code
func (o *IdentityAccessManagementGetRoleByNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this identity access management get role by name default response a status code equal to that given
func (o *IdentityAccessManagementGetRoleByNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the identity access management get role by name default response
func (o *IdentityAccessManagementGetRoleByNameDefault) Code() int {
	return o._statusCode
}

func (o *IdentityAccessManagementGetRoleByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/roles/name/{name}][%d] IdentityAccessManagement_GetRoleByName default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementGetRoleByNameDefault) String() string {
	return fmt.Sprintf("[GET /v1/roles/name/{name}][%d] IdentityAccessManagement_GetRoleByName default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementGetRoleByNameDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementGetRoleByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
