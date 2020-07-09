// Code generated by go-swagger; DO NOT EDIT.

package cfdestinationsets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// CreateanewCFDestinationSetReader is a Reader for the CreateanewCFDestinationSet structure.
type CreateanewCFDestinationSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateanewCFDestinationSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateanewCFDestinationSetCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateanewCFDestinationSetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateanewCFDestinationSetCreated creates a CreateanewCFDestinationSetCreated with default headers values
func NewCreateanewCFDestinationSetCreated() *CreateanewCFDestinationSetCreated {
	return &CreateanewCFDestinationSetCreated{
		Location: "nostrud velit non",
	}
}

/*CreateanewCFDestinationSetCreated handles this case with default header values.

CreateanewCFDestinationSetCreated createanew c f destination set created
*/
type CreateanewCFDestinationSetCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty28
}

func (o *CreateanewCFDestinationSetCreated) Error() string {
	return fmt.Sprintf("[POST /cfdestinationsets][%d] createanewCFDestinationSetCreated  %+v", 201, o.Payload)
}

func (o *CreateanewCFDestinationSetCreated) GetPayload() []*models.Thenewlycreateditemorempty28 {
	return o.Payload
}

func (o *CreateanewCFDestinationSetCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateanewCFDestinationSetUnprocessableEntity creates a CreateanewCFDestinationSetUnprocessableEntity with default headers values
func NewCreateanewCFDestinationSetUnprocessableEntity() *CreateanewCFDestinationSetUnprocessableEntity {
	return &CreateanewCFDestinationSetUnprocessableEntity{}
}

/*CreateanewCFDestinationSetUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type CreateanewCFDestinationSetUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *CreateanewCFDestinationSetUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /cfdestinationsets][%d] createanewCFDestinationSetUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateanewCFDestinationSetUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *CreateanewCFDestinationSetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}