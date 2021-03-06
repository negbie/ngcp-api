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

// NewReplaceChangeaspecificPbxDeviceConfigParams creates a new ReplaceChangeaspecificPbxDeviceConfigParams object
// with the default values initialized.
func NewReplaceChangeaspecificPbxDeviceConfigParams() *ReplaceChangeaspecificPbxDeviceConfigParams {
	var ()
	return &ReplaceChangeaspecificPbxDeviceConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceChangeaspecificPbxDeviceConfigParamsWithTimeout creates a new ReplaceChangeaspecificPbxDeviceConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceChangeaspecificPbxDeviceConfigParamsWithTimeout(timeout time.Duration) *ReplaceChangeaspecificPbxDeviceConfigParams {
	var ()
	return &ReplaceChangeaspecificPbxDeviceConfigParams{

		timeout: timeout,
	}
}

// NewReplaceChangeaspecificPbxDeviceConfigParamsWithContext creates a new ReplaceChangeaspecificPbxDeviceConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceChangeaspecificPbxDeviceConfigParamsWithContext(ctx context.Context) *ReplaceChangeaspecificPbxDeviceConfigParams {
	var ()
	return &ReplaceChangeaspecificPbxDeviceConfigParams{

		Context: ctx,
	}
}

// NewReplaceChangeaspecificPbxDeviceConfigParamsWithHTTPClient creates a new ReplaceChangeaspecificPbxDeviceConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceChangeaspecificPbxDeviceConfigParamsWithHTTPClient(client *http.Client) *ReplaceChangeaspecificPbxDeviceConfigParams {
	var ()
	return &ReplaceChangeaspecificPbxDeviceConfigParams{
		HTTPClient: client,
	}
}

/*ReplaceChangeaspecificPbxDeviceConfigParams contains all the parameters to send to the API endpoint
for the replace changeaspecific pbx device config operation typically these are written to a http.Request
*/
type ReplaceChangeaspecificPbxDeviceConfigParams struct {

	/*Body*/
	Body *Replace1changeaspecificPbxDeviceConfigRequest
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

// WithTimeout adds the timeout to the replace changeaspecific pbx device config params
func (o *ReplaceChangeaspecificPbxDeviceConfigParams) WithTimeout(timeout time.Duration) *ReplaceChangeaspecificPbxDeviceConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace changeaspecific pbx device config params
func (o *ReplaceChangeaspecificPbxDeviceConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace changeaspecific pbx device config params
func (o *ReplaceChangeaspecificPbxDeviceConfigParams) WithContext(ctx context.Context) *ReplaceChangeaspecificPbxDeviceConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace changeaspecific pbx device config params
func (o *ReplaceChangeaspecificPbxDeviceConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace changeaspecific pbx device config params
func (o *ReplaceChangeaspecificPbxDeviceConfigParams) WithHTTPClient(client *http.Client) *ReplaceChangeaspecificPbxDeviceConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace changeaspecific pbx device config params
func (o *ReplaceChangeaspecificPbxDeviceConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace changeaspecific pbx device config params
func (o *ReplaceChangeaspecificPbxDeviceConfigParams) WithBody(body *Replace1changeaspecificPbxDeviceConfigRequest) *ReplaceChangeaspecificPbxDeviceConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace changeaspecific pbx device config params
func (o *ReplaceChangeaspecificPbxDeviceConfigParams) SetBody(body *Replace1changeaspecificPbxDeviceConfigRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the replace changeaspecific pbx device config params
func (o *ReplaceChangeaspecificPbxDeviceConfigParams) WithContentType(contentType string) *ReplaceChangeaspecificPbxDeviceConfigParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the replace changeaspecific pbx device config params
func (o *ReplaceChangeaspecificPbxDeviceConfigParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the replace changeaspecific pbx device config params
func (o *ReplaceChangeaspecificPbxDeviceConfigParams) WithID(id string) *ReplaceChangeaspecificPbxDeviceConfigParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replace changeaspecific pbx device config params
func (o *ReplaceChangeaspecificPbxDeviceConfigParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceChangeaspecificPbxDeviceConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
