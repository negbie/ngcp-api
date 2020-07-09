// Code generated by go-swagger; DO NOT EDIT.

package soundhandles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// GetaspecificSoundHandleReader is a Reader for the GetaspecificSoundHandle structure.
type GetaspecificSoundHandleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificSoundHandleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificSoundHandleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificSoundHandleOK creates a GetaspecificSoundHandleOK with default headers values
func NewGetaspecificSoundHandleOK() *GetaspecificSoundHandleOK {
	return &GetaspecificSoundHandleOK{}
}

/*GetaspecificSoundHandleOK handles this case with default header values.

GetaspecificSoundHandleOK getaspecific sound handle o k
*/
type GetaspecificSoundHandleOK struct {
	Payload *models.SoundHandles
}

func (o *GetaspecificSoundHandleOK) Error() string {
	return fmt.Sprintf("[GET /soundhandles/{id}][%d] getaspecificSoundHandleOK  %+v", 200, o.Payload)
}

func (o *GetaspecificSoundHandleOK) GetPayload() *models.SoundHandles {
	return o.Payload
}

func (o *GetaspecificSoundHandleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SoundHandles)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
