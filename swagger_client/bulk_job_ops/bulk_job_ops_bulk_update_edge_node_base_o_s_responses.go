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

// BulkJobOpsBulkUpdateEdgeNodeBaseOSReader is a Reader for the BulkJobOpsBulkUpdateEdgeNodeBaseOS structure.
type BulkJobOpsBulkUpdateEdgeNodeBaseOSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBulkJobOpsBulkUpdateEdgeNodeBaseOSOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBulkJobOpsBulkUpdateEdgeNodeBaseOSBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBulkJobOpsBulkUpdateEdgeNodeBaseOSUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBulkJobOpsBulkUpdateEdgeNodeBaseOSForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBulkJobOpsBulkUpdateEdgeNodeBaseOSNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewBulkJobOpsBulkUpdateEdgeNodeBaseOSConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBulkJobOpsBulkUpdateEdgeNodeBaseOSInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewBulkJobOpsBulkUpdateEdgeNodeBaseOSGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewBulkJobOpsBulkUpdateEdgeNodeBaseOSDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBulkJobOpsBulkUpdateEdgeNodeBaseOSOK creates a BulkJobOpsBulkUpdateEdgeNodeBaseOSOK with default headers values
func NewBulkJobOpsBulkUpdateEdgeNodeBaseOSOK() *BulkJobOpsBulkUpdateEdgeNodeBaseOSOK {
	return &BulkJobOpsBulkUpdateEdgeNodeBaseOSOK{}
}

/*
BulkJobOpsBulkUpdateEdgeNodeBaseOSOK describes a response with status code 200, with default header values.

A successful response.
*/
type BulkJobOpsBulkUpdateEdgeNodeBaseOSOK struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this bulk job ops bulk update edge node base o s o k response has a 2xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this bulk job ops bulk update edge node base o s o k response has a 3xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk job ops bulk update edge node base o s o k response has a 4xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this bulk job ops bulk update edge node base o s o k response has a 5xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSOK) IsServerError() bool {
	return false
}

// IsCode returns true when this bulk job ops bulk update edge node base o s o k response a status code equal to that given
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSOK) IsCode(code int) bool {
	return code == 200
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSOK) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/devices/baseos/upgrade][%d] bulkJobOpsBulkUpdateEdgeNodeBaseOSOK  %+v", 200, o.Payload)
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSOK) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/devices/baseos/upgrade][%d] bulkJobOpsBulkUpdateEdgeNodeBaseOSOK  %+v", 200, o.Payload)
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkUpdateEdgeNodeBaseOSBadRequest creates a BulkJobOpsBulkUpdateEdgeNodeBaseOSBadRequest with default headers values
func NewBulkJobOpsBulkUpdateEdgeNodeBaseOSBadRequest() *BulkJobOpsBulkUpdateEdgeNodeBaseOSBadRequest {
	return &BulkJobOpsBulkUpdateEdgeNodeBaseOSBadRequest{}
}

