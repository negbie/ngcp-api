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

// Thenewlycreateditemorempty47 Thenewlycreateditemorempty47
//
// swagger:model Thenewlycreateditemorempty47
type Thenewlycreateditemorempty47 struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// ncos level id
	// Required: true
	NcosLevelID *string `json:"ncos_level_id"`
}

// Validate validates this thenewlycreateditemorempty47
func (m *Thenewlycreateditemorempty47) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNcosLevelID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Thenewlycreateditemorempty47) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty47) validateNcosLevelID(formats strfmt.Registry) error {

	if err := validate.Required("ncos_level_id", "body", m.NcosLevelID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Thenewlycreateditemorempty47) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Thenewlycreateditemorempty47) UnmarshalBinary(b []byte) error {
	var res Thenewlycreateditemorempty47
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
