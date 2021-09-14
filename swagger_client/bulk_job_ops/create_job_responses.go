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

// CreateJobReader is a Reader for the CreateJob structure.
type CreateJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateJobConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewCreateJobGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateJobOK creates a CreateJobOK with default headers values
func NewCreateJobOK() *CreateJobOK {
	return &CreateJobOK{}
}

/* CreateJobOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateJobOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateJobOK) Error() string {
	return fmt.Sprintf("[POST /v1/jobs][%d] createJobOK  %+v", 200, o.Payload)
}
func (o *CreateJobOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateJobBadRequest creates a CreateJobBadRequest with default headers values
func NewCreateJobBadRequest() *CreateJobBadRequest {
	return &CreateJobBadRequest{}
}

/* CreateJobBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type CreateJobBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateJobBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/jobs][%d] createJobBadRequest  %+v", 400, o.Payload)
}
func (o *CreateJobBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateJobUnauthorized creates a CreateJobUnauthorized with default headers values
func NewCreateJobUnauthorized() *CreateJobUnauthorized {
	return &CreateJobUnauthorized{}
}

/* CreateJobUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type CreateJobUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateJobUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/jobs][%d] createJobUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateJobUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateJobForbidden creates a CreateJobForbidden with default headers values
func NewCreateJobForbidden() *CreateJobForbidden {
	return &CreateJobForbidden{}
}

/* CreateJobForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-app level access permission for the operation or does not have access scope to the project.
*/
type CreateJobForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateJobForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/jobs][%d] createJobForbidden  %+v", 403, o.Payload)
}
func (o *CreateJobForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateJobConflict creates a CreateJobConflict with default headers values
func NewCreateJobConflict() *CreateJobConflict {
	return &CreateJobConflict{}
}

/* CreateJobConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this job record will conflict with an already existing job record.
*/
type CreateJobConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateJobConflict) Error() string {
	return fmt.Sprintf("[POST /v1/jobs][%d] createJobConflict  %+v", 409, o.Payload)
}
func (o *CreateJobConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateJobConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateJobInternalServerError creates a CreateJobInternalServerError with default headers values
func NewCreateJobInternalServerError() *CreateJobInternalServerError {
	return &CreateJobInternalServerError{}
}

/* CreateJobInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type CreateJobInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateJobInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/jobs][%d] createJobInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateJobInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateJobGatewayTimeout creates a CreateJobGatewayTimeout with default headers values
func NewCreateJobGatewayTimeout() *CreateJobGatewayTimeout {
	return &CreateJobGatewayTimeout{}
}

/* CreateJobGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type CreateJobGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateJobGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/jobs][%d] createJobGatewayTimeout  %+v", 504, o.Payload)
}
func (o *CreateJobGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateJobGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
