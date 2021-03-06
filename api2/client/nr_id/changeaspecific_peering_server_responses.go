// Code generated by go-swagger; DO NOT EDIT.

package nr_id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// ChangeaspecificPeeringServerReader is a Reader for the ChangeaspecificPeeringServer structure.
type ChangeaspecificPeeringServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeaspecificPeeringServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeaspecificPeeringServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeaspecificPeeringServerOK creates a ChangeaspecificPeeringServerOK with default headers values
func NewChangeaspecificPeeringServerOK() *ChangeaspecificPeeringServerOK {
	return &ChangeaspecificPeeringServerOK{}
}

/*ChangeaspecificPeeringServerOK handles this case with default header values.

ChangeaspecificPeeringServerOK changeaspecific peering server o k
*/
type ChangeaspecificPeeringServerOK struct {
	Payload *models.PeeringServers
}

func (o *ChangeaspecificPeeringServerOK) Error() string {
	return fmt.Sprintf("[PATCH /peeringservers/{id}][%d] changeaspecificPeeringServerOK  %+v", 200, o.Payload)
}

func (o *ChangeaspecificPeeringServerOK) GetPayload() *models.PeeringServers {
	return o.Payload
}

func (o *ChangeaspecificPeeringServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PeeringServers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
