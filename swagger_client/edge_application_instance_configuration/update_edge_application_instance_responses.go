// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_application_instance_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// UpdateEdgeApplicationInstanceReader is a Reader for the UpdateEdgeApplicationInstance structure.
type UpdateEdgeApplicationInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateEdgeApplicationInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateEdgeApplicationInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateEdgeApplicationInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateEdgeApplicationInstanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateEdgeApplicationInstanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateEdgeApplicationInstanceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateEdgeApplicationInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewUpdateEdgeApplicationInstanceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateEdgeApplicationInstanceOK creates a UpdateEdgeApplicationInstanceOK with default headers values
func NewUpdateEdgeApplicationInstanceOK() *UpdateEdgeApplicationInstanceOK {
	return &UpdateEdgeApplicationInstanceOK{}
}

/* UpdateEdgeApplicationInstanceOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateEdgeApplicationInstanceOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateEdgeApplicationInstanceOK) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}][%d] updateEdgeApplicationInstanceOK  %+v", 200, o.Payload)
}
func (o *UpdateEdgeApplicationInstanceOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateEdgeApplicationInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEdgeApplicationInstanceUnauthorized creates a UpdateEdgeApplicationInstanceUnauthorized with default headers values
func NewUpdateEdgeApplicationInstanceUnauthorized() *UpdateEdgeApplicationInstanceUnauthorized {
	return &UpdateEdgeApplicationInstanceUnauthorized{}
}

/* UpdateEdgeApplicationInstanceUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type UpdateEdgeApplicationInstanceUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateEdgeApplicationInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}][%d] updateEdgeApplicationInstanceUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateEdgeApplicationInstanceUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateEdgeApplicationInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEdgeApplicationInstanceForbidden creates a UpdateEdgeApplicationInstanceForbidden with default headers values
func NewUpdateEdgeApplicationInstanceForbidden() *UpdateEdgeApplicationInstanceForbidden {
	return &UpdateEdgeApplicationInstanceForbidden{}
}

/* UpdateEdgeApplicationInstanceForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type UpdateEdgeApplicationInstanceForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateEdgeApplicationInstanceForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}][%d] updateEdgeApplicationInstanceForbidden  %+v", 403, o.Payload)
}
func (o *UpdateEdgeApplicationInstanceForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateEdgeApplicationInstanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEdgeApplicationInstanceNotFound creates a UpdateEdgeApplicationInstanceNotFound with default headers values
func NewUpdateEdgeApplicationInstanceNotFound() *UpdateEdgeApplicationInstanceNotFound {
	return &UpdateEdgeApplicationInstanceNotFound{}
}

/* UpdateEdgeApplicationInstanceNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type UpdateEdgeApplicationInstanceNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateEdgeApplicationInstanceNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}][%d] updateEdgeApplicationInstanceNotFound  %+v", 404, o.Payload)
}
func (o *UpdateEdgeApplicationInstanceNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateEdgeApplicationInstanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEdgeApplicationInstanceConflict creates a UpdateEdgeApplicationInstanceConflict with default headers values
func NewUpdateEdgeApplicationInstanceConflict() *UpdateEdgeApplicationInstanceConflict {
	return &UpdateEdgeApplicationInstanceConflict{}
}

/* UpdateEdgeApplicationInstanceConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing edge network record.
*/
type UpdateEdgeApplicationInstanceConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateEdgeApplicationInstanceConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}][%d] updateEdgeApplicationInstanceConflict  %+v", 409, o.Payload)
}
func (o *UpdateEdgeApplicationInstanceConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateEdgeApplicationInstanceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEdgeApplicationInstanceInternalServerError creates a UpdateEdgeApplicationInstanceInternalServerError with default headers values
func NewUpdateEdgeApplicationInstanceInternalServerError() *UpdateEdgeApplicationInstanceInternalServerError {
	return &UpdateEdgeApplicationInstanceInternalServerError{}
}

/* UpdateEdgeApplicationInstanceInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type UpdateEdgeApplicationInstanceInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateEdgeApplicationInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}][%d] updateEdgeApplicationInstanceInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateEdgeApplicationInstanceInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateEdgeApplicationInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEdgeApplicationInstanceGatewayTimeout creates a UpdateEdgeApplicationInstanceGatewayTimeout with default headers values
func NewUpdateEdgeApplicationInstanceGatewayTimeout() *UpdateEdgeApplicationInstanceGatewayTimeout {
	return &UpdateEdgeApplicationInstanceGatewayTimeout{}
}

/* UpdateEdgeApplicationInstanceGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type UpdateEdgeApplicationInstanceGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateEdgeApplicationInstanceGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/instances/id/{id}][%d] updateEdgeApplicationInstanceGatewayTimeout  %+v", 504, o.Payload)
}
func (o *UpdateEdgeApplicationInstanceGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateEdgeApplicationInstanceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
