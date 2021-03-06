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

// NewGetaspecificAutoAttendantParams creates a new GetaspecificAutoAttendantParams object
// with the default values initialized.
func NewGetaspecificAutoAttendantParams() *GetaspecificAutoAttendantParams {
	var ()
	return &GetaspecificAutoAttendantParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificAutoAttendantParamsWithTimeout creates a new GetaspecificAutoAttendantParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificAutoAttendantParamsWithTimeout(timeout time.Duration) *GetaspecificAutoAttendantParams {
	var ()
	return &GetaspecificAutoAttendantParams{

		timeout: timeout,
	}
}

// NewGetaspecificAutoAttendantParamsWithContext creates a new GetaspecificAutoAttendantParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificAutoAttendantParamsWithContext(ctx context.Context) *GetaspecificAutoAttendantParams {
	var ()
	return &GetaspecificAutoAttendantParams{

		Context: ctx,
	}
}

// NewGetaspecificAutoAttendantParamsWithHTTPClient creates a new GetaspecificAutoAttendantParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificAutoAttendantParamsWithHTTPClient(client *http.Client) *GetaspecificAutoAttendantParams {
	var ()
	return &GetaspecificAutoAttendantParams{
		HTTPClient: client,
	}
}

/*GetaspecificAutoAttendantParams contains all the parameters to send to the API endpoint
for the getaspecific auto attendant operation typically these are written to a http.Request
*/
type GetaspecificAutoAttendantParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific auto attendant params
func (o *GetaspecificAutoAttendantParams) WithTimeout(timeout time.Duration) *GetaspecificAutoAttendantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific auto attendant params
func (o *GetaspecificAutoAttendantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific auto attendant params
func (o *GetaspecificAutoAttendantParams) WithContext(ctx context.Context) *GetaspecificAutoAttendantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific auto attendant params
func (o *GetaspecificAutoAttendantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific auto attendant params
func (o *GetaspecificAutoAttendantParams) WithHTTPClient(client *http.Client) *GetaspecificAutoAttendantParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific auto attendant params
func (o *GetaspecificAutoAttendantParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific auto attendant params
func (o *GetaspecificAutoAttendantParams) WithID(id string) *GetaspecificAutoAttendantParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific auto attendant params
func (o *GetaspecificAutoAttendantParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificAutoAttendantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
