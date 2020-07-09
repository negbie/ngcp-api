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

// ReplaceChangeaspecificReminderReader is a Reader for the ReplaceChangeaspecificReminder structure.
type ReplaceChangeaspecificReminderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceChangeaspecificReminderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewReplaceChangeaspecificReminderNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceChangeaspecificReminderNoContent creates a ReplaceChangeaspecificReminderNoContent with default headers values
func NewReplaceChangeaspecificReminderNoContent() *ReplaceChangeaspecificReminderNoContent {
	return &ReplaceChangeaspecificReminderNoContent{}
}

/*ReplaceChangeaspecificReminderNoContent handles this case with default header values.

ReplaceChangeaspecificReminderNoContent replace changeaspecific reminder no content
*/
type ReplaceChangeaspecificReminderNoContent struct {
	Payload interface{}
}

func (o *ReplaceChangeaspecificReminderNoContent) Error() string {
	return fmt.Sprintf("[PUT /reminders/{id}][%d] replaceChangeaspecificReminderNoContent  %+v", 204, o.Payload)
}

func (o *ReplaceChangeaspecificReminderNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ReplaceChangeaspecificReminderNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
