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

// NewPeeringrulesPostParams creates a new PeeringrulesPostParams object
// with the default values initialized.
func NewPeeringrulesPostParams() *PeeringrulesPostParams {
	var ()
	return &PeeringrulesPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPeeringrulesPostParamsWithTimeout creates a new PeeringrulesPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPeeringrulesPostParamsWithTimeout(timeout time.Duration) *PeeringrulesPostParams {
	var ()
	return &PeeringrulesPostParams{

		timeout: timeout,
	}
}

// NewPeeringrulesPostParamsWithContext creates a new PeeringrulesPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewPeeringrulesPostParamsWithContext(ctx context.Context) *PeeringrulesPostParams {
	var ()
	return &PeeringrulesPostParams{

		Context: ctx,
	}
}

// NewPeeringrulesPostParamsWithHTTPClient creates a new PeeringrulesPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPeeringrulesPostParamsWithHTTPClient(client *http.Client) *PeeringrulesPostParams {
	var ()
	return &PeeringrulesPostParams{
		HTTPClient: client,
	}
}

/*PeeringrulesPostParams contains all the parameters to send to the API endpoint
for the peeringrules post operation typically these are written to a http.Request
*/
type PeeringrulesPostParams struct {

	/*Body*/
	Body *models.CreateanewPeeringRuleRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the peeringrules post params
func (o *PeeringrulesPostParams) WithTimeout(timeout time.Duration) *PeeringrulesPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the peeringrules post params
func (o *PeeringrulesPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the peeringrules post params
func (o *PeeringrulesPostParams) WithContext(ctx context.Context) *PeeringrulesPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the peeringrules post params
func (o *PeeringrulesPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the peeringrules post params
func (o *PeeringrulesPostParams) WithHTTPClient(client *http.Client) *PeeringrulesPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the peeringrules post params
func (o *PeeringrulesPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the peeringrules post params
func (o *PeeringrulesPostParams) WithBody(body *models.CreateanewPeeringRuleRequest) *PeeringrulesPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the peeringrules post params
func (o *PeeringrulesPostParams) SetBody(body *models.CreateanewPeeringRuleRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the peeringrules post params
func (o *PeeringrulesPostParams) WithContentType(contentType string) *PeeringrulesPostParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the peeringrules post params
func (o *PeeringrulesPostParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *PeeringrulesPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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