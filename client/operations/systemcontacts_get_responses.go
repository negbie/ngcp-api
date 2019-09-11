// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/negbie/ngcp-api/models"
)

// SystemcontactsGetReader is a Reader for the SystemcontactsGet structure.
type SystemcontactsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemcontactsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSystemcontactsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSystemcontactsGetOK creates a SystemcontactsGetOK with default headers values
func NewSystemcontactsGetOK() *SystemcontactsGetOK {
	return &SystemcontactsGetOK{}
}

/*SystemcontactsGetOK handles this case with default header values.

SystemcontactsGetOK systemcontacts get o k
*/
type SystemcontactsGetOK struct {
	Payload []*models.SystemContacts1
}

func (o *SystemcontactsGetOK) Error() string {
	return fmt.Sprintf("[GET /systemcontacts][%d] systemcontactsGetOK  %+v", 200, o.Payload)
}

func (o *SystemcontactsGetOK) GetPayload() []*models.SystemContacts1 {
	return o.Payload
}

func (o *SystemcontactsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}