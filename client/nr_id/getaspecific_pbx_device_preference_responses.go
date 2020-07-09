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

// GetaspecificPbxDevicePreferenceReader is a Reader for the GetaspecificPbxDevicePreference structure.
type GetaspecificPbxDevicePreferenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificPbxDevicePreferenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificPbxDevicePreferenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificPbxDevicePreferenceOK creates a GetaspecificPbxDevicePreferenceOK with default headers values
func NewGetaspecificPbxDevicePreferenceOK() *GetaspecificPbxDevicePreferenceOK {
	return &GetaspecificPbxDevicePreferenceOK{}
}

/*GetaspecificPbxDevicePreferenceOK handles this case with default header values.

GetaspecificPbxDevicePreferenceOK getaspecific pbx device preference o k
*/
type GetaspecificPbxDevicePreferenceOK struct {
	Payload string
}

func (o *GetaspecificPbxDevicePreferenceOK) Error() string {
	return fmt.Sprintf("[GET /pbxdevicepreferences/{id}][%d] getaspecificPbxDevicePreferenceOK  %+v", 200, o.Payload)
}

func (o *GetaspecificPbxDevicePreferenceOK) GetPayload() string {
	return o.Payload
}

func (o *GetaspecificPbxDevicePreferenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
