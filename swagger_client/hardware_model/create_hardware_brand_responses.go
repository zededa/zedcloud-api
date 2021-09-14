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

// CreateHardwareBrandReader is a Reader for the CreateHardwareBrand structure.
type CreateHardwareBrandReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateHardwareBrandReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateHardwareBrandOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateHardwareBrandBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateHardwareBrandUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateHardwareBrandForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateHardwareBrandConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateHardwareBrandInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewCreateHardwareBrandGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateHardwareBrandOK creates a CreateHardwareBrandOK with default headers values
func NewCreateHardwareBrandOK() *CreateHardwareBrandOK {
	return &CreateHardwareBrandOK{}
}

/* CreateHardwareBrandOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateHardwareBrandOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateHardwareBrandOK) Error() string {
	return fmt.Sprintf("[POST /v1/brands][%d] createHardwareBrandOK  %+v", 200, o.Payload)
}
func (o *CreateHardwareBrandOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateHardwareBrandOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHardwareBrandBadRequest creates a CreateHardwareBrandBadRequest with default headers values
func NewCreateHardwareBrandBadRequest() *CreateHardwareBrandBadRequest {
	return &CreateHardwareBrandBadRequest{}
}

/* CreateHardwareBrandBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type CreateHardwareBrandBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateHardwareBrandBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/brands][%d] createHardwareBrandBadRequest  %+v", 400, o.Payload)
}
func (o *CreateHardwareBrandBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateHardwareBrandBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHardwareBrandUnauthorized creates a CreateHardwareBrandUnauthorized with default headers values
func NewCreateHardwareBrandUnauthorized() *CreateHardwareBrandUnauthorized {
	return &CreateHardwareBrandUnauthorized{}
}

/* CreateHardwareBrandUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type CreateHardwareBrandUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateHardwareBrandUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/brands][%d] createHardwareBrandUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateHardwareBrandUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateHardwareBrandUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHardwareBrandForbidden creates a CreateHardwareBrandForbidden with default headers values
func NewCreateHardwareBrandForbidden() *CreateHardwareBrandForbidden {
	return &CreateHardwareBrandForbidden{}
}

/* CreateHardwareBrandForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type CreateHardwareBrandForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateHardwareBrandForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/brands][%d] createHardwareBrandForbidden  %+v", 403, o.Payload)
}
func (o *CreateHardwareBrandForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateHardwareBrandForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHardwareBrandConflict creates a CreateHardwareBrandConflict with default headers values
func NewCreateHardwareBrandConflict() *CreateHardwareBrandConflict {
	return &CreateHardwareBrandConflict{}
}

/* CreateHardwareBrandConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this hardware brand record will conflict with an already existing hardware brand record.
*/
type CreateHardwareBrandConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateHardwareBrandConflict) Error() string {
	return fmt.Sprintf("[POST /v1/brands][%d] createHardwareBrandConflict  %+v", 409, o.Payload)
}
func (o *CreateHardwareBrandConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateHardwareBrandConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHardwareBrandInternalServerError creates a CreateHardwareBrandInternalServerError with default headers values
func NewCreateHardwareBrandInternalServerError() *CreateHardwareBrandInternalServerError {
	return &CreateHardwareBrandInternalServerError{}
}

/* CreateHardwareBrandInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type CreateHardwareBrandInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateHardwareBrandInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/brands][%d] createHardwareBrandInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateHardwareBrandInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateHardwareBrandInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHardwareBrandGatewayTimeout creates a CreateHardwareBrandGatewayTimeout with default headers values
func NewCreateHardwareBrandGatewayTimeout() *CreateHardwareBrandGatewayTimeout {
	return &CreateHardwareBrandGatewayTimeout{}
}

/* CreateHardwareBrandGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type CreateHardwareBrandGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateHardwareBrandGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/brands][%d] createHardwareBrandGatewayTimeout  %+v", 504, o.Payload)
}
func (o *CreateHardwareBrandGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateHardwareBrandGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
