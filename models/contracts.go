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

// Contracts Contracts
//
// swagger:model Contracts
type Contracts struct {

	// billing profile definition
	// Required: true
	BillingProfileDefinition *string `json:"billing_profile_definition"`

	// billing profile id
	// Required: true
	BillingProfileID *string `json:"billing_profile_id"`

	// billing profiles
	// Required: true
	BillingProfiles []*BillingProfiles3 `json:"billing_profiles"`

	// contact id
	// Required: true
	ContactID *string `json:"contact_id"`

	// external id
	// Required: true
	ExternalID *string `json:"external_id"`

	// status
	// Required: true
	Status *string `json:"status"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this contracts
func (m *Contracts) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingProfileDefinition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingProfileID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingProfiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Contracts) validateBillingProfileDefinition(formats strfmt.Registry) error {

	if err := validate.Required("billing_profile_definition", "body", m.BillingProfileDefinition); err != nil {
		return err
	}

	return nil
}

func (m *Contracts) validateBillingProfileID(formats strfmt.Registry) error {

	if err := validate.Required("billing_profile_id", "body", m.BillingProfileID); err != nil {
		return err
	}

	return nil
}

func (m *Contracts) validateBillingProfiles(formats strfmt.Registry) error {

	if err := validate.Required("billing_profiles", "body", m.BillingProfiles); err != nil {
		return err
	}

	for i := 0; i < len(m.BillingProfiles); i++ {
		if swag.IsZero(m.BillingProfiles[i]) { // not required
			continue
		}

		if m.BillingProfiles[i] != nil {
			if err := m.BillingProfiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("billing_profiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Contracts) validateContactID(formats strfmt.Registry) error {

	if err := validate.Required("contact_id", "body", m.ContactID); err != nil {
		return err
	}

	return nil
}

func (m *Contracts) validateExternalID(formats strfmt.Registry) error {

	if err := validate.Required("external_id", "body", m.ExternalID); err != nil {
		return err
	}

	return nil
}

func (m *Contracts) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Contracts) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Contracts) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Contracts) UnmarshalBinary(b []byte) error {
	var res Contracts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
