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

// BulkJobOpsBulkImportEdgeApplicationsReader is a Reader for the BulkJobOpsBulkImportEdgeApplications structure.
type BulkJobOpsBulkImportEdgeApplicationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkJobOpsBulkImportEdgeApplicationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBulkJobOpsBulkImportEdgeApplicationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBulkJobOpsBulkImportEdgeApplicationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBulkJobOpsBulkImportEdgeApplicationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBulkJobOpsBulkImportEdgeApplicationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBulkJobOpsBulkImportEdgeApplicationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewBulkJobOpsBulkImportEdgeApplicationsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBulkJobOpsBulkImportEdgeApplicationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewBulkJobOpsBulkImportEdgeApplicationsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewBulkJobOpsBulkImportEdgeApplicationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBulkJobOpsBulkImportEdgeApplicationsOK creates a BulkJobOpsBulkImportEdgeApplicationsOK with default headers values
func NewBulkJobOpsBulkImportEdgeApplicationsOK() *BulkJobOpsBulkImportEdgeApplicationsOK {
	return &BulkJobOpsBulkImportEdgeApplicationsOK{}
}

/* BulkJobOpsBulkImportEdgeApplicationsOK describes a response with status code 200, with default header values.

A successful response.
*/
type BulkJobOpsBulkImportEdgeApplicationsOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsBulkImportEdgeApplicationsOK) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/bundles/import][%d] bulkJobOpsBulkImportEdgeApplicationsOK  %+v", 200, o.Payload)
}
func (o *BulkJobOpsBulkImportEdgeApplicationsOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkImportEdgeApplicationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkImportEdgeApplicationsBadRequest creates a BulkJobOpsBulkImportEdgeApplicationsBadRequest with default headers values
func NewBulkJobOpsBulkImportEdgeApplicationsBadRequest() *BulkJobOpsBulkImportEdgeApplicationsBadRequest {
	return &BulkJobOpsBulkImportEdgeApplicationsBadRequest{}
}

/* BulkJobOpsBulkImportEdgeApplicationsBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type BulkJobOpsBulkImportEdgeApplicationsBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsBulkImportEdgeApplicationsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/bundles/import][%d] bulkJobOpsBulkImportEdgeApplicationsBadRequest  %+v", 400, o.Payload)
}
func (o *BulkJobOpsBulkImportEdgeApplicationsBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkImportEdgeApplicationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkImportEdgeApplicationsUnauthorized creates a BulkJobOpsBulkImportEdgeApplicationsUnauthorized with default headers values
func NewBulkJobOpsBulkImportEdgeApplicationsUnauthorized() *BulkJobOpsBulkImportEdgeApplicationsUnauthorized {
	return &BulkJobOpsBulkImportEdgeApplicationsUnauthorized{}
}

/* BulkJobOpsBulkImportEdgeApplicationsUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type BulkJobOpsBulkImportEdgeApplicationsUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsBulkImportEdgeApplicationsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/bundles/import][%d] bulkJobOpsBulkImportEdgeApplicationsUnauthorized  %+v", 401, o.Payload)
}
func (o *BulkJobOpsBulkImportEdgeApplicationsUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkImportEdgeApplicationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkImportEdgeApplicationsForbidden creates a BulkJobOpsBulkImportEdgeApplicationsForbidden with default headers values
func NewBulkJobOpsBulkImportEdgeApplicationsForbidden() *BulkJobOpsBulkImportEdgeApplicationsForbidden {
	return &BulkJobOpsBulkImportEdgeApplicationsForbidden{}
}

/* BulkJobOpsBulkImportEdgeApplicationsForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-app level access permission for the operation or does not have access scope to the project.
*/
type BulkJobOpsBulkImportEdgeApplicationsForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsBulkImportEdgeApplicationsForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/bundles/import][%d] bulkJobOpsBulkImportEdgeApplicationsForbidden  %+v", 403, o.Payload)
}
func (o *BulkJobOpsBulkImportEdgeApplicationsForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkImportEdgeApplicationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkImportEdgeApplicationsNotFound creates a BulkJobOpsBulkImportEdgeApplicationsNotFound with default headers values
func NewBulkJobOpsBulkImportEdgeApplicationsNotFound() *BulkJobOpsBulkImportEdgeApplicationsNotFound {
	return &BulkJobOpsBulkImportEdgeApplicationsNotFound{}
}

/* BulkJobOpsBulkImportEdgeApplicationsNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type BulkJobOpsBulkImportEdgeApplicationsNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsBulkImportEdgeApplicationsNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/bundles/import][%d] bulkJobOpsBulkImportEdgeApplicationsNotFound  %+v", 404, o.Payload)
}
func (o *BulkJobOpsBulkImportEdgeApplicationsNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkImportEdgeApplicationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkImportEdgeApplicationsConflict creates a BulkJobOpsBulkImportEdgeApplicationsConflict with default headers values
func NewBulkJobOpsBulkImportEdgeApplicationsConflict() *BulkJobOpsBulkImportEdgeApplicationsConflict {
	return &BulkJobOpsBulkImportEdgeApplicationsConflict{}
}

/* BulkJobOpsBulkImportEdgeApplicationsConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing asynchronous job request record.
*/
type BulkJobOpsBulkImportEdgeApplicationsConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsBulkImportEdgeApplicationsConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/bundles/import][%d] bulkJobOpsBulkImportEdgeApplicationsConflict  %+v", 409, o.Payload)
}
func (o *BulkJobOpsBulkImportEdgeApplicationsConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkImportEdgeApplicationsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkImportEdgeApplicationsInternalServerError creates a BulkJobOpsBulkImportEdgeApplicationsInternalServerError with default headers values
func NewBulkJobOpsBulkImportEdgeApplicationsInternalServerError() *BulkJobOpsBulkImportEdgeApplicationsInternalServerError {
	return &BulkJobOpsBulkImportEdgeApplicationsInternalServerError{}
}

/* BulkJobOpsBulkImportEdgeApplicationsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type BulkJobOpsBulkImportEdgeApplicationsInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsBulkImportEdgeApplicationsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/bundles/import][%d] bulkJobOpsBulkImportEdgeApplicationsInternalServerError  %+v", 500, o.Payload)
}
func (o *BulkJobOpsBulkImportEdgeApplicationsInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkImportEdgeApplicationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkImportEdgeApplicationsGatewayTimeout creates a BulkJobOpsBulkImportEdgeApplicationsGatewayTimeout with default headers values
func NewBulkJobOpsBulkImportEdgeApplicationsGatewayTimeout() *BulkJobOpsBulkImportEdgeApplicationsGatewayTimeout {
	return &BulkJobOpsBulkImportEdgeApplicationsGatewayTimeout{}
}

/* BulkJobOpsBulkImportEdgeApplicationsGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type BulkJobOpsBulkImportEdgeApplicationsGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsBulkImportEdgeApplicationsGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/bundles/import][%d] bulkJobOpsBulkImportEdgeApplicationsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *BulkJobOpsBulkImportEdgeApplicationsGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkImportEdgeApplicationsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkImportEdgeApplicationsDefault creates a BulkJobOpsBulkImportEdgeApplicationsDefault with default headers values
func NewBulkJobOpsBulkImportEdgeApplicationsDefault(code int) *BulkJobOpsBulkImportEdgeApplicationsDefault {
	return &BulkJobOpsBulkImportEdgeApplicationsDefault{
		_statusCode: code,
	}
}

/* BulkJobOpsBulkImportEdgeApplicationsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BulkJobOpsBulkImportEdgeApplicationsDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the bulk job ops bulk import edge applications default response
func (o *BulkJobOpsBulkImportEdgeApplicationsDefault) Code() int {
	return o._statusCode
}

func (o *BulkJobOpsBulkImportEdgeApplicationsDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/bundles/import][%d] BulkJobOps_BulkImportEdgeApplications default  %+v", o._statusCode, o.Payload)
}
func (o *BulkJobOpsBulkImportEdgeApplicationsDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *BulkJobOpsBulkImportEdgeApplicationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}