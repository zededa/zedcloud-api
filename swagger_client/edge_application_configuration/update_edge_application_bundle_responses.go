// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_application_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// UpdateEdgeApplicationBundleReader is a Reader for the UpdateEdgeApplicationBundle structure.
type UpdateEdgeApplicationBundleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateEdgeApplicationBundleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateEdgeApplicationBundleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateEdgeApplicationBundleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateEdgeApplicationBundleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateEdgeApplicationBundleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateEdgeApplicationBundleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateEdgeApplicationBundleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewUpdateEdgeApplicationBundleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateEdgeApplicationBundleOK creates a UpdateEdgeApplicationBundleOK with default headers values
func NewUpdateEdgeApplicationBundleOK() *UpdateEdgeApplicationBundleOK {
	return &UpdateEdgeApplicationBundleOK{}
}

/* UpdateEdgeApplicationBundleOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateEdgeApplicationBundleOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateEdgeApplicationBundleOK) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/id/{id}][%d] updateEdgeApplicationBundleOK  %+v", 200, o.Payload)
}
func (o *UpdateEdgeApplicationBundleOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateEdgeApplicationBundleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEdgeApplicationBundleUnauthorized creates a UpdateEdgeApplicationBundleUnauthorized with default headers values
func NewUpdateEdgeApplicationBundleUnauthorized() *UpdateEdgeApplicationBundleUnauthorized {
	return &UpdateEdgeApplicationBundleUnauthorized{}
}

/* UpdateEdgeApplicationBundleUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type UpdateEdgeApplicationBundleUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateEdgeApplicationBundleUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/id/{id}][%d] updateEdgeApplicationBundleUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateEdgeApplicationBundleUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateEdgeApplicationBundleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEdgeApplicationBundleForbidden creates a UpdateEdgeApplicationBundleForbidden with default headers values
func NewUpdateEdgeApplicationBundleForbidden() *UpdateEdgeApplicationBundleForbidden {
	return &UpdateEdgeApplicationBundleForbidden{}
}

/* UpdateEdgeApplicationBundleForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type UpdateEdgeApplicationBundleForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateEdgeApplicationBundleForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/id/{id}][%d] updateEdgeApplicationBundleForbidden  %+v", 403, o.Payload)
}
func (o *UpdateEdgeApplicationBundleForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateEdgeApplicationBundleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEdgeApplicationBundleNotFound creates a UpdateEdgeApplicationBundleNotFound with default headers values
func NewUpdateEdgeApplicationBundleNotFound() *UpdateEdgeApplicationBundleNotFound {
	return &UpdateEdgeApplicationBundleNotFound{}
}

/* UpdateEdgeApplicationBundleNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type UpdateEdgeApplicationBundleNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateEdgeApplicationBundleNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/id/{id}][%d] updateEdgeApplicationBundleNotFound  %+v", 404, o.Payload)
}
func (o *UpdateEdgeApplicationBundleNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateEdgeApplicationBundleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEdgeApplicationBundleConflict creates a UpdateEdgeApplicationBundleConflict with default headers values
func NewUpdateEdgeApplicationBundleConflict() *UpdateEdgeApplicationBundleConflict {
	return &UpdateEdgeApplicationBundleConflict{}
}

/* UpdateEdgeApplicationBundleConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing edge network record.
*/
type UpdateEdgeApplicationBundleConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateEdgeApplicationBundleConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/id/{id}][%d] updateEdgeApplicationBundleConflict  %+v", 409, o.Payload)
}
func (o *UpdateEdgeApplicationBundleConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateEdgeApplicationBundleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEdgeApplicationBundleInternalServerError creates a UpdateEdgeApplicationBundleInternalServerError with default headers values
func NewUpdateEdgeApplicationBundleInternalServerError() *UpdateEdgeApplicationBundleInternalServerError {
	return &UpdateEdgeApplicationBundleInternalServerError{}
}

/* UpdateEdgeApplicationBundleInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type UpdateEdgeApplicationBundleInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateEdgeApplicationBundleInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/id/{id}][%d] updateEdgeApplicationBundleInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateEdgeApplicationBundleInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateEdgeApplicationBundleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEdgeApplicationBundleGatewayTimeout creates a UpdateEdgeApplicationBundleGatewayTimeout with default headers values
func NewUpdateEdgeApplicationBundleGatewayTimeout() *UpdateEdgeApplicationBundleGatewayTimeout {
	return &UpdateEdgeApplicationBundleGatewayTimeout{}
}

/* UpdateEdgeApplicationBundleGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type UpdateEdgeApplicationBundleGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateEdgeApplicationBundleGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/id/{id}][%d] updateEdgeApplicationBundleGatewayTimeout  %+v", 504, o.Payload)
}
func (o *UpdateEdgeApplicationBundleGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateEdgeApplicationBundleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
