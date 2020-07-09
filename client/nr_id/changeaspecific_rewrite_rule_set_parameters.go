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

// NewChangeaspecificRewriteRuleSetParams creates a new ChangeaspecificRewriteRuleSetParams object
// with the default values initialized.
func NewChangeaspecificRewriteRuleSetParams() *ChangeaspecificRewriteRuleSetParams {
	var ()
	return &ChangeaspecificRewriteRuleSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeaspecificRewriteRuleSetParamsWithTimeout creates a new ChangeaspecificRewriteRuleSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeaspecificRewriteRuleSetParamsWithTimeout(timeout time.Duration) *ChangeaspecificRewriteRuleSetParams {
	var ()
	return &ChangeaspecificRewriteRuleSetParams{

		timeout: timeout,
	}
}

// NewChangeaspecificRewriteRuleSetParamsWithContext creates a new ChangeaspecificRewriteRuleSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeaspecificRewriteRuleSetParamsWithContext(ctx context.Context) *ChangeaspecificRewriteRuleSetParams {
	var ()
	return &ChangeaspecificRewriteRuleSetParams{

		Context: ctx,
	}
}

// NewChangeaspecificRewriteRuleSetParamsWithHTTPClient creates a new ChangeaspecificRewriteRuleSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeaspecificRewriteRuleSetParamsWithHTTPClient(client *http.Client) *ChangeaspecificRewriteRuleSetParams {
	var ()
	return &ChangeaspecificRewriteRuleSetParams{
		HTTPClient: client,
	}
}

/*ChangeaspecificRewriteRuleSetParams contains all the parameters to send to the API endpoint
for the changeaspecific rewrite rule set operation typically these are written to a http.Request
*/
type ChangeaspecificRewriteRuleSetParams struct {

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

// WithTimeout adds the timeout to the changeaspecific rewrite rule set params
func (o *ChangeaspecificRewriteRuleSetParams) WithTimeout(timeout time.Duration) *ChangeaspecificRewriteRuleSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the changeaspecific rewrite rule set params
func (o *ChangeaspecificRewriteRuleSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the changeaspecific rewrite rule set params
func (o *ChangeaspecificRewriteRuleSetParams) WithContext(ctx context.Context) *ChangeaspecificRewriteRuleSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the changeaspecific rewrite rule set params
func (o *ChangeaspecificRewriteRuleSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the changeaspecific rewrite rule set params
func (o *ChangeaspecificRewriteRuleSetParams) WithHTTPClient(client *http.Client) *ChangeaspecificRewriteRuleSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the changeaspecific rewrite rule set params
func (o *ChangeaspecificRewriteRuleSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the changeaspecific rewrite rule set params
func (o *ChangeaspecificRewriteRuleSetParams) WithBody(body string) *ChangeaspecificRewriteRuleSetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the changeaspecific rewrite rule set params
func (o *ChangeaspecificRewriteRuleSetParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the changeaspecific rewrite rule set params
func (o *ChangeaspecificRewriteRuleSetParams) WithContentType(contentType string) *ChangeaspecificRewriteRuleSetParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the changeaspecific rewrite rule set params
func (o *ChangeaspecificRewriteRuleSetParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the changeaspecific rewrite rule set params
func (o *ChangeaspecificRewriteRuleSetParams) WithID(id string) *ChangeaspecificRewriteRuleSetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the changeaspecific rewrite rule set params
func (o *ChangeaspecificRewriteRuleSetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeaspecificRewriteRuleSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