/*
BulkJobOpsBulkUpdateEdgeNodeBaseOSBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type BulkJobOpsBulkUpdateEdgeNodeBaseOSBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this bulk job ops bulk update edge node base o s bad request response has a 2xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bulk job ops bulk update edge node base o s bad request response has a 3xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk job ops bulk update edge node base o s bad request response has a 4xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this bulk job ops bulk update edge node base o s bad request response has a 5xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this bulk job ops bulk update edge node base o s bad request response a status code equal to that given
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/devices/baseos/upgrade][%d] bulkJobOpsBulkUpdateEdgeNodeBaseOSBadRequest  %+v", 400, o.Payload)
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/devices/baseos/upgrade][%d] bulkJobOpsBulkUpdateEdgeNodeBaseOSBadRequest  %+v", 400, o.Payload)
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkUpdateEdgeNodeBaseOSUnauthorized creates a BulkJobOpsBulkUpdateEdgeNodeBaseOSUnauthorized with default headers values
func NewBulkJobOpsBulkUpdateEdgeNodeBaseOSUnauthorized() *BulkJobOpsBulkUpdateEdgeNodeBaseOSUnauthorized {
	return &BulkJobOpsBulkUpdateEdgeNodeBaseOSUnauthorized{}
}

/*
BulkJobOpsBulkUpdateEdgeNodeBaseOSUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type BulkJobOpsBulkUpdateEdgeNodeBaseOSUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this bulk job ops bulk update edge node base o s unauthorized response has a 2xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bulk job ops bulk update edge node base o s unauthorized response has a 3xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk job ops bulk update edge node base o s unauthorized response has a 4xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this bulk job ops bulk update edge node base o s unauthorized response has a 5xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this bulk job ops bulk update edge node base o s unauthorized response a status code equal to that given
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/devices/baseos/upgrade][%d] bulkJobOpsBulkUpdateEdgeNodeBaseOSUnauthorized  %+v", 401, o.Payload)
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/devices/baseos/upgrade][%d] bulkJobOpsBulkUpdateEdgeNodeBaseOSUnauthorized  %+v", 401, o.Payload)
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkUpdateEdgeNodeBaseOSForbidden creates a BulkJobOpsBulkUpdateEdgeNodeBaseOSForbidden with default headers values
func NewBulkJobOpsBulkUpdateEdgeNodeBaseOSForbidden() *BulkJobOpsBulkUpdateEdgeNodeBaseOSForbidden {
	return &BulkJobOpsBulkUpdateEdgeNodeBaseOSForbidden{}
}

/*
BulkJobOpsBulkUpdateEdgeNodeBaseOSForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-app level access permission for the operation or does not have access scope to the project.
*/
type BulkJobOpsBulkUpdateEdgeNodeBaseOSForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this bulk job ops bulk update edge node base o s forbidden response has a 2xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bulk job ops bulk update edge node base o s forbidden response has a 3xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk job ops bulk update edge node base o s forbidden response has a 4xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this bulk job ops bulk update edge node base o s forbidden response has a 5xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this bulk job ops bulk update edge node base o s forbidden response a status code equal to that given
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/devices/baseos/upgrade][%d] bulkJobOpsBulkUpdateEdgeNodeBaseOSForbidden  %+v", 403, o.Payload)
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/devices/baseos/upgrade][%d] bulkJobOpsBulkUpdateEdgeNodeBaseOSForbidden  %+v", 403, o.Payload)
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkUpdateEdgeNodeBaseOSNotFound creates a BulkJobOpsBulkUpdateEdgeNodeBaseOSNotFound with default headers values
func NewBulkJobOpsBulkUpdateEdgeNodeBaseOSNotFound() *BulkJobOpsBulkUpdateEdgeNodeBaseOSNotFound {
	return &BulkJobOpsBulkUpdateEdgeNodeBaseOSNotFound{}
}

/*
BulkJobOpsBulkUpdateEdgeNodeBaseOSNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type BulkJobOpsBulkUpdateEdgeNodeBaseOSNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this bulk job ops bulk update edge node base o s not found response has a 2xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bulk job ops bulk update edge node base o s not found response has a 3xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk job ops bulk update edge node base o s not found response has a 4xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this bulk job ops bulk update edge node base o s not found response has a 5xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this bulk job ops bulk update edge node base o s not found response a status code equal to that given
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/devices/baseos/upgrade][%d] bulkJobOpsBulkUpdateEdgeNodeBaseOSNotFound  %+v", 404, o.Payload)
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/devices/baseos/upgrade][%d] bulkJobOpsBulkUpdateEdgeNodeBaseOSNotFound  %+v", 404, o.Payload)
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkUpdateEdgeNodeBaseOSConflict creates a BulkJobOpsBulkUpdateEdgeNodeBaseOSConflict with default headers values
func NewBulkJobOpsBulkUpdateEdgeNodeBaseOSConflict() *BulkJobOpsBulkUpdateEdgeNodeBaseOSConflict {
	return &BulkJobOpsBulkUpdateEdgeNodeBaseOSConflict{}
}

/*
BulkJobOpsBulkUpdateEdgeNodeBaseOSConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing asynchronous job request record.
*/
type BulkJobOpsBulkUpdateEdgeNodeBaseOSConflict struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this bulk job ops bulk update edge node base o s conflict response has a 2xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bulk job ops bulk update edge node base o s conflict response has a 3xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk job ops bulk update edge node base o s conflict response has a 4xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this bulk job ops bulk update edge node base o s conflict response has a 5xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this bulk job ops bulk update edge node base o s conflict response a status code equal to that given
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSConflict) IsCode(code int) bool {
	return code == 409
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/devices/baseos/upgrade][%d] bulkJobOpsBulkUpdateEdgeNodeBaseOSConflict  %+v", 409, o.Payload)
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSConflict) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/devices/baseos/upgrade][%d] bulkJobOpsBulkUpdateEdgeNodeBaseOSConflict  %+v", 409, o.Payload)
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkUpdateEdgeNodeBaseOSInternalServerError creates a BulkJobOpsBulkUpdateEdgeNodeBaseOSInternalServerError with default headers values
func NewBulkJobOpsBulkUpdateEdgeNodeBaseOSInternalServerError() *BulkJobOpsBulkUpdateEdgeNodeBaseOSInternalServerError {
	return &BulkJobOpsBulkUpdateEdgeNodeBaseOSInternalServerError{}
}

