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

// ActivateEdgeNodeReader is a Reader for the ActivateEdgeNode structure.
type ActivateEdgeNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActivateEdgeNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActivateEdgeNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewActivateEdgeNodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewActivateEdgeNodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewActivateEdgeNodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewActivateEdgeNodeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewActivateEdgeNodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewActivateEdgeNodeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewActivateEdgeNodeOK creates a ActivateEdgeNodeOK with default headers values
func NewActivateEdgeNodeOK() *ActivateEdgeNodeOK {
	return &ActivateEdgeNodeOK{}
}

/* ActivateEdgeNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type ActivateEdgeNodeOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ActivateEdgeNodeOK) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/activate][%d] activateEdgeNodeOK  %+v", 200, o.Payload)
}
func (o *ActivateEdgeNodeOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ActivateEdgeNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActivateEdgeNodeUnauthorized creates a ActivateEdgeNodeUnauthorized with default headers values
func NewActivateEdgeNodeUnauthorized() *ActivateEdgeNodeUnauthorized {
	return &ActivateEdgeNodeUnauthorized{}
}

/* ActivateEdgeNodeUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ActivateEdgeNodeUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ActivateEdgeNodeUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/activate][%d] activateEdgeNodeUnauthorized  %+v", 401, o.Payload)
}
func (o *ActivateEdgeNodeUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ActivateEdgeNodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActivateEdgeNodeForbidden creates a ActivateEdgeNodeForbidden with default headers values
func NewActivateEdgeNodeForbidden() *ActivateEdgeNodeForbidden {
	return &ActivateEdgeNodeForbidden{}
}

/* ActivateEdgeNodeForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ActivateEdgeNodeForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ActivateEdgeNodeForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/activate][%d] activateEdgeNodeForbidden  %+v", 403, o.Payload)
}
func (o *ActivateEdgeNodeForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ActivateEdgeNodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActivateEdgeNodeNotFound creates a ActivateEdgeNodeNotFound with default headers values
func NewActivateEdgeNodeNotFound() *ActivateEdgeNodeNotFound {
	return &ActivateEdgeNodeNotFound{}
}

/* ActivateEdgeNodeNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type ActivateEdgeNodeNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ActivateEdgeNodeNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/activate][%d] activateEdgeNodeNotFound  %+v", 404, o.Payload)
}
func (o *ActivateEdgeNodeNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ActivateEdgeNodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActivateEdgeNodeConflict creates a ActivateEdgeNodeConflict with default headers values
func NewActivateEdgeNodeConflict() *ActivateEdgeNodeConflict {
	return &ActivateEdgeNodeConflict{}
}

/* ActivateEdgeNodeConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing edge node record.
*/
type ActivateEdgeNodeConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ActivateEdgeNodeConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/activate][%d] activateEdgeNodeConflict  %+v", 409, o.Payload)
}
func (o *ActivateEdgeNodeConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ActivateEdgeNodeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActivateEdgeNodeInternalServerError creates a ActivateEdgeNodeInternalServerError with default headers values
func NewActivateEdgeNodeInternalServerError() *ActivateEdgeNodeInternalServerError {
	return &ActivateEdgeNodeInternalServerError{}
}

/* ActivateEdgeNodeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ActivateEdgeNodeInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ActivateEdgeNodeInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/activate][%d] activateEdgeNodeInternalServerError  %+v", 500, o.Payload)
}
func (o *ActivateEdgeNodeInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ActivateEdgeNodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActivateEdgeNodeGatewayTimeout creates a ActivateEdgeNodeGatewayTimeout with default headers values
func NewActivateEdgeNodeGatewayTimeout() *ActivateEdgeNodeGatewayTimeout {
	return &ActivateEdgeNodeGatewayTimeout{}
}

/* ActivateEdgeNodeGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ActivateEdgeNodeGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *ActivateEdgeNodeGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/activate][%d] activateEdgeNodeGatewayTimeout  %+v", 504, o.Payload)
}
func (o *ActivateEdgeNodeGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *ActivateEdgeNodeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
