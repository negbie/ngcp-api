// Code generated by go-swagger; DO NOT EDIT.

package pbxdeviceprofiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new pbxdeviceprofiles API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pbxdeviceprofiles API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateanewPbxDeviceProfile(params *CreateanewPbxDeviceProfileParams) (*CreateanewPbxDeviceProfileCreated, error)

	GetPbxDeviceProfileitems(params *GetPbxDeviceProfileitemsParams) (*GetPbxDeviceProfileitemsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateanewPbxDeviceProfile creates a new pbx device profile
*/
func (a *Client) CreateanewPbxDeviceProfile(params *CreateanewPbxDeviceProfileParams) (*CreateanewPbxDeviceProfileCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateanewPbxDeviceProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateanewPbxDeviceProfile",
		Method:             "POST",
		PathPattern:        "/pbxdeviceprofiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateanewPbxDeviceProfileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateanewPbxDeviceProfileCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateanewPbxDeviceProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPbxDeviceProfileitems gets pbx device profile items
*/
func (a *Client) GetPbxDeviceProfileitems(params *GetPbxDeviceProfileitemsParams) (*GetPbxDeviceProfileitemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPbxDeviceProfileitemsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPbxDeviceProfileitems",
		Method:             "GET",
		PathPattern:        "/pbxdeviceprofiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPbxDeviceProfileitemsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPbxDeviceProfileitemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPbxDeviceProfileitems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
