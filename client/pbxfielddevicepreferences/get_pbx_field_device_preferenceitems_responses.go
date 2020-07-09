// Code generated by go-swagger; DO NOT EDIT.

package pbxfielddevicepreferences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetPbxFieldDevicePreferenceitemsReader is a Reader for the GetPbxFieldDevicePreferenceitems structure.
type GetPbxFieldDevicePreferenceitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPbxFieldDevicePreferenceitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPbxFieldDevicePreferenceitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPbxFieldDevicePreferenceitemsOK creates a GetPbxFieldDevicePreferenceitemsOK with default headers values
func NewGetPbxFieldDevicePreferenceitemsOK() *GetPbxFieldDevicePreferenceitemsOK {
	return &GetPbxFieldDevicePreferenceitemsOK{}
}

/*GetPbxFieldDevicePreferenceitemsOK handles this case with default header values.

GetPbxFieldDevicePreferenceitemsOK get pbx field device preferenceitems o k
*/
type GetPbxFieldDevicePreferenceitemsOK struct {
	Payload []string
}

func (o *GetPbxFieldDevicePreferenceitemsOK) Error() string {
	return fmt.Sprintf("[GET /pbxfielddevicepreferences][%d] getPbxFieldDevicePreferenceitemsOK  %+v", 200, o.Payload)
}

func (o *GetPbxFieldDevicePreferenceitemsOK) GetPayload() []string {
	return o.Payload
}

func (o *GetPbxFieldDevicePreferenceitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
