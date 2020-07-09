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

// GetaspecificPbxDeviceProfileReader is a Reader for the GetaspecificPbxDeviceProfile structure.
type GetaspecificPbxDeviceProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificPbxDeviceProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificPbxDeviceProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificPbxDeviceProfileOK creates a GetaspecificPbxDeviceProfileOK with default headers values
func NewGetaspecificPbxDeviceProfileOK() *GetaspecificPbxDeviceProfileOK {
	return &GetaspecificPbxDeviceProfileOK{}
}

/*GetaspecificPbxDeviceProfileOK handles this case with default header values.

GetaspecificPbxDeviceProfileOK getaspecific pbx device profile o k
*/
type GetaspecificPbxDeviceProfileOK struct {
	Payload *models.PbxDeviceProfiles
}

func (o *GetaspecificPbxDeviceProfileOK) Error() string {
	return fmt.Sprintf("[GET /pbxdeviceprofiles/{id}][%d] getaspecificPbxDeviceProfileOK  %+v", 200, o.Payload)
}

func (o *GetaspecificPbxDeviceProfileOK) GetPayload() *models.PbxDeviceProfiles {
	return o.Payload
}

func (o *GetaspecificPbxDeviceProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PbxDeviceProfiles)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
