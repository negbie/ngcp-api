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

// CustomerfraudpreferencesByIDPatchReader is a Reader for the CustomerfraudpreferencesByIDPatch structure.
type CustomerfraudpreferencesByIDPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerfraudpreferencesByIDPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerfraudpreferencesByIDPatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomerfraudpreferencesByIDPatchOK creates a CustomerfraudpreferencesByIDPatchOK with default headers values
func NewCustomerfraudpreferencesByIDPatchOK() *CustomerfraudpreferencesByIDPatchOK {
	return &CustomerfraudpreferencesByIDPatchOK{}
}

/*CustomerfraudpreferencesByIDPatchOK handles this case with default header values.

CustomerfraudpreferencesByIDPatchOK customerfraudpreferences by Id patch o k
*/
type CustomerfraudpreferencesByIDPatchOK struct {
	Payload *models.CustomerFraudPreferences
}

func (o *CustomerfraudpreferencesByIDPatchOK) Error() string {
	return fmt.Sprintf("[PATCH /customerfraudpreferences/{id}][%d] customerfraudpreferencesByIdPatchOK  %+v", 200, o.Payload)
}

func (o *CustomerfraudpreferencesByIDPatchOK) GetPayload() *models.CustomerFraudPreferences {
	return o.Payload
}

func (o *CustomerfraudpreferencesByIDPatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerFraudPreferences)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}