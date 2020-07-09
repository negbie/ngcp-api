// Code generated by go-swagger; DO NOT EDIT.

package nr_id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// GetaspecificCustomerBalanceReader is a Reader for the GetaspecificCustomerBalance structure.
type GetaspecificCustomerBalanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificCustomerBalanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificCustomerBalanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificCustomerBalanceOK creates a GetaspecificCustomerBalanceOK with default headers values
func NewGetaspecificCustomerBalanceOK() *GetaspecificCustomerBalanceOK {
	return &GetaspecificCustomerBalanceOK{}
}

/*GetaspecificCustomerBalanceOK handles this case with default header values.

GetaspecificCustomerBalanceOK getaspecific customer balance o k
*/
type GetaspecificCustomerBalanceOK struct {
	Payload *models.CustomerBalances
}

func (o *GetaspecificCustomerBalanceOK) Error() string {
	return fmt.Sprintf("[GET /customerbalances/{id}][%d] getaspecificCustomerBalanceOK  %+v", 200, o.Payload)
}

func (o *GetaspecificCustomerBalanceOK) GetPayload() *models.CustomerBalances {
	return o.Payload
}

func (o *GetaspecificCustomerBalanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerBalances)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
