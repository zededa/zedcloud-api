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

// IdentityAccessManagementDeleteRealmReader is a Reader for the IdentityAccessManagementDeleteRealm structure.
type IdentityAccessManagementDeleteRealmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementDeleteRealmReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementDeleteRealmOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIdentityAccessManagementDeleteRealmUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementDeleteRealmForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIdentityAccessManagementDeleteRealmNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementDeleteRealmInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementDeleteRealmGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementDeleteRealmDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementDeleteRealmOK creates a IdentityAccessManagementDeleteRealmOK with default headers values
func NewIdentityAccessManagementDeleteRealmOK() *IdentityAccessManagementDeleteRealmOK {
	return &IdentityAccessManagementDeleteRealmOK{}
}

/*
IdentityAccessManagementDeleteRealmOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementDeleteRealmOK struct {
	Payload *swagger_models.CrudResponse
}

// IsSuccess returns true when this identity access management delete realm o k response has a 2xx status code
func (o *IdentityAccessManagementDeleteRealmOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this identity access management delete realm o k response has a 3xx status code
func (o *IdentityAccessManagementDeleteRealmOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management delete realm o k response has a 4xx status code
func (o *IdentityAccessManagementDeleteRealmOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management delete realm o k response has a 5xx status code
func (o *IdentityAccessManagementDeleteRealmOK) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management delete realm o k response a status code equal to that given
func (o *IdentityAccessManagementDeleteRealmOK) IsCode(code int) bool {
	return code == 200
}

func (o *IdentityAccessManagementDeleteRealmOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/realms/id/{id}][%d] identityAccessManagementDeleteRealmOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementDeleteRealmOK) String() string {
	return fmt.Sprintf("[DELETE /v1/realms/id/{id}][%d] identityAccessManagementDeleteRealmOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementDeleteRealmOK) GetPayload() *swagger_models.CrudResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteRealmOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.CrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteRealmUnauthorized creates a IdentityAccessManagementDeleteRealmUnauthorized with default headers values
func NewIdentityAccessManagementDeleteRealmUnauthorized() *IdentityAccessManagementDeleteRealmUnauthorized {
	return &IdentityAccessManagementDeleteRealmUnauthorized{}
}

/*
IdentityAccessManagementDeleteRealmUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementDeleteRealmUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management delete realm unauthorized response has a 2xx status code
func (o *IdentityAccessManagementDeleteRealmUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management delete realm unauthorized response has a 3xx status code
func (o *IdentityAccessManagementDeleteRealmUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management delete realm unauthorized response has a 4xx status code
func (o *IdentityAccessManagementDeleteRealmUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management delete realm unauthorized response has a 5xx status code
func (o *IdentityAccessManagementDeleteRealmUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management delete realm unauthorized response a status code equal to that given
func (o *IdentityAccessManagementDeleteRealmUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *IdentityAccessManagementDeleteRealmUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/realms/id/{id}][%d] identityAccessManagementDeleteRealmUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementDeleteRealmUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/realms/id/{id}][%d] identityAccessManagementDeleteRealmUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementDeleteRealmUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteRealmUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteRealmForbidden creates a IdentityAccessManagementDeleteRealmForbidden with default headers values
func NewIdentityAccessManagementDeleteRealmForbidden() *IdentityAccessManagementDeleteRealmForbidden {
	return &IdentityAccessManagementDeleteRealmForbidden{}
}

/*
IdentityAccessManagementDeleteRealmForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementDeleteRealmForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management delete realm forbidden response has a 2xx status code
func (o *IdentityAccessManagementDeleteRealmForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management delete realm forbidden response has a 3xx status code
func (o *IdentityAccessManagementDeleteRealmForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management delete realm forbidden response has a 4xx status code
func (o *IdentityAccessManagementDeleteRealmForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management delete realm forbidden response has a 5xx status code
func (o *IdentityAccessManagementDeleteRealmForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management delete realm forbidden response a status code equal to that given
func (o *IdentityAccessManagementDeleteRealmForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *IdentityAccessManagementDeleteRealmForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/realms/id/{id}][%d] identityAccessManagementDeleteRealmForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementDeleteRealmForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/realms/id/{id}][%d] identityAccessManagementDeleteRealmForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementDeleteRealmForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteRealmForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteRealmNotFound creates a IdentityAccessManagementDeleteRealmNotFound with default headers values
func NewIdentityAccessManagementDeleteRealmNotFound() *IdentityAccessManagementDeleteRealmNotFound {
	return &IdentityAccessManagementDeleteRealmNotFound{}
}

/*
IdentityAccessManagementDeleteRealmNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type IdentityAccessManagementDeleteRealmNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management delete realm not found response has a 2xx status code
func (o *IdentityAccessManagementDeleteRealmNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management delete realm not found response has a 3xx status code
func (o *IdentityAccessManagementDeleteRealmNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management delete realm not found response has a 4xx status code
func (o *IdentityAccessManagementDeleteRealmNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management delete realm not found response has a 5xx status code
func (o *IdentityAccessManagementDeleteRealmNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management delete realm not found response a status code equal to that given
func (o *IdentityAccessManagementDeleteRealmNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *IdentityAccessManagementDeleteRealmNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/realms/id/{id}][%d] identityAccessManagementDeleteRealmNotFound  %+v", 404, o.Payload)
}

func (o *IdentityAccessManagementDeleteRealmNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/realms/id/{id}][%d] identityAccessManagementDeleteRealmNotFound  %+v", 404, o.Payload)
}

func (o *IdentityAccessManagementDeleteRealmNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteRealmNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteRealmInternalServerError creates a IdentityAccessManagementDeleteRealmInternalServerError with default headers values
func NewIdentityAccessManagementDeleteRealmInternalServerError() *IdentityAccessManagementDeleteRealmInternalServerError {
	return &IdentityAccessManagementDeleteRealmInternalServerError{}
}

/*
IdentityAccessManagementDeleteRealmInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementDeleteRealmInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management delete realm internal server error response has a 2xx status code
func (o *IdentityAccessManagementDeleteRealmInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management delete realm internal server error response has a 3xx status code
func (o *IdentityAccessManagementDeleteRealmInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management delete realm internal server error response has a 4xx status code
func (o *IdentityAccessManagementDeleteRealmInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management delete realm internal server error response has a 5xx status code
func (o *IdentityAccessManagementDeleteRealmInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management delete realm internal server error response a status code equal to that given
func (o *IdentityAccessManagementDeleteRealmInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *IdentityAccessManagementDeleteRealmInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/realms/id/{id}][%d] identityAccessManagementDeleteRealmInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementDeleteRealmInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/realms/id/{id}][%d] identityAccessManagementDeleteRealmInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementDeleteRealmInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteRealmInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteRealmGatewayTimeout creates a IdentityAccessManagementDeleteRealmGatewayTimeout with default headers values
func NewIdentityAccessManagementDeleteRealmGatewayTimeout() *IdentityAccessManagementDeleteRealmGatewayTimeout {
	return &IdentityAccessManagementDeleteRealmGatewayTimeout{}
}

/*
IdentityAccessManagementDeleteRealmGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementDeleteRealmGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this identity access management delete realm gateway timeout response has a 2xx status code
func (o *IdentityAccessManagementDeleteRealmGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management delete realm gateway timeout response has a 3xx status code
func (o *IdentityAccessManagementDeleteRealmGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management delete realm gateway timeout response has a 4xx status code
func (o *IdentityAccessManagementDeleteRealmGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management delete realm gateway timeout response has a 5xx status code
func (o *IdentityAccessManagementDeleteRealmGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management delete realm gateway timeout response a status code equal to that given
func (o *IdentityAccessManagementDeleteRealmGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *IdentityAccessManagementDeleteRealmGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /v1/realms/id/{id}][%d] identityAccessManagementDeleteRealmGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementDeleteRealmGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /v1/realms/id/{id}][%d] identityAccessManagementDeleteRealmGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementDeleteRealmGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteRealmGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteRealmDefault creates a IdentityAccessManagementDeleteRealmDefault with default headers values
func NewIdentityAccessManagementDeleteRealmDefault(code int) *IdentityAccessManagementDeleteRealmDefault {
	return &IdentityAccessManagementDeleteRealmDefault{
		_statusCode: code,
	}
}

/*
IdentityAccessManagementDeleteRealmDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementDeleteRealmDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the identity access management delete realm default response
func (o *IdentityAccessManagementDeleteRealmDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this identity access management delete realm default response has a 2xx status code
func (o *IdentityAccessManagementDeleteRealmDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this identity access management delete realm default response has a 3xx status code
func (o *IdentityAccessManagementDeleteRealmDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this identity access management delete realm default response has a 4xx status code
func (o *IdentityAccessManagementDeleteRealmDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this identity access management delete realm default response has a 5xx status code
func (o *IdentityAccessManagementDeleteRealmDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this identity access management delete realm default response a status code equal to that given
func (o *IdentityAccessManagementDeleteRealmDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IdentityAccessManagementDeleteRealmDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/realms/id/{id}][%d] IdentityAccessManagement_DeleteRealm default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementDeleteRealmDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/realms/id/{id}][%d] IdentityAccessManagement_DeleteRealm default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementDeleteRealmDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteRealmDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
