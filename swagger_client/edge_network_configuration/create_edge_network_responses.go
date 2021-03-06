// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_network_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// CreateEdgeNetworkReader is a Reader for the CreateEdgeNetwork structure.
type CreateEdgeNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateEdgeNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateEdgeNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateEdgeNetworkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateEdgeNetworkUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateEdgeNetworkForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateEdgeNetworkConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateEdgeNetworkInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewCreateEdgeNetworkGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateEdgeNetworkOK creates a CreateEdgeNetworkOK with default headers values
func NewCreateEdgeNetworkOK() *CreateEdgeNetworkOK {
	return &CreateEdgeNetworkOK{}
}

/* CreateEdgeNetworkOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateEdgeNetworkOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeNetworkOK) Error() string {
	return fmt.Sprintf("[POST /v1/networks][%d] createEdgeNetworkOK  %+v", 200, o.Payload)
}
func (o *CreateEdgeNetworkOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEdgeNetworkBadRequest creates a CreateEdgeNetworkBadRequest with default headers values
func NewCreateEdgeNetworkBadRequest() *CreateEdgeNetworkBadRequest {
	return &CreateEdgeNetworkBadRequest{}
}

/* CreateEdgeNetworkBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type CreateEdgeNetworkBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeNetworkBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/networks][%d] createEdgeNetworkBadRequest  %+v", 400, o.Payload)
}
func (o *CreateEdgeNetworkBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeNetworkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEdgeNetworkUnauthorized creates a CreateEdgeNetworkUnauthorized with default headers values
func NewCreateEdgeNetworkUnauthorized() *CreateEdgeNetworkUnauthorized {
	return &CreateEdgeNetworkUnauthorized{}
}

/* CreateEdgeNetworkUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type CreateEdgeNetworkUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeNetworkUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/networks][%d] createEdgeNetworkUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateEdgeNetworkUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeNetworkUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEdgeNetworkForbidden creates a CreateEdgeNetworkForbidden with default headers values
func NewCreateEdgeNetworkForbidden() *CreateEdgeNetworkForbidden {
	return &CreateEdgeNetworkForbidden{}
}

/* CreateEdgeNetworkForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type CreateEdgeNetworkForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeNetworkForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/networks][%d] createEdgeNetworkForbidden  %+v", 403, o.Payload)
}
func (o *CreateEdgeNetworkForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeNetworkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEdgeNetworkConflict creates a CreateEdgeNetworkConflict with default headers values
func NewCreateEdgeNetworkConflict() *CreateEdgeNetworkConflict {
	return &CreateEdgeNetworkConflict{}
}

/* CreateEdgeNetworkConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this edge network record will conflict with an already existing edge network record.
*/
type CreateEdgeNetworkConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeNetworkConflict) Error() string {
	return fmt.Sprintf("[POST /v1/networks][%d] createEdgeNetworkConflict  %+v", 409, o.Payload)
}
func (o *CreateEdgeNetworkConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeNetworkConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEdgeNetworkInternalServerError creates a CreateEdgeNetworkInternalServerError with default headers values
func NewCreateEdgeNetworkInternalServerError() *CreateEdgeNetworkInternalServerError {
	return &CreateEdgeNetworkInternalServerError{}
}

/* CreateEdgeNetworkInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type CreateEdgeNetworkInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeNetworkInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/networks][%d] createEdgeNetworkInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateEdgeNetworkInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeNetworkInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEdgeNetworkGatewayTimeout creates a CreateEdgeNetworkGatewayTimeout with default headers values
func NewCreateEdgeNetworkGatewayTimeout() *CreateEdgeNetworkGatewayTimeout {
	return &CreateEdgeNetworkGatewayTimeout{}
}

/* CreateEdgeNetworkGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type CreateEdgeNetworkGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeNetworkGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/networks][%d] createEdgeNetworkGatewayTimeout  %+v", 504, o.Payload)
}
func (o *CreateEdgeNetworkGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeNetworkGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
