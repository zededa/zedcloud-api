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

// UpdateJobReader is a Reader for the UpdateJob structure.
type UpdateJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateJobConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewUpdateJobGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateJobOK creates a UpdateJobOK with default headers values
func NewUpdateJobOK() *UpdateJobOK {
	return &UpdateJobOK{}
}

/* UpdateJobOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateJobOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateJobOK) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/id/{id}][%d] updateJobOK  %+v", 200, o.Payload)
}
func (o *UpdateJobOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateJobBadRequest creates a UpdateJobBadRequest with default headers values
func NewUpdateJobBadRequest() *UpdateJobBadRequest {
	return &UpdateJobBadRequest{}
}

/* UpdateJobBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type UpdateJobBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateJobBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/id/{id}][%d] updateJobBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateJobBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateJobUnauthorized creates a UpdateJobUnauthorized with default headers values
func NewUpdateJobUnauthorized() *UpdateJobUnauthorized {
	return &UpdateJobUnauthorized{}
}

/* UpdateJobUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type UpdateJobUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateJobUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/id/{id}][%d] updateJobUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateJobUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateJobForbidden creates a UpdateJobForbidden with default headers values
func NewUpdateJobForbidden() *UpdateJobForbidden {
	return &UpdateJobForbidden{}
}

/* UpdateJobForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-app level access permission for the operation or does not have access scope to the project.
*/
type UpdateJobForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateJobForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/id/{id}][%d] updateJobForbidden  %+v", 403, o.Payload)
}
func (o *UpdateJobForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateJobNotFound creates a UpdateJobNotFound with default headers values
func NewUpdateJobNotFound() *UpdateJobNotFound {
	return &UpdateJobNotFound{}
}

/* UpdateJobNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type UpdateJobNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateJobNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/id/{id}][%d] updateJobNotFound  %+v", 404, o.Payload)
}
func (o *UpdateJobNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateJobConflict creates a UpdateJobConflict with default headers values
func NewUpdateJobConflict() *UpdateJobConflict {
	return &UpdateJobConflict{}
}

/* UpdateJobConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing asynchronous job request record.
*/
type UpdateJobConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateJobConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/id/{id}][%d] updateJobConflict  %+v", 409, o.Payload)
}
func (o *UpdateJobConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateJobConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateJobInternalServerError creates a UpdateJobInternalServerError with default headers values
func NewUpdateJobInternalServerError() *UpdateJobInternalServerError {
	return &UpdateJobInternalServerError{}
}

/* UpdateJobInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type UpdateJobInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateJobInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/id/{id}][%d] updateJobInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateJobInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateJobGatewayTimeout creates a UpdateJobGatewayTimeout with default headers values
func NewUpdateJobGatewayTimeout() *UpdateJobGatewayTimeout {
	return &UpdateJobGatewayTimeout{}
}

/* UpdateJobGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type UpdateJobGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateJobGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/jobs/id/{id}][%d] updateJobGatewayTimeout  %+v", 504, o.Payload)
}
func (o *UpdateJobGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateJobGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
