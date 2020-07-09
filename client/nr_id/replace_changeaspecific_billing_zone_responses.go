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

// ReplaceChangeaspecificBillingZoneReader is a Reader for the ReplaceChangeaspecificBillingZone structure.
type ReplaceChangeaspecificBillingZoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceChangeaspecificBillingZoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewReplaceChangeaspecificBillingZoneNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceChangeaspecificBillingZoneNoContent creates a ReplaceChangeaspecificBillingZoneNoContent with default headers values
func NewReplaceChangeaspecificBillingZoneNoContent() *ReplaceChangeaspecificBillingZoneNoContent {
	return &ReplaceChangeaspecificBillingZoneNoContent{}
}

/*ReplaceChangeaspecificBillingZoneNoContent handles this case with default header values.

ReplaceChangeaspecificBillingZoneNoContent replace changeaspecific billing zone no content
*/
type ReplaceChangeaspecificBillingZoneNoContent struct {
	Payload interface{}
}

func (o *ReplaceChangeaspecificBillingZoneNoContent) Error() string {
	return fmt.Sprintf("[PUT /billingzones/{id}][%d] replaceChangeaspecificBillingZoneNoContent  %+v", 204, o.Payload)
}

func (o *ReplaceChangeaspecificBillingZoneNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ReplaceChangeaspecificBillingZoneNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}