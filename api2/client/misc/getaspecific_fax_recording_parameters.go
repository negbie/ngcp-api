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

// NewGetaspecificFaxRecordingParams creates a new GetaspecificFaxRecordingParams object
// with the default values initialized.
func NewGetaspecificFaxRecordingParams() *GetaspecificFaxRecordingParams {
	var ()
	return &GetaspecificFaxRecordingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificFaxRecordingParamsWithTimeout creates a new GetaspecificFaxRecordingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificFaxRecordingParamsWithTimeout(timeout time.Duration) *GetaspecificFaxRecordingParams {
	var ()
	return &GetaspecificFaxRecordingParams{

		timeout: timeout,
	}
}

// NewGetaspecificFaxRecordingParamsWithContext creates a new GetaspecificFaxRecordingParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificFaxRecordingParamsWithContext(ctx context.Context) *GetaspecificFaxRecordingParams {
	var ()
	return &GetaspecificFaxRecordingParams{

		Context: ctx,
	}
}

// NewGetaspecificFaxRecordingParamsWithHTTPClient creates a new GetaspecificFaxRecordingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificFaxRecordingParamsWithHTTPClient(client *http.Client) *GetaspecificFaxRecordingParams {
	var ()
	return &GetaspecificFaxRecordingParams{
		HTTPClient: client,
	}
}

/*GetaspecificFaxRecordingParams contains all the parameters to send to the API endpoint
for the getaspecific fax recording operation typically these are written to a http.Request
*/
type GetaspecificFaxRecordingParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific fax recording params
func (o *GetaspecificFaxRecordingParams) WithTimeout(timeout time.Duration) *GetaspecificFaxRecordingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific fax recording params
func (o *GetaspecificFaxRecordingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific fax recording params
func (o *GetaspecificFaxRecordingParams) WithContext(ctx context.Context) *GetaspecificFaxRecordingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific fax recording params
func (o *GetaspecificFaxRecordingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific fax recording params
func (o *GetaspecificFaxRecordingParams) WithHTTPClient(client *http.Client) *GetaspecificFaxRecordingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific fax recording params
func (o *GetaspecificFaxRecordingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific fax recording params
func (o *GetaspecificFaxRecordingParams) WithID(id string) *GetaspecificFaxRecordingParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific fax recording params
func (o *GetaspecificFaxRecordingParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificFaxRecordingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
