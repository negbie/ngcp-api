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

// NewReplaceChangeaspecificSystemContactParams creates a new ReplaceChangeaspecificSystemContactParams object
// with the default values initialized.
func NewReplaceChangeaspecificSystemContactParams() *ReplaceChangeaspecificSystemContactParams {
	var ()
	return &ReplaceChangeaspecificSystemContactParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceChangeaspecificSystemContactParamsWithTimeout creates a new ReplaceChangeaspecificSystemContactParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceChangeaspecificSystemContactParamsWithTimeout(timeout time.Duration) *ReplaceChangeaspecificSystemContactParams {
	var ()
	return &ReplaceChangeaspecificSystemContactParams{

		timeout: timeout,
	}
}

// NewReplaceChangeaspecificSystemContactParamsWithContext creates a new ReplaceChangeaspecificSystemContactParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceChangeaspecificSystemContactParamsWithContext(ctx context.Context) *ReplaceChangeaspecificSystemContactParams {
	var ()
	return &ReplaceChangeaspecificSystemContactParams{

		Context: ctx,
	}
}

// NewReplaceChangeaspecificSystemContactParamsWithHTTPClient creates a new ReplaceChangeaspecificSystemContactParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceChangeaspecificSystemContactParamsWithHTTPClient(client *http.Client) *ReplaceChangeaspecificSystemContactParams {
	var ()
	return &ReplaceChangeaspecificSystemContactParams{
		HTTPClient: client,
	}
}

/*ReplaceChangeaspecificSystemContactParams contains all the parameters to send to the API endpoint
for the replace changeaspecific system contact operation typically these are written to a http.Request
*/
type ReplaceChangeaspecificSystemContactParams struct {

	/*Body*/
	Body *Replace1changeaspecificSystemContactRequest
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

// WithTimeout adds the timeout to the replace changeaspecific system contact params
func (o *ReplaceChangeaspecificSystemContactParams) WithTimeout(timeout time.Duration) *ReplaceChangeaspecificSystemContactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace changeaspecific system contact params
func (o *ReplaceChangeaspecificSystemContactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace changeaspecific system contact params
func (o *ReplaceChangeaspecificSystemContactParams) WithContext(ctx context.Context) *ReplaceChangeaspecificSystemContactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace changeaspecific system contact params
func (o *ReplaceChangeaspecificSystemContactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace changeaspecific system contact params
func (o *ReplaceChangeaspecificSystemContactParams) WithHTTPClient(client *http.Client) *ReplaceChangeaspecificSystemContactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace changeaspecific system contact params
func (o *ReplaceChangeaspecificSystemContactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace changeaspecific system contact params
func (o *ReplaceChangeaspecificSystemContactParams) WithBody(body *Replace1changeaspecificSystemContactRequest) *ReplaceChangeaspecificSystemContactParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace changeaspecific system contact params
func (o *ReplaceChangeaspecificSystemContactParams) SetBody(body *Replace1changeaspecificSystemContactRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the replace changeaspecific system contact params
func (o *ReplaceChangeaspecificSystemContactParams) WithContentType(contentType string) *ReplaceChangeaspecificSystemContactParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the replace changeaspecific system contact params
func (o *ReplaceChangeaspecificSystemContactParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the replace changeaspecific system contact params
func (o *ReplaceChangeaspecificSystemContactParams) WithID(id string) *ReplaceChangeaspecificSystemContactParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replace changeaspecific system contact params
func (o *ReplaceChangeaspecificSystemContactParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceChangeaspecificSystemContactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
