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

// ChangeaspecificCFSourceSetReader is a Reader for the ChangeaspecificCFSourceSet structure.
type ChangeaspecificCFSourceSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeaspecificCFSourceSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeaspecificCFSourceSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeaspecificCFSourceSetOK creates a ChangeaspecificCFSourceSetOK with default headers values
func NewChangeaspecificCFSourceSetOK() *ChangeaspecificCFSourceSetOK {
	return &ChangeaspecificCFSourceSetOK{}
}

/*ChangeaspecificCFSourceSetOK handles this case with default header values.

ChangeaspecificCFSourceSetOK changeaspecific c f source set o k
*/
type ChangeaspecificCFSourceSetOK struct {
	Payload *models.CFSourceSets
}

func (o *ChangeaspecificCFSourceSetOK) Error() string {
	return fmt.Sprintf("[PATCH /cfsourcesets/{id}][%d] changeaspecificCFSourceSetOK  %+v", 200, o.Payload)
}

func (o *ChangeaspecificCFSourceSetOK) GetPayload() *models.CFSourceSets {
	return o.Payload
}

func (o *ChangeaspecificCFSourceSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CFSourceSets)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}