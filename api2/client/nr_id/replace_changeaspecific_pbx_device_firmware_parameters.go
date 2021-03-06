// Code generated by go-swagger; DO NOT EDIT.

package nr_id

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

// NewReplaceChangeaspecificPbxDeviceFirmwareParams creates a new ReplaceChangeaspecificPbxDeviceFirmwareParams object
// with the default values initialized.
func NewReplaceChangeaspecificPbxDeviceFirmwareParams() *ReplaceChangeaspecificPbxDeviceFirmwareParams {
	var ()
	return &ReplaceChangeaspecificPbxDeviceFirmwareParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceChangeaspecificPbxDeviceFirmwareParamsWithTimeout creates a new ReplaceChangeaspecificPbxDeviceFirmwareParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceChangeaspecificPbxDeviceFirmwareParamsWithTimeout(timeout time.Duration) *ReplaceChangeaspecificPbxDeviceFirmwareParams {
	var ()
	return &ReplaceChangeaspecificPbxDeviceFirmwareParams{

		timeout: timeout,
	}
}

// NewReplaceChangeaspecificPbxDeviceFirmwareParamsWithContext creates a new ReplaceChangeaspecificPbxDeviceFirmwareParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceChangeaspecificPbxDeviceFirmwareParamsWithContext(ctx context.Context) *ReplaceChangeaspecificPbxDeviceFirmwareParams {
	var ()
	return &ReplaceChangeaspecificPbxDeviceFirmwareParams{

		Context: ctx,
	}
}

// NewReplaceChangeaspecificPbxDeviceFirmwareParamsWithHTTPClient creates a new ReplaceChangeaspecificPbxDeviceFirmwareParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceChangeaspecificPbxDeviceFirmwareParamsWithHTTPClient(client *http.Client) *ReplaceChangeaspecificPbxDeviceFirmwareParams {
	var ()
	return &ReplaceChangeaspecificPbxDeviceFirmwareParams{
		HTTPClient: client,
	}
}

/*ReplaceChangeaspecificPbxDeviceFirmwareParams contains all the parameters to send to the API endpoint
for the replace changeaspecific pbx device firmware operation typically these are written to a http.Request
*/
type ReplaceChangeaspecificPbxDeviceFirmwareParams struct {

	/*Body*/
	Body *Replace1changeaspecificPbxDeviceFirmwareRequest
	/*ContentType*/
	ContentType string
	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace changeaspecific pbx device firmware params
func (o *ReplaceChangeaspecificPbxDeviceFirmwareParams) WithTimeout(timeout time.Duration) *ReplaceChangeaspecificPbxDeviceFirmwareParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace changeaspecific pbx device firmware params
func (o *ReplaceChangeaspecificPbxDeviceFirmwareParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace changeaspecific pbx device firmware params
func (o *ReplaceChangeaspecificPbxDeviceFirmwareParams) WithContext(ctx context.Context) *ReplaceChangeaspecificPbxDeviceFirmwareParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace changeaspecific pbx device firmware params
func (o *ReplaceChangeaspecificPbxDeviceFirmwareParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace changeaspecific pbx device firmware params
func (o *ReplaceChangeaspecificPbxDeviceFirmwareParams) WithHTTPClient(client *http.Client) *ReplaceChangeaspecificPbxDeviceFirmwareParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace changeaspecific pbx device firmware params
func (o *ReplaceChangeaspecificPbxDeviceFirmwareParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace changeaspecific pbx device firmware params
func (o *ReplaceChangeaspecificPbxDeviceFirmwareParams) WithBody(body *Replace1changeaspecificPbxDeviceFirmwareRequest) *ReplaceChangeaspecificPbxDeviceFirmwareParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace changeaspecific pbx device firmware params
func (o *ReplaceChangeaspecificPbxDeviceFirmwareParams) SetBody(body *Replace1changeaspecificPbxDeviceFirmwareRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the replace changeaspecific pbx device firmware params
func (o *ReplaceChangeaspecificPbxDeviceFirmwareParams) WithContentType(contentType string) *ReplaceChangeaspecificPbxDeviceFirmwareParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the replace changeaspecific pbx device firmware params
func (o *ReplaceChangeaspecificPbxDeviceFirmwareParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the replace changeaspecific pbx device firmware params
func (o *ReplaceChangeaspecificPbxDeviceFirmwareParams) WithID(id string) *ReplaceChangeaspecificPbxDeviceFirmwareParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replace changeaspecific pbx device firmware params
func (o *ReplaceChangeaspecificPbxDeviceFirmwareParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceChangeaspecificPbxDeviceFirmwareParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
