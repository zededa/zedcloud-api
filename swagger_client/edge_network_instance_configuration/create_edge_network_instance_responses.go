// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_network_instance_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// CreateEdgeNetworkInstanceReader is a Reader for the CreateEdgeNetworkInstance structure.
type CreateEdgeNetworkInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateEdgeNetworkInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateEdgeNetworkInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateEdgeNetworkInstanceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateEdgeNetworkInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateEdgeNetworkInstanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateEdgeNetworkInstanceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateEdgeNetworkInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewCreateEdgeNetworkInstanceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateEdgeNetworkInstanceOK creates a CreateEdgeNetworkInstanceOK with default headers values
func NewCreateEdgeNetworkInstanceOK() *CreateEdgeNetworkInstanceOK {
	return &CreateEdgeNetworkInstanceOK{}
}

/* CreateEdgeNetworkInstanceOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateEdgeNetworkInstanceOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeNetworkInstanceOK) Error() string {
	return fmt.Sprintf("[POST /v1/netinsts][%d] createEdgeNetworkInstanceOK  %+v", 200, o.Payload)
}
func (o *CreateEdgeNetworkInstanceOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeNetworkInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEdgeNetworkInstanceBadRequest creates a CreateEdgeNetworkInstanceBadRequest with default headers values
func NewCreateEdgeNetworkInstanceBadRequest() *CreateEdgeNetworkInstanceBadRequest {
	return &CreateEdgeNetworkInstanceBadRequest{}
}

/* CreateEdgeNetworkInstanceBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type CreateEdgeNetworkInstanceBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeNetworkInstanceBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/netinsts][%d] createEdgeNetworkInstanceBadRequest  %+v", 400, o.Payload)
}
func (o *CreateEdgeNetworkInstanceBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeNetworkInstanceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEdgeNetworkInstanceUnauthorized creates a CreateEdgeNetworkInstanceUnauthorized with default headers values
func NewCreateEdgeNetworkInstanceUnauthorized() *CreateEdgeNetworkInstanceUnauthorized {
	return &CreateEdgeNetworkInstanceUnauthorized{}
}

/* CreateEdgeNetworkInstanceUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type CreateEdgeNetworkInstanceUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeNetworkInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/netinsts][%d] createEdgeNetworkInstanceUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateEdgeNetworkInstanceUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeNetworkInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEdgeNetworkInstanceForbidden creates a CreateEdgeNetworkInstanceForbidden with default headers values
func NewCreateEdgeNetworkInstanceForbidden() *CreateEdgeNetworkInstanceForbidden {
	return &CreateEdgeNetworkInstanceForbidden{}
}

/* CreateEdgeNetworkInstanceForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type CreateEdgeNetworkInstanceForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeNetworkInstanceForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/netinsts][%d] createEdgeNetworkInstanceForbidden  %+v", 403, o.Payload)
}
func (o *CreateEdgeNetworkInstanceForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeNetworkInstanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEdgeNetworkInstanceConflict creates a CreateEdgeNetworkInstanceConflict with default headers values
func NewCreateEdgeNetworkInstanceConflict() *CreateEdgeNetworkInstanceConflict {
	return &CreateEdgeNetworkInstanceConflict{}
}

/* CreateEdgeNetworkInstanceConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this edge network record will conflict with an already existing edge network record.
*/
type CreateEdgeNetworkInstanceConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeNetworkInstanceConflict) Error() string {
	return fmt.Sprintf("[POST /v1/netinsts][%d] createEdgeNetworkInstanceConflict  %+v", 409, o.Payload)
}
func (o *CreateEdgeNetworkInstanceConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeNetworkInstanceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEdgeNetworkInstanceInternalServerError creates a CreateEdgeNetworkInstanceInternalServerError with default headers values
func NewCreateEdgeNetworkInstanceInternalServerError() *CreateEdgeNetworkInstanceInternalServerError {
	return &CreateEdgeNetworkInstanceInternalServerError{}
}

/* CreateEdgeNetworkInstanceInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type CreateEdgeNetworkInstanceInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeNetworkInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/netinsts][%d] createEdgeNetworkInstanceInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateEdgeNetworkInstanceInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeNetworkInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEdgeNetworkInstanceGatewayTimeout creates a CreateEdgeNetworkInstanceGatewayTimeout with default headers values
func NewCreateEdgeNetworkInstanceGatewayTimeout() *CreateEdgeNetworkInstanceGatewayTimeout {
	return &CreateEdgeNetworkInstanceGatewayTimeout{}
}

/* CreateEdgeNetworkInstanceGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type CreateEdgeNetworkInstanceGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeNetworkInstanceGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/netinsts][%d] createEdgeNetworkInstanceGatewayTimeout  %+v", 504, o.Payload)
}
func (o *CreateEdgeNetworkInstanceGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeNetworkInstanceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
