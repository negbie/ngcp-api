// Code generated by go-swagger; DO NOT EDIT.

package systemcontacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new systemcontacts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for systemcontacts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateanewSystemContact(params *CreateanewSystemContactParams) (*CreateanewSystemContactCreated, error)

	GetSystemContactitems(params *GetSystemContactitemsParams) (*GetSystemContactitemsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateanewSystemContact creates a new system contact
*/
func (a *Client) CreateanewSystemContact(params *CreateanewSystemContactParams) (*CreateanewSystemContactCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateanewSystemContactParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateanewSystemContact",
		Method:             "POST",
		PathPattern:        "/systemcontacts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateanewSystemContactReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateanewSystemContactCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateanewSystemContact: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSystemContactitems gets system contact items
*/
func (a *Client) GetSystemContactitems(params *GetSystemContactitemsParams) (*GetSystemContactitemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSystemContactitemsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSystemContactitems",
		Method:             "GET",
		PathPattern:        "/systemcontacts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSystemContactitemsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSystemContactitemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSystemContactitems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
