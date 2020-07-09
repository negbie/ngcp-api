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

// GetaspecificPeeringGroupReader is a Reader for the GetaspecificPeeringGroup structure.
type GetaspecificPeeringGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificPeeringGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificPeeringGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificPeeringGroupOK creates a GetaspecificPeeringGroupOK with default headers values
func NewGetaspecificPeeringGroupOK() *GetaspecificPeeringGroupOK {
	return &GetaspecificPeeringGroupOK{}
}

/*GetaspecificPeeringGroupOK handles this case with default header values.

GetaspecificPeeringGroupOK getaspecific peering group o k
*/
type GetaspecificPeeringGroupOK struct {
	Payload *models.PeeringGroups
}

func (o *GetaspecificPeeringGroupOK) Error() string {
	return fmt.Sprintf("[GET /peeringgroups/{id}][%d] getaspecificPeeringGroupOK  %+v", 200, o.Payload)
}

func (o *GetaspecificPeeringGroupOK) GetPayload() *models.PeeringGroups {
	return o.Payload
}

func (o *GetaspecificPeeringGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PeeringGroups)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
