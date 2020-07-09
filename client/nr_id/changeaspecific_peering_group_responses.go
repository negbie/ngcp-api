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

// ChangeaspecificPeeringGroupReader is a Reader for the ChangeaspecificPeeringGroup structure.
type ChangeaspecificPeeringGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeaspecificPeeringGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeaspecificPeeringGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeaspecificPeeringGroupOK creates a ChangeaspecificPeeringGroupOK with default headers values
func NewChangeaspecificPeeringGroupOK() *ChangeaspecificPeeringGroupOK {
	return &ChangeaspecificPeeringGroupOK{}
}

/*ChangeaspecificPeeringGroupOK handles this case with default header values.

ChangeaspecificPeeringGroupOK changeaspecific peering group o k
*/
type ChangeaspecificPeeringGroupOK struct {
	Payload *models.PeeringGroups
}

func (o *ChangeaspecificPeeringGroupOK) Error() string {
	return fmt.Sprintf("[PATCH /peeringgroups/{id}][%d] changeaspecificPeeringGroupOK  %+v", 200, o.Payload)
}

func (o *ChangeaspecificPeeringGroupOK) GetPayload() *models.PeeringGroups {
	return o.Payload
}

func (o *ChangeaspecificPeeringGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PeeringGroups)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
