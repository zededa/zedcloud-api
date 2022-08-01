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

// IdentityAccessManagementQueryRealmsReader is a Reader for the IdentityAccessManagementQueryRealms structure.
type IdentityAccessManagementQueryRealmsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementQueryRealmsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementQueryRealmsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIdentityAccessManagementQueryRealmsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewIdentityAccessManagementQueryRealmsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementQueryRealmsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementQueryRealmsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementQueryRealmsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementQueryRealmsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementQueryRealmsOK creates a IdentityAccessManagementQueryRealmsOK with default headers values
func NewIdentityAccessManagementQueryRealmsOK() *IdentityAccessManagementQueryRealmsOK {
	return &IdentityAccessManagementQueryRealmsOK{}
}

/* IdentityAccessManagementQueryRealmsOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementQueryRealmsOK struct {
	Payload *swagger_models.CrudResponse
}

func (o *IdentityAccessManagementQueryRealmsOK) Error() string {
	return fmt.Sprintf("[GET /v1/realms][%d] identityAccessManagementQueryRealmsOK  %+v", 200, o.Payload)
}
func (o *IdentityAccessManagementQueryRealmsOK) GetPayload() *swagger_models.CrudResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryRealmsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.CrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryRealmsBadRequest creates a IdentityAccessManagementQueryRealmsBadRequest with default headers values
func NewIdentityAccessManagementQueryRealmsBadRequest() *IdentityAccessManagementQueryRealmsBadRequest {
	return &IdentityAccessManagementQueryRealmsBadRequest{}
}

/* IdentityAccessManagementQueryRealmsBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type IdentityAccessManagementQueryRealmsBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementQueryRealmsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/realms][%d] identityAccessManagementQueryRealmsBadRequest  %+v", 400, o.Payload)
}
func (o *IdentityAccessManagementQueryRealmsBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryRealmsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryRealmsUnauthorized creates a IdentityAccessManagementQueryRealmsUnauthorized with default headers values
func NewIdentityAccessManagementQueryRealmsUnauthorized() *IdentityAccessManagementQueryRealmsUnauthorized {
	return &IdentityAccessManagementQueryRealmsUnauthorized{}
}

/* IdentityAccessManagementQueryRealmsUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementQueryRealmsUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementQueryRealmsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/realms][%d] identityAccessManagementQueryRealmsUnauthorized  %+v", 401, o.Payload)
}
func (o *IdentityAccessManagementQueryRealmsUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryRealmsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryRealmsForbidden creates a IdentityAccessManagementQueryRealmsForbidden with default headers values
func NewIdentityAccessManagementQueryRealmsForbidden() *IdentityAccessManagementQueryRealmsForbidden {
	return &IdentityAccessManagementQueryRealmsForbidden{}
}

/* IdentityAccessManagementQueryRealmsForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementQueryRealmsForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementQueryRealmsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/realms][%d] identityAccessManagementQueryRealmsForbidden  %+v", 403, o.Payload)
}
func (o *IdentityAccessManagementQueryRealmsForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryRealmsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryRealmsInternalServerError creates a IdentityAccessManagementQueryRealmsInternalServerError with default headers values
func NewIdentityAccessManagementQueryRealmsInternalServerError() *IdentityAccessManagementQueryRealmsInternalServerError {
	return &IdentityAccessManagementQueryRealmsInternalServerError{}
}

/* IdentityAccessManagementQueryRealmsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementQueryRealmsInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementQueryRealmsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/realms][%d] identityAccessManagementQueryRealmsInternalServerError  %+v", 500, o.Payload)
}
func (o *IdentityAccessManagementQueryRealmsInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryRealmsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryRealmsGatewayTimeout creates a IdentityAccessManagementQueryRealmsGatewayTimeout with default headers values
func NewIdentityAccessManagementQueryRealmsGatewayTimeout() *IdentityAccessManagementQueryRealmsGatewayTimeout {
	return &IdentityAccessManagementQueryRealmsGatewayTimeout{}
}

/* IdentityAccessManagementQueryRealmsGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementQueryRealmsGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementQueryRealmsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/realms][%d] identityAccessManagementQueryRealmsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *IdentityAccessManagementQueryRealmsGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementQueryRealmsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementQueryRealmsDefault creates a IdentityAccessManagementQueryRealmsDefault with default headers values
func NewIdentityAccessManagementQueryRealmsDefault(code int) *IdentityAccessManagementQueryRealmsDefault {
	return &IdentityAccessManagementQueryRealmsDefault{
		_statusCode: code,
	}
}

/* IdentityAccessManagementQueryRealmsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementQueryRealmsDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the identity access management query realms default response
func (o *IdentityAccessManagementQueryRealmsDefault) Code() int {
	return o._statusCode
}

func (o *IdentityAccessManagementQueryRealmsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/realms][%d] IdentityAccessManagement_QueryRealms default  %+v", o._statusCode, o.Payload)
}
func (o *IdentityAccessManagementQueryRealmsDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementQueryRealmsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
