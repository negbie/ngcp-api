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

	"github.com/negbie/ngcp-api/models"
)

// NewCreateanewCallControlParams creates a new CreateanewCallControlParams object
// with the default values initialized.
func NewCreateanewCallControlParams() *CreateanewCallControlParams {
	var ()
	return &CreateanewCallControlParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateanewCallControlParamsWithTimeout creates a new CreateanewCallControlParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateanewCallControlParamsWithTimeout(timeout time.Duration) *CreateanewCallControlParams {
	var ()
	return &CreateanewCallControlParams{

		timeout: timeout,
	}
}

// NewCreateanewCallControlParamsWithContext creates a new CreateanewCallControlParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateanewCallControlParamsWithContext(ctx context.Context) *CreateanewCallControlParams {
	var ()
	return &CreateanewCallControlParams{

		Context: ctx,
	}
}

// NewCreateanewCallControlParamsWithHTTPClient creates a new CreateanewCallControlParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateanewCallControlParamsWithHTTPClient(client *http.Client) *CreateanewCallControlParams {
	var ()
	return &CreateanewCallControlParams{
		HTTPClient: client,
	}
}

/*CreateanewCallControlParams contains all the parameters to send to the API endpoint
for the createanew call control operation typically these are written to a http.Request
*/
type CreateanewCallControlParams struct {

	/*Body*/
	Body *models.CreateanewCallControlRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the createanew call control params
func (o *CreateanewCallControlParams) WithTimeout(timeout time.Duration) *CreateanewCallControlParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the createanew call control params
func (o *CreateanewCallControlParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the createanew call control params
func (o *CreateanewCallControlParams) WithContext(ctx context.Context) *CreateanewCallControlParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the createanew call control params
func (o *CreateanewCallControlParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the createanew call control params
func (o *CreateanewCallControlParams) WithHTTPClient(client *http.Client) *CreateanewCallControlParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the createanew call control params
func (o *CreateanewCallControlParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the createanew call control params
func (o *CreateanewCallControlParams) WithBody(body *models.CreateanewCallControlRequest) *CreateanewCallControlParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the createanew call control params
func (o *CreateanewCallControlParams) SetBody(body *models.CreateanewCallControlRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the createanew call control params
func (o *CreateanewCallControlParams) WithContentType(contentType string) *CreateanewCallControlParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the createanew call control params
func (o *CreateanewCallControlParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateanewCallControlParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