/*
BulkJobOpsBulkUpdateEdgeNodeBaseOSInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type BulkJobOpsBulkUpdateEdgeNodeBaseOSInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this bulk job ops bulk update edge node base o s internal server error response has a 2xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bulk job ops bulk update edge node base o s internal server error response has a 3xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk job ops bulk update edge node base o s internal server error response has a 4xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this bulk job ops bulk update edge node base o s internal server error response has a 5xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this bulk job ops bulk update edge node base o s internal server error response a status code equal to that given
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/devices/baseos/upgrade][%d] bulkJobOpsBulkUpdateEdgeNodeBaseOSInternalServerError  %+v", 500, o.Payload)
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/devices/baseos/upgrade][%d] bulkJobOpsBulkUpdateEdgeNodeBaseOSInternalServerError  %+v", 500, o.Payload)
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkUpdateEdgeNodeBaseOSGatewayTimeout creates a BulkJobOpsBulkUpdateEdgeNodeBaseOSGatewayTimeout with default headers values
func NewBulkJobOpsBulkUpdateEdgeNodeBaseOSGatewayTimeout() *BulkJobOpsBulkUpdateEdgeNodeBaseOSGatewayTimeout {
	return &BulkJobOpsBulkUpdateEdgeNodeBaseOSGatewayTimeout{}
}

/*
BulkJobOpsBulkUpdateEdgeNodeBaseOSGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type BulkJobOpsBulkUpdateEdgeNodeBaseOSGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this bulk job ops bulk update edge node base o s gateway timeout response has a 2xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bulk job ops bulk update edge node base o s gateway timeout response has a 3xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk job ops bulk update edge node base o s gateway timeout response has a 4xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this bulk job ops bulk update edge node base o s gateway timeout response has a 5xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this bulk job ops bulk update edge node base o s gateway timeout response a status code equal to that given
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/devices/baseos/upgrade][%d] bulkJobOpsBulkUpdateEdgeNodeBaseOSGatewayTimeout  %+v", 504, o.Payload)
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/devices/baseos/upgrade][%d] bulkJobOpsBulkUpdateEdgeNodeBaseOSGatewayTimeout  %+v", 504, o.Payload)
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkJobOpsBulkUpdateEdgeNodeBaseOSDefault creates a BulkJobOpsBulkUpdateEdgeNodeBaseOSDefault with default headers values
func NewBulkJobOpsBulkUpdateEdgeNodeBaseOSDefault(code int) *BulkJobOpsBulkUpdateEdgeNodeBaseOSDefault {
	return &BulkJobOpsBulkUpdateEdgeNodeBaseOSDefault{
		_statusCode: code,
	}
}

/*
BulkJobOpsBulkUpdateEdgeNodeBaseOSDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BulkJobOpsBulkUpdateEdgeNodeBaseOSDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the bulk job ops bulk update edge node base o s default response
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this bulk job ops bulk update edge node base o s default response has a 2xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this bulk job ops bulk update edge node base o s default response has a 3xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this bulk job ops bulk update edge node base o s default response has a 4xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this bulk job ops bulk update edge node base o s default response has a 5xx status code
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this bulk job ops bulk update edge node base o s default response a status code equal to that given
func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/devices/baseos/upgrade][%d] BulkJobOps_BulkUpdateEdgeNodeBaseOS default  %+v", o._statusCode, o.Payload)
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSDefault) String() string {
	return fmt.Sprintf("[PUT /v1/jobs/devices/baseos/upgrade][%d] BulkJobOps_BulkUpdateEdgeNodeBaseOS default  %+v", o._statusCode, o.Payload)
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *BulkJobOpsBulkUpdateEdgeNodeBaseOSDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
