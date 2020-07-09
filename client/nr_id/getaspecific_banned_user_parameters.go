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

// NewGetaspecificBannedUserParams creates a new GetaspecificBannedUserParams object
// with the default values initialized.
func NewGetaspecificBannedUserParams() *GetaspecificBannedUserParams {
	var ()
	return &GetaspecificBannedUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetaspecificBannedUserParamsWithTimeout creates a new GetaspecificBannedUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetaspecificBannedUserParamsWithTimeout(timeout time.Duration) *GetaspecificBannedUserParams {
	var ()
	return &GetaspecificBannedUserParams{

		timeout: timeout,
	}
}

// NewGetaspecificBannedUserParamsWithContext creates a new GetaspecificBannedUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetaspecificBannedUserParamsWithContext(ctx context.Context) *GetaspecificBannedUserParams {
	var ()
	return &GetaspecificBannedUserParams{

		Context: ctx,
	}
}

// NewGetaspecificBannedUserParamsWithHTTPClient creates a new GetaspecificBannedUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetaspecificBannedUserParamsWithHTTPClient(client *http.Client) *GetaspecificBannedUserParams {
	var ()
	return &GetaspecificBannedUserParams{
		HTTPClient: client,
	}
}

/*GetaspecificBannedUserParams contains all the parameters to send to the API endpoint
for the getaspecific banned user operation typically these are written to a http.Request
*/
type GetaspecificBannedUserParams struct {

	/*ID
	  (Required)

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the getaspecific banned user params
func (o *GetaspecificBannedUserParams) WithTimeout(timeout time.Duration) *GetaspecificBannedUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getaspecific banned user params
func (o *GetaspecificBannedUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getaspecific banned user params
func (o *GetaspecificBannedUserParams) WithContext(ctx context.Context) *GetaspecificBannedUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getaspecific banned user params
func (o *GetaspecificBannedUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getaspecific banned user params
func (o *GetaspecificBannedUserParams) WithHTTPClient(client *http.Client) *GetaspecificBannedUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getaspecific banned user params
func (o *GetaspecificBannedUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the getaspecific banned user params
func (o *GetaspecificBannedUserParams) WithID(id string) *GetaspecificBannedUserParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the getaspecific banned user params
func (o *GetaspecificBannedUserParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetaspecificBannedUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
