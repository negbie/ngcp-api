// Code generated by go-swagger; DO NOT EDIT.

package id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewEmergencymappingcontainersByIDGetParams creates a new EmergencymappingcontainersByIDGetParams object
// with the default values initialized.
func NewEmergencymappingcontainersByIDGetParams() *EmergencymappingcontainersByIDGetParams {
	var ()
	return &EmergencymappingcontainersByIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEmergencymappingcontainersByIDGetParamsWithTimeout creates a new EmergencymappingcontainersByIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmergencymappingcontainersByIDGetParamsWithTimeout(timeout time.Duration) *EmergencymappingcontainersByIDGetParams {
	var ()
	return &EmergencymappingcontainersByIDGetParams{

		timeout: timeout,
	}
}

// NewEmergencymappingcontainersByIDGetParamsWithContext creates a new EmergencymappingcontainersByIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmergencymappingcontainersByIDGetParamsWithContext(ctx context.Context) *EmergencymappingcontainersByIDGetParams {
	var ()
	return &EmergencymappingcontainersByIDGetParams{

		Context: ctx,
	}
}

// NewEmergencymappingcontainersByIDGetParamsWithHTTPClient creates a new EmergencymappingcontainersByIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmergencymappingcontainersByIDGetParamsWithHTTPClient(client *http.Client) *EmergencymappingcontainersByIDGetParams {
	var ()
	return &EmergencymappingcontainersByIDGetParams{
		HTTPClient: client,
	}
}

/*EmergencymappingcontainersByIDGetParams contains all the parameters to send to the API endpoint
for the emergencymappingcontainers by Id get operation typically these are written to a http.Request
*/
type EmergencymappingcontainersByIDGetParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the emergencymappingcontainers by Id get params
func (o *EmergencymappingcontainersByIDGetParams) WithTimeout(timeout time.Duration) *EmergencymappingcontainersByIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the emergencymappingcontainers by Id get params
func (o *EmergencymappingcontainersByIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the emergencymappingcontainers by Id get params
func (o *EmergencymappingcontainersByIDGetParams) WithContext(ctx context.Context) *EmergencymappingcontainersByIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the emergencymappingcontainers by Id get params
func (o *EmergencymappingcontainersByIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the emergencymappingcontainers by Id get params
func (o *EmergencymappingcontainersByIDGetParams) WithHTTPClient(client *http.Client) *EmergencymappingcontainersByIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the emergencymappingcontainers by Id get params
func (o *EmergencymappingcontainersByIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the emergencymappingcontainers by Id get params
func (o *EmergencymappingcontainersByIDGetParams) WithID(id string) *EmergencymappingcontainersByIDGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the emergencymappingcontainers by Id get params
func (o *EmergencymappingcontainersByIDGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EmergencymappingcontainersByIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}