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

// NewChangeaspecificSubscriberProfileParams creates a new ChangeaspecificSubscriberProfileParams object
// with the default values initialized.
func NewChangeaspecificSubscriberProfileParams() *ChangeaspecificSubscriberProfileParams {
	var ()
	return &ChangeaspecificSubscriberProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeaspecificSubscriberProfileParamsWithTimeout creates a new ChangeaspecificSubscriberProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeaspecificSubscriberProfileParamsWithTimeout(timeout time.Duration) *ChangeaspecificSubscriberProfileParams {
	var ()
	return &ChangeaspecificSubscriberProfileParams{

		timeout: timeout,
	}
}

// NewChangeaspecificSubscriberProfileParamsWithContext creates a new ChangeaspecificSubscriberProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeaspecificSubscriberProfileParamsWithContext(ctx context.Context) *ChangeaspecificSubscriberProfileParams {
	var ()
	return &ChangeaspecificSubscriberProfileParams{

		Context: ctx,
	}
}

// NewChangeaspecificSubscriberProfileParamsWithHTTPClient creates a new ChangeaspecificSubscriberProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeaspecificSubscriberProfileParamsWithHTTPClient(client *http.Client) *ChangeaspecificSubscriberProfileParams {
	var ()
	return &ChangeaspecificSubscriberProfileParams{
		HTTPClient: client,
	}
}

/*ChangeaspecificSubscriberProfileParams contains all the parameters to send to the API endpoint
for the changeaspecific subscriber profile operation typically these are written to a http.Request
*/
type ChangeaspecificSubscriberProfileParams struct {

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

// WithTimeout adds the timeout to the changeaspecific subscriber profile params
func (o *ChangeaspecificSubscriberProfileParams) WithTimeout(timeout time.Duration) *ChangeaspecificSubscriberProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the changeaspecific subscriber profile params
func (o *ChangeaspecificSubscriberProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the changeaspecific subscriber profile params
func (o *ChangeaspecificSubscriberProfileParams) WithContext(ctx context.Context) *ChangeaspecificSubscriberProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the changeaspecific subscriber profile params
func (o *ChangeaspecificSubscriberProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the changeaspecific subscriber profile params
func (o *ChangeaspecificSubscriberProfileParams) WithHTTPClient(client *http.Client) *ChangeaspecificSubscriberProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the changeaspecific subscriber profile params
func (o *ChangeaspecificSubscriberProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the changeaspecific subscriber profile params
func (o *ChangeaspecificSubscriberProfileParams) WithBody(body string) *ChangeaspecificSubscriberProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the changeaspecific subscriber profile params
func (o *ChangeaspecificSubscriberProfileParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the changeaspecific subscriber profile params
func (o *ChangeaspecificSubscriberProfileParams) WithContentType(contentType string) *ChangeaspecificSubscriberProfileParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the changeaspecific subscriber profile params
func (o *ChangeaspecificSubscriberProfileParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the changeaspecific subscriber profile params
func (o *ChangeaspecificSubscriberProfileParams) WithID(id string) *ChangeaspecificSubscriberProfileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the changeaspecific subscriber profile params
func (o *ChangeaspecificSubscriberProfileParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeaspecificSubscriberProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
