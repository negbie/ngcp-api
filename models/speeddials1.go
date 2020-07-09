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

// Speeddials1 Speeddials1
//
// swagger:model Speeddials1
type Speeddials1 struct {

	// destination
	// Required: true
	Destination *string `json:"destination"`

	// slot
	// Required: true
	Slot *string `json:"slot"`
}

// Validate validates this speeddials1
func (m *Speeddials1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlot(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Speeddials1) validateDestination(formats strfmt.Registry) error {

	if err := validate.Required("destination", "body", m.Destination); err != nil {
		return err
	}

	return nil
}

func (m *Speeddials1) validateSlot(formats strfmt.Registry) error {

	if err := validate.Required("slot", "body", m.Slot); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Speeddials1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Speeddials1) UnmarshalBinary(b []byte) error {
	var res Speeddials1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
