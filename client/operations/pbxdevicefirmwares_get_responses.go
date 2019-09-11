// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/negbie/ngcp-api/models"
)

// PbxdevicefirmwaresGetReader is a Reader for the PbxdevicefirmwaresGet structure.
type PbxdevicefirmwaresGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PbxdevicefirmwaresGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPbxdevicefirmwaresGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPbxdevicefirmwaresGetOK creates a PbxdevicefirmwaresGetOK with default headers values
func NewPbxdevicefirmwaresGetOK() *PbxdevicefirmwaresGetOK {
	return &PbxdevicefirmwaresGetOK{}
}

/*PbxdevicefirmwaresGetOK handles this case with default header values.

PbxdevicefirmwaresGetOK pbxdevicefirmwares get o k
*/
type PbxdevicefirmwaresGetOK struct {
	Payload []*models.PbxDeviceFirmware
}

func (o *PbxdevicefirmwaresGetOK) Error() string {
	return fmt.Sprintf("[GET /pbxdevicefirmwares][%d] pbxdevicefirmwaresGetOK  %+v", 200, o.Payload)
}

func (o *PbxdevicefirmwaresGetOK) GetPayload() []*models.PbxDeviceFirmware {
	return o.Payload
}

func (o *PbxdevicefirmwaresGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}