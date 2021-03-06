// Code generated by go-swagger; DO NOT EDIT.

package customercontacts

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

	"github.com/negbie/ngcp-api/models"
)

// NewCreateanewCustomerContactParams creates a new CreateanewCustomerContactParams object
// with the default values initialized.
func NewCreateanewCustomerContactParams() *CreateanewCustomerContactParams {
	var ()
	return &CreateanewCustomerContactParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateanewCustomerContactParamsWithTimeout creates a new CreateanewCustomerContactParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateanewCustomerContactParamsWithTimeout(timeout time.Duration) *CreateanewCustomerContactParams {
	var ()
	return &CreateanewCustomerContactParams{

		timeout: timeout,
	}
}

// NewCreateanewCustomerContactParamsWithContext creates a new CreateanewCustomerContactParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateanewCustomerContactParamsWithContext(ctx context.Context) *CreateanewCustomerContactParams {
	var ()
	return &CreateanewCustomerContactParams{

		Context: ctx,
	}
}

// NewCreateanewCustomerContactParamsWithHTTPClient creates a new CreateanewCustomerContactParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateanewCustomerContactParamsWithHTTPClient(client *http.Client) *CreateanewCustomerContactParams {
	var ()
	return &CreateanewCustomerContactParams{
		HTTPClient: client,
	}
}

/*CreateanewCustomerContactParams contains all the parameters to send to the API endpoint
for the createanew customer contact operation typically these are written to a http.Request
*/
type CreateanewCustomerContactParams struct {

	/*Body*/
	Body *models.CreateanewCustomerContactRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the createanew customer contact params
func (o *CreateanewCustomerContactParams) WithTimeout(timeout time.Duration) *CreateanewCustomerContactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the createanew customer contact params
func (o *CreateanewCustomerContactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the createanew customer contact params
func (o *CreateanewCustomerContactParams) WithContext(ctx context.Context) *CreateanewCustomerContactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the createanew customer contact params
func (o *CreateanewCustomerContactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the createanew customer contact params
func (o *CreateanewCustomerContactParams) WithHTTPClient(client *http.Client) *CreateanewCustomerContactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the createanew customer contact params
func (o *CreateanewCustomerContactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the createanew customer contact params
func (o *CreateanewCustomerContactParams) WithBody(body *models.CreateanewCustomerContactRequest) *CreateanewCustomerContactParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the createanew customer contact params
func (o *CreateanewCustomerContactParams) SetBody(body *models.CreateanewCustomerContactRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the createanew customer contact params
func (o *CreateanewCustomerContactParams) WithContentType(contentType string) *CreateanewCustomerContactParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the createanew customer contact params
func (o *CreateanewCustomerContactParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateanewCustomerContactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
