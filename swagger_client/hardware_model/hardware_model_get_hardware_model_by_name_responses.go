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

// HardwareModelGetHardwareModelByNameReader is a Reader for the HardwareModelGetHardwareModelByName structure.
type HardwareModelGetHardwareModelByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HardwareModelGetHardwareModelByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHardwareModelGetHardwareModelByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewHardwareModelGetHardwareModelByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHardwareModelGetHardwareModelByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHardwareModelGetHardwareModelByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHardwareModelGetHardwareModelByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewHardwareModelGetHardwareModelByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewHardwareModelGetHardwareModelByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHardwareModelGetHardwareModelByNameOK creates a HardwareModelGetHardwareModelByNameOK with default headers values
func NewHardwareModelGetHardwareModelByNameOK() *HardwareModelGetHardwareModelByNameOK {
	return &HardwareModelGetHardwareModelByNameOK{}
}

/*
HardwareModelGetHardwareModelByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type HardwareModelGetHardwareModelByNameOK struct {
	Payload *swagger_models.SysModel
}

// IsSuccess returns true when this hardware model get hardware model by name o k response has a 2xx status code
func (o *HardwareModelGetHardwareModelByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hardware model get hardware model by name o k response has a 3xx status code
func (o *HardwareModelGetHardwareModelByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get hardware model by name o k response has a 4xx status code
func (o *HardwareModelGetHardwareModelByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get hardware model by name o k response has a 5xx status code
func (o *HardwareModelGetHardwareModelByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get hardware model by name o k response a status code equal to that given
func (o *HardwareModelGetHardwareModelByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the hardware model get hardware model by name o k response
func (o *HardwareModelGetHardwareModelByNameOK) Code() int {
	return 200
}

func (o *HardwareModelGetHardwareModelByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] hardwareModelGetHardwareModelByNameOK  %+v", 200, o.Payload)
}

func (o *HardwareModelGetHardwareModelByNameOK) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] hardwareModelGetHardwareModelByNameOK  %+v", 200, o.Payload)
}

func (o *HardwareModelGetHardwareModelByNameOK) GetPayload() *swagger_models.SysModel {
	return o.Payload
}

func (o *HardwareModelGetHardwareModelByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.SysModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareModelByNameUnauthorized creates a HardwareModelGetHardwareModelByNameUnauthorized with default headers values
func NewHardwareModelGetHardwareModelByNameUnauthorized() *HardwareModelGetHardwareModelByNameUnauthorized {
	return &HardwareModelGetHardwareModelByNameUnauthorized{}
}

/*
HardwareModelGetHardwareModelByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type HardwareModelGetHardwareModelByNameUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get hardware model by name unauthorized response has a 2xx status code
func (o *HardwareModelGetHardwareModelByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get hardware model by name unauthorized response has a 3xx status code
func (o *HardwareModelGetHardwareModelByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get hardware model by name unauthorized response has a 4xx status code
func (o *HardwareModelGetHardwareModelByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get hardware model by name unauthorized response has a 5xx status code
func (o *HardwareModelGetHardwareModelByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get hardware model by name unauthorized response a status code equal to that given
func (o *HardwareModelGetHardwareModelByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the hardware model get hardware model by name unauthorized response
func (o *HardwareModelGetHardwareModelByNameUnauthorized) Code() int {
	return 401
}

func (o *HardwareModelGetHardwareModelByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] hardwareModelGetHardwareModelByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelGetHardwareModelByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] hardwareModelGetHardwareModelByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelGetHardwareModelByNameUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetHardwareModelByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareModelByNameForbidden creates a HardwareModelGetHardwareModelByNameForbidden with default headers values
func NewHardwareModelGetHardwareModelByNameForbidden() *HardwareModelGetHardwareModelByNameForbidden {
	return &HardwareModelGetHardwareModelByNameForbidden{}
}

/*
HardwareModelGetHardwareModelByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type HardwareModelGetHardwareModelByNameForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get hardware model by name forbidden response has a 2xx status code
func (o *HardwareModelGetHardwareModelByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get hardware model by name forbidden response has a 3xx status code
func (o *HardwareModelGetHardwareModelByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get hardware model by name forbidden response has a 4xx status code
func (o *HardwareModelGetHardwareModelByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get hardware model by name forbidden response has a 5xx status code
func (o *HardwareModelGetHardwareModelByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get hardware model by name forbidden response a status code equal to that given
func (o *HardwareModelGetHardwareModelByNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the hardware model get hardware model by name forbidden response
func (o *HardwareModelGetHardwareModelByNameForbidden) Code() int {
	return 403
}

func (o *HardwareModelGetHardwareModelByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] hardwareModelGetHardwareModelByNameForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelGetHardwareModelByNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] hardwareModelGetHardwareModelByNameForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelGetHardwareModelByNameForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetHardwareModelByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareModelByNameNotFound creates a HardwareModelGetHardwareModelByNameNotFound with default headers values
func NewHardwareModelGetHardwareModelByNameNotFound() *HardwareModelGetHardwareModelByNameNotFound {
	return &HardwareModelGetHardwareModelByNameNotFound{}
}

/*
HardwareModelGetHardwareModelByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type HardwareModelGetHardwareModelByNameNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get hardware model by name not found response has a 2xx status code
func (o *HardwareModelGetHardwareModelByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get hardware model by name not found response has a 3xx status code
func (o *HardwareModelGetHardwareModelByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get hardware model by name not found response has a 4xx status code
func (o *HardwareModelGetHardwareModelByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get hardware model by name not found response has a 5xx status code
func (o *HardwareModelGetHardwareModelByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get hardware model by name not found response a status code equal to that given
func (o *HardwareModelGetHardwareModelByNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the hardware model get hardware model by name not found response
func (o *HardwareModelGetHardwareModelByNameNotFound) Code() int {
	return 404
}

func (o *HardwareModelGetHardwareModelByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] hardwareModelGetHardwareModelByNameNotFound  %+v", 404, o.Payload)
}

func (o *HardwareModelGetHardwareModelByNameNotFound) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] hardwareModelGetHardwareModelByNameNotFound  %+v", 404, o.Payload)
}

func (o *HardwareModelGetHardwareModelByNameNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetHardwareModelByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareModelByNameInternalServerError creates a HardwareModelGetHardwareModelByNameInternalServerError with default headers values
func NewHardwareModelGetHardwareModelByNameInternalServerError() *HardwareModelGetHardwareModelByNameInternalServerError {
	return &HardwareModelGetHardwareModelByNameInternalServerError{}
}

/*
HardwareModelGetHardwareModelByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type HardwareModelGetHardwareModelByNameInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get hardware model by name internal server error response has a 2xx status code
func (o *HardwareModelGetHardwareModelByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get hardware model by name internal server error response has a 3xx status code
func (o *HardwareModelGetHardwareModelByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get hardware model by name internal server error response has a 4xx status code
func (o *HardwareModelGetHardwareModelByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get hardware model by name internal server error response has a 5xx status code
func (o *HardwareModelGetHardwareModelByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model get hardware model by name internal server error response a status code equal to that given
func (o *HardwareModelGetHardwareModelByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the hardware model get hardware model by name internal server error response
func (o *HardwareModelGetHardwareModelByNameInternalServerError) Code() int {
	return 500
}

func (o *HardwareModelGetHardwareModelByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] hardwareModelGetHardwareModelByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelGetHardwareModelByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] hardwareModelGetHardwareModelByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelGetHardwareModelByNameInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetHardwareModelByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareModelByNameGatewayTimeout creates a HardwareModelGetHardwareModelByNameGatewayTimeout with default headers values
func NewHardwareModelGetHardwareModelByNameGatewayTimeout() *HardwareModelGetHardwareModelByNameGatewayTimeout {
	return &HardwareModelGetHardwareModelByNameGatewayTimeout{}
}

/*
HardwareModelGetHardwareModelByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type HardwareModelGetHardwareModelByNameGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get hardware model by name gateway timeout response has a 2xx status code
func (o *HardwareModelGetHardwareModelByNameGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get hardware model by name gateway timeout response has a 3xx status code
func (o *HardwareModelGetHardwareModelByNameGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get hardware model by name gateway timeout response has a 4xx status code
func (o *HardwareModelGetHardwareModelByNameGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get hardware model by name gateway timeout response has a 5xx status code
func (o *HardwareModelGetHardwareModelByNameGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model get hardware model by name gateway timeout response a status code equal to that given
func (o *HardwareModelGetHardwareModelByNameGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the hardware model get hardware model by name gateway timeout response
func (o *HardwareModelGetHardwareModelByNameGatewayTimeout) Code() int {
	return 504
}

func (o *HardwareModelGetHardwareModelByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] hardwareModelGetHardwareModelByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelGetHardwareModelByNameGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] hardwareModelGetHardwareModelByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelGetHardwareModelByNameGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetHardwareModelByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareModelByNameDefault creates a HardwareModelGetHardwareModelByNameDefault with default headers values
func NewHardwareModelGetHardwareModelByNameDefault(code int) *HardwareModelGetHardwareModelByNameDefault {
	return &HardwareModelGetHardwareModelByNameDefault{
		_statusCode: code,
	}
}

/*
HardwareModelGetHardwareModelByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HardwareModelGetHardwareModelByNameDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// IsSuccess returns true when this hardware model get hardware model by name default response has a 2xx status code
func (o *HardwareModelGetHardwareModelByNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this hardware model get hardware model by name default response has a 3xx status code
func (o *HardwareModelGetHardwareModelByNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this hardware model get hardware model by name default response has a 4xx status code
func (o *HardwareModelGetHardwareModelByNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this hardware model get hardware model by name default response has a 5xx status code
func (o *HardwareModelGetHardwareModelByNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this hardware model get hardware model by name default response a status code equal to that given
func (o *HardwareModelGetHardwareModelByNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the hardware model get hardware model by name default response
func (o *HardwareModelGetHardwareModelByNameDefault) Code() int {
	return o._statusCode
}

func (o *HardwareModelGetHardwareModelByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] HardwareModel_GetHardwareModelByName default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelGetHardwareModelByNameDefault) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/name/{name}][%d] HardwareModel_GetHardwareModelByName default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelGetHardwareModelByNameDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *HardwareModelGetHardwareModelByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
