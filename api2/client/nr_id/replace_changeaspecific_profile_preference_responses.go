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

// ReplaceChangeaspecificProfilePreferenceReader is a Reader for the ReplaceChangeaspecificProfilePreference structure.
type ReplaceChangeaspecificProfilePreferenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceChangeaspecificProfilePreferenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewReplaceChangeaspecificProfilePreferenceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceChangeaspecificProfilePreferenceNoContent creates a ReplaceChangeaspecificProfilePreferenceNoContent with default headers values
func NewReplaceChangeaspecificProfilePreferenceNoContent() *ReplaceChangeaspecificProfilePreferenceNoContent {
	return &ReplaceChangeaspecificProfilePreferenceNoContent{}
}

/*ReplaceChangeaspecificProfilePreferenceNoContent handles this case with default header values.

ReplaceChangeaspecificProfilePreferenceNoContent replace changeaspecific profile preference no content
*/
type ReplaceChangeaspecificProfilePreferenceNoContent struct {
	Payload interface{}
}

func (o *ReplaceChangeaspecificProfilePreferenceNoContent) Error() string {
	return fmt.Sprintf("[PUT /profilepreferences/{id}][%d] replaceChangeaspecificProfilePreferenceNoContent  %+v", 204, o.Payload)
}

func (o *ReplaceChangeaspecificProfilePreferenceNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ReplaceChangeaspecificProfilePreferenceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
