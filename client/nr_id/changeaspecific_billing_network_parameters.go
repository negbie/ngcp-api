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

// NewChangeaspecificBillingNetworkParams creates a new ChangeaspecificBillingNetworkParams object
// with the default values initialized.
func NewChangeaspecificBillingNetworkParams() *ChangeaspecificBillingNetworkParams {
	var ()
	return &ChangeaspecificBillingNetworkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeaspecificBillingNetworkParamsWithTimeout creates a new ChangeaspecificBillingNetworkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeaspecificBillingNetworkParamsWithTimeout(timeout time.Duration) *ChangeaspecificBillingNetworkParams {
	var ()
	return &ChangeaspecificBillingNetworkParams{

		timeout: timeout,
	}
}

// NewChangeaspecificBillingNetworkParamsWithContext creates a new ChangeaspecificBillingNetworkParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeaspecificBillingNetworkParamsWithContext(ctx context.Context) *ChangeaspecificBillingNetworkParams {
	var ()
	return &ChangeaspecificBillingNetworkParams{

		Context: ctx,
	}
}

// NewChangeaspecificBillingNetworkParamsWithHTTPClient creates a new ChangeaspecificBillingNetworkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeaspecificBillingNetworkParamsWithHTTPClient(client *http.Client) *ChangeaspecificBillingNetworkParams {
	var ()
	return &ChangeaspecificBillingNetworkParams{
		HTTPClient: client,
	}
}

/*ChangeaspecificBillingNetworkParams contains all the parameters to send to the API endpoint
for the changeaspecific billing network operation typically these are written to a http.Request
*/
type ChangeaspecificBillingNetworkParams struct {

	/*Body*/
	Body string
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

// WithTimeout adds the timeout to the changeaspecific billing network params
func (o *ChangeaspecificBillingNetworkParams) WithTimeout(timeout time.Duration) *ChangeaspecificBillingNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the changeaspecific billing network params
func (o *ChangeaspecificBillingNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the changeaspecific billing network params
func (o *ChangeaspecificBillingNetworkParams) WithContext(ctx context.Context) *ChangeaspecificBillingNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the changeaspecific billing network params
func (o *ChangeaspecificBillingNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the changeaspecific billing network params
func (o *ChangeaspecificBillingNetworkParams) WithHTTPClient(client *http.Client) *ChangeaspecificBillingNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the changeaspecific billing network params
func (o *ChangeaspecificBillingNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the changeaspecific billing network params
func (o *ChangeaspecificBillingNetworkParams) WithBody(body string) *ChangeaspecificBillingNetworkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the changeaspecific billing network params
func (o *ChangeaspecificBillingNetworkParams) SetBody(body string) {
	o.Body = body
}

// WithContentType adds the contentType to the changeaspecific billing network params
func (o *ChangeaspecificBillingNetworkParams) WithContentType(contentType string) *ChangeaspecificBillingNetworkParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the changeaspecific billing network params
func (o *ChangeaspecificBillingNetworkParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithID adds the id to the changeaspecific billing network params
func (o *ChangeaspecificBillingNetworkParams) WithID(id string) *ChangeaspecificBillingNetworkParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the changeaspecific billing network params
func (o *ChangeaspecificBillingNetworkParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeaspecificBillingNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
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
