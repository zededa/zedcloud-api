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

// EdgeNodeConfigurationPreparePowerOffReader is a Reader for the EdgeNodeConfigurationPreparePowerOff structure.
type EdgeNodeConfigurationPreparePowerOffReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeConfigurationPreparePowerOffReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeConfigurationPreparePowerOffOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeConfigurationPreparePowerOffUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeConfigurationPreparePowerOffForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeConfigurationPreparePowerOffNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewEdgeNodeConfigurationPreparePowerOffConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeConfigurationPreparePowerOffInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeConfigurationPreparePowerOffGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeConfigurationPreparePowerOffDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeConfigurationPreparePowerOffOK creates a EdgeNodeConfigurationPreparePowerOffOK with default headers values
func NewEdgeNodeConfigurationPreparePowerOffOK() *EdgeNodeConfigurationPreparePowerOffOK {
	return &EdgeNodeConfigurationPreparePowerOffOK{}
}

/* EdgeNodeConfigurationPreparePowerOffOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeConfigurationPreparePowerOffOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationPreparePowerOffOK) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/preparepoweroff][%d] edgeNodeConfigurationPreparePowerOffOK  %+v", 200, o.Payload)
}
func (o *EdgeNodeConfigurationPreparePowerOffOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationPreparePowerOffOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationPreparePowerOffUnauthorized creates a EdgeNodeConfigurationPreparePowerOffUnauthorized with default headers values
func NewEdgeNodeConfigurationPreparePowerOffUnauthorized() *EdgeNodeConfigurationPreparePowerOffUnauthorized {
	return &EdgeNodeConfigurationPreparePowerOffUnauthorized{}
}

/* EdgeNodeConfigurationPreparePowerOffUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeConfigurationPreparePowerOffUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationPreparePowerOffUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/preparepoweroff][%d] edgeNodeConfigurationPreparePowerOffUnauthorized  %+v", 401, o.Payload)
}
func (o *EdgeNodeConfigurationPreparePowerOffUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationPreparePowerOffUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationPreparePowerOffForbidden creates a EdgeNodeConfigurationPreparePowerOffForbidden with default headers values
func NewEdgeNodeConfigurationPreparePowerOffForbidden() *EdgeNodeConfigurationPreparePowerOffForbidden {
	return &EdgeNodeConfigurationPreparePowerOffForbidden{}
}

/* EdgeNodeConfigurationPreparePowerOffForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeConfigurationPreparePowerOffForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationPreparePowerOffForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/preparepoweroff][%d] edgeNodeConfigurationPreparePowerOffForbidden  %+v", 403, o.Payload)
}
func (o *EdgeNodeConfigurationPreparePowerOffForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationPreparePowerOffForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationPreparePowerOffNotFound creates a EdgeNodeConfigurationPreparePowerOffNotFound with default headers values
func NewEdgeNodeConfigurationPreparePowerOffNotFound() *EdgeNodeConfigurationPreparePowerOffNotFound {
	return &EdgeNodeConfigurationPreparePowerOffNotFound{}
}

/* EdgeNodeConfigurationPreparePowerOffNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeConfigurationPreparePowerOffNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationPreparePowerOffNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/preparepoweroff][%d] edgeNodeConfigurationPreparePowerOffNotFound  %+v", 404, o.Payload)
}
func (o *EdgeNodeConfigurationPreparePowerOffNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationPreparePowerOffNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationPreparePowerOffConflict creates a EdgeNodeConfigurationPreparePowerOffConflict with default headers values
func NewEdgeNodeConfigurationPreparePowerOffConflict() *EdgeNodeConfigurationPreparePowerOffConflict {
	return &EdgeNodeConfigurationPreparePowerOffConflict{}
}

/* EdgeNodeConfigurationPreparePowerOffConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing edge node record.
*/
type EdgeNodeConfigurationPreparePowerOffConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationPreparePowerOffConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/preparepoweroff][%d] edgeNodeConfigurationPreparePowerOffConflict  %+v", 409, o.Payload)
}
func (o *EdgeNodeConfigurationPreparePowerOffConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationPreparePowerOffConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationPreparePowerOffInternalServerError creates a EdgeNodeConfigurationPreparePowerOffInternalServerError with default headers values
func NewEdgeNodeConfigurationPreparePowerOffInternalServerError() *EdgeNodeConfigurationPreparePowerOffInternalServerError {
	return &EdgeNodeConfigurationPreparePowerOffInternalServerError{}
}

/* EdgeNodeConfigurationPreparePowerOffInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeConfigurationPreparePowerOffInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationPreparePowerOffInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/preparepoweroff][%d] edgeNodeConfigurationPreparePowerOffInternalServerError  %+v", 500, o.Payload)
}
func (o *EdgeNodeConfigurationPreparePowerOffInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationPreparePowerOffInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationPreparePowerOffGatewayTimeout creates a EdgeNodeConfigurationPreparePowerOffGatewayTimeout with default headers values
func NewEdgeNodeConfigurationPreparePowerOffGatewayTimeout() *EdgeNodeConfigurationPreparePowerOffGatewayTimeout {
	return &EdgeNodeConfigurationPreparePowerOffGatewayTimeout{}
}

/* EdgeNodeConfigurationPreparePowerOffGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeConfigurationPreparePowerOffGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationPreparePowerOffGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/preparepoweroff][%d] edgeNodeConfigurationPreparePowerOffGatewayTimeout  %+v", 504, o.Payload)
}
func (o *EdgeNodeConfigurationPreparePowerOffGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationPreparePowerOffGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationPreparePowerOffDefault creates a EdgeNodeConfigurationPreparePowerOffDefault with default headers values
func NewEdgeNodeConfigurationPreparePowerOffDefault(code int) *EdgeNodeConfigurationPreparePowerOffDefault {
	return &EdgeNodeConfigurationPreparePowerOffDefault{
		_statusCode: code,
	}
}

/* EdgeNodeConfigurationPreparePowerOffDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeConfigurationPreparePowerOffDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the edge node configuration prepare power off default response
func (o *EdgeNodeConfigurationPreparePowerOffDefault) Code() int {
	return o._statusCode
}

func (o *EdgeNodeConfigurationPreparePowerOffDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/preparepoweroff][%d] EdgeNodeConfiguration_PreparePowerOff default  %+v", o._statusCode, o.Payload)
}
func (o *EdgeNodeConfigurationPreparePowerOffDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeConfigurationPreparePowerOffDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}