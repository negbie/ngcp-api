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

// NewReplaceChangeaspecificCustomerLocationParams creates a new ReplaceChangeaspecificCustomerLocationParams object
// with the default values initialized.
func NewReplaceChangeaspecificCustomerLocationParams() *ReplaceChangeaspecificCustomerLocationParams {
	var ()
	return &ReplaceChangeaspecificCustomerLocationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceChangeaspecificCustomerLocationParamsWithTimeout creates a new ReplaceChangeaspecificCustomerLocationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceChangeaspecificCustomerLocationParamsWithTimeout(timeout time.Duration) *ReplaceChangeaspecificCustomerLocationParams {
	var ()
	return &ReplaceChangeaspecificCustomerLocationParams{

		timeout: timeout,
	}
}

// NewReplaceChangeaspecificCustomerLocationParamsWithContext creates a new ReplaceChangeaspecificCustomerLocationParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceChangeaspecificCustomerLocationParamsWithContext(ctx context.Context) *ReplaceChangeaspecificCustomerLocationParams {
	var ()
	return &ReplaceChangeaspecificCustomerLocationParams{

		Context: ctx,
	}
}

// NewReplaceChangeaspecificCustomerLocationParamsWithHTTPClient creates a new ReplaceChangeaspecificCustomerLocationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceChangeaspecificCustomerLocationParamsWithHTTPClient(client *http.Client) *ReplaceChangeaspecificCustomerLocationParams {
	var ()
	return &ReplaceChangeaspecificCustomerLocationParams{
		HTTPClient: client,
	}
}

/*ReplaceChangeaspecificCustomerLocationParams contains all the parameters to send to the API endpoint
for the replace changeaspecific customer location operation typically these are written to a http.Request
*/
type ReplaceChangeaspecificCustomerLocationParams struct {

	/*Body*/
	Body *Replace1changeaspecificCustomerLocationRequest
	/*ContentType*/
	ContentType string
	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace changeaspecific customer location params
func (o *ReplaceChangeaspecificCustomerLocationParams) WithTimeout(timeout time.Duration) *ReplaceChangeaspecificCustomerLocationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace changeaspecific customer location params
func (o *ReplaceChangeaspecificCustomerLocationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace changeaspecific customer location params
func (o *ReplaceChangeaspecificCustomerLocationParams) WithContext(ctx context.Context) *ReplaceChangeaspecificCustomerLocationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace changeaspecific customer location params
func (o *ReplaceChangeaspecificCustomerLocationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace changeaspecific customer location params
func (o *ReplaceChangeaspecificCustomerLocationParams) WithHTTPClient(client *http.Client) *ReplaceChangeaspecificCustomerLocationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace changeaspecific customer location params
func (o *ReplaceChangeaspecificCustomerLocationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace changeaspecific customer location params
func (o *ReplaceChangeaspecificCustomerLocationParams) WithBody(body *Replace1changeaspecificCustomerLocationRequest) *ReplaceChangeaspecificCustomerLocationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace changeaspecific customer location params
func (o *ReplaceChangeaspecificCustomerLocationParams) SetBody(body *Replace1changeaspecificCustomerLocationRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the replace changeaspecific customer location params
func (o *ReplaceChangeaspecificCustomerLocationParams) WithContentType(contentType string) *ReplaceChangeaspecificCustomerLocationParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the replace changeaspecific customer location params
func (o *ReplaceChangeaspecificCustomerLocationParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the replace changeaspecific customer location params
func (o *ReplaceChangeaspecificCustomerLocationParams) WithID(id string) *ReplaceChangeaspecificCustomerLocationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replace changeaspecific customer location params
func (o *ReplaceChangeaspecificCustomerLocationParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceChangeaspecificCustomerLocationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
