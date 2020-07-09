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

// Thenewlycreateditemorempty26 Thenewlycreateditemorempty26
//
// swagger:model Thenewlycreateditemorempty26
type Thenewlycreateditemorempty26 struct {

	// attribute
	// Required: true
	Attribute *Attribute `json:"attribute"`

	// description
	// Required: true
	Description *string `json:"description"`

	// name
	// Required: true
	Name *string `json:"name"`

	// profile set id
	// Required: true
	ProfileSetID *string `json:"profile_set_id"`

	// set default
	// Required: true
	SetDefault *string `json:"set_default"`
}

// Validate validates this thenewlycreateditemorempty26
func (m *Thenewlycreateditemorempty26) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttribute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfileSetID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSetDefault(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Thenewlycreateditemorempty26) validateAttribute(formats strfmt.Registry) error {

	if err := validate.Required("attribute", "body", m.Attribute); err != nil {
		return err
	}

	if m.Attribute != nil {
		if err := m.Attribute.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attribute")
			}
			return err
		}
	}

	return nil
}

func (m *Thenewlycreateditemorempty26) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty26) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty26) validateProfileSetID(formats strfmt.Registry) error {

	if err := validate.Required("profile_set_id", "body", m.ProfileSetID); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty26) validateSetDefault(formats strfmt.Registry) error {

	if err := validate.Required("set_default", "body", m.SetDefault); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Thenewlycreateditemorempty26) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Thenewlycreateditemorempty26) UnmarshalBinary(b []byte) error {
	var res Thenewlycreateditemorempty26
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
