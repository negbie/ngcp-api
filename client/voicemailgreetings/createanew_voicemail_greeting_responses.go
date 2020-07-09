// Code generated by go-swagger; DO NOT EDIT.

package voicemailgreetings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// CreateanewVoicemailGreetingReader is a Reader for the CreateanewVoicemailGreeting structure.
type CreateanewVoicemailGreetingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateanewVoicemailGreetingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateanewVoicemailGreetingCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateanewVoicemailGreetingUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateanewVoicemailGreetingCreated creates a CreateanewVoicemailGreetingCreated with default headers values
func NewCreateanewVoicemailGreetingCreated() *CreateanewVoicemailGreetingCreated {
	return &CreateanewVoicemailGreetingCreated{
		Location: "nostrud velit non",
	}
}

/*CreateanewVoicemailGreetingCreated handles this case with default header values.

CreateanewVoicemailGreetingCreated createanew voicemail greeting created
*/
type CreateanewVoicemailGreetingCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty7
}

func (o *CreateanewVoicemailGreetingCreated) Error() string {
	return fmt.Sprintf("[POST /voicemailgreetings][%d] createanewVoicemailGreetingCreated  %+v", 201, o.Payload)
}

func (o *CreateanewVoicemailGreetingCreated) GetPayload() []*models.Thenewlycreateditemorempty7 {
	return o.Payload
}

func (o *CreateanewVoicemailGreetingCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateanewVoicemailGreetingUnprocessableEntity creates a CreateanewVoicemailGreetingUnprocessableEntity with default headers values
func NewCreateanewVoicemailGreetingUnprocessableEntity() *CreateanewVoicemailGreetingUnprocessableEntity {
	return &CreateanewVoicemailGreetingUnprocessableEntity{}
}

/*CreateanewVoicemailGreetingUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type CreateanewVoicemailGreetingUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *CreateanewVoicemailGreetingUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /voicemailgreetings][%d] createanewVoicemailGreetingUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateanewVoicemailGreetingUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *CreateanewVoicemailGreetingUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
