// Code generated by go-swagger; DO NOT EDIT.

package preferencesmetaentries

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

// NewGetPreferencesMetaEntryitemsParams creates a new GetPreferencesMetaEntryitemsParams object
// with the default values initialized.
func NewGetPreferencesMetaEntryitemsParams() *GetPreferencesMetaEntryitemsParams {
	var ()
	return &GetPreferencesMetaEntryitemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPreferencesMetaEntryitemsParamsWithTimeout creates a new GetPreferencesMetaEntryitemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPreferencesMetaEntryitemsParamsWithTimeout(timeout time.Duration) *GetPreferencesMetaEntryitemsParams {
	var ()
	return &GetPreferencesMetaEntryitemsParams{

		timeout: timeout,
	}
}

// NewGetPreferencesMetaEntryitemsParamsWithContext creates a new GetPreferencesMetaEntryitemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPreferencesMetaEntryitemsParamsWithContext(ctx context.Context) *GetPreferencesMetaEntryitemsParams {
	var ()
	return &GetPreferencesMetaEntryitemsParams{

		Context: ctx,
	}
}

// NewGetPreferencesMetaEntryitemsParamsWithHTTPClient creates a new GetPreferencesMetaEntryitemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPreferencesMetaEntryitemsParamsWithHTTPClient(client *http.Client) *GetPreferencesMetaEntryitemsParams {
	var ()
	return &GetPreferencesMetaEntryitemsParams{
		HTTPClient: client,
	}
}

/*GetPreferencesMetaEntryitemsParams contains all the parameters to send to the API endpoint
for the get preferences meta entryitems operation typically these are written to a http.Request
*/
type GetPreferencesMetaEntryitemsParams struct {

	/*Attribute
	  Filter for dynamic preference with a specific name

	*/
	Attribute string
	/*ModelID
	  Filter for dynamic preference relevant to the spcified pbx device model id

	*/
	ModelID string
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
	  Filter for dynamic preference relevant to the spcified reseller id

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

// WithTimeout adds the timeout to the get preferences meta entryitems params
func (o *GetPreferencesMetaEntryitemsParams) WithTimeout(timeout time.Duration) *GetPreferencesMetaEntryitemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get preferences meta entryitems params
func (o *GetPreferencesMetaEntryitemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get preferences meta entryitems params
func (o *GetPreferencesMetaEntryitemsParams) WithContext(ctx context.Context) *GetPreferencesMetaEntryitemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get preferences meta entryitems params
func (o *GetPreferencesMetaEntryitemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get preferences meta entryitems params
func (o *GetPreferencesMetaEntryitemsParams) WithHTTPClient(client *http.Client) *GetPreferencesMetaEntryitemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get preferences meta entryitems params
func (o *GetPreferencesMetaEntryitemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttribute adds the attribute to the get preferences meta entryitems params
func (o *GetPreferencesMetaEntryitemsParams) WithAttribute(attribute string) *GetPreferencesMetaEntryitemsParams {
	o.SetAttribute(attribute)
	return o
}

// SetAttribute adds the attribute to the get preferences meta entryitems params
func (o *GetPreferencesMetaEntryitemsParams) SetAttribute(attribute string) {
	o.Attribute = attribute
}

// WithModelID adds the modelID to the get preferences meta entryitems params
func (o *GetPreferencesMetaEntryitemsParams) WithModelID(modelID string) *GetPreferencesMetaEntryitemsParams {
	o.SetModelID(modelID)
	return o
}

// SetModelID adds the modelId to the get preferences meta entryitems params
func (o *GetPreferencesMetaEntryitemsParams) SetModelID(modelID string) {
	o.ModelID = modelID
}

// WithOrderBy adds the orderBy to the get preferences meta entryitems params
func (o *GetPreferencesMetaEntryitemsParams) WithOrderBy(orderBy string) *GetPreferencesMetaEntryitemsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get preferences meta entryitems params
func (o *GetPreferencesMetaEntryitemsParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithOrderByDirection adds the orderByDirection to the get preferences meta entryitems params
func (o *GetPreferencesMetaEntryitemsParams) WithOrderByDirection(orderByDirection string) *GetPreferencesMetaEntryitemsParams {
	o.SetOrderByDirection(orderByDirection)
	return o
}

// SetOrderByDirection adds the orderByDirection to the get preferences meta entryitems params
func (o *GetPreferencesMetaEntryitemsParams) SetOrderByDirection(orderByDirection string) {
	o.OrderByDirection = orderByDirection
}

// WithPage adds the page to the get preferences meta entryitems params
func (o *GetPreferencesMetaEntryitemsParams) WithPage(page string) *GetPreferencesMetaEntryitemsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get preferences meta entryitems params
func (o *GetPreferencesMetaEntryitemsParams) SetPage(page string) {
	o.Page = page
}

// WithResellerID adds the resellerID to the get preferences meta entryitems params
func (o *GetPreferencesMetaEntryitemsParams) WithResellerID(resellerID string) *GetPreferencesMetaEntryitemsParams {
	o.SetResellerID(resellerID)
	return o
}

// SetResellerID adds the resellerId to the get preferences meta entryitems params
func (o *GetPreferencesMetaEntryitemsParams) SetResellerID(resellerID string) {
	o.ResellerID = resellerID
}

// WithRows adds the rows to the get preferences meta entryitems params
func (o *GetPreferencesMetaEntryitemsParams) WithRows(rows string) *GetPreferencesMetaEntryitemsParams {
	o.SetRows(rows)
	return o
}

// SetRows adds the rows to the get preferences meta entryitems params
func (o *GetPreferencesMetaEntryitemsParams) SetRows(rows string) {
	o.Rows = rows
}

// WriteToRequest writes these params to a swagger request
func (o *GetPreferencesMetaEntryitemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param attribute
	qrAttribute := o.Attribute
	qAttribute := qrAttribute
	if qAttribute != "" {
		if err := r.SetQueryParam("attribute", qAttribute); err != nil {
			return err
		}
	}

	// query param model_id
	qrModelID := o.ModelID
	qModelID := qrModelID
	if qModelID != "" {
		if err := r.SetQueryParam("model_id", qModelID); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
