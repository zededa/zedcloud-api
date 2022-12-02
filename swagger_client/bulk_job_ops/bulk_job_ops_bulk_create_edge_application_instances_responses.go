// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
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

// BulkJobOpsBulkCreateEdgeApplicationInstancesReader is a Reader for the BulkJobOpsBulkCreateEdgeApplicationInstances structure.
type BulkJobOpsBulkCreateEdgeApplicationInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBulkJobOpsBulkCreateEdgeApplicationInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBulkJobOpsBulkCreateEdgeApplicationInstancesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBulkJobOpsBulkCreateEdgeApplicationInstancesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBulkJobOpsBulkCreateEdgeApplicationInstancesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBulkJobOpsBulkCreateEdgeApplicationInstancesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewBulkJobOpsBulkCreateEdgeApplicationInstancesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBulkJobOpsBulkCreateEdgeApplicationInstancesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewBulkJobOpsBulkCreateEdgeApplicationInstancesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewBulkJobOpsBulkCreateEdgeApplicationInstancesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBulkJobOpsBulkCreateEdgeApplicationInstancesOK creates a BulkJobOpsBulkCreateEdgeApplicationInstancesOK with default headers values
func NewBulkJobOpsBulkCreateEdgeApplicationInstancesOK() *BulkJobOpsBulkCreateEdgeApplicationInstancesOK {
	return &BulkJobOpsBulkCreateEdgeApplicationInstancesOK{}
}

/*
BulkJobOpsBulkCreateEdgeApplicationInstancesOK describes a response with status code 200, with default header values.

A successful response.
*/
type BulkJobOpsBulkCreateEdgeApplicationInstancesOK struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this bulk job ops bulk create edge application instances o k response has a 2xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this bulk job ops bulk create edge application instances o k response has a 3xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk job ops bulk create edge application instances o k response has a 4xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this bulk job ops bulk create edge application instances o k response has a 5xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this bulk job ops bulk create edge application instances o k response a status code equal to that given
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesOK) IsCode(code int) bool {
	return code == 200
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesOK) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/create][%d] bulkJobOpsBulkCreateEdgeApplicationInstancesOK  %+v", 200, o.Payload)
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesOK) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/create][%d] bulkJobOpsBulkCreateEdgeApplicationInstancesOK  %+v", 200, o.Payload)
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkCreateEdgeApplicationInstancesBadRequest creates a BulkJobOpsBulkCreateEdgeApplicationInstancesBadRequest with default headers values
func NewBulkJobOpsBulkCreateEdgeApplicationInstancesBadRequest() *BulkJobOpsBulkCreateEdgeApplicationInstancesBadRequest {
	return &BulkJobOpsBulkCreateEdgeApplicationInstancesBadRequest{}
}

