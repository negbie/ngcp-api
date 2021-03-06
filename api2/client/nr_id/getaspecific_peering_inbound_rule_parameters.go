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

// NewGetaspecificPeeringInboundRuleParams creates a new GetaspecificPeeringInboundRuleParams object
// with the default values initialized.
func NewGetaspecificPeeringInboundRuleParams() *GetaspecificPeeringInboundRuleParams {
	var ()
	return &GetaspecificPeeringInboundRuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificPeeringInboundRuleParamsWithTimeout creates a new GetaspecificPeeringInboundRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificPeeringInboundRuleParamsWithTimeout(timeout time.Duration) *GetaspecificPeeringInboundRuleParams {
	var ()
	return &GetaspecificPeeringInboundRuleParams{

		timeout: timeout,
	}
}

// NewGetaspecificPeeringInboundRuleParamsWithContext creates a new GetaspecificPeeringInboundRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificPeeringInboundRuleParamsWithContext(ctx context.Context) *GetaspecificPeeringInboundRuleParams {
	var ()
	return &GetaspecificPeeringInboundRuleParams{

		Context: ctx,
	}
}

// NewGetaspecificPeeringInboundRuleParamsWithHTTPClient creates a new GetaspecificPeeringInboundRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificPeeringInboundRuleParamsWithHTTPClient(client *http.Client) *GetaspecificPeeringInboundRuleParams {
	var ()
	return &GetaspecificPeeringInboundRuleParams{
		HTTPClient: client,
	}
}

/*GetaspecificPeeringInboundRuleParams contains all the parameters to send to the API endpoint
for the getaspecific peering inbound rule operation typically these are written to a http.Request
*/
type GetaspecificPeeringInboundRuleParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific peering inbound rule params
func (o *GetaspecificPeeringInboundRuleParams) WithTimeout(timeout time.Duration) *GetaspecificPeeringInboundRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific peering inbound rule params
func (o *GetaspecificPeeringInboundRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific peering inbound rule params
func (o *GetaspecificPeeringInboundRuleParams) WithContext(ctx context.Context) *GetaspecificPeeringInboundRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific peering inbound rule params
func (o *GetaspecificPeeringInboundRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific peering inbound rule params
func (o *GetaspecificPeeringInboundRuleParams) WithHTTPClient(client *http.Client) *GetaspecificPeeringInboundRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific peering inbound rule params
func (o *GetaspecificPeeringInboundRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific peering inbound rule params
func (o *GetaspecificPeeringInboundRuleParams) WithID(id string) *GetaspecificPeeringInboundRuleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific peering inbound rule params
func (o *GetaspecificPeeringInboundRuleParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificPeeringInboundRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
