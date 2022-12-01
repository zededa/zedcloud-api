// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package edge_node_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// EdgeNodeStatusGetEdgeNodeEdgeviewStatusReader is a Reader for the EdgeNodeStatusGetEdgeNodeEdgeviewStatus structure.
type EdgeNodeStatusGetEdgeNodeEdgeviewStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusOK creates a EdgeNodeStatusGetEdgeNodeEdgeviewStatusOK with default headers values
func NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusOK() *EdgeNodeStatusGetEdgeNodeEdgeviewStatusOK {
	return &EdgeNodeStatusGetEdgeNodeEdgeviewStatusOK{}
}

/*
EdgeNodeStatusGetEdgeNodeEdgeviewStatusOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeStatusGetEdgeNodeEdgeviewStatusOK struct {
	Payload *swagger_models.DeviceRawMetrics
}

// IsSuccess returns true when this edge node status get edge node edgeview status o k response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node status get edge node edgeview status o k response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node edgeview status o k response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node status get edge node edgeview status o k response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node edgeview status o k response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/edgeview][%d] edgeNodeStatusGetEdgeNodeEdgeviewStatusOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusOK) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/edgeview][%d] edgeNodeStatusGetEdgeNodeEdgeviewStatusOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusOK) GetPayload() *swagger_models.DeviceRawMetrics {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.DeviceRawMetrics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusUnauthorized creates a EdgeNodeStatusGetEdgeNodeEdgeviewStatusUnauthorized with default headers values
func NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusUnauthorized() *EdgeNodeStatusGetEdgeNodeEdgeviewStatusUnauthorized {
	return &EdgeNodeStatusGetEdgeNodeEdgeviewStatusUnauthorized{}
}

/*
EdgeNodeStatusGetEdgeNodeEdgeviewStatusUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeStatusGetEdgeNodeEdgeviewStatusUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node edgeview status unauthorized response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node edgeview status unauthorized response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node edgeview status unauthorized response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node status get edge node edgeview status unauthorized response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node edgeview status unauthorized response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/edgeview][%d] edgeNodeStatusGetEdgeNodeEdgeviewStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/edgeview][%d] edgeNodeStatusGetEdgeNodeEdgeviewStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusForbidden creates a EdgeNodeStatusGetEdgeNodeEdgeviewStatusForbidden with default headers values
func NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusForbidden() *EdgeNodeStatusGetEdgeNodeEdgeviewStatusForbidden {
	return &EdgeNodeStatusGetEdgeNodeEdgeviewStatusForbidden{}
}

/*
EdgeNodeStatusGetEdgeNodeEdgeviewStatusForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeStatusGetEdgeNodeEdgeviewStatusForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node edgeview status forbidden response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node edgeview status forbidden response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node edgeview status forbidden response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node status get edge node edgeview status forbidden response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node edgeview status forbidden response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/edgeview][%d] edgeNodeStatusGetEdgeNodeEdgeviewStatusForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusForbidden) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/edgeview][%d] edgeNodeStatusGetEdgeNodeEdgeviewStatusForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusNotFound creates a EdgeNodeStatusGetEdgeNodeEdgeviewStatusNotFound with default headers values
func NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusNotFound() *EdgeNodeStatusGetEdgeNodeEdgeviewStatusNotFound {
	return &EdgeNodeStatusGetEdgeNodeEdgeviewStatusNotFound{}
}

/*
EdgeNodeStatusGetEdgeNodeEdgeviewStatusNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeStatusGetEdgeNodeEdgeviewStatusNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node edgeview status not found response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node edgeview status not found response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node edgeview status not found response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node status get edge node edgeview status not found response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node edgeview status not found response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/edgeview][%d] edgeNodeStatusGetEdgeNodeEdgeviewStatusNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusNotFound) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/edgeview][%d] edgeNodeStatusGetEdgeNodeEdgeviewStatusNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusInternalServerError creates a EdgeNodeStatusGetEdgeNodeEdgeviewStatusInternalServerError with default headers values
func NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusInternalServerError() *EdgeNodeStatusGetEdgeNodeEdgeviewStatusInternalServerError {
	return &EdgeNodeStatusGetEdgeNodeEdgeviewStatusInternalServerError{}
}

/*
EdgeNodeStatusGetEdgeNodeEdgeviewStatusInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeStatusGetEdgeNodeEdgeviewStatusInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node edgeview status internal server error response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node edgeview status internal server error response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node edgeview status internal server error response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node status get edge node edgeview status internal server error response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node status get edge node edgeview status internal server error response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/edgeview][%d] edgeNodeStatusGetEdgeNodeEdgeviewStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/edgeview][%d] edgeNodeStatusGetEdgeNodeEdgeviewStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusGatewayTimeout creates a EdgeNodeStatusGetEdgeNodeEdgeviewStatusGatewayTimeout with default headers values
func NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusGatewayTimeout() *EdgeNodeStatusGetEdgeNodeEdgeviewStatusGatewayTimeout {
	return &EdgeNodeStatusGetEdgeNodeEdgeviewStatusGatewayTimeout{}
}

/*
EdgeNodeStatusGetEdgeNodeEdgeviewStatusGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeStatusGetEdgeNodeEdgeviewStatusGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node edgeview status gateway timeout response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node edgeview status gateway timeout response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node edgeview status gateway timeout response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node status get edge node edgeview status gateway timeout response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node status get edge node edgeview status gateway timeout response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/edgeview][%d] edgeNodeStatusGetEdgeNodeEdgeviewStatusGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/edgeview][%d] edgeNodeStatusGetEdgeNodeEdgeviewStatusGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusDefault creates a EdgeNodeStatusGetEdgeNodeEdgeviewStatusDefault with default headers values
func NewEdgeNodeStatusGetEdgeNodeEdgeviewStatusDefault(code int) *EdgeNodeStatusGetEdgeNodeEdgeviewStatusDefault {
	return &EdgeNodeStatusGetEdgeNodeEdgeviewStatusDefault{
		_statusCode: code,
	}
}

/*
EdgeNodeStatusGetEdgeNodeEdgeviewStatusDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeStatusGetEdgeNodeEdgeviewStatusDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the edge node status get edge node edgeview status default response
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge node status get edge node edgeview status default response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node status get edge node edgeview status default response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node status get edge node edgeview status default response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node status get edge node edgeview status default response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node status get edge node edgeview status default response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusDefault) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/edgeview][%d] EdgeNodeStatus_GetEdgeNodeEdgeviewStatus default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusDefault) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/status/edgeview][%d] EdgeNodeStatus_GetEdgeNodeEdgeviewStatus default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeEdgeviewStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}