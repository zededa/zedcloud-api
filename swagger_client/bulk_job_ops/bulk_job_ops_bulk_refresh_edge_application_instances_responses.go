// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package bulk_job_ops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// BulkJobOpsBulkRefreshEdgeApplicationInstancesReader is a Reader for the BulkJobOpsBulkRefreshEdgeApplicationInstances structure.
type BulkJobOpsBulkRefreshEdgeApplicationInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBulkJobOpsBulkRefreshEdgeApplicationInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBulkJobOpsBulkRefreshEdgeApplicationInstancesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBulkJobOpsBulkRefreshEdgeApplicationInstancesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBulkJobOpsBulkRefreshEdgeApplicationInstancesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBulkJobOpsBulkRefreshEdgeApplicationInstancesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewBulkJobOpsBulkRefreshEdgeApplicationInstancesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBulkJobOpsBulkRefreshEdgeApplicationInstancesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewBulkJobOpsBulkRefreshEdgeApplicationInstancesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewBulkJobOpsBulkRefreshEdgeApplicationInstancesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBulkJobOpsBulkRefreshEdgeApplicationInstancesOK creates a BulkJobOpsBulkRefreshEdgeApplicationInstancesOK with default headers values
func NewBulkJobOpsBulkRefreshEdgeApplicationInstancesOK() *BulkJobOpsBulkRefreshEdgeApplicationInstancesOK {
	return &BulkJobOpsBulkRefreshEdgeApplicationInstancesOK{}
}

/* BulkJobOpsBulkRefreshEdgeApplicationInstancesOK describes a response with status code 200, with default header values.

A successful response.
*/
type BulkJobOpsBulkRefreshEdgeApplicationInstancesOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesOK) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh][%d] bulkJobOpsBulkRefreshEdgeApplicationInstancesOK  %+v", 200, o.Payload)
}
func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkRefreshEdgeApplicationInstancesBadRequest creates a BulkJobOpsBulkRefreshEdgeApplicationInstancesBadRequest with default headers values
func NewBulkJobOpsBulkRefreshEdgeApplicationInstancesBadRequest() *BulkJobOpsBulkRefreshEdgeApplicationInstancesBadRequest {
	return &BulkJobOpsBulkRefreshEdgeApplicationInstancesBadRequest{}
}

/* BulkJobOpsBulkRefreshEdgeApplicationInstancesBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type BulkJobOpsBulkRefreshEdgeApplicationInstancesBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh][%d] bulkJobOpsBulkRefreshEdgeApplicationInstancesBadRequest  %+v", 400, o.Payload)
}
func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkRefreshEdgeApplicationInstancesUnauthorized creates a BulkJobOpsBulkRefreshEdgeApplicationInstancesUnauthorized with default headers values
func NewBulkJobOpsBulkRefreshEdgeApplicationInstancesUnauthorized() *BulkJobOpsBulkRefreshEdgeApplicationInstancesUnauthorized {
	return &BulkJobOpsBulkRefreshEdgeApplicationInstancesUnauthorized{}
}

/* BulkJobOpsBulkRefreshEdgeApplicationInstancesUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type BulkJobOpsBulkRefreshEdgeApplicationInstancesUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh][%d] bulkJobOpsBulkRefreshEdgeApplicationInstancesUnauthorized  %+v", 401, o.Payload)
}
func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkRefreshEdgeApplicationInstancesForbidden creates a BulkJobOpsBulkRefreshEdgeApplicationInstancesForbidden with default headers values
func NewBulkJobOpsBulkRefreshEdgeApplicationInstancesForbidden() *BulkJobOpsBulkRefreshEdgeApplicationInstancesForbidden {
	return &BulkJobOpsBulkRefreshEdgeApplicationInstancesForbidden{}
}

/* BulkJobOpsBulkRefreshEdgeApplicationInstancesForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-app level access permission for the operation or does not have access scope to the project.
*/
type BulkJobOpsBulkRefreshEdgeApplicationInstancesForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh][%d] bulkJobOpsBulkRefreshEdgeApplicationInstancesForbidden  %+v", 403, o.Payload)
}
func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkRefreshEdgeApplicationInstancesNotFound creates a BulkJobOpsBulkRefreshEdgeApplicationInstancesNotFound with default headers values
func NewBulkJobOpsBulkRefreshEdgeApplicationInstancesNotFound() *BulkJobOpsBulkRefreshEdgeApplicationInstancesNotFound {
	return &BulkJobOpsBulkRefreshEdgeApplicationInstancesNotFound{}
}

