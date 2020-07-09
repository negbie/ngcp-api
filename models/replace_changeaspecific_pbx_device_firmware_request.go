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

// ReplaceChangeaspecificPbxDeviceFirmwareRequest Replace/changeaspecificPbxDeviceFirmwareRequest
//
// swagger:model Replace/changeaspecificPbxDeviceFirmwareRequest
type ReplaceChangeaspecificPbxDeviceFirmwareRequest struct {

	// device id
	// Required: true
	DeviceID *string `json:"device_id"`

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

// Validate validates this replace changeaspecific pbx device firmware request
func (m *ReplaceChangeaspecificPbxDeviceFirmwareRequest) Validate(formats strfmt.Registry) error {
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

func (m *ReplaceChangeaspecificPbxDeviceFirmwareRequest) validateDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("device_id", "body", m.DeviceID); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificPbxDeviceFirmwareRequest) validateFilename(formats strfmt.Registry) error {

	if err := validate.Required("filename", "body", m.Filename); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificPbxDeviceFirmwareRequest) validateTag(formats strfmt.Registry) error {

	if err := validate.Required("tag", "body", m.Tag); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificPbxDeviceFirmwareRequest) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReplaceChangeaspecificPbxDeviceFirmwareRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplaceChangeaspecificPbxDeviceFirmwareRequest) UnmarshalBinary(b []byte) error {
	var res ReplaceChangeaspecificPbxDeviceFirmwareRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
