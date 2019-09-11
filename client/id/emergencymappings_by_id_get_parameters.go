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

// NewEmergencymappingsByIDGetParams creates a new EmergencymappingsByIDGetParams object
// with the default values initialized.
func NewEmergencymappingsByIDGetParams() *EmergencymappingsByIDGetParams {
	var ()
	return &EmergencymappingsByIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEmergencymappingsByIDGetParamsWithTimeout creates a new EmergencymappingsByIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmergencymappingsByIDGetParamsWithTimeout(timeout time.Duration) *EmergencymappingsByIDGetParams {
	var ()
	return &EmergencymappingsByIDGetParams{

		timeout: timeout,
	}
}

// NewEmergencymappingsByIDGetParamsWithContext creates a new EmergencymappingsByIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmergencymappingsByIDGetParamsWithContext(ctx context.Context) *EmergencymappingsByIDGetParams {
	var ()
	return &EmergencymappingsByIDGetParams{

		Context: ctx,
	}
}

// NewEmergencymappingsByIDGetParamsWithHTTPClient creates a new EmergencymappingsByIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmergencymappingsByIDGetParamsWithHTTPClient(client *http.Client) *EmergencymappingsByIDGetParams {
	var ()
	return &EmergencymappingsByIDGetParams{
		HTTPClient: client,
	}
}

/*EmergencymappingsByIDGetParams contains all the parameters to send to the API endpoint
for the emergencymappings by Id get operation typically these are written to a http.Request
*/
type EmergencymappingsByIDGetParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the emergencymappings by Id get params
func (o *EmergencymappingsByIDGetParams) WithTimeout(timeout time.Duration) *EmergencymappingsByIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the emergencymappings by Id get params
func (o *EmergencymappingsByIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the emergencymappings by Id get params
func (o *EmergencymappingsByIDGetParams) WithContext(ctx context.Context) *EmergencymappingsByIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the emergencymappings by Id get params
func (o *EmergencymappingsByIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the emergencymappings by Id get params
func (o *EmergencymappingsByIDGetParams) WithHTTPClient(client *http.Client) *EmergencymappingsByIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the emergencymappings by Id get params
func (o *EmergencymappingsByIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the emergencymappings by Id get params
func (o *EmergencymappingsByIDGetParams) WithID(id string) *EmergencymappingsByIDGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the emergencymappings by Id get params
func (o *EmergencymappingsByIDGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EmergencymappingsByIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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