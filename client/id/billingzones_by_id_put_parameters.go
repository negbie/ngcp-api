// Code generated by go-swagger; DO NOT EDIT.

package id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewBillingzonesByIDPutParams creates a new BillingzonesByIDPutParams object
// with the default values initialized.
func NewBillingzonesByIDPutParams() *BillingzonesByIDPutParams {
	var ()
	return &BillingzonesByIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBillingzonesByIDPutParamsWithTimeout creates a new BillingzonesByIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBillingzonesByIDPutParamsWithTimeout(timeout time.Duration) *BillingzonesByIDPutParams {
	var ()
	return &BillingzonesByIDPutParams{

		timeout: timeout,
	}
}

// NewBillingzonesByIDPutParamsWithContext creates a new BillingzonesByIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewBillingzonesByIDPutParamsWithContext(ctx context.Context) *BillingzonesByIDPutParams {
	var ()
	return &BillingzonesByIDPutParams{

		Context: ctx,
	}
}

// NewBillingzonesByIDPutParamsWithHTTPClient creates a new BillingzonesByIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBillingzonesByIDPutParamsWithHTTPClient(client *http.Client) *BillingzonesByIDPutParams {
	var ()
	return &BillingzonesByIDPutParams{
		HTTPClient: client,
	}
}

/*BillingzonesByIDPutParams contains all the parameters to send to the API endpoint
for the billingzones by Id put operation typically these are written to a http.Request
*/
type BillingzonesByIDPutParams struct {

	/*Body*/
	Body *Replace1changeaspecificBillingZoneRequest
	/*ContentType*/
	ContentType string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the billingzones by Id put params
func (o *BillingzonesByIDPutParams) WithTimeout(timeout time.Duration) *BillingzonesByIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the billingzones by Id put params
func (o *BillingzonesByIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the billingzones by Id put params
func (o *BillingzonesByIDPutParams) WithContext(ctx context.Context) *BillingzonesByIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the billingzones by Id put params
func (o *BillingzonesByIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the billingzones by Id put params
func (o *BillingzonesByIDPutParams) WithHTTPClient(client *http.Client) *BillingzonesByIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the billingzones by Id put params
func (o *BillingzonesByIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the billingzones by Id put params
func (o *BillingzonesByIDPutParams) WithBody(body *Replace1changeaspecificBillingZoneRequest) *BillingzonesByIDPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the billingzones by Id put params
func (o *BillingzonesByIDPutParams) SetBody(body *Replace1changeaspecificBillingZoneRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the billingzones by Id put params
func (o *BillingzonesByIDPutParams) WithContentType(contentType string) *BillingzonesByIDPutParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the billingzones by Id put params
func (o *BillingzonesByIDPutParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the billingzones by Id put params
func (o *BillingzonesByIDPutParams) WithID(id string) *BillingzonesByIDPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the billingzones by Id put params
func (o *BillingzonesByIDPutParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *BillingzonesByIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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