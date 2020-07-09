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

// NewGetaspecificCustomerContactParams creates a new GetaspecificCustomerContactParams object
// with the default values initialized.
func NewGetaspecificCustomerContactParams() *GetaspecificCustomerContactParams {
	var ()
	return &GetaspecificCustomerContactParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificCustomerContactParamsWithTimeout creates a new GetaspecificCustomerContactParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificCustomerContactParamsWithTimeout(timeout time.Duration) *GetaspecificCustomerContactParams {
	var ()
	return &GetaspecificCustomerContactParams{

		timeout: timeout,
	}
}

// NewGetaspecificCustomerContactParamsWithContext creates a new GetaspecificCustomerContactParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificCustomerContactParamsWithContext(ctx context.Context) *GetaspecificCustomerContactParams {
	var ()
	return &GetaspecificCustomerContactParams{

		Context: ctx,
	}
}

// NewGetaspecificCustomerContactParamsWithHTTPClient creates a new GetaspecificCustomerContactParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificCustomerContactParamsWithHTTPClient(client *http.Client) *GetaspecificCustomerContactParams {
	var ()
	return &GetaspecificCustomerContactParams{
		HTTPClient: client,
	}
}

/*GetaspecificCustomerContactParams contains all the parameters to send to the API endpoint
for the getaspecific customer contact operation typically these are written to a http.Request
*/
type GetaspecificCustomerContactParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific customer contact params
func (o *GetaspecificCustomerContactParams) WithTimeout(timeout time.Duration) *GetaspecificCustomerContactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific customer contact params
func (o *GetaspecificCustomerContactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific customer contact params
func (o *GetaspecificCustomerContactParams) WithContext(ctx context.Context) *GetaspecificCustomerContactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific customer contact params
func (o *GetaspecificCustomerContactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific customer contact params
func (o *GetaspecificCustomerContactParams) WithHTTPClient(client *http.Client) *GetaspecificCustomerContactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific customer contact params
func (o *GetaspecificCustomerContactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific customer contact params
func (o *GetaspecificCustomerContactParams) WithID(id string) *GetaspecificCustomerContactParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific customer contact params
func (o *GetaspecificCustomerContactParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificCustomerContactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
