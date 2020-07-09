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

// NewDeleteaspecificEmergencyMappingParams creates a new DeleteaspecificEmergencyMappingParams object
// with the default values initialized.
func NewDeleteaspecificEmergencyMappingParams() *DeleteaspecificEmergencyMappingParams {
	var ()
	return &DeleteaspecificEmergencyMappingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteaspecificEmergencyMappingParamsWithTimeout creates a new DeleteaspecificEmergencyMappingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteaspecificEmergencyMappingParamsWithTimeout(timeout time.Duration) *DeleteaspecificEmergencyMappingParams {
	var ()
	return &DeleteaspecificEmergencyMappingParams{

		timeout: timeout,
	}
}

// NewDeleteaspecificEmergencyMappingParamsWithContext creates a new DeleteaspecificEmergencyMappingParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteaspecificEmergencyMappingParamsWithContext(ctx context.Context) *DeleteaspecificEmergencyMappingParams {
	var ()
	return &DeleteaspecificEmergencyMappingParams{

		Context: ctx,
	}
}

// NewDeleteaspecificEmergencyMappingParamsWithHTTPClient creates a new DeleteaspecificEmergencyMappingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteaspecificEmergencyMappingParamsWithHTTPClient(client *http.Client) *DeleteaspecificEmergencyMappingParams {
	var ()
	return &DeleteaspecificEmergencyMappingParams{
		HTTPClient: client,
	}
}

/*DeleteaspecificEmergencyMappingParams contains all the parameters to send to the API endpoint
for the deleteaspecific emergency mapping operation typically these are written to a http.Request
*/
type DeleteaspecificEmergencyMappingParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the deleteaspecific emergency mapping params
func (o *DeleteaspecificEmergencyMappingParams) WithTimeout(timeout time.Duration) *DeleteaspecificEmergencyMappingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deleteaspecific emergency mapping params
func (o *DeleteaspecificEmergencyMappingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deleteaspecific emergency mapping params
func (o *DeleteaspecificEmergencyMappingParams) WithContext(ctx context.Context) *DeleteaspecificEmergencyMappingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deleteaspecific emergency mapping params
func (o *DeleteaspecificEmergencyMappingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deleteaspecific emergency mapping params
func (o *DeleteaspecificEmergencyMappingParams) WithHTTPClient(client *http.Client) *DeleteaspecificEmergencyMappingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deleteaspecific emergency mapping params
func (o *DeleteaspecificEmergencyMappingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the deleteaspecific emergency mapping params
func (o *DeleteaspecificEmergencyMappingParams) WithID(id string) *DeleteaspecificEmergencyMappingParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the deleteaspecific emergency mapping params
func (o *DeleteaspecificEmergencyMappingParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteaspecificEmergencyMappingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
