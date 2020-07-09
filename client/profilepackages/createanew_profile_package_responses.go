// Code generated by go-swagger; DO NOT EDIT.

package profilepackages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// CreateanewProfilePackageReader is a Reader for the CreateanewProfilePackage structure.
type CreateanewProfilePackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateanewProfilePackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateanewProfilePackageCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateanewProfilePackageUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateanewProfilePackageCreated creates a CreateanewProfilePackageCreated with default headers values
func NewCreateanewProfilePackageCreated() *CreateanewProfilePackageCreated {
	return &CreateanewProfilePackageCreated{
		Location: "nostrud velit non",
	}
}

/*CreateanewProfilePackageCreated handles this case with default header values.

CreateanewProfilePackageCreated createanew profile package created
*/
type CreateanewProfilePackageCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty40
}

func (o *CreateanewProfilePackageCreated) Error() string {
	return fmt.Sprintf("[POST /profilepackages][%d] createanewProfilePackageCreated  %+v", 201, o.Payload)
}

func (o *CreateanewProfilePackageCreated) GetPayload() []*models.Thenewlycreateditemorempty40 {
	return o.Payload
}

func (o *CreateanewProfilePackageCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateanewProfilePackageUnprocessableEntity creates a CreateanewProfilePackageUnprocessableEntity with default headers values
func NewCreateanewProfilePackageUnprocessableEntity() *CreateanewProfilePackageUnprocessableEntity {
	return &CreateanewProfilePackageUnprocessableEntity{}
}

/*CreateanewProfilePackageUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type CreateanewProfilePackageUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *CreateanewProfilePackageUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /profilepackages][%d] createanewProfilePackageUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateanewProfilePackageUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *CreateanewProfilePackageUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
