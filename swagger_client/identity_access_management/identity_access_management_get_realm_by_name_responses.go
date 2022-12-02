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

// IdentityAccessManagementGetRealmByNameReader is a Reader for the IdentityAccessManagementGetRealmByName structure.
type IdentityAccessManagementGetRealmByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementGetRealmByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementGetRealmByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIdentityAccessManagementGetRealmByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementGetRealmByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIdentityAccessManagementGetRealmByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementGetRealmByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementGetRealmByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementGetRealmByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementGetRealmByNameOK creates a IdentityAccessManagementGetRealmByNameOK with default headers values
func NewIdentityAccessManagementGetRealmByNameOK() *IdentityAccessManagementGetRealmByNameOK {
	return &IdentityAccessManagementGetRealmByNameOK{}
}

/*
IdentityAccessManagementGetRealmByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementGetRealmByNameOK struct {
	Payload *swagger_models.CrudResponseRead
}

// IsSuccess returns true when this identity access management get realm by name o k response has a 2xx status code
func (o *IdentityAccessManagementGetRealmByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this identity access management get realm by name o k response has a 3xx status code
func (o *IdentityAccessManagementGetRealmByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get realm by name o k response has a 4xx status code
func (o *IdentityAccessManagementGetRealmByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management get realm by name o k response has a 5xx status code
func (o *IdentityAccessManagementGetRealmByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get realm by name o k response a status code equal to that given
func (o *IdentityAccessManagementGetRealmByNameOK) IsCode(code int) bool {
	return code == 200
}

func (o *IdentityAccessManagementGetRealmByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/realms/name/{name}][%d] identityAccessManagementGetRealmByNameOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementGetRealmByNameOK) String() string {
	return fmt.Sprintf("[GET /v1/realms/name/{name}][%d] identityAccessManagementGetRealmByNameOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementGetRealmByNameOK) GetPayload() *swagger_models.CrudResponseRead {
	return o.Payload
}

func (o *IdentityAccessManagementGetRealmByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.CrudResponseRead)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetRealmByNameUnauthorized creates a IdentityAccessManagementGetRealmByNameUnauthorized with default headers values
func NewIdentityAccessManagementGetRealmByNameUnauthorized() *IdentityAccessManagementGetRealmByNameUnauthorized {
	return &IdentityAccessManagementGetRealmByNameUnauthorized{}
}

/*
IdentityAccessManagementGetRealmByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementGetRealmByNameUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get realm by name unauthorized response has a 2xx status code
func (o *IdentityAccessManagementGetRealmByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get realm by name unauthorized response has a 3xx status code
func (o *IdentityAccessManagementGetRealmByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get realm by name unauthorized response has a 4xx status code
func (o *IdentityAccessManagementGetRealmByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management get realm by name unauthorized response has a 5xx status code
func (o *IdentityAccessManagementGetRealmByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get realm by name unauthorized response a status code equal to that given
func (o *IdentityAccessManagementGetRealmByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *IdentityAccessManagementGetRealmByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/realms/name/{name}][%d] identityAccessManagementGetRealmByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementGetRealmByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/realms/name/{name}][%d] identityAccessManagementGetRealmByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementGetRealmByNameUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetRealmByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetRealmByNameForbidden creates a IdentityAccessManagementGetRealmByNameForbidden with default headers values
func NewIdentityAccessManagementGetRealmByNameForbidden() *IdentityAccessManagementGetRealmByNameForbidden {
	return &IdentityAccessManagementGetRealmByNameForbidden{}
}

/*
IdentityAccessManagementGetRealmByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementGetRealmByNameForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get realm by name forbidden response has a 2xx status code
func (o *IdentityAccessManagementGetRealmByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get realm by name forbidden response has a 3xx status code
func (o *IdentityAccessManagementGetRealmByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get realm by name forbidden response has a 4xx status code
func (o *IdentityAccessManagementGetRealmByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management get realm by name forbidden response has a 5xx status code
func (o *IdentityAccessManagementGetRealmByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get realm by name forbidden response a status code equal to that given
func (o *IdentityAccessManagementGetRealmByNameForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *IdentityAccessManagementGetRealmByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/realms/name/{name}][%d] identityAccessManagementGetRealmByNameForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementGetRealmByNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/realms/name/{name}][%d] identityAccessManagementGetRealmByNameForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementGetRealmByNameForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetRealmByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetRealmByNameNotFound creates a IdentityAccessManagementGetRealmByNameNotFound with default headers values
func NewIdentityAccessManagementGetRealmByNameNotFound() *IdentityAccessManagementGetRealmByNameNotFound {
	return &IdentityAccessManagementGetRealmByNameNotFound{}
}

/*
IdentityAccessManagementGetRealmByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type IdentityAccessManagementGetRealmByNameNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get realm by name not found response has a 2xx status code
func (o *IdentityAccessManagementGetRealmByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get realm by name not found response has a 3xx status code
func (o *IdentityAccessManagementGetRealmByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get realm by name not found response has a 4xx status code
func (o *IdentityAccessManagementGetRealmByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management get realm by name not found response has a 5xx status code
func (o *IdentityAccessManagementGetRealmByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management get realm by name not found response a status code equal to that given
func (o *IdentityAccessManagementGetRealmByNameNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *IdentityAccessManagementGetRealmByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/realms/name/{name}][%d] identityAccessManagementGetRealmByNameNotFound  %+v", 404, o.Payload)
}

func (o *IdentityAccessManagementGetRealmByNameNotFound) String() string {
	return fmt.Sprintf("[GET /v1/realms/name/{name}][%d] identityAccessManagementGetRealmByNameNotFound  %+v", 404, o.Payload)
}

func (o *IdentityAccessManagementGetRealmByNameNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetRealmByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetRealmByNameInternalServerError creates a IdentityAccessManagementGetRealmByNameInternalServerError with default headers values
func NewIdentityAccessManagementGetRealmByNameInternalServerError() *IdentityAccessManagementGetRealmByNameInternalServerError {
	return &IdentityAccessManagementGetRealmByNameInternalServerError{}
}

/*
IdentityAccessManagementGetRealmByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementGetRealmByNameInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get realm by name internal server error response has a 2xx status code
func (o *IdentityAccessManagementGetRealmByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get realm by name internal server error response has a 3xx status code
func (o *IdentityAccessManagementGetRealmByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get realm by name internal server error response has a 4xx status code
func (o *IdentityAccessManagementGetRealmByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management get realm by name internal server error response has a 5xx status code
func (o *IdentityAccessManagementGetRealmByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management get realm by name internal server error response a status code equal to that given
func (o *IdentityAccessManagementGetRealmByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *IdentityAccessManagementGetRealmByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/realms/name/{name}][%d] identityAccessManagementGetRealmByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementGetRealmByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/realms/name/{name}][%d] identityAccessManagementGetRealmByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementGetRealmByNameInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetRealmByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetRealmByNameGatewayTimeout creates a IdentityAccessManagementGetRealmByNameGatewayTimeout with default headers values
func NewIdentityAccessManagementGetRealmByNameGatewayTimeout() *IdentityAccessManagementGetRealmByNameGatewayTimeout {
	return &IdentityAccessManagementGetRealmByNameGatewayTimeout{}
}

/*
IdentityAccessManagementGetRealmByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementGetRealmByNameGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management get realm by name gateway timeout response has a 2xx status code
func (o *IdentityAccessManagementGetRealmByNameGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management get realm by name gateway timeout response has a 3xx status code
func (o *IdentityAccessManagementGetRealmByNameGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management get realm by name gateway timeout response has a 4xx status code
func (o *IdentityAccessManagementGetRealmByNameGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management get realm by name gateway timeout response has a 5xx status code
func (o *IdentityAccessManagementGetRealmByNameGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management get realm by name gateway timeout response a status code equal to that given
func (o *IdentityAccessManagementGetRealmByNameGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *IdentityAccessManagementGetRealmByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/realms/name/{name}][%d] identityAccessManagementGetRealmByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementGetRealmByNameGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/realms/name/{name}][%d] identityAccessManagementGetRealmByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementGetRealmByNameGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementGetRealmByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementGetRealmByNameDefault creates a IdentityAccessManagementGetRealmByNameDefault with default headers values
func NewIdentityAccessManagementGetRealmByNameDefault(code int) *IdentityAccessManagementGetRealmByNameDefault {
	return &IdentityAccessManagementGetRealmByNameDefault{
		_statusCode: code,
	}
}

/*
IdentityAccessManagementGetRealmByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementGetRealmByNameDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the identity access management get realm by name default response
func (o *IdentityAccessManagementGetRealmByNameDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this identity access management get realm by name default response has a 2xx status code
func (o *IdentityAccessManagementGetRealmByNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this identity access management get realm by name default response has a 3xx status code
func (o *IdentityAccessManagementGetRealmByNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this identity access management get realm by name default response has a 4xx status code
func (o *IdentityAccessManagementGetRealmByNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this identity access management get realm by name default response has a 5xx status code
func (o *IdentityAccessManagementGetRealmByNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this identity access management get realm by name default response a status code equal to that given
func (o *IdentityAccessManagementGetRealmByNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IdentityAccessManagementGetRealmByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/realms/name/{name}][%d] IdentityAccessManagement_GetRealmByName default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementGetRealmByNameDefault) String() string {
	return fmt.Sprintf("[GET /v1/realms/name/{name}][%d] IdentityAccessManagement_GetRealmByName default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementGetRealmByNameDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementGetRealmByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
