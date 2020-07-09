// Code generated by go-swagger; DO NOT EDIT.

package nr_id

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

// NewGetaspecificCallForwardParams creates a new GetaspecificCallForwardParams object
// with the default values initialized.
func NewGetaspecificCallForwardParams() *GetaspecificCallForwardParams {
	var ()
	return &GetaspecificCallForwardParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificCallForwardParamsWithTimeout creates a new GetaspecificCallForwardParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificCallForwardParamsWithTimeout(timeout time.Duration) *GetaspecificCallForwardParams {
	var ()
	return &GetaspecificCallForwardParams{

		timeout: timeout,
	}
}

// NewGetaspecificCallForwardParamsWithContext creates a new GetaspecificCallForwardParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificCallForwardParamsWithContext(ctx context.Context) *GetaspecificCallForwardParams {
	var ()
	return &GetaspecificCallForwardParams{

		Context: ctx,
	}
}

// NewGetaspecificCallForwardParamsWithHTTPClient creates a new GetaspecificCallForwardParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificCallForwardParamsWithHTTPClient(client *http.Client) *GetaspecificCallForwardParams {
	var ()
	return &GetaspecificCallForwardParams{
		HTTPClient: client,
	}
}

/*GetaspecificCallForwardParams contains all the parameters to send to the API endpoint
for the getaspecific call forward operation typically these are written to a http.Request
*/
type GetaspecificCallForwardParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific call forward params
func (o *GetaspecificCallForwardParams) WithTimeout(timeout time.Duration) *GetaspecificCallForwardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific call forward params
func (o *GetaspecificCallForwardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific call forward params
func (o *GetaspecificCallForwardParams) WithContext(ctx context.Context) *GetaspecificCallForwardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific call forward params
func (o *GetaspecificCallForwardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific call forward params
func (o *GetaspecificCallForwardParams) WithHTTPClient(client *http.Client) *GetaspecificCallForwardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific call forward params
func (o *GetaspecificCallForwardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific call forward params
func (o *GetaspecificCallForwardParams) WithID(id string) *GetaspecificCallForwardParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific call forward params
func (o *GetaspecificCallForwardParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificCallForwardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
