// Code generated by go-swagger; DO NOT EDIT.

package ncoslevels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ncoslevels API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ncoslevels API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateanewNcosLevel(params *CreateanewNcosLevelParams) (*CreateanewNcosLevelCreated, error)

	GetNcosLevelitems(params *GetNcosLevelitemsParams) (*GetNcosLevelitemsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateanewNcosLevel creates a new ncos level
*/
func (a *Client) CreateanewNcosLevel(params *CreateanewNcosLevelParams) (*CreateanewNcosLevelCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateanewNcosLevelParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateanewNcosLevel",
		Method:             "POST",
		PathPattern:        "/ncoslevels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateanewNcosLevelReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateanewNcosLevelCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateanewNcosLevel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNcosLevelitems gets ncos level items
*/
func (a *Client) GetNcosLevelitems(params *GetNcosLevelitemsParams) (*GetNcosLevelitemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNcosLevelitemsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNcosLevelitems",
		Method:             "GET",
		PathPattern:        "/ncoslevels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNcosLevelitemsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNcosLevelitemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetNcosLevelitems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
