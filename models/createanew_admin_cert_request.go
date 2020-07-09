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

// CreateanewAdminCertRequest CreateanewAdminCertRequest
//
// swagger:model CreateanewAdminCertRequest
type CreateanewAdminCertRequest struct {

	// login
	// Required: true
	Login *string `json:"login"`
}

// Validate validates this createanew admin cert request
func (m *CreateanewAdminCertRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateanewAdminCertRequest) validateLogin(formats strfmt.Registry) error {

	if err := validate.Required("login", "body", m.Login); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateanewAdminCertRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateanewAdminCertRequest) UnmarshalBinary(b []byte) error {
	var res CreateanewAdminCertRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
