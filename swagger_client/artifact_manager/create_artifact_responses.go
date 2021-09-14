// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package artifact_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// CreateArtifactReader is a Reader for the CreateArtifact structure.
type CreateArtifactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateArtifactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateArtifactOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateArtifactBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateArtifactUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateArtifactForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateArtifactInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewCreateArtifactGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateArtifactOK creates a CreateArtifactOK with default headers values
func NewCreateArtifactOK() *CreateArtifactOK {
	return &CreateArtifactOK{}
}

/* CreateArtifactOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateArtifactOK struct {
	Payload *swagger_models.Artifact
}

func (o *CreateArtifactOK) Error() string {
	return fmt.Sprintf("[POST /v1/artifacts][%d] createArtifactOK  %+v", 200, o.Payload)
}
func (o *CreateArtifactOK) GetPayload() *swagger_models.Artifact {
	return o.Payload
}

func (o *CreateArtifactOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.Artifact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateArtifactBadRequest creates a CreateArtifactBadRequest with default headers values
func NewCreateArtifactBadRequest() *CreateArtifactBadRequest {
	return &CreateArtifactBadRequest{}
}

/* CreateArtifactBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type CreateArtifactBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateArtifactBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/artifacts][%d] createArtifactBadRequest  %+v", 400, o.Payload)
}
func (o *CreateArtifactBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateArtifactBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateArtifactUnauthorized creates a CreateArtifactUnauthorized with default headers values
func NewCreateArtifactUnauthorized() *CreateArtifactUnauthorized {
	return &CreateArtifactUnauthorized{}
}

/* CreateArtifactUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type CreateArtifactUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateArtifactUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/artifacts][%d] createArtifactUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateArtifactUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateArtifactUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateArtifactForbidden creates a CreateArtifactForbidden with default headers values
func NewCreateArtifactForbidden() *CreateArtifactForbidden {
	return &CreateArtifactForbidden{}
}

/* CreateArtifactForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type CreateArtifactForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateArtifactForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/artifacts][%d] createArtifactForbidden  %+v", 403, o.Payload)
}
func (o *CreateArtifactForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateArtifactForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateArtifactInternalServerError creates a CreateArtifactInternalServerError with default headers values
func NewCreateArtifactInternalServerError() *CreateArtifactInternalServerError {
	return &CreateArtifactInternalServerError{}
}

/* CreateArtifactInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type CreateArtifactInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateArtifactInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/artifacts][%d] createArtifactInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateArtifactInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateArtifactInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateArtifactGatewayTimeout creates a CreateArtifactGatewayTimeout with default headers values
func NewCreateArtifactGatewayTimeout() *CreateArtifactGatewayTimeout {
	return &CreateArtifactGatewayTimeout{}
}

/* CreateArtifactGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type CreateArtifactGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateArtifactGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/artifacts][%d] createArtifactGatewayTimeout  %+v", 504, o.Payload)
}
func (o *CreateArtifactGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateArtifactGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
