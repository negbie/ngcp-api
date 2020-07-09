// Code generated by go-swagger; DO NOT EDIT.

package ncoslevels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// CreateanewNcosLevelReader is a Reader for the CreateanewNcosLevel structure.
type CreateanewNcosLevelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateanewNcosLevelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateanewNcosLevelCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateanewNcosLevelUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateanewNcosLevelCreated creates a CreateanewNcosLevelCreated with default headers values
func NewCreateanewNcosLevelCreated() *CreateanewNcosLevelCreated {
	return &CreateanewNcosLevelCreated{
		Location: "nostrud velit non",
	}
}

/*CreateanewNcosLevelCreated handles this case with default header values.

CreateanewNcosLevelCreated createanew ncos level created
*/
type CreateanewNcosLevelCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty32
}

func (o *CreateanewNcosLevelCreated) Error() string {
	return fmt.Sprintf("[POST /ncoslevels][%d] createanewNcosLevelCreated  %+v", 201, o.Payload)
}

func (o *CreateanewNcosLevelCreated) GetPayload() []*models.Thenewlycreateditemorempty32 {
	return o.Payload
}

func (o *CreateanewNcosLevelCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateanewNcosLevelUnprocessableEntity creates a CreateanewNcosLevelUnprocessableEntity with default headers values
func NewCreateanewNcosLevelUnprocessableEntity() *CreateanewNcosLevelUnprocessableEntity {
	return &CreateanewNcosLevelUnprocessableEntity{}
}

/*CreateanewNcosLevelUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type CreateanewNcosLevelUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *CreateanewNcosLevelUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /ncoslevels][%d] createanewNcosLevelUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateanewNcosLevelUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *CreateanewNcosLevelUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}