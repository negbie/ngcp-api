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

// NewDeleteaspecificVoucherParams creates a new DeleteaspecificVoucherParams object
// with the default values initialized.
func NewDeleteaspecificVoucherParams() *DeleteaspecificVoucherParams {
	var ()
	return &DeleteaspecificVoucherParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteaspecificVoucherParamsWithTimeout creates a new DeleteaspecificVoucherParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteaspecificVoucherParamsWithTimeout(timeout time.Duration) *DeleteaspecificVoucherParams {
	var ()
	return &DeleteaspecificVoucherParams{

		timeout: timeout,
	}
}

// NewDeleteaspecificVoucherParamsWithContext creates a new DeleteaspecificVoucherParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteaspecificVoucherParamsWithContext(ctx context.Context) *DeleteaspecificVoucherParams {
	var ()
	return &DeleteaspecificVoucherParams{

		Context: ctx,
	}
}

// NewDeleteaspecificVoucherParamsWithHTTPClient creates a new DeleteaspecificVoucherParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteaspecificVoucherParamsWithHTTPClient(client *http.Client) *DeleteaspecificVoucherParams {
	var ()
	return &DeleteaspecificVoucherParams{
		HTTPClient: client,
	}
}

/*DeleteaspecificVoucherParams contains all the parameters to send to the API endpoint
for the deleteaspecific voucher operation typically these are written to a http.Request
*/
type DeleteaspecificVoucherParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the deleteaspecific voucher params
func (o *DeleteaspecificVoucherParams) WithTimeout(timeout time.Duration) *DeleteaspecificVoucherParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deleteaspecific voucher params
func (o *DeleteaspecificVoucherParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deleteaspecific voucher params
func (o *DeleteaspecificVoucherParams) WithContext(ctx context.Context) *DeleteaspecificVoucherParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deleteaspecific voucher params
func (o *DeleteaspecificVoucherParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deleteaspecific voucher params
func (o *DeleteaspecificVoucherParams) WithHTTPClient(client *http.Client) *DeleteaspecificVoucherParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deleteaspecific voucher params
func (o *DeleteaspecificVoucherParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the deleteaspecific voucher params
func (o *DeleteaspecificVoucherParams) WithID(id string) *DeleteaspecificVoucherParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the deleteaspecific voucher params
func (o *DeleteaspecificVoucherParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteaspecificVoucherParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
