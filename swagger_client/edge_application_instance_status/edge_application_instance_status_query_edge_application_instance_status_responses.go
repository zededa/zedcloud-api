// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_application_instance_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusReader is a Reader for the EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatus structure.
type EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusOK creates a EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusOK with default headers values
func NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusOK() *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusOK {
	return &EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusOK{}
}

/* EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusOK struct {
	Payload *swagger_models.AppInstStatusListMsg
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/status][%d] edgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusOK  %+v", 200, o.Payload)
}
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusOK) GetPayload() *swagger_models.AppInstStatusListMsg {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.AppInstStatusListMsg)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusBadRequest creates a EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusBadRequest with default headers values
func NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusBadRequest() *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusBadRequest {
	return &EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusBadRequest{}
}

/* EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/status][%d] edgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusBadRequest  %+v", 400, o.Payload)
}
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusUnauthorized creates a EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusUnauthorized with default headers values
func NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusUnauthorized() *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusUnauthorized {
	return &EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusUnauthorized{}
}

/* EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/status][%d] edgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusUnauthorized  %+v", 401, o.Payload)
}
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusForbidden creates a EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusForbidden with default headers values
func NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusForbidden() *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusForbidden {
	return &EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusForbidden{}
}

/* EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/status][%d] edgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusForbidden  %+v", 403, o.Payload)
}
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusInternalServerError creates a EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusInternalServerError with default headers values
func NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusInternalServerError() *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusInternalServerError {
	return &EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusInternalServerError{}
}

/* EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/status][%d] edgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusInternalServerError  %+v", 500, o.Payload)
}
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusGatewayTimeout creates a EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusGatewayTimeout with default headers values
func NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusGatewayTimeout() *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusGatewayTimeout {
	return &EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusGatewayTimeout{}
}

/* EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/status][%d] edgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusGatewayTimeout  %+v", 504, o.Payload)
}
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusDefault creates a EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusDefault with default headers values
func NewEdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusDefault(code int) *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusDefault {
	return &EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusDefault{
		_statusCode: code,
	}
}

/* EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the edge application instance status query edge application instance status default response
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusDefault) Code() int {
	return o._statusCode
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusDefault) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/status][%d] EdgeApplicationInstanceStatus_QueryEdgeApplicationInstanceStatus default  %+v", o._statusCode, o.Payload)
}
func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusQueryEdgeApplicationInstanceStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
