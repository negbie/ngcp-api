// Code generated by go-swagger; DO NOT EDIT.

package id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPhonebookentriesByIDPutParams creates a new PhonebookentriesByIDPutParams object
// with the default values initialized.
func NewPhonebookentriesByIDPutParams() *PhonebookentriesByIDPutParams {
	var ()
	return &PhonebookentriesByIDPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPhonebookentriesByIDPutParamsWithTimeout creates a new PhonebookentriesByIDPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPhonebookentriesByIDPutParamsWithTimeout(timeout time.Duration) *PhonebookentriesByIDPutParams {
	var ()
	return &PhonebookentriesByIDPutParams{

		timeout: timeout,
	}
}

// NewPhonebookentriesByIDPutParamsWithContext creates a new PhonebookentriesByIDPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewPhonebookentriesByIDPutParamsWithContext(ctx context.Context) *PhonebookentriesByIDPutParams {
	var ()
	return &PhonebookentriesByIDPutParams{

		Context: ctx,
	}
}

// NewPhonebookentriesByIDPutParamsWithHTTPClient creates a new PhonebookentriesByIDPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPhonebookentriesByIDPutParamsWithHTTPClient(client *http.Client) *PhonebookentriesByIDPutParams {
	var ()
	return &PhonebookentriesByIDPutParams{
		HTTPClient: client,
	}
}

/*PhonebookentriesByIDPutParams contains all the parameters to send to the API endpoint
for the phonebookentries by Id put operation typically these are written to a http.Request
*/
type PhonebookentriesByIDPutParams struct {

	/*Body*/
	Body *Replace1changeaspecificPhonebookEntryRequest
	/*ContentType*/
	ContentType string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the phonebookentries by Id put params
func (o *PhonebookentriesByIDPutParams) WithTimeout(timeout time.Duration) *PhonebookentriesByIDPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the phonebookentries by Id put params
func (o *PhonebookentriesByIDPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the phonebookentries by Id put params
func (o *PhonebookentriesByIDPutParams) WithContext(ctx context.Context) *PhonebookentriesByIDPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the phonebookentries by Id put params
func (o *PhonebookentriesByIDPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the phonebookentries by Id put params
func (o *PhonebookentriesByIDPutParams) WithHTTPClient(client *http.Client) *PhonebookentriesByIDPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the phonebookentries by Id put params
func (o *PhonebookentriesByIDPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the phonebookentries by Id put params
func (o *PhonebookentriesByIDPutParams) WithBody(body *Replace1changeaspecificPhonebookEntryRequest) *PhonebookentriesByIDPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the phonebookentries by Id put params
func (o *PhonebookentriesByIDPutParams) SetBody(body *Replace1changeaspecificPhonebookEntryRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the phonebookentries by Id put params
func (o *PhonebookentriesByIDPutParams) WithContentType(contentType string) *PhonebookentriesByIDPutParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the phonebookentries by Id put params
func (o *PhonebookentriesByIDPutParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the phonebookentries by Id put params
func (o *PhonebookentriesByIDPutParams) WithID(id string) *PhonebookentriesByIDPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the phonebookentries by Id put params
func (o *PhonebookentriesByIDPutParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PhonebookentriesByIDPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// header param Content-Type
	if err := r.SetHeaderParam("Content-Type", o.ContentType); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}