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

// NewDeleteaspecificBillingFeeParams creates a new DeleteaspecificBillingFeeParams object
// with the default values initialized.
func NewDeleteaspecificBillingFeeParams() *DeleteaspecificBillingFeeParams {
	var ()
	return &DeleteaspecificBillingFeeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteaspecificBillingFeeParamsWithTimeout creates a new DeleteaspecificBillingFeeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteaspecificBillingFeeParamsWithTimeout(timeout time.Duration) *DeleteaspecificBillingFeeParams {
	var ()
	return &DeleteaspecificBillingFeeParams{

		timeout: timeout,
	}
}

// NewDeleteaspecificBillingFeeParamsWithContext creates a new DeleteaspecificBillingFeeParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteaspecificBillingFeeParamsWithContext(ctx context.Context) *DeleteaspecificBillingFeeParams {
	var ()
	return &DeleteaspecificBillingFeeParams{

		Context: ctx,
	}
}

// NewDeleteaspecificBillingFeeParamsWithHTTPClient creates a new DeleteaspecificBillingFeeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteaspecificBillingFeeParamsWithHTTPClient(client *http.Client) *DeleteaspecificBillingFeeParams {
	var ()
	return &DeleteaspecificBillingFeeParams{
		HTTPClient: client,
	}
}

/*DeleteaspecificBillingFeeParams contains all the parameters to send to the API endpoint
for the deleteaspecific billing fee operation typically these are written to a http.Request
*/
type DeleteaspecificBillingFeeParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the deleteaspecific billing fee params
func (o *DeleteaspecificBillingFeeParams) WithTimeout(timeout time.Duration) *DeleteaspecificBillingFeeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deleteaspecific billing fee params
func (o *DeleteaspecificBillingFeeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deleteaspecific billing fee params
func (o *DeleteaspecificBillingFeeParams) WithContext(ctx context.Context) *DeleteaspecificBillingFeeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deleteaspecific billing fee params
func (o *DeleteaspecificBillingFeeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deleteaspecific billing fee params
func (o *DeleteaspecificBillingFeeParams) WithHTTPClient(client *http.Client) *DeleteaspecificBillingFeeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deleteaspecific billing fee params
func (o *DeleteaspecificBillingFeeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the deleteaspecific billing fee params
func (o *DeleteaspecificBillingFeeParams) WithID(id string) *DeleteaspecificBillingFeeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the deleteaspecific billing fee params
func (o *DeleteaspecificBillingFeeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteaspecificBillingFeeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
