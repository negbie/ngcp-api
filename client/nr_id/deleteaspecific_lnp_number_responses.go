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

// DeleteaspecificLnpNumberReader is a Reader for the DeleteaspecificLnpNumber structure.
type DeleteaspecificLnpNumberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteaspecificLnpNumberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteaspecificLnpNumberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteaspecificLnpNumberNoContent creates a DeleteaspecificLnpNumberNoContent with default headers values
func NewDeleteaspecificLnpNumberNoContent() *DeleteaspecificLnpNumberNoContent {
	return &DeleteaspecificLnpNumberNoContent{}
}

/*DeleteaspecificLnpNumberNoContent handles this case with default header values.

DeleteaspecificLnpNumberNoContent deleteaspecific lnp number no content
*/
type DeleteaspecificLnpNumberNoContent struct {
	Payload interface{}
}

func (o *DeleteaspecificLnpNumberNoContent) Error() string {
	return fmt.Sprintf("[DELETE /lnpnumbers/{id}][%d] deleteaspecificLnpNumberNoContent  %+v", 204, o.Payload)
}

func (o *DeleteaspecificLnpNumberNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteaspecificLnpNumberNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}