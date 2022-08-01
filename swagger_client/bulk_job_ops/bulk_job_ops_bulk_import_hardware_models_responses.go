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

// BulkJobOpsBulkImportHardwareModelsReader is a Reader for the BulkJobOpsBulkImportHardwareModels structure.
type BulkJobOpsBulkImportHardwareModelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkJobOpsBulkImportHardwareModelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBulkJobOpsBulkImportHardwareModelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBulkJobOpsBulkImportHardwareModelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBulkJobOpsBulkImportHardwareModelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBulkJobOpsBulkImportHardwareModelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBulkJobOpsBulkImportHardwareModelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewBulkJobOpsBulkImportHardwareModelsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBulkJobOpsBulkImportHardwareModelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewBulkJobOpsBulkImportHardwareModelsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewBulkJobOpsBulkImportHardwareModelsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBulkJobOpsBulkImportHardwareModelsOK creates a BulkJobOpsBulkImportHardwareModelsOK with default headers values
func NewBulkJobOpsBulkImportHardwareModelsOK() *BulkJobOpsBulkImportHardwareModelsOK {
	return &BulkJobOpsBulkImportHardwareModelsOK{}
}

/* BulkJobOpsBulkImportHardwareModelsOK describes a response with status code 200, with default header values.

A successful response.
*/
type BulkJobOpsBulkImportHardwareModelsOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsBulkImportHardwareModelsOK) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/models/import][%d] bulkJobOpsBulkImportHardwareModelsOK  %+v", 200, o.Payload)
}
func (o *BulkJobOpsBulkImportHardwareModelsOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkImportHardwareModelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkImportHardwareModelsBadRequest creates a BulkJobOpsBulkImportHardwareModelsBadRequest with default headers values
func NewBulkJobOpsBulkImportHardwareModelsBadRequest() *BulkJobOpsBulkImportHardwareModelsBadRequest {
	return &BulkJobOpsBulkImportHardwareModelsBadRequest{}
}

/* BulkJobOpsBulkImportHardwareModelsBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type BulkJobOpsBulkImportHardwareModelsBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsBulkImportHardwareModelsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/models/import][%d] bulkJobOpsBulkImportHardwareModelsBadRequest  %+v", 400, o.Payload)
}
func (o *BulkJobOpsBulkImportHardwareModelsBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkImportHardwareModelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkImportHardwareModelsUnauthorized creates a BulkJobOpsBulkImportHardwareModelsUnauthorized with default headers values
func NewBulkJobOpsBulkImportHardwareModelsUnauthorized() *BulkJobOpsBulkImportHardwareModelsUnauthorized {
	return &BulkJobOpsBulkImportHardwareModelsUnauthorized{}
}

/* BulkJobOpsBulkImportHardwareModelsUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type BulkJobOpsBulkImportHardwareModelsUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsBulkImportHardwareModelsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/models/import][%d] bulkJobOpsBulkImportHardwareModelsUnauthorized  %+v", 401, o.Payload)
}
func (o *BulkJobOpsBulkImportHardwareModelsUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkImportHardwareModelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkImportHardwareModelsForbidden creates a BulkJobOpsBulkImportHardwareModelsForbidden with default headers values
func NewBulkJobOpsBulkImportHardwareModelsForbidden() *BulkJobOpsBulkImportHardwareModelsForbidden {
	return &BulkJobOpsBulkImportHardwareModelsForbidden{}
}

/* BulkJobOpsBulkImportHardwareModelsForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-app level access permission for the operation or does not have access scope to the project.
*/
type BulkJobOpsBulkImportHardwareModelsForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsBulkImportHardwareModelsForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/models/import][%d] bulkJobOpsBulkImportHardwareModelsForbidden  %+v", 403, o.Payload)
}
func (o *BulkJobOpsBulkImportHardwareModelsForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkImportHardwareModelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkImportHardwareModelsNotFound creates a BulkJobOpsBulkImportHardwareModelsNotFound with default headers values
func NewBulkJobOpsBulkImportHardwareModelsNotFound() *BulkJobOpsBulkImportHardwareModelsNotFound {
	return &BulkJobOpsBulkImportHardwareModelsNotFound{}
}

/* BulkJobOpsBulkImportHardwareModelsNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type BulkJobOpsBulkImportHardwareModelsNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsBulkImportHardwareModelsNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/models/import][%d] bulkJobOpsBulkImportHardwareModelsNotFound  %+v", 404, o.Payload)
}
func (o *BulkJobOpsBulkImportHardwareModelsNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkImportHardwareModelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkImportHardwareModelsConflict creates a BulkJobOpsBulkImportHardwareModelsConflict with default headers values
func NewBulkJobOpsBulkImportHardwareModelsConflict() *BulkJobOpsBulkImportHardwareModelsConflict {
	return &BulkJobOpsBulkImportHardwareModelsConflict{}
}

/* BulkJobOpsBulkImportHardwareModelsConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing asynchronous job request record.
*/
type BulkJobOpsBulkImportHardwareModelsConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsBulkImportHardwareModelsConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/models/import][%d] bulkJobOpsBulkImportHardwareModelsConflict  %+v", 409, o.Payload)
}
func (o *BulkJobOpsBulkImportHardwareModelsConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkImportHardwareModelsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkImportHardwareModelsInternalServerError creates a BulkJobOpsBulkImportHardwareModelsInternalServerError with default headers values
func NewBulkJobOpsBulkImportHardwareModelsInternalServerError() *BulkJobOpsBulkImportHardwareModelsInternalServerError {
	return &BulkJobOpsBulkImportHardwareModelsInternalServerError{}
}

/* BulkJobOpsBulkImportHardwareModelsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type BulkJobOpsBulkImportHardwareModelsInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsBulkImportHardwareModelsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/models/import][%d] bulkJobOpsBulkImportHardwareModelsInternalServerError  %+v", 500, o.Payload)
}
func (o *BulkJobOpsBulkImportHardwareModelsInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkImportHardwareModelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkImportHardwareModelsGatewayTimeout creates a BulkJobOpsBulkImportHardwareModelsGatewayTimeout with default headers values
func NewBulkJobOpsBulkImportHardwareModelsGatewayTimeout() *BulkJobOpsBulkImportHardwareModelsGatewayTimeout {
	return &BulkJobOpsBulkImportHardwareModelsGatewayTimeout{}
}

/* BulkJobOpsBulkImportHardwareModelsGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type BulkJobOpsBulkImportHardwareModelsGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *BulkJobOpsBulkImportHardwareModelsGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/models/import][%d] bulkJobOpsBulkImportHardwareModelsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *BulkJobOpsBulkImportHardwareModelsGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkImportHardwareModelsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkImportHardwareModelsDefault creates a BulkJobOpsBulkImportHardwareModelsDefault with default headers values
func NewBulkJobOpsBulkImportHardwareModelsDefault(code int) *BulkJobOpsBulkImportHardwareModelsDefault {
	return &BulkJobOpsBulkImportHardwareModelsDefault{
		_statusCode: code,
	}
}

/* BulkJobOpsBulkImportHardwareModelsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BulkJobOpsBulkImportHardwareModelsDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the bulk job ops bulk import hardware models default response
func (o *BulkJobOpsBulkImportHardwareModelsDefault) Code() int {
	return o._statusCode
}

func (o *BulkJobOpsBulkImportHardwareModelsDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/models/import][%d] BulkJobOps_BulkImportHardwareModels default  %+v", o._statusCode, o.Payload)
}
func (o *BulkJobOpsBulkImportHardwareModelsDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *BulkJobOpsBulkImportHardwareModelsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}