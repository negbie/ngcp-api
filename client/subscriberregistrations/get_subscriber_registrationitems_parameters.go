// Code generated by go-swagger; DO NOT EDIT.

package subscriberregistrations

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

// NewGetSubscriberRegistrationitemsParams creates a new GetSubscriberRegistrationitemsParams object
// with the default values initialized.
func NewGetSubscriberRegistrationitemsParams() *GetSubscriberRegistrationitemsParams {
	var ()
	return &GetSubscriberRegistrationitemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSubscriberRegistrationitemsParamsWithTimeout creates a new GetSubscriberRegistrationitemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSubscriberRegistrationitemsParamsWithTimeout(timeout time.Duration) *GetSubscriberRegistrationitemsParams {
	var ()
	return &GetSubscriberRegistrationitemsParams{

		timeout: timeout,
	}
}

// NewGetSubscriberRegistrationitemsParamsWithContext creates a new GetSubscriberRegistrationitemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSubscriberRegistrationitemsParamsWithContext(ctx context.Context) *GetSubscriberRegistrationitemsParams {
	var ()
	return &GetSubscriberRegistrationitemsParams{

		Context: ctx,
	}
}

// NewGetSubscriberRegistrationitemsParamsWithHTTPClient creates a new GetSubscriberRegistrationitemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSubscriberRegistrationitemsParamsWithHTTPClient(client *http.Client) *GetSubscriberRegistrationitemsParams {
	var ()
	return &GetSubscriberRegistrationitemsParams{
		HTTPClient: client,
	}
}

/*GetSubscriberRegistrationitemsParams contains all the parameters to send to the API endpoint
for the get subscriber registrationitems operation typically these are written to a http.Request
*/
type GetSubscriberRegistrationitemsParams struct {

	/*Page
	  Pagination page which should be displayed (default: 1)

	*/
	Page string
	/*Rows
	  Number of rows in one pagination page (default: 10)

	*/
	Rows string
	/*SubscriberID
	  Filter for registrations of a specific subscriber

	*/
	SubscriberID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get subscriber registrationitems params
func (o *GetSubscriberRegistrationitemsParams) WithTimeout(timeout time.Duration) *GetSubscriberRegistrationitemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get subscriber registrationitems params
func (o *GetSubscriberRegistrationitemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get subscriber registrationitems params
func (o *GetSubscriberRegistrationitemsParams) WithContext(ctx context.Context) *GetSubscriberRegistrationitemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get subscriber registrationitems params
func (o *GetSubscriberRegistrationitemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get subscriber registrationitems params
func (o *GetSubscriberRegistrationitemsParams) WithHTTPClient(client *http.Client) *GetSubscriberRegistrationitemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get subscriber registrationitems params
func (o *GetSubscriberRegistrationitemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get subscriber registrationitems params
func (o *GetSubscriberRegistrationitemsParams) WithPage(page string) *GetSubscriberRegistrationitemsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get subscriber registrationitems params
func (o *GetSubscriberRegistrationitemsParams) SetPage(page string) {
	o.Page = page
}

// WithRows adds the rows to the get subscriber registrationitems params
func (o *GetSubscriberRegistrationitemsParams) WithRows(rows string) *GetSubscriberRegistrationitemsParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the get subscriber registrationitems params
func (o *GetSubscriberRegistrationitemsParams) SetRows(rows string) {
	o.Rows = rows
}

// WithSubscriberID adds the subscriberID to the get subscriber registrationitems params
func (o *GetSubscriberRegistrationitemsParams) WithSubscriberID(subscriberID string) *GetSubscriberRegistrationitemsParams {
	o.SetSubscriberID(subscriberID)
	return o
}

// SetSubscriberID adds the subscriberId to the get subscriber registrationitems params
func (o *GetSubscriberRegistrationitemsParams) SetSubscriberID(subscriberID string) {
	o.SubscriberID = subscriberID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSubscriberRegistrationitemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
