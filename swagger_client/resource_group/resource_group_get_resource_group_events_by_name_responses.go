// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package resource_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// ResourceGroupGetResourceGroupEventsByNameReader is a Reader for the ResourceGroupGetResourceGroupEventsByName structure.
type ResourceGroupGetResourceGroupEventsByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResourceGroupGetResourceGroupEventsByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResourceGroupGetResourceGroupEventsByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewResourceGroupGetResourceGroupEventsByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewResourceGroupGetResourceGroupEventsByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewResourceGroupGetResourceGroupEventsByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewResourceGroupGetResourceGroupEventsByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewResourceGroupGetResourceGroupEventsByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewResourceGroupGetResourceGroupEventsByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResourceGroupGetResourceGroupEventsByNameOK creates a ResourceGroupGetResourceGroupEventsByNameOK with default headers values
func NewResourceGroupGetResourceGroupEventsByNameOK() *ResourceGroupGetResourceGroupEventsByNameOK {
	return &ResourceGroupGetResourceGroupEventsByNameOK{}
}

/* ResourceGroupGetResourceGroupEventsByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type ResourceGroupGetResourceGroupEventsByNameOK struct {
	Payload *swagger_models.EventQueryResponse
}

func (o *ResourceGroupGetResourceGroupEventsByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/projects/name/{objname}/events][%d] resourceGroupGetResourceGroupEventsByNameOK  %+v", 200, o.Payload)
}
func (o *ResourceGroupGetResourceGroupEventsByNameOK) GetPayload() *swagger_models.EventQueryResponse {
	return o.Payload
}

func (o *ResourceGroupGetResourceGroupEventsByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.EventQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupGetResourceGroupEventsByNameUnauthorized creates a ResourceGroupGetResourceGroupEventsByNameUnauthorized with default headers values
func NewResourceGroupGetResourceGroupEventsByNameUnauthorized() *ResourceGroupGetResourceGroupEventsByNameUnauthorized {
	return &ResourceGroupGetResourceGroupEventsByNameUnauthorized{}
}

/* ResourceGroupGetResourceGroupEventsByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ResourceGroupGetResourceGroupEventsByNameUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ResourceGroupGetResourceGroupEventsByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/projects/name/{objname}/events][%d] resourceGroupGetResourceGroupEventsByNameUnauthorized  %+v", 401, o.Payload)
}
func (o *ResourceGroupGetResourceGroupEventsByNameUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupGetResourceGroupEventsByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupGetResourceGroupEventsByNameForbidden creates a ResourceGroupGetResourceGroupEventsByNameForbidden with default headers values
func NewResourceGroupGetResourceGroupEventsByNameForbidden() *ResourceGroupGetResourceGroupEventsByNameForbidden {
	return &ResourceGroupGetResourceGroupEventsByNameForbidden{}
}

/* ResourceGroupGetResourceGroupEventsByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ResourceGroupGetResourceGroupEventsByNameForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ResourceGroupGetResourceGroupEventsByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/projects/name/{objname}/events][%d] resourceGroupGetResourceGroupEventsByNameForbidden  %+v", 403, o.Payload)
}
func (o *ResourceGroupGetResourceGroupEventsByNameForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupGetResourceGroupEventsByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupGetResourceGroupEventsByNameNotFound creates a ResourceGroupGetResourceGroupEventsByNameNotFound with default headers values
func NewResourceGroupGetResourceGroupEventsByNameNotFound() *ResourceGroupGetResourceGroupEventsByNameNotFound {
	return &ResourceGroupGetResourceGroupEventsByNameNotFound{}
}

/* ResourceGroupGetResourceGroupEventsByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type ResourceGroupGetResourceGroupEventsByNameNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ResourceGroupGetResourceGroupEventsByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/projects/name/{objname}/events][%d] resourceGroupGetResourceGroupEventsByNameNotFound  %+v", 404, o.Payload)
}
func (o *ResourceGroupGetResourceGroupEventsByNameNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupGetResourceGroupEventsByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupGetResourceGroupEventsByNameInternalServerError creates a ResourceGroupGetResourceGroupEventsByNameInternalServerError with default headers values
func NewResourceGroupGetResourceGroupEventsByNameInternalServerError() *ResourceGroupGetResourceGroupEventsByNameInternalServerError {
	return &ResourceGroupGetResourceGroupEventsByNameInternalServerError{}
}

/* ResourceGroupGetResourceGroupEventsByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ResourceGroupGetResourceGroupEventsByNameInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ResourceGroupGetResourceGroupEventsByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/projects/name/{objname}/events][%d] resourceGroupGetResourceGroupEventsByNameInternalServerError  %+v", 500, o.Payload)
}
func (o *ResourceGroupGetResourceGroupEventsByNameInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupGetResourceGroupEventsByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupGetResourceGroupEventsByNameGatewayTimeout creates a ResourceGroupGetResourceGroupEventsByNameGatewayTimeout with default headers values
func NewResourceGroupGetResourceGroupEventsByNameGatewayTimeout() *ResourceGroupGetResourceGroupEventsByNameGatewayTimeout {
	return &ResourceGroupGetResourceGroupEventsByNameGatewayTimeout{}
}

/* ResourceGroupGetResourceGroupEventsByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ResourceGroupGetResourceGroupEventsByNameGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ResourceGroupGetResourceGroupEventsByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/projects/name/{objname}/events][%d] resourceGroupGetResourceGroupEventsByNameGatewayTimeout  %+v", 504, o.Payload)
}
func (o *ResourceGroupGetResourceGroupEventsByNameGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupGetResourceGroupEventsByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupGetResourceGroupEventsByNameDefault creates a ResourceGroupGetResourceGroupEventsByNameDefault with default headers values
func NewResourceGroupGetResourceGroupEventsByNameDefault(code int) *ResourceGroupGetResourceGroupEventsByNameDefault {
	return &ResourceGroupGetResourceGroupEventsByNameDefault{
		_statusCode: code,
	}
}

/* ResourceGroupGetResourceGroupEventsByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ResourceGroupGetResourceGroupEventsByNameDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the resource group get resource group events by name default response
func (o *ResourceGroupGetResourceGroupEventsByNameDefault) Code() int {
	return o._statusCode
}

func (o *ResourceGroupGetResourceGroupEventsByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/projects/name/{objname}/events][%d] ResourceGroup_GetResourceGroupEventsByName default  %+v", o._statusCode, o.Payload)
}
func (o *ResourceGroupGetResourceGroupEventsByNameDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *ResourceGroupGetResourceGroupEventsByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
