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

// Call Call
//
// swagger:model Call
type Call struct {

	// call code
	// Required: true
	CallCode *string `json:"call_code"`

	// call id
	// Required: true
	CallID *string `json:"call_id"`

	// call status
	// Required: true
	CallStatus *string `json:"call_status"`

	// call type
	// Required: true
	CallType *string `json:"call_type"`

	// destination carrier billing fee id
	// Required: true
	DestinationCarrierBillingFeeID *float64 `json:"destination_carrier_billing_fee_id"`

	// destination carrier billing zone id
	// Required: true
	DestinationCarrierBillingZoneID *float64 `json:"destination_carrier_billing_zone_id"`

	// destination carrier cost
	// Required: true
	DestinationCarrierCost *float64 `json:"destination_carrier_cost"`

	// destination carrier free time
	// Required: true
	DestinationCarrierFreeTime *float64 `json:"destination_carrier_free_time"`

	// destination customer billing fee id
	// Required: true
	DestinationCustomerBillingFeeID *float64 `json:"destination_customer_billing_fee_id"`

	// destination customer billing zone id
	// Required: true
	DestinationCustomerBillingZoneID *float64 `json:"destination_customer_billing_zone_id"`

	// destination customer cost
	// Required: true
	DestinationCustomerCost *float64 `json:"destination_customer_cost"`

	// destination customer free time
	// Required: true
	DestinationCustomerFreeTime *float64 `json:"destination_customer_free_time"`

	// destination customer id
	// Required: true
	DestinationCustomerID *string `json:"destination_customer_id"`

	// destination domain
	// Required: true
	DestinationDomain *string `json:"destination_domain"`

	// destination domain in
	// Required: true
	DestinationDomainIn *string `json:"destination_domain_in"`

	// destination external contract id
	// Required: true
	DestinationExternalContractID *string `json:"destination_external_contract_id"`

	// destination external subscriber id
	// Required: true
	DestinationExternalSubscriberID *string `json:"destination_external_subscriber_id"`

	// destination gpp0
	// Required: true
	DestinationGpp0 *string `json:"destination_gpp0"`

	// destination gpp1
	// Required: true
	DestinationGpp1 *string `json:"destination_gpp1"`

	// destination gpp2
	// Required: true
	DestinationGpp2 *string `json:"destination_gpp2"`

	// destination gpp3
	// Required: true
	DestinationGpp3 *string `json:"destination_gpp3"`

	// destination gpp4
	// Required: true
	DestinationGpp4 *string `json:"destination_gpp4"`

	// destination gpp5
	// Required: true
	DestinationGpp5 *string `json:"destination_gpp5"`

	// destination gpp6
	// Required: true
	DestinationGpp6 *string `json:"destination_gpp6"`

	// destination gpp7
	// Required: true
	DestinationGpp7 *string `json:"destination_gpp7"`

	// destination gpp8
	// Required: true
	DestinationGpp8 *string `json:"destination_gpp8"`

	// destination gpp9
	// Required: true
	DestinationGpp9 *string `json:"destination_gpp9"`

	// destination lnp prefix
	// Required: true
	DestinationLnpPrefix *string `json:"destination_lnp_prefix"`

	// destination provider id
	// Required: true
	DestinationProviderID *string `json:"destination_provider_id"`

	// destination reseller billing fee id
	// Required: true
	DestinationResellerBillingFeeID *float64 `json:"destination_reseller_billing_fee_id"`

	// destination reseller billing zone id
	// Required: true
	DestinationResellerBillingZoneID *float64 `json:"destination_reseller_billing_zone_id"`

	// destination reseller cost
	// Required: true
	DestinationResellerCost *float64 `json:"destination_reseller_cost"`

	// destination reseller free time
	// Required: true
	DestinationResellerFreeTime *float64 `json:"destination_reseller_free_time"`

	// destination user
	// Required: true
	DestinationUser *string `json:"destination_user"`

	// destination user dialed
	// Required: true
	DestinationUserDialed *string `json:"destination_user_dialed"`

	// destination user id
	// Required: true
	DestinationUserID *string `json:"destination_user_id"`

	// destination user in
	// Required: true
	DestinationUserIn *string `json:"destination_user_in"`

	// destination user out
	// Required: true
	DestinationUserOut *string `json:"destination_user_out"`

	// duration
	// Required: true
	Duration *float64 `json:"duration"`

	// export status
	// Required: true
	ExportStatus *string `json:"export_status"`

	// exported at
	// Required: true
	ExportedAt *string `json:"exported_at"`

	// init time
	// Required: true
	InitTime *string `json:"init_time"`

	// peer auth realm
	// Required: true
	PeerAuthRealm *string `json:"peer_auth_realm"`

	// peer auth user
	// Required: true
	PeerAuthUser *string `json:"peer_auth_user"`

	// rated at
	// Required: true
	RatedAt *string `json:"rated_at"`

	// rating status
	// Required: true
	RatingStatus *string `json:"rating_status"`

	// source carrier billing fee id
	// Required: true
	SourceCarrierBillingFeeID *float64 `json:"source_carrier_billing_fee_id"`

	// source carrier billing zone id
	// Required: true
	SourceCarrierBillingZoneID *float64 `json:"source_carrier_billing_zone_id"`

	// source carrier cost
	// Required: true
	SourceCarrierCost *float64 `json:"source_carrier_cost"`

	// source carrier free time
	// Required: true
	SourceCarrierFreeTime *float64 `json:"source_carrier_free_time"`

	// source cli
	// Required: true
	SourceCli *string `json:"source_cli"`

	// source clir
	// Required: true
	SourceClir *bool `json:"source_clir"`

	// source customer billing fee id
	// Required: true
	SourceCustomerBillingFeeID *float64 `json:"source_customer_billing_fee_id"`

	// source customer billing zone id
	// Required: true
	SourceCustomerBillingZoneID *float64 `json:"source_customer_billing_zone_id"`

	// source customer cost
	// Required: true
	SourceCustomerCost *float64 `json:"source_customer_cost"`

	// source customer free time
	// Required: true
	SourceCustomerFreeTime *float64 `json:"source_customer_free_time"`

	// source customer id
	// Required: true
	SourceCustomerID *string `json:"source_customer_id"`

	// source domain
	// Required: true
	SourceDomain *string `json:"source_domain"`

	// source external contract id
	// Required: true
	SourceExternalContractID *string `json:"source_external_contract_id"`

	// source external subscriber id
	// Required: true
	SourceExternalSubscriberID *string `json:"source_external_subscriber_id"`

	// source gpp0
	// Required: true
	SourceGpp0 *string `json:"source_gpp0"`

	// source gpp1
	// Required: true
	SourceGpp1 *string `json:"source_gpp1"`

	// source gpp2
	// Required: true
	SourceGpp2 *string `json:"source_gpp2"`

	// source gpp3
	// Required: true
	SourceGpp3 *string `json:"source_gpp3"`

	// source gpp4
	// Required: true
	SourceGpp4 *string `json:"source_gpp4"`

	// source gpp5
	// Required: true
	SourceGpp5 *string `json:"source_gpp5"`

	// source gpp6
	// Required: true
	SourceGpp6 *string `json:"source_gpp6"`

	// source gpp7
	// Required: true
	SourceGpp7 *string `json:"source_gpp7"`

	// source gpp8
	// Required: true
	SourceGpp8 *string `json:"source_gpp8"`

	// source gpp9
	// Required: true
	SourceGpp9 *string `json:"source_gpp9"`

	// source ip
	// Required: true
	SourceIP *string `json:"source_ip"`

	// source lnp prefix
	// Required: true
	SourceLnpPrefix *string `json:"source_lnp_prefix"`

	// source provider id
	// Required: true
	SourceProviderID *string `json:"source_provider_id"`

	// source reseller billing fee id
	// Required: true
	SourceResellerBillingFeeID *float64 `json:"source_reseller_billing_fee_id"`

	// source reseller billing zone id
	// Required: true
	SourceResellerBillingZoneID *float64 `json:"source_reseller_billing_zone_id"`

	// source reseller cost
	// Required: true
	SourceResellerCost *float64 `json:"source_reseller_cost"`

	// source reseller free time
	// Required: true
	SourceResellerFreeTime *float64 `json:"source_reseller_free_time"`

	// source user
	// Required: true
	SourceUser *string `json:"source_user"`

	// source user id
	// Required: true
	SourceUserID *string `json:"source_user_id"`

	// source user out
	// Required: true
	SourceUserOut *string `json:"source_user_out"`

	// start time
	// Required: true
	StartTime *string `json:"start_time"`
}

