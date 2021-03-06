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

// NewDeleteaspecificPbxDeviceParams creates a new DeleteaspecificPbxDeviceParams object
// with the default values initialized.
func NewDeleteaspecificPbxDeviceParams() *DeleteaspecificPbxDeviceParams {
	var ()
	return &DeleteaspecificPbxDeviceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteaspecificPbxDeviceParamsWithTimeout creates a new DeleteaspecificPbxDeviceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteaspecificPbxDeviceParamsWithTimeout(timeout time.Duration) *DeleteaspecificPbxDeviceParams {
	var ()
	return &DeleteaspecificPbxDeviceParams{

		timeout: timeout,
	}
}

// NewDeleteaspecificPbxDeviceParamsWithContext creates a new DeleteaspecificPbxDeviceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteaspecificPbxDeviceParamsWithContext(ctx context.Context) *DeleteaspecificPbxDeviceParams {
	var ()
	return &DeleteaspecificPbxDeviceParams{

		Context: ctx,
	}
}

// NewDeleteaspecificPbxDeviceParamsWithHTTPClient creates a new DeleteaspecificPbxDeviceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteaspecificPbxDeviceParamsWithHTTPClient(client *http.Client) *DeleteaspecificPbxDeviceParams {
	var ()
	return &DeleteaspecificPbxDeviceParams{
		HTTPClient: client,
	}
}

/*DeleteaspecificPbxDeviceParams contains all the parameters to send to the API endpoint
for the deleteaspecific pbx device operation typically these are written to a http.Request
*/
type DeleteaspecificPbxDeviceParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the deleteaspecific pbx device params
func (o *DeleteaspecificPbxDeviceParams) WithTimeout(timeout time.Duration) *DeleteaspecificPbxDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deleteaspecific pbx device params
func (o *DeleteaspecificPbxDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deleteaspecific pbx device params
func (o *DeleteaspecificPbxDeviceParams) WithContext(ctx context.Context) *DeleteaspecificPbxDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deleteaspecific pbx device params
func (o *DeleteaspecificPbxDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deleteaspecific pbx device params
func (o *DeleteaspecificPbxDeviceParams) WithHTTPClient(client *http.Client) *DeleteaspecificPbxDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deleteaspecific pbx device params
func (o *DeleteaspecificPbxDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the deleteaspecific pbx device params
func (o *DeleteaspecificPbxDeviceParams) WithID(id string) *DeleteaspecificPbxDeviceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the deleteaspecific pbx device params
func (o *DeleteaspecificPbxDeviceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteaspecificPbxDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
