// Code generated by go-swagger; DO NOT EDIT.

package peeringrules

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

	"github.com/negbie/ngcp-api1/models"
)

// NewCreateanewPeeringRuleParams creates a new CreateanewPeeringRuleParams object
// with the default values initialized.
func NewCreateanewPeeringRuleParams() *CreateanewPeeringRuleParams {
	var ()
	return &CreateanewPeeringRuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateanewPeeringRuleParamsWithTimeout creates a new CreateanewPeeringRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateanewPeeringRuleParamsWithTimeout(timeout time.Duration) *CreateanewPeeringRuleParams {
	var ()
	return &CreateanewPeeringRuleParams{

		timeout: timeout,
	}
}

// NewCreateanewPeeringRuleParamsWithContext creates a new CreateanewPeeringRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateanewPeeringRuleParamsWithContext(ctx context.Context) *CreateanewPeeringRuleParams {
	var ()
	return &CreateanewPeeringRuleParams{

		Context: ctx,
	}
}

// NewCreateanewPeeringRuleParamsWithHTTPClient creates a new CreateanewPeeringRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateanewPeeringRuleParamsWithHTTPClient(client *http.Client) *CreateanewPeeringRuleParams {
	var ()
	return &CreateanewPeeringRuleParams{
		HTTPClient: client,
	}
}

/*CreateanewPeeringRuleParams contains all the parameters to send to the API endpoint
for the createanew peering rule operation typically these are written to a http.Request
*/
type CreateanewPeeringRuleParams struct {

	/*Body*/
	Body *models.CreateanewPeeringRuleRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the createanew peering rule params
func (o *CreateanewPeeringRuleParams) WithTimeout(timeout time.Duration) *CreateanewPeeringRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the createanew peering rule params
func (o *CreateanewPeeringRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the createanew peering rule params
func (o *CreateanewPeeringRuleParams) WithContext(ctx context.Context) *CreateanewPeeringRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the createanew peering rule params
func (o *CreateanewPeeringRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the createanew peering rule params
func (o *CreateanewPeeringRuleParams) WithHTTPClient(client *http.Client) *CreateanewPeeringRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the createanew peering rule params
func (o *CreateanewPeeringRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the createanew peering rule params
func (o *CreateanewPeeringRuleParams) WithBody(body *models.CreateanewPeeringRuleRequest) *CreateanewPeeringRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the createanew peering rule params
func (o *CreateanewPeeringRuleParams) SetBody(body *models.CreateanewPeeringRuleRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the createanew peering rule params
func (o *CreateanewPeeringRuleParams) WithContentType(contentType string) *CreateanewPeeringRuleParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the createanew peering rule params
func (o *CreateanewPeeringRuleParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateanewPeeringRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
