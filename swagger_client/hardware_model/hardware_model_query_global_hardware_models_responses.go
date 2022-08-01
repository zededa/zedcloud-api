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

// HardwareModelQueryGlobalHardwareModelsReader is a Reader for the HardwareModelQueryGlobalHardwareModels structure.
type HardwareModelQueryGlobalHardwareModelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HardwareModelQueryGlobalHardwareModelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHardwareModelQueryGlobalHardwareModelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHardwareModelQueryGlobalHardwareModelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewHardwareModelQueryGlobalHardwareModelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHardwareModelQueryGlobalHardwareModelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHardwareModelQueryGlobalHardwareModelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewHardwareModelQueryGlobalHardwareModelsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewHardwareModelQueryGlobalHardwareModelsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHardwareModelQueryGlobalHardwareModelsOK creates a HardwareModelQueryGlobalHardwareModelsOK with default headers values
func NewHardwareModelQueryGlobalHardwareModelsOK() *HardwareModelQueryGlobalHardwareModelsOK {
	return &HardwareModelQueryGlobalHardwareModelsOK{}
}

/* HardwareModelQueryGlobalHardwareModelsOK describes a response with status code 200, with default header values.

A successful response.
*/
type HardwareModelQueryGlobalHardwareModelsOK struct {
	Payload *swagger_models.SysModels
}

func (o *HardwareModelQueryGlobalHardwareModelsOK) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global][%d] hardwareModelQueryGlobalHardwareModelsOK  %+v", 200, o.Payload)
}
func (o *HardwareModelQueryGlobalHardwareModelsOK) GetPayload() *swagger_models.SysModels {
	return o.Payload
}

func (o *HardwareModelQueryGlobalHardwareModelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.SysModels)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelQueryGlobalHardwareModelsBadRequest creates a HardwareModelQueryGlobalHardwareModelsBadRequest with default headers values
func NewHardwareModelQueryGlobalHardwareModelsBadRequest() *HardwareModelQueryGlobalHardwareModelsBadRequest {
	return &HardwareModelQueryGlobalHardwareModelsBadRequest{}
}

/* HardwareModelQueryGlobalHardwareModelsBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type HardwareModelQueryGlobalHardwareModelsBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *HardwareModelQueryGlobalHardwareModelsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global][%d] hardwareModelQueryGlobalHardwareModelsBadRequest  %+v", 400, o.Payload)
}
func (o *HardwareModelQueryGlobalHardwareModelsBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelQueryGlobalHardwareModelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelQueryGlobalHardwareModelsUnauthorized creates a HardwareModelQueryGlobalHardwareModelsUnauthorized with default headers values
func NewHardwareModelQueryGlobalHardwareModelsUnauthorized() *HardwareModelQueryGlobalHardwareModelsUnauthorized {
	return &HardwareModelQueryGlobalHardwareModelsUnauthorized{}
}

/* HardwareModelQueryGlobalHardwareModelsUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type HardwareModelQueryGlobalHardwareModelsUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *HardwareModelQueryGlobalHardwareModelsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global][%d] hardwareModelQueryGlobalHardwareModelsUnauthorized  %+v", 401, o.Payload)
}
func (o *HardwareModelQueryGlobalHardwareModelsUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelQueryGlobalHardwareModelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelQueryGlobalHardwareModelsForbidden creates a HardwareModelQueryGlobalHardwareModelsForbidden with default headers values
func NewHardwareModelQueryGlobalHardwareModelsForbidden() *HardwareModelQueryGlobalHardwareModelsForbidden {
	return &HardwareModelQueryGlobalHardwareModelsForbidden{}
}

/* HardwareModelQueryGlobalHardwareModelsForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type HardwareModelQueryGlobalHardwareModelsForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *HardwareModelQueryGlobalHardwareModelsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global][%d] hardwareModelQueryGlobalHardwareModelsForbidden  %+v", 403, o.Payload)
}
func (o *HardwareModelQueryGlobalHardwareModelsForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelQueryGlobalHardwareModelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelQueryGlobalHardwareModelsInternalServerError creates a HardwareModelQueryGlobalHardwareModelsInternalServerError with default headers values
func NewHardwareModelQueryGlobalHardwareModelsInternalServerError() *HardwareModelQueryGlobalHardwareModelsInternalServerError {
	return &HardwareModelQueryGlobalHardwareModelsInternalServerError{}
}

/* HardwareModelQueryGlobalHardwareModelsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type HardwareModelQueryGlobalHardwareModelsInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *HardwareModelQueryGlobalHardwareModelsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global][%d] hardwareModelQueryGlobalHardwareModelsInternalServerError  %+v", 500, o.Payload)
}
func (o *HardwareModelQueryGlobalHardwareModelsInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelQueryGlobalHardwareModelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelQueryGlobalHardwareModelsGatewayTimeout creates a HardwareModelQueryGlobalHardwareModelsGatewayTimeout with default headers values
func NewHardwareModelQueryGlobalHardwareModelsGatewayTimeout() *HardwareModelQueryGlobalHardwareModelsGatewayTimeout {
	return &HardwareModelQueryGlobalHardwareModelsGatewayTimeout{}
}

/* HardwareModelQueryGlobalHardwareModelsGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type HardwareModelQueryGlobalHardwareModelsGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *HardwareModelQueryGlobalHardwareModelsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global][%d] hardwareModelQueryGlobalHardwareModelsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *HardwareModelQueryGlobalHardwareModelsGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelQueryGlobalHardwareModelsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelQueryGlobalHardwareModelsDefault creates a HardwareModelQueryGlobalHardwareModelsDefault with default headers values
func NewHardwareModelQueryGlobalHardwareModelsDefault(code int) *HardwareModelQueryGlobalHardwareModelsDefault {
	return &HardwareModelQueryGlobalHardwareModelsDefault{
		_statusCode: code,
	}
}

/* HardwareModelQueryGlobalHardwareModelsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HardwareModelQueryGlobalHardwareModelsDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the hardware model query global hardware models default response
func (o *HardwareModelQueryGlobalHardwareModelsDefault) Code() int {
	return o._statusCode
}

func (o *HardwareModelQueryGlobalHardwareModelsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global][%d] HardwareModel_QueryGlobalHardwareModels default  %+v", o._statusCode, o.Payload)
}
func (o *HardwareModelQueryGlobalHardwareModelsDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *HardwareModelQueryGlobalHardwareModelsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
