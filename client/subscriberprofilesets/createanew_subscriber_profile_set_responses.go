// Code generated by go-swagger; DO NOT EDIT.

package subscriberprofilesets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// CreateanewSubscriberProfileSetReader is a Reader for the CreateanewSubscriberProfileSet structure.
type CreateanewSubscriberProfileSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateanewSubscriberProfileSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateanewSubscriberProfileSetCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateanewSubscriberProfileSetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateanewSubscriberProfileSetCreated creates a CreateanewSubscriberProfileSetCreated with default headers values
func NewCreateanewSubscriberProfileSetCreated() *CreateanewSubscriberProfileSetCreated {
	return &CreateanewSubscriberProfileSetCreated{
		Location: "nostrud velit non",
	}
}

/*CreateanewSubscriberProfileSetCreated handles this case with default header values.

CreateanewSubscriberProfileSetCreated createanew subscriber profile set created
*/
type CreateanewSubscriberProfileSetCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty11
}

func (o *CreateanewSubscriberProfileSetCreated) Error() string {
	return fmt.Sprintf("[POST /subscriberprofilesets][%d] createanewSubscriberProfileSetCreated  %+v", 201, o.Payload)
}

func (o *CreateanewSubscriberProfileSetCreated) GetPayload() []*models.Thenewlycreateditemorempty11 {
	return o.Payload
}

func (o *CreateanewSubscriberProfileSetCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateanewSubscriberProfileSetUnprocessableEntity creates a CreateanewSubscriberProfileSetUnprocessableEntity with default headers values
func NewCreateanewSubscriberProfileSetUnprocessableEntity() *CreateanewSubscriberProfileSetUnprocessableEntity {
	return &CreateanewSubscriberProfileSetUnprocessableEntity{}
}

/*CreateanewSubscriberProfileSetUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type CreateanewSubscriberProfileSetUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *CreateanewSubscriberProfileSetUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /subscriberprofilesets][%d] createanewSubscriberProfileSetUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateanewSubscriberProfileSetUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *CreateanewSubscriberProfileSetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
