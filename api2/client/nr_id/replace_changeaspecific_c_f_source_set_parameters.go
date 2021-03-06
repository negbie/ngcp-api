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

// NewReplaceChangeaspecificCFSourceSetParams creates a new ReplaceChangeaspecificCFSourceSetParams object
// with the default values initialized.
func NewReplaceChangeaspecificCFSourceSetParams() *ReplaceChangeaspecificCFSourceSetParams {
	var ()
	return &ReplaceChangeaspecificCFSourceSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceChangeaspecificCFSourceSetParamsWithTimeout creates a new ReplaceChangeaspecificCFSourceSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceChangeaspecificCFSourceSetParamsWithTimeout(timeout time.Duration) *ReplaceChangeaspecificCFSourceSetParams {
	var ()
	return &ReplaceChangeaspecificCFSourceSetParams{

		timeout: timeout,
	}
}

// NewReplaceChangeaspecificCFSourceSetParamsWithContext creates a new ReplaceChangeaspecificCFSourceSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceChangeaspecificCFSourceSetParamsWithContext(ctx context.Context) *ReplaceChangeaspecificCFSourceSetParams {
	var ()
	return &ReplaceChangeaspecificCFSourceSetParams{

		Context: ctx,
	}
}

// NewReplaceChangeaspecificCFSourceSetParamsWithHTTPClient creates a new ReplaceChangeaspecificCFSourceSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceChangeaspecificCFSourceSetParamsWithHTTPClient(client *http.Client) *ReplaceChangeaspecificCFSourceSetParams {
	var ()
	return &ReplaceChangeaspecificCFSourceSetParams{
		HTTPClient: client,
	}
}

/*ReplaceChangeaspecificCFSourceSetParams contains all the parameters to send to the API endpoint
for the replace changeaspecific c f source set operation typically these are written to a http.Request
*/
type ReplaceChangeaspecificCFSourceSetParams struct {

	/*Body*/
	Body *Replace1changeaspecificCFSourceSetRequest
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

// WithTimeout adds the timeout to the replace changeaspecific c f source set params
func (o *ReplaceChangeaspecificCFSourceSetParams) WithTimeout(timeout time.Duration) *ReplaceChangeaspecificCFSourceSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace changeaspecific c f source set params
func (o *ReplaceChangeaspecificCFSourceSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace changeaspecific c f source set params
func (o *ReplaceChangeaspecificCFSourceSetParams) WithContext(ctx context.Context) *ReplaceChangeaspecificCFSourceSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace changeaspecific c f source set params
func (o *ReplaceChangeaspecificCFSourceSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace changeaspecific c f source set params
func (o *ReplaceChangeaspecificCFSourceSetParams) WithHTTPClient(client *http.Client) *ReplaceChangeaspecificCFSourceSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace changeaspecific c f source set params
func (o *ReplaceChangeaspecificCFSourceSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace changeaspecific c f source set params
func (o *ReplaceChangeaspecificCFSourceSetParams) WithBody(body *Replace1changeaspecificCFSourceSetRequest) *ReplaceChangeaspecificCFSourceSetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace changeaspecific c f source set params
func (o *ReplaceChangeaspecificCFSourceSetParams) SetBody(body *Replace1changeaspecificCFSourceSetRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the replace changeaspecific c f source set params
func (o *ReplaceChangeaspecificCFSourceSetParams) WithContentType(contentType string) *ReplaceChangeaspecificCFSourceSetParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the replace changeaspecific c f source set params
func (o *ReplaceChangeaspecificCFSourceSetParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the replace changeaspecific c f source set params
func (o *ReplaceChangeaspecificCFSourceSetParams) WithID(id string) *ReplaceChangeaspecificCFSourceSetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replace changeaspecific c f source set params
func (o *ReplaceChangeaspecificCFSourceSetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceChangeaspecificCFSourceSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
