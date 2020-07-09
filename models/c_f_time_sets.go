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

// CFTimeSets CFTimeSets
//
// swagger:model CFTimeSets
type CFTimeSets struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// subscriber id
	// Required: true
	SubscriberID *string `json:"subscriber_id"`

	// times
	// Required: true
	Times []*Time `json:"times"`
}

// Validate validates this c f time sets
func (m *CFTimeSets) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CFTimeSets) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CFTimeSets) validateSubscriberID(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_id", "body", m.SubscriberID); err != nil {
		return err
	}

	return nil
}

func (m *CFTimeSets) validateTimes(formats strfmt.Registry) error {

	if err := validate.Required("times", "body", m.Times); err != nil {
		return err
	}

	for i := 0; i < len(m.Times); i++ {
		if swag.IsZero(m.Times[i]) { // not required
			continue
		}

		if m.Times[i] != nil {
			if err := m.Times[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("times" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CFTimeSets) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CFTimeSets) UnmarshalBinary(b []byte) error {
	var res CFTimeSets
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
