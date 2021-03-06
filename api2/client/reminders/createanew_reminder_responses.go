// Code generated by go-swagger; DO NOT EDIT.

package reminders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// CreateanewReminderReader is a Reader for the CreateanewReminder structure.
type CreateanewReminderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateanewReminderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateanewReminderCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateanewReminderUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateanewReminderCreated creates a CreateanewReminderCreated with default headers values
func NewCreateanewReminderCreated() *CreateanewReminderCreated {
	return &CreateanewReminderCreated{
		Location: "nostrud velit non",
	}
}

/*CreateanewReminderCreated handles this case with default header values.

CreateanewReminderCreated createanew reminder created
*/
type CreateanewReminderCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty2
}

func (o *CreateanewReminderCreated) Error() string {
	return fmt.Sprintf("[POST /reminders][%d] createanewReminderCreated  %+v", 201, o.Payload)
}

func (o *CreateanewReminderCreated) GetPayload() []*models.Thenewlycreateditemorempty2 {
	return o.Payload
}

func (o *CreateanewReminderCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateanewReminderUnprocessableEntity creates a CreateanewReminderUnprocessableEntity with default headers values
func NewCreateanewReminderUnprocessableEntity() *CreateanewReminderUnprocessableEntity {
	return &CreateanewReminderUnprocessableEntity{}
}

/*CreateanewReminderUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type CreateanewReminderUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *CreateanewReminderUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /reminders][%d] createanewReminderUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateanewReminderUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *CreateanewReminderUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
