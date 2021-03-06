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

// VoicemailSettings2 VoicemailSettings2
//
// swagger:model VoicemailSettings2
type VoicemailSettings2 struct {

	// attach
	// Required: true
	Attach *bool `json:"attach"`

	// delete
	// Required: true
	Delete *bool `json:"delete"`

	// email
	// Required: true
	Email *string `json:"email"`

	// pin
	// Required: true
	Pin *string `json:"pin"`

	// sms number
	// Required: true
	SmsNumber *string `json:"sms_number"`
}

// Validate validates this voicemail settings2
func (m *VoicemailSettings2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttach(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDelete(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSmsNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VoicemailSettings2) validateAttach(formats strfmt.Registry) error {

	if err := validate.Required("attach", "body", m.Attach); err != nil {
		return err
	}

	return nil
}

func (m *VoicemailSettings2) validateDelete(formats strfmt.Registry) error {

	if err := validate.Required("delete", "body", m.Delete); err != nil {
		return err
	}

	return nil
}

func (m *VoicemailSettings2) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *VoicemailSettings2) validatePin(formats strfmt.Registry) error {

	if err := validate.Required("pin", "body", m.Pin); err != nil {
		return err
	}

	return nil
}

func (m *VoicemailSettings2) validateSmsNumber(formats strfmt.Registry) error {

	if err := validate.Required("sms_number", "body", m.SmsNumber); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VoicemailSettings2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VoicemailSettings2) UnmarshalBinary(b []byte) error {
	var res VoicemailSettings2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