/*
BulkJobOpsBulkCreateEdgeApplicationInstancesBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type BulkJobOpsBulkCreateEdgeApplicationInstancesBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this bulk job ops bulk create edge application instances bad request response has a 2xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bulk job ops bulk create edge application instances bad request response has a 3xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk job ops bulk create edge application instances bad request response has a 4xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this bulk job ops bulk create edge application instances bad request response has a 5xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this bulk job ops bulk create edge application instances bad request response a status code equal to that given
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/create][%d] bulkJobOpsBulkCreateEdgeApplicationInstancesBadRequest  %+v", 400, o.Payload)
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/create][%d] bulkJobOpsBulkCreateEdgeApplicationInstancesBadRequest  %+v", 400, o.Payload)
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkCreateEdgeApplicationInstancesUnauthorized creates a BulkJobOpsBulkCreateEdgeApplicationInstancesUnauthorized with default headers values
func NewBulkJobOpsBulkCreateEdgeApplicationInstancesUnauthorized() *BulkJobOpsBulkCreateEdgeApplicationInstancesUnauthorized {
	return &BulkJobOpsBulkCreateEdgeApplicationInstancesUnauthorized{}
}

/*
BulkJobOpsBulkCreateEdgeApplicationInstancesUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type BulkJobOpsBulkCreateEdgeApplicationInstancesUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this bulk job ops bulk create edge application instances unauthorized response has a 2xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bulk job ops bulk create edge application instances unauthorized response has a 3xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk job ops bulk create edge application instances unauthorized response has a 4xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this bulk job ops bulk create edge application instances unauthorized response has a 5xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this bulk job ops bulk create edge application instances unauthorized response a status code equal to that given
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/create][%d] bulkJobOpsBulkCreateEdgeApplicationInstancesUnauthorized  %+v", 401, o.Payload)
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/create][%d] bulkJobOpsBulkCreateEdgeApplicationInstancesUnauthorized  %+v", 401, o.Payload)
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkCreateEdgeApplicationInstancesForbidden creates a BulkJobOpsBulkCreateEdgeApplicationInstancesForbidden with default headers values
func NewBulkJobOpsBulkCreateEdgeApplicationInstancesForbidden() *BulkJobOpsBulkCreateEdgeApplicationInstancesForbidden {
	return &BulkJobOpsBulkCreateEdgeApplicationInstancesForbidden{}
}

/*
BulkJobOpsBulkCreateEdgeApplicationInstancesForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-app level access permission for the operation or does not have access scope to the project.
*/
type BulkJobOpsBulkCreateEdgeApplicationInstancesForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this bulk job ops bulk create edge application instances forbidden response has a 2xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bulk job ops bulk create edge application instances forbidden response has a 3xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk job ops bulk create edge application instances forbidden response has a 4xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this bulk job ops bulk create edge application instances forbidden response has a 5xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this bulk job ops bulk create edge application instances forbidden response a status code equal to that given
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/create][%d] bulkJobOpsBulkCreateEdgeApplicationInstancesForbidden  %+v", 403, o.Payload)
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/create][%d] bulkJobOpsBulkCreateEdgeApplicationInstancesForbidden  %+v", 403, o.Payload)
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkCreateEdgeApplicationInstancesNotFound creates a BulkJobOpsBulkCreateEdgeApplicationInstancesNotFound with default headers values
func NewBulkJobOpsBulkCreateEdgeApplicationInstancesNotFound() *BulkJobOpsBulkCreateEdgeApplicationInstancesNotFound {
	return &BulkJobOpsBulkCreateEdgeApplicationInstancesNotFound{}
}

/*
BulkJobOpsBulkCreateEdgeApplicationInstancesNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type BulkJobOpsBulkCreateEdgeApplicationInstancesNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this bulk job ops bulk create edge application instances not found response has a 2xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bulk job ops bulk create edge application instances not found response has a 3xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk job ops bulk create edge application instances not found response has a 4xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this bulk job ops bulk create edge application instances not found response has a 5xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this bulk job ops bulk create edge application instances not found response a status code equal to that given
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/create][%d] bulkJobOpsBulkCreateEdgeApplicationInstancesNotFound  %+v", 404, o.Payload)
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/create][%d] bulkJobOpsBulkCreateEdgeApplicationInstancesNotFound  %+v", 404, o.Payload)
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkCreateEdgeApplicationInstancesConflict creates a BulkJobOpsBulkCreateEdgeApplicationInstancesConflict with default headers values
func NewBulkJobOpsBulkCreateEdgeApplicationInstancesConflict() *BulkJobOpsBulkCreateEdgeApplicationInstancesConflict {
	return &BulkJobOpsBulkCreateEdgeApplicationInstancesConflict{}
}

/*
BulkJobOpsBulkCreateEdgeApplicationInstancesConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing asynchronous job request record.
*/
type BulkJobOpsBulkCreateEdgeApplicationInstancesConflict struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this bulk job ops bulk create edge application instances conflict response has a 2xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bulk job ops bulk create edge application instances conflict response has a 3xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk job ops bulk create edge application instances conflict response has a 4xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this bulk job ops bulk create edge application instances conflict response has a 5xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this bulk job ops bulk create edge application instances conflict response a status code equal to that given
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesConflict) IsCode(code int) bool {
	return code == 409
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/create][%d] bulkJobOpsBulkCreateEdgeApplicationInstancesConflict  %+v", 409, o.Payload)
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesConflict) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/create][%d] bulkJobOpsBulkCreateEdgeApplicationInstancesConflict  %+v", 409, o.Payload)
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkCreateEdgeApplicationInstancesInternalServerError creates a BulkJobOpsBulkCreateEdgeApplicationInstancesInternalServerError with default headers values
func NewBulkJobOpsBulkCreateEdgeApplicationInstancesInternalServerError() *BulkJobOpsBulkCreateEdgeApplicationInstancesInternalServerError {
	return &BulkJobOpsBulkCreateEdgeApplicationInstancesInternalServerError{}
}

