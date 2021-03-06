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

// NewReplaceChangeaspecificBillingZoneParams creates a new ReplaceChangeaspecificBillingZoneParams object
// with the default values initialized.
func NewReplaceChangeaspecificBillingZoneParams() *ReplaceChangeaspecificBillingZoneParams {
	var ()
	return &ReplaceChangeaspecificBillingZoneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceChangeaspecificBillingZoneParamsWithTimeout creates a new ReplaceChangeaspecificBillingZoneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceChangeaspecificBillingZoneParamsWithTimeout(timeout time.Duration) *ReplaceChangeaspecificBillingZoneParams {
	var ()
	return &ReplaceChangeaspecificBillingZoneParams{

		timeout: timeout,
	}
}

// NewReplaceChangeaspecificBillingZoneParamsWithContext creates a new ReplaceChangeaspecificBillingZoneParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceChangeaspecificBillingZoneParamsWithContext(ctx context.Context) *ReplaceChangeaspecificBillingZoneParams {
	var ()
	return &ReplaceChangeaspecificBillingZoneParams{

		Context: ctx,
	}
}

// NewReplaceChangeaspecificBillingZoneParamsWithHTTPClient creates a new ReplaceChangeaspecificBillingZoneParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceChangeaspecificBillingZoneParamsWithHTTPClient(client *http.Client) *ReplaceChangeaspecificBillingZoneParams {
	var ()
	return &ReplaceChangeaspecificBillingZoneParams{
		HTTPClient: client,
	}
}

/*ReplaceChangeaspecificBillingZoneParams contains all the parameters to send to the API endpoint
for the replace changeaspecific billing zone operation typically these are written to a http.Request
*/
type ReplaceChangeaspecificBillingZoneParams struct {

	/*Body*/
	Body *Replace1changeaspecificBillingZoneRequest
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

// WithTimeout adds the timeout to the replace changeaspecific billing zone params
func (o *ReplaceChangeaspecificBillingZoneParams) WithTimeout(timeout time.Duration) *ReplaceChangeaspecificBillingZoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace changeaspecific billing zone params
func (o *ReplaceChangeaspecificBillingZoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace changeaspecific billing zone params
func (o *ReplaceChangeaspecificBillingZoneParams) WithContext(ctx context.Context) *ReplaceChangeaspecificBillingZoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace changeaspecific billing zone params
func (o *ReplaceChangeaspecificBillingZoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace changeaspecific billing zone params
func (o *ReplaceChangeaspecificBillingZoneParams) WithHTTPClient(client *http.Client) *ReplaceChangeaspecificBillingZoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace changeaspecific billing zone params
func (o *ReplaceChangeaspecificBillingZoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace changeaspecific billing zone params
func (o *ReplaceChangeaspecificBillingZoneParams) WithBody(body *Replace1changeaspecificBillingZoneRequest) *ReplaceChangeaspecificBillingZoneParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace changeaspecific billing zone params
func (o *ReplaceChangeaspecificBillingZoneParams) SetBody(body *Replace1changeaspecificBillingZoneRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the replace changeaspecific billing zone params
func (o *ReplaceChangeaspecificBillingZoneParams) WithContentType(contentType string) *ReplaceChangeaspecificBillingZoneParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the replace changeaspecific billing zone params
func (o *ReplaceChangeaspecificBillingZoneParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the replace changeaspecific billing zone params
func (o *ReplaceChangeaspecificBillingZoneParams) WithID(id string) *ReplaceChangeaspecificBillingZoneParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replace changeaspecific billing zone params
func (o *ReplaceChangeaspecificBillingZoneParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceChangeaspecificBillingZoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
