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

// CreateanewPreferencesMetaEntryRequest CreateanewPreferencesMetaEntryRequest
//
// swagger:model CreateanewPreferencesMetaEntryRequest
type CreateanewPreferencesMetaEntryRequest struct {

	// attribute
	// Required: true
	Attribute *string `json:"attribute"`

	// data type
	// Required: true
	DataType *string `json:"data_type"`

	// description
	// Required: true
	Description *string `json:"description"`

	// enum
	// Required: true
	Enum []string `json:"enum"`

	// fielddev pref
	// Required: true
	FielddevPref *string `json:"fielddev_pref"`

	// label
	// Required: true
	Label *string `json:"label"`

	// max occur
	// Required: true
	MaxOccur *string `json:"max_occur"`
}

// Validate validates this createanew preferences meta entry request
func (m *CreateanewPreferencesMetaEntryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttribute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFielddevPref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxOccur(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateanewPreferencesMetaEntryRequest) validateAttribute(formats strfmt.Registry) error {

	if err := validate.Required("attribute", "body", m.Attribute); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewPreferencesMetaEntryRequest) validateDataType(formats strfmt.Registry) error {

	if err := validate.Required("data_type", "body", m.DataType); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewPreferencesMetaEntryRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewPreferencesMetaEntryRequest) validateEnum(formats strfmt.Registry) error {

	if err := validate.Required("enum", "body", m.Enum); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewPreferencesMetaEntryRequest) validateFielddevPref(formats strfmt.Registry) error {

	if err := validate.Required("fielddev_pref", "body", m.FielddevPref); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewPreferencesMetaEntryRequest) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewPreferencesMetaEntryRequest) validateMaxOccur(formats strfmt.Registry) error {

	if err := validate.Required("max_occur", "body", m.MaxOccur); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateanewPreferencesMetaEntryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateanewPreferencesMetaEntryRequest) UnmarshalBinary(b []byte) error {
	var res CreateanewPreferencesMetaEntryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
