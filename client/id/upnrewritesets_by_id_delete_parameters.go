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

// NewUpnrewritesetsByIDDeleteParams creates a new UpnrewritesetsByIDDeleteParams object
// with the default values initialized.
func NewUpnrewritesetsByIDDeleteParams() *UpnrewritesetsByIDDeleteParams {
	var ()
	return &UpnrewritesetsByIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpnrewritesetsByIDDeleteParamsWithTimeout creates a new UpnrewritesetsByIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpnrewritesetsByIDDeleteParamsWithTimeout(timeout time.Duration) *UpnrewritesetsByIDDeleteParams {
	var ()
	return &UpnrewritesetsByIDDeleteParams{

		timeout: timeout,
	}
}

// NewUpnrewritesetsByIDDeleteParamsWithContext creates a new UpnrewritesetsByIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpnrewritesetsByIDDeleteParamsWithContext(ctx context.Context) *UpnrewritesetsByIDDeleteParams {
	var ()
	return &UpnrewritesetsByIDDeleteParams{

		Context: ctx,
	}
}

// NewUpnrewritesetsByIDDeleteParamsWithHTTPClient creates a new UpnrewritesetsByIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpnrewritesetsByIDDeleteParamsWithHTTPClient(client *http.Client) *UpnrewritesetsByIDDeleteParams {
	var ()
	return &UpnrewritesetsByIDDeleteParams{
		HTTPClient: client,
	}
}

/*UpnrewritesetsByIDDeleteParams contains all the parameters to send to the API endpoint
for the upnrewritesets by Id delete operation typically these are written to a http.Request
*/
type UpnrewritesetsByIDDeleteParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the upnrewritesets by Id delete params
func (o *UpnrewritesetsByIDDeleteParams) WithTimeout(timeout time.Duration) *UpnrewritesetsByIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upnrewritesets by Id delete params
func (o *UpnrewritesetsByIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upnrewritesets by Id delete params
func (o *UpnrewritesetsByIDDeleteParams) WithContext(ctx context.Context) *UpnrewritesetsByIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upnrewritesets by Id delete params
func (o *UpnrewritesetsByIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upnrewritesets by Id delete params
func (o *UpnrewritesetsByIDDeleteParams) WithHTTPClient(client *http.Client) *UpnrewritesetsByIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upnrewritesets by Id delete params
func (o *UpnrewritesetsByIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the upnrewritesets by Id delete params
func (o *UpnrewritesetsByIDDeleteParams) WithID(id string) *UpnrewritesetsByIDDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the upnrewritesets by Id delete params
func (o *UpnrewritesetsByIDDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpnrewritesetsByIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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