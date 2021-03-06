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

// ManagerSecretary ManagerSecretary
//
// swagger:model ManagerSecretary
type ManagerSecretary struct {

	// secretary numbers
	// Required: true
	SecretaryNumbers []*SecretaryNumber `json:"secretary_numbers"`

	// uuid
	// Required: true
	UUID *string `json:"uuid"`
}

// Validate validates this manager secretary
func (m *ManagerSecretary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecretaryNumbers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagerSecretary) validateSecretaryNumbers(formats strfmt.Registry) error {

	if err := validate.Required("secretary_numbers", "body", m.SecretaryNumbers); err != nil {
		return err
	}

	for i := 0; i < len(m.SecretaryNumbers); i++ {
		if swag.IsZero(m.SecretaryNumbers[i]) { // not required
			continue
		}

		if m.SecretaryNumbers[i] != nil {
			if err := m.SecretaryNumbers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("secretary_numbers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ManagerSecretary) validateUUID(formats strfmt.Registry) error {

	if err := validate.Required("uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ManagerSecretary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagerSecretary) UnmarshalBinary(b []byte) error {
	var res ManagerSecretary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
