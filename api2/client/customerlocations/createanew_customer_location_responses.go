// Code generated by go-swagger; DO NOT EDIT.

package customerlocations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// CreateanewCustomerLocationReader is a Reader for the CreateanewCustomerLocation structure.
type CreateanewCustomerLocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateanewCustomerLocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateanewCustomerLocationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateanewCustomerLocationUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateanewCustomerLocationCreated creates a CreateanewCustomerLocationCreated with default headers values
func NewCreateanewCustomerLocationCreated() *CreateanewCustomerLocationCreated {
	return &CreateanewCustomerLocationCreated{
		Location: "nostrud velit non",
	}
}

/*CreateanewCustomerLocationCreated handles this case with default header values.

CreateanewCustomerLocationCreated createanew customer location created
*/
type CreateanewCustomerLocationCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty27
}

func (o *CreateanewCustomerLocationCreated) Error() string {
	return fmt.Sprintf("[POST /customerlocations][%d] createanewCustomerLocationCreated  %+v", 201, o.Payload)
}

func (o *CreateanewCustomerLocationCreated) GetPayload() []*models.Thenewlycreateditemorempty27 {
	return o.Payload
}

func (o *CreateanewCustomerLocationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateanewCustomerLocationUnprocessableEntity creates a CreateanewCustomerLocationUnprocessableEntity with default headers values
func NewCreateanewCustomerLocationUnprocessableEntity() *CreateanewCustomerLocationUnprocessableEntity {
	return &CreateanewCustomerLocationUnprocessableEntity{}
}

/*CreateanewCustomerLocationUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type CreateanewCustomerLocationUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *CreateanewCustomerLocationUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /customerlocations][%d] createanewCustomerLocationUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateanewCustomerLocationUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *CreateanewCustomerLocationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
