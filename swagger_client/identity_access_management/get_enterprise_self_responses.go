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

// GetEnterpriseSelfReader is a Reader for the GetEnterpriseSelf structure.
type GetEnterpriseSelfReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEnterpriseSelfReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEnterpriseSelfOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEnterpriseSelfUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEnterpriseSelfForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEnterpriseSelfNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEnterpriseSelfInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetEnterpriseSelfGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEnterpriseSelfOK creates a GetEnterpriseSelfOK with default headers values
func NewGetEnterpriseSelfOK() *GetEnterpriseSelfOK {
	return &GetEnterpriseSelfOK{}
}

/* GetEnterpriseSelfOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetEnterpriseSelfOK struct {
	Payload *swagger_models.CrudResponseRead
}

func (o *GetEnterpriseSelfOK) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises/self][%d] getEnterpriseSelfOK  %+v", 200, o.Payload)
}
func (o *GetEnterpriseSelfOK) GetPayload() *swagger_models.CrudResponseRead {
	return o.Payload
}

func (o *GetEnterpriseSelfOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.CrudResponseRead)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEnterpriseSelfUnauthorized creates a GetEnterpriseSelfUnauthorized with default headers values
func NewGetEnterpriseSelfUnauthorized() *GetEnterpriseSelfUnauthorized {
	return &GetEnterpriseSelfUnauthorized{}
}

/* GetEnterpriseSelfUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetEnterpriseSelfUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEnterpriseSelfUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises/self][%d] getEnterpriseSelfUnauthorized  %+v", 401, o.Payload)
}
func (o *GetEnterpriseSelfUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEnterpriseSelfUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEnterpriseSelfForbidden creates a GetEnterpriseSelfForbidden with default headers values
func NewGetEnterpriseSelfForbidden() *GetEnterpriseSelfForbidden {
	return &GetEnterpriseSelfForbidden{}
}

/* GetEnterpriseSelfForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type GetEnterpriseSelfForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEnterpriseSelfForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises/self][%d] getEnterpriseSelfForbidden  %+v", 403, o.Payload)
}
func (o *GetEnterpriseSelfForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEnterpriseSelfForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEnterpriseSelfNotFound creates a GetEnterpriseSelfNotFound with default headers values
func NewGetEnterpriseSelfNotFound() *GetEnterpriseSelfNotFound {
	return &GetEnterpriseSelfNotFound{}
}

/* GetEnterpriseSelfNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetEnterpriseSelfNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEnterpriseSelfNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises/self][%d] getEnterpriseSelfNotFound  %+v", 404, o.Payload)
}
func (o *GetEnterpriseSelfNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEnterpriseSelfNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEnterpriseSelfInternalServerError creates a GetEnterpriseSelfInternalServerError with default headers values
func NewGetEnterpriseSelfInternalServerError() *GetEnterpriseSelfInternalServerError {
	return &GetEnterpriseSelfInternalServerError{}
}

/* GetEnterpriseSelfInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetEnterpriseSelfInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEnterpriseSelfInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises/self][%d] getEnterpriseSelfInternalServerError  %+v", 500, o.Payload)
}
func (o *GetEnterpriseSelfInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEnterpriseSelfInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEnterpriseSelfGatewayTimeout creates a GetEnterpriseSelfGatewayTimeout with default headers values
func NewGetEnterpriseSelfGatewayTimeout() *GetEnterpriseSelfGatewayTimeout {
	return &GetEnterpriseSelfGatewayTimeout{}
}

/* GetEnterpriseSelfGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetEnterpriseSelfGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEnterpriseSelfGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/enterprises/self][%d] getEnterpriseSelfGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetEnterpriseSelfGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEnterpriseSelfGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
