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

// Thenewlycreateditemorempty51 Thenewlycreateditemorempty51
//
// swagger:model Thenewlycreateditemorempty51
type Thenewlycreateditemorempty51 struct {

	// amount
	// Required: true
	Amount *string `json:"amount"`

	// package id
	// Required: true
	PackageID *string `json:"package_id"`

	// request token
	// Required: true
	RequestToken *string `json:"request_token"`

	// subscriber id
	// Required: true
	SubscriberID *string `json:"subscriber_id"`
}

// Validate validates this thenewlycreateditemorempty51
func (m *Thenewlycreateditemorempty51) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackageID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Thenewlycreateditemorempty51) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty51) validatePackageID(formats strfmt.Registry) error {

	if err := validate.Required("package_id", "body", m.PackageID); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty51) validateRequestToken(formats strfmt.Registry) error {

	if err := validate.Required("request_token", "body", m.RequestToken); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty51) validateSubscriberID(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_id", "body", m.SubscriberID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Thenewlycreateditemorempty51) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Thenewlycreateditemorempty51) UnmarshalBinary(b []byte) error {
	var res Thenewlycreateditemorempty51
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
