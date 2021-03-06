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

// ReplaceChangeaspecificEmailTemplateRequest Replace/changeaspecificEmailTemplateRequest
//
// swagger:model Replace/changeaspecificEmailTemplateRequest
type ReplaceChangeaspecificEmailTemplateRequest struct {

	// attachment name
	// Required: true
	AttachmentName *string `json:"attachment_name"`

	// body
	// Required: true
	Body *string `json:"body"`

	// from email
	// Required: true
	FromEmail *string `json:"from_email"`

	// name
	// Required: true
	Name *string `json:"name"`

	// reseller id
	// Required: true
	ResellerID *string `json:"reseller_id"`

	// subject
	// Required: true
	Subject *string `json:"subject"`
}

// Validate validates this replace changeaspecific email template request
func (m *ReplaceChangeaspecificEmailTemplateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachmentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFromEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResellerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReplaceChangeaspecificEmailTemplateRequest) validateAttachmentName(formats strfmt.Registry) error {

	if err := validate.Required("attachment_name", "body", m.AttachmentName); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificEmailTemplateRequest) validateBody(formats strfmt.Registry) error {

	if err := validate.Required("body", "body", m.Body); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificEmailTemplateRequest) validateFromEmail(formats strfmt.Registry) error {

	if err := validate.Required("from_email", "body", m.FromEmail); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificEmailTemplateRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificEmailTemplateRequest) validateResellerID(formats strfmt.Registry) error {

	if err := validate.Required("reseller_id", "body", m.ResellerID); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceChangeaspecificEmailTemplateRequest) validateSubject(formats strfmt.Registry) error {

	if err := validate.Required("subject", "body", m.Subject); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReplaceChangeaspecificEmailTemplateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplaceChangeaspecificEmailTemplateRequest) UnmarshalBinary(b []byte) error {
	var res ReplaceChangeaspecificEmailTemplateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
