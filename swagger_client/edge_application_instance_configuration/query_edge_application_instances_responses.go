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

// QueryEdgeApplicationInstancesReader is a Reader for the QueryEdgeApplicationInstances structure.
type QueryEdgeApplicationInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryEdgeApplicationInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryEdgeApplicationInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryEdgeApplicationInstancesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewQueryEdgeApplicationInstancesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryEdgeApplicationInstancesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryEdgeApplicationInstancesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewQueryEdgeApplicationInstancesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryEdgeApplicationInstancesOK creates a QueryEdgeApplicationInstancesOK with default headers values
func NewQueryEdgeApplicationInstancesOK() *QueryEdgeApplicationInstancesOK {
	return &QueryEdgeApplicationInstancesOK{}
}

/* QueryEdgeApplicationInstancesOK describes a response with status code 200, with default header values.

A successful response.
*/
type QueryEdgeApplicationInstancesOK struct {
	Payload *swagger_models.AppInstances
}

func (o *QueryEdgeApplicationInstancesOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] queryEdgeApplicationInstancesOK  %+v", 200, o.Payload)
}
func (o *QueryEdgeApplicationInstancesOK) GetPayload() *swagger_models.AppInstances {
	return o.Payload
}

func (o *QueryEdgeApplicationInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.AppInstances)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryEdgeApplicationInstancesBadRequest creates a QueryEdgeApplicationInstancesBadRequest with default headers values
func NewQueryEdgeApplicationInstancesBadRequest() *QueryEdgeApplicationInstancesBadRequest {
	return &QueryEdgeApplicationInstancesBadRequest{}
}

/* QueryEdgeApplicationInstancesBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type QueryEdgeApplicationInstancesBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryEdgeApplicationInstancesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] queryEdgeApplicationInstancesBadRequest  %+v", 400, o.Payload)
}
func (o *QueryEdgeApplicationInstancesBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryEdgeApplicationInstancesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryEdgeApplicationInstancesUnauthorized creates a QueryEdgeApplicationInstancesUnauthorized with default headers values
func NewQueryEdgeApplicationInstancesUnauthorized() *QueryEdgeApplicationInstancesUnauthorized {
	return &QueryEdgeApplicationInstancesUnauthorized{}
}

/* QueryEdgeApplicationInstancesUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type QueryEdgeApplicationInstancesUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryEdgeApplicationInstancesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] queryEdgeApplicationInstancesUnauthorized  %+v", 401, o.Payload)
}
func (o *QueryEdgeApplicationInstancesUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryEdgeApplicationInstancesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryEdgeApplicationInstancesForbidden creates a QueryEdgeApplicationInstancesForbidden with default headers values
func NewQueryEdgeApplicationInstancesForbidden() *QueryEdgeApplicationInstancesForbidden {
	return &QueryEdgeApplicationInstancesForbidden{}
}

/* QueryEdgeApplicationInstancesForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type QueryEdgeApplicationInstancesForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryEdgeApplicationInstancesForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] queryEdgeApplicationInstancesForbidden  %+v", 403, o.Payload)
}
func (o *QueryEdgeApplicationInstancesForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryEdgeApplicationInstancesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryEdgeApplicationInstancesInternalServerError creates a QueryEdgeApplicationInstancesInternalServerError with default headers values
func NewQueryEdgeApplicationInstancesInternalServerError() *QueryEdgeApplicationInstancesInternalServerError {
	return &QueryEdgeApplicationInstancesInternalServerError{}
}

/* QueryEdgeApplicationInstancesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type QueryEdgeApplicationInstancesInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryEdgeApplicationInstancesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] queryEdgeApplicationInstancesInternalServerError  %+v", 500, o.Payload)
}
func (o *QueryEdgeApplicationInstancesInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryEdgeApplicationInstancesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryEdgeApplicationInstancesGatewayTimeout creates a QueryEdgeApplicationInstancesGatewayTimeout with default headers values
func NewQueryEdgeApplicationInstancesGatewayTimeout() *QueryEdgeApplicationInstancesGatewayTimeout {
	return &QueryEdgeApplicationInstancesGatewayTimeout{}
}

/* QueryEdgeApplicationInstancesGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type QueryEdgeApplicationInstancesGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryEdgeApplicationInstancesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances][%d] queryEdgeApplicationInstancesGatewayTimeout  %+v", 504, o.Payload)
}
func (o *QueryEdgeApplicationInstancesGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryEdgeApplicationInstancesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
