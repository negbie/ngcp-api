// Code generated by go-swagger; DO NOT EDIT.

package pbxdevices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// GetPbxDeviceitemsReader is a Reader for the GetPbxDeviceitems structure.
type GetPbxDeviceitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPbxDeviceitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPbxDeviceitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPbxDeviceitemsOK creates a GetPbxDeviceitemsOK with default headers values
func NewGetPbxDeviceitemsOK() *GetPbxDeviceitemsOK {
	return &GetPbxDeviceitemsOK{}
}

/*GetPbxDeviceitemsOK handles this case with default header values.

GetPbxDeviceitemsOK get pbx deviceitems o k
*/
type GetPbxDeviceitemsOK struct {
	Payload []*models.PbxDevices
}

func (o *GetPbxDeviceitemsOK) Error() string {
	return fmt.Sprintf("[GET /pbxdevices][%d] getPbxDeviceitemsOK  %+v", 200, o.Payload)
}

func (o *GetPbxDeviceitemsOK) GetPayload() []*models.PbxDevices {
	return o.Payload
}

func (o *GetPbxDeviceitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
