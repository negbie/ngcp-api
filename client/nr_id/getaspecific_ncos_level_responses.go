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

// GetaspecificNcosLevelReader is a Reader for the GetaspecificNcosLevel structure.
type GetaspecificNcosLevelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificNcosLevelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificNcosLevelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificNcosLevelOK creates a GetaspecificNcosLevelOK with default headers values
func NewGetaspecificNcosLevelOK() *GetaspecificNcosLevelOK {
	return &GetaspecificNcosLevelOK{}
}

/*GetaspecificNcosLevelOK handles this case with default header values.

GetaspecificNcosLevelOK getaspecific ncos level o k
*/
type GetaspecificNcosLevelOK struct {
	Payload *models.NcosLevels
}

func (o *GetaspecificNcosLevelOK) Error() string {
	return fmt.Sprintf("[GET /ncoslevels/{id}][%d] getaspecificNcosLevelOK  %+v", 200, o.Payload)
}

func (o *GetaspecificNcosLevelOK) GetPayload() *models.NcosLevels {
	return o.Payload
}

func (o *GetaspecificNcosLevelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NcosLevels)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
