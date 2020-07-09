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

// NewGetaspecificLnpNumberParams creates a new GetaspecificLnpNumberParams object
// with the default values initialized.
func NewGetaspecificLnpNumberParams() *GetaspecificLnpNumberParams {
	var ()
	return &GetaspecificLnpNumberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificLnpNumberParamsWithTimeout creates a new GetaspecificLnpNumberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificLnpNumberParamsWithTimeout(timeout time.Duration) *GetaspecificLnpNumberParams {
	var ()
	return &GetaspecificLnpNumberParams{

		timeout: timeout,
	}
}

// NewGetaspecificLnpNumberParamsWithContext creates a new GetaspecificLnpNumberParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificLnpNumberParamsWithContext(ctx context.Context) *GetaspecificLnpNumberParams {
	var ()
	return &GetaspecificLnpNumberParams{

		Context: ctx,
	}
}

// NewGetaspecificLnpNumberParamsWithHTTPClient creates a new GetaspecificLnpNumberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificLnpNumberParamsWithHTTPClient(client *http.Client) *GetaspecificLnpNumberParams {
	var ()
	return &GetaspecificLnpNumberParams{
		HTTPClient: client,
	}
}

/*GetaspecificLnpNumberParams contains all the parameters to send to the API endpoint
for the getaspecific lnp number operation typically these are written to a http.Request
*/
type GetaspecificLnpNumberParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific lnp number params
func (o *GetaspecificLnpNumberParams) WithTimeout(timeout time.Duration) *GetaspecificLnpNumberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific lnp number params
func (o *GetaspecificLnpNumberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific lnp number params
func (o *GetaspecificLnpNumberParams) WithContext(ctx context.Context) *GetaspecificLnpNumberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific lnp number params
func (o *GetaspecificLnpNumberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific lnp number params
func (o *GetaspecificLnpNumberParams) WithHTTPClient(client *http.Client) *GetaspecificLnpNumberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific lnp number params
func (o *GetaspecificLnpNumberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific lnp number params
func (o *GetaspecificLnpNumberParams) WithID(id string) *GetaspecificLnpNumberParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific lnp number params
func (o *GetaspecificLnpNumberParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificLnpNumberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
