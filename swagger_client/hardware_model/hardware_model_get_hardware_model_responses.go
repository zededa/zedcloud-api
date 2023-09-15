// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
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

// HardwareModelGetHardwareModelReader is a Reader for the HardwareModelGetHardwareModel structure.
type HardwareModelGetHardwareModelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HardwareModelGetHardwareModelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHardwareModelGetHardwareModelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewHardwareModelGetHardwareModelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHardwareModelGetHardwareModelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHardwareModelGetHardwareModelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHardwareModelGetHardwareModelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewHardwareModelGetHardwareModelGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewHardwareModelGetHardwareModelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHardwareModelGetHardwareModelOK creates a HardwareModelGetHardwareModelOK with default headers values
func NewHardwareModelGetHardwareModelOK() *HardwareModelGetHardwareModelOK {
	return &HardwareModelGetHardwareModelOK{}
}

/*
HardwareModelGetHardwareModelOK describes a response with status code 200, with default header values.

A successful response.
*/
type HardwareModelGetHardwareModelOK struct {
	Payload *swagger_models.SysModel
}

// IsSuccess returns true when this hardware model get hardware model o k response has a 2xx status code
func (o *HardwareModelGetHardwareModelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hardware model get hardware model o k response has a 3xx status code
func (o *HardwareModelGetHardwareModelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get hardware model o k response has a 4xx status code
func (o *HardwareModelGetHardwareModelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get hardware model o k response has a 5xx status code
func (o *HardwareModelGetHardwareModelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get hardware model o k response a status code equal to that given
func (o *HardwareModelGetHardwareModelOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the hardware model get hardware model o k response
func (o *HardwareModelGetHardwareModelOK) Code() int {
	return 200
}

func (o *HardwareModelGetHardwareModelOK) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/id/{id}][%d] hardwareModelGetHardwareModelOK  %+v", 200, o.Payload)
}

func (o *HardwareModelGetHardwareModelOK) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/id/{id}][%d] hardwareModelGetHardwareModelOK  %+v", 200, o.Payload)
}

func (o *HardwareModelGetHardwareModelOK) GetPayload() *swagger_models.SysModel {
	return o.Payload
}

func (o *HardwareModelGetHardwareModelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.SysModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareModelUnauthorized creates a HardwareModelGetHardwareModelUnauthorized with default headers values
func NewHardwareModelGetHardwareModelUnauthorized() *HardwareModelGetHardwareModelUnauthorized {
	return &HardwareModelGetHardwareModelUnauthorized{}
}

/*
HardwareModelGetHardwareModelUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type HardwareModelGetHardwareModelUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get hardware model unauthorized response has a 2xx status code
func (o *HardwareModelGetHardwareModelUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get hardware model unauthorized response has a 3xx status code
func (o *HardwareModelGetHardwareModelUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get hardware model unauthorized response has a 4xx status code
func (o *HardwareModelGetHardwareModelUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get hardware model unauthorized response has a 5xx status code
func (o *HardwareModelGetHardwareModelUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get hardware model unauthorized response a status code equal to that given
func (o *HardwareModelGetHardwareModelUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the hardware model get hardware model unauthorized response
func (o *HardwareModelGetHardwareModelUnauthorized) Code() int {
	return 401
}

func (o *HardwareModelGetHardwareModelUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/id/{id}][%d] hardwareModelGetHardwareModelUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelGetHardwareModelUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/id/{id}][%d] hardwareModelGetHardwareModelUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelGetHardwareModelUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetHardwareModelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareModelForbidden creates a HardwareModelGetHardwareModelForbidden with default headers values
func NewHardwareModelGetHardwareModelForbidden() *HardwareModelGetHardwareModelForbidden {
	return &HardwareModelGetHardwareModelForbidden{}
}

/*
HardwareModelGetHardwareModelForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type HardwareModelGetHardwareModelForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get hardware model forbidden response has a 2xx status code
func (o *HardwareModelGetHardwareModelForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get hardware model forbidden response has a 3xx status code
func (o *HardwareModelGetHardwareModelForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get hardware model forbidden response has a 4xx status code
func (o *HardwareModelGetHardwareModelForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get hardware model forbidden response has a 5xx status code
func (o *HardwareModelGetHardwareModelForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get hardware model forbidden response a status code equal to that given
func (o *HardwareModelGetHardwareModelForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the hardware model get hardware model forbidden response
func (o *HardwareModelGetHardwareModelForbidden) Code() int {
	return 403
}

func (o *HardwareModelGetHardwareModelForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/id/{id}][%d] hardwareModelGetHardwareModelForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelGetHardwareModelForbidden) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/id/{id}][%d] hardwareModelGetHardwareModelForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelGetHardwareModelForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetHardwareModelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareModelNotFound creates a HardwareModelGetHardwareModelNotFound with default headers values
func NewHardwareModelGetHardwareModelNotFound() *HardwareModelGetHardwareModelNotFound {
	return &HardwareModelGetHardwareModelNotFound{}
}

/*
HardwareModelGetHardwareModelNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type HardwareModelGetHardwareModelNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get hardware model not found response has a 2xx status code
func (o *HardwareModelGetHardwareModelNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get hardware model not found response has a 3xx status code
func (o *HardwareModelGetHardwareModelNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get hardware model not found response has a 4xx status code
func (o *HardwareModelGetHardwareModelNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get hardware model not found response has a 5xx status code
func (o *HardwareModelGetHardwareModelNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get hardware model not found response a status code equal to that given
func (o *HardwareModelGetHardwareModelNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the hardware model get hardware model not found response
func (o *HardwareModelGetHardwareModelNotFound) Code() int {
	return 404
}

func (o *HardwareModelGetHardwareModelNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/id/{id}][%d] hardwareModelGetHardwareModelNotFound  %+v", 404, o.Payload)
}

func (o *HardwareModelGetHardwareModelNotFound) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/id/{id}][%d] hardwareModelGetHardwareModelNotFound  %+v", 404, o.Payload)
}

func (o *HardwareModelGetHardwareModelNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetHardwareModelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareModelInternalServerError creates a HardwareModelGetHardwareModelInternalServerError with default headers values
func NewHardwareModelGetHardwareModelInternalServerError() *HardwareModelGetHardwareModelInternalServerError {
	return &HardwareModelGetHardwareModelInternalServerError{}
}

/*
HardwareModelGetHardwareModelInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type HardwareModelGetHardwareModelInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get hardware model internal server error response has a 2xx status code
func (o *HardwareModelGetHardwareModelInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get hardware model internal server error response has a 3xx status code
func (o *HardwareModelGetHardwareModelInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get hardware model internal server error response has a 4xx status code
func (o *HardwareModelGetHardwareModelInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get hardware model internal server error response has a 5xx status code
func (o *HardwareModelGetHardwareModelInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model get hardware model internal server error response a status code equal to that given
func (o *HardwareModelGetHardwareModelInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the hardware model get hardware model internal server error response
func (o *HardwareModelGetHardwareModelInternalServerError) Code() int {
	return 500
}

func (o *HardwareModelGetHardwareModelInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/id/{id}][%d] hardwareModelGetHardwareModelInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelGetHardwareModelInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/id/{id}][%d] hardwareModelGetHardwareModelInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelGetHardwareModelInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetHardwareModelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareModelGatewayTimeout creates a HardwareModelGetHardwareModelGatewayTimeout with default headers values
func NewHardwareModelGetHardwareModelGatewayTimeout() *HardwareModelGetHardwareModelGatewayTimeout {
	return &HardwareModelGetHardwareModelGatewayTimeout{}
}

/*
HardwareModelGetHardwareModelGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type HardwareModelGetHardwareModelGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get hardware model gateway timeout response has a 2xx status code
func (o *HardwareModelGetHardwareModelGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get hardware model gateway timeout response has a 3xx status code
func (o *HardwareModelGetHardwareModelGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get hardware model gateway timeout response has a 4xx status code
func (o *HardwareModelGetHardwareModelGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get hardware model gateway timeout response has a 5xx status code
func (o *HardwareModelGetHardwareModelGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model get hardware model gateway timeout response a status code equal to that given
func (o *HardwareModelGetHardwareModelGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the hardware model get hardware model gateway timeout response
func (o *HardwareModelGetHardwareModelGatewayTimeout) Code() int {
	return 504
}

func (o *HardwareModelGetHardwareModelGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/id/{id}][%d] hardwareModelGetHardwareModelGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelGetHardwareModelGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/id/{id}][%d] hardwareModelGetHardwareModelGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelGetHardwareModelGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetHardwareModelGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareModelDefault creates a HardwareModelGetHardwareModelDefault with default headers values
func NewHardwareModelGetHardwareModelDefault(code int) *HardwareModelGetHardwareModelDefault {
	return &HardwareModelGetHardwareModelDefault{
		_statusCode: code,
	}
}

/*
HardwareModelGetHardwareModelDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HardwareModelGetHardwareModelDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// IsSuccess returns true when this hardware model get hardware model default response has a 2xx status code
func (o *HardwareModelGetHardwareModelDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this hardware model get hardware model default response has a 3xx status code
func (o *HardwareModelGetHardwareModelDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this hardware model get hardware model default response has a 4xx status code
func (o *HardwareModelGetHardwareModelDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this hardware model get hardware model default response has a 5xx status code
func (o *HardwareModelGetHardwareModelDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this hardware model get hardware model default response a status code equal to that given
func (o *HardwareModelGetHardwareModelDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the hardware model get hardware model default response
func (o *HardwareModelGetHardwareModelDefault) Code() int {
	return o._statusCode
}

func (o *HardwareModelGetHardwareModelDefault) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/id/{id}][%d] HardwareModel_GetHardwareModel default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelGetHardwareModelDefault) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/id/{id}][%d] HardwareModel_GetHardwareModel default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelGetHardwareModelDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *HardwareModelGetHardwareModelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
