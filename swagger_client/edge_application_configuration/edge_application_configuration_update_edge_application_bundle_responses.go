// Copyright (c) 2018-2021 Zededa, Inc.\n// SPDX-License-Identifier: Apache-2.0\n
// Code generated by go-swagger; DO NOT EDIT.

package edge_application_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// EdgeApplicationConfigurationUpdateEdgeApplicationBundleReader is a Reader for the EdgeApplicationConfigurationUpdateEdgeApplicationBundle structure.
type EdgeApplicationConfigurationUpdateEdgeApplicationBundleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleOK creates a EdgeApplicationConfigurationUpdateEdgeApplicationBundleOK with default headers values
func NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleOK() *EdgeApplicationConfigurationUpdateEdgeApplicationBundleOK {
	return &EdgeApplicationConfigurationUpdateEdgeApplicationBundleOK{}
}

/*
EdgeApplicationConfigurationUpdateEdgeApplicationBundleOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeApplicationConfigurationUpdateEdgeApplicationBundleOK struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration update edge application bundle o k response has a 2xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge application configuration update edge application bundle o k response has a 3xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration update edge application bundle o k response has a 4xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application configuration update edge application bundle o k response has a 5xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration update edge application bundle o k response a status code equal to that given
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleOK) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/id/{id}][%d] edgeApplicationConfigurationUpdateEdgeApplicationBundleOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleOK) String() string {
	return fmt.Sprintf("[PUT /v1/apps/id/{id}][%d] edgeApplicationConfigurationUpdateEdgeApplicationBundleOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleUnauthorized creates a EdgeApplicationConfigurationUpdateEdgeApplicationBundleUnauthorized with default headers values
func NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleUnauthorized() *EdgeApplicationConfigurationUpdateEdgeApplicationBundleUnauthorized {
	return &EdgeApplicationConfigurationUpdateEdgeApplicationBundleUnauthorized{}
}

/*
EdgeApplicationConfigurationUpdateEdgeApplicationBundleUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeApplicationConfigurationUpdateEdgeApplicationBundleUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration update edge application bundle unauthorized response has a 2xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration update edge application bundle unauthorized response has a 3xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration update edge application bundle unauthorized response has a 4xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application configuration update edge application bundle unauthorized response has a 5xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration update edge application bundle unauthorized response a status code equal to that given
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/id/{id}][%d] edgeApplicationConfigurationUpdateEdgeApplicationBundleUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/apps/id/{id}][%d] edgeApplicationConfigurationUpdateEdgeApplicationBundleUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleForbidden creates a EdgeApplicationConfigurationUpdateEdgeApplicationBundleForbidden with default headers values
func NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleForbidden() *EdgeApplicationConfigurationUpdateEdgeApplicationBundleForbidden {
	return &EdgeApplicationConfigurationUpdateEdgeApplicationBundleForbidden{}
}

/*
EdgeApplicationConfigurationUpdateEdgeApplicationBundleForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type EdgeApplicationConfigurationUpdateEdgeApplicationBundleForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration update edge application bundle forbidden response has a 2xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration update edge application bundle forbidden response has a 3xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration update edge application bundle forbidden response has a 4xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application configuration update edge application bundle forbidden response has a 5xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration update edge application bundle forbidden response a status code equal to that given
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/id/{id}][%d] edgeApplicationConfigurationUpdateEdgeApplicationBundleForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/apps/id/{id}][%d] edgeApplicationConfigurationUpdateEdgeApplicationBundleForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleNotFound creates a EdgeApplicationConfigurationUpdateEdgeApplicationBundleNotFound with default headers values
func NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleNotFound() *EdgeApplicationConfigurationUpdateEdgeApplicationBundleNotFound {
	return &EdgeApplicationConfigurationUpdateEdgeApplicationBundleNotFound{}
}

/*
EdgeApplicationConfigurationUpdateEdgeApplicationBundleNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeApplicationConfigurationUpdateEdgeApplicationBundleNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration update edge application bundle not found response has a 2xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration update edge application bundle not found response has a 3xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration update edge application bundle not found response has a 4xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application configuration update edge application bundle not found response has a 5xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration update edge application bundle not found response a status code equal to that given
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/id/{id}][%d] edgeApplicationConfigurationUpdateEdgeApplicationBundleNotFound  %+v", 404, o.Payload)
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/apps/id/{id}][%d] edgeApplicationConfigurationUpdateEdgeApplicationBundleNotFound  %+v", 404, o.Payload)
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleConflict creates a EdgeApplicationConfigurationUpdateEdgeApplicationBundleConflict with default headers values
func NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleConflict() *EdgeApplicationConfigurationUpdateEdgeApplicationBundleConflict {
	return &EdgeApplicationConfigurationUpdateEdgeApplicationBundleConflict{}
}

/*
EdgeApplicationConfigurationUpdateEdgeApplicationBundleConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing edge network record.
*/
type EdgeApplicationConfigurationUpdateEdgeApplicationBundleConflict struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration update edge application bundle conflict response has a 2xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration update edge application bundle conflict response has a 3xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration update edge application bundle conflict response has a 4xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application configuration update edge application bundle conflict response has a 5xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration update edge application bundle conflict response a status code equal to that given
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleConflict) IsCode(code int) bool {
	return code == 409
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/id/{id}][%d] edgeApplicationConfigurationUpdateEdgeApplicationBundleConflict  %+v", 409, o.Payload)
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleConflict) String() string {
	return fmt.Sprintf("[PUT /v1/apps/id/{id}][%d] edgeApplicationConfigurationUpdateEdgeApplicationBundleConflict  %+v", 409, o.Payload)
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleInternalServerError creates a EdgeApplicationConfigurationUpdateEdgeApplicationBundleInternalServerError with default headers values
func NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleInternalServerError() *EdgeApplicationConfigurationUpdateEdgeApplicationBundleInternalServerError {
	return &EdgeApplicationConfigurationUpdateEdgeApplicationBundleInternalServerError{}
}

/*
EdgeApplicationConfigurationUpdateEdgeApplicationBundleInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeApplicationConfigurationUpdateEdgeApplicationBundleInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration update edge application bundle internal server error response has a 2xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration update edge application bundle internal server error response has a 3xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration update edge application bundle internal server error response has a 4xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application configuration update edge application bundle internal server error response has a 5xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application configuration update edge application bundle internal server error response a status code equal to that given
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/id/{id}][%d] edgeApplicationConfigurationUpdateEdgeApplicationBundleInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/apps/id/{id}][%d] edgeApplicationConfigurationUpdateEdgeApplicationBundleInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleGatewayTimeout creates a EdgeApplicationConfigurationUpdateEdgeApplicationBundleGatewayTimeout with default headers values
func NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleGatewayTimeout() *EdgeApplicationConfigurationUpdateEdgeApplicationBundleGatewayTimeout {
	return &EdgeApplicationConfigurationUpdateEdgeApplicationBundleGatewayTimeout{}
}

/*
EdgeApplicationConfigurationUpdateEdgeApplicationBundleGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeApplicationConfigurationUpdateEdgeApplicationBundleGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration update edge application bundle gateway timeout response has a 2xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration update edge application bundle gateway timeout response has a 3xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration update edge application bundle gateway timeout response has a 4xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application configuration update edge application bundle gateway timeout response has a 5xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application configuration update edge application bundle gateway timeout response a status code equal to that given
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/id/{id}][%d] edgeApplicationConfigurationUpdateEdgeApplicationBundleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/apps/id/{id}][%d] edgeApplicationConfigurationUpdateEdgeApplicationBundleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleDefault creates a EdgeApplicationConfigurationUpdateEdgeApplicationBundleDefault with default headers values
func NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleDefault(code int) *EdgeApplicationConfigurationUpdateEdgeApplicationBundleDefault {
	return &EdgeApplicationConfigurationUpdateEdgeApplicationBundleDefault{
		_statusCode: code,
	}
}

/*
EdgeApplicationConfigurationUpdateEdgeApplicationBundleDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeApplicationConfigurationUpdateEdgeApplicationBundleDefault struct {
	_statusCode int

	Payload *swagger_models.GooglerpcStatus
}

// Code gets the status code for the edge application configuration update edge application bundle default response
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge application configuration update edge application bundle default response has a 2xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge application configuration update edge application bundle default response has a 3xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge application configuration update edge application bundle default response has a 4xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge application configuration update edge application bundle default response has a 5xx status code
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge application configuration update edge application bundle default response a status code equal to that given
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/id/{id}][%d] EdgeApplicationConfiguration_UpdateEdgeApplicationBundle default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleDefault) String() string {
	return fmt.Sprintf("[PUT /v1/apps/id/{id}][%d] EdgeApplicationConfiguration_UpdateEdgeApplicationBundle default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleDefault) GetPayload() *swagger_models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
EdgeApplicationConfigurationUpdateEdgeApplicationBundleBody Edge application detailed configuration
//
// Edge application gets installed/uninstalled to/from edge node(s) and perform specific edge computing tasks. Lifecycle of Edge application (upgrade/restart) on Edge node can be managed and monitored by ZEDEDA Cloud controller through this detailed configuration.
// Example: {"name":"sample-app","originType":"ORIGIN_LOCAL","title":"Sample Edge Applications"}
swagger:model EdgeApplicationConfigurationUpdateEdgeApplicationBundleBody
*/
type EdgeApplicationConfigurationUpdateEdgeApplicationBundleBody struct {

	// user defined cpus for bundle
	Cpus int64 `json:"cpus,omitempty"`

	// Detailed description of the edge application
	// Max Length: 256
	Description string `json:"description,omitempty"`

	// user defined drives
	// Read Only: true
	Drives int64 `json:"drives,omitempty"`

	// Flag to represent where app bundle is already imported
	IsImported bool `json:"isImported,omitempty"`

	// user defined manifest in JSON format
	ManifestJSON *swagger_models.VMManifest `json:"manifestJSON,omitempty"`

	// user defined memory for bundle
	Memory int64 `json:"memory,omitempty"`

	// User defined name of the edge application, unique across the enterprise. Once object is created, name can’t be changed
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	Name *string `json:"name"`

	// user defined network options
	Networks int64 `json:"networks,omitempty"`

	// origin of object
	// Required: true
	OriginType *swagger_models.Origin `json:"originType"`

	// origin and parent related details
	ParentDetail *swagger_models.ObjectParentDetail `json:"parentDetail,omitempty"`

	// project access list of the app bundle
	ProjectAccessList []string `json:"projectAccessList"`

	// system defined info
	Revision *swagger_models.ObjectRevision `json:"revision,omitempty"`

	// user defined storage for bundle
	Storage int64 `json:"storage,omitempty"`

	// User defined title of the edge application. Title can be changed at any time
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: ^[a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+$
	Title *string `json:"title"`

	// User defined version for the given edge-app
	UserDefinedVersion string `json:"userDefinedVersion,omitempty"`
}

// Validate validates this edge application configuration update edge application bundle body
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateManifestJSON(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOriginType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateParentDetail(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleBody) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(o.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("body"+"."+"description", "body", o.Description, 256); err != nil {
		return err
	}

	return nil
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleBody) validateManifestJSON(formats strfmt.Registry) error {
	if swag.IsZero(o.ManifestJSON) { // not required
		return nil
	}

	if o.ManifestJSON != nil {
		if err := o.ManifestJSON.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "manifestJSON")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "manifestJSON")
			}
			return err
		}
	}

	return nil
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	if err := validate.MinLength("body"+"."+"name", "body", *o.Name, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("body"+"."+"name", "body", *o.Name, 256); err != nil {
		return err
	}

	if err := validate.Pattern("body"+"."+"name", "body", *o.Name, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleBody) validateOriginType(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"originType", "body", o.OriginType); err != nil {
		return err
	}

	if err := validate.Required("body"+"."+"originType", "body", o.OriginType); err != nil {
		return err
	}

	if o.OriginType != nil {
		if err := o.OriginType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "originType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "originType")
			}
			return err
		}
	}

	return nil
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleBody) validateParentDetail(formats strfmt.Registry) error {
	if swag.IsZero(o.ParentDetail) { // not required
		return nil
	}

	if o.ParentDetail != nil {
		if err := o.ParentDetail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "parentDetail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "parentDetail")
			}
			return err
		}
	}

	return nil
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleBody) validateRevision(formats strfmt.Registry) error {
	if swag.IsZero(o.Revision) { // not required
		return nil
	}

	if o.Revision != nil {
		if err := o.Revision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "revision")
			}
			return err
		}
	}

	return nil
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleBody) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"title", "body", o.Title); err != nil {
		return err
	}

	if err := validate.MinLength("body"+"."+"title", "body", *o.Title, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("body"+"."+"title", "body", *o.Title, 256); err != nil {
		return err
	}

	if err := validate.Pattern("body"+"."+"title", "body", *o.Title, `^[a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this edge application configuration update edge application bundle body based on the context it is used
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDrives(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateManifestJSON(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateOriginType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateParentDetail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleBody) contextValidateDrives(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "body"+"."+"drives", "body", int64(o.Drives)); err != nil {
		return err
	}

	return nil
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleBody) contextValidateManifestJSON(ctx context.Context, formats strfmt.Registry) error {

	if o.ManifestJSON != nil {
		if err := o.ManifestJSON.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "manifestJSON")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "manifestJSON")
			}
			return err
		}
	}

	return nil
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleBody) contextValidateOriginType(ctx context.Context, formats strfmt.Registry) error {

	if o.OriginType != nil {
		if err := o.OriginType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "originType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "originType")
			}
			return err
		}
	}

	return nil
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleBody) contextValidateParentDetail(ctx context.Context, formats strfmt.Registry) error {

	if o.ParentDetail != nil {
		if err := o.ParentDetail.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "parentDetail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "parentDetail")
			}
			return err
		}
	}

	return nil
}

func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleBody) contextValidateRevision(ctx context.Context, formats strfmt.Registry) error {

	if o.Revision != nil {
		if err := o.Revision.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "revision")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleBody) UnmarshalBinary(b []byte) error {
	var res EdgeApplicationConfigurationUpdateEdgeApplicationBundleBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
