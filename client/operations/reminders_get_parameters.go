// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewRemindersGetParams creates a new RemindersGetParams object
// with the default values initialized.
func NewRemindersGetParams() *RemindersGetParams {
	var ()
	return &RemindersGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemindersGetParamsWithTimeout creates a new RemindersGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemindersGetParamsWithTimeout(timeout time.Duration) *RemindersGetParams {
	var ()
	return &RemindersGetParams{

		timeout: timeout,
	}
}

// NewRemindersGetParamsWithContext creates a new RemindersGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemindersGetParamsWithContext(ctx context.Context) *RemindersGetParams {
	var ()
	return &RemindersGetParams{

		Context: ctx,
	}
}

// NewRemindersGetParamsWithHTTPClient creates a new RemindersGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemindersGetParamsWithHTTPClient(client *http.Client) *RemindersGetParams {
	var ()
	return &RemindersGetParams{
		HTTPClient: client,
	}
}

/*RemindersGetParams contains all the parameters to send to the API endpoint
for the reminders get operation typically these are written to a http.Request
*/
type RemindersGetParams struct {

	/*Active
	  Filter for active or inactive reminders (0|1)

	*/
	Active string
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
	  Filter for reminders of a specific subscriber

	*/
	SubscriberID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the reminders get params
func (o *RemindersGetParams) WithTimeout(timeout time.Duration) *RemindersGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reminders get params
func (o *RemindersGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reminders get params
func (o *RemindersGetParams) WithContext(ctx context.Context) *RemindersGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reminders get params
func (o *RemindersGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reminders get params
func (o *RemindersGetParams) WithHTTPClient(client *http.Client) *RemindersGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reminders get params
func (o *RemindersGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActive adds the active to the reminders get params
func (o *RemindersGetParams) WithActive(active string) *RemindersGetParams {
	o.SetActive(active)
	return o
}

// SetActive adds the active to the reminders get params
func (o *RemindersGetParams) SetActive(active string) {
	o.Active = active
}

// WithOrderBy adds the orderBy to the reminders get params
func (o *RemindersGetParams) WithOrderBy(orderBy string) *RemindersGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the reminders get params
func (o *RemindersGetParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the reminders get params
func (o *RemindersGetParams) WithOrderByDirection(orderByDirection string) *RemindersGetParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the reminders get params
func (o *RemindersGetParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the reminders get params
func (o *RemindersGetParams) WithPage(page string) *RemindersGetParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the reminders get params
func (o *RemindersGetParams) SetPage(page string) {
	o.Page = page
}

// WithRows adds the rows to the reminders get params
func (o *RemindersGetParams) WithRows(rows string) *RemindersGetParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the reminders get params
func (o *RemindersGetParams) SetRows(rows string) {
	o.Rows = rows
}

// WithSubscriberID adds the subscriberID to the reminders get params
func (o *RemindersGetParams) WithSubscriberID(subscriberID string) *RemindersGetParams {
	o.SetSubscriberID(subscriberID)
	return o
}

// SetSubscriberID adds the subscriberId to the reminders get params
func (o *RemindersGetParams) SetSubscriberID(subscriberID string) {
	o.SubscriberID = subscriberID
}

// WriteToRequest writes these params to a swagger request
func (o *RemindersGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param active
	qrActive := o.Active
	qActive := qrActive
	if qActive != "" {
		if err := r.SetQueryParam("active", qActive); err != nil {
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