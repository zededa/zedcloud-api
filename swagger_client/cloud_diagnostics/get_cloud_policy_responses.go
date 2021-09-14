// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package cloud_diagnostics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// GetCloudPolicyReader is a Reader for the GetCloudPolicy structure.
type GetCloudPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCloudPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetCloudPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCloudPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCloudPolicyGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCloudPolicyOK creates a GetCloudPolicyOK with default headers values
func NewGetCloudPolicyOK() *GetCloudPolicyOK {
	return &GetCloudPolicyOK{}
}

/* GetCloudPolicyOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetCloudPolicyOK struct {
	Payload *swagger_models.PolicyDocVersionResp
}

func (o *GetCloudPolicyOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/name/{policy}][%d] getCloudPolicyOK  %+v", 200, o.Payload)
}
func (o *GetCloudPolicyOK) GetPayload() *swagger_models.PolicyDocVersionResp {
	return o.Payload
}

func (o *GetCloudPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.PolicyDocVersionResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudPolicyNotFound creates a GetCloudPolicyNotFound with default headers values
func NewGetCloudPolicyNotFound() *GetCloudPolicyNotFound {
	return &GetCloudPolicyNotFound{}
}

/* GetCloudPolicyNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetCloudPolicyNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetCloudPolicyNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/name/{policy}][%d] getCloudPolicyNotFound  %+v", 404, o.Payload)
}
func (o *GetCloudPolicyNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetCloudPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudPolicyInternalServerError creates a GetCloudPolicyInternalServerError with default headers values
func NewGetCloudPolicyInternalServerError() *GetCloudPolicyInternalServerError {
	return &GetCloudPolicyInternalServerError{}
}

/* GetCloudPolicyInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetCloudPolicyInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetCloudPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/name/{policy}][%d] getCloudPolicyInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCloudPolicyInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetCloudPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudPolicyGatewayTimeout creates a GetCloudPolicyGatewayTimeout with default headers values
func NewGetCloudPolicyGatewayTimeout() *GetCloudPolicyGatewayTimeout {
	return &GetCloudPolicyGatewayTimeout{}
}

/* GetCloudPolicyGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetCloudPolicyGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetCloudPolicyGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/cloud/policies/name/{policy}][%d] getCloudPolicyGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCloudPolicyGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetCloudPolicyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
