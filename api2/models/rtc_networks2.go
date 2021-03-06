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

// RtcNetworks2 RtcNetworks2
//
// swagger:model RtcNetworks2
type RtcNetworks2 struct {

	// networks
	// Required: true
	Networks []*Network `json:"networks"`

	// reseller id
	// Required: true
	ResellerID *string `json:"reseller_id"`

	// rtc user id
	// Required: true
	RtcUserID *string `json:"rtc_user_id"`
}

// Validate validates this rtc networks2
func (m *RtcNetworks2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResellerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRtcUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RtcNetworks2) validateNetworks(formats strfmt.Registry) error {

	if err := validate.Required("networks", "body", m.Networks); err != nil {
		return err
	}

	for i := 0; i < len(m.Networks); i++ {
		if swag.IsZero(m.Networks[i]) { // not required
			continue
		}

		if m.Networks[i] != nil {
			if err := m.Networks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RtcNetworks2) validateResellerID(formats strfmt.Registry) error {

	if err := validate.Required("reseller_id", "body", m.ResellerID); err != nil {
		return err
	}

	return nil
}

func (m *RtcNetworks2) validateRtcUserID(formats strfmt.Registry) error {

	if err := validate.Required("rtc_user_id", "body", m.RtcUserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RtcNetworks2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RtcNetworks2) UnmarshalBinary(b []byte) error {
	var res RtcNetworks2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
