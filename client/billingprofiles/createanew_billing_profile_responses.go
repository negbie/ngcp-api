// Code generated by go-swagger; DO NOT EDIT.

package billingprofiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// CreateanewBillingProfileReader is a Reader for the CreateanewBillingProfile structure.
type CreateanewBillingProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateanewBillingProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateanewBillingProfileCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateanewBillingProfileUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateanewBillingProfileCreated creates a CreateanewBillingProfileCreated with default headers values
func NewCreateanewBillingProfileCreated() *CreateanewBillingProfileCreated {
	return &CreateanewBillingProfileCreated{
		Location: "nostrud velit non",
	}
}

/*CreateanewBillingProfileCreated handles this case with default header values.

CreateanewBillingProfileCreated createanew billing profile created
*/
type CreateanewBillingProfileCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty1
}

func (o *CreateanewBillingProfileCreated) Error() string {
	return fmt.Sprintf("[POST /billingprofiles][%d] createanewBillingProfileCreated  %+v", 201, o.Payload)
}

func (o *CreateanewBillingProfileCreated) GetPayload() []*models.Thenewlycreateditemorempty1 {
	return o.Payload
}

func (o *CreateanewBillingProfileCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateanewBillingProfileUnprocessableEntity creates a CreateanewBillingProfileUnprocessableEntity with default headers values
func NewCreateanewBillingProfileUnprocessableEntity() *CreateanewBillingProfileUnprocessableEntity {
	return &CreateanewBillingProfileUnprocessableEntity{}
}

/*CreateanewBillingProfileUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type CreateanewBillingProfileUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *CreateanewBillingProfileUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /billingprofiles][%d] createanewBillingProfileUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateanewBillingProfileUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *CreateanewBillingProfileUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
