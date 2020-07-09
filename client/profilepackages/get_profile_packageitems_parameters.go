// Code generated by go-swagger; DO NOT EDIT.

package profilepackages

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

// NewGetProfilePackageitemsParams creates a new GetProfilePackageitemsParams object
// with the default values initialized.
func NewGetProfilePackageitemsParams() *GetProfilePackageitemsParams {
	var ()
	return &GetProfilePackageitemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProfilePackageitemsParamsWithTimeout creates a new GetProfilePackageitemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProfilePackageitemsParamsWithTimeout(timeout time.Duration) *GetProfilePackageitemsParams {
	var ()
	return &GetProfilePackageitemsParams{

		timeout: timeout,
	}
}

// NewGetProfilePackageitemsParamsWithContext creates a new GetProfilePackageitemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProfilePackageitemsParamsWithContext(ctx context.Context) *GetProfilePackageitemsParams {
	var ()
	return &GetProfilePackageitemsParams{

		Context: ctx,
	}
}

// NewGetProfilePackageitemsParamsWithHTTPClient creates a new GetProfilePackageitemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProfilePackageitemsParamsWithHTTPClient(client *http.Client) *GetProfilePackageitemsParams {
	var ()
	return &GetProfilePackageitemsParams{
		HTTPClient: client,
	}
}

/*GetProfilePackageitemsParams contains all the parameters to send to the API endpoint
for the get profile packageitems operation typically these are written to a http.Request
*/
type GetProfilePackageitemsParams struct {

	/*Name
	  Filter for profile packages with a specific name

	*/
	Name string
	/*NetworkName
	  Filter for profile packages containing a billing network with specific name

	*/
	NetworkName string
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
	/*ProfileName
	  Filter for profile packages containing a billing profile with specific name

	*/
	ProfileName string
	/*ResellerID
	  Filter for profile packages belonging to a specific reseller

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

// WithTimeout adds the timeout to the get profile packageitems params
func (o *GetProfilePackageitemsParams) WithTimeout(timeout time.Duration) *GetProfilePackageitemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get profile packageitems params
func (o *GetProfilePackageitemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get profile packageitems params
func (o *GetProfilePackageitemsParams) WithContext(ctx context.Context) *GetProfilePackageitemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get profile packageitems params
func (o *GetProfilePackageitemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get profile packageitems params
func (o *GetProfilePackageitemsParams) WithHTTPClient(client *http.Client) *GetProfilePackageitemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get profile packageitems params
func (o *GetProfilePackageitemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get profile packageitems params
func (o *GetProfilePackageitemsParams) WithName(name string) *GetProfilePackageitemsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get profile packageitems params
func (o *GetProfilePackageitemsParams) SetName(name string) {
	o.Name = name
}

// WithNetworkName adds the networkName to the get profile packageitems params
func (o *GetProfilePackageitemsParams) WithNetworkName(networkName string) *GetProfilePackageitemsParams {
	o.SetNetworkName(networkName)
	return o
}

// SetNetworkName adds the networkName to the get profile packageitems params
func (o *GetProfilePackageitemsParams) SetNetworkName(networkName string) {
	o.NetworkName = networkName
}

// WithOrderBy adds the orderBy to the get profile packageitems params
func (o *GetProfilePackageitemsParams) WithOrderBy(orderBy string) *GetProfilePackageitemsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get profile packageitems params
func (o *GetProfilePackageitemsParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the get profile packageitems params
func (o *GetProfilePackageitemsParams) WithOrderByDirection(orderByDirection string) *GetProfilePackageitemsParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the get profile packageitems params
func (o *GetProfilePackageitemsParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the get profile packageitems params
func (o *GetProfilePackageitemsParams) WithPage(page string) *GetProfilePackageitemsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get profile packageitems params
func (o *GetProfilePackageitemsParams) SetPage(page string) {
	o.Page = page
}

// WithProfileName adds the profileName to the get profile packageitems params
func (o *GetProfilePackageitemsParams) WithProfileName(profileName string) *GetProfilePackageitemsParams {
	o.SetProfileName(profileName)
	return o
}

// SetProfileName adds the profileName to the get profile packageitems params
func (o *GetProfilePackageitemsParams) SetProfileName(profileName string) {
	o.ProfileName = profileName
}

// WithResellerID adds the resellerID to the get profile packageitems params
func (o *GetProfilePackageitemsParams) WithResellerID(resellerID string) *GetProfilePackageitemsParams {
	o.SetResellerID(resellerID)
	return o
}

// SetResellerID adds the resellerId to the get profile packageitems params
func (o *GetProfilePackageitemsParams) SetResellerID(resellerID string) {
	o.ResellerID = resellerID
}

// WithRows adds the rows to the get profile packageitems params
func (o *GetProfilePackageitemsParams) WithRows(rows string) *GetProfilePackageitemsParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the get profile packageitems params
func (o *GetProfilePackageitemsParams) SetRows(rows string) {
	o.Rows = rows
}

// WriteToRequest writes these params to a swagger request
func (o *GetProfilePackageitemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param network_name
	qrNetworkName := o.NetworkName
	qNetworkName := qrNetworkName
	if qNetworkName != "" {
		if err := r.SetQueryParam("network_name", qNetworkName); err != nil {
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

	// query param profile_name
	qrProfileName := o.ProfileName
	qProfileName := qrProfileName
	if qProfileName != "" {
		if err := r.SetQueryParam("profile_name", qProfileName); err != nil {
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