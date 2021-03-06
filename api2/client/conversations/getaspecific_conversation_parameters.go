// Code generated by go-swagger; DO NOT EDIT.

package conversations

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

// NewGetaspecificConversationParams creates a new GetaspecificConversationParams object
// with the default values initialized.
func NewGetaspecificConversationParams() *GetaspecificConversationParams {
	var ()
	return &GetaspecificConversationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificConversationParamsWithTimeout creates a new GetaspecificConversationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificConversationParamsWithTimeout(timeout time.Duration) *GetaspecificConversationParams {
	var ()
	return &GetaspecificConversationParams{

		timeout: timeout,
	}
}

// NewGetaspecificConversationParamsWithContext creates a new GetaspecificConversationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificConversationParamsWithContext(ctx context.Context) *GetaspecificConversationParams {
	var ()
	return &GetaspecificConversationParams{

		Context: ctx,
	}
}

// NewGetaspecificConversationParamsWithHTTPClient creates a new GetaspecificConversationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificConversationParamsWithHTTPClient(client *http.Client) *GetaspecificConversationParams {
	var ()
	return &GetaspecificConversationParams{
		HTTPClient: client,
	}
}

/*GetaspecificConversationParams contains all the parameters to send to the API endpoint
for the getaspecific conversation operation typically these are written to a http.Request
*/
type GetaspecificConversationParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific conversation params
func (o *GetaspecificConversationParams) WithTimeout(timeout time.Duration) *GetaspecificConversationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific conversation params
func (o *GetaspecificConversationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific conversation params
func (o *GetaspecificConversationParams) WithContext(ctx context.Context) *GetaspecificConversationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific conversation params
func (o *GetaspecificConversationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific conversation params
func (o *GetaspecificConversationParams) WithHTTPClient(client *http.Client) *GetaspecificConversationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific conversation params
func (o *GetaspecificConversationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific conversation params
func (o *GetaspecificConversationParams) WithID(id string) *GetaspecificConversationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific conversation params
func (o *GetaspecificConversationParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificConversationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
