// Code generated by go-swagger; DO NOT EDIT.

package rewriterules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new rewriterules API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for rewriterules API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateanewRewriteRule(params *CreateanewRewriteRuleParams) (*CreateanewRewriteRuleCreated, error)

	GetRewriteRuleitems(params *GetRewriteRuleitemsParams) (*GetRewriteRuleitemsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateanewRewriteRule creates a new rewrite rule
*/
func (a *Client) CreateanewRewriteRule(params *CreateanewRewriteRuleParams) (*CreateanewRewriteRuleCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateanewRewriteRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateanewRewriteRule",
		Method:             "POST",
		PathPattern:        "/rewriterules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateanewRewriteRuleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateanewRewriteRuleCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateanewRewriteRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRewriteRuleitems gets rewrite rule items
*/
func (a *Client) GetRewriteRuleitems(params *GetRewriteRuleitemsParams) (*GetRewriteRuleitemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRewriteRuleitemsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRewriteRuleitems",
		Method:             "GET",
		PathPattern:        "/rewriterules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRewriteRuleitemsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRewriteRuleitemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetRewriteRuleitems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
