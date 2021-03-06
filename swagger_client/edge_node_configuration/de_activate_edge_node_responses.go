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

// DeActivateEdgeNodeReader is a Reader for the DeActivateEdgeNode structure.
type DeActivateEdgeNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeActivateEdgeNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeActivateEdgeNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeActivateEdgeNodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeActivateEdgeNodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeActivateEdgeNodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeActivateEdgeNodeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeActivateEdgeNodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeActivateEdgeNodeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeActivateEdgeNodeOK creates a DeActivateEdgeNodeOK with default headers values
func NewDeActivateEdgeNodeOK() *DeActivateEdgeNodeOK {
	return &DeActivateEdgeNodeOK{}
}

/* DeActivateEdgeNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeActivateEdgeNodeOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *DeActivateEdgeNodeOK) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/deactivate][%d] deActivateEdgeNodeOK  %+v", 200, o.Payload)
}
func (o *DeActivateEdgeNodeOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DeActivateEdgeNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeActivateEdgeNodeUnauthorized creates a DeActivateEdgeNodeUnauthorized with default headers values
func NewDeActivateEdgeNodeUnauthorized() *DeActivateEdgeNodeUnauthorized {
	return &DeActivateEdgeNodeUnauthorized{}
}

/* DeActivateEdgeNodeUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type DeActivateEdgeNodeUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *DeActivateEdgeNodeUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/deactivate][%d] deActivateEdgeNodeUnauthorized  %+v", 401, o.Payload)
}
func (o *DeActivateEdgeNodeUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DeActivateEdgeNodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeActivateEdgeNodeForbidden creates a DeActivateEdgeNodeForbidden with default headers values
func NewDeActivateEdgeNodeForbidden() *DeActivateEdgeNodeForbidden {
	return &DeActivateEdgeNodeForbidden{}
}

/* DeActivateEdgeNodeForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type DeActivateEdgeNodeForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *DeActivateEdgeNodeForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/deactivate][%d] deActivateEdgeNodeForbidden  %+v", 403, o.Payload)
}
func (o *DeActivateEdgeNodeForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DeActivateEdgeNodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeActivateEdgeNodeNotFound creates a DeActivateEdgeNodeNotFound with default headers values
func NewDeActivateEdgeNodeNotFound() *DeActivateEdgeNodeNotFound {
	return &DeActivateEdgeNodeNotFound{}
}

/* DeActivateEdgeNodeNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type DeActivateEdgeNodeNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *DeActivateEdgeNodeNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/deactivate][%d] deActivateEdgeNodeNotFound  %+v", 404, o.Payload)
}
func (o *DeActivateEdgeNodeNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DeActivateEdgeNodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeActivateEdgeNodeConflict creates a DeActivateEdgeNodeConflict with default headers values
func NewDeActivateEdgeNodeConflict() *DeActivateEdgeNodeConflict {
	return &DeActivateEdgeNodeConflict{}
}

/* DeActivateEdgeNodeConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing edge node record.
*/
type DeActivateEdgeNodeConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *DeActivateEdgeNodeConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/deactivate][%d] deActivateEdgeNodeConflict  %+v", 409, o.Payload)
}
func (o *DeActivateEdgeNodeConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DeActivateEdgeNodeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeActivateEdgeNodeInternalServerError creates a DeActivateEdgeNodeInternalServerError with default headers values
func NewDeActivateEdgeNodeInternalServerError() *DeActivateEdgeNodeInternalServerError {
	return &DeActivateEdgeNodeInternalServerError{}
}

/* DeActivateEdgeNodeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type DeActivateEdgeNodeInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *DeActivateEdgeNodeInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/deactivate][%d] deActivateEdgeNodeInternalServerError  %+v", 500, o.Payload)
}
func (o *DeActivateEdgeNodeInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DeActivateEdgeNodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeActivateEdgeNodeGatewayTimeout creates a DeActivateEdgeNodeGatewayTimeout with default headers values
func NewDeActivateEdgeNodeGatewayTimeout() *DeActivateEdgeNodeGatewayTimeout {
	return &DeActivateEdgeNodeGatewayTimeout{}
}

/* DeActivateEdgeNodeGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type DeActivateEdgeNodeGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *DeActivateEdgeNodeGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/deactivate][%d] deActivateEdgeNodeGatewayTimeout  %+v", 504, o.Payload)
}
func (o *DeActivateEdgeNodeGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DeActivateEdgeNodeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
