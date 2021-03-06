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

// CreateanewPbxDeviceConfigRequest CreateanewPbxDeviceConfigRequest
//
// swagger:model CreateanewPbxDeviceConfigRequest
type CreateanewPbxDeviceConfigRequest struct {

	// content type
	// Required: true
	ContentType *string `json:"content_type"`

	// device id
	// Required: true
	DeviceID *string `json:"device_id"`

	// version
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this createanew pbx device config request
func (m *CreateanewPbxDeviceConfigRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceID(formats); err != nil {
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

func (m *CreateanewPbxDeviceConfigRequest) validateContentType(formats strfmt.Registry) error {

	if err := validate.Required("content_type", "body", m.ContentType); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewPbxDeviceConfigRequest) validateDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("device_id", "body", m.DeviceID); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewPbxDeviceConfigRequest) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateanewPbxDeviceConfigRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateanewPbxDeviceConfigRequest) UnmarshalBinary(b []byte) error {
	var res CreateanewPbxDeviceConfigRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
