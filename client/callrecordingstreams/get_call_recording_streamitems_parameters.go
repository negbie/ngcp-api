// Code generated by go-swagger; DO NOT EDIT.

package callrecordingstreams

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

// NewGetCallRecordingStreamitemsParams creates a new GetCallRecordingStreamitemsParams object
// with the default values initialized.
func NewGetCallRecordingStreamitemsParams() *GetCallRecordingStreamitemsParams {
	var ()
	return &GetCallRecordingStreamitemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCallRecordingStreamitemsParamsWithTimeout creates a new GetCallRecordingStreamitemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCallRecordingStreamitemsParamsWithTimeout(timeout time.Duration) *GetCallRecordingStreamitemsParams {
	var ()
	return &GetCallRecordingStreamitemsParams{

		timeout: timeout,
	}
}

// NewGetCallRecordingStreamitemsParamsWithContext creates a new GetCallRecordingStreamitemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCallRecordingStreamitemsParamsWithContext(ctx context.Context) *GetCallRecordingStreamitemsParams {
	var ()
	return &GetCallRecordingStreamitemsParams{

		Context: ctx,
	}
}

// NewGetCallRecordingStreamitemsParamsWithHTTPClient creates a new GetCallRecordingStreamitemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCallRecordingStreamitemsParamsWithHTTPClient(client *http.Client) *GetCallRecordingStreamitemsParams {
	var ()
	return &GetCallRecordingStreamitemsParams{
		HTTPClient: client,
	}
}

/*GetCallRecordingStreamitemsParams contains all the parameters to send to the API endpoint
for the get call recording streamitems operation typically these are written to a http.Request
*/
type GetCallRecordingStreamitemsParams struct {

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
	/*RecordingID
	  Filter for callrecording streams belonging to a specific recording session.

	*/
	RecordingID string
	/*Rows
	  Number of rows in one pagination page (default: 10)

	*/
	Rows string
	/*Type
	  Filter for callrecording streams with a specific type ("single" or "mixed")

	*/
	Type string
	/*Tz
	  Format start_time according to the optional time zone provided here, e.g. Europe/Berlin.

	*/
	Tz string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get call recording streamitems params
func (o *GetCallRecordingStreamitemsParams) WithTimeout(timeout time.Duration) *GetCallRecordingStreamitemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get call recording streamitems params
func (o *GetCallRecordingStreamitemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get call recording streamitems params
func (o *GetCallRecordingStreamitemsParams) WithContext(ctx context.Context) *GetCallRecordingStreamitemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get call recording streamitems params
func (o *GetCallRecordingStreamitemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get call recording streamitems params
func (o *GetCallRecordingStreamitemsParams) WithHTTPClient(client *http.Client) *GetCallRecordingStreamitemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get call recording streamitems params
func (o *GetCallRecordingStreamitemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderBy adds the orderBy to the get call recording streamitems params
func (o *GetCallRecordingStreamitemsParams) WithOrderBy(orderBy string) *GetCallRecordingStreamitemsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get call recording streamitems params
func (o *GetCallRecordingStreamitemsParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the get call recording streamitems params
func (o *GetCallRecordingStreamitemsParams) WithOrderByDirection(orderByDirection string) *GetCallRecordingStreamitemsParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the get call recording streamitems params
func (o *GetCallRecordingStreamitemsParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the get call recording streamitems params
func (o *GetCallRecordingStreamitemsParams) WithPage(page string) *GetCallRecordingStreamitemsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get call recording streamitems params
func (o *GetCallRecordingStreamitemsParams) SetPage(page string) {
	o.Page = page
}

// WithRecordingID adds the recordingID to the get call recording streamitems params
func (o *GetCallRecordingStreamitemsParams) WithRecordingID(recordingID string) *GetCallRecordingStreamitemsParams {
	o.SetRecordingID(recordingID)
	return o
}

// SetRecordingID adds the recordingId to the get call recording streamitems params
func (o *GetCallRecordingStreamitemsParams) SetRecordingID(recordingID string) {
	o.RecordingID = recordingID
}

// WithRows adds the rows to the get call recording streamitems params
func (o *GetCallRecordingStreamitemsParams) WithRows(rows string) *GetCallRecordingStreamitemsParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the get call recording streamitems params
func (o *GetCallRecordingStreamitemsParams) SetRows(rows string) {
	o.Rows = rows
}

// WithType adds the typeVar to the get call recording streamitems params
func (o *GetCallRecordingStreamitemsParams) WithType(typeVar string) *GetCallRecordingStreamitemsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get call recording streamitems params
func (o *GetCallRecordingStreamitemsParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WithTz adds the tz to the get call recording streamitems params
func (o *GetCallRecordingStreamitemsParams) WithTz(tz string) *GetCallRecordingStreamitemsParams {
	o.SetTz(tz)
	return o
}

// SetTz adds the tz to the get call recording streamitems params
func (o *GetCallRecordingStreamitemsParams) SetTz(tz string) {
	o.Tz = tz
}

// WriteToRequest writes these params to a swagger request
func (o *GetCallRecordingStreamitemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param recording_id
	qrRecordingID := o.RecordingID
	qRecordingID := qrRecordingID
	if qRecordingID != "" {
		if err := r.SetQueryParam("recording_id", qRecordingID); err != nil {
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

	// query param type
	qrType := o.Type
	qType := qrType
	if qType != "" {
		if err := r.SetQueryParam("type", qType); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}