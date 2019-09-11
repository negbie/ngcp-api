// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/negbie/ngcp-api/models"
)

// InvoicesGetReader is a Reader for the InvoicesGet structure.
type InvoicesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InvoicesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInvoicesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInvoicesGetOK creates a InvoicesGetOK with default headers values
func NewInvoicesGetOK() *InvoicesGetOK {
	return &InvoicesGetOK{}
}

/*InvoicesGetOK handles this case with default header values.

InvoicesGetOK invoices get o k
*/
type InvoicesGetOK struct {
	Payload []*models.Invoices
}

func (o *InvoicesGetOK) Error() string {
	return fmt.Sprintf("[GET /invoices][%d] invoicesGetOK  %+v", 200, o.Payload)
}

func (o *InvoicesGetOK) GetPayload() []*models.Invoices {
	return o.Payload
}

func (o *InvoicesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}