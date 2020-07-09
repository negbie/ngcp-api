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

// GetaspecificCustomerPreferenceReader is a Reader for the GetaspecificCustomerPreference structure.
type GetaspecificCustomerPreferenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificCustomerPreferenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificCustomerPreferenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificCustomerPreferenceOK creates a GetaspecificCustomerPreferenceOK with default headers values
func NewGetaspecificCustomerPreferenceOK() *GetaspecificCustomerPreferenceOK {
	return &GetaspecificCustomerPreferenceOK{}
}

/*GetaspecificCustomerPreferenceOK handles this case with default header values.

GetaspecificCustomerPreferenceOK getaspecific customer preference o k
*/
type GetaspecificCustomerPreferenceOK struct {
	Payload string
}

func (o *GetaspecificCustomerPreferenceOK) Error() string {
	return fmt.Sprintf("[GET /customerpreferences/{id}][%d] getaspecificCustomerPreferenceOK  %+v", 200, o.Payload)
}

func (o *GetaspecificCustomerPreferenceOK) GetPayload() string {
	return o.Payload
}

func (o *GetaspecificCustomerPreferenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}