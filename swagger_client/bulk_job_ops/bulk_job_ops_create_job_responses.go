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

// BulkJobOpsCreateJobReader is a Reader for the BulkJobOpsCreateJob structure.
type BulkJobOpsCreateJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkJobOpsCreateJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBulkJobOpsCreateJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBulkJobOpsCreateJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBulkJobOpsCreateJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBulkJobOpsCreateJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewBulkJobOpsCreateJobConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBulkJobOpsCreateJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewBulkJobOpsCreateJobGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewBulkJobOpsCreateJobDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBulkJobOpsCreateJobOK creates a BulkJobOpsCreateJobOK with default headers values
func NewBulkJobOpsCreateJobOK() *BulkJobOpsCreateJobOK {
	return &BulkJobOpsCreateJobOK{}
}

/* BulkJobOpsCreateJobOK describes a response with status code 200, with default header values.

A successful response.
*/
type BulkJobOpsCreateJobOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsCreateJobOK) Error() string {
	return fmt.Sprintf("[POST /v1/jobs][%d] bulkJobOpsCreateJobOK  %+v", 200, o.Payload)
}
func (o *BulkJobOpsCreateJobOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsCreateJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsCreateJobBadRequest creates a BulkJobOpsCreateJobBadRequest with default headers values
func NewBulkJobOpsCreateJobBadRequest() *BulkJobOpsCreateJobBadRequest {
	return &BulkJobOpsCreateJobBadRequest{}
}

/* BulkJobOpsCreateJobBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type BulkJobOpsCreateJobBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsCreateJobBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/jobs][%d] bulkJobOpsCreateJobBadRequest  %+v", 400, o.Payload)
}
func (o *BulkJobOpsCreateJobBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsCreateJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsCreateJobUnauthorized creates a BulkJobOpsCreateJobUnauthorized with default headers values
func NewBulkJobOpsCreateJobUnauthorized() *BulkJobOpsCreateJobUnauthorized {
	return &BulkJobOpsCreateJobUnauthorized{}
}

/* BulkJobOpsCreateJobUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type BulkJobOpsCreateJobUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsCreateJobUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/jobs][%d] bulkJobOpsCreateJobUnauthorized  %+v", 401, o.Payload)
}
func (o *BulkJobOpsCreateJobUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsCreateJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsCreateJobForbidden creates a BulkJobOpsCreateJobForbidden with default headers values
func NewBulkJobOpsCreateJobForbidden() *BulkJobOpsCreateJobForbidden {
	return &BulkJobOpsCreateJobForbidden{}
}

/* BulkJobOpsCreateJobForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-app level access permission for the operation or does not have access scope to the project.
*/
type BulkJobOpsCreateJobForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsCreateJobForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/jobs][%d] bulkJobOpsCreateJobForbidden  %+v", 403, o.Payload)
}
func (o *BulkJobOpsCreateJobForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsCreateJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsCreateJobConflict creates a BulkJobOpsCreateJobConflict with default headers values
func NewBulkJobOpsCreateJobConflict() *BulkJobOpsCreateJobConflict {
	return &BulkJobOpsCreateJobConflict{}
}

/* BulkJobOpsCreateJobConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this job record will conflict with an already existing job record.
*/
type BulkJobOpsCreateJobConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsCreateJobConflict) Error() string {
	return fmt.Sprintf("[POST /v1/jobs][%d] bulkJobOpsCreateJobConflict  %+v", 409, o.Payload)
}
func (o *BulkJobOpsCreateJobConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsCreateJobConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsCreateJobInternalServerError creates a BulkJobOpsCreateJobInternalServerError with default headers values
func NewBulkJobOpsCreateJobInternalServerError() *BulkJobOpsCreateJobInternalServerError {
	return &BulkJobOpsCreateJobInternalServerError{}
}

/* BulkJobOpsCreateJobInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type BulkJobOpsCreateJobInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsCreateJobInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/jobs][%d] bulkJobOpsCreateJobInternalServerError  %+v", 500, o.Payload)
}
func (o *BulkJobOpsCreateJobInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsCreateJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsCreateJobGatewayTimeout creates a BulkJobOpsCreateJobGatewayTimeout with default headers values
func NewBulkJobOpsCreateJobGatewayTimeout() *BulkJobOpsCreateJobGatewayTimeout {
	return &BulkJobOpsCreateJobGatewayTimeout{}
}

/* BulkJobOpsCreateJobGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type BulkJobOpsCreateJobGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsCreateJobGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/jobs][%d] bulkJobOpsCreateJobGatewayTimeout  %+v", 504, o.Payload)
}
func (o *BulkJobOpsCreateJobGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsCreateJobGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsCreateJobDefault creates a BulkJobOpsCreateJobDefault with default headers values
func NewBulkJobOpsCreateJobDefault(code int) *BulkJobOpsCreateJobDefault {
	return &BulkJobOpsCreateJobDefault{
		_statusCode: code,
	}
}

/* BulkJobOpsCreateJobDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BulkJobOpsCreateJobDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the bulk job ops create job default response
func (o *BulkJobOpsCreateJobDefault) Code() int {
	return o._statusCode
}

func (o *BulkJobOpsCreateJobDefault) Error() string {
	return fmt.Sprintf("[POST /v1/jobs][%d] BulkJobOps_CreateJob default  %+v", o._statusCode, o.Payload)
}
func (o *BulkJobOpsCreateJobDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *BulkJobOpsCreateJobDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
