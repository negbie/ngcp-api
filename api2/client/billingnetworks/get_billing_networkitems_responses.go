// Code generated by go-swagger; DO NOT EDIT.

package billingnetworks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// GetBillingNetworkitemsReader is a Reader for the GetBillingNetworkitems structure.
type GetBillingNetworkitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBillingNetworkitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBillingNetworkitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBillingNetworkitemsOK creates a GetBillingNetworkitemsOK with default headers values
func NewGetBillingNetworkitemsOK() *GetBillingNetworkitemsOK {
	return &GetBillingNetworkitemsOK{}
}

/*GetBillingNetworkitemsOK handles this case with default header values.

GetBillingNetworkitemsOK get billing networkitems o k
*/
type GetBillingNetworkitemsOK struct {
	Payload []*models.BillingNetworks2
}

func (o *GetBillingNetworkitemsOK) Error() string {
	return fmt.Sprintf("[GET /billingnetworks][%d] getBillingNetworkitemsOK  %+v", 200, o.Payload)
}

func (o *GetBillingNetworkitemsOK) GetPayload() []*models.BillingNetworks2 {
	return o.Payload
}

func (o *GetBillingNetworkitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}