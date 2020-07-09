// Code generated by go-swagger; DO NOT EDIT.

package subscriberprofiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// CreateanewSubscriberProfileReader is a Reader for the CreateanewSubscriberProfile structure.
type CreateanewSubscriberProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateanewSubscriberProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateanewSubscriberProfileCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateanewSubscriberProfileUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateanewSubscriberProfileCreated creates a CreateanewSubscriberProfileCreated with default headers values
func NewCreateanewSubscriberProfileCreated() *CreateanewSubscriberProfileCreated {
	return &CreateanewSubscriberProfileCreated{
		Location: "nostrud velit non",
	}
}

/*CreateanewSubscriberProfileCreated handles this case with default header values.

CreateanewSubscriberProfileCreated createanew subscriber profile created
*/
type CreateanewSubscriberProfileCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty26
}

func (o *CreateanewSubscriberProfileCreated) Error() string {
	return fmt.Sprintf("[POST /subscriberprofiles][%d] createanewSubscriberProfileCreated  %+v", 201, o.Payload)
}

func (o *CreateanewSubscriberProfileCreated) GetPayload() []*models.Thenewlycreateditemorempty26 {
	return o.Payload
}

func (o *CreateanewSubscriberProfileCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateanewSubscriberProfileUnprocessableEntity creates a CreateanewSubscriberProfileUnprocessableEntity with default headers values
func NewCreateanewSubscriberProfileUnprocessableEntity() *CreateanewSubscriberProfileUnprocessableEntity {
	return &CreateanewSubscriberProfileUnprocessableEntity{}
}

/*CreateanewSubscriberProfileUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type CreateanewSubscriberProfileUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *CreateanewSubscriberProfileUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /subscriberprofiles][%d] createanewSubscriberProfileUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateanewSubscriberProfileUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *CreateanewSubscriberProfileUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
