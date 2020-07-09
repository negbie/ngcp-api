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

// NewReplaceChangeaspecificPreferencesMetaEntryParams creates a new ReplaceChangeaspecificPreferencesMetaEntryParams object
// with the default values initialized.
func NewReplaceChangeaspecificPreferencesMetaEntryParams() *ReplaceChangeaspecificPreferencesMetaEntryParams {
	var ()
	return &ReplaceChangeaspecificPreferencesMetaEntryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceChangeaspecificPreferencesMetaEntryParamsWithTimeout creates a new ReplaceChangeaspecificPreferencesMetaEntryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceChangeaspecificPreferencesMetaEntryParamsWithTimeout(timeout time.Duration) *ReplaceChangeaspecificPreferencesMetaEntryParams {
	var ()
	return &ReplaceChangeaspecificPreferencesMetaEntryParams{

		timeout: timeout,
	}
}

// NewReplaceChangeaspecificPreferencesMetaEntryParamsWithContext creates a new ReplaceChangeaspecificPreferencesMetaEntryParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceChangeaspecificPreferencesMetaEntryParamsWithContext(ctx context.Context) *ReplaceChangeaspecificPreferencesMetaEntryParams {
	var ()
	return &ReplaceChangeaspecificPreferencesMetaEntryParams{

		Context: ctx,
	}
}

// NewReplaceChangeaspecificPreferencesMetaEntryParamsWithHTTPClient creates a new ReplaceChangeaspecificPreferencesMetaEntryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceChangeaspecificPreferencesMetaEntryParamsWithHTTPClient(client *http.Client) *ReplaceChangeaspecificPreferencesMetaEntryParams {
	var ()
	return &ReplaceChangeaspecificPreferencesMetaEntryParams{
		HTTPClient: client,
	}
}

/*ReplaceChangeaspecificPreferencesMetaEntryParams contains all the parameters to send to the API endpoint
for the replace changeaspecific preferences meta entry operation typically these are written to a http.Request
*/
type ReplaceChangeaspecificPreferencesMetaEntryParams struct {

	/*Body*/
	Body *Replace1changeaspecificPreferencesMetaEntryRequest
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

// WithTimeout adds the timeout to the replace changeaspecific preferences meta entry params
func (o *ReplaceChangeaspecificPreferencesMetaEntryParams) WithTimeout(timeout time.Duration) *ReplaceChangeaspecificPreferencesMetaEntryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace changeaspecific preferences meta entry params
func (o *ReplaceChangeaspecificPreferencesMetaEntryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace changeaspecific preferences meta entry params
func (o *ReplaceChangeaspecificPreferencesMetaEntryParams) WithContext(ctx context.Context) *ReplaceChangeaspecificPreferencesMetaEntryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace changeaspecific preferences meta entry params
func (o *ReplaceChangeaspecificPreferencesMetaEntryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace changeaspecific preferences meta entry params
func (o *ReplaceChangeaspecificPreferencesMetaEntryParams) WithHTTPClient(client *http.Client) *ReplaceChangeaspecificPreferencesMetaEntryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace changeaspecific preferences meta entry params
func (o *ReplaceChangeaspecificPreferencesMetaEntryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace changeaspecific preferences meta entry params
func (o *ReplaceChangeaspecificPreferencesMetaEntryParams) WithBody(body *Replace1changeaspecificPreferencesMetaEntryRequest) *ReplaceChangeaspecificPreferencesMetaEntryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace changeaspecific preferences meta entry params
func (o *ReplaceChangeaspecificPreferencesMetaEntryParams) SetBody(body *Replace1changeaspecificPreferencesMetaEntryRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the replace changeaspecific preferences meta entry params
func (o *ReplaceChangeaspecificPreferencesMetaEntryParams) WithContentType(contentType string) *ReplaceChangeaspecificPreferencesMetaEntryParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the replace changeaspecific preferences meta entry params
func (o *ReplaceChangeaspecificPreferencesMetaEntryParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the replace changeaspecific preferences meta entry params
func (o *ReplaceChangeaspecificPreferencesMetaEntryParams) WithID(id string) *ReplaceChangeaspecificPreferencesMetaEntryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replace changeaspecific preferences meta entry params
func (o *ReplaceChangeaspecificPreferencesMetaEntryParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceChangeaspecificPreferencesMetaEntryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
