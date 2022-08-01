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

// EdgeNodeConfigurationDeActivateEdgeNodeReader is a Reader for the EdgeNodeConfigurationDeActivateEdgeNode structure.
type EdgeNodeConfigurationDeActivateEdgeNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeConfigurationDeActivateEdgeNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeConfigurationDeActivateEdgeNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeConfigurationDeActivateEdgeNodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeConfigurationDeActivateEdgeNodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeConfigurationDeActivateEdgeNodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewEdgeNodeConfigurationDeActivateEdgeNodeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeConfigurationDeActivateEdgeNodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeConfigurationDeActivateEdgeNodeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeConfigurationDeActivateEdgeNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeConfigurationDeActivateEdgeNodeOK creates a EdgeNodeConfigurationDeActivateEdgeNodeOK with default headers values
func NewEdgeNodeConfigurationDeActivateEdgeNodeOK() *EdgeNodeConfigurationDeActivateEdgeNodeOK {
	return &EdgeNodeConfigurationDeActivateEdgeNodeOK{}
}

/* EdgeNodeConfigurationDeActivateEdgeNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeConfigurationDeActivateEdgeNodeOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationDeActivateEdgeNodeOK) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/deactivate][%d] edgeNodeConfigurationDeActivateEdgeNodeOK  %+v", 200, o.Payload)
}
func (o *EdgeNodeConfigurationDeActivateEdgeNodeOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationDeActivateEdgeNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationDeActivateEdgeNodeUnauthorized creates a EdgeNodeConfigurationDeActivateEdgeNodeUnauthorized with default headers values
func NewEdgeNodeConfigurationDeActivateEdgeNodeUnauthorized() *EdgeNodeConfigurationDeActivateEdgeNodeUnauthorized {
	return &EdgeNodeConfigurationDeActivateEdgeNodeUnauthorized{}
}

/* EdgeNodeConfigurationDeActivateEdgeNodeUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeConfigurationDeActivateEdgeNodeUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationDeActivateEdgeNodeUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/deactivate][%d] edgeNodeConfigurationDeActivateEdgeNodeUnauthorized  %+v", 401, o.Payload)
}
func (o *EdgeNodeConfigurationDeActivateEdgeNodeUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationDeActivateEdgeNodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationDeActivateEdgeNodeForbidden creates a EdgeNodeConfigurationDeActivateEdgeNodeForbidden with default headers values
func NewEdgeNodeConfigurationDeActivateEdgeNodeForbidden() *EdgeNodeConfigurationDeActivateEdgeNodeForbidden {
	return &EdgeNodeConfigurationDeActivateEdgeNodeForbidden{}
}

/* EdgeNodeConfigurationDeActivateEdgeNodeForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeConfigurationDeActivateEdgeNodeForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationDeActivateEdgeNodeForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/deactivate][%d] edgeNodeConfigurationDeActivateEdgeNodeForbidden  %+v", 403, o.Payload)
}
func (o *EdgeNodeConfigurationDeActivateEdgeNodeForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationDeActivateEdgeNodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationDeActivateEdgeNodeNotFound creates a EdgeNodeConfigurationDeActivateEdgeNodeNotFound with default headers values
func NewEdgeNodeConfigurationDeActivateEdgeNodeNotFound() *EdgeNodeConfigurationDeActivateEdgeNodeNotFound {
	return &EdgeNodeConfigurationDeActivateEdgeNodeNotFound{}
}

/* EdgeNodeConfigurationDeActivateEdgeNodeNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeConfigurationDeActivateEdgeNodeNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationDeActivateEdgeNodeNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/deactivate][%d] edgeNodeConfigurationDeActivateEdgeNodeNotFound  %+v", 404, o.Payload)
}
func (o *EdgeNodeConfigurationDeActivateEdgeNodeNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationDeActivateEdgeNodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationDeActivateEdgeNodeConflict creates a EdgeNodeConfigurationDeActivateEdgeNodeConflict with default headers values
func NewEdgeNodeConfigurationDeActivateEdgeNodeConflict() *EdgeNodeConfigurationDeActivateEdgeNodeConflict {
	return &EdgeNodeConfigurationDeActivateEdgeNodeConflict{}
}

/* EdgeNodeConfigurationDeActivateEdgeNodeConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing edge node record.
*/
type EdgeNodeConfigurationDeActivateEdgeNodeConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationDeActivateEdgeNodeConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/deactivate][%d] edgeNodeConfigurationDeActivateEdgeNodeConflict  %+v", 409, o.Payload)
}
func (o *EdgeNodeConfigurationDeActivateEdgeNodeConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationDeActivateEdgeNodeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationDeActivateEdgeNodeInternalServerError creates a EdgeNodeConfigurationDeActivateEdgeNodeInternalServerError with default headers values
func NewEdgeNodeConfigurationDeActivateEdgeNodeInternalServerError() *EdgeNodeConfigurationDeActivateEdgeNodeInternalServerError {
	return &EdgeNodeConfigurationDeActivateEdgeNodeInternalServerError{}
}

/* EdgeNodeConfigurationDeActivateEdgeNodeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeConfigurationDeActivateEdgeNodeInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationDeActivateEdgeNodeInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/deactivate][%d] edgeNodeConfigurationDeActivateEdgeNodeInternalServerError  %+v", 500, o.Payload)
}
func (o *EdgeNodeConfigurationDeActivateEdgeNodeInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationDeActivateEdgeNodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationDeActivateEdgeNodeGatewayTimeout creates a EdgeNodeConfigurationDeActivateEdgeNodeGatewayTimeout with default headers values
func NewEdgeNodeConfigurationDeActivateEdgeNodeGatewayTimeout() *EdgeNodeConfigurationDeActivateEdgeNodeGatewayTimeout {
	return &EdgeNodeConfigurationDeActivateEdgeNodeGatewayTimeout{}
}

/* EdgeNodeConfigurationDeActivateEdgeNodeGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeConfigurationDeActivateEdgeNodeGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeNodeConfigurationDeActivateEdgeNodeGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/deactivate][%d] edgeNodeConfigurationDeActivateEdgeNodeGatewayTimeout  %+v", 504, o.Payload)
}
func (o *EdgeNodeConfigurationDeActivateEdgeNodeGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationDeActivateEdgeNodeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationDeActivateEdgeNodeDefault creates a EdgeNodeConfigurationDeActivateEdgeNodeDefault with default headers values
func NewEdgeNodeConfigurationDeActivateEdgeNodeDefault(code int) *EdgeNodeConfigurationDeActivateEdgeNodeDefault {
	return &EdgeNodeConfigurationDeActivateEdgeNodeDefault{
		_statusCode: code,
	}
}

/* EdgeNodeConfigurationDeActivateEdgeNodeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeConfigurationDeActivateEdgeNodeDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the edge node configuration de activate edge node default response
func (o *EdgeNodeConfigurationDeActivateEdgeNodeDefault) Code() int {
	return o._statusCode
}

func (o *EdgeNodeConfigurationDeActivateEdgeNodeDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/deactivate][%d] EdgeNodeConfiguration_DeActivateEdgeNode default  %+v", o._statusCode, o.Payload)
}
func (o *EdgeNodeConfigurationDeActivateEdgeNodeDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeConfigurationDeActivateEdgeNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}