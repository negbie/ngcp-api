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

// NewReplaceChangeaspecificLnpCarrierParams creates a new ReplaceChangeaspecificLnpCarrierParams object
// with the default values initialized.
func NewReplaceChangeaspecificLnpCarrierParams() *ReplaceChangeaspecificLnpCarrierParams {
	var ()
	return &ReplaceChangeaspecificLnpCarrierParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceChangeaspecificLnpCarrierParamsWithTimeout creates a new ReplaceChangeaspecificLnpCarrierParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceChangeaspecificLnpCarrierParamsWithTimeout(timeout time.Duration) *ReplaceChangeaspecificLnpCarrierParams {
	var ()
	return &ReplaceChangeaspecificLnpCarrierParams{

		timeout: timeout,
	}
}

// NewReplaceChangeaspecificLnpCarrierParamsWithContext creates a new ReplaceChangeaspecificLnpCarrierParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceChangeaspecificLnpCarrierParamsWithContext(ctx context.Context) *ReplaceChangeaspecificLnpCarrierParams {
	var ()
	return &ReplaceChangeaspecificLnpCarrierParams{

		Context: ctx,
	}
}

// NewReplaceChangeaspecificLnpCarrierParamsWithHTTPClient creates a new ReplaceChangeaspecificLnpCarrierParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceChangeaspecificLnpCarrierParamsWithHTTPClient(client *http.Client) *ReplaceChangeaspecificLnpCarrierParams {
	var ()
	return &ReplaceChangeaspecificLnpCarrierParams{
		HTTPClient: client,
	}
}

/*ReplaceChangeaspecificLnpCarrierParams contains all the parameters to send to the API endpoint
for the replace changeaspecific lnp carrier operation typically these are written to a http.Request
*/
type ReplaceChangeaspecificLnpCarrierParams struct {

	/*Body*/
	Body *Replace1changeaspecificLnpCarrierRequest
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

// WithTimeout adds the timeout to the replace changeaspecific lnp carrier params
func (o *ReplaceChangeaspecificLnpCarrierParams) WithTimeout(timeout time.Duration) *ReplaceChangeaspecificLnpCarrierParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace changeaspecific lnp carrier params
func (o *ReplaceChangeaspecificLnpCarrierParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace changeaspecific lnp carrier params
func (o *ReplaceChangeaspecificLnpCarrierParams) WithContext(ctx context.Context) *ReplaceChangeaspecificLnpCarrierParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace changeaspecific lnp carrier params
func (o *ReplaceChangeaspecificLnpCarrierParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace changeaspecific lnp carrier params
func (o *ReplaceChangeaspecificLnpCarrierParams) WithHTTPClient(client *http.Client) *ReplaceChangeaspecificLnpCarrierParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace changeaspecific lnp carrier params
func (o *ReplaceChangeaspecificLnpCarrierParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace changeaspecific lnp carrier params
func (o *ReplaceChangeaspecificLnpCarrierParams) WithBody(body *Replace1changeaspecificLnpCarrierRequest) *ReplaceChangeaspecificLnpCarrierParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace changeaspecific lnp carrier params
func (o *ReplaceChangeaspecificLnpCarrierParams) SetBody(body *Replace1changeaspecificLnpCarrierRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the replace changeaspecific lnp carrier params
func (o *ReplaceChangeaspecificLnpCarrierParams) WithContentType(contentType string) *ReplaceChangeaspecificLnpCarrierParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the replace changeaspecific lnp carrier params
func (o *ReplaceChangeaspecificLnpCarrierParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the replace changeaspecific lnp carrier params
func (o *ReplaceChangeaspecificLnpCarrierParams) WithID(id string) *ReplaceChangeaspecificLnpCarrierParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replace changeaspecific lnp carrier params
func (o *ReplaceChangeaspecificLnpCarrierParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceChangeaspecificLnpCarrierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
