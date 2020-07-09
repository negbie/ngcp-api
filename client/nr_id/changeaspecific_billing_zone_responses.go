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

// ChangeaspecificBillingZoneReader is a Reader for the ChangeaspecificBillingZone structure.
type ChangeaspecificBillingZoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeaspecificBillingZoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewChangeaspecificBillingZoneNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeaspecificBillingZoneNoContent creates a ChangeaspecificBillingZoneNoContent with default headers values
func NewChangeaspecificBillingZoneNoContent() *ChangeaspecificBillingZoneNoContent {
	return &ChangeaspecificBillingZoneNoContent{}
}

/*ChangeaspecificBillingZoneNoContent handles this case with default header values.

ChangeaspecificBillingZoneNoContent changeaspecific billing zone no content
*/
type ChangeaspecificBillingZoneNoContent struct {
	Payload interface{}
}

func (o *ChangeaspecificBillingZoneNoContent) Error() string {
	return fmt.Sprintf("[PATCH /billingzones/{id}][%d] changeaspecificBillingZoneNoContent  %+v", 204, o.Payload)
}

func (o *ChangeaspecificBillingZoneNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ChangeaspecificBillingZoneNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
