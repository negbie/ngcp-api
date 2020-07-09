// Code generated by go-swagger; DO NOT EDIT.

package billingzones

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

// NewGetBillingZoneitemsParams creates a new GetBillingZoneitemsParams object
// with the default values initialized.
func NewGetBillingZoneitemsParams() *GetBillingZoneitemsParams {
	var ()
	return &GetBillingZoneitemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBillingZoneitemsParamsWithTimeout creates a new GetBillingZoneitemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBillingZoneitemsParamsWithTimeout(timeout time.Duration) *GetBillingZoneitemsParams {
	var ()
	return &GetBillingZoneitemsParams{

		timeout: timeout,
	}
}

// NewGetBillingZoneitemsParamsWithContext creates a new GetBillingZoneitemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBillingZoneitemsParamsWithContext(ctx context.Context) *GetBillingZoneitemsParams {
	var ()
	return &GetBillingZoneitemsParams{

		Context: ctx,
	}
}

// NewGetBillingZoneitemsParamsWithHTTPClient creates a new GetBillingZoneitemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBillingZoneitemsParamsWithHTTPClient(client *http.Client) *GetBillingZoneitemsParams {
	var ()
	return &GetBillingZoneitemsParams{
		HTTPClient: client,
	}
}

/*GetBillingZoneitemsParams contains all the parameters to send to the API endpoint
for the get billing zoneitems operation typically these are written to a http.Request
*/
type GetBillingZoneitemsParams struct {

	/*BillingProfileID
	  Filter for zones belonging to a specific billing profile

	*/
	BillingProfileID string
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
	/*Zone
	  Filter for zone name

	*/
	Zone string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get billing zoneitems params
func (o *GetBillingZoneitemsParams) WithTimeout(timeout time.Duration) *GetBillingZoneitemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get billing zoneitems params
func (o *GetBillingZoneitemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get billing zoneitems params
func (o *GetBillingZoneitemsParams) WithContext(ctx context.Context) *GetBillingZoneitemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get billing zoneitems params
func (o *GetBillingZoneitemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get billing zoneitems params
func (o *GetBillingZoneitemsParams) WithHTTPClient(client *http.Client) *GetBillingZoneitemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get billing zoneitems params
func (o *GetBillingZoneitemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBillingProfileID adds the billingProfileID to the get billing zoneitems params
func (o *GetBillingZoneitemsParams) WithBillingProfileID(billingProfileID string) *GetBillingZoneitemsParams {
	o.SetBillingProfileID(billingProfileID)
	return o
}

// SetBillingProfileID adds the billingProfileId to the get billing zoneitems params
func (o *GetBillingZoneitemsParams) SetBillingProfileID(billingProfileID string) {
	o.BillingProfileID = billingProfileID
}

// WithOrderBy adds the orderBy to the get billing zoneitems params
func (o *GetBillingZoneitemsParams) WithOrderBy(orderBy string) *GetBillingZoneitemsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get billing zoneitems params
func (o *GetBillingZoneitemsParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the get billing zoneitems params
func (o *GetBillingZoneitemsParams) WithOrderByDirection(orderByDirection string) *GetBillingZoneitemsParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the get billing zoneitems params
func (o *GetBillingZoneitemsParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the get billing zoneitems params
func (o *GetBillingZoneitemsParams) WithPage(page string) *GetBillingZoneitemsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get billing zoneitems params
func (o *GetBillingZoneitemsParams) SetPage(page string) {
	o.Page = page
}

// WithRows adds the rows to the get billing zoneitems params
func (o *GetBillingZoneitemsParams) WithRows(rows string) *GetBillingZoneitemsParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the get billing zoneitems params
func (o *GetBillingZoneitemsParams) SetRows(rows string) {
	o.Rows = rows
}

// WithZone adds the zone to the get billing zoneitems params
func (o *GetBillingZoneitemsParams) WithZone(zone string) *GetBillingZoneitemsParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the get billing zoneitems params
func (o *GetBillingZoneitemsParams) SetZone(zone string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *GetBillingZoneitemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param billing_profile_id
	qrBillingProfileID := o.BillingProfileID
	qBillingProfileID := qrBillingProfileID
	if qBillingProfileID != "" {
		if err := r.SetQueryParam("billing_profile_id", qBillingProfileID); err != nil {
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

	// query param zone
	qrZone := o.Zone
	qZone := qrZone
	if qZone != "" {
		if err := r.SetQueryParam("zone", qZone); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
