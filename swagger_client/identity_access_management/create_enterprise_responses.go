// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package identity_access_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// CreateEnterpriseReader is a Reader for the CreateEnterprise structure.
type CreateEnterpriseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateEnterpriseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateEnterpriseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateEnterpriseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateEnterpriseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateEnterpriseForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateEnterpriseConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateEnterpriseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewCreateEnterpriseGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateEnterpriseOK creates a CreateEnterpriseOK with default headers values
func NewCreateEnterpriseOK() *CreateEnterpriseOK {
	return &CreateEnterpriseOK{}
}

/* CreateEnterpriseOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateEnterpriseOK struct {
	Payload *swagger_models.CrudResponse
}

func (o *CreateEnterpriseOK) Error() string {
	return fmt.Sprintf("[POST /v1/enterprises][%d] createEnterpriseOK  %+v", 200, o.Payload)
}
func (o *CreateEnterpriseOK) GetPayload() *swagger_models.CrudResponse {
	return o.Payload
}

func (o *CreateEnterpriseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.CrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEnterpriseBadRequest creates a CreateEnterpriseBadRequest with default headers values
func NewCreateEnterpriseBadRequest() *CreateEnterpriseBadRequest {
	return &CreateEnterpriseBadRequest{}
}

/* CreateEnterpriseBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type CreateEnterpriseBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEnterpriseBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/enterprises][%d] createEnterpriseBadRequest  %+v", 400, o.Payload)
}
func (o *CreateEnterpriseBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEnterpriseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEnterpriseUnauthorized creates a CreateEnterpriseUnauthorized with default headers values
func NewCreateEnterpriseUnauthorized() *CreateEnterpriseUnauthorized {
	return &CreateEnterpriseUnauthorized{}
}

/* CreateEnterpriseUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type CreateEnterpriseUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEnterpriseUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/enterprises][%d] createEnterpriseUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateEnterpriseUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEnterpriseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEnterpriseForbidden creates a CreateEnterpriseForbidden with default headers values
func NewCreateEnterpriseForbidden() *CreateEnterpriseForbidden {
	return &CreateEnterpriseForbidden{}
}

/* CreateEnterpriseForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type CreateEnterpriseForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEnterpriseForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/enterprises][%d] createEnterpriseForbidden  %+v", 403, o.Payload)
}
func (o *CreateEnterpriseForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEnterpriseForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEnterpriseConflict creates a CreateEnterpriseConflict with default headers values
func NewCreateEnterpriseConflict() *CreateEnterpriseConflict {
	return &CreateEnterpriseConflict{}
}

/* CreateEnterpriseConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this IAM role record will conflict with an already existing IAM role record.
*/
type CreateEnterpriseConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEnterpriseConflict) Error() string {
	return fmt.Sprintf("[POST /v1/enterprises][%d] createEnterpriseConflict  %+v", 409, o.Payload)
}
func (o *CreateEnterpriseConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEnterpriseConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEnterpriseInternalServerError creates a CreateEnterpriseInternalServerError with default headers values
func NewCreateEnterpriseInternalServerError() *CreateEnterpriseInternalServerError {
	return &CreateEnterpriseInternalServerError{}
}

/* CreateEnterpriseInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type CreateEnterpriseInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEnterpriseInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/enterprises][%d] createEnterpriseInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateEnterpriseInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEnterpriseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEnterpriseGatewayTimeout creates a CreateEnterpriseGatewayTimeout with default headers values
func NewCreateEnterpriseGatewayTimeout() *CreateEnterpriseGatewayTimeout {
	return &CreateEnterpriseGatewayTimeout{}
}

/* CreateEnterpriseGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type CreateEnterpriseGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEnterpriseGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/enterprises][%d] createEnterpriseGatewayTimeout  %+v", 504, o.Payload)
}
func (o *CreateEnterpriseGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEnterpriseGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
