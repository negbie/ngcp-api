// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SoundFileRecordings SoundFileRecordings
// swagger:model SoundFileRecordings
type SoundFileRecordings struct {

	// filename
	// Required: true
	Filename *string `json:"filename"`

	// handle
	// Required: true
	Handle *string `json:"handle"`

	// loopplay
	// Required: true
	Loopplay *string `json:"loopplay"`

	// set id
	// Required: true
	SetID *string `json:"set_id"`
}

// Validate validates this sound file recordings
func (m *SoundFileRecordings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilename(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHandle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoopplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSetID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoundFileRecordings) validateFilename(formats strfmt.Registry) error {

	if err := validate.Required("filename", "body", m.Filename); err != nil {
		return err
	}

	return nil
}

func (m *SoundFileRecordings) validateHandle(formats strfmt.Registry) error {

	if err := validate.Required("handle", "body", m.Handle); err != nil {
		return err
	}

	return nil
}

func (m *SoundFileRecordings) validateLoopplay(formats strfmt.Registry) error {

	if err := validate.Required("loopplay", "body", m.Loopplay); err != nil {
		return err
	}

	return nil
}

func (m *SoundFileRecordings) validateSetID(formats strfmt.Registry) error {

	if err := validate.Required("set_id", "body", m.SetID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SoundFileRecordings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoundFileRecordings) UnmarshalBinary(b []byte) error {
	var res SoundFileRecordings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}