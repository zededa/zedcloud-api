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

// BulkJobOpsDeleteJobReader is a Reader for the BulkJobOpsDeleteJob structure.
type BulkJobOpsDeleteJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkJobOpsDeleteJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBulkJobOpsDeleteJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBulkJobOpsDeleteJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBulkJobOpsDeleteJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBulkJobOpsDeleteJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBulkJobOpsDeleteJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewBulkJobOpsDeleteJobConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBulkJobOpsDeleteJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewBulkJobOpsDeleteJobGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewBulkJobOpsDeleteJobDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBulkJobOpsDeleteJobOK creates a BulkJobOpsDeleteJobOK with default headers values
func NewBulkJobOpsDeleteJobOK() *BulkJobOpsDeleteJobOK {
	return &BulkJobOpsDeleteJobOK{}
}

/* BulkJobOpsDeleteJobOK describes a response with status code 200, with default header values.

A successful response.
*/
type BulkJobOpsDeleteJobOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsDeleteJobOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/jobs/id/{id}/objectType/{objectType}][%d] bulkJobOpsDeleteJobOK  %+v", 200, o.Payload)
}
func (o *BulkJobOpsDeleteJobOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsDeleteJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsDeleteJobBadRequest creates a BulkJobOpsDeleteJobBadRequest with default headers values
func NewBulkJobOpsDeleteJobBadRequest() *BulkJobOpsDeleteJobBadRequest {
	return &BulkJobOpsDeleteJobBadRequest{}
}

/* BulkJobOpsDeleteJobBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type BulkJobOpsDeleteJobBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsDeleteJobBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/jobs/id/{id}/objectType/{objectType}][%d] bulkJobOpsDeleteJobBadRequest  %+v", 400, o.Payload)
}
func (o *BulkJobOpsDeleteJobBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsDeleteJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsDeleteJobUnauthorized creates a BulkJobOpsDeleteJobUnauthorized with default headers values
func NewBulkJobOpsDeleteJobUnauthorized() *BulkJobOpsDeleteJobUnauthorized {
	return &BulkJobOpsDeleteJobUnauthorized{}
}

/* BulkJobOpsDeleteJobUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type BulkJobOpsDeleteJobUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsDeleteJobUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/jobs/id/{id}/objectType/{objectType}][%d] bulkJobOpsDeleteJobUnauthorized  %+v", 401, o.Payload)
}
func (o *BulkJobOpsDeleteJobUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsDeleteJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsDeleteJobForbidden creates a BulkJobOpsDeleteJobForbidden with default headers values
func NewBulkJobOpsDeleteJobForbidden() *BulkJobOpsDeleteJobForbidden {
	return &BulkJobOpsDeleteJobForbidden{}
}

/* BulkJobOpsDeleteJobForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-app level access permission for the operation or does not have access scope to the project.
*/
type BulkJobOpsDeleteJobForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsDeleteJobForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/jobs/id/{id}/objectType/{objectType}][%d] bulkJobOpsDeleteJobForbidden  %+v", 403, o.Payload)
}
func (o *BulkJobOpsDeleteJobForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsDeleteJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsDeleteJobNotFound creates a BulkJobOpsDeleteJobNotFound with default headers values
func NewBulkJobOpsDeleteJobNotFound() *BulkJobOpsDeleteJobNotFound {
	return &BulkJobOpsDeleteJobNotFound{}
}

/* BulkJobOpsDeleteJobNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type BulkJobOpsDeleteJobNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsDeleteJobNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/jobs/id/{id}/objectType/{objectType}][%d] bulkJobOpsDeleteJobNotFound  %+v", 404, o.Payload)
}
func (o *BulkJobOpsDeleteJobNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsDeleteJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsDeleteJobConflict creates a BulkJobOpsDeleteJobConflict with default headers values
func NewBulkJobOpsDeleteJobConflict() *BulkJobOpsDeleteJobConflict {
	return &BulkJobOpsDeleteJobConflict{}
}

/* BulkJobOpsDeleteJobConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this job is in progress, dependent tasks are still running
*/
type BulkJobOpsDeleteJobConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsDeleteJobConflict) Error() string {
	return fmt.Sprintf("[DELETE /v1/jobs/id/{id}/objectType/{objectType}][%d] bulkJobOpsDeleteJobConflict  %+v", 409, o.Payload)
}
func (o *BulkJobOpsDeleteJobConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsDeleteJobConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsDeleteJobInternalServerError creates a BulkJobOpsDeleteJobInternalServerError with default headers values
func NewBulkJobOpsDeleteJobInternalServerError() *BulkJobOpsDeleteJobInternalServerError {
	return &BulkJobOpsDeleteJobInternalServerError{}
}

/* BulkJobOpsDeleteJobInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type BulkJobOpsDeleteJobInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsDeleteJobInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/jobs/id/{id}/objectType/{objectType}][%d] bulkJobOpsDeleteJobInternalServerError  %+v", 500, o.Payload)
}
func (o *BulkJobOpsDeleteJobInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsDeleteJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsDeleteJobGatewayTimeout creates a BulkJobOpsDeleteJobGatewayTimeout with default headers values
func NewBulkJobOpsDeleteJobGatewayTimeout() *BulkJobOpsDeleteJobGatewayTimeout {
	return &BulkJobOpsDeleteJobGatewayTimeout{}
}

/* BulkJobOpsDeleteJobGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type BulkJobOpsDeleteJobGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsDeleteJobGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /v1/jobs/id/{id}/objectType/{objectType}][%d] bulkJobOpsDeleteJobGatewayTimeout  %+v", 504, o.Payload)
}
func (o *BulkJobOpsDeleteJobGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsDeleteJobGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsDeleteJobDefault creates a BulkJobOpsDeleteJobDefault with default headers values
func NewBulkJobOpsDeleteJobDefault(code int) *BulkJobOpsDeleteJobDefault {
	return &BulkJobOpsDeleteJobDefault{
		_statusCode: code,
	}
}

/* BulkJobOpsDeleteJobDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BulkJobOpsDeleteJobDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the bulk job ops delete job default response
func (o *BulkJobOpsDeleteJobDefault) Code() int {
	return o._statusCode
}

func (o *BulkJobOpsDeleteJobDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/jobs/id/{id}/objectType/{objectType}][%d] BulkJobOps_DeleteJob default  %+v", o._statusCode, o.Payload)
}
func (o *BulkJobOpsDeleteJobDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *BulkJobOpsDeleteJobDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
