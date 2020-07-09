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

// ReplaceChangeaspecificLnpCarrierReader is a Reader for the ReplaceChangeaspecificLnpCarrier structure.
type ReplaceChangeaspecificLnpCarrierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceChangeaspecificLnpCarrierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceChangeaspecificLnpCarrierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceChangeaspecificLnpCarrierOK creates a ReplaceChangeaspecificLnpCarrierOK with default headers values
func NewReplaceChangeaspecificLnpCarrierOK() *ReplaceChangeaspecificLnpCarrierOK {
	return &ReplaceChangeaspecificLnpCarrierOK{}
}

/*ReplaceChangeaspecificLnpCarrierOK handles this case with default header values.

ReplaceChangeaspecificLnpCarrierOK replace changeaspecific lnp carrier o k
*/
type ReplaceChangeaspecificLnpCarrierOK struct {
	Payload *models.LnpCarriers
}

func (o *ReplaceChangeaspecificLnpCarrierOK) Error() string {
	return fmt.Sprintf("[PUT /lnpcarriers/{id}][%d] replaceChangeaspecificLnpCarrierOK  %+v", 200, o.Payload)
}

func (o *ReplaceChangeaspecificLnpCarrierOK) GetPayload() *models.LnpCarriers {
	return o.Payload
}

func (o *ReplaceChangeaspecificLnpCarrierOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LnpCarriers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
