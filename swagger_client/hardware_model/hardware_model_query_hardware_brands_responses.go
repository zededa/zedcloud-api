// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package hardware_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// HardwareModelQueryHardwareBrandsReader is a Reader for the HardwareModelQueryHardwareBrands structure.
type HardwareModelQueryHardwareBrandsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HardwareModelQueryHardwareBrandsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHardwareModelQueryHardwareBrandsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHardwareModelQueryHardwareBrandsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewHardwareModelQueryHardwareBrandsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHardwareModelQueryHardwareBrandsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHardwareModelQueryHardwareBrandsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewHardwareModelQueryHardwareBrandsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewHardwareModelQueryHardwareBrandsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHardwareModelQueryHardwareBrandsOK creates a HardwareModelQueryHardwareBrandsOK with default headers values
func NewHardwareModelQueryHardwareBrandsOK() *HardwareModelQueryHardwareBrandsOK {
	return &HardwareModelQueryHardwareBrandsOK{}
}

/* HardwareModelQueryHardwareBrandsOK describes a response with status code 200, with default header values.

A successful response.
*/
type HardwareModelQueryHardwareBrandsOK struct {
	Payload *swagger_models.SysBrands
}

func (o *HardwareModelQueryHardwareBrandsOK) Error() string {
	return fmt.Sprintf("[GET /v1/brands][%d] hardwareModelQueryHardwareBrandsOK  %+v", 200, o.Payload)
}
func (o *HardwareModelQueryHardwareBrandsOK) GetPayload() *swagger_models.SysBrands {
	return o.Payload
}

func (o *HardwareModelQueryHardwareBrandsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.SysBrands)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelQueryHardwareBrandsBadRequest creates a HardwareModelQueryHardwareBrandsBadRequest with default headers values
func NewHardwareModelQueryHardwareBrandsBadRequest() *HardwareModelQueryHardwareBrandsBadRequest {
	return &HardwareModelQueryHardwareBrandsBadRequest{}
}

/* HardwareModelQueryHardwareBrandsBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type HardwareModelQueryHardwareBrandsBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *HardwareModelQueryHardwareBrandsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/brands][%d] hardwareModelQueryHardwareBrandsBadRequest  %+v", 400, o.Payload)
}
func (o *HardwareModelQueryHardwareBrandsBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelQueryHardwareBrandsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelQueryHardwareBrandsUnauthorized creates a HardwareModelQueryHardwareBrandsUnauthorized with default headers values
func NewHardwareModelQueryHardwareBrandsUnauthorized() *HardwareModelQueryHardwareBrandsUnauthorized {
	return &HardwareModelQueryHardwareBrandsUnauthorized{}
}

/* HardwareModelQueryHardwareBrandsUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type HardwareModelQueryHardwareBrandsUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *HardwareModelQueryHardwareBrandsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/brands][%d] hardwareModelQueryHardwareBrandsUnauthorized  %+v", 401, o.Payload)
}
func (o *HardwareModelQueryHardwareBrandsUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelQueryHardwareBrandsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelQueryHardwareBrandsForbidden creates a HardwareModelQueryHardwareBrandsForbidden with default headers values
func NewHardwareModelQueryHardwareBrandsForbidden() *HardwareModelQueryHardwareBrandsForbidden {
	return &HardwareModelQueryHardwareBrandsForbidden{}
}

/* HardwareModelQueryHardwareBrandsForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type HardwareModelQueryHardwareBrandsForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *HardwareModelQueryHardwareBrandsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/brands][%d] hardwareModelQueryHardwareBrandsForbidden  %+v", 403, o.Payload)
}
func (o *HardwareModelQueryHardwareBrandsForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelQueryHardwareBrandsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelQueryHardwareBrandsInternalServerError creates a HardwareModelQueryHardwareBrandsInternalServerError with default headers values
func NewHardwareModelQueryHardwareBrandsInternalServerError() *HardwareModelQueryHardwareBrandsInternalServerError {
	return &HardwareModelQueryHardwareBrandsInternalServerError{}
}

/* HardwareModelQueryHardwareBrandsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type HardwareModelQueryHardwareBrandsInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *HardwareModelQueryHardwareBrandsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/brands][%d] hardwareModelQueryHardwareBrandsInternalServerError  %+v", 500, o.Payload)
}
func (o *HardwareModelQueryHardwareBrandsInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelQueryHardwareBrandsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelQueryHardwareBrandsGatewayTimeout creates a HardwareModelQueryHardwareBrandsGatewayTimeout with default headers values
func NewHardwareModelQueryHardwareBrandsGatewayTimeout() *HardwareModelQueryHardwareBrandsGatewayTimeout {
	return &HardwareModelQueryHardwareBrandsGatewayTimeout{}
}

/* HardwareModelQueryHardwareBrandsGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type HardwareModelQueryHardwareBrandsGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *HardwareModelQueryHardwareBrandsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/brands][%d] hardwareModelQueryHardwareBrandsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *HardwareModelQueryHardwareBrandsGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelQueryHardwareBrandsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelQueryHardwareBrandsDefault creates a HardwareModelQueryHardwareBrandsDefault with default headers values
func NewHardwareModelQueryHardwareBrandsDefault(code int) *HardwareModelQueryHardwareBrandsDefault {
	return &HardwareModelQueryHardwareBrandsDefault{
		_statusCode: code,
	}
}

/* HardwareModelQueryHardwareBrandsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HardwareModelQueryHardwareBrandsDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the hardware model query hardware brands default response
func (o *HardwareModelQueryHardwareBrandsDefault) Code() int {
	return o._statusCode
}

func (o *HardwareModelQueryHardwareBrandsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/brands][%d] HardwareModel_QueryHardwareBrands default  %+v", o._statusCode, o.Payload)
}
func (o *HardwareModelQueryHardwareBrandsDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *HardwareModelQueryHardwareBrandsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
