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

// NewReplaceChangeaspecificPeeringServerPreferenceParams creates a new ReplaceChangeaspecificPeeringServerPreferenceParams object
// with the default values initialized.
func NewReplaceChangeaspecificPeeringServerPreferenceParams() *ReplaceChangeaspecificPeeringServerPreferenceParams {
	var ()
	return &ReplaceChangeaspecificPeeringServerPreferenceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceChangeaspecificPeeringServerPreferenceParamsWithTimeout creates a new ReplaceChangeaspecificPeeringServerPreferenceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceChangeaspecificPeeringServerPreferenceParamsWithTimeout(timeout time.Duration) *ReplaceChangeaspecificPeeringServerPreferenceParams {
	var ()
	return &ReplaceChangeaspecificPeeringServerPreferenceParams{

		timeout: timeout,
	}
}

// NewReplaceChangeaspecificPeeringServerPreferenceParamsWithContext creates a new ReplaceChangeaspecificPeeringServerPreferenceParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceChangeaspecificPeeringServerPreferenceParamsWithContext(ctx context.Context) *ReplaceChangeaspecificPeeringServerPreferenceParams {
	var ()
	return &ReplaceChangeaspecificPeeringServerPreferenceParams{

		Context: ctx,
	}
}

// NewReplaceChangeaspecificPeeringServerPreferenceParamsWithHTTPClient creates a new ReplaceChangeaspecificPeeringServerPreferenceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceChangeaspecificPeeringServerPreferenceParamsWithHTTPClient(client *http.Client) *ReplaceChangeaspecificPeeringServerPreferenceParams {
	var ()
	return &ReplaceChangeaspecificPeeringServerPreferenceParams{
		HTTPClient: client,
	}
}

/*ReplaceChangeaspecificPeeringServerPreferenceParams contains all the parameters to send to the API endpoint
for the replace changeaspecific peering server preference operation typically these are written to a http.Request
*/
type ReplaceChangeaspecificPeeringServerPreferenceParams struct {

	/*Body*/
	Body string
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

// WithTimeout adds the timeout to the replace changeaspecific peering server preference params
func (o *ReplaceChangeaspecificPeeringServerPreferenceParams) WithTimeout(timeout time.Duration) *ReplaceChangeaspecificPeeringServerPreferenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace changeaspecific peering server preference params
func (o *ReplaceChangeaspecificPeeringServerPreferenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace changeaspecific peering server preference params
func (o *ReplaceChangeaspecificPeeringServerPreferenceParams) WithContext(ctx context.Context) *ReplaceChangeaspecificPeeringServerPreferenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace changeaspecific peering server preference params
func (o *ReplaceChangeaspecificPeeringServerPreferenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace changeaspecific peering server preference params
func (o *ReplaceChangeaspecificPeeringServerPreferenceParams) WithHTTPClient(client *http.Client) *ReplaceChangeaspecificPeeringServerPreferenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace changeaspecific peering server preference params
func (o *ReplaceChangeaspecificPeeringServerPreferenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace changeaspecific peering server preference params
func (o *ReplaceChangeaspecificPeeringServerPreferenceParams) WithBody(body string) *ReplaceChangeaspecificPeeringServerPreferenceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace changeaspecific peering server preference params
func (o *ReplaceChangeaspecificPeeringServerPreferenceParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the replace changeaspecific peering server preference params
func (o *ReplaceChangeaspecificPeeringServerPreferenceParams) WithContentType(contentType string) *ReplaceChangeaspecificPeeringServerPreferenceParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the replace changeaspecific peering server preference params
func (o *ReplaceChangeaspecificPeeringServerPreferenceParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the replace changeaspecific peering server preference params
func (o *ReplaceChangeaspecificPeeringServerPreferenceParams) WithID(id string) *ReplaceChangeaspecificPeeringServerPreferenceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replace changeaspecific peering server preference params
func (o *ReplaceChangeaspecificPeeringServerPreferenceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceChangeaspecificPeeringServerPreferenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
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
