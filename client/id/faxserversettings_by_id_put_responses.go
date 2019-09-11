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

// FaxserversettingsByIDPutReader is a Reader for the FaxserversettingsByIDPut structure.
type FaxserversettingsByIDPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FaxserversettingsByIDPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFaxserversettingsByIDPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewFaxserversettingsByIDPutOK creates a FaxserversettingsByIDPutOK with default headers values
func NewFaxserversettingsByIDPutOK() *FaxserversettingsByIDPutOK {
	return &FaxserversettingsByIDPutOK{}
}

/*FaxserversettingsByIDPutOK handles this case with default header values.

FaxserversettingsByIDPutOK faxserversettings by Id put o k
*/
type FaxserversettingsByIDPutOK struct {
	Payload *models.FaxserverSettings
}

func (o *FaxserversettingsByIDPutOK) Error() string {
	return fmt.Sprintf("[PUT /faxserversettings/{id}][%d] faxserversettingsByIdPutOK  %+v", 200, o.Payload)
}

func (o *FaxserversettingsByIDPutOK) GetPayload() *models.FaxserverSettings {
	return o.Payload
}

func (o *FaxserversettingsByIDPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FaxserverSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}