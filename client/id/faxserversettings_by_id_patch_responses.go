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

// FaxserversettingsByIDPatchReader is a Reader for the FaxserversettingsByIDPatch structure.
type FaxserversettingsByIDPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FaxserversettingsByIDPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFaxserversettingsByIDPatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewFaxserversettingsByIDPatchOK creates a FaxserversettingsByIDPatchOK with default headers values
func NewFaxserversettingsByIDPatchOK() *FaxserversettingsByIDPatchOK {
	return &FaxserversettingsByIDPatchOK{}
}

/*FaxserversettingsByIDPatchOK handles this case with default header values.

FaxserversettingsByIDPatchOK faxserversettings by Id patch o k
*/
type FaxserversettingsByIDPatchOK struct {
	Payload *models.FaxserverSettings
}

func (o *FaxserversettingsByIDPatchOK) Error() string {
	return fmt.Sprintf("[PATCH /faxserversettings/{id}][%d] faxserversettingsByIdPatchOK  %+v", 200, o.Payload)
}

func (o *FaxserversettingsByIDPatchOK) GetPayload() *models.FaxserverSettings {
	return o.Payload
}

func (o *FaxserversettingsByIDPatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FaxserverSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}