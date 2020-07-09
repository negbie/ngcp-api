// Code generated by go-swagger; DO NOT EDIT.

package billingprofiles

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

// NewCreateanewBillingProfileParams creates a new CreateanewBillingProfileParams object
// with the default values initialized.
func NewCreateanewBillingProfileParams() *CreateanewBillingProfileParams {
	var ()
	return &CreateanewBillingProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateanewBillingProfileParamsWithTimeout creates a new CreateanewBillingProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateanewBillingProfileParamsWithTimeout(timeout time.Duration) *CreateanewBillingProfileParams {
	var ()
	return &CreateanewBillingProfileParams{

		timeout: timeout,
	}
}

// NewCreateanewBillingProfileParamsWithContext creates a new CreateanewBillingProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateanewBillingProfileParamsWithContext(ctx context.Context) *CreateanewBillingProfileParams {
	var ()
	return &CreateanewBillingProfileParams{

		Context: ctx,
	}
}

// NewCreateanewBillingProfileParamsWithHTTPClient creates a new CreateanewBillingProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateanewBillingProfileParamsWithHTTPClient(client *http.Client) *CreateanewBillingProfileParams {
	var ()
	return &CreateanewBillingProfileParams{
		HTTPClient: client,
	}
}

/*CreateanewBillingProfileParams contains all the parameters to send to the API endpoint
for the createanew billing profile operation typically these are written to a http.Request
*/
type CreateanewBillingProfileParams struct {

	/*Body*/
	Body *models.CreateanewBillingProfileRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the createanew billing profile params
func (o *CreateanewBillingProfileParams) WithTimeout(timeout time.Duration) *CreateanewBillingProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the createanew billing profile params
func (o *CreateanewBillingProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the createanew billing profile params
func (o *CreateanewBillingProfileParams) WithContext(ctx context.Context) *CreateanewBillingProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the createanew billing profile params
func (o *CreateanewBillingProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the createanew billing profile params
func (o *CreateanewBillingProfileParams) WithHTTPClient(client *http.Client) *CreateanewBillingProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the createanew billing profile params
func (o *CreateanewBillingProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the createanew billing profile params
func (o *CreateanewBillingProfileParams) WithBody(body *models.CreateanewBillingProfileRequest) *CreateanewBillingProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the createanew billing profile params
func (o *CreateanewBillingProfileParams) SetBody(body *models.CreateanewBillingProfileRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the createanew billing profile params
func (o *CreateanewBillingProfileParams) WithContentType(contentType string) *CreateanewBillingProfileParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the createanew billing profile params
func (o *CreateanewBillingProfileParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateanewBillingProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
