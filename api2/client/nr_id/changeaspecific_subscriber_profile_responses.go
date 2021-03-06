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

// ChangeaspecificSubscriberProfileReader is a Reader for the ChangeaspecificSubscriberProfile structure.
type ChangeaspecificSubscriberProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeaspecificSubscriberProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewChangeaspecificSubscriberProfileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeaspecificSubscriberProfileNoContent creates a ChangeaspecificSubscriberProfileNoContent with default headers values
func NewChangeaspecificSubscriberProfileNoContent() *ChangeaspecificSubscriberProfileNoContent {
	return &ChangeaspecificSubscriberProfileNoContent{}
}

/*ChangeaspecificSubscriberProfileNoContent handles this case with default header values.

ChangeaspecificSubscriberProfileNoContent changeaspecific subscriber profile no content
*/
type ChangeaspecificSubscriberProfileNoContent struct {
	Payload interface{}
}

func (o *ChangeaspecificSubscriberProfileNoContent) Error() string {
	return fmt.Sprintf("[PATCH /subscriberprofiles/{id}][%d] changeaspecificSubscriberProfileNoContent  %+v", 204, o.Payload)
}

func (o *ChangeaspecificSubscriberProfileNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ChangeaspecificSubscriberProfileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
