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

// NewGetaspecificSubscriberRegistrationParams creates a new GetaspecificSubscriberRegistrationParams object
// with the default values initialized.
func NewGetaspecificSubscriberRegistrationParams() *GetaspecificSubscriberRegistrationParams {
	var ()
	return &GetaspecificSubscriberRegistrationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificSubscriberRegistrationParamsWithTimeout creates a new GetaspecificSubscriberRegistrationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificSubscriberRegistrationParamsWithTimeout(timeout time.Duration) *GetaspecificSubscriberRegistrationParams {
	var ()
	return &GetaspecificSubscriberRegistrationParams{

		timeout: timeout,
	}
}

// NewGetaspecificSubscriberRegistrationParamsWithContext creates a new GetaspecificSubscriberRegistrationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificSubscriberRegistrationParamsWithContext(ctx context.Context) *GetaspecificSubscriberRegistrationParams {
	var ()
	return &GetaspecificSubscriberRegistrationParams{

		Context: ctx,
	}
}

// NewGetaspecificSubscriberRegistrationParamsWithHTTPClient creates a new GetaspecificSubscriberRegistrationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificSubscriberRegistrationParamsWithHTTPClient(client *http.Client) *GetaspecificSubscriberRegistrationParams {
	var ()
	return &GetaspecificSubscriberRegistrationParams{
		HTTPClient: client,
	}
}

/*GetaspecificSubscriberRegistrationParams contains all the parameters to send to the API endpoint
for the getaspecific subscriber registration operation typically these are written to a http.Request
*/
type GetaspecificSubscriberRegistrationParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific subscriber registration params
func (o *GetaspecificSubscriberRegistrationParams) WithTimeout(timeout time.Duration) *GetaspecificSubscriberRegistrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific subscriber registration params
func (o *GetaspecificSubscriberRegistrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific subscriber registration params
func (o *GetaspecificSubscriberRegistrationParams) WithContext(ctx context.Context) *GetaspecificSubscriberRegistrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific subscriber registration params
func (o *GetaspecificSubscriberRegistrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific subscriber registration params
func (o *GetaspecificSubscriberRegistrationParams) WithHTTPClient(client *http.Client) *GetaspecificSubscriberRegistrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific subscriber registration params
func (o *GetaspecificSubscriberRegistrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific subscriber registration params
func (o *GetaspecificSubscriberRegistrationParams) WithID(id string) *GetaspecificSubscriberRegistrationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific subscriber registration params
func (o *GetaspecificSubscriberRegistrationParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificSubscriberRegistrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
