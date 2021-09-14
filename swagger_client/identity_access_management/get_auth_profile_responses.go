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

// GetAuthProfileReader is a Reader for the GetAuthProfile structure.
type GetAuthProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuthProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAuthProfileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAuthProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAuthProfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAuthProfileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAuthProfileGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAuthProfileOK creates a GetAuthProfileOK with default headers values
func NewGetAuthProfileOK() *GetAuthProfileOK {
	return &GetAuthProfileOK{}
}

/* GetAuthProfileOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetAuthProfileOK struct {
	Payload *swagger_models.CrudResponseRead
}

func (o *GetAuthProfileOK) Error() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles/id/{id}][%d] getAuthProfileOK  %+v", 200, o.Payload)
}
func (o *GetAuthProfileOK) GetPayload() *swagger_models.CrudResponseRead {
	return o.Payload
}

func (o *GetAuthProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.CrudResponseRead)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthProfileUnauthorized creates a GetAuthProfileUnauthorized with default headers values
func NewGetAuthProfileUnauthorized() *GetAuthProfileUnauthorized {
	return &GetAuthProfileUnauthorized{}
}

/* GetAuthProfileUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetAuthProfileUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetAuthProfileUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles/id/{id}][%d] getAuthProfileUnauthorized  %+v", 401, o.Payload)
}
func (o *GetAuthProfileUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetAuthProfileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthProfileForbidden creates a GetAuthProfileForbidden with default headers values
func NewGetAuthProfileForbidden() *GetAuthProfileForbidden {
	return &GetAuthProfileForbidden{}
}

/* GetAuthProfileForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type GetAuthProfileForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetAuthProfileForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles/id/{id}][%d] getAuthProfileForbidden  %+v", 403, o.Payload)
}
func (o *GetAuthProfileForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetAuthProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthProfileNotFound creates a GetAuthProfileNotFound with default headers values
func NewGetAuthProfileNotFound() *GetAuthProfileNotFound {
	return &GetAuthProfileNotFound{}
}

/* GetAuthProfileNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetAuthProfileNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetAuthProfileNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles/id/{id}][%d] getAuthProfileNotFound  %+v", 404, o.Payload)
}
func (o *GetAuthProfileNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetAuthProfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthProfileInternalServerError creates a GetAuthProfileInternalServerError with default headers values
func NewGetAuthProfileInternalServerError() *GetAuthProfileInternalServerError {
	return &GetAuthProfileInternalServerError{}
}

/* GetAuthProfileInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetAuthProfileInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetAuthProfileInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles/id/{id}][%d] getAuthProfileInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAuthProfileInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetAuthProfileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthProfileGatewayTimeout creates a GetAuthProfileGatewayTimeout with default headers values
func NewGetAuthProfileGatewayTimeout() *GetAuthProfileGatewayTimeout {
	return &GetAuthProfileGatewayTimeout{}
}

/* GetAuthProfileGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetAuthProfileGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetAuthProfileGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles/id/{id}][%d] getAuthProfileGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetAuthProfileGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetAuthProfileGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
