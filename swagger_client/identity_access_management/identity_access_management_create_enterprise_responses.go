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

// IdentityAccessManagementCreateEnterpriseReader is a Reader for the IdentityAccessManagementCreateEnterprise structure.
type IdentityAccessManagementCreateEnterpriseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementCreateEnterpriseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementCreateEnterpriseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIdentityAccessManagementCreateEnterpriseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewIdentityAccessManagementCreateEnterpriseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementCreateEnterpriseForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewIdentityAccessManagementCreateEnterpriseConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementCreateEnterpriseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementCreateEnterpriseGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementCreateEnterpriseDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementCreateEnterpriseOK creates a IdentityAccessManagementCreateEnterpriseOK with default headers values
func NewIdentityAccessManagementCreateEnterpriseOK() *IdentityAccessManagementCreateEnterpriseOK {
	return &IdentityAccessManagementCreateEnterpriseOK{}
}

/* IdentityAccessManagementCreateEnterpriseOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementCreateEnterpriseOK struct {
	Payload *swagger_models.CrudResponse
}

func (o *IdentityAccessManagementCreateEnterpriseOK) Error() string {
	return fmt.Sprintf("[POST /v1/enterprises][%d] identityAccessManagementCreateEnterpriseOK  %+v", 200, o.Payload)
}
func (o *IdentityAccessManagementCreateEnterpriseOK) GetPayload() *swagger_models.CrudResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateEnterpriseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.CrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateEnterpriseBadRequest creates a IdentityAccessManagementCreateEnterpriseBadRequest with default headers values
func NewIdentityAccessManagementCreateEnterpriseBadRequest() *IdentityAccessManagementCreateEnterpriseBadRequest {
	return &IdentityAccessManagementCreateEnterpriseBadRequest{}
}

/* IdentityAccessManagementCreateEnterpriseBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type IdentityAccessManagementCreateEnterpriseBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementCreateEnterpriseBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/enterprises][%d] identityAccessManagementCreateEnterpriseBadRequest  %+v", 400, o.Payload)
}
func (o *IdentityAccessManagementCreateEnterpriseBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateEnterpriseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateEnterpriseUnauthorized creates a IdentityAccessManagementCreateEnterpriseUnauthorized with default headers values
func NewIdentityAccessManagementCreateEnterpriseUnauthorized() *IdentityAccessManagementCreateEnterpriseUnauthorized {
	return &IdentityAccessManagementCreateEnterpriseUnauthorized{}
}

/* IdentityAccessManagementCreateEnterpriseUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementCreateEnterpriseUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementCreateEnterpriseUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/enterprises][%d] identityAccessManagementCreateEnterpriseUnauthorized  %+v", 401, o.Payload)
}
func (o *IdentityAccessManagementCreateEnterpriseUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateEnterpriseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateEnterpriseForbidden creates a IdentityAccessManagementCreateEnterpriseForbidden with default headers values
func NewIdentityAccessManagementCreateEnterpriseForbidden() *IdentityAccessManagementCreateEnterpriseForbidden {
	return &IdentityAccessManagementCreateEnterpriseForbidden{}
}

/* IdentityAccessManagementCreateEnterpriseForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementCreateEnterpriseForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementCreateEnterpriseForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/enterprises][%d] identityAccessManagementCreateEnterpriseForbidden  %+v", 403, o.Payload)
}
func (o *IdentityAccessManagementCreateEnterpriseForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateEnterpriseForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateEnterpriseConflict creates a IdentityAccessManagementCreateEnterpriseConflict with default headers values
func NewIdentityAccessManagementCreateEnterpriseConflict() *IdentityAccessManagementCreateEnterpriseConflict {
	return &IdentityAccessManagementCreateEnterpriseConflict{}
}

/* IdentityAccessManagementCreateEnterpriseConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this IAM role record will conflict with an already existing IAM role record.
*/
type IdentityAccessManagementCreateEnterpriseConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementCreateEnterpriseConflict) Error() string {
	return fmt.Sprintf("[POST /v1/enterprises][%d] identityAccessManagementCreateEnterpriseConflict  %+v", 409, o.Payload)
}
func (o *IdentityAccessManagementCreateEnterpriseConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateEnterpriseConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateEnterpriseInternalServerError creates a IdentityAccessManagementCreateEnterpriseInternalServerError with default headers values
func NewIdentityAccessManagementCreateEnterpriseInternalServerError() *IdentityAccessManagementCreateEnterpriseInternalServerError {
	return &IdentityAccessManagementCreateEnterpriseInternalServerError{}
}

/* IdentityAccessManagementCreateEnterpriseInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementCreateEnterpriseInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementCreateEnterpriseInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/enterprises][%d] identityAccessManagementCreateEnterpriseInternalServerError  %+v", 500, o.Payload)
}
func (o *IdentityAccessManagementCreateEnterpriseInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateEnterpriseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateEnterpriseGatewayTimeout creates a IdentityAccessManagementCreateEnterpriseGatewayTimeout with default headers values
func NewIdentityAccessManagementCreateEnterpriseGatewayTimeout() *IdentityAccessManagementCreateEnterpriseGatewayTimeout {
	return &IdentityAccessManagementCreateEnterpriseGatewayTimeout{}
}

/* IdentityAccessManagementCreateEnterpriseGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementCreateEnterpriseGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *IdentityAccessManagementCreateEnterpriseGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/enterprises][%d] identityAccessManagementCreateEnterpriseGatewayTimeout  %+v", 504, o.Payload)
}
func (o *IdentityAccessManagementCreateEnterpriseGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementCreateEnterpriseGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementCreateEnterpriseDefault creates a IdentityAccessManagementCreateEnterpriseDefault with default headers values
func NewIdentityAccessManagementCreateEnterpriseDefault(code int) *IdentityAccessManagementCreateEnterpriseDefault {
	return &IdentityAccessManagementCreateEnterpriseDefault{
		_statusCode: code,
	}
}

/* IdentityAccessManagementCreateEnterpriseDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementCreateEnterpriseDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the identity access management create enterprise default response
func (o *IdentityAccessManagementCreateEnterpriseDefault) Code() int {
	return o._statusCode
}

func (o *IdentityAccessManagementCreateEnterpriseDefault) Error() string {
	return fmt.Sprintf("[POST /v1/enterprises][%d] IdentityAccessManagement_CreateEnterprise default  %+v", o._statusCode, o.Payload)
}
func (o *IdentityAccessManagementCreateEnterpriseDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementCreateEnterpriseDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}