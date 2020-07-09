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

// ReplaceChangeaspecificCFBNumberSetRequest Replace/changeaspecificCFBNumberSetRequest
//
// swagger:model Replace/changeaspecificCFBNumberSetRequest
type ReplaceChangeaspecificCFBNumberSetRequest struct {

	// bnumbers
	// Required: true
	Bnumbers []*Bnumber `json:"bnumbers"`

	// is regex
	// Required: true
	IsRegex *string `json:"is_regex"`

	// mode
	// Required: true
	Mode *string `json:"mode"`

	// name
	// Required: true
	Name *string `json:"name"`

	// subscriber id
	// Required: true
	SubscriberID *string `json:"subscriber_id"`
}

// Validate validates this replace changeaspecific c f b number set request
func (m *ReplaceChangeaspecificCFBNumberSetRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBnumbers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsRegex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

func (m *ReplaceChangeaspecificCFBNumberSetRequest) validateBnumbers(formats strfmt.Registry) error {

	if err := validate.Required("bnumbers", "body", m.Bnumbers); err != nil {
		return err
	}

	for i := 0; i < len(m.Bnumbers); i++ {
		if swag.IsZero(m.Bnumbers[i]) { // not required
			continue
		}

		if m.Bnumbers[i] != nil {
			if err := m.Bnumbers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bnumbers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReplaceChangeaspecificCFBNumberSetRequest) validateIsRegex(formats strfmt.Registry) error {

	if err := validate.Required("is_regex", "body", m.IsRegex); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificCFBNumberSetRequest) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificCFBNumberSetRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificCFBNumberSetRequest) validateSubscriberID(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_id", "body", m.SubscriberID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReplaceChangeaspecificCFBNumberSetRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplaceChangeaspecificCFBNumberSetRequest) UnmarshalBinary(b []byte) error {
	var res ReplaceChangeaspecificCFBNumberSetRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
