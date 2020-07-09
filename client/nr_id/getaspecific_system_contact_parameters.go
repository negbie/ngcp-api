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

// NewGetaspecificSystemContactParams creates a new GetaspecificSystemContactParams object
// with the default values initialized.
func NewGetaspecificSystemContactParams() *GetaspecificSystemContactParams {
	var ()
	return &GetaspecificSystemContactParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificSystemContactParamsWithTimeout creates a new GetaspecificSystemContactParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificSystemContactParamsWithTimeout(timeout time.Duration) *GetaspecificSystemContactParams {
	var ()
	return &GetaspecificSystemContactParams{

		timeout: timeout,
	}
}

// NewGetaspecificSystemContactParamsWithContext creates a new GetaspecificSystemContactParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificSystemContactParamsWithContext(ctx context.Context) *GetaspecificSystemContactParams {
	var ()
	return &GetaspecificSystemContactParams{

		Context: ctx,
	}
}

// NewGetaspecificSystemContactParamsWithHTTPClient creates a new GetaspecificSystemContactParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificSystemContactParamsWithHTTPClient(client *http.Client) *GetaspecificSystemContactParams {
	var ()
	return &GetaspecificSystemContactParams{
		HTTPClient: client,
	}
}

/*GetaspecificSystemContactParams contains all the parameters to send to the API endpoint
for the getaspecific system contact operation typically these are written to a http.Request
*/
type GetaspecificSystemContactParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific system contact params
func (o *GetaspecificSystemContactParams) WithTimeout(timeout time.Duration) *GetaspecificSystemContactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific system contact params
func (o *GetaspecificSystemContactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific system contact params
func (o *GetaspecificSystemContactParams) WithContext(ctx context.Context) *GetaspecificSystemContactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific system contact params
func (o *GetaspecificSystemContactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific system contact params
func (o *GetaspecificSystemContactParams) WithHTTPClient(client *http.Client) *GetaspecificSystemContactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific system contact params
func (o *GetaspecificSystemContactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific system contact params
func (o *GetaspecificSystemContactParams) WithID(id string) *GetaspecificSystemContactParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific system contact params
func (o *GetaspecificSystemContactParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificSystemContactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
