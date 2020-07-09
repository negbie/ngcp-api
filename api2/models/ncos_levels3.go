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

// NcosLevels3 NcosLevels3
//
// swagger:model NcosLevels3
type NcosLevels3 struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// intra pbx
	// Required: true
	IntraPbx *bool `json:"intra_pbx"`

	// level
	// Required: true
	Level *string `json:"level"`

	// local ac
	// Required: true
	LocalAc *bool `json:"local_ac"`

	// mode
	// Required: true
	Mode *string `json:"mode"`

	// reseller id
	// Required: true
	ResellerID *float64 `json:"reseller_id"`
}

// Validate validates this ncos levels3
func (m *NcosLevels3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntraPbx(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalAc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResellerID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NcosLevels3) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *NcosLevels3) validateIntraPbx(formats strfmt.Registry) error {

	if err := validate.Required("intra_pbx", "body", m.IntraPbx); err != nil {
		return err
	}

	return nil
}

func (m *NcosLevels3) validateLevel(formats strfmt.Registry) error {

	if err := validate.Required("level", "body", m.Level); err != nil {
		return err
	}

	return nil
}

func (m *NcosLevels3) validateLocalAc(formats strfmt.Registry) error {

	if err := validate.Required("local_ac", "body", m.LocalAc); err != nil {
		return err
	}

	return nil
}

func (m *NcosLevels3) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *NcosLevels3) validateResellerID(formats strfmt.Registry) error {

	if err := validate.Required("reseller_id", "body", m.ResellerID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NcosLevels3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NcosLevels3) UnmarshalBinary(b []byte) error {
	var res NcosLevels3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}