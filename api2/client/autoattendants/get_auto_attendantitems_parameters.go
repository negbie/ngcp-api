// Code generated by go-swagger; DO NOT EDIT.

package autoattendants

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

// NewGetAutoAttendantitemsParams creates a new GetAutoAttendantitemsParams object
// with the default values initialized.
func NewGetAutoAttendantitemsParams() *GetAutoAttendantitemsParams {
	var ()
	return &GetAutoAttendantitemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAutoAttendantitemsParamsWithTimeout creates a new GetAutoAttendantitemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAutoAttendantitemsParamsWithTimeout(timeout time.Duration) *GetAutoAttendantitemsParams {
	var ()
	return &GetAutoAttendantitemsParams{

		timeout: timeout,
	}
}

// NewGetAutoAttendantitemsParamsWithContext creates a new GetAutoAttendantitemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAutoAttendantitemsParamsWithContext(ctx context.Context) *GetAutoAttendantitemsParams {
	var ()
	return &GetAutoAttendantitemsParams{

		Context: ctx,
	}
}

// NewGetAutoAttendantitemsParamsWithHTTPClient creates a new GetAutoAttendantitemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAutoAttendantitemsParamsWithHTTPClient(client *http.Client) *GetAutoAttendantitemsParams {
	var ()
	return &GetAutoAttendantitemsParams{
		HTTPClient: client,
	}
}

/*GetAutoAttendantitemsParams contains all the parameters to send to the API endpoint
for the get auto attendantitems operation typically these are written to a http.Request
*/
type GetAutoAttendantitemsParams struct {

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

// WithTimeout adds the timeout to the get auto attendantitems params
func (o *GetAutoAttendantitemsParams) WithTimeout(timeout time.Duration) *GetAutoAttendantitemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get auto attendantitems params
func (o *GetAutoAttendantitemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get auto attendantitems params
func (o *GetAutoAttendantitemsParams) WithContext(ctx context.Context) *GetAutoAttendantitemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get auto attendantitems params
func (o *GetAutoAttendantitemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get auto attendantitems params
func (o *GetAutoAttendantitemsParams) WithHTTPClient(client *http.Client) *GetAutoAttendantitemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get auto attendantitems params
func (o *GetAutoAttendantitemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderBy adds the orderBy to the get auto attendantitems params
func (o *GetAutoAttendantitemsParams) WithOrderBy(orderBy string) *GetAutoAttendantitemsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get auto attendantitems params
func (o *GetAutoAttendantitemsParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the get auto attendantitems params
func (o *GetAutoAttendantitemsParams) WithOrderByDirection(orderByDirection string) *GetAutoAttendantitemsParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the get auto attendantitems params
func (o *GetAutoAttendantitemsParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the get auto attendantitems params
func (o *GetAutoAttendantitemsParams) WithPage(page string) *GetAutoAttendantitemsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get auto attendantitems params
func (o *GetAutoAttendantitemsParams) SetPage(page string) {
	o.Page = page
}

// WithRows adds the rows to the get auto attendantitems params
func (o *GetAutoAttendantitemsParams) WithRows(rows string) *GetAutoAttendantitemsParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the get auto attendantitems params
func (o *GetAutoAttendantitemsParams) SetRows(rows string) {
	o.Rows = rows
}

// WriteToRequest writes these params to a swagger request
func (o *GetAutoAttendantitemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
