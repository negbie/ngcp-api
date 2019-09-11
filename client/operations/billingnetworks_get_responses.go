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

// BillingnetworksGetReader is a Reader for the BillingnetworksGet structure.
type BillingnetworksGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingnetworksGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBillingnetworksGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBillingnetworksGetOK creates a BillingnetworksGetOK with default headers values
func NewBillingnetworksGetOK() *BillingnetworksGetOK {
	return &BillingnetworksGetOK{}
}

/*BillingnetworksGetOK handles this case with default header values.

BillingnetworksGetOK billingnetworks get o k
*/
type BillingnetworksGetOK struct {
	Payload []*models.BillingNetwork
}

func (o *BillingnetworksGetOK) Error() string {
	return fmt.Sprintf("[GET /billingnetworks][%d] billingnetworksGetOK  %+v", 200, o.Payload)
}

func (o *BillingnetworksGetOK) GetPayload() []*models.BillingNetwork {
	return o.Payload
}

func (o *BillingnetworksGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}