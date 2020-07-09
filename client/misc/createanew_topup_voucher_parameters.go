// Code generated by go-swagger; DO NOT EDIT.

package misc

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

	"github.com/negbie/ngcp-api1/models"
)

// NewCreateanewTopupVoucherParams creates a new CreateanewTopupVoucherParams object
// with the default values initialized.
func NewCreateanewTopupVoucherParams() *CreateanewTopupVoucherParams {
	var ()
	return &CreateanewTopupVoucherParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateanewTopupVoucherParamsWithTimeout creates a new CreateanewTopupVoucherParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateanewTopupVoucherParamsWithTimeout(timeout time.Duration) *CreateanewTopupVoucherParams {
	var ()
	return &CreateanewTopupVoucherParams{

		timeout: timeout,
	}
}

// NewCreateanewTopupVoucherParamsWithContext creates a new CreateanewTopupVoucherParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateanewTopupVoucherParamsWithContext(ctx context.Context) *CreateanewTopupVoucherParams {
	var ()
	return &CreateanewTopupVoucherParams{

		Context: ctx,
	}
}

// NewCreateanewTopupVoucherParamsWithHTTPClient creates a new CreateanewTopupVoucherParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateanewTopupVoucherParamsWithHTTPClient(client *http.Client) *CreateanewTopupVoucherParams {
	var ()
	return &CreateanewTopupVoucherParams{
		HTTPClient: client,
	}
}

/*CreateanewTopupVoucherParams contains all the parameters to send to the API endpoint
for the createanew topup voucher operation typically these are written to a http.Request
*/
type CreateanewTopupVoucherParams struct {

	/*Body*/
	Body *models.CreateanewTopupVoucherRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the createanew topup voucher params
func (o *CreateanewTopupVoucherParams) WithTimeout(timeout time.Duration) *CreateanewTopupVoucherParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the createanew topup voucher params
func (o *CreateanewTopupVoucherParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the createanew topup voucher params
func (o *CreateanewTopupVoucherParams) WithContext(ctx context.Context) *CreateanewTopupVoucherParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the createanew topup voucher params
func (o *CreateanewTopupVoucherParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the createanew topup voucher params
func (o *CreateanewTopupVoucherParams) WithHTTPClient(client *http.Client) *CreateanewTopupVoucherParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the createanew topup voucher params
func (o *CreateanewTopupVoucherParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the createanew topup voucher params
func (o *CreateanewTopupVoucherParams) WithBody(body *models.CreateanewTopupVoucherRequest) *CreateanewTopupVoucherParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the createanew topup voucher params
func (o *CreateanewTopupVoucherParams) SetBody(body *models.CreateanewTopupVoucherRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the createanew topup voucher params
func (o *CreateanewTopupVoucherParams) WithContentType(contentType string) *CreateanewTopupVoucherParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the createanew topup voucher params
func (o *CreateanewTopupVoucherParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateanewTopupVoucherParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
