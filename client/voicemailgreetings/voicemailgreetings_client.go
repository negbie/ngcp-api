// Code generated by go-swagger; DO NOT EDIT.

package voicemailgreetings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new voicemailgreetings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for voicemailgreetings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateanewVoicemailGreeting(params *CreateanewVoicemailGreetingParams) (*CreateanewVoicemailGreetingCreated, error)

	GetVoicemailGreetingitems(params *GetVoicemailGreetingitemsParams) (*GetVoicemailGreetingitemsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateanewVoicemailGreeting creates a new voicemail greeting
*/
func (a *Client) CreateanewVoicemailGreeting(params *CreateanewVoicemailGreetingParams) (*CreateanewVoicemailGreetingCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateanewVoicemailGreetingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateanewVoicemailGreeting",
		Method:             "POST",
		PathPattern:        "/voicemailgreetings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateanewVoicemailGreetingReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateanewVoicemailGreetingCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateanewVoicemailGreeting: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVoicemailGreetingitems gets voicemail greeting items
*/
func (a *Client) GetVoicemailGreetingitems(params *GetVoicemailGreetingitemsParams) (*GetVoicemailGreetingitemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVoicemailGreetingitemsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetVoicemailGreetingitems",
		Method:             "GET",
		PathPattern:        "/voicemailgreetings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVoicemailGreetingitemsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVoicemailGreetingitemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetVoicemailGreetingitems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}