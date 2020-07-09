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

// ChangeaspecificSubscriberProfileSetReader is a Reader for the ChangeaspecificSubscriberProfileSet structure.
type ChangeaspecificSubscriberProfileSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeaspecificSubscriberProfileSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewChangeaspecificSubscriberProfileSetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeaspecificSubscriberProfileSetNoContent creates a ChangeaspecificSubscriberProfileSetNoContent with default headers values
func NewChangeaspecificSubscriberProfileSetNoContent() *ChangeaspecificSubscriberProfileSetNoContent {
	return &ChangeaspecificSubscriberProfileSetNoContent{}
}

/*ChangeaspecificSubscriberProfileSetNoContent handles this case with default header values.

ChangeaspecificSubscriberProfileSetNoContent changeaspecific subscriber profile set no content
*/
type ChangeaspecificSubscriberProfileSetNoContent struct {
	Payload interface{}
}

func (o *ChangeaspecificSubscriberProfileSetNoContent) Error() string {
	return fmt.Sprintf("[PATCH /subscriberprofilesets/{id}][%d] changeaspecificSubscriberProfileSetNoContent  %+v", 204, o.Payload)
}

func (o *ChangeaspecificSubscriberProfileSetNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ChangeaspecificSubscriberProfileSetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
