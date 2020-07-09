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

// GetaspecificVoicemailReader is a Reader for the GetaspecificVoicemail structure.
type GetaspecificVoicemailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificVoicemailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificVoicemailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificVoicemailOK creates a GetaspecificVoicemailOK with default headers values
func NewGetaspecificVoicemailOK() *GetaspecificVoicemailOK {
	return &GetaspecificVoicemailOK{}
}

/*GetaspecificVoicemailOK handles this case with default header values.

GetaspecificVoicemailOK getaspecific voicemail o k
*/
type GetaspecificVoicemailOK struct {
	Payload *models.Voicemails
}

func (o *GetaspecificVoicemailOK) Error() string {
	return fmt.Sprintf("[GET /voicemails/{id}][%d] getaspecificVoicemailOK  %+v", 200, o.Payload)
}

func (o *GetaspecificVoicemailOK) GetPayload() *models.Voicemails {
	return o.Payload
}

func (o *GetaspecificVoicemailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Voicemails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
