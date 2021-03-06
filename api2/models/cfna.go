// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Cfna Cfna
//
// swagger:model Cfna
type Cfna struct {

	// bnumbers
	// Required: true
	Bnumbers []*Bnumber `json:"bnumbers"`

	// bnumbers is regex
	// Required: true
	BnumbersIsRegex *string `json:"bnumbers_is_regex"`

	// bnumbers mode
	// Required: true
	BnumbersMode *string `json:"bnumbers_mode"`

	// destinations
	// Required: true
	Destinations []*Destination `json:"destinations"`

	// sources
	// Required: true
	Sources []*Source `json:"sources"`

	// sources is regex
	// Required: true
	SourcesIsRegex *string `json:"sources_is_regex"`

	// sources mode
	// Required: true
	SourcesMode *string `json:"sources_mode"`

	// times
	// Required: true
	Times []*Time `json:"times"`
}

// Validate validates this cfna
func (m *Cfna) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBnumbers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBnumbersIsRegex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBnumbersMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourcesIsRegex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourcesMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cfna) validateBnumbers(formats strfmt.Registry) error {

	if err := validate.Required("bnumbers", "body", m.Bnumbers); err != nil {
		return err
	}

	for i := 0; i < len(m.Bnumbers); i++ {
		if swag.IsZero(m.Bnumbers[i]) { // not required
			continue
		}

		if m.Bnumbers[i] != nil {
			if err := m.Bnumbers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bnumbers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Cfna) validateBnumbersIsRegex(formats strfmt.Registry) error {

	if err := validate.Required("bnumbers_is_regex", "body", m.BnumbersIsRegex); err != nil {
		return err
	}

	return nil
}

func (m *Cfna) validateBnumbersMode(formats strfmt.Registry) error {

	if err := validate.Required("bnumbers_mode", "body", m.BnumbersMode); err != nil {
		return err
	}

	return nil
}

func (m *Cfna) validateDestinations(formats strfmt.Registry) error {

	if err := validate.Required("destinations", "body", m.Destinations); err != nil {
		return err
	}

	for i := 0; i < len(m.Destinations); i++ {
		if swag.IsZero(m.Destinations[i]) { // not required
			continue
		}

		if m.Destinations[i] != nil {
			if err := m.Destinations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("destinations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Cfna) validateSources(formats strfmt.Registry) error {

	if err := validate.Required("sources", "body", m.Sources); err != nil {
		return err
	}

	for i := 0; i < len(m.Sources); i++ {
		if swag.IsZero(m.Sources[i]) { // not required
			continue
		}

		if m.Sources[i] != nil {
			if err := m.Sources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Cfna) validateSourcesIsRegex(formats strfmt.Registry) error {

	if err := validate.Required("sources_is_regex", "body", m.SourcesIsRegex); err != nil {
		return err
	}

	return nil
}

func (m *Cfna) validateSourcesMode(formats strfmt.Registry) error {

	if err := validate.Required("sources_mode", "body", m.SourcesMode); err != nil {
		return err
	}

	return nil
}

func (m *Cfna) validateTimes(formats strfmt.Registry) error {

	if err := validate.Required("times", "body", m.Times); err != nil {
		return err
	}

	for i := 0; i < len(m.Times); i++ {
		if swag.IsZero(m.Times[i]) { // not required
			continue
		}

		if m.Times[i] != nil {
			if err := m.Times[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("times" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Cfna) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cfna) UnmarshalBinary(b []byte) error {
	var res Cfna
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
