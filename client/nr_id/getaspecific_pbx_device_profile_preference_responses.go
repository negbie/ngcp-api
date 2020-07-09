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

// GetaspecificPbxDeviceProfilePreferenceReader is a Reader for the GetaspecificPbxDeviceProfilePreference structure.
type GetaspecificPbxDeviceProfilePreferenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificPbxDeviceProfilePreferenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificPbxDeviceProfilePreferenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificPbxDeviceProfilePreferenceOK creates a GetaspecificPbxDeviceProfilePreferenceOK with default headers values
func NewGetaspecificPbxDeviceProfilePreferenceOK() *GetaspecificPbxDeviceProfilePreferenceOK {
	return &GetaspecificPbxDeviceProfilePreferenceOK{}
}

/*GetaspecificPbxDeviceProfilePreferenceOK handles this case with default header values.

GetaspecificPbxDeviceProfilePreferenceOK getaspecific pbx device profile preference o k
*/
type GetaspecificPbxDeviceProfilePreferenceOK struct {
	Payload string
}

func (o *GetaspecificPbxDeviceProfilePreferenceOK) Error() string {
	return fmt.Sprintf("[GET /pbxdeviceprofilepreferences/{id}][%d] getaspecificPbxDeviceProfilePreferenceOK  %+v", 200, o.Payload)
}

func (o *GetaspecificPbxDeviceProfilePreferenceOK) GetPayload() string {
	return o.Payload
}

func (o *GetaspecificPbxDeviceProfilePreferenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
