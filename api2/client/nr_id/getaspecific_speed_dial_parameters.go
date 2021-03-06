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

// NewGetaspecificSpeedDialParams creates a new GetaspecificSpeedDialParams object
// with the default values initialized.
func NewGetaspecificSpeedDialParams() *GetaspecificSpeedDialParams {
	var ()
	return &GetaspecificSpeedDialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificSpeedDialParamsWithTimeout creates a new GetaspecificSpeedDialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificSpeedDialParamsWithTimeout(timeout time.Duration) *GetaspecificSpeedDialParams {
	var ()
	return &GetaspecificSpeedDialParams{

		timeout: timeout,
	}
}

// NewGetaspecificSpeedDialParamsWithContext creates a new GetaspecificSpeedDialParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificSpeedDialParamsWithContext(ctx context.Context) *GetaspecificSpeedDialParams {
	var ()
	return &GetaspecificSpeedDialParams{

		Context: ctx,
	}
}

// NewGetaspecificSpeedDialParamsWithHTTPClient creates a new GetaspecificSpeedDialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificSpeedDialParamsWithHTTPClient(client *http.Client) *GetaspecificSpeedDialParams {
	var ()
	return &GetaspecificSpeedDialParams{
		HTTPClient: client,
	}
}

/*GetaspecificSpeedDialParams contains all the parameters to send to the API endpoint
for the getaspecific speed dial operation typically these are written to a http.Request
*/
type GetaspecificSpeedDialParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific speed dial params
func (o *GetaspecificSpeedDialParams) WithTimeout(timeout time.Duration) *GetaspecificSpeedDialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific speed dial params
func (o *GetaspecificSpeedDialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific speed dial params
func (o *GetaspecificSpeedDialParams) WithContext(ctx context.Context) *GetaspecificSpeedDialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific speed dial params
func (o *GetaspecificSpeedDialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific speed dial params
func (o *GetaspecificSpeedDialParams) WithHTTPClient(client *http.Client) *GetaspecificSpeedDialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific speed dial params
func (o *GetaspecificSpeedDialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific speed dial params
func (o *GetaspecificSpeedDialParams) WithID(id string) *GetaspecificSpeedDialParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific speed dial params
func (o *GetaspecificSpeedDialParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificSpeedDialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
