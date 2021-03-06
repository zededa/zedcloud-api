// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_node_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// QueryEdgeNodeStatusReader is a Reader for the QueryEdgeNodeStatus structure.
type QueryEdgeNodeStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryEdgeNodeStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryEdgeNodeStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryEdgeNodeStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewQueryEdgeNodeStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryEdgeNodeStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryEdgeNodeStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewQueryEdgeNodeStatusGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryEdgeNodeStatusOK creates a QueryEdgeNodeStatusOK with default headers values
func NewQueryEdgeNodeStatusOK() *QueryEdgeNodeStatusOK {
	return &QueryEdgeNodeStatusOK{}
}

/* QueryEdgeNodeStatusOK describes a response with status code 200, with default header values.

A successful response.
*/
type QueryEdgeNodeStatusOK struct {
	Payload *swagger_models.DeviceStatusListMsg
}

func (o *QueryEdgeNodeStatusOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/status][%d] queryEdgeNodeStatusOK  %+v", 200, o.Payload)
}
func (o *QueryEdgeNodeStatusOK) GetPayload() *swagger_models.DeviceStatusListMsg {
	return o.Payload
}

func (o *QueryEdgeNodeStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.DeviceStatusListMsg)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryEdgeNodeStatusBadRequest creates a QueryEdgeNodeStatusBadRequest with default headers values
func NewQueryEdgeNodeStatusBadRequest() *QueryEdgeNodeStatusBadRequest {
	return &QueryEdgeNodeStatusBadRequest{}
}

/* QueryEdgeNodeStatusBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type QueryEdgeNodeStatusBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryEdgeNodeStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/devices/status][%d] queryEdgeNodeStatusBadRequest  %+v", 400, o.Payload)
}
func (o *QueryEdgeNodeStatusBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryEdgeNodeStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryEdgeNodeStatusUnauthorized creates a QueryEdgeNodeStatusUnauthorized with default headers values
func NewQueryEdgeNodeStatusUnauthorized() *QueryEdgeNodeStatusUnauthorized {
	return &QueryEdgeNodeStatusUnauthorized{}
}

/* QueryEdgeNodeStatusUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type QueryEdgeNodeStatusUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryEdgeNodeStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/status][%d] queryEdgeNodeStatusUnauthorized  %+v", 401, o.Payload)
}
func (o *QueryEdgeNodeStatusUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryEdgeNodeStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryEdgeNodeStatusForbidden creates a QueryEdgeNodeStatusForbidden with default headers values
func NewQueryEdgeNodeStatusForbidden() *QueryEdgeNodeStatusForbidden {
	return &QueryEdgeNodeStatusForbidden{}
}

/* QueryEdgeNodeStatusForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type QueryEdgeNodeStatusForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryEdgeNodeStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/status][%d] queryEdgeNodeStatusForbidden  %+v", 403, o.Payload)
}
func (o *QueryEdgeNodeStatusForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryEdgeNodeStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryEdgeNodeStatusInternalServerError creates a QueryEdgeNodeStatusInternalServerError with default headers values
func NewQueryEdgeNodeStatusInternalServerError() *QueryEdgeNodeStatusInternalServerError {
	return &QueryEdgeNodeStatusInternalServerError{}
}

/* QueryEdgeNodeStatusInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type QueryEdgeNodeStatusInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryEdgeNodeStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/status][%d] queryEdgeNodeStatusInternalServerError  %+v", 500, o.Payload)
}
func (o *QueryEdgeNodeStatusInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryEdgeNodeStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryEdgeNodeStatusGatewayTimeout creates a QueryEdgeNodeStatusGatewayTimeout with default headers values
func NewQueryEdgeNodeStatusGatewayTimeout() *QueryEdgeNodeStatusGatewayTimeout {
	return &QueryEdgeNodeStatusGatewayTimeout{}
}

/* QueryEdgeNodeStatusGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type QueryEdgeNodeStatusGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryEdgeNodeStatusGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/status][%d] queryEdgeNodeStatusGatewayTimeout  %+v", 504, o.Payload)
}
func (o *QueryEdgeNodeStatusGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryEdgeNodeStatusGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
