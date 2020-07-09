// Code generated by go-swagger; DO NOT EDIT.

package cftimesets

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

// NewGetCFTimeSetitemsParams creates a new GetCFTimeSetitemsParams object
// with the default values initialized.
func NewGetCFTimeSetitemsParams() *GetCFTimeSetitemsParams {
	var ()
	return &GetCFTimeSetitemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCFTimeSetitemsParamsWithTimeout creates a new GetCFTimeSetitemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCFTimeSetitemsParamsWithTimeout(timeout time.Duration) *GetCFTimeSetitemsParams {
	var ()
	return &GetCFTimeSetitemsParams{

		timeout: timeout,
	}
}

// NewGetCFTimeSetitemsParamsWithContext creates a new GetCFTimeSetitemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCFTimeSetitemsParamsWithContext(ctx context.Context) *GetCFTimeSetitemsParams {
	var ()
	return &GetCFTimeSetitemsParams{

		Context: ctx,
	}
}

// NewGetCFTimeSetitemsParamsWithHTTPClient creates a new GetCFTimeSetitemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCFTimeSetitemsParamsWithHTTPClient(client *http.Client) *GetCFTimeSetitemsParams {
	var ()
	return &GetCFTimeSetitemsParams{
		HTTPClient: client,
	}
}

/*GetCFTimeSetitemsParams contains all the parameters to send to the API endpoint
for the get c f time setitems operation typically these are written to a http.Request
*/
type GetCFTimeSetitemsParams struct {

	/*Name
	  Filter for contacts matching a timeset name pattern

	*/
	Name string
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
	/*SubscriberID
	  Filter for timesets belonging to a specific subscriber

	*/
	SubscriberID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get c f time setitems params
func (o *GetCFTimeSetitemsParams) WithTimeout(timeout time.Duration) *GetCFTimeSetitemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get c f time setitems params
func (o *GetCFTimeSetitemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get c f time setitems params
func (o *GetCFTimeSetitemsParams) WithContext(ctx context.Context) *GetCFTimeSetitemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get c f time setitems params
func (o *GetCFTimeSetitemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get c f time setitems params
func (o *GetCFTimeSetitemsParams) WithHTTPClient(client *http.Client) *GetCFTimeSetitemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get c f time setitems params
func (o *GetCFTimeSetitemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get c f time setitems params
func (o *GetCFTimeSetitemsParams) WithName(name string) *GetCFTimeSetitemsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get c f time setitems params
func (o *GetCFTimeSetitemsParams) SetName(name string) {
	o.Name = name
}

// WithOrderBy adds the orderBy to the get c f time setitems params
func (o *GetCFTimeSetitemsParams) WithOrderBy(orderBy string) *GetCFTimeSetitemsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get c f time setitems params
func (o *GetCFTimeSetitemsParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the get c f time setitems params
func (o *GetCFTimeSetitemsParams) WithOrderByDirection(orderByDirection string) *GetCFTimeSetitemsParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the get c f time setitems params
func (o *GetCFTimeSetitemsParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the get c f time setitems params
func (o *GetCFTimeSetitemsParams) WithPage(page string) *GetCFTimeSetitemsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get c f time setitems params
func (o *GetCFTimeSetitemsParams) SetPage(page string) {
	o.Page = page
}

// WithRows adds the rows to the get c f time setitems params
func (o *GetCFTimeSetitemsParams) WithRows(rows string) *GetCFTimeSetitemsParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the get c f time setitems params
func (o *GetCFTimeSetitemsParams) SetRows(rows string) {
	o.Rows = rows
}

// WithSubscriberID adds the subscriberID to the get c f time setitems params
func (o *GetCFTimeSetitemsParams) WithSubscriberID(subscriberID string) *GetCFTimeSetitemsParams {
	o.SetSubscriberID(subscriberID)
	return o
}

// SetSubscriberID adds the subscriberId to the get c f time setitems params
func (o *GetCFTimeSetitemsParams) SetSubscriberID(subscriberID string) {
	o.SubscriberID = subscriberID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCFTimeSetitemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param name
	qrName := o.Name
	qName := qrName
	if qName != "" {
		if err := r.SetQueryParam("name", qName); err != nil {
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

	// query param subscriber_id
	qrSubscriberID := o.SubscriberID
	qSubscriberID := qrSubscriberID
	if qSubscriberID != "" {
		if err := r.SetQueryParam("subscriber_id", qSubscriberID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
