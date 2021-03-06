// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_node_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// GetEdgeNodeOnboardingReader is a Reader for the GetEdgeNodeOnboarding structure.
type GetEdgeNodeOnboardingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEdgeNodeOnboardingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEdgeNodeOnboardingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEdgeNodeOnboardingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEdgeNodeOnboardingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEdgeNodeOnboardingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEdgeNodeOnboardingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetEdgeNodeOnboardingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEdgeNodeOnboardingOK creates a GetEdgeNodeOnboardingOK with default headers values
func NewGetEdgeNodeOnboardingOK() *GetEdgeNodeOnboardingOK {
	return &GetEdgeNodeOnboardingOK{}
}

/* GetEdgeNodeOnboardingOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetEdgeNodeOnboardingOK struct {
	Payload *swagger_models.DeviceConfig
}

func (o *GetEdgeNodeOnboardingOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/onboarding][%d] getEdgeNodeOnboardingOK  %+v", 200, o.Payload)
}
func (o *GetEdgeNodeOnboardingOK) GetPayload() *swagger_models.DeviceConfig {
	return o.Payload
}

func (o *GetEdgeNodeOnboardingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.DeviceConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNodeOnboardingUnauthorized creates a GetEdgeNodeOnboardingUnauthorized with default headers values
func NewGetEdgeNodeOnboardingUnauthorized() *GetEdgeNodeOnboardingUnauthorized {
	return &GetEdgeNodeOnboardingUnauthorized{}
}

/* GetEdgeNodeOnboardingUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetEdgeNodeOnboardingUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNodeOnboardingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/onboarding][%d] getEdgeNodeOnboardingUnauthorized  %+v", 401, o.Payload)
}
func (o *GetEdgeNodeOnboardingUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNodeOnboardingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNodeOnboardingForbidden creates a GetEdgeNodeOnboardingForbidden with default headers values
func NewGetEdgeNodeOnboardingForbidden() *GetEdgeNodeOnboardingForbidden {
	return &GetEdgeNodeOnboardingForbidden{}
}

/* GetEdgeNodeOnboardingForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type GetEdgeNodeOnboardingForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNodeOnboardingForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/onboarding][%d] getEdgeNodeOnboardingForbidden  %+v", 403, o.Payload)
}
func (o *GetEdgeNodeOnboardingForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNodeOnboardingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNodeOnboardingNotFound creates a GetEdgeNodeOnboardingNotFound with default headers values
func NewGetEdgeNodeOnboardingNotFound() *GetEdgeNodeOnboardingNotFound {
	return &GetEdgeNodeOnboardingNotFound{}
}

/* GetEdgeNodeOnboardingNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetEdgeNodeOnboardingNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNodeOnboardingNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/onboarding][%d] getEdgeNodeOnboardingNotFound  %+v", 404, o.Payload)
}
func (o *GetEdgeNodeOnboardingNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNodeOnboardingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNodeOnboardingInternalServerError creates a GetEdgeNodeOnboardingInternalServerError with default headers values
func NewGetEdgeNodeOnboardingInternalServerError() *GetEdgeNodeOnboardingInternalServerError {
	return &GetEdgeNodeOnboardingInternalServerError{}
}

/* GetEdgeNodeOnboardingInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetEdgeNodeOnboardingInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNodeOnboardingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/onboarding][%d] getEdgeNodeOnboardingInternalServerError  %+v", 500, o.Payload)
}
func (o *GetEdgeNodeOnboardingInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNodeOnboardingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEdgeNodeOnboardingGatewayTimeout creates a GetEdgeNodeOnboardingGatewayTimeout with default headers values
func NewGetEdgeNodeOnboardingGatewayTimeout() *GetEdgeNodeOnboardingGatewayTimeout {
	return &GetEdgeNodeOnboardingGatewayTimeout{}
}

/* GetEdgeNodeOnboardingGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetEdgeNodeOnboardingGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetEdgeNodeOnboardingGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}/onboarding][%d] getEdgeNodeOnboardingGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetEdgeNodeOnboardingGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetEdgeNodeOnboardingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
