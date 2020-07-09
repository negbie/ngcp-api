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

// TopupLogs TopupLogs
//
// swagger:model TopupLogs
type TopupLogs struct {

	// amount
	// Required: true
	Amount *float64 `json:"amount"`

	// cash balance after
	// Required: true
	CashBalanceAfter *float64 `json:"cash_balance_after"`

	// cash balance before
	// Required: true
	CashBalanceBefore *float64 `json:"cash_balance_before"`

	// contract balance after id
	// Required: true
	ContractBalanceAfterID *float64 `json:"contract_balance_after_id"`

	// contract balance before id
	// Required: true
	ContractBalanceBeforeID *float64 `json:"contract_balance_before_id"`

	// contract id
	// Required: true
	ContractID *float64 `json:"contract_id"`

	// lock level after
	// Required: true
	LockLevelAfter *string `json:"lock_level_after"`

	// lock level before
	// Required: true
	LockLevelBefore *string `json:"lock_level_before"`

	// message
	// Required: true
	Message *string `json:"message"`

	// outcome
	// Required: true
	Outcome *string `json:"outcome"`

	// package after id
	// Required: true
	PackageAfterID *float64 `json:"package_after_id"`

	// package before id
	// Required: true
	PackageBeforeID *float64 `json:"package_before_id"`

	// profile after id
	// Required: true
	ProfileAfterID *float64 `json:"profile_after_id"`

	// profile before id
	// Required: true
	ProfileBeforeID *float64 `json:"profile_before_id"`

	// request token
	// Required: true
	RequestToken *string `json:"request_token"`

	// subscriber id
	// Required: true
	SubscriberID *float64 `json:"subscriber_id"`

	// timestamp
	// Required: true
	Timestamp *string `json:"timestamp"`

	// type
	// Required: true
	Type *string `json:"type"`

	// username
	// Required: true
	Username *string `json:"username"`

	// voucher id
	// Required: true
	VoucherID *float64 `json:"voucher_id"`
}

// Validate validates this topup logs
func (m *TopupLogs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCashBalanceAfter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCashBalanceBefore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContractBalanceAfterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContractBalanceBeforeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContractID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLockLevelAfter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLockLevelBefore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutcome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackageAfterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackageBeforeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfileAfterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfileBeforeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVoucherID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopupLogs) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *TopupLogs) validateCashBalanceAfter(formats strfmt.Registry) error {

	if err := validate.Required("cash_balance_after", "body", m.CashBalanceAfter); err != nil {
		return err
	}

	return nil
}

func (m *TopupLogs) validateCashBalanceBefore(formats strfmt.Registry) error {

	if err := validate.Required("cash_balance_before", "body", m.CashBalanceBefore); err != nil {
		return err
	}

	return nil
}

func (m *TopupLogs) validateContractBalanceAfterID(formats strfmt.Registry) error {

	if err := validate.Required("contract_balance_after_id", "body", m.ContractBalanceAfterID); err != nil {
		return err
	}

	return nil
}

func (m *TopupLogs) validateContractBalanceBeforeID(formats strfmt.Registry) error {

	if err := validate.Required("contract_balance_before_id", "body", m.ContractBalanceBeforeID); err != nil {
		return err
	}

	return nil
}

func (m *TopupLogs) validateContractID(formats strfmt.Registry) error {

	if err := validate.Required("contract_id", "body", m.ContractID); err != nil {
		return err
	}

	return nil
}

func (m *TopupLogs) validateLockLevelAfter(formats strfmt.Registry) error {

	if err := validate.Required("lock_level_after", "body", m.LockLevelAfter); err != nil {
		return err
	}

	return nil
}

func (m *TopupLogs) validateLockLevelBefore(formats strfmt.Registry) error {

	if err := validate.Required("lock_level_before", "body", m.LockLevelBefore); err != nil {
		return err
	}

	return nil
}

func (m *TopupLogs) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *TopupLogs) validateOutcome(formats strfmt.Registry) error {

	if err := validate.Required("outcome", "body", m.Outcome); err != nil {
		return err
	}

	return nil
}

func (m *TopupLogs) validatePackageAfterID(formats strfmt.Registry) error {

	if err := validate.Required("package_after_id", "body", m.PackageAfterID); err != nil {
		return err
	}

	return nil
}

func (m *TopupLogs) validatePackageBeforeID(formats strfmt.Registry) error {

	if err := validate.Required("package_before_id", "body", m.PackageBeforeID); err != nil {
		return err
	}

	return nil
}

func (m *TopupLogs) validateProfileAfterID(formats strfmt.Registry) error {

	if err := validate.Required("profile_after_id", "body", m.ProfileAfterID); err != nil {
		return err
	}

	return nil
}

func (m *TopupLogs) validateProfileBeforeID(formats strfmt.Registry) error {

	if err := validate.Required("profile_before_id", "body", m.ProfileBeforeID); err != nil {
		return err
	}

	return nil
}

func (m *TopupLogs) validateRequestToken(formats strfmt.Registry) error {

	if err := validate.Required("request_token", "body", m.RequestToken); err != nil {
		return err
	}

	return nil
}

func (m *TopupLogs) validateSubscriberID(formats strfmt.Registry) error {

	if err := validate.Required("subscriber_id", "body", m.SubscriberID); err != nil {
		return err
	}

	return nil
}

func (m *TopupLogs) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

func (m *TopupLogs) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *TopupLogs) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

func (m *TopupLogs) validateVoucherID(formats strfmt.Registry) error {

	if err := validate.Required("voucher_id", "body", m.VoucherID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopupLogs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopupLogs) UnmarshalBinary(b []byte) error {
	var res TopupLogs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
