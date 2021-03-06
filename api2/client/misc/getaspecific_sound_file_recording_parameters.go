// Code generated by go-swagger; DO NOT EDIT.

package misc

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

// NewGetaspecificSoundFileRecordingParams creates a new GetaspecificSoundFileRecordingParams object
// with the default values initialized.
func NewGetaspecificSoundFileRecordingParams() *GetaspecificSoundFileRecordingParams {
	var ()
	return &GetaspecificSoundFileRecordingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificSoundFileRecordingParamsWithTimeout creates a new GetaspecificSoundFileRecordingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificSoundFileRecordingParamsWithTimeout(timeout time.Duration) *GetaspecificSoundFileRecordingParams {
	var ()
	return &GetaspecificSoundFileRecordingParams{

		timeout: timeout,
	}
}

// NewGetaspecificSoundFileRecordingParamsWithContext creates a new GetaspecificSoundFileRecordingParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificSoundFileRecordingParamsWithContext(ctx context.Context) *GetaspecificSoundFileRecordingParams {
	var ()
	return &GetaspecificSoundFileRecordingParams{

		Context: ctx,
	}
}

// NewGetaspecificSoundFileRecordingParamsWithHTTPClient creates a new GetaspecificSoundFileRecordingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificSoundFileRecordingParamsWithHTTPClient(client *http.Client) *GetaspecificSoundFileRecordingParams {
	var ()
	return &GetaspecificSoundFileRecordingParams{
		HTTPClient: client,
	}
}

/*GetaspecificSoundFileRecordingParams contains all the parameters to send to the API endpoint
for the getaspecific sound file recording operation typically these are written to a http.Request
*/
type GetaspecificSoundFileRecordingParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific sound file recording params
func (o *GetaspecificSoundFileRecordingParams) WithTimeout(timeout time.Duration) *GetaspecificSoundFileRecordingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific sound file recording params
func (o *GetaspecificSoundFileRecordingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific sound file recording params
func (o *GetaspecificSoundFileRecordingParams) WithContext(ctx context.Context) *GetaspecificSoundFileRecordingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific sound file recording params
func (o *GetaspecificSoundFileRecordingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific sound file recording params
func (o *GetaspecificSoundFileRecordingParams) WithHTTPClient(client *http.Client) *GetaspecificSoundFileRecordingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific sound file recording params
func (o *GetaspecificSoundFileRecordingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific sound file recording params
func (o *GetaspecificSoundFileRecordingParams) WithID(id string) *GetaspecificSoundFileRecordingParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific sound file recording params
func (o *GetaspecificSoundFileRecordingParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificSoundFileRecordingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
