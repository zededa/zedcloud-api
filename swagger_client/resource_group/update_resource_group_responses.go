// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package resource_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// UpdateResourceGroupReader is a Reader for the UpdateResourceGroup structure.
type UpdateResourceGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateResourceGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateResourceGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateResourceGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateResourceGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateResourceGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateResourceGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateResourceGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewUpdateResourceGroupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateResourceGroupOK creates a UpdateResourceGroupOK with default headers values
func NewUpdateResourceGroupOK() *UpdateResourceGroupOK {
	return &UpdateResourceGroupOK{}
}

/* UpdateResourceGroupOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateResourceGroupOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateResourceGroupOK) Error() string {
	return fmt.Sprintf("[PUT /v1/projects/id/{id}][%d] updateResourceGroupOK  %+v", 200, o.Payload)
}
func (o *UpdateResourceGroupOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateResourceGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateResourceGroupUnauthorized creates a UpdateResourceGroupUnauthorized with default headers values
func NewUpdateResourceGroupUnauthorized() *UpdateResourceGroupUnauthorized {
	return &UpdateResourceGroupUnauthorized{}
}

/* UpdateResourceGroupUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type UpdateResourceGroupUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateResourceGroupUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/projects/id/{id}][%d] updateResourceGroupUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateResourceGroupUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateResourceGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateResourceGroupForbidden creates a UpdateResourceGroupForbidden with default headers values
func NewUpdateResourceGroupForbidden() *UpdateResourceGroupForbidden {
	return &UpdateResourceGroupForbidden{}
}

/* UpdateResourceGroupForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type UpdateResourceGroupForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateResourceGroupForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/projects/id/{id}][%d] updateResourceGroupForbidden  %+v", 403, o.Payload)
}
func (o *UpdateResourceGroupForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateResourceGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateResourceGroupNotFound creates a UpdateResourceGroupNotFound with default headers values
func NewUpdateResourceGroupNotFound() *UpdateResourceGroupNotFound {
	return &UpdateResourceGroupNotFound{}
}

/* UpdateResourceGroupNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type UpdateResourceGroupNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateResourceGroupNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/projects/id/{id}][%d] updateResourceGroupNotFound  %+v", 404, o.Payload)
}
func (o *UpdateResourceGroupNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateResourceGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateResourceGroupConflict creates a UpdateResourceGroupConflict with default headers values
func NewUpdateResourceGroupConflict() *UpdateResourceGroupConflict {
	return &UpdateResourceGroupConflict{}
}

/* UpdateResourceGroupConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing resource group record.
*/
type UpdateResourceGroupConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateResourceGroupConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/projects/id/{id}][%d] updateResourceGroupConflict  %+v", 409, o.Payload)
}
func (o *UpdateResourceGroupConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateResourceGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateResourceGroupInternalServerError creates a UpdateResourceGroupInternalServerError with default headers values
func NewUpdateResourceGroupInternalServerError() *UpdateResourceGroupInternalServerError {
	return &UpdateResourceGroupInternalServerError{}
}

/* UpdateResourceGroupInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type UpdateResourceGroupInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateResourceGroupInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/projects/id/{id}][%d] updateResourceGroupInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateResourceGroupInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateResourceGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateResourceGroupGatewayTimeout creates a UpdateResourceGroupGatewayTimeout with default headers values
func NewUpdateResourceGroupGatewayTimeout() *UpdateResourceGroupGatewayTimeout {
	return &UpdateResourceGroupGatewayTimeout{}
}

/* UpdateResourceGroupGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type UpdateResourceGroupGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *UpdateResourceGroupGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/projects/id/{id}][%d] updateResourceGroupGatewayTimeout  %+v", 504, o.Payload)
}
func (o *UpdateResourceGroupGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *UpdateResourceGroupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
