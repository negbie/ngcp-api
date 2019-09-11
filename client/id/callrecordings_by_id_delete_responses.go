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

// CallrecordingsByIDDeleteReader is a Reader for the CallrecordingsByIDDelete structure.
type CallrecordingsByIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CallrecordingsByIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCallrecordingsByIDDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCallrecordingsByIDDeleteNoContent creates a CallrecordingsByIDDeleteNoContent with default headers values
func NewCallrecordingsByIDDeleteNoContent() *CallrecordingsByIDDeleteNoContent {
	return &CallrecordingsByIDDeleteNoContent{}
}

/*CallrecordingsByIDDeleteNoContent handles this case with default header values.

CallrecordingsByIDDeleteNoContent callrecordings by Id delete no content
*/
type CallrecordingsByIDDeleteNoContent struct {
	Payload interface{}
}

func (o *CallrecordingsByIDDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /callrecordings/{id}][%d] callrecordingsByIdDeleteNoContent  %+v", 204, o.Payload)
}

func (o *CallrecordingsByIDDeleteNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CallrecordingsByIDDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}