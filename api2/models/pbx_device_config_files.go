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

// PbxDeviceConfigFiles PbxDeviceConfigFiles
//
// swagger:model PbxDeviceConfigFiles
type PbxDeviceConfigFiles struct {

	// device id
	// Required: true
	DeviceID *float64 `json:"device_id"`

	// filename
	// Required: true
	Filename *string `json:"filename"`

	// tag
	// Required: true
	Tag *string `json:"tag"`

	// version
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this pbx device config files
func (m *PbxDeviceConfigFiles) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilename(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PbxDeviceConfigFiles) validateDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("device_id", "body", m.DeviceID); err != nil {
		return err
	}

	return nil
}

func (m *PbxDeviceConfigFiles) validateFilename(formats strfmt.Registry) error {

	if err := validate.Required("filename", "body", m.Filename); err != nil {
		return err
	}

	return nil
}

func (m *PbxDeviceConfigFiles) validateTag(formats strfmt.Registry) error {

	if err := validate.Required("tag", "body", m.Tag); err != nil {
		return err
	}

	return nil
}

func (m *PbxDeviceConfigFiles) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PbxDeviceConfigFiles) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PbxDeviceConfigFiles) UnmarshalBinary(b []byte) error {
	var res PbxDeviceConfigFiles
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
