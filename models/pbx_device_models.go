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

// PbxDeviceModels PbxDeviceModels
//
// swagger:model PbxDeviceModels
type PbxDeviceModels struct {

	// bootstrap config http sync method
	// Required: true
	BootstrapConfigHTTPSyncMethod *string `json:"bootstrap_config_http_sync_method"`

	// bootstrap config http sync params
	// Required: true
	BootstrapConfigHTTPSyncParams *string `json:"bootstrap_config_http_sync_params"`

	// bootstrap config http sync uri
	// Required: true
	BootstrapConfigHTTPSyncURI *string `json:"bootstrap_config_http_sync_uri"`

	// bootstrap config redirect grandstream cid
	// Required: true
	BootstrapConfigRedirectGrandstreamCid *string `json:"bootstrap_config_redirect_grandstream_cid"`

	// bootstrap config redirect grandstream key
	// Required: true
	BootstrapConfigRedirectGrandstreamKey *string `json:"bootstrap_config_redirect_grandstream_key"`

	// bootstrap config redirect panasonic password
	// Required: true
	BootstrapConfigRedirectPanasonicPassword *string `json:"bootstrap_config_redirect_panasonic_password"`

	// bootstrap config redirect panasonic user
	// Required: true
	BootstrapConfigRedirectPanasonicUser *string `json:"bootstrap_config_redirect_panasonic_user"`

	// bootstrap config redirect polycom password
	// Required: true
	BootstrapConfigRedirectPolycomPassword *string `json:"bootstrap_config_redirect_polycom_password"`

	// bootstrap config redirect polycom profile
	// Required: true
	BootstrapConfigRedirectPolycomProfile *string `json:"bootstrap_config_redirect_polycom_profile"`

	// bootstrap config redirect polycom user
	// Required: true
	BootstrapConfigRedirectPolycomUser *string `json:"bootstrap_config_redirect_polycom_user"`

	// bootstrap config redirect snom password
	// Required: true
	BootstrapConfigRedirectSnomPassword *string `json:"bootstrap_config_redirect_snom_password"`

	// bootstrap config redirect snom user
	// Required: true
	BootstrapConfigRedirectSnomUser *string `json:"bootstrap_config_redirect_snom_user"`

	// bootstrap config redirect yealink password
	// Required: true
	BootstrapConfigRedirectYealinkPassword *string `json:"bootstrap_config_redirect_yealink_password"`

	// bootstrap config redirect yealink user
	// Required: true
	BootstrapConfigRedirectYealinkUser *string `json:"bootstrap_config_redirect_yealink_user"`

	// bootstrap method
	// Required: true
	BootstrapMethod *string `json:"bootstrap_method"`

	// bootstrap uri
	// Required: true
	BootstrapURI *string `json:"bootstrap_uri"`

	// connectable models id
	// Required: true
	ConnectableModelsID *string `json:"connectable_models_id"`

	// extensions num
	// Required: true
	ExtensionsNum *string `json:"extensions_num"`

	// linerange
	// Required: true
	Linerange []*Linerange `json:"linerange"`

	// model
	// Required: true
	Model *string `json:"model"`

	// reseller id
	// Required: true
	ResellerID *string `json:"reseller_id"`

	// type
	// Required: true
	Type *string `json:"type"`

	// vendor
	// Required: true
	Vendor *string `json:"vendor"`
}

// Validate validates this pbx device models
func (m *PbxDeviceModels) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBootstrapConfigHTTPSyncMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBootstrapConfigHTTPSyncParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBootstrapConfigHTTPSyncURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBootstrapConfigRedirectGrandstreamCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBootstrapConfigRedirectGrandstreamKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBootstrapConfigRedirectPanasonicPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBootstrapConfigRedirectPanasonicUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBootstrapConfigRedirectPolycomPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBootstrapConfigRedirectPolycomProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBootstrapConfigRedirectPolycomUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBootstrapConfigRedirectSnomPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBootstrapConfigRedirectSnomUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBootstrapConfigRedirectYealinkPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBootstrapConfigRedirectYealinkUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBootstrapMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBootstrapURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectableModelsID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtensionsNum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinerange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResellerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVendor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PbxDeviceModels) validateBootstrapConfigHTTPSyncMethod(formats strfmt.Registry) error {

	if err := validate.Required("bootstrap_config_http_sync_method", "body", m.BootstrapConfigHTTPSyncMethod); err != nil {
		return err
	}

	return nil
}

func (m *PbxDeviceModels) validateBootstrapConfigHTTPSyncParams(formats strfmt.Registry) error {

	if err := validate.Required("bootstrap_config_http_sync_params", "body", m.BootstrapConfigHTTPSyncParams); err != nil {
		return err
	}

	return nil
}

func (m *PbxDeviceModels) validateBootstrapConfigHTTPSyncURI(formats strfmt.Registry) error {

	if err := validate.Required("bootstrap_config_http_sync_uri", "body", m.BootstrapConfigHTTPSyncURI); err != nil {
		return err
	}

	return nil
}

func (m *PbxDeviceModels) validateBootstrapConfigRedirectGrandstreamCid(formats strfmt.Registry) error {

	if err := validate.Required("bootstrap_config_redirect_grandstream_cid", "body", m.BootstrapConfigRedirectGrandstreamCid); err != nil {
		return err
	}

	return nil
}

