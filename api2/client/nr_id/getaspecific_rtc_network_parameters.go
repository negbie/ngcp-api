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

// NewGetaspecificRtcNetworkParams creates a new GetaspecificRtcNetworkParams object
// with the default values initialized.
func NewGetaspecificRtcNetworkParams() *GetaspecificRtcNetworkParams {
	var ()
	return &GetaspecificRtcNetworkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificRtcNetworkParamsWithTimeout creates a new GetaspecificRtcNetworkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificRtcNetworkParamsWithTimeout(timeout time.Duration) *GetaspecificRtcNetworkParams {
	var ()
	return &GetaspecificRtcNetworkParams{

		timeout: timeout,
	}
}

// NewGetaspecificRtcNetworkParamsWithContext creates a new GetaspecificRtcNetworkParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificRtcNetworkParamsWithContext(ctx context.Context) *GetaspecificRtcNetworkParams {
	var ()
	return &GetaspecificRtcNetworkParams{

		Context: ctx,
	}
}

// NewGetaspecificRtcNetworkParamsWithHTTPClient creates a new GetaspecificRtcNetworkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificRtcNetworkParamsWithHTTPClient(client *http.Client) *GetaspecificRtcNetworkParams {
	var ()
	return &GetaspecificRtcNetworkParams{
		HTTPClient: client,
	}
}

/*GetaspecificRtcNetworkParams contains all the parameters to send to the API endpoint
for the getaspecific rtc network operation typically these are written to a http.Request
*/
type GetaspecificRtcNetworkParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific rtc network params
func (o *GetaspecificRtcNetworkParams) WithTimeout(timeout time.Duration) *GetaspecificRtcNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific rtc network params
func (o *GetaspecificRtcNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific rtc network params
func (o *GetaspecificRtcNetworkParams) WithContext(ctx context.Context) *GetaspecificRtcNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific rtc network params
func (o *GetaspecificRtcNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific rtc network params
func (o *GetaspecificRtcNetworkParams) WithHTTPClient(client *http.Client) *GetaspecificRtcNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific rtc network params
func (o *GetaspecificRtcNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific rtc network params
func (o *GetaspecificRtcNetworkParams) WithID(id string) *GetaspecificRtcNetworkParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific rtc network params
func (o *GetaspecificRtcNetworkParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificRtcNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
