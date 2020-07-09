// Code generated by go-swagger; DO NOT EDIT.

package profilepackages

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

// NewCreateanewProfilePackageParams creates a new CreateanewProfilePackageParams object
// with the default values initialized.
func NewCreateanewProfilePackageParams() *CreateanewProfilePackageParams {
	var ()
	return &CreateanewProfilePackageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateanewProfilePackageParamsWithTimeout creates a new CreateanewProfilePackageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateanewProfilePackageParamsWithTimeout(timeout time.Duration) *CreateanewProfilePackageParams {
	var ()
	return &CreateanewProfilePackageParams{

		timeout: timeout,
	}
}

// NewCreateanewProfilePackageParamsWithContext creates a new CreateanewProfilePackageParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateanewProfilePackageParamsWithContext(ctx context.Context) *CreateanewProfilePackageParams {
	var ()
	return &CreateanewProfilePackageParams{

		Context: ctx,
	}
}

// NewCreateanewProfilePackageParamsWithHTTPClient creates a new CreateanewProfilePackageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateanewProfilePackageParamsWithHTTPClient(client *http.Client) *CreateanewProfilePackageParams {
	var ()
	return &CreateanewProfilePackageParams{
		HTTPClient: client,
	}
}

/*CreateanewProfilePackageParams contains all the parameters to send to the API endpoint
for the createanew profile package operation typically these are written to a http.Request
*/
type CreateanewProfilePackageParams struct {

	/*Body*/
	Body *models.CreateanewProfilePackageRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the createanew profile package params
func (o *CreateanewProfilePackageParams) WithTimeout(timeout time.Duration) *CreateanewProfilePackageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the createanew profile package params
func (o *CreateanewProfilePackageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the createanew profile package params
func (o *CreateanewProfilePackageParams) WithContext(ctx context.Context) *CreateanewProfilePackageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the createanew profile package params
func (o *CreateanewProfilePackageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the createanew profile package params
func (o *CreateanewProfilePackageParams) WithHTTPClient(client *http.Client) *CreateanewProfilePackageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the createanew profile package params
func (o *CreateanewProfilePackageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the createanew profile package params
func (o *CreateanewProfilePackageParams) WithBody(body *models.CreateanewProfilePackageRequest) *CreateanewProfilePackageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the createanew profile package params
func (o *CreateanewProfilePackageParams) SetBody(body *models.CreateanewProfilePackageRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the createanew profile package params
func (o *CreateanewProfilePackageParams) WithContentType(contentType string) *CreateanewProfilePackageParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the createanew profile package params
func (o *CreateanewProfilePackageParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateanewProfilePackageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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