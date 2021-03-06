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

// DeleteaspecificCFBNumberSetReader is a Reader for the DeleteaspecificCFBNumberSet structure.
type DeleteaspecificCFBNumberSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteaspecificCFBNumberSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteaspecificCFBNumberSetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteaspecificCFBNumberSetNoContent creates a DeleteaspecificCFBNumberSetNoContent with default headers values
func NewDeleteaspecificCFBNumberSetNoContent() *DeleteaspecificCFBNumberSetNoContent {
	return &DeleteaspecificCFBNumberSetNoContent{}
}

/*DeleteaspecificCFBNumberSetNoContent handles this case with default header values.

DeleteaspecificCFBNumberSetNoContent deleteaspecific c f b number set no content
*/
type DeleteaspecificCFBNumberSetNoContent struct {
	Payload interface{}
}

func (o *DeleteaspecificCFBNumberSetNoContent) Error() string {
	return fmt.Sprintf("[DELETE /cfbnumbersets/{id}][%d] deleteaspecificCFBNumberSetNoContent  %+v", 204, o.Payload)
}

func (o *DeleteaspecificCFBNumberSetNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteaspecificCFBNumberSetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
