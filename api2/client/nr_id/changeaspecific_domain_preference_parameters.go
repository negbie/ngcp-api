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

// NewChangeaspecificDomainPreferenceParams creates a new ChangeaspecificDomainPreferenceParams object
// with the default values initialized.
func NewChangeaspecificDomainPreferenceParams() *ChangeaspecificDomainPreferenceParams {
	var ()
	return &ChangeaspecificDomainPreferenceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeaspecificDomainPreferenceParamsWithTimeout creates a new ChangeaspecificDomainPreferenceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeaspecificDomainPreferenceParamsWithTimeout(timeout time.Duration) *ChangeaspecificDomainPreferenceParams {
	var ()
	return &ChangeaspecificDomainPreferenceParams{

		timeout: timeout,
	}
}

// NewChangeaspecificDomainPreferenceParamsWithContext creates a new ChangeaspecificDomainPreferenceParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeaspecificDomainPreferenceParamsWithContext(ctx context.Context) *ChangeaspecificDomainPreferenceParams {
	var ()
	return &ChangeaspecificDomainPreferenceParams{

		Context: ctx,
	}
}

// NewChangeaspecificDomainPreferenceParamsWithHTTPClient creates a new ChangeaspecificDomainPreferenceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeaspecificDomainPreferenceParamsWithHTTPClient(client *http.Client) *ChangeaspecificDomainPreferenceParams {
	var ()
	return &ChangeaspecificDomainPreferenceParams{
		HTTPClient: client,
	}
}

/*ChangeaspecificDomainPreferenceParams contains all the parameters to send to the API endpoint
for the changeaspecific domain preference operation typically these are written to a http.Request
*/
type ChangeaspecificDomainPreferenceParams struct {

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

// WithTimeout adds the timeout to the changeaspecific domain preference params
func (o *ChangeaspecificDomainPreferenceParams) WithTimeout(timeout time.Duration) *ChangeaspecificDomainPreferenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the changeaspecific domain preference params
func (o *ChangeaspecificDomainPreferenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the changeaspecific domain preference params
func (o *ChangeaspecificDomainPreferenceParams) WithContext(ctx context.Context) *ChangeaspecificDomainPreferenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the changeaspecific domain preference params
func (o *ChangeaspecificDomainPreferenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the changeaspecific domain preference params
func (o *ChangeaspecificDomainPreferenceParams) WithHTTPClient(client *http.Client) *ChangeaspecificDomainPreferenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the changeaspecific domain preference params
func (o *ChangeaspecificDomainPreferenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the changeaspecific domain preference params
func (o *ChangeaspecificDomainPreferenceParams) WithBody(body string) *ChangeaspecificDomainPreferenceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the changeaspecific domain preference params
func (o *ChangeaspecificDomainPreferenceParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the changeaspecific domain preference params
func (o *ChangeaspecificDomainPreferenceParams) WithContentType(contentType string) *ChangeaspecificDomainPreferenceParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the changeaspecific domain preference params
func (o *ChangeaspecificDomainPreferenceParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the changeaspecific domain preference params
func (o *ChangeaspecificDomainPreferenceParams) WithID(id string) *ChangeaspecificDomainPreferenceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the changeaspecific domain preference params
func (o *ChangeaspecificDomainPreferenceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeaspecificDomainPreferenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
