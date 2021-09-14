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

// DeleteEdgeApplicationInstanceReader is a Reader for the DeleteEdgeApplicationInstance structure.
type DeleteEdgeApplicationInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEdgeApplicationInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteEdgeApplicationInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteEdgeApplicationInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteEdgeApplicationInstanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteEdgeApplicationInstanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteEdgeApplicationInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteEdgeApplicationInstanceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteEdgeApplicationInstanceOK creates a DeleteEdgeApplicationInstanceOK with default headers values
func NewDeleteEdgeApplicationInstanceOK() *DeleteEdgeApplicationInstanceOK {
	return &DeleteEdgeApplicationInstanceOK{}
}

/* DeleteEdgeApplicationInstanceOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteEdgeApplicationInstanceOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *DeleteEdgeApplicationInstanceOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/instances/id/{id}][%d] deleteEdgeApplicationInstanceOK  %+v", 200, o.Payload)
}
func (o *DeleteEdgeApplicationInstanceOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteEdgeApplicationInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEdgeApplicationInstanceUnauthorized creates a DeleteEdgeApplicationInstanceUnauthorized with default headers values
func NewDeleteEdgeApplicationInstanceUnauthorized() *DeleteEdgeApplicationInstanceUnauthorized {
	return &DeleteEdgeApplicationInstanceUnauthorized{}
}

/* DeleteEdgeApplicationInstanceUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type DeleteEdgeApplicationInstanceUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *DeleteEdgeApplicationInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/instances/id/{id}][%d] deleteEdgeApplicationInstanceUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteEdgeApplicationInstanceUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteEdgeApplicationInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEdgeApplicationInstanceForbidden creates a DeleteEdgeApplicationInstanceForbidden with default headers values
func NewDeleteEdgeApplicationInstanceForbidden() *DeleteEdgeApplicationInstanceForbidden {
	return &DeleteEdgeApplicationInstanceForbidden{}
}

/* DeleteEdgeApplicationInstanceForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type DeleteEdgeApplicationInstanceForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *DeleteEdgeApplicationInstanceForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/instances/id/{id}][%d] deleteEdgeApplicationInstanceForbidden  %+v", 403, o.Payload)
}
func (o *DeleteEdgeApplicationInstanceForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteEdgeApplicationInstanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEdgeApplicationInstanceNotFound creates a DeleteEdgeApplicationInstanceNotFound with default headers values
func NewDeleteEdgeApplicationInstanceNotFound() *DeleteEdgeApplicationInstanceNotFound {
	return &DeleteEdgeApplicationInstanceNotFound{}
}

/* DeleteEdgeApplicationInstanceNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type DeleteEdgeApplicationInstanceNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *DeleteEdgeApplicationInstanceNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/instances/id/{id}][%d] deleteEdgeApplicationInstanceNotFound  %+v", 404, o.Payload)
}
func (o *DeleteEdgeApplicationInstanceNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteEdgeApplicationInstanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEdgeApplicationInstanceInternalServerError creates a DeleteEdgeApplicationInstanceInternalServerError with default headers values
func NewDeleteEdgeApplicationInstanceInternalServerError() *DeleteEdgeApplicationInstanceInternalServerError {
	return &DeleteEdgeApplicationInstanceInternalServerError{}
}

/* DeleteEdgeApplicationInstanceInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type DeleteEdgeApplicationInstanceInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *DeleteEdgeApplicationInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/instances/id/{id}][%d] deleteEdgeApplicationInstanceInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteEdgeApplicationInstanceInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteEdgeApplicationInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEdgeApplicationInstanceGatewayTimeout creates a DeleteEdgeApplicationInstanceGatewayTimeout with default headers values
func NewDeleteEdgeApplicationInstanceGatewayTimeout() *DeleteEdgeApplicationInstanceGatewayTimeout {
	return &DeleteEdgeApplicationInstanceGatewayTimeout{}
}

/* DeleteEdgeApplicationInstanceGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type DeleteEdgeApplicationInstanceGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *DeleteEdgeApplicationInstanceGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/instances/id/{id}][%d] deleteEdgeApplicationInstanceGatewayTimeout  %+v", 504, o.Payload)
}
func (o *DeleteEdgeApplicationInstanceGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteEdgeApplicationInstanceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
