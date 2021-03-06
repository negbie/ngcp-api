// Code generated by go-swagger; DO NOT EDIT.

package faxes

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

// NewGetaspecificFaxParams creates a new GetaspecificFaxParams object
// with the default values initialized.
func NewGetaspecificFaxParams() *GetaspecificFaxParams {
	var ()
	return &GetaspecificFaxParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificFaxParamsWithTimeout creates a new GetaspecificFaxParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificFaxParamsWithTimeout(timeout time.Duration) *GetaspecificFaxParams {
	var ()
	return &GetaspecificFaxParams{

		timeout: timeout,
	}
}

// NewGetaspecificFaxParamsWithContext creates a new GetaspecificFaxParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificFaxParamsWithContext(ctx context.Context) *GetaspecificFaxParams {
	var ()
	return &GetaspecificFaxParams{

		Context: ctx,
	}
}

// NewGetaspecificFaxParamsWithHTTPClient creates a new GetaspecificFaxParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificFaxParamsWithHTTPClient(client *http.Client) *GetaspecificFaxParams {
	var ()
	return &GetaspecificFaxParams{
		HTTPClient: client,
	}
}

/*GetaspecificFaxParams contains all the parameters to send to the API endpoint
for the getaspecific fax operation typically these are written to a http.Request
*/
type GetaspecificFaxParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific fax params
func (o *GetaspecificFaxParams) WithTimeout(timeout time.Duration) *GetaspecificFaxParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific fax params
func (o *GetaspecificFaxParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific fax params
func (o *GetaspecificFaxParams) WithContext(ctx context.Context) *GetaspecificFaxParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific fax params
func (o *GetaspecificFaxParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific fax params
func (o *GetaspecificFaxParams) WithHTTPClient(client *http.Client) *GetaspecificFaxParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific fax params
func (o *GetaspecificFaxParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific fax params
func (o *GetaspecificFaxParams) WithID(id string) *GetaspecificFaxParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific fax params
func (o *GetaspecificFaxParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificFaxParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
