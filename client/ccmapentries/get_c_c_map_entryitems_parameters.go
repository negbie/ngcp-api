// Code generated by go-swagger; DO NOT EDIT.

package ccmapentries

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

// NewGetCCMapEntryitemsParams creates a new GetCCMapEntryitemsParams object
// with the default values initialized.
func NewGetCCMapEntryitemsParams() *GetCCMapEntryitemsParams {
	var ()
	return &GetCCMapEntryitemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCCMapEntryitemsParamsWithTimeout creates a new GetCCMapEntryitemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCCMapEntryitemsParamsWithTimeout(timeout time.Duration) *GetCCMapEntryitemsParams {
	var ()
	return &GetCCMapEntryitemsParams{

		timeout: timeout,
	}
}

// NewGetCCMapEntryitemsParamsWithContext creates a new GetCCMapEntryitemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCCMapEntryitemsParamsWithContext(ctx context.Context) *GetCCMapEntryitemsParams {
	var ()
	return &GetCCMapEntryitemsParams{

		Context: ctx,
	}
}

// NewGetCCMapEntryitemsParamsWithHTTPClient creates a new GetCCMapEntryitemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCCMapEntryitemsParamsWithHTTPClient(client *http.Client) *GetCCMapEntryitemsParams {
	var ()
	return &GetCCMapEntryitemsParams{
		HTTPClient: client,
	}
}

/*GetCCMapEntryitemsParams contains all the parameters to send to the API endpoint
for the get c c map entryitems operation typically these are written to a http.Request
*/
type GetCCMapEntryitemsParams struct {

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

// WithTimeout adds the timeout to the get c c map entryitems params
func (o *GetCCMapEntryitemsParams) WithTimeout(timeout time.Duration) *GetCCMapEntryitemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get c c map entryitems params
func (o *GetCCMapEntryitemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get c c map entryitems params
func (o *GetCCMapEntryitemsParams) WithContext(ctx context.Context) *GetCCMapEntryitemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get c c map entryitems params
func (o *GetCCMapEntryitemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get c c map entryitems params
func (o *GetCCMapEntryitemsParams) WithHTTPClient(client *http.Client) *GetCCMapEntryitemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get c c map entryitems params
func (o *GetCCMapEntryitemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderBy adds the orderBy to the get c c map entryitems params
func (o *GetCCMapEntryitemsParams) WithOrderBy(orderBy string) *GetCCMapEntryitemsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get c c map entryitems params
func (o *GetCCMapEntryitemsParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the get c c map entryitems params
func (o *GetCCMapEntryitemsParams) WithOrderByDirection(orderByDirection string) *GetCCMapEntryitemsParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the get c c map entryitems params
func (o *GetCCMapEntryitemsParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the get c c map entryitems params
func (o *GetCCMapEntryitemsParams) WithPage(page string) *GetCCMapEntryitemsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get c c map entryitems params
func (o *GetCCMapEntryitemsParams) SetPage(page string) {
	o.Page = page
}

// WithRows adds the rows to the get c c map entryitems params
func (o *GetCCMapEntryitemsParams) WithRows(rows string) *GetCCMapEntryitemsParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the get c c map entryitems params
func (o *GetCCMapEntryitemsParams) SetRows(rows string) {
	o.Rows = rows
}

// WriteToRequest writes these params to a swagger request
func (o *GetCCMapEntryitemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
