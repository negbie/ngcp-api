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

// NewChangeaspecificSubscriberProfileSetParams creates a new ChangeaspecificSubscriberProfileSetParams object
// with the default values initialized.
func NewChangeaspecificSubscriberProfileSetParams() *ChangeaspecificSubscriberProfileSetParams {
	var ()
	return &ChangeaspecificSubscriberProfileSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeaspecificSubscriberProfileSetParamsWithTimeout creates a new ChangeaspecificSubscriberProfileSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeaspecificSubscriberProfileSetParamsWithTimeout(timeout time.Duration) *ChangeaspecificSubscriberProfileSetParams {
	var ()
	return &ChangeaspecificSubscriberProfileSetParams{

		timeout: timeout,
	}
}

// NewChangeaspecificSubscriberProfileSetParamsWithContext creates a new ChangeaspecificSubscriberProfileSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeaspecificSubscriberProfileSetParamsWithContext(ctx context.Context) *ChangeaspecificSubscriberProfileSetParams {
	var ()
	return &ChangeaspecificSubscriberProfileSetParams{

		Context: ctx,
	}
}

// NewChangeaspecificSubscriberProfileSetParamsWithHTTPClient creates a new ChangeaspecificSubscriberProfileSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeaspecificSubscriberProfileSetParamsWithHTTPClient(client *http.Client) *ChangeaspecificSubscriberProfileSetParams {
	var ()
	return &ChangeaspecificSubscriberProfileSetParams{
		HTTPClient: client,
	}
}

/*ChangeaspecificSubscriberProfileSetParams contains all the parameters to send to the API endpoint
for the changeaspecific subscriber profile set operation typically these are written to a http.Request
*/
type ChangeaspecificSubscriberProfileSetParams struct {

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

// WithTimeout adds the timeout to the changeaspecific subscriber profile set params
func (o *ChangeaspecificSubscriberProfileSetParams) WithTimeout(timeout time.Duration) *ChangeaspecificSubscriberProfileSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the changeaspecific subscriber profile set params
func (o *ChangeaspecificSubscriberProfileSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the changeaspecific subscriber profile set params
func (o *ChangeaspecificSubscriberProfileSetParams) WithContext(ctx context.Context) *ChangeaspecificSubscriberProfileSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the changeaspecific subscriber profile set params
func (o *ChangeaspecificSubscriberProfileSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the changeaspecific subscriber profile set params
func (o *ChangeaspecificSubscriberProfileSetParams) WithHTTPClient(client *http.Client) *ChangeaspecificSubscriberProfileSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the changeaspecific subscriber profile set params
func (o *ChangeaspecificSubscriberProfileSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the changeaspecific subscriber profile set params
func (o *ChangeaspecificSubscriberProfileSetParams) WithBody(body string) *ChangeaspecificSubscriberProfileSetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the changeaspecific subscriber profile set params
func (o *ChangeaspecificSubscriberProfileSetParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the changeaspecific subscriber profile set params
func (o *ChangeaspecificSubscriberProfileSetParams) WithContentType(contentType string) *ChangeaspecificSubscriberProfileSetParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the changeaspecific subscriber profile set params
func (o *ChangeaspecificSubscriberProfileSetParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the changeaspecific subscriber profile set params
func (o *ChangeaspecificSubscriberProfileSetParams) WithID(id string) *ChangeaspecificSubscriberProfileSetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the changeaspecific subscriber profile set params
func (o *ChangeaspecificSubscriberProfileSetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeaspecificSubscriberProfileSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
