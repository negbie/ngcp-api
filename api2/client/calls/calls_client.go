// Code generated by go-swagger; DO NOT EDIT.

package calls

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new calls API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for calls API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetCallitems(params *GetCallitemsParams) (*GetCallitemsOK, error)

	GetaspecificCall(params *GetaspecificCallParams) (*GetaspecificCallOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetCallitems gets call items
*/
func (a *Client) GetCallitems(params *GetCallitemsParams) (*GetCallitemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCallitemsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCallitems",
		Method:             "GET",
		PathPattern:        "/calls",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCallitemsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCallitemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCallitems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetaspecificCall gets a specific call
*/
func (a *Client) GetaspecificCall(params *GetaspecificCallParams) (*GetaspecificCallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetaspecificCallParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetaspecificCall",
		Method:             "GET",
		PathPattern:        "/calls/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetaspecificCallReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetaspecificCallOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetaspecificCall: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}