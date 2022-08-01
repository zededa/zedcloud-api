// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_node_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// EdgeNodeConfigurationRebootReader is a Reader for the EdgeNodeConfigurationReboot structure.
type EdgeNodeConfigurationRebootReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeConfigurationRebootReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeConfigurationRebootOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeConfigurationRebootUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeConfigurationRebootForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeConfigurationRebootNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewEdgeNodeConfigurationRebootConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeConfigurationRebootInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeConfigurationRebootGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeConfigurationRebootDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeConfigurationRebootOK creates a EdgeNodeConfigurationRebootOK with default headers values
func NewEdgeNodeConfigurationRebootOK() *EdgeNodeConfigurationRebootOK {
	return &EdgeNodeConfigurationRebootOK{}
}

/* EdgeNodeConfigurationRebootOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeConfigurationRebootOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationRebootOK) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/reboot][%d] edgeNodeConfigurationRebootOK  %+v", 200, o.Payload)
}
func (o *EdgeNodeConfigurationRebootOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationRebootOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationRebootUnauthorized creates a EdgeNodeConfigurationRebootUnauthorized with default headers values
func NewEdgeNodeConfigurationRebootUnauthorized() *EdgeNodeConfigurationRebootUnauthorized {
	return &EdgeNodeConfigurationRebootUnauthorized{}
}

/* EdgeNodeConfigurationRebootUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeConfigurationRebootUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationRebootUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/reboot][%d] edgeNodeConfigurationRebootUnauthorized  %+v", 401, o.Payload)
}
func (o *EdgeNodeConfigurationRebootUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationRebootUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationRebootForbidden creates a EdgeNodeConfigurationRebootForbidden with default headers values
func NewEdgeNodeConfigurationRebootForbidden() *EdgeNodeConfigurationRebootForbidden {
	return &EdgeNodeConfigurationRebootForbidden{}
}

/* EdgeNodeConfigurationRebootForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeConfigurationRebootForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationRebootForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/reboot][%d] edgeNodeConfigurationRebootForbidden  %+v", 403, o.Payload)
}
func (o *EdgeNodeConfigurationRebootForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationRebootForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationRebootNotFound creates a EdgeNodeConfigurationRebootNotFound with default headers values
func NewEdgeNodeConfigurationRebootNotFound() *EdgeNodeConfigurationRebootNotFound {
	return &EdgeNodeConfigurationRebootNotFound{}
}

/* EdgeNodeConfigurationRebootNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeConfigurationRebootNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationRebootNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/reboot][%d] edgeNodeConfigurationRebootNotFound  %+v", 404, o.Payload)
}
func (o *EdgeNodeConfigurationRebootNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationRebootNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationRebootConflict creates a EdgeNodeConfigurationRebootConflict with default headers values
func NewEdgeNodeConfigurationRebootConflict() *EdgeNodeConfigurationRebootConflict {
	return &EdgeNodeConfigurationRebootConflict{}
}

/* EdgeNodeConfigurationRebootConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing edge node record.
*/
type EdgeNodeConfigurationRebootConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationRebootConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/reboot][%d] edgeNodeConfigurationRebootConflict  %+v", 409, o.Payload)
}
func (o *EdgeNodeConfigurationRebootConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationRebootConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationRebootInternalServerError creates a EdgeNodeConfigurationRebootInternalServerError with default headers values
func NewEdgeNodeConfigurationRebootInternalServerError() *EdgeNodeConfigurationRebootInternalServerError {
	return &EdgeNodeConfigurationRebootInternalServerError{}
}

/* EdgeNodeConfigurationRebootInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeConfigurationRebootInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationRebootInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/reboot][%d] edgeNodeConfigurationRebootInternalServerError  %+v", 500, o.Payload)
}
func (o *EdgeNodeConfigurationRebootInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationRebootInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationRebootGatewayTimeout creates a EdgeNodeConfigurationRebootGatewayTimeout with default headers values
func NewEdgeNodeConfigurationRebootGatewayTimeout() *EdgeNodeConfigurationRebootGatewayTimeout {
	return &EdgeNodeConfigurationRebootGatewayTimeout{}
}

/* EdgeNodeConfigurationRebootGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeConfigurationRebootGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationRebootGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/reboot][%d] edgeNodeConfigurationRebootGatewayTimeout  %+v", 504, o.Payload)
}
func (o *EdgeNodeConfigurationRebootGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationRebootGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationRebootDefault creates a EdgeNodeConfigurationRebootDefault with default headers values
func NewEdgeNodeConfigurationRebootDefault(code int) *EdgeNodeConfigurationRebootDefault {
	return &EdgeNodeConfigurationRebootDefault{
		_statusCode: code,
	}
}

/* EdgeNodeConfigurationRebootDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeConfigurationRebootDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the edge node configuration reboot default response
func (o *EdgeNodeConfigurationRebootDefault) Code() int {
	return o._statusCode
}

func (o *EdgeNodeConfigurationRebootDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/reboot][%d] EdgeNodeConfiguration_Reboot default  %+v", o._statusCode, o.Payload)
}
func (o *EdgeNodeConfigurationRebootDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeConfigurationRebootDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}