// Code generated by go-swagger; DO NOT EDIT.

package balanceintervals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/negbie/ngcp-api/models"
)

// BalanceintervalsByIDGetReader is a Reader for the BalanceintervalsByIDGet structure.
type BalanceintervalsByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BalanceintervalsByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBalanceintervalsByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBalanceintervalsByIDGetOK creates a BalanceintervalsByIDGetOK with default headers values
func NewBalanceintervalsByIDGetOK() *BalanceintervalsByIDGetOK {
	return &BalanceintervalsByIDGetOK{}
}

/*BalanceintervalsByIDGetOK handles this case with default header values.

BalanceintervalsByIDGetOK balanceintervals by Id get o k
*/
type BalanceintervalsByIDGetOK struct {
	Payload *models.BalanceIntervals
}

func (o *BalanceintervalsByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /balanceintervals/{id}][%d] balanceintervalsByIdGetOK  %+v", 200, o.Payload)
}

func (o *BalanceintervalsByIDGetOK) GetPayload() *models.BalanceIntervals {
	return o.Payload
}

func (o *BalanceintervalsByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BalanceIntervals)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}