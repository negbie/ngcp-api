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

// NewChangeaspecificEmergencyMappingParams creates a new ChangeaspecificEmergencyMappingParams object
// with the default values initialized.
func NewChangeaspecificEmergencyMappingParams() *ChangeaspecificEmergencyMappingParams {
	var ()
	return &ChangeaspecificEmergencyMappingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeaspecificEmergencyMappingParamsWithTimeout creates a new ChangeaspecificEmergencyMappingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeaspecificEmergencyMappingParamsWithTimeout(timeout time.Duration) *ChangeaspecificEmergencyMappingParams {
	var ()
	return &ChangeaspecificEmergencyMappingParams{

		timeout: timeout,
	}
}

// NewChangeaspecificEmergencyMappingParamsWithContext creates a new ChangeaspecificEmergencyMappingParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeaspecificEmergencyMappingParamsWithContext(ctx context.Context) *ChangeaspecificEmergencyMappingParams {
	var ()
	return &ChangeaspecificEmergencyMappingParams{

		Context: ctx,
	}
}

// NewChangeaspecificEmergencyMappingParamsWithHTTPClient creates a new ChangeaspecificEmergencyMappingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeaspecificEmergencyMappingParamsWithHTTPClient(client *http.Client) *ChangeaspecificEmergencyMappingParams {
	var ()
	return &ChangeaspecificEmergencyMappingParams{
		HTTPClient: client,
	}
}

/*ChangeaspecificEmergencyMappingParams contains all the parameters to send to the API endpoint
for the changeaspecific emergency mapping operation typically these are written to a http.Request
*/
type ChangeaspecificEmergencyMappingParams struct {

	/*Body*/
	Body string
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

// WithTimeout adds the timeout to the changeaspecific emergency mapping params
func (o *ChangeaspecificEmergencyMappingParams) WithTimeout(timeout time.Duration) *ChangeaspecificEmergencyMappingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the changeaspecific emergency mapping params
func (o *ChangeaspecificEmergencyMappingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the changeaspecific emergency mapping params
func (o *ChangeaspecificEmergencyMappingParams) WithContext(ctx context.Context) *ChangeaspecificEmergencyMappingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the changeaspecific emergency mapping params
func (o *ChangeaspecificEmergencyMappingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the changeaspecific emergency mapping params
func (o *ChangeaspecificEmergencyMappingParams) WithHTTPClient(client *http.Client) *ChangeaspecificEmergencyMappingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the changeaspecific emergency mapping params
func (o *ChangeaspecificEmergencyMappingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the changeaspecific emergency mapping params
func (o *ChangeaspecificEmergencyMappingParams) WithBody(body string) *ChangeaspecificEmergencyMappingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the changeaspecific emergency mapping params
func (o *ChangeaspecificEmergencyMappingParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the changeaspecific emergency mapping params
func (o *ChangeaspecificEmergencyMappingParams) WithContentType(contentType string) *ChangeaspecificEmergencyMappingParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the changeaspecific emergency mapping params
func (o *ChangeaspecificEmergencyMappingParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the changeaspecific emergency mapping params
func (o *ChangeaspecificEmergencyMappingParams) WithID(id string) *ChangeaspecificEmergencyMappingParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the changeaspecific emergency mapping params
func (o *ChangeaspecificEmergencyMappingParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeaspecificEmergencyMappingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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