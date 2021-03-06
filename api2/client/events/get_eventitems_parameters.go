// Code generated by go-swagger; DO NOT EDIT.

package events

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

// NewGetEventitemsParams creates a new GetEventitemsParams object
// with the default values initialized.
func NewGetEventitemsParams() *GetEventitemsParams {
	var ()
	return &GetEventitemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEventitemsParamsWithTimeout creates a new GetEventitemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEventitemsParamsWithTimeout(timeout time.Duration) *GetEventitemsParams {
	var ()
	return &GetEventitemsParams{

		timeout: timeout,
	}
}

// NewGetEventitemsParamsWithContext creates a new GetEventitemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEventitemsParamsWithContext(ctx context.Context) *GetEventitemsParams {
	var ()
	return &GetEventitemsParams{

		Context: ctx,
	}
}

// NewGetEventitemsParamsWithHTTPClient creates a new GetEventitemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEventitemsParamsWithHTTPClient(client *http.Client) *GetEventitemsParams {
	var ()
	return &GetEventitemsParams{
		HTTPClient: client,
	}
}

/*GetEventitemsParams contains all the parameters to send to the API endpoint
for the get eventitems operation typically these are written to a http.Request
*/
type GetEventitemsParams struct {

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
	  Filter for events for customers/subscribers of a specific reseller.

	*/
	ResellerID string
	/*Rows
	  Number of rows in one pagination page (default: 10)

	*/
	Rows string
	/*SubscriberID
	  Filter for events of a specific subscriber.

	*/
	SubscriberID string
	/*TimestampFrom
	  Filter for events occurred after or at the given time stamp.

	*/
	TimestampFrom string
	/*TimestampTo
	  Filter for events occurred before or at the given time stamp.

	*/
	TimestampTo string
	/*Type
	  Filter for events of a specific type.

	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get eventitems params
func (o *GetEventitemsParams) WithTimeout(timeout time.Duration) *GetEventitemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get eventitems params
func (o *GetEventitemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get eventitems params
func (o *GetEventitemsParams) WithContext(ctx context.Context) *GetEventitemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get eventitems params
func (o *GetEventitemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get eventitems params
func (o *GetEventitemsParams) WithHTTPClient(client *http.Client) *GetEventitemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get eventitems params
func (o *GetEventitemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderBy adds the orderBy to the get eventitems params
func (o *GetEventitemsParams) WithOrderBy(orderBy string) *GetEventitemsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get eventitems params
func (o *GetEventitemsParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the get eventitems params
func (o *GetEventitemsParams) WithOrderByDirection(orderByDirection string) *GetEventitemsParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the get eventitems params
func (o *GetEventitemsParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the get eventitems params
func (o *GetEventitemsParams) WithPage(page string) *GetEventitemsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get eventitems params
func (o *GetEventitemsParams) SetPage(page string) {
	o.Page = page
}

// WithResellerID adds the resellerID to the get eventitems params
func (o *GetEventitemsParams) WithResellerID(resellerID string) *GetEventitemsParams {
	o.SetResellerID(resellerID)
	return o
}

// SetResellerID adds the resellerId to the get eventitems params
func (o *GetEventitemsParams) SetResellerID(resellerID string) {
	o.ResellerID = resellerID
}

// WithRows adds the rows to the get eventitems params
func (o *GetEventitemsParams) WithRows(rows string) *GetEventitemsParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the get eventitems params
func (o *GetEventitemsParams) SetRows(rows string) {
	o.Rows = rows
}

// WithSubscriberID adds the subscriberID to the get eventitems params
func (o *GetEventitemsParams) WithSubscriberID(subscriberID string) *GetEventitemsParams {
	o.SetSubscriberID(subscriberID)
	return o
}

// SetSubscriberID adds the subscriberId to the get eventitems params
func (o *GetEventitemsParams) SetSubscriberID(subscriberID string) {
	o.SubscriberID = subscriberID
}

// WithTimestampFrom adds the timestampFrom to the get eventitems params
func (o *GetEventitemsParams) WithTimestampFrom(timestampFrom string) *GetEventitemsParams {
	o.SetTimestampFrom(timestampFrom)
	return o
}

// SetTimestampFrom adds the timestampFrom to the get eventitems params
func (o *GetEventitemsParams) SetTimestampFrom(timestampFrom string) {
	o.TimestampFrom = timestampFrom
}

// WithTimestampTo adds the timestampTo to the get eventitems params
func (o *GetEventitemsParams) WithTimestampTo(timestampTo string) *GetEventitemsParams {
	o.SetTimestampTo(timestampTo)
	return o
}

// SetTimestampTo adds the timestampTo to the get eventitems params
func (o *GetEventitemsParams) SetTimestampTo(timestampTo string) {
	o.TimestampTo = timestampTo
}

// WithType adds the typeVar to the get eventitems params
func (o *GetEventitemsParams) WithType(typeVar string) *GetEventitemsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get eventitems params
func (o *GetEventitemsParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetEventitemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// query param timestamp_from
	qrTimestampFrom := o.TimestampFrom
	qTimestampFrom := qrTimestampFrom
	if qTimestampFrom != "" {
		if err := r.SetQueryParam("timestamp_from", qTimestampFrom); err != nil {
			return err
		}
	}

	// query param timestamp_to
	qrTimestampTo := o.TimestampTo
	qTimestampTo := qrTimestampTo
	if qTimestampTo != "" {
		if err := r.SetQueryParam("timestamp_to", qTimestampTo); err != nil {
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
