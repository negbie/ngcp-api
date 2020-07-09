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

// NewReplaceChangeaspecificNumberParams creates a new ReplaceChangeaspecificNumberParams object
// with the default values initialized.
func NewReplaceChangeaspecificNumberParams() *ReplaceChangeaspecificNumberParams {
	var ()
	return &ReplaceChangeaspecificNumberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceChangeaspecificNumberParamsWithTimeout creates a new ReplaceChangeaspecificNumberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceChangeaspecificNumberParamsWithTimeout(timeout time.Duration) *ReplaceChangeaspecificNumberParams {
	var ()
	return &ReplaceChangeaspecificNumberParams{

		timeout: timeout,
	}
}

// NewReplaceChangeaspecificNumberParamsWithContext creates a new ReplaceChangeaspecificNumberParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceChangeaspecificNumberParamsWithContext(ctx context.Context) *ReplaceChangeaspecificNumberParams {
	var ()
	return &ReplaceChangeaspecificNumberParams{

		Context: ctx,
	}
}

// NewReplaceChangeaspecificNumberParamsWithHTTPClient creates a new ReplaceChangeaspecificNumberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceChangeaspecificNumberParamsWithHTTPClient(client *http.Client) *ReplaceChangeaspecificNumberParams {
	var ()
	return &ReplaceChangeaspecificNumberParams{
		HTTPClient: client,
	}
}

/*ReplaceChangeaspecificNumberParams contains all the parameters to send to the API endpoint
for the replace changeaspecific number operation typically these are written to a http.Request
*/
type ReplaceChangeaspecificNumberParams struct {

	/*Body*/
	Body *Replace1changeaspecificNumberRequest
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

// WithTimeout adds the timeout to the replace changeaspecific number params
func (o *ReplaceChangeaspecificNumberParams) WithTimeout(timeout time.Duration) *ReplaceChangeaspecificNumberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace changeaspecific number params
func (o *ReplaceChangeaspecificNumberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace changeaspecific number params
func (o *ReplaceChangeaspecificNumberParams) WithContext(ctx context.Context) *ReplaceChangeaspecificNumberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace changeaspecific number params
func (o *ReplaceChangeaspecificNumberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace changeaspecific number params
func (o *ReplaceChangeaspecificNumberParams) WithHTTPClient(client *http.Client) *ReplaceChangeaspecificNumberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace changeaspecific number params
func (o *ReplaceChangeaspecificNumberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace changeaspecific number params
func (o *ReplaceChangeaspecificNumberParams) WithBody(body *Replace1changeaspecificNumberRequest) *ReplaceChangeaspecificNumberParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace changeaspecific number params
func (o *ReplaceChangeaspecificNumberParams) SetBody(body *Replace1changeaspecificNumberRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the replace changeaspecific number params
func (o *ReplaceChangeaspecificNumberParams) WithContentType(contentType string) *ReplaceChangeaspecificNumberParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the replace changeaspecific number params
func (o *ReplaceChangeaspecificNumberParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the replace changeaspecific number params
func (o *ReplaceChangeaspecificNumberParams) WithID(id string) *ReplaceChangeaspecificNumberParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replace changeaspecific number params
func (o *ReplaceChangeaspecificNumberParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceChangeaspecificNumberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
