// Code generated by go-swagger; DO NOT EDIT.

package nr_id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// GetaspecificBillingFeeReader is a Reader for the GetaspecificBillingFee structure.
type GetaspecificBillingFeeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificBillingFeeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificBillingFeeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificBillingFeeOK creates a GetaspecificBillingFeeOK with default headers values
func NewGetaspecificBillingFeeOK() *GetaspecificBillingFeeOK {
	return &GetaspecificBillingFeeOK{}
}

/*GetaspecificBillingFeeOK handles this case with default header values.

GetaspecificBillingFeeOK getaspecific billing fee o k
*/
type GetaspecificBillingFeeOK struct {
	Payload *models.BillingFees
}

func (o *GetaspecificBillingFeeOK) Error() string {
	return fmt.Sprintf("[GET /billingfees/{id}][%d] getaspecificBillingFeeOK  %+v", 200, o.Payload)
}

func (o *GetaspecificBillingFeeOK) GetPayload() *models.BillingFees {
	return o.Payload
}

func (o *GetaspecificBillingFeeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BillingFees)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
