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

// EdgeNodeStatusGetEdgeNodeEventsReader is a Reader for the EdgeNodeStatusGetEdgeNodeEvents structure.
type EdgeNodeStatusGetEdgeNodeEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeStatusGetEdgeNodeEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeStatusGetEdgeNodeEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeStatusGetEdgeNodeEventsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeStatusGetEdgeNodeEventsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeStatusGetEdgeNodeEventsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeStatusGetEdgeNodeEventsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeStatusGetEdgeNodeEventsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeStatusGetEdgeNodeEventsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeStatusGetEdgeNodeEventsOK creates a EdgeNodeStatusGetEdgeNodeEventsOK with default headers values
func NewEdgeNodeStatusGetEdgeNodeEventsOK() *EdgeNodeStatusGetEdgeNodeEventsOK {
	return &EdgeNodeStatusGetEdgeNodeEventsOK{}
}

/*
EdgeNodeStatusGetEdgeNodeEventsOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeStatusGetEdgeNodeEventsOK struct {
	Payload *swagger_models.EventQueryResponse
}

// IsSuccess returns true when this edge node status get edge node events o k response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node status get edge node events o k response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node events o k response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node status get edge node events o k response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node events o k response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeEventsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge node status get edge node events o k response
func (o *EdgeNodeStatusGetEdgeNodeEventsOK) Code() int {
	return 200
}

func (o *EdgeNodeStatusGetEdgeNodeEventsOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/events][%d] edgeNodeStatusGetEdgeNodeEventsOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEventsOK) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/events][%d] edgeNodeStatusGetEdgeNodeEventsOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEventsOK) GetPayload() *swagger_models.EventQueryResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.EventQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeEventsUnauthorized creates a EdgeNodeStatusGetEdgeNodeEventsUnauthorized with default headers values
func NewEdgeNodeStatusGetEdgeNodeEventsUnauthorized() *EdgeNodeStatusGetEdgeNodeEventsUnauthorized {
	return &EdgeNodeStatusGetEdgeNodeEventsUnauthorized{}
}

/*
EdgeNodeStatusGetEdgeNodeEventsUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeStatusGetEdgeNodeEventsUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node events unauthorized response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node events unauthorized response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node events unauthorized response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node status get edge node events unauthorized response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node events unauthorized response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeEventsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge node status get edge node events unauthorized response
func (o *EdgeNodeStatusGetEdgeNodeEventsUnauthorized) Code() int {
	return 401
}

func (o *EdgeNodeStatusGetEdgeNodeEventsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/events][%d] edgeNodeStatusGetEdgeNodeEventsUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEventsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/events][%d] edgeNodeStatusGetEdgeNodeEventsUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEventsUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeEventsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeEventsForbidden creates a EdgeNodeStatusGetEdgeNodeEventsForbidden with default headers values
func NewEdgeNodeStatusGetEdgeNodeEventsForbidden() *EdgeNodeStatusGetEdgeNodeEventsForbidden {
	return &EdgeNodeStatusGetEdgeNodeEventsForbidden{}
}

/*
EdgeNodeStatusGetEdgeNodeEventsForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeStatusGetEdgeNodeEventsForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node events forbidden response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node events forbidden response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node events forbidden response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node status get edge node events forbidden response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node events forbidden response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeEventsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge node status get edge node events forbidden response
func (o *EdgeNodeStatusGetEdgeNodeEventsForbidden) Code() int {
	return 403
}

func (o *EdgeNodeStatusGetEdgeNodeEventsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/events][%d] edgeNodeStatusGetEdgeNodeEventsForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEventsForbidden) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/events][%d] edgeNodeStatusGetEdgeNodeEventsForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEventsForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeEventsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeEventsNotFound creates a EdgeNodeStatusGetEdgeNodeEventsNotFound with default headers values
func NewEdgeNodeStatusGetEdgeNodeEventsNotFound() *EdgeNodeStatusGetEdgeNodeEventsNotFound {
	return &EdgeNodeStatusGetEdgeNodeEventsNotFound{}
}

/*
EdgeNodeStatusGetEdgeNodeEventsNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeStatusGetEdgeNodeEventsNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node events not found response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node events not found response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node events not found response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node status get edge node events not found response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node status get edge node events not found response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeEventsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the edge node status get edge node events not found response
func (o *EdgeNodeStatusGetEdgeNodeEventsNotFound) Code() int {
	return 404
}

func (o *EdgeNodeStatusGetEdgeNodeEventsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/events][%d] edgeNodeStatusGetEdgeNodeEventsNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEventsNotFound) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/events][%d] edgeNodeStatusGetEdgeNodeEventsNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEventsNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeEventsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeEventsInternalServerError creates a EdgeNodeStatusGetEdgeNodeEventsInternalServerError with default headers values
func NewEdgeNodeStatusGetEdgeNodeEventsInternalServerError() *EdgeNodeStatusGetEdgeNodeEventsInternalServerError {
	return &EdgeNodeStatusGetEdgeNodeEventsInternalServerError{}
}

/*
EdgeNodeStatusGetEdgeNodeEventsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeStatusGetEdgeNodeEventsInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node events internal server error response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node events internal server error response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node events internal server error response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node status get edge node events internal server error response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node status get edge node events internal server error response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeEventsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge node status get edge node events internal server error response
func (o *EdgeNodeStatusGetEdgeNodeEventsInternalServerError) Code() int {
	return 500
}

func (o *EdgeNodeStatusGetEdgeNodeEventsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/events][%d] edgeNodeStatusGetEdgeNodeEventsInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEventsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/events][%d] edgeNodeStatusGetEdgeNodeEventsInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEventsInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeEventsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeEventsGatewayTimeout creates a EdgeNodeStatusGetEdgeNodeEventsGatewayTimeout with default headers values
func NewEdgeNodeStatusGetEdgeNodeEventsGatewayTimeout() *EdgeNodeStatusGetEdgeNodeEventsGatewayTimeout {
	return &EdgeNodeStatusGetEdgeNodeEventsGatewayTimeout{}
}

/*
EdgeNodeStatusGetEdgeNodeEventsGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeStatusGetEdgeNodeEventsGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge node status get edge node events gateway timeout response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node status get edge node events gateway timeout response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node status get edge node events gateway timeout response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node status get edge node events gateway timeout response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node status get edge node events gateway timeout response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeEventsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge node status get edge node events gateway timeout response
func (o *EdgeNodeStatusGetEdgeNodeEventsGatewayTimeout) Code() int {
	return 504
}

func (o *EdgeNodeStatusGetEdgeNodeEventsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/events][%d] edgeNodeStatusGetEdgeNodeEventsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEventsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/events][%d] edgeNodeStatusGetEdgeNodeEventsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEventsGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeEventsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeStatusGetEdgeNodeEventsDefault creates a EdgeNodeStatusGetEdgeNodeEventsDefault with default headers values
func NewEdgeNodeStatusGetEdgeNodeEventsDefault(code int) *EdgeNodeStatusGetEdgeNodeEventsDefault {
	return &EdgeNodeStatusGetEdgeNodeEventsDefault{
		_statusCode: code,
	}
}

/*
EdgeNodeStatusGetEdgeNodeEventsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeStatusGetEdgeNodeEventsDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// IsSuccess returns true when this edge node status get edge node events default response has a 2xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node status get edge node events default response has a 3xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node status get edge node events default response has a 4xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node status get edge node events default response has a 5xx status code
func (o *EdgeNodeStatusGetEdgeNodeEventsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node status get edge node events default response a status code equal to that given
func (o *EdgeNodeStatusGetEdgeNodeEventsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge node status get edge node events default response
func (o *EdgeNodeStatusGetEdgeNodeEventsDefault) Code() int {
	return o._statusCode
}

func (o *EdgeNodeStatusGetEdgeNodeEventsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/events][%d] EdgeNodeStatus_GetEdgeNodeEvents default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEventsDefault) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{objid}/events][%d] EdgeNodeStatus_GetEdgeNodeEvents default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeStatusGetEdgeNodeEventsDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeStatusGetEdgeNodeEventsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
