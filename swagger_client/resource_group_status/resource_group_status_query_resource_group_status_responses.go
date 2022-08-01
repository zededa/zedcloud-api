// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package resource_group_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// ResourceGroupStatusQueryResourceGroupStatusReader is a Reader for the ResourceGroupStatusQueryResourceGroupStatus structure.
type ResourceGroupStatusQueryResourceGroupStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResourceGroupStatusQueryResourceGroupStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResourceGroupStatusQueryResourceGroupStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewResourceGroupStatusQueryResourceGroupStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewResourceGroupStatusQueryResourceGroupStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewResourceGroupStatusQueryResourceGroupStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewResourceGroupStatusQueryResourceGroupStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewResourceGroupStatusQueryResourceGroupStatusGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewResourceGroupStatusQueryResourceGroupStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResourceGroupStatusQueryResourceGroupStatusOK creates a ResourceGroupStatusQueryResourceGroupStatusOK with default headers values
func NewResourceGroupStatusQueryResourceGroupStatusOK() *ResourceGroupStatusQueryResourceGroupStatusOK {
	return &ResourceGroupStatusQueryResourceGroupStatusOK{}
}

/* ResourceGroupStatusQueryResourceGroupStatusOK describes a response with status code 200, with default header values.

A successful response.
*/
type ResourceGroupStatusQueryResourceGroupStatusOK struct {
	Payload *swagger_models.TagStatusListMsg
}

func (o *ResourceGroupStatusQueryResourceGroupStatusOK) Error() string {
	return fmt.Sprintf("[GET /v1/projects/status][%d] resourceGroupStatusQueryResourceGroupStatusOK  %+v", 200, o.Payload)
}
func (o *ResourceGroupStatusQueryResourceGroupStatusOK) GetPayload() *swagger_models.TagStatusListMsg {
	return o.Payload
}

func (o *ResourceGroupStatusQueryResourceGroupStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.TagStatusListMsg)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupStatusQueryResourceGroupStatusBadRequest creates a ResourceGroupStatusQueryResourceGroupStatusBadRequest with default headers values
func NewResourceGroupStatusQueryResourceGroupStatusBadRequest() *ResourceGroupStatusQueryResourceGroupStatusBadRequest {
	return &ResourceGroupStatusQueryResourceGroupStatusBadRequest{}
}

/* ResourceGroupStatusQueryResourceGroupStatusBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type ResourceGroupStatusQueryResourceGroupStatusBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ResourceGroupStatusQueryResourceGroupStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/projects/status][%d] resourceGroupStatusQueryResourceGroupStatusBadRequest  %+v", 400, o.Payload)
}
func (o *ResourceGroupStatusQueryResourceGroupStatusBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupStatusQueryResourceGroupStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupStatusQueryResourceGroupStatusUnauthorized creates a ResourceGroupStatusQueryResourceGroupStatusUnauthorized with default headers values
func NewResourceGroupStatusQueryResourceGroupStatusUnauthorized() *ResourceGroupStatusQueryResourceGroupStatusUnauthorized {
	return &ResourceGroupStatusQueryResourceGroupStatusUnauthorized{}
}

/* ResourceGroupStatusQueryResourceGroupStatusUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ResourceGroupStatusQueryResourceGroupStatusUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ResourceGroupStatusQueryResourceGroupStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/projects/status][%d] resourceGroupStatusQueryResourceGroupStatusUnauthorized  %+v", 401, o.Payload)
}
func (o *ResourceGroupStatusQueryResourceGroupStatusUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupStatusQueryResourceGroupStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupStatusQueryResourceGroupStatusForbidden creates a ResourceGroupStatusQueryResourceGroupStatusForbidden with default headers values
func NewResourceGroupStatusQueryResourceGroupStatusForbidden() *ResourceGroupStatusQueryResourceGroupStatusForbidden {
	return &ResourceGroupStatusQueryResourceGroupStatusForbidden{}
}

/* ResourceGroupStatusQueryResourceGroupStatusForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ResourceGroupStatusQueryResourceGroupStatusForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ResourceGroupStatusQueryResourceGroupStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/projects/status][%d] resourceGroupStatusQueryResourceGroupStatusForbidden  %+v", 403, o.Payload)
}
func (o *ResourceGroupStatusQueryResourceGroupStatusForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupStatusQueryResourceGroupStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupStatusQueryResourceGroupStatusInternalServerError creates a ResourceGroupStatusQueryResourceGroupStatusInternalServerError with default headers values
func NewResourceGroupStatusQueryResourceGroupStatusInternalServerError() *ResourceGroupStatusQueryResourceGroupStatusInternalServerError {
	return &ResourceGroupStatusQueryResourceGroupStatusInternalServerError{}
}

/* ResourceGroupStatusQueryResourceGroupStatusInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ResourceGroupStatusQueryResourceGroupStatusInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ResourceGroupStatusQueryResourceGroupStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/projects/status][%d] resourceGroupStatusQueryResourceGroupStatusInternalServerError  %+v", 500, o.Payload)
}
func (o *ResourceGroupStatusQueryResourceGroupStatusInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupStatusQueryResourceGroupStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupStatusQueryResourceGroupStatusGatewayTimeout creates a ResourceGroupStatusQueryResourceGroupStatusGatewayTimeout with default headers values
func NewResourceGroupStatusQueryResourceGroupStatusGatewayTimeout() *ResourceGroupStatusQueryResourceGroupStatusGatewayTimeout {
	return &ResourceGroupStatusQueryResourceGroupStatusGatewayTimeout{}
}

/* ResourceGroupStatusQueryResourceGroupStatusGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ResourceGroupStatusQueryResourceGroupStatusGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ResourceGroupStatusQueryResourceGroupStatusGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/projects/status][%d] resourceGroupStatusQueryResourceGroupStatusGatewayTimeout  %+v", 504, o.Payload)
}
func (o *ResourceGroupStatusQueryResourceGroupStatusGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ResourceGroupStatusQueryResourceGroupStatusGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceGroupStatusQueryResourceGroupStatusDefault creates a ResourceGroupStatusQueryResourceGroupStatusDefault with default headers values
func NewResourceGroupStatusQueryResourceGroupStatusDefault(code int) *ResourceGroupStatusQueryResourceGroupStatusDefault {
	return &ResourceGroupStatusQueryResourceGroupStatusDefault{
		_statusCode: code,
	}
}

/* ResourceGroupStatusQueryResourceGroupStatusDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ResourceGroupStatusQueryResourceGroupStatusDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the resource group status query resource group status default response
func (o *ResourceGroupStatusQueryResourceGroupStatusDefault) Code() int {
	return o._statusCode
}

func (o *ResourceGroupStatusQueryResourceGroupStatusDefault) Error() string {
	return fmt.Sprintf("[GET /v1/projects/status][%d] ResourceGroupStatus_QueryResourceGroupStatus default  %+v", o._statusCode, o.Payload)
}
func (o *ResourceGroupStatusQueryResourceGroupStatusDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *ResourceGroupStatusQueryResourceGroupStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
