// Code generated by go-swagger; DO NOT EDIT.

package billingzones

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// GetBillingZoneitemsReader is a Reader for the GetBillingZoneitems structure.
type GetBillingZoneitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBillingZoneitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBillingZoneitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBillingZoneitemsOK creates a GetBillingZoneitemsOK with default headers values
func NewGetBillingZoneitemsOK() *GetBillingZoneitemsOK {
	return &GetBillingZoneitemsOK{}
}

/*GetBillingZoneitemsOK handles this case with default header values.

GetBillingZoneitemsOK get billing zoneitems o k
*/
type GetBillingZoneitemsOK struct {
	Payload []*models.BillingZones
}

func (o *GetBillingZoneitemsOK) Error() string {
	return fmt.Sprintf("[GET /billingzones][%d] getBillingZoneitemsOK  %+v", 200, o.Payload)
}

func (o *GetBillingZoneitemsOK) GetPayload() []*models.BillingZones {
	return o.Payload
}

func (o *GetBillingZoneitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
