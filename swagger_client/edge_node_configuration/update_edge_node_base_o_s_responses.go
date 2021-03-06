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

// UpdateEdgeNodeBaseOSReader is a Reader for the UpdateEdgeNodeBaseOS structure.
type UpdateEdgeNodeBaseOSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateEdgeNodeBaseOSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateEdgeNodeBaseOSOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateEdgeNodeBaseOSUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateEdgeNodeBaseOSForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateEdgeNodeBaseOSNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateEdgeNodeBaseOSConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateEdgeNodeBaseOSInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewUpdateEdgeNodeBaseOSGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateEdgeNodeBaseOSOK creates a UpdateEdgeNodeBaseOSOK with default headers values
func NewUpdateEdgeNodeBaseOSOK() *UpdateEdgeNodeBaseOSOK {
	return &UpdateEdgeNodeBaseOSOK{}
}

/* UpdateEdgeNodeBaseOSOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateEdgeNodeBaseOSOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateEdgeNodeBaseOSOK) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/apply][%d] updateEdgeNodeBaseOSOK  %+v", 200, o.Payload)
}
func (o *UpdateEdgeNodeBaseOSOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateEdgeNodeBaseOSOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEdgeNodeBaseOSUnauthorized creates a UpdateEdgeNodeBaseOSUnauthorized with default headers values
func NewUpdateEdgeNodeBaseOSUnauthorized() *UpdateEdgeNodeBaseOSUnauthorized {
	return &UpdateEdgeNodeBaseOSUnauthorized{}
}

/* UpdateEdgeNodeBaseOSUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type UpdateEdgeNodeBaseOSUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateEdgeNodeBaseOSUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/apply][%d] updateEdgeNodeBaseOSUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateEdgeNodeBaseOSUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateEdgeNodeBaseOSUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEdgeNodeBaseOSForbidden creates a UpdateEdgeNodeBaseOSForbidden with default headers values
func NewUpdateEdgeNodeBaseOSForbidden() *UpdateEdgeNodeBaseOSForbidden {
	return &UpdateEdgeNodeBaseOSForbidden{}
}

/* UpdateEdgeNodeBaseOSForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type UpdateEdgeNodeBaseOSForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateEdgeNodeBaseOSForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/apply][%d] updateEdgeNodeBaseOSForbidden  %+v", 403, o.Payload)
}
func (o *UpdateEdgeNodeBaseOSForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateEdgeNodeBaseOSForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEdgeNodeBaseOSNotFound creates a UpdateEdgeNodeBaseOSNotFound with default headers values
func NewUpdateEdgeNodeBaseOSNotFound() *UpdateEdgeNodeBaseOSNotFound {
	return &UpdateEdgeNodeBaseOSNotFound{}
}

/* UpdateEdgeNodeBaseOSNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type UpdateEdgeNodeBaseOSNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateEdgeNodeBaseOSNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/apply][%d] updateEdgeNodeBaseOSNotFound  %+v", 404, o.Payload)
}
func (o *UpdateEdgeNodeBaseOSNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateEdgeNodeBaseOSNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEdgeNodeBaseOSConflict creates a UpdateEdgeNodeBaseOSConflict with default headers values
func NewUpdateEdgeNodeBaseOSConflict() *UpdateEdgeNodeBaseOSConflict {
	return &UpdateEdgeNodeBaseOSConflict{}
}

/* UpdateEdgeNodeBaseOSConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing edge node record.
*/
type UpdateEdgeNodeBaseOSConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateEdgeNodeBaseOSConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/apply][%d] updateEdgeNodeBaseOSConflict  %+v", 409, o.Payload)
}
func (o *UpdateEdgeNodeBaseOSConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateEdgeNodeBaseOSConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEdgeNodeBaseOSInternalServerError creates a UpdateEdgeNodeBaseOSInternalServerError with default headers values
func NewUpdateEdgeNodeBaseOSInternalServerError() *UpdateEdgeNodeBaseOSInternalServerError {
	return &UpdateEdgeNodeBaseOSInternalServerError{}
}

/* UpdateEdgeNodeBaseOSInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type UpdateEdgeNodeBaseOSInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateEdgeNodeBaseOSInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/apply][%d] updateEdgeNodeBaseOSInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateEdgeNodeBaseOSInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateEdgeNodeBaseOSInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEdgeNodeBaseOSGatewayTimeout creates a UpdateEdgeNodeBaseOSGatewayTimeout with default headers values
func NewUpdateEdgeNodeBaseOSGatewayTimeout() *UpdateEdgeNodeBaseOSGatewayTimeout {
	return &UpdateEdgeNodeBaseOSGatewayTimeout{}
}

/* UpdateEdgeNodeBaseOSGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type UpdateEdgeNodeBaseOSGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateEdgeNodeBaseOSGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/apply][%d] updateEdgeNodeBaseOSGatewayTimeout  %+v", 504, o.Payload)
}
func (o *UpdateEdgeNodeBaseOSGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateEdgeNodeBaseOSGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
