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

// QueryRolesReader is a Reader for the QueryRoles structure.
type QueryRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewQueryRolesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewQueryRolesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryRolesOK creates a QueryRolesOK with default headers values
func NewQueryRolesOK() *QueryRolesOK {
	return &QueryRolesOK{}
}

/* QueryRolesOK describes a response with status code 200, with default header values.

A successful response.
*/
type QueryRolesOK struct {
	Payload *swagger_models.CrudResponse
}

func (o *QueryRolesOK) Error() string {
	return fmt.Sprintf("[GET /v1/roles][%d] queryRolesOK  %+v", 200, o.Payload)
}
func (o *QueryRolesOK) GetPayload() *swagger_models.CrudResponse {
	return o.Payload
}

func (o *QueryRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.CrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryRolesBadRequest creates a QueryRolesBadRequest with default headers values
func NewQueryRolesBadRequest() *QueryRolesBadRequest {
	return &QueryRolesBadRequest{}
}

/* QueryRolesBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type QueryRolesBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryRolesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/roles][%d] queryRolesBadRequest  %+v", 400, o.Payload)
}
func (o *QueryRolesBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryRolesUnauthorized creates a QueryRolesUnauthorized with default headers values
func NewQueryRolesUnauthorized() *QueryRolesUnauthorized {
	return &QueryRolesUnauthorized{}
}

/* QueryRolesUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type QueryRolesUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryRolesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/roles][%d] queryRolesUnauthorized  %+v", 401, o.Payload)
}
func (o *QueryRolesUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryRolesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryRolesForbidden creates a QueryRolesForbidden with default headers values
func NewQueryRolesForbidden() *QueryRolesForbidden {
	return &QueryRolesForbidden{}
}

/* QueryRolesForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type QueryRolesForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryRolesForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/roles][%d] queryRolesForbidden  %+v", 403, o.Payload)
}
func (o *QueryRolesForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryRolesInternalServerError creates a QueryRolesInternalServerError with default headers values
func NewQueryRolesInternalServerError() *QueryRolesInternalServerError {
	return &QueryRolesInternalServerError{}
}

/* QueryRolesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type QueryRolesInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryRolesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/roles][%d] queryRolesInternalServerError  %+v", 500, o.Payload)
}
func (o *QueryRolesInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryRolesGatewayTimeout creates a QueryRolesGatewayTimeout with default headers values
func NewQueryRolesGatewayTimeout() *QueryRolesGatewayTimeout {
	return &QueryRolesGatewayTimeout{}
}

/* QueryRolesGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type QueryRolesGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryRolesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/roles][%d] queryRolesGatewayTimeout  %+v", 504, o.Payload)
}
func (o *QueryRolesGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryRolesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
