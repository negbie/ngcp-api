// Code generated by go-swagger; DO NOT EDIT.

package sms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetaspecificSMSParams creates a new GetaspecificSMSParams object
// with the default values initialized.
func NewGetaspecificSMSParams() *GetaspecificSMSParams {
	var ()
	return &GetaspecificSMSParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificSMSParamsWithTimeout creates a new GetaspecificSMSParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificSMSParamsWithTimeout(timeout time.Duration) *GetaspecificSMSParams {
	var ()
	return &GetaspecificSMSParams{

		timeout: timeout,
	}
}

// NewGetaspecificSMSParamsWithContext creates a new GetaspecificSMSParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificSMSParamsWithContext(ctx context.Context) *GetaspecificSMSParams {
	var ()
	return &GetaspecificSMSParams{

		Context: ctx,
	}
}

// NewGetaspecificSMSParamsWithHTTPClient creates a new GetaspecificSMSParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificSMSParamsWithHTTPClient(client *http.Client) *GetaspecificSMSParams {
	var ()
	return &GetaspecificSMSParams{
		HTTPClient: client,
	}
}

/*GetaspecificSMSParams contains all the parameters to send to the API endpoint
for the getaspecific s m s operation typically these are written to a http.Request
*/
type GetaspecificSMSParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific s m s params
func (o *GetaspecificSMSParams) WithTimeout(timeout time.Duration) *GetaspecificSMSParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific s m s params
func (o *GetaspecificSMSParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific s m s params
func (o *GetaspecificSMSParams) WithContext(ctx context.Context) *GetaspecificSMSParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific s m s params
func (o *GetaspecificSMSParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific s m s params
func (o *GetaspecificSMSParams) WithHTTPClient(client *http.Client) *GetaspecificSMSParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific s m s params
func (o *GetaspecificSMSParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific s m s params
func (o *GetaspecificSMSParams) WithID(id string) *GetaspecificSMSParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific s m s params
func (o *GetaspecificSMSParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificSMSParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
