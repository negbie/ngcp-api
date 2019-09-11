// Code generated by go-swagger; DO NOT EDIT.

package numbers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewNumbersGetParams creates a new NumbersGetParams object
// with the default values initialized.
func NewNumbersGetParams() *NumbersGetParams {
	var ()
	return &NumbersGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNumbersGetParamsWithTimeout creates a new NumbersGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNumbersGetParamsWithTimeout(timeout time.Duration) *NumbersGetParams {
	var ()
	return &NumbersGetParams{

		timeout: timeout,
	}
}

// NewNumbersGetParamsWithContext creates a new NumbersGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewNumbersGetParamsWithContext(ctx context.Context) *NumbersGetParams {
	var ()
	return &NumbersGetParams{

		Context: ctx,
	}
}

// NewNumbersGetParamsWithHTTPClient creates a new NumbersGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNumbersGetParamsWithHTTPClient(client *http.Client) *NumbersGetParams {
	var ()
	return &NumbersGetParams{
		HTTPClient: client,
	}
}

/*NumbersGetParams contains all the parameters to send to the API endpoint
for the numbers get operation typically these are written to a http.Request
*/
type NumbersGetParams struct {

	/*CustomerID
	  Filter for numbers assigned to subscribers of a specific customer.

	*/
	CustomerID string
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
	  Filter for numbers assigned to subscribers belonging to a specific reseller

	*/
	ResellerID string
	/*Rows
	  Number of rows in one pagination page (default: 10)

	*/
	Rows string
	/*SubscriberID
	  Filter for numbers assigned to a specific subscriber.

	*/
	SubscriberID string
	/*Type
	  Filter for number type, either "primary" or "alias".

	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the numbers get params
func (o *NumbersGetParams) WithTimeout(timeout time.Duration) *NumbersGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the numbers get params
func (o *NumbersGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the numbers get params
func (o *NumbersGetParams) WithContext(ctx context.Context) *NumbersGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the numbers get params
func (o *NumbersGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the numbers get params
func (o *NumbersGetParams) WithHTTPClient(client *http.Client) *NumbersGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the numbers get params
func (o *NumbersGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the numbers get params
func (o *NumbersGetParams) WithCustomerID(customerID string) *NumbersGetParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the numbers get params
func (o *NumbersGetParams) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WithOrderBy adds the orderBy to the numbers get params
func (o *NumbersGetParams) WithOrderBy(orderBy string) *NumbersGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the numbers get params
func (o *NumbersGetParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the numbers get params
func (o *NumbersGetParams) WithOrderByDirection(orderByDirection string) *NumbersGetParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the numbers get params
func (o *NumbersGetParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the numbers get params
func (o *NumbersGetParams) WithPage(page string) *NumbersGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the numbers get params
func (o *NumbersGetParams) SetPage(page string) {
	o.Page = page
}

// WithResellerID adds the resellerID to the numbers get params
func (o *NumbersGetParams) WithResellerID(resellerID string) *NumbersGetParams {
	o.SetResellerID(resellerID)
	return o
}

// SetResellerID adds the resellerId to the numbers get params
func (o *NumbersGetParams) SetResellerID(resellerID string) {
	o.ResellerID = resellerID
}

// WithRows adds the rows to the numbers get params
func (o *NumbersGetParams) WithRows(rows string) *NumbersGetParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the numbers get params
func (o *NumbersGetParams) SetRows(rows string) {
	o.Rows = rows
}

// WithSubscriberID adds the subscriberID to the numbers get params
func (o *NumbersGetParams) WithSubscriberID(subscriberID string) *NumbersGetParams {
	o.SetSubscriberID(subscriberID)
	return o
}

// SetSubscriberID adds the subscriberId to the numbers get params
func (o *NumbersGetParams) SetSubscriberID(subscriberID string) {
	o.SubscriberID = subscriberID
}

// WithType adds the typeVar to the numbers get params
func (o *NumbersGetParams) WithType(typeVar string) *NumbersGetParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the numbers get params
func (o *NumbersGetParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *NumbersGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param customer_id
	qrCustomerID := o.CustomerID
	qCustomerID := qrCustomerID
	if qCustomerID != "" {
		if err := r.SetQueryParam("customer_id", qCustomerID); err != nil {
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

	// query param subscriber_id
	qrSubscriberID := o.SubscriberID
	qSubscriberID := qrSubscriberID
	if qSubscriberID != "" {
		if err := r.SetQueryParam("subscriber_id", qSubscriberID); err != nil {
			return err
		}
	}

	// query param type
	qrType := o.Type
	qType := qrType
	if qType != "" {
		if err := r.SetQueryParam("type", qType); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}