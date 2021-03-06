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

// ReplaceChangeaspecificCCMapEntryReader is a Reader for the ReplaceChangeaspecificCCMapEntry structure.
type ReplaceChangeaspecificCCMapEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceChangeaspecificCCMapEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewReplaceChangeaspecificCCMapEntryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceChangeaspecificCCMapEntryNoContent creates a ReplaceChangeaspecificCCMapEntryNoContent with default headers values
func NewReplaceChangeaspecificCCMapEntryNoContent() *ReplaceChangeaspecificCCMapEntryNoContent {
	return &ReplaceChangeaspecificCCMapEntryNoContent{}
}

/*ReplaceChangeaspecificCCMapEntryNoContent handles this case with default header values.

ReplaceChangeaspecificCCMapEntryNoContent replace changeaspecific c c map entry no content
*/
type ReplaceChangeaspecificCCMapEntryNoContent struct {
	Payload interface{}
}

func (o *ReplaceChangeaspecificCCMapEntryNoContent) Error() string {
	return fmt.Sprintf("[PUT /ccmapentries/{id}][%d] replaceChangeaspecificCCMapEntryNoContent  %+v", 204, o.Payload)
}

func (o *ReplaceChangeaspecificCCMapEntryNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ReplaceChangeaspecificCCMapEntryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
