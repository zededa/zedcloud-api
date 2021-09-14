// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package datastore_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// UpdateDatastoreReader is a Reader for the UpdateDatastore structure.
type UpdateDatastoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDatastoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDatastoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateDatastoreUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateDatastoreForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateDatastoreNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateDatastoreConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateDatastoreInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewUpdateDatastoreGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateDatastoreOK creates a UpdateDatastoreOK with default headers values
func NewUpdateDatastoreOK() *UpdateDatastoreOK {
	return &UpdateDatastoreOK{}
}

/* UpdateDatastoreOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateDatastoreOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateDatastoreOK) Error() string {
	return fmt.Sprintf("[PUT /v1/datastores/id/{id}][%d] updateDatastoreOK  %+v", 200, o.Payload)
}
func (o *UpdateDatastoreOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateDatastoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDatastoreUnauthorized creates a UpdateDatastoreUnauthorized with default headers values
func NewUpdateDatastoreUnauthorized() *UpdateDatastoreUnauthorized {
	return &UpdateDatastoreUnauthorized{}
}

/* UpdateDatastoreUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type UpdateDatastoreUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateDatastoreUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/datastores/id/{id}][%d] updateDatastoreUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateDatastoreUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateDatastoreUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDatastoreForbidden creates a UpdateDatastoreForbidden with default headers values
func NewUpdateDatastoreForbidden() *UpdateDatastoreForbidden {
	return &UpdateDatastoreForbidden{}
}

/* UpdateDatastoreForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type UpdateDatastoreForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateDatastoreForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/datastores/id/{id}][%d] updateDatastoreForbidden  %+v", 403, o.Payload)
}
func (o *UpdateDatastoreForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateDatastoreForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDatastoreNotFound creates a UpdateDatastoreNotFound with default headers values
func NewUpdateDatastoreNotFound() *UpdateDatastoreNotFound {
	return &UpdateDatastoreNotFound{}
}

/* UpdateDatastoreNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type UpdateDatastoreNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateDatastoreNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/datastores/id/{id}][%d] updateDatastoreNotFound  %+v", 404, o.Payload)
}
func (o *UpdateDatastoreNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateDatastoreNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDatastoreConflict creates a UpdateDatastoreConflict with default headers values
func NewUpdateDatastoreConflict() *UpdateDatastoreConflict {
	return &UpdateDatastoreConflict{}
}

/* UpdateDatastoreConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing datastore record.
*/
type UpdateDatastoreConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateDatastoreConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/datastores/id/{id}][%d] updateDatastoreConflict  %+v", 409, o.Payload)
}
func (o *UpdateDatastoreConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateDatastoreConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDatastoreInternalServerError creates a UpdateDatastoreInternalServerError with default headers values
func NewUpdateDatastoreInternalServerError() *UpdateDatastoreInternalServerError {
	return &UpdateDatastoreInternalServerError{}
}

/* UpdateDatastoreInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type UpdateDatastoreInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateDatastoreInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/datastores/id/{id}][%d] updateDatastoreInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateDatastoreInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateDatastoreInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDatastoreGatewayTimeout creates a UpdateDatastoreGatewayTimeout with default headers values
func NewUpdateDatastoreGatewayTimeout() *UpdateDatastoreGatewayTimeout {
	return &UpdateDatastoreGatewayTimeout{}
}

/* UpdateDatastoreGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type UpdateDatastoreGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateDatastoreGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/datastores/id/{id}][%d] updateDatastoreGatewayTimeout  %+v", 504, o.Payload)
}
func (o *UpdateDatastoreGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateDatastoreGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
