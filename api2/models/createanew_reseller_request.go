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

// CreateanewResellerRequest CreateanewResellerRequest
//
// swagger:model CreateanewResellerRequest
type CreateanewResellerRequest struct {

	// contract id
	// Required: true
	ContractID *string `json:"contract_id"`

	// enable rtc
	// Required: true
	EnableRtc *string `json:"enable_rtc"`

	// name
	// Required: true
	Name *string `json:"name"`

	// rtc networks
	// Required: true
	RtcNetworks *string `json:"rtc_networks"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this createanew reseller request
func (m *CreateanewResellerRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContractID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnableRtc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRtcNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateanewResellerRequest) validateContractID(formats strfmt.Registry) error {

	if err := validate.Required("contract_id", "body", m.ContractID); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewResellerRequest) validateEnableRtc(formats strfmt.Registry) error {

	if err := validate.Required("enable_rtc", "body", m.EnableRtc); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewResellerRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewResellerRequest) validateRtcNetworks(formats strfmt.Registry) error {

	if err := validate.Required("rtc_networks", "body", m.RtcNetworks); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewResellerRequest) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateanewResellerRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateanewResellerRequest) UnmarshalBinary(b []byte) error {
	var res CreateanewResellerRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
