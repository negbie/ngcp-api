// Code generated by go-swagger; DO NOT EDIT.

package topuplogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new topuplogs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for topuplogs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetTopupLogitems(params *GetTopupLogitemsParams) (*GetTopupLogitemsOK, error)

	GetaspecificTopupLog(params *GetaspecificTopupLogParams) (*GetaspecificTopupLogOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetTopupLogitems gets topup log items
*/
func (a *Client) GetTopupLogitems(params *GetTopupLogitemsParams) (*GetTopupLogitemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTopupLogitemsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTopupLogitems",
		Method:             "GET",
		PathPattern:        "/topuplogs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTopupLogitemsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTopupLogitemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetTopupLogitems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetaspecificTopupLog gets a specific topup log
*/
func (a *Client) GetaspecificTopupLog(params *GetaspecificTopupLogParams) (*GetaspecificTopupLogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetaspecificTopupLogParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetaspecificTopupLog",
		Method:             "GET",
		PathPattern:        "/topuplogs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetaspecificTopupLogReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetaspecificTopupLogOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetaspecificTopupLog: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
