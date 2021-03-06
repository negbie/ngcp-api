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

// CreateanewSubscriberProfileRequest CreateanewSubscriberProfileRequest
//
// swagger:model CreateanewSubscriberProfileRequest
type CreateanewSubscriberProfileRequest struct {

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

// Validate validates this createanew subscriber profile request
func (m *CreateanewSubscriberProfileRequest) Validate(formats strfmt.Registry) error {
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

func (m *CreateanewSubscriberProfileRequest) validateAttribute(formats strfmt.Registry) error {

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

func (m *CreateanewSubscriberProfileRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewSubscriberProfileRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewSubscriberProfileRequest) validateProfileSetID(formats strfmt.Registry) error {

	if err := validate.Required("profile_set_id", "body", m.ProfileSetID); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewSubscriberProfileRequest) validateSetDefault(formats strfmt.Registry) error {

	if err := validate.Required("set_default", "body", m.SetDefault); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateanewSubscriberProfileRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateanewSubscriberProfileRequest) UnmarshalBinary(b []byte) error {
	var res CreateanewSubscriberProfileRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
