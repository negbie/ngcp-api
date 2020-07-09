// Code generated by go-swagger; DO NOT EDIT.

package emailtemplates

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

// NewCreateanewEmailTemplateParams creates a new CreateanewEmailTemplateParams object
// with the default values initialized.
func NewCreateanewEmailTemplateParams() *CreateanewEmailTemplateParams {
	var ()
	return &CreateanewEmailTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateanewEmailTemplateParamsWithTimeout creates a new CreateanewEmailTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateanewEmailTemplateParamsWithTimeout(timeout time.Duration) *CreateanewEmailTemplateParams {
	var ()
	return &CreateanewEmailTemplateParams{

		timeout: timeout,
	}
}

// NewCreateanewEmailTemplateParamsWithContext creates a new CreateanewEmailTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateanewEmailTemplateParamsWithContext(ctx context.Context) *CreateanewEmailTemplateParams {
	var ()
	return &CreateanewEmailTemplateParams{

		Context: ctx,
	}
}

// NewCreateanewEmailTemplateParamsWithHTTPClient creates a new CreateanewEmailTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateanewEmailTemplateParamsWithHTTPClient(client *http.Client) *CreateanewEmailTemplateParams {
	var ()
	return &CreateanewEmailTemplateParams{
		HTTPClient: client,
	}
}

/*CreateanewEmailTemplateParams contains all the parameters to send to the API endpoint
for the createanew email template operation typically these are written to a http.Request
*/
type CreateanewEmailTemplateParams struct {

	/*Body*/
	Body *models.CreateanewEmailTemplateRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the createanew email template params
func (o *CreateanewEmailTemplateParams) WithTimeout(timeout time.Duration) *CreateanewEmailTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the createanew email template params
func (o *CreateanewEmailTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the createanew email template params
func (o *CreateanewEmailTemplateParams) WithContext(ctx context.Context) *CreateanewEmailTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the createanew email template params
func (o *CreateanewEmailTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the createanew email template params
func (o *CreateanewEmailTemplateParams) WithHTTPClient(client *http.Client) *CreateanewEmailTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the createanew email template params
func (o *CreateanewEmailTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the createanew email template params
func (o *CreateanewEmailTemplateParams) WithBody(body *models.CreateanewEmailTemplateRequest) *CreateanewEmailTemplateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the createanew email template params
func (o *CreateanewEmailTemplateParams) SetBody(body *models.CreateanewEmailTemplateRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the createanew email template params
func (o *CreateanewEmailTemplateParams) WithContentType(contentType string) *CreateanewEmailTemplateParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the createanew email template params
func (o *CreateanewEmailTemplateParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateanewEmailTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