/*
BulkJobOpsBulkCreateEdgeApplicationInstancesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type BulkJobOpsBulkCreateEdgeApplicationInstancesInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this bulk job ops bulk create edge application instances internal server error response has a 2xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bulk job ops bulk create edge application instances internal server error response has a 3xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk job ops bulk create edge application instances internal server error response has a 4xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this bulk job ops bulk create edge application instances internal server error response has a 5xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this bulk job ops bulk create edge application instances internal server error response a status code equal to that given
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/create][%d] bulkJobOpsBulkCreateEdgeApplicationInstancesInternalServerError  %+v", 500, o.Payload)
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/create][%d] bulkJobOpsBulkCreateEdgeApplicationInstancesInternalServerError  %+v", 500, o.Payload)
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkCreateEdgeApplicationInstancesGatewayTimeout creates a BulkJobOpsBulkCreateEdgeApplicationInstancesGatewayTimeout with default headers values
func NewBulkJobOpsBulkCreateEdgeApplicationInstancesGatewayTimeout() *BulkJobOpsBulkCreateEdgeApplicationInstancesGatewayTimeout {
	return &BulkJobOpsBulkCreateEdgeApplicationInstancesGatewayTimeout{}
}

/*
BulkJobOpsBulkCreateEdgeApplicationInstancesGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type BulkJobOpsBulkCreateEdgeApplicationInstancesGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this bulk job ops bulk create edge application instances gateway timeout response has a 2xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bulk job ops bulk create edge application instances gateway timeout response has a 3xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk job ops bulk create edge application instances gateway timeout response has a 4xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this bulk job ops bulk create edge application instances gateway timeout response has a 5xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this bulk job ops bulk create edge application instances gateway timeout response a status code equal to that given
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/create][%d] bulkJobOpsBulkCreateEdgeApplicationInstancesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/create][%d] bulkJobOpsBulkCreateEdgeApplicationInstancesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkCreateEdgeApplicationInstancesDefault creates a BulkJobOpsBulkCreateEdgeApplicationInstancesDefault with default headers values
func NewBulkJobOpsBulkCreateEdgeApplicationInstancesDefault(code int) *BulkJobOpsBulkCreateEdgeApplicationInstancesDefault {
	return &BulkJobOpsBulkCreateEdgeApplicationInstancesDefault{
		_statusCode: code,
	}
}

/*
BulkJobOpsBulkCreateEdgeApplicationInstancesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BulkJobOpsBulkCreateEdgeApplicationInstancesDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the bulk job ops bulk create edge application instances default response
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this bulk job ops bulk create edge application instances default response has a 2xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this bulk job ops bulk create edge application instances default response has a 3xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this bulk job ops bulk create edge application instances default response has a 4xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this bulk job ops bulk create edge application instances default response has a 5xx status code
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this bulk job ops bulk create edge application instances default response a status code equal to that given
func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/create][%d] BulkJobOps_BulkCreateEdgeApplicationInstances default  %+v", o._statusCode, o.Payload)
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesDefault) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/create][%d] BulkJobOps_BulkCreateEdgeApplicationInstances default  %+v", o._statusCode, o.Payload)
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *BulkJobOpsBulkCreateEdgeApplicationInstancesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
