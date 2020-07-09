// Code generated by go-swagger; DO NOT EDIT.

package balanceintervals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new balanceintervals API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for balanceintervals API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetBalanceIntervalitems(params *GetBalanceIntervalitemsParams) (*GetBalanceIntervalitemsOK, error)

	GetaspecificBalanceInterval(params *GetaspecificBalanceIntervalParams) (*GetaspecificBalanceIntervalOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetBalanceIntervalitems gets balance interval items
*/
func (a *Client) GetBalanceIntervalitems(params *GetBalanceIntervalitemsParams) (*GetBalanceIntervalitemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBalanceIntervalitemsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetBalanceIntervalitems",
		Method:             "GET",
		PathPattern:        "/balanceintervals",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBalanceIntervalitemsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBalanceIntervalitemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetBalanceIntervalitems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetaspecificBalanceInterval gets a specific balance interval
*/
func (a *Client) GetaspecificBalanceInterval(params *GetaspecificBalanceIntervalParams) (*GetaspecificBalanceIntervalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetaspecificBalanceIntervalParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetaspecificBalanceInterval",
		Method:             "GET",
		PathPattern:        "/balanceintervals/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetaspecificBalanceIntervalReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetaspecificBalanceIntervalOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetaspecificBalanceInterval: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
