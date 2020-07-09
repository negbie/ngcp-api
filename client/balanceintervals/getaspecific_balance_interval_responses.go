// Code generated by go-swagger; DO NOT EDIT.

package balanceintervals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// GetaspecificBalanceIntervalReader is a Reader for the GetaspecificBalanceInterval structure.
type GetaspecificBalanceIntervalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificBalanceIntervalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificBalanceIntervalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificBalanceIntervalOK creates a GetaspecificBalanceIntervalOK with default headers values
func NewGetaspecificBalanceIntervalOK() *GetaspecificBalanceIntervalOK {
	return &GetaspecificBalanceIntervalOK{}
}

/*GetaspecificBalanceIntervalOK handles this case with default header values.

GetaspecificBalanceIntervalOK getaspecific balance interval o k
*/
type GetaspecificBalanceIntervalOK struct {
	Payload *models.BalanceIntervals
}

func (o *GetaspecificBalanceIntervalOK) Error() string {
	return fmt.Sprintf("[GET /balanceintervals/{id}][%d] getaspecificBalanceIntervalOK  %+v", 200, o.Payload)
}

func (o *GetaspecificBalanceIntervalOK) GetPayload() *models.BalanceIntervals {
	return o.Payload
}

func (o *GetaspecificBalanceIntervalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BalanceIntervals)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
