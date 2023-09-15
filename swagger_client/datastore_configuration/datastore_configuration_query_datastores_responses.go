// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package datastore_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// DatastoreConfigurationQueryDatastoresReader is a Reader for the DatastoreConfigurationQueryDatastores structure.
type DatastoreConfigurationQueryDatastoresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DatastoreConfigurationQueryDatastoresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDatastoreConfigurationQueryDatastoresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDatastoreConfigurationQueryDatastoresBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDatastoreConfigurationQueryDatastoresUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDatastoreConfigurationQueryDatastoresForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDatastoreConfigurationQueryDatastoresInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDatastoreConfigurationQueryDatastoresGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDatastoreConfigurationQueryDatastoresDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDatastoreConfigurationQueryDatastoresOK creates a DatastoreConfigurationQueryDatastoresOK with default headers values
func NewDatastoreConfigurationQueryDatastoresOK() *DatastoreConfigurationQueryDatastoresOK {
	return &DatastoreConfigurationQueryDatastoresOK{}
}

/*
DatastoreConfigurationQueryDatastoresOK describes a response with status code 200, with default header values.

A successful response.
*/
type DatastoreConfigurationQueryDatastoresOK struct {
	Payload *swagger_models.Datastores
}

// IsSuccess returns true when this datastore configuration query datastores o k response has a 2xx status code
func (o *DatastoreConfigurationQueryDatastoresOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this datastore configuration query datastores o k response has a 3xx status code
func (o *DatastoreConfigurationQueryDatastoresOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datastore configuration query datastores o k response has a 4xx status code
func (o *DatastoreConfigurationQueryDatastoresOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this datastore configuration query datastores o k response has a 5xx status code
func (o *DatastoreConfigurationQueryDatastoresOK) IsServerError() bool {
	return false
}

// IsCode returns true when this datastore configuration query datastores o k response a status code equal to that given
func (o *DatastoreConfigurationQueryDatastoresOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the datastore configuration query datastores o k response
func (o *DatastoreConfigurationQueryDatastoresOK) Code() int {
	return 200
}

func (o *DatastoreConfigurationQueryDatastoresOK) Error() string {
	return fmt.Sprintf("[GET /v1/datastores][%d] datastoreConfigurationQueryDatastoresOK  %+v", 200, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoresOK) String() string {
	return fmt.Sprintf("[GET /v1/datastores][%d] datastoreConfigurationQueryDatastoresOK  %+v", 200, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoresOK) GetPayload() *swagger_models.Datastores {
	return o.Payload
}

func (o *DatastoreConfigurationQueryDatastoresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.Datastores)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatastoreConfigurationQueryDatastoresBadRequest creates a DatastoreConfigurationQueryDatastoresBadRequest with default headers values
func NewDatastoreConfigurationQueryDatastoresBadRequest() *DatastoreConfigurationQueryDatastoresBadRequest {
	return &DatastoreConfigurationQueryDatastoresBadRequest{}
}

/*
DatastoreConfigurationQueryDatastoresBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type DatastoreConfigurationQueryDatastoresBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this datastore configuration query datastores bad request response has a 2xx status code
func (o *DatastoreConfigurationQueryDatastoresBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datastore configuration query datastores bad request response has a 3xx status code
func (o *DatastoreConfigurationQueryDatastoresBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datastore configuration query datastores bad request response has a 4xx status code
func (o *DatastoreConfigurationQueryDatastoresBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this datastore configuration query datastores bad request response has a 5xx status code
func (o *DatastoreConfigurationQueryDatastoresBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this datastore configuration query datastores bad request response a status code equal to that given
func (o *DatastoreConfigurationQueryDatastoresBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the datastore configuration query datastores bad request response
func (o *DatastoreConfigurationQueryDatastoresBadRequest) Code() int {
	return 400
}

func (o *DatastoreConfigurationQueryDatastoresBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/datastores][%d] datastoreConfigurationQueryDatastoresBadRequest  %+v", 400, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoresBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/datastores][%d] datastoreConfigurationQueryDatastoresBadRequest  %+v", 400, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoresBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DatastoreConfigurationQueryDatastoresBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatastoreConfigurationQueryDatastoresUnauthorized creates a DatastoreConfigurationQueryDatastoresUnauthorized with default headers values
func NewDatastoreConfigurationQueryDatastoresUnauthorized() *DatastoreConfigurationQueryDatastoresUnauthorized {
	return &DatastoreConfigurationQueryDatastoresUnauthorized{}
}

/*
DatastoreConfigurationQueryDatastoresUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type DatastoreConfigurationQueryDatastoresUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this datastore configuration query datastores unauthorized response has a 2xx status code
func (o *DatastoreConfigurationQueryDatastoresUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datastore configuration query datastores unauthorized response has a 3xx status code
func (o *DatastoreConfigurationQueryDatastoresUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datastore configuration query datastores unauthorized response has a 4xx status code
func (o *DatastoreConfigurationQueryDatastoresUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this datastore configuration query datastores unauthorized response has a 5xx status code
func (o *DatastoreConfigurationQueryDatastoresUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this datastore configuration query datastores unauthorized response a status code equal to that given
func (o *DatastoreConfigurationQueryDatastoresUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the datastore configuration query datastores unauthorized response
func (o *DatastoreConfigurationQueryDatastoresUnauthorized) Code() int {
	return 401
}

func (o *DatastoreConfigurationQueryDatastoresUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/datastores][%d] datastoreConfigurationQueryDatastoresUnauthorized  %+v", 401, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoresUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/datastores][%d] datastoreConfigurationQueryDatastoresUnauthorized  %+v", 401, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoresUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DatastoreConfigurationQueryDatastoresUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatastoreConfigurationQueryDatastoresForbidden creates a DatastoreConfigurationQueryDatastoresForbidden with default headers values
func NewDatastoreConfigurationQueryDatastoresForbidden() *DatastoreConfigurationQueryDatastoresForbidden {
	return &DatastoreConfigurationQueryDatastoresForbidden{}
}

/*
DatastoreConfigurationQueryDatastoresForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type DatastoreConfigurationQueryDatastoresForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this datastore configuration query datastores forbidden response has a 2xx status code
func (o *DatastoreConfigurationQueryDatastoresForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datastore configuration query datastores forbidden response has a 3xx status code
func (o *DatastoreConfigurationQueryDatastoresForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datastore configuration query datastores forbidden response has a 4xx status code
func (o *DatastoreConfigurationQueryDatastoresForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this datastore configuration query datastores forbidden response has a 5xx status code
func (o *DatastoreConfigurationQueryDatastoresForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this datastore configuration query datastores forbidden response a status code equal to that given
func (o *DatastoreConfigurationQueryDatastoresForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the datastore configuration query datastores forbidden response
func (o *DatastoreConfigurationQueryDatastoresForbidden) Code() int {
	return 403
}

func (o *DatastoreConfigurationQueryDatastoresForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/datastores][%d] datastoreConfigurationQueryDatastoresForbidden  %+v", 403, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoresForbidden) String() string {
	return fmt.Sprintf("[GET /v1/datastores][%d] datastoreConfigurationQueryDatastoresForbidden  %+v", 403, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoresForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DatastoreConfigurationQueryDatastoresForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatastoreConfigurationQueryDatastoresInternalServerError creates a DatastoreConfigurationQueryDatastoresInternalServerError with default headers values
func NewDatastoreConfigurationQueryDatastoresInternalServerError() *DatastoreConfigurationQueryDatastoresInternalServerError {
	return &DatastoreConfigurationQueryDatastoresInternalServerError{}
}

/*
DatastoreConfigurationQueryDatastoresInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type DatastoreConfigurationQueryDatastoresInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this datastore configuration query datastores internal server error response has a 2xx status code
func (o *DatastoreConfigurationQueryDatastoresInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datastore configuration query datastores internal server error response has a 3xx status code
func (o *DatastoreConfigurationQueryDatastoresInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datastore configuration query datastores internal server error response has a 4xx status code
func (o *DatastoreConfigurationQueryDatastoresInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this datastore configuration query datastores internal server error response has a 5xx status code
func (o *DatastoreConfigurationQueryDatastoresInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this datastore configuration query datastores internal server error response a status code equal to that given
func (o *DatastoreConfigurationQueryDatastoresInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the datastore configuration query datastores internal server error response
func (o *DatastoreConfigurationQueryDatastoresInternalServerError) Code() int {
	return 500
}

func (o *DatastoreConfigurationQueryDatastoresInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/datastores][%d] datastoreConfigurationQueryDatastoresInternalServerError  %+v", 500, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoresInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/datastores][%d] datastoreConfigurationQueryDatastoresInternalServerError  %+v", 500, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoresInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DatastoreConfigurationQueryDatastoresInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatastoreConfigurationQueryDatastoresGatewayTimeout creates a DatastoreConfigurationQueryDatastoresGatewayTimeout with default headers values
func NewDatastoreConfigurationQueryDatastoresGatewayTimeout() *DatastoreConfigurationQueryDatastoresGatewayTimeout {
	return &DatastoreConfigurationQueryDatastoresGatewayTimeout{}
}

/*
DatastoreConfigurationQueryDatastoresGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type DatastoreConfigurationQueryDatastoresGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this datastore configuration query datastores gateway timeout response has a 2xx status code
func (o *DatastoreConfigurationQueryDatastoresGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datastore configuration query datastores gateway timeout response has a 3xx status code
func (o *DatastoreConfigurationQueryDatastoresGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datastore configuration query datastores gateway timeout response has a 4xx status code
func (o *DatastoreConfigurationQueryDatastoresGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this datastore configuration query datastores gateway timeout response has a 5xx status code
func (o *DatastoreConfigurationQueryDatastoresGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this datastore configuration query datastores gateway timeout response a status code equal to that given
func (o *DatastoreConfigurationQueryDatastoresGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the datastore configuration query datastores gateway timeout response
func (o *DatastoreConfigurationQueryDatastoresGatewayTimeout) Code() int {
	return 504
}

func (o *DatastoreConfigurationQueryDatastoresGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/datastores][%d] datastoreConfigurationQueryDatastoresGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoresGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/datastores][%d] datastoreConfigurationQueryDatastoresGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoresGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *DatastoreConfigurationQueryDatastoresGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatastoreConfigurationQueryDatastoresDefault creates a DatastoreConfigurationQueryDatastoresDefault with default headers values
func NewDatastoreConfigurationQueryDatastoresDefault(code int) *DatastoreConfigurationQueryDatastoresDefault {
	return &DatastoreConfigurationQueryDatastoresDefault{
		_statusCode: code,
	}
}

/*
DatastoreConfigurationQueryDatastoresDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DatastoreConfigurationQueryDatastoresDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// IsSuccess returns true when this datastore configuration query datastores default response has a 2xx status code
func (o *DatastoreConfigurationQueryDatastoresDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this datastore configuration query datastores default response has a 3xx status code
func (o *DatastoreConfigurationQueryDatastoresDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this datastore configuration query datastores default response has a 4xx status code
func (o *DatastoreConfigurationQueryDatastoresDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this datastore configuration query datastores default response has a 5xx status code
func (o *DatastoreConfigurationQueryDatastoresDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this datastore configuration query datastores default response a status code equal to that given
func (o *DatastoreConfigurationQueryDatastoresDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the datastore configuration query datastores default response
func (o *DatastoreConfigurationQueryDatastoresDefault) Code() int {
	return o._statusCode
}

func (o *DatastoreConfigurationQueryDatastoresDefault) Error() string {
	return fmt.Sprintf("[GET /v1/datastores][%d] DatastoreConfiguration_QueryDatastores default  %+v", o._statusCode, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoresDefault) String() string {
	return fmt.Sprintf("[GET /v1/datastores][%d] DatastoreConfiguration_QueryDatastores default  %+v", o._statusCode, o.Payload)
}

func (o *DatastoreConfigurationQueryDatastoresDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *DatastoreConfigurationQueryDatastoresDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
