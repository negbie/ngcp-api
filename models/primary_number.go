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

// PrimaryNumber PrimaryNumber
//
// swagger:model PrimaryNumber
type PrimaryNumber struct {

	// ac
	// Required: true
	Ac *string `json:"ac"`

	// cc
	// Required: true
	Cc *string `json:"cc"`

	// sn
	// Required: true
	Sn *string `json:"sn"`
}

// Validate validates this primary number
func (m *PrimaryNumber) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrimaryNumber) validateAc(formats strfmt.Registry) error {

	if err := validate.Required("ac", "body", m.Ac); err != nil {
		return err
	}

	return nil
}

func (m *PrimaryNumber) validateCc(formats strfmt.Registry) error {

	if err := validate.Required("cc", "body", m.Cc); err != nil {
		return err
	}

	return nil
}

func (m *PrimaryNumber) validateSn(formats strfmt.Registry) error {

	if err := validate.Required("sn", "body", m.Sn); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrimaryNumber) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrimaryNumber) UnmarshalBinary(b []byte) error {
	var res PrimaryNumber
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
