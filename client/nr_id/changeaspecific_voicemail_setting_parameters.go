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

// NewChangeaspecificVoicemailSettingParams creates a new ChangeaspecificVoicemailSettingParams object
// with the default values initialized.
func NewChangeaspecificVoicemailSettingParams() *ChangeaspecificVoicemailSettingParams {
	var ()
	return &ChangeaspecificVoicemailSettingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeaspecificVoicemailSettingParamsWithTimeout creates a new ChangeaspecificVoicemailSettingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeaspecificVoicemailSettingParamsWithTimeout(timeout time.Duration) *ChangeaspecificVoicemailSettingParams {
	var ()
	return &ChangeaspecificVoicemailSettingParams{

		timeout: timeout,
	}
}

// NewChangeaspecificVoicemailSettingParamsWithContext creates a new ChangeaspecificVoicemailSettingParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeaspecificVoicemailSettingParamsWithContext(ctx context.Context) *ChangeaspecificVoicemailSettingParams {
	var ()
	return &ChangeaspecificVoicemailSettingParams{

		Context: ctx,
	}
}

// NewChangeaspecificVoicemailSettingParamsWithHTTPClient creates a new ChangeaspecificVoicemailSettingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeaspecificVoicemailSettingParamsWithHTTPClient(client *http.Client) *ChangeaspecificVoicemailSettingParams {
	var ()
	return &ChangeaspecificVoicemailSettingParams{
		HTTPClient: client,
	}
}

/*ChangeaspecificVoicemailSettingParams contains all the parameters to send to the API endpoint
for the changeaspecific voicemail setting operation typically these are written to a http.Request
*/
type ChangeaspecificVoicemailSettingParams struct {

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

// WithTimeout adds the timeout to the changeaspecific voicemail setting params
func (o *ChangeaspecificVoicemailSettingParams) WithTimeout(timeout time.Duration) *ChangeaspecificVoicemailSettingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the changeaspecific voicemail setting params
func (o *ChangeaspecificVoicemailSettingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the changeaspecific voicemail setting params
func (o *ChangeaspecificVoicemailSettingParams) WithContext(ctx context.Context) *ChangeaspecificVoicemailSettingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the changeaspecific voicemail setting params
func (o *ChangeaspecificVoicemailSettingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the changeaspecific voicemail setting params
func (o *ChangeaspecificVoicemailSettingParams) WithHTTPClient(client *http.Client) *ChangeaspecificVoicemailSettingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the changeaspecific voicemail setting params
func (o *ChangeaspecificVoicemailSettingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the changeaspecific voicemail setting params
func (o *ChangeaspecificVoicemailSettingParams) WithBody(body string) *ChangeaspecificVoicemailSettingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the changeaspecific voicemail setting params
func (o *ChangeaspecificVoicemailSettingParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the changeaspecific voicemail setting params
func (o *ChangeaspecificVoicemailSettingParams) WithContentType(contentType string) *ChangeaspecificVoicemailSettingParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the changeaspecific voicemail setting params
func (o *ChangeaspecificVoicemailSettingParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the changeaspecific voicemail setting params
func (o *ChangeaspecificVoicemailSettingParams) WithID(id string) *ChangeaspecificVoicemailSettingParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the changeaspecific voicemail setting params
func (o *ChangeaspecificVoicemailSettingParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeaspecificVoicemailSettingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
