// Code generated by go-swagger; DO NOT EDIT.

package peeringservers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// GetPeeringServeritemsReader is a Reader for the GetPeeringServeritems structure.
type GetPeeringServeritemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPeeringServeritemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPeeringServeritemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPeeringServeritemsOK creates a GetPeeringServeritemsOK with default headers values
func NewGetPeeringServeritemsOK() *GetPeeringServeritemsOK {
	return &GetPeeringServeritemsOK{}
}

/*GetPeeringServeritemsOK handles this case with default header values.

GetPeeringServeritemsOK get peering serveritems o k
*/
type GetPeeringServeritemsOK struct {
	Payload []*models.PeeringServers3
}

func (o *GetPeeringServeritemsOK) Error() string {
	return fmt.Sprintf("[GET /peeringservers][%d] getPeeringServeritemsOK  %+v", 200, o.Payload)
}

func (o *GetPeeringServeritemsOK) GetPayload() []*models.PeeringServers3 {
	return o.Payload
}

func (o *GetPeeringServeritemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
