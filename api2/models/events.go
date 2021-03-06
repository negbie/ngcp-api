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

// Events Events
//
// swagger:model Events
type Events struct {

	// export status
	// Required: true
	ExportStatus *string `json:"export_status"`

	// exported at
	// Required: true
	ExportedAt *string `json:"exported_at"`

	// first non primary alias username after
	// Required: true
	FirstNonPrimaryAliasUsernameAfter *string `json:"first_non_primary_alias_username_after"`

	// first non primary alias username before
	// Required: true
	FirstNonPrimaryAliasUsernameBefore *string `json:"first_non_primary_alias_username_before"`

	// new status
	// Required: true
	NewStatus *string `json:"new_status"`

	// non primary alias username
	// Required: true
	NonPrimaryAliasUsername *string `json:"non_primary_alias_username"`

	// old status
	// Required: true
	OldStatus *string `json:"old_status"`

	// pilot first non primary alias username after
	// Required: true
	PilotFirstNonPrimaryAliasUsernameAfter *string `json:"pilot_first_non_primary_alias_username_after"`

	// pilot first non primary alias username before
	// Required: true
	PilotFirstNonPrimaryAliasUsernameBefore *string `json:"pilot_first_non_primary_alias_username_before"`

	// pilot primary alias username after
	// Required: true
	PilotPrimaryAliasUsernameAfter *string `json:"pilot_primary_alias_username_after"`

	// pilot primary alias username before
	// Required: true
	PilotPrimaryAliasUsernameBefore *string `json:"pilot_primary_alias_username_before"`

	// pilot primary number ac
	// Required: true
	PilotPrimaryNumberAc *string `json:"pilot_primary_number_ac"`

	// pilot primary number cc
	// Required: true
	PilotPrimaryNumberCc *string `json:"pilot_primary_number_cc"`

	// pilot primary number id
	// Required: true
	PilotPrimaryNumberID *float64 `json:"pilot_primary_number_id"`

	// pilot primary number sn
	// Required: true
	PilotPrimaryNumberSn *string `json:"pilot_primary_number_sn"`

	// pilot subscriber id
	// Required: true
	PilotSubscriberID *float64 `json:"pilot_subscriber_id"`

	// pilot subscriber profile id
	// Required: true
	PilotSubscriberProfileID *float64 `json:"pilot_subscriber_profile_id"`

	// pilot subscriber profile name
	// Required: true
	PilotSubscriberProfileName *string `json:"pilot_subscriber_profile_name"`

	// pilot subscriber profile set id
	// Required: true
	PilotSubscriberProfileSetID *float64 `json:"pilot_subscriber_profile_set_id"`

	// pilot subscriber profile set name
	// Required: true
	PilotSubscriberProfileSetName *string `json:"pilot_subscriber_profile_set_name"`

	// primary alias username after
	// Required: true
	PrimaryAliasUsernameAfter *string `json:"primary_alias_username_after"`

	// primary alias username before
	// Required: true
	PrimaryAliasUsernameBefore *string `json:"primary_alias_username_before"`

	// primary number ac
	// Required: true
	PrimaryNumberAc *string `json:"primary_number_ac"`

	// primary number cc
	// Required: true
	PrimaryNumberCc *string `json:"primary_number_cc"`

	// primary number id
	// Required: true
	PrimaryNumberID *float64 `json:"primary_number_id"`

	// primary number sn
	// Required: true
	PrimaryNumberSn *string `json:"primary_number_sn"`

	// reseller id
	// Required: true
	ResellerID *float64 `json:"reseller_id"`

	// subscriber id
	// Required: true
	SubscriberID *float64 `json:"subscriber_id"`

	// subscriber profile id
	// Required: true
	SubscriberProfileID *float64 `json:"subscriber_profile_id"`

	// subscriber profile name
	// Required: true
	SubscriberProfileName *string `json:"subscriber_profile_name"`

	// subscriber profile set id
	// Required: true
	SubscriberProfileSetID *float64 `json:"subscriber_profile_set_id"`

	// subscriber profile set name
	// Required: true
	SubscriberProfileSetName *string `json:"subscriber_profile_set_name"`

	// timestamp
	// Required: true
	Timestamp *string `json:"timestamp"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this events
func (m *Events) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExportStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExportedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstNonPrimaryAliasUsernameAfter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstNonPrimaryAliasUsernameBefore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNonPrimaryAliasUsername(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOldStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotFirstNonPrimaryAliasUsernameAfter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotFirstNonPrimaryAliasUsernameBefore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotPrimaryAliasUsernameAfter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotPrimaryAliasUsernameBefore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotPrimaryNumberAc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotPrimaryNumberCc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotPrimaryNumberID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotPrimaryNumberSn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotSubscriberID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotSubscriberProfileID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotSubscriberProfileName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotSubscriberProfileSetID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotSubscriberProfileSetName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryAliasUsernameAfter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryAliasUsernameBefore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryNumberAc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryNumberCc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryNumberID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryNumberSn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResellerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberProfileID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberProfileName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberProfileSetID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberProfileSetName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
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

func (m *Events) validateExportStatus(formats strfmt.Registry) error {

	if err := validate.Required("export_status", "body", m.ExportStatus); err != nil {
		return err
	}

	return nil
}

func (m *Events) validateExportedAt(formats strfmt.Registry) error {

	if err := validate.Required("exported_at", "body", m.ExportedAt); err != nil {
		return err
	}

	return nil
}

func (m *Events) validateFirstNonPrimaryAliasUsernameAfter(formats strfmt.Registry) error {

	if err := validate.Required("first_non_primary_alias_username_after", "body", m.FirstNonPrimaryAliasUsernameAfter); err != nil {
		return err
	}

	return nil
}

func (m *Events) validateFirstNonPrimaryAliasUsernameBefore(formats strfmt.Registry) error {

	if err := validate.Required("first_non_primary_alias_username_before", "body", m.FirstNonPrimaryAliasUsernameBefore); err != nil {
		return err
	}

	return nil
}

func (m *Events) validateNewStatus(formats strfmt.Registry) error {

	if err := validate.Required("new_status", "body", m.NewStatus); err != nil {
		return err
	}

	return nil
}

func (m *Events) validateNonPrimaryAliasUsername(formats strfmt.Registry) error {

	if err := validate.Required("non_primary_alias_username", "body", m.NonPrimaryAliasUsername); err != nil {
		return err
	}

	return nil
}

func (m *Events) validateOldStatus(formats strfmt.Registry) error {

	if err := validate.Required("old_status", "body", m.OldStatus); err != nil {
		return err
	}

	return nil
}

func (m *Events) validatePilotFirstNonPrimaryAliasUsernameAfter(formats strfmt.Registry) error {

	if err := validate.Required("pilot_first_non_primary_alias_username_after", "body", m.PilotFirstNonPrimaryAliasUsernameAfter); err != nil {
		return err
	}

	return nil
}

func (m *Events) validatePilotFirstNonPrimaryAliasUsernameBefore(formats strfmt.Registry) error {

	if err := validate.Required("pilot_first_non_primary_alias_username_before", "body", m.PilotFirstNonPrimaryAliasUsernameBefore); err != nil {
		return err
	}

	return nil
}

func (m *Events) validatePilotPrimaryAliasUsernameAfter(formats strfmt.Registry) error {

	if err := validate.Required("pilot_primary_alias_username_after", "body", m.PilotPrimaryAliasUsernameAfter); err != nil {
		return err
	}

	return nil
}

func (m *Events) validatePilotPrimaryAliasUsernameBefore(formats strfmt.Registry) error {

	if err := validate.Required("pilot_primary_alias_username_before", "body", m.PilotPrimaryAliasUsernameBefore); err != nil {
		return err
	}

	return nil
}

func (m *Events) validatePilotPrimaryNumberAc(formats strfmt.Registry) error {

	if err := validate.Required("pilot_primary_number_ac", "body", m.PilotPrimaryNumberAc); err != nil {
		return err
	}

	return nil
}

func (m *Events) validatePilotPrimaryNumberCc(formats strfmt.Registry) error {

	if err := validate.Required("pilot_primary_number_cc", "body", m.PilotPrimaryNumberCc); err != nil {
		return err
	}

	return nil
}

func (m *Events) validatePilotPrimaryNumberID(formats strfmt.Registry) error {

	if err := validate.Required("pilot_primary_number_id", "body", m.PilotPrimaryNumberID); err != nil {
		return err
	}

	return nil
}

func (m *Events) validatePilotPrimaryNumberSn(formats strfmt.Registry) error {

	if err := validate.Required("pilot_primary_number_sn", "body", m.PilotPrimaryNumberSn); err != nil {
		return err
	}

	return nil
}

func (m *Events) validatePilotSubscriberID(formats strfmt.Registry) error {

	if err := validate.Required("pilot_subscriber_id", "body", m.PilotSubscriberID); err != nil {
		return err
	}

	return nil
}

func (m *Events) validatePilotSubscriberProfileID(formats strfmt.Registry) error {

	if err := validate.Required("pilot_subscriber_profile_id", "body", m.PilotSubscriberProfileID); err != nil {
		return err
	}

	return nil
}

func (m *Events) validatePilotSubscriberProfileName(formats strfmt.Registry) error {

	if err := validate.Required("pilot_subscriber_profile_name", "body", m.PilotSubscriberProfileName); err != nil {
		return err
	}

	return nil
}

func (m *Events) validatePilotSubscriberProfileSetID(formats strfmt.Registry) error {

	if err := validate.Required("pilot_subscriber_profile_set_id", "body", m.PilotSubscriberProfileSetID); err != nil {
		return err
	}

	return nil
}

func (m *Events) validatePilotSubscriberProfileSetName(formats strfmt.Registry) error {

	if err := validate.Required("pilot_subscriber_profile_set_name", "body", m.PilotSubscriberProfileSetName); err != nil {
		return err
	}

	return nil
}

func (m *Events) validatePrimaryAliasUsernameAfter(formats strfmt.Registry) error {

	if err := validate.Required("primary_alias_username_after", "body", m.PrimaryAliasUsernameAfter); err != nil {
		return err
	}

	return nil
}

func (m *Events) validatePrimaryAliasUsernameBefore(formats strfmt.Registry) error {

	if err := validate.Required("primary_alias_username_before", "body", m.PrimaryAliasUsernameBefore); err != nil {
		return err
	}

	return nil
}

func (m *Events) validatePrimaryNumberAc(formats strfmt.Registry) error {

	if err := validate.Required("primary_number_ac", "body", m.PrimaryNumberAc); err != nil {
		return err
	}

	return nil
}

func (m *Events) validatePrimaryNumberCc(formats strfmt.Registry) error {

	if err := validate.Required("primary_number_cc", "body", m.PrimaryNumberCc); err != nil {
		return err
	}

	return nil
}

func (m *Events) validatePrimaryNumberID(formats strfmt.Registry) error {

	if err := validate.Required("primary_number_id", "body", m.PrimaryNumberID); err != nil {
		return err
	}

	return nil
}

func (m *Events) validatePrimaryNumberSn(formats strfmt.Registry) error {

	if err := validate.Required("primary_number_sn", "body", m.PrimaryNumberSn); err != nil {
		return err
	}

	return nil
}

func (m *Events) validateResellerID(formats strfmt.Registry) error {

	if err := validate.Required("reseller_id", "body", m.ResellerID); err != nil {
		return err
	}

	return nil
}

func (m *Events) validateSubscriberID(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_id", "body", m.SubscriberID); err != nil {
		return err
	}

	return nil
}

func (m *Events) validateSubscriberProfileID(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_profile_id", "body", m.SubscriberProfileID); err != nil {
		return err
	}

	return nil
}

func (m *Events) validateSubscriberProfileName(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_profile_name", "body", m.SubscriberProfileName); err != nil {
		return err
	}

	return nil
}

func (m *Events) validateSubscriberProfileSetID(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_profile_set_id", "body", m.SubscriberProfileSetID); err != nil {
		return err
	}

	return nil
}

func (m *Events) validateSubscriberProfileSetName(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_profile_set_name", "body", m.SubscriberProfileSetName); err != nil {
		return err
	}

	return nil
}

func (m *Events) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

func (m *Events) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Events) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Events) UnmarshalBinary(b []byte) error {
	var res Events
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
