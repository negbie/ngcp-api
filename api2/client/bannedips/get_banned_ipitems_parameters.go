// Code generated by go-swagger; DO NOT EDIT.

package bannedips

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

// NewGetBannedIpitemsParams creates a new GetBannedIpitemsParams object
// with the default values initialized.
func NewGetBannedIpitemsParams() *GetBannedIpitemsParams {
	var ()
	return &GetBannedIpitemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBannedIpitemsParamsWithTimeout creates a new GetBannedIpitemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBannedIpitemsParamsWithTimeout(timeout time.Duration) *GetBannedIpitemsParams {
	var ()
	return &GetBannedIpitemsParams{

		timeout: timeout,
	}
}

// NewGetBannedIpitemsParamsWithContext creates a new GetBannedIpitemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBannedIpitemsParamsWithContext(ctx context.Context) *GetBannedIpitemsParams {
	var ()
	return &GetBannedIpitemsParams{

		Context: ctx,
	}
}

// NewGetBannedIpitemsParamsWithHTTPClient creates a new GetBannedIpitemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBannedIpitemsParamsWithHTTPClient(client *http.Client) *GetBannedIpitemsParams {
	var ()
	return &GetBannedIpitemsParams{
		HTTPClient: client,
	}
}

/*GetBannedIpitemsParams contains all the parameters to send to the API endpoint
for the get banned ipitems operation typically these are written to a http.Request
*/
type GetBannedIpitemsParams struct {

	/*Page
	  Pagination page which should be displayed (default: 1)

	*/
	Page string
	/*Rows
	  Number of rows in one pagination page (default: 10)

	*/
	Rows string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get banned ipitems params
func (o *GetBannedIpitemsParams) WithTimeout(timeout time.Duration) *GetBannedIpitemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get banned ipitems params
func (o *GetBannedIpitemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get banned ipitems params
func (o *GetBannedIpitemsParams) WithContext(ctx context.Context) *GetBannedIpitemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get banned ipitems params
func (o *GetBannedIpitemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get banned ipitems params
func (o *GetBannedIpitemsParams) WithHTTPClient(client *http.Client) *GetBannedIpitemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get banned ipitems params
func (o *GetBannedIpitemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get banned ipitems params
func (o *GetBannedIpitemsParams) WithPage(page string) *GetBannedIpitemsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get banned ipitems params
func (o *GetBannedIpitemsParams) SetPage(page string) {
	o.Page = page
}

// WithRows adds the rows to the get banned ipitems params
func (o *GetBannedIpitemsParams) WithRows(rows string) *GetBannedIpitemsParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the get banned ipitems params
func (o *GetBannedIpitemsParams) SetRows(rows string) {
	o.Rows = rows
}

// WriteToRequest writes these params to a swagger request
func (o *GetBannedIpitemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param page
	qrPage := o.Page
	qPage := qrPage
	if qPage != "" {
		if err := r.SetQueryParam("page", qPage); err != nil {
			return err
		}
	}

	// query param rows
	qrRows := o.Rows
	qRows := qrRows
	if qRows != "" {
		if err := r.SetQueryParam("rows", qRows); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
