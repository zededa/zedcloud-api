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

// IdentityAccessManagementUpdateEnterpriseReader is a Reader for the IdentityAccessManagementUpdateEnterprise structure.
type IdentityAccessManagementUpdateEnterpriseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementUpdateEnterpriseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementUpdateEnterpriseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIdentityAccessManagementUpdateEnterpriseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementUpdateEnterpriseForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIdentityAccessManagementUpdateEnterpriseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewIdentityAccessManagementUpdateEnterpriseConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementUpdateEnterpriseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementUpdateEnterpriseGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementUpdateEnterpriseDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementUpdateEnterpriseOK creates a IdentityAccessManagementUpdateEnterpriseOK with default headers values
func NewIdentityAccessManagementUpdateEnterpriseOK() *IdentityAccessManagementUpdateEnterpriseOK {
	return &IdentityAccessManagementUpdateEnterpriseOK{}
}

/* IdentityAccessManagementUpdateEnterpriseOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementUpdateEnterpriseOK struct {
	Payload *swagger_models.CrudResponse
}

func (o *IdentityAccessManagementUpdateEnterpriseOK) Error() string {
	return fmt.Sprintf("[PUT /v1/enterprises/self][%d] identityAccessManagementUpdateEnterpriseOK  %+v", 200, o.Payload)
}
func (o *IdentityAccessManagementUpdateEnterpriseOK) GetPayload() *swagger_models.CrudResponse {
	return o.Payload
}

func (o *IdentityAccessManagementUpdateEnterpriseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.CrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementUpdateEnterpriseUnauthorized creates a IdentityAccessManagementUpdateEnterpriseUnauthorized with default headers values
func NewIdentityAccessManagementUpdateEnterpriseUnauthorized() *IdentityAccessManagementUpdateEnterpriseUnauthorized {
	return &IdentityAccessManagementUpdateEnterpriseUnauthorized{}
}

/* IdentityAccessManagementUpdateEnterpriseUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementUpdateEnterpriseUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementUpdateEnterpriseUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/enterprises/self][%d] identityAccessManagementUpdateEnterpriseUnauthorized  %+v", 401, o.Payload)
}
func (o *IdentityAccessManagementUpdateEnterpriseUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementUpdateEnterpriseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementUpdateEnterpriseForbidden creates a IdentityAccessManagementUpdateEnterpriseForbidden with default headers values
func NewIdentityAccessManagementUpdateEnterpriseForbidden() *IdentityAccessManagementUpdateEnterpriseForbidden {
	return &IdentityAccessManagementUpdateEnterpriseForbidden{}
}

/* IdentityAccessManagementUpdateEnterpriseForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementUpdateEnterpriseForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementUpdateEnterpriseForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/enterprises/self][%d] identityAccessManagementUpdateEnterpriseForbidden  %+v", 403, o.Payload)
}
func (o *IdentityAccessManagementUpdateEnterpriseForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementUpdateEnterpriseForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementUpdateEnterpriseNotFound creates a IdentityAccessManagementUpdateEnterpriseNotFound with default headers values
func NewIdentityAccessManagementUpdateEnterpriseNotFound() *IdentityAccessManagementUpdateEnterpriseNotFound {
	return &IdentityAccessManagementUpdateEnterpriseNotFound{}
}

/* IdentityAccessManagementUpdateEnterpriseNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type IdentityAccessManagementUpdateEnterpriseNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementUpdateEnterpriseNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/enterprises/self][%d] identityAccessManagementUpdateEnterpriseNotFound  %+v", 404, o.Payload)
}
func (o *IdentityAccessManagementUpdateEnterpriseNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementUpdateEnterpriseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementUpdateEnterpriseConflict creates a IdentityAccessManagementUpdateEnterpriseConflict with default headers values
func NewIdentityAccessManagementUpdateEnterpriseConflict() *IdentityAccessManagementUpdateEnterpriseConflict {
	return &IdentityAccessManagementUpdateEnterpriseConflict{}
}

/* IdentityAccessManagementUpdateEnterpriseConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing enterprise record.
*/
type IdentityAccessManagementUpdateEnterpriseConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementUpdateEnterpriseConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/enterprises/self][%d] identityAccessManagementUpdateEnterpriseConflict  %+v", 409, o.Payload)
}
func (o *IdentityAccessManagementUpdateEnterpriseConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementUpdateEnterpriseConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementUpdateEnterpriseInternalServerError creates a IdentityAccessManagementUpdateEnterpriseInternalServerError with default headers values
func NewIdentityAccessManagementUpdateEnterpriseInternalServerError() *IdentityAccessManagementUpdateEnterpriseInternalServerError {
	return &IdentityAccessManagementUpdateEnterpriseInternalServerError{}
}

/* IdentityAccessManagementUpdateEnterpriseInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementUpdateEnterpriseInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementUpdateEnterpriseInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/enterprises/self][%d] identityAccessManagementUpdateEnterpriseInternalServerError  %+v", 500, o.Payload)
}
func (o *IdentityAccessManagementUpdateEnterpriseInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementUpdateEnterpriseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementUpdateEnterpriseGatewayTimeout creates a IdentityAccessManagementUpdateEnterpriseGatewayTimeout with default headers values
func NewIdentityAccessManagementUpdateEnterpriseGatewayTimeout() *IdentityAccessManagementUpdateEnterpriseGatewayTimeout {
	return &IdentityAccessManagementUpdateEnterpriseGatewayTimeout{}
}

/* IdentityAccessManagementUpdateEnterpriseGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementUpdateEnterpriseGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementUpdateEnterpriseGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/enterprises/self][%d] identityAccessManagementUpdateEnterpriseGatewayTimeout  %+v", 504, o.Payload)
}
func (o *IdentityAccessManagementUpdateEnterpriseGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementUpdateEnterpriseGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementUpdateEnterpriseDefault creates a IdentityAccessManagementUpdateEnterpriseDefault with default headers values
func NewIdentityAccessManagementUpdateEnterpriseDefault(code int) *IdentityAccessManagementUpdateEnterpriseDefault {
	return &IdentityAccessManagementUpdateEnterpriseDefault{
		_statusCode: code,
	}
}

/* IdentityAccessManagementUpdateEnterpriseDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementUpdateEnterpriseDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the identity access management update enterprise default response
func (o *IdentityAccessManagementUpdateEnterpriseDefault) Code() int {
	return o._statusCode
}

func (o *IdentityAccessManagementUpdateEnterpriseDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/enterprises/self][%d] IdentityAccessManagement_UpdateEnterprise default  %+v", o._statusCode, o.Payload)
}
func (o *IdentityAccessManagementUpdateEnterpriseDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementUpdateEnterpriseDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
