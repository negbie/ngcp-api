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

// ChangeaspecificCFDestinationSetReader is a Reader for the ChangeaspecificCFDestinationSet structure.
type ChangeaspecificCFDestinationSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeaspecificCFDestinationSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeaspecificCFDestinationSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeaspecificCFDestinationSetOK creates a ChangeaspecificCFDestinationSetOK with default headers values
func NewChangeaspecificCFDestinationSetOK() *ChangeaspecificCFDestinationSetOK {
	return &ChangeaspecificCFDestinationSetOK{}
}

/*ChangeaspecificCFDestinationSetOK handles this case with default header values.

ChangeaspecificCFDestinationSetOK changeaspecific c f destination set o k
*/
type ChangeaspecificCFDestinationSetOK struct {
	Payload *models.CFDestinationSets
}

func (o *ChangeaspecificCFDestinationSetOK) Error() string {
	return fmt.Sprintf("[PATCH /cfdestinationsets/{id}][%d] changeaspecificCFDestinationSetOK  %+v", 200, o.Payload)
}

func (o *ChangeaspecificCFDestinationSetOK) GetPayload() *models.CFDestinationSets {
	return o.Payload
}

func (o *ChangeaspecificCFDestinationSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CFDestinationSets)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
