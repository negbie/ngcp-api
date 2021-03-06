// Code generated by go-swagger; DO NOT EDIT.

package vouchers

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

	"github.com/negbie/ngcp-api/models"
)

// NewCreateanewVoucherParams creates a new CreateanewVoucherParams object
// with the default values initialized.
func NewCreateanewVoucherParams() *CreateanewVoucherParams {
	var ()
	return &CreateanewVoucherParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateanewVoucherParamsWithTimeout creates a new CreateanewVoucherParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateanewVoucherParamsWithTimeout(timeout time.Duration) *CreateanewVoucherParams {
	var ()
	return &CreateanewVoucherParams{

		timeout: timeout,
	}
}

// NewCreateanewVoucherParamsWithContext creates a new CreateanewVoucherParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateanewVoucherParamsWithContext(ctx context.Context) *CreateanewVoucherParams {
	var ()
	return &CreateanewVoucherParams{

		Context: ctx,
	}
}

// NewCreateanewVoucherParamsWithHTTPClient creates a new CreateanewVoucherParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateanewVoucherParamsWithHTTPClient(client *http.Client) *CreateanewVoucherParams {
	var ()
	return &CreateanewVoucherParams{
		HTTPClient: client,
	}
}

/*CreateanewVoucherParams contains all the parameters to send to the API endpoint
for the createanew voucher operation typically these are written to a http.Request
*/
type CreateanewVoucherParams struct {

	/*Body*/
	Body *models.CreateanewVoucherRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the createanew voucher params
func (o *CreateanewVoucherParams) WithTimeout(timeout time.Duration) *CreateanewVoucherParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the createanew voucher params
func (o *CreateanewVoucherParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the createanew voucher params
func (o *CreateanewVoucherParams) WithContext(ctx context.Context) *CreateanewVoucherParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the createanew voucher params
func (o *CreateanewVoucherParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the createanew voucher params
func (o *CreateanewVoucherParams) WithHTTPClient(client *http.Client) *CreateanewVoucherParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the createanew voucher params
func (o *CreateanewVoucherParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the createanew voucher params
func (o *CreateanewVoucherParams) WithBody(body *models.CreateanewVoucherRequest) *CreateanewVoucherParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the createanew voucher params
func (o *CreateanewVoucherParams) SetBody(body *models.CreateanewVoucherRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the createanew voucher params
func (o *CreateanewVoucherParams) WithContentType(contentType string) *CreateanewVoucherParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the createanew voucher params
func (o *CreateanewVoucherParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateanewVoucherParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
