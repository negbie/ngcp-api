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

// DeleteaspecificBillingFeeReader is a Reader for the DeleteaspecificBillingFee structure.
type DeleteaspecificBillingFeeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteaspecificBillingFeeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteaspecificBillingFeeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteaspecificBillingFeeNoContent creates a DeleteaspecificBillingFeeNoContent with default headers values
func NewDeleteaspecificBillingFeeNoContent() *DeleteaspecificBillingFeeNoContent {
	return &DeleteaspecificBillingFeeNoContent{}
}

/*DeleteaspecificBillingFeeNoContent handles this case with default header values.

DeleteaspecificBillingFeeNoContent deleteaspecific billing fee no content
*/
type DeleteaspecificBillingFeeNoContent struct {
	Payload interface{}
}

func (o *DeleteaspecificBillingFeeNoContent) Error() string {
	return fmt.Sprintf("[DELETE /billingfees/{id}][%d] deleteaspecificBillingFeeNoContent  %+v", 204, o.Payload)
}

func (o *DeleteaspecificBillingFeeNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteaspecificBillingFeeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
