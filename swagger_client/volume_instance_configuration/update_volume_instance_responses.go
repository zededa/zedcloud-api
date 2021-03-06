// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package volume_instance_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// UpdateVolumeInstanceReader is a Reader for the UpdateVolumeInstance structure.
type UpdateVolumeInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVolumeInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVolumeInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateVolumeInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateVolumeInstanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateVolumeInstanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateVolumeInstanceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateVolumeInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewUpdateVolumeInstanceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateVolumeInstanceOK creates a UpdateVolumeInstanceOK with default headers values
func NewUpdateVolumeInstanceOK() *UpdateVolumeInstanceOK {
	return &UpdateVolumeInstanceOK{}
}

/* UpdateVolumeInstanceOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateVolumeInstanceOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateVolumeInstanceOK) Error() string {
	return fmt.Sprintf("[PUT /v1/volumes/instances/id/{id}][%d] updateVolumeInstanceOK  %+v", 200, o.Payload)
}
func (o *UpdateVolumeInstanceOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateVolumeInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVolumeInstanceUnauthorized creates a UpdateVolumeInstanceUnauthorized with default headers values
func NewUpdateVolumeInstanceUnauthorized() *UpdateVolumeInstanceUnauthorized {
	return &UpdateVolumeInstanceUnauthorized{}
}

/* UpdateVolumeInstanceUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type UpdateVolumeInstanceUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateVolumeInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/volumes/instances/id/{id}][%d] updateVolumeInstanceUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateVolumeInstanceUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateVolumeInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVolumeInstanceForbidden creates a UpdateVolumeInstanceForbidden with default headers values
func NewUpdateVolumeInstanceForbidden() *UpdateVolumeInstanceForbidden {
	return &UpdateVolumeInstanceForbidden{}
}

/* UpdateVolumeInstanceForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type UpdateVolumeInstanceForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateVolumeInstanceForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/volumes/instances/id/{id}][%d] updateVolumeInstanceForbidden  %+v", 403, o.Payload)
}
func (o *UpdateVolumeInstanceForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateVolumeInstanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVolumeInstanceNotFound creates a UpdateVolumeInstanceNotFound with default headers values
func NewUpdateVolumeInstanceNotFound() *UpdateVolumeInstanceNotFound {
	return &UpdateVolumeInstanceNotFound{}
}

/* UpdateVolumeInstanceNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type UpdateVolumeInstanceNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateVolumeInstanceNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/volumes/instances/id/{id}][%d] updateVolumeInstanceNotFound  %+v", 404, o.Payload)
}
func (o *UpdateVolumeInstanceNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateVolumeInstanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVolumeInstanceConflict creates a UpdateVolumeInstanceConflict with default headers values
func NewUpdateVolumeInstanceConflict() *UpdateVolumeInstanceConflict {
	return &UpdateVolumeInstanceConflict{}
}

/* UpdateVolumeInstanceConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing edge volume instance record.
*/
type UpdateVolumeInstanceConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateVolumeInstanceConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/volumes/instances/id/{id}][%d] updateVolumeInstanceConflict  %+v", 409, o.Payload)
}
func (o *UpdateVolumeInstanceConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateVolumeInstanceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVolumeInstanceInternalServerError creates a UpdateVolumeInstanceInternalServerError with default headers values
func NewUpdateVolumeInstanceInternalServerError() *UpdateVolumeInstanceInternalServerError {
	return &UpdateVolumeInstanceInternalServerError{}
}

/* UpdateVolumeInstanceInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type UpdateVolumeInstanceInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateVolumeInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/volumes/instances/id/{id}][%d] updateVolumeInstanceInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateVolumeInstanceInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateVolumeInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVolumeInstanceGatewayTimeout creates a UpdateVolumeInstanceGatewayTimeout with default headers values
func NewUpdateVolumeInstanceGatewayTimeout() *UpdateVolumeInstanceGatewayTimeout {
	return &UpdateVolumeInstanceGatewayTimeout{}
}

/* UpdateVolumeInstanceGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type UpdateVolumeInstanceGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateVolumeInstanceGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/volumes/instances/id/{id}][%d] updateVolumeInstanceGatewayTimeout  %+v", 504, o.Payload)
}
func (o *UpdateVolumeInstanceGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateVolumeInstanceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
