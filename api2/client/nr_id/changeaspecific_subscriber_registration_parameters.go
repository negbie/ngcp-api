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

// NewChangeaspecificSubscriberRegistrationParams creates a new ChangeaspecificSubscriberRegistrationParams object
// with the default values initialized.
func NewChangeaspecificSubscriberRegistrationParams() *ChangeaspecificSubscriberRegistrationParams {
	var ()
	return &ChangeaspecificSubscriberRegistrationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeaspecificSubscriberRegistrationParamsWithTimeout creates a new ChangeaspecificSubscriberRegistrationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeaspecificSubscriberRegistrationParamsWithTimeout(timeout time.Duration) *ChangeaspecificSubscriberRegistrationParams {
	var ()
	return &ChangeaspecificSubscriberRegistrationParams{

		timeout: timeout,
	}
}

// NewChangeaspecificSubscriberRegistrationParamsWithContext creates a new ChangeaspecificSubscriberRegistrationParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeaspecificSubscriberRegistrationParamsWithContext(ctx context.Context) *ChangeaspecificSubscriberRegistrationParams {
	var ()
	return &ChangeaspecificSubscriberRegistrationParams{

		Context: ctx,
	}
}

// NewChangeaspecificSubscriberRegistrationParamsWithHTTPClient creates a new ChangeaspecificSubscriberRegistrationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeaspecificSubscriberRegistrationParamsWithHTTPClient(client *http.Client) *ChangeaspecificSubscriberRegistrationParams {
	var ()
	return &ChangeaspecificSubscriberRegistrationParams{
		HTTPClient: client,
	}
}

/*ChangeaspecificSubscriberRegistrationParams contains all the parameters to send to the API endpoint
for the changeaspecific subscriber registration operation typically these are written to a http.Request
*/
type ChangeaspecificSubscriberRegistrationParams struct {

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

// WithTimeout adds the timeout to the changeaspecific subscriber registration params
func (o *ChangeaspecificSubscriberRegistrationParams) WithTimeout(timeout time.Duration) *ChangeaspecificSubscriberRegistrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the changeaspecific subscriber registration params
func (o *ChangeaspecificSubscriberRegistrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the changeaspecific subscriber registration params
func (o *ChangeaspecificSubscriberRegistrationParams) WithContext(ctx context.Context) *ChangeaspecificSubscriberRegistrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the changeaspecific subscriber registration params
func (o *ChangeaspecificSubscriberRegistrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the changeaspecific subscriber registration params
func (o *ChangeaspecificSubscriberRegistrationParams) WithHTTPClient(client *http.Client) *ChangeaspecificSubscriberRegistrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the changeaspecific subscriber registration params
func (o *ChangeaspecificSubscriberRegistrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the changeaspecific subscriber registration params
func (o *ChangeaspecificSubscriberRegistrationParams) WithBody(body string) *ChangeaspecificSubscriberRegistrationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the changeaspecific subscriber registration params
func (o *ChangeaspecificSubscriberRegistrationParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the changeaspecific subscriber registration params
func (o *ChangeaspecificSubscriberRegistrationParams) WithContentType(contentType string) *ChangeaspecificSubscriberRegistrationParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the changeaspecific subscriber registration params
func (o *ChangeaspecificSubscriberRegistrationParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the changeaspecific subscriber registration params
func (o *ChangeaspecificSubscriberRegistrationParams) WithID(id string) *ChangeaspecificSubscriberRegistrationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the changeaspecific subscriber registration params
func (o *ChangeaspecificSubscriberRegistrationParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeaspecificSubscriberRegistrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
