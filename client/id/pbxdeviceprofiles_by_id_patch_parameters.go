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

// NewPbxdeviceprofilesByIDPatchParams creates a new PbxdeviceprofilesByIDPatchParams object
// with the default values initialized.
func NewPbxdeviceprofilesByIDPatchParams() *PbxdeviceprofilesByIDPatchParams {
	var ()
	return &PbxdeviceprofilesByIDPatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPbxdeviceprofilesByIDPatchParamsWithTimeout creates a new PbxdeviceprofilesByIDPatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPbxdeviceprofilesByIDPatchParamsWithTimeout(timeout time.Duration) *PbxdeviceprofilesByIDPatchParams {
	var ()
	return &PbxdeviceprofilesByIDPatchParams{

		timeout: timeout,
	}
}

// NewPbxdeviceprofilesByIDPatchParamsWithContext creates a new PbxdeviceprofilesByIDPatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPbxdeviceprofilesByIDPatchParamsWithContext(ctx context.Context) *PbxdeviceprofilesByIDPatchParams {
	var ()
	return &PbxdeviceprofilesByIDPatchParams{

		Context: ctx,
	}
}

// NewPbxdeviceprofilesByIDPatchParamsWithHTTPClient creates a new PbxdeviceprofilesByIDPatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPbxdeviceprofilesByIDPatchParamsWithHTTPClient(client *http.Client) *PbxdeviceprofilesByIDPatchParams {
	var ()
	return &PbxdeviceprofilesByIDPatchParams{
		HTTPClient: client,
	}
}

/*PbxdeviceprofilesByIDPatchParams contains all the parameters to send to the API endpoint
for the pbxdeviceprofiles by Id patch operation typically these are written to a http.Request
*/
type PbxdeviceprofilesByIDPatchParams struct {

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

// WithTimeout adds the timeout to the pbxdeviceprofiles by Id patch params
func (o *PbxdeviceprofilesByIDPatchParams) WithTimeout(timeout time.Duration) *PbxdeviceprofilesByIDPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pbxdeviceprofiles by Id patch params
func (o *PbxdeviceprofilesByIDPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pbxdeviceprofiles by Id patch params
func (o *PbxdeviceprofilesByIDPatchParams) WithContext(ctx context.Context) *PbxdeviceprofilesByIDPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pbxdeviceprofiles by Id patch params
func (o *PbxdeviceprofilesByIDPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pbxdeviceprofiles by Id patch params
func (o *PbxdeviceprofilesByIDPatchParams) WithHTTPClient(client *http.Client) *PbxdeviceprofilesByIDPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pbxdeviceprofiles by Id patch params
func (o *PbxdeviceprofilesByIDPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pbxdeviceprofiles by Id patch params
func (o *PbxdeviceprofilesByIDPatchParams) WithBody(body string) *PbxdeviceprofilesByIDPatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pbxdeviceprofiles by Id patch params
func (o *PbxdeviceprofilesByIDPatchParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the pbxdeviceprofiles by Id patch params
func (o *PbxdeviceprofilesByIDPatchParams) WithContentType(contentType string) *PbxdeviceprofilesByIDPatchParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the pbxdeviceprofiles by Id patch params
func (o *PbxdeviceprofilesByIDPatchParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the pbxdeviceprofiles by Id patch params
func (o *PbxdeviceprofilesByIDPatchParams) WithID(id string) *PbxdeviceprofilesByIDPatchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the pbxdeviceprofiles by Id patch params
func (o *PbxdeviceprofilesByIDPatchParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PbxdeviceprofilesByIDPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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