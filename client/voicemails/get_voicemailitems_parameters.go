// Code generated by go-swagger; DO NOT EDIT.

package voicemails

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

// NewGetVoicemailitemsParams creates a new GetVoicemailitemsParams object
// with the default values initialized.
func NewGetVoicemailitemsParams() *GetVoicemailitemsParams {
	var ()
	return &GetVoicemailitemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVoicemailitemsParamsWithTimeout creates a new GetVoicemailitemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVoicemailitemsParamsWithTimeout(timeout time.Duration) *GetVoicemailitemsParams {
	var ()
	return &GetVoicemailitemsParams{

		timeout: timeout,
	}
}

// NewGetVoicemailitemsParamsWithContext creates a new GetVoicemailitemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVoicemailitemsParamsWithContext(ctx context.Context) *GetVoicemailitemsParams {
	var ()
	return &GetVoicemailitemsParams{

		Context: ctx,
	}
}

// NewGetVoicemailitemsParamsWithHTTPClient creates a new GetVoicemailitemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVoicemailitemsParamsWithHTTPClient(client *http.Client) *GetVoicemailitemsParams {
	var ()
	return &GetVoicemailitemsParams{
		HTTPClient: client,
	}
}

/*GetVoicemailitemsParams contains all the parameters to send to the API endpoint
for the get voicemailitems operation typically these are written to a http.Request
*/
type GetVoicemailitemsParams struct {

	/*Folder
	  Filter for voicemails in a specific folder (one of INBOX, Old, Friends, Family, Cust1 to Cust4)

	*/
	Folder string
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
	  Filter for voicemails belonging to a specific subscriber

	*/
	SubscriberID string
	/*Tz
	  Format timestamp according to the optional time zone provided here, e.g. Europe/Berlin.

	*/
	Tz string
	/*UseOwnerTz
	  Format timestamp according to the filtered customer's/subscribers's inherited time zone.

	*/
	UseOwnerTz string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get voicemailitems params
func (o *GetVoicemailitemsParams) WithTimeout(timeout time.Duration) *GetVoicemailitemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get voicemailitems params
func (o *GetVoicemailitemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get voicemailitems params
func (o *GetVoicemailitemsParams) WithContext(ctx context.Context) *GetVoicemailitemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get voicemailitems params
func (o *GetVoicemailitemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get voicemailitems params
func (o *GetVoicemailitemsParams) WithHTTPClient(client *http.Client) *GetVoicemailitemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get voicemailitems params
func (o *GetVoicemailitemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFolder adds the folder to the get voicemailitems params
func (o *GetVoicemailitemsParams) WithFolder(folder string) *GetVoicemailitemsParams {
	o.SetFolder(folder)
	return o
}

// SetFolder adds the folder to the get voicemailitems params
func (o *GetVoicemailitemsParams) SetFolder(folder string) {
	o.Folder = folder
}

// WithOrderBy adds the orderBy to the get voicemailitems params
func (o *GetVoicemailitemsParams) WithOrderBy(orderBy string) *GetVoicemailitemsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get voicemailitems params
func (o *GetVoicemailitemsParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the get voicemailitems params
func (o *GetVoicemailitemsParams) WithOrderByDirection(orderByDirection string) *GetVoicemailitemsParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the get voicemailitems params
func (o *GetVoicemailitemsParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the get voicemailitems params
func (o *GetVoicemailitemsParams) WithPage(page string) *GetVoicemailitemsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get voicemailitems params
func (o *GetVoicemailitemsParams) SetPage(page string) {
	o.Page = page
}

// WithRows adds the rows to the get voicemailitems params
func (o *GetVoicemailitemsParams) WithRows(rows string) *GetVoicemailitemsParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the get voicemailitems params
func (o *GetVoicemailitemsParams) SetRows(rows string) {
	o.Rows = rows
}

// WithSubscriberID adds the subscriberID to the get voicemailitems params
func (o *GetVoicemailitemsParams) WithSubscriberID(subscriberID string) *GetVoicemailitemsParams {
	o.SetSubscriberID(subscriberID)
	return o
}

// SetSubscriberID adds the subscriberId to the get voicemailitems params
func (o *GetVoicemailitemsParams) SetSubscriberID(subscriberID string) {
	o.SubscriberID = subscriberID
}

// WithTz adds the tz to the get voicemailitems params
func (o *GetVoicemailitemsParams) WithTz(tz string) *GetVoicemailitemsParams {
	o.SetTz(tz)
	return o
}

// SetTz adds the tz to the get voicemailitems params
func (o *GetVoicemailitemsParams) SetTz(tz string) {
	o.Tz = tz
}

// WithUseOwnerTz adds the useOwnerTz to the get voicemailitems params
func (o *GetVoicemailitemsParams) WithUseOwnerTz(useOwnerTz string) *GetVoicemailitemsParams {
	o.SetUseOwnerTz(useOwnerTz)
	return o
}

// SetUseOwnerTz adds the useOwnerTz to the get voicemailitems params
func (o *GetVoicemailitemsParams) SetUseOwnerTz(useOwnerTz string) {
	o.UseOwnerTz = useOwnerTz
}

// WriteToRequest writes these params to a swagger request
func (o *GetVoicemailitemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param folder
	qrFolder := o.Folder
	qFolder := qrFolder
	if qFolder != "" {
		if err := r.SetQueryParam("folder", qFolder); err != nil {
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

	// query param tz
	qrTz := o.Tz
	qTz := qrTz
	if qTz != "" {
		if err := r.SetQueryParam("tz", qTz); err != nil {
			return err
		}
	}

	// query param use_owner_tz
	qrUseOwnerTz := o.UseOwnerTz
	qUseOwnerTz := qrUseOwnerTz
	if qUseOwnerTz != "" {
		if err := r.SetQueryParam("use_owner_tz", qUseOwnerTz); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
