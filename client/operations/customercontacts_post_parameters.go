// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	models "github.com/negbie/ngcp-api/models"
)

// NewCustomercontactsPostParams creates a new CustomercontactsPostParams object
// with the default values initialized.
func NewCustomercontactsPostParams() *CustomercontactsPostParams {
	var ()
	return &CustomercontactsPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomercontactsPostParamsWithTimeout creates a new CustomercontactsPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomercontactsPostParamsWithTimeout(timeout time.Duration) *CustomercontactsPostParams {
	var ()
	return &CustomercontactsPostParams{

		timeout: timeout,
	}
}

// NewCustomercontactsPostParamsWithContext creates a new CustomercontactsPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomercontactsPostParamsWithContext(ctx context.Context) *CustomercontactsPostParams {
	var ()
	return &CustomercontactsPostParams{

		Context: ctx,
	}
}

// NewCustomercontactsPostParamsWithHTTPClient creates a new CustomercontactsPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomercontactsPostParamsWithHTTPClient(client *http.Client) *CustomercontactsPostParams {
	var ()
	return &CustomercontactsPostParams{
		HTTPClient: client,
	}
}

/*CustomercontactsPostParams contains all the parameters to send to the API endpoint
for the customercontacts post operation typically these are written to a http.Request
*/
type CustomercontactsPostParams struct {

	/*Body*/
	Body *models.CreateanewCustomerContactRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the customercontacts post params
func (o *CustomercontactsPostParams) WithTimeout(timeout time.Duration) *CustomercontactsPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customercontacts post params
func (o *CustomercontactsPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customercontacts post params
func (o *CustomercontactsPostParams) WithContext(ctx context.Context) *CustomercontactsPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customercontacts post params
func (o *CustomercontactsPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customercontacts post params
func (o *CustomercontactsPostParams) WithHTTPClient(client *http.Client) *CustomercontactsPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customercontacts post params
func (o *CustomercontactsPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the customercontacts post params
func (o *CustomercontactsPostParams) WithBody(body *models.CreateanewCustomerContactRequest) *CustomercontactsPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the customercontacts post params
func (o *CustomercontactsPostParams) SetBody(body *models.CreateanewCustomerContactRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the customercontacts post params
func (o *CustomercontactsPostParams) WithContentType(contentType string) *CustomercontactsPostParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the customercontacts post params
func (o *CustomercontactsPostParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *CustomercontactsPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}