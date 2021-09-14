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

// QueryAuthProfilesReader is a Reader for the QueryAuthProfiles structure.
type QueryAuthProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryAuthProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryAuthProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryAuthProfilesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewQueryAuthProfilesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryAuthProfilesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryAuthProfilesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewQueryAuthProfilesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryAuthProfilesOK creates a QueryAuthProfilesOK with default headers values
func NewQueryAuthProfilesOK() *QueryAuthProfilesOK {
	return &QueryAuthProfilesOK{}
}

/* QueryAuthProfilesOK describes a response with status code 200, with default header values.

A successful response.
*/
type QueryAuthProfilesOK struct {
	Payload *swagger_models.CrudResponse
}

func (o *QueryAuthProfilesOK) Error() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles][%d] queryAuthProfilesOK  %+v", 200, o.Payload)
}
func (o *QueryAuthProfilesOK) GetPayload() *swagger_models.CrudResponse {
	return o.Payload
}

func (o *QueryAuthProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.CrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryAuthProfilesBadRequest creates a QueryAuthProfilesBadRequest with default headers values
func NewQueryAuthProfilesBadRequest() *QueryAuthProfilesBadRequest {
	return &QueryAuthProfilesBadRequest{}
}

/* QueryAuthProfilesBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type QueryAuthProfilesBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryAuthProfilesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles][%d] queryAuthProfilesBadRequest  %+v", 400, o.Payload)
}
func (o *QueryAuthProfilesBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryAuthProfilesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryAuthProfilesUnauthorized creates a QueryAuthProfilesUnauthorized with default headers values
func NewQueryAuthProfilesUnauthorized() *QueryAuthProfilesUnauthorized {
	return &QueryAuthProfilesUnauthorized{}
}

/* QueryAuthProfilesUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type QueryAuthProfilesUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryAuthProfilesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles][%d] queryAuthProfilesUnauthorized  %+v", 401, o.Payload)
}
func (o *QueryAuthProfilesUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryAuthProfilesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryAuthProfilesForbidden creates a QueryAuthProfilesForbidden with default headers values
func NewQueryAuthProfilesForbidden() *QueryAuthProfilesForbidden {
	return &QueryAuthProfilesForbidden{}
}

/* QueryAuthProfilesForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type QueryAuthProfilesForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryAuthProfilesForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles][%d] queryAuthProfilesForbidden  %+v", 403, o.Payload)
}
func (o *QueryAuthProfilesForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryAuthProfilesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryAuthProfilesInternalServerError creates a QueryAuthProfilesInternalServerError with default headers values
func NewQueryAuthProfilesInternalServerError() *QueryAuthProfilesInternalServerError {
	return &QueryAuthProfilesInternalServerError{}
}

/* QueryAuthProfilesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type QueryAuthProfilesInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryAuthProfilesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles][%d] queryAuthProfilesInternalServerError  %+v", 500, o.Payload)
}
func (o *QueryAuthProfilesInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryAuthProfilesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryAuthProfilesGatewayTimeout creates a QueryAuthProfilesGatewayTimeout with default headers values
func NewQueryAuthProfilesGatewayTimeout() *QueryAuthProfilesGatewayTimeout {
	return &QueryAuthProfilesGatewayTimeout{}
}

/* QueryAuthProfilesGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type QueryAuthProfilesGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryAuthProfilesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/authorization/profiles][%d] queryAuthProfilesGatewayTimeout  %+v", 504, o.Payload)
}
func (o *QueryAuthProfilesGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryAuthProfilesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
