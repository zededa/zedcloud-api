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

// GetJobByNameReader is a Reader for the GetJobByName structure.
type GetJobByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJobByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJobByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetJobByNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetJobByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetJobByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetJobByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetJobByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetJobByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetJobByNameOK creates a GetJobByNameOK with default headers values
func NewGetJobByNameOK() *GetJobByNameOK {
	return &GetJobByNameOK{}
}

/* GetJobByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetJobByNameOK struct {
	Payload *swagger_models.JobConfig
}

func (o *GetJobByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/jobs/name/{name}/objectType/{objectType}][%d] getJobByNameOK  %+v", 200, o.Payload)
}
func (o *GetJobByNameOK) GetPayload() *swagger_models.JobConfig {
	return o.Payload
}

func (o *GetJobByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.JobConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobByNameBadRequest creates a GetJobByNameBadRequest with default headers values
func NewGetJobByNameBadRequest() *GetJobByNameBadRequest {
	return &GetJobByNameBadRequest{}
}

/* GetJobByNameBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type GetJobByNameBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetJobByNameBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/jobs/name/{name}/objectType/{objectType}][%d] getJobByNameBadRequest  %+v", 400, o.Payload)
}
func (o *GetJobByNameBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetJobByNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobByNameUnauthorized creates a GetJobByNameUnauthorized with default headers values
func NewGetJobByNameUnauthorized() *GetJobByNameUnauthorized {
	return &GetJobByNameUnauthorized{}
}

/* GetJobByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetJobByNameUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetJobByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/jobs/name/{name}/objectType/{objectType}][%d] getJobByNameUnauthorized  %+v", 401, o.Payload)
}
func (o *GetJobByNameUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetJobByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobByNameForbidden creates a GetJobByNameForbidden with default headers values
func NewGetJobByNameForbidden() *GetJobByNameForbidden {
	return &GetJobByNameForbidden{}
}

/* GetJobByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-app level access permission for the operation or does not have access scope to the project.
*/
type GetJobByNameForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetJobByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/jobs/name/{name}/objectType/{objectType}][%d] getJobByNameForbidden  %+v", 403, o.Payload)
}
func (o *GetJobByNameForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetJobByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobByNameNotFound creates a GetJobByNameNotFound with default headers values
func NewGetJobByNameNotFound() *GetJobByNameNotFound {
	return &GetJobByNameNotFound{}
}

/* GetJobByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetJobByNameNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetJobByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/jobs/name/{name}/objectType/{objectType}][%d] getJobByNameNotFound  %+v", 404, o.Payload)
}
func (o *GetJobByNameNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetJobByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobByNameInternalServerError creates a GetJobByNameInternalServerError with default headers values
func NewGetJobByNameInternalServerError() *GetJobByNameInternalServerError {
	return &GetJobByNameInternalServerError{}
}

/* GetJobByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetJobByNameInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetJobByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/jobs/name/{name}/objectType/{objectType}][%d] getJobByNameInternalServerError  %+v", 500, o.Payload)
}
func (o *GetJobByNameInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetJobByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobByNameGatewayTimeout creates a GetJobByNameGatewayTimeout with default headers values
func NewGetJobByNameGatewayTimeout() *GetJobByNameGatewayTimeout {
	return &GetJobByNameGatewayTimeout{}
}

/* GetJobByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetJobByNameGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetJobByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/jobs/name/{name}/objectType/{objectType}][%d] getJobByNameGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetJobByNameGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetJobByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
