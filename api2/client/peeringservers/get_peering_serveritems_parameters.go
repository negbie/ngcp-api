// Code generated by go-swagger; DO NOT EDIT.

package peeringservers

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

// NewGetPeeringServeritemsParams creates a new GetPeeringServeritemsParams object
// with the default values initialized.
func NewGetPeeringServeritemsParams() *GetPeeringServeritemsParams {
	var ()
	return &GetPeeringServeritemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPeeringServeritemsParamsWithTimeout creates a new GetPeeringServeritemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPeeringServeritemsParamsWithTimeout(timeout time.Duration) *GetPeeringServeritemsParams {
	var ()
	return &GetPeeringServeritemsParams{

		timeout: timeout,
	}
}

// NewGetPeeringServeritemsParamsWithContext creates a new GetPeeringServeritemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPeeringServeritemsParamsWithContext(ctx context.Context) *GetPeeringServeritemsParams {
	var ()
	return &GetPeeringServeritemsParams{

		Context: ctx,
	}
}

// NewGetPeeringServeritemsParamsWithHTTPClient creates a new GetPeeringServeritemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPeeringServeritemsParamsWithHTTPClient(client *http.Client) *GetPeeringServeritemsParams {
	var ()
	return &GetPeeringServeritemsParams{
		HTTPClient: client,
	}
}

/*GetPeeringServeritemsParams contains all the parameters to send to the API endpoint
for the get peering serveritems operation typically these are written to a http.Request
*/
type GetPeeringServeritemsParams struct {

	/*Enabled
	  Filter for peering server enabled flag

	*/
	Enabled string
	/*GroupID
	  Filter for peering server group

	*/
	GroupID string
	/*Host
	  Filter for peering server host

	*/
	Host string
	/*IP
	  Filter for peering server ip

	*/
	IP string
	/*Name
	  Filter for peering server name

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

// WithTimeout adds the timeout to the get peering serveritems params
func (o *GetPeeringServeritemsParams) WithTimeout(timeout time.Duration) *GetPeeringServeritemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get peering serveritems params
func (o *GetPeeringServeritemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get peering serveritems params
func (o *GetPeeringServeritemsParams) WithContext(ctx context.Context) *GetPeeringServeritemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get peering serveritems params
func (o *GetPeeringServeritemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get peering serveritems params
func (o *GetPeeringServeritemsParams) WithHTTPClient(client *http.Client) *GetPeeringServeritemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get peering serveritems params
func (o *GetPeeringServeritemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnabled adds the enabled to the get peering serveritems params
func (o *GetPeeringServeritemsParams) WithEnabled(enabled string) *GetPeeringServeritemsParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the get peering serveritems params
func (o *GetPeeringServeritemsParams) SetEnabled(enabled string) {
	o.Enabled = enabled
}

// WithGroupID adds the groupID to the get peering serveritems params
func (o *GetPeeringServeritemsParams) WithGroupID(groupID string) *GetPeeringServeritemsParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the get peering serveritems params
func (o *GetPeeringServeritemsParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithHost adds the host to the get peering serveritems params
func (o *GetPeeringServeritemsParams) WithHost(host string) *GetPeeringServeritemsParams {
	o.SetHost(host)
	return o
}

// SetHost adds the host to the get peering serveritems params
func (o *GetPeeringServeritemsParams) SetHost(host string) {
	o.Host = host
}

// WithIP adds the ip to the get peering serveritems params
func (o *GetPeeringServeritemsParams) WithIP(ip string) *GetPeeringServeritemsParams {
	o.SetIP(ip)
	return o
}

// SetIP adds the ip to the get peering serveritems params
func (o *GetPeeringServeritemsParams) SetIP(ip string) {
	o.IP = ip
}

// WithName adds the name to the get peering serveritems params
func (o *GetPeeringServeritemsParams) WithName(name string) *GetPeeringServeritemsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get peering serveritems params
func (o *GetPeeringServeritemsParams) SetName(name string) {
	o.Name = name
}

// WithOrderBy adds the orderBy to the get peering serveritems params
func (o *GetPeeringServeritemsParams) WithOrderBy(orderBy string) *GetPeeringServeritemsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get peering serveritems params
func (o *GetPeeringServeritemsParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the get peering serveritems params
func (o *GetPeeringServeritemsParams) WithOrderByDirection(orderByDirection string) *GetPeeringServeritemsParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the get peering serveritems params
func (o *GetPeeringServeritemsParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the get peering serveritems params
func (o *GetPeeringServeritemsParams) WithPage(page string) *GetPeeringServeritemsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get peering serveritems params
func (o *GetPeeringServeritemsParams) SetPage(page string) {
	o.Page = page
}

// WithRows adds the rows to the get peering serveritems params
func (o *GetPeeringServeritemsParams) WithRows(rows string) *GetPeeringServeritemsParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the get peering serveritems params
func (o *GetPeeringServeritemsParams) SetRows(rows string) {
	o.Rows = rows
}

// WriteToRequest writes these params to a swagger request
func (o *GetPeeringServeritemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param enabled
	qrEnabled := o.Enabled
	qEnabled := qrEnabled
	if qEnabled != "" {
		if err := r.SetQueryParam("enabled", qEnabled); err != nil {
			return err
		}
	}

	// query param group_id
	qrGroupID := o.GroupID
	qGroupID := qrGroupID
	if qGroupID != "" {
		if err := r.SetQueryParam("group_id", qGroupID); err != nil {
			return err
		}
	}

	// query param host
	qrHost := o.Host
	qHost := qrHost
	if qHost != "" {
		if err := r.SetQueryParam("host", qHost); err != nil {
			return err
		}
	}

	// query param ip
	qrIP := o.IP
	qIP := qrIP
	if qIP != "" {
		if err := r.SetQueryParam("ip", qIP); err != nil {
			return err
		}
	}

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
