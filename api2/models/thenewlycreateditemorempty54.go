// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Thenewlycreateditemorempty54 Thenewlycreateditemorempty54
//
// swagger:model Thenewlycreateditemorempty54
type Thenewlycreateditemorempty54 struct {

	// login
	// Required: true
	Login *string `json:"login"`
}

// Validate validates this thenewlycreateditemorempty54
func (m *Thenewlycreateditemorempty54) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Thenewlycreateditemorempty54) validateLogin(formats strfmt.Registry) error {

	if err := validate.Required("login", "body", m.Login); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Thenewlycreateditemorempty54) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Thenewlycreateditemorempty54) UnmarshalBinary(b []byte) error {
	var res Thenewlycreateditemorempty54
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
