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

// NewChangeaspecificBillingFeeParams creates a new ChangeaspecificBillingFeeParams object
// with the default values initialized.
func NewChangeaspecificBillingFeeParams() *ChangeaspecificBillingFeeParams {
	var ()
	return &ChangeaspecificBillingFeeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeaspecificBillingFeeParamsWithTimeout creates a new ChangeaspecificBillingFeeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeaspecificBillingFeeParamsWithTimeout(timeout time.Duration) *ChangeaspecificBillingFeeParams {
	var ()
	return &ChangeaspecificBillingFeeParams{

		timeout: timeout,
	}
}

// NewChangeaspecificBillingFeeParamsWithContext creates a new ChangeaspecificBillingFeeParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeaspecificBillingFeeParamsWithContext(ctx context.Context) *ChangeaspecificBillingFeeParams {
	var ()
	return &ChangeaspecificBillingFeeParams{

		Context: ctx,
	}
}

// NewChangeaspecificBillingFeeParamsWithHTTPClient creates a new ChangeaspecificBillingFeeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeaspecificBillingFeeParamsWithHTTPClient(client *http.Client) *ChangeaspecificBillingFeeParams {
	var ()
	return &ChangeaspecificBillingFeeParams{
		HTTPClient: client,
	}
}

/*ChangeaspecificBillingFeeParams contains all the parameters to send to the API endpoint
for the changeaspecific billing fee operation typically these are written to a http.Request
*/
type ChangeaspecificBillingFeeParams struct {

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

// WithTimeout adds the timeout to the changeaspecific billing fee params
func (o *ChangeaspecificBillingFeeParams) WithTimeout(timeout time.Duration) *ChangeaspecificBillingFeeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the changeaspecific billing fee params
func (o *ChangeaspecificBillingFeeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the changeaspecific billing fee params
func (o *ChangeaspecificBillingFeeParams) WithContext(ctx context.Context) *ChangeaspecificBillingFeeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the changeaspecific billing fee params
func (o *ChangeaspecificBillingFeeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the changeaspecific billing fee params
func (o *ChangeaspecificBillingFeeParams) WithHTTPClient(client *http.Client) *ChangeaspecificBillingFeeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the changeaspecific billing fee params
func (o *ChangeaspecificBillingFeeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the changeaspecific billing fee params
func (o *ChangeaspecificBillingFeeParams) WithBody(body string) *ChangeaspecificBillingFeeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the changeaspecific billing fee params
func (o *ChangeaspecificBillingFeeParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the changeaspecific billing fee params
func (o *ChangeaspecificBillingFeeParams) WithContentType(contentType string) *ChangeaspecificBillingFeeParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the changeaspecific billing fee params
func (o *ChangeaspecificBillingFeeParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the changeaspecific billing fee params
func (o *ChangeaspecificBillingFeeParams) WithID(id string) *ChangeaspecificBillingFeeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the changeaspecific billing fee params
func (o *ChangeaspecificBillingFeeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeaspecificBillingFeeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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