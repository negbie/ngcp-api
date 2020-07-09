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

// NewReplaceChangeaspecificSoundFileParams creates a new ReplaceChangeaspecificSoundFileParams object
// with the default values initialized.
func NewReplaceChangeaspecificSoundFileParams() *ReplaceChangeaspecificSoundFileParams {
	var ()
	return &ReplaceChangeaspecificSoundFileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceChangeaspecificSoundFileParamsWithTimeout creates a new ReplaceChangeaspecificSoundFileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceChangeaspecificSoundFileParamsWithTimeout(timeout time.Duration) *ReplaceChangeaspecificSoundFileParams {
	var ()
	return &ReplaceChangeaspecificSoundFileParams{

		timeout: timeout,
	}
}

// NewReplaceChangeaspecificSoundFileParamsWithContext creates a new ReplaceChangeaspecificSoundFileParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceChangeaspecificSoundFileParamsWithContext(ctx context.Context) *ReplaceChangeaspecificSoundFileParams {
	var ()
	return &ReplaceChangeaspecificSoundFileParams{

		Context: ctx,
	}
}

// NewReplaceChangeaspecificSoundFileParamsWithHTTPClient creates a new ReplaceChangeaspecificSoundFileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceChangeaspecificSoundFileParamsWithHTTPClient(client *http.Client) *ReplaceChangeaspecificSoundFileParams {
	var ()
	return &ReplaceChangeaspecificSoundFileParams{
		HTTPClient: client,
	}
}

/*ReplaceChangeaspecificSoundFileParams contains all the parameters to send to the API endpoint
for the replace changeaspecific sound file operation typically these are written to a http.Request
*/
type ReplaceChangeaspecificSoundFileParams struct {

	/*Body*/
	Body *Replace1changeaspecificSoundFileRequest
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

// WithTimeout adds the timeout to the replace changeaspecific sound file params
func (o *ReplaceChangeaspecificSoundFileParams) WithTimeout(timeout time.Duration) *ReplaceChangeaspecificSoundFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace changeaspecific sound file params
func (o *ReplaceChangeaspecificSoundFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace changeaspecific sound file params
func (o *ReplaceChangeaspecificSoundFileParams) WithContext(ctx context.Context) *ReplaceChangeaspecificSoundFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace changeaspecific sound file params
func (o *ReplaceChangeaspecificSoundFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace changeaspecific sound file params
func (o *ReplaceChangeaspecificSoundFileParams) WithHTTPClient(client *http.Client) *ReplaceChangeaspecificSoundFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace changeaspecific sound file params
func (o *ReplaceChangeaspecificSoundFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace changeaspecific sound file params
func (o *ReplaceChangeaspecificSoundFileParams) WithBody(body *Replace1changeaspecificSoundFileRequest) *ReplaceChangeaspecificSoundFileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace changeaspecific sound file params
func (o *ReplaceChangeaspecificSoundFileParams) SetBody(body *Replace1changeaspecificSoundFileRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the replace changeaspecific sound file params
func (o *ReplaceChangeaspecificSoundFileParams) WithContentType(contentType string) *ReplaceChangeaspecificSoundFileParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the replace changeaspecific sound file params
func (o *ReplaceChangeaspecificSoundFileParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the replace changeaspecific sound file params
func (o *ReplaceChangeaspecificSoundFileParams) WithID(id string) *ReplaceChangeaspecificSoundFileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replace changeaspecific sound file params
func (o *ReplaceChangeaspecificSoundFileParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceChangeaspecificSoundFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