/* BulkJobOpsBulkRefreshEdgeApplicationInstancesNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type BulkJobOpsBulkRefreshEdgeApplicationInstancesNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh][%d] bulkJobOpsBulkRefreshEdgeApplicationInstancesNotFound  %+v", 404, o.Payload)
}
func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkRefreshEdgeApplicationInstancesConflict creates a BulkJobOpsBulkRefreshEdgeApplicationInstancesConflict with default headers values
func NewBulkJobOpsBulkRefreshEdgeApplicationInstancesConflict() *BulkJobOpsBulkRefreshEdgeApplicationInstancesConflict {
	return &BulkJobOpsBulkRefreshEdgeApplicationInstancesConflict{}
}

/* BulkJobOpsBulkRefreshEdgeApplicationInstancesConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing asynchronous job request record.
*/
type BulkJobOpsBulkRefreshEdgeApplicationInstancesConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh][%d] bulkJobOpsBulkRefreshEdgeApplicationInstancesConflict  %+v", 409, o.Payload)
}
func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkRefreshEdgeApplicationInstancesInternalServerError creates a BulkJobOpsBulkRefreshEdgeApplicationInstancesInternalServerError with default headers values
func NewBulkJobOpsBulkRefreshEdgeApplicationInstancesInternalServerError() *BulkJobOpsBulkRefreshEdgeApplicationInstancesInternalServerError {
	return &BulkJobOpsBulkRefreshEdgeApplicationInstancesInternalServerError{}
}

/* BulkJobOpsBulkRefreshEdgeApplicationInstancesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type BulkJobOpsBulkRefreshEdgeApplicationInstancesInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh][%d] bulkJobOpsBulkRefreshEdgeApplicationInstancesInternalServerError  %+v", 500, o.Payload)
}
func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkRefreshEdgeApplicationInstancesGatewayTimeout creates a BulkJobOpsBulkRefreshEdgeApplicationInstancesGatewayTimeout with default headers values
func NewBulkJobOpsBulkRefreshEdgeApplicationInstancesGatewayTimeout() *BulkJobOpsBulkRefreshEdgeApplicationInstancesGatewayTimeout {
	return &BulkJobOpsBulkRefreshEdgeApplicationInstancesGatewayTimeout{}
}

/* BulkJobOpsBulkRefreshEdgeApplicationInstancesGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type BulkJobOpsBulkRefreshEdgeApplicationInstancesGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh][%d] bulkJobOpsBulkRefreshEdgeApplicationInstancesGatewayTimeout  %+v", 504, o.Payload)
}
func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkRefreshEdgeApplicationInstancesDefault creates a BulkJobOpsBulkRefreshEdgeApplicationInstancesDefault with default headers values
func NewBulkJobOpsBulkRefreshEdgeApplicationInstancesDefault(code int) *BulkJobOpsBulkRefreshEdgeApplicationInstancesDefault {
	return &BulkJobOpsBulkRefreshEdgeApplicationInstancesDefault{
		_statusCode: code,
	}
}

/* BulkJobOpsBulkRefreshEdgeApplicationInstancesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BulkJobOpsBulkRefreshEdgeApplicationInstancesDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the bulk job ops bulk refresh edge application instances default response
func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesDefault) Code() int {
	return o._statusCode
}

func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh][%d] BulkJobOps_BulkRefreshEdgeApplicationInstances default  %+v", o._statusCode, o.Payload)
}
func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *BulkJobOpsBulkRefreshEdgeApplicationInstancesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
