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

// DeleteaspecificCCMapEntryReader is a Reader for the DeleteaspecificCCMapEntry structure.
type DeleteaspecificCCMapEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteaspecificCCMapEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteaspecificCCMapEntryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteaspecificCCMapEntryNoContent creates a DeleteaspecificCCMapEntryNoContent with default headers values
func NewDeleteaspecificCCMapEntryNoContent() *DeleteaspecificCCMapEntryNoContent {
	return &DeleteaspecificCCMapEntryNoContent{}
}

/*DeleteaspecificCCMapEntryNoContent handles this case with default header values.

DeleteaspecificCCMapEntryNoContent deleteaspecific c c map entry no content
*/
type DeleteaspecificCCMapEntryNoContent struct {
	Payload interface{}
}

func (o *DeleteaspecificCCMapEntryNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ccmapentries/{id}][%d] deleteaspecificCCMapEntryNoContent  %+v", 204, o.Payload)
}

func (o *DeleteaspecificCCMapEntryNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteaspecificCCMapEntryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
