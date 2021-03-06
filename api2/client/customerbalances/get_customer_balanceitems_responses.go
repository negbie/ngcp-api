// Code generated by go-swagger; DO NOT EDIT.

package customerbalances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// GetCustomerBalanceitemsReader is a Reader for the GetCustomerBalanceitems structure.
type GetCustomerBalanceitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomerBalanceitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomerBalanceitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCustomerBalanceitemsOK creates a GetCustomerBalanceitemsOK with default headers values
func NewGetCustomerBalanceitemsOK() *GetCustomerBalanceitemsOK {
	return &GetCustomerBalanceitemsOK{}
}

/*GetCustomerBalanceitemsOK handles this case with default header values.

GetCustomerBalanceitemsOK get customer balanceitems o k
*/
type GetCustomerBalanceitemsOK struct {
	Payload []*models.CustomerBalances
}

func (o *GetCustomerBalanceitemsOK) Error() string {
	return fmt.Sprintf("[GET /customerbalances][%d] getCustomerBalanceitemsOK  %+v", 200, o.Payload)
}

func (o *GetCustomerBalanceitemsOK) GetPayload() []*models.CustomerBalances {
	return o.Payload
}

func (o *GetCustomerBalanceitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
