// Code generated by go-swagger; DO NOT EDIT.

package customercontacts

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

// NewGetCustomerContactitemsParams creates a new GetCustomerContactitemsParams object
// with the default values initialized.
func NewGetCustomerContactitemsParams() *GetCustomerContactitemsParams {
	var ()
	return &GetCustomerContactitemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomerContactitemsParamsWithTimeout creates a new GetCustomerContactitemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCustomerContactitemsParamsWithTimeout(timeout time.Duration) *GetCustomerContactitemsParams {
	var ()
	return &GetCustomerContactitemsParams{

		timeout: timeout,
	}
}

// NewGetCustomerContactitemsParamsWithContext creates a new GetCustomerContactitemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCustomerContactitemsParamsWithContext(ctx context.Context) *GetCustomerContactitemsParams {
	var ()
	return &GetCustomerContactitemsParams{

		Context: ctx,
	}
}

// NewGetCustomerContactitemsParamsWithHTTPClient creates a new GetCustomerContactitemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCustomerContactitemsParamsWithHTTPClient(client *http.Client) *GetCustomerContactitemsParams {
	var ()
	return &GetCustomerContactitemsParams{
		HTTPClient: client,
	}
}

/*GetCustomerContactitemsParams contains all the parameters to send to the API endpoint
for the get customer contactitems operation typically these are written to a http.Request
*/
type GetCustomerContactitemsParams struct {

	/*Email
	  Filter for contacts matching an email pattern

	*/
	Email string
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
	  Filter for contacts belonging to a specific reseller

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

// WithTimeout adds the timeout to the get customer contactitems params
func (o *GetCustomerContactitemsParams) WithTimeout(timeout time.Duration) *GetCustomerContactitemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get customer contactitems params
func (o *GetCustomerContactitemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get customer contactitems params
func (o *GetCustomerContactitemsParams) WithContext(ctx context.Context) *GetCustomerContactitemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get customer contactitems params
func (o *GetCustomerContactitemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get customer contactitems params
func (o *GetCustomerContactitemsParams) WithHTTPClient(client *http.Client) *GetCustomerContactitemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get customer contactitems params
func (o *GetCustomerContactitemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmail adds the email to the get customer contactitems params
func (o *GetCustomerContactitemsParams) WithEmail(email string) *GetCustomerContactitemsParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the get customer contactitems params
func (o *GetCustomerContactitemsParams) SetEmail(email string) {
	o.Email = email
}

// WithOrderBy adds the orderBy to the get customer contactitems params
func (o *GetCustomerContactitemsParams) WithOrderBy(orderBy string) *GetCustomerContactitemsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get customer contactitems params
func (o *GetCustomerContactitemsParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the get customer contactitems params
func (o *GetCustomerContactitemsParams) WithOrderByDirection(orderByDirection string) *GetCustomerContactitemsParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the get customer contactitems params
func (o *GetCustomerContactitemsParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the get customer contactitems params
func (o *GetCustomerContactitemsParams) WithPage(page string) *GetCustomerContactitemsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get customer contactitems params
func (o *GetCustomerContactitemsParams) SetPage(page string) {
	o.Page = page
}

// WithResellerID adds the resellerID to the get customer contactitems params
func (o *GetCustomerContactitemsParams) WithResellerID(resellerID string) *GetCustomerContactitemsParams {
	o.SetResellerID(resellerID)
	return o
}

// SetResellerID adds the resellerId to the get customer contactitems params
func (o *GetCustomerContactitemsParams) SetResellerID(resellerID string) {
	o.ResellerID = resellerID
}

// WithRows adds the rows to the get customer contactitems params
func (o *GetCustomerContactitemsParams) WithRows(rows string) *GetCustomerContactitemsParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the get customer contactitems params
func (o *GetCustomerContactitemsParams) SetRows(rows string) {
	o.Rows = rows
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomerContactitemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param email
	qrEmail := o.Email
	qEmail := qrEmail
	if qEmail != "" {
		if err := r.SetQueryParam("email", qEmail); err != nil {
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
