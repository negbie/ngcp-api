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

// NewDeleteaspecificCustomerContactParams creates a new DeleteaspecificCustomerContactParams object
// with the default values initialized.
func NewDeleteaspecificCustomerContactParams() *DeleteaspecificCustomerContactParams {
	var ()
	return &DeleteaspecificCustomerContactParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteaspecificCustomerContactParamsWithTimeout creates a new DeleteaspecificCustomerContactParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteaspecificCustomerContactParamsWithTimeout(timeout time.Duration) *DeleteaspecificCustomerContactParams {
	var ()
	return &DeleteaspecificCustomerContactParams{

		timeout: timeout,
	}
}

// NewDeleteaspecificCustomerContactParamsWithContext creates a new DeleteaspecificCustomerContactParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteaspecificCustomerContactParamsWithContext(ctx context.Context) *DeleteaspecificCustomerContactParams {
	var ()
	return &DeleteaspecificCustomerContactParams{

		Context: ctx,
	}
}

// NewDeleteaspecificCustomerContactParamsWithHTTPClient creates a new DeleteaspecificCustomerContactParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteaspecificCustomerContactParamsWithHTTPClient(client *http.Client) *DeleteaspecificCustomerContactParams {
	var ()
	return &DeleteaspecificCustomerContactParams{
		HTTPClient: client,
	}
}

/*DeleteaspecificCustomerContactParams contains all the parameters to send to the API endpoint
for the deleteaspecific customer contact operation typically these are written to a http.Request
*/
type DeleteaspecificCustomerContactParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the deleteaspecific customer contact params
func (o *DeleteaspecificCustomerContactParams) WithTimeout(timeout time.Duration) *DeleteaspecificCustomerContactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deleteaspecific customer contact params
func (o *DeleteaspecificCustomerContactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deleteaspecific customer contact params
func (o *DeleteaspecificCustomerContactParams) WithContext(ctx context.Context) *DeleteaspecificCustomerContactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deleteaspecific customer contact params
func (o *DeleteaspecificCustomerContactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deleteaspecific customer contact params
func (o *DeleteaspecificCustomerContactParams) WithHTTPClient(client *http.Client) *DeleteaspecificCustomerContactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deleteaspecific customer contact params
func (o *DeleteaspecificCustomerContactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the deleteaspecific customer contact params
func (o *DeleteaspecificCustomerContactParams) WithID(id string) *DeleteaspecificCustomerContactParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the deleteaspecific customer contact params
func (o *DeleteaspecificCustomerContactParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteaspecificCustomerContactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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