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

// Block8 Block8
//
// swagger:model Block8
type Block8 struct {

	// ip
	// Required: true
	IP *string `json:"ip"`

	// mask
	// Required: true
	Mask *float64 `json:"mask"`
}

// Validate validates this block8
func (m *Block8) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMask(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Block8) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	return nil
}

func (m *Block8) validateMask(formats strfmt.Registry) error {

	if err := validate.Required("mask", "body", m.Mask); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Block8) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Block8) UnmarshalBinary(b []byte) error {
	var res Block8
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
