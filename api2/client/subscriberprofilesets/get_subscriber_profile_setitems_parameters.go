// Code generated by go-swagger; DO NOT EDIT.

package subscriberprofilesets

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

// NewGetSubscriberProfileSetitemsParams creates a new GetSubscriberProfileSetitemsParams object
// with the default values initialized.
func NewGetSubscriberProfileSetitemsParams() *GetSubscriberProfileSetitemsParams {
	var ()
	return &GetSubscriberProfileSetitemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSubscriberProfileSetitemsParamsWithTimeout creates a new GetSubscriberProfileSetitemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSubscriberProfileSetitemsParamsWithTimeout(timeout time.Duration) *GetSubscriberProfileSetitemsParams {
	var ()
	return &GetSubscriberProfileSetitemsParams{

		timeout: timeout,
	}
}

// NewGetSubscriberProfileSetitemsParamsWithContext creates a new GetSubscriberProfileSetitemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSubscriberProfileSetitemsParamsWithContext(ctx context.Context) *GetSubscriberProfileSetitemsParams {
	var ()
	return &GetSubscriberProfileSetitemsParams{

		Context: ctx,
	}
}

// NewGetSubscriberProfileSetitemsParamsWithHTTPClient creates a new GetSubscriberProfileSetitemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSubscriberProfileSetitemsParamsWithHTTPClient(client *http.Client) *GetSubscriberProfileSetitemsParams {
	var ()
	return &GetSubscriberProfileSetitemsParams{
		HTTPClient: client,
	}
}

/*GetSubscriberProfileSetitemsParams contains all the parameters to send to the API endpoint
for the get subscriber profile setitems operation typically these are written to a http.Request
*/
type GetSubscriberProfileSetitemsParams struct {

	/*Name
	  Filter for profile sets with a specific name

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
	/*ResellerID
	  Filter for profile sets belonging to a specific reseller

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

// WithTimeout adds the timeout to the get subscriber profile setitems params
func (o *GetSubscriberProfileSetitemsParams) WithTimeout(timeout time.Duration) *GetSubscriberProfileSetitemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get subscriber profile setitems params
func (o *GetSubscriberProfileSetitemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get subscriber profile setitems params
func (o *GetSubscriberProfileSetitemsParams) WithContext(ctx context.Context) *GetSubscriberProfileSetitemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get subscriber profile setitems params
func (o *GetSubscriberProfileSetitemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get subscriber profile setitems params
func (o *GetSubscriberProfileSetitemsParams) WithHTTPClient(client *http.Client) *GetSubscriberProfileSetitemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get subscriber profile setitems params
func (o *GetSubscriberProfileSetitemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get subscriber profile setitems params
func (o *GetSubscriberProfileSetitemsParams) WithName(name string) *GetSubscriberProfileSetitemsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get subscriber profile setitems params
func (o *GetSubscriberProfileSetitemsParams) SetName(name string) {
	o.Name = name
}

// WithOrderBy adds the orderBy to the get subscriber profile setitems params
func (o *GetSubscriberProfileSetitemsParams) WithOrderBy(orderBy string) *GetSubscriberProfileSetitemsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get subscriber profile setitems params
func (o *GetSubscriberProfileSetitemsParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the get subscriber profile setitems params
func (o *GetSubscriberProfileSetitemsParams) WithOrderByDirection(orderByDirection string) *GetSubscriberProfileSetitemsParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the get subscriber profile setitems params
func (o *GetSubscriberProfileSetitemsParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the get subscriber profile setitems params
func (o *GetSubscriberProfileSetitemsParams) WithPage(page string) *GetSubscriberProfileSetitemsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get subscriber profile setitems params
func (o *GetSubscriberProfileSetitemsParams) SetPage(page string) {
	o.Page = page
}

// WithResellerID adds the resellerID to the get subscriber profile setitems params
func (o *GetSubscriberProfileSetitemsParams) WithResellerID(resellerID string) *GetSubscriberProfileSetitemsParams {
	o.SetResellerID(resellerID)
	return o
}

// SetResellerID adds the resellerId to the get subscriber profile setitems params
func (o *GetSubscriberProfileSetitemsParams) SetResellerID(resellerID string) {
	o.ResellerID = resellerID
}

// WithRows adds the rows to the get subscriber profile setitems params
func (o *GetSubscriberProfileSetitemsParams) WithRows(rows string) *GetSubscriberProfileSetitemsParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the get subscriber profile setitems params
func (o *GetSubscriberProfileSetitemsParams) SetRows(rows string) {
	o.Rows = rows
}

// WriteToRequest writes these params to a swagger request
func (o *GetSubscriberProfileSetitemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
