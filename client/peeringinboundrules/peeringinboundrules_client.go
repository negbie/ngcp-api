// Code generated by go-swagger; DO NOT EDIT.

package peeringinboundrules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new peeringinboundrules API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for peeringinboundrules API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateanewPeeringInboundRule(params *CreateanewPeeringInboundRuleParams) (*CreateanewPeeringInboundRuleCreated, error)

	GetPeeringInboundRuleitems(params *GetPeeringInboundRuleitemsParams) (*GetPeeringInboundRuleitemsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateanewPeeringInboundRule creates a new peering inbound rule
*/
func (a *Client) CreateanewPeeringInboundRule(params *CreateanewPeeringInboundRuleParams) (*CreateanewPeeringInboundRuleCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateanewPeeringInboundRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateanewPeeringInboundRule",
		Method:             "POST",
		PathPattern:        "/peeringinboundrules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateanewPeeringInboundRuleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateanewPeeringInboundRuleCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateanewPeeringInboundRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPeeringInboundRuleitems gets peering inbound rule items
*/
func (a *Client) GetPeeringInboundRuleitems(params *GetPeeringInboundRuleitemsParams) (*GetPeeringInboundRuleitemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPeeringInboundRuleitemsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPeeringInboundRuleitems",
		Method:             "GET",
		PathPattern:        "/peeringinboundrules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPeeringInboundRuleitemsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPeeringInboundRuleitemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPeeringInboundRuleitems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
