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

// GetaspecificMaliciousCallReader is a Reader for the GetaspecificMaliciousCall structure.
type GetaspecificMaliciousCallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificMaliciousCallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificMaliciousCallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificMaliciousCallOK creates a GetaspecificMaliciousCallOK with default headers values
func NewGetaspecificMaliciousCallOK() *GetaspecificMaliciousCallOK {
	return &GetaspecificMaliciousCallOK{}
}

/*GetaspecificMaliciousCallOK handles this case with default header values.

GetaspecificMaliciousCallOK getaspecific malicious call o k
*/
type GetaspecificMaliciousCallOK struct {
	Payload *models.MaliciousCalls
}

func (o *GetaspecificMaliciousCallOK) Error() string {
	return fmt.Sprintf("[GET /maliciouscalls/{id}][%d] getaspecificMaliciousCallOK  %+v", 200, o.Payload)
}

func (o *GetaspecificMaliciousCallOK) GetPayload() *models.MaliciousCalls {
	return o.Payload
}

func (o *GetaspecificMaliciousCallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MaliciousCalls)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}