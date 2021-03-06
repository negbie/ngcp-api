// Code generated by go-swagger; DO NOT EDIT.

package nr_id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ChangeaspecificPbxFieldDevicePreferenceReader is a Reader for the ChangeaspecificPbxFieldDevicePreference structure.
type ChangeaspecificPbxFieldDevicePreferenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeaspecificPbxFieldDevicePreferenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewChangeaspecificPbxFieldDevicePreferenceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeaspecificPbxFieldDevicePreferenceNoContent creates a ChangeaspecificPbxFieldDevicePreferenceNoContent with default headers values
func NewChangeaspecificPbxFieldDevicePreferenceNoContent() *ChangeaspecificPbxFieldDevicePreferenceNoContent {
	return &ChangeaspecificPbxFieldDevicePreferenceNoContent{}
}

/*ChangeaspecificPbxFieldDevicePreferenceNoContent handles this case with default header values.

ChangeaspecificPbxFieldDevicePreferenceNoContent changeaspecific pbx field device preference no content
*/
type ChangeaspecificPbxFieldDevicePreferenceNoContent struct {
	Payload interface{}
}

func (o *ChangeaspecificPbxFieldDevicePreferenceNoContent) Error() string {
	return fmt.Sprintf("[PATCH /pbxfielddevicepreferences/{id}][%d] changeaspecificPbxFieldDevicePreferenceNoContent  %+v", 204, o.Payload)
}

func (o *ChangeaspecificPbxFieldDevicePreferenceNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ChangeaspecificPbxFieldDevicePreferenceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
