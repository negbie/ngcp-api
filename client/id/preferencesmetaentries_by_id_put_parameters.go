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

// NewPreferencesmetaentriesByIDPutParams creates a new PreferencesmetaentriesByIDPutParams object
// with the default values initialized.
func NewPreferencesmetaentriesByIDPutParams() *PreferencesmetaentriesByIDPutParams {
	var ()
	return &PreferencesmetaentriesByIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPreferencesmetaentriesByIDPutParamsWithTimeout creates a new PreferencesmetaentriesByIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPreferencesmetaentriesByIDPutParamsWithTimeout(timeout time.Duration) *PreferencesmetaentriesByIDPutParams {
	var ()
	return &PreferencesmetaentriesByIDPutParams{

		timeout: timeout,
	}
}

// NewPreferencesmetaentriesByIDPutParamsWithContext creates a new PreferencesmetaentriesByIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewPreferencesmetaentriesByIDPutParamsWithContext(ctx context.Context) *PreferencesmetaentriesByIDPutParams {
	var ()
	return &PreferencesmetaentriesByIDPutParams{

		Context: ctx,
	}
}

// NewPreferencesmetaentriesByIDPutParamsWithHTTPClient creates a new PreferencesmetaentriesByIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPreferencesmetaentriesByIDPutParamsWithHTTPClient(client *http.Client) *PreferencesmetaentriesByIDPutParams {
	var ()
	return &PreferencesmetaentriesByIDPutParams{
		HTTPClient: client,
	}
}

/*PreferencesmetaentriesByIDPutParams contains all the parameters to send to the API endpoint
for the preferencesmetaentries by Id put operation typically these are written to a http.Request
*/
type PreferencesmetaentriesByIDPutParams struct {

	/*Body*/
	Body *Replace1changeaspecificPreferencesMetaEntryRequest
	/*ContentType*/
	ContentType string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the preferencesmetaentries by Id put params
func (o *PreferencesmetaentriesByIDPutParams) WithTimeout(timeout time.Duration) *PreferencesmetaentriesByIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the preferencesmetaentries by Id put params
func (o *PreferencesmetaentriesByIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the preferencesmetaentries by Id put params
func (o *PreferencesmetaentriesByIDPutParams) WithContext(ctx context.Context) *PreferencesmetaentriesByIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the preferencesmetaentries by Id put params
func (o *PreferencesmetaentriesByIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the preferencesmetaentries by Id put params
func (o *PreferencesmetaentriesByIDPutParams) WithHTTPClient(client *http.Client) *PreferencesmetaentriesByIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the preferencesmetaentries by Id put params
func (o *PreferencesmetaentriesByIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the preferencesmetaentries by Id put params
func (o *PreferencesmetaentriesByIDPutParams) WithBody(body *Replace1changeaspecificPreferencesMetaEntryRequest) *PreferencesmetaentriesByIDPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the preferencesmetaentries by Id put params
func (o *PreferencesmetaentriesByIDPutParams) SetBody(body *Replace1changeaspecificPreferencesMetaEntryRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the preferencesmetaentries by Id put params
func (o *PreferencesmetaentriesByIDPutParams) WithContentType(contentType string) *PreferencesmetaentriesByIDPutParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the preferencesmetaentries by Id put params
func (o *PreferencesmetaentriesByIDPutParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the preferencesmetaentries by Id put params
func (o *PreferencesmetaentriesByIDPutParams) WithID(id string) *PreferencesmetaentriesByIDPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the preferencesmetaentries by Id put params
func (o *PreferencesmetaentriesByIDPutParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PreferencesmetaentriesByIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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