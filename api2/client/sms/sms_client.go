// Code generated by go-swagger; DO NOT EDIT.

package sms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new sms API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sms API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateanewSMS(params *CreateanewSMSParams) (*CreateanewSMSCreated, error)

	GetSMSitems(params *GetSMSitemsParams) (*GetSMSitemsOK, error)

	GetaspecificSMS(params *GetaspecificSMSParams) (*GetaspecificSMSOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateanewSMS creates a new s m s
*/
func (a *Client) CreateanewSMS(params *CreateanewSMSParams) (*CreateanewSMSCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateanewSMSParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateanewSMS",
		Method:             "POST",
		PathPattern:        "/sms",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateanewSMSReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateanewSMSCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateanewSMS: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSMSitems gets s m s items
*/
func (a *Client) GetSMSitems(params *GetSMSitemsParams) (*GetSMSitemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSMSitemsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSMSitems",
		Method:             "GET",
		PathPattern:        "/sms",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSMSitemsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSMSitemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSMSitems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetaspecificSMS gets a specific s m s
*/
func (a *Client) GetaspecificSMS(params *GetaspecificSMSParams) (*GetaspecificSMSOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetaspecificSMSParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetaspecificSMS",
		Method:             "GET",
		PathPattern:        "/sms/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetaspecificSMSReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetaspecificSMSOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetaspecificSMS: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}