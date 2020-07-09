// Code generated by go-swagger; DO NOT EDIT.

package misc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// GetaspecificPbxDeviceFirmwareBinaryReader is a Reader for the GetaspecificPbxDeviceFirmwareBinary structure.
type GetaspecificPbxDeviceFirmwareBinaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificPbxDeviceFirmwareBinaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificPbxDeviceFirmwareBinaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificPbxDeviceFirmwareBinaryOK creates a GetaspecificPbxDeviceFirmwareBinaryOK with default headers values
func NewGetaspecificPbxDeviceFirmwareBinaryOK() *GetaspecificPbxDeviceFirmwareBinaryOK {
	return &GetaspecificPbxDeviceFirmwareBinaryOK{}
}

/*GetaspecificPbxDeviceFirmwareBinaryOK handles this case with default header values.

GetaspecificPbxDeviceFirmwareBinaryOK getaspecific pbx device firmware binary o k
*/
type GetaspecificPbxDeviceFirmwareBinaryOK struct {
	Payload *models.PbxDeviceFirmwareBinaries
}

func (o *GetaspecificPbxDeviceFirmwareBinaryOK) Error() string {
	return fmt.Sprintf("[GET /pbxdevicefirmwarebinaries/{id}][%d] getaspecificPbxDeviceFirmwareBinaryOK  %+v", 200, o.Payload)
}

func (o *GetaspecificPbxDeviceFirmwareBinaryOK) GetPayload() *models.PbxDeviceFirmwareBinaries {
	return o.Payload
}

func (o *GetaspecificPbxDeviceFirmwareBinaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PbxDeviceFirmwareBinaries)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}