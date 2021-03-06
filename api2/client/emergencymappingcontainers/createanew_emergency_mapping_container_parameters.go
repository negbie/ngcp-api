// Code generated by go-swagger; DO NOT EDIT.

package emergencymappingcontainers

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

// NewCreateanewEmergencyMappingContainerParams creates a new CreateanewEmergencyMappingContainerParams object
// with the default values initialized.
func NewCreateanewEmergencyMappingContainerParams() *CreateanewEmergencyMappingContainerParams {
	var ()
	return &CreateanewEmergencyMappingContainerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateanewEmergencyMappingContainerParamsWithTimeout creates a new CreateanewEmergencyMappingContainerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateanewEmergencyMappingContainerParamsWithTimeout(timeout time.Duration) *CreateanewEmergencyMappingContainerParams {
	var ()
	return &CreateanewEmergencyMappingContainerParams{

		timeout: timeout,
	}
}

// NewCreateanewEmergencyMappingContainerParamsWithContext creates a new CreateanewEmergencyMappingContainerParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateanewEmergencyMappingContainerParamsWithContext(ctx context.Context) *CreateanewEmergencyMappingContainerParams {
	var ()
	return &CreateanewEmergencyMappingContainerParams{

		Context: ctx,
	}
}

// NewCreateanewEmergencyMappingContainerParamsWithHTTPClient creates a new CreateanewEmergencyMappingContainerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateanewEmergencyMappingContainerParamsWithHTTPClient(client *http.Client) *CreateanewEmergencyMappingContainerParams {
	var ()
	return &CreateanewEmergencyMappingContainerParams{
		HTTPClient: client,
	}
}

/*CreateanewEmergencyMappingContainerParams contains all the parameters to send to the API endpoint
for the createanew emergency mapping container operation typically these are written to a http.Request
*/
type CreateanewEmergencyMappingContainerParams struct {

	/*Body*/
	Body *models.CreateanewEmergencyMappingContainerRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the createanew emergency mapping container params
func (o *CreateanewEmergencyMappingContainerParams) WithTimeout(timeout time.Duration) *CreateanewEmergencyMappingContainerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the createanew emergency mapping container params
func (o *CreateanewEmergencyMappingContainerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the createanew emergency mapping container params
func (o *CreateanewEmergencyMappingContainerParams) WithContext(ctx context.Context) *CreateanewEmergencyMappingContainerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the createanew emergency mapping container params
func (o *CreateanewEmergencyMappingContainerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the createanew emergency mapping container params
func (o *CreateanewEmergencyMappingContainerParams) WithHTTPClient(client *http.Client) *CreateanewEmergencyMappingContainerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the createanew emergency mapping container params
func (o *CreateanewEmergencyMappingContainerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the createanew emergency mapping container params
func (o *CreateanewEmergencyMappingContainerParams) WithBody(body *models.CreateanewEmergencyMappingContainerRequest) *CreateanewEmergencyMappingContainerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the createanew emergency mapping container params
func (o *CreateanewEmergencyMappingContainerParams) SetBody(body *models.CreateanewEmergencyMappingContainerRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the createanew emergency mapping container params
func (o *CreateanewEmergencyMappingContainerParams) WithContentType(contentType string) *CreateanewEmergencyMappingContainerParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the createanew emergency mapping container params
func (o *CreateanewEmergencyMappingContainerParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateanewEmergencyMappingContainerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