func (m *PbxDeviceModels) validateBootstrapConfigRedirectGrandstreamKey(formats strfmt.Registry) error {

	if err := validate.Required("bootstrap_config_redirect_grandstream_key", "body", m.BootstrapConfigRedirectGrandstreamKey); err != nil {
		return err
	}

	return nil
}

func (m *PbxDeviceModels) validateBootstrapConfigRedirectPanasonicPassword(formats strfmt.Registry) error {

	if err := validate.Required("bootstrap_config_redirect_panasonic_password", "body", m.BootstrapConfigRedirectPanasonicPassword); err != nil {
		return err
	}

	return nil
}

func (m *PbxDeviceModels) validateBootstrapConfigRedirectPanasonicUser(formats strfmt.Registry) error {

	if err := validate.Required("bootstrap_config_redirect_panasonic_user", "body", m.BootstrapConfigRedirectPanasonicUser); err != nil {
		return err
	}

	return nil
}

func (m *PbxDeviceModels) validateBootstrapConfigRedirectPolycomPassword(formats strfmt.Registry) error {

	if err := validate.Required("bootstrap_config_redirect_polycom_password", "body", m.BootstrapConfigRedirectPolycomPassword); err != nil {
		return err
	}

	return nil
}

func (m *PbxDeviceModels) validateBootstrapConfigRedirectPolycomProfile(formats strfmt.Registry) error {

	if err := validate.Required("bootstrap_config_redirect_polycom_profile", "body", m.BootstrapConfigRedirectPolycomProfile); err != nil {
		return err
	}

	return nil
}

func (m *PbxDeviceModels) validateBootstrapConfigRedirectPolycomUser(formats strfmt.Registry) error {

	if err := validate.Required("bootstrap_config_redirect_polycom_user", "body", m.BootstrapConfigRedirectPolycomUser); err != nil {
		return err
	}

	return nil
}

func (m *PbxDeviceModels) validateBootstrapConfigRedirectSnomPassword(formats strfmt.Registry) error {

	if err := validate.Required("bootstrap_config_redirect_snom_password", "body", m.BootstrapConfigRedirectSnomPassword); err != nil {
		return err
	}

	return nil
}

func (m *PbxDeviceModels) validateBootstrapConfigRedirectSnomUser(formats strfmt.Registry) error {

	if err := validate.Required("bootstrap_config_redirect_snom_user", "body", m.BootstrapConfigRedirectSnomUser); err != nil {
		return err
	}

	return nil
}

func (m *PbxDeviceModels) validateBootstrapConfigRedirectYealinkPassword(formats strfmt.Registry) error {

	if err := validate.Required("bootstrap_config_redirect_yealink_password", "body", m.BootstrapConfigRedirectYealinkPassword); err != nil {
		return err
	}

	return nil
}

func (m *PbxDeviceModels) validateBootstrapConfigRedirectYealinkUser(formats strfmt.Registry) error {

	if err := validate.Required("bootstrap_config_redirect_yealink_user", "body", m.BootstrapConfigRedirectYealinkUser); err != nil {
		return err
	}

	return nil
}

func (m *PbxDeviceModels) validateBootstrapMethod(formats strfmt.Registry) error {

	if err := validate.Required("bootstrap_method", "body", m.BootstrapMethod); err != nil {
		return err
	}

	return nil
}

func (m *PbxDeviceModels) validateBootstrapURI(formats strfmt.Registry) error {

	if err := validate.Required("bootstrap_uri", "body", m.BootstrapURI); err != nil {
		return err
	}

	return nil
}

func (m *PbxDeviceModels) validateConnectableModelsID(formats strfmt.Registry) error {

	if err := validate.Required("connectable_models_id", "body", m.ConnectableModelsID); err != nil {
		return err
	}

	return nil
}

func (m *PbxDeviceModels) validateExtensionsNum(formats strfmt.Registry) error {

	if err := validate.Required("extensions_num", "body", m.ExtensionsNum); err != nil {
		return err
	}

	return nil
}

func (m *PbxDeviceModels) validateLinerange(formats strfmt.Registry) error {

	if err := validate.Required("linerange", "body", m.Linerange); err != nil {
		return err
	}

	for i := 0; i < len(m.Linerange); i++ {
		if swag.IsZero(m.Linerange[i]) { // not required
			continue
		}

		if m.Linerange[i] != nil {
			if err := m.Linerange[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("linerange" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PbxDeviceModels) validateModel(formats strfmt.Registry) error {

	if err := validate.Required("model", "body", m.Model); err != nil {
		return err
	}

	return nil
}

func (m *PbxDeviceModels) validateResellerID(formats strfmt.Registry) error {

	if err := validate.Required("reseller_id", "body", m.ResellerID); err != nil {
		return err
	}

	return nil
}

func (m *PbxDeviceModels) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *PbxDeviceModels) validateVendor(formats strfmt.Registry) error {

	if err := validate.Required("vendor", "body", m.Vendor); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PbxDeviceModels) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PbxDeviceModels) UnmarshalBinary(b []byte) error {
	var res PbxDeviceModels
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
