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

// BillingProfiles2 BillingProfiles2
//
// swagger:model BillingProfiles2
type BillingProfiles2 struct {

	// advice of charge
	// Required: true
	AdviceOfCharge *bool `json:"advice_of_charge"`

	// currency
	// Required: true
	Currency *string `json:"currency"`

	// fraud daily limit
	// Required: true
	FraudDailyLimit *float64 `json:"fraud_daily_limit"`

	// fraud daily lock
	// Required: true
	FraudDailyLock *string `json:"fraud_daily_lock"`

	// fraud daily notify
	// Required: true
	FraudDailyNotify *string `json:"fraud_daily_notify"`

	// fraud interval limit
	// Required: true
	FraudIntervalLimit *float64 `json:"fraud_interval_limit"`

	// fraud interval lock
	// Required: true
	FraudIntervalLock *string `json:"fraud_interval_lock"`

	// fraud interval notify
	// Required: true
	FraudIntervalNotify *string `json:"fraud_interval_notify"`

	// fraud use reseller rates
	// Required: true
	FraudUseResellerRates *bool `json:"fraud_use_reseller_rates"`

	// handle
	// Required: true
	Handle *string `json:"handle"`

	// interval charge
	// Required: true
	IntervalCharge *float64 `json:"interval_charge"`

	// interval free cash
	// Required: true
	IntervalFreeCash *float64 `json:"interval_free_cash"`

	// interval free time
	// Required: true
	IntervalFreeTime *float64 `json:"interval_free_time"`

	// name
	// Required: true
	Name *string `json:"name"`

	// peaktime special
	// Required: true
	PeaktimeSpecial []*PeaktimeSpecial `json:"peaktime_special"`

	// peaktime weekdays
	// Required: true
	PeaktimeWeekdays []*PeaktimeWeekday3 `json:"peaktime_weekdays"`

	// prepaid
	// Required: true
	Prepaid *bool `json:"prepaid"`

	// prepaid library
	// Required: true
	PrepaidLibrary *string `json:"prepaid_library"`

	// reseller id
	// Required: true
	ResellerID *float64 `json:"reseller_id"`
}

// Validate validates this billing profiles2
func (m *BillingProfiles2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdviceOfCharge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFraudDailyLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFraudDailyLock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFraudDailyNotify(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFraudIntervalLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFraudIntervalLock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFraudIntervalNotify(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFraudUseResellerRates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHandle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntervalCharge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntervalFreeCash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntervalFreeTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeaktimeSpecial(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeaktimeWeekdays(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrepaid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrepaidLibrary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResellerID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BillingProfiles2) validateAdviceOfCharge(formats strfmt.Registry) error {

	if err := validate.Required("advice_of_charge", "body", m.AdviceOfCharge); err != nil {
		return err
	}

	return nil
}

func (m *BillingProfiles2) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("currency", "body", m.Currency); err != nil {
		return err
	}

	return nil
}

func (m *BillingProfiles2) validateFraudDailyLimit(formats strfmt.Registry) error {

	if err := validate.Required("fraud_daily_limit", "body", m.FraudDailyLimit); err != nil {
		return err
	}

	return nil
}

func (m *BillingProfiles2) validateFraudDailyLock(formats strfmt.Registry) error {

	if err := validate.Required("fraud_daily_lock", "body", m.FraudDailyLock); err != nil {
		return err
	}

	return nil
}

func (m *BillingProfiles2) validateFraudDailyNotify(formats strfmt.Registry) error {

	if err := validate.Required("fraud_daily_notify", "body", m.FraudDailyNotify); err != nil {
		return err
	}

	return nil
}

func (m *BillingProfiles2) validateFraudIntervalLimit(formats strfmt.Registry) error {

	if err := validate.Required("fraud_interval_limit", "body", m.FraudIntervalLimit); err != nil {
		return err
	}

	return nil
}

func (m *BillingProfiles2) validateFraudIntervalLock(formats strfmt.Registry) error {

	if err := validate.Required("fraud_interval_lock", "body", m.FraudIntervalLock); err != nil {
		return err
	}

	return nil
}

func (m *BillingProfiles2) validateFraudIntervalNotify(formats strfmt.Registry) error {

	if err := validate.Required("fraud_interval_notify", "body", m.FraudIntervalNotify); err != nil {
		return err
	}

	return nil
}

func (m *BillingProfiles2) validateFraudUseResellerRates(formats strfmt.Registry) error {

	if err := validate.Required("fraud_use_reseller_rates", "body", m.FraudUseResellerRates); err != nil {
		return err
	}

	return nil
}

func (m *BillingProfiles2) validateHandle(formats strfmt.Registry) error {

	if err := validate.Required("handle", "body", m.Handle); err != nil {
		return err
	}

	return nil
}

func (m *BillingProfiles2) validateIntervalCharge(formats strfmt.Registry) error {

	if err := validate.Required("interval_charge", "body", m.IntervalCharge); err != nil {
		return err
	}

	return nil
}

func (m *BillingProfiles2) validateIntervalFreeCash(formats strfmt.Registry) error {

	if err := validate.Required("interval_free_cash", "body", m.IntervalFreeCash); err != nil {
		return err
	}

	return nil
}

func (m *BillingProfiles2) validateIntervalFreeTime(formats strfmt.Registry) error {

	if err := validate.Required("interval_free_time", "body", m.IntervalFreeTime); err != nil {
		return err
	}

	return nil
}

func (m *BillingProfiles2) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *BillingProfiles2) validatePeaktimeSpecial(formats strfmt.Registry) error {

	if err := validate.Required("peaktime_special", "body", m.PeaktimeSpecial); err != nil {
		return err
	}

	for i := 0; i < len(m.PeaktimeSpecial); i++ {
		if swag.IsZero(m.PeaktimeSpecial[i]) { // not required
			continue
		}

		if m.PeaktimeSpecial[i] != nil {
			if err := m.PeaktimeSpecial[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("peaktime_special" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BillingProfiles2) validatePeaktimeWeekdays(formats strfmt.Registry) error {

	if err := validate.Required("peaktime_weekdays", "body", m.PeaktimeWeekdays); err != nil {
		return err
	}

	for i := 0; i < len(m.PeaktimeWeekdays); i++ {
		if swag.IsZero(m.PeaktimeWeekdays[i]) { // not required
			continue
		}

		if m.PeaktimeWeekdays[i] != nil {
			if err := m.PeaktimeWeekdays[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("peaktime_weekdays" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BillingProfiles2) validatePrepaid(formats strfmt.Registry) error {

	if err := validate.Required("prepaid", "body", m.Prepaid); err != nil {
		return err
	}

	return nil
}

func (m *BillingProfiles2) validatePrepaidLibrary(formats strfmt.Registry) error {

	if err := validate.Required("prepaid_library", "body", m.PrepaidLibrary); err != nil {
		return err
	}

	return nil
}

func (m *BillingProfiles2) validateResellerID(formats strfmt.Registry) error {

	if err := validate.Required("reseller_id", "body", m.ResellerID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BillingProfiles2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingProfiles2) UnmarshalBinary(b []byte) error {
	var res BillingProfiles2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}