// Validate validates this call
func (m *Call) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCallID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCallStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCallType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationCarrierBillingFeeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationCarrierBillingZoneID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationCarrierCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationCarrierFreeTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationCustomerBillingFeeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationCustomerBillingZoneID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationCustomerCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationCustomerFreeTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationDomainIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationExternalContractID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationExternalSubscriberID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationGpp0(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationGpp1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationGpp2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationGpp3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationGpp4(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationGpp5(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationGpp6(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationGpp7(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationGpp8(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationGpp9(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationLnpPrefix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationProviderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationResellerBillingFeeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationResellerBillingZoneID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationResellerCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationResellerFreeTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationUserDialed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationUserID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationUserIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationUserOut(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExportStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExportedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeerAuthRealm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeerAuthUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRatingStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceCarrierBillingFeeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceCarrierBillingZoneID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceCarrierCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceCarrierFreeTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceCli(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceClir(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceCustomerBillingFeeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceCustomerBillingZoneID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceCustomerCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceCustomerFreeTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceExternalContractID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceExternalSubscriberID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceGpp0(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceGpp1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceGpp2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceGpp3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceGpp4(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceGpp5(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceGpp6(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceGpp7(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceGpp8(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceGpp9(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceLnpPrefix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceProviderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceResellerBillingFeeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceResellerBillingZoneID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceResellerCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceResellerFreeTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceUserID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceUserOut(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Call) validateCallCode(formats strfmt.Registry) error {

	if err := validate.Required("call_code", "body", m.CallCode); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateCallID(formats strfmt.Registry) error {

	if err := validate.Required("call_id", "body", m.CallID); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateCallStatus(formats strfmt.Registry) error {

	if err := validate.Required("call_status", "body", m.CallStatus); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateCallType(formats strfmt.Registry) error {

	if err := validate.Required("call_type", "body", m.CallType); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationCarrierBillingFeeID(formats strfmt.Registry) error {

	if err := validate.Required("destination_carrier_billing_fee_id", "body", m.DestinationCarrierBillingFeeID); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationCarrierBillingZoneID(formats strfmt.Registry) error {

	if err := validate.Required("destination_carrier_billing_zone_id", "body", m.DestinationCarrierBillingZoneID); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationCarrierCost(formats strfmt.Registry) error {

	if err := validate.Required("destination_carrier_cost", "body", m.DestinationCarrierCost); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationCarrierFreeTime(formats strfmt.Registry) error {

	if err := validate.Required("destination_carrier_free_time", "body", m.DestinationCarrierFreeTime); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationCustomerBillingFeeID(formats strfmt.Registry) error {

	if err := validate.Required("destination_customer_billing_fee_id", "body", m.DestinationCustomerBillingFeeID); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationCustomerBillingZoneID(formats strfmt.Registry) error {

	if err := validate.Required("destination_customer_billing_zone_id", "body", m.DestinationCustomerBillingZoneID); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationCustomerCost(formats strfmt.Registry) error {

	if err := validate.Required("destination_customer_cost", "body", m.DestinationCustomerCost); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationCustomerFreeTime(formats strfmt.Registry) error {

	if err := validate.Required("destination_customer_free_time", "body", m.DestinationCustomerFreeTime); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationCustomerID(formats strfmt.Registry) error {

	if err := validate.Required("destination_customer_id", "body", m.DestinationCustomerID); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationDomain(formats strfmt.Registry) error {

	if err := validate.Required("destination_domain", "body", m.DestinationDomain); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationDomainIn(formats strfmt.Registry) error {

	if err := validate.Required("destination_domain_in", "body", m.DestinationDomainIn); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationExternalContractID(formats strfmt.Registry) error {

	if err := validate.Required("destination_external_contract_id", "body", m.DestinationExternalContractID); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationExternalSubscriberID(formats strfmt.Registry) error {

	if err := validate.Required("destination_external_subscriber_id", "body", m.DestinationExternalSubscriberID); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationGpp0(formats strfmt.Registry) error {

	if err := validate.Required("destination_gpp0", "body", m.DestinationGpp0); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationGpp1(formats strfmt.Registry) error {

	if err := validate.Required("destination_gpp1", "body", m.DestinationGpp1); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationGpp2(formats strfmt.Registry) error {

	if err := validate.Required("destination_gpp2", "body", m.DestinationGpp2); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationGpp3(formats strfmt.Registry) error {

	if err := validate.Required("destination_gpp3", "body", m.DestinationGpp3); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationGpp4(formats strfmt.Registry) error {

	if err := validate.Required("destination_gpp4", "body", m.DestinationGpp4); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationGpp5(formats strfmt.Registry) error {

	if err := validate.Required("destination_gpp5", "body", m.DestinationGpp5); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationGpp6(formats strfmt.Registry) error {

	if err := validate.Required("destination_gpp6", "body", m.DestinationGpp6); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationGpp7(formats strfmt.Registry) error {

	if err := validate.Required("destination_gpp7", "body", m.DestinationGpp7); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationGpp8(formats strfmt.Registry) error {

	if err := validate.Required("destination_gpp8", "body", m.DestinationGpp8); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationGpp9(formats strfmt.Registry) error {

	if err := validate.Required("destination_gpp9", "body", m.DestinationGpp9); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationLnpPrefix(formats strfmt.Registry) error {

	if err := validate.Required("destination_lnp_prefix", "body", m.DestinationLnpPrefix); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationProviderID(formats strfmt.Registry) error {

	if err := validate.Required("destination_provider_id", "body", m.DestinationProviderID); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationResellerBillingFeeID(formats strfmt.Registry) error {

	if err := validate.Required("destination_reseller_billing_fee_id", "body", m.DestinationResellerBillingFeeID); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationResellerBillingZoneID(formats strfmt.Registry) error {

	if err := validate.Required("destination_reseller_billing_zone_id", "body", m.DestinationResellerBillingZoneID); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationResellerCost(formats strfmt.Registry) error {

	if err := validate.Required("destination_reseller_cost", "body", m.DestinationResellerCost); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationResellerFreeTime(formats strfmt.Registry) error {

	if err := validate.Required("destination_reseller_free_time", "body", m.DestinationResellerFreeTime); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationUser(formats strfmt.Registry) error {

	if err := validate.Required("destination_user", "body", m.DestinationUser); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationUserDialed(formats strfmt.Registry) error {

	if err := validate.Required("destination_user_dialed", "body", m.DestinationUserDialed); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationUserID(formats strfmt.Registry) error {

	if err := validate.Required("destination_user_id", "body", m.DestinationUserID); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationUserIn(formats strfmt.Registry) error {

	if err := validate.Required("destination_user_in", "body", m.DestinationUserIn); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDestinationUserOut(formats strfmt.Registry) error {

	if err := validate.Required("destination_user_out", "body", m.DestinationUserOut); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateExportStatus(formats strfmt.Registry) error {

	if err := validate.Required("export_status", "body", m.ExportStatus); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateExportedAt(formats strfmt.Registry) error {

	if err := validate.Required("exported_at", "body", m.ExportedAt); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateInitTime(formats strfmt.Registry) error {

	if err := validate.Required("init_time", "body", m.InitTime); err != nil {
		return err
	}

	return nil
}

func (m *Call) validatePeerAuthRealm(formats strfmt.Registry) error {

	if err := validate.Required("peer_auth_realm", "body", m.PeerAuthRealm); err != nil {
		return err
	}

	return nil
}

func (m *Call) validatePeerAuthUser(formats strfmt.Registry) error {

	if err := validate.Required("peer_auth_user", "body", m.PeerAuthUser); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateRatedAt(formats strfmt.Registry) error {

	if err := validate.Required("rated_at", "body", m.RatedAt); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateRatingStatus(formats strfmt.Registry) error {

	if err := validate.Required("rating_status", "body", m.RatingStatus); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceCarrierBillingFeeID(formats strfmt.Registry) error {

	if err := validate.Required("source_carrier_billing_fee_id", "body", m.SourceCarrierBillingFeeID); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceCarrierBillingZoneID(formats strfmt.Registry) error {

	if err := validate.Required("source_carrier_billing_zone_id", "body", m.SourceCarrierBillingZoneID); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceCarrierCost(formats strfmt.Registry) error {

	if err := validate.Required("source_carrier_cost", "body", m.SourceCarrierCost); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceCarrierFreeTime(formats strfmt.Registry) error {

	if err := validate.Required("source_carrier_free_time", "body", m.SourceCarrierFreeTime); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceCli(formats strfmt.Registry) error {

	if err := validate.Required("source_cli", "body", m.SourceCli); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceClir(formats strfmt.Registry) error {

	if err := validate.Required("source_clir", "body", m.SourceClir); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceCustomerBillingFeeID(formats strfmt.Registry) error {

	if err := validate.Required("source_customer_billing_fee_id", "body", m.SourceCustomerBillingFeeID); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceCustomerBillingZoneID(formats strfmt.Registry) error {

	if err := validate.Required("source_customer_billing_zone_id", "body", m.SourceCustomerBillingZoneID); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceCustomerCost(formats strfmt.Registry) error {

	if err := validate.Required("source_customer_cost", "body", m.SourceCustomerCost); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceCustomerFreeTime(formats strfmt.Registry) error {

	if err := validate.Required("source_customer_free_time", "body", m.SourceCustomerFreeTime); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceCustomerID(formats strfmt.Registry) error {

	if err := validate.Required("source_customer_id", "body", m.SourceCustomerID); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceDomain(formats strfmt.Registry) error {

	if err := validate.Required("source_domain", "body", m.SourceDomain); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceExternalContractID(formats strfmt.Registry) error {

	if err := validate.Required("source_external_contract_id", "body", m.SourceExternalContractID); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceExternalSubscriberID(formats strfmt.Registry) error {

	if err := validate.Required("source_external_subscriber_id", "body", m.SourceExternalSubscriberID); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceGpp0(formats strfmt.Registry) error {

	if err := validate.Required("source_gpp0", "body", m.SourceGpp0); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceGpp1(formats strfmt.Registry) error {

	if err := validate.Required("source_gpp1", "body", m.SourceGpp1); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceGpp2(formats strfmt.Registry) error {

	if err := validate.Required("source_gpp2", "body", m.SourceGpp2); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceGpp3(formats strfmt.Registry) error {

	if err := validate.Required("source_gpp3", "body", m.SourceGpp3); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceGpp4(formats strfmt.Registry) error {

	if err := validate.Required("source_gpp4", "body", m.SourceGpp4); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceGpp5(formats strfmt.Registry) error {

	if err := validate.Required("source_gpp5", "body", m.SourceGpp5); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceGpp6(formats strfmt.Registry) error {

	if err := validate.Required("source_gpp6", "body", m.SourceGpp6); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceGpp7(formats strfmt.Registry) error {

	if err := validate.Required("source_gpp7", "body", m.SourceGpp7); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceGpp8(formats strfmt.Registry) error {

	if err := validate.Required("source_gpp8", "body", m.SourceGpp8); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceGpp9(formats strfmt.Registry) error {

	if err := validate.Required("source_gpp9", "body", m.SourceGpp9); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceIP(formats strfmt.Registry) error {

	if err := validate.Required("source_ip", "body", m.SourceIP); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceLnpPrefix(formats strfmt.Registry) error {

	if err := validate.Required("source_lnp_prefix", "body", m.SourceLnpPrefix); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceProviderID(formats strfmt.Registry) error {

	if err := validate.Required("source_provider_id", "body", m.SourceProviderID); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceResellerBillingFeeID(formats strfmt.Registry) error {

	if err := validate.Required("source_reseller_billing_fee_id", "body", m.SourceResellerBillingFeeID); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceResellerBillingZoneID(formats strfmt.Registry) error {

	if err := validate.Required("source_reseller_billing_zone_id", "body", m.SourceResellerBillingZoneID); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceResellerCost(formats strfmt.Registry) error {

	if err := validate.Required("source_reseller_cost", "body", m.SourceResellerCost); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceResellerFreeTime(formats strfmt.Registry) error {

	if err := validate.Required("source_reseller_free_time", "body", m.SourceResellerFreeTime); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceUser(formats strfmt.Registry) error {

	if err := validate.Required("source_user", "body", m.SourceUser); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceUserID(formats strfmt.Registry) error {

	if err := validate.Required("source_user_id", "body", m.SourceUserID); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateSourceUserOut(formats strfmt.Registry) error {

	if err := validate.Required("source_user_out", "body", m.SourceUserOut); err != nil {
		return err
	}

	return nil
}

func (m *Call) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("start_time", "body", m.StartTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Call) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Call) UnmarshalBinary(b []byte) error {
	var res Call
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}