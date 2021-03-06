// Code generated by go-swagger; DO NOT EDIT.

package ncoslevels

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

// NewGetNcosLevelitemsParams creates a new GetNcosLevelitemsParams object
// with the default values initialized.
func NewGetNcosLevelitemsParams() *GetNcosLevelitemsParams {
	var ()
	return &GetNcosLevelitemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNcosLevelitemsParamsWithTimeout creates a new GetNcosLevelitemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNcosLevelitemsParamsWithTimeout(timeout time.Duration) *GetNcosLevelitemsParams {
	var ()
	return &GetNcosLevelitemsParams{

		timeout: timeout,
	}
}

// NewGetNcosLevelitemsParamsWithContext creates a new GetNcosLevelitemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNcosLevelitemsParamsWithContext(ctx context.Context) *GetNcosLevelitemsParams {
	var ()
	return &GetNcosLevelitemsParams{

		Context: ctx,
	}
}

// NewGetNcosLevelitemsParamsWithHTTPClient creates a new GetNcosLevelitemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNcosLevelitemsParamsWithHTTPClient(client *http.Client) *GetNcosLevelitemsParams {
	var ()
	return &GetNcosLevelitemsParams{
		HTTPClient: client,
	}
}

/*GetNcosLevelitemsParams contains all the parameters to send to the API endpoint
for the get ncos levelitems operation typically these are written to a http.Request
*/
type GetNcosLevelitemsParams struct {

	/*Level
	  Filter for levels matching the given pattern

	*/
	Level string
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
	/*ResellerID
	  Filter for ncos levels belonging to a specific reseller

	*/
	ResellerID string
	/*Rows
	  Number of rows in one pagination page (default: 10)

	*/
	Rows string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ncos levelitems params
func (o *GetNcosLevelitemsParams) WithTimeout(timeout time.Duration) *GetNcosLevelitemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ncos levelitems params
func (o *GetNcosLevelitemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ncos levelitems params
func (o *GetNcosLevelitemsParams) WithContext(ctx context.Context) *GetNcosLevelitemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ncos levelitems params
func (o *GetNcosLevelitemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ncos levelitems params
func (o *GetNcosLevelitemsParams) WithHTTPClient(client *http.Client) *GetNcosLevelitemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ncos levelitems params
func (o *GetNcosLevelitemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLevel adds the level to the get ncos levelitems params
func (o *GetNcosLevelitemsParams) WithLevel(level string) *GetNcosLevelitemsParams {
	o.SetLevel(level)
	return o
}

// SetLevel adds the level to the get ncos levelitems params
func (o *GetNcosLevelitemsParams) SetLevel(level string) {
	o.Level = level
}

// WithOrderBy adds the orderBy to the get ncos levelitems params
func (o *GetNcosLevelitemsParams) WithOrderBy(orderBy string) *GetNcosLevelitemsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get ncos levelitems params
func (o *GetNcosLevelitemsParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the get ncos levelitems params
func (o *GetNcosLevelitemsParams) WithOrderByDirection(orderByDirection string) *GetNcosLevelitemsParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the get ncos levelitems params
func (o *GetNcosLevelitemsParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the get ncos levelitems params
func (o *GetNcosLevelitemsParams) WithPage(page string) *GetNcosLevelitemsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get ncos levelitems params
func (o *GetNcosLevelitemsParams) SetPage(page string) {
	o.Page = page
}

// WithResellerID adds the resellerID to the get ncos levelitems params
func (o *GetNcosLevelitemsParams) WithResellerID(resellerID string) *GetNcosLevelitemsParams {
	o.SetResellerID(resellerID)
	return o
}

// SetResellerID adds the resellerId to the get ncos levelitems params
func (o *GetNcosLevelitemsParams) SetResellerID(resellerID string) {
	o.ResellerID = resellerID
}

// WithRows adds the rows to the get ncos levelitems params
func (o *GetNcosLevelitemsParams) WithRows(rows string) *GetNcosLevelitemsParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the get ncos levelitems params
func (o *GetNcosLevelitemsParams) SetRows(rows string) {
	o.Rows = rows
}

// WriteToRequest writes these params to a swagger request
func (o *GetNcosLevelitemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param level
	qrLevel := o.Level
	qLevel := qrLevel
	if qLevel != "" {
		if err := r.SetQueryParam("level", qLevel); err != nil {
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

	// query param reseller_id
	qrResellerID := o.ResellerID
	qResellerID := qrResellerID
	if qResellerID != "" {
		if err := r.SetQueryParam("reseller_id", qResellerID); err != nil {
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
