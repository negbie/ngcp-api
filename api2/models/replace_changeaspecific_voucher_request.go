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

// ReplaceChangeaspecificVoucherRequest Replace/changeaspecificVoucherRequest
//
// swagger:model Replace/changeaspecificVoucherRequest
type ReplaceChangeaspecificVoucherRequest struct {

	// amount
	// Required: true
	Amount *string `json:"amount"`

	// code
	// Required: true
	Code *string `json:"code"`

	// customer id
	// Required: true
	CustomerID *string `json:"customer_id"`

	// package id
	// Required: true
	PackageID *string `json:"package_id"`

	// reseller id
	// Required: true
	ResellerID *string `json:"reseller_id"`

	// valid until
	// Required: true
	ValidUntil *string `json:"valid_until"`
}

// Validate validates this replace changeaspecific voucher request
func (m *ReplaceChangeaspecificVoucherRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackageID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResellerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidUntil(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReplaceChangeaspecificVoucherRequest) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificVoucherRequest) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificVoucherRequest) validateCustomerID(formats strfmt.Registry) error {

	if err := validate.Required("customer_id", "body", m.CustomerID); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificVoucherRequest) validatePackageID(formats strfmt.Registry) error {

	if err := validate.Required("package_id", "body", m.PackageID); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificVoucherRequest) validateResellerID(formats strfmt.Registry) error {

	if err := validate.Required("reseller_id", "body", m.ResellerID); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificVoucherRequest) validateValidUntil(formats strfmt.Registry) error {

	if err := validate.Required("valid_until", "body", m.ValidUntil); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReplaceChangeaspecificVoucherRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplaceChangeaspecificVoucherRequest) UnmarshalBinary(b []byte) error {
	var res ReplaceChangeaspecificVoucherRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
