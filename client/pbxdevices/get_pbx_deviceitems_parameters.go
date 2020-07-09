// Code generated by go-swagger; DO NOT EDIT.

package pbxdevices

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

// NewGetPbxDeviceitemsParams creates a new GetPbxDeviceitemsParams object
// with the default values initialized.
func NewGetPbxDeviceitemsParams() *GetPbxDeviceitemsParams {
	var ()
	return &GetPbxDeviceitemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPbxDeviceitemsParamsWithTimeout creates a new GetPbxDeviceitemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPbxDeviceitemsParamsWithTimeout(timeout time.Duration) *GetPbxDeviceitemsParams {
	var ()
	return &GetPbxDeviceitemsParams{

		timeout: timeout,
	}
}

// NewGetPbxDeviceitemsParamsWithContext creates a new GetPbxDeviceitemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPbxDeviceitemsParamsWithContext(ctx context.Context) *GetPbxDeviceitemsParams {
	var ()
	return &GetPbxDeviceitemsParams{

		Context: ctx,
	}
}

// NewGetPbxDeviceitemsParamsWithHTTPClient creates a new GetPbxDeviceitemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPbxDeviceitemsParamsWithHTTPClient(client *http.Client) *GetPbxDeviceitemsParams {
	var ()
	return &GetPbxDeviceitemsParams{
		HTTPClient: client,
	}
}

/*GetPbxDeviceitemsParams contains all the parameters to send to the API endpoint
for the get pbx deviceitems operation typically these are written to a http.Request
*/
type GetPbxDeviceitemsParams struct {

	/*CustomerID
	  Search for PBX devices belonging to a specific customer

	*/
	CustomerID string
	/*Identifier
	  Search for PBX devices matching an identifier/MAC pattern (wildcards possible)

	*/
	Identifier string
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
	/*ProfileID
	  Search for PBX devices with a specific autoprovisioning device profile

	*/
	ProfileID string
	/*Rows
	  Number of rows in one pagination page (default: 10)

	*/
	Rows string
	/*StationName
	  Search for PBX devices matching a station_name pattern (wildcards possible)

	*/
	StationName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get pbx deviceitems params
func (o *GetPbxDeviceitemsParams) WithTimeout(timeout time.Duration) *GetPbxDeviceitemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pbx deviceitems params
func (o *GetPbxDeviceitemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pbx deviceitems params
func (o *GetPbxDeviceitemsParams) WithContext(ctx context.Context) *GetPbxDeviceitemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pbx deviceitems params
func (o *GetPbxDeviceitemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pbx deviceitems params
func (o *GetPbxDeviceitemsParams) WithHTTPClient(client *http.Client) *GetPbxDeviceitemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pbx deviceitems params
func (o *GetPbxDeviceitemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the get pbx deviceitems params
func (o *GetPbxDeviceitemsParams) WithCustomerID(customerID string) *GetPbxDeviceitemsParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get pbx deviceitems params
func (o *GetPbxDeviceitemsParams) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WithIdentifier adds the identifier to the get pbx deviceitems params
func (o *GetPbxDeviceitemsParams) WithIdentifier(identifier string) *GetPbxDeviceitemsParams {
	o.SetIdentifier(identifier)
	return o
}

// SetIdentifier adds the identifier to the get pbx deviceitems params
func (o *GetPbxDeviceitemsParams) SetIdentifier(identifier string) {
	o.Identifier = identifier
}

// WithOrderBy adds the orderBy to the get pbx deviceitems params
func (o *GetPbxDeviceitemsParams) WithOrderBy(orderBy string) *GetPbxDeviceitemsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get pbx deviceitems params
func (o *GetPbxDeviceitemsParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the get pbx deviceitems params
func (o *GetPbxDeviceitemsParams) WithOrderByDirection(orderByDirection string) *GetPbxDeviceitemsParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the get pbx deviceitems params
func (o *GetPbxDeviceitemsParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the get pbx deviceitems params
func (o *GetPbxDeviceitemsParams) WithPage(page string) *GetPbxDeviceitemsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get pbx deviceitems params
func (o *GetPbxDeviceitemsParams) SetPage(page string) {
	o.Page = page
}

// WithProfileID adds the profileID to the get pbx deviceitems params
func (o *GetPbxDeviceitemsParams) WithProfileID(profileID string) *GetPbxDeviceitemsParams {
	o.SetProfileID(profileID)
	return o
}

// SetProfileID adds the profileId to the get pbx deviceitems params
func (o *GetPbxDeviceitemsParams) SetProfileID(profileID string) {
	o.ProfileID = profileID
}

// WithRows adds the rows to the get pbx deviceitems params
func (o *GetPbxDeviceitemsParams) WithRows(rows string) *GetPbxDeviceitemsParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the get pbx deviceitems params
func (o *GetPbxDeviceitemsParams) SetRows(rows string) {
	o.Rows = rows
}

// WithStationName adds the stationName to the get pbx deviceitems params
func (o *GetPbxDeviceitemsParams) WithStationName(stationName string) *GetPbxDeviceitemsParams {
	o.SetStationName(stationName)
	return o
}

// SetStationName adds the stationName to the get pbx deviceitems params
func (o *GetPbxDeviceitemsParams) SetStationName(stationName string) {
	o.StationName = stationName
}

// WriteToRequest writes these params to a swagger request
func (o *GetPbxDeviceitemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param identifier
	qrIdentifier := o.Identifier
	qIdentifier := qrIdentifier
	if qIdentifier != "" {
		if err := r.SetQueryParam("identifier", qIdentifier); err != nil {
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

	// query param profile_id
	qrProfileID := o.ProfileID
	qProfileID := qrProfileID
	if qProfileID != "" {
		if err := r.SetQueryParam("profile_id", qProfileID); err != nil {
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

	// query param station_name
	qrStationName := o.StationName
	qStationName := qrStationName
	if qStationName != "" {
		if err := r.SetQueryParam("station_name", qStationName); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}