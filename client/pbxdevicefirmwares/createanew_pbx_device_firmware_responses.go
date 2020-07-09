// Code generated by go-swagger; DO NOT EDIT.

package pbxdevicefirmwares

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// CreateanewPbxDeviceFirmwareReader is a Reader for the CreateanewPbxDeviceFirmware structure.
type CreateanewPbxDeviceFirmwareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateanewPbxDeviceFirmwareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateanewPbxDeviceFirmwareCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateanewPbxDeviceFirmwareUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateanewPbxDeviceFirmwareCreated creates a CreateanewPbxDeviceFirmwareCreated with default headers values
func NewCreateanewPbxDeviceFirmwareCreated() *CreateanewPbxDeviceFirmwareCreated {
	return &CreateanewPbxDeviceFirmwareCreated{
		Location: "nostrud velit non",
	}
}

/*CreateanewPbxDeviceFirmwareCreated handles this case with default header values.

CreateanewPbxDeviceFirmwareCreated createanew pbx device firmware created
*/
type CreateanewPbxDeviceFirmwareCreated struct {
	Location string

	Payload []*models.Thenewlycreateditemorempty15
}

func (o *CreateanewPbxDeviceFirmwareCreated) Error() string {
	return fmt.Sprintf("[POST /pbxdevicefirmwares][%d] createanewPbxDeviceFirmwareCreated  %+v", 201, o.Payload)
}

func (o *CreateanewPbxDeviceFirmwareCreated) GetPayload() []*models.Thenewlycreateditemorempty15 {
	return o.Payload
}

func (o *CreateanewPbxDeviceFirmwareCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateanewPbxDeviceFirmwareUnprocessableEntity creates a CreateanewPbxDeviceFirmwareUnprocessableEntity with default headers values
func NewCreateanewPbxDeviceFirmwareUnprocessableEntity() *CreateanewPbxDeviceFirmwareUnprocessableEntity {
	return &CreateanewPbxDeviceFirmwareUnprocessableEntity{}
}

/*CreateanewPbxDeviceFirmwareUnprocessableEntity handles this case with default header values.

Unprocessable Entity (WebDAV) (RFC 4918)
*/
type CreateanewPbxDeviceFirmwareUnprocessableEntity struct {
	Payload *models.Anerror
}

func (o *CreateanewPbxDeviceFirmwareUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pbxdevicefirmwares][%d] createanewPbxDeviceFirmwareUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateanewPbxDeviceFirmwareUnprocessableEntity) GetPayload() *models.Anerror {
	return o.Payload
}

func (o *CreateanewPbxDeviceFirmwareUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Anerror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
