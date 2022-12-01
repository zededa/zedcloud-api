// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package edge_network_instance_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsReader is a Reader for the EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTags structure.
type EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsOK creates a EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsOK with default headers values
func NewEdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsOK() *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsOK {
	return &EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsOK{}
}

/*
EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsOK struct {
	Payload *swagger_models.ObjectTagsList
}

// IsSuccess returns true when this edge network instance status get network instance interface tags o k response has a 2xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge network instance status get network instance interface tags o k response has a 3xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance status get network instance interface tags o k response has a 4xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network instance status get network instance interface tags o k response has a 5xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network instance status get network instance interface tags o k response a status code equal to that given
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsOK) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/tags][%d] edgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsOK  %+v", 200, o.Payload)
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsOK) String() string {
	return fmt.Sprintf("[GET /v1/netinsts/tags][%d] edgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsOK  %+v", 200, o.Payload)
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsOK) GetPayload() *swagger_models.ObjectTagsList {
	return o.Payload
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ObjectTagsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsBadRequest creates a EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsBadRequest with default headers values
func NewEdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsBadRequest() *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsBadRequest {
	return &EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsBadRequest{}
}

/*
EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance status get network instance interface tags bad request response has a 2xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network instance status get network instance interface tags bad request response has a 3xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance status get network instance interface tags bad request response has a 4xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network instance status get network instance interface tags bad request response has a 5xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network instance status get network instance interface tags bad request response a status code equal to that given
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/tags][%d] edgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsBadRequest  %+v", 400, o.Payload)
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/netinsts/tags][%d] edgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsBadRequest  %+v", 400, o.Payload)
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsUnauthorized creates a EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsUnauthorized with default headers values
func NewEdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsUnauthorized() *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsUnauthorized {
	return &EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsUnauthorized{}
}

/*
EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance status get network instance interface tags unauthorized response has a 2xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network instance status get network instance interface tags unauthorized response has a 3xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance status get network instance interface tags unauthorized response has a 4xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network instance status get network instance interface tags unauthorized response has a 5xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network instance status get network instance interface tags unauthorized response a status code equal to that given
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/tags][%d] edgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/netinsts/tags][%d] edgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsForbidden creates a EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsForbidden with default headers values
func NewEdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsForbidden() *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsForbidden {
	return &EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsForbidden{}
}

/*
EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance status get network instance interface tags forbidden response has a 2xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network instance status get network instance interface tags forbidden response has a 3xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance status get network instance interface tags forbidden response has a 4xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge network instance status get network instance interface tags forbidden response has a 5xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge network instance status get network instance interface tags forbidden response a status code equal to that given
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/tags][%d] edgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsForbidden) String() string {
	return fmt.Sprintf("[GET /v1/netinsts/tags][%d] edgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsInternalServerError creates a EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsInternalServerError with default headers values
func NewEdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsInternalServerError() *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsInternalServerError {
	return &EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsInternalServerError{}
}

/*
EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance status get network instance interface tags internal server error response has a 2xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network instance status get network instance interface tags internal server error response has a 3xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance status get network instance interface tags internal server error response has a 4xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network instance status get network instance interface tags internal server error response has a 5xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge network instance status get network instance interface tags internal server error response a status code equal to that given
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/tags][%d] edgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/netinsts/tags][%d] edgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsGatewayTimeout creates a EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsGatewayTimeout with default headers values
func NewEdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsGatewayTimeout() *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsGatewayTimeout {
	return &EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsGatewayTimeout{}
}

/*
EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge network instance status get network instance interface tags gateway timeout response has a 2xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge network instance status get network instance interface tags gateway timeout response has a 3xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge network instance status get network instance interface tags gateway timeout response has a 4xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge network instance status get network instance interface tags gateway timeout response has a 5xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge network instance status get network instance interface tags gateway timeout response a status code equal to that given
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/tags][%d] edgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/netinsts/tags][%d] edgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsDefault creates a EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsDefault with default headers values
func NewEdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsDefault(code int) *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsDefault {
	return &EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsDefault{
		_statusCode: code,
	}
}

/*
EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the edge network instance status get network instance interface tags default response
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge network instance status get network instance interface tags default response has a 2xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge network instance status get network instance interface tags default response has a 3xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge network instance status get network instance interface tags default response has a 4xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge network instance status get network instance interface tags default response has a 5xx status code
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge network instance status get network instance interface tags default response a status code equal to that given
func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/netinsts/tags][%d] EdgeNetworkInstanceStatus_GetNetworkInstanceInterfaceTags default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsDefault) String() string {
	return fmt.Sprintf("[GET /v1/netinsts/tags][%d] EdgeNetworkInstanceStatus_GetNetworkInstanceInterfaceTags default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNetworkInstanceStatusGetNetworkInstanceInterfaceTagsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}