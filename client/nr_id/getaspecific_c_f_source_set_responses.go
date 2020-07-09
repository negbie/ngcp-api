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

// GetaspecificCFSourceSetReader is a Reader for the GetaspecificCFSourceSet structure.
type GetaspecificCFSourceSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificCFSourceSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificCFSourceSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificCFSourceSetOK creates a GetaspecificCFSourceSetOK with default headers values
func NewGetaspecificCFSourceSetOK() *GetaspecificCFSourceSetOK {
	return &GetaspecificCFSourceSetOK{}
}

/*GetaspecificCFSourceSetOK handles this case with default header values.

GetaspecificCFSourceSetOK getaspecific c f source set o k
*/
type GetaspecificCFSourceSetOK struct {
	Payload *models.CFSourceSets
}

func (o *GetaspecificCFSourceSetOK) Error() string {
	return fmt.Sprintf("[GET /cfsourcesets/{id}][%d] getaspecificCFSourceSetOK  %+v", 200, o.Payload)
}

func (o *GetaspecificCFSourceSetOK) GetPayload() *models.CFSourceSets {
	return o.Payload
}

func (o *GetaspecificCFSourceSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CFSourceSets)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
