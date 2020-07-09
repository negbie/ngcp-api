// Code generated by go-swagger; DO NOT EDIT.

package phonebookentries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// CreateanewPhonebookEntryReader is a Reader for the CreateanewPhonebookEntry structure.
type CreateanewPhonebookEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateanewPhonebookEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateanewPhonebookEntryCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateanewPhonebookEntryUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateanewPhonebookEntryCreated creates a CreateanewPhonebookEntryCreated with default headers values
func NewCreateanewPhonebookEntryCreated() *CreateanewPhonebookEntryCreated {
	return &CreateanewPhonebookEntryCreated{
		Location: "nostrud velit non",
	}
}

/*CreateanewPhonebookEntryCreated handles this case with default header values.

CreateanewPhonebookEntryCreated createanew phonebook entry created
*/
type CreateanewPhonebookEntryCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty14
}

func (o *CreateanewPhonebookEntryCreated) Error() string {
	return fmt.Sprintf("[POST /phonebookentries][%d] createanewPhonebookEntryCreated  %+v", 201, o.Payload)
}

func (o *CreateanewPhonebookEntryCreated) GetPayload() []*models.Thenewlycreateditemorempty14 {
	return o.Payload
}

func (o *CreateanewPhonebookEntryCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateanewPhonebookEntryUnprocessableEntity creates a CreateanewPhonebookEntryUnprocessableEntity with default headers values
func NewCreateanewPhonebookEntryUnprocessableEntity() *CreateanewPhonebookEntryUnprocessableEntity {
	return &CreateanewPhonebookEntryUnprocessableEntity{}
}

/*CreateanewPhonebookEntryUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type CreateanewPhonebookEntryUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *CreateanewPhonebookEntryUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /phonebookentries][%d] createanewPhonebookEntryUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateanewPhonebookEntryUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *CreateanewPhonebookEntryUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}