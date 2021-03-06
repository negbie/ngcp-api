// Code generated by go-swagger; DO NOT EDIT.

package interceptions

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

// NewGetInterceptionitemsParams creates a new GetInterceptionitemsParams object
// with the default values initialized.
func NewGetInterceptionitemsParams() *GetInterceptionitemsParams {
	var ()
	return &GetInterceptionitemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInterceptionitemsParamsWithTimeout creates a new GetInterceptionitemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInterceptionitemsParamsWithTimeout(timeout time.Duration) *GetInterceptionitemsParams {
	var ()
	return &GetInterceptionitemsParams{

		timeout: timeout,
	}
}

// NewGetInterceptionitemsParamsWithContext creates a new GetInterceptionitemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInterceptionitemsParamsWithContext(ctx context.Context) *GetInterceptionitemsParams {
	var ()
	return &GetInterceptionitemsParams{

		Context: ctx,
	}
}

// NewGetInterceptionitemsParamsWithHTTPClient creates a new GetInterceptionitemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInterceptionitemsParamsWithHTTPClient(client *http.Client) *GetInterceptionitemsParams {
	var ()
	return &GetInterceptionitemsParams{
		HTTPClient: client,
	}
}

/*GetInterceptionitemsParams contains all the parameters to send to the API endpoint
for the get interceptionitems operation typically these are written to a http.Request
*/
type GetInterceptionitemsParams struct {

	/*Liid
	  Filter for interceptions of a specific interception id

	*/
	Liid string
	/*Number
	  Filter for interceptions of a specific number (in E.164 format)

	*/
	Number string
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

// WithTimeout adds the timeout to the get interceptionitems params
func (o *GetInterceptionitemsParams) WithTimeout(timeout time.Duration) *GetInterceptionitemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get interceptionitems params
func (o *GetInterceptionitemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get interceptionitems params
func (o *GetInterceptionitemsParams) WithContext(ctx context.Context) *GetInterceptionitemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get interceptionitems params
func (o *GetInterceptionitemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get interceptionitems params
func (o *GetInterceptionitemsParams) WithHTTPClient(client *http.Client) *GetInterceptionitemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get interceptionitems params
func (o *GetInterceptionitemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLiid adds the liid to the get interceptionitems params
func (o *GetInterceptionitemsParams) WithLiid(liid string) *GetInterceptionitemsParams {
	o.SetLiid(liid)
	return o
}

// SetLiid adds the liid to the get interceptionitems params
func (o *GetInterceptionitemsParams) SetLiid(liid string) {
	o.Liid = liid
}

// WithNumber adds the number to the get interceptionitems params
func (o *GetInterceptionitemsParams) WithNumber(number string) *GetInterceptionitemsParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the get interceptionitems params
func (o *GetInterceptionitemsParams) SetNumber(number string) {
	o.Number = number
}

// WithOrderBy adds the orderBy to the get interceptionitems params
func (o *GetInterceptionitemsParams) WithOrderBy(orderBy string) *GetInterceptionitemsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get interceptionitems params
func (o *GetInterceptionitemsParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the get interceptionitems params
func (o *GetInterceptionitemsParams) WithOrderByDirection(orderByDirection string) *GetInterceptionitemsParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the get interceptionitems params
func (o *GetInterceptionitemsParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the get interceptionitems params
func (o *GetInterceptionitemsParams) WithPage(page string) *GetInterceptionitemsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get interceptionitems params
func (o *GetInterceptionitemsParams) SetPage(page string) {
	o.Page = page
}

// WithRows adds the rows to the get interceptionitems params
func (o *GetInterceptionitemsParams) WithRows(rows string) *GetInterceptionitemsParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the get interceptionitems params
func (o *GetInterceptionitemsParams) SetRows(rows string) {
	o.Rows = rows
}

// WriteToRequest writes these params to a swagger request
func (o *GetInterceptionitemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param liid
	qrLiid := o.Liid
	qLiid := qrLiid
	if qLiid != "" {
		if err := r.SetQueryParam("liid", qLiid); err != nil {
			return err
		}
	}

	// query param number
	qrNumber := o.Number
	qNumber := qrNumber
	if qNumber != "" {
		if err := r.SetQueryParam("number", qNumber); err != nil {
			return err
		}
	}

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
