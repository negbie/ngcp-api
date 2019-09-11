// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Thenewlycreateditemorempty37 Thenewlycreateditemorempty37
// swagger:model Thenewlycreateditemorempty37
type Thenewlycreateditemorempty37 struct {

	// billing profile id
	// Required: true
	BillingProfileID *string `json:"billing_profile_id"`

	// detail
	// Required: true
	Detail *string `json:"detail"`

	// zone
	// Required: true
	Zone *string `json:"zone"`
}

// Validate validates this thenewlycreateditemorempty37
func (m *Thenewlycreateditemorempty37) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingProfileID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDetail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Thenewlycreateditemorempty37) validateBillingProfileID(formats strfmt.Registry) error {

	if err := validate.Required("billing_profile_id", "body", m.BillingProfileID); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty37) validateDetail(formats strfmt.Registry) error {

	if err := validate.Required("detail", "body", m.Detail); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty37) validateZone(formats strfmt.Registry) error {

	if err := validate.Required("zone", "body", m.Zone); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Thenewlycreateditemorempty37) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Thenewlycreateditemorempty37) UnmarshalBinary(b []byte) error {
	var res Thenewlycreateditemorempty37
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}