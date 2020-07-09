// Code generated by go-swagger; DO NOT EDIT.

package pbxdevices

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

// NewCreateanewPbxDeviceParams creates a new CreateanewPbxDeviceParams object
// with the default values initialized.
func NewCreateanewPbxDeviceParams() *CreateanewPbxDeviceParams {
	var ()
	return &CreateanewPbxDeviceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateanewPbxDeviceParamsWithTimeout creates a new CreateanewPbxDeviceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateanewPbxDeviceParamsWithTimeout(timeout time.Duration) *CreateanewPbxDeviceParams {
	var ()
	return &CreateanewPbxDeviceParams{

		timeout: timeout,
	}
}

// NewCreateanewPbxDeviceParamsWithContext creates a new CreateanewPbxDeviceParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateanewPbxDeviceParamsWithContext(ctx context.Context) *CreateanewPbxDeviceParams {
	var ()
	return &CreateanewPbxDeviceParams{

		Context: ctx,
	}
}

// NewCreateanewPbxDeviceParamsWithHTTPClient creates a new CreateanewPbxDeviceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateanewPbxDeviceParamsWithHTTPClient(client *http.Client) *CreateanewPbxDeviceParams {
	var ()
	return &CreateanewPbxDeviceParams{
		HTTPClient: client,
	}
}

/*CreateanewPbxDeviceParams contains all the parameters to send to the API endpoint
for the createanew pbx device operation typically these are written to a http.Request
*/
type CreateanewPbxDeviceParams struct {

	/*Body*/
	Body *models.CreateanewPbxDeviceRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the createanew pbx device params
func (o *CreateanewPbxDeviceParams) WithTimeout(timeout time.Duration) *CreateanewPbxDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the createanew pbx device params
func (o *CreateanewPbxDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the createanew pbx device params
func (o *CreateanewPbxDeviceParams) WithContext(ctx context.Context) *CreateanewPbxDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the createanew pbx device params
func (o *CreateanewPbxDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the createanew pbx device params
func (o *CreateanewPbxDeviceParams) WithHTTPClient(client *http.Client) *CreateanewPbxDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the createanew pbx device params
func (o *CreateanewPbxDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the createanew pbx device params
func (o *CreateanewPbxDeviceParams) WithBody(body *models.CreateanewPbxDeviceRequest) *CreateanewPbxDeviceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the createanew pbx device params
func (o *CreateanewPbxDeviceParams) SetBody(body *models.CreateanewPbxDeviceRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the createanew pbx device params
func (o *CreateanewPbxDeviceParams) WithContentType(contentType string) *CreateanewPbxDeviceParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the createanew pbx device params
func (o *CreateanewPbxDeviceParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateanewPbxDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
