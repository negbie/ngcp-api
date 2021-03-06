// Code generated by go-swagger; DO NOT EDIT.

package upnrewritesets

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

// NewCreateanewUpnRewriteSetParams creates a new CreateanewUpnRewriteSetParams object
// with the default values initialized.
func NewCreateanewUpnRewriteSetParams() *CreateanewUpnRewriteSetParams {
	var ()
	return &CreateanewUpnRewriteSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateanewUpnRewriteSetParamsWithTimeout creates a new CreateanewUpnRewriteSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateanewUpnRewriteSetParamsWithTimeout(timeout time.Duration) *CreateanewUpnRewriteSetParams {
	var ()
	return &CreateanewUpnRewriteSetParams{

		timeout: timeout,
	}
}

// NewCreateanewUpnRewriteSetParamsWithContext creates a new CreateanewUpnRewriteSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateanewUpnRewriteSetParamsWithContext(ctx context.Context) *CreateanewUpnRewriteSetParams {
	var ()
	return &CreateanewUpnRewriteSetParams{

		Context: ctx,
	}
}

// NewCreateanewUpnRewriteSetParamsWithHTTPClient creates a new CreateanewUpnRewriteSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateanewUpnRewriteSetParamsWithHTTPClient(client *http.Client) *CreateanewUpnRewriteSetParams {
	var ()
	return &CreateanewUpnRewriteSetParams{
		HTTPClient: client,
	}
}

/*CreateanewUpnRewriteSetParams contains all the parameters to send to the API endpoint
for the createanew upn rewrite set operation typically these are written to a http.Request
*/
type CreateanewUpnRewriteSetParams struct {

	/*Body*/
	Body string
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the createanew upn rewrite set params
func (o *CreateanewUpnRewriteSetParams) WithTimeout(timeout time.Duration) *CreateanewUpnRewriteSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the createanew upn rewrite set params
func (o *CreateanewUpnRewriteSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the createanew upn rewrite set params
func (o *CreateanewUpnRewriteSetParams) WithContext(ctx context.Context) *CreateanewUpnRewriteSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the createanew upn rewrite set params
func (o *CreateanewUpnRewriteSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the createanew upn rewrite set params
func (o *CreateanewUpnRewriteSetParams) WithHTTPClient(client *http.Client) *CreateanewUpnRewriteSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the createanew upn rewrite set params
func (o *CreateanewUpnRewriteSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the createanew upn rewrite set params
func (o *CreateanewUpnRewriteSetParams) WithBody(body string) *CreateanewUpnRewriteSetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the createanew upn rewrite set params
func (o *CreateanewUpnRewriteSetParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the createanew upn rewrite set params
func (o *CreateanewUpnRewriteSetParams) WithContentType(contentType string) *CreateanewUpnRewriteSetParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the createanew upn rewrite set params
func (o *CreateanewUpnRewriteSetParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateanewUpnRewriteSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
