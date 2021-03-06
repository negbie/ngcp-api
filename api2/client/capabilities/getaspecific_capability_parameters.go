// Code generated by go-swagger; DO NOT EDIT.

package capabilities

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

// NewGetaspecificCapabilityParams creates a new GetaspecificCapabilityParams object
// with the default values initialized.
func NewGetaspecificCapabilityParams() *GetaspecificCapabilityParams {
	var ()
	return &GetaspecificCapabilityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificCapabilityParamsWithTimeout creates a new GetaspecificCapabilityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificCapabilityParamsWithTimeout(timeout time.Duration) *GetaspecificCapabilityParams {
	var ()
	return &GetaspecificCapabilityParams{

		timeout: timeout,
	}
}

// NewGetaspecificCapabilityParamsWithContext creates a new GetaspecificCapabilityParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificCapabilityParamsWithContext(ctx context.Context) *GetaspecificCapabilityParams {
	var ()
	return &GetaspecificCapabilityParams{

		Context: ctx,
	}
}

// NewGetaspecificCapabilityParamsWithHTTPClient creates a new GetaspecificCapabilityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificCapabilityParamsWithHTTPClient(client *http.Client) *GetaspecificCapabilityParams {
	var ()
	return &GetaspecificCapabilityParams{
		HTTPClient: client,
	}
}

/*GetaspecificCapabilityParams contains all the parameters to send to the API endpoint
for the getaspecific capability operation typically these are written to a http.Request
*/
type GetaspecificCapabilityParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific capability params
func (o *GetaspecificCapabilityParams) WithTimeout(timeout time.Duration) *GetaspecificCapabilityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific capability params
func (o *GetaspecificCapabilityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific capability params
func (o *GetaspecificCapabilityParams) WithContext(ctx context.Context) *GetaspecificCapabilityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific capability params
func (o *GetaspecificCapabilityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific capability params
func (o *GetaspecificCapabilityParams) WithHTTPClient(client *http.Client) *GetaspecificCapabilityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific capability params
func (o *GetaspecificCapabilityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific capability params
func (o *GetaspecificCapabilityParams) WithID(id string) *GetaspecificCapabilityParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific capability params
func (o *GetaspecificCapabilityParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificCapabilityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
