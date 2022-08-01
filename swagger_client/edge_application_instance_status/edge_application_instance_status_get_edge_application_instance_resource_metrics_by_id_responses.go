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

// EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDReader is a Reader for the EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByID structure.
type EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDOK creates a EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDOK with default headers values
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDOK() *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDOK {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDOK{}
}

/* EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDOK struct {
	Payload *swagger_models.MetricQueryResponse
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{objid}/timeSeries/{mType}][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIdOK  %+v", 200, o.Payload)
}
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDOK) GetPayload() *swagger_models.MetricQueryResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.MetricQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDUnauthorized creates a EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDUnauthorized with default headers values
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDUnauthorized() *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDUnauthorized {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDUnauthorized{}
}

/* EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{objid}/timeSeries/{mType}][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIdUnauthorized  %+v", 401, o.Payload)
}
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDForbidden creates a EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDForbidden with default headers values
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDForbidden() *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDForbidden {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDForbidden{}
}

/* EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{objid}/timeSeries/{mType}][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIdForbidden  %+v", 403, o.Payload)
}
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDNotFound creates a EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDNotFound with default headers values
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDNotFound() *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDNotFound {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDNotFound{}
}

/* EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{objid}/timeSeries/{mType}][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIdNotFound  %+v", 404, o.Payload)
}
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDInternalServerError creates a EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDInternalServerError with default headers values
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDInternalServerError() *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDInternalServerError {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDInternalServerError{}
}

/* EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{objid}/timeSeries/{mType}][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIdInternalServerError  %+v", 500, o.Payload)
}
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDGatewayTimeout creates a EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDGatewayTimeout with default headers values
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDGatewayTimeout() *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDGatewayTimeout {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDGatewayTimeout{}
}

/* EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{objid}/timeSeries/{mType}][%d] edgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIdGatewayTimeout  %+v", 504, o.Payload)
}
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDDefault creates a EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDDefault with default headers values
func NewEdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDDefault(code int) *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDDefault {
	return &EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDDefault{
		_statusCode: code,
	}
}

/* EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the edge application instance status get edge application instance resource metrics by Id default response
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDDefault) Code() int {
	return o._statusCode
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDDefault) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/id/{objid}/timeSeries/{mType}][%d] EdgeApplicationInstanceStatus_GetEdgeApplicationInstanceResourceMetricsById default  %+v", o._statusCode, o.Payload)
}
func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeApplicationInstanceStatusGetEdgeApplicationInstanceResourceMetricsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
