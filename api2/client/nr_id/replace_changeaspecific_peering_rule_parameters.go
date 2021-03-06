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

// NewReplaceChangeaspecificPeeringRuleParams creates a new ReplaceChangeaspecificPeeringRuleParams object
// with the default values initialized.
func NewReplaceChangeaspecificPeeringRuleParams() *ReplaceChangeaspecificPeeringRuleParams {
	var ()
	return &ReplaceChangeaspecificPeeringRuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceChangeaspecificPeeringRuleParamsWithTimeout creates a new ReplaceChangeaspecificPeeringRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceChangeaspecificPeeringRuleParamsWithTimeout(timeout time.Duration) *ReplaceChangeaspecificPeeringRuleParams {
	var ()
	return &ReplaceChangeaspecificPeeringRuleParams{

		timeout: timeout,
	}
}

// NewReplaceChangeaspecificPeeringRuleParamsWithContext creates a new ReplaceChangeaspecificPeeringRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceChangeaspecificPeeringRuleParamsWithContext(ctx context.Context) *ReplaceChangeaspecificPeeringRuleParams {
	var ()
	return &ReplaceChangeaspecificPeeringRuleParams{

		Context: ctx,
	}
}

// NewReplaceChangeaspecificPeeringRuleParamsWithHTTPClient creates a new ReplaceChangeaspecificPeeringRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceChangeaspecificPeeringRuleParamsWithHTTPClient(client *http.Client) *ReplaceChangeaspecificPeeringRuleParams {
	var ()
	return &ReplaceChangeaspecificPeeringRuleParams{
		HTTPClient: client,
	}
}

/*ReplaceChangeaspecificPeeringRuleParams contains all the parameters to send to the API endpoint
for the replace changeaspecific peering rule operation typically these are written to a http.Request
*/
type ReplaceChangeaspecificPeeringRuleParams struct {

	/*Body*/
	Body *Replace1changeaspecificPeeringRuleRequest
	/*ContentType*/
	ContentType string
	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace changeaspecific peering rule params
func (o *ReplaceChangeaspecificPeeringRuleParams) WithTimeout(timeout time.Duration) *ReplaceChangeaspecificPeeringRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace changeaspecific peering rule params
func (o *ReplaceChangeaspecificPeeringRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace changeaspecific peering rule params
func (o *ReplaceChangeaspecificPeeringRuleParams) WithContext(ctx context.Context) *ReplaceChangeaspecificPeeringRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace changeaspecific peering rule params
func (o *ReplaceChangeaspecificPeeringRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace changeaspecific peering rule params
func (o *ReplaceChangeaspecificPeeringRuleParams) WithHTTPClient(client *http.Client) *ReplaceChangeaspecificPeeringRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace changeaspecific peering rule params
func (o *ReplaceChangeaspecificPeeringRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace changeaspecific peering rule params
func (o *ReplaceChangeaspecificPeeringRuleParams) WithBody(body *Replace1changeaspecificPeeringRuleRequest) *ReplaceChangeaspecificPeeringRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace changeaspecific peering rule params
func (o *ReplaceChangeaspecificPeeringRuleParams) SetBody(body *Replace1changeaspecificPeeringRuleRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the replace changeaspecific peering rule params
func (o *ReplaceChangeaspecificPeeringRuleParams) WithContentType(contentType string) *ReplaceChangeaspecificPeeringRuleParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the replace changeaspecific peering rule params
func (o *ReplaceChangeaspecificPeeringRuleParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the replace changeaspecific peering rule params
func (o *ReplaceChangeaspecificPeeringRuleParams) WithID(id string) *ReplaceChangeaspecificPeeringRuleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replace changeaspecific peering rule params
func (o *ReplaceChangeaspecificPeeringRuleParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceChangeaspecificPeeringRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// header param Content-Type
	if err := r.SetHeaderParam("Content-Type", o.ContentType); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
