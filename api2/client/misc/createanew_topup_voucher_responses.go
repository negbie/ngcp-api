// Code generated by go-swagger; DO NOT EDIT.

package misc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// CreateanewTopupVoucherReader is a Reader for the CreateanewTopupVoucher structure.
type CreateanewTopupVoucherReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateanewTopupVoucherReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateanewTopupVoucherCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateanewTopupVoucherUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateanewTopupVoucherCreated creates a CreateanewTopupVoucherCreated with default headers values
func NewCreateanewTopupVoucherCreated() *CreateanewTopupVoucherCreated {
	return &CreateanewTopupVoucherCreated{
		Location: "nostrud velit non",
	}
}

/*CreateanewTopupVoucherCreated handles this case with default header values.

CreateanewTopupVoucherCreated createanew topup voucher created
*/
type CreateanewTopupVoucherCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty50
}

func (o *CreateanewTopupVoucherCreated) Error() string {
	return fmt.Sprintf("[POST /topupvouchers][%d] createanewTopupVoucherCreated  %+v", 201, o.Payload)
}

func (o *CreateanewTopupVoucherCreated) GetPayload() []*models.Thenewlycreateditemorempty50 {
	return o.Payload
}

func (o *CreateanewTopupVoucherCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateanewTopupVoucherUnprocessableEntity creates a CreateanewTopupVoucherUnprocessableEntity with default headers values
func NewCreateanewTopupVoucherUnprocessableEntity() *CreateanewTopupVoucherUnprocessableEntity {
	return &CreateanewTopupVoucherUnprocessableEntity{}
}

/*CreateanewTopupVoucherUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type CreateanewTopupVoucherUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *CreateanewTopupVoucherUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /topupvouchers][%d] createanewTopupVoucherUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateanewTopupVoucherUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *CreateanewTopupVoucherUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}