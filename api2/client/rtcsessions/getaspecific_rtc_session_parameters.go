// Code generated by go-swagger; DO NOT EDIT.

package rtcsessions

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

// NewGetaspecificRtcSessionParams creates a new GetaspecificRtcSessionParams object
// with the default values initialized.
func NewGetaspecificRtcSessionParams() *GetaspecificRtcSessionParams {
	var ()
	return &GetaspecificRtcSessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificRtcSessionParamsWithTimeout creates a new GetaspecificRtcSessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificRtcSessionParamsWithTimeout(timeout time.Duration) *GetaspecificRtcSessionParams {
	var ()
	return &GetaspecificRtcSessionParams{

		timeout: timeout,
	}
}

// NewGetaspecificRtcSessionParamsWithContext creates a new GetaspecificRtcSessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificRtcSessionParamsWithContext(ctx context.Context) *GetaspecificRtcSessionParams {
	var ()
	return &GetaspecificRtcSessionParams{

		Context: ctx,
	}
}

// NewGetaspecificRtcSessionParamsWithHTTPClient creates a new GetaspecificRtcSessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificRtcSessionParamsWithHTTPClient(client *http.Client) *GetaspecificRtcSessionParams {
	var ()
	return &GetaspecificRtcSessionParams{
		HTTPClient: client,
	}
}

/*GetaspecificRtcSessionParams contains all the parameters to send to the API endpoint
for the getaspecific rtc session operation typically these are written to a http.Request
*/
type GetaspecificRtcSessionParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific rtc session params
func (o *GetaspecificRtcSessionParams) WithTimeout(timeout time.Duration) *GetaspecificRtcSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific rtc session params
func (o *GetaspecificRtcSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific rtc session params
func (o *GetaspecificRtcSessionParams) WithContext(ctx context.Context) *GetaspecificRtcSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific rtc session params
func (o *GetaspecificRtcSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific rtc session params
func (o *GetaspecificRtcSessionParams) WithHTTPClient(client *http.Client) *GetaspecificRtcSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific rtc session params
func (o *GetaspecificRtcSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific rtc session params
func (o *GetaspecificRtcSessionParams) WithID(id string) *GetaspecificRtcSessionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific rtc session params
func (o *GetaspecificRtcSessionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificRtcSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
