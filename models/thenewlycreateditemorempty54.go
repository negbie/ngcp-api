// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Thenewlycreateditemorempty54 Thenewlycreateditemorempty54
// swagger:model Thenewlycreateditemorempty54
type Thenewlycreateditemorempty54 struct {

	// code
	// Required: true
	Code *string `json:"code"`

	// request token
	// Required: true
	RequestToken *string `json:"request_token"`

	// subscriber id
	// Required: true
	SubscriberID *string `json:"subscriber_id"`
}

// Validate validates this thenewlycreateditemorempty54
func (m *Thenewlycreateditemorempty54) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
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

func (m *Thenewlycreateditemorempty54) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty54) validateRequestToken(formats strfmt.Registry) error {

	if err := validate.Required("request_token", "body", m.RequestToken); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty54) validateSubscriberID(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_id", "body", m.SubscriberID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Thenewlycreateditemorempty54) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Thenewlycreateditemorempty54) UnmarshalBinary(b []byte) error {
	var res Thenewlycreateditemorempty54
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}