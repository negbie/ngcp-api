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

// ReplaceChangeaspecificNumberRequest Replace/changeaspecificNumberRequest
//
// swagger:model Replace/changeaspecificNumberRequest
type ReplaceChangeaspecificNumberRequest struct {

	// ac
	// Required: true
	Ac *string `json:"ac"`

	// cc
	// Required: true
	Cc *string `json:"cc"`

	// is primary
	// Required: true
	IsPrimary *string `json:"is_primary"`

	// reseller id
	// Required: true
	ResellerID *string `json:"reseller_id"`

	// sn
	// Required: true
	Sn *string `json:"sn"`

	// subscriber id
	// Required: true
	SubscriberID *string `json:"subscriber_id"`
}

// Validate validates this replace changeaspecific number request
func (m *ReplaceChangeaspecificNumberRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsPrimary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResellerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSn(formats); err != nil {
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

func (m *ReplaceChangeaspecificNumberRequest) validateAc(formats strfmt.Registry) error {

	if err := validate.Required("ac", "body", m.Ac); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificNumberRequest) validateCc(formats strfmt.Registry) error {

	if err := validate.Required("cc", "body", m.Cc); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificNumberRequest) validateIsPrimary(formats strfmt.Registry) error {

	if err := validate.Required("is_primary", "body", m.IsPrimary); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificNumberRequest) validateResellerID(formats strfmt.Registry) error {

	if err := validate.Required("reseller_id", "body", m.ResellerID); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificNumberRequest) validateSn(formats strfmt.Registry) error {

	if err := validate.Required("sn", "body", m.Sn); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificNumberRequest) validateSubscriberID(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_id", "body", m.SubscriberID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReplaceChangeaspecificNumberRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplaceChangeaspecificNumberRequest) UnmarshalBinary(b []byte) error {
	var res ReplaceChangeaspecificNumberRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
