// Code generated by go-swagger; DO NOT EDIT.

package contracts

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

// NewCreateanewContractParams creates a new CreateanewContractParams object
// with the default values initialized.
func NewCreateanewContractParams() *CreateanewContractParams {
	var ()
	return &CreateanewContractParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateanewContractParamsWithTimeout creates a new CreateanewContractParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateanewContractParamsWithTimeout(timeout time.Duration) *CreateanewContractParams {
	var ()
	return &CreateanewContractParams{

		timeout: timeout,
	}
}

// NewCreateanewContractParamsWithContext creates a new CreateanewContractParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateanewContractParamsWithContext(ctx context.Context) *CreateanewContractParams {
	var ()
	return &CreateanewContractParams{

		Context: ctx,
	}
}

// NewCreateanewContractParamsWithHTTPClient creates a new CreateanewContractParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateanewContractParamsWithHTTPClient(client *http.Client) *CreateanewContractParams {
	var ()
	return &CreateanewContractParams{
		HTTPClient: client,
	}
}

/*CreateanewContractParams contains all the parameters to send to the API endpoint
for the createanew contract operation typically these are written to a http.Request
*/
type CreateanewContractParams struct {

	/*Body*/
	Body *models.CreateanewContractRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the createanew contract params
func (o *CreateanewContractParams) WithTimeout(timeout time.Duration) *CreateanewContractParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the createanew contract params
func (o *CreateanewContractParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the createanew contract params
func (o *CreateanewContractParams) WithContext(ctx context.Context) *CreateanewContractParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the createanew contract params
func (o *CreateanewContractParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the createanew contract params
func (o *CreateanewContractParams) WithHTTPClient(client *http.Client) *CreateanewContractParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the createanew contract params
func (o *CreateanewContractParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the createanew contract params
func (o *CreateanewContractParams) WithBody(body *models.CreateanewContractRequest) *CreateanewContractParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the createanew contract params
func (o *CreateanewContractParams) SetBody(body *models.CreateanewContractRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the createanew contract params
func (o *CreateanewContractParams) WithContentType(contentType string) *CreateanewContractParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the createanew contract params
func (o *CreateanewContractParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateanewContractParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
