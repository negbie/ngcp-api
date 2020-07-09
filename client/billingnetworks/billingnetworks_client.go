// Code generated by go-swagger; DO NOT EDIT.

package billingnetworks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new billingnetworks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for billingnetworks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateanewBillingNetwork(params *CreateanewBillingNetworkParams) (*CreateanewBillingNetworkCreated, error)

	GetBillingNetworkitems(params *GetBillingNetworkitemsParams) (*GetBillingNetworkitemsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateanewBillingNetwork creates a new billing network
*/
func (a *Client) CreateanewBillingNetwork(params *CreateanewBillingNetworkParams) (*CreateanewBillingNetworkCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateanewBillingNetworkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateanewBillingNetwork",
		Method:             "POST",
		PathPattern:        "/billingnetworks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateanewBillingNetworkReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateanewBillingNetworkCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateanewBillingNetwork: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetBillingNetworkitems gets billing network items
*/
func (a *Client) GetBillingNetworkitems(params *GetBillingNetworkitemsParams) (*GetBillingNetworkitemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBillingNetworkitemsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetBillingNetworkitems",
		Method:             "GET",
		PathPattern:        "/billingnetworks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBillingNetworkitemsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBillingNetworkitemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetBillingNetworkitems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
