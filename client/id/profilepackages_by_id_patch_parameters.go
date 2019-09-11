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

// NewProfilepackagesByIDPatchParams creates a new ProfilepackagesByIDPatchParams object
// with the default values initialized.
func NewProfilepackagesByIDPatchParams() *ProfilepackagesByIDPatchParams {
	var ()
	return &ProfilepackagesByIDPatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProfilepackagesByIDPatchParamsWithTimeout creates a new ProfilepackagesByIDPatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProfilepackagesByIDPatchParamsWithTimeout(timeout time.Duration) *ProfilepackagesByIDPatchParams {
	var ()
	return &ProfilepackagesByIDPatchParams{

		timeout: timeout,
	}
}

// NewProfilepackagesByIDPatchParamsWithContext creates a new ProfilepackagesByIDPatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewProfilepackagesByIDPatchParamsWithContext(ctx context.Context) *ProfilepackagesByIDPatchParams {
	var ()
	return &ProfilepackagesByIDPatchParams{

		Context: ctx,
	}
}

// NewProfilepackagesByIDPatchParamsWithHTTPClient creates a new ProfilepackagesByIDPatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProfilepackagesByIDPatchParamsWithHTTPClient(client *http.Client) *ProfilepackagesByIDPatchParams {
	var ()
	return &ProfilepackagesByIDPatchParams{
		HTTPClient: client,
	}
}

/*ProfilepackagesByIDPatchParams contains all the parameters to send to the API endpoint
for the profilepackages by Id patch operation typically these are written to a http.Request
*/
type ProfilepackagesByIDPatchParams struct {

	/*Body*/
	Body string
	/*ContentType*/
	ContentType string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the profilepackages by Id patch params
func (o *ProfilepackagesByIDPatchParams) WithTimeout(timeout time.Duration) *ProfilepackagesByIDPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the profilepackages by Id patch params
func (o *ProfilepackagesByIDPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the profilepackages by Id patch params
func (o *ProfilepackagesByIDPatchParams) WithContext(ctx context.Context) *ProfilepackagesByIDPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the profilepackages by Id patch params
func (o *ProfilepackagesByIDPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the profilepackages by Id patch params
func (o *ProfilepackagesByIDPatchParams) WithHTTPClient(client *http.Client) *ProfilepackagesByIDPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the profilepackages by Id patch params
func (o *ProfilepackagesByIDPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the profilepackages by Id patch params
func (o *ProfilepackagesByIDPatchParams) WithBody(body string) *ProfilepackagesByIDPatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the profilepackages by Id patch params
func (o *ProfilepackagesByIDPatchParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the profilepackages by Id patch params
func (o *ProfilepackagesByIDPatchParams) WithContentType(contentType string) *ProfilepackagesByIDPatchParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the profilepackages by Id patch params
func (o *ProfilepackagesByIDPatchParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the profilepackages by Id patch params
func (o *ProfilepackagesByIDPatchParams) WithID(id string) *ProfilepackagesByIDPatchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the profilepackages by Id patch params
func (o *ProfilepackagesByIDPatchParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ProfilepackagesByIDPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}