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

// NewGetaspecificPeeringGroupParams creates a new GetaspecificPeeringGroupParams object
// with the default values initialized.
func NewGetaspecificPeeringGroupParams() *GetaspecificPeeringGroupParams {
	var ()
	return &GetaspecificPeeringGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificPeeringGroupParamsWithTimeout creates a new GetaspecificPeeringGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificPeeringGroupParamsWithTimeout(timeout time.Duration) *GetaspecificPeeringGroupParams {
	var ()
	return &GetaspecificPeeringGroupParams{

		timeout: timeout,
	}
}

// NewGetaspecificPeeringGroupParamsWithContext creates a new GetaspecificPeeringGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificPeeringGroupParamsWithContext(ctx context.Context) *GetaspecificPeeringGroupParams {
	var ()
	return &GetaspecificPeeringGroupParams{

		Context: ctx,
	}
}

// NewGetaspecificPeeringGroupParamsWithHTTPClient creates a new GetaspecificPeeringGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificPeeringGroupParamsWithHTTPClient(client *http.Client) *GetaspecificPeeringGroupParams {
	var ()
	return &GetaspecificPeeringGroupParams{
		HTTPClient: client,
	}
}

/*GetaspecificPeeringGroupParams contains all the parameters to send to the API endpoint
for the getaspecific peering group operation typically these are written to a http.Request
*/
type GetaspecificPeeringGroupParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific peering group params
func (o *GetaspecificPeeringGroupParams) WithTimeout(timeout time.Duration) *GetaspecificPeeringGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific peering group params
func (o *GetaspecificPeeringGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific peering group params
func (o *GetaspecificPeeringGroupParams) WithContext(ctx context.Context) *GetaspecificPeeringGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific peering group params
func (o *GetaspecificPeeringGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific peering group params
func (o *GetaspecificPeeringGroupParams) WithHTTPClient(client *http.Client) *GetaspecificPeeringGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific peering group params
func (o *GetaspecificPeeringGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific peering group params
func (o *GetaspecificPeeringGroupParams) WithID(id string) *GetaspecificPeeringGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific peering group params
func (o *GetaspecificPeeringGroupParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificPeeringGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
