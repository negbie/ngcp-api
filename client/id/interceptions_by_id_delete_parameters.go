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

// NewInterceptionsByIDDeleteParams creates a new InterceptionsByIDDeleteParams object
// with the default values initialized.
func NewInterceptionsByIDDeleteParams() *InterceptionsByIDDeleteParams {
	var ()
	return &InterceptionsByIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInterceptionsByIDDeleteParamsWithTimeout creates a new InterceptionsByIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInterceptionsByIDDeleteParamsWithTimeout(timeout time.Duration) *InterceptionsByIDDeleteParams {
	var ()
	return &InterceptionsByIDDeleteParams{

		timeout: timeout,
	}
}

// NewInterceptionsByIDDeleteParamsWithContext creates a new InterceptionsByIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewInterceptionsByIDDeleteParamsWithContext(ctx context.Context) *InterceptionsByIDDeleteParams {
	var ()
	return &InterceptionsByIDDeleteParams{

		Context: ctx,
	}
}

// NewInterceptionsByIDDeleteParamsWithHTTPClient creates a new InterceptionsByIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInterceptionsByIDDeleteParamsWithHTTPClient(client *http.Client) *InterceptionsByIDDeleteParams {
	var ()
	return &InterceptionsByIDDeleteParams{
		HTTPClient: client,
	}
}

/*InterceptionsByIDDeleteParams contains all the parameters to send to the API endpoint
for the interceptions by Id delete operation typically these are written to a http.Request
*/
type InterceptionsByIDDeleteParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the interceptions by Id delete params
func (o *InterceptionsByIDDeleteParams) WithTimeout(timeout time.Duration) *InterceptionsByIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the interceptions by Id delete params
func (o *InterceptionsByIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the interceptions by Id delete params
func (o *InterceptionsByIDDeleteParams) WithContext(ctx context.Context) *InterceptionsByIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the interceptions by Id delete params
func (o *InterceptionsByIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the interceptions by Id delete params
func (o *InterceptionsByIDDeleteParams) WithHTTPClient(client *http.Client) *InterceptionsByIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the interceptions by Id delete params
func (o *InterceptionsByIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the interceptions by Id delete params
func (o *InterceptionsByIDDeleteParams) WithID(id string) *InterceptionsByIDDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the interceptions by Id delete params
func (o *InterceptionsByIDDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *InterceptionsByIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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