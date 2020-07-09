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

// DeleteaspecificPeeringGroupReader is a Reader for the DeleteaspecificPeeringGroup structure.
type DeleteaspecificPeeringGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteaspecificPeeringGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteaspecificPeeringGroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteaspecificPeeringGroupNoContent creates a DeleteaspecificPeeringGroupNoContent with default headers values
func NewDeleteaspecificPeeringGroupNoContent() *DeleteaspecificPeeringGroupNoContent {
	return &DeleteaspecificPeeringGroupNoContent{}
}

/*DeleteaspecificPeeringGroupNoContent handles this case with default header values.

DeleteaspecificPeeringGroupNoContent deleteaspecific peering group no content
*/
type DeleteaspecificPeeringGroupNoContent struct {
	Payload interface{}
}

func (o *DeleteaspecificPeeringGroupNoContent) Error() string {
	return fmt.Sprintf("[DELETE /peeringgroups/{id}][%d] deleteaspecificPeeringGroupNoContent  %+v", 204, o.Payload)
}

func (o *DeleteaspecificPeeringGroupNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteaspecificPeeringGroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}