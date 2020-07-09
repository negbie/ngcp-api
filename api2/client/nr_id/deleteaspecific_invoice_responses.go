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

// DeleteaspecificInvoiceReader is a Reader for the DeleteaspecificInvoice structure.
type DeleteaspecificInvoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteaspecificInvoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteaspecificInvoiceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteaspecificInvoiceNoContent creates a DeleteaspecificInvoiceNoContent with default headers values
func NewDeleteaspecificInvoiceNoContent() *DeleteaspecificInvoiceNoContent {
	return &DeleteaspecificInvoiceNoContent{}
}

/*DeleteaspecificInvoiceNoContent handles this case with default header values.

DeleteaspecificInvoiceNoContent deleteaspecific invoice no content
*/
type DeleteaspecificInvoiceNoContent struct {
	Payload interface{}
}

func (o *DeleteaspecificInvoiceNoContent) Error() string {
	return fmt.Sprintf("[DELETE /invoices/{id}][%d] deleteaspecificInvoiceNoContent  %+v", 204, o.Payload)
}

func (o *DeleteaspecificInvoiceNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteaspecificInvoiceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}