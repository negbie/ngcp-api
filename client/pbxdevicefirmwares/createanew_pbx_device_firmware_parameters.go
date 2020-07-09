// Code generated by go-swagger; DO NOT EDIT.

package pbxdevicefirmwares

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

// NewCreateanewPbxDeviceFirmwareParams creates a new CreateanewPbxDeviceFirmwareParams object
// with the default values initialized.
func NewCreateanewPbxDeviceFirmwareParams() *CreateanewPbxDeviceFirmwareParams {
	var ()
	return &CreateanewPbxDeviceFirmwareParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateanewPbxDeviceFirmwareParamsWithTimeout creates a new CreateanewPbxDeviceFirmwareParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateanewPbxDeviceFirmwareParamsWithTimeout(timeout time.Duration) *CreateanewPbxDeviceFirmwareParams {
	var ()
	return &CreateanewPbxDeviceFirmwareParams{

		timeout: timeout,
	}
}

// NewCreateanewPbxDeviceFirmwareParamsWithContext creates a new CreateanewPbxDeviceFirmwareParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateanewPbxDeviceFirmwareParamsWithContext(ctx context.Context) *CreateanewPbxDeviceFirmwareParams {
	var ()
	return &CreateanewPbxDeviceFirmwareParams{

		Context: ctx,
	}
}

// NewCreateanewPbxDeviceFirmwareParamsWithHTTPClient creates a new CreateanewPbxDeviceFirmwareParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateanewPbxDeviceFirmwareParamsWithHTTPClient(client *http.Client) *CreateanewPbxDeviceFirmwareParams {
	var ()
	return &CreateanewPbxDeviceFirmwareParams{
		HTTPClient: client,
	}
}

/*CreateanewPbxDeviceFirmwareParams contains all the parameters to send to the API endpoint
for the createanew pbx device firmware operation typically these are written to a http.Request
*/
type CreateanewPbxDeviceFirmwareParams struct {

	/*Body*/
	Body *models.CreateanewPbxDeviceFirmwareRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the createanew pbx device firmware params
func (o *CreateanewPbxDeviceFirmwareParams) WithTimeout(timeout time.Duration) *CreateanewPbxDeviceFirmwareParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the createanew pbx device firmware params
func (o *CreateanewPbxDeviceFirmwareParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the createanew pbx device firmware params
func (o *CreateanewPbxDeviceFirmwareParams) WithContext(ctx context.Context) *CreateanewPbxDeviceFirmwareParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the createanew pbx device firmware params
func (o *CreateanewPbxDeviceFirmwareParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the createanew pbx device firmware params
func (o *CreateanewPbxDeviceFirmwareParams) WithHTTPClient(client *http.Client) *CreateanewPbxDeviceFirmwareParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the createanew pbx device firmware params
func (o *CreateanewPbxDeviceFirmwareParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the createanew pbx device firmware params
func (o *CreateanewPbxDeviceFirmwareParams) WithBody(body *models.CreateanewPbxDeviceFirmwareRequest) *CreateanewPbxDeviceFirmwareParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the createanew pbx device firmware params
func (o *CreateanewPbxDeviceFirmwareParams) SetBody(body *models.CreateanewPbxDeviceFirmwareRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the createanew pbx device firmware params
func (o *CreateanewPbxDeviceFirmwareParams) WithContentType(contentType string) *CreateanewPbxDeviceFirmwareParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the createanew pbx device firmware params
func (o *CreateanewPbxDeviceFirmwareParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateanewPbxDeviceFirmwareParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
