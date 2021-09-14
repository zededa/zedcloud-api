// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package hardware_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// DeleteHardwareModelReader is a Reader for the DeleteHardwareModel structure.
type DeleteHardwareModelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteHardwareModelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteHardwareModelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteHardwareModelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteHardwareModelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteHardwareModelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteHardwareModelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteHardwareModelGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteHardwareModelOK creates a DeleteHardwareModelOK with default headers values
func NewDeleteHardwareModelOK() *DeleteHardwareModelOK {
	return &DeleteHardwareModelOK{}
}

/* DeleteHardwareModelOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteHardwareModelOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *DeleteHardwareModelOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/sysmodels/id/{id}][%d] deleteHardwareModelOK  %+v", 200, o.Payload)
}
func (o *DeleteHardwareModelOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteHardwareModelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteHardwareModelUnauthorized creates a DeleteHardwareModelUnauthorized with default headers values
func NewDeleteHardwareModelUnauthorized() *DeleteHardwareModelUnauthorized {
	return &DeleteHardwareModelUnauthorized{}
}

/* DeleteHardwareModelUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type DeleteHardwareModelUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *DeleteHardwareModelUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/sysmodels/id/{id}][%d] deleteHardwareModelUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteHardwareModelUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteHardwareModelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteHardwareModelForbidden creates a DeleteHardwareModelForbidden with default headers values
func NewDeleteHardwareModelForbidden() *DeleteHardwareModelForbidden {
	return &DeleteHardwareModelForbidden{}
}

/* DeleteHardwareModelForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type DeleteHardwareModelForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *DeleteHardwareModelForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/sysmodels/id/{id}][%d] deleteHardwareModelForbidden  %+v", 403, o.Payload)
}
func (o *DeleteHardwareModelForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteHardwareModelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteHardwareModelNotFound creates a DeleteHardwareModelNotFound with default headers values
func NewDeleteHardwareModelNotFound() *DeleteHardwareModelNotFound {
	return &DeleteHardwareModelNotFound{}
}

/* DeleteHardwareModelNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type DeleteHardwareModelNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *DeleteHardwareModelNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/sysmodels/id/{id}][%d] deleteHardwareModelNotFound  %+v", 404, o.Payload)
}
func (o *DeleteHardwareModelNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteHardwareModelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteHardwareModelInternalServerError creates a DeleteHardwareModelInternalServerError with default headers values
func NewDeleteHardwareModelInternalServerError() *DeleteHardwareModelInternalServerError {
	return &DeleteHardwareModelInternalServerError{}
}

/* DeleteHardwareModelInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type DeleteHardwareModelInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *DeleteHardwareModelInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/sysmodels/id/{id}][%d] deleteHardwareModelInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteHardwareModelInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteHardwareModelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteHardwareModelGatewayTimeout creates a DeleteHardwareModelGatewayTimeout with default headers values
func NewDeleteHardwareModelGatewayTimeout() *DeleteHardwareModelGatewayTimeout {
	return &DeleteHardwareModelGatewayTimeout{}
}

/* DeleteHardwareModelGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type DeleteHardwareModelGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *DeleteHardwareModelGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /v1/sysmodels/id/{id}][%d] deleteHardwareModelGatewayTimeout  %+v", 504, o.Payload)
}
func (o *DeleteHardwareModelGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DeleteHardwareModelGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
