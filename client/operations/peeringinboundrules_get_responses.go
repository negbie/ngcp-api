// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PeeringinboundrulesGetReader is a Reader for the PeeringinboundrulesGet structure.
type PeeringinboundrulesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PeeringinboundrulesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPeeringinboundrulesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPeeringinboundrulesGetOK creates a PeeringinboundrulesGetOK with default headers values
func NewPeeringinboundrulesGetOK() *PeeringinboundrulesGetOK {
	return &PeeringinboundrulesGetOK{}
}

/*PeeringinboundrulesGetOK handles this case with default header values.

PeeringinboundrulesGetOK peeringinboundrules get o k
*/
type PeeringinboundrulesGetOK struct {
	Payload string
}

func (o *PeeringinboundrulesGetOK) Error() string {
	return fmt.Sprintf("[GET /peeringinboundrules][%d] peeringinboundrulesGetOK  %+v", 200, o.Payload)
}

func (o *PeeringinboundrulesGetOK) GetPayload() string {
	return o.Payload
}

func (o *PeeringinboundrulesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}