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

// NewGetaspecificCFDestinationSetParams creates a new GetaspecificCFDestinationSetParams object
// with the default values initialized.
func NewGetaspecificCFDestinationSetParams() *GetaspecificCFDestinationSetParams {
	var ()
	return &GetaspecificCFDestinationSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificCFDestinationSetParamsWithTimeout creates a new GetaspecificCFDestinationSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificCFDestinationSetParamsWithTimeout(timeout time.Duration) *GetaspecificCFDestinationSetParams {
	var ()
	return &GetaspecificCFDestinationSetParams{

		timeout: timeout,
	}
}

// NewGetaspecificCFDestinationSetParamsWithContext creates a new GetaspecificCFDestinationSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificCFDestinationSetParamsWithContext(ctx context.Context) *GetaspecificCFDestinationSetParams {
	var ()
	return &GetaspecificCFDestinationSetParams{

		Context: ctx,
	}
}

// NewGetaspecificCFDestinationSetParamsWithHTTPClient creates a new GetaspecificCFDestinationSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificCFDestinationSetParamsWithHTTPClient(client *http.Client) *GetaspecificCFDestinationSetParams {
	var ()
	return &GetaspecificCFDestinationSetParams{
		HTTPClient: client,
	}
}

/*GetaspecificCFDestinationSetParams contains all the parameters to send to the API endpoint
for the getaspecific c f destination set operation typically these are written to a http.Request
*/
type GetaspecificCFDestinationSetParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific c f destination set params
func (o *GetaspecificCFDestinationSetParams) WithTimeout(timeout time.Duration) *GetaspecificCFDestinationSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific c f destination set params
func (o *GetaspecificCFDestinationSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific c f destination set params
func (o *GetaspecificCFDestinationSetParams) WithContext(ctx context.Context) *GetaspecificCFDestinationSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific c f destination set params
func (o *GetaspecificCFDestinationSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific c f destination set params
func (o *GetaspecificCFDestinationSetParams) WithHTTPClient(client *http.Client) *GetaspecificCFDestinationSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific c f destination set params
func (o *GetaspecificCFDestinationSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific c f destination set params
func (o *GetaspecificCFDestinationSetParams) WithID(id string) *GetaspecificCFDestinationSetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific c f destination set params
func (o *GetaspecificCFDestinationSetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificCFDestinationSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
