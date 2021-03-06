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

// DeleteaspecificSoundSetReader is a Reader for the DeleteaspecificSoundSet structure.
type DeleteaspecificSoundSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteaspecificSoundSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteaspecificSoundSetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteaspecificSoundSetNoContent creates a DeleteaspecificSoundSetNoContent with default headers values
func NewDeleteaspecificSoundSetNoContent() *DeleteaspecificSoundSetNoContent {
	return &DeleteaspecificSoundSetNoContent{}
}

/*DeleteaspecificSoundSetNoContent handles this case with default header values.

DeleteaspecificSoundSetNoContent deleteaspecific sound set no content
*/
type DeleteaspecificSoundSetNoContent struct {
	Payload interface{}
}

func (o *DeleteaspecificSoundSetNoContent) Error() string {
	return fmt.Sprintf("[DELETE /soundsets/{id}][%d] deleteaspecificSoundSetNoContent  %+v", 204, o.Payload)
}

func (o *DeleteaspecificSoundSetNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteaspecificSoundSetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
