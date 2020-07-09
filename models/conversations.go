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

// Conversations Conversations
//
// swagger:model Conversations
type Conversations struct {

	// call id
	// Required: true
	CallID *string `json:"call_id"`

	// call type
	// Required: true
	CallType *string `json:"call_type"`

	// callee
	// Required: true
	Callee *string `json:"callee"`

	// caller
	// Required: true
	Caller *string `json:"caller"`

	// context
	// Required: true
	Context *string `json:"context"`

	// currency
	// Required: true
	Currency *string `json:"currency"`

	// customer cost
	// Required: true
	CustomerCost *float64 `json:"customer_cost"`

	// direction
	// Required: true
	Direction *string `json:"direction"`

	// duration
	// Required: true
	Duration *string `json:"duration"`

	// filename
	// Required: true
	Filename *string `json:"filename"`

	// folder
	// Required: true
	Folder *string `json:"folder"`

	// id
	// Required: true
	ID *float64 `json:"id"`

	// pages
	// Required: true
	Pages *float64 `json:"pages"`

	// rating status
	// Required: true
	RatingStatus *string `json:"rating_status"`

	// start time
	// Required: true
	StartTime *string `json:"start_time"`

	// status
	// Required: true
	Status *string `json:"status"`

	// subscriber id
	// Required: true
	SubscriberID *float64 `json:"subscriber_id"`

	// total customer cost
	// Required: true
	TotalCustomerCost *float64 `json:"total_customer_cost"`

	// type
	// Required: true
	Type *string `json:"type"`

	// voicemail subscriber id
	// Required: true
	VoicemailSubscriberID *float64 `json:"voicemail_subscriber_id"`
}

// Validate validates this conversations
func (m *Conversations) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCallType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCallee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCaller(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilename(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFolder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRatingStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCustomerCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVoicemailSubscriberID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Conversations) validateCallID(formats strfmt.Registry) error {

	if err := validate.Required("call_id", "body", m.CallID); err != nil {
		return err
	}

	return nil
}

func (m *Conversations) validateCallType(formats strfmt.Registry) error {

	if err := validate.Required("call_type", "body", m.CallType); err != nil {
		return err
	}

	return nil
}

func (m *Conversations) validateCallee(formats strfmt.Registry) error {

	if err := validate.Required("callee", "body", m.Callee); err != nil {
		return err
	}

	return nil
}

func (m *Conversations) validateCaller(formats strfmt.Registry) error {

	if err := validate.Required("caller", "body", m.Caller); err != nil {
		return err
	}

	return nil
}

func (m *Conversations) validateContext(formats strfmt.Registry) error {

	if err := validate.Required("context", "body", m.Context); err != nil {
		return err
	}

	return nil
}

func (m *Conversations) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("currency", "body", m.Currency); err != nil {
		return err
	}

	return nil
}

func (m *Conversations) validateCustomerCost(formats strfmt.Registry) error {

	if err := validate.Required("customer_cost", "body", m.CustomerCost); err != nil {
		return err
	}

	return nil
}

func (m *Conversations) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *Conversations) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *Conversations) validateFilename(formats strfmt.Registry) error {

	if err := validate.Required("filename", "body", m.Filename); err != nil {
		return err
	}

	return nil
}

func (m *Conversations) validateFolder(formats strfmt.Registry) error {

	if err := validate.Required("folder", "body", m.Folder); err != nil {
		return err
	}

	return nil
}

func (m *Conversations) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Conversations) validatePages(formats strfmt.Registry) error {

	if err := validate.Required("pages", "body", m.Pages); err != nil {
		return err
	}

	return nil
}

func (m *Conversations) validateRatingStatus(formats strfmt.Registry) error {

	if err := validate.Required("rating_status", "body", m.RatingStatus); err != nil {
		return err
	}

	return nil
}

func (m *Conversations) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("start_time", "body", m.StartTime); err != nil {
		return err
	}

	return nil
}

func (m *Conversations) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Conversations) validateSubscriberID(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_id", "body", m.SubscriberID); err != nil {
		return err
	}

	return nil
}

func (m *Conversations) validateTotalCustomerCost(formats strfmt.Registry) error {

	if err := validate.Required("total_customer_cost", "body", m.TotalCustomerCost); err != nil {
		return err
	}

	return nil
}

func (m *Conversations) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *Conversations) validateVoicemailSubscriberID(formats strfmt.Registry) error {

	if err := validate.Required("voicemail_subscriber_id", "body", m.VoicemailSubscriberID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Conversations) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Conversations) UnmarshalBinary(b []byte) error {
	var res Conversations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
