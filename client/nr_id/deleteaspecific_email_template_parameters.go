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

// NewDeleteaspecificEmailTemplateParams creates a new DeleteaspecificEmailTemplateParams object
// with the default values initialized.
func NewDeleteaspecificEmailTemplateParams() *DeleteaspecificEmailTemplateParams {
	var ()
	return &DeleteaspecificEmailTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteaspecificEmailTemplateParamsWithTimeout creates a new DeleteaspecificEmailTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteaspecificEmailTemplateParamsWithTimeout(timeout time.Duration) *DeleteaspecificEmailTemplateParams {
	var ()
	return &DeleteaspecificEmailTemplateParams{

		timeout: timeout,
	}
}

// NewDeleteaspecificEmailTemplateParamsWithContext creates a new DeleteaspecificEmailTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteaspecificEmailTemplateParamsWithContext(ctx context.Context) *DeleteaspecificEmailTemplateParams {
	var ()
	return &DeleteaspecificEmailTemplateParams{

		Context: ctx,
	}
}

// NewDeleteaspecificEmailTemplateParamsWithHTTPClient creates a new DeleteaspecificEmailTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteaspecificEmailTemplateParamsWithHTTPClient(client *http.Client) *DeleteaspecificEmailTemplateParams {
	var ()
	return &DeleteaspecificEmailTemplateParams{
		HTTPClient: client,
	}
}

/*DeleteaspecificEmailTemplateParams contains all the parameters to send to the API endpoint
for the deleteaspecific email template operation typically these are written to a http.Request
*/
type DeleteaspecificEmailTemplateParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the deleteaspecific email template params
func (o *DeleteaspecificEmailTemplateParams) WithTimeout(timeout time.Duration) *DeleteaspecificEmailTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deleteaspecific email template params
func (o *DeleteaspecificEmailTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deleteaspecific email template params
func (o *DeleteaspecificEmailTemplateParams) WithContext(ctx context.Context) *DeleteaspecificEmailTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deleteaspecific email template params
func (o *DeleteaspecificEmailTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deleteaspecific email template params
func (o *DeleteaspecificEmailTemplateParams) WithHTTPClient(client *http.Client) *DeleteaspecificEmailTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deleteaspecific email template params
func (o *DeleteaspecificEmailTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the deleteaspecific email template params
func (o *DeleteaspecificEmailTemplateParams) WithID(id string) *DeleteaspecificEmailTemplateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the deleteaspecific email template params
func (o *DeleteaspecificEmailTemplateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteaspecificEmailTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
