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

// EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusReader is a Reader for the EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatus structure.
type EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusOK creates a EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusOK with default headers values
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusOK() *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusOK {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusOK{}
}

/*
EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusOK struct {
	Payload *swagger_models.AppInstStatusMsg
}

// IsSuccess returns true when this edge application instance status get edge application instance status o k response has a 2xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge application instance status get edge application instance status o k response has a 3xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance status get edge application instance status o k response has a 4xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance status get edge application instance status o k response has a 5xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance status get edge application instance status o k response a status code equal to that given
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}/status][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}/status][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusOK) GetPayload() *swagger_models.AppInstStatusMsg {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.AppInstStatusMsg)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusUnauthorized creates a EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusUnauthorized with default headers values
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusUnauthorized() *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusUnauthorized {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusUnauthorized{}
}

/*
EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance status get edge application instance status unauthorized response has a 2xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance status get edge application instance status unauthorized response has a 3xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance status get edge application instance status unauthorized response has a 4xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance status get edge application instance status unauthorized response has a 5xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance status get edge application instance status unauthorized response a status code equal to that given
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}/status][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}/status][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusForbidden creates a EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusForbidden with default headers values
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusForbidden() *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusForbidden {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusForbidden{}
}

/*
EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance status get edge application instance status forbidden response has a 2xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance status get edge application instance status forbidden response has a 3xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance status get edge application instance status forbidden response has a 4xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance status get edge application instance status forbidden response has a 5xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance status get edge application instance status forbidden response a status code equal to that given
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}/status][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}/status][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusNotFound creates a EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusNotFound with default headers values
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusNotFound() *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusNotFound {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusNotFound{}
}

/*
EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance status get edge application instance status not found response has a 2xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance status get edge application instance status not found response has a 3xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance status get edge application instance status not found response has a 4xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance status get edge application instance status not found response has a 5xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance status get edge application instance status not found response a status code equal to that given
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}/status][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusNotFound  %+v", 404, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusNotFound) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}/status][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusNotFound  %+v", 404, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusInternalServerError creates a EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusInternalServerError with default headers values
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusInternalServerError() *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusInternalServerError {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusInternalServerError{}
}

/*
EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance status get edge application instance status internal server error response has a 2xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance status get edge application instance status internal server error response has a 3xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance status get edge application instance status internal server error response has a 4xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance status get edge application instance status internal server error response has a 5xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application instance status get edge application instance status internal server error response a status code equal to that given
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}/status][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}/status][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusGatewayTimeout creates a EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusGatewayTimeout with default headers values
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusGatewayTimeout() *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusGatewayTimeout {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusGatewayTimeout{}
}

/*
EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance status get edge application instance status gateway timeout response has a 2xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance status get edge application instance status gateway timeout response has a 3xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance status get edge application instance status gateway timeout response has a 4xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance status get edge application instance status gateway timeout response has a 5xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application instance status get edge application instance status gateway timeout response a status code equal to that given
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}/status][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}/status][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusDefault creates a EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusDefault with default headers values
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusDefault(code int) *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusDefault {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusDefault{
		_statusCode: code,
	}
}

/*
EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the edge application instance status get edge application instance status default response
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge application instance status get edge application instance status default response has a 2xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge application instance status get edge application instance status default response has a 3xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge application instance status get edge application instance status default response has a 4xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge application instance status get edge application instance status default response has a 5xx status code
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge application instance status get edge application instance status default response a status code equal to that given
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusDefault) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}/status][%d] EdgeApplicationInstanceStatus_GetEdgeApplicationInstanceStatus default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusDefault) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{id}/status][%d] EdgeApplicationInstanceStatus_GetEdgeApplicationInstanceStatus default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
