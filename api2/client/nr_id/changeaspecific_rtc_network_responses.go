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

// ChangeaspecificRtcNetworkReader is a Reader for the ChangeaspecificRtcNetwork structure.
type ChangeaspecificRtcNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeaspecificRtcNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeaspecificRtcNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeaspecificRtcNetworkOK creates a ChangeaspecificRtcNetworkOK with default headers values
func NewChangeaspecificRtcNetworkOK() *ChangeaspecificRtcNetworkOK {
	return &ChangeaspecificRtcNetworkOK{}
}

/*ChangeaspecificRtcNetworkOK handles this case with default header values.

ChangeaspecificRtcNetworkOK changeaspecific rtc network o k
*/
type ChangeaspecificRtcNetworkOK struct {
	Payload *models.RtcNetworks
}

func (o *ChangeaspecificRtcNetworkOK) Error() string {
	return fmt.Sprintf("[PATCH /rtcnetworks/{id}][%d] changeaspecificRtcNetworkOK  %+v", 200, o.Payload)
}

func (o *ChangeaspecificRtcNetworkOK) GetPayload() *models.RtcNetworks {
	return o.Payload
}

func (o *ChangeaspecificRtcNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RtcNetworks)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}