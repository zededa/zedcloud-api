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

// GetEdgeApplicationInstanceReader is a Reader for the GetEdgeApplicationInstance structure.
type GetEdgeApplicationInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEdgeApplicationInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEdgeApplicationInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEdgeApplicationInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEdgeApplicationInstanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEdgeApplicationInstanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEdgeApplicationInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetEdgeApplicationInstanceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEdgeApplicationInstanceOK creates a GetEdgeApplicationInstanceOK with default headers values
func NewGetEdgeApplicationInstanceOK() *GetEdgeApplicationInstanceOK {
	return &GetEdgeApplicationInstanceOK{}
}

/* GetEdgeApplicationInstanceOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetEdgeApplicationInstanceOK struct {
	Payload *swagger_models.AppInstance
}

func (o *GetEdgeApplicationInstanceOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}][%d] getEdgeApplicationInstanceOK  %+v", 200, o.Payload)
}
func (o *GetEdgeApplicationInstanceOK) GetPayload() *swagger_models.AppInstance {
	return o.Payload
}

func (o *GetEdgeApplicationInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.AppInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeApplicationInstanceUnauthorized creates a GetEdgeApplicationInstanceUnauthorized with default headers values
func NewGetEdgeApplicationInstanceUnauthorized() *GetEdgeApplicationInstanceUnauthorized {
	return &GetEdgeApplicationInstanceUnauthorized{}
}

/* GetEdgeApplicationInstanceUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetEdgeApplicationInstanceUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeApplicationInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}][%d] getEdgeApplicationInstanceUnauthorized  %+v", 401, o.Payload)
}
func (o *GetEdgeApplicationInstanceUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeApplicationInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeApplicationInstanceForbidden creates a GetEdgeApplicationInstanceForbidden with default headers values
func NewGetEdgeApplicationInstanceForbidden() *GetEdgeApplicationInstanceForbidden {
	return &GetEdgeApplicationInstanceForbidden{}
}

/* GetEdgeApplicationInstanceForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type GetEdgeApplicationInstanceForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeApplicationInstanceForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}][%d] getEdgeApplicationInstanceForbidden  %+v", 403, o.Payload)
}
func (o *GetEdgeApplicationInstanceForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeApplicationInstanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeApplicationInstanceNotFound creates a GetEdgeApplicationInstanceNotFound with default headers values
func NewGetEdgeApplicationInstanceNotFound() *GetEdgeApplicationInstanceNotFound {
	return &GetEdgeApplicationInstanceNotFound{}
}

/* GetEdgeApplicationInstanceNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetEdgeApplicationInstanceNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeApplicationInstanceNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}][%d] getEdgeApplicationInstanceNotFound  %+v", 404, o.Payload)
}
func (o *GetEdgeApplicationInstanceNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeApplicationInstanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeApplicationInstanceInternalServerError creates a GetEdgeApplicationInstanceInternalServerError with default headers values
func NewGetEdgeApplicationInstanceInternalServerError() *GetEdgeApplicationInstanceInternalServerError {
	return &GetEdgeApplicationInstanceInternalServerError{}
}

/* GetEdgeApplicationInstanceInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetEdgeApplicationInstanceInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeApplicationInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}][%d] getEdgeApplicationInstanceInternalServerError  %+v", 500, o.Payload)
}
func (o *GetEdgeApplicationInstanceInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeApplicationInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeApplicationInstanceGatewayTimeout creates a GetEdgeApplicationInstanceGatewayTimeout with default headers values
func NewGetEdgeApplicationInstanceGatewayTimeout() *GetEdgeApplicationInstanceGatewayTimeout {
	return &GetEdgeApplicationInstanceGatewayTimeout{}
}

/* GetEdgeApplicationInstanceGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetEdgeApplicationInstanceGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeApplicationInstanceGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}][%d] getEdgeApplicationInstanceGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetEdgeApplicationInstanceGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeApplicationInstanceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
