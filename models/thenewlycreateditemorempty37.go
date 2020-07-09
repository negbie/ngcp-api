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

// Thenewlycreateditemorempty37 Thenewlycreateditemorempty37
//
// swagger:model Thenewlycreateditemorempty37
type Thenewlycreateditemorempty37 struct {

	// callee
	// Required: true
	Callee *string `json:"callee"`

	// caller
	// Required: true
	Caller *string `json:"caller"`

	// direction
	// Required: true
	Direction *string `json:"direction"`

	// reason
	// Required: true
	Reason *string `json:"reason"`

	// status
	// Required: true
	Status *string `json:"status"`

	// subscriber id
	// Required: true
	SubscriberID *string `json:"subscriber_id"`

	// text
	// Required: true
	Text *string `json:"text"`
}

// Validate validates this thenewlycreateditemorempty37
func (m *Thenewlycreateditemorempty37) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCaller(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateText(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Thenewlycreateditemorempty37) validateCallee(formats strfmt.Registry) error {

	if err := validate.Required("callee", "body", m.Callee); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty37) validateCaller(formats strfmt.Registry) error {

	if err := validate.Required("caller", "body", m.Caller); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty37) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty37) validateReason(formats strfmt.Registry) error {

	if err := validate.Required("reason", "body", m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty37) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty37) validateSubscriberID(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_id", "body", m.SubscriberID); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty37) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", m.Text); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Thenewlycreateditemorempty37) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Thenewlycreateditemorempty37) UnmarshalBinary(b []byte) error {
	var res Thenewlycreateditemorempty37
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
