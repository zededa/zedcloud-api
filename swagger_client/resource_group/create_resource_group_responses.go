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

// CreateResourceGroupReader is a Reader for the CreateResourceGroup structure.
type CreateResourceGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateResourceGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateResourceGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateResourceGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateResourceGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateResourceGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateResourceGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateResourceGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewCreateResourceGroupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateResourceGroupOK creates a CreateResourceGroupOK with default headers values
func NewCreateResourceGroupOK() *CreateResourceGroupOK {
	return &CreateResourceGroupOK{}
}

/* CreateResourceGroupOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateResourceGroupOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateResourceGroupOK) Error() string {
	return fmt.Sprintf("[POST /v1/projects][%d] createResourceGroupOK  %+v", 200, o.Payload)
}
func (o *CreateResourceGroupOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateResourceGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateResourceGroupBadRequest creates a CreateResourceGroupBadRequest with default headers values
func NewCreateResourceGroupBadRequest() *CreateResourceGroupBadRequest {
	return &CreateResourceGroupBadRequest{}
}

/* CreateResourceGroupBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type CreateResourceGroupBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateResourceGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/projects][%d] createResourceGroupBadRequest  %+v", 400, o.Payload)
}
func (o *CreateResourceGroupBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateResourceGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateResourceGroupUnauthorized creates a CreateResourceGroupUnauthorized with default headers values
func NewCreateResourceGroupUnauthorized() *CreateResourceGroupUnauthorized {
	return &CreateResourceGroupUnauthorized{}
}

/* CreateResourceGroupUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type CreateResourceGroupUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateResourceGroupUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/projects][%d] createResourceGroupUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateResourceGroupUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateResourceGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateResourceGroupForbidden creates a CreateResourceGroupForbidden with default headers values
func NewCreateResourceGroupForbidden() *CreateResourceGroupForbidden {
	return &CreateResourceGroupForbidden{}
}

/* CreateResourceGroupForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type CreateResourceGroupForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateResourceGroupForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/projects][%d] createResourceGroupForbidden  %+v", 403, o.Payload)
}
func (o *CreateResourceGroupForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateResourceGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateResourceGroupConflict creates a CreateResourceGroupConflict with default headers values
func NewCreateResourceGroupConflict() *CreateResourceGroupConflict {
	return &CreateResourceGroupConflict{}
}

/* CreateResourceGroupConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this resource group record will conflict with an already existing resource group record.
*/
type CreateResourceGroupConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateResourceGroupConflict) Error() string {
	return fmt.Sprintf("[POST /v1/projects][%d] createResourceGroupConflict  %+v", 409, o.Payload)
}
func (o *CreateResourceGroupConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateResourceGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateResourceGroupInternalServerError creates a CreateResourceGroupInternalServerError with default headers values
func NewCreateResourceGroupInternalServerError() *CreateResourceGroupInternalServerError {
	return &CreateResourceGroupInternalServerError{}
}

/* CreateResourceGroupInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type CreateResourceGroupInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateResourceGroupInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/projects][%d] createResourceGroupInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateResourceGroupInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateResourceGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateResourceGroupGatewayTimeout creates a CreateResourceGroupGatewayTimeout with default headers values
func NewCreateResourceGroupGatewayTimeout() *CreateResourceGroupGatewayTimeout {
	return &CreateResourceGroupGatewayTimeout{}
}

/* CreateResourceGroupGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type CreateResourceGroupGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateResourceGroupGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/projects][%d] createResourceGroupGatewayTimeout  %+v", 504, o.Payload)
}
func (o *CreateResourceGroupGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateResourceGroupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
