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

// PbxfielddevicepreferencesByIDPutReader is a Reader for the PbxfielddevicepreferencesByIDPut structure.
type PbxfielddevicepreferencesByIDPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PbxfielddevicepreferencesByIDPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPbxfielddevicepreferencesByIDPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPbxfielddevicepreferencesByIDPutOK creates a PbxfielddevicepreferencesByIDPutOK with default headers values
func NewPbxfielddevicepreferencesByIDPutOK() *PbxfielddevicepreferencesByIDPutOK {
	return &PbxfielddevicepreferencesByIDPutOK{}
}

/*PbxfielddevicepreferencesByIDPutOK handles this case with default header values.

PbxfielddevicepreferencesByIDPutOK pbxfielddevicepreferences by Id put o k
*/
type PbxfielddevicepreferencesByIDPutOK struct {
	Payload string
}

func (o *PbxfielddevicepreferencesByIDPutOK) Error() string {
	return fmt.Sprintf("[PUT /pbxfielddevicepreferences/{id}][%d] pbxfielddevicepreferencesByIdPutOK  %+v", 200, o.Payload)
}

func (o *PbxfielddevicepreferencesByIDPutOK) GetPayload() string {
	return o.Payload
}

func (o *PbxfielddevicepreferencesByIDPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}