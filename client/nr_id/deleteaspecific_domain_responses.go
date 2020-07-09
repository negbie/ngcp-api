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

// DeleteaspecificDomainReader is a Reader for the DeleteaspecificDomain structure.
type DeleteaspecificDomainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteaspecificDomainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteaspecificDomainNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteaspecificDomainNoContent creates a DeleteaspecificDomainNoContent with default headers values
func NewDeleteaspecificDomainNoContent() *DeleteaspecificDomainNoContent {
	return &DeleteaspecificDomainNoContent{}
}

/*DeleteaspecificDomainNoContent handles this case with default header values.

DeleteaspecificDomainNoContent deleteaspecific domain no content
*/
type DeleteaspecificDomainNoContent struct {
	Payload interface{}
}

func (o *DeleteaspecificDomainNoContent) Error() string {
	return fmt.Sprintf("[DELETE /domains/{id}][%d] deleteaspecificDomainNoContent  %+v", 204, o.Payload)
}

func (o *DeleteaspecificDomainNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteaspecificDomainNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}