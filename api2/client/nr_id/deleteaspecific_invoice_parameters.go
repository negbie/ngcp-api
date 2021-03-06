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

// NewDeleteaspecificInvoiceParams creates a new DeleteaspecificInvoiceParams object
// with the default values initialized.
func NewDeleteaspecificInvoiceParams() *DeleteaspecificInvoiceParams {
	var ()
	return &DeleteaspecificInvoiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteaspecificInvoiceParamsWithTimeout creates a new DeleteaspecificInvoiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteaspecificInvoiceParamsWithTimeout(timeout time.Duration) *DeleteaspecificInvoiceParams {
	var ()
	return &DeleteaspecificInvoiceParams{

		timeout: timeout,
	}
}

// NewDeleteaspecificInvoiceParamsWithContext creates a new DeleteaspecificInvoiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteaspecificInvoiceParamsWithContext(ctx context.Context) *DeleteaspecificInvoiceParams {
	var ()
	return &DeleteaspecificInvoiceParams{

		Context: ctx,
	}
}

// NewDeleteaspecificInvoiceParamsWithHTTPClient creates a new DeleteaspecificInvoiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteaspecificInvoiceParamsWithHTTPClient(client *http.Client) *DeleteaspecificInvoiceParams {
	var ()
	return &DeleteaspecificInvoiceParams{
		HTTPClient: client,
	}
}

/*DeleteaspecificInvoiceParams contains all the parameters to send to the API endpoint
for the deleteaspecific invoice operation typically these are written to a http.Request
*/
type DeleteaspecificInvoiceParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the deleteaspecific invoice params
func (o *DeleteaspecificInvoiceParams) WithTimeout(timeout time.Duration) *DeleteaspecificInvoiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deleteaspecific invoice params
func (o *DeleteaspecificInvoiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deleteaspecific invoice params
func (o *DeleteaspecificInvoiceParams) WithContext(ctx context.Context) *DeleteaspecificInvoiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deleteaspecific invoice params
func (o *DeleteaspecificInvoiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deleteaspecific invoice params
func (o *DeleteaspecificInvoiceParams) WithHTTPClient(client *http.Client) *DeleteaspecificInvoiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deleteaspecific invoice params
func (o *DeleteaspecificInvoiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the deleteaspecific invoice params
func (o *DeleteaspecificInvoiceParams) WithID(id string) *DeleteaspecificInvoiceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the deleteaspecific invoice params
func (o *DeleteaspecificInvoiceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteaspecificInvoiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
