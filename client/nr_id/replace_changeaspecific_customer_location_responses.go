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

// ReplaceChangeaspecificCustomerLocationReader is a Reader for the ReplaceChangeaspecificCustomerLocation structure.
type ReplaceChangeaspecificCustomerLocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceChangeaspecificCustomerLocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewReplaceChangeaspecificCustomerLocationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceChangeaspecificCustomerLocationNoContent creates a ReplaceChangeaspecificCustomerLocationNoContent with default headers values
func NewReplaceChangeaspecificCustomerLocationNoContent() *ReplaceChangeaspecificCustomerLocationNoContent {
	return &ReplaceChangeaspecificCustomerLocationNoContent{}
}

/*ReplaceChangeaspecificCustomerLocationNoContent handles this case with default header values.

ReplaceChangeaspecificCustomerLocationNoContent replace changeaspecific customer location no content
*/
type ReplaceChangeaspecificCustomerLocationNoContent struct {
	Payload interface{}
}

func (o *ReplaceChangeaspecificCustomerLocationNoContent) Error() string {
	return fmt.Sprintf("[PUT /customerlocations/{id}][%d] replaceChangeaspecificCustomerLocationNoContent  %+v", 204, o.Payload)
}

func (o *ReplaceChangeaspecificCustomerLocationNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ReplaceChangeaspecificCustomerLocationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
