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

// NewGetaspecificCFSourceSetParams creates a new GetaspecificCFSourceSetParams object
// with the default values initialized.
func NewGetaspecificCFSourceSetParams() *GetaspecificCFSourceSetParams {
	var ()
	return &GetaspecificCFSourceSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificCFSourceSetParamsWithTimeout creates a new GetaspecificCFSourceSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificCFSourceSetParamsWithTimeout(timeout time.Duration) *GetaspecificCFSourceSetParams {
	var ()
	return &GetaspecificCFSourceSetParams{

		timeout: timeout,
	}
}

// NewGetaspecificCFSourceSetParamsWithContext creates a new GetaspecificCFSourceSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificCFSourceSetParamsWithContext(ctx context.Context) *GetaspecificCFSourceSetParams {
	var ()
	return &GetaspecificCFSourceSetParams{

		Context: ctx,
	}
}

// NewGetaspecificCFSourceSetParamsWithHTTPClient creates a new GetaspecificCFSourceSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificCFSourceSetParamsWithHTTPClient(client *http.Client) *GetaspecificCFSourceSetParams {
	var ()
	return &GetaspecificCFSourceSetParams{
		HTTPClient: client,
	}
}

/*GetaspecificCFSourceSetParams contains all the parameters to send to the API endpoint
for the getaspecific c f source set operation typically these are written to a http.Request
*/
type GetaspecificCFSourceSetParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific c f source set params
func (o *GetaspecificCFSourceSetParams) WithTimeout(timeout time.Duration) *GetaspecificCFSourceSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific c f source set params
func (o *GetaspecificCFSourceSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific c f source set params
func (o *GetaspecificCFSourceSetParams) WithContext(ctx context.Context) *GetaspecificCFSourceSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific c f source set params
func (o *GetaspecificCFSourceSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific c f source set params
func (o *GetaspecificCFSourceSetParams) WithHTTPClient(client *http.Client) *GetaspecificCFSourceSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific c f source set params
func (o *GetaspecificCFSourceSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific c f source set params
func (o *GetaspecificCFSourceSetParams) WithID(id string) *GetaspecificCFSourceSetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific c f source set params
func (o *GetaspecificCFSourceSetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificCFSourceSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
