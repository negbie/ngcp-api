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

// NewReplaceChangeaspecificVoucherParams creates a new ReplaceChangeaspecificVoucherParams object
// with the default values initialized.
func NewReplaceChangeaspecificVoucherParams() *ReplaceChangeaspecificVoucherParams {
	var ()
	return &ReplaceChangeaspecificVoucherParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceChangeaspecificVoucherParamsWithTimeout creates a new ReplaceChangeaspecificVoucherParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceChangeaspecificVoucherParamsWithTimeout(timeout time.Duration) *ReplaceChangeaspecificVoucherParams {
	var ()
	return &ReplaceChangeaspecificVoucherParams{

		timeout: timeout,
	}
}

// NewReplaceChangeaspecificVoucherParamsWithContext creates a new ReplaceChangeaspecificVoucherParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceChangeaspecificVoucherParamsWithContext(ctx context.Context) *ReplaceChangeaspecificVoucherParams {
	var ()
	return &ReplaceChangeaspecificVoucherParams{

		Context: ctx,
	}
}

// NewReplaceChangeaspecificVoucherParamsWithHTTPClient creates a new ReplaceChangeaspecificVoucherParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceChangeaspecificVoucherParamsWithHTTPClient(client *http.Client) *ReplaceChangeaspecificVoucherParams {
	var ()
	return &ReplaceChangeaspecificVoucherParams{
		HTTPClient: client,
	}
}

/*ReplaceChangeaspecificVoucherParams contains all the parameters to send to the API endpoint
for the replace changeaspecific voucher operation typically these are written to a http.Request
*/
type ReplaceChangeaspecificVoucherParams struct {

	/*Body*/
	Body *Replace1changeaspecificVoucherRequest
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

// WithTimeout adds the timeout to the replace changeaspecific voucher params
func (o *ReplaceChangeaspecificVoucherParams) WithTimeout(timeout time.Duration) *ReplaceChangeaspecificVoucherParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace changeaspecific voucher params
func (o *ReplaceChangeaspecificVoucherParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace changeaspecific voucher params
func (o *ReplaceChangeaspecificVoucherParams) WithContext(ctx context.Context) *ReplaceChangeaspecificVoucherParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace changeaspecific voucher params
func (o *ReplaceChangeaspecificVoucherParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace changeaspecific voucher params
func (o *ReplaceChangeaspecificVoucherParams) WithHTTPClient(client *http.Client) *ReplaceChangeaspecificVoucherParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace changeaspecific voucher params
func (o *ReplaceChangeaspecificVoucherParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace changeaspecific voucher params
func (o *ReplaceChangeaspecificVoucherParams) WithBody(body *Replace1changeaspecificVoucherRequest) *ReplaceChangeaspecificVoucherParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace changeaspecific voucher params
func (o *ReplaceChangeaspecificVoucherParams) SetBody(body *Replace1changeaspecificVoucherRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the replace changeaspecific voucher params
func (o *ReplaceChangeaspecificVoucherParams) WithContentType(contentType string) *ReplaceChangeaspecificVoucherParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the replace changeaspecific voucher params
func (o *ReplaceChangeaspecificVoucherParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the replace changeaspecific voucher params
func (o *ReplaceChangeaspecificVoucherParams) WithID(id string) *ReplaceChangeaspecificVoucherParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replace changeaspecific voucher params
func (o *ReplaceChangeaspecificVoucherParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceChangeaspecificVoucherParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
