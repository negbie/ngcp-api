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

// NewEmergencymappingcontainersByIDPatchParams creates a new EmergencymappingcontainersByIDPatchParams object
// with the default values initialized.
func NewEmergencymappingcontainersByIDPatchParams() *EmergencymappingcontainersByIDPatchParams {
	var ()
	return &EmergencymappingcontainersByIDPatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEmergencymappingcontainersByIDPatchParamsWithTimeout creates a new EmergencymappingcontainersByIDPatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmergencymappingcontainersByIDPatchParamsWithTimeout(timeout time.Duration) *EmergencymappingcontainersByIDPatchParams {
	var ()
	return &EmergencymappingcontainersByIDPatchParams{

		timeout: timeout,
	}
}

// NewEmergencymappingcontainersByIDPatchParamsWithContext creates a new EmergencymappingcontainersByIDPatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmergencymappingcontainersByIDPatchParamsWithContext(ctx context.Context) *EmergencymappingcontainersByIDPatchParams {
	var ()
	return &EmergencymappingcontainersByIDPatchParams{

		Context: ctx,
	}
}

// NewEmergencymappingcontainersByIDPatchParamsWithHTTPClient creates a new EmergencymappingcontainersByIDPatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmergencymappingcontainersByIDPatchParamsWithHTTPClient(client *http.Client) *EmergencymappingcontainersByIDPatchParams {
	var ()
	return &EmergencymappingcontainersByIDPatchParams{
		HTTPClient: client,
	}
}

/*EmergencymappingcontainersByIDPatchParams contains all the parameters to send to the API endpoint
for the emergencymappingcontainers by Id patch operation typically these are written to a http.Request
*/
type EmergencymappingcontainersByIDPatchParams struct {

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

// WithTimeout adds the timeout to the emergencymappingcontainers by Id patch params
func (o *EmergencymappingcontainersByIDPatchParams) WithTimeout(timeout time.Duration) *EmergencymappingcontainersByIDPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the emergencymappingcontainers by Id patch params
func (o *EmergencymappingcontainersByIDPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the emergencymappingcontainers by Id patch params
func (o *EmergencymappingcontainersByIDPatchParams) WithContext(ctx context.Context) *EmergencymappingcontainersByIDPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the emergencymappingcontainers by Id patch params
func (o *EmergencymappingcontainersByIDPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the emergencymappingcontainers by Id patch params
func (o *EmergencymappingcontainersByIDPatchParams) WithHTTPClient(client *http.Client) *EmergencymappingcontainersByIDPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the emergencymappingcontainers by Id patch params
func (o *EmergencymappingcontainersByIDPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the emergencymappingcontainers by Id patch params
func (o *EmergencymappingcontainersByIDPatchParams) WithBody(body string) *EmergencymappingcontainersByIDPatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the emergencymappingcontainers by Id patch params
func (o *EmergencymappingcontainersByIDPatchParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the emergencymappingcontainers by Id patch params
func (o *EmergencymappingcontainersByIDPatchParams) WithContentType(contentType string) *EmergencymappingcontainersByIDPatchParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the emergencymappingcontainers by Id patch params
func (o *EmergencymappingcontainersByIDPatchParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the emergencymappingcontainers by Id patch params
func (o *EmergencymappingcontainersByIDPatchParams) WithID(id string) *EmergencymappingcontainersByIDPatchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the emergencymappingcontainers by Id patch params
func (o *EmergencymappingcontainersByIDPatchParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EmergencymappingcontainersByIDPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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