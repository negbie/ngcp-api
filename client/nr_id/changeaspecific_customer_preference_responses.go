// Code generated by go-swagger; DO NOT EDIT.

package nr_id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ChangeaspecificCustomerPreferenceReader is a Reader for the ChangeaspecificCustomerPreference structure.
type ChangeaspecificCustomerPreferenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeaspecificCustomerPreferenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeaspecificCustomerPreferenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeaspecificCustomerPreferenceOK creates a ChangeaspecificCustomerPreferenceOK with default headers values
func NewChangeaspecificCustomerPreferenceOK() *ChangeaspecificCustomerPreferenceOK {
	return &ChangeaspecificCustomerPreferenceOK{}
}

/*ChangeaspecificCustomerPreferenceOK handles this case with default header values.

ChangeaspecificCustomerPreferenceOK changeaspecific customer preference o k
*/
type ChangeaspecificCustomerPreferenceOK struct {
	Payload string
}

func (o *ChangeaspecificCustomerPreferenceOK) Error() string {
	return fmt.Sprintf("[PATCH /customerpreferences/{id}][%d] changeaspecificCustomerPreferenceOK  %+v", 200, o.Payload)
}

func (o *ChangeaspecificCustomerPreferenceOK) GetPayload() string {
	return o.Payload
}

func (o *ChangeaspecificCustomerPreferenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}