// Code generated by go-swagger; DO NOT EDIT.

package ncoslevels

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

// NewCreateanewNcosLevelParams creates a new CreateanewNcosLevelParams object
// with the default values initialized.
func NewCreateanewNcosLevelParams() *CreateanewNcosLevelParams {
	var ()
	return &CreateanewNcosLevelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateanewNcosLevelParamsWithTimeout creates a new CreateanewNcosLevelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateanewNcosLevelParamsWithTimeout(timeout time.Duration) *CreateanewNcosLevelParams {
	var ()
	return &CreateanewNcosLevelParams{

		timeout: timeout,
	}
}

// NewCreateanewNcosLevelParamsWithContext creates a new CreateanewNcosLevelParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateanewNcosLevelParamsWithContext(ctx context.Context) *CreateanewNcosLevelParams {
	var ()
	return &CreateanewNcosLevelParams{

		Context: ctx,
	}
}

// NewCreateanewNcosLevelParamsWithHTTPClient creates a new CreateanewNcosLevelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateanewNcosLevelParamsWithHTTPClient(client *http.Client) *CreateanewNcosLevelParams {
	var ()
	return &CreateanewNcosLevelParams{
		HTTPClient: client,
	}
}

/*CreateanewNcosLevelParams contains all the parameters to send to the API endpoint
for the createanew ncos level operation typically these are written to a http.Request
*/
type CreateanewNcosLevelParams struct {

	/*Body*/
	Body *models.CreateanewNcosLevelRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the createanew ncos level params
func (o *CreateanewNcosLevelParams) WithTimeout(timeout time.Duration) *CreateanewNcosLevelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the createanew ncos level params
func (o *CreateanewNcosLevelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the createanew ncos level params
func (o *CreateanewNcosLevelParams) WithContext(ctx context.Context) *CreateanewNcosLevelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the createanew ncos level params
func (o *CreateanewNcosLevelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the createanew ncos level params
func (o *CreateanewNcosLevelParams) WithHTTPClient(client *http.Client) *CreateanewNcosLevelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the createanew ncos level params
func (o *CreateanewNcosLevelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the createanew ncos level params
func (o *CreateanewNcosLevelParams) WithBody(body *models.CreateanewNcosLevelRequest) *CreateanewNcosLevelParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the createanew ncos level params
func (o *CreateanewNcosLevelParams) SetBody(body *models.CreateanewNcosLevelRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the createanew ncos level params
func (o *CreateanewNcosLevelParams) WithContentType(contentType string) *CreateanewNcosLevelParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the createanew ncos level params
func (o *CreateanewNcosLevelParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateanewNcosLevelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
