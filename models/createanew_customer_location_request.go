// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateanewCustomerLocationRequest CreateanewCustomerLocationRequest
//
// swagger:model CreateanewCustomerLocationRequest
type CreateanewCustomerLocationRequest struct {

	// blocks
	// Required: true
	Blocks []*Block `json:"blocks"`

	// contract id
	// Required: true
	ContractID *string `json:"contract_id"`

	// description
	// Required: true
	Description *string `json:"description"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this createanew customer location request
func (m *CreateanewCustomerLocationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlocks(formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateanewCustomerLocationRequest) validateBlocks(formats strfmt.Registry) error {

	if err := validate.Required("blocks", "body", m.Blocks); err != nil {
		return err
	}

	for i := 0; i < len(m.Blocks); i++ {
		if swag.IsZero(m.Blocks[i]) { // not required
			continue
		}

		if m.Blocks[i] != nil {
			if err := m.Blocks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("blocks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateanewCustomerLocationRequest) validateContractID(formats strfmt.Registry) error {

	if err := validate.Required("contract_id", "body", m.ContractID); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewCustomerLocationRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *CreateanewCustomerLocationRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateanewCustomerLocationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateanewCustomerLocationRequest) UnmarshalBinary(b []byte) error {
	var res CreateanewCustomerLocationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
