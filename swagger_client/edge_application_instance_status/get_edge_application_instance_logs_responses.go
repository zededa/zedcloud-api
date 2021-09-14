// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_application_instance_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// GetEdgeApplicationInstanceLogsReader is a Reader for the GetEdgeApplicationInstanceLogs structure.
type GetEdgeApplicationInstanceLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEdgeApplicationInstanceLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEdgeApplicationInstanceLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEdgeApplicationInstanceLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEdgeApplicationInstanceLogsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEdgeApplicationInstanceLogsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEdgeApplicationInstanceLogsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetEdgeApplicationInstanceLogsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEdgeApplicationInstanceLogsOK creates a GetEdgeApplicationInstanceLogsOK with default headers values
func NewGetEdgeApplicationInstanceLogsOK() *GetEdgeApplicationInstanceLogsOK {
	return &GetEdgeApplicationInstanceLogsOK{}
}

/* GetEdgeApplicationInstanceLogsOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetEdgeApplicationInstanceLogsOK struct {
	Payload *swagger_models.AppInstanceLogsResponse
}

func (o *GetEdgeApplicationInstanceLogsOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}/logs][%d] getEdgeApplicationInstanceLogsOK  %+v", 200, o.Payload)
}
func (o *GetEdgeApplicationInstanceLogsOK) GetPayload() *swagger_models.AppInstanceLogsResponse {
	return o.Payload
}

func (o *GetEdgeApplicationInstanceLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.AppInstanceLogsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeApplicationInstanceLogsUnauthorized creates a GetEdgeApplicationInstanceLogsUnauthorized with default headers values
func NewGetEdgeApplicationInstanceLogsUnauthorized() *GetEdgeApplicationInstanceLogsUnauthorized {
	return &GetEdgeApplicationInstanceLogsUnauthorized{}
}

/* GetEdgeApplicationInstanceLogsUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetEdgeApplicationInstanceLogsUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeApplicationInstanceLogsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}/logs][%d] getEdgeApplicationInstanceLogsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetEdgeApplicationInstanceLogsUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeApplicationInstanceLogsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeApplicationInstanceLogsForbidden creates a GetEdgeApplicationInstanceLogsForbidden with default headers values
func NewGetEdgeApplicationInstanceLogsForbidden() *GetEdgeApplicationInstanceLogsForbidden {
	return &GetEdgeApplicationInstanceLogsForbidden{}
}

/* GetEdgeApplicationInstanceLogsForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type GetEdgeApplicationInstanceLogsForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeApplicationInstanceLogsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}/logs][%d] getEdgeApplicationInstanceLogsForbidden  %+v", 403, o.Payload)
}
func (o *GetEdgeApplicationInstanceLogsForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeApplicationInstanceLogsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeApplicationInstanceLogsNotFound creates a GetEdgeApplicationInstanceLogsNotFound with default headers values
func NewGetEdgeApplicationInstanceLogsNotFound() *GetEdgeApplicationInstanceLogsNotFound {
	return &GetEdgeApplicationInstanceLogsNotFound{}
}

/* GetEdgeApplicationInstanceLogsNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetEdgeApplicationInstanceLogsNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeApplicationInstanceLogsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}/logs][%d] getEdgeApplicationInstanceLogsNotFound  %+v", 404, o.Payload)
}
func (o *GetEdgeApplicationInstanceLogsNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeApplicationInstanceLogsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeApplicationInstanceLogsInternalServerError creates a GetEdgeApplicationInstanceLogsInternalServerError with default headers values
func NewGetEdgeApplicationInstanceLogsInternalServerError() *GetEdgeApplicationInstanceLogsInternalServerError {
	return &GetEdgeApplicationInstanceLogsInternalServerError{}
}

/* GetEdgeApplicationInstanceLogsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetEdgeApplicationInstanceLogsInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeApplicationInstanceLogsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}/logs][%d] getEdgeApplicationInstanceLogsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetEdgeApplicationInstanceLogsInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeApplicationInstanceLogsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeApplicationInstanceLogsGatewayTimeout creates a GetEdgeApplicationInstanceLogsGatewayTimeout with default headers values
func NewGetEdgeApplicationInstanceLogsGatewayTimeout() *GetEdgeApplicationInstanceLogsGatewayTimeout {
	return &GetEdgeApplicationInstanceLogsGatewayTimeout{}
}

/* GetEdgeApplicationInstanceLogsGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetEdgeApplicationInstanceLogsGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeApplicationInstanceLogsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}/logs][%d] getEdgeApplicationInstanceLogsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetEdgeApplicationInstanceLogsGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeApplicationInstanceLogsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
