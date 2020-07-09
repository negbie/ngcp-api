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

// CreateanewInterceptionRequest CreateanewInterceptionRequest
//
// swagger:model CreateanewInterceptionRequest
type CreateanewInterceptionRequest struct {

	// liid
	// Required: true
	Liid *string `json:"liid"`

	// number
	// Required: true
	Number *string `json:"number"`

	// x2 host
	// Required: true
	X2Host *string `json:"x2_host"`

	// x2 password
	// Required: true
	X2Password *string `json:"x2_password"`

	// x2 port
	// Required: true
	X2Port *string `json:"x2_port"`

	// x2 user
	// Required: true
	X2User *string `json:"x2_user"`

	// x3 host
	// Required: true
	X3Host *string `json:"x3_host"`

	// x3 port
	// Required: true
	X3Port *string `json:"x3_port"`

	// x3 required
	// Required: true
	X3Required *string `json:"x3_required"`
}

// Validate validates this createanew interception request
func (m *CreateanewInterceptionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLiid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateX2Host(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateX2Password(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateX2Port(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateX2User(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateX3Host(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateX3Port(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateX3Required(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateanewInterceptionRequest) validateLiid(formats strfmt.Registry) error {

	if err := validate.Required("liid", "body", m.Liid); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewInterceptionRequest) validateNumber(formats strfmt.Registry) error {

	if err := validate.Required("number", "body", m.Number); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewInterceptionRequest) validateX2Host(formats strfmt.Registry) error {

	if err := validate.Required("x2_host", "body", m.X2Host); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewInterceptionRequest) validateX2Password(formats strfmt.Registry) error {

	if err := validate.Required("x2_password", "body", m.X2Password); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewInterceptionRequest) validateX2Port(formats strfmt.Registry) error {

	if err := validate.Required("x2_port", "body", m.X2Port); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewInterceptionRequest) validateX2User(formats strfmt.Registry) error {

	if err := validate.Required("x2_user", "body", m.X2User); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewInterceptionRequest) validateX3Host(formats strfmt.Registry) error {

	if err := validate.Required("x3_host", "body", m.X3Host); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewInterceptionRequest) validateX3Port(formats strfmt.Registry) error {

	if err := validate.Required("x3_port", "body", m.X3Port); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewInterceptionRequest) validateX3Required(formats strfmt.Registry) error {

	if err := validate.Required("x3_required", "body", m.X3Required); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateanewInterceptionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateanewInterceptionRequest) UnmarshalBinary(b []byte) error {
	var res CreateanewInterceptionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
