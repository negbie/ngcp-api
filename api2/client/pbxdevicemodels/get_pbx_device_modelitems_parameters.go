// Code generated by go-swagger; DO NOT EDIT.

package pbxdevicemodels

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

// NewGetPbxDeviceModelitemsParams creates a new GetPbxDeviceModelitemsParams object
// with the default values initialized.
func NewGetPbxDeviceModelitemsParams() *GetPbxDeviceModelitemsParams {
	var ()
	return &GetPbxDeviceModelitemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPbxDeviceModelitemsParamsWithTimeout creates a new GetPbxDeviceModelitemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPbxDeviceModelitemsParamsWithTimeout(timeout time.Duration) *GetPbxDeviceModelitemsParams {
	var ()
	return &GetPbxDeviceModelitemsParams{

		timeout: timeout,
	}
}

// NewGetPbxDeviceModelitemsParamsWithContext creates a new GetPbxDeviceModelitemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPbxDeviceModelitemsParamsWithContext(ctx context.Context) *GetPbxDeviceModelitemsParams {
	var ()
	return &GetPbxDeviceModelitemsParams{

		Context: ctx,
	}
}

// NewGetPbxDeviceModelitemsParamsWithHTTPClient creates a new GetPbxDeviceModelitemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPbxDeviceModelitemsParamsWithHTTPClient(client *http.Client) *GetPbxDeviceModelitemsParams {
	var ()
	return &GetPbxDeviceModelitemsParams{
		HTTPClient: client,
	}
}

/*GetPbxDeviceModelitemsParams contains all the parameters to send to the API endpoint
for the get pbx device modelitems operation typically these are written to a http.Request
*/
type GetPbxDeviceModelitemsParams struct {

	/*Model
	  Filter for models matching a model name pattern

	*/
	Model string
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
	  Filter for models belonging to a certain reseller

	*/
	ResellerID string
	/*Rows
	  Number of rows in one pagination page (default: 10)

	*/
	Rows string
	/*Vendor
	  Filter for vendor matching a vendor name pattern

	*/
	Vendor string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get pbx device modelitems params
func (o *GetPbxDeviceModelitemsParams) WithTimeout(timeout time.Duration) *GetPbxDeviceModelitemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pbx device modelitems params
func (o *GetPbxDeviceModelitemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pbx device modelitems params
func (o *GetPbxDeviceModelitemsParams) WithContext(ctx context.Context) *GetPbxDeviceModelitemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pbx device modelitems params
func (o *GetPbxDeviceModelitemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pbx device modelitems params
func (o *GetPbxDeviceModelitemsParams) WithHTTPClient(client *http.Client) *GetPbxDeviceModelitemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pbx device modelitems params
func (o *GetPbxDeviceModelitemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModel adds the model to the get pbx device modelitems params
func (o *GetPbxDeviceModelitemsParams) WithModel(model string) *GetPbxDeviceModelitemsParams {
	o.SetModel(model)
	return o
}

// SetModel adds the model to the get pbx device modelitems params
func (o *GetPbxDeviceModelitemsParams) SetModel(model string) {
	o.Model = model
}

// WithOrderBy adds the orderBy to the get pbx device modelitems params
func (o *GetPbxDeviceModelitemsParams) WithOrderBy(orderBy string) *GetPbxDeviceModelitemsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get pbx device modelitems params
func (o *GetPbxDeviceModelitemsParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the get pbx device modelitems params
func (o *GetPbxDeviceModelitemsParams) WithOrderByDirection(orderByDirection string) *GetPbxDeviceModelitemsParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the get pbx device modelitems params
func (o *GetPbxDeviceModelitemsParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the get pbx device modelitems params
func (o *GetPbxDeviceModelitemsParams) WithPage(page string) *GetPbxDeviceModelitemsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get pbx device modelitems params
func (o *GetPbxDeviceModelitemsParams) SetPage(page string) {
	o.Page = page
}

// WithResellerID adds the resellerID to the get pbx device modelitems params
func (o *GetPbxDeviceModelitemsParams) WithResellerID(resellerID string) *GetPbxDeviceModelitemsParams {
	o.SetResellerID(resellerID)
	return o
}

// SetResellerID adds the resellerId to the get pbx device modelitems params
func (o *GetPbxDeviceModelitemsParams) SetResellerID(resellerID string) {
	o.ResellerID = resellerID
}

// WithRows adds the rows to the get pbx device modelitems params
func (o *GetPbxDeviceModelitemsParams) WithRows(rows string) *GetPbxDeviceModelitemsParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the get pbx device modelitems params
func (o *GetPbxDeviceModelitemsParams) SetRows(rows string) {
	o.Rows = rows
}

// WithVendor adds the vendor to the get pbx device modelitems params
func (o *GetPbxDeviceModelitemsParams) WithVendor(vendor string) *GetPbxDeviceModelitemsParams {
	o.SetVendor(vendor)
	return o
}

// SetVendor adds the vendor to the get pbx device modelitems params
func (o *GetPbxDeviceModelitemsParams) SetVendor(vendor string) {
	o.Vendor = vendor
}

// WriteToRequest writes these params to a swagger request
func (o *GetPbxDeviceModelitemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param model
	qrModel := o.Model
	qModel := qrModel
	if qModel != "" {
		if err := r.SetQueryParam("model", qModel); err != nil {
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

	// query param vendor
	qrVendor := o.Vendor
	qVendor := qrVendor
	if qVendor != "" {
		if err := r.SetQueryParam("vendor", qVendor); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
