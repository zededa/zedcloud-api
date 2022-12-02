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

// BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesReader is a Reader for the BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstances structure.
type BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesOK creates a BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesOK with default headers values
func NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesOK() *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesOK {
	return &BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesOK{}
}

/*
BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesOK describes a response with status code 200, with default header values.

A successful response.
*/
type BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesOK struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this bulk job ops bulk refresh and purge edge application instances o k response has a 2xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this bulk job ops bulk refresh and purge edge application instances o k response has a 3xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk job ops bulk refresh and purge edge application instances o k response has a 4xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this bulk job ops bulk refresh and purge edge application instances o k response has a 5xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this bulk job ops bulk refresh and purge edge application instances o k response a status code equal to that given
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesOK) IsCode(code int) bool {
	return code == 200
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesOK) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh/purge][%d] bulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesOK  %+v", 200, o.Payload)
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesOK) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh/purge][%d] bulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesOK  %+v", 200, o.Payload)
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesBadRequest creates a BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesBadRequest with default headers values
func NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesBadRequest() *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesBadRequest {
	return &BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesBadRequest{}
}

/*
BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this bulk job ops bulk refresh and purge edge application instances bad request response has a 2xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bulk job ops bulk refresh and purge edge application instances bad request response has a 3xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk job ops bulk refresh and purge edge application instances bad request response has a 4xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this bulk job ops bulk refresh and purge edge application instances bad request response has a 5xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this bulk job ops bulk refresh and purge edge application instances bad request response a status code equal to that given
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh/purge][%d] bulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesBadRequest  %+v", 400, o.Payload)
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh/purge][%d] bulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesBadRequest  %+v", 400, o.Payload)
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized creates a BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized with default headers values
func NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized() *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized {
	return &BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized{}
}

/*
BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this bulk job ops bulk refresh and purge edge application instances unauthorized response has a 2xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bulk job ops bulk refresh and purge edge application instances unauthorized response has a 3xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk job ops bulk refresh and purge edge application instances unauthorized response has a 4xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this bulk job ops bulk refresh and purge edge application instances unauthorized response has a 5xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this bulk job ops bulk refresh and purge edge application instances unauthorized response a status code equal to that given
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh/purge][%d] bulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized  %+v", 401, o.Payload)
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh/purge][%d] bulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized  %+v", 401, o.Payload)
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesForbidden creates a BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesForbidden with default headers values
func NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesForbidden() *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesForbidden {
	return &BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesForbidden{}
}

/*
BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-app level access permission for the operation or does not have access scope to the project.
*/
type BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this bulk job ops bulk refresh and purge edge application instances forbidden response has a 2xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bulk job ops bulk refresh and purge edge application instances forbidden response has a 3xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk job ops bulk refresh and purge edge application instances forbidden response has a 4xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this bulk job ops bulk refresh and purge edge application instances forbidden response has a 5xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this bulk job ops bulk refresh and purge edge application instances forbidden response a status code equal to that given
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh/purge][%d] bulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesForbidden  %+v", 403, o.Payload)
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh/purge][%d] bulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesForbidden  %+v", 403, o.Payload)
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesNotFound creates a BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesNotFound with default headers values
func NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesNotFound() *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesNotFound {
	return &BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesNotFound{}
}

/*
BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this bulk job ops bulk refresh and purge edge application instances not found response has a 2xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bulk job ops bulk refresh and purge edge application instances not found response has a 3xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk job ops bulk refresh and purge edge application instances not found response has a 4xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this bulk job ops bulk refresh and purge edge application instances not found response has a 5xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this bulk job ops bulk refresh and purge edge application instances not found response a status code equal to that given
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh/purge][%d] bulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesNotFound  %+v", 404, o.Payload)
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh/purge][%d] bulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesNotFound  %+v", 404, o.Payload)
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesConflict creates a BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesConflict with default headers values
func NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesConflict() *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesConflict {
	return &BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesConflict{}
}

/*
BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing asynchronous job request record.
*/
type BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesConflict struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this bulk job ops bulk refresh and purge edge application instances conflict response has a 2xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bulk job ops bulk refresh and purge edge application instances conflict response has a 3xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk job ops bulk refresh and purge edge application instances conflict response has a 4xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this bulk job ops bulk refresh and purge edge application instances conflict response has a 5xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this bulk job ops bulk refresh and purge edge application instances conflict response a status code equal to that given
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesConflict) IsCode(code int) bool {
	return code == 409
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh/purge][%d] bulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesConflict  %+v", 409, o.Payload)
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesConflict) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh/purge][%d] bulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesConflict  %+v", 409, o.Payload)
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError creates a BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError with default headers values
func NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError() *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError {
	return &BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError{}
}

/*
BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this bulk job ops bulk refresh and purge edge application instances internal server error response has a 2xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bulk job ops bulk refresh and purge edge application instances internal server error response has a 3xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk job ops bulk refresh and purge edge application instances internal server error response has a 4xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this bulk job ops bulk refresh and purge edge application instances internal server error response has a 5xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this bulk job ops bulk refresh and purge edge application instances internal server error response a status code equal to that given
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh/purge][%d] bulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError  %+v", 500, o.Payload)
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh/purge][%d] bulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError  %+v", 500, o.Payload)
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout creates a BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout with default headers values
func NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout() *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout {
	return &BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout{}
}

/*
BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this bulk job ops bulk refresh and purge edge application instances gateway timeout response has a 2xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bulk job ops bulk refresh and purge edge application instances gateway timeout response has a 3xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk job ops bulk refresh and purge edge application instances gateway timeout response has a 4xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this bulk job ops bulk refresh and purge edge application instances gateway timeout response has a 5xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this bulk job ops bulk refresh and purge edge application instances gateway timeout response a status code equal to that given
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh/purge][%d] bulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh/purge][%d] bulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesDefault creates a BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesDefault with default headers values
func NewBulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesDefault(code int) *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesDefault {
	return &BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesDefault{
		_statusCode: code,
	}
}

/*
BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the bulk job ops bulk refresh and purge edge application instances default response
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this bulk job ops bulk refresh and purge edge application instances default response has a 2xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this bulk job ops bulk refresh and purge edge application instances default response has a 3xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this bulk job ops bulk refresh and purge edge application instances default response has a 4xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this bulk job ops bulk refresh and purge edge application instances default response has a 5xx status code
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this bulk job ops bulk refresh and purge edge application instances default response a status code equal to that given
func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh/purge][%d] BulkJobOps_BulkRefreshAndPurgeEdgeApplicationInstances default  %+v", o._statusCode, o.Payload)
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesDefault) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/apps/instances/refresh/purge][%d] BulkJobOps_BulkRefreshAndPurgeEdgeApplicationInstances default  %+v", o._statusCode, o.Payload)
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *BulkJobOpsBulkRefreshAndPurgeEdgeApplicationInstancesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
