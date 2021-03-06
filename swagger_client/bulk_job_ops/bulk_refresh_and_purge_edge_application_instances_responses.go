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

// BulkRefreshAndPurgeEdgeApplicationInstancesReader is a Reader for the BulkRefreshAndPurgeEdgeApplicationInstances structure.
type BulkRefreshAndPurgeEdgeApplicationInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkRefreshAndPurgeEdgeApplicationInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBulkRefreshAndPurgeEdgeApplicationInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBulkRefreshAndPurgeEdgeApplicationInstancesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBulkRefreshAndPurgeEdgeApplicationInstancesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBulkRefreshAndPurgeEdgeApplicationInstancesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewBulkRefreshAndPurgeEdgeApplicationInstancesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewBulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBulkRefreshAndPurgeEdgeApplicationInstancesOK creates a BulkRefreshAndPurgeEdgeApplicationInstancesOK with default headers values
func NewBulkRefreshAndPurgeEdgeApplicationInstancesOK() *BulkRefreshAndPurgeEdgeApplicationInstancesOK {
	return &BulkRefreshAndPurgeEdgeApplicationInstancesOK{}
}

/* BulkRefreshAndPurgeEdgeApplicationInstancesOK describes a response with status code 200, with default header values.

A successful response.
*/
type BulkRefreshAndPurgeEdgeApplicationInstancesOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkRefreshAndPurgeEdgeApplicationInstancesOK) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh/purge][%d] bulkRefreshAndPurgeEdgeApplicationInstancesOK  %+v", 200, o.Payload)
}
func (o *BulkRefreshAndPurgeEdgeApplicationInstancesOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkRefreshAndPurgeEdgeApplicationInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkRefreshAndPurgeEdgeApplicationInstancesBadRequest creates a BulkRefreshAndPurgeEdgeApplicationInstancesBadRequest with default headers values
func NewBulkRefreshAndPurgeEdgeApplicationInstancesBadRequest() *BulkRefreshAndPurgeEdgeApplicationInstancesBadRequest {
	return &BulkRefreshAndPurgeEdgeApplicationInstancesBadRequest{}
}

/* BulkRefreshAndPurgeEdgeApplicationInstancesBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type BulkRefreshAndPurgeEdgeApplicationInstancesBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkRefreshAndPurgeEdgeApplicationInstancesBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh/purge][%d] bulkRefreshAndPurgeEdgeApplicationInstancesBadRequest  %+v", 400, o.Payload)
}
func (o *BulkRefreshAndPurgeEdgeApplicationInstancesBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkRefreshAndPurgeEdgeApplicationInstancesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized creates a BulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized with default headers values
func NewBulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized() *BulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized {
	return &BulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized{}
}

/* BulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type BulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh/purge][%d] bulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized  %+v", 401, o.Payload)
}
func (o *BulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkRefreshAndPurgeEdgeApplicationInstancesForbidden creates a BulkRefreshAndPurgeEdgeApplicationInstancesForbidden with default headers values
func NewBulkRefreshAndPurgeEdgeApplicationInstancesForbidden() *BulkRefreshAndPurgeEdgeApplicationInstancesForbidden {
	return &BulkRefreshAndPurgeEdgeApplicationInstancesForbidden{}
}

/* BulkRefreshAndPurgeEdgeApplicationInstancesForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-app level access permission for the operation or does not have access scope to the project.
*/
type BulkRefreshAndPurgeEdgeApplicationInstancesForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkRefreshAndPurgeEdgeApplicationInstancesForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh/purge][%d] bulkRefreshAndPurgeEdgeApplicationInstancesForbidden  %+v", 403, o.Payload)
}
func (o *BulkRefreshAndPurgeEdgeApplicationInstancesForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkRefreshAndPurgeEdgeApplicationInstancesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkRefreshAndPurgeEdgeApplicationInstancesNotFound creates a BulkRefreshAndPurgeEdgeApplicationInstancesNotFound with default headers values
func NewBulkRefreshAndPurgeEdgeApplicationInstancesNotFound() *BulkRefreshAndPurgeEdgeApplicationInstancesNotFound {
	return &BulkRefreshAndPurgeEdgeApplicationInstancesNotFound{}
}

/* BulkRefreshAndPurgeEdgeApplicationInstancesNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type BulkRefreshAndPurgeEdgeApplicationInstancesNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkRefreshAndPurgeEdgeApplicationInstancesNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh/purge][%d] bulkRefreshAndPurgeEdgeApplicationInstancesNotFound  %+v", 404, o.Payload)
}
func (o *BulkRefreshAndPurgeEdgeApplicationInstancesNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkRefreshAndPurgeEdgeApplicationInstancesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkRefreshAndPurgeEdgeApplicationInstancesConflict creates a BulkRefreshAndPurgeEdgeApplicationInstancesConflict with default headers values
func NewBulkRefreshAndPurgeEdgeApplicationInstancesConflict() *BulkRefreshAndPurgeEdgeApplicationInstancesConflict {
	return &BulkRefreshAndPurgeEdgeApplicationInstancesConflict{}
}

/* BulkRefreshAndPurgeEdgeApplicationInstancesConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing asynchronous job request record.
*/
type BulkRefreshAndPurgeEdgeApplicationInstancesConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkRefreshAndPurgeEdgeApplicationInstancesConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh/purge][%d] bulkRefreshAndPurgeEdgeApplicationInstancesConflict  %+v", 409, o.Payload)
}
func (o *BulkRefreshAndPurgeEdgeApplicationInstancesConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkRefreshAndPurgeEdgeApplicationInstancesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError creates a BulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError with default headers values
func NewBulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError() *BulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError {
	return &BulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError{}
}

/* BulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type BulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh/purge][%d] bulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError  %+v", 500, o.Payload)
}
func (o *BulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout creates a BulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout with default headers values
func NewBulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout() *BulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout {
	return &BulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout{}
}

/* BulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type BulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh/purge][%d] bulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout  %+v", 504, o.Payload)
}
func (o *BulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
