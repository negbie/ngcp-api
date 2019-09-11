// Code generated by go-swagger; DO NOT EDIT.

package id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/negbie/ngcp-api/models"
)

// RtcappsByIDPatchReader is a Reader for the RtcappsByIDPatch structure.
type RtcappsByIDPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RtcappsByIDPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRtcappsByIDPatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRtcappsByIDPatchOK creates a RtcappsByIDPatchOK with default headers values
func NewRtcappsByIDPatchOK() *RtcappsByIDPatchOK {
	return &RtcappsByIDPatchOK{}
}

/*RtcappsByIDPatchOK handles this case with default header values.

RtcappsByIDPatchOK rtcapps by Id patch o k
*/
type RtcappsByIDPatchOK struct {
	Payload *models.RtcApps
}

func (o *RtcappsByIDPatchOK) Error() string {
	return fmt.Sprintf("[PATCH /rtcapps/{id}][%d] rtcappsByIdPatchOK  %+v", 200, o.Payload)
}

func (o *RtcappsByIDPatchOK) GetPayload() *models.RtcApps {
	return o.Payload
}

func (o *RtcappsByIDPatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RtcApps)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}