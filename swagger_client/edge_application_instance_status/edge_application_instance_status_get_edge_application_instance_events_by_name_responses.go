// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
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

// EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameReader is a Reader for the EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByName structure.
type EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameOK creates a EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameOK with default headers values
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameOK() *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameOK {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameOK{}
}

/*
EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameOK struct {
	Payload *swagger_models.EventQueryResponse
}

// IsSuccess returns true when this edge application instance status get edge application instance events by name o k response has a 2xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge application instance status get edge application instance events by name o k response has a 3xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance status get edge application instance events by name o k response has a 4xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance status get edge application instance events by name o k response has a 5xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance status get edge application instance events by name o k response a status code equal to that given
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{objname}/events][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{objname}/events][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameOK) GetPayload() *swagger_models.EventQueryResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.EventQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameUnauthorized creates a EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameUnauthorized with default headers values
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameUnauthorized() *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameUnauthorized {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameUnauthorized{}
}

/*
EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance status get edge application instance events by name unauthorized response has a 2xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance status get edge application instance events by name unauthorized response has a 3xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance status get edge application instance events by name unauthorized response has a 4xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance status get edge application instance events by name unauthorized response has a 5xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance status get edge application instance events by name unauthorized response a status code equal to that given
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{objname}/events][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{objname}/events][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameForbidden creates a EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameForbidden with default headers values
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameForbidden() *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameForbidden {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameForbidden{}
}

/*
EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance status get edge application instance events by name forbidden response has a 2xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance status get edge application instance events by name forbidden response has a 3xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance status get edge application instance events by name forbidden response has a 4xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance status get edge application instance events by name forbidden response has a 5xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance status get edge application instance events by name forbidden response a status code equal to that given
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{objname}/events][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{objname}/events][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameNotFound creates a EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameNotFound with default headers values
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameNotFound() *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameNotFound {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameNotFound{}
}

/*
EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance status get edge application instance events by name not found response has a 2xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance status get edge application instance events by name not found response has a 3xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance status get edge application instance events by name not found response has a 4xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance status get edge application instance events by name not found response has a 5xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance status get edge application instance events by name not found response a status code equal to that given
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{objname}/events][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameNotFound  %+v", 404, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameNotFound) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{objname}/events][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameNotFound  %+v", 404, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameInternalServerError creates a EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameInternalServerError with default headers values
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameInternalServerError() *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameInternalServerError {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameInternalServerError{}
}

/*
EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance status get edge application instance events by name internal server error response has a 2xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance status get edge application instance events by name internal server error response has a 3xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance status get edge application instance events by name internal server error response has a 4xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance status get edge application instance events by name internal server error response has a 5xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application instance status get edge application instance events by name internal server error response a status code equal to that given
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{objname}/events][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{objname}/events][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameGatewayTimeout creates a EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameGatewayTimeout with default headers values
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameGatewayTimeout() *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameGatewayTimeout {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameGatewayTimeout{}
}

/*
EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance status get edge application instance events by name gateway timeout response has a 2xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance status get edge application instance events by name gateway timeout response has a 3xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance status get edge application instance events by name gateway timeout response has a 4xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance status get edge application instance events by name gateway timeout response has a 5xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application instance status get edge application instance events by name gateway timeout response a status code equal to that given
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{objname}/events][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{objname}/events][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameDefault creates a EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameDefault with default headers values
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameDefault(code int) *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameDefault {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameDefault{
		_statusCode: code,
	}
}

/*
EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the edge application instance status get edge application instance events by name default response
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge application instance status get edge application instance events by name default response has a 2xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge application instance status get edge application instance events by name default response has a 3xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge application instance status get edge application instance events by name default response has a 4xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge application instance status get edge application instance events by name default response has a 5xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge application instance status get edge application instance events by name default response a status code equal to that given
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{objname}/events][%d] EdgeApplicationInstanceStatus_GetEdgeApplicationInstanceEventsByName default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameDefault) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{objname}/events][%d] EdgeApplicationInstanceStatus_GetEdgeApplicationInstanceEventsByName default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceEventsByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
