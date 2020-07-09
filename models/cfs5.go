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

// Cfs5 Cfs5
//
// swagger:model Cfs5
type Cfs5 struct {

	// bnumberset
	// Required: true
	Bnumberset *string `json:"bnumberset"`

	// bnumberset id
	// Required: true
	BnumbersetID *float64 `json:"bnumberset_id"`

	// destinationset
	// Required: true
	Destinationset *string `json:"destinationset"`

	// destinationset id
	// Required: true
	DestinationsetID *float64 `json:"destinationset_id"`

	// sourceset
	// Required: true
	Sourceset *string `json:"sourceset"`

	// sourceset id
	// Required: true
	SourcesetID *float64 `json:"sourceset_id"`

	// timeset
	// Required: true
	Timeset *string `json:"timeset"`

	// timeset id
	// Required: true
	TimesetID *float64 `json:"timeset_id"`
}

// Validate validates this cfs5
func (m *Cfs5) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBnumberset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBnumbersetID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationsetID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourcesetID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimesetID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cfs5) validateBnumberset(formats strfmt.Registry) error {

	if err := validate.Required("bnumberset", "body", m.Bnumberset); err != nil {
		return err
	}

	return nil
}

func (m *Cfs5) validateBnumbersetID(formats strfmt.Registry) error {

	if err := validate.Required("bnumberset_id", "body", m.BnumbersetID); err != nil {
		return err
	}

	return nil
}

func (m *Cfs5) validateDestinationset(formats strfmt.Registry) error {

	if err := validate.Required("destinationset", "body", m.Destinationset); err != nil {
		return err
	}

	return nil
}

func (m *Cfs5) validateDestinationsetID(formats strfmt.Registry) error {

	if err := validate.Required("destinationset_id", "body", m.DestinationsetID); err != nil {
		return err
	}

	return nil
}

func (m *Cfs5) validateSourceset(formats strfmt.Registry) error {

	if err := validate.Required("sourceset", "body", m.Sourceset); err != nil {
		return err
	}

	return nil
}

func (m *Cfs5) validateSourcesetID(formats strfmt.Registry) error {

	if err := validate.Required("sourceset_id", "body", m.SourcesetID); err != nil {
		return err
	}

	return nil
}

func (m *Cfs5) validateTimeset(formats strfmt.Registry) error {

	if err := validate.Required("timeset", "body", m.Timeset); err != nil {
		return err
	}

	return nil
}

func (m *Cfs5) validateTimesetID(formats strfmt.Registry) error {

	if err := validate.Required("timeset_id", "body", m.TimesetID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Cfs5) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cfs5) UnmarshalBinary(b []byte) error {
	var res Cfs5
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
