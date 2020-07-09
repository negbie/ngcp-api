// Code generated by go-swagger; DO NOT EDIT.

package billingprofiles

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

// NewGetBillingProfileitemsParams creates a new GetBillingProfileitemsParams object
// with the default values initialized.
func NewGetBillingProfileitemsParams() *GetBillingProfileitemsParams {
	var ()
	return &GetBillingProfileitemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBillingProfileitemsParamsWithTimeout creates a new GetBillingProfileitemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBillingProfileitemsParamsWithTimeout(timeout time.Duration) *GetBillingProfileitemsParams {
	var ()
	return &GetBillingProfileitemsParams{

		timeout: timeout,
	}
}

// NewGetBillingProfileitemsParamsWithContext creates a new GetBillingProfileitemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBillingProfileitemsParamsWithContext(ctx context.Context) *GetBillingProfileitemsParams {
	var ()
	return &GetBillingProfileitemsParams{

		Context: ctx,
	}
}

// NewGetBillingProfileitemsParamsWithHTTPClient creates a new GetBillingProfileitemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBillingProfileitemsParamsWithHTTPClient(client *http.Client) *GetBillingProfileitemsParams {
	var ()
	return &GetBillingProfileitemsParams{
		HTTPClient: client,
	}
}

/*GetBillingProfileitemsParams contains all the parameters to send to the API endpoint
for the get billing profileitems operation typically these are written to a http.Request
*/
type GetBillingProfileitemsParams struct {

	/*Handle
	  Filter for billing profiles with a specific handle

	*/
	Handle string
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
	  Filter for billing profiles belonging to a specific reseller

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

// WithTimeout adds the timeout to the get billing profileitems params
func (o *GetBillingProfileitemsParams) WithTimeout(timeout time.Duration) *GetBillingProfileitemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get billing profileitems params
func (o *GetBillingProfileitemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get billing profileitems params
func (o *GetBillingProfileitemsParams) WithContext(ctx context.Context) *GetBillingProfileitemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get billing profileitems params
func (o *GetBillingProfileitemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get billing profileitems params
func (o *GetBillingProfileitemsParams) WithHTTPClient(client *http.Client) *GetBillingProfileitemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get billing profileitems params
func (o *GetBillingProfileitemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHandle adds the handle to the get billing profileitems params
func (o *GetBillingProfileitemsParams) WithHandle(handle string) *GetBillingProfileitemsParams {
	o.SetHandle(handle)
	return o
}

// SetHandle adds the handle to the get billing profileitems params
func (o *GetBillingProfileitemsParams) SetHandle(handle string) {
	o.Handle = handle
}

// WithOrderBy adds the orderBy to the get billing profileitems params
func (o *GetBillingProfileitemsParams) WithOrderBy(orderBy string) *GetBillingProfileitemsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get billing profileitems params
func (o *GetBillingProfileitemsParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the get billing profileitems params
func (o *GetBillingProfileitemsParams) WithOrderByDirection(orderByDirection string) *GetBillingProfileitemsParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the get billing profileitems params
func (o *GetBillingProfileitemsParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the get billing profileitems params
func (o *GetBillingProfileitemsParams) WithPage(page string) *GetBillingProfileitemsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get billing profileitems params
func (o *GetBillingProfileitemsParams) SetPage(page string) {
	o.Page = page
}

// WithResellerID adds the resellerID to the get billing profileitems params
func (o *GetBillingProfileitemsParams) WithResellerID(resellerID string) *GetBillingProfileitemsParams {
	o.SetResellerID(resellerID)
	return o
}

// SetResellerID adds the resellerId to the get billing profileitems params
func (o *GetBillingProfileitemsParams) SetResellerID(resellerID string) {
	o.ResellerID = resellerID
}

// WithRows adds the rows to the get billing profileitems params
func (o *GetBillingProfileitemsParams) WithRows(rows string) *GetBillingProfileitemsParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the get billing profileitems params
func (o *GetBillingProfileitemsParams) SetRows(rows string) {
	o.Rows = rows
}

// WriteToRequest writes these params to a swagger request
func (o *GetBillingProfileitemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param handle
	qrHandle := o.Handle
	qHandle := qrHandle
	if qHandle != "" {
		if err := r.SetQueryParam("handle", qHandle); err != nil {
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
