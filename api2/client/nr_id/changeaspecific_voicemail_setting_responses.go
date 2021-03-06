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

// ChangeaspecificVoicemailSettingReader is a Reader for the ChangeaspecificVoicemailSetting structure.
type ChangeaspecificVoicemailSettingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeaspecificVoicemailSettingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeaspecificVoicemailSettingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeaspecificVoicemailSettingOK creates a ChangeaspecificVoicemailSettingOK with default headers values
func NewChangeaspecificVoicemailSettingOK() *ChangeaspecificVoicemailSettingOK {
	return &ChangeaspecificVoicemailSettingOK{}
}

/*ChangeaspecificVoicemailSettingOK handles this case with default header values.

ChangeaspecificVoicemailSettingOK changeaspecific voicemail setting o k
*/
type ChangeaspecificVoicemailSettingOK struct {
	Payload *models.VoicemailSettings
}

func (o *ChangeaspecificVoicemailSettingOK) Error() string {
	return fmt.Sprintf("[PATCH /voicemailsettings/{id}][%d] changeaspecificVoicemailSettingOK  %+v", 200, o.Payload)
}

func (o *ChangeaspecificVoicemailSettingOK) GetPayload() *models.VoicemailSettings {
	return o.Payload
}

func (o *ChangeaspecificVoicemailSettingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VoicemailSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
