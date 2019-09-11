// Code generated by go-swagger; DO NOT EDIT.

package id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NcospatternsByIDGetReader is a Reader for the NcospatternsByIDGet structure.
type NcospatternsByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NcospatternsByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNcospatternsByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNcospatternsByIDGetOK creates a NcospatternsByIDGetOK with default headers values
func NewNcospatternsByIDGetOK() *NcospatternsByIDGetOK {
	return &NcospatternsByIDGetOK{}
}

/*NcospatternsByIDGetOK handles this case with default header values.

NcospatternsByIDGetOK ncospatterns by Id get o k
*/
type NcospatternsByIDGetOK struct {
	Payload string
}

func (o *NcospatternsByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /ncospatterns/{id}][%d] ncospatternsByIdGetOK  %+v", 200, o.Payload)
}

func (o *NcospatternsByIDGetOK) GetPayload() string {
	return o.Payload
}

func (o *NcospatternsByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}