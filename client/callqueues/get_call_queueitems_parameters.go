// Code generated by go-swagger; DO NOT EDIT.

package callqueues

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

// NewGetCallQueueitemsParams creates a new GetCallQueueitemsParams object
// with the default values initialized.
func NewGetCallQueueitemsParams() *GetCallQueueitemsParams {
	var ()
	return &GetCallQueueitemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCallQueueitemsParamsWithTimeout creates a new GetCallQueueitemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCallQueueitemsParamsWithTimeout(timeout time.Duration) *GetCallQueueitemsParams {
	var ()
	return &GetCallQueueitemsParams{

		timeout: timeout,
	}
}

// NewGetCallQueueitemsParamsWithContext creates a new GetCallQueueitemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCallQueueitemsParamsWithContext(ctx context.Context) *GetCallQueueitemsParams {
	var ()
	return &GetCallQueueitemsParams{

		Context: ctx,
	}
}

// NewGetCallQueueitemsParamsWithHTTPClient creates a new GetCallQueueitemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCallQueueitemsParamsWithHTTPClient(client *http.Client) *GetCallQueueitemsParams {
	var ()
	return &GetCallQueueitemsParams{
		HTTPClient: client,
	}
}

/*GetCallQueueitemsParams contains all the parameters to send to the API endpoint
for the get call queueitems operation typically these are written to a http.Request
*/
type GetCallQueueitemsParams struct {

	/*Number
	  Filter for callqueues of subscribers with numbers matching the given pattern.

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
	/*ResellerID
	  Filter for callqueues of subscribers belonging to a specific reseller

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

// WithTimeout adds the timeout to the get call queueitems params
func (o *GetCallQueueitemsParams) WithTimeout(timeout time.Duration) *GetCallQueueitemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get call queueitems params
func (o *GetCallQueueitemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get call queueitems params
func (o *GetCallQueueitemsParams) WithContext(ctx context.Context) *GetCallQueueitemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get call queueitems params
func (o *GetCallQueueitemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get call queueitems params
func (o *GetCallQueueitemsParams) WithHTTPClient(client *http.Client) *GetCallQueueitemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get call queueitems params
func (o *GetCallQueueitemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNumber adds the number to the get call queueitems params
func (o *GetCallQueueitemsParams) WithNumber(number string) *GetCallQueueitemsParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the get call queueitems params
func (o *GetCallQueueitemsParams) SetNumber(number string) {
	o.Number = number
}

// WithOrderBy adds the orderBy to the get call queueitems params
func (o *GetCallQueueitemsParams) WithOrderBy(orderBy string) *GetCallQueueitemsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get call queueitems params
func (o *GetCallQueueitemsParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the get call queueitems params
func (o *GetCallQueueitemsParams) WithOrderByDirection(orderByDirection string) *GetCallQueueitemsParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the get call queueitems params
func (o *GetCallQueueitemsParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the get call queueitems params
func (o *GetCallQueueitemsParams) WithPage(page string) *GetCallQueueitemsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get call queueitems params
func (o *GetCallQueueitemsParams) SetPage(page string) {
	o.Page = page
}

// WithResellerID adds the resellerID to the get call queueitems params
func (o *GetCallQueueitemsParams) WithResellerID(resellerID string) *GetCallQueueitemsParams {
	o.SetResellerID(resellerID)
	return o
}

// SetResellerID adds the resellerId to the get call queueitems params
func (o *GetCallQueueitemsParams) SetResellerID(resellerID string) {
	o.ResellerID = resellerID
}

// WithRows adds the rows to the get call queueitems params
func (o *GetCallQueueitemsParams) WithRows(rows string) *GetCallQueueitemsParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the get call queueitems params
func (o *GetCallQueueitemsParams) SetRows(rows string) {
	o.Rows = rows
}

// WriteToRequest writes these params to a swagger request
func (o *GetCallQueueitemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
