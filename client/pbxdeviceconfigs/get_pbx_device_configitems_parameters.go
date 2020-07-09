// Code generated by go-swagger; DO NOT EDIT.

package pbxdeviceconfigs

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

// NewGetPbxDeviceConfigitemsParams creates a new GetPbxDeviceConfigitemsParams object
// with the default values initialized.
func NewGetPbxDeviceConfigitemsParams() *GetPbxDeviceConfigitemsParams {
	var ()
	return &GetPbxDeviceConfigitemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPbxDeviceConfigitemsParamsWithTimeout creates a new GetPbxDeviceConfigitemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPbxDeviceConfigitemsParamsWithTimeout(timeout time.Duration) *GetPbxDeviceConfigitemsParams {
	var ()
	return &GetPbxDeviceConfigitemsParams{

		timeout: timeout,
	}
}

// NewGetPbxDeviceConfigitemsParamsWithContext creates a new GetPbxDeviceConfigitemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPbxDeviceConfigitemsParamsWithContext(ctx context.Context) *GetPbxDeviceConfigitemsParams {
	var ()
	return &GetPbxDeviceConfigitemsParams{

		Context: ctx,
	}
}

// NewGetPbxDeviceConfigitemsParamsWithHTTPClient creates a new GetPbxDeviceConfigitemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPbxDeviceConfigitemsParamsWithHTTPClient(client *http.Client) *GetPbxDeviceConfigitemsParams {
	var ()
	return &GetPbxDeviceConfigitemsParams{
		HTTPClient: client,
	}
}

/*GetPbxDeviceConfigitemsParams contains all the parameters to send to the API endpoint
for the get pbx device configitems operation typically these are written to a http.Request
*/
type GetPbxDeviceConfigitemsParams struct {

	/*ContentType
	  Filter for configs by a specific content type

	*/
	ContentType string
	/*DeviceID
	  Filter for configs of a specific device model

	*/
	DeviceID string
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
	/*Version
	  Filter for configs by a specific version

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get pbx device configitems params
func (o *GetPbxDeviceConfigitemsParams) WithTimeout(timeout time.Duration) *GetPbxDeviceConfigitemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pbx device configitems params
func (o *GetPbxDeviceConfigitemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pbx device configitems params
func (o *GetPbxDeviceConfigitemsParams) WithContext(ctx context.Context) *GetPbxDeviceConfigitemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pbx device configitems params
func (o *GetPbxDeviceConfigitemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pbx device configitems params
func (o *GetPbxDeviceConfigitemsParams) WithHTTPClient(client *http.Client) *GetPbxDeviceConfigitemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pbx device configitems params
func (o *GetPbxDeviceConfigitemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the get pbx device configitems params
func (o *GetPbxDeviceConfigitemsParams) WithContentType(contentType string) *GetPbxDeviceConfigitemsParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the get pbx device configitems params
func (o *GetPbxDeviceConfigitemsParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithDeviceID adds the deviceID to the get pbx device configitems params
func (o *GetPbxDeviceConfigitemsParams) WithDeviceID(deviceID string) *GetPbxDeviceConfigitemsParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get pbx device configitems params
func (o *GetPbxDeviceConfigitemsParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithOrderBy adds the orderBy to the get pbx device configitems params
func (o *GetPbxDeviceConfigitemsParams) WithOrderBy(orderBy string) *GetPbxDeviceConfigitemsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get pbx device configitems params
func (o *GetPbxDeviceConfigitemsParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the get pbx device configitems params
func (o *GetPbxDeviceConfigitemsParams) WithOrderByDirection(orderByDirection string) *GetPbxDeviceConfigitemsParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the get pbx device configitems params
func (o *GetPbxDeviceConfigitemsParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the get pbx device configitems params
func (o *GetPbxDeviceConfigitemsParams) WithPage(page string) *GetPbxDeviceConfigitemsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get pbx device configitems params
func (o *GetPbxDeviceConfigitemsParams) SetPage(page string) {
	o.Page = page
}

// WithRows adds the rows to the get pbx device configitems params
func (o *GetPbxDeviceConfigitemsParams) WithRows(rows string) *GetPbxDeviceConfigitemsParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the get pbx device configitems params
func (o *GetPbxDeviceConfigitemsParams) SetRows(rows string) {
	o.Rows = rows
}

// WithVersion adds the version to the get pbx device configitems params
func (o *GetPbxDeviceConfigitemsParams) WithVersion(version string) *GetPbxDeviceConfigitemsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get pbx device configitems params
func (o *GetPbxDeviceConfigitemsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetPbxDeviceConfigitemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param content_type
	qrContentType := o.ContentType
	qContentType := qrContentType
	if qContentType != "" {
		if err := r.SetQueryParam("content_type", qContentType); err != nil {
			return err
		}
	}

	// query param device_id
	qrDeviceID := o.DeviceID
	qDeviceID := qrDeviceID
	if qDeviceID != "" {
		if err := r.SetQueryParam("device_id", qDeviceID); err != nil {
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

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
