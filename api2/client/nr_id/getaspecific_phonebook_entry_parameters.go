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

// NewGetaspecificPhonebookEntryParams creates a new GetaspecificPhonebookEntryParams object
// with the default values initialized.
func NewGetaspecificPhonebookEntryParams() *GetaspecificPhonebookEntryParams {
	var ()
	return &GetaspecificPhonebookEntryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificPhonebookEntryParamsWithTimeout creates a new GetaspecificPhonebookEntryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificPhonebookEntryParamsWithTimeout(timeout time.Duration) *GetaspecificPhonebookEntryParams {
	var ()
	return &GetaspecificPhonebookEntryParams{

		timeout: timeout,
	}
}

// NewGetaspecificPhonebookEntryParamsWithContext creates a new GetaspecificPhonebookEntryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificPhonebookEntryParamsWithContext(ctx context.Context) *GetaspecificPhonebookEntryParams {
	var ()
	return &GetaspecificPhonebookEntryParams{

		Context: ctx,
	}
}

// NewGetaspecificPhonebookEntryParamsWithHTTPClient creates a new GetaspecificPhonebookEntryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificPhonebookEntryParamsWithHTTPClient(client *http.Client) *GetaspecificPhonebookEntryParams {
	var ()
	return &GetaspecificPhonebookEntryParams{
		HTTPClient: client,
	}
}

/*GetaspecificPhonebookEntryParams contains all the parameters to send to the API endpoint
for the getaspecific phonebook entry operation typically these are written to a http.Request
*/
type GetaspecificPhonebookEntryParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific phonebook entry params
func (o *GetaspecificPhonebookEntryParams) WithTimeout(timeout time.Duration) *GetaspecificPhonebookEntryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific phonebook entry params
func (o *GetaspecificPhonebookEntryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific phonebook entry params
func (o *GetaspecificPhonebookEntryParams) WithContext(ctx context.Context) *GetaspecificPhonebookEntryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific phonebook entry params
func (o *GetaspecificPhonebookEntryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific phonebook entry params
func (o *GetaspecificPhonebookEntryParams) WithHTTPClient(client *http.Client) *GetaspecificPhonebookEntryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific phonebook entry params
func (o *GetaspecificPhonebookEntryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific phonebook entry params
func (o *GetaspecificPhonebookEntryParams) WithID(id string) *GetaspecificPhonebookEntryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific phonebook entry params
func (o *GetaspecificPhonebookEntryParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificPhonebookEntryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
