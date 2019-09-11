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

// SubscriberregistrationsByIDDeleteReader is a Reader for the SubscriberregistrationsByIDDelete structure.
type SubscriberregistrationsByIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubscriberregistrationsByIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSubscriberregistrationsByIDDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSubscriberregistrationsByIDDeleteNoContent creates a SubscriberregistrationsByIDDeleteNoContent with default headers values
func NewSubscriberregistrationsByIDDeleteNoContent() *SubscriberregistrationsByIDDeleteNoContent {
	return &SubscriberregistrationsByIDDeleteNoContent{}
}

/*SubscriberregistrationsByIDDeleteNoContent handles this case with default header values.

SubscriberregistrationsByIDDeleteNoContent subscriberregistrations by Id delete no content
*/
type SubscriberregistrationsByIDDeleteNoContent struct {
	Payload interface{}
}

func (o *SubscriberregistrationsByIDDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /subscriberregistrations/{id}][%d] subscriberregistrationsByIdDeleteNoContent  %+v", 204, o.Payload)
}

func (o *SubscriberregistrationsByIDDeleteNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *SubscriberregistrationsByIDDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}