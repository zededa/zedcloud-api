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

// EdgeNodeConfigurationOffboardReader is a Reader for the EdgeNodeConfigurationOffboard structure.
type EdgeNodeConfigurationOffboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeConfigurationOffboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeConfigurationOffboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeConfigurationOffboardUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeConfigurationOffboardForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeConfigurationOffboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeConfigurationOffboardInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeConfigurationOffboardGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeConfigurationOffboardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeConfigurationOffboardOK creates a EdgeNodeConfigurationOffboardOK with default headers values
func NewEdgeNodeConfigurationOffboardOK() *EdgeNodeConfigurationOffboardOK {
	return &EdgeNodeConfigurationOffboardOK{}
}

/* EdgeNodeConfigurationOffboardOK describes a response with status code 200, with default header values.

Success. The API gateway offboarded the edge-node successfully..
*/
type EdgeNodeConfigurationOffboardOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationOffboardOK) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/Offboard][%d] edgeNodeConfigurationOffboardOK  %+v", 200, o.Payload)
}
func (o *EdgeNodeConfigurationOffboardOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationOffboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationOffboardUnauthorized creates a EdgeNodeConfigurationOffboardUnauthorized with default headers values
func NewEdgeNodeConfigurationOffboardUnauthorized() *EdgeNodeConfigurationOffboardUnauthorized {
	return &EdgeNodeConfigurationOffboardUnauthorized{}
}

/* EdgeNodeConfigurationOffboardUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeConfigurationOffboardUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationOffboardUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/Offboard][%d] edgeNodeConfigurationOffboardUnauthorized  %+v", 401, o.Payload)
}
func (o *EdgeNodeConfigurationOffboardUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationOffboardUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationOffboardForbidden creates a EdgeNodeConfigurationOffboardForbidden with default headers values
func NewEdgeNodeConfigurationOffboardForbidden() *EdgeNodeConfigurationOffboardForbidden {
	return &EdgeNodeConfigurationOffboardForbidden{}
}

/* EdgeNodeConfigurationOffboardForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeConfigurationOffboardForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationOffboardForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/Offboard][%d] edgeNodeConfigurationOffboardForbidden  %+v", 403, o.Payload)
}
func (o *EdgeNodeConfigurationOffboardForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationOffboardForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationOffboardNotFound creates a EdgeNodeConfigurationOffboardNotFound with default headers values
func NewEdgeNodeConfigurationOffboardNotFound() *EdgeNodeConfigurationOffboardNotFound {
	return &EdgeNodeConfigurationOffboardNotFound{}
}

/* EdgeNodeConfigurationOffboardNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeConfigurationOffboardNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationOffboardNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/Offboard][%d] edgeNodeConfigurationOffboardNotFound  %+v", 404, o.Payload)
}
func (o *EdgeNodeConfigurationOffboardNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationOffboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationOffboardInternalServerError creates a EdgeNodeConfigurationOffboardInternalServerError with default headers values
func NewEdgeNodeConfigurationOffboardInternalServerError() *EdgeNodeConfigurationOffboardInternalServerError {
	return &EdgeNodeConfigurationOffboardInternalServerError{}
}

/* EdgeNodeConfigurationOffboardInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeConfigurationOffboardInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationOffboardInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/Offboard][%d] edgeNodeConfigurationOffboardInternalServerError  %+v", 500, o.Payload)
}
func (o *EdgeNodeConfigurationOffboardInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationOffboardInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationOffboardGatewayTimeout creates a EdgeNodeConfigurationOffboardGatewayTimeout with default headers values
func NewEdgeNodeConfigurationOffboardGatewayTimeout() *EdgeNodeConfigurationOffboardGatewayTimeout {
	return &EdgeNodeConfigurationOffboardGatewayTimeout{}
}

/* EdgeNodeConfigurationOffboardGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeConfigurationOffboardGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationOffboardGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/Offboard][%d] edgeNodeConfigurationOffboardGatewayTimeout  %+v", 504, o.Payload)
}
func (o *EdgeNodeConfigurationOffboardGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationOffboardGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationOffboardDefault creates a EdgeNodeConfigurationOffboardDefault with default headers values
func NewEdgeNodeConfigurationOffboardDefault(code int) *EdgeNodeConfigurationOffboardDefault {
	return &EdgeNodeConfigurationOffboardDefault{
		_statusCode: code,
	}
}

/* EdgeNodeConfigurationOffboardDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeConfigurationOffboardDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the edge node configuration offboard default response
func (o *EdgeNodeConfigurationOffboardDefault) Code() int {
	return o._statusCode
}

func (o *EdgeNodeConfigurationOffboardDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/Offboard][%d] EdgeNodeConfiguration_Offboard default  %+v", o._statusCode, o.Payload)
}
func (o *EdgeNodeConfigurationOffboardDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeConfigurationOffboardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
