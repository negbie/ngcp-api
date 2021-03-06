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

// NewDeleteaspecificCFDestinationSetParams creates a new DeleteaspecificCFDestinationSetParams object
// with the default values initialized.
func NewDeleteaspecificCFDestinationSetParams() *DeleteaspecificCFDestinationSetParams {
	var ()
	return &DeleteaspecificCFDestinationSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteaspecificCFDestinationSetParamsWithTimeout creates a new DeleteaspecificCFDestinationSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteaspecificCFDestinationSetParamsWithTimeout(timeout time.Duration) *DeleteaspecificCFDestinationSetParams {
	var ()
	return &DeleteaspecificCFDestinationSetParams{

		timeout: timeout,
	}
}

// NewDeleteaspecificCFDestinationSetParamsWithContext creates a new DeleteaspecificCFDestinationSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteaspecificCFDestinationSetParamsWithContext(ctx context.Context) *DeleteaspecificCFDestinationSetParams {
	var ()
	return &DeleteaspecificCFDestinationSetParams{

		Context: ctx,
	}
}

// NewDeleteaspecificCFDestinationSetParamsWithHTTPClient creates a new DeleteaspecificCFDestinationSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteaspecificCFDestinationSetParamsWithHTTPClient(client *http.Client) *DeleteaspecificCFDestinationSetParams {
	var ()
	return &DeleteaspecificCFDestinationSetParams{
		HTTPClient: client,
	}
}

/*DeleteaspecificCFDestinationSetParams contains all the parameters to send to the API endpoint
for the deleteaspecific c f destination set operation typically these are written to a http.Request
*/
type DeleteaspecificCFDestinationSetParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the deleteaspecific c f destination set params
func (o *DeleteaspecificCFDestinationSetParams) WithTimeout(timeout time.Duration) *DeleteaspecificCFDestinationSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deleteaspecific c f destination set params
func (o *DeleteaspecificCFDestinationSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deleteaspecific c f destination set params
func (o *DeleteaspecificCFDestinationSetParams) WithContext(ctx context.Context) *DeleteaspecificCFDestinationSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deleteaspecific c f destination set params
func (o *DeleteaspecificCFDestinationSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deleteaspecific c f destination set params
func (o *DeleteaspecificCFDestinationSetParams) WithHTTPClient(client *http.Client) *DeleteaspecificCFDestinationSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deleteaspecific c f destination set params
func (o *DeleteaspecificCFDestinationSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the deleteaspecific c f destination set params
func (o *DeleteaspecificCFDestinationSetParams) WithID(id string) *DeleteaspecificCFDestinationSetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the deleteaspecific c f destination set params
func (o *DeleteaspecificCFDestinationSetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteaspecificCFDestinationSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
