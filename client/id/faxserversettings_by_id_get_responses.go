// Code generated by go-swagger; DO NOT EDIT.

package id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/negbie/ngcp-api/models"
)

// FaxserversettingsByIDGetReader is a Reader for the FaxserversettingsByIDGet structure.
type FaxserversettingsByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FaxserversettingsByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFaxserversettingsByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewFaxserversettingsByIDGetOK creates a FaxserversettingsByIDGetOK with default headers values
func NewFaxserversettingsByIDGetOK() *FaxserversettingsByIDGetOK {
	return &FaxserversettingsByIDGetOK{}
}

/*FaxserversettingsByIDGetOK handles this case with default header values.

FaxserversettingsByIDGetOK faxserversettings by Id get o k
*/
type FaxserversettingsByIDGetOK struct {
	Payload *models.FaxserverSettings
}

func (o *FaxserversettingsByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /faxserversettings/{id}][%d] faxserversettingsByIdGetOK  %+v", 200, o.Payload)
}

func (o *FaxserversettingsByIDGetOK) GetPayload() *models.FaxserverSettings {
	return o.Payload
}

func (o *FaxserversettingsByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FaxserverSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}