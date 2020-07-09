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

// DeleteaspecificCustomerContactReader is a Reader for the DeleteaspecificCustomerContact structure.
type DeleteaspecificCustomerContactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteaspecificCustomerContactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteaspecificCustomerContactNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteaspecificCustomerContactNoContent creates a DeleteaspecificCustomerContactNoContent with default headers values
func NewDeleteaspecificCustomerContactNoContent() *DeleteaspecificCustomerContactNoContent {
	return &DeleteaspecificCustomerContactNoContent{}
}

/*DeleteaspecificCustomerContactNoContent handles this case with default header values.

DeleteaspecificCustomerContactNoContent deleteaspecific customer contact no content
*/
type DeleteaspecificCustomerContactNoContent struct {
	Payload interface{}
}

func (o *DeleteaspecificCustomerContactNoContent) Error() string {
	return fmt.Sprintf("[DELETE /customercontacts/{id}][%d] deleteaspecificCustomerContactNoContent  %+v", 204, o.Payload)
}

func (o *DeleteaspecificCustomerContactNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteaspecificCustomerContactNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
