// Code generated by go-swagger; DO NOT EDIT.

package peeringinboundrules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// CreateanewPeeringInboundRuleReader is a Reader for the CreateanewPeeringInboundRule structure.
type CreateanewPeeringInboundRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateanewPeeringInboundRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateanewPeeringInboundRuleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateanewPeeringInboundRuleUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateanewPeeringInboundRuleCreated creates a CreateanewPeeringInboundRuleCreated with default headers values
func NewCreateanewPeeringInboundRuleCreated() *CreateanewPeeringInboundRuleCreated {
	return &CreateanewPeeringInboundRuleCreated{
		Location: "nostrud velit non",
	}
}

/*CreateanewPeeringInboundRuleCreated handles this case with default header values.

CreateanewPeeringInboundRuleCreated createanew peering inbound rule created
*/
type CreateanewPeeringInboundRuleCreated struct {
	Location string

	Payload string
}

func (o *CreateanewPeeringInboundRuleCreated) Error() string {
	return fmt.Sprintf("[POST /peeringinboundrules][%d] createanewPeeringInboundRuleCreated  %+v", 201, o.Payload)
}

func (o *CreateanewPeeringInboundRuleCreated) GetPayload() string {
	return o.Payload
}

func (o *CreateanewPeeringInboundRuleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateanewPeeringInboundRuleUnprocessableEntity creates a CreateanewPeeringInboundRuleUnprocessableEntity with default headers values
func NewCreateanewPeeringInboundRuleUnprocessableEntity() *CreateanewPeeringInboundRuleUnprocessableEntity {
	return &CreateanewPeeringInboundRuleUnprocessableEntity{}
}

/*CreateanewPeeringInboundRuleUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type CreateanewPeeringInboundRuleUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *CreateanewPeeringInboundRuleUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /peeringinboundrules][%d] createanewPeeringInboundRuleUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateanewPeeringInboundRuleUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *CreateanewPeeringInboundRuleUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
