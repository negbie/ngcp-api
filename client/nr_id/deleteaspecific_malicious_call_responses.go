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

// DeleteaspecificMaliciousCallReader is a Reader for the DeleteaspecificMaliciousCall structure.
type DeleteaspecificMaliciousCallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteaspecificMaliciousCallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteaspecificMaliciousCallNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteaspecificMaliciousCallNoContent creates a DeleteaspecificMaliciousCallNoContent with default headers values
func NewDeleteaspecificMaliciousCallNoContent() *DeleteaspecificMaliciousCallNoContent {
	return &DeleteaspecificMaliciousCallNoContent{}
}

/*DeleteaspecificMaliciousCallNoContent handles this case with default header values.

DeleteaspecificMaliciousCallNoContent deleteaspecific malicious call no content
*/
type DeleteaspecificMaliciousCallNoContent struct {
	Payload interface{}
}

func (o *DeleteaspecificMaliciousCallNoContent) Error() string {
	return fmt.Sprintf("[DELETE /maliciouscalls/{id}][%d] deleteaspecificMaliciousCallNoContent  %+v", 204, o.Payload)
}

func (o *DeleteaspecificMaliciousCallNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteaspecificMaliciousCallNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}