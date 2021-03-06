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

// NewGetaspecificCallRecordingParams creates a new GetaspecificCallRecordingParams object
// with the default values initialized.
func NewGetaspecificCallRecordingParams() *GetaspecificCallRecordingParams {
	var ()
	return &GetaspecificCallRecordingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificCallRecordingParamsWithTimeout creates a new GetaspecificCallRecordingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificCallRecordingParamsWithTimeout(timeout time.Duration) *GetaspecificCallRecordingParams {
	var ()
	return &GetaspecificCallRecordingParams{

		timeout: timeout,
	}
}

// NewGetaspecificCallRecordingParamsWithContext creates a new GetaspecificCallRecordingParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificCallRecordingParamsWithContext(ctx context.Context) *GetaspecificCallRecordingParams {
	var ()
	return &GetaspecificCallRecordingParams{

		Context: ctx,
	}
}

// NewGetaspecificCallRecordingParamsWithHTTPClient creates a new GetaspecificCallRecordingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificCallRecordingParamsWithHTTPClient(client *http.Client) *GetaspecificCallRecordingParams {
	var ()
	return &GetaspecificCallRecordingParams{
		HTTPClient: client,
	}
}

/*GetaspecificCallRecordingParams contains all the parameters to send to the API endpoint
for the getaspecific call recording operation typically these are written to a http.Request
*/
type GetaspecificCallRecordingParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific call recording params
func (o *GetaspecificCallRecordingParams) WithTimeout(timeout time.Duration) *GetaspecificCallRecordingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific call recording params
func (o *GetaspecificCallRecordingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific call recording params
func (o *GetaspecificCallRecordingParams) WithContext(ctx context.Context) *GetaspecificCallRecordingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific call recording params
func (o *GetaspecificCallRecordingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific call recording params
func (o *GetaspecificCallRecordingParams) WithHTTPClient(client *http.Client) *GetaspecificCallRecordingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific call recording params
func (o *GetaspecificCallRecordingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific call recording params
func (o *GetaspecificCallRecordingParams) WithID(id string) *GetaspecificCallRecordingParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific call recording params
func (o *GetaspecificCallRecordingParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificCallRecordingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
