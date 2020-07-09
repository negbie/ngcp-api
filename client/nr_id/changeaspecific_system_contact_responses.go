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

// ChangeaspecificSystemContactReader is a Reader for the ChangeaspecificSystemContact structure.
type ChangeaspecificSystemContactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeaspecificSystemContactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeaspecificSystemContactOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeaspecificSystemContactOK creates a ChangeaspecificSystemContactOK with default headers values
func NewChangeaspecificSystemContactOK() *ChangeaspecificSystemContactOK {
	return &ChangeaspecificSystemContactOK{}
}

/*ChangeaspecificSystemContactOK handles this case with default header values.

ChangeaspecificSystemContactOK changeaspecific system contact o k
*/
type ChangeaspecificSystemContactOK struct {
	Payload *models.SystemContacts
}

func (o *ChangeaspecificSystemContactOK) Error() string {
	return fmt.Sprintf("[PATCH /systemcontacts/{id}][%d] changeaspecificSystemContactOK  %+v", 200, o.Payload)
}

func (o *ChangeaspecificSystemContactOK) GetPayload() *models.SystemContacts {
	return o.Payload
}

func (o *ChangeaspecificSystemContactOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemContacts)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
