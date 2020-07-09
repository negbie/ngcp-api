// Code generated by go-swagger; DO NOT EDIT.

package maliciouscalls

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

// NewGetMaliciousCallitemsParams creates a new GetMaliciousCallitemsParams object
// with the default values initialized.
func NewGetMaliciousCallitemsParams() *GetMaliciousCallitemsParams {
	var ()
	return &GetMaliciousCallitemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMaliciousCallitemsParamsWithTimeout creates a new GetMaliciousCallitemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMaliciousCallitemsParamsWithTimeout(timeout time.Duration) *GetMaliciousCallitemsParams {
	var ()
	return &GetMaliciousCallitemsParams{

		timeout: timeout,
	}
}

// NewGetMaliciousCallitemsParamsWithContext creates a new GetMaliciousCallitemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMaliciousCallitemsParamsWithContext(ctx context.Context) *GetMaliciousCallitemsParams {
	var ()
	return &GetMaliciousCallitemsParams{

		Context: ctx,
	}
}

// NewGetMaliciousCallitemsParamsWithHTTPClient creates a new GetMaliciousCallitemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMaliciousCallitemsParamsWithHTTPClient(client *http.Client) *GetMaliciousCallitemsParams {
	var ()
	return &GetMaliciousCallitemsParams{
		HTTPClient: client,
	}
}

/*GetMaliciousCallitemsParams contains all the parameters to send to the API endpoint
for the get malicious callitems operation typically these are written to a http.Request
*/
type GetMaliciousCallitemsParams struct {

	/*Callee
	  Filter by the callee number

	*/
	Callee string
	/*Caller
	  Filter by the caller number

	*/
	Caller string
	/*Callid
	  Filter by the call id

	*/
	Callid string
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
	  Filter for malicious calls belonging to a specific reseller

	*/
	ResellerID string
	/*Rows
	  Number of rows in one pagination page (default: 10)

	*/
	Rows string
	/*StartGe
	  Filter by records with greater or equal the specified time stamp.

	*/
	StartGe string
	/*StartLe
	  Filter by records with lower or equal than the specified time stamp.

	*/
	StartLe string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get malicious callitems params
func (o *GetMaliciousCallitemsParams) WithTimeout(timeout time.Duration) *GetMaliciousCallitemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get malicious callitems params
func (o *GetMaliciousCallitemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get malicious callitems params
func (o *GetMaliciousCallitemsParams) WithContext(ctx context.Context) *GetMaliciousCallitemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get malicious callitems params
func (o *GetMaliciousCallitemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get malicious callitems params
func (o *GetMaliciousCallitemsParams) WithHTTPClient(client *http.Client) *GetMaliciousCallitemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get malicious callitems params
func (o *GetMaliciousCallitemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCallee adds the callee to the get malicious callitems params
func (o *GetMaliciousCallitemsParams) WithCallee(callee string) *GetMaliciousCallitemsParams {
	o.SetCallee(callee)
	return o
}

// SetCallee adds the callee to the get malicious callitems params
func (o *GetMaliciousCallitemsParams) SetCallee(callee string) {
	o.Callee = callee
}

// WithCaller adds the caller to the get malicious callitems params
func (o *GetMaliciousCallitemsParams) WithCaller(caller string) *GetMaliciousCallitemsParams {
	o.SetCaller(caller)
	return o
}

// SetCaller adds the caller to the get malicious callitems params
func (o *GetMaliciousCallitemsParams) SetCaller(caller string) {
	o.Caller = caller
}

// WithCallid adds the callid to the get malicious callitems params
func (o *GetMaliciousCallitemsParams) WithCallid(callid string) *GetMaliciousCallitemsParams {
	o.SetCallid(callid)
	return o
}

// SetCallid adds the callid to the get malicious callitems params
func (o *GetMaliciousCallitemsParams) SetCallid(callid string) {
	o.Callid = callid
}

// WithOrderBy adds the orderBy to the get malicious callitems params
func (o *GetMaliciousCallitemsParams) WithOrderBy(orderBy string) *GetMaliciousCallitemsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get malicious callitems params
func (o *GetMaliciousCallitemsParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the get malicious callitems params
func (o *GetMaliciousCallitemsParams) WithOrderByDirection(orderByDirection string) *GetMaliciousCallitemsParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the get malicious callitems params
func (o *GetMaliciousCallitemsParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the get malicious callitems params
func (o *GetMaliciousCallitemsParams) WithPage(page string) *GetMaliciousCallitemsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get malicious callitems params
func (o *GetMaliciousCallitemsParams) SetPage(page string) {
	o.Page = page
}

// WithResellerID adds the resellerID to the get malicious callitems params
func (o *GetMaliciousCallitemsParams) WithResellerID(resellerID string) *GetMaliciousCallitemsParams {
	o.SetResellerID(resellerID)
	return o
}

// SetResellerID adds the resellerId to the get malicious callitems params
func (o *GetMaliciousCallitemsParams) SetResellerID(resellerID string) {
	o.ResellerID = resellerID
}

// WithRows adds the rows to the get malicious callitems params
func (o *GetMaliciousCallitemsParams) WithRows(rows string) *GetMaliciousCallitemsParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the get malicious callitems params
func (o *GetMaliciousCallitemsParams) SetRows(rows string) {
	o.Rows = rows
}

// WithStartGe adds the startGe to the get malicious callitems params
func (o *GetMaliciousCallitemsParams) WithStartGe(startGe string) *GetMaliciousCallitemsParams {
	o.SetStartGe(startGe)
	return o
}

// SetStartGe adds the startGe to the get malicious callitems params
func (o *GetMaliciousCallitemsParams) SetStartGe(startGe string) {
	o.StartGe = startGe
}

// WithStartLe adds the startLe to the get malicious callitems params
func (o *GetMaliciousCallitemsParams) WithStartLe(startLe string) *GetMaliciousCallitemsParams {
	o.SetStartLe(startLe)
	return o
}

// SetStartLe adds the startLe to the get malicious callitems params
func (o *GetMaliciousCallitemsParams) SetStartLe(startLe string) {
	o.StartLe = startLe
}

// WriteToRequest writes these params to a swagger request
func (o *GetMaliciousCallitemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param callee
	qrCallee := o.Callee
	qCallee := qrCallee
	if qCallee != "" {
		if err := r.SetQueryParam("callee", qCallee); err != nil {
			return err
		}
	}

	// query param caller
	qrCaller := o.Caller
	qCaller := qrCaller
	if qCaller != "" {
		if err := r.SetQueryParam("caller", qCaller); err != nil {
			return err
		}
	}

	// query param callid
	qrCallid := o.Callid
	qCallid := qrCallid
	if qCallid != "" {
		if err := r.SetQueryParam("callid", qCallid); err != nil {
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

	// query param start_ge
	qrStartGe := o.StartGe
	qStartGe := qrStartGe
	if qStartGe != "" {
		if err := r.SetQueryParam("start_ge", qStartGe); err != nil {
			return err
		}
	}

	// query param start_le
	qrStartLe := o.StartLe
	qStartLe := qrStartLe
	if qStartLe != "" {
		if err := r.SetQueryParam("start_le", qStartLe); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
