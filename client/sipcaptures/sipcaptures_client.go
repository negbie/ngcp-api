// Code generated by go-swagger; DO NOT EDIT.

package sipcaptures

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new sipcaptures API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sipcaptures API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetSIPCaptureitems(params *GetSIPCaptureitemsParams) (*GetSIPCaptureitemsOK, error)

	GetaspecificSIPCapture(params *GetaspecificSIPCaptureParams) (*GetaspecificSIPCaptureOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetSIPCaptureitems gets s IP capture items
*/
func (a *Client) GetSIPCaptureitems(params *GetSIPCaptureitemsParams) (*GetSIPCaptureitemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSIPCaptureitemsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSIPCaptureitems",
		Method:             "GET",
		PathPattern:        "/sipcaptures",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSIPCaptureitemsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSIPCaptureitemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSIPCaptureitems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetaspecificSIPCapture gets a specific s IP capture
*/
func (a *Client) GetaspecificSIPCapture(params *GetaspecificSIPCaptureParams) (*GetaspecificSIPCaptureOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetaspecificSIPCaptureParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetaspecificSIPCapture",
		Method:             "GET",
		PathPattern:        "/sipcaptures/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetaspecificSIPCaptureReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetaspecificSIPCaptureOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetaspecificSIPCapture: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
