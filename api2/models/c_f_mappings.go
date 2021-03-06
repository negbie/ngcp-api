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

// CFMappings CFMappings
//
// swagger:model CFMappings
type CFMappings struct {

	// cfb
	// Required: true
	Cfb []*Cfb5 `json:"cfb"`

	// cfna
	// Required: true
	Cfna []*Cfna5 `json:"cfna"`

	// cfr
	// Required: true
	Cfr []*Cfr5 `json:"cfr"`

	// cfs
	// Required: true
	Cfs []*Cfs5 `json:"cfs"`

	// cft
	// Required: true
	Cft []*Cft5 `json:"cft"`

	// cft ringtimeout
	// Required: true
	CftRingtimeout *float64 `json:"cft_ringtimeout"`

	// cfu
	// Required: true
	Cfu []*Cfu5 `json:"cfu"`
}

// Validate validates this c f mappings
func (m *CFMappings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCfb(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCfna(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCfr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCfs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCft(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCftRingtimeout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCfu(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CFMappings) validateCfb(formats strfmt.Registry) error {

	if err := validate.Required("cfb", "body", m.Cfb); err != nil {
		return err
	}

	for i := 0; i < len(m.Cfb); i++ {
		if swag.IsZero(m.Cfb[i]) { // not required
			continue
		}

		if m.Cfb[i] != nil {
			if err := m.Cfb[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cfb" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CFMappings) validateCfna(formats strfmt.Registry) error {

	if err := validate.Required("cfna", "body", m.Cfna); err != nil {
		return err
	}

	for i := 0; i < len(m.Cfna); i++ {
		if swag.IsZero(m.Cfna[i]) { // not required
			continue
		}

		if m.Cfna[i] != nil {
			if err := m.Cfna[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cfna" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CFMappings) validateCfr(formats strfmt.Registry) error {

	if err := validate.Required("cfr", "body", m.Cfr); err != nil {
		return err
	}

	for i := 0; i < len(m.Cfr); i++ {
		if swag.IsZero(m.Cfr[i]) { // not required
			continue
		}

		if m.Cfr[i] != nil {
			if err := m.Cfr[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cfr" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CFMappings) validateCfs(formats strfmt.Registry) error {

	if err := validate.Required("cfs", "body", m.Cfs); err != nil {
		return err
	}

	for i := 0; i < len(m.Cfs); i++ {
		if swag.IsZero(m.Cfs[i]) { // not required
			continue
		}

		if m.Cfs[i] != nil {
			if err := m.Cfs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cfs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CFMappings) validateCft(formats strfmt.Registry) error {

	if err := validate.Required("cft", "body", m.Cft); err != nil {
		return err
	}

	for i := 0; i < len(m.Cft); i++ {
		if swag.IsZero(m.Cft[i]) { // not required
			continue
		}

		if m.Cft[i] != nil {
			if err := m.Cft[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cft" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CFMappings) validateCftRingtimeout(formats strfmt.Registry) error {

	if err := validate.Required("cft_ringtimeout", "body", m.CftRingtimeout); err != nil {
		return err
	}

	return nil
}

func (m *CFMappings) validateCfu(formats strfmt.Registry) error {

	if err := validate.Required("cfu", "body", m.Cfu); err != nil {
		return err
	}

	for i := 0; i < len(m.Cfu); i++ {
		if swag.IsZero(m.Cfu[i]) { // not required
			continue
		}

		if m.Cfu[i] != nil {
			if err := m.Cfu[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cfu" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CFMappings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CFMappings) UnmarshalBinary(b []byte) error {
	var res CFMappings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
