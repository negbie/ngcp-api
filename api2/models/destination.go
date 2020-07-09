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

// Destination Destination
//
// swagger:model Destination
type Destination struct {

	// announcement id
	// Required: true
	AnnouncementID *string `json:"announcement_id"`

	// destination
	// Required: true
	Destination *string `json:"destination"`

	// priority
	// Required: true
	Priority *string `json:"priority"`

	// simple destination
	// Required: true
	SimpleDestination *string `json:"simple_destination"`

	// timeout
	// Required: true
	Timeout *string `json:"timeout"`
}

// Validate validates this destination
func (m *Destination) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnnouncementID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSimpleDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeout(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Destination) validateAnnouncementID(formats strfmt.Registry) error {

	if err := validate.Required("announcement_id", "body", m.AnnouncementID); err != nil {
		return err
	}

	return nil
}

func (m *Destination) validateDestination(formats strfmt.Registry) error {

	if err := validate.Required("destination", "body", m.Destination); err != nil {
		return err
	}

	return nil
}

func (m *Destination) validatePriority(formats strfmt.Registry) error {

	if err := validate.Required("priority", "body", m.Priority); err != nil {
		return err
	}

	return nil
}

func (m *Destination) validateSimpleDestination(formats strfmt.Registry) error {

	if err := validate.Required("simple_destination", "body", m.SimpleDestination); err != nil {
		return err
	}

	return nil
}

func (m *Destination) validateTimeout(formats strfmt.Registry) error {

	if err := validate.Required("timeout", "body", m.Timeout); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Destination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Destination) UnmarshalBinary(b []byte) error {
	var res Destination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}