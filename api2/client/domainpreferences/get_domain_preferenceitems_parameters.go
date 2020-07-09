// Code generated by go-swagger; DO NOT EDIT.

package domainpreferences

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

// NewGetDomainPreferenceitemsParams creates a new GetDomainPreferenceitemsParams object
// with the default values initialized.
func NewGetDomainPreferenceitemsParams() *GetDomainPreferenceitemsParams {
	var ()
	return &GetDomainPreferenceitemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDomainPreferenceitemsParamsWithTimeout creates a new GetDomainPreferenceitemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDomainPreferenceitemsParamsWithTimeout(timeout time.Duration) *GetDomainPreferenceitemsParams {
	var ()
	return &GetDomainPreferenceitemsParams{

		timeout: timeout,
	}
}

// NewGetDomainPreferenceitemsParamsWithContext creates a new GetDomainPreferenceitemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDomainPreferenceitemsParamsWithContext(ctx context.Context) *GetDomainPreferenceitemsParams {
	var ()
	return &GetDomainPreferenceitemsParams{

		Context: ctx,
	}
}

// NewGetDomainPreferenceitemsParamsWithHTTPClient creates a new GetDomainPreferenceitemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDomainPreferenceitemsParamsWithHTTPClient(client *http.Client) *GetDomainPreferenceitemsParams {
	var ()
	return &GetDomainPreferenceitemsParams{
		HTTPClient: client,
	}
}

/*GetDomainPreferenceitemsParams contains all the parameters to send to the API endpoint
for the get domain preferenceitems operation typically these are written to a http.Request
*/
type GetDomainPreferenceitemsParams struct {

	/*OrderBy
	  Order collection by a specific attribute.

	*/
	OrderBy string
	/*OrderByDirection
	  Direction which the collection should be ordered by. Possible values are: asc (default), desc.

	*/
	OrderByDirection string
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

// WithTimeout adds the timeout to the get domain preferenceitems params
func (o *GetDomainPreferenceitemsParams) WithTimeout(timeout time.Duration) *GetDomainPreferenceitemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get domain preferenceitems params
func (o *GetDomainPreferenceitemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get domain preferenceitems params
func (o *GetDomainPreferenceitemsParams) WithContext(ctx context.Context) *GetDomainPreferenceitemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get domain preferenceitems params
func (o *GetDomainPreferenceitemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get domain preferenceitems params
func (o *GetDomainPreferenceitemsParams) WithHTTPClient(client *http.Client) *GetDomainPreferenceitemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get domain preferenceitems params
func (o *GetDomainPreferenceitemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderBy adds the orderBy to the get domain preferenceitems params
func (o *GetDomainPreferenceitemsParams) WithOrderBy(orderBy string) *GetDomainPreferenceitemsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get domain preferenceitems params
func (o *GetDomainPreferenceitemsParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the get domain preferenceitems params
func (o *GetDomainPreferenceitemsParams) WithOrderByDirection(orderByDirection string) *GetDomainPreferenceitemsParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the get domain preferenceitems params
func (o *GetDomainPreferenceitemsParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the get domain preferenceitems params
func (o *GetDomainPreferenceitemsParams) WithPage(page string) *GetDomainPreferenceitemsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get domain preferenceitems params
func (o *GetDomainPreferenceitemsParams) SetPage(page string) {
	o.Page = page
}

// WithRows adds the rows to the get domain preferenceitems params
func (o *GetDomainPreferenceitemsParams) WithRows(rows string) *GetDomainPreferenceitemsParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the get domain preferenceitems params
func (o *GetDomainPreferenceitemsParams) SetRows(rows string) {
	o.Rows = rows
}

// WriteToRequest writes these params to a swagger request
func (o *GetDomainPreferenceitemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param order_by
	qrOrderBy := o.OrderBy
	qOrderBy := qrOrderBy
	if qOrderBy != "" {
		if err := r.SetQueryParam("order_by", qOrderBy); err != nil {
			return err
		}
	}

	// query param order_by_direction
	qrOrderByDirection := o.OrderByDirection
	qOrderByDirection := qrOrderByDirection
	if qOrderByDirection != "" {
		if err := r.SetQueryParam("order_by_direction", qOrderByDirection); err != nil {
			return err
		}
	}

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