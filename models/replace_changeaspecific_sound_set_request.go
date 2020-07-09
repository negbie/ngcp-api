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

// ReplaceChangeaspecificSoundSetRequest Replace/changeaspecificSoundSetRequest
//
// swagger:model Replace/changeaspecificSoundSetRequest
type ReplaceChangeaspecificSoundSetRequest struct {

	// contract default
	// Required: true
	ContractDefault *string `json:"contract_default"`

	// contract id
	// Required: true
	ContractID *string `json:"contract_id"`

	// description
	// Required: true
	Description *string `json:"description"`

	// name
	// Required: true
	Name *string `json:"name"`

	// reseller id
	// Required: true
	ResellerID *string `json:"reseller_id"`
}

// Validate validates this replace changeaspecific sound set request
func (m *ReplaceChangeaspecificSoundSetRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContractDefault(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContractID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResellerID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReplaceChangeaspecificSoundSetRequest) validateContractDefault(formats strfmt.Registry) error {

	if err := validate.Required("contract_default", "body", m.ContractDefault); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificSoundSetRequest) validateContractID(formats strfmt.Registry) error {

	if err := validate.Required("contract_id", "body", m.ContractID); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificSoundSetRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificSoundSetRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificSoundSetRequest) validateResellerID(formats strfmt.Registry) error {

	if err := validate.Required("reseller_id", "body", m.ResellerID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReplaceChangeaspecificSoundSetRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplaceChangeaspecificSoundSetRequest) UnmarshalBinary(b []byte) error {
	var res ReplaceChangeaspecificSoundSetRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
