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

// NewPbxdeviceprofilepreferencesByIDPutParams creates a new PbxdeviceprofilepreferencesByIDPutParams object
// with the default values initialized.
func NewPbxdeviceprofilepreferencesByIDPutParams() *PbxdeviceprofilepreferencesByIDPutParams {
	var ()
	return &PbxdeviceprofilepreferencesByIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPbxdeviceprofilepreferencesByIDPutParamsWithTimeout creates a new PbxdeviceprofilepreferencesByIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPbxdeviceprofilepreferencesByIDPutParamsWithTimeout(timeout time.Duration) *PbxdeviceprofilepreferencesByIDPutParams {
	var ()
	return &PbxdeviceprofilepreferencesByIDPutParams{

		timeout: timeout,
	}
}

// NewPbxdeviceprofilepreferencesByIDPutParamsWithContext creates a new PbxdeviceprofilepreferencesByIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewPbxdeviceprofilepreferencesByIDPutParamsWithContext(ctx context.Context) *PbxdeviceprofilepreferencesByIDPutParams {
	var ()
	return &PbxdeviceprofilepreferencesByIDPutParams{

		Context: ctx,
	}
}

// NewPbxdeviceprofilepreferencesByIDPutParamsWithHTTPClient creates a new PbxdeviceprofilepreferencesByIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPbxdeviceprofilepreferencesByIDPutParamsWithHTTPClient(client *http.Client) *PbxdeviceprofilepreferencesByIDPutParams {
	var ()
	return &PbxdeviceprofilepreferencesByIDPutParams{
		HTTPClient: client,
	}
}

/*PbxdeviceprofilepreferencesByIDPutParams contains all the parameters to send to the API endpoint
for the pbxdeviceprofilepreferences by Id put operation typically these are written to a http.Request
*/
type PbxdeviceprofilepreferencesByIDPutParams struct {

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

// WithTimeout adds the timeout to the pbxdeviceprofilepreferences by Id put params
func (o *PbxdeviceprofilepreferencesByIDPutParams) WithTimeout(timeout time.Duration) *PbxdeviceprofilepreferencesByIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pbxdeviceprofilepreferences by Id put params
func (o *PbxdeviceprofilepreferencesByIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pbxdeviceprofilepreferences by Id put params
func (o *PbxdeviceprofilepreferencesByIDPutParams) WithContext(ctx context.Context) *PbxdeviceprofilepreferencesByIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pbxdeviceprofilepreferences by Id put params
func (o *PbxdeviceprofilepreferencesByIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pbxdeviceprofilepreferences by Id put params
func (o *PbxdeviceprofilepreferencesByIDPutParams) WithHTTPClient(client *http.Client) *PbxdeviceprofilepreferencesByIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pbxdeviceprofilepreferences by Id put params
func (o *PbxdeviceprofilepreferencesByIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pbxdeviceprofilepreferences by Id put params
func (o *PbxdeviceprofilepreferencesByIDPutParams) WithBody(body string) *PbxdeviceprofilepreferencesByIDPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pbxdeviceprofilepreferences by Id put params
func (o *PbxdeviceprofilepreferencesByIDPutParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the pbxdeviceprofilepreferences by Id put params
func (o *PbxdeviceprofilepreferencesByIDPutParams) WithContentType(contentType string) *PbxdeviceprofilepreferencesByIDPutParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the pbxdeviceprofilepreferences by Id put params
func (o *PbxdeviceprofilepreferencesByIDPutParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the pbxdeviceprofilepreferences by Id put params
func (o *PbxdeviceprofilepreferencesByIDPutParams) WithID(id string) *PbxdeviceprofilepreferencesByIDPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the pbxdeviceprofilepreferences by Id put params
func (o *PbxdeviceprofilepreferencesByIDPutParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PbxdeviceprofilepreferencesByIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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