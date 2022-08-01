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

// IdentityAccessManagementDeleteRoleReader is a Reader for the IdentityAccessManagementDeleteRole structure.
type IdentityAccessManagementDeleteRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementDeleteRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementDeleteRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIdentityAccessManagementDeleteRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementDeleteRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIdentityAccessManagementDeleteRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewIdentityAccessManagementDeleteRoleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementDeleteRoleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementDeleteRoleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementDeleteRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementDeleteRoleOK creates a IdentityAccessManagementDeleteRoleOK with default headers values
func NewIdentityAccessManagementDeleteRoleOK() *IdentityAccessManagementDeleteRoleOK {
	return &IdentityAccessManagementDeleteRoleOK{}
}

/* IdentityAccessManagementDeleteRoleOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementDeleteRoleOK struct {
	Payload *swagger_models.CrudResponse
}

func (o *IdentityAccessManagementDeleteRoleOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/roles/id/{id}][%d] identityAccessManagementDeleteRoleOK  %+v", 200, o.Payload)
}
func (o *IdentityAccessManagementDeleteRoleOK) GetPayload() *swagger_models.CrudResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.CrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteRoleUnauthorized creates a IdentityAccessManagementDeleteRoleUnauthorized with default headers values
func NewIdentityAccessManagementDeleteRoleUnauthorized() *IdentityAccessManagementDeleteRoleUnauthorized {
	return &IdentityAccessManagementDeleteRoleUnauthorized{}
}

/* IdentityAccessManagementDeleteRoleUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementDeleteRoleUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementDeleteRoleUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/roles/id/{id}][%d] identityAccessManagementDeleteRoleUnauthorized  %+v", 401, o.Payload)
}
func (o *IdentityAccessManagementDeleteRoleUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteRoleForbidden creates a IdentityAccessManagementDeleteRoleForbidden with default headers values
func NewIdentityAccessManagementDeleteRoleForbidden() *IdentityAccessManagementDeleteRoleForbidden {
	return &IdentityAccessManagementDeleteRoleForbidden{}
}

/* IdentityAccessManagementDeleteRoleForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementDeleteRoleForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementDeleteRoleForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/roles/id/{id}][%d] identityAccessManagementDeleteRoleForbidden  %+v", 403, o.Payload)
}
func (o *IdentityAccessManagementDeleteRoleForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteRoleNotFound creates a IdentityAccessManagementDeleteRoleNotFound with default headers values
func NewIdentityAccessManagementDeleteRoleNotFound() *IdentityAccessManagementDeleteRoleNotFound {
	return &IdentityAccessManagementDeleteRoleNotFound{}
}

/* IdentityAccessManagementDeleteRoleNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type IdentityAccessManagementDeleteRoleNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementDeleteRoleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/roles/id/{id}][%d] identityAccessManagementDeleteRoleNotFound  %+v", 404, o.Payload)
}
func (o *IdentityAccessManagementDeleteRoleNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteRoleConflict creates a IdentityAccessManagementDeleteRoleConflict with default headers values
func NewIdentityAccessManagementDeleteRoleConflict() *IdentityAccessManagementDeleteRoleConflict {
	return &IdentityAccessManagementDeleteRoleConflict{}
}

/* IdentityAccessManagementDeleteRoleConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because there are IAM users of this IAM role
*/
type IdentityAccessManagementDeleteRoleConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementDeleteRoleConflict) Error() string {
	return fmt.Sprintf("[DELETE /v1/roles/id/{id}][%d] identityAccessManagementDeleteRoleConflict  %+v", 409, o.Payload)
}
func (o *IdentityAccessManagementDeleteRoleConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteRoleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteRoleInternalServerError creates a IdentityAccessManagementDeleteRoleInternalServerError with default headers values
func NewIdentityAccessManagementDeleteRoleInternalServerError() *IdentityAccessManagementDeleteRoleInternalServerError {
	return &IdentityAccessManagementDeleteRoleInternalServerError{}
}

/* IdentityAccessManagementDeleteRoleInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementDeleteRoleInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementDeleteRoleInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/roles/id/{id}][%d] identityAccessManagementDeleteRoleInternalServerError  %+v", 500, o.Payload)
}
func (o *IdentityAccessManagementDeleteRoleInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteRoleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteRoleGatewayTimeout creates a IdentityAccessManagementDeleteRoleGatewayTimeout with default headers values
func NewIdentityAccessManagementDeleteRoleGatewayTimeout() *IdentityAccessManagementDeleteRoleGatewayTimeout {
	return &IdentityAccessManagementDeleteRoleGatewayTimeout{}
}

/* IdentityAccessManagementDeleteRoleGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementDeleteRoleGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementDeleteRoleGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /v1/roles/id/{id}][%d] identityAccessManagementDeleteRoleGatewayTimeout  %+v", 504, o.Payload)
}
func (o *IdentityAccessManagementDeleteRoleGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteRoleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteRoleDefault creates a IdentityAccessManagementDeleteRoleDefault with default headers values
func NewIdentityAccessManagementDeleteRoleDefault(code int) *IdentityAccessManagementDeleteRoleDefault {
	return &IdentityAccessManagementDeleteRoleDefault{
		_statusCode: code,
	}
}

/* IdentityAccessManagementDeleteRoleDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementDeleteRoleDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the identity access management delete role default response
func (o *IdentityAccessManagementDeleteRoleDefault) Code() int {
	return o._statusCode
}

func (o *IdentityAccessManagementDeleteRoleDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/roles/id/{id}][%d] IdentityAccessManagement_DeleteRole default  %+v", o._statusCode, o.Payload)
}
func (o *IdentityAccessManagementDeleteRoleDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}