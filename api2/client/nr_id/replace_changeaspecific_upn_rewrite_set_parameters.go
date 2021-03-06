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

// NewReplaceChangeaspecificUpnRewriteSetParams creates a new ReplaceChangeaspecificUpnRewriteSetParams object
// with the default values initialized.
func NewReplaceChangeaspecificUpnRewriteSetParams() *ReplaceChangeaspecificUpnRewriteSetParams {
	var ()
	return &ReplaceChangeaspecificUpnRewriteSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceChangeaspecificUpnRewriteSetParamsWithTimeout creates a new ReplaceChangeaspecificUpnRewriteSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceChangeaspecificUpnRewriteSetParamsWithTimeout(timeout time.Duration) *ReplaceChangeaspecificUpnRewriteSetParams {
	var ()
	return &ReplaceChangeaspecificUpnRewriteSetParams{

		timeout: timeout,
	}
}

// NewReplaceChangeaspecificUpnRewriteSetParamsWithContext creates a new ReplaceChangeaspecificUpnRewriteSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceChangeaspecificUpnRewriteSetParamsWithContext(ctx context.Context) *ReplaceChangeaspecificUpnRewriteSetParams {
	var ()
	return &ReplaceChangeaspecificUpnRewriteSetParams{

		Context: ctx,
	}
}

// NewReplaceChangeaspecificUpnRewriteSetParamsWithHTTPClient creates a new ReplaceChangeaspecificUpnRewriteSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceChangeaspecificUpnRewriteSetParamsWithHTTPClient(client *http.Client) *ReplaceChangeaspecificUpnRewriteSetParams {
	var ()
	return &ReplaceChangeaspecificUpnRewriteSetParams{
		HTTPClient: client,
	}
}

/*ReplaceChangeaspecificUpnRewriteSetParams contains all the parameters to send to the API endpoint
for the replace changeaspecific upn rewrite set operation typically these are written to a http.Request
*/
type ReplaceChangeaspecificUpnRewriteSetParams struct {

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

// WithTimeout adds the timeout to the replace changeaspecific upn rewrite set params
func (o *ReplaceChangeaspecificUpnRewriteSetParams) WithTimeout(timeout time.Duration) *ReplaceChangeaspecificUpnRewriteSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace changeaspecific upn rewrite set params
func (o *ReplaceChangeaspecificUpnRewriteSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace changeaspecific upn rewrite set params
func (o *ReplaceChangeaspecificUpnRewriteSetParams) WithContext(ctx context.Context) *ReplaceChangeaspecificUpnRewriteSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace changeaspecific upn rewrite set params
func (o *ReplaceChangeaspecificUpnRewriteSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace changeaspecific upn rewrite set params
func (o *ReplaceChangeaspecificUpnRewriteSetParams) WithHTTPClient(client *http.Client) *ReplaceChangeaspecificUpnRewriteSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace changeaspecific upn rewrite set params
func (o *ReplaceChangeaspecificUpnRewriteSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace changeaspecific upn rewrite set params
func (o *ReplaceChangeaspecificUpnRewriteSetParams) WithBody(body string) *ReplaceChangeaspecificUpnRewriteSetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace changeaspecific upn rewrite set params
func (o *ReplaceChangeaspecificUpnRewriteSetParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the replace changeaspecific upn rewrite set params
func (o *ReplaceChangeaspecificUpnRewriteSetParams) WithContentType(contentType string) *ReplaceChangeaspecificUpnRewriteSetParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the replace changeaspecific upn rewrite set params
func (o *ReplaceChangeaspecificUpnRewriteSetParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the replace changeaspecific upn rewrite set params
func (o *ReplaceChangeaspecificUpnRewriteSetParams) WithID(id string) *ReplaceChangeaspecificUpnRewriteSetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replace changeaspecific upn rewrite set params
func (o *ReplaceChangeaspecificUpnRewriteSetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceChangeaspecificUpnRewriteSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
