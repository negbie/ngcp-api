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

// NewReplaceChangeaspecificDomainPreferenceParams creates a new ReplaceChangeaspecificDomainPreferenceParams object
// with the default values initialized.
func NewReplaceChangeaspecificDomainPreferenceParams() *ReplaceChangeaspecificDomainPreferenceParams {
	var ()
	return &ReplaceChangeaspecificDomainPreferenceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceChangeaspecificDomainPreferenceParamsWithTimeout creates a new ReplaceChangeaspecificDomainPreferenceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceChangeaspecificDomainPreferenceParamsWithTimeout(timeout time.Duration) *ReplaceChangeaspecificDomainPreferenceParams {
	var ()
	return &ReplaceChangeaspecificDomainPreferenceParams{

		timeout: timeout,
	}
}

// NewReplaceChangeaspecificDomainPreferenceParamsWithContext creates a new ReplaceChangeaspecificDomainPreferenceParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceChangeaspecificDomainPreferenceParamsWithContext(ctx context.Context) *ReplaceChangeaspecificDomainPreferenceParams {
	var ()
	return &ReplaceChangeaspecificDomainPreferenceParams{

		Context: ctx,
	}
}

// NewReplaceChangeaspecificDomainPreferenceParamsWithHTTPClient creates a new ReplaceChangeaspecificDomainPreferenceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceChangeaspecificDomainPreferenceParamsWithHTTPClient(client *http.Client) *ReplaceChangeaspecificDomainPreferenceParams {
	var ()
	return &ReplaceChangeaspecificDomainPreferenceParams{
		HTTPClient: client,
	}
}

/*ReplaceChangeaspecificDomainPreferenceParams contains all the parameters to send to the API endpoint
for the replace changeaspecific domain preference operation typically these are written to a http.Request
*/
type ReplaceChangeaspecificDomainPreferenceParams struct {

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

// WithTimeout adds the timeout to the replace changeaspecific domain preference params
func (o *ReplaceChangeaspecificDomainPreferenceParams) WithTimeout(timeout time.Duration) *ReplaceChangeaspecificDomainPreferenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace changeaspecific domain preference params
func (o *ReplaceChangeaspecificDomainPreferenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace changeaspecific domain preference params
func (o *ReplaceChangeaspecificDomainPreferenceParams) WithContext(ctx context.Context) *ReplaceChangeaspecificDomainPreferenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace changeaspecific domain preference params
func (o *ReplaceChangeaspecificDomainPreferenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace changeaspecific domain preference params
func (o *ReplaceChangeaspecificDomainPreferenceParams) WithHTTPClient(client *http.Client) *ReplaceChangeaspecificDomainPreferenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace changeaspecific domain preference params
func (o *ReplaceChangeaspecificDomainPreferenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace changeaspecific domain preference params
func (o *ReplaceChangeaspecificDomainPreferenceParams) WithBody(body string) *ReplaceChangeaspecificDomainPreferenceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace changeaspecific domain preference params
func (o *ReplaceChangeaspecificDomainPreferenceParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the replace changeaspecific domain preference params
func (o *ReplaceChangeaspecificDomainPreferenceParams) WithContentType(contentType string) *ReplaceChangeaspecificDomainPreferenceParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the replace changeaspecific domain preference params
func (o *ReplaceChangeaspecificDomainPreferenceParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the replace changeaspecific domain preference params
func (o *ReplaceChangeaspecificDomainPreferenceParams) WithID(id string) *ReplaceChangeaspecificDomainPreferenceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replace changeaspecific domain preference params
func (o *ReplaceChangeaspecificDomainPreferenceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceChangeaspecificDomainPreferenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
