// Code generated by go-swagger; DO NOT EDIT.

package maliciouscalls

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// GetMaliciousCallitemsReader is a Reader for the GetMaliciousCallitems structure.
type GetMaliciousCallitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMaliciousCallitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMaliciousCallitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMaliciousCallitemsOK creates a GetMaliciousCallitemsOK with default headers values
func NewGetMaliciousCallitemsOK() *GetMaliciousCallitemsOK {
	return &GetMaliciousCallitemsOK{}
}

/*GetMaliciousCallitemsOK handles this case with default header values.

GetMaliciousCallitemsOK get malicious callitems o k
*/
type GetMaliciousCallitemsOK struct {
	Payload []*models.MaliciousCalls
}

func (o *GetMaliciousCallitemsOK) Error() string {
	return fmt.Sprintf("[GET /maliciouscalls][%d] getMaliciousCallitemsOK  %+v", 200, o.Payload)
}

func (o *GetMaliciousCallitemsOK) GetPayload() []*models.MaliciousCalls {
	return o.Payload
}

func (o *GetMaliciousCallitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
