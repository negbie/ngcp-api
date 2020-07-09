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

// NewGetaspecificInvoiceParams creates a new GetaspecificInvoiceParams object
// with the default values initialized.
func NewGetaspecificInvoiceParams() *GetaspecificInvoiceParams {
	var ()
	return &GetaspecificInvoiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificInvoiceParamsWithTimeout creates a new GetaspecificInvoiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificInvoiceParamsWithTimeout(timeout time.Duration) *GetaspecificInvoiceParams {
	var ()
	return &GetaspecificInvoiceParams{

		timeout: timeout,
	}
}

// NewGetaspecificInvoiceParamsWithContext creates a new GetaspecificInvoiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificInvoiceParamsWithContext(ctx context.Context) *GetaspecificInvoiceParams {
	var ()
	return &GetaspecificInvoiceParams{

		Context: ctx,
	}
}

// NewGetaspecificInvoiceParamsWithHTTPClient creates a new GetaspecificInvoiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificInvoiceParamsWithHTTPClient(client *http.Client) *GetaspecificInvoiceParams {
	var ()
	return &GetaspecificInvoiceParams{
		HTTPClient: client,
	}
}

/*GetaspecificInvoiceParams contains all the parameters to send to the API endpoint
for the getaspecific invoice operation typically these are written to a http.Request
*/
type GetaspecificInvoiceParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific invoice params
func (o *GetaspecificInvoiceParams) WithTimeout(timeout time.Duration) *GetaspecificInvoiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific invoice params
func (o *GetaspecificInvoiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific invoice params
func (o *GetaspecificInvoiceParams) WithContext(ctx context.Context) *GetaspecificInvoiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific invoice params
func (o *GetaspecificInvoiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific invoice params
func (o *GetaspecificInvoiceParams) WithHTTPClient(client *http.Client) *GetaspecificInvoiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific invoice params
func (o *GetaspecificInvoiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific invoice params
func (o *GetaspecificInvoiceParams) WithID(id string) *GetaspecificInvoiceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific invoice params
func (o *GetaspecificInvoiceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificInvoiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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