// Code generated by go-swagger; DO NOT EDIT.

package autoattendants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new autoattendants API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for autoattendants API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetAutoAttendantitems(params *GetAutoAttendantitemsParams) (*GetAutoAttendantitemsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAutoAttendantitems gets auto attendant items
*/
func (a *Client) GetAutoAttendantitems(params *GetAutoAttendantitemsParams) (*GetAutoAttendantitemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAutoAttendantitemsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAutoAttendantitems",
		Method:             "GET",
		PathPattern:        "/autoattendants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAutoAttendantitemsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAutoAttendantitemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAutoAttendantitems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
