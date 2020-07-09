// Code generated by go-swagger; DO NOT EDIT.

package emailtemplates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// CreateanewEmailTemplateReader is a Reader for the CreateanewEmailTemplate structure.
type CreateanewEmailTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateanewEmailTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateanewEmailTemplateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateanewEmailTemplateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateanewEmailTemplateCreated creates a CreateanewEmailTemplateCreated with default headers values
func NewCreateanewEmailTemplateCreated() *CreateanewEmailTemplateCreated {
	return &CreateanewEmailTemplateCreated{
		Location: "nostrud velit non",
	}
}

/*CreateanewEmailTemplateCreated handles this case with default header values.

CreateanewEmailTemplateCreated createanew email template created
*/
type CreateanewEmailTemplateCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty21
}

func (o *CreateanewEmailTemplateCreated) Error() string {
	return fmt.Sprintf("[POST /emailtemplates][%d] createanewEmailTemplateCreated  %+v", 201, o.Payload)
}

func (o *CreateanewEmailTemplateCreated) GetPayload() []*models.Thenewlycreateditemorempty21 {
	return o.Payload
}

func (o *CreateanewEmailTemplateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateanewEmailTemplateUnprocessableEntity creates a CreateanewEmailTemplateUnprocessableEntity with default headers values
func NewCreateanewEmailTemplateUnprocessableEntity() *CreateanewEmailTemplateUnprocessableEntity {
	return &CreateanewEmailTemplateUnprocessableEntity{}
}

/*CreateanewEmailTemplateUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type CreateanewEmailTemplateUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *CreateanewEmailTemplateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /emailtemplates][%d] createanewEmailTemplateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateanewEmailTemplateUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *CreateanewEmailTemplateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
