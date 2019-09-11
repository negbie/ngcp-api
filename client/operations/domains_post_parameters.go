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

// NewDomainsPostParams creates a new DomainsPostParams object
// with the default values initialized.
func NewDomainsPostParams() *DomainsPostParams {
	var ()
	return &DomainsPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDomainsPostParamsWithTimeout creates a new DomainsPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDomainsPostParamsWithTimeout(timeout time.Duration) *DomainsPostParams {
	var ()
	return &DomainsPostParams{

		timeout: timeout,
	}
}

// NewDomainsPostParamsWithContext creates a new DomainsPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewDomainsPostParamsWithContext(ctx context.Context) *DomainsPostParams {
	var ()
	return &DomainsPostParams{

		Context: ctx,
	}
}

// NewDomainsPostParamsWithHTTPClient creates a new DomainsPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDomainsPostParamsWithHTTPClient(client *http.Client) *DomainsPostParams {
	var ()
	return &DomainsPostParams{
		HTTPClient: client,
	}
}

/*DomainsPostParams contains all the parameters to send to the API endpoint
for the domains post operation typically these are written to a http.Request
*/
type DomainsPostParams struct {

	/*Body*/
	Body *models.CreateanewDomainRequest
	/*ContentType*/
	ContentType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the domains post params
func (o *DomainsPostParams) WithTimeout(timeout time.Duration) *DomainsPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the domains post params
func (o *DomainsPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the domains post params
func (o *DomainsPostParams) WithContext(ctx context.Context) *DomainsPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the domains post params
func (o *DomainsPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the domains post params
func (o *DomainsPostParams) WithHTTPClient(client *http.Client) *DomainsPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the domains post params
func (o *DomainsPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the domains post params
func (o *DomainsPostParams) WithBody(body *models.CreateanewDomainRequest) *DomainsPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the domains post params
func (o *DomainsPostParams) SetBody(body *models.CreateanewDomainRequest) {
	o.Body = body
}

// WithContentType adds the contentType to the domains post params
func (o *DomainsPostParams) WithContentType(contentType string) *DomainsPostParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the domains post params
func (o *DomainsPostParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WriteToRequest writes these params to a swagger request
func (o *DomainsPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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