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

// NewReplaceChangeaspecificRewriteRuleSetParams creates a new ReplaceChangeaspecificRewriteRuleSetParams object
// with the default values initialized.
func NewReplaceChangeaspecificRewriteRuleSetParams() *ReplaceChangeaspecificRewriteRuleSetParams {
	var ()
	return &ReplaceChangeaspecificRewriteRuleSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceChangeaspecificRewriteRuleSetParamsWithTimeout creates a new ReplaceChangeaspecificRewriteRuleSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceChangeaspecificRewriteRuleSetParamsWithTimeout(timeout time.Duration) *ReplaceChangeaspecificRewriteRuleSetParams {
	var ()
	return &ReplaceChangeaspecificRewriteRuleSetParams{

		timeout: timeout,
	}
}

// NewReplaceChangeaspecificRewriteRuleSetParamsWithContext creates a new ReplaceChangeaspecificRewriteRuleSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceChangeaspecificRewriteRuleSetParamsWithContext(ctx context.Context) *ReplaceChangeaspecificRewriteRuleSetParams {
	var ()
	return &ReplaceChangeaspecificRewriteRuleSetParams{

		Context: ctx,
	}
}

// NewReplaceChangeaspecificRewriteRuleSetParamsWithHTTPClient creates a new ReplaceChangeaspecificRewriteRuleSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceChangeaspecificRewriteRuleSetParamsWithHTTPClient(client *http.Client) *ReplaceChangeaspecificRewriteRuleSetParams {
	var ()
	return &ReplaceChangeaspecificRewriteRuleSetParams{
		HTTPClient: client,
	}
}

/*ReplaceChangeaspecificRewriteRuleSetParams contains all the parameters to send to the API endpoint
for the replace changeaspecific rewrite rule set operation typically these are written to a http.Request
*/
type ReplaceChangeaspecificRewriteRuleSetParams struct {

	/*Body*/
	Body *Replace1changeaspecificRewriteRuleSetRequest
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

// WithTimeout adds the timeout to the replace changeaspecific rewrite rule set params
func (o *ReplaceChangeaspecificRewriteRuleSetParams) WithTimeout(timeout time.Duration) *ReplaceChangeaspecificRewriteRuleSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace changeaspecific rewrite rule set params
func (o *ReplaceChangeaspecificRewriteRuleSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace changeaspecific rewrite rule set params
func (o *ReplaceChangeaspecificRewriteRuleSetParams) WithContext(ctx context.Context) *ReplaceChangeaspecificRewriteRuleSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace changeaspecific rewrite rule set params
func (o *ReplaceChangeaspecificRewriteRuleSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace changeaspecific rewrite rule set params
func (o *ReplaceChangeaspecificRewriteRuleSetParams) WithHTTPClient(client *http.Client) *ReplaceChangeaspecificRewriteRuleSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace changeaspecific rewrite rule set params
func (o *ReplaceChangeaspecificRewriteRuleSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace changeaspecific rewrite rule set params
func (o *ReplaceChangeaspecificRewriteRuleSetParams) WithBody(body *Replace1changeaspecificRewriteRuleSetRequest) *ReplaceChangeaspecificRewriteRuleSetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace changeaspecific rewrite rule set params
func (o *ReplaceChangeaspecificRewriteRuleSetParams) SetBody(body *Replace1changeaspecificRewriteRuleSetRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the replace changeaspecific rewrite rule set params
func (o *ReplaceChangeaspecificRewriteRuleSetParams) WithContentType(contentType string) *ReplaceChangeaspecificRewriteRuleSetParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the replace changeaspecific rewrite rule set params
func (o *ReplaceChangeaspecificRewriteRuleSetParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the replace changeaspecific rewrite rule set params
func (o *ReplaceChangeaspecificRewriteRuleSetParams) WithID(id string) *ReplaceChangeaspecificRewriteRuleSetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replace changeaspecific rewrite rule set params
func (o *ReplaceChangeaspecificRewriteRuleSetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceChangeaspecificRewriteRuleSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
