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

// PeeringGroups PeeringGroups
//
// swagger:model PeeringGroups
type PeeringGroups struct {

	// contract id
	// Required: true
	ContractID *string `json:"contract_id"`

	// description
	// Required: true
	Description *string `json:"description"`

	// name
	// Required: true
	Name *string `json:"name"`

	// priority
	// Required: true
	Priority *string `json:"priority"`
}

// Validate validates this peering groups
func (m *PeeringGroups) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContractID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PeeringGroups) validateContractID(formats strfmt.Registry) error {

	if err := validate.Required("contract_id", "body", m.ContractID); err != nil {
		return err
	}

	return nil
}

func (m *PeeringGroups) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *PeeringGroups) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PeeringGroups) validatePriority(formats strfmt.Registry) error {

	if err := validate.Required("priority", "body", m.Priority); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PeeringGroups) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PeeringGroups) UnmarshalBinary(b []byte) error {
	var res PeeringGroups
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
