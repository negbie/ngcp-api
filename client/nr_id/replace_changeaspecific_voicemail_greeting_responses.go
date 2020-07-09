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

// ReplaceChangeaspecificVoicemailGreetingReader is a Reader for the ReplaceChangeaspecificVoicemailGreeting structure.
type ReplaceChangeaspecificVoicemailGreetingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceChangeaspecificVoicemailGreetingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewReplaceChangeaspecificVoicemailGreetingNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceChangeaspecificVoicemailGreetingNoContent creates a ReplaceChangeaspecificVoicemailGreetingNoContent with default headers values
func NewReplaceChangeaspecificVoicemailGreetingNoContent() *ReplaceChangeaspecificVoicemailGreetingNoContent {
	return &ReplaceChangeaspecificVoicemailGreetingNoContent{}
}

/*ReplaceChangeaspecificVoicemailGreetingNoContent handles this case with default header values.

ReplaceChangeaspecificVoicemailGreetingNoContent replace changeaspecific voicemail greeting no content
*/
type ReplaceChangeaspecificVoicemailGreetingNoContent struct {
	Payload interface{}
}

func (o *ReplaceChangeaspecificVoicemailGreetingNoContent) Error() string {
	return fmt.Sprintf("[PUT /voicemailgreetings/{id}][%d] replaceChangeaspecificVoicemailGreetingNoContent  %+v", 204, o.Payload)
}

func (o *ReplaceChangeaspecificVoicemailGreetingNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ReplaceChangeaspecificVoicemailGreetingNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
