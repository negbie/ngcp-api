// Code generated by go-swagger; DO NOT EDIT.

package subscriberprofilesets

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

// NewCreateanewSubscriberProfileSetParams creates a new CreateanewSubscriberProfileSetParams object
// with the default values initialized.
func NewCreateanewSubscriberProfileSetParams() *CreateanewSubscriberProfileSetParams {
	var ()
	return &CreateanewSubscriberProfileSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateanewSubscriberProfileSetParamsWithTimeout creates a new CreateanewSubscriberProfileSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateanewSubscriberProfileSetParamsWithTimeout(timeout time.Duration) *CreateanewSubscriberProfileSetParams {
	var ()
	return &CreateanewSubscriberProfileSetParams{

		timeout: timeout,
	}
}

// NewCreateanewSubscriberProfileSetParamsWithContext creates a new CreateanewSubscriberProfileSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateanewSubscriberProfileSetParamsWithContext(ctx context.Context) *CreateanewSubscriberProfileSetParams {
	var ()
	return &CreateanewSubscriberProfileSetParams{

		Context: ctx,
	}
}

// NewCreateanewSubscriberProfileSetParamsWithHTTPClient creates a new CreateanewSubscriberProfileSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateanewSubscriberProfileSetParamsWithHTTPClient(client *http.Client) *CreateanewSubscriberProfileSetParams {
	var ()
	return &CreateanewSubscriberProfileSetParams{
		HTTPClient: client,
	}
}

/*CreateanewSubscriberProfileSetParams contains all the parameters to send to the API endpoint
for the createanew subscriber profile set operation typically these are written to a http.Request
*/
type CreateanewSubscriberProfileSetParams struct {

	/*Body*/
	Body *models.CreateanewSubscriberProfileSetRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the createanew subscriber profile set params
func (o *CreateanewSubscriberProfileSetParams) WithTimeout(timeout time.Duration) *CreateanewSubscriberProfileSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the createanew subscriber profile set params
func (o *CreateanewSubscriberProfileSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the createanew subscriber profile set params
func (o *CreateanewSubscriberProfileSetParams) WithContext(ctx context.Context) *CreateanewSubscriberProfileSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the createanew subscriber profile set params
func (o *CreateanewSubscriberProfileSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the createanew subscriber profile set params
func (o *CreateanewSubscriberProfileSetParams) WithHTTPClient(client *http.Client) *CreateanewSubscriberProfileSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the createanew subscriber profile set params
func (o *CreateanewSubscriberProfileSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the createanew subscriber profile set params
func (o *CreateanewSubscriberProfileSetParams) WithBody(body *models.CreateanewSubscriberProfileSetRequest) *CreateanewSubscriberProfileSetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the createanew subscriber profile set params
func (o *CreateanewSubscriberProfileSetParams) SetBody(body *models.CreateanewSubscriberProfileSetRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the createanew subscriber profile set params
func (o *CreateanewSubscriberProfileSetParams) WithContentType(contentType string) *CreateanewSubscriberProfileSetParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the createanew subscriber profile set params
func (o *CreateanewSubscriberProfileSetParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateanewSubscriberProfileSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
