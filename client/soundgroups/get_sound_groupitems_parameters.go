// Code generated by go-swagger; DO NOT EDIT.

package soundgroups

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

// NewGetSoundGroupitemsParams creates a new GetSoundGroupitemsParams object
// with the default values initialized.
func NewGetSoundGroupitemsParams() *GetSoundGroupitemsParams {
	var ()
	return &GetSoundGroupitemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSoundGroupitemsParamsWithTimeout creates a new GetSoundGroupitemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSoundGroupitemsParamsWithTimeout(timeout time.Duration) *GetSoundGroupitemsParams {
	var ()
	return &GetSoundGroupitemsParams{

		timeout: timeout,
	}
}

// NewGetSoundGroupitemsParamsWithContext creates a new GetSoundGroupitemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSoundGroupitemsParamsWithContext(ctx context.Context) *GetSoundGroupitemsParams {
	var ()
	return &GetSoundGroupitemsParams{

		Context: ctx,
	}
}

// NewGetSoundGroupitemsParamsWithHTTPClient creates a new GetSoundGroupitemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSoundGroupitemsParamsWithHTTPClient(client *http.Client) *GetSoundGroupitemsParams {
	var ()
	return &GetSoundGroupitemsParams{
		HTTPClient: client,
	}
}

/*GetSoundGroupitemsParams contains all the parameters to send to the API endpoint
for the get sound groupitems operation typically these are written to a http.Request
*/
type GetSoundGroupitemsParams struct {

	/*Name
	  Filter for sound groups with a specific name

	*/
	Name string
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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sound groupitems params
func (o *GetSoundGroupitemsParams) WithTimeout(timeout time.Duration) *GetSoundGroupitemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sound groupitems params
func (o *GetSoundGroupitemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sound groupitems params
func (o *GetSoundGroupitemsParams) WithContext(ctx context.Context) *GetSoundGroupitemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sound groupitems params
func (o *GetSoundGroupitemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sound groupitems params
func (o *GetSoundGroupitemsParams) WithHTTPClient(client *http.Client) *GetSoundGroupitemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sound groupitems params
func (o *GetSoundGroupitemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get sound groupitems params
func (o *GetSoundGroupitemsParams) WithName(name string) *GetSoundGroupitemsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get sound groupitems params
func (o *GetSoundGroupitemsParams) SetName(name string) {
	o.Name = name
}

// WithOrderBy adds the orderBy to the get sound groupitems params
func (o *GetSoundGroupitemsParams) WithOrderBy(orderBy string) *GetSoundGroupitemsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get sound groupitems params
func (o *GetSoundGroupitemsParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the get sound groupitems params
func (o *GetSoundGroupitemsParams) WithOrderByDirection(orderByDirection string) *GetSoundGroupitemsParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the get sound groupitems params
func (o *GetSoundGroupitemsParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the get sound groupitems params
func (o *GetSoundGroupitemsParams) WithPage(page string) *GetSoundGroupitemsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get sound groupitems params
func (o *GetSoundGroupitemsParams) SetPage(page string) {
	o.Page = page
}

// WithRows adds the rows to the get sound groupitems params
func (o *GetSoundGroupitemsParams) WithRows(rows string) *GetSoundGroupitemsParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the get sound groupitems params
func (o *GetSoundGroupitemsParams) SetRows(rows string) {
	o.Rows = rows
}

// WriteToRequest writes these params to a swagger request
func (o *GetSoundGroupitemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param name
	qrName := o.Name
	qName := qrName
	if qName != "" {
		if err := r.SetQueryParam("name", qName); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
