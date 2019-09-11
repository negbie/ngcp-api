// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Thenewlycreateditemorempty36 Thenewlycreateditemorempty36
// swagger:model Thenewlycreateditemorempty36
type Thenewlycreateditemorempty36 struct {

	// is regex
	// Required: true
	IsRegex *string `json:"is_regex"`

	// mode
	// Required: true
	Mode *string `json:"mode"`

	// name
	// Required: true
	Name *string `json:"name"`

	// sources
	// Required: true
	Sources []*Source `json:"sources"`

	// subscriber id
	// Required: true
	SubscriberID *string `json:"subscriber_id"`
}

// Validate validates this thenewlycreateditemorempty36
func (m *Thenewlycreateditemorempty36) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsRegex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSources(formats); err != nil {
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

func (m *Thenewlycreateditemorempty36) validateIsRegex(formats strfmt.Registry) error {

	if err := validate.Required("is_regex", "body", m.IsRegex); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty36) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty36) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Thenewlycreateditemorempty36) validateSources(formats strfmt.Registry) error {

	if err := validate.Required("sources", "body", m.Sources); err != nil {
		return err
	}

	for i := 0; i < len(m.Sources); i++ {
		if swag.IsZero(m.Sources[i]) { // not required
			continue
		}

		if m.Sources[i] != nil {
			if err := m.Sources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Thenewlycreateditemorempty36) validateSubscriberID(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_id", "body", m.SubscriberID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Thenewlycreateditemorempty36) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Thenewlycreateditemorempty36) UnmarshalBinary(b []byte) error {
	var res Thenewlycreateditemorempty36
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}