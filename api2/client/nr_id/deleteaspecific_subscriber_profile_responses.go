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

// DeleteaspecificSubscriberProfileReader is a Reader for the DeleteaspecificSubscriberProfile structure.
type DeleteaspecificSubscriberProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteaspecificSubscriberProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteaspecificSubscriberProfileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteaspecificSubscriberProfileNoContent creates a DeleteaspecificSubscriberProfileNoContent with default headers values
func NewDeleteaspecificSubscriberProfileNoContent() *DeleteaspecificSubscriberProfileNoContent {
	return &DeleteaspecificSubscriberProfileNoContent{}
}

/*DeleteaspecificSubscriberProfileNoContent handles this case with default header values.

DeleteaspecificSubscriberProfileNoContent deleteaspecific subscriber profile no content
*/
type DeleteaspecificSubscriberProfileNoContent struct {
	Payload interface{}
}

func (o *DeleteaspecificSubscriberProfileNoContent) Error() string {
	return fmt.Sprintf("[DELETE /subscriberprofiles/{id}][%d] deleteaspecificSubscriberProfileNoContent  %+v", 204, o.Payload)
}

func (o *DeleteaspecificSubscriberProfileNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteaspecificSubscriberProfileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